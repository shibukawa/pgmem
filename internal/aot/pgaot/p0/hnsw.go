package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_HnswCheckNorm(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 float64
	_ = v9
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v5 = F_FunctionCall1Coll(m, v3, v4, l1)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = *(*float64)(unsafe.Add(mBase, uint32(v5)))
		return base.F64_gt(v9, float64(0))
	}
}
func F_HnswFindElementNeighbors(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) {
	mBase := m.M
	_ = mBase
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v83 int32
	_ = v83
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v114 int32
	_ = v114
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v139 int32
	_ = v139
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 float64
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v204 int32
	_ = v204
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v282 int32
	_ = v282
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v346 int32
	_ = v346
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v369 int32
	_ = v369
	var v373 int32
	_ = v373
	var v374 int64
	_ = v374
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	v19 = m.G0
	v21 = v19 - int32(16)
	m.G0 = v21
	if l0 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+65)))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+12)) = v34
	if l3 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L2:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	v34 = v25
	goto L1
L3:
	;
	goto L4
L4:
	;
	v26 = int32(0)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	if v27 == v26 {
		v34 = v26
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v34 = l0 + v27 - int32(1)
	goto L1
L6:
	;
	v39 = l1 - l0
	v40 = int32(16)
	v44 = (int32(base.Ui32(v39)>>(uint(v40)%32)) ^ v39) * int32(-2048144789)
	v49 = (int32(base.Ui32(v44)>>(uint(int32(13))%32)) ^ v44) * int32(-1028477387)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+68)) = int32(base.Ui32(v49)>>(uint(v40)%32)) ^ v49
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
	m.G0 = v21 + int32(16)
	return
L12:
	;
	v56 = l1
	goto L14
L13:
	;
	v56 = int32(0)
	goto L14
L14:
	;
	v60 = F_HnswEntryCandidate(m, l0, l2, v21+int32(12), l3, l4, int32(1))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	return
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = v60
	v67 = F_list_make1_impl(m, int32(1), v21+int32(4))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+65)))
	if base.Ui32(v35) < base.Ui32(v69) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v73 = v69
	v83 = v67
	goto L21
L19:
	;
	v114 = v67
	goto L20
L20:
	;
	if base.Ui32(v35) < base.Ui32(v69) {
		goto L25
	} else {
		goto L26
	}
L21:
	;
	v91 = int32(1)
	v93 = int32(0)
	v97 = F_HnswSearchLayer(m, l0, v21+int32(12), v83, v91, v73, l3, l4, l5, v91, v56, v93, v93, v91, v93)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L15
	} else {
		goto L23
	}
L22:
	;
	v114 = v97
	goto L20
L23:
	;
	v100 = v73 - int32(1)
	if v35 < v100 {
		v73 = v100
		v83 = v97
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	v121 = v35
	goto L27
L26:
	;
	v121 = v69
	goto L27
L27:
	;
	v125 = l0 - int32(1)
	v129 = v121
	v139 = v114
	goto L28
L28:
	;
	v147 = int32(1)
	v148 = int32(0)
	v152 = F_HnswSearchLayer(m, l0, v21+int32(12), v139, l6+l7, v129, l3, l4, l5, v147, v56, v148, v148, v147, v148)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L15
	} else {
		goto L33
	}
L29:
	;
	goto L11
L30:
	;
	v296 = l5 << (uint(base.B2i32(v129 == int32(0))) % 32)
	if l0 != 0 {
		goto L68
	} else {
		goto L69
	}
L31:
	;
	if l3 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L32:
	;
	v162 = v154
	v167 = v154
	goto L38
L33:
	;
	if v152 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v154 = int32(0)
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v152)+4))
	if v154 < v156 {
		goto L32
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v282 = int32(0)
	goto L30
L37:
	;
	v204 = v154
	goto L31
L38:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v152)+12))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v178+v162<<(uint(int32(2))%32))))
	v184 = F_palloc(m, int32(12))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L15
	} else {
		goto L40
	}
L39:
	;
	v204 = v191
	goto L31
L40:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v182)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v184))) = v186
	v188 = *(*float64)(unsafe.Add(mBase, uint32(v182)+32))
	*(*float32)(unsafe.Add(mBase, uint32(v184)+4)) = base.F32_demote_f64(v188)
	v191 = F_lappend(m, v167, v184)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L15
	} else {
		goto L41
	}
L41:
	;
	v194 = v162 + int32(1)
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v152)+4))
	if v194 < v195 {
		v162 = v194
		v167 = v191
		goto L38
	} else {
		goto L42
	}
L42:
	;
	goto L39
L43:
	;
	v282 = v204
	goto L30
L44:
	;
	goto L45
L45:
	;
	if v204 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v282 = int32(0)
	goto L30
L47:
	;
	goto L48
L48:
	;
	v220 = int32(0)
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v204)+4))
	if v222 <= v220 {
		v282 = v220
		goto L30
	} else {
		goto L49
	}
L49:
	;
	v227 = v220
	v231 = v220
	v235 = v222
	goto L50
L50:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v204)+12))
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v243+v227<<(uint(int32(2))%32))))
	if l0 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	v282 = v271
	goto L30
L52:
	;
	if v56 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L53:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v247)))
	v256 = v250
	goto L52
L54:
	;
	goto L55
L55:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v247)))
	if v251 != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v254 = v125 + v251
	goto L58
L57:
	;
	v254 = int32(0)
	goto L58
L58:
	;
	v256 = v254
	goto L52
L59:
	;
	v274 = v227 + int32(1)
	if v274 < v272 {
		v227 = v274
		v231 = v271
		v235 = v272
		goto L50
	} else {
		goto L66
	}
L60:
	;
	v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v256)+64)))
	if v265 == int32(0) {
		v271 = v231
		v272 = v235
		goto L59
	} else {
		goto L64
	}
L61:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v256)+76))
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v56)+76))
	if v259 != v260 {
		goto L60
	} else {
		goto L62
	}
L62:
	;
	v262 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v256)+80)))
	v263 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v56)+80)))
	if v262 == v263 {
		v271 = v231
		v272 = v235
		goto L59
	} else {
		goto L63
	}
L63:
	;
	goto L60
L64:
	;
	v268 = F_lappend(m, v231, v247)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L15
	} else {
		goto L65
	}
L65:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v204)+4))
	v271 = v268
	v272 = v270
	goto L59
L66:
	;
	goto L51
L67:
	;
	if v333 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L68:
	;
	v298 = v129 << (uint(int32(2)) % 32)
	v299 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v298+(v125+v299))))
	v304 = int32(0)
	v307 = F_SelectNeighbors(m, l0, v282, v296, l4, l0+int32(3)+v302, v304, v304, v304)
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L15
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v316 = int32(0)
	v318 = v129 << (uint(int32(2)) % 32)
	v319 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v318+v319)))
	v327 = F_SelectNeighbors(m, v316, v282, v296, l4, v321+int32(4), v316, v316, v316)
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L15
	} else {
		goto L75
	}
L71:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v125+v309+v298)))
	if v312 != 0 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v315 = v125 + v312
	goto L74
L73:
	;
	v315 = int32(0)
	goto L74
L74:
	;
	v333 = v307
	v335 = v315
	goto L67
L75:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v329+v318)))
	v333 = v327
	v335 = v331
	goto L67
L76:
	;
	if int32(0) < v129 {
		v129 = v129 - int32(1)
		v139 = v152
		goto L28
	} else {
		goto L82
	}
L77:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v333)+4))
	if v338 <= int32(0) {
		goto L76
	} else {
		goto L78
	}
L78:
	;
	v346 = int32(0)
	goto L79
L79:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v333)+12))
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v335)))
	v364 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v335))) = v363 + v364
	v369 = v335 + int32(8) + v363*int32(12)
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v362+v346<<(uint(int32(2))%32))))
	v374 = *(*int64)(unsafe.Add(mBase, uint32(v373)))
	*(*int64)(unsafe.Add(mBase, uint32(v369))) = v374
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v373)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v369)+8)) = v376
	v379 = v346 + v364
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v333)+4))
	if v379 < v380 {
		v346 = v379
		goto L79
	} else {
		goto L81
	}
L80:
	;
	goto L76
L81:
	;
	goto L80
L82:
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
		F_LockBuffer(m, v8, int32(1))
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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v96 int32
	_ = v96
	var v101 float64
	_ = v101
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
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
			v22 = F_ShmemInitStruct(m, int32(_a_F_HnswInit_0), int32(4), v5+int32(15))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return
			} else {
				v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+15)))
				if v24 == int32(1) {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
					v31 = v27
					*(*int32)(unsafe.Add(mBase, _c_F_HnswInit[2])) = v31
					v35 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInit[1]))
					F_LWLockRelease(m, v35+int32(2688))
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return
					} else {
						v41 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInit[2]))
						F_LWLockRegisterTranche(m, v41, int32(_a_F_HnswInit_1))
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return
						} else {
							v47 = F_add_reloption_kind(m)
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, _c_F_HnswInit[3])) = v47
								F_add_int_reloption(m, v47, int32(_a_F_HnswInit_2), int32(_a_F_HnswInit_3), int32(16), int32(2), int32(100))
								mBase = m.M
								v56 = m.ExcPending
								if v56 != 0 {
									return
								} else {
									v58 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInit[3]))
									F_add_int_reloption(m, v58, int32(_a_F_HnswInit_4), int32(_a_F_HnswInit_5), int32(64), int32(4), int32(1000))
									mBase = m.M
									v65 = m.ExcPending
									if v65 != 0 {
										return
									} else {
										F_DefineCustomIntVariable(m, int32(_a_F_HnswInit_6), int32(_a_F_HnswInit_7), int32(_a_F_HnswInit_8), int32(_a_F_HnswInit_9), int32(40), int32(1), int32(1000), int32(6), int32(0))
										mBase = m.M
										v76 = m.ExcPending
										if v76 != 0 {
											return
										} else {
											v79 = int32(0)
											F_DefineCustomEnumVariable(m, int32(_a_F_HnswInit_10), int32(_a_F_HnswInit_11), v79, int32(_a_F_HnswInit_12), v79, int32(_a_F_HnswInit_13), int32(6))
											mBase = m.M
											v85 = m.ExcPending
											if v85 != 0 {
												return
											} else {
												v88 = int32(0)
												F_DefineCustomIntVariable(m, int32(_a_F_HnswInit_14), int32(_a_F_HnswInit_15), v88, int32(_a_F_HnswInit_16), int32(_a_F_HnswInit_17), int32(1), int32(2147483647), int32(6), v88)
												mBase = m.M
												v96 = m.ExcPending
												if v96 != 0 {
													return
												} else {
													v101 = float64(1)
													F_DefineCustomRealVariable(m, int32(_a_F_HnswInit_18), int32(_a_F_HnswInit_19), int32(0), int32(_a_F_HnswInit_20), v101, v101, float64(1000))
													mBase = m.M
													v105 = m.ExcPending
													if v105 != 0 {
														return
													} else {
														F_MarkGUCPrefixReserved(m, int32(_a_F_HnswInit_21))
														mBase = m.M
														v108 = m.ExcPending
														if v108 != 0 {
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
				} else {
					v28 = F_LWLockNewTrancheId(m)
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v22))) = v28
						v31 = v28
						*(*int32)(unsafe.Add(mBase, _c_F_HnswInit[2])) = v31
						v35 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInit[1]))
						F_LWLockRelease(m, v35+int32(2688))
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return
						} else {
							v41 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInit[2]))
							F_LWLockRegisterTranche(m, v41, int32(_a_F_HnswInit_1))
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return
							} else {
								v47 = F_add_reloption_kind(m)
								mBase = m.M
								v48 = m.ExcPending
								if v48 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, _c_F_HnswInit[3])) = v47
									F_add_int_reloption(m, v47, int32(_a_F_HnswInit_2), int32(_a_F_HnswInit_3), int32(16), int32(2), int32(100))
									mBase = m.M
									v56 = m.ExcPending
									if v56 != 0 {
										return
									} else {
										v58 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInit[3]))
										F_add_int_reloption(m, v58, int32(_a_F_HnswInit_4), int32(_a_F_HnswInit_5), int32(64), int32(4), int32(1000))
										mBase = m.M
										v65 = m.ExcPending
										if v65 != 0 {
											return
										} else {
											F_DefineCustomIntVariable(m, int32(_a_F_HnswInit_6), int32(_a_F_HnswInit_7), int32(_a_F_HnswInit_8), int32(_a_F_HnswInit_9), int32(40), int32(1), int32(1000), int32(6), int32(0))
											mBase = m.M
											v76 = m.ExcPending
											if v76 != 0 {
												return
											} else {
												v79 = int32(0)
												F_DefineCustomEnumVariable(m, int32(_a_F_HnswInit_10), int32(_a_F_HnswInit_11), v79, int32(_a_F_HnswInit_12), v79, int32(_a_F_HnswInit_13), int32(6))
												mBase = m.M
												v85 = m.ExcPending
												if v85 != 0 {
													return
												} else {
													v88 = int32(0)
													F_DefineCustomIntVariable(m, int32(_a_F_HnswInit_14), int32(_a_F_HnswInit_15), v88, int32(_a_F_HnswInit_16), int32(_a_F_HnswInit_17), int32(1), int32(2147483647), int32(6), v88)
													mBase = m.M
													v96 = m.ExcPending
													if v96 != 0 {
														return
													} else {
														v101 = float64(1)
														F_DefineCustomRealVariable(m, int32(_a_F_HnswInit_18), int32(_a_F_HnswInit_19), int32(0), int32(_a_F_HnswInit_20), v101, v101, float64(1000))
														mBase = m.M
														v105 = m.ExcPending
														if v105 != 0 {
															return
														} else {
															F_MarkGUCPrefixReserved(m, int32(_a_F_HnswInit_21))
															mBase = m.M
															v108 = m.ExcPending
															if v108 != 0 {
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
		}
	} else {
		v47 = F_add_reloption_kind(m)
		mBase = m.M
		v48 = m.ExcPending
		if v48 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_HnswInit[3])) = v47
			F_add_int_reloption(m, v47, int32(_a_F_HnswInit_2), int32(_a_F_HnswInit_3), int32(16), int32(2), int32(100))
			mBase = m.M
			v56 = m.ExcPending
			if v56 != 0 {
				return
			} else {
				v58 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInit[3]))
				F_add_int_reloption(m, v58, int32(_a_F_HnswInit_4), int32(_a_F_HnswInit_5), int32(64), int32(4), int32(1000))
				mBase = m.M
				v65 = m.ExcPending
				if v65 != 0 {
					return
				} else {
					F_DefineCustomIntVariable(m, int32(_a_F_HnswInit_6), int32(_a_F_HnswInit_7), int32(_a_F_HnswInit_8), int32(_a_F_HnswInit_9), int32(40), int32(1), int32(1000), int32(6), int32(0))
					mBase = m.M
					v76 = m.ExcPending
					if v76 != 0 {
						return
					} else {
						v79 = int32(0)
						F_DefineCustomEnumVariable(m, int32(_a_F_HnswInit_10), int32(_a_F_HnswInit_11), v79, int32(_a_F_HnswInit_12), v79, int32(_a_F_HnswInit_13), int32(6))
						mBase = m.M
						v85 = m.ExcPending
						if v85 != 0 {
							return
						} else {
							v88 = int32(0)
							F_DefineCustomIntVariable(m, int32(_a_F_HnswInit_14), int32(_a_F_HnswInit_15), v88, int32(_a_F_HnswInit_16), int32(_a_F_HnswInit_17), int32(1), int32(2147483647), int32(6), v88)
							mBase = m.M
							v96 = m.ExcPending
							if v96 != 0 {
								return
							} else {
								v101 = float64(1)
								F_DefineCustomRealVariable(m, int32(_a_F_HnswInit_18), int32(_a_F_HnswInit_19), int32(0), int32(_a_F_HnswInit_20), v101, v101, float64(1000))
								mBase = m.M
								v105 = m.ExcPending
								if v105 != 0 {
									return
								} else {
									F_MarkGUCPrefixReserved(m, int32(_a_F_HnswInit_21))
									mBase = m.M
									v108 = m.ExcPending
									if v108 != 0 {
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
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
		v18 = F_ShmemInitStruct(m, int32(_a_F_HnswInitLockTranche_0), int32(4), v5+int32(15))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return
		} else {
			v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+15)))
			if v20 == int32(1) {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
				v27 = v23
				*(*int32)(unsafe.Add(mBase, _c_F_HnswInitLockTranche[1])) = v27
				v31 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInitLockTranche[0]))
				F_LWLockRelease(m, v31+int32(2688))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return
				} else {
					v37 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInitLockTranche[1]))
					F_LWLockRegisterTranche(m, v37, int32(_a_F_HnswInitLockTranche_1))
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return
					} else {
						m.G0 = v5 + int32(16)
						return
					}
				}
			} else {
				v24 = F_LWLockNewTrancheId(m)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v18))) = v24
					v27 = v24
					*(*int32)(unsafe.Add(mBase, _c_F_HnswInitLockTranche[1])) = v27
					v31 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInitLockTranche[0]))
					F_LWLockRelease(m, v31+int32(2688))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return
					} else {
						v37 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInitLockTranche[1]))
						F_LWLockRegisterTranche(m, v37, int32(_a_F_HnswInitLockTranche_1))
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
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
func F_HnswLoadElementImpl(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 float64
	_ = v55
	var v57 float64
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
	v13 = F_ReadBuffer(m, l4, l0)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		F_LockBuffer(m, v13, int32(1))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			if v13 < int32(0) {
				v21 = *(*int32)(unsafe.Add(mBase, _c_F_HnswLoadElementImpl[0]))
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v21+(v13^int32(-1))<<(uint(int32(2))%32))))
				v35 = v27
			} else {
				v29 = *(*int32)(unsafe.Add(mBase, _c_F_HnswLoadElementImpl[1]))
				v35 = v29 + v13<<(uint(int32(13))%32) + int32(-8192)
			}
			v39 = *(*int32)(unsafe.Add(mBase, uint32(v35+v2<<(uint(int32(2))%32))+20))
			v42 = v39&int32(_a_F_HnswLoadElementImpl_0) + v35
			v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+2)))
			if v43 == int32(0) {
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
							F_HnswLoadElementFromTuple(m, v80, v42, int32(1), l6)
							mBase = m.M
							v83 = m.ExcPending
							if v83 != 0 {
								return
							} else {
								F_UnlockReleaseBuffer(m, v13)
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
						F_HnswLoadElementFromTuple(m, v80, v42, int32(1), l6)
						mBase = m.M
						v83 = m.ExcPending
						if v83 != 0 {
							return
						} else {
							F_UnlockReleaseBuffer(m, v13)
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
					v48 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
					if v48 != 0 {
						v49 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
						v50 = *(*int32)(unsafe.Add(mBase, uint32(l5)+8))
						v53 = F_FunctionCall2Coll(m, v49, v50, v48, v42+int32(72))
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return
						} else {
							v55 = *(*float64)(unsafe.Add(mBase, uint32(v53)))
							v57 = v55
							*(*float64)(unsafe.Add(mBase, uint32(l2))) = v57
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
										F_HnswLoadElementFromTuple(m, v80, v42, int32(1), l6)
										mBase = m.M
										v83 = m.ExcPending
										if v83 != 0 {
											return
										} else {
											F_UnlockReleaseBuffer(m, v13)
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
									F_HnswLoadElementFromTuple(m, v80, v42, int32(1), l6)
									mBase = m.M
									v83 = m.ExcPending
									if v83 != 0 {
										return
									} else {
										F_UnlockReleaseBuffer(m, v13)
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
								if base.F64_lt(v57, v61) == int32(0) {
									F_UnlockReleaseBuffer(m, v13)
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
											F_HnswLoadElementFromTuple(m, v80, v42, int32(1), l6)
											mBase = m.M
											v83 = m.ExcPending
											if v83 != 0 {
												return
											} else {
												F_UnlockReleaseBuffer(m, v13)
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
										F_HnswLoadElementFromTuple(m, v80, v42, int32(1), l6)
										mBase = m.M
										v83 = m.ExcPending
										if v83 != 0 {
											return
										} else {
											F_UnlockReleaseBuffer(m, v13)
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
						v57 = float64(0)
						*(*float64)(unsafe.Add(mBase, uint32(l2))) = v57
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
									F_HnswLoadElementFromTuple(m, v80, v42, int32(1), l6)
									mBase = m.M
									v83 = m.ExcPending
									if v83 != 0 {
										return
									} else {
										F_UnlockReleaseBuffer(m, v13)
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
								F_HnswLoadElementFromTuple(m, v80, v42, int32(1), l6)
								mBase = m.M
								v83 = m.ExcPending
								if v83 != 0 {
									return
								} else {
									F_UnlockReleaseBuffer(m, v13)
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
							if base.F64_lt(v57, v61) == int32(0) {
								F_UnlockReleaseBuffer(m, v13)
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
										F_HnswLoadElementFromTuple(m, v80, v42, int32(1), l6)
										mBase = m.M
										v83 = m.ExcPending
										if v83 != 0 {
											return
										} else {
											F_UnlockReleaseBuffer(m, v13)
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
									F_HnswLoadElementFromTuple(m, v80, v42, int32(1), l6)
									mBase = m.M
									v83 = m.ExcPending
									if v83 != 0 {
										return
									} else {
										F_UnlockReleaseBuffer(m, v13)
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
								F_sequence_close(m, v24, v23)
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
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 float64
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v62 float64
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int64
	_ = v72
	var v74 int64
	_ = v74
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	v9 = m.G0
	v11 = v9 - int32(224)
	m.G0 = v11
	v13 = F_BuildIndexInfo(m, l1)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+8)))
		*(*uint8)(unsafe.Add(mBase, uint32(v13)+121)) = uint8(v15)
		F_InitBuildState_1(m, v11+int32(16), l0, l1, v13, int32(0))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v11)+220)) = l3
			*(*int32)(unsafe.Add(mBase, uint32(v11)+176)) = l2 + int32(40)
			*(*int32)(unsafe.Add(mBase, uint32(v11)+204)) = int32(_a_F_HnswParallelScanAndInsert_0)
			v29 = v11 + int32(16)
			*(*int32)(unsafe.Add(mBase, uint32(v11)+208)) = v29
			v32 = int32(0)
			v40 = F_table_beginscan_parallel(m, l0, l2+int32(160))
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return
			} else {
				v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
				v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+140))
				v44 = m.T0[v43].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32) float64)(m, l0, l1, v13, int32(1), v32, l4, v32, int32(-1), int32(_a_F_HnswParallelScanAndInsert_1), v29, v40)
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return
				} else {
					v46 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
					*(*int32)(unsafe.Add(mBase, uint32(l2)+24)) = int32(1)
					if v46 != 0 {
						F_s_lock(m, l2+int32(24), int32(_a_F_HnswParallelScanAndInsert_2), int32(818), int32(_a_F_HnswParallelScanAndInsert_3))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return
						} else {
							v56 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(l2)+24)) = v56
							v58 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
							*(*int32)(unsafe.Add(mBase, uint32(l2)+28)) = v58 + int32(1)
							v62 = *(*float64)(unsafe.Add(mBase, uint32(l2)+32))
							*(*float64)(unsafe.Add(mBase, uint32(l2)+32)) = base.F64_add(v62, v44)
							v67 = F_errstart(m, int32(14), v56)
							mBase = m.M
							v68 = m.ExcPending
							if v68 != 0 {
								return
							} else {
								if v67 != 0 {
									if base.F64_lt(base.F64_abs(v44), float64(9.223372036854776e+18)) != 0 {
										v72 = base.I64_trunc_f64_s(v44)
										v74 = v72
									} else {
										v74 = int64(-9223372036854775807 - 1)
									}
									*(*int64)(unsafe.Add(mBase, uint32(v11))) = v74
									if l4 != 0 {
										v78 = int32(_a_F_HnswParallelScanAndInsert_4)
									} else {
										v78 = int32(_a_F_HnswParallelScanAndInsert_5)
									}
									F_errmsg(m, v78, v11)
									mBase = m.M
									v80 = m.ExcPending
									if v80 != 0 {
										return
									} else {
										if l4 != 0 {
											v84 = int32(825)
										} else {
											v84 = int32(827)
										}
										F_errfinish(m, int32(_a_F_HnswParallelScanAndInsert_2), v84, int32(_a_F_HnswParallelScanAndInsert_3))
										mBase = m.M
										v87 = m.ExcPending
										if v87 != 0 {
											return
										} else {
											F_ConditionVariableSignal(m, l2+int32(12))
											mBase = m.M
											v91 = m.ExcPending
											if v91 != 0 {
												return
											} else {
												v92 = *(*int32)(unsafe.Add(mBase, uint32(v11)+196))
												F_MemoryContextDelete(m, v92)
												mBase = m.M
												v94 = m.ExcPending
												if v94 != 0 {
													return
												} else {
													v95 = *(*int32)(unsafe.Add(mBase, uint32(v11)+200))
													F_MemoryContextDelete(m, v95)
													mBase = m.M
													v97 = m.ExcPending
													if v97 != 0 {
														return
													} else {
														m.G0 = v11 + int32(224)
														return
													}
												}
											}
										}
									}
								} else {
									F_ConditionVariableSignal(m, l2+int32(12))
									mBase = m.M
									v91 = m.ExcPending
									if v91 != 0 {
										return
									} else {
										v92 = *(*int32)(unsafe.Add(mBase, uint32(v11)+196))
										F_MemoryContextDelete(m, v92)
										mBase = m.M
										v94 = m.ExcPending
										if v94 != 0 {
											return
										} else {
											v95 = *(*int32)(unsafe.Add(mBase, uint32(v11)+200))
											F_MemoryContextDelete(m, v95)
											mBase = m.M
											v97 = m.ExcPending
											if v97 != 0 {
												return
											} else {
												m.G0 = v11 + int32(224)
												return
											}
										}
									}
								}
							}
						}
					} else {
						v56 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(l2)+24)) = v56
						v58 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
						*(*int32)(unsafe.Add(mBase, uint32(l2)+28)) = v58 + int32(1)
						v62 = *(*float64)(unsafe.Add(mBase, uint32(l2)+32))
						*(*float64)(unsafe.Add(mBase, uint32(l2)+32)) = base.F64_add(v62, v44)
						v67 = F_errstart(m, int32(14), v56)
						mBase = m.M
						v68 = m.ExcPending
						if v68 != 0 {
							return
						} else {
							if v67 != 0 {
								if base.F64_lt(base.F64_abs(v44), float64(9.223372036854776e+18)) != 0 {
									v72 = base.I64_trunc_f64_s(v44)
									v74 = v72
								} else {
									v74 = int64(-9223372036854775807 - 1)
								}
								*(*int64)(unsafe.Add(mBase, uint32(v11))) = v74
								if l4 != 0 {
									v78 = int32(_a_F_HnswParallelScanAndInsert_4)
								} else {
									v78 = int32(_a_F_HnswParallelScanAndInsert_5)
								}
								F_errmsg(m, v78, v11)
								mBase = m.M
								v80 = m.ExcPending
								if v80 != 0 {
									return
								} else {
									if l4 != 0 {
										v84 = int32(825)
									} else {
										v84 = int32(827)
									}
									F_errfinish(m, int32(_a_F_HnswParallelScanAndInsert_2), v84, int32(_a_F_HnswParallelScanAndInsert_3))
									mBase = m.M
									v87 = m.ExcPending
									if v87 != 0 {
										return
									} else {
										F_ConditionVariableSignal(m, l2+int32(12))
										mBase = m.M
										v91 = m.ExcPending
										if v91 != 0 {
											return
										} else {
											v92 = *(*int32)(unsafe.Add(mBase, uint32(v11)+196))
											F_MemoryContextDelete(m, v92)
											mBase = m.M
											v94 = m.ExcPending
											if v94 != 0 {
												return
											} else {
												v95 = *(*int32)(unsafe.Add(mBase, uint32(v11)+200))
												F_MemoryContextDelete(m, v95)
												mBase = m.M
												v97 = m.ExcPending
												if v97 != 0 {
													return
												} else {
													m.G0 = v11 + int32(224)
													return
												}
											}
										}
									}
								}
							} else {
								F_ConditionVariableSignal(m, l2+int32(12))
								mBase = m.M
								v91 = m.ExcPending
								if v91 != 0 {
									return
								} else {
									v92 = *(*int32)(unsafe.Add(mBase, uint32(v11)+196))
									F_MemoryContextDelete(m, v92)
									mBase = m.M
									v94 = m.ExcPending
									if v94 != 0 {
										return
									} else {
										v95 = *(*int32)(unsafe.Add(mBase, uint32(v11)+200))
										F_MemoryContextDelete(m, v95)
										mBase = m.M
										v97 = m.ExcPending
										if v97 != 0 {
											return
										} else {
											m.G0 = v11 + int32(224)
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
func F_hnsw_bit_support(m *base.Module, l0 int32) int32 {
	return int32(_a_F_hnsw_bit_support_0)
}
