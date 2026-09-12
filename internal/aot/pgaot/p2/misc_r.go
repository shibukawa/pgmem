package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_R(m *base.Module, l0 float64) float64 {
	return base.F64_div(base.F64_mul(l0, base.F64_add(base.F64_mul(l0, base.F64_add(base.F64_mul(l0, base.F64_add(base.F64_mul(l0, base.F64_add(base.F64_mul(l0, base.F64_add(base.F64_mul(l0, float64(3.479331075960212e-05)), float64(0.0007915349942898145))), float64(-0.04005553450067941))), float64(0.20121253213486293))), float64(-0.3255658186224009))), float64(0.16666666666666666))), base.F64_add(base.F64_mul(l0, base.F64_add(base.F64_mul(l0, base.F64_add(base.F64_mul(l0, base.F64_add(base.F64_mul(l0, float64(0.07703815055590194)), float64(-0.6882839716054533))), float64(2.0209457602335057))), float64(-2.403394911734414))), float64(1)))
}
func F_ReadDirExtended(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
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
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	if l0 == int32(0) {
		v11 = int32(0)
		v13 = F_errstart(m, l2, v11)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			if v13 == int32(0) {
				v50 = v11
				m.G0 = v7 + int32(16)
				return v50
			} else {
				v38 = int32(291629)
				v39 = int32(3003)
				F_errcode_for_file_access(m)
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = l1
					F_errmsg(m, v38, v7)
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(491162), v39, int32(453897))
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return int32(0)
						} else {
							v50 = int32(0)
							m.G0 = v7 + int32(16)
							return v50
						}
					}
				}
			}
		}
	} else {
		*(*int32)(unsafe.Add(mBase, _consts[155])) = int32(0)
		v24 = F_readdir(m, l0)
		mBase = m.M
		if v24 != 0 {
			v50 = v24
			m.G0 = v7 + int32(16)
			return v50
		} else {
			v25 = int32(0)
			v27 = *(*int32)(unsafe.Add(mBase, _consts[155]))
			if v27 == v25 {
				v50 = v25
				m.G0 = v7 + int32(16)
				return v50
			} else {
				v31 = F_errstart(m, l2, int32(0))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					if v31 == int32(0) {
						v50 = v25
						m.G0 = v7 + int32(16)
						return v50
					} else {
						v38 = int32(291828)
						v39 = int32(3015)
						F_errcode_for_file_access(m)
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = l1
							F_errmsg(m, v38, v7)
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(491162), v39, int32(453897))
								mBase = m.M
								v48 = m.ExcPending
								if v48 != 0 {
									return int32(0)
								} else {
									v50 = int32(0)
									m.G0 = v7 + int32(16)
									return v50
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_RegisterRelcacheInvalidation(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v14 = *(*int32)(unsafe.Add(mBase, _consts[1156]))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v15 < v16 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v99 = F_GetCurrentCommandId(m, int32(1))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L20
	} else {
		goto L23
	}
L2:
	;
	v21 = v15
	goto L5
L3:
	;
	goto L4
L4:
	;
	v49 = *(*int32)(unsafe.Add(mBase, _consts[1157]))
	if v49 <= v16 {
		goto L13
	} else {
		goto L14
	}
L5:
	;
	v28 = v14 + v21<<(uint(int32(4))%32)
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
	if v29 == int32(254) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	goto L4
L7:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
	if v32 == l2 {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v38 = v21 + int32(1)
	if v38 != v16 {
		v21 = v38
		goto L5
	} else {
		goto L12
	}
L10:
	;
	if v32 == int32(0) {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	goto L6
L13:
	;
	if v14 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	v71 = v14
	goto L15
L15:
	;
	v75 = v71 + v16<<(uint(int32(4))%32)
	v76 = int32(254)
	*(*uint8)(unsafe.Add(mBase, uint32(v75))) = uint8(v76)
	v78 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11)+13)))
	*(*uint16)(unsafe.Add(mBase, uint32(v75)+1)) = uint16(v78)
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)))
	*(*uint8)(unsafe.Add(mBase, uint32(v75)+3)) = uint8(v80)
	*(*int32)(unsafe.Add(mBase, uint32(v75)+8)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v75)+4)) = l1
	v85 = l0 + int32(12)
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	*(*int32)(unsafe.Add(mBase, uint32(v85))) = v86 + int32(1)
	goto L1
L16:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1157])) = v65
	*(*int32)(unsafe.Add(mBase, _consts[1156])) = v66
	v71 = v66
	goto L15
L17:
	;
	v55 = *(*int32)(unsafe.Add(mBase, _consts[68]))
	v57 = F_MemoryContextAlloc(m, v55, int32(512))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	goto L19
L19:
	;
	v63 = F_repalloc(m, v14, v49<<(uint(int32(5))%32))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L20
	} else {
		goto L22
	}
L20:
	;
	return
L21:
	;
	v65 = int32(32)
	v66 = v57
	goto L16
L22:
	;
	v65 = v49 << (uint(int32(1)) % 32)
	v66 = v63
	goto L16
L23:
	;
	if l2 != 0 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	m.G0 = v11 + int32(16)
	return
L25:
	;
	v101 = int32(1)
	if base.Ui32(l2-int32(3592)) < base.Ui32(int32(2)) {
		v151 = v101
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L27
L27:
	;
	v155 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v155)
	goto L24
L28:
	;
	if v151 == int32(0) {
		goto L24
	} else {
		goto L45
	}
L29:
	;
	if l2 == int32(2671) {
		v151 = v101
		goto L28
	} else {
		goto L30
	}
L30:
	;
	if l2 == int32(2701) {
		v151 = v101
		goto L28
	} else {
		goto L31
	}
L31:
	;
	v110 = int32(0)
	v116 = *(*int32)(unsafe.Add(mBase, _consts[1158]))
	v118 = v116 - int32(1)
	if v118 < v110 {
		v149 = v110
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v151 = v149
	goto L28
L33:
	;
	goto L32
L34:
	;
	v122 = v110
	v123 = v118
	goto L35
L35:
	;
	v128 = int32(2)
	v129 = base.I32_div_s(v123-v122, v128)
	v130 = v129 + v122
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v130<<(uint(v128)%32))+uint32(_consts[1159])))
	v136 = base.B2i32(v135 == l2)
	if v135 == l2 {
		v149 = v136
		goto L33
	} else {
		goto L37
	}
L36:
	;
	v149 = v136
	goto L33
L37:
	;
	v139 = base.B2i32(base.Ui32(v135) < base.Ui32(l2))
	if base.Ui32(v135) < base.Ui32(l2) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v140 = v130 + int32(1)
	goto L40
L39:
	;
	v140 = v122
	goto L40
L40:
	;
	if base.Ui32(v135) < base.Ui32(l2) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v143 = v123
	goto L43
L42:
	;
	v143 = v130 - int32(1)
	goto L43
L43:
	;
	if v140 <= v143 {
		v122 = v140
		v123 = v143
		goto L35
	} else {
		goto L44
	}
L44:
	;
	goto L36
L45:
	;
	goto L27
}
func F_ReleaseOneSerializableXact(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
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
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v60 int64
	_ = v60
	var v62 int32
	_ = v62
	var v63 int64
	_ = v63
	var v65 int64
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v126 int64
	_ = v126
	var v127 int64
	_ = v127
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v156 int32
	_ = v156
	var v160 int64
	_ = v160
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v180 int32
	_ = v180
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v268 int32
	_ = v268
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v334 int32
	_ = v334
	var v354 int32
	_ = v354
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v376 int32
	_ = v376
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v402 int32
	_ = v402
	var v407 int32
	_ = v407
	var v412 int32
	_ = v412
	v13 = m.G0
	v15 = v13 - int32(32)
	m.G0 = v15
	v18 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v22 = F_LWLockAcquire(m, v18+int32(3840), int32(1))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v28 = *(*int32)(unsafe.Add(mBase, _consts[61]))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+72))
	if v29 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	if v31&int32(1) != 0 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	v31 = int32(1)
	goto L6
L5:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+76)))
	v31 = v30
	goto L6
L6:
	;
	goto L3
L7:
	;
	v37 = F_LWLockAcquire(m, l0+int32(72), int32(0))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v40 = l0 + int32(48)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v41 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L9
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L1
	} else {
		goto L89
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v40
	v200 = *(*int32)(unsafe.Add(mBase, _consts[61]))
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v200)+72))
	if v201 != 0 {
		goto L44
	} else {
		goto L45
	}
L13:
	;
	if v41 == v40 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v52 = v41
	goto L15
L15:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	v60 = *(*int64)(unsafe.Add(mBase, uint32(v52-int32(16))))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+24)) = v60
	v62 = base.I32_wrap_i64(v60)
	v63 = *(*int64)(unsafe.Add(mBase, uint32(v62)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+16)) = v63
	v65 = *(*int64)(unsafe.Add(mBase, uint32(v62)))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+8)) = v65
	v68 = *(*int32)(unsafe.Add(mBase, _consts[817]))
	v71 = F_get_hash_value(m, v68, v15+int32(8))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L17
	}
L16:
	;
	goto L12
L17:
	;
	v74 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v81 = v74 + v71&int32(15)<<(uint(int32(7))%32) + int32(25344)
	v83 = F_LWLockAcquire(m, v81, int32(0))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v86 = v52 - int32(8)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	v88 = int32(4)
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v52-v88)))
	*(*int32)(unsafe.Add(mBase, uint32(v87)+4)) = v90
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	*(*int32)(unsafe.Add(mBase, uint32(v90))) = v92
	v95 = *(*int32)(unsafe.Add(mBase, _consts[818]))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
	v104 = F_hash_search_with_hash_value(m, v95, v15+int32(24), v71^v98<<(uint(v88)%32), int32(2), int32(0))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	if l2 != 0 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	F_LWLockRelease(m, v81)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L41
	}
L21:
	;
	v107 = *(*int32)(unsafe.Add(mBase, _consts[816]))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = v107
	v110 = *(*int32)(unsafe.Add(mBase, _consts[818]))
	v119 = F_hash_search_with_hash_value(m, v110, v15+int32(24), v107<<(uint(int32(4))%32)^v71, int32(3), v15+int32(7))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v62)+20))
	if v162 != v62+int32(16) {
		goto L36
	} else {
		goto L37
	}
L24:
	;
	if v119 == int32(0) {
		goto L11
	} else {
		goto L25
	}
L25:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+7)))
	if v123 == int32(1) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v126 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	v127 = *(*int64)(unsafe.Add(mBase, uint32(v119)+24))
	if base.Ui64(v126) <= base.Ui64(v127) {
		goto L20
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v131 = v119 + int32(8)
	v133 = v62 + int32(16)
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v62)+20))
	if v134 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v119)+24)) = v126
	goto L20
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v62)+20)) = v133
	*(*int32)(unsafe.Add(mBase, uint32(v62)+16)) = v133
	goto L32
L31:
	;
	goto L32
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v119)+12)) = v133
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v133)))
	*(*int32)(unsafe.Add(mBase, uint32(v119)+8)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v140)+4)) = v131
	*(*int32)(unsafe.Add(mBase, uint32(v133))) = v131
	v145 = v119 + int32(16)
	v147 = *(*int32)(unsafe.Add(mBase, _consts[816]))
	v149 = v147 + int32(48)
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v147)+52))
	if v150 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v147)+52)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v147)+48)) = v149
	goto L35
L34:
	;
	goto L35
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v119)+20)) = v149
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v149)))
	*(*int32)(unsafe.Add(mBase, uint32(v119)+16)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v156)+4)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v149))) = v145
	v160 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v119)+24)) = v160
	goto L20
L36:
	;
	v167 = v162
	goto L38
L37:
	;
	v167 = int32(0)
	goto L38
L38:
	;
	if v167 != 0 {
		goto L20
	} else {
		goto L39
	}
L39:
	;
	v169 = *(*int32)(unsafe.Add(mBase, _consts[817]))
	v172 = F_hash_search_with_hash_value(m, v169, v62, v71, int32(2), int32(0))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	goto L20
L41:
	;
	if v57 != v40 {
		v52 = v57
		goto L15
	} else {
		goto L42
	}
L42:
	;
	goto L16
L43:
	;
	if v203&int32(1) != 0 {
		goto L47
	} else {
		goto L48
	}
L44:
	;
	v203 = int32(1)
	goto L46
L45:
	;
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v200)+76)))
	v203 = v202
	goto L46
L46:
	;
	goto L43
L47:
	;
	F_LWLockRelease(m, l0+int32(72))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L1
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	v211 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v211+int32(3840))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L1
	} else {
		goto L51
	}
L50:
	;
	goto L49
L51:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v216
	v219 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v223 = F_LWLockAcquire(m, v219+int32(3584), int32(0))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	if l1 != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v285 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L54:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v225 == int32(0) {
		goto L53
	} else {
		goto L55
	}
L55:
	;
	v229 = l0 + int32(32)
	if v225 == v229 {
		goto L53
	} else {
		goto L56
	}
L56:
	;
	v234 = v225
	goto L57
L57:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v234)+4))
	if l2 != 0 {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	goto L53
L59:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v234)+20))
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v244)+108))
	*(*int32)(unsafe.Add(mBase, uint32(v244)+108)) = v245 | int32(512)
	goto L61
L60:
	;
	goto L61
L61:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v234)+8))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v234)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v250)+4)) = v251
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v234)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v251))) = v253
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v234)))
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v234)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v255)+4)) = v256
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v234)))
	*(*int32)(unsafe.Add(mBase, uint32(v256))) = v258
	v261 = *(*int32)(unsafe.Add(mBase, _consts[810]))
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v261)+4))
	if v262 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v261)+4)) = v261
	*(*int32)(unsafe.Add(mBase, uint32(v261))) = v261
	goto L64
L63:
	;
	goto L64
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v234)+4)) = v261
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v261)))
	*(*int32)(unsafe.Add(mBase, uint32(v234))) = v268
	*(*int32)(unsafe.Add(mBase, uint32(v268)+4)) = v234
	*(*int32)(unsafe.Add(mBase, uint32(v261))) = v234
	if v229 != v243 {
		v234 = v243
		goto L57
	} else {
		goto L65
	}
L65:
	;
	goto L58
L66:
	;
	if l1 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L67:
	;
	v289 = l0 + int32(40)
	if v285 == v289 {
		goto L66
	} else {
		goto L68
	}
L68:
	;
	v294 = v285
	goto L69
L69:
	;
	v304 = v294 + int32(4)
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v304)))
	if l2 != 0 {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	goto L66
L71:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v294)+8))
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v306)+108))
	*(*int32)(unsafe.Add(mBase, uint32(v306)+108)) = v307 | int32(1024)
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v304)))
	v312 = v311
	goto L73
L72:
	;
	v312 = v305
	goto L73
L73:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v294)))
	*(*int32)(unsafe.Add(mBase, uint32(v313)+4)) = v312
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v294)))
	*(*int32)(unsafe.Add(mBase, uint32(v312))) = v315
	v318 = v294 - int32(8)
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v318)))
	v321 = v294 - int32(4)
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v321)))
	*(*int32)(unsafe.Add(mBase, uint32(v319)+4)) = v322
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v318)))
	*(*int32)(unsafe.Add(mBase, uint32(v322))) = v324
	v327 = *(*int32)(unsafe.Add(mBase, _consts[810]))
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v327)+4))
	if v328 == int32(0) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v327)+4)) = v327
	*(*int32)(unsafe.Add(mBase, uint32(v327))) = v327
	goto L76
L75:
	;
	goto L76
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v321))) = v327
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v327)))
	*(*int32)(unsafe.Add(mBase, uint32(v318))) = v334
	*(*int32)(unsafe.Add(mBase, uint32(v334)+4)) = v318
	*(*int32)(unsafe.Add(mBase, uint32(v327))) = v318
	if v289 != v305 {
		v294 = v305
		goto L69
	} else {
		goto L77
	}
L77:
	;
	goto L70
L78:
	;
	if v216 != 0 {
		goto L81
	} else {
		goto L82
	}
L79:
	;
	goto L80
L80:
	;
	v384 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v384+int32(3584))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L1
	} else {
		goto L88
	}
L81:
	;
	v354 = *(*int32)(unsafe.Add(mBase, _consts[819]))
	v359 = F_hash_search(m, v354, v15+int32(8), int32(2), int32(0))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L1
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v362 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v361)+4)) = v362
	v364 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v362))) = v364
	v367 = l0 - int32(-64)
	v369 = *(*int32)(unsafe.Add(mBase, _consts[804]))
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v369)+4))
	if v370 == int32(0) {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	goto L83
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v369)+4)) = v369
	*(*int32)(unsafe.Add(mBase, uint32(v369))) = v369
	goto L87
L86:
	;
	goto L87
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v369
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v369)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v376
	*(*int32)(unsafe.Add(mBase, uint32(v376)+4)) = v367
	*(*int32)(unsafe.Add(mBase, uint32(v369))) = v367
	goto L80
L88:
	;
	m.G0 = v15 + int32(32)
	return
L89:
	;
	F_errcode(m, int32(8389))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	F_errmsg(m, int32(13982), int32(0))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(252026)
	F_errhint(m, int32(641357), v15)
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	F_errfinish(m, int32(490099), int32(3891), int32(109931))
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ReleaseSemaphores(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	v3 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, _consts[511]))
	if v3 < v5 {
		v8 = v3
		for {
			v13 = v8 + int32(1)
			v15 = *(*int32)(unsafe.Add(mBase, _consts[511]))
			if v13 < v15 {
				v8 = v13
				continue
			} else {
				break
			}
			break
		}
	} else {
	}
	return
}
func F_RemoveOldXlogFiles(m *base.Module, l0 int64, l1 int64, l2 int64, l3 int32) {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int64
	_ = v23
	var v24 int64
	_ = v24
	var v27 int64
	_ = v27
	var v28 int64
	_ = v28
	var v31 int64
	_ = v31
	var v34 int32
	_ = v34
	var v36 float64
	_ = v36
	var v38 float64
	_ = v38
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int64
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v62 float64
	_ = v62
	var v68 int64
	_ = v68
	var v70 int64
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int64
	_ = v96
	var v98 int64
	_ = v98
	var v100 int64
	_ = v100
	var v102 int64
	_ = v102
	var v104 int64
	_ = v104
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v124 int32
	_ = v124
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int64
	_ = v191
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v247 int32
	_ = v247
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v277 int64
	_ = v277
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v333 int32
	_ = v333
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v353 int32
	_ = v353
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v418 int64
	_ = v418
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int64
	_ = v431
	var v432 int64
	_ = v432
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v439 int64
	_ = v439
	var v441 int64
	_ = v441
	var v443 int32
	_ = v443
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v453 int64
	_ = v453
	var v461 int32
	_ = v461
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v485 int32
	_ = v485
	v15 = m.G0
	v17 = v15 - int32(144)
	m.G0 = v17
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = int32(0)
	v22 = *(*int32)(unsafe.Add(mBase, _consts[176]))
	v23 = base.I64_extend_i32_s(v22)
	v24 = base.I64_div_u_s(l2, v23)
	*(*int64)(unsafe.Add(mBase, uint32(v17)+56)) = v24
	v27 = base.I64_div_u_s(int64(4294967296), v23)
	v28 = base.I64_div_u_s(l0, v27)
	*(*uint32)(unsafe.Add(mBase, uint32(v17)+36)) = uint32(v28)
	v31 = l0 - v27*v28
	*(*uint32)(unsafe.Add(mBase, uint32(v17)+40)) = uint32(v31)
	v34 = *(*int32)(unsafe.Add(mBase, _consts[216]))
	v36 = *(*float64)(unsafe.Add(mBase, _consts[214]))
	v38 = *(*float64)(unsafe.Add(mBase, _consts[217]))
	v40 = *(*int32)(unsafe.Add(mBase, _consts[218]))
	v47 = F_pg_snprintf(m, v17-int32(-64), int32(64), int32(501258), v17+int32(32))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v49 = base.I64_div_u_s(l1, v23)
	v51 = base.I32_div_s(v22, int32(1048576))
	v52 = base.I32_div_s(v40, v51)
	v62 = base.F64_ceil(base.F64_div(base.F64_add(base.F64_mul(base.F64_mul(v36, base.F64_add(v38, float64(1))), float64(1.1)), base.F64_convert_i64_u(l1)), base.F64_convert_i32_s(v22)))
	if base.F64_lt(v62, float64(1.8446744073709552e+19))&base.F64_ge(v62, float64(0)) != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v71 = base.I32_div_s(v34, v51)
	v74 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L7
	}
L4:
	;
	v68 = base.I64_trunc_f64_u(v62)
	v70 = v68
	goto L3
L5:
	;
	goto L6
L6:
	;
	v70 = int64(0)
	goto L3
L7:
	;
	if v74 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v17 - int32(-64)
	F_errmsg_internal(m, int32(192080), v17+int32(16))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v90 = F_AllocateDir(m, int32(303460))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L13
	}
L11:
	;
	F_errfinish(m, int32(489567), int32(3884), int32(162535))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	v93 = F_ReadDir(m, v90, int32(303460))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	if v93 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v96 = v49 - int64(1)
	v98 = v96 + base.I64_extend_i32_s(v52)
	if base.Ui64(v70) < base.Ui64(v98) {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L17
L17:
	;
	F_FreeDir(m, v90)
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L1
	} else {
		goto L122
	}
L18:
	;
	v100 = v98
	goto L20
L19:
	;
	v100 = v70
	goto L20
L20:
	;
	v102 = v96 + base.I64_extend_i32_s(v71)
	if base.Ui64(v100) < base.Ui64(v102) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v104 = v100
	goto L23
L22:
	;
	v104 = v102
	goto L23
L23:
	;
	v108 = v17 - int32(-64) | int32(8)
	v115 = v93
	goto L24
L24:
	;
	v124 = v115 + int32(19)
	if v124&int32(3) == int32(0) {
		v148 = v124
		goto L32
	} else {
		goto L33
	}
L25:
	;
	goto L17
L26:
	;
	v468 = F_ReadDir(m, v90, int32(303460))
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L1
	} else {
		goto L120
	}
L27:
	;
	v385 = v115 + int32(27)
	v388 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108))))
	v389 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v385))))
	if v389 == int32(0) {
		v408 = v388
		v409 = v389
		goto L101
	} else {
		goto L102
	}
L28:
	;
	v270 = int32(528047)
	v274 = m.G0
	v276 = v274 - int32(32)
	v277 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v276)+24)) = v277
	*(*int64)(unsafe.Add(mBase, uint32(v276)+16)) = v277
	*(*int64)(unsafe.Add(mBase, uint32(v276)+8)) = v277
	*(*int64)(unsafe.Add(mBase, uint32(v276))) = v277
	v285 = int32(*(*uint8)(unsafe.Add(mBase, _consts[219])))
	if v285 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L29:
	;
	v184 = int32(528047)
	v188 = m.G0
	v190 = v188 - int32(32)
	v191 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v190)+24)) = v191
	*(*int64)(unsafe.Add(mBase, uint32(v190)+16)) = v191
	*(*int64)(unsafe.Add(mBase, uint32(v190)+8)) = v191
	*(*int64)(unsafe.Add(mBase, uint32(v190))) = v191
	v199 = int32(*(*uint8)(unsafe.Add(mBase, _consts[219])))
	if v199 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L30:
	;
	switch v181 - int32(24) {
	case 0:
		goto L29
	default:
		goto L26
	case 8:
		goto L28
	}
L31:
	;
	v181 = v173 - v124
	goto L30
L32:
	;
	v152 = v148
	goto L41
L33:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124))))
	if v132 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v181 = int32(0)
	goto L30
L35:
	;
	goto L36
L36:
	;
	v137 = v124
	goto L37
L37:
	;
	v141 = v137 + int32(1)
	if v141&int32(3) == int32(0) {
		v148 = v141
		goto L32
	} else {
		goto L39
	}
L38:
	;
	v173 = v141
	goto L31
L39:
	;
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141))))
	if v146 != 0 {
		v137 = v141
		goto L37
	} else {
		goto L40
	}
L40:
	;
	goto L38
L41:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v152)))
	v161 = int32(-2139062144)
	if (int32(16843008)-v158|v158)&v161 == v161 {
		v152 = v152 + int32(4)
		goto L41
	} else {
		goto L43
	}
L42:
	;
	v167 = v152
	goto L44
L43:
	;
	goto L42
L44:
	;
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167))))
	if v171 != 0 {
		v167 = v167 + int32(1)
		goto L44
	} else {
		goto L46
	}
L45:
	;
	v173 = v167
	goto L31
L46:
	;
	goto L45
L47:
	;
	if v267 == int32(24) {
		goto L27
	} else {
		goto L68
	}
L48:
	;
	v267 = int32(0)
	goto L47
L49:
	;
	goto L50
L50:
	;
	v203 = int32(*(*uint8)(unsafe.Add(mBase, _consts[220])))
	if v203 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v207 = v124
	goto L54
L52:
	;
	goto L53
L53:
	;
	v217 = v184
	v218 = v199
	goto L57
L54:
	;
	v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207))))
	if v213 == v199 {
		v207 = v207 + int32(1)
		goto L54
	} else {
		goto L56
	}
L55:
	;
	v267 = v207 - v124
	goto L47
L56:
	;
	goto L55
L57:
	;
	v225 = v190 + int32(base.Ui32(v218)>>(uint(int32(3))%32))&int32(28)
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v225)))
	v227 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v225))) = v226 | v227<<(uint(v218)%32)
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217)+1)))
	if v231 != 0 {
		v217 = v217 + v227
		v218 = v231
		goto L57
	} else {
		goto L59
	}
L58:
	;
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124))))
	if v234 == int32(0) {
		v259 = v124
		goto L60
	} else {
		goto L61
	}
L59:
	;
	goto L58
L60:
	;
	v267 = v259 - v124
	goto L47
L61:
	;
	v238 = v124
	v239 = v234
	goto L62
L62:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v190+int32(base.Ui32(v239)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v247)>>(uint(v239)%32))&int32(1) == int32(0) {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	v259 = v255
	goto L60
L64:
	;
	v259 = v238
	goto L60
L65:
	;
	goto L66
L66:
	;
	v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v238)+1)))
	v255 = v238 + int32(1)
	if v253 != 0 {
		v238 = v255
		v239 = v253
		goto L62
	} else {
		goto L67
	}
L67:
	;
	goto L63
L68:
	;
	goto L26
L69:
	;
	if v353 != int32(24) {
		goto L26
	} else {
		goto L90
	}
L70:
	;
	v353 = int32(0)
	goto L69
L71:
	;
	goto L72
L72:
	;
	v289 = int32(*(*uint8)(unsafe.Add(mBase, _consts[220])))
	if v289 == int32(0) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v293 = v124
	goto L76
L74:
	;
	goto L75
L75:
	;
	v303 = v270
	v304 = v285
	goto L79
L76:
	;
	v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v293))))
	if v299 == v285 {
		v293 = v293 + int32(1)
		goto L76
	} else {
		goto L78
	}
L77:
	;
	v353 = v293 - v124
	goto L69
L78:
	;
	goto L77
L79:
	;
	v311 = v276 + int32(base.Ui32(v304)>>(uint(int32(3))%32))&int32(28)
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v311)))
	v313 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v311))) = v312 | v313<<(uint(v304)%32)
	v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v303)+1)))
	if v317 != 0 {
		v303 = v303 + v313
		v304 = v317
		goto L79
	} else {
		goto L81
	}
L80:
	;
	v320 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124))))
	if v320 == int32(0) {
		v345 = v124
		goto L82
	} else {
		goto L83
	}
L81:
	;
	goto L80
L82:
	;
	v353 = v345 - v124
	goto L69
L83:
	;
	v324 = v124
	v325 = v320
	goto L84
L84:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v276+int32(base.Ui32(v325)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v333)>>(uint(v325)%32))&int32(1) == int32(0) {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	v345 = v341
	goto L82
L86:
	;
	v345 = v324
	goto L82
L87:
	;
	goto L88
L88:
	;
	v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v324)+1)))
	v341 = v324 + int32(1)
	if v339 != 0 {
		v324 = v341
		v325 = v339
		goto L84
	} else {
		goto L89
	}
L89:
	;
	goto L85
L90:
	;
	v357 = v115 + int32(43)
	v358 = int32(308703)
	v361 = int32(*(*uint8)(unsafe.Add(mBase, _consts[221])))
	v362 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v357))))
	if v362 == int32(0) {
		v381 = v361
		v382 = v362
		goto L92
	} else {
		goto L93
	}
L91:
	;
	if v382-v381 != 0 {
		goto L26
	} else {
		goto L99
	}
L92:
	;
	goto L91
L93:
	;
	if v361 != v362 {
		v381 = v361
		v382 = v362
		goto L92
	} else {
		goto L94
	}
L94:
	;
	v366 = v357
	v367 = v358
	goto L95
L95:
	;
	v370 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v367)+1)))
	v371 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v366)+1)))
	if v371 == int32(0) {
		v381 = v370
		v382 = v371
		goto L92
	} else {
		goto L97
	}
L96:
	;
	v381 = v370
	v382 = v371
	goto L92
L97:
	;
	v374 = int32(1)
	if v370 == v371 {
		v366 = v366 + v374
		v367 = v367 + v374
		goto L95
	} else {
		goto L98
	}
L98:
	;
	goto L96
L99:
	;
	goto L27
L100:
	;
	if int32(0) < v409-v408 {
		goto L26
	} else {
		goto L108
	}
L101:
	;
	goto L100
L102:
	;
	if v388 != v389 {
		v408 = v388
		v409 = v389
		goto L101
	} else {
		goto L103
	}
L103:
	;
	v393 = v385
	v394 = v108
	goto L104
L104:
	;
	v397 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v394)+1)))
	v398 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v393)+1)))
	if v398 == int32(0) {
		v408 = v397
		v409 = v398
		goto L101
	} else {
		goto L106
	}
L105:
	;
	v408 = v397
	v409 = v398
	goto L101
L106:
	;
	v401 = int32(1)
	if v397 == v398 {
		v393 = v393 + v401
		v394 = v394 + v401
		goto L104
	} else {
		goto L107
	}
L107:
	;
	goto L105
L108:
	;
	v413 = F_XLogArchiveCheckDone(m, v124)
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	if v413 == int32(0) {
		goto L26
	} else {
		goto L110
	}
L110:
	;
	v418 = int64(*(*int32)(unsafe.Add(mBase, _consts[176])))
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v17 + int32(132)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v17 + int32(140)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = v17 + int32(136)
	v429 = F_sscanf(m, v124, int32(501258), v17)
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	v431 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v17)+136)))
	v432 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v17)+140)))
	v434 = *(*int32)(unsafe.Add(mBase, _consts[175]))
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v434)+440))
	*(*int32)(unsafe.Add(mBase, uint32(v434)+440)) = int32(1)
	v439 = base.I64_div_u_s(int64(4294967296), v418)
	v441 = v431 + v432*v439
	if v435 != 0 {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v443 = *(*int32)(unsafe.Add(mBase, _consts[175]))
	F_s_lock(m, v443+int32(440), int32(489567), int32(3817), int32(203536))
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L1
	} else {
		goto L115
	}
L113:
	;
	goto L114
L114:
	;
	v452 = *(*int32)(unsafe.Add(mBase, _consts[175]))
	v453 = *(*int64)(unsafe.Add(mBase, uint32(v452)+232))
	if base.Ui64(v453) < base.Ui64(v441) {
		goto L116
	} else {
		goto L117
	}
L115:
	;
	goto L114
L116:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v452)+232)) = v441
	goto L118
L117:
	;
	goto L118
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v452)+440)) = int32(0)
	F_RemoveXlogFile(m, v115, v104, v17+int32(56), l3)
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	goto L26
L120:
	;
	if v468 != 0 {
		v115 = v468
		goto L24
	} else {
		goto L121
	}
L121:
	;
	goto L25
L122:
	;
	m.G0 = v17 + int32(144)
	return
}
func F_ResetUsage(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v29 int32
	_ = v29
	var v30 int64
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int64
	_ = v36
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	v1 = int32(4394200)
	v2 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = int32(4394216)
	v16 = F___memset(m, int32(4394224), v2, int32(144))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _consts[845])) = int32(4)
	*(*int64)(unsafe.Add(mBase, _consts[846])) = int64(3)
	*(*int32)(unsafe.Add(mBase, _consts[831])) = int32(2)
	*(*int64)(unsafe.Add(mBase, _consts[834])) = int64(1)
	v29 = F___memcpy(m, v8, v11, int32(16))
	mBase = m.M
	v30 = int64(*(*int32)(unsafe.Add(mBase, uint32(v8))))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v32 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[847])) = v32
	*(*int32)(unsafe.Add(mBase, _consts[830])) = v31
	*(*int64)(unsafe.Add(mBase, _consts[833])) = v30
	v36 = int64(*(*int32)(unsafe.Add(mBase, uint32(v8)+8)))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	*(*int32)(unsafe.Add(mBase, _consts[848])) = v32
	*(*int32)(unsafe.Add(mBase, _consts[831])) = v37
	*(*int64)(unsafe.Add(mBase, _consts[834])) = v36
	v44 = F___syscall_ret(m, v2)
	mBase = m.M
	m.G0 = v8 + int32(16)
	F___gettimeofday(m, int32(4394352))
	mBase = m.M
	return
}
func F_RestoreArchivedFile(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int64
	_ = v80
	var v82 int64
	_ = v82
	var v83 int64
	_ = v83
	var v85 int64
	_ = v85
	var v86 int64
	_ = v86
	var v89 int64
	_ = v89
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v201 int32
	_ = v201
	var v206 int64
	_ = v206
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v357 int32
	_ = v357
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v381 int32
	_ = v381
	var v385 int32
	_ = v385
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v418 int32
	_ = v418
	var v423 int32
	_ = v423
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v440 int32
	_ = v440
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v461 int32
	_ = v461
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	v6 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(2336)
	m.G0 = v15
	v18 = int32(*(*uint8)(unsafe.Add(mBase, _consts[227])))
	if v18 != int32(1) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_proc_exit(m, int32(1))
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L8
	} else {
		goto L136
	}
L2:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L8
	} else {
		goto L132
	}
L3:
	;
	m.G0 = v15 + int32(2336)
	return v440
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = l1
	v434 = F_pg_snprintf(m, l0, int32(1024), int32(174240), v15)
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L8
	} else {
		goto L131
	}
L5:
	;
	v22 = *(*int32)(unsafe.Add(mBase, _consts[228]))
	if v22 == int32(0) {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	if v25 == int32(0) {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+176)) = l2
	v35 = F_pg_snprintf(m, v15+int32(1312), int32(1024), int32(174240), v15+int32(176))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return int32(0)
L9:
	;
	v45 = F___fstatat(m, int32(-100), v15+int32(1312), v15+int32(192), int32(0))
	mBase = m.M
	goto L11
L10:
	;
	if l4 != 0 {
		goto L22
	} else {
		goto L23
	}
L11:
	;
	if v45 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v47 = *(*int32)(unsafe.Add(mBase, _consts[155]))
	if v47 == int32(44) {
		goto L10
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v71 = F_unlink(m, v15+int32(1312))
	mBase = m.M
	if v71 != 0 {
		goto L2
	} else {
		goto L20
	}
L15:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L8
	} else {
		goto L16
	}
L16:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L8
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+160)) = v15 + int32(1312)
	F_errmsg(m, int32(292901), v15+int32(160))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L8
	} else {
		goto L18
	}
L18:
	;
	F_errfinish(m, int32(489869), int32(112), int32(383584))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L8
	} else {
		goto L19
	}
L19:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L20:
	;
	goto L10
L21:
	;
	v115 = *(*int32)(unsafe.Add(mBase, _consts[228]))
	v117 = v15 + int32(288)
	v118 = m.G0
	v120 = v118 - int32(32)
	m.G0 = v120
	v123 = v15 + int32(1312)
	if v123 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L22:
	;
	F_GetOldestRestartPoint(m, v15+int32(184), v15+int32(180))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L8
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+136)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+128)) = int64(0)
	v109 = F_pg_snprintf(m, v15+int32(288), int32(64), int32(501258), v15+int32(128))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L8
	} else {
		goto L27
	}
L25:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v15)+180))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+112)) = v78
	v80 = *(*int64)(unsafe.Add(mBase, uint32(v15)+184))
	v82 = int64(*(*int32)(unsafe.Add(mBase, _consts[176])))
	v83 = base.I64_div_u_s(v80, v82)
	v85 = base.I64_div_u_s(int64(4294967296), v82)
	v86 = base.I64_div_u_s(v83, v85)
	*(*uint32)(unsafe.Add(mBase, uint32(v15)+116)) = uint32(v86)
	v89 = v83 - v85*v86
	*(*uint32)(unsafe.Add(mBase, uint32(v15)+120)) = uint32(v89)
	v97 = F_pg_snprintf(m, v15+int32(288), int32(64), int32(501258), v15+int32(112))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L8
	} else {
		goto L26
	}
L26:
	;
	goto L21
L27:
	;
	goto L21
L28:
	;
	m.G0 = v120 + int32(32)
	v156 = F_errstart(m, int32(12), int32(0))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L8
	} else {
		goto L37
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v120)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v120)+4)) = v117
	*(*int32)(unsafe.Add(mBase, uint32(v120))) = l1
	v132 = F_replace_percent_placeholders(m, v115, int32(420582), int32(230427), v120)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L8
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v134 = F_pstrdup(m, v123)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L8
	} else {
		goto L33
	}
L32:
	;
	v149 = v132
	goto L28
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v120)+24)) = v134
	*(*int32)(unsafe.Add(mBase, uint32(v120)+20)) = v117
	*(*int32)(unsafe.Add(mBase, uint32(v120)+16)) = l1
	v143 = F_replace_percent_placeholders(m, v115, int32(420582), int32(230427), v120+int32(16))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L8
	} else {
		goto L34
	}
L34:
	;
	if v134 == int32(0) {
		v149 = v143
		goto L28
	} else {
		goto L35
	}
L35:
	;
	F_pfree(m, v134)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L8
	} else {
		goto L36
	}
L36:
	;
	v149 = v143
	goto L28
L37:
	;
	if v156 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+96)) = v149
	F_errmsg_internal(m, int32(697108), v15+int32(96))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L8
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v170 = F_fflush(m, int32(0))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L8
	} else {
		goto L43
	}
L41:
	;
	F_errfinish(m, int32(489869), int32(159), int32(383584))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L8
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	v173 = *(*int32)(unsafe.Add(mBase, _consts[189]))
	*(*int32)(unsafe.Add(mBase, uint32(v173))) = int32(134217778)
	*(*int32)(unsafe.Add(mBase, _consts[229])) = int32(1)
	v180 = *(*int32)(unsafe.Add(mBase, _consts[230]))
	if v180 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	F_proc_exit(m, int32(1))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L8
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v184 = F_pgl_system(m, v149)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L8
	} else {
		goto L48
	}
L47:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L48:
	;
	v187 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[229])) = v187
	v190 = *(*int32)(unsafe.Add(mBase, _consts[189]))
	*(*int32)(unsafe.Add(mBase, uint32(v190))) = v187
	F_pfree(m, v149)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L8
	} else {
		goto L49
	}
L49:
	;
	if v184 != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v361 = int32(1)
	v363 = v184 & int32(127)
	if int32(15) == v363 {
		goto L107
	} else {
		goto L108
	}
L51:
	;
	v201 = F___fstatat(m, int32(-100), v15+int32(1312), v15+int32(192), int32(0))
	mBase = m.M
	goto L52
L52:
	;
	if v201 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	if l3 <= int64(0) {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	goto L55
L55:
	;
	v330 = *(*int32)(unsafe.Add(mBase, _consts[155]))
	if v330 == int32(44) {
		goto L96
	} else {
		goto L97
	}
L56:
	;
	v237 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L8
	} else {
		goto L69
	}
L57:
	;
	v206 = *(*int64)(unsafe.Add(mBase, uint32(v15)+216))
	if v206 == l3 {
		goto L56
	} else {
		goto L58
	}
L58:
	;
	v208 = int32(22)
	v212 = int32(*(*uint8)(unsafe.Add(mBase, _consts[231])))
	if v212 != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v213 = int32(14)
	goto L61
L60:
	;
	v213 = v208
	goto L61
L61:
	;
	if l3 <= v206 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v215 = v208
	goto L64
L63:
	;
	v215 = v213
	goto L64
L64:
	;
	v217 = F_errstart(m, v215, int32(0))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L8
	} else {
		goto L65
	}
L65:
	;
	if v217 == int32(0) {
		v440 = v6
		goto L3
	} else {
		goto L66
	}
L66:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v15)+48)) = l3
	*(*int64)(unsafe.Add(mBase, uint32(v15)+40)) = v206
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = l1
	F_errmsg(m, int32(422925), v15+int32(32))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L8
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(489869), int32(216), int32(383584))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L8
	} else {
		goto L68
	}
L68:
	;
	v440 = v6
	goto L3
L69:
	;
	if v237 != 0 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = l1
	F_errmsg(m, int32(338052), v15+int32(16))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L8
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	v251 = v15 + int32(1312)
	if (v251^l0)&int32(3) != 0 {
		goto L78
	} else {
		goto L79
	}
L73:
	;
	F_errfinish(m, int32(489869), int32(223), int32(383584))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L8
	} else {
		goto L74
	}
L74:
	;
	goto L72
L75:
	;
	v440 = int32(1)
	goto L3
L76:
	;
	goto L75
L77:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v306))) = uint8(v305)
	if v305&int32(255) == int32(0) {
		goto L76
	} else {
		goto L92
	}
L78:
	;
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v251))))
	v304 = v251
	v305 = v257
	v306 = l0
	goto L77
L79:
	;
	goto L80
L80:
	;
	if v251&int32(3) != 0 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v261 = v251
	v263 = l0
	goto L84
L82:
	;
	v275 = v251
	v277 = l0
	goto L83
L83:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v275)))
	v282 = int32(-2139062144)
	if (int32(16843008)-v279|v279)&v282 != v282 {
		v304 = v275
		v305 = v279
		v306 = v277
		goto L77
	} else {
		goto L88
	}
L84:
	;
	v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v261))))
	*(*uint8)(unsafe.Add(mBase, uint32(v263))) = uint8(v264)
	if v264 == int32(0) {
		goto L76
	} else {
		goto L86
	}
L85:
	;
	v275 = v271
	v277 = v269
	goto L83
L86:
	;
	v268 = int32(1)
	v269 = v263 + v268
	v271 = v261 + v268
	if v271&int32(3) != 0 {
		v261 = v271
		v263 = v269
		goto L84
	} else {
		goto L87
	}
L87:
	;
	goto L85
L88:
	;
	v287 = v275
	v288 = v279
	v289 = v277
	goto L89
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v289))) = v288
	v291 = int32(4)
	v292 = v289 + v291
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v287)+4))
	v295 = v287 + v291
	v299 = int32(-2139062144)
	if (v293|(int32(16843008)-v293))&v299 == v299 {
		v287 = v295
		v288 = v293
		v289 = v292
		goto L89
	} else {
		goto L91
	}
L90:
	;
	v304 = v295
	v305 = v293
	v306 = v292
	goto L77
L91:
	;
	goto L90
L92:
	;
	v313 = v304
	v315 = v306
	goto L93
L93:
	;
	v316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v313)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v315)+1)) = uint8(v316)
	v318 = int32(1)
	if v316 != 0 {
		v313 = v313 + v318
		v315 = v315 + v318
		goto L93
	} else {
		goto L95
	}
L94:
	;
	goto L76
L95:
	;
	goto L94
L96:
	;
	v333 = int32(15)
	goto L98
L97:
	;
	v333 = int32(22)
	goto L98
L98:
	;
	v335 = F_errstart(m, v333, int32(0))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L8
	} else {
		goto L99
	}
L99:
	;
	if v335 == int32(0) {
		goto L50
	} else {
		goto L100
	}
L100:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L8
	} else {
		goto L101
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+80)) = v15 + int32(1312)
	F_errmsg(m, int32(292901), v15+int32(80))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L8
	} else {
		goto L102
	}
L102:
	;
	F_errdetail(m, int32(624581), int32(0))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L8
	} else {
		goto L103
	}
L103:
	;
	F_errfinish(m, int32(489869), int32(236), int32(383584))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L8
	} else {
		goto L104
	}
L104:
	;
	goto L50
L105:
	;
	if v381 != 0 {
		goto L1
	} else {
		goto L115
	}
L106:
	;
	goto L105
L107:
	;
	if base.Ui32(v184&int32(65535)-int32(1)) < base.Ui32(int32(255)) {
		v381 = v361
		goto L106
	} else {
		goto L110
	}
L108:
	;
	goto L109
L109:
	;
	if v363 == int32(0) {
		goto L111
	} else {
		goto L112
	}
L110:
	;
	goto L109
L111:
	;
	if int32(143) == int32(base.Ui32(v184)>>(uint(int32(8))%32))&int32(255) {
		v381 = v361
		goto L106
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	v381 = int32(0)
	goto L106
L114:
	;
	goto L113
L115:
	;
	v385 = int32(1)
	if base.Ui32(v184&int32(65535)-v385) < base.Ui32(int32(255)) {
		v403 = v385
		goto L117
	} else {
		goto L118
	}
L116:
	;
	if v403 != 0 {
		goto L123
	} else {
		goto L124
	}
L117:
	;
	goto L116
L118:
	;
	if v184&int32(127) == int32(0) {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	if base.Ui32(int32(125)) < base.Ui32(int32(base.Ui32(v184)>>(uint(int32(8))%32))&int32(255)) {
		v403 = v385
		goto L117
	} else {
		goto L122
	}
L120:
	;
	goto L121
L121:
	;
	v403 = int32(0)
	goto L117
L122:
	;
	goto L121
L123:
	;
	v404 = int32(22)
	goto L125
L124:
	;
	v404 = int32(13)
	goto L125
L125:
	;
	v406 = F_errstart(m, v404, int32(0))
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L8
	} else {
		goto L126
	}
L126:
	;
	if v406 == int32(0) {
		goto L4
	} else {
		goto L127
	}
L127:
	;
	v410 = F_wait_result_to_str(m, v184)
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L8
	} else {
		goto L128
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v410
	*(*int32)(unsafe.Add(mBase, uint32(v15)+64)) = l1
	F_errmsg(m, int32(199885), v15-int32(-64))
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L8
	} else {
		goto L129
	}
L129:
	;
	F_errfinish(m, int32(489869), int32(269), int32(383584))
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L8
	} else {
		goto L130
	}
L130:
	;
	goto L4
L131:
	;
	v440 = v6
	goto L3
L132:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L8
	} else {
		goto L133
	}
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+144)) = v15 + int32(1312)
	F_errmsg(m, int32(294353), v15+int32(144))
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L8
	} else {
		goto L134
	}
L134:
	;
	F_errfinish(m, int32(489869), int32(120), int32(383584))
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L8
	} else {
		goto L135
	}
L135:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L136:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_RestrictInfoIsTidQual(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
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
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	v4 = int32(0)
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+10)))
	if v8 != 0 {
		v87 = v4
		return v87
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l2)+208))
		if base.Ui32(v12) < base.Ui32(v11) {
			v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+13)))
			v15 = v14
		} else {
			v15 = int32(1)
		}
		if v15&int32(1) == int32(0) {
			v87 = v4
			return v87
		} else {
			v20 = F_IsBinaryTidClause(m, l1, l2)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				if v20 != 0 {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
					if v25 != int32(387) {
						v32 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
						if v32 == int32(20) {
							v35 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
							if v35 != int32(387) {
								v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
								if v72 == int32(0) {
									v87 = v4
								} else {
									v75 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
									v76 = v72
									v79 = v75
									if v79 != int32(58) {
										v87 = v4
									} else {
										v82 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
										v83 = *(*int32)(unsafe.Add(mBase, uint32(l2)+68))
										v87 = base.B2i32(v82 == v83)
									}
								}
								return v87
							} else {
								v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+20)))
								if v38 != int32(1) {
									v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
									if v72 == int32(0) {
										v87 = v4
									} else {
										v75 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
										v76 = v72
										v79 = v75
										if v79 != int32(58) {
											v87 = v4
										} else {
											v82 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
											v83 = *(*int32)(unsafe.Add(mBase, uint32(l2)+68))
											v87 = base.B2i32(v82 == v83)
										}
									}
									return v87
								} else {
									v41 = *(*int32)(unsafe.Add(mBase, uint32(v24)+28))
									v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+12))
									v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
									if v43 == int32(0) {
										v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
										if v72 == int32(0) {
											v87 = v4
										} else {
											v75 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
											v76 = v72
											v79 = v75
											if v79 != int32(58) {
												v87 = v4
											} else {
												v82 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
												v83 = *(*int32)(unsafe.Add(mBase, uint32(l2)+68))
												v87 = base.B2i32(v82 == v83)
											}
										}
										return v87
									} else {
										v46 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
										if v46 != int32(6) {
											v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
											if v72 == int32(0) {
												v87 = v4
											} else {
												v75 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
												v76 = v72
												v79 = v75
												if v79 != int32(58) {
													v87 = v4
												} else {
													v82 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
													v83 = *(*int32)(unsafe.Add(mBase, uint32(l2)+68))
													v87 = base.B2i32(v82 == v83)
												}
											}
											return v87
										} else {
											v49 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v43)+8)))
											if v49 != int32(65535) {
												v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
												if v72 == int32(0) {
													v87 = v4
												} else {
													v75 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
													v76 = v72
													v79 = v75
													if v79 != int32(58) {
														v87 = v4
													} else {
														v82 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
														v83 = *(*int32)(unsafe.Add(mBase, uint32(l2)+68))
														v87 = base.B2i32(v82 == v83)
													}
												}
												return v87
											} else {
												v52 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
												if v52 != int32(27) {
													v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
													if v72 == int32(0) {
														v87 = v4
													} else {
														v75 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
														v76 = v72
														v79 = v75
														if v79 != int32(58) {
															v87 = v4
														} else {
															v82 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
															v83 = *(*int32)(unsafe.Add(mBase, uint32(l2)+68))
															v87 = base.B2i32(v82 == v83)
														}
													}
													return v87
												} else {
													v55 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
													v56 = *(*int32)(unsafe.Add(mBase, uint32(l2)+68))
													if v55 != v56 {
														v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
														if v72 == int32(0) {
															v87 = v4
														} else {
															v75 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
															v76 = v72
															v79 = v75
															if v79 != int32(58) {
																v87 = v4
															} else {
																v82 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
																v83 = *(*int32)(unsafe.Add(mBase, uint32(l2)+68))
																v87 = base.B2i32(v82 == v83)
															}
														}
														return v87
													} else {
														v58 = *(*int32)(unsafe.Add(mBase, uint32(v43)+24))
														if v58 != 0 {
															v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
															if v72 == int32(0) {
																v87 = v4
															} else {
																v75 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
																v76 = v72
																v79 = v75
																if v79 != int32(58) {
																	v87 = v4
																} else {
																	v82 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
																	v83 = *(*int32)(unsafe.Add(mBase, uint32(l2)+68))
																	v87 = base.B2i32(v82 == v83)
																}
															}
															return v87
														} else {
															v59 = *(*int32)(unsafe.Add(mBase, uint32(v43)+28))
															if v59 != 0 {
																v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
																if v72 == int32(0) {
																	v87 = v4
																} else {
																	v75 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
																	v76 = v72
																	v79 = v75
																	if v79 != int32(58) {
																		v87 = v4
																	} else {
																		v82 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
																		v83 = *(*int32)(unsafe.Add(mBase, uint32(l2)+68))
																		v87 = base.B2i32(v82 == v83)
																	}
																}
																return v87
															} else {
																v60 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
																v61 = F_pull_varnos(m, l0, v60)
																mBase = m.M
																v62 = m.ExcPending
																if v62 != 0 {
																	return int32(0)
																} else {
																	v63 = F_bms_is_member(m, v55, v61)
																	mBase = m.M
																	v64 = m.ExcPending
																	if v64 != 0 {
																		return int32(0)
																	} else {
																		if v63 != 0 {
																			v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
																			if v72 == int32(0) {
																				v87 = v4
																			} else {
																				v75 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
																				v76 = v72
																				v79 = v75
																				if v79 != int32(58) {
																					v87 = v4
																				} else {
																					v82 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
																					v83 = *(*int32)(unsafe.Add(mBase, uint32(l2)+68))
																					v87 = base.B2i32(v82 == v83)
																				}
																			}
																			return v87
																		} else {
																			v65 = F_contain_volatile_functions(m, v60)
																			mBase = m.M
																			v66 = m.ExcPending
																			if v66 != 0 {
																				return int32(0)
																			} else {
																				if v65 != 0 {
																					v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
																					if v72 == int32(0) {
																						v87 = v4
																					} else {
																						v75 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
																						v76 = v72
																						v79 = v75
																						if v79 != int32(58) {
																							v87 = v4
																						} else {
																							v82 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
																							v83 = *(*int32)(unsafe.Add(mBase, uint32(l2)+68))
																							v87 = base.B2i32(v82 == v83)
																						}
																					}
																					return v87
																				} else {
																					return int32(1)
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
							v76 = v24
							v79 = v32
							if v79 != int32(58) {
								v87 = v4
							} else {
								v82 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
								v83 = *(*int32)(unsafe.Add(mBase, uint32(l2)+68))
								v87 = base.B2i32(v82 == v83)
							}
							return v87
						}
					} else {
						return int32(1)
					}
				} else {
					if v24 == int32(0) {
						v87 = v4
						return v87
					} else {
						v32 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
						if v32 == int32(20) {
							v35 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
							if v35 != int32(387) {
								v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
								if v72 == int32(0) {
									v87 = v4
								} else {
									v75 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
									v76 = v72
									v79 = v75
									if v79 != int32(58) {
										v87 = v4
									} else {
										v82 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
										v83 = *(*int32)(unsafe.Add(mBase, uint32(l2)+68))
										v87 = base.B2i32(v82 == v83)
									}
								}
								return v87
							} else {
								v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+20)))
								if v38 != int32(1) {
									v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
									if v72 == int32(0) {
										v87 = v4
									} else {
										v75 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
										v76 = v72
										v79 = v75
										if v79 != int32(58) {
											v87 = v4
										} else {
											v82 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
											v83 = *(*int32)(unsafe.Add(mBase, uint32(l2)+68))
											v87 = base.B2i32(v82 == v83)
										}
									}
									return v87
								} else {
									v41 = *(*int32)(unsafe.Add(mBase, uint32(v24)+28))
									v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+12))
									v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
									if v43 == int32(0) {
										v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
										if v72 == int32(0) {
											v87 = v4
										} else {
											v75 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
											v76 = v72
											v79 = v75
											if v79 != int32(58) {
												v87 = v4
											} else {
												v82 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
												v83 = *(*int32)(unsafe.Add(mBase, uint32(l2)+68))
												v87 = base.B2i32(v82 == v83)
											}
										}
										return v87
									} else {
										v46 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
										if v46 != int32(6) {
											v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
											if v72 == int32(0) {
												v87 = v4
											} else {
												v75 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
												v76 = v72
												v79 = v75
												if v79 != int32(58) {
													v87 = v4
												} else {
													v82 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
													v83 = *(*int32)(unsafe.Add(mBase, uint32(l2)+68))
													v87 = base.B2i32(v82 == v83)
												}
											}
											return v87
										} else {
											v49 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v43)+8)))
											if v49 != int32(65535) {
												v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
												if v72 == int32(0) {
													v87 = v4
												} else {
													v75 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
													v76 = v72
													v79 = v75
													if v79 != int32(58) {
														v87 = v4
													} else {
														v82 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
														v83 = *(*int32)(unsafe.Add(mBase, uint32(l2)+68))
														v87 = base.B2i32(v82 == v83)
													}
												}
												return v87
											} else {
												v52 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
												if v52 != int32(27) {
													v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
													if v72 == int32(0) {
														v87 = v4
													} else {
														v75 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
														v76 = v72
														v79 = v75
														if v79 != int32(58) {
															v87 = v4
														} else {
															v82 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
															v83 = *(*int32)(unsafe.Add(mBase, uint32(l2)+68))
															v87 = base.B2i32(v82 == v83)
														}
													}
													return v87
												} else {
													v55 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
													v56 = *(*int32)(unsafe.Add(mBase, uint32(l2)+68))
													if v55 != v56 {
														v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
														if v72 == int32(0) {
															v87 = v4
														} else {
															v75 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
															v76 = v72
															v79 = v75
															if v79 != int32(58) {
																v87 = v4
															} else {
																v82 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
																v83 = *(*int32)(unsafe.Add(mBase, uint32(l2)+68))
																v87 = base.B2i32(v82 == v83)
															}
														}
														return v87
													} else {
														v58 = *(*int32)(unsafe.Add(mBase, uint32(v43)+24))
														if v58 != 0 {
															v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
															if v72 == int32(0) {
																v87 = v4
															} else {
																v75 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
																v76 = v72
																v79 = v75
																if v79 != int32(58) {
																	v87 = v4
																} else {
																	v82 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
																	v83 = *(*int32)(unsafe.Add(mBase, uint32(l2)+68))
																	v87 = base.B2i32(v82 == v83)
																}
															}
															return v87
														} else {
															v59 = *(*int32)(unsafe.Add(mBase, uint32(v43)+28))
															if v59 != 0 {
																v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
																if v72 == int32(0) {
																	v87 = v4
																} else {
																	v75 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
																	v76 = v72
																	v79 = v75
																	if v79 != int32(58) {
																		v87 = v4
																	} else {
																		v82 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
																		v83 = *(*int32)(unsafe.Add(mBase, uint32(l2)+68))
																		v87 = base.B2i32(v82 == v83)
																	}
																}
																return v87
															} else {
																v60 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
																v61 = F_pull_varnos(m, l0, v60)
																mBase = m.M
																v62 = m.ExcPending
																if v62 != 0 {
																	return int32(0)
																} else {
																	v63 = F_bms_is_member(m, v55, v61)
																	mBase = m.M
																	v64 = m.ExcPending
																	if v64 != 0 {
																		return int32(0)
																	} else {
																		if v63 != 0 {
																			v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
																			if v72 == int32(0) {
																				v87 = v4
																			} else {
																				v75 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
																				v76 = v72
																				v79 = v75
																				if v79 != int32(58) {
																					v87 = v4
																				} else {
																					v82 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
																					v83 = *(*int32)(unsafe.Add(mBase, uint32(l2)+68))
																					v87 = base.B2i32(v82 == v83)
																				}
																			}
																			return v87
																		} else {
																			v65 = F_contain_volatile_functions(m, v60)
																			mBase = m.M
																			v66 = m.ExcPending
																			if v66 != 0 {
																				return int32(0)
																			} else {
																				if v65 != 0 {
																					v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
																					if v72 == int32(0) {
																						v87 = v4
																					} else {
																						v75 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
																						v76 = v72
																						v79 = v75
																						if v79 != int32(58) {
																							v87 = v4
																						} else {
																							v82 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
																							v83 = *(*int32)(unsafe.Add(mBase, uint32(l2)+68))
																							v87 = base.B2i32(v82 == v83)
																						}
																					}
																					return v87
																				} else {
																					return int32(1)
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
							v76 = v24
							v79 = v32
							if v79 != int32(58) {
								v87 = v4
							} else {
								v82 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
								v83 = *(*int32)(unsafe.Add(mBase, uint32(l2)+68))
								v87 = base.B2i32(v82 == v83)
							}
							return v87
						}
					}
				}
			}
		}
	}
}
func F_RoleidCallback(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, _consts[1235])) = int32(0)
	return
}
func F_r_CONSONANT(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v88 int32
	_ = v88
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v18 <= v19 {
		v123 = int32(-1)
		v130 = v123
	} else {
		v36 = int32(1)
		v37 = v18 - v36
		v39 = int32(*(*int8)(unsafe.Add(mBase, uint32(v15+v37))))
		v41 = v39 & int32(255)
		if v37 == v19 {
			v96 = v41
			v97 = v36
		} else {
			if int32(0) <= v39 {
				v96 = v41
				v97 = v36
			} else {
				v47 = v41 & int32(63)
				v49 = v18 - int32(2)
				v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15+v49))))
				v53 = v51 << (uint(int32(6)) % 32)
				if base.B2i32(v49 != v19)&base.B2i32(base.Ui32(v51) < base.Ui32(int32(192))) == int32(0) {
					v96 = v53&int32(1984) | v47
					v97 = int32(2)
				} else {
					v66 = v53&int32(4032) | v47
					v68 = v18 - int32(3)
					v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15+v68))))
					if base.B2i32(v68 != v19)&base.B2i32(base.Ui32(v70) < base.Ui32(int32(224))) == int32(0) {
						v96 = v70<<(uint(int32(12))%32)&int32(61440) | v66
						v97 = int32(3)
					} else {
						v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18+(v15-int32(4))))))
						v96 = v70<<(uint(int32(12))%32)&int32(258048) | v88&int32(7)<<(uint(int32(18))%32) | v66
						v97 = int32(4)
					}
				}
			}
		}
		if int32(2399) < v96 {
			v130 = v97
		} else {
			v101 = v96 - int32(2325)
			if v101 < int32(0) {
				v130 = v97
			} else {
				v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v101)>>(uint(int32(3))%32)))+uint32(_consts[1322]))))
				if int32(base.Ui32(v107)>>(uint(v101&int32(7))%32))&int32(1) == int32(0) {
					v130 = v97
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v18 - v97
					v123 = int32(0)
					v130 = v123
				}
			}
		}
	}
	return base.B2i32(v130 == int32(0))
}
func F_r_KER_2(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v75 int32
	_ = v75
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v112 int32
	_ = v112
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v15 <= v14 {
		v119 = int32(-1)
	} else {
		v31 = int32(1)
		v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14+v16))))
		if base.Ui32(v33) < base.Ui32(int32(192)) {
			v90 = v33
			v91 = v31
		} else {
			v37 = v14 + int32(1)
			if v37 == v15 {
				v90 = v33
				v91 = v31
			} else {
				v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37+v16))))
				v42 = v40 & int32(63)
				if base.Ui32(int32(224)) <= base.Ui32(v33) {
					v46 = v14 + int32(2)
					if v46 != v15 {
						v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46+v16))))
						v58 = v56 & int32(63)
						if base.Ui32(int32(240)) <= base.Ui32(v33) {
							v62 = v14 + int32(3)
							if v62 != v15 {
								v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16+v62))))
								v90 = v75&int32(63) | (v33<<(uint(int32(18))%32)&int32(1835008) | v42<<(uint(int32(12))%32) | v58<<(uint(int32(6))%32))
								v91 = int32(4)
							} else {
								v90 = v33<<(uint(int32(12))%32)&int32(61440) | v42<<(uint(int32(6))%32) | v58
								v91 = int32(3)
							}
						} else {
							v90 = v33<<(uint(int32(12))%32)&int32(61440) | v42<<(uint(int32(6))%32) | v58
							v91 = int32(3)
						}
					} else {
						v90 = v33<<(uint(int32(6))%32)&int32(1984) | v42
						v91 = int32(2)
					}
				} else {
					v90 = v33<<(uint(int32(6))%32)&int32(1984) | v42
					v91 = int32(2)
				}
			}
		}
		if int32(117) < v90 {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v91 + v14
			v112 = int32(0)
		} else {
			v95 = v90 - int32(97)
			if v95 < int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v91 + v14
				v112 = int32(0)
			} else {
				v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v95)>>(uint(int32(3))%32)))+uint32(_consts[1324]))))
				if int32(base.Ui32(v101)>>(uint(v95&int32(7))%32))&int32(1) != 0 {
					v112 = v91
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v91 + v14
					v112 = int32(0)
				}
			}
		}
		v119 = v112
	}
	if v119 != 0 {
		v137 = int32(0)
	} else {
		v120 = int32(2)
		v122 = int32(0)
		v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v124-v125 < v120 {
			v134 = v122
		} else {
			v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v130 = F_memcmp(m, v128+v125, int32(2176213), v120)
			mBase = m.M
			if v130 != 0 {
				v134 = v122
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v120 + v125
				v134 = int32(1)
			}
		}
		v137 = base.B2i32(v134 != int32(0))
	}
	return v137
}
func F_r_Suffix_Noun_Step2b(m *base.Module, l0 int32) int32 {
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
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	v2 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v4
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v4-int32(3) <= v6 {
		v117 = v2
		return v117
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10+v4-int32(1)))))
		if v14 != int32(170) {
			v117 = v2
			return v117
		} else {
			v19 = F_find_among_b(m, l0, int32(4201104), int32(1))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				if v19 == int32(0) {
					v117 = v2
					return v117
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v25
					v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v28 = int32(0)
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v27-int32(4))))
					if v35 == v28 {
						v107 = int32(0)
					} else {
						v40 = v35 & int32(3)
						if base.Ui32(v35) < base.Ui32(int32(4)) {
							v74 = v27
							v75 = int32(0)
						} else {
							v47 = v27
							v48 = int32(0)
							v51 = v28
							for {
								v53 = int32(*(*int8)(unsafe.Add(mBase, uint32(v47))))
								v54 = int32(-65)
								v57 = int32(*(*int8)(unsafe.Add(mBase, uint32(v47)+1)))
								v61 = int32(*(*int8)(unsafe.Add(mBase, uint32(v47)+2)))
								v65 = int32(*(*int8)(unsafe.Add(mBase, uint32(v47)+3)))
								v68 = v48 + base.B2i32(v54 < v53) + base.B2i32(v54 < v57) + base.B2i32(v54 < v61) + base.B2i32(v54 < v65)
								v69 = int32(4)
								v70 = v47 + v69
								v72 = v51 + v69
								if v72 != v35&int32(-4) {
									v47 = v70
									v48 = v68
									v51 = v72
									continue
								} else {
									break
								}
								break
							}
							v74 = v70
							v75 = v68
						}
						if v40 != 0 {
							v80 = v74
							v81 = v75
							v83 = v28
							for {
								v86 = int32(*(*int8)(unsafe.Add(mBase, uint32(v80))))
								v89 = v81 + base.B2i32(int32(-65) < v86)
								v90 = int32(1)
								v93 = v83 + v90
								if v93 != v40 {
									v80 = v80 + v90
									v81 = v89
									v83 = v93
									continue
								} else {
									break
								}
								break
							}
							v96 = v89
						} else {
							v96 = v75
						}
						v107 = v96
					}
					if v107 < int32(5) {
						v117 = v2
						return v117
					} else {
						v111 = F_slice_del(m, l0)
						mBase = m.M
						v112 = m.ExcPending
						if v112 != 0 {
							return int32(0)
						} else {
							if int32(0) <= v111 {
								v115 = int32(1)
							} else {
								v115 = v111
							}
							v117 = v115
							return v117
						}
					}
				}
			}
		}
	}
}
func F_r_Suffix_Noun_Step2c1(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	v2 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v4
	v7 = v4 - int32(1)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v7 <= v8 {
		v115 = v2
		return v115
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10+v7))))
		if v12 != int32(170) {
			v115 = v2
			return v115
		} else {
			v17 = F_find_among_b(m, l0, int32(4201136), int32(1))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				if v17 == int32(0) {
					v115 = v2
					return v115
				} else {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v23
					v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v26 = int32(0)
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v25-int32(4))))
					if v33 == v26 {
						v105 = int32(0)
					} else {
						v38 = v33 & int32(3)
						if base.Ui32(v33) < base.Ui32(int32(4)) {
							v72 = v25
							v73 = int32(0)
						} else {
							v45 = v25
							v46 = int32(0)
							v49 = v26
							for {
								v51 = int32(*(*int8)(unsafe.Add(mBase, uint32(v45))))
								v52 = int32(-65)
								v55 = int32(*(*int8)(unsafe.Add(mBase, uint32(v45)+1)))
								v59 = int32(*(*int8)(unsafe.Add(mBase, uint32(v45)+2)))
								v63 = int32(*(*int8)(unsafe.Add(mBase, uint32(v45)+3)))
								v66 = v46 + base.B2i32(v52 < v51) + base.B2i32(v52 < v55) + base.B2i32(v52 < v59) + base.B2i32(v52 < v63)
								v67 = int32(4)
								v68 = v45 + v67
								v70 = v49 + v67
								if v70 != v33&int32(-4) {
									v45 = v68
									v46 = v66
									v49 = v70
									continue
								} else {
									break
								}
								break
							}
							v72 = v68
							v73 = v66
						}
						if v38 != 0 {
							v78 = v72
							v79 = v73
							v81 = v26
							for {
								v84 = int32(*(*int8)(unsafe.Add(mBase, uint32(v78))))
								v87 = v79 + base.B2i32(int32(-65) < v84)
								v88 = int32(1)
								v91 = v81 + v88
								if v91 != v38 {
									v78 = v78 + v88
									v79 = v87
									v81 = v91
									continue
								} else {
									break
								}
								break
							}
							v94 = v87
						} else {
							v94 = v73
						}
						v105 = v94
					}
					if v105 < int32(4) {
						v115 = v2
						return v115
					} else {
						v109 = F_slice_del(m, l0)
						mBase = m.M
						v110 = m.ExcPending
						if v110 != 0 {
							return int32(0)
						} else {
							if int32(0) <= v109 {
								v113 = int32(1)
							} else {
								v113 = v109
							}
							v115 = v113
							return v115
						}
					}
				}
			}
		}
	}
}
func F_r_Suffix_Verb_Step2a(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v287 int32
	_ = v287
	var v292 int32
	_ = v292
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v368 int32
	_ = v368
	v2 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v4
	v8 = F_find_among_b(m, l0, int32(4200480), int32(11))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		if v8 == int32(0) {
			v368 = v2
			return v368
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v14
			switch v8 - int32(1) {
			case 0:
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v19 = int32(0)
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v18-int32(4))))
				if v26 == v19 {
					v98 = int32(0)
				} else {
					v31 = v26 & int32(3)
					if base.Ui32(v26) < base.Ui32(int32(4)) {
						v65 = v18
						v66 = int32(0)
					} else {
						v38 = v18
						v39 = int32(0)
						v42 = v19
						for {
							v44 = int32(*(*int8)(unsafe.Add(mBase, uint32(v38))))
							v45 = int32(-65)
							v48 = int32(*(*int8)(unsafe.Add(mBase, uint32(v38)+1)))
							v52 = int32(*(*int8)(unsafe.Add(mBase, uint32(v38)+2)))
							v56 = int32(*(*int8)(unsafe.Add(mBase, uint32(v38)+3)))
							v59 = v39 + base.B2i32(v45 < v44) + base.B2i32(v45 < v48) + base.B2i32(v45 < v52) + base.B2i32(v45 < v56)
							v60 = int32(4)
							v61 = v38 + v60
							v63 = v42 + v60
							if v63 != v26&int32(-4) {
								v38 = v61
								v39 = v59
								v42 = v63
								continue
							} else {
								break
							}
							break
						}
						v65 = v61
						v66 = v59
					}
					if v31 != 0 {
						v71 = v65
						v72 = v66
						v74 = v19
						for {
							v77 = int32(*(*int8)(unsafe.Add(mBase, uint32(v71))))
							v80 = v72 + base.B2i32(int32(-65) < v77)
							v81 = int32(1)
							v84 = v74 + v81
							if v84 != v31 {
								v71 = v71 + v81
								v72 = v80
								v74 = v84
								continue
							} else {
								break
							}
							break
						}
						v87 = v80
					} else {
						v87 = v66
					}
					v98 = v87
				}
				if v98 < int32(4) {
					v368 = v2
					return v368
				} else {
					v101 = F_slice_del(m, l0)
					mBase = m.M
					v102 = m.ExcPending
					if v102 != 0 {
						return int32(0)
					} else {
						if int32(0) <= v101 {
							v368 = int32(1)
						} else {
							v368 = v101
						}
						return v368
					}
				}
			case 1:
				v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v106 = int32(0)
				v113 = *(*int32)(unsafe.Add(mBase, uint32(v105-int32(4))))
				if v113 == v106 {
					v185 = int32(0)
				} else {
					v118 = v113 & int32(3)
					if base.Ui32(v113) < base.Ui32(int32(4)) {
						v152 = v105
						v153 = int32(0)
					} else {
						v125 = v105
						v126 = int32(0)
						v129 = v106
						for {
							v131 = int32(*(*int8)(unsafe.Add(mBase, uint32(v125))))
							v132 = int32(-65)
							v135 = int32(*(*int8)(unsafe.Add(mBase, uint32(v125)+1)))
							v139 = int32(*(*int8)(unsafe.Add(mBase, uint32(v125)+2)))
							v143 = int32(*(*int8)(unsafe.Add(mBase, uint32(v125)+3)))
							v146 = v126 + base.B2i32(v132 < v131) + base.B2i32(v132 < v135) + base.B2i32(v132 < v139) + base.B2i32(v132 < v143)
							v147 = int32(4)
							v148 = v125 + v147
							v150 = v129 + v147
							if v150 != v113&int32(-4) {
								v125 = v148
								v126 = v146
								v129 = v150
								continue
							} else {
								break
							}
							break
						}
						v152 = v148
						v153 = v146
					}
					if v118 != 0 {
						v158 = v152
						v159 = v153
						v161 = v106
						for {
							v164 = int32(*(*int8)(unsafe.Add(mBase, uint32(v158))))
							v167 = v159 + base.B2i32(int32(-65) < v164)
							v168 = int32(1)
							v171 = v161 + v168
							if v171 != v118 {
								v158 = v158 + v168
								v159 = v167
								v161 = v171
								continue
							} else {
								break
							}
							break
						}
						v174 = v167
					} else {
						v174 = v153
					}
					v185 = v174
				}
				if v185 < int32(5) {
					v368 = v2
					return v368
				} else {
					v188 = F_slice_del(m, l0)
					mBase = m.M
					v189 = m.ExcPending
					if v189 != 0 {
						return int32(0)
					} else {
						if int32(0) <= v188 {
							v368 = int32(1)
						} else {
							v368 = v188
						}
						return v368
					}
				}
			case 2:
				v192 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v193 = int32(0)
				v200 = *(*int32)(unsafe.Add(mBase, uint32(v192-int32(4))))
				if v200 == v193 {
					v272 = int32(0)
				} else {
					v205 = v200 & int32(3)
					if base.Ui32(v200) < base.Ui32(int32(4)) {
						v239 = v192
						v240 = int32(0)
					} else {
						v212 = v192
						v213 = int32(0)
						v216 = v193
						for {
							v218 = int32(*(*int8)(unsafe.Add(mBase, uint32(v212))))
							v219 = int32(-65)
							v222 = int32(*(*int8)(unsafe.Add(mBase, uint32(v212)+1)))
							v226 = int32(*(*int8)(unsafe.Add(mBase, uint32(v212)+2)))
							v230 = int32(*(*int8)(unsafe.Add(mBase, uint32(v212)+3)))
							v233 = v213 + base.B2i32(v219 < v218) + base.B2i32(v219 < v222) + base.B2i32(v219 < v226) + base.B2i32(v219 < v230)
							v234 = int32(4)
							v235 = v212 + v234
							v237 = v216 + v234
							if v237 != v200&int32(-4) {
								v212 = v235
								v213 = v233
								v216 = v237
								continue
							} else {
								break
							}
							break
						}
						v239 = v235
						v240 = v233
					}
					if v205 != 0 {
						v245 = v239
						v246 = v240
						v248 = v193
						for {
							v251 = int32(*(*int8)(unsafe.Add(mBase, uint32(v245))))
							v254 = v246 + base.B2i32(int32(-65) < v251)
							v255 = int32(1)
							v258 = v248 + v255
							if v258 != v205 {
								v245 = v245 + v255
								v246 = v254
								v248 = v258
								continue
							} else {
								break
							}
							break
						}
						v261 = v254
					} else {
						v261 = v240
					}
					v272 = v261
				}
				if v272 < int32(6) {
					v368 = v2
					return v368
				} else {
					v275 = F_slice_del(m, l0)
					mBase = m.M
					v276 = m.ExcPending
					if v276 != 0 {
						return int32(0)
					} else {
						if int32(0) <= v275 {
							v368 = int32(1)
						} else {
							v368 = v275
						}
						return v368
					}
				}
			case 3:
				v279 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v280 = int32(0)
				v287 = *(*int32)(unsafe.Add(mBase, uint32(v279-int32(4))))
				if v287 == v280 {
					v359 = int32(0)
				} else {
					v292 = v287 & int32(3)
					if base.Ui32(v287) < base.Ui32(int32(4)) {
						v326 = v279
						v327 = int32(0)
					} else {
						v299 = v279
						v300 = int32(0)
						v303 = v280
						for {
							v305 = int32(*(*int8)(unsafe.Add(mBase, uint32(v299))))
							v306 = int32(-65)
							v309 = int32(*(*int8)(unsafe.Add(mBase, uint32(v299)+1)))
							v313 = int32(*(*int8)(unsafe.Add(mBase, uint32(v299)+2)))
							v317 = int32(*(*int8)(unsafe.Add(mBase, uint32(v299)+3)))
							v320 = v300 + base.B2i32(v306 < v305) + base.B2i32(v306 < v309) + base.B2i32(v306 < v313) + base.B2i32(v306 < v317)
							v321 = int32(4)
							v322 = v299 + v321
							v324 = v303 + v321
							if v324 != v287&int32(-4) {
								v299 = v322
								v300 = v320
								v303 = v324
								continue
							} else {
								break
							}
							break
						}
						v326 = v322
						v327 = v320
					}
					if v292 != 0 {
						v332 = v326
						v333 = v327
						v335 = v280
						for {
							v338 = int32(*(*int8)(unsafe.Add(mBase, uint32(v332))))
							v341 = v333 + base.B2i32(int32(-65) < v338)
							v342 = int32(1)
							v345 = v335 + v342
							if v345 != v292 {
								v332 = v332 + v342
								v333 = v341
								v335 = v345
								continue
							} else {
								break
							}
							break
						}
						v348 = v341
					} else {
						v348 = v327
					}
					v359 = v348
				}
				if v359 < int32(6) {
					v368 = v2
					return v368
				} else {
					v362 = F_slice_del(m, l0)
					mBase = m.M
					v363 = m.ExcPending
					if v363 != 0 {
						return int32(0)
					} else {
						if v362 < int32(0) {
							v368 = v362
						} else {
							v368 = int32(1)
						}
						return v368
					}
				}
			default:
				v368 = int32(1)
				return v368
			}
		}
	}
}
func F_r_check_vowel_harmony(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v36 int32
	_ = v36
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v96 int32
	_ = v96
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v124 int32
	_ = v124
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v182 int32
	_ = v182
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v242 int32
	_ = v242
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v261 int32
	_ = v261
	var v270 int32
	_ = v270
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v332 int32
	_ = v332
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v392 int32
	_ = v392
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v411 int32
	_ = v411
	var v420 int32
	_ = v420
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v450 int32
	_ = v450
	var v453 int32
	_ = v453
	var v457 int32
	_ = v457
	var v471 int32
	_ = v471
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v484 int32
	_ = v484
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v526 int32
	_ = v526
	var v544 int32
	_ = v544
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v557 int32
	_ = v557
	var v563 int32
	_ = v563
	var v572 int32
	_ = v572
	var v587 int32
	_ = v587
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v593 int32
	_ = v593
	var v595 int32
	_ = v595
	var v599 int32
	_ = v599
	var v603 int32
	_ = v603
	var v618 int32
	_ = v618
	var v622 int32
	_ = v622
	var v631 int32
	_ = v631
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v642 int32
	_ = v642
	var v644 int32
	_ = v644
	var v650 int32
	_ = v650
	var v652 int32
	_ = v652
	var v654 int32
	_ = v654
	var v656 int32
	_ = v656
	var v669 int32
	_ = v669
	var v671 int32
	_ = v671
	var v673 int32
	_ = v673
	var v691 int32
	_ = v691
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v704 int32
	_ = v704
	var v710 int32
	_ = v710
	var v719 int32
	_ = v719
	var v734 int32
	_ = v734
	var v737 int32
	_ = v737
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v744 int32
	_ = v744
	var v748 int32
	_ = v748
	var v752 int32
	_ = v752
	var v767 int32
	_ = v767
	var v771 int32
	_ = v771
	var v780 int32
	_ = v780
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v791 int32
	_ = v791
	var v793 int32
	_ = v793
	var v799 int32
	_ = v799
	var v801 int32
	_ = v801
	var v803 int32
	_ = v803
	var v805 int32
	_ = v805
	var v818 int32
	_ = v818
	var v820 int32
	_ = v820
	var v822 int32
	_ = v822
	var v840 int32
	_ = v840
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v853 int32
	_ = v853
	var v859 int32
	_ = v859
	var v868 int32
	_ = v868
	var v883 int32
	_ = v883
	var v886 int32
	_ = v886
	var v888 int32
	_ = v888
	var v890 int32
	_ = v890
	var v892 int32
	_ = v892
	var v895 int32
	_ = v895
	var v898 int32
	_ = v898
	var v901 int32
	_ = v901
	var v905 int32
	_ = v905
	var v919 int32
	_ = v919
	var v922 int32
	_ = v922
	var v923 int32
	_ = v923
	var v932 int32
	_ = v932
	var v940 int32
	_ = v940
	var v941 int32
	_ = v941
	var v943 int32
	_ = v943
	var v945 int32
	_ = v945
	var v951 int32
	_ = v951
	var v953 int32
	_ = v953
	var v955 int32
	_ = v955
	var v957 int32
	_ = v957
	var v970 int32
	_ = v970
	var v972 int32
	_ = v972
	var v974 int32
	_ = v974
	var v992 int32
	_ = v992
	var v1000 int32
	_ = v1000
	var v1001 int32
	_ = v1001
	var v1005 int32
	_ = v1005
	var v1011 int32
	_ = v1011
	var v1020 int32
	_ = v1020
	var v1035 int32
	_ = v1035
	var v1038 int32
	_ = v1038
	var v1039 int32
	_ = v1039
	var v1041 int32
	_ = v1041
	var v1043 int32
	_ = v1043
	var v1047 int32
	_ = v1047
	var v1051 int32
	_ = v1051
	var v1066 int32
	_ = v1066
	var v1070 int32
	_ = v1070
	var v1079 int32
	_ = v1079
	var v1087 int32
	_ = v1087
	var v1088 int32
	_ = v1088
	var v1090 int32
	_ = v1090
	var v1092 int32
	_ = v1092
	var v1098 int32
	_ = v1098
	var v1100 int32
	_ = v1100
	var v1102 int32
	_ = v1102
	var v1104 int32
	_ = v1104
	var v1117 int32
	_ = v1117
	var v1119 int32
	_ = v1119
	var v1121 int32
	_ = v1121
	var v1139 int32
	_ = v1139
	var v1147 int32
	_ = v1147
	var v1148 int32
	_ = v1148
	var v1152 int32
	_ = v1152
	var v1158 int32
	_ = v1158
	var v1167 int32
	_ = v1167
	var v1182 int32
	_ = v1182
	var v1185 int32
	_ = v1185
	var v1187 int32
	_ = v1187
	var v1189 int32
	_ = v1189
	var v1191 int32
	_ = v1191
	var v1194 int32
	_ = v1194
	var v1197 int32
	_ = v1197
	var v1200 int32
	_ = v1200
	var v1204 int32
	_ = v1204
	var v1220 int32
	_ = v1220
	var v1223 int32
	_ = v1223
	var v1224 int32
	_ = v1224
	var v1233 int32
	_ = v1233
	var v1241 int32
	_ = v1241
	var v1242 int32
	_ = v1242
	var v1244 int32
	_ = v1244
	var v1246 int32
	_ = v1246
	var v1252 int32
	_ = v1252
	var v1254 int32
	_ = v1254
	var v1256 int32
	_ = v1256
	var v1258 int32
	_ = v1258
	var v1271 int32
	_ = v1271
	var v1273 int32
	_ = v1273
	var v1275 int32
	_ = v1275
	var v1293 int32
	_ = v1293
	var v1301 int32
	_ = v1301
	var v1302 int32
	_ = v1302
	var v1306 int32
	_ = v1306
	var v1312 int32
	_ = v1312
	var v1321 int32
	_ = v1321
	var v1336 int32
	_ = v1336
	var v1342 int32
	_ = v1342
	var v1350 int32
	_ = v1350
	v2 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v36 = v9
	goto L4
L1:
	;
	return v1350
L2:
	;
	if v139 < int32(0) {
		v1350 = v2
		goto L1
	} else {
		goto L21
	}
L3:
	;
	v139 = int32(-1)
	goto L2
L4:
	;
	if v36 <= v27 {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v44 = int32(1)
	v45 = v36 - v44
	v47 = int32(*(*int8)(unsafe.Add(mBase, uint32(v23+v45))))
	v49 = v47 & int32(255)
	if v45 == v27 {
		v104 = v49
		v105 = v44
		goto L7
	} else {
		goto L8
	}
L7:
	;
	if int32(305) < v104 {
		goto L16
	} else {
		goto L17
	}
L8:
	;
	if int32(0) <= v47 {
		v104 = v49
		v105 = v44
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v55 = v49 & int32(63)
	v57 = v36 - int32(2)
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23+v57))))
	v61 = v59 << (uint(int32(6)) % 32)
	if base.B2i32(v57 != v27)&base.B2i32(base.Ui32(v59) < base.Ui32(int32(192))) == int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v104 = v61&int32(1984) | v55
	v105 = int32(2)
	goto L7
L11:
	;
	goto L12
L12:
	;
	v74 = v61&int32(4032) | v55
	v76 = v36 - int32(3)
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23+v76))))
	if base.B2i32(v76 != v27)&base.B2i32(base.Ui32(v78) < base.Ui32(int32(224))) == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v104 = v78<<(uint(int32(12))%32)&int32(61440) | v74
	v105 = int32(3)
	goto L7
L14:
	;
	goto L15
L15:
	;
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36+(v23-int32(4))))))
	v104 = v78<<(uint(int32(12))%32)&int32(258048) | v96&int32(7)<<(uint(int32(18))%32) | v74
	v105 = int32(4)
	goto L7
L16:
	;
	v124 = v36 - v105
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v124
	v36 = v124
	goto L4
L17:
	;
	v109 = v104 - int32(97)
	if v109 < int32(0) {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v109)>>(uint(int32(3))%32)))+uint32(_consts[1327]))))
	if int32(base.Ui32(v115)>>(uint(v109&int32(7))%32))&int32(1) == int32(0) {
		goto L16
	} else {
		goto L19
	}
L19:
	;
	v139 = v105
	goto L2
L21:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v143 <= v144 {
		v290 = v144
		v291 = v142
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v1342 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1342 + (v9 - v8)
	v1350 = int32(1)
	goto L1
L23:
	;
	v292 = v142 - v143
	v293 = v291 - v292
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v293
	if v293 <= v290 {
		v440 = v293
		goto L46
	} else {
		goto L47
	}
L24:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146+v143-int32(1)))))
	if v150 != int32(97) {
		v290 = v144
		v291 = v142
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v154 = v143 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v154
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v182 = v154
	goto L28
L26:
	;
	if int32(0) <= v285 {
		goto L22
	} else {
		goto L45
	}
L27:
	;
	v285 = int32(-1)
	goto L26
L28:
	;
	if v182 <= v173 {
		goto L27
	} else {
		goto L30
	}
L30:
	;
	v190 = int32(1)
	v191 = v182 - v190
	v193 = int32(*(*int8)(unsafe.Add(mBase, uint32(v169+v191))))
	v195 = v193 & int32(255)
	if v191 == v173 {
		v250 = v195
		v251 = v190
		goto L31
	} else {
		goto L32
	}
L31:
	;
	if int32(305) < v250 {
		goto L40
	} else {
		goto L41
	}
L32:
	;
	if int32(0) <= v193 {
		v250 = v195
		v251 = v190
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v201 = v195 & int32(63)
	v203 = v182 - int32(2)
	v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169+v203))))
	v207 = v205 << (uint(int32(6)) % 32)
	if base.B2i32(v203 != v173)&base.B2i32(base.Ui32(v205) < base.Ui32(int32(192))) == int32(0) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v250 = v207&int32(1984) | v201
	v251 = int32(2)
	goto L31
L35:
	;
	goto L36
L36:
	;
	v220 = v207&int32(4032) | v201
	v222 = v182 - int32(3)
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169+v222))))
	if base.B2i32(v222 != v173)&base.B2i32(base.Ui32(v224) < base.Ui32(int32(224))) == int32(0) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v250 = v224<<(uint(int32(12))%32)&int32(61440) | v220
	v251 = int32(3)
	goto L31
L38:
	;
	goto L39
L39:
	;
	v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182+(v169-int32(4))))))
	v250 = v224<<(uint(int32(12))%32)&int32(258048) | v242&int32(7)<<(uint(int32(18))%32) | v220
	v251 = int32(4)
	goto L31
L40:
	;
	v270 = v182 - v251
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v270
	v182 = v270
	goto L28
L41:
	;
	v255 = v250 - int32(97)
	if v255 < int32(0) {
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v255)>>(uint(int32(3))%32)))+uint32(_consts[1328]))))
	if int32(base.Ui32(v261)>>(uint(v255&int32(7))%32))&int32(1) == int32(0) {
		goto L40
	} else {
		goto L43
	}
L43:
	;
	v285 = v251
	goto L26
L45:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v289 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v290 = v288
	v291 = v289
	goto L23
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v440
	v442 = int32(2)
	v444 = int32(0)
	v447 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v440-v447 < v442 {
		v457 = v444
		goto L70
	} else {
		goto L71
	}
L47:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v296+v293-int32(1)))))
	if v300 != int32(101) {
		v440 = v293
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v304 = v293 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v304
	v319 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v323 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v332 = v304
	goto L51
L49:
	;
	if int32(0) <= v435 {
		goto L22
	} else {
		goto L68
	}
L50:
	;
	v435 = int32(-1)
	goto L49
L51:
	;
	if v332 <= v323 {
		goto L50
	} else {
		goto L53
	}
L53:
	;
	v340 = int32(1)
	v341 = v332 - v340
	v343 = int32(*(*int8)(unsafe.Add(mBase, uint32(v319+v341))))
	v345 = v343 & int32(255)
	if v341 == v323 {
		v400 = v345
		v401 = v340
		goto L54
	} else {
		goto L55
	}
L54:
	;
	if int32(252) < v400 {
		goto L63
	} else {
		goto L64
	}
L55:
	;
	if int32(0) <= v343 {
		v400 = v345
		v401 = v340
		goto L54
	} else {
		goto L56
	}
L56:
	;
	v351 = v345 & int32(63)
	v353 = v332 - int32(2)
	v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v319+v353))))
	v357 = v355 << (uint(int32(6)) % 32)
	if base.B2i32(v353 != v323)&base.B2i32(base.Ui32(v355) < base.Ui32(int32(192))) == int32(0) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v400 = v357&int32(1984) | v351
	v401 = int32(2)
	goto L54
L58:
	;
	goto L59
L59:
	;
	v370 = v357&int32(4032) | v351
	v372 = v332 - int32(3)
	v374 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v319+v372))))
	if base.B2i32(v372 != v323)&base.B2i32(base.Ui32(v374) < base.Ui32(int32(224))) == int32(0) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v400 = v374<<(uint(int32(12))%32)&int32(61440) | v370
	v401 = int32(3)
	goto L54
L61:
	;
	goto L62
L62:
	;
	v392 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v332+(v319-int32(4))))))
	v400 = v374<<(uint(int32(12))%32)&int32(258048) | v392&int32(7)<<(uint(int32(18))%32) | v370
	v401 = int32(4)
	goto L54
L63:
	;
	v420 = v332 - v401
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v420
	v332 = v420
	goto L51
L64:
	;
	v405 = v400 - int32(101)
	if v405 < int32(0) {
		goto L63
	} else {
		goto L65
	}
L65:
	;
	v411 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v405)>>(uint(int32(3))%32)))+uint32(_consts[1329]))))
	if int32(base.Ui32(v411)>>(uint(v405&int32(7))%32))&int32(1) == int32(0) {
		goto L63
	} else {
		goto L66
	}
L66:
	;
	v435 = v401
	goto L49
L68:
	;
	v438 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v440 = v438 - v292
	goto L46
L69:
	;
	if v457 != 0 {
		goto L73
	} else {
		goto L74
	}
L70:
	;
	goto L69
L71:
	;
	v450 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v453 = F_memcmp(m, v450+v440-v442, int32(2198515), v442)
	mBase = m.M
	if v453 != 0 {
		v457 = v444
		goto L70
	} else {
		goto L72
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v440 - v442
	v457 = int32(1)
	goto L70
L73:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v474 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v475 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v484 = v474
	goto L78
L74:
	;
	goto L75
L75:
	;
	v590 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v591 = v590 - v292
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v591
	v593 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v591 <= v593 {
		v740 = v591
		v741 = v593
		goto L96
	} else {
		goto L97
	}
L76:
	;
	if int32(0) <= v587 {
		goto L22
	} else {
		goto L95
	}
L77:
	;
	v587 = int32(-1)
	goto L76
L78:
	;
	if v484 <= v475 {
		goto L77
	} else {
		goto L80
	}
L80:
	;
	v492 = int32(1)
	v493 = v484 - v492
	v495 = int32(*(*int8)(unsafe.Add(mBase, uint32(v471+v493))))
	v497 = v495 & int32(255)
	if v493 == v475 {
		v552 = v497
		v553 = v492
		goto L81
	} else {
		goto L82
	}
L81:
	;
	if int32(305) < v552 {
		goto L90
	} else {
		goto L91
	}
L82:
	;
	if int32(0) <= v495 {
		v552 = v497
		v553 = v492
		goto L81
	} else {
		goto L83
	}
L83:
	;
	v503 = v497 & int32(63)
	v505 = v484 - int32(2)
	v507 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v471+v505))))
	v509 = v507 << (uint(int32(6)) % 32)
	if base.B2i32(v505 != v475)&base.B2i32(base.Ui32(v507) < base.Ui32(int32(192))) == int32(0) {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v552 = v509&int32(1984) | v503
	v553 = int32(2)
	goto L81
L85:
	;
	goto L86
L86:
	;
	v522 = v509&int32(4032) | v503
	v524 = v484 - int32(3)
	v526 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v471+v524))))
	if base.B2i32(v524 != v475)&base.B2i32(base.Ui32(v526) < base.Ui32(int32(224))) == int32(0) {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v552 = v526<<(uint(int32(12))%32)&int32(61440) | v522
	v553 = int32(3)
	goto L81
L88:
	;
	goto L89
L89:
	;
	v544 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v484+(v471-int32(4))))))
	v552 = v526<<(uint(int32(12))%32)&int32(258048) | v544&int32(7)<<(uint(int32(18))%32) | v522
	v553 = int32(4)
	goto L81
L90:
	;
	v572 = v484 - v553
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v572
	v484 = v572
	goto L78
L91:
	;
	v557 = v552 - int32(97)
	if v557 < int32(0) {
		goto L90
	} else {
		goto L92
	}
L92:
	;
	v563 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v557)>>(uint(int32(3))%32)))+uint32(_consts[1330]))))
	if int32(base.Ui32(v563)>>(uint(v557&int32(7))%32))&int32(1) == int32(0) {
		goto L90
	} else {
		goto L93
	}
L93:
	;
	v587 = v553
	goto L76
L95:
	;
	goto L75
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v740
	if v740 <= v741 {
		v888 = v740
		goto L119
	} else {
		goto L120
	}
L97:
	;
	v595 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v599 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v595+v591-int32(1)))))
	if v599 != int32(105) {
		v740 = v591
		v741 = v593
		goto L96
	} else {
		goto L98
	}
L98:
	;
	v603 = v591 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v603
	v618 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v622 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v631 = v603
	goto L101
L99:
	;
	if int32(0) <= v734 {
		goto L22
	} else {
		goto L118
	}
L100:
	;
	v734 = int32(-1)
	goto L99
L101:
	;
	if v631 <= v622 {
		goto L100
	} else {
		goto L103
	}
L103:
	;
	v639 = int32(1)
	v640 = v631 - v639
	v642 = int32(*(*int8)(unsafe.Add(mBase, uint32(v618+v640))))
	v644 = v642 & int32(255)
	if v640 == v622 {
		v699 = v644
		v700 = v639
		goto L104
	} else {
		goto L105
	}
L104:
	;
	if int32(105) < v699 {
		goto L113
	} else {
		goto L114
	}
L105:
	;
	if int32(0) <= v642 {
		v699 = v644
		v700 = v639
		goto L104
	} else {
		goto L106
	}
L106:
	;
	v650 = v644 & int32(63)
	v652 = v631 - int32(2)
	v654 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v618+v652))))
	v656 = v654 << (uint(int32(6)) % 32)
	if base.B2i32(v652 != v622)&base.B2i32(base.Ui32(v654) < base.Ui32(int32(192))) == int32(0) {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v699 = v656&int32(1984) | v650
	v700 = int32(2)
	goto L104
L108:
	;
	goto L109
L109:
	;
	v669 = v656&int32(4032) | v650
	v671 = v631 - int32(3)
	v673 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v618+v671))))
	if base.B2i32(v671 != v622)&base.B2i32(base.Ui32(v673) < base.Ui32(int32(224))) == int32(0) {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v699 = v673<<(uint(int32(12))%32)&int32(61440) | v669
	v700 = int32(3)
	goto L104
L111:
	;
	goto L112
L112:
	;
	v691 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v631+(v618-int32(4))))))
	v699 = v673<<(uint(int32(12))%32)&int32(258048) | v691&int32(7)<<(uint(int32(18))%32) | v669
	v700 = int32(4)
	goto L104
L113:
	;
	v719 = v631 - v700
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v719
	v631 = v719
	goto L101
L114:
	;
	v704 = v699 - int32(101)
	if v704 < int32(0) {
		goto L113
	} else {
		goto L115
	}
L115:
	;
	v710 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v704)>>(uint(int32(3))%32)))+uint32(_consts[1331]))))
	if int32(base.Ui32(v710)>>(uint(v704&int32(7))%32))&int32(1) == int32(0) {
		goto L113
	} else {
		goto L116
	}
L116:
	;
	v734 = v700
	goto L99
L118:
	;
	v737 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v739 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v740 = v737 - v292
	v741 = v739
	goto L96
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v888
	v890 = int32(2)
	v892 = int32(0)
	v895 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v888-v895 < v890 {
		v905 = v892
		goto L143
	} else {
		goto L144
	}
L120:
	;
	v744 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v748 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744+v740-int32(1)))))
	if v748 != int32(111) {
		v888 = v740
		goto L119
	} else {
		goto L121
	}
L121:
	;
	v752 = v740 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v752
	v767 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v771 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v780 = v752
	goto L124
L122:
	;
	if int32(0) <= v883 {
		goto L22
	} else {
		goto L141
	}
L123:
	;
	v883 = int32(-1)
	goto L122
L124:
	;
	if v780 <= v771 {
		goto L123
	} else {
		goto L126
	}
L126:
	;
	v788 = int32(1)
	v789 = v780 - v788
	v791 = int32(*(*int8)(unsafe.Add(mBase, uint32(v767+v789))))
	v793 = v791 & int32(255)
	if v789 == v771 {
		v848 = v793
		v849 = v788
		goto L127
	} else {
		goto L128
	}
L127:
	;
	if int32(117) < v848 {
		goto L136
	} else {
		goto L137
	}
L128:
	;
	if int32(0) <= v791 {
		v848 = v793
		v849 = v788
		goto L127
	} else {
		goto L129
	}
L129:
	;
	v799 = v793 & int32(63)
	v801 = v780 - int32(2)
	v803 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v767+v801))))
	v805 = v803 << (uint(int32(6)) % 32)
	if base.B2i32(v801 != v771)&base.B2i32(base.Ui32(v803) < base.Ui32(int32(192))) == int32(0) {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v848 = v805&int32(1984) | v799
	v849 = int32(2)
	goto L127
L131:
	;
	goto L132
L132:
	;
	v818 = v805&int32(4032) | v799
	v820 = v780 - int32(3)
	v822 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v767+v820))))
	if base.B2i32(v820 != v771)&base.B2i32(base.Ui32(v822) < base.Ui32(int32(224))) == int32(0) {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v848 = v822<<(uint(int32(12))%32)&int32(61440) | v818
	v849 = int32(3)
	goto L127
L134:
	;
	goto L135
L135:
	;
	v840 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v780+(v767-int32(4))))))
	v848 = v822<<(uint(int32(12))%32)&int32(258048) | v840&int32(7)<<(uint(int32(18))%32) | v818
	v849 = int32(4)
	goto L127
L136:
	;
	v868 = v780 - v849
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v868
	v780 = v868
	goto L124
L137:
	;
	v853 = v848 - int32(111)
	if v853 < int32(0) {
		goto L136
	} else {
		goto L138
	}
L138:
	;
	v859 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v853)>>(uint(int32(3))%32)))+uint32(_consts[1332]))))
	if int32(base.Ui32(v859)>>(uint(v853&int32(7))%32))&int32(1) == int32(0) {
		goto L136
	} else {
		goto L139
	}
L139:
	;
	v883 = v849
	goto L122
L141:
	;
	v886 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v888 = v886 - v292
	goto L119
L142:
	;
	if v905 != 0 {
		goto L146
	} else {
		goto L147
	}
L143:
	;
	goto L142
L144:
	;
	v898 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v901 = F_memcmp(m, v898+v888-v890, int32(2198557), v890)
	mBase = m.M
	if v901 != 0 {
		v905 = v892
		goto L143
	} else {
		goto L145
	}
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v888 - v890
	v905 = int32(1)
	goto L143
L146:
	;
	v919 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v922 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v923 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v932 = v922
	goto L151
L147:
	;
	goto L148
L148:
	;
	v1038 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1039 = v1038 - v292
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1039
	v1041 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1039 <= v1041 {
		v1187 = v1039
		goto L169
	} else {
		goto L170
	}
L149:
	;
	if int32(0) <= v1035 {
		goto L22
	} else {
		goto L168
	}
L150:
	;
	v1035 = int32(-1)
	goto L149
L151:
	;
	if v932 <= v923 {
		goto L150
	} else {
		goto L153
	}
L153:
	;
	v940 = int32(1)
	v941 = v932 - v940
	v943 = int32(*(*int8)(unsafe.Add(mBase, uint32(v919+v941))))
	v945 = v943 & int32(255)
	if v941 == v923 {
		v1000 = v945
		v1001 = v940
		goto L154
	} else {
		goto L155
	}
L154:
	;
	if int32(252) < v1000 {
		goto L163
	} else {
		goto L164
	}
L155:
	;
	if int32(0) <= v943 {
		v1000 = v945
		v1001 = v940
		goto L154
	} else {
		goto L156
	}
L156:
	;
	v951 = v945 & int32(63)
	v953 = v932 - int32(2)
	v955 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v919+v953))))
	v957 = v955 << (uint(int32(6)) % 32)
	if base.B2i32(v953 != v923)&base.B2i32(base.Ui32(v955) < base.Ui32(int32(192))) == int32(0) {
		goto L157
	} else {
		goto L158
	}
L157:
	;
	v1000 = v957&int32(1984) | v951
	v1001 = int32(2)
	goto L154
L158:
	;
	goto L159
L159:
	;
	v970 = v957&int32(4032) | v951
	v972 = v932 - int32(3)
	v974 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v919+v972))))
	if base.B2i32(v972 != v923)&base.B2i32(base.Ui32(v974) < base.Ui32(int32(224))) == int32(0) {
		goto L160
	} else {
		goto L161
	}
L160:
	;
	v1000 = v974<<(uint(int32(12))%32)&int32(61440) | v970
	v1001 = int32(3)
	goto L154
L161:
	;
	goto L162
L162:
	;
	v992 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v932+(v919-int32(4))))))
	v1000 = v974<<(uint(int32(12))%32)&int32(258048) | v992&int32(7)<<(uint(int32(18))%32) | v970
	v1001 = int32(4)
	goto L154
L163:
	;
	v1020 = v932 - v1001
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1020
	v932 = v1020
	goto L151
L164:
	;
	v1005 = v1000 - int32(246)
	if v1005 < int32(0) {
		goto L163
	} else {
		goto L165
	}
L165:
	;
	v1011 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1005)>>(uint(int32(3))%32)))+uint32(_consts[1333]))))
	if int32(base.Ui32(v1011)>>(uint(v1005&int32(7))%32))&int32(1) == int32(0) {
		goto L163
	} else {
		goto L166
	}
L166:
	;
	v1035 = v1001
	goto L149
L168:
	;
	goto L148
L169:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1187
	v1189 = int32(2)
	v1191 = int32(0)
	v1194 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1187-v1194 < v1189 {
		v1204 = v1191
		goto L193
	} else {
		goto L194
	}
L170:
	;
	v1043 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1047 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1043+v1039-int32(1)))))
	if v1047 != int32(117) {
		v1187 = v1039
		goto L169
	} else {
		goto L171
	}
L171:
	;
	v1051 = v1039 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1051
	v1066 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1070 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1079 = v1051
	goto L174
L172:
	;
	if int32(0) <= v1182 {
		goto L22
	} else {
		goto L191
	}
L173:
	;
	v1182 = int32(-1)
	goto L172
L174:
	;
	if v1079 <= v1070 {
		goto L173
	} else {
		goto L176
	}
L176:
	;
	v1087 = int32(1)
	v1088 = v1079 - v1087
	v1090 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1066+v1088))))
	v1092 = v1090 & int32(255)
	if v1088 == v1070 {
		v1147 = v1092
		v1148 = v1087
		goto L177
	} else {
		goto L178
	}
L177:
	;
	if int32(117) < v1147 {
		goto L186
	} else {
		goto L187
	}
L178:
	;
	if int32(0) <= v1090 {
		v1147 = v1092
		v1148 = v1087
		goto L177
	} else {
		goto L179
	}
L179:
	;
	v1098 = v1092 & int32(63)
	v1100 = v1079 - int32(2)
	v1102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1066+v1100))))
	v1104 = v1102 << (uint(int32(6)) % 32)
	if base.B2i32(v1100 != v1070)&base.B2i32(base.Ui32(v1102) < base.Ui32(int32(192))) == int32(0) {
		goto L180
	} else {
		goto L181
	}
L180:
	;
	v1147 = v1104&int32(1984) | v1098
	v1148 = int32(2)
	goto L177
L181:
	;
	goto L182
L182:
	;
	v1117 = v1104&int32(4032) | v1098
	v1119 = v1079 - int32(3)
	v1121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1066+v1119))))
	if base.B2i32(v1119 != v1070)&base.B2i32(base.Ui32(v1121) < base.Ui32(int32(224))) == int32(0) {
		goto L183
	} else {
		goto L184
	}
L183:
	;
	v1147 = v1121<<(uint(int32(12))%32)&int32(61440) | v1117
	v1148 = int32(3)
	goto L177
L184:
	;
	goto L185
L185:
	;
	v1139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1079+(v1066-int32(4))))))
	v1147 = v1121<<(uint(int32(12))%32)&int32(258048) | v1139&int32(7)<<(uint(int32(18))%32) | v1117
	v1148 = int32(4)
	goto L177
L186:
	;
	v1167 = v1079 - v1148
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1167
	v1079 = v1167
	goto L174
L187:
	;
	v1152 = v1147 - int32(111)
	if v1152 < int32(0) {
		goto L186
	} else {
		goto L188
	}
L188:
	;
	v1158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1152)>>(uint(int32(3))%32)))+uint32(_consts[1332]))))
	if int32(base.Ui32(v1158)>>(uint(v1152&int32(7))%32))&int32(1) == int32(0) {
		goto L186
	} else {
		goto L189
	}
L189:
	;
	v1182 = v1148
	goto L172
L191:
	;
	v1185 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1187 = v1185 - v292
	goto L169
L192:
	;
	if v1204 == int32(0) {
		v1350 = v2
		goto L1
	} else {
		goto L196
	}
L193:
	;
	goto L192
L194:
	;
	v1197 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1200 = F_memcmp(m, v1197+v1187-v1189, int32(2198560), v1189)
	mBase = m.M
	if v1200 != 0 {
		v1204 = v1191
		goto L193
	} else {
		goto L195
	}
L195:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1187 - v1189
	v1204 = int32(1)
	goto L193
L196:
	;
	v1220 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1223 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1224 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1233 = v1223
	goto L199
L197:
	;
	if v1336 < int32(0) {
		v1350 = v2
		goto L1
	} else {
		goto L216
	}
L198:
	;
	v1336 = int32(-1)
	goto L197
L199:
	;
	if v1233 <= v1224 {
		goto L198
	} else {
		goto L201
	}
L201:
	;
	v1241 = int32(1)
	v1242 = v1233 - v1241
	v1244 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1220+v1242))))
	v1246 = v1244 & int32(255)
	if v1242 == v1224 {
		v1301 = v1246
		v1302 = v1241
		goto L202
	} else {
		goto L203
	}
L202:
	;
	if int32(252) < v1301 {
		goto L211
	} else {
		goto L212
	}
L203:
	;
	if int32(0) <= v1244 {
		v1301 = v1246
		v1302 = v1241
		goto L202
	} else {
		goto L204
	}
L204:
	;
	v1252 = v1246 & int32(63)
	v1254 = v1233 - int32(2)
	v1256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1220+v1254))))
	v1258 = v1256 << (uint(int32(6)) % 32)
	if base.B2i32(v1254 != v1224)&base.B2i32(base.Ui32(v1256) < base.Ui32(int32(192))) == int32(0) {
		goto L205
	} else {
		goto L206
	}
L205:
	;
	v1301 = v1258&int32(1984) | v1252
	v1302 = int32(2)
	goto L202
L206:
	;
	goto L207
L207:
	;
	v1271 = v1258&int32(4032) | v1252
	v1273 = v1233 - int32(3)
	v1275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1220+v1273))))
	if base.B2i32(v1273 != v1224)&base.B2i32(base.Ui32(v1275) < base.Ui32(int32(224))) == int32(0) {
		goto L208
	} else {
		goto L209
	}
L208:
	;
	v1301 = v1275<<(uint(int32(12))%32)&int32(61440) | v1271
	v1302 = int32(3)
	goto L202
L209:
	;
	goto L210
L210:
	;
	v1293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1233+(v1220-int32(4))))))
	v1301 = v1275<<(uint(int32(12))%32)&int32(258048) | v1293&int32(7)<<(uint(int32(18))%32) | v1271
	v1302 = int32(4)
	goto L202
L211:
	;
	v1321 = v1233 - v1302
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1321
	v1233 = v1321
	goto L199
L212:
	;
	v1306 = v1301 - int32(246)
	if v1306 < int32(0) {
		goto L211
	} else {
		goto L213
	}
L213:
	;
	v1312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1306)>>(uint(int32(3))%32)))+uint32(_consts[1333]))))
	if int32(base.Ui32(v1312)>>(uint(v1306&int32(7))%32))&int32(1) == int32(0) {
		goto L211
	} else {
		goto L214
	}
L214:
	;
	v1336 = v1302
	goto L197
L216:
	;
	goto L22
}
func F_r_consonant_pair_2(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	if v7 < v9 {
		v112 = int32(0)
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v112
L2:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v7
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v9
	v16 = v7 - int32(1)
	if v16 <= v9 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v28 = F_find_among_b(m, l0, int32(4226080), int32(4))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v13
	return int32(0)
L5:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18+v16))))
	switch v20 - int32(100) {
	case 0, 16:
		goto L3
	default:
		goto L4
	}
L6:
	;
	return int32(0)
L7:
	;
	if v28 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v13
	return int32(0)
L9:
	;
	goto L10
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v13
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v38
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v42 = v40 + (v7 - v11)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v42
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L13
L11:
	;
	if v96 < int32(0) {
		v112 = int32(0)
		goto L1
	} else {
		goto L31
	}
L13:
	;
	goto L14
L14:
	;
	goto L15
L15:
	;
	v52 = v42
	v54 = int32(1)
	goto L18
L17:
	;
	v96 = v78
	goto L11
L18:
	;
	if v52 <= v13 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	goto L17
L20:
	;
	v96 = int32(-1)
	goto L11
L21:
	;
	goto L22
L22:
	;
	v59 = v52 - int32(1)
	v61 = int32(*(*int8)(unsafe.Add(mBase, uint32(v45+v59))))
	if int32(0) <= v61 {
		v78 = v59
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v82 = int32(1)
	if v82 < v54 {
		v52 = v78
		v54 = v54 - v82
		goto L18
	} else {
		goto L30
	}
L24:
	;
	if v59 <= v13 {
		v78 = v59
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v66 = v59
	goto L26
L26:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45+v66))))
	if base.Ui32(int32(191)) < base.Ui32(v71) {
		v78 = v66
		goto L23
	} else {
		goto L28
	}
L27:
	;
	v78 = v13
	goto L23
L28:
	;
	v75 = v66 - int32(1)
	if v13 < v75 {
		v66 = v75
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	goto L19
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v96
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v96
	v102 = F_slice_del(m, l0)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L6
	} else {
		goto L32
	}
L32:
	;
	if int32(0) <= v102 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v109 = int32(1)
	goto L35
L34:
	;
	v109 = v102 >> (uint(int32(31)) % 32) & v102
	goto L35
L35:
	;
	v112 = v109
	goto L1
}
func F_r_remove_second_order_prefix_1(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
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
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
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
	var v66 int32
	_ = v66
	v2 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v4
	v7 = v4 + int32(1)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v8 <= v7 {
		v66 = v2
		return v66
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10+v7))))
		if v12 != int32(101) {
			v66 = v2
			return v66
		} else {
			v17 = F_find_among(m, l0, int32(4174048), int32(6))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				if v17 == int32(0) {
					v66 = v2
					return v66
				} else {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v23
					v25 = int32(1)
					switch v17 - v25 {
					case 0:
						v28 = F_slice_del(m, l0)
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return int32(0)
						} else {
							if v28 < int32(0) {
								v66 = v28
							} else {
								v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
								*(*int32)(unsafe.Add(mBase, uint32(v32))) = int32(2)
								v58 = v32
								v60 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
								v61 = int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(v58)+4)) = v60 - v61
								v66 = v61
							}
							return v66
						}
					case 1:
						v37 = F_slice_from_s(m, l0, int32(4), int32(2150665))
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int32(0)
						} else {
							if v37 < int32(0) {
								v66 = v37
							} else {
								v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
								v58 = v41
								v60 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
								v61 = int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(v58)+4)) = v60 - v61
								v66 = v61
							}
							return v66
						}
					case 2:
						v42 = F_slice_del(m, l0)
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return int32(0)
						} else {
							if v42 < int32(0) {
								v66 = v42
							} else {
								v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
								*(*int32)(unsafe.Add(mBase, uint32(v46))) = int32(4)
								v58 = v46
								v60 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
								v61 = int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(v58)+4)) = v60 - v61
								v66 = v61
							}
							return v66
						}
					case 3:
						v51 = F_slice_from_s(m, l0, int32(4), int32(2150669))
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return int32(0)
						} else {
							if v51 < int32(0) {
								v66 = v51
							} else {
								v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
								*(*int32)(unsafe.Add(mBase, uint32(v55))) = int32(4)
								v58 = v55
								v60 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
								v61 = int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(v58)+4)) = v60 - v61
								v66 = v61
							}
							return v66
						}
					default:
						v66 = v25
						return v66
					}
				}
			}
		}
	}
}
func F_raw_expression_tree_walker_impl(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
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
	var v47 int32
	_ = v47
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
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
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
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
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
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
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
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
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
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
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
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
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
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
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
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
	var v224 int32
	_ = v224
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
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
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
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
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
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
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
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
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
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v744 int32
	_ = v744
	var v749 int32
	_ = v749
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v773 int32
	_ = v773
	v4 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	if l0 == v4 {
		v773 = v4
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v11 + int32(16)
	return v773
L2:
	;
	F_check_stack_depth(m)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v19 - int32(1) {
	case 0:
		goto L78
	case 1, 12, 39, 41, 56, 57, 68, 69, 71, 76, 116, 464, 465, 466, 467, 468:
		goto L5
	case 2:
		goto L77
	default:
		goto L7
	case 4:
		goto L53
	case 9:
		goto L76
	case 15:
		goto L40
	case 20:
		goto L42
	case 21:
		goto L75
	case 31:
		goto L74
	case 35:
		goto L73
	case 37:
		goto L72
	case 38:
		goto L71
	case 40:
		goto L70
	case 42:
		goto L69
	case 43:
		goto L68
	case 44:
		goto L64
	case 45:
		goto L63
	case 46:
		goto L60
	case 51:
		goto L56
	case 52:
		goto L55
	case 63:
		goto L54
	case 67:
		goto L25
	case 70:
		goto L43
	case 72:
		goto L34
	case 73:
		goto L33
	case 75:
		goto L41
	case 77:
		goto L39
	case 78:
		goto L38
	case 79:
		goto L37
	case 80:
		goto L36
	case 81:
		goto L35
	case 82:
		goto L32
	case 83:
		goto L31
	case 84:
		goto L30
	case 85:
		goto L29
	case 86:
		goto L27
	case 87:
		goto L26
	case 88:
		goto L28
	case 89:
		goto L24
	case 91:
		goto L23
	case 93:
		goto L21
	case 94:
		goto L20
	case 106:
		goto L22
	case 109:
		goto L19
	case 110:
		goto L18
	case 111:
		goto L17
	case 114:
		goto L16
	case 115:
		goto L47
	case 117:
		goto L46
	case 119:
		goto L15
	case 120:
		goto L62
	case 121:
		goto L61
	case 122:
		goto L57
	case 123:
		goto L59
	case 124:
		goto L58
	case 125:
		goto L14
	case 126:
		goto L67
	case 127:
		goto L66
	case 128:
		goto L65
	case 129:
		goto L13
	case 130:
		goto L12
	case 131:
		goto L8
	case 132:
		goto L11
	case 133:
		goto L10
	case 134:
		goto L9
	case 136:
		goto L51
	case 137:
		goto L50
	case 138:
		goto L49
	case 139:
		goto L48
	case 140:
		goto L45
	case 143:
		goto L44
	}
L5:
	;
	v773 = int32(0)
	goto L1
L6:
	;
	v758 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v759 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v758, l2)
	mBase = m.M
	v760 = m.ExcPending
	if v760 != 0 {
		goto L3
	} else {
		goto L425
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v739 = m.ExcPending
	if v739 != 0 {
		goto L3
	} else {
		goto L422
	}
L8:
	;
	v727 = int32(1)
	v728 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v729 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v728, l2)
	mBase = m.M
	v730 = m.ExcPending
	if v730 != 0 {
		goto L3
	} else {
		goto L418
	}
L9:
	;
	v718 = int32(1)
	v719 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v720 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v719, l2)
	mBase = m.M
	v721 = m.ExcPending
	if v721 != 0 {
		goto L3
	} else {
		goto L414
	}
L10:
	;
	v709 = int32(1)
	v710 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v711 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v710, l2)
	mBase = m.M
	v712 = m.ExcPending
	if v712 != 0 {
		goto L3
	} else {
		goto L410
	}
L11:
	;
	v694 = int32(1)
	v695 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v696 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v695, l2)
	mBase = m.M
	v697 = m.ExcPending
	if v697 != 0 {
		goto L3
	} else {
		goto L402
	}
L12:
	;
	v685 = int32(1)
	v686 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v687 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v686, l2)
	mBase = m.M
	v688 = m.ExcPending
	if v688 != 0 {
		goto L3
	} else {
		goto L398
	}
L13:
	;
	v676 = int32(1)
	v677 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v678 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v677, l2)
	mBase = m.M
	v679 = m.ExcPending
	if v679 != 0 {
		goto L3
	} else {
		goto L394
	}
L14:
	;
	v667 = int32(1)
	v668 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v669 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v668, l2)
	mBase = m.M
	v670 = m.ExcPending
	if v670 != 0 {
		goto L3
	} else {
		goto L390
	}
L15:
	;
	v658 = int32(1)
	v659 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v660 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v659, l2)
	mBase = m.M
	v661 = m.ExcPending
	if v661 != 0 {
		goto L3
	} else {
		goto L386
	}
L16:
	;
	v655 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v656 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v655, l2)
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		goto L3
	} else {
		goto L385
	}
L17:
	;
	v643 = int32(1)
	v644 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v645 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v644, l2)
	mBase = m.M
	v646 = m.ExcPending
	if v646 != 0 {
		goto L3
	} else {
		goto L379
	}
L18:
	;
	v634 = int32(1)
	v635 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v636 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v635, l2)
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L3
	} else {
		goto L375
	}
L19:
	;
	v631 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v632 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v631, l2)
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
		goto L3
	} else {
		goto L374
	}
L20:
	;
	v622 = int32(1)
	v623 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v624 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v623, l2)
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L3
	} else {
		goto L370
	}
L21:
	;
	v619 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v620 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v619, l2)
	mBase = m.M
	v621 = m.ExcPending
	if v621 != 0 {
		goto L3
	} else {
		goto L369
	}
L22:
	;
	v616 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v617 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v616, l2)
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L3
	} else {
		goto L368
	}
L23:
	;
	v610 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v611 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v610, l2)
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
		goto L3
	} else {
		goto L366
	}
L24:
	;
	v598 = int32(1)
	v599 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v600 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v599, l2)
	mBase = m.M
	v601 = m.ExcPending
	if v601 != 0 {
		goto L3
	} else {
		goto L360
	}
L25:
	;
	v589 = int32(1)
	v590 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v591 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v590, l2)
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L3
	} else {
		goto L356
	}
L26:
	;
	v580 = int32(1)
	v581 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v582 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v581, l2)
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L3
	} else {
		goto L352
	}
L27:
	;
	v562 = int32(1)
	v563 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v564 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v563, l2)
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L3
	} else {
		goto L342
	}
L28:
	;
	v550 = int32(1)
	v551 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v552 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v551, l2)
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
		goto L3
	} else {
		goto L336
	}
L29:
	;
	v538 = int32(1)
	v539 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v540 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v539, l2)
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L3
	} else {
		goto L330
	}
L30:
	;
	v529 = int32(1)
	v530 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v531 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v530, l2)
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L3
	} else {
		goto L326
	}
L31:
	;
	v514 = int32(1)
	v515 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v516 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v515, l2)
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L3
	} else {
		goto L318
	}
L32:
	;
	v511 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v512 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v511, l2)
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L3
	} else {
		goto L317
	}
L33:
	;
	v508 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v509 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v508, l2)
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L3
	} else {
		goto L316
	}
L34:
	;
	v499 = int32(1)
	v500 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v501 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v500, l2)
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L3
	} else {
		goto L312
	}
L35:
	;
	v496 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v497 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v496, l2)
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L3
	} else {
		goto L311
	}
L36:
	;
	v487 = int32(1)
	v488 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v489 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v488, l2)
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L3
	} else {
		goto L307
	}
L37:
	;
	v484 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v485 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v484, l2)
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L3
	} else {
		goto L306
	}
L38:
	;
	v475 = int32(1)
	v476 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v477 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v476, l2)
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L3
	} else {
		goto L302
	}
L39:
	;
	v466 = int32(1)
	v467 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v468 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v467, l2)
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L3
	} else {
		goto L298
	}
L40:
	;
	v463 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v464 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v463, l2)
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L3
	} else {
		goto L297
	}
L41:
	;
	v448 = int32(1)
	v449 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v450 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v449, l2)
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L3
	} else {
		goto L289
	}
L42:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v443 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v442, l2)
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L3
	} else {
		goto L287
	}
L43:
	;
	v433 = int32(1)
	v434 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v435 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v434, l2)
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L3
	} else {
		goto L283
	}
L44:
	;
	v424 = int32(1)
	v425 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v426 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v425, l2)
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L3
	} else {
		goto L279
	}
L45:
	;
	v373 = int32(1)
	v374 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v375 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v374, l2)
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L3
	} else {
		goto L247
	}
L46:
	;
	v364 = int32(1)
	v365 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v366 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v365, l2)
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L3
	} else {
		goto L243
	}
L47:
	;
	v352 = int32(1)
	v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v354 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v353, l2)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L3
	} else {
		goto L237
	}
L48:
	;
	v331 = int32(1)
	v332 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v333 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v332, l2)
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L3
	} else {
		goto L225
	}
L49:
	;
	v310 = int32(1)
	v311 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v312 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v311, l2)
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L3
	} else {
		goto L213
	}
L50:
	;
	v292 = int32(1)
	v293 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v294 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v293, l2)
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L3
	} else {
		goto L203
	}
L51:
	;
	v271 = int32(1)
	v272 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v273 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v272, l2)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L3
	} else {
		goto L191
	}
L52:
	;
	v252 = v4
	goto L184
L53:
	;
	v240 = int32(1)
	v241 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v242 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v241, l2)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L3
	} else {
		goto L180
	}
L54:
	;
	v225 = int32(1)
	v226 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v227 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v226, l2)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L3
	} else {
		goto L172
	}
L55:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v223 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v222, l2)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L3
	} else {
		goto L171
	}
L56:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v220 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v219, l2)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L3
	} else {
		goto L170
	}
L57:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v217 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v216, l2)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L3
	} else {
		goto L169
	}
L58:
	;
	v201 = int32(1)
	v202 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v203 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v202, l2)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L3
	} else {
		goto L161
	}
L59:
	;
	v183 = int32(1)
	v184 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v185 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v184, l2)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L3
	} else {
		goto L151
	}
L60:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v178 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v177, l2)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L3
	} else {
		goto L149
	}
L61:
	;
	v156 = int32(1)
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v158 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v157, l2)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L3
	} else {
		goto L137
	}
L62:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v154 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v153, l2)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L3
	} else {
		goto L136
	}
L63:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v151 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v150, l2)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L3
	} else {
		goto L135
	}
L64:
	;
	v135 = int32(1)
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v137 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v136, l2)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L3
	} else {
		goto L127
	}
L65:
	;
	v126 = int32(1)
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v128 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v127, l2)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L3
	} else {
		goto L123
	}
L66:
	;
	v117 = int32(1)
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v119 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v118, l2)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L3
	} else {
		goto L119
	}
L67:
	;
	v108 = int32(1)
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v110 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v109, l2)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L3
	} else {
		goto L115
	}
L68:
	;
	v96 = int32(1)
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v98 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v97, l2)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L3
	} else {
		goto L109
	}
L69:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v94 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v93, l2)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L3
	} else {
		goto L108
	}
L70:
	;
	v84 = int32(1)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v86 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v85, l2)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L3
	} else {
		goto L104
	}
L71:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v82 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v81, l2)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L3
	} else {
		goto L103
	}
L72:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v79 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v78, l2)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L3
	} else {
		goto L102
	}
L73:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v76 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v75, l2)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L3
	} else {
		goto L101
	}
L74:
	;
	v40 = int32(1)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v42 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v41, l2)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L3
	} else {
		goto L86
	}
L75:
	;
	v31 = int32(1)
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v33 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v32, l2)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L3
	} else {
		goto L82
	}
L76:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v29 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v28, l2)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L3
	} else {
		goto L81
	}
L77:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v26 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v25, l2)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L3
	} else {
		goto L80
	}
L78:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v22 <= int32(0) {
		goto L5
	} else {
		goto L79
	}
L79:
	;
	goto L52
L80:
	;
	v773 = v26
	goto L1
L81:
	;
	v773 = v29
	goto L1
L82:
	;
	if v33 != 0 {
		v773 = v31
		goto L1
	} else {
		goto L83
	}
L83:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v36 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v35, l2)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L3
	} else {
		goto L84
	}
L84:
	;
	if v36 == int32(0) {
		goto L5
	} else {
		goto L85
	}
L85:
	;
	v773 = v31
	goto L1
L86:
	;
	if v42 != 0 {
		v773 = v40
		goto L1
	} else {
		goto L87
	}
L87:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v44 == int32(0) {
		goto L6
	} else {
		goto L88
	}
L88:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	if v47 <= int32(0) {
		goto L6
	} else {
		goto L89
	}
L89:
	;
	v56 = v4
	goto L90
L90:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v44)+12))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v58+v56<<(uint(int32(2))%32))))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
	v64 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v63, l2)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L3
	} else {
		goto L92
	}
L91:
	;
	v773 = v40
	goto L1
L92:
	;
	if v64 != 0 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v773 = v40
	goto L1
L94:
	;
	goto L95
L95:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v62)+8))
	v67 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v66, l2)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L3
	} else {
		goto L96
	}
L96:
	;
	if v67 == int32(0) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v72 = v56 + int32(1)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	if v73 <= v72 {
		goto L6
	} else {
		goto L100
	}
L98:
	;
	goto L99
L99:
	;
	goto L91
L100:
	;
	v56 = v72
	goto L90
L101:
	;
	v773 = v76
	goto L1
L102:
	;
	v773 = v79
	goto L1
L103:
	;
	v773 = v82
	goto L1
L104:
	;
	if v86 != 0 {
		v773 = v84
		goto L1
	} else {
		goto L105
	}
L105:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v89 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v88, l2)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L3
	} else {
		goto L106
	}
L106:
	;
	if v89 == int32(0) {
		goto L5
	} else {
		goto L107
	}
L107:
	;
	v773 = v84
	goto L1
L108:
	;
	v773 = v94
	goto L1
L109:
	;
	if v98 != 0 {
		v773 = v96
		goto L1
	} else {
		goto L110
	}
L110:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v101 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v100, l2)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L3
	} else {
		goto L111
	}
L111:
	;
	if v101 != 0 {
		v773 = v96
		goto L1
	} else {
		goto L112
	}
L112:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v104 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v103, l2)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L3
	} else {
		goto L113
	}
L113:
	;
	if v104 == int32(0) {
		goto L5
	} else {
		goto L114
	}
L114:
	;
	v773 = v96
	goto L1
L115:
	;
	if v110 != 0 {
		v773 = v108
		goto L1
	} else {
		goto L116
	}
L116:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v113 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v112, l2)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L3
	} else {
		goto L117
	}
L117:
	;
	if v113 == int32(0) {
		goto L5
	} else {
		goto L118
	}
L118:
	;
	v773 = v108
	goto L1
L119:
	;
	if v119 != 0 {
		v773 = v117
		goto L1
	} else {
		goto L120
	}
L120:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v122 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v121, l2)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L3
	} else {
		goto L121
	}
L121:
	;
	if v122 == int32(0) {
		goto L5
	} else {
		goto L122
	}
L122:
	;
	v773 = v117
	goto L1
L123:
	;
	if v128 != 0 {
		v773 = v126
		goto L1
	} else {
		goto L124
	}
L124:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v131 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v130, l2)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L3
	} else {
		goto L125
	}
L125:
	;
	if v131 == int32(0) {
		goto L5
	} else {
		goto L126
	}
L126:
	;
	v773 = v126
	goto L1
L127:
	;
	if v137 != 0 {
		v773 = v135
		goto L1
	} else {
		goto L128
	}
L128:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v140 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v139, l2)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L3
	} else {
		goto L129
	}
L129:
	;
	if v140 != 0 {
		v773 = v135
		goto L1
	} else {
		goto L130
	}
L130:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v143 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v142, l2)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L3
	} else {
		goto L131
	}
L131:
	;
	if v143 != 0 {
		v773 = v135
		goto L1
	} else {
		goto L132
	}
L132:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v146 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v145, l2)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L3
	} else {
		goto L133
	}
L133:
	;
	if v146 == int32(0) {
		goto L5
	} else {
		goto L134
	}
L134:
	;
	v773 = v135
	goto L1
L135:
	;
	v773 = v151
	goto L1
L136:
	;
	v773 = v154
	goto L1
L137:
	;
	if v158 != 0 {
		v773 = v156
		goto L1
	} else {
		goto L138
	}
L138:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v161 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v160, l2)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L3
	} else {
		goto L139
	}
L139:
	;
	if v161 != 0 {
		v773 = v156
		goto L1
	} else {
		goto L140
	}
L140:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v164 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v163, l2)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L3
	} else {
		goto L141
	}
L141:
	;
	if v164 != 0 {
		v773 = v156
		goto L1
	} else {
		goto L142
	}
L142:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v167 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v166, l2)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L3
	} else {
		goto L143
	}
L143:
	;
	if v167 != 0 {
		v773 = v156
		goto L1
	} else {
		goto L144
	}
L144:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v170 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v169, l2)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L3
	} else {
		goto L145
	}
L145:
	;
	if v170 != 0 {
		v773 = v156
		goto L1
	} else {
		goto L146
	}
L146:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v173 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v172, l2)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L3
	} else {
		goto L147
	}
L147:
	;
	if v173 == int32(0) {
		goto L5
	} else {
		goto L148
	}
L148:
	;
	v773 = v156
	goto L1
L149:
	;
	if v178 == int32(0) {
		goto L5
	} else {
		goto L150
	}
L150:
	;
	v773 = int32(1)
	goto L1
L151:
	;
	if v185 != 0 {
		v773 = v183
		goto L1
	} else {
		goto L152
	}
L152:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v188 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v187, l2)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L3
	} else {
		goto L153
	}
L153:
	;
	if v188 != 0 {
		v773 = v183
		goto L1
	} else {
		goto L154
	}
L154:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v191 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v190, l2)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L3
	} else {
		goto L155
	}
L155:
	;
	if v191 != 0 {
		v773 = v183
		goto L1
	} else {
		goto L156
	}
L156:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v194 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v193, l2)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L3
	} else {
		goto L157
	}
L157:
	;
	if v194 != 0 {
		v773 = v183
		goto L1
	} else {
		goto L158
	}
L158:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v197 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v196, l2)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L3
	} else {
		goto L159
	}
L159:
	;
	if v197 == int32(0) {
		goto L5
	} else {
		goto L160
	}
L160:
	;
	v773 = v183
	goto L1
L161:
	;
	if v203 != 0 {
		v773 = v201
		goto L1
	} else {
		goto L162
	}
L162:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v206 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v205, l2)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L3
	} else {
		goto L163
	}
L163:
	;
	if v206 != 0 {
		v773 = v201
		goto L1
	} else {
		goto L164
	}
L164:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v209 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v208, l2)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L3
	} else {
		goto L165
	}
L165:
	;
	if v209 != 0 {
		v773 = v201
		goto L1
	} else {
		goto L166
	}
L166:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v212 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v211, l2)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L3
	} else {
		goto L167
	}
L167:
	;
	if v212 == int32(0) {
		goto L5
	} else {
		goto L168
	}
L168:
	;
	v773 = v201
	goto L1
L169:
	;
	v773 = v217
	goto L1
L170:
	;
	v773 = v220
	goto L1
L171:
	;
	v773 = v223
	goto L1
L172:
	;
	if v227 != 0 {
		v773 = v225
		goto L1
	} else {
		goto L173
	}
L173:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v230 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v229, l2)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L3
	} else {
		goto L174
	}
L174:
	;
	if v230 != 0 {
		v773 = v225
		goto L1
	} else {
		goto L175
	}
L175:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v233 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v232, l2)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L3
	} else {
		goto L176
	}
L176:
	;
	if v233 != 0 {
		v773 = v225
		goto L1
	} else {
		goto L177
	}
L177:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v236 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v235, l2)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L3
	} else {
		goto L178
	}
L178:
	;
	if v236 == int32(0) {
		goto L5
	} else {
		goto L179
	}
L179:
	;
	v773 = v225
	goto L1
L180:
	;
	if v242 != 0 {
		v773 = v240
		goto L1
	} else {
		goto L181
	}
L181:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v245 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v244, l2)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L3
	} else {
		goto L182
	}
L182:
	;
	if v245 == int32(0) {
		goto L5
	} else {
		goto L183
	}
L183:
	;
	v773 = v240
	goto L1
L184:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v257+v252<<(uint(int32(2))%32))))
	v262 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v261, l2)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L3
	} else {
		goto L186
	}
L185:
	;
	v773 = int32(1)
	goto L1
L186:
	;
	if v262 == int32(0) {
		goto L187
	} else {
		goto L188
	}
L187:
	;
	v267 = v252 + int32(1)
	v268 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v267 < v268 {
		v252 = v267
		goto L184
	} else {
		goto L190
	}
L188:
	;
	goto L189
L189:
	;
	goto L185
L190:
	;
	goto L5
L191:
	;
	if v273 != 0 {
		v773 = v271
		goto L1
	} else {
		goto L192
	}
L192:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v276 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v275, l2)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L3
	} else {
		goto L193
	}
L193:
	;
	if v276 != 0 {
		v773 = v271
		goto L1
	} else {
		goto L194
	}
L194:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v279 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v278, l2)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L3
	} else {
		goto L195
	}
L195:
	;
	if v279 != 0 {
		v773 = v271
		goto L1
	} else {
		goto L196
	}
L196:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v282 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v281, l2)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L3
	} else {
		goto L197
	}
L197:
	;
	if v282 != 0 {
		v773 = v271
		goto L1
	} else {
		goto L198
	}
L198:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v285 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v284, l2)
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L3
	} else {
		goto L199
	}
L199:
	;
	if v285 != 0 {
		v773 = v271
		goto L1
	} else {
		goto L200
	}
L200:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v288 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v287, l2)
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L3
	} else {
		goto L201
	}
L201:
	;
	if v288 == int32(0) {
		goto L5
	} else {
		goto L202
	}
L202:
	;
	v773 = v271
	goto L1
L203:
	;
	if v294 != 0 {
		v773 = v292
		goto L1
	} else {
		goto L204
	}
L204:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v297 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v296, l2)
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L3
	} else {
		goto L205
	}
L205:
	;
	if v297 != 0 {
		v773 = v292
		goto L1
	} else {
		goto L206
	}
L206:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v300 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v299, l2)
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L3
	} else {
		goto L207
	}
L207:
	;
	if v300 != 0 {
		v773 = v292
		goto L1
	} else {
		goto L208
	}
L208:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v303 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v302, l2)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L3
	} else {
		goto L209
	}
L209:
	;
	if v303 != 0 {
		v773 = v292
		goto L1
	} else {
		goto L210
	}
L210:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v306 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v305, l2)
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L3
	} else {
		goto L211
	}
L211:
	;
	if v306 == int32(0) {
		goto L5
	} else {
		goto L212
	}
L212:
	;
	v773 = v292
	goto L1
L213:
	;
	if v312 != 0 {
		v773 = v310
		goto L1
	} else {
		goto L214
	}
L214:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v315 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v314, l2)
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L3
	} else {
		goto L215
	}
L215:
	;
	if v315 != 0 {
		v773 = v310
		goto L1
	} else {
		goto L216
	}
L216:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v318 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v317, l2)
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L3
	} else {
		goto L217
	}
L217:
	;
	if v318 != 0 {
		v773 = v310
		goto L1
	} else {
		goto L218
	}
L218:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v321 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v320, l2)
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L3
	} else {
		goto L219
	}
L219:
	;
	if v321 != 0 {
		v773 = v310
		goto L1
	} else {
		goto L220
	}
L220:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v324 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v323, l2)
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L3
	} else {
		goto L221
	}
L221:
	;
	if v324 != 0 {
		v773 = v310
		goto L1
	} else {
		goto L222
	}
L222:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v327 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v326, l2)
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L3
	} else {
		goto L223
	}
L223:
	;
	if v327 == int32(0) {
		goto L5
	} else {
		goto L224
	}
L224:
	;
	v773 = v310
	goto L1
L225:
	;
	if v333 != 0 {
		v773 = v331
		goto L1
	} else {
		goto L226
	}
L226:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v336 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v335, l2)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L3
	} else {
		goto L227
	}
L227:
	;
	if v336 != 0 {
		v773 = v331
		goto L1
	} else {
		goto L228
	}
L228:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v339 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v338, l2)
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L3
	} else {
		goto L229
	}
L229:
	;
	if v339 != 0 {
		v773 = v331
		goto L1
	} else {
		goto L230
	}
L230:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v342 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v341, l2)
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L3
	} else {
		goto L231
	}
L231:
	;
	if v342 != 0 {
		v773 = v331
		goto L1
	} else {
		goto L232
	}
L232:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v345 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v344, l2)
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L3
	} else {
		goto L233
	}
L233:
	;
	if v345 != 0 {
		v773 = v331
		goto L1
	} else {
		goto L234
	}
L234:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v348 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v347, l2)
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L3
	} else {
		goto L235
	}
L235:
	;
	if v348 == int32(0) {
		goto L5
	} else {
		goto L236
	}
L236:
	;
	v773 = v331
	goto L1
L237:
	;
	if v354 != 0 {
		v773 = v352
		goto L1
	} else {
		goto L238
	}
L238:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v357 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v356, l2)
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L3
	} else {
		goto L239
	}
L239:
	;
	if v357 != 0 {
		v773 = v352
		goto L1
	} else {
		goto L240
	}
L240:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v360 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v359, l2)
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L3
	} else {
		goto L241
	}
L241:
	;
	if v360 == int32(0) {
		goto L5
	} else {
		goto L242
	}
L242:
	;
	v773 = v352
	goto L1
L243:
	;
	if v366 != 0 {
		v773 = v364
		goto L1
	} else {
		goto L244
	}
L244:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v369 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v368, l2)
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L3
	} else {
		goto L245
	}
L245:
	;
	if v369 == int32(0) {
		goto L5
	} else {
		goto L246
	}
L246:
	;
	v773 = v364
	goto L1
L247:
	;
	if v375 != 0 {
		v773 = v373
		goto L1
	} else {
		goto L248
	}
L248:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v378 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v377, l2)
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L3
	} else {
		goto L249
	}
L249:
	;
	if v378 != 0 {
		v773 = v373
		goto L1
	} else {
		goto L250
	}
L250:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v381 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v380, l2)
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L3
	} else {
		goto L251
	}
L251:
	;
	if v381 != 0 {
		v773 = v373
		goto L1
	} else {
		goto L252
	}
L252:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v384 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v383, l2)
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L3
	} else {
		goto L253
	}
L253:
	;
	if v384 != 0 {
		v773 = v373
		goto L1
	} else {
		goto L254
	}
L254:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v387 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v386, l2)
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L3
	} else {
		goto L255
	}
L255:
	;
	if v387 != 0 {
		v773 = v373
		goto L1
	} else {
		goto L256
	}
L256:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v390 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v389, l2)
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L3
	} else {
		goto L257
	}
L257:
	;
	if v390 != 0 {
		v773 = v373
		goto L1
	} else {
		goto L258
	}
L258:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v393 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v392, l2)
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L3
	} else {
		goto L259
	}
L259:
	;
	if v393 != 0 {
		v773 = v373
		goto L1
	} else {
		goto L260
	}
L260:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v396 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v395, l2)
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L3
	} else {
		goto L261
	}
L261:
	;
	if v396 != 0 {
		v773 = v373
		goto L1
	} else {
		goto L262
	}
L262:
	;
	v398 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v399 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v398, l2)
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L3
	} else {
		goto L263
	}
L263:
	;
	if v399 != 0 {
		v773 = v373
		goto L1
	} else {
		goto L264
	}
L264:
	;
	v401 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v402 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v401, l2)
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L3
	} else {
		goto L265
	}
L265:
	;
	if v402 != 0 {
		v773 = v373
		goto L1
	} else {
		goto L266
	}
L266:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v405 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v404, l2)
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L3
	} else {
		goto L267
	}
L267:
	;
	if v405 != 0 {
		v773 = v373
		goto L1
	} else {
		goto L268
	}
L268:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v408 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v407, l2)
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L3
	} else {
		goto L269
	}
L269:
	;
	if v408 != 0 {
		v773 = v373
		goto L1
	} else {
		goto L270
	}
L270:
	;
	v410 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v411 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v410, l2)
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L3
	} else {
		goto L271
	}
L271:
	;
	if v411 != 0 {
		v773 = v373
		goto L1
	} else {
		goto L272
	}
L272:
	;
	v413 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v414 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v413, l2)
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L3
	} else {
		goto L273
	}
L273:
	;
	if v414 != 0 {
		v773 = v373
		goto L1
	} else {
		goto L274
	}
L274:
	;
	v416 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v417 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v416, l2)
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L3
	} else {
		goto L275
	}
L275:
	;
	if v417 != 0 {
		v773 = v373
		goto L1
	} else {
		goto L276
	}
L276:
	;
	v419 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v420 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v419, l2)
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L3
	} else {
		goto L277
	}
L277:
	;
	if v420 == int32(0) {
		goto L5
	} else {
		goto L278
	}
L278:
	;
	v773 = v373
	goto L1
L279:
	;
	if v426 != 0 {
		v773 = v424
		goto L1
	} else {
		goto L280
	}
L280:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v429 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v428, l2)
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L3
	} else {
		goto L281
	}
L281:
	;
	if v429 == int32(0) {
		goto L5
	} else {
		goto L282
	}
L282:
	;
	v773 = v424
	goto L1
L283:
	;
	if v435 != 0 {
		v773 = v433
		goto L1
	} else {
		goto L284
	}
L284:
	;
	v437 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v438 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v437, l2)
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L3
	} else {
		goto L285
	}
L285:
	;
	if v438 == int32(0) {
		goto L5
	} else {
		goto L286
	}
L286:
	;
	v773 = v433
	goto L1
L287:
	;
	if v443 == int32(0) {
		goto L5
	} else {
		goto L288
	}
L288:
	;
	v773 = int32(1)
	goto L1
L289:
	;
	if v450 != 0 {
		v773 = v448
		goto L1
	} else {
		goto L290
	}
L290:
	;
	v452 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v453 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v452, l2)
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L3
	} else {
		goto L291
	}
L291:
	;
	if v453 != 0 {
		v773 = v448
		goto L1
	} else {
		goto L292
	}
L292:
	;
	v455 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v456 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v455, l2)
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L3
	} else {
		goto L293
	}
L293:
	;
	if v456 != 0 {
		v773 = v448
		goto L1
	} else {
		goto L294
	}
L294:
	;
	v458 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v459 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v458, l2)
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L3
	} else {
		goto L295
	}
L295:
	;
	if v459 == int32(0) {
		goto L5
	} else {
		goto L296
	}
L296:
	;
	v773 = v448
	goto L1
L297:
	;
	v773 = v464
	goto L1
L298:
	;
	if v468 != 0 {
		v773 = v466
		goto L1
	} else {
		goto L299
	}
L299:
	;
	v470 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v471 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v470, l2)
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L3
	} else {
		goto L300
	}
L300:
	;
	if v471 == int32(0) {
		goto L5
	} else {
		goto L301
	}
L301:
	;
	v773 = v466
	goto L1
L302:
	;
	if v477 != 0 {
		v773 = v475
		goto L1
	} else {
		goto L303
	}
L303:
	;
	v479 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v480 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v479, l2)
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L3
	} else {
		goto L304
	}
L304:
	;
	if v480 == int32(0) {
		goto L5
	} else {
		goto L305
	}
L305:
	;
	v773 = v475
	goto L1
L306:
	;
	v773 = v485
	goto L1
L307:
	;
	if v489 != 0 {
		v773 = v487
		goto L1
	} else {
		goto L308
	}
L308:
	;
	v491 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v492 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v491, l2)
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L3
	} else {
		goto L309
	}
L309:
	;
	if v492 == int32(0) {
		goto L5
	} else {
		goto L310
	}
L310:
	;
	v773 = v487
	goto L1
L311:
	;
	v773 = v497
	goto L1
L312:
	;
	if v501 != 0 {
		v773 = v499
		goto L1
	} else {
		goto L313
	}
L313:
	;
	v503 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v504 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v503, l2)
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L3
	} else {
		goto L314
	}
L314:
	;
	if v504 == int32(0) {
		goto L5
	} else {
		goto L315
	}
L315:
	;
	v773 = v499
	goto L1
L316:
	;
	v773 = v509
	goto L1
L317:
	;
	v773 = v512
	goto L1
L318:
	;
	if v516 != 0 {
		v773 = v514
		goto L1
	} else {
		goto L319
	}
L319:
	;
	v518 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v519 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v518, l2)
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L3
	} else {
		goto L320
	}
L320:
	;
	if v519 != 0 {
		v773 = v514
		goto L1
	} else {
		goto L321
	}
L321:
	;
	v521 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v522 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v521, l2)
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L3
	} else {
		goto L322
	}
L322:
	;
	if v522 != 0 {
		v773 = v514
		goto L1
	} else {
		goto L323
	}
L323:
	;
	v524 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v525 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v524, l2)
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L3
	} else {
		goto L324
	}
L324:
	;
	if v525 == int32(0) {
		goto L5
	} else {
		goto L325
	}
L325:
	;
	v773 = v514
	goto L1
L326:
	;
	if v531 != 0 {
		v773 = v529
		goto L1
	} else {
		goto L327
	}
L327:
	;
	v533 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v534 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v533, l2)
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L3
	} else {
		goto L328
	}
L328:
	;
	if v534 == int32(0) {
		goto L5
	} else {
		goto L329
	}
L329:
	;
	v773 = v529
	goto L1
L330:
	;
	if v540 != 0 {
		v773 = v538
		goto L1
	} else {
		goto L331
	}
L331:
	;
	v542 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v543 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v542, l2)
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L3
	} else {
		goto L332
	}
L332:
	;
	if v543 != 0 {
		v773 = v538
		goto L1
	} else {
		goto L333
	}
L333:
	;
	v545 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v546 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v545, l2)
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L3
	} else {
		goto L334
	}
L334:
	;
	if v546 == int32(0) {
		goto L5
	} else {
		goto L335
	}
L335:
	;
	v773 = v538
	goto L1
L336:
	;
	if v552 != 0 {
		v773 = v550
		goto L1
	} else {
		goto L337
	}
L337:
	;
	v554 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v555 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v554, l2)
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L3
	} else {
		goto L338
	}
L338:
	;
	if v555 != 0 {
		v773 = v550
		goto L1
	} else {
		goto L339
	}
L339:
	;
	v557 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v558 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v557, l2)
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L3
	} else {
		goto L340
	}
L340:
	;
	if v558 == int32(0) {
		goto L5
	} else {
		goto L341
	}
L341:
	;
	v773 = v550
	goto L1
L342:
	;
	if v564 != 0 {
		v773 = v562
		goto L1
	} else {
		goto L343
	}
L343:
	;
	v566 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v567 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v566, l2)
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L3
	} else {
		goto L344
	}
L344:
	;
	if v567 != 0 {
		v773 = v562
		goto L1
	} else {
		goto L345
	}
L345:
	;
	v569 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v570 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v569, l2)
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L3
	} else {
		goto L346
	}
L346:
	;
	if v570 != 0 {
		v773 = v562
		goto L1
	} else {
		goto L347
	}
L347:
	;
	v572 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v573 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v572, l2)
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L3
	} else {
		goto L348
	}
L348:
	;
	if v573 != 0 {
		v773 = v562
		goto L1
	} else {
		goto L349
	}
L349:
	;
	v575 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v576 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v575, l2)
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L3
	} else {
		goto L350
	}
L350:
	;
	if v576 == int32(0) {
		goto L5
	} else {
		goto L351
	}
L351:
	;
	v773 = v562
	goto L1
L352:
	;
	if v582 != 0 {
		v773 = v580
		goto L1
	} else {
		goto L353
	}
L353:
	;
	v584 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v585 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v584, l2)
	mBase = m.M
	v586 = m.ExcPending
	if v586 != 0 {
		goto L3
	} else {
		goto L354
	}
L354:
	;
	if v585 == int32(0) {
		goto L5
	} else {
		goto L355
	}
L355:
	;
	v773 = v580
	goto L1
L356:
	;
	if v591 != 0 {
		v773 = v589
		goto L1
	} else {
		goto L357
	}
L357:
	;
	v593 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v594 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v593, l2)
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L3
	} else {
		goto L358
	}
L358:
	;
	if v594 == int32(0) {
		goto L5
	} else {
		goto L359
	}
L359:
	;
	v773 = v589
	goto L1
L360:
	;
	if v600 != 0 {
		v773 = v598
		goto L1
	} else {
		goto L361
	}
L361:
	;
	v602 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v603 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v602, l2)
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L3
	} else {
		goto L362
	}
L362:
	;
	if v603 != 0 {
		v773 = v598
		goto L1
	} else {
		goto L363
	}
L363:
	;
	v605 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v606 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v605, l2)
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L3
	} else {
		goto L364
	}
L364:
	;
	if v606 == int32(0) {
		goto L5
	} else {
		goto L365
	}
L365:
	;
	v773 = v598
	goto L1
L366:
	;
	if v611 == int32(0) {
		goto L5
	} else {
		goto L367
	}
L367:
	;
	v773 = int32(1)
	goto L1
L368:
	;
	v773 = v617
	goto L1
L369:
	;
	v773 = v620
	goto L1
L370:
	;
	if v624 != 0 {
		v773 = v622
		goto L1
	} else {
		goto L371
	}
L371:
	;
	v626 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v627 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v626, l2)
	mBase = m.M
	v628 = m.ExcPending
	if v628 != 0 {
		goto L3
	} else {
		goto L372
	}
L372:
	;
	if v627 == int32(0) {
		goto L5
	} else {
		goto L373
	}
L373:
	;
	v773 = v622
	goto L1
L374:
	;
	v773 = v632
	goto L1
L375:
	;
	if v636 != 0 {
		v773 = v634
		goto L1
	} else {
		goto L376
	}
L376:
	;
	v638 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v639 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v638, l2)
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		goto L3
	} else {
		goto L377
	}
L377:
	;
	if v639 == int32(0) {
		goto L5
	} else {
		goto L378
	}
L378:
	;
	v773 = v634
	goto L1
L379:
	;
	if v645 != 0 {
		v773 = v643
		goto L1
	} else {
		goto L380
	}
L380:
	;
	v647 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v648 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v647, l2)
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L3
	} else {
		goto L381
	}
L381:
	;
	if v648 != 0 {
		v773 = v643
		goto L1
	} else {
		goto L382
	}
L382:
	;
	v650 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v651 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v650, l2)
	mBase = m.M
	v652 = m.ExcPending
	if v652 != 0 {
		goto L3
	} else {
		goto L383
	}
L383:
	;
	if v651 == int32(0) {
		goto L5
	} else {
		goto L384
	}
L384:
	;
	v773 = v643
	goto L1
L385:
	;
	v773 = v656
	goto L1
L386:
	;
	if v660 != 0 {
		v773 = v658
		goto L1
	} else {
		goto L387
	}
L387:
	;
	v662 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v663 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v662, l2)
	mBase = m.M
	v664 = m.ExcPending
	if v664 != 0 {
		goto L3
	} else {
		goto L388
	}
L388:
	;
	if v663 == int32(0) {
		goto L5
	} else {
		goto L389
	}
L389:
	;
	v773 = v658
	goto L1
L390:
	;
	if v669 != 0 {
		v773 = v667
		goto L1
	} else {
		goto L391
	}
L391:
	;
	v671 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v672 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v671, l2)
	mBase = m.M
	v673 = m.ExcPending
	if v673 != 0 {
		goto L3
	} else {
		goto L392
	}
L392:
	;
	if v672 == int32(0) {
		goto L5
	} else {
		goto L393
	}
L393:
	;
	v773 = v667
	goto L1
L394:
	;
	if v678 != 0 {
		v773 = v676
		goto L1
	} else {
		goto L395
	}
L395:
	;
	v680 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v681 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v680, l2)
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
		goto L3
	} else {
		goto L396
	}
L396:
	;
	if v681 == int32(0) {
		goto L5
	} else {
		goto L397
	}
L397:
	;
	v773 = v676
	goto L1
L398:
	;
	if v687 != 0 {
		v773 = v685
		goto L1
	} else {
		goto L399
	}
L399:
	;
	v689 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v690 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v689, l2)
	mBase = m.M
	v691 = m.ExcPending
	if v691 != 0 {
		goto L3
	} else {
		goto L400
	}
L400:
	;
	if v690 == int32(0) {
		goto L5
	} else {
		goto L401
	}
L401:
	;
	v773 = v685
	goto L1
L402:
	;
	if v696 != 0 {
		v773 = v694
		goto L1
	} else {
		goto L403
	}
L403:
	;
	v698 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v699 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v698, l2)
	mBase = m.M
	v700 = m.ExcPending
	if v700 != 0 {
		goto L3
	} else {
		goto L404
	}
L404:
	;
	if v699 != 0 {
		v773 = v694
		goto L1
	} else {
		goto L405
	}
L405:
	;
	v701 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v702 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v701, l2)
	mBase = m.M
	v703 = m.ExcPending
	if v703 != 0 {
		goto L3
	} else {
		goto L406
	}
L406:
	;
	if v702 != 0 {
		v773 = v694
		goto L1
	} else {
		goto L407
	}
L407:
	;
	v704 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v705 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v704, l2)
	mBase = m.M
	v706 = m.ExcPending
	if v706 != 0 {
		goto L3
	} else {
		goto L408
	}
L408:
	;
	if v705 == int32(0) {
		goto L5
	} else {
		goto L409
	}
L409:
	;
	v773 = v694
	goto L1
L410:
	;
	if v711 != 0 {
		v773 = v709
		goto L1
	} else {
		goto L411
	}
L411:
	;
	v713 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v714 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v713, l2)
	mBase = m.M
	v715 = m.ExcPending
	if v715 != 0 {
		goto L3
	} else {
		goto L412
	}
L412:
	;
	if v714 == int32(0) {
		goto L5
	} else {
		goto L413
	}
L413:
	;
	v773 = v709
	goto L1
L414:
	;
	if v720 != 0 {
		v773 = v718
		goto L1
	} else {
		goto L415
	}
L415:
	;
	v722 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v723 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v722, l2)
	mBase = m.M
	v724 = m.ExcPending
	if v724 != 0 {
		goto L3
	} else {
		goto L416
	}
L416:
	;
	if v723 == int32(0) {
		goto L5
	} else {
		goto L417
	}
L417:
	;
	v773 = v718
	goto L1
L418:
	;
	if v729 != 0 {
		v773 = v727
		goto L1
	} else {
		goto L419
	}
L419:
	;
	v731 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v732 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v731, l2)
	mBase = m.M
	v733 = m.ExcPending
	if v733 != 0 {
		goto L3
	} else {
		goto L420
	}
L420:
	;
	if v732 == int32(0) {
		goto L5
	} else {
		goto L421
	}
L421:
	;
	v773 = v727
	goto L1
L422:
	;
	v740 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v740
	F_errmsg_internal(m, int32(477745), v11)
	mBase = m.M
	v744 = m.ExcPending
	if v744 != 0 {
		goto L3
	} else {
		goto L423
	}
L423:
	;
	F_errfinish(m, int32(486502), int32(4706), int32(296260))
	mBase = m.M
	v749 = m.ExcPending
	if v749 != 0 {
		goto L3
	} else {
		goto L424
	}
L424:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L425:
	;
	if v759 != 0 {
		v773 = v40
		goto L1
	} else {
		goto L426
	}
L426:
	;
	goto L5
}
func F_readIntCols(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
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
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
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
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v12 = F_pg_strtok(m, v8+int32(12))
	mBase = m.M
	if v12 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L10
	} else {
		goto L48
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L10
	} else {
		goto L45
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L10
	} else {
		goto L42
	}
L4:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	switch v13 {
	case 0:
		v113 = int32(0)
		goto L7
	case 1:
		goto L8
	default:
		goto L3
	}
L5:
	;
	goto L6
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L10
	} else {
		goto L39
	}
L7:
	;
	m.G0 = v8 + int32(16)
	return v113
L8:
	;
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v14 != int32(40) {
		goto L3
	} else {
		goto L9
	}
L9:
	;
	v20 = F_palloc(m, l0<<(uint(int32(2))%32))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return int32(0)
L11:
	;
	if int32(0) < l0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v28 = int32(0)
	goto L15
L13:
	;
	goto L14
L14:
	;
	v101 = F_pg_strtok(m, v8+int32(12))
	mBase = m.M
	if v101 == int32(0) {
		goto L2
	} else {
		goto L36
	}
L15:
	;
	v33 = F_pg_strtok(m, v8+int32(12))
	mBase = m.M
	if v33 == int32(0) {
		goto L1
	} else {
		goto L17
	}
L16:
	;
	goto L14
L17:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
	if v36 == int32(41) {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v45 = v33
	goto L20
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20+v28<<(uint(int32(2))%32)))) = v89
	v92 = v28 + int32(1)
	if v92 != l0 {
		v28 = v92
		goto L15
	} else {
		goto L35
	}
L20:
	;
	v50 = v45 + int32(1)
	v51 = int32(*(*int8)(unsafe.Add(mBase, uint32(v45))))
	v52 = F___isspace(m, v51)
	mBase = m.M
	if v52 != 0 {
		v45 = v50
		goto L20
	} else {
		goto L22
	}
L21:
	;
	v53 = int32(1)
	switch v51&int32(255) - int32(43) {
	case 0:
		v59 = v53
		goto L24
	default:
		v61 = v51
		v62 = v45
		v63 = v53
		goto L23
	case 2:
		goto L25
	}
L22:
	;
	goto L21
L23:
	;
	v64 = int32(0)
	v66 = v61 - int32(48)
	if base.Ui32(v66) <= base.Ui32(int32(9)) {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	v60 = int32(*(*int8)(unsafe.Add(mBase, uint32(v50))))
	v61 = v60
	v62 = v50
	v63 = v59
	goto L23
L25:
	;
	v59 = int32(0)
	goto L24
L26:
	;
	v69 = v64
	v70 = v66
	v71 = v62
	goto L29
L27:
	;
	v83 = v64
	goto L28
L28:
	;
	if v63 != 0 {
		goto L32
	} else {
		goto L33
	}
L29:
	;
	v73 = int32(10)
	v75 = v69*v73 - v70
	v76 = int32(*(*int8)(unsafe.Add(mBase, uint32(v71)+1)))
	v80 = v76 - int32(48)
	if base.Ui32(v80) < base.Ui32(v73) {
		v69 = v75
		v70 = v80
		v71 = v71 + int32(1)
		goto L29
	} else {
		goto L31
	}
L30:
	;
	v83 = v75
	goto L28
L31:
	;
	goto L30
L32:
	;
	v89 = int32(0) - v83
	goto L34
L33:
	;
	v89 = v83
	goto L34
L34:
	;
	goto L19
L35:
	;
	goto L16
L36:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	if v104 != int32(1) {
		goto L2
	} else {
		goto L37
	}
L37:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101))))
	if v107 != int32(41) {
		goto L2
	} else {
		goto L38
	}
L38:
	;
	v113 = v20
	goto L7
L39:
	;
	F_errmsg_internal(m, int32(24646), int32(0))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L10
	} else {
		goto L40
	}
L40:
	;
	F_errfinish(m, int32(486455), int32(696), int32(149336))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L10
	} else {
		goto L41
	}
L41:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v12
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v137
	F_errmsg_internal(m, int32(664866), v8)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L10
	} else {
		goto L43
	}
L43:
	;
	F_errfinish(m, int32(486455), int32(696), int32(149336))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L10
	} else {
		goto L44
	}
L44:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L45:
	;
	F_errmsg_internal(m, int32(24646), int32(0))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L10
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(486455), int32(696), int32(149336))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L10
	} else {
		goto L47
	}
L47:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L48:
	;
	F_errmsg_internal(m, int32(24646), int32(0))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L10
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(486455), int32(696), int32(149336))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L10
	} else {
		goto L50
	}
L50:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_readstoplist(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
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
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v203 int32
	_ = v203
	v4 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v4
	if l0 == v4 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v11 + int32(48)
	return
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = int32(0)
	goto L1
L3:
	;
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v17 == int32(0) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v23 = F_get_tsearch_config_filename(m, l0, int32(230617))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	F_tsearch_readline_end(m, v11+int32(4))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L7
	} else {
		goto L64
	}
L6:
	;
	v50 = v4
	v51 = v29
	v53 = v4
	goto L19
L7:
	;
	return
L8:
	;
	v25 = F_tsearch_readline_begin(m, v11+int32(4), v23)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	if v25 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v29 = F_tsearch_readline(m, v11+int32(4))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L7
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L7
	} else {
		goto L15
	}
L13:
	;
	if v29 != 0 {
		goto L6
	} else {
		goto L14
	}
L14:
	;
	v183 = v4
	goto L5
L15:
	;
	F_errcode(m, int32(22))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L7
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v23
	F_errmsg(m, int32(294648), v11)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L7
	} else {
		goto L17
	}
L17:
	;
	F_errfinish(m, int32(485524), int32(85), int32(72532))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L7
	} else {
		goto L18
	}
L18:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L19:
	;
	v55 = v51
	goto L21
L20:
	;
	v183 = v174
	goto L5
L21:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
	switch v63 {
	case 0, 9, 10, 11, 12, 13, 32:
		goto L23
	default:
		goto L24
	}
L22:
	;
	v67 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v55))) = uint8(v67)
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	if v69 == v67 {
		goto L27
	} else {
		goto L28
	}
L23:
	;
	goto L22
L24:
	;
	v64 = F_pg_mblen_cstr(m, v55)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L7
	} else {
		goto L25
	}
L25:
	;
	v55 = v64 + v55
	goto L21
L26:
	;
	v178 = F_tsearch_readline(m, v11+int32(4))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L7
	} else {
		goto L62
	}
L27:
	;
	F_pfree(m, v51)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L7
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v74 < v53 {
		v88 = v50
		v89 = v53
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v174 = v50
	v175 = v53
	goto L26
L31:
	;
	if l2 != 0 {
		goto L39
	} else {
		goto L40
	}
L32:
	;
	if v53 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v80 = F_palloc(m, int32(256))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L7
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v84 = F_repalloc(m, v50, v53<<(uint(int32(3))%32))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L7
	} else {
		goto L37
	}
L36:
	;
	v88 = v80
	v89 = int32(64)
	goto L31
L37:
	;
	v88 = v84
	v89 = v53 << (uint(int32(1)) % 32)
	goto L31
L38:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v169 + int32(1)
	v174 = v88
	v175 = v89
	goto L26
L39:
	;
	if v51&int32(3) == int32(0) {
		v113 = v51
		goto L44
	} else {
		goto L45
	}
L40:
	;
	goto L41
L41:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v88+v163<<(uint(int32(2))%32)))) = v51
	goto L38
L42:
	;
	v148 = m.T0[l2].(func(*base.Module, int32, int32, int32) int32)(m, v51, v146, int32(100))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L7
	} else {
		goto L59
	}
L43:
	;
	v146 = v138 - v51
	goto L42
L44:
	;
	v117 = v113
	goto L53
L45:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	if v97 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v146 = int32(0)
	goto L42
L47:
	;
	goto L48
L48:
	;
	v102 = v51
	goto L49
L49:
	;
	v106 = v102 + int32(1)
	if v106&int32(3) == int32(0) {
		v113 = v106
		goto L44
	} else {
		goto L51
	}
L50:
	;
	v138 = v106
	goto L43
L51:
	;
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106))))
	if v111 != 0 {
		v102 = v106
		goto L49
	} else {
		goto L52
	}
L52:
	;
	goto L50
L53:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
	v126 = int32(-2139062144)
	if (int32(16843008)-v123|v123)&v126 == v126 {
		v117 = v117 + int32(4)
		goto L53
	} else {
		goto L55
	}
L54:
	;
	v132 = v117
	goto L56
L55:
	;
	goto L54
L56:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132))))
	if v136 != 0 {
		v132 = v132 + int32(1)
		goto L56
	} else {
		goto L58
	}
L57:
	;
	v138 = v132
	goto L43
L58:
	;
	goto L57
L59:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v151 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v88+v150<<(uint(v151)%32)))) = v148
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v88+v155<<(uint(v151)%32))))
	if v159 == v51 {
		goto L38
	} else {
		goto L60
	}
L60:
	;
	F_pfree(m, v51)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L7
	} else {
		goto L61
	}
L61:
	;
	goto L38
L62:
	;
	if v178 != 0 {
		v50 = v174
		v51 = v178
		v53 = v175
		goto L19
	} else {
		goto L63
	}
L63:
	;
	goto L20
L64:
	;
	F_pfree(m, v23)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L7
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v183
	if v183 == int32(0) {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v197 <= int32(0) {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	F_pg_qsort(m, v183, v197, int32(4), int32(1185))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L7
	} else {
		goto L68
	}
L68:
	;
	goto L1
}
func F_readtup_index(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v14 = l3 - int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v14
	v16 = F_tuplesort_readtup_alloc(m, l0, v14)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return
	} else {
		v18 = F_LogicalTapeRead(m, l2, v16, v14)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return
		} else {
			if v18 == v14 {
				v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)))
				if v21&int32(1) != 0 {
					v27 = F_LogicalTapeRead(m, l2, v10+int32(12), int32(4))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return
					} else {
						if v27 != int32(4) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
								return
							} else {
								F_errmsg_internal(m, int32(496256), int32(0))
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return
								} else {
									F_errfinish(m, int32(485168), int32(1833), int32(27676))
									mBase = m.M
									v68 = m.ExcPending
									if v68 != 0 {
										return
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l1))) = v16
							v33 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
							v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+52))
							v37 = F_index_getattr_2(m, v16, int32(1), v34, l1+int32(8))
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v37
								m.G0 = v10 + int32(16)
								return
							}
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v16
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
					v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+52))
					v37 = F_index_getattr_2(m, v16, int32(1), v34, l1+int32(8))
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v37
						m.G0 = v10 + int32(16)
						return
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return
				} else {
					F_errmsg_internal(m, int32(496256), int32(0))
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return
					} else {
						F_errfinish(m, int32(485168), int32(1831), int32(27676))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
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
func F_readtup_index_brin(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v12 = l3 - int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v12
	v14 = F_tuplesort_readtup_alloc(m, l0, l3)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v14))) = v12
		v19 = F_LogicalTapeRead(m, l2, v14+int32(4), v12)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return
		} else {
			if v19 == v12 {
				v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)))
				if v22&int32(1) != 0 {
					v28 = F_LogicalTapeRead(m, l2, v9+int32(12), int32(4))
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return
					} else {
						if v28 != int32(4) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return
							} else {
								F_errmsg_internal(m, int32(496256), int32(0))
								mBase = m.M
								v58 = m.ExcPending
								if v58 != 0 {
									return
								} else {
									F_errfinish(m, int32(485168), int32(1909), int32(271074))
									mBase = m.M
									v63 = m.ExcPending
									if v63 != 0 {
										return
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l1))) = v14
							v33 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
							*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v33
							m.G0 = v9 + int32(16)
							return
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v14
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v33
					m.G0 = v9 + int32(16)
					return
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return
				} else {
					F_errmsg_internal(m, int32(496256), int32(0))
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return
					} else {
						F_errfinish(m, int32(485168), int32(1907), int32(271074))
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
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
func F_rebin_segment(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
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
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
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
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
	v9 = int32(15)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
	v14 = int32(32) - base.I32_clz(v12)
	if base.Ui32(v9) <= base.Ui32(v14) {
		v17 = v9
	} else {
		v17 = v14
	}
	if v12 != 0 {
		v19 = v17
	} else {
		v19 = int32(0)
	}
	if v8 == v19 {
		return
	} else {
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
		if v21 != int32(-1) {
			v24 = F_get_segment_by_index(m, l0, v21)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return
			} else {
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
				v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+16))
				*(*int32)(unsafe.Add(mBase, uint32(v26)+16)) = v28
				v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+16))
				if v37 != int32(-1) {
					v40 = F_get_segment_by_index(m, l0, v37)
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return
					} else {
						v42 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
						v43 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
						v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v42)+12)) = v44
						v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
						v47 = v46
						v48 = int32(-1)
						*(*int32)(unsafe.Add(mBase, uint32(v47)+12)) = v48
						v50 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
						v52 = v19 << (uint(int32(2)) % 32)
						v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v55 = *(*int32)(unsafe.Add(mBase, uint32(v52+v53)+160))
						*(*int32)(unsafe.Add(mBase, uint32(v50)+16)) = v55
						v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v57)+20)) = v19
						v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v65 = base.I32_div_s(l1-l0-int32(8), int32(20))
						*(*int32)(unsafe.Add(mBase, uint32(v59+v52)+160)) = v65
						v67 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
						v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+16))
						if v68 == v48 {
							return
						} else {
							v71 = F_get_segment_by_index(m, l0, v68)
							mBase = m.M
							v72 = m.ExcPending
							if v72 != 0 {
								return
							} else {
								v73 = *(*int32)(unsafe.Add(mBase, uint32(v71)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v73)+12)) = v65
								return
							}
						}
					}
				} else {
					v47 = v36
					v48 = int32(-1)
					*(*int32)(unsafe.Add(mBase, uint32(v47)+12)) = v48
					v50 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					v52 = v19 << (uint(int32(2)) % 32)
					v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v55 = *(*int32)(unsafe.Add(mBase, uint32(v52+v53)+160))
					*(*int32)(unsafe.Add(mBase, uint32(v50)+16)) = v55
					v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v57)+20)) = v19
					v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v65 = base.I32_div_s(l1-l0-int32(8), int32(20))
					*(*int32)(unsafe.Add(mBase, uint32(v59+v52)+160)) = v65
					v67 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+16))
					if v68 == v48 {
						return
					} else {
						v71 = F_get_segment_by_index(m, l0, v68)
						mBase = m.M
						v72 = m.ExcPending
						if v72 != 0 {
							return
						} else {
							v73 = *(*int32)(unsafe.Add(mBase, uint32(v71)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v73)+12)) = v65
							return
						}
					}
				}
			}
		} else {
			v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v34 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
			*(*int32)(unsafe.Add(mBase, uint32(v30+v8<<(uint(int32(2))%32))+160)) = v34
			v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+16))
			if v37 != int32(-1) {
				v40 = F_get_segment_by_index(m, l0, v37)
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return
				} else {
					v42 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
					v43 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v42)+12)) = v44
					v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					v47 = v46
					v48 = int32(-1)
					*(*int32)(unsafe.Add(mBase, uint32(v47)+12)) = v48
					v50 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					v52 = v19 << (uint(int32(2)) % 32)
					v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v55 = *(*int32)(unsafe.Add(mBase, uint32(v52+v53)+160))
					*(*int32)(unsafe.Add(mBase, uint32(v50)+16)) = v55
					v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v57)+20)) = v19
					v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v65 = base.I32_div_s(l1-l0-int32(8), int32(20))
					*(*int32)(unsafe.Add(mBase, uint32(v59+v52)+160)) = v65
					v67 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+16))
					if v68 == v48 {
						return
					} else {
						v71 = F_get_segment_by_index(m, l0, v68)
						mBase = m.M
						v72 = m.ExcPending
						if v72 != 0 {
							return
						} else {
							v73 = *(*int32)(unsafe.Add(mBase, uint32(v71)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v73)+12)) = v65
							return
						}
					}
				}
			} else {
				v47 = v36
				v48 = int32(-1)
				*(*int32)(unsafe.Add(mBase, uint32(v47)+12)) = v48
				v50 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				v52 = v19 << (uint(int32(2)) % 32)
				v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v55 = *(*int32)(unsafe.Add(mBase, uint32(v52+v53)+160))
				*(*int32)(unsafe.Add(mBase, uint32(v50)+16)) = v55
				v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v57)+20)) = v19
				v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v65 = base.I32_div_s(l1-l0-int32(8), int32(20))
				*(*int32)(unsafe.Add(mBase, uint32(v59+v52)+160)) = v65
				v67 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+16))
				if v68 == v48 {
					return
				} else {
					v71 = F_get_segment_by_index(m, l0, v68)
					mBase = m.M
					v72 = m.ExcPending
					if v72 != 0 {
						return
					} else {
						v73 = *(*int32)(unsafe.Add(mBase, uint32(v71)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v73)+12)) = v65
						return
					}
				}
			}
		}
	}
}
func F_record_eq(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
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
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int64
	_ = v92
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v219 int32
	_ = v219
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v240 int32
	_ = v240
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v301 int32
	_ = v301
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v389 int32
	_ = v389
	var v394 int32
	_ = v394
	var v398 int32
	_ = v398
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
	var v413 int32
	_ = v413
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v425 int32
	_ = v425
	var v430 int32
	_ = v430
	var v435 int32
	_ = v435
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v484 int32
	_ = v484
	v26 = m.G0
	v28 = v26 - int32(112)
	m.G0 = v28
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v31 = F_pg_detoast_datum(m, v30)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v36 = F_pg_detoast_datum(m, v35)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	F_check_stack_depth(m)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	v42 = F_lookup_rowtype_tupdesc(m, v40, v41)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v36)+8))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	v47 = F_lookup_rowtype_tupdesc(m, v45, v46)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+108)) = v31
	v52 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+104)) = v52
	*(*uint16)(unsafe.Add(mBase, uint32(v28)+100)) = uint16(v52)
	v56 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+96)) = v56
	v58 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+92)) = int32(base.Ui32(v50) >> (uint(v58) % 32))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+88)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v28)+84)) = v52
	*(*uint16)(unsafe.Add(mBase, uint32(v28)+80)) = uint16(v52)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+76)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v28)+72)) = int32(base.Ui32(v61) >> (uint(v58) % 32))
	if v49 < v44 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v73 = v44
	goto L9
L8:
	;
	v73 = v49
	goto L9
L9:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+16))
	if v75 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	if v98 != v40 {
		goto L16
	} else {
		goto L17
	}
L11:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v74)+20))
	v86 = F_MemoryContextAlloc(m, v81, v73<<(uint(int32(2))%32)+int32(20))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L14
	}
L12:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	if v78 < v73 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v75)+4))
	v97 = v75
	v98 = v80
	goto L10
L14:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v88)+16)) = v86
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)+16))
	v92 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v91)+4)) = v92
	*(*int32)(unsafe.Add(mBase, uint32(v91))) = v73
	*(*int64)(unsafe.Add(mBase, uint32(v91)+12)) = v92
	v97 = v91
	v98 = int32(0)
	goto L10
L15:
	;
	v149 = F_palloc(m, v44<<(uint(int32(2))%32))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L31
	}
L16:
	;
	v107 = v73 << (uint(int32(2)) % 32)
	v109 = v97 + int32(20)
	if v109&int32(3) != 0 {
		goto L22
	} else {
		goto L23
	}
L17:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v97)+8))
	if v100 != v41 {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v97)+12))
	if v102 != v45 {
		goto L16
	} else {
		goto L19
	}
L19:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v97)+16))
	if v104 == v46 {
		goto L15
	} else {
		goto L20
	}
L20:
	;
	goto L16
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v97)+16)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v97)+12)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v97)+8)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v97)+4)) = v40
	goto L15
L22:
	;
	v135 = F__emscripten_memset_bulkmem(m, v109, base.I32_extend8_s(int32(0)), v107)
	mBase = m.M
	goto L30
L23:
	;
	if base.Ui32(int32(1024)) < base.Ui32(v107) {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	if base.Ui32(v107+v109) <= base.Ui32(v109) {
		goto L21
	} else {
		goto L25
	}
L25:
	;
	v119 = v97 + v107 + int32(20)
	v121 = v97 + int32(24)
	if base.Ui32(v121) < base.Ui32(v119) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v123 = v119
	goto L28
L27:
	;
	v123 = v121
	goto L28
L28:
	;
	v132 = F__emscripten_memset_bulkmem(m, v109, base.I32_extend8_s(int32(0)), (v123-v97-int32(21))&int32(-4)+int32(4))
	mBase = m.M
	goto L29
L29:
	;
	goto L21
L30:
	;
	goto L21
L31:
	;
	v151 = F_palloc(m, v44)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	F_heap_deform_tuple(m, v28+int32(92), v42, v149, v151)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v159 = F_palloc(m, v49<<(uint(int32(2))%32))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v161 = F_palloc(m, v49)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	F_heap_deform_tuple(m, v28+int32(72), v47, v159, v161)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v165 = int32(0)
	v167 = base.B2i32(v165 < v44)
	if v165 < v44 {
		goto L43
	} else {
		goto L44
	}
L37:
	;
	F_pfree(m, v149)
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L1
	} else {
		goto L98
	}
L38:
	;
	v435 = int32(0)
	goto L37
L39:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L1
	} else {
		goto L94
	}
L40:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L1
	} else {
		goto L89
	}
L41:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L1
	} else {
		goto L83
	}
L42:
	;
	if v340 != v44 {
		goto L39
	} else {
		goto L81
	}
L43:
	;
	v171 = int32(0)
	v173 = int32(20)
	v177 = int32(111)
	v186 = v171
	v188 = v165
	v189 = base.B2i32(v171 < v49)
	v190 = v167
	v192 = v171
	goto L46
L44:
	;
	if int32(0) < v49 {
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v340 = int32(0)
	v342 = v165
	goto L42
L46:
	;
	v211 = v190 & int32(1)
	if v211 != 0 {
		goto L50
	} else {
		goto L51
	}
L47:
	;
	v340 = v328
	v342 = v329
	goto L42
L48:
	;
	v336 = base.B2i32(v329 < v49)
	v337 = base.B2i32(v328 < v44)
	if v328 < v44 {
		v186 = v328
		v188 = v329
		v189 = v336
		v190 = v337
		v192 = v332
		goto L46
	} else {
		goto L79
	}
L49:
	;
	v328 = v186 + int32(1)
	v329 = v319
	v332 = v322
	goto L48
L50:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42+v177+v212<<(uint(int32(4))%32)+v186*int32(100)))))
	if v219 == int32(1) {
		v319 = v188
		v322 = v192
		goto L49
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	if v189&int32(1) == int32(0) {
		v340 = v186
		v342 = v188
		goto L42
	} else {
		goto L54
	}
L53:
	;
	goto L52
L54:
	;
	v227 = v188 * int32(100)
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	v230 = v228 << (uint(int32(4)) % 32)
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v227+(v47+v177+v230)))))
	if v233 == int32(1) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v328 = v186
	v329 = v188 + int32(1)
	v332 = v192
	goto L48
L56:
	;
	goto L57
L57:
	;
	if v211 == int32(0) {
		v340 = v186
		v342 = v188
		goto L42
	} else {
		goto L58
	}
L58:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	v246 = v42 + v173 + v240<<(uint(int32(4))%32) + v186*int32(100)
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v246)+68))
	v249 = v230 + (v47 + v173) + v227
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v249)+68))
	if v247 != v250 {
		goto L41
	} else {
		goto L59
	}
L59:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v249)+96))
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v246)+96))
	v256 = v97 + v173 + v192<<(uint(int32(2))%32)
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v256)))
	if v257 != 0 {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188+v161))))
	v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v186+v151))))
	if v271 == int32(1) {
		goto L68
	} else {
		goto L69
	}
L61:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v257)))
	if v258 == v247 {
		v267 = v257
		goto L60
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	v261 = F_lookup_type_cache(m, v247, int32(32))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L1
	} else {
		goto L65
	}
L64:
	;
	goto L63
L65:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v261)+80))
	if v263 == int32(0) {
		goto L40
	} else {
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v256))) = v261
	v267 = v261
	goto L60
L67:
	;
	v315 = int32(1)
	v319 = v188 + v315
	v322 = v192 + v315
	goto L49
L68:
	;
	if v269&int32(1) != 0 {
		goto L67
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	if v269&int32(1) != 0 {
		goto L38
	} else {
		goto L72
	}
L71:
	;
	goto L38
L72:
	;
	v278 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+52)) = uint8(v278)
	if v253 == v252 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v282 = v253
	goto L75
L74:
	;
	v282 = v278
	goto L75
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+48)) = v282
	*(*int64)(unsafe.Add(mBase, uint32(v28)+40)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+36)) = v267 + int32(76)
	v289 = int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v28)+54)) = uint16(v289)
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v149+v186<<(uint(v289)%32))))
	v295 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+60)) = uint8(v295)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+56)) = v294
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v159+v188<<(uint(v289)%32))))
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+68)) = uint8(v295)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+64)) = v301
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v267)+76))
	v308 = m.T0[v307].(func(*base.Module, int32) int32)(m, v28+int32(36))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+52)))
	if v310 != 0 {
		goto L38
	} else {
		goto L77
	}
L77:
	;
	if v308 == int32(0) {
		goto L38
	} else {
		goto L78
	}
L78:
	;
	goto L67
L79:
	;
	if v329 < v49 {
		v186 = v328
		v188 = v329
		v189 = v336
		v190 = v337
		v192 = v332
		goto L46
	} else {
		goto L80
	}
L80:
	;
	goto L47
L81:
	;
	if v342 != v49 {
		goto L39
	} else {
		goto L82
	}
L82:
	;
	v435 = int32(1)
	goto L37
L83:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v246)+68))
	v375 = F_format_type_be(m, v374)
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v249)+68))
	v378 = F_format_type_be(m, v377)
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+24)) = v192 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+20)) = v378
	*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = v375
	F_errmsg(m, int32(465268), v28+int32(16))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	F_errfinish(m, int32(485676), int32(1198), int32(227788))
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L89:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v261)))
	v403 = F_format_type_be(m, v402)
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28))) = v403
	F_errmsg(m, int32(186468), v28)
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	F_errfinish(m, int32(485676), int32(1221), int32(227788))
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L94:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	F_errmsg(m, int32(145667), int32(0))
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	F_errfinish(m, int32(485676), int32(1265), int32(227788))
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L98:
	;
	F_pfree(m, v151)
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	F_pfree(m, v159)
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	F_pfree(m, v161)
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
	if int32(0) <= v467 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	F_DecrTupleDescRefCount(m, v42)
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L1
	} else {
		goto L105
	}
L103:
	;
	goto L104
L104:
	;
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v47)+12))
	if int32(0) <= v472 {
		goto L106
	} else {
		goto L107
	}
L105:
	;
	goto L104
L106:
	;
	F_DecrTupleDescRefCount(m, v47)
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L1
	} else {
		goto L109
	}
L107:
	;
	goto L108
L108:
	;
	v477 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v477 != v31 {
		goto L110
	} else {
		goto L111
	}
L109:
	;
	goto L108
L110:
	;
	F_pfree(m, v31)
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L1
	} else {
		goto L113
	}
L111:
	;
	goto L112
L112:
	;
	v481 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v481 != v36 {
		goto L114
	} else {
		goto L115
	}
L113:
	;
	goto L112
L114:
	;
	F_pfree(m, v36)
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L1
	} else {
		goto L117
	}
L115:
	;
	goto L116
L116:
	;
	m.G0 = v28 + int32(112)
	return v435
L117:
	;
	goto L116
}
func F_record_ge(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_record_cmp(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return int32(base.Ui32(v2^int32(-1)) >> (uint(int32(31)) % 32))
	}
}
func F_record_image_lt(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_record_image_cmp(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return int32(base.Ui32(v2) >> (uint(int32(31)) % 32))
	}
}
func F_refnameNamespaceItem(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
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
	var v178 int32
	_ = v178
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v217 int32
	_ = v217
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	v6 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(32)
	m.G0 = v17
	if l4 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(0)
	goto L3
L2:
	;
	goto L3
L3:
	;
	if l1 != 0 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L10
	} else {
		goto L85
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L10
	} else {
		goto L80
	}
L6:
	;
	m.G0 = v17 + int32(32)
	return v217
L7:
	;
	v21 = F_LookupNamespaceNoError(m, l1)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v32 = v6
	goto L9
L9:
	;
	v33 = l0
	v43 = v6
	goto L15
L10:
	;
	return int32(0)
L11:
	;
	if v21 == int32(0) {
		v217 = v6
		goto L6
	} else {
		goto L12
	}
L12:
	;
	v27 = F_get_relname_relid(m, l2, v21)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L10
	} else {
		goto L13
	}
L13:
	;
	if v27 == int32(0) {
		v217 = v6
		goto L6
	} else {
		goto L14
	}
L14:
	;
	v32 = v27
	goto L9
L15:
	;
	if v33 == int32(0) {
		v217 = v6
		goto L6
	} else {
		goto L17
	}
L16:
	;
	if v199 == int32(3) {
		v217 = v6
		goto L6
	} else {
		goto L79
	}
L17:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v33)+28))
	if v32 != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	if v178 != 0 {
		goto L65
	} else {
		goto L66
	}
L19:
	;
	if v49 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L21
L21:
	;
	if v49 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L22:
	;
	v178 = int32(0)
	goto L18
L23:
	;
	goto L24
L24:
	;
	v53 = int32(0)
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	if v55 <= v53 {
		v178 = v53
		goto L18
	} else {
		goto L25
	}
L25:
	;
	v59 = v53
	v64 = v53
	v67 = v55
	goto L26
L26:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v49)+12))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v72+v59<<(uint(int32(2))%32))))
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+20)))
	if v77 != int32(1) {
		v95 = v64
		v96 = v67
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v178 = v95
	goto L18
L28:
	;
	v99 = v59 + int32(1)
	if v99 < v96 {
		v59 = v99
		v64 = v95
		v67 = v96
		goto L26
	} else {
		goto L40
	}
L29:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+22)))
	if v81 == int32(1) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+32)))
	if v84 != int32(1) {
		v95 = v64
		v96 = v67
		goto L28
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v76)+24))
	if v87 != 0 {
		v95 = v64
		v96 = v67
		goto L28
	} else {
		goto L34
	}
L33:
	;
	goto L32
L34:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
	if v88 != 0 {
		v95 = v64
		v96 = v67
		goto L28
	} else {
		goto L35
	}
L35:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
	if v89 != v32 {
		v95 = v64
		v96 = v67
		goto L28
	} else {
		goto L36
	}
L36:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
	if v91 != 0 {
		v95 = v64
		v96 = v67
		goto L28
	} else {
		goto L37
	}
L37:
	;
	if v64 != 0 {
		goto L5
	} else {
		goto L38
	}
L38:
	;
	F_check_lateral_ref_ok(m, v33, v76, l3)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L10
	} else {
		goto L39
	}
L39:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	v95 = v76
	v96 = v94
	goto L28
L40:
	;
	goto L27
L41:
	;
	v178 = int32(0)
	goto L18
L42:
	;
	goto L43
L43:
	;
	v104 = int32(0)
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	if v106 <= v104 {
		v178 = v104
		goto L18
	} else {
		goto L44
	}
L44:
	;
	v110 = v104
	v115 = v104
	v118 = v106
	goto L45
L45:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v49)+12))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v123+v110<<(uint(int32(2))%32))))
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+20)))
	if v128 != int32(1) {
		v167 = v115
		v168 = v118
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v178 = v167
	goto L18
L47:
	;
	v170 = v110 + int32(1)
	if v170 < v168 {
		v110 = v170
		v115 = v167
		v118 = v168
		goto L45
	} else {
		goto L64
	}
L48:
	;
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+22)))
	if v131 == int32(1) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+32)))
	if v134 != int32(1) {
		v167 = v115
		v168 = v118
		goto L47
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v127)))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v137)+4))
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138))))
	if v142 == int32(0) {
		v161 = v141
		v162 = v142
		goto L54
	} else {
		goto L55
	}
L52:
	;
	goto L51
L53:
	;
	if v162-v161 != 0 {
		v167 = v115
		v168 = v118
		goto L47
	} else {
		goto L61
	}
L54:
	;
	goto L53
L55:
	;
	if v141 != v142 {
		v161 = v141
		v162 = v142
		goto L54
	} else {
		goto L56
	}
L56:
	;
	v146 = v138
	v147 = l2
	goto L57
L57:
	;
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147)+1)))
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+1)))
	if v151 == int32(0) {
		v161 = v150
		v162 = v151
		goto L54
	} else {
		goto L59
	}
L58:
	;
	v161 = v150
	v162 = v151
	goto L54
L59:
	;
	v154 = int32(1)
	if v150 == v151 {
		v146 = v146 + v154
		v147 = v147 + v154
		goto L57
	} else {
		goto L60
	}
L60:
	;
	goto L58
L61:
	;
	if v115 != 0 {
		goto L4
	} else {
		goto L62
	}
L62:
	;
	F_check_lateral_ref_ok(m, v33, v127, l3)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L10
	} else {
		goto L63
	}
L63:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	v167 = v127
	v168 = v166
	goto L47
L64:
	;
	goto L46
L65:
	;
	v188 = int32(1)
	goto L67
L66:
	;
	v188 = int32(3)
	goto L67
L67:
	;
	if v178 != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v189 = v178
	goto L70
L69:
	;
	v189 = v43
	goto L70
L70:
	;
	if l4 == int32(0) {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	if v199 == int32(0) {
		v33 = v198
		v43 = v200
		goto L15
	} else {
		goto L78
	}
L72:
	;
	v198 = v33
	v199 = v188
	v200 = v189
	goto L71
L73:
	;
	goto L74
L74:
	;
	if v178 != 0 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v198 = v33
	v199 = v188
	v200 = v189
	goto L71
L76:
	;
	goto L77
L77:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v192 + int32(1)
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v198 = v196
	v199 = int32(0)
	v200 = v43
	goto L71
L78:
	;
	goto L16
L79:
	;
	v217 = v200
	goto L6
L80:
	;
	F_errcode(m, int32(151126148))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L10
	} else {
		goto L81
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v32
	F_errmsg(m, int32(112838), v17+int32(16))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L10
	} else {
		goto L82
	}
L82:
	;
	F_parser_errposition(m, v33, l3)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L10
	} else {
		goto L83
	}
L83:
	;
	F_errfinish(m, int32(487844), int32(275), int32(428128))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L10
	} else {
		goto L84
	}
L84:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L85:
	;
	F_errcode(m, int32(151126148))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L10
	} else {
		goto L86
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = l2
	F_errmsg(m, int32(112926), v17)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L10
	} else {
		goto L87
	}
L87:
	;
	F_parser_errposition(m, v33, l3)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L10
	} else {
		goto L88
	}
L88:
	;
	F_errfinish(m, int32(487844), int32(228), int32(371329))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L10
	} else {
		goto L89
	}
L89:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_regcomp_auth_token(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
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
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	v10 = m.G0
	v12 = v10 - int32(160)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	if v15 == int32(47) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v19 = F_palloc0(m, int32(32))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v209 = int32(0)
	goto L3
L3:
	;
	m.G0 = v12 + int32(160)
	return v209
L4:
	;
	return int32(0)
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v19
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v26 = v24 + int32(1)
	if v26&int32(3) == int32(0) {
		v50 = v26
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v88 = F_palloc(m, v83<<(uint(int32(2))%32)+int32(4))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L4
	} else {
		goto L23
	}
L7:
	;
	v83 = v75 - v26
	goto L6
L8:
	;
	v54 = v50
	goto L17
L9:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	if v34 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v83 = int32(0)
	goto L6
L11:
	;
	goto L12
L12:
	;
	v39 = v26
	goto L13
L13:
	;
	v43 = v39 + int32(1)
	if v43&int32(3) == int32(0) {
		v50 = v43
		goto L8
	} else {
		goto L15
	}
L14:
	;
	v75 = v43
	goto L7
L15:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
	if v48 != 0 {
		v39 = v43
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	v63 = int32(-2139062144)
	if (int32(16843008)-v60|v60)&v63 == v63 {
		v54 = v54 + int32(4)
		goto L17
	} else {
		goto L19
	}
L18:
	;
	v69 = v54
	goto L20
L19:
	;
	goto L18
L20:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69))))
	if v73 != 0 {
		v69 = v69 + int32(1)
		goto L20
	} else {
		goto L22
	}
L21:
	;
	v75 = v69
	goto L7
L22:
	;
	goto L21
L23:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v92 = v90 + int32(1)
	if v92&int32(3) == int32(0) {
		v116 = v92
		goto L26
	} else {
		goto L27
	}
L24:
	;
	v150 = F_pg_mb2wchar_with_len(m, v92, v88, v149)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L4
	} else {
		goto L41
	}
L25:
	;
	v149 = v141 - v92
	goto L24
L26:
	;
	v120 = v116
	goto L35
L27:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
	if v100 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v149 = int32(0)
	goto L24
L29:
	;
	goto L30
L30:
	;
	v105 = v92
	goto L31
L31:
	;
	v109 = v105 + int32(1)
	if v109&int32(3) == int32(0) {
		v116 = v109
		goto L26
	} else {
		goto L33
	}
L32:
	;
	v141 = v109
	goto L25
L33:
	;
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109))))
	if v114 != 0 {
		v105 = v109
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
	v129 = int32(-2139062144)
	if (int32(16843008)-v126|v126)&v129 == v129 {
		v120 = v120 + int32(4)
		goto L35
	} else {
		goto L37
	}
L36:
	;
	v135 = v120
	goto L38
L37:
	;
	goto L36
L38:
	;
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135))))
	if v139 != 0 {
		v135 = v135 + int32(1)
		goto L38
	} else {
		goto L40
	}
L39:
	;
	v141 = v135
	goto L25
L40:
	;
	goto L39
L41:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v155 = F_pg_regcomp(m, v152, v88, v150, int32(3), int32(950))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L4
	} else {
		goto L42
	}
L42:
	;
	if v155 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v160 = F_pg_regerror(m, v155, v12+int32(48))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L4
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	F_pfree(m, v88)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L4
	} else {
		goto L57
	}
L46:
	;
	v163 = F_errstart(m, l4, int32(0))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L4
	} else {
		goto L47
	}
L47:
	;
	if v163 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	F_errcode(m, int32(302252162))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L4
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v195 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v12 + int32(48)
	v203 = F_psprintf(m, int32(202585), v12)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L4
	} else {
		goto L56
	}
L51:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v168 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v12 + int32(48)
	F_errmsg(m, int32(202585), v12+int32(32))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L4
	} else {
		goto L52
	}
L52:
	;
	F_set_errcontext_domain(m, int32(0))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L4
	} else {
		goto L53
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = l2
	F_errcontext_msg(m, int32(690554), v12+int32(16))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L4
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(491559), int32(332), int32(277610))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L4
	} else {
		goto L55
	}
L55:
	;
	goto L50
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v203
	goto L45
L57:
	;
	v209 = v155
	goto L3
}
func F_regdictionaryout(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
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
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
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
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v10 == int32(0) {
		v14 = F_pstrdup(m, int32(644528))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v45 = v14
			m.G0 = v8 + int32(16)
			return v45
		}
	} else {
		v19 = F_SearchSysCache1(m, int32(76), v10)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			if v19 != 0 {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
				v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+22)))
				v23 = v21 + v22
				v26 = F_TSDictionaryIsVisible(m, v10)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					if v26 != 0 {
						v32 = int32(0)
						v33 = F_quote_qualified_identifier(m, v32, v23+int32(4))
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int32(0)
						} else {
							F_ReleaseCatCache(m, v19)
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return int32(0)
							} else {
								v45 = v33
								m.G0 = v8 + int32(16)
								return v45
							}
						}
					} else {
						v29 = *(*int32)(unsafe.Add(mBase, uint32(v23)+68))
						v30 = F_get_namespace_name(m, v29)
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return int32(0)
						} else {
							v32 = v30
							v33 = F_quote_qualified_identifier(m, v32, v23+int32(4))
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return int32(0)
							} else {
								F_ReleaseCatCache(m, v19)
								mBase = m.M
								v36 = m.ExcPending
								if v36 != 0 {
									return int32(0)
								} else {
									v45 = v33
									m.G0 = v8 + int32(16)
									return v45
								}
							}
						}
					}
				}
			} else {
				v38 = F_palloc(m, int32(64))
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = v10
					v43 = F_pg_snprintf(m, v38, int32(64), int32(58775), v8)
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int32(0)
					} else {
						v45 = v38
						m.G0 = v8 + int32(16)
						return v45
					}
				}
			}
		}
	}
}
func F_regexnejoinsel(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_Float8GetDatum(m, float64(0.995))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_regoperatorout(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v2 == int32(0) {
		v6 = F_pstrdup(m, int32(546158))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			return v6
		}
	} else {
		v12 = F_format_operator_extended(m, v2, int32(0))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			return v12
		}
	}
}
func F_regroleout(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
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
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v8 == int32(0) {
		v12 = F_pstrdup(m, int32(644528))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v31 = v12
			m.G0 = v6 + int32(16)
			return v31
		}
	} else {
		v17 = F_GetUserNameFromId(m, v8, int32(1))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			if v17 != 0 {
				v19 = F_quote_identifier(m, v17)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					v21 = F_pstrdup(m, v19)
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return int32(0)
					} else {
						v31 = v21
						m.G0 = v6 + int32(16)
						return v31
					}
				}
			} else {
				v24 = F_palloc(m, int32(64))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v6))) = v8
					v29 = F_pg_snprintf(m, v24, int32(64), int32(58775), v6)
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
					} else {
						v31 = v24
						m.G0 = v6 + int32(16)
						return v31
					}
				}
			}
		}
	}
}
func F_remove_useless_results_recurse(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
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
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
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
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
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
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v304 int32
	_ = v304
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v387 int32
	_ = v387
	var v394 int32
	_ = v394
	var v402 int32
	_ = v402
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v458 int32
	_ = v458
	var v472 int32
	_ = v472
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	v5 = int32(0)
	v12 = m.G0
	v14 = v12 + int32(-64)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v16 - int32(63) {
	case 0:
		v495 = l1
		goto L1
	case 1:
		goto L4
	case 2:
		goto L5
	default:
		goto L3
	}
L1:
	;
	m.G0 = v14 - int32(-64)
	return v495
L2:
	;
	if v327 == int32(0) {
		goto L96
	} else {
		goto L97
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L10
	} else {
		goto L93
	}
L4:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v89 == int32(1) {
		goto L23
	} else {
		goto L24
	}
L5:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v19 == int32(0) {
		v495 = l1
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v28 = v5
	v29 = v19
	v32 = v5
	goto L7
L7:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	if v35 <= v28 {
		v327 = v32
		goto L2
	} else {
		goto L9
	}
L8:
	;
	v327 = v84
	goto L2
L9:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	v40 = v37 + v28<<(uint(int32(2))%32)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	v42 = F_remove_useless_results_recurse(m, l0, v41, l1+int32(8), l3)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return int32(0)
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40))) = v42
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v47 == int32(0) {
		v81 = v28
		v82 = v29
		v84 = v32
		goto L12
	} else {
		goto L13
	}
L12:
	;
	if v82 != 0 {
		v28 = v81 + int32(1)
		v29 = v82
		v32 = v84
		goto L7
	} else {
		goto L22
	}
L13:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	if v50 < int32(2) {
		v81 = v28
		v82 = v29
		v84 = v32
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	if v53 != int32(63) {
		v81 = v28
		v82 = v29
		v84 = v32
		goto L12
	} else {
		goto L15
	}
L15:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	if v56 == int32(0) {
		v81 = v28
		v82 = v29
		v84 = v32
		goto L12
	} else {
		goto L16
	}
L16:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+52))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+12))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v61+v56<<(uint(int32(2))%32)-int32(4))))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+12))
	if v68 != int32(8) {
		v81 = v28
		v82 = v29
		v84 = v32
		goto L12
	} else {
		goto L17
	}
L17:
	;
	v71 = F_find_dependent_phvs_in_jointree(m, l0, l1, v56)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L10
	} else {
		goto L18
	}
L18:
	;
	if v71 != 0 {
		v81 = v28
		v82 = v29
		v84 = v32
		goto L12
	} else {
		goto L19
	}
L19:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v74 = F_list_delete_nth_cell(m, v73, v28)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L10
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v74
	v79 = F_bms_add_member(m, v32, v56)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L10
	} else {
		goto L21
	}
L21:
	;
	v81 = v28 - int32(1)
	v82 = v74
	v84 = v79
	goto L12
L22:
	;
	goto L8
L23:
	;
	v92 = l2
	goto L25
L24:
	;
	v92 = int32(0)
	goto L25
L25:
	;
	v94 = l1 + int32(28)
	if v89 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v95 = v92
	goto L28
L27:
	;
	v95 = v94
	goto L28
L28:
	;
	v96 = F_remove_useless_results_recurse(m, l0, v87, v95, l3)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L10
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v96
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if base.Ui32(v101) < base.Ui32(int32(2)) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v104 = v94
	goto L32
L31:
	;
	v104 = int32(0)
	goto L32
L32:
	;
	v105 = F_remove_useless_results_recurse(m, l0, v99, v104, l3)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L10
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v105
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	switch v108 {
	case 0:
		goto L38
	case 1:
		goto L37
	case 2, 5:
		v495 = l1
		goto L1
	default:
		goto L34
	case 4:
		goto L36
	}
L34:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L10
	} else {
		goto L90
	}
L35:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v495 = v293
	goto L1
L36:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
	if v248 != int32(63) {
		v495 = l1
		goto L1
	} else {
		goto L79
	}
L37:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
	if v196 != int32(63) {
		v495 = l1
		goto L1
	} else {
		goto L66
	}
L38:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	if v110 != int32(63) {
		v153 = v105
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v153)))
	if v155 != int32(63) {
		v495 = l1
		goto L1
	} else {
		goto L55
	}
L40:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v109)+4))
	if v113 == int32(0) {
		v153 = v105
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)+52))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)+12))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v118+v113<<(uint(int32(2))%32)-int32(4))))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v124)+12))
	if v125 != int32(8) {
		v153 = v105
		goto L39
	} else {
		goto L42
	}
L42:
	;
	v128 = F_find_dependent_phvs_in_jointree(m, l0, v105, v113)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L10
	} else {
		goto L43
	}
L43:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v128 != 0 {
		v153 = v130
		goto L39
	} else {
		goto L44
	}
L44:
	;
	F_remove_result_refs(m, l0, v113, v130)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L10
	} else {
		goto L45
	}
L45:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
	if l2 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	if v133 != 0 {
		goto L51
	} else {
		goto L52
	}
L47:
	;
	if v133 == int32(0) {
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v136
	*(*int32)(unsafe.Add(mBase, uint32(v14)+52)) = v136
	v142 = F_list_make1_impl(m, int32(1), v12+int32(-28))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L10
	} else {
		goto L49
	}
L49:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v145 = F_makeFromExpr(m, v142, v144)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L10
	} else {
		goto L50
	}
L50:
	;
	v495 = v145
	goto L1
L51:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v148 = F_list_concat(m, v133, v147)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L10
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v495 = v151
	goto L1
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v148
	goto L53
L55:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v153)+4))
	if v158 == int32(0) {
		v495 = l1
		goto L1
	} else {
		goto L56
	}
L56:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v161)+52))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v162)+12))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v163+v158<<(uint(int32(2))%32)-int32(4))))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v169)+12))
	if v170 != int32(8) {
		v495 = l1
		goto L1
	} else {
		goto L57
	}
L57:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	F_remove_result_refs(m, l0, v158, v173)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L10
	} else {
		goto L58
	}
L58:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if l2 != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	if v176 == int32(0) {
		goto L35
	} else {
		goto L64
	}
L60:
	;
	if v176 == int32(0) {
		goto L59
	} else {
		goto L61
	}
L61:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v179
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v179
	v185 = F_list_make1_impl(m, int32(1), v12+int32(-32))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L10
	} else {
		goto L62
	}
L62:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v188 = F_makeFromExpr(m, v185, v187)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L10
	} else {
		goto L63
	}
L63:
	;
	v495 = v188
	goto L1
L64:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v193 = F_list_concat(m, v176, v192)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L10
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v193
	goto L35
L66:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v105)+4))
	if v199 == int32(0) {
		v495 = l1
		goto L1
	} else {
		goto L67
	}
L67:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v202)+52))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v203)+12))
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v204+v199<<(uint(int32(2))%32)-int32(4))))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v210)+12))
	if v211 != int32(8) {
		v495 = l1
		goto L1
	} else {
		goto L68
	}
L68:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
	if v214 == int32(0) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	F_remove_result_refs(m, l0, v199, v240)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L10
	} else {
		goto L77
	}
L70:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v217)+68))
	if v218 == int32(0) {
		goto L69
	} else {
		goto L71
	}
L71:
	;
	v221 = F_bms_make_singleton(m, v199)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L10
	} else {
		goto L72
	}
L72:
	;
	v223 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+60)) = v223
	*(*int32)(unsafe.Add(mBase, uint32(v14)+56)) = v221
	v226 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v231 = F_query_tree_walker_impl(m, v226, int32(852), v12+int32(-8), v223)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L10
	} else {
		goto L73
	}
L73:
	;
	if v231 != 0 {
		v495 = l1
		goto L1
	} else {
		goto L74
	}
L74:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v237 = F_expression_tree_walker_impl(m, v233, int32(852), v12+int32(-8))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L10
	} else {
		goto L75
	}
L75:
	;
	if v237 != 0 {
		v495 = l1
		goto L1
	} else {
		goto L76
	}
L76:
	;
	goto L69
L77:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v244 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v245 = F_bms_add_member(m, v243, v244)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L10
	} else {
		goto L78
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v245
	goto L35
L79:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v105)+4))
	if v251 == int32(0) {
		v495 = l1
		goto L1
	} else {
		goto L80
	}
L80:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v254)+52))
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v255)+12))
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v256+v251<<(uint(int32(2))%32)-int32(4))))
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v262)+12))
	if v263 != int32(8) {
		v495 = l1
		goto L1
	} else {
		goto L81
	}
L81:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	F_remove_result_refs(m, l0, v251, v266)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L10
	} else {
		goto L82
	}
L82:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if l2 != 0 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	if v269 == int32(0) {
		goto L35
	} else {
		goto L88
	}
L84:
	;
	if v269 == int32(0) {
		goto L83
	} else {
		goto L85
	}
L85:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+40)) = v272
	*(*int32)(unsafe.Add(mBase, uint32(v14)+44)) = v272
	v278 = F_list_make1_impl(m, int32(1), v12+int32(-24))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L10
	} else {
		goto L86
	}
L86:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v281 = F_makeFromExpr(m, v278, v280)
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L10
	} else {
		goto L87
	}
L87:
	;
	v495 = v281
	goto L1
L88:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v286 = F_list_concat(m, v269, v285)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L10
	} else {
		goto L89
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v286
	goto L35
L90:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v298
	F_errmsg_internal(m, int32(476940), v12+int32(-48))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L10
	} else {
		goto L91
	}
L91:
	;
	F_errfinish(m, int32(490818), int32(3924), int32(354249))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L10
	} else {
		goto L92
	}
L92:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L93:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v314
	F_errmsg_internal(m, int32(477745), v14)
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L10
	} else {
		goto L94
	}
L94:
	;
	F_errfinish(m, int32(490818), int32(3930), int32(354249))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L10
	} else {
		goto L95
	}
L95:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L96:
	;
	v472 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v472 == int32(0) {
		v495 = l1
		goto L1
	} else {
		goto L125
	}
L97:
	;
	if v327 == int32(0) {
		goto L100
	} else {
		goto L101
	}
L98:
	;
	if v387 < int32(0) {
		goto L96
	} else {
		goto L109
	}
L99:
	;
	v387 = base.I32_ctz(v373) | v374<<(uint(int32(5))%32)
	goto L98
L100:
	;
	v387 = int32(-2)
	goto L98
L101:
	;
	v340 = base.I32_div_s(int32(0), int32(32))
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v327)+4))
	if v341 <= v340 {
		goto L100
	} else {
		goto L102
	}
L102:
	;
	v344 = v327 + int32(8)
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v344+v340<<(uint(int32(2))%32))))
	v351 = v348 & int32(-1)
	if v351 != 0 {
		v373 = v351
		v374 = v340
		goto L99
	} else {
		goto L103
	}
L103:
	;
	v353 = v340 + int32(1)
	if v353 == v341 {
		goto L100
	} else {
		goto L104
	}
L104:
	;
	v356 = v353
	goto L105
L105:
	;
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v344+v356<<(uint(int32(2))%32))))
	if v363 != 0 {
		v373 = v363
		v374 = v356
		goto L99
	} else {
		goto L107
	}
L106:
	;
	goto L100
L107:
	;
	v365 = v356 + int32(1)
	if v365 != v341 {
		v356 = v365
		goto L105
	} else {
		goto L108
	}
L108:
	;
	goto L106
L109:
	;
	v394 = v387
	goto L110
L110:
	;
	F_remove_result_refs(m, l0, v394, l1)
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L10
	} else {
		goto L112
	}
L111:
	;
	goto L96
L112:
	;
	if v327 == int32(0) {
		goto L115
	} else {
		goto L116
	}
L113:
	;
	if int32(0) <= v458 {
		v394 = v458
		goto L110
	} else {
		goto L124
	}
L114:
	;
	v458 = base.I32_ctz(v444) | v445<<(uint(int32(5))%32)
	goto L113
L115:
	;
	v458 = int32(-2)
	goto L113
L116:
	;
	v409 = v394 + int32(1)
	v411 = base.I32_div_s(v409, int32(32))
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v327)+4))
	if v412 <= v411 {
		goto L115
	} else {
		goto L117
	}
L117:
	;
	v415 = v327 + int32(8)
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v415+v411<<(uint(int32(2))%32))))
	v422 = v419 & (int32(-1) << (uint(v409) % 32))
	if v422 != 0 {
		v444 = v422
		v445 = v411
		goto L114
	} else {
		goto L118
	}
L118:
	;
	v424 = v411 + int32(1)
	if v424 == v412 {
		goto L115
	} else {
		goto L119
	}
L119:
	;
	v427 = v424
	goto L120
L120:
	;
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v415+v427<<(uint(int32(2))%32))))
	if v434 != 0 {
		v444 = v434
		v445 = v427
		goto L114
	} else {
		goto L122
	}
L121:
	;
	goto L115
L122:
	;
	v436 = v427 + int32(1)
	if v436 != v412 {
		v427 = v436
		goto L120
	} else {
		goto L123
	}
L123:
	;
	goto L121
L124:
	;
	goto L111
L125:
	;
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v472)+4))
	if v475 != int32(1) {
		v495 = l1
		goto L1
	} else {
		goto L126
	}
L126:
	;
	v478 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v478)+60))
	if l1 == v479 {
		v495 = l1
		goto L1
	} else {
		goto L127
	}
L127:
	;
	v482 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v482 != 0 {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v483 = l2
	goto L130
L129:
	;
	v483 = int32(1)
	goto L130
L130:
	;
	if v483 == int32(0) {
		v495 = l1
		goto L1
	} else {
		goto L131
	}
L131:
	;
	if v482 != 0 {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v486 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v487 = F_list_concat(m, v482, v486)
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L10
	} else {
		goto L135
	}
L133:
	;
	v491 = v472
	goto L134
L134:
	;
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v491)+12))
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v492)))
	v495 = v493
	goto L1
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v487
	v490 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v491 = v490
	goto L134
}
func F_removeabbrev_index_gin(m *base.Module, l0 int32, l1 int32, l2 int32) {
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	F_errstart_cold(m, int32(21), int32(0))
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		F_errmsg_internal(m, int32(438753), int32(0))
		v11 = m.ExcPending
		if v11 != 0 {
			return
		} else {
			F_errfinish(m, int32(485168), int32(1924), int32(272573))
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		}
	}
}
func F_renametrig_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
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
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
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
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	v9 = m.G0
	v11 = v9 - int32(128)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+22)))
	v17 = v13 + v14 + int32(12)
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if v21 == int32(0) {
		v40 = v20
		v41 = v21
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L13
	} else {
		goto L42
	}
L2:
	;
	if v41-v40 != 0 {
		goto L10
	} else {
		goto L11
	}
L3:
	;
	goto L2
L4:
	;
	if v20 != v21 {
		v40 = v20
		v41 = v21
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v25 = v17
	v26 = l3
	goto L6
L6:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+1)))
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+1)))
	if v30 == int32(0) {
		v40 = v29
		v41 = v30
		goto L3
	} else {
		goto L8
	}
L7:
	;
	v40 = v29
	v41 = v30
	goto L3
L8:
	;
	v33 = int32(1)
	if v29 == v30 {
		v25 = v25 + v33
		v26 = v26 + v33
		goto L6
	} else {
		goto L9
	}
L9:
	;
	goto L7
L10:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	F_ScanKeyInit(m, v11+int32(32), int32(2), int32(3), int32(184), v48)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	m.G0 = v11 + int32(128)
	return
L13:
	;
	return
L14:
	;
	F_ScanKeyInit(m, v11+int32(80), int32(4), int32(3), int32(62), l3)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v64 = F_systable_beginscan(m, l0, int32(2701), int32(1), int32(0), int32(2), v11+int32(32))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	v66 = F_systable_getnext(m, v64)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L13
	} else {
		goto L17
	}
L17:
	;
	if v66 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	F_systable_endscan(m, v64)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L13
	} else {
		goto L19
	}
L19:
	;
	v70 = F_heap_copytuple(m, l2)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L13
	} else {
		goto L21
	}
L20:
	;
	v125 = F_strncpy(m, v76, l3, int32(64))
	mBase = m.M
	v126 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v125)+63)) = uint8(v126)
	goto L35
L21:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v70)+16))
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+22)))
	v74 = v72 + v73
	v76 = v74 + int32(12)
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76))))
	if v80 == int32(0) {
		v99 = v79
		v100 = v80
		goto L23
	} else {
		goto L24
	}
L22:
	;
	if v100-v99 == int32(0) {
		goto L20
	} else {
		goto L30
	}
L23:
	;
	goto L22
L24:
	;
	if v79 != v80 {
		v99 = v79
		v100 = v80
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v84 = v76
	v85 = l4
	goto L26
L26:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+1)))
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+1)))
	if v89 == int32(0) {
		v99 = v88
		v100 = v89
		goto L23
	} else {
		goto L28
	}
L27:
	;
	v99 = v88
	v100 = v89
	goto L23
L28:
	;
	v92 = int32(1)
	if v88 == v89 {
		v84 = v84 + v92
		v85 = v85 + v92
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	v106 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L13
	} else {
		goto L31
	}
L31:
	;
	if v106 == int32(0) {
		goto L20
	} else {
		goto L32
	}
L32:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v76
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v110 + int32(4)
	F_errmsg(m, int32(680682), v11)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L13
	} else {
		goto L33
	}
L33:
	;
	F_errfinish(m, int32(487030), int32(1633), int32(306888))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L13
	} else {
		goto L34
	}
L34:
	;
	goto L20
L35:
	;
	F_CatalogTupleUpdate(m, l0, v70+int32(4), v70)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L13
	} else {
		goto L36
	}
L36:
	;
	v133 = *(*int32)(unsafe.Add(mBase, _consts[275]))
	if v133 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	v136 = int32(0)
	F_RunObjectPostAlterHook(m, int32(2620), v135, v136, v136, v136)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L13
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	F_CacheInvalidateRelcache(m, l1)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L13
	} else {
		goto L41
	}
L40:
	;
	goto L39
L41:
	;
	goto L12
L42:
	;
	F_errcode(m, int32(290948))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L13
	} else {
		goto L43
	}
L43:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v157 + int32(4)
	F_errmsg(m, int32(114486), v11+int32(16))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L13
	} else {
		goto L44
	}
L44:
	;
	F_errfinish(m, int32(487030), int32(1615), int32(306888))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L13
	} else {
		goto L45
	}
L45:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_repalloc(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0-int32(8))))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v6&int32(15)*int32(36))+uint32(_consts[1237])))
	v14 = m.T0[v13].(func(*base.Module, int32, int32, int32) int32)(m, l0, l1, int32(0))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		return v14
	}
}
func F_repeat_2(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v51 int64
	_ = v51
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_pg_detoast_datum_packed(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v14 = int32(0)
	if v14 < v13 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L39
	}
L4:
	;
	v17 = v13
	goto L6
L5:
	;
	v17 = v14
	goto L6
L6:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
	if v19 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v51 = base.I64_extend_i32_s(v17) * base.I64_extend_i32_s(v49)
	v55 = base.I32_wrap_i64(v51)
	if base.I32_wrap_i64(int64(base.Ui64(v51)>>(uint(int64(32))%64))) != v55>>(uint(int32(31))%32) {
		goto L3
	} else {
		goto L18
	}
L8:
	;
	v22 = int32(4)
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+1)))
	if v24&int32(254) == int32(2) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v37 = int32(1)
	if v19&v37 != 0 {
		v49 = int32(base.Ui32(v19)>>(uint(v37)%32)) - v37
		goto L7
	} else {
		goto L17
	}
L11:
	;
	v33 = v22
	goto L13
L12:
	;
	v33 = base.B2i32(v24 == int32(18)) << (uint(v22) % 32)
	goto L13
L13:
	;
	if v24 == int32(1) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v36 = v22
	goto L16
L15:
	;
	v36 = v33
	goto L16
L16:
	;
	v49 = v36
	goto L7
L17:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v49 = int32(base.Ui32(v43)>>(uint(int32(2))%32)) - int32(4)
	goto L7
L18:
	;
	v60 = v55 + int32(4)
	if v60 < v55 {
		goto L3
	} else {
		goto L19
	}
L19:
	;
	if base.Ui32(int32(1073741824)) <= base.Ui32(v60) {
		goto L3
	} else {
		goto L20
	}
L20:
	;
	v64 = F_palloc(m, v60)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v64))) = v60 << (uint(int32(2)) % 32)
	if int32(0) < v13 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v71 = int32(1)
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
	if v73&v71 != 0 {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L24
L24:
	;
	return v64
L25:
	;
	v76 = v71
	goto L27
L26:
	;
	v76 = int32(4)
	goto L27
L27:
	;
	v82 = v64 + int32(4)
	v83 = int32(0)
	goto L28
L28:
	;
	if v49 != 0 {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	goto L24
L30:
	;
	v91 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v91 != 0 {
		goto L34
	} else {
		goto L35
	}
L31:
	;
	v88 = F__emscripten_memcpy_bulkmem(m, v82, v9+v76, v49)
	mBase = m.M
	v89 = v88
	goto L33
L32:
	;
	v89 = v82
	goto L33
L33:
	;
	goto L30
L34:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v96 = v83 + int32(1)
	if v96 != v13 {
		v82 = v49 + v89
		v83 = v96
		goto L28
	} else {
		goto L38
	}
L37:
	;
	goto L36
L38:
	;
	goto L29
L39:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	F_errmsg(m, int32(393985), int32(0))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(485060), int32(1167), int32(110428))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_report_invalid_encoding(m *base.Module, l0 int32, l1 int32, l2 int32) {
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	v4 = F_pg_encoding_mblen_or_incomplete(m, l0, l1, l2)
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		F_report_invalid_encoding_int(m, l0, l1, v4, l2)
		v7 = m.ExcPending
		if v7 != 0 {
			return
		} else {
			base.Wasm_trap_unreachable()
			for {
			}
		}
	}
}
func F_report_invalid_record(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = l2
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1252))
	v12 = F_pg_vsnprintf(m, v10, int32(1000), l1, l2)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		v14 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+1256)) = uint8(v14)
		m.G0 = v7 + int32(16)
		return
	}
}
func F_report_newlocale_failure(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	v9 = *(*int32)(unsafe.Add(mBase, _consts[155]))
	if v9 == int32(0) {
		v13 = int32(44)
		*(*int32)(unsafe.Add(mBase, _consts[155])) = v13
		v16 = v13
	} else {
		v16 = v9
	}
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return
	} else {
		F_errcode(m, int32(50856066))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = l0
			F_errmsg(m, int32(294938), v6+int32(16))
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return
			} else {
				if v16 == int32(44) {
					*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
					F_errdetail(m, int32(641445), v6)
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return
					} else {
						F_errfinish(m, int32(491466), int32(829), int32(357133))
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				} else {
					F_errfinish(m, int32(491466), int32(829), int32(357133))
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
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
func F_report_untranslatable_char(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
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
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	v9 = m.G0
	v11 = v9 - int32(80)
	m.G0 = v11
	v13 = F_pg_encoding_mblen_or_incomplete(m, l0, l2, l3)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if v13 < l3 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v16 = v13
	goto L5
L4:
	;
	v16 = l3
	goto L5
L5:
	;
	if int32(0) < v16 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v19 = int32(8)
	if v19 <= v16 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L20
	}
L9:
	;
	v22 = v19
	goto L11
L10:
	;
	v22 = v16
	goto L11
L11:
	;
	v31 = v11 + int32(32)
	v33 = int32(0)
	goto L12
L12:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v33))))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v37
	v42 = F_pg_sprintf(m, v31, int32(29221), v11+int32(16))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L14
	}
L13:
	;
	goto L8
L14:
	;
	v44 = v42 + v31
	if v33 < v22-int32(1) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v48 = F_pg_sprintf(m, v44, int32(720559), int32(0))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L18
	}
L16:
	;
	v51 = v44
	goto L17
L17:
	;
	v53 = v33 + int32(1)
	if v53 != v22 {
		v31 = v51
		v33 = v53
		goto L12
	} else {
		goto L19
	}
L18:
	;
	v51 = v48 + v44
	goto L17
L19:
	;
	goto L13
L20:
	;
	F_errcode(m, int32(84017282))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v70 = int32(3)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l1<<(uint(v70)%32))+uint32(_consts[359])))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v74
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(v70)%32))+uint32(_consts[359])))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v11 + int32(32)
	F_errmsg(m, int32(687230), v11)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(485501), int32(1901), int32(226770))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_rescanLatestTimeLine(m *base.Module, l0 int32, l1 int64) int32 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v3 int32
	_ = v3
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
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
	var v42 int32
	_ = v42
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int64
	_ = v59
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v72 int64
	_ = v72
	var v78 int32
	_ = v78
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v139 int32
	_ = v139
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	v2 = l1
	v3 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(48)
	m.G0 = v14
	v17 = *(*int32)(unsafe.Add(mBase, _consts[246]))
	v18 = F_findNewestTimeLine(m, v17)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v14 + int32(48)
	return v154
L2:
	;
	return int32(0)
L3:
	;
	v23 = *(*int32)(unsafe.Add(mBase, _consts[246]))
	if v18 == v23 {
		v154 = v3
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v25 = F_readTimeLineHistory(m, v18)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L2
	} else {
		goto L8
	}
L5:
	;
	F_errfinish(m, int32(483944), v139, int32(367620))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L2
	} else {
		goto L32
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, _consts[246])) = v18
	v109 = *(*int32)(unsafe.Add(mBase, _consts[247]))
	F_list_free_deep(m, v109)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L2
	} else {
		goto L27
	}
L7:
	;
	v92 = int32(0)
	v95 = F_errstart(m, int32(15), v92)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L2
	} else {
		goto L24
	}
L8:
	;
	if v25 == int32(0) {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	if v29 <= int32(0) {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v32 = int32(0)
	if v32 < v29 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v35 = v29
	goto L13
L12:
	;
	v35 = v32
	goto L13
L13:
	;
	v37 = *(*int32)(unsafe.Add(mBase, _consts[246]))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	v42 = v3
	goto L14
L14:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v38+v42<<(uint(int32(2))%32))))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	if v37 != v54 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v59 = *(*int64)(unsafe.Add(mBase, uint32(v53)+16))
	if base.Ui64(v2) <= base.Ui64(v59) {
		goto L6
	} else {
		goto L20
	}
L16:
	;
	v57 = v42 + int32(1)
	if v35 != v57 {
		v42 = v57
		goto L14
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	goto L15
L19:
	;
	goto L7
L20:
	;
	v61 = int32(0)
	v64 = F_errstart(m, int32(15), v61)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L2
	} else {
		goto L21
	}
L21:
	;
	if v64 == int32(0) {
		v154 = v61
		goto L1
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v18
	*(*uint32)(unsafe.Add(mBase, uint32(v14)+28)) = uint32(v2)
	v72 = int64(base.Ui64(v2) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v14)+24)) = uint32(v72)
	F_errmsg(m, int32(502332), v14+int32(16))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L2
	} else {
		goto L23
	}
L23:
	;
	v139 = int32(4205)
	v146 = int32(0)
	goto L5
L24:
	;
	if v95 == int32(0) {
		v154 = v92
		goto L1
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v18
	F_errmsg(m, int32(50665), v14)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L2
	} else {
		goto L26
	}
L26:
	;
	v139 = int32(4190)
	v146 = int32(0)
	goto L5
L27:
	;
	*(*int32)(unsafe.Add(mBase, _consts[247])) = v25
	v114 = int32(1)
	F_restoreTimeLineHistoryFiles(m, v17+v114, v18)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L2
	} else {
		goto L28
	}
L28:
	;
	v121 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L2
	} else {
		goto L29
	}
L29:
	;
	if v121 == int32(0) {
		v154 = v114
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v126 = *(*int32)(unsafe.Add(mBase, _consts[246]))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v126
	F_errmsg(m, int32(41834), v14+int32(32))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L2
	} else {
		goto L31
	}
L31:
	;
	v139 = int32(4222)
	v146 = int32(1)
	goto L5
L32:
	;
	v154 = v146
	goto L1
}
func F_restriction_is_securely_promotable(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+208))
	if base.Ui32(v6) < base.Ui32(v5) {
		v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)))
		v9 = v8
	} else {
		v9 = int32(1)
	}
	return v9 & int32(1)
}
func F_restriction_selectivity(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) float64 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v44 int32
	_ = v44
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v79 float64
	_ = v79
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v102 float64
	_ = v102
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = F_get_oprrest(m, l1)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return float64(0)
	} else {
		if v14 == int32(0) {
			v102 = float64(0.5)
			m.G0 = v12 + int32(16)
			return v102
		} else {
			v21 = m.G0
			v23 = v21 - int32(96)
			m.G0 = v23
			v28 = *(*int32)(unsafe.Add(mBase, _consts[0]))
			F_fmgr_info_cxt_security(m, v14, v23+int32(16), v28, int32(0))
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return float64(0)
			} else {
				v32 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v23)+92)) = uint8(v32)
				*(*int32)(unsafe.Add(mBase, uint32(v23)+88)) = l4
				*(*uint8)(unsafe.Add(mBase, uint32(v23)+84)) = uint8(v32)
				*(*int32)(unsafe.Add(mBase, uint32(v23)+80)) = l2
				*(*uint8)(unsafe.Add(mBase, uint32(v23)+76)) = uint8(v32)
				*(*int32)(unsafe.Add(mBase, uint32(v23)+72)) = l1
				*(*uint8)(unsafe.Add(mBase, uint32(v23)+68)) = uint8(v32)
				*(*int32)(unsafe.Add(mBase, uint32(v23)+64)) = l0
				v44 = int32(4)
				*(*uint16)(unsafe.Add(mBase, uint32(v23)+62)) = uint16(v44)
				*(*uint8)(unsafe.Add(mBase, uint32(v23)+60)) = uint8(v32)
				*(*int32)(unsafe.Add(mBase, uint32(v23)+56)) = l3
				*(*int64)(unsafe.Add(mBase, uint32(v23)+48)) = int64(0)
				*(*int32)(unsafe.Add(mBase, uint32(v23)+44)) = v23 + int32(16)
				v56 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
				v57 = m.T0[v56].(func(*base.Module, int32) int32)(m, v23+int32(44))
				mBase = m.M
				v58 = m.ExcPending
				if v58 != 0 {
					return float64(0)
				} else {
					v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+60)))
					if v59 == int32(1) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
							return float64(0)
						} else {
							v66 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
							*(*int32)(unsafe.Add(mBase, uint32(v23))) = v66
							F_errmsg_internal(m, int32(523458), v23)
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return float64(0)
							} else {
								F_errfinish(m, int32(486681), int32(1217), int32(299135))
								mBase = m.M
								v75 = m.ExcPending
								if v75 != 0 {
									return float64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						m.G0 = v23 + int32(96)
						v79 = *(*float64)(unsafe.Add(mBase, uint32(v57)))
						if base.F64_lt(v79, float64(0))|base.F64_gt(v79, float64(1)) == int32(0) {
							v102 = v79
							m.G0 = v12 + int32(16)
							return v102
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v90 = m.ExcPending
							if v90 != 0 {
								return float64(0)
							} else {
								*(*float64)(unsafe.Add(mBase, uint32(v12))) = v79
								F_errmsg_internal(m, int32(334346), v12)
								mBase = m.M
								v94 = m.ExcPending
								if v94 != 0 {
									return float64(0)
								} else {
									F_errfinish(m, int32(485095), int32(2007), int32(9797))
									mBase = m.M
									v99 = m.ExcPending
									if v99 != 0 {
										return float64(0)
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
func F_rollback_prepared_cb_wrapper(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int64
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int64
	_ = v34
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = int32(444064)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v12
	v16 = *(*int64)(unsafe.Add(mBase, uint32(l1)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = int32(992)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = v16
	v20 = int32(4463464)
	v21 = *(*int32)(unsafe.Add(mBase, _consts[84]))
	*(*int32)(unsafe.Add(mBase, _consts[84])) = v10 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v21
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v10 + int32(16)
	v30 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+147)) = uint8(v30)
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+160)) = v32
	v34 = *(*int64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+164)) = uint8(v30)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+152)) = v34
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v12)+72))
	if v38 == int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v44 = m.ExcPending
		if v44 != 0 {
			return
		} else {
			F_errcode(m, int32(325))
			mBase = m.M
			v47 = m.ExcPending
			if v47 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(494988)
				F_errmsg(m, int32(313577), v10)
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return
				} else {
					F_errfinish(m, int32(489070), int32(1077), int32(214858))
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
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
		m.T0[v38].(func(*base.Module, int32, int32, int64, int64))(m, v12, l1, l2, l3)
		mBase = m.M
		v59 = m.ExcPending
		if v59 != 0 {
			return
		} else {
			v61 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
			*(*int32)(unsafe.Add(mBase, _consts[84])) = v61
			m.G0 = v10 + int32(32)
			return
		}
	}
}
func F_round_var(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
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
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l1
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v12 = l1 + v9<<(uint(int32(2))%32)
	if v12+int32(4) < int32(0) {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
		*(*int64)(unsafe.Add(mBase, uint32(l0))) = int64(0)
		return
	} else {
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v23 = l1 & int32(3)
		v27 = base.I32_div_s(v12+int32(7), int32(4))
		v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v28 <= v27 {
			if v23 == int32(0) {
			} else {
				if v27 != v28 {
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0))) = v27
					v41 = int32(1)
					v42 = v27 - v41
					v45 = v21 + v42<<(uint(v41)%32)
					v46 = int32(*(*int16)(unsafe.Add(mBase, uint32(v45))))
					v47 = int32(2)
					v51 = *(*int32)(unsafe.Add(mBase, uint32(v23<<(uint(v47)%32))+uint32(_consts[1116])))
					v52 = base.I32_rem_s(v46, v51)
					v53 = v46 - v52
					*(*uint16)(unsafe.Add(mBase, uint32(v45))) = uint16(v53)
					v56 = base.I32_div_s(v51, v47)
					if v52 < v56 {
						v94 = v42
					} else {
						v59 = v51 + base.I32_extend16_s(v53)
						if int32(9999) < v59 {
							v64 = v59 + int32(55536)
						} else {
							v64 = v59
						}
						*(*uint16)(unsafe.Add(mBase, uint32(v45))) = uint16(v64)
						if v59 < int32(10000) {
							v94 = v42
						} else {
							v68 = v42
							v74 = v68
							for {
								v80 = int32(1)
								v81 = v74 - v80
								v84 = v21 + v81<<(uint(v80)%32)
								v87 = int32(*(*int16)(unsafe.Add(mBase, uint32(v84))))
								v89 = base.B2i32(int32(9998) < v87)
								if int32(9998) < v87 {
									v90 = int32(-9999)
								} else {
									v90 = v80
								}
								v91 = v90 + v87
								*(*uint16)(unsafe.Add(mBase, uint32(v84))) = uint16(v91)
								if int32(9998) < v87 {
									v74 = v81
									continue
								} else {
									break
								}
								break
							}
							v94 = v81
						}
					}
					if int32(0) <= v94 {
					} else {
						v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v102 - int32(2)
						v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v107 = int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(l0))) = v106 + v107
						v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v110 + v107
					}
				}
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = v27
			if v23 != 0 {
				v41 = int32(1)
				v42 = v27 - v41
				v45 = v21 + v42<<(uint(v41)%32)
				v46 = int32(*(*int16)(unsafe.Add(mBase, uint32(v45))))
				v47 = int32(2)
				v51 = *(*int32)(unsafe.Add(mBase, uint32(v23<<(uint(v47)%32))+uint32(_consts[1116])))
				v52 = base.I32_rem_s(v46, v51)
				v53 = v46 - v52
				*(*uint16)(unsafe.Add(mBase, uint32(v45))) = uint16(v53)
				v56 = base.I32_div_s(v51, v47)
				if v52 < v56 {
					v94 = v42
				} else {
					v59 = v51 + base.I32_extend16_s(v53)
					if int32(9999) < v59 {
						v64 = v59 + int32(55536)
					} else {
						v64 = v59
					}
					*(*uint16)(unsafe.Add(mBase, uint32(v45))) = uint16(v64)
					if v59 < int32(10000) {
						v94 = v42
					} else {
						v68 = v42
						v74 = v68
						for {
							v80 = int32(1)
							v81 = v74 - v80
							v84 = v21 + v81<<(uint(v80)%32)
							v87 = int32(*(*int16)(unsafe.Add(mBase, uint32(v84))))
							v89 = base.B2i32(int32(9998) < v87)
							if int32(9998) < v87 {
								v90 = int32(-9999)
							} else {
								v90 = v80
							}
							v91 = v90 + v87
							*(*uint16)(unsafe.Add(mBase, uint32(v84))) = uint16(v91)
							if int32(9998) < v87 {
								v74 = v81
								continue
							} else {
								break
							}
							break
						}
						v94 = v81
					}
				}
			} else {
				v38 = int32(*(*int16)(unsafe.Add(mBase, uint32(v21+v27<<(uint(int32(1))%32)))))
				if v38 <= int32(4999) {
					v94 = v27
				} else {
					v68 = v27
					v74 = v68
					for {
						v80 = int32(1)
						v81 = v74 - v80
						v84 = v21 + v81<<(uint(v80)%32)
						v87 = int32(*(*int16)(unsafe.Add(mBase, uint32(v84))))
						v89 = base.B2i32(int32(9998) < v87)
						if int32(9998) < v87 {
							v90 = int32(-9999)
						} else {
							v90 = v80
						}
						v91 = v90 + v87
						*(*uint16)(unsafe.Add(mBase, uint32(v84))) = uint16(v91)
						if int32(9998) < v87 {
							v74 = v81
							continue
						} else {
							break
						}
						break
					}
					v94 = v81
				}
			}
			if int32(0) <= v94 {
			} else {
				v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v102 - int32(2)
				v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v107 = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(l0))) = v106 + v107
				v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v110 + v107
			}
		}
		return
	}
}
func F_russian_KOI8_R_stem(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v93 int32
	_ = v93
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v152 int32
	_ = v152
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v207 int32
	_ = v207
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v239 int32
	_ = v239
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v255 int32
	_ = v255
	var v264 int32
	_ = v264
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
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
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v404 int32
	_ = v404
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v431 int32
	_ = v431
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v447 int32
	_ = v447
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v468 int32
	_ = v468
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v485 int32
	_ = v485
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v495 int32
	_ = v495
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v521 int32
	_ = v521
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v538 int32
	_ = v538
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v545 int32
	_ = v545
	var v547 int32
	_ = v547
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v562 int32
	_ = v562
	var v565 int32
	_ = v565
	var v567 int32
	_ = v567
	var v571 int32
	_ = v571
	var v575 int32
	_ = v575
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v592 int32
	_ = v592
	var v596 int32
	_ = v596
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v605 int32
	_ = v605
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v615 int32
	_ = v615
	var v617 int32
	_ = v617
	var v620 int32
	_ = v620
	var v622 int32
	_ = v622
	var v626 int32
	_ = v626
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v643 int32
	_ = v643
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v651 int32
	_ = v651
	var v653 int32
	_ = v653
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v659 int32
	_ = v659
	var v663 int32
	_ = v663
	var v669 int32
	_ = v669
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v679 int32
	_ = v679
	var v681 int32
	_ = v681
	var v685 int32
	_ = v685
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v706 int32
	_ = v706
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	goto L2
L1:
	;
	return v706
L2:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v15 < v14 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v7
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v48))) = v15
	*(*int32)(unsafe.Add(mBase, uint32(v48)+4)) = v15
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v61 < v51 {
		goto L21
	} else {
		goto L22
	}
L4:
	;
	v17 = v14
	goto L6
L5:
	;
	v17 = v15
	goto L6
L6:
	;
	v19 = v14
	goto L8
L7:
	;
	goto L3
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v19
	if v19 != v15 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v19
	v36 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v19 + v36
	v41 = F_slice_from_s(m, l0, v36, int32(2154834))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L16
	} else {
		goto L17
	}
L10:
	;
	goto L9
L11:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26+v19))))
	if v28 == int32(163) {
		goto L10
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	if v19 == v17 {
		goto L7
	} else {
		goto L15
	}
L14:
	;
	goto L13
L15:
	;
	v33 = v19 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v33
	v19 = v33
	goto L8
L16:
	;
	return int32(0)
L17:
	;
	if int32(0) <= v41 {
		goto L2
	} else {
		goto L18
	}
L18:
	;
	v706 = v41
	goto L1
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v51
	v282 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v282
	v284 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v284)+4))
	if v282 < v285 {
		v706 = int32(0)
		goto L1
	} else {
		goto L82
	}
L20:
	;
	if v101 < int32(0) {
		goto L19
	} else {
		goto L35
	}
L21:
	;
	v63 = v51
	goto L23
L22:
	;
	v63 = v61
	goto L23
L23:
	;
	v70 = v51
	goto L25
L24:
	;
	v101 = v81
	goto L20
L25:
	;
	if v70 == v63 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v101 = int32(-1)
	goto L20
L28:
	;
	goto L29
L29:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74+v70))))
	if int32(220) < v76 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v93 = v70 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v93
	v70 = v93
	goto L25
L31:
	;
	v78 = v76 - int32(192)
	if v78 < int32(0) {
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v81 = int32(1)
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v78)>>(uint(int32(3))%32)))+uint32(_consts[1316]))))
	if int32(base.Ui32(v85)>>(uint(v78&int32(7))%32))&v81 != 0 {
		goto L24
	} else {
		goto L33
	}
L33:
	;
	goto L30
L35:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v105 = v104 + v101
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v105
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v107)+4)) = v105
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v118 < v117 {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	if v161 < int32(0) {
		goto L19
	} else {
		goto L50
	}
L37:
	;
	v120 = v117
	goto L39
L38:
	;
	v120 = v118
	goto L39
L39:
	;
	v127 = v117
	goto L41
L40:
	;
	v161 = int32(1)
	goto L36
L41:
	;
	if v127 == v120 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v161 = int32(-1)
	goto L36
L44:
	;
	goto L45
L45:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133+v127))))
	if int32(220) < v135 {
		goto L40
	} else {
		goto L46
	}
L46:
	;
	v137 = v135 - int32(192)
	if v137 < int32(0) {
		goto L40
	} else {
		goto L47
	}
L47:
	;
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v137)>>(uint(int32(3))%32)))+uint32(_consts[1316]))))
	if int32(base.Ui32(v143)>>(uint(v137&int32(7))%32))&int32(1) == int32(0) {
		goto L40
	} else {
		goto L48
	}
L48:
	;
	v152 = v127 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v152
	v127 = v152
	goto L41
L50:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v165 = v164 + v161
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v165
	v175 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v175 < v165 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	if v215 < int32(0) {
		goto L19
	} else {
		goto L66
	}
L52:
	;
	v177 = v165
	goto L54
L53:
	;
	v177 = v175
	goto L54
L54:
	;
	v184 = v165
	goto L56
L55:
	;
	v215 = v195
	goto L51
L56:
	;
	if v184 == v177 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v215 = int32(-1)
	goto L51
L59:
	;
	goto L60
L60:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188+v184))))
	if int32(220) < v190 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v207 = v184 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v207
	v184 = v207
	goto L56
L62:
	;
	v192 = v190 - int32(192)
	if v192 < int32(0) {
		goto L61
	} else {
		goto L63
	}
L63:
	;
	v195 = int32(1)
	v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v192)>>(uint(int32(3))%32)))+uint32(_consts[1316]))))
	if int32(base.Ui32(v199)>>(uint(v192&int32(7))%32))&v195 != 0 {
		goto L55
	} else {
		goto L64
	}
L64:
	;
	goto L61
L66:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v219 = v218 + v215
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v219
	v230 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v230 < v219 {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	if v273 < int32(0) {
		goto L19
	} else {
		goto L81
	}
L68:
	;
	v232 = v219
	goto L70
L69:
	;
	v232 = v230
	goto L70
L70:
	;
	v239 = v219
	goto L72
L71:
	;
	v273 = int32(1)
	goto L67
L72:
	;
	if v239 == v232 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v273 = int32(-1)
	goto L67
L75:
	;
	goto L76
L76:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v245+v239))))
	if int32(220) < v247 {
		goto L71
	} else {
		goto L77
	}
L77:
	;
	v249 = v247 - int32(192)
	if v249 < int32(0) {
		goto L71
	} else {
		goto L78
	}
L78:
	;
	v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v249)>>(uint(int32(3))%32)))+uint32(_consts[1316]))))
	if int32(base.Ui32(v255)>>(uint(v249&int32(7))%32))&int32(1) == int32(0) {
		goto L71
	} else {
		goto L79
	}
L79:
	;
	v264 = v239 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v264
	v239 = v264
	goto L72
L81:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v277 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v276))) = v277 + v273
	goto L19
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v282
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v285
	if v282 <= v285 {
		v338 = v285
		goto L84
	} else {
		goto L85
	}
L83:
	;
	v562 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v562
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v562
	v565 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v562 <= v565 {
		v585 = v565
		v586 = v562
		v587 = v562
		goto L166
	} else {
		goto L167
	}
L84:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v339
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v339
	v343 = v339 - int32(1)
	if v343 <= v338 {
		v355 = v339
		goto L102
	} else {
		goto L103
	}
L85:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v290+v282-int32(1)))))
	if v294&int32(224) != int32(192) {
		v338 = v285
		goto L84
	} else {
		goto L86
	}
L86:
	;
	if int32(1)<<(uint(v294)%32)&int32(25166336) == int32(0) {
		v338 = v285
		goto L84
	} else {
		goto L87
	}
L87:
	;
	v307 = F_find_among_b(m, l0, int32(4194576), int32(9))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L16
	} else {
		goto L88
	}
L88:
	;
	if v307 == int32(0) {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v338 = v311
	goto L84
L90:
	;
	goto L91
L91:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v312
	switch v307 - int32(1) {
	case 0:
		goto L93
	case 1:
		goto L92
	default:
		goto L83
	}
L92:
	;
	v332 = F_slice_del(m, l0)
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L16
	} else {
		goto L98
	}
L93:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v312 <= v316 {
		v338 = v316
		goto L84
	} else {
		goto L94
	}
L94:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v318+v312-int32(1)))))
	switch v322 - int32(193) {
	case 0, 16:
		goto L95
	default:
		v338 = v316
		goto L84
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v312 - int32(1)
	v328 = F_slice_del(m, l0)
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L16
	} else {
		goto L96
	}
L96:
	;
	if int32(0) <= v328 {
		goto L83
	} else {
		goto L97
	}
L97:
	;
	v706 = v328
	goto L1
L98:
	;
	if int32(0) <= v332 {
		goto L83
	} else {
		goto L99
	}
L99:
	;
	v706 = v332
	goto L1
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v365
	v369 = v365 - int32(1)
	v370 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v369 <= v370 {
		v391 = v366
		goto L112
	} else {
		goto L113
	}
L101:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v357
	v359 = F_slice_del(m, l0)
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L16
	} else {
		goto L107
	}
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v355
	v365 = v355
	v366 = v355
	goto L100
L103:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v347 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v345+v343))))
	switch v347 - int32(209) {
	case 0, 7:
		goto L104
	default:
		v355 = v339
		goto L102
	}
L104:
	;
	v352 = F_find_among_b(m, l0, int32(4194768), int32(2))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L16
	} else {
		goto L105
	}
L105:
	;
	if v352 != 0 {
		goto L101
	} else {
		goto L106
	}
L106:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v355 = v354
	goto L102
L107:
	;
	if v359 < int32(0) {
		v706 = v359
		goto L1
	} else {
		goto L108
	}
L108:
	;
	v363 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v364 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v365 = v363
	v366 = v364
	goto L100
L109:
	;
	if v556 != 0 {
		v706 = v555
		goto L1
	} else {
		goto L165
	}
L110:
	;
	v509 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v510 = v509 - v392
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v510
	v512 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v510
	v515 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v510 <= v515 {
		v547 = v512
		goto L151
	} else {
		goto L152
	}
L111:
	;
	v454 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v454
	v456 = F_slice_del(m, l0)
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L16
	} else {
		goto L136
	}
L112:
	;
	v392 = v366 - v365
	v393 = v391 - v392
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v393
	v395 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v393
	v398 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v393 <= v398 {
		v447 = v395
		goto L118
	} else {
		goto L119
	}
L113:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v374 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v372+v369))))
	if v374&int32(224) != int32(192) {
		v391 = v366
		goto L112
	} else {
		goto L114
	}
L114:
	;
	if int32(1)<<(uint(v374)%32)&int32(2271009) == int32(0) {
		v391 = v366
		goto L112
	} else {
		goto L115
	}
L115:
	;
	v387 = F_find_among_b(m, l0, int32(4194976), int32(26))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L16
	} else {
		goto L116
	}
L116:
	;
	if v387 != 0 {
		goto L111
	} else {
		goto L117
	}
L117:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v391 = v389
	goto L112
L118:
	;
	v451 = int32(base.Ui32(v447) >> (uint(int32(31)) % 32))
	if v447 != 0 {
		goto L133
	} else {
		goto L134
	}
L119:
	;
	v400 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v404 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v400+v393-int32(1)))))
	if v404&int32(224) != int32(192) {
		v447 = v395
		goto L118
	} else {
		goto L120
	}
L120:
	;
	if int32(1)<<(uint(v404)%32)&int32(51443235) == int32(0) {
		v447 = v395
		goto L118
	} else {
		goto L121
	}
L121:
	;
	v417 = F_find_among_b(m, l0, int32(4195504), int32(46))
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L16
	} else {
		goto L122
	}
L122:
	;
	if v417 == int32(0) {
		v447 = v395
		goto L118
	} else {
		goto L123
	}
L123:
	;
	v421 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v421
	switch v417 - int32(1) {
	case 0:
		goto L126
	case 1:
		goto L125
	default:
		goto L124
	}
L124:
	;
	v447 = int32(1)
	goto L118
L125:
	;
	v441 = F_slice_del(m, l0)
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L16
	} else {
		goto L131
	}
L126:
	;
	v425 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v421 <= v425 {
		v447 = v395
		goto L118
	} else {
		goto L127
	}
L127:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v431 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v427+v421-int32(1)))))
	switch v431 - int32(193) {
	case 0, 16:
		goto L128
	default:
		v447 = v395
		goto L118
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v421 - int32(1)
	v437 = F_slice_del(m, l0)
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L16
	} else {
		goto L129
	}
L129:
	;
	if int32(0) <= v437 {
		goto L124
	} else {
		goto L130
	}
L130:
	;
	v447 = v437
	goto L118
L131:
	;
	if v441 < int32(0) {
		v447 = v441
		goto L118
	} else {
		goto L132
	}
L132:
	;
	goto L124
L133:
	;
	v453 = v451
	goto L135
L134:
	;
	v453 = int32(13)
	goto L135
L135:
	;
	switch v453 {
	case 0:
		goto L83
	default:
		v555 = v447
		v556 = v451
		goto L109
	case 13:
		goto L110
	}
L136:
	;
	if v456 < int32(0) {
		v706 = v456
		goto L1
	} else {
		goto L137
	}
L137:
	;
	v460 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v460
	v462 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v460 <= v462 {
		goto L83
	} else {
		goto L138
	}
L138:
	;
	v464 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v468 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v464+v460-int32(1)))))
	if v468&int32(224) != int32(192) {
		goto L83
	} else {
		goto L139
	}
L139:
	;
	if int32(1)<<(uint(v468)%32)&int32(671113216) == int32(0) {
		goto L83
	} else {
		goto L140
	}
L140:
	;
	v481 = F_find_among_b(m, l0, int32(4194816), int32(8))
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L16
	} else {
		goto L141
	}
L141:
	;
	if v481 == int32(0) {
		goto L83
	} else {
		goto L142
	}
L142:
	;
	v485 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v485
	switch v481 - int32(1) {
	case 0:
		goto L144
	case 1:
		goto L143
	default:
		goto L83
	}
L143:
	;
	v505 = F_slice_del(m, l0)
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L16
	} else {
		goto L149
	}
L144:
	;
	v489 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v485 <= v489 {
		goto L83
	} else {
		goto L145
	}
L145:
	;
	v491 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v495 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v491+v485-int32(1)))))
	switch v495 - int32(193) {
	case 0, 16:
		goto L146
	default:
		goto L83
	}
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v485 - int32(1)
	v501 = F_slice_del(m, l0)
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L16
	} else {
		goto L147
	}
L147:
	;
	if int32(0) <= v501 {
		goto L83
	} else {
		goto L148
	}
L148:
	;
	v706 = v501
	goto L1
L149:
	;
	if int32(0) <= v505 {
		goto L83
	} else {
		goto L150
	}
L150:
	;
	v706 = v505
	goto L1
L151:
	;
	if v547 == int32(0) {
		goto L83
	} else {
		goto L161
	}
L152:
	;
	v517 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v521 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v517+v510-int32(1)))))
	if v521&int32(224) != int32(192) {
		v547 = v512
		goto L151
	} else {
		goto L153
	}
L153:
	;
	if int32(1)<<(uint(v521)%32)&int32(60991267) == int32(0) {
		v547 = v512
		goto L151
	} else {
		goto L154
	}
L154:
	;
	v534 = F_find_among_b(m, l0, int32(4196432), int32(36))
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L16
	} else {
		goto L155
	}
L155:
	;
	if v534 == int32(0) {
		v547 = v512
		goto L151
	} else {
		goto L156
	}
L156:
	;
	v538 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v538
	v541 = F_slice_del(m, l0)
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L16
	} else {
		goto L157
	}
L157:
	;
	if int32(0) <= v541 {
		goto L158
	} else {
		goto L159
	}
L158:
	;
	v545 = int32(1)
	goto L160
L159:
	;
	v545 = v541
	goto L160
L160:
	;
	v547 = v545
	goto L151
L161:
	;
	if v547 < int32(0) {
		goto L162
	} else {
		goto L163
	}
L162:
	;
	v552 = v547
	goto L164
L163:
	;
	v552 = v447
	goto L164
L164:
	;
	v555 = v552
	v556 = int32(base.Ui32(v547) >> (uint(int32(31)) % 32))
	goto L109
L165:
	;
	goto L83
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v586
	if v586-int32(2) <= v585 {
		goto L171
	} else {
		goto L172
	}
L167:
	;
	v567 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v571 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v567+v562-int32(1)))))
	if v571 != int32(201) {
		v585 = v565
		v586 = v562
		v587 = v562
		goto L166
	} else {
		goto L168
	}
L168:
	;
	v575 = v562 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v575
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v575
	v578 = F_slice_del(m, l0)
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L16
	} else {
		goto L169
	}
L169:
	;
	if v578 < int32(0) {
		v706 = v578
		goto L1
	} else {
		goto L170
	}
L170:
	;
	v582 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v583 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v584 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v585 = v582
	v586 = v583
	v587 = v584
	goto L166
L171:
	;
	v615 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v617 = v615 + (v586 - v587)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v617
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v617
	v620 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v617 <= v620 {
		goto L179
	} else {
		goto L180
	}
L172:
	;
	v592 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v596 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v592+v586-int32(1)))))
	switch v596 - int32(212) {
	case 0, 4:
		goto L173
	default:
		goto L171
	}
L173:
	;
	v601 = F_find_among_b(m, l0, int32(4197152), int32(2))
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L16
	} else {
		goto L174
	}
L174:
	;
	if v601 == int32(0) {
		goto L171
	} else {
		goto L175
	}
L175:
	;
	v605 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v605
	v607 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v608 = *(*int32)(unsafe.Add(mBase, uint32(v607)))
	if v605 < v608 {
		goto L171
	} else {
		goto L176
	}
L176:
	;
	v610 = F_slice_del(m, l0)
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L16
	} else {
		goto L177
	}
L177:
	;
	if v610 < int32(0) {
		v706 = v610
		goto L1
	} else {
		goto L178
	}
L178:
	;
	goto L171
L179:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v51
	v706 = int32(1)
	goto L1
L180:
	;
	v622 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v626 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v622+v617-int32(1)))))
	if v626&int32(224) != int32(192) {
		goto L179
	} else {
		goto L181
	}
L181:
	;
	if int32(1)<<(uint(v626)%32)&int32(151011360) == int32(0) {
		goto L179
	} else {
		goto L182
	}
L182:
	;
	v639 = F_find_among_b(m, l0, int32(4197200), int32(4))
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		goto L16
	} else {
		goto L183
	}
L183:
	;
	if v639 == int32(0) {
		goto L179
	} else {
		goto L184
	}
L184:
	;
	v643 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v643
	switch v639 - int32(1) {
	case 0:
		goto L187
	case 1:
		goto L186
	case 2:
		goto L185
	default:
		goto L179
	}
L185:
	;
	v695 = F_slice_del(m, l0)
	mBase = m.M
	v696 = m.ExcPending
	if v696 != 0 {
		goto L16
	} else {
		goto L200
	}
L186:
	;
	v679 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v643 <= v679 {
		goto L179
	} else {
		goto L196
	}
L187:
	;
	v647 = F_slice_del(m, l0)
	mBase = m.M
	v648 = m.ExcPending
	if v648 != 0 {
		goto L16
	} else {
		goto L188
	}
L188:
	;
	if v647 < int32(0) {
		v706 = v647
		goto L1
	} else {
		goto L189
	}
L189:
	;
	v651 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v651
	v653 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v651 <= v653 {
		goto L179
	} else {
		goto L190
	}
L190:
	;
	v655 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v656 = v655 + v651
	v659 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v656-int32(1)))))
	if v659 != int32(206) {
		goto L179
	} else {
		goto L191
	}
L191:
	;
	v663 = v651 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v663
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v663
	if v663 <= v653 {
		goto L179
	} else {
		goto L192
	}
L192:
	;
	v669 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v656-int32(2)))))
	if v669 != int32(206) {
		goto L179
	} else {
		goto L193
	}
L193:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v651 - int32(2)
	v675 = F_slice_del(m, l0)
	mBase = m.M
	v676 = m.ExcPending
	if v676 != 0 {
		goto L16
	} else {
		goto L194
	}
L194:
	;
	if int32(0) <= v675 {
		goto L179
	} else {
		goto L195
	}
L195:
	;
	v706 = v675
	goto L1
L196:
	;
	v681 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v685 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v681+v643-int32(1)))))
	if v685 != int32(206) {
		goto L179
	} else {
		goto L197
	}
L197:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v643 - int32(1)
	v691 = F_slice_del(m, l0)
	mBase = m.M
	v692 = m.ExcPending
	if v692 != 0 {
		goto L16
	} else {
		goto L198
	}
L198:
	;
	if int32(0) <= v691 {
		goto L179
	} else {
		goto L199
	}
L199:
	;
	v706 = v691
	goto L1
L200:
	;
	if v695 < int32(0) {
		v706 = v695
		goto L1
	} else {
		goto L201
	}
L201:
	;
	goto L179
}
func F_russian_UTF_8_stem(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v160 int32
	_ = v160
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v211 int32
	_ = v211
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v237 int32
	_ = v237
	var v244 int32
	_ = v244
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v284 int32
	_ = v284
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v335 int32
	_ = v335
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v361 int32
	_ = v361
	var v369 int32
	_ = v369
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v407 int32
	_ = v407
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v420 int32
	_ = v420
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v429 int32
	_ = v429
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v445 int32
	_ = v445
	var v458 int32
	_ = v458
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v478 int32
	_ = v478
	var v484 int32
	_ = v484
	var v491 int32
	_ = v491
	var v502 int32
	_ = v502
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v529 int32
	_ = v529
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v542 int32
	_ = v542
	var v545 int32
	_ = v545
	var v547 int32
	_ = v547
	var v551 int32
	_ = v551
	var v561 int32
	_ = v561
	var v563 int32
	_ = v563
	var v567 int32
	_ = v567
	var v580 int32
	_ = v580
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v600 int32
	_ = v600
	var v606 int32
	_ = v606
	var v614 int32
	_ = v614
	var v625 int32
	_ = v625
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v634 int32
	_ = v634
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v647 int32
	_ = v647
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v654 int32
	_ = v654
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v660 int32
	_ = v660
	var v663 int32
	_ = v663
	var v667 int32
	_ = v667
	var v670 int32
	_ = v670
	var v672 int32
	_ = v672
	var v674 int32
	_ = v674
	var v676 int32
	_ = v676
	var v679 int32
	_ = v679
	var v682 int32
	_ = v682
	var v685 int32
	_ = v685
	var v689 int32
	_ = v689
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v702 int32
	_ = v702
	var v705 int32
	_ = v705
	var v709 int32
	_ = v709
	var v713 int32
	_ = v713
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v723 int32
	_ = v723
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v744 int32
	_ = v744
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v752 int32
	_ = v752
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v759 int32
	_ = v759
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v765 int32
	_ = v765
	var v768 int32
	_ = v768
	var v772 int32
	_ = v772
	var v775 int32
	_ = v775
	var v777 int32
	_ = v777
	var v779 int32
	_ = v779
	var v781 int32
	_ = v781
	var v784 int32
	_ = v784
	var v787 int32
	_ = v787
	var v790 int32
	_ = v790
	var v794 int32
	_ = v794
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v809 int32
	_ = v809
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v818 int32
	_ = v818
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v825 int32
	_ = v825
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v831 int32
	_ = v831
	var v834 int32
	_ = v834
	var v838 int32
	_ = v838
	var v841 int32
	_ = v841
	var v843 int32
	_ = v843
	var v845 int32
	_ = v845
	var v847 int32
	_ = v847
	var v850 int32
	_ = v850
	var v853 int32
	_ = v853
	var v856 int32
	_ = v856
	var v860 int32
	_ = v860
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v875 int32
	_ = v875
	var v878 int32
	_ = v878
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v889 int32
	_ = v889
	var v890 int32
	_ = v890
	var v893 int32
	_ = v893
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v900 int32
	_ = v900
	var v902 int32
	_ = v902
	var v907 int32
	_ = v907
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v916 int32
	_ = v916
	var v919 int32
	_ = v919
	var v921 int32
	_ = v921
	var v924 int32
	_ = v924
	var v927 int32
	_ = v927
	var v930 int32
	_ = v930
	var v934 int32
	_ = v934
	var v937 int32
	_ = v937
	var v939 int32
	_ = v939
	var v941 int32
	_ = v941
	var v942 int32
	_ = v942
	var v945 int32
	_ = v945
	var v946 int32
	_ = v946
	var v948 int32
	_ = v948
	var v949 int32
	_ = v949
	var v951 int32
	_ = v951
	var v955 int32
	_ = v955
	var v959 int32
	_ = v959
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v968 int32
	_ = v968
	var v970 int32
	_ = v970
	var v971 int32
	_ = v971
	var v973 int32
	_ = v973
	var v974 int32
	_ = v974
	var v978 int32
	_ = v978
	var v980 int32
	_ = v980
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v989 int32
	_ = v989
	var v993 int32
	_ = v993
	var v994 int32
	_ = v994
	var v997 int32
	_ = v997
	var v999 int32
	_ = v999
	var v1001 int32
	_ = v1001
	var v1004 int32
	_ = v1004
	var v1007 int32
	_ = v1007
	var v1010 int32
	_ = v1010
	var v1014 int32
	_ = v1014
	var v1017 int32
	_ = v1017
	var v1019 int32
	_ = v1019
	var v1021 int32
	_ = v1021
	var v1024 int32
	_ = v1024
	var v1027 int32
	_ = v1027
	var v1030 int32
	_ = v1030
	var v1034 int32
	_ = v1034
	var v1037 int32
	_ = v1037
	var v1038 int32
	_ = v1038
	var v1041 int32
	_ = v1041
	var v1043 int32
	_ = v1043
	var v1045 int32
	_ = v1045
	var v1046 int32
	_ = v1046
	var v1049 int32
	_ = v1049
	var v1052 int32
	_ = v1052
	var v1056 int32
	_ = v1056
	var v1059 int32
	_ = v1059
	var v1060 int32
	_ = v1060
	var v1063 int32
	_ = v1063
	var v1064 int32
	_ = v1064
	var v1072 int32
	_ = v1072
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	goto L2
L1:
	;
	return v1072
L2:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v14
	v16 = int32(2)
	v18 = int32(0)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v20-v14 < v16 {
		v30 = v18
		goto L6
	} else {
		goto L7
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v7
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v133))) = v134
	*(*int32)(unsafe.Add(mBase, uint32(v133)+4)) = v134
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v160 = v137
	goto L46
L4:
	;
	goto L3
L5:
	;
	if v30 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	goto L5
L7:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v26 = F_memcmp(m, v24+v14, int32(2182195), v16)
	mBase = m.M
	if v26 != 0 {
		v30 = v18
		goto L6
	} else {
		goto L8
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v16 + v14
	v30 = int32(1)
	goto L6
L9:
	;
	v34 = v14
	goto L12
L10:
	;
	v116 = v14
	goto L11
L11:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v116
	v126 = F_slice_from_s(m, l0, int32(2), int32(2182197))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L40
	} else {
		goto L41
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v34
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L16
L13:
	;
	v116 = v93
	goto L11
L14:
	;
	if v93 < int32(0) {
		goto L4
	} else {
		goto L34
	}
L16:
	;
	goto L17
L17:
	;
	goto L18
L18:
	;
	v48 = v34
	v50 = int32(1)
	goto L21
L20:
	;
	v93 = v78
	goto L14
L21:
	;
	if v41 <= v48 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	goto L20
L23:
	;
	v93 = int32(-1)
	goto L14
L24:
	;
	goto L25
L25:
	;
	v55 = v48 + int32(1)
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40+v48))))
	if base.Ui32(v57) < base.Ui32(int32(192)) {
		v78 = v55
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v79 = int32(1)
	if v79 < v50 {
		v48 = v78
		v50 = v50 - v79
		goto L21
	} else {
		goto L33
	}
L27:
	;
	if v41 <= v55 {
		v78 = v55
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v64 = v55
	goto L29
L29:
	;
	v67 = int32(*(*int8)(unsafe.Add(mBase, uint32(v40+v64))))
	if int32(-65) < v67 {
		v78 = v64
		goto L26
	} else {
		goto L31
	}
L30:
	;
	v78 = v41
	goto L26
L31:
	;
	v71 = v64 + int32(1)
	if v71 != v41 {
		v64 = v71
		goto L29
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	goto L22
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v93
	v98 = int32(2)
	v100 = int32(0)
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v102-v93 < v98 {
		v112 = v100
		goto L36
	} else {
		goto L37
	}
L35:
	;
	if v112 == int32(0) {
		v34 = v93
		goto L12
	} else {
		goto L39
	}
L36:
	;
	goto L35
L37:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v108 = F_memcmp(m, v106+v93, int32(2182195), v98)
	mBase = m.M
	if v108 != 0 {
		v112 = v100
		goto L36
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v98 + v93
	v112 = int32(1)
	goto L36
L39:
	;
	goto L13
L40:
	;
	return int32(0)
L41:
	;
	if int32(0) <= v126 {
		goto L2
	} else {
		goto L42
	}
L42:
	;
	v1072 = v126
	goto L1
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v137
	v634 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v634
	v636 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v637 = *(*int32)(unsafe.Add(mBase, uint32(v636)+4))
	if v634 < v637 {
		v1072 = int32(0)
		goto L1
	} else {
		goto L146
	}
L44:
	;
	if v255 < int32(0) {
		goto L43
	} else {
		goto L69
	}
L45:
	;
	v255 = v227
	goto L44
L46:
	;
	if v151 <= v160 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v255 = int32(-1)
	goto L44
L49:
	;
	goto L50
L50:
	;
	v167 = int32(1)
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160+v152))))
	if base.Ui32(v169) < base.Ui32(int32(192)) {
		v226 = v169
		v227 = v167
		goto L51
	} else {
		goto L52
	}
L51:
	;
	if int32(1103) < v226 {
		goto L64
	} else {
		goto L65
	}
L52:
	;
	v173 = v160 + int32(1)
	if v173 == v151 {
		v226 = v169
		v227 = v167
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173+v152))))
	v178 = v176 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v169) {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182+v152))))
	v194 = v192 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v169) {
		goto L60
	} else {
		goto L61
	}
L55:
	;
	v182 = v160 + int32(2)
	if v182 != v151 {
		goto L54
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v226 = v169<<(uint(int32(6))%32)&int32(1984) | v178
	v227 = int32(2)
	goto L51
L58:
	;
	goto L57
L59:
	;
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152+v198))))
	v226 = v211&int32(63) | (v169<<(uint(int32(18))%32)&int32(1835008) | v178<<(uint(int32(12))%32) | v194<<(uint(int32(6))%32))
	v227 = int32(4)
	goto L51
L60:
	;
	v198 = v160 + int32(3)
	if v198 != v151 {
		goto L59
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v226 = v169<<(uint(int32(12))%32)&int32(61440) | v178<<(uint(int32(6))%32) | v194
	v227 = int32(3)
	goto L51
L63:
	;
	goto L62
L64:
	;
	v244 = v227 + v160
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v244
	v160 = v244
	goto L46
L65:
	;
	v231 = v226 - int32(1072)
	if v231 < int32(0) {
		goto L64
	} else {
		goto L66
	}
L66:
	;
	v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v231)>>(uint(int32(3))%32)))+uint32(_consts[1326]))))
	if int32(base.Ui32(v237)>>(uint(v231&int32(7))%32))&int32(1) != 0 {
		goto L45
	} else {
		goto L67
	}
L67:
	;
	goto L64
L69:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v259 = v258 + v255
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v259
	v261 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v261)+4)) = v259
	v274 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v275 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v276 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v284 = v274
	goto L72
L70:
	;
	if v380 < int32(0) {
		goto L43
	} else {
		goto L94
	}
L71:
	;
	v380 = v351
	goto L70
L72:
	;
	if v275 <= v284 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v380 = int32(-1)
	goto L70
L75:
	;
	goto L76
L76:
	;
	v291 = int32(1)
	v293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v284+v276))))
	if base.Ui32(v293) < base.Ui32(int32(192)) {
		v350 = v293
		v351 = v291
		goto L77
	} else {
		goto L78
	}
L77:
	;
	if int32(1103) < v350 {
		goto L71
	} else {
		goto L90
	}
L78:
	;
	v297 = v284 + int32(1)
	if v297 == v275 {
		v350 = v293
		v351 = v291
		goto L77
	} else {
		goto L79
	}
L79:
	;
	v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v297+v276))))
	v302 = v300 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v293) {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	v316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v306+v276))))
	v318 = v316 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v293) {
		goto L86
	} else {
		goto L87
	}
L81:
	;
	v306 = v284 + int32(2)
	if v306 != v275 {
		goto L80
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	v350 = v293<<(uint(int32(6))%32)&int32(1984) | v302
	v351 = int32(2)
	goto L77
L84:
	;
	goto L83
L85:
	;
	v335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v276+v322))))
	v350 = v335&int32(63) | (v293<<(uint(int32(18))%32)&int32(1835008) | v302<<(uint(int32(12))%32) | v318<<(uint(int32(6))%32))
	v351 = int32(4)
	goto L77
L86:
	;
	v322 = v284 + int32(3)
	if v322 != v275 {
		goto L85
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	v350 = v293<<(uint(int32(12))%32)&int32(61440) | v302<<(uint(int32(6))%32) | v318
	v351 = int32(3)
	goto L77
L89:
	;
	goto L88
L90:
	;
	v355 = v350 - int32(1072)
	if v355 < int32(0) {
		goto L71
	} else {
		goto L91
	}
L91:
	;
	v361 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v355)>>(uint(int32(3))%32)))+uint32(_consts[1326]))))
	if int32(base.Ui32(v361)>>(uint(v355&int32(7))%32))&int32(1) == int32(0) {
		goto L71
	} else {
		goto L92
	}
L92:
	;
	v369 = v351 + v284
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v369
	v284 = v369
	goto L72
L94:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v384 = v383 + v380
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v384
	v398 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v399 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v407 = v384
	goto L97
L95:
	;
	if v502 < int32(0) {
		goto L43
	} else {
		goto L120
	}
L96:
	;
	v502 = v474
	goto L95
L97:
	;
	if v398 <= v407 {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v502 = int32(-1)
	goto L95
L100:
	;
	goto L101
L101:
	;
	v414 = int32(1)
	v416 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v407+v399))))
	if base.Ui32(v416) < base.Ui32(int32(192)) {
		v473 = v416
		v474 = v414
		goto L102
	} else {
		goto L103
	}
L102:
	;
	if int32(1103) < v473 {
		goto L115
	} else {
		goto L116
	}
L103:
	;
	v420 = v407 + int32(1)
	if v420 == v398 {
		v473 = v416
		v474 = v414
		goto L102
	} else {
		goto L104
	}
L104:
	;
	v423 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v420+v399))))
	v425 = v423 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v416) {
		goto L106
	} else {
		goto L107
	}
L105:
	;
	v439 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v429+v399))))
	v441 = v439 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v416) {
		goto L111
	} else {
		goto L112
	}
L106:
	;
	v429 = v407 + int32(2)
	if v429 != v398 {
		goto L105
	} else {
		goto L109
	}
L107:
	;
	goto L108
L108:
	;
	v473 = v416<<(uint(int32(6))%32)&int32(1984) | v425
	v474 = int32(2)
	goto L102
L109:
	;
	goto L108
L110:
	;
	v458 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v399+v445))))
	v473 = v458&int32(63) | (v416<<(uint(int32(18))%32)&int32(1835008) | v425<<(uint(int32(12))%32) | v441<<(uint(int32(6))%32))
	v474 = int32(4)
	goto L102
L111:
	;
	v445 = v407 + int32(3)
	if v445 != v398 {
		goto L110
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	v473 = v416<<(uint(int32(12))%32)&int32(61440) | v425<<(uint(int32(6))%32) | v441
	v474 = int32(3)
	goto L102
L114:
	;
	goto L113
L115:
	;
	v491 = v474 + v407
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v491
	v407 = v491
	goto L97
L116:
	;
	v478 = v473 - int32(1072)
	if v478 < int32(0) {
		goto L115
	} else {
		goto L117
	}
L117:
	;
	v484 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v478)>>(uint(int32(3))%32)))+uint32(_consts[1326]))))
	if int32(base.Ui32(v484)>>(uint(v478&int32(7))%32))&int32(1) != 0 {
		goto L96
	} else {
		goto L118
	}
L118:
	;
	goto L115
L120:
	;
	v505 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v506 = v505 + v502
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v506
	v520 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v521 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v529 = v506
	goto L123
L121:
	;
	if v625 < int32(0) {
		goto L43
	} else {
		goto L145
	}
L122:
	;
	v625 = v596
	goto L121
L123:
	;
	if v520 <= v529 {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v625 = int32(-1)
	goto L121
L126:
	;
	goto L127
L127:
	;
	v536 = int32(1)
	v538 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v529+v521))))
	if base.Ui32(v538) < base.Ui32(int32(192)) {
		v595 = v538
		v596 = v536
		goto L128
	} else {
		goto L129
	}
L128:
	;
	if int32(1103) < v595 {
		goto L122
	} else {
		goto L141
	}
L129:
	;
	v542 = v529 + int32(1)
	if v542 == v520 {
		v595 = v538
		v596 = v536
		goto L128
	} else {
		goto L130
	}
L130:
	;
	v545 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v542+v521))))
	v547 = v545 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v538) {
		goto L132
	} else {
		goto L133
	}
L131:
	;
	v561 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v551+v521))))
	v563 = v561 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v538) {
		goto L137
	} else {
		goto L138
	}
L132:
	;
	v551 = v529 + int32(2)
	if v551 != v520 {
		goto L131
	} else {
		goto L135
	}
L133:
	;
	goto L134
L134:
	;
	v595 = v538<<(uint(int32(6))%32)&int32(1984) | v547
	v596 = int32(2)
	goto L128
L135:
	;
	goto L134
L136:
	;
	v580 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v521+v567))))
	v595 = v580&int32(63) | (v538<<(uint(int32(18))%32)&int32(1835008) | v547<<(uint(int32(12))%32) | v563<<(uint(int32(6))%32))
	v596 = int32(4)
	goto L128
L137:
	;
	v567 = v529 + int32(3)
	if v567 != v520 {
		goto L136
	} else {
		goto L140
	}
L138:
	;
	goto L139
L139:
	;
	v595 = v538<<(uint(int32(12))%32)&int32(61440) | v547<<(uint(int32(6))%32) | v563
	v596 = int32(3)
	goto L128
L140:
	;
	goto L139
L141:
	;
	v600 = v595 - int32(1072)
	if v600 < int32(0) {
		goto L122
	} else {
		goto L142
	}
L142:
	;
	v606 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v600)>>(uint(int32(3))%32)))+uint32(_consts[1326]))))
	if int32(base.Ui32(v606)>>(uint(v600&int32(7))%32))&int32(1) == int32(0) {
		goto L122
	} else {
		goto L143
	}
L143:
	;
	v614 = v596 + v529
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v614
	v529 = v614
	goto L123
L145:
	;
	v628 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v629 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v628))) = v629 + v625
	goto L43
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v634
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v637
	v643 = F_find_among_b(m, l0, int32(4291776), int32(9))
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
		goto L40
	} else {
		goto L149
	}
L147:
	;
	v916 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v916
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v916
	v919 = int32(2)
	v921 = int32(0)
	v924 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v916-v924 < v919 {
		v934 = v921
		goto L245
	} else {
		goto L246
	}
L148:
	;
	v702 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v702
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v702
	v705 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v702-int32(3) <= v705 {
		v721 = v702
		goto L171
	} else {
		goto L172
	}
L149:
	;
	if v643 == int32(0) {
		goto L148
	} else {
		goto L150
	}
L150:
	;
	v647 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v647
	switch v643 - int32(1) {
	case 0:
		goto L152
	case 1:
		goto L151
	default:
		goto L147
	}
L151:
	;
	v696 = F_slice_del(m, l0)
	mBase = m.M
	v697 = m.ExcPending
	if v697 != 0 {
		goto L40
	} else {
		goto L167
	}
L152:
	;
	v651 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v652 = int32(2)
	v654 = int32(0)
	v656 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v657 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v656-v657 < v652 {
		v667 = v654
		goto L154
	} else {
		goto L155
	}
L153:
	;
	if v667 == int32(0) {
		goto L157
	} else {
		goto L158
	}
L154:
	;
	goto L153
L155:
	;
	v660 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v663 = F_memcmp(m, v660+v656-v652, int32(2182205), v652)
	mBase = m.M
	if v663 != 0 {
		v667 = v654
		goto L154
	} else {
		goto L156
	}
L156:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v656 - v652
	v667 = int32(1)
	goto L154
L157:
	;
	v670 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v672 = v670 + (v647 - v651)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v672
	v674 = int32(2)
	v676 = int32(0)
	v679 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v672-v679 < v674 {
		v689 = v676
		goto L161
	} else {
		goto L162
	}
L158:
	;
	goto L159
L159:
	;
	v692 = F_slice_del(m, l0)
	mBase = m.M
	v693 = m.ExcPending
	if v693 != 0 {
		goto L40
	} else {
		goto L165
	}
L160:
	;
	if v689 == int32(0) {
		goto L148
	} else {
		goto L164
	}
L161:
	;
	goto L160
L162:
	;
	v682 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v685 = F_memcmp(m, v682+v672-v674, int32(2182207), v674)
	mBase = m.M
	if v685 != 0 {
		v689 = v676
		goto L161
	} else {
		goto L163
	}
L163:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v672 - v674
	v689 = int32(1)
	goto L161
L164:
	;
	goto L159
L165:
	;
	if int32(0) <= v692 {
		goto L147
	} else {
		goto L166
	}
L166:
	;
	v1072 = v692
	goto L1
L167:
	;
	if int32(0) <= v696 {
		goto L147
	} else {
		goto L168
	}
L168:
	;
	v1072 = v696
	goto L1
L169:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v731
	v736 = F_find_among_b(m, l0, int32(4292176), int32(26))
	mBase = m.M
	v737 = m.ExcPending
	if v737 != 0 {
		goto L40
	} else {
		goto L178
	}
L170:
	;
	v723 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v723
	v725 = F_slice_del(m, l0)
	mBase = m.M
	v726 = m.ExcPending
	if v726 != 0 {
		goto L40
	} else {
		goto L176
	}
L171:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v721
	v731 = v721
	v732 = v721
	goto L169
L172:
	;
	v709 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v713 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v709+v702-int32(1)))))
	switch v713 - int32(140) {
	case 0, 3:
		goto L173
	default:
		v721 = v702
		goto L171
	}
L173:
	;
	v718 = F_find_among_b(m, l0, int32(4291968), int32(2))
	mBase = m.M
	v719 = m.ExcPending
	if v719 != 0 {
		goto L40
	} else {
		goto L174
	}
L174:
	;
	if v718 != 0 {
		goto L170
	} else {
		goto L175
	}
L175:
	;
	v720 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v721 = v720
	goto L171
L176:
	;
	if v725 < int32(0) {
		v1072 = v725
		goto L1
	} else {
		goto L177
	}
L177:
	;
	v729 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v730 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v731 = v729
	v732 = v730
	goto L169
L178:
	;
	if v736 != 0 {
		goto L179
	} else {
		goto L180
	}
L179:
	;
	v738 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v738
	v740 = F_slice_del(m, l0)
	mBase = m.M
	v741 = m.ExcPending
	if v741 != 0 {
		goto L40
	} else {
		goto L182
	}
L180:
	;
	goto L181
L181:
	;
	v805 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v806 = v732 - v731
	v807 = v805 - v806
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v807
	v809 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v807
	v814 = F_find_among_b(m, l0, int32(4292704), int32(46))
	mBase = m.M
	v815 = m.ExcPending
	if v815 != 0 {
		goto L40
	} else {
		goto L205
	}
L182:
	;
	if v740 < int32(0) {
		v1072 = v740
		goto L1
	} else {
		goto L183
	}
L183:
	;
	v744 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v744
	v748 = F_find_among_b(m, l0, int32(4292016), int32(8))
	mBase = m.M
	v749 = m.ExcPending
	if v749 != 0 {
		goto L40
	} else {
		goto L184
	}
L184:
	;
	if v748 == int32(0) {
		goto L147
	} else {
		goto L185
	}
L185:
	;
	v752 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v752
	switch v748 - int32(1) {
	case 0:
		goto L187
	case 1:
		goto L186
	default:
		goto L147
	}
L186:
	;
	v801 = F_slice_del(m, l0)
	mBase = m.M
	v802 = m.ExcPending
	if v802 != 0 {
		goto L40
	} else {
		goto L202
	}
L187:
	;
	v756 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v757 = int32(2)
	v759 = int32(0)
	v761 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v762 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v761-v762 < v757 {
		v772 = v759
		goto L189
	} else {
		goto L190
	}
L188:
	;
	if v772 == int32(0) {
		goto L192
	} else {
		goto L193
	}
L189:
	;
	goto L188
L190:
	;
	v765 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v768 = F_memcmp(m, v765+v761-v757, int32(2182283), v757)
	mBase = m.M
	if v768 != 0 {
		v772 = v759
		goto L189
	} else {
		goto L191
	}
L191:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v761 - v757
	v772 = int32(1)
	goto L189
L192:
	;
	v775 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v777 = v775 + (v752 - v756)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v777
	v779 = int32(2)
	v781 = int32(0)
	v784 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v777-v784 < v779 {
		v794 = v781
		goto L196
	} else {
		goto L197
	}
L193:
	;
	goto L194
L194:
	;
	v797 = F_slice_del(m, l0)
	mBase = m.M
	v798 = m.ExcPending
	if v798 != 0 {
		goto L40
	} else {
		goto L200
	}
L195:
	;
	if v794 == int32(0) {
		goto L147
	} else {
		goto L199
	}
L196:
	;
	goto L195
L197:
	;
	v787 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v790 = F_memcmp(m, v787+v777-v779, int32(2182285), v779)
	mBase = m.M
	if v790 != 0 {
		v794 = v781
		goto L196
	} else {
		goto L198
	}
L198:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v777 - v779
	v794 = int32(1)
	goto L196
L199:
	;
	goto L194
L200:
	;
	if int32(0) <= v797 {
		goto L147
	} else {
		goto L201
	}
L201:
	;
	v1072 = v797
	goto L1
L202:
	;
	if int32(0) <= v801 {
		goto L147
	} else {
		goto L203
	}
L203:
	;
	v1072 = v801
	goto L1
L204:
	;
	v878 = int32(base.Ui32(v875) >> (uint(int32(31)) % 32))
	if v875 != 0 {
		goto L228
	} else {
		goto L229
	}
L205:
	;
	if v814 == int32(0) {
		v875 = v809
		goto L204
	} else {
		goto L206
	}
L206:
	;
	v818 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v818
	switch v814 - int32(1) {
	case 0:
		goto L209
	case 1:
		goto L208
	default:
		goto L207
	}
L207:
	;
	v875 = int32(1)
	goto L204
L208:
	;
	v867 = F_slice_del(m, l0)
	mBase = m.M
	v868 = m.ExcPending
	if v868 != 0 {
		goto L40
	} else {
		goto L224
	}
L209:
	;
	v822 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v823 = int32(2)
	v825 = int32(0)
	v827 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v828 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v827-v828 < v823 {
		v838 = v825
		goto L211
	} else {
		goto L212
	}
L210:
	;
	if v838 == int32(0) {
		goto L214
	} else {
		goto L215
	}
L211:
	;
	goto L210
L212:
	;
	v831 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v834 = F_memcmp(m, v831+v827-v823, int32(2182439), v823)
	mBase = m.M
	if v834 != 0 {
		v838 = v825
		goto L211
	} else {
		goto L213
	}
L213:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v827 - v823
	v838 = int32(1)
	goto L211
L214:
	;
	v841 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v843 = v841 + (v818 - v822)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v843
	v845 = int32(2)
	v847 = int32(0)
	v850 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v843-v850 < v845 {
		v860 = v847
		goto L218
	} else {
		goto L219
	}
L215:
	;
	goto L216
L216:
	;
	v863 = F_slice_del(m, l0)
	mBase = m.M
	v864 = m.ExcPending
	if v864 != 0 {
		goto L40
	} else {
		goto L222
	}
L217:
	;
	if v860 == int32(0) {
		v875 = v809
		goto L204
	} else {
		goto L221
	}
L218:
	;
	goto L217
L219:
	;
	v853 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v856 = F_memcmp(m, v853+v843-v845, int32(2182441), v845)
	mBase = m.M
	if v856 != 0 {
		v860 = v847
		goto L218
	} else {
		goto L220
	}
L220:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v843 - v845
	v860 = int32(1)
	goto L218
L221:
	;
	goto L216
L222:
	;
	if int32(0) <= v863 {
		goto L207
	} else {
		goto L223
	}
L223:
	;
	v875 = v863
	goto L204
L224:
	;
	if v867 < int32(0) {
		v875 = v867
		goto L204
	} else {
		goto L225
	}
L225:
	;
	goto L207
L226:
	;
	if v911 != 0 {
		v1072 = v910
		goto L1
	} else {
		goto L242
	}
L227:
	;
	v881 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v882 = v881 - v806
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v882
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v882
	v889 = F_find_among_b(m, l0, int32(4293632), int32(36))
	mBase = m.M
	v890 = m.ExcPending
	if v890 != 0 {
		goto L40
	} else {
		goto L232
	}
L228:
	;
	v880 = v878
	goto L230
L229:
	;
	v880 = int32(13)
	goto L230
L230:
	;
	switch v880 {
	case 0:
		goto L147
	default:
		v910 = v875
		v911 = v878
		goto L226
	case 13:
		goto L227
	}
L231:
	;
	if v902 == int32(0) {
		goto L147
	} else {
		goto L238
	}
L232:
	;
	if v889 == int32(0) {
		v902 = int32(0)
		goto L231
	} else {
		goto L233
	}
L233:
	;
	v893 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v893
	v896 = F_slice_del(m, l0)
	mBase = m.M
	v897 = m.ExcPending
	if v897 != 0 {
		goto L40
	} else {
		goto L234
	}
L234:
	;
	if int32(0) <= v896 {
		goto L235
	} else {
		goto L236
	}
L235:
	;
	v900 = int32(1)
	goto L237
L236:
	;
	v900 = v896
	goto L237
L237:
	;
	v902 = v900
	goto L231
L238:
	;
	if v902 < int32(0) {
		goto L239
	} else {
		goto L240
	}
L239:
	;
	v907 = v902
	goto L241
L240:
	;
	v907 = v875
	goto L241
L241:
	;
	v910 = v907
	v911 = int32(base.Ui32(v902) >> (uint(int32(31)) % 32))
	goto L226
L242:
	;
	goto L147
L243:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v948
	v951 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v948-int32(5) <= v951 {
		goto L253
	} else {
		goto L254
	}
L244:
	;
	if v934 == int32(0) {
		goto L248
	} else {
		goto L249
	}
L245:
	;
	goto L244
L246:
	;
	v927 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v930 = F_memcmp(m, v927+v916-v919, int32(2182199), v919)
	mBase = m.M
	if v930 != 0 {
		v934 = v921
		goto L245
	} else {
		goto L247
	}
L247:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v916 - v919
	v934 = int32(1)
	goto L245
L248:
	;
	v937 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v937
	v948 = v937
	v949 = v937
	goto L243
L249:
	;
	goto L250
L250:
	;
	v939 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v939
	v941 = F_slice_del(m, l0)
	mBase = m.M
	v942 = m.ExcPending
	if v942 != 0 {
		goto L40
	} else {
		goto L251
	}
L251:
	;
	if v941 < int32(0) {
		v1072 = v941
		goto L1
	} else {
		goto L252
	}
L252:
	;
	v945 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v946 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v948 = v945
	v949 = v946
	goto L243
L253:
	;
	v978 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v980 = v978 + (v948 - v949)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v980
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v980
	v985 = F_find_among_b(m, l0, int32(4294400), int32(4))
	mBase = m.M
	v986 = m.ExcPending
	if v986 != 0 {
		goto L40
	} else {
		goto L262
	}
L254:
	;
	v955 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v959 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v955+v948-int32(1)))))
	switch v959 - int32(130) {
	case 0, 10:
		goto L255
	default:
		goto L253
	}
L255:
	;
	v964 = F_find_among_b(m, l0, int32(4294352), int32(2))
	mBase = m.M
	v965 = m.ExcPending
	if v965 != 0 {
		goto L40
	} else {
		goto L256
	}
L256:
	;
	if v964 == int32(0) {
		goto L253
	} else {
		goto L257
	}
L257:
	;
	v968 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v968
	v970 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v971 = *(*int32)(unsafe.Add(mBase, uint32(v970)))
	if v968 < v971 {
		goto L253
	} else {
		goto L258
	}
L258:
	;
	v973 = F_slice_del(m, l0)
	mBase = m.M
	v974 = m.ExcPending
	if v974 != 0 {
		goto L40
	} else {
		goto L259
	}
L259:
	;
	if v973 < int32(0) {
		v1072 = v973
		goto L1
	} else {
		goto L260
	}
L260:
	;
	goto L253
L261:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v137
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v137
	v1072 = int32(1)
	goto L1
L262:
	;
	if v985 == int32(0) {
		goto L261
	} else {
		goto L263
	}
L263:
	;
	v989 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v989
	switch v985 - int32(1) {
	case 0:
		goto L266
	case 1:
		goto L265
	case 2:
		goto L264
	default:
		goto L261
	}
L264:
	;
	v1063 = F_slice_del(m, l0)
	mBase = m.M
	v1064 = m.ExcPending
	if v1064 != 0 {
		goto L40
	} else {
		goto L288
	}
L265:
	;
	v1041 = int32(2)
	v1043 = int32(0)
	v1045 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1046 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1045-v1046 < v1041 {
		v1056 = v1043
		goto L282
	} else {
		goto L283
	}
L266:
	;
	v993 = F_slice_del(m, l0)
	mBase = m.M
	v994 = m.ExcPending
	if v994 != 0 {
		goto L40
	} else {
		goto L267
	}
L267:
	;
	if v993 < int32(0) {
		v1072 = v993
		goto L1
	} else {
		goto L268
	}
L268:
	;
	v997 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v997
	v999 = int32(2)
	v1001 = int32(0)
	v1004 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v997-v1004 < v999 {
		v1014 = v1001
		goto L270
	} else {
		goto L271
	}
L269:
	;
	if v1014 == int32(0) {
		goto L261
	} else {
		goto L273
	}
L270:
	;
	goto L269
L271:
	;
	v1007 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1010 = F_memcmp(m, v1007+v997-v999, int32(2182819), v999)
	mBase = m.M
	if v1010 != 0 {
		v1014 = v1001
		goto L270
	} else {
		goto L272
	}
L272:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v997 - v999
	v1014 = int32(1)
	goto L270
L273:
	;
	v1017 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1017
	v1019 = int32(2)
	v1021 = int32(0)
	v1024 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1017-v1024 < v1019 {
		v1034 = v1021
		goto L275
	} else {
		goto L276
	}
L274:
	;
	if v1034 == int32(0) {
		goto L261
	} else {
		goto L278
	}
L275:
	;
	goto L274
L276:
	;
	v1027 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1030 = F_memcmp(m, v1027+v1017-v1019, int32(2182821), v1019)
	mBase = m.M
	if v1030 != 0 {
		v1034 = v1021
		goto L275
	} else {
		goto L277
	}
L277:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1017 - v1019
	v1034 = int32(1)
	goto L275
L278:
	;
	v1037 = F_slice_del(m, l0)
	mBase = m.M
	v1038 = m.ExcPending
	if v1038 != 0 {
		goto L40
	} else {
		goto L279
	}
L279:
	;
	if int32(0) <= v1037 {
		goto L261
	} else {
		goto L280
	}
L280:
	;
	v1072 = v1037
	goto L1
L281:
	;
	if v1056 == int32(0) {
		goto L261
	} else {
		goto L285
	}
L282:
	;
	goto L281
L283:
	;
	v1049 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1052 = F_memcmp(m, v1049+v1045-v1041, int32(2182823), v1041)
	mBase = m.M
	if v1052 != 0 {
		v1056 = v1043
		goto L282
	} else {
		goto L284
	}
L284:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1045 - v1041
	v1056 = int32(1)
	goto L282
L285:
	;
	v1059 = F_slice_del(m, l0)
	mBase = m.M
	v1060 = m.ExcPending
	if v1060 != 0 {
		goto L40
	} else {
		goto L286
	}
L286:
	;
	if int32(0) <= v1059 {
		goto L261
	} else {
		goto L287
	}
L287:
	;
	v1072 = v1059
	goto L1
L288:
	;
	if v1063 < int32(0) {
		v1072 = v1063
		goto L1
	} else {
		goto L289
	}
L289:
	;
	goto L261
}
