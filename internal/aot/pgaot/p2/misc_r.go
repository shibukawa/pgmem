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
				v38 = int32(_a_F_ReadDirExtended_0)
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
						F_errfinish(m, int32(_a_F_ReadDirExtended_1), v39, int32(_a_F_ReadDirExtended_2))
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
		*(*int32)(unsafe.Add(mBase, _c_F_ReadDirExtended[0])) = int32(0)
		v24 = F_readdir(m, l0)
		mBase = m.M
		if v24 != 0 {
			v50 = v24
			m.G0 = v7 + int32(16)
			return v50
		} else {
			v25 = int32(0)
			v27 = *(*int32)(unsafe.Add(mBase, _c_F_ReadDirExtended[0]))
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
						v38 = int32(_a_F_ReadDirExtended_3)
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
								F_errfinish(m, int32(_a_F_ReadDirExtended_1), v39, int32(_a_F_ReadDirExtended_2))
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
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_RegisterRelcacheInvalidation[0]))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v10 < v11 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v86 = F_GetCurrentCommandId(m, int32(1))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L19
	} else {
		goto L22
	}
L2:
	;
	v17 = v10
	goto L5
L3:
	;
	goto L4
L4:
	;
	v43 = *(*int32)(unsafe.Add(mBase, _c_F_RegisterRelcacheInvalidation[1]))
	if v43 <= v11 {
		goto L12
	} else {
		goto L13
	}
L5:
	;
	v22 = v9 + v17<<(uint(int32(4))%32)
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	if v23 == int32(254) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	goto L4
L7:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	if base.B2i32(v26 == l2)|base.B2i32(v26 == int32(0)) != 0 {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v33 = v17 + int32(1)
	if v33 != v11 {
		v17 = v33
		goto L5
	} else {
		goto L11
	}
L10:
	;
	goto L9
L11:
	;
	goto L6
L12:
	;
	if v9 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	v65 = v9
	goto L14
L14:
	;
	v69 = v65 + v11<<(uint(int32(4))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v69)+8)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v69)+4)) = l1
	v72 = int32(254)
	*(*uint8)(unsafe.Add(mBase, uint32(v69))) = uint8(v72)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v74 + int32(1)
	goto L1
L15:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_RegisterRelcacheInvalidation[1])) = v59
	*(*int32)(unsafe.Add(mBase, _c_F_RegisterRelcacheInvalidation[0])) = v60
	v65 = v60
	goto L14
L16:
	;
	v49 = *(*int32)(unsafe.Add(mBase, _c_F_RegisterRelcacheInvalidation[2]))
	v51 = F_MemoryContextAlloc(m, v49, int32(512))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	v57 = F_repalloc(m, v9, v43<<(uint(int32(5))%32))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L19
	} else {
		goto L21
	}
L19:
	;
	return
L20:
	;
	v59 = int32(32)
	v60 = v51
	goto L15
L21:
	;
	v59 = v43 << (uint(int32(1)) % 32)
	v60 = v57
	goto L15
L22:
	;
	if l2 != 0 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	return
L24:
	;
	if base.B2i32(l2 == int32(2671))|base.B2i32(base.Ui32(l2-int32(3592)) < base.Ui32(int32(2)))|base.B2i32(l2 == int32(2701)) != 0 {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	goto L26
L26:
	;
	v145 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v145)
	goto L23
L27:
	;
	v141 = int32(1)
	goto L29
L28:
	;
	v100 = int32(0)
	v106 = *(*int32)(unsafe.Add(mBase, _c_F_RegisterRelcacheInvalidation[3]))
	v108 = v106 - int32(1)
	if v108 < v100 {
		v140 = v100
		goto L31
	} else {
		goto L32
	}
L29:
	;
	if v141 == int32(0) {
		goto L23
	} else {
		goto L43
	}
L30:
	;
	v141 = v140
	goto L29
L31:
	;
	goto L30
L32:
	;
	v112 = v108
	v113 = v100
	goto L33
L33:
	;
	v118 = int32(2)
	v119 = base.I32_div_s(v112-v113, v118)
	v120 = v119 + v113
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v120<<(uint(v118)%32))+uint32(_c_F_RegisterRelcacheInvalidation[4])))
	v126 = base.B2i32(v125 == l2)
	if v125 == l2 {
		v140 = v126
		goto L31
	} else {
		goto L35
	}
L34:
	;
	v140 = v126
	goto L31
L35:
	;
	v129 = base.B2i32(base.Ui32(v125) < base.Ui32(l2))
	if base.Ui32(v125) < base.Ui32(l2) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v130 = v120 + int32(1)
	goto L38
L37:
	;
	v130 = v113
	goto L38
L38:
	;
	if base.Ui32(v125) < base.Ui32(l2) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v133 = v112
	goto L41
L40:
	;
	v133 = v120 - int32(1)
	goto L41
L41:
	;
	if v130 <= v133 {
		v112 = v133
		v113 = v130
		goto L33
	} else {
		goto L42
	}
L42:
	;
	goto L34
L43:
	;
	goto L26
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
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v51 int32
	_ = v51
	var v59 int32
	_ = v59
	var v62 int64
	_ = v62
	var v64 int32
	_ = v64
	var v65 int64
	_ = v65
	var v67 int64
	_ = v67
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
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
	var v132 int32
	_ = v132
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
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
	var v179 int32
	_ = v179
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v233 int32
	_ = v233
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
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
	var v266 int32
	_ = v266
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v293 int32
	_ = v293
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v330 int32
	_ = v330
	var v350 int32
	_ = v350
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v398 int32
	_ = v398
	var v403 int32
	_ = v403
	var v408 int32
	_ = v408
	v13 = m.G0
	v15 = v13 - int32(32)
	m.G0 = v15
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_ReleaseOneSerializableXact[0]))
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
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_ReleaseOneSerializableXact[1]))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+72))
	if v27 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	if v30&int32(1) != 0 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	v30 = int32(1)
	goto L6
L5:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+76)))
	v30 = v29
	goto L6
L6:
	;
	goto L3
L7:
	;
	v36 = F_LWLockAcquire(m, l0+int32(72), int32(0))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v39 = int32(0)
	v42 = l0 + int32(48)
	if base.B2i32(v38 == v39)|base.B2i32(v42 == v38) == v39 {
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
	v391 = m.ExcPending
	if v391 != 0 {
		goto L1
	} else {
		goto L89
	}
L12:
	;
	v51 = v38
	goto L15
L13:
	;
	goto L14
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v42
	v197 = *(*int32)(unsafe.Add(mBase, _c_F_ReleaseOneSerializableXact[1]))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v197)+72))
	if v198 != 0 {
		goto L44
	} else {
		goto L45
	}
L15:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	v62 = *(*int64)(unsafe.Add(mBase, uint32(v51-int32(16))))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+24)) = v62
	v64 = base.I32_wrap_i64(v62)
	v65 = *(*int64)(unsafe.Add(mBase, uint32(v64)))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+8)) = v65
	v67 = *(*int64)(unsafe.Add(mBase, uint32(v64)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+16)) = v67
	v70 = *(*int32)(unsafe.Add(mBase, _c_F_ReleaseOneSerializableXact[2]))
	v73 = F_get_hash_value(m, v70, v15+int32(8))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L17
	}
L16:
	;
	goto L14
L17:
	;
	v76 = *(*int32)(unsafe.Add(mBase, _c_F_ReleaseOneSerializableXact[0]))
	v83 = v76 + v73&int32(15)<<(uint(int32(7))%32) + int32(_a_F_ReleaseOneSerializableXact_0)
	v85 = F_LWLockAcquire(m, v83, int32(0))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v88 = v51 - int32(8)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	v90 = int32(4)
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v51-v90)))
	*(*int32)(unsafe.Add(mBase, uint32(v89)+4)) = v92
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	*(*int32)(unsafe.Add(mBase, uint32(v92))) = v94
	v97 = *(*int32)(unsafe.Add(mBase, _c_F_ReleaseOneSerializableXact[3]))
	v99 = v15 + int32(24)
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
	v106 = F_hash_search_with_hash_value(m, v97, v99, v73^v100<<(uint(v90)%32), int32(2), int32(0))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
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
	F_LWLockRelease(m, v83)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L1
	} else {
		goto L41
	}
L21:
	;
	v109 = *(*int32)(unsafe.Add(mBase, _c_F_ReleaseOneSerializableXact[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = v109
	v112 = *(*int32)(unsafe.Add(mBase, _c_F_ReleaseOneSerializableXact[3]))
	v119 = F_hash_search_with_hash_value(m, v112, v99, v109<<(uint(int32(4))%32)^v73, int32(3), v15+int32(7))
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
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v64)+20))
	if v162 != v64+int32(16) {
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
	v131 = v64 + int32(16)
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v64)+20))
	if v132 == int32(0) {
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
	*(*int32)(unsafe.Add(mBase, uint32(v64)+20)) = v131
	*(*int32)(unsafe.Add(mBase, uint32(v64)+16)) = v131
	goto L32
L31:
	;
	goto L32
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v119)+12)) = v131
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v131)))
	*(*int32)(unsafe.Add(mBase, uint32(v119)+8)) = v138
	v141 = v119 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v138)+4)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v131))) = v141
	v145 = *(*int32)(unsafe.Add(mBase, _c_F_ReleaseOneSerializableXact[4]))
	v147 = v145 + int32(48)
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v145)+52))
	if v148 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v145)+52)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v145)+48)) = v147
	goto L35
L34:
	;
	goto L35
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v119)+20)) = v147
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	*(*int32)(unsafe.Add(mBase, uint32(v119)+16)) = v154
	v157 = v119 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v154)+4)) = v157
	*(*int32)(unsafe.Add(mBase, uint32(v147))) = v157
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
	v169 = *(*int32)(unsafe.Add(mBase, _c_F_ReleaseOneSerializableXact[2]))
	v172 = F_hash_search_with_hash_value(m, v169, v64, v73, int32(2), int32(0))
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
	if v59 != v42 {
		v51 = v59
		goto L15
	} else {
		goto L42
	}
L42:
	;
	goto L16
L43:
	;
	if v201&int32(1) != 0 {
		goto L47
	} else {
		goto L48
	}
L44:
	;
	v201 = int32(1)
	goto L46
L45:
	;
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197)+76)))
	v201 = v200
	goto L46
L46:
	;
	goto L43
L47:
	;
	F_LWLockRelease(m, l0+int32(72))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L1
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	v209 = *(*int32)(unsafe.Add(mBase, _c_F_ReleaseOneSerializableXact[0]))
	F_LWLockRelease(m, v209+int32(3840))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L1
	} else {
		goto L51
	}
L50:
	;
	goto L49
L51:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v214
	v217 = *(*int32)(unsafe.Add(mBase, _c_F_ReleaseOneSerializableXact[0]))
	v221 = F_LWLockAcquire(m, v217+int32(3584), int32(0))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
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
	v283 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v283 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L54:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v223 == int32(0) {
		goto L53
	} else {
		goto L55
	}
L55:
	;
	v227 = l0 + int32(32)
	if v223 == v227 {
		goto L53
	} else {
		goto L56
	}
L56:
	;
	v233 = v223
	goto L57
L57:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v233)+4))
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
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v233)+20))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v242)+108))
	*(*int32)(unsafe.Add(mBase, uint32(v242)+108)) = v243 | int32(512)
	goto L61
L60:
	;
	goto L61
L61:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v233)+8))
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v233)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v248)+4)) = v249
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v233)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v249))) = v251
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v233)))
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v233)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v253)+4)) = v254
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v233)))
	*(*int32)(unsafe.Add(mBase, uint32(v254))) = v256
	v259 = *(*int32)(unsafe.Add(mBase, _c_F_ReleaseOneSerializableXact[5]))
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v259)+4))
	if v260 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v259)+4)) = v259
	*(*int32)(unsafe.Add(mBase, uint32(v259))) = v259
	goto L64
L63:
	;
	goto L64
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v233)+4)) = v259
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v259)))
	*(*int32)(unsafe.Add(mBase, uint32(v233))) = v266
	*(*int32)(unsafe.Add(mBase, uint32(v266)+4)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v259))) = v233
	if v227 != v241 {
		v233 = v241
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
	v287 = l0 + int32(40)
	if v283 == v287 {
		goto L66
	} else {
		goto L68
	}
L68:
	;
	v293 = v283
	goto L69
L69:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v293)+4))
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
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v293)+8))
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v302)+108))
	*(*int32)(unsafe.Add(mBase, uint32(v302)+108)) = v303 | int32(1024)
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v293)+4))
	v308 = v307
	goto L73
L72:
	;
	v308 = v301
	goto L73
L73:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v293)))
	*(*int32)(unsafe.Add(mBase, uint32(v309)+4)) = v308
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v293)))
	*(*int32)(unsafe.Add(mBase, uint32(v308))) = v311
	v314 = v293 - int32(8)
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v314)))
	v317 = v293 - int32(4)
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v317)))
	*(*int32)(unsafe.Add(mBase, uint32(v315)+4)) = v318
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v314)))
	*(*int32)(unsafe.Add(mBase, uint32(v318))) = v320
	v323 = *(*int32)(unsafe.Add(mBase, _c_F_ReleaseOneSerializableXact[5]))
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v323)+4))
	if v324 == int32(0) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v323)+4)) = v323
	*(*int32)(unsafe.Add(mBase, uint32(v323))) = v323
	goto L76
L75:
	;
	goto L76
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v317))) = v323
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v323)))
	*(*int32)(unsafe.Add(mBase, uint32(v314))) = v330
	*(*int32)(unsafe.Add(mBase, uint32(v330)+4)) = v314
	*(*int32)(unsafe.Add(mBase, uint32(v323))) = v314
	if v287 != v301 {
		v293 = v301
		goto L69
	} else {
		goto L77
	}
L77:
	;
	goto L70
L78:
	;
	if v214 != 0 {
		goto L81
	} else {
		goto L82
	}
L79:
	;
	goto L80
L80:
	;
	v380 = *(*int32)(unsafe.Add(mBase, _c_F_ReleaseOneSerializableXact[0]))
	F_LWLockRelease(m, v380+int32(3584))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L1
	} else {
		goto L88
	}
L81:
	;
	v350 = *(*int32)(unsafe.Add(mBase, _c_F_ReleaseOneSerializableXact[6]))
	v355 = F_hash_search(m, v350, v15+int32(8), int32(2), int32(0))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L1
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v358 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v357)+4)) = v358
	v360 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v358))) = v360
	v363 = *(*int32)(unsafe.Add(mBase, _c_F_ReleaseOneSerializableXact[7]))
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v363)+4))
	if v364 == int32(0) {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	goto L83
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v363)+4)) = v363
	*(*int32)(unsafe.Add(mBase, uint32(v363))) = v363
	goto L87
L86:
	;
	goto L87
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v363
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v363)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v370
	v373 = l0 - int32(-64)
	*(*int32)(unsafe.Add(mBase, uint32(v370)+4)) = v373
	*(*int32)(unsafe.Add(mBase, uint32(v363))) = v373
	goto L80
L88:
	;
	m.G0 = v15 + int32(32)
	return
L89:
	;
	F_errcode(m, int32(_a_F_ReleaseOneSerializableXact_1))
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	F_errmsg(m, int32(_a_F_ReleaseOneSerializableXact_2), int32(0))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(_a_F_ReleaseOneSerializableXact_3)
	F_errhint(m, int32(_a_F_ReleaseOneSerializableXact_4), v15)
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	F_errfinish(m, int32(_a_F_ReleaseOneSerializableXact_5), int32(3891), int32(_a_F_ReleaseOneSerializableXact_6))
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
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
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_ReleaseSemaphores[0]))
	if v3 < v5 {
		v8 = v3
		for {
			v13 = v8 + int32(1)
			v15 = *(*int32)(unsafe.Add(mBase, _c_F_ReleaseSemaphores[0]))
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
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v24 int64
	_ = v24
	var v25 int64
	_ = v25
	var v28 int64
	_ = v28
	var v29 int64
	_ = v29
	var v32 int64
	_ = v32
	var v35 int32
	_ = v35
	var v37 float64
	_ = v37
	var v39 float64
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int64
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int64
	_ = v77
	var v79 int64
	_ = v79
	var v90 int64
	_ = v90
	var v92 int64
	_ = v92
	var v94 int64
	_ = v94
	var v96 int64
	_ = v96
	var v100 int32
	_ = v100
	var v107 int32
	_ = v107
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int64
	_ = v128
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v184 int32
	_ = v184
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v214 int64
	_ = v214
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v270 int32
	_ = v270
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v357 int64
	_ = v357
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int64
	_ = v370
	var v371 int64
	_ = v371
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v378 int64
	_ = v378
	var v380 int64
	_ = v380
	var v382 int32
	_ = v382
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v392 int64
	_ = v392
	var v400 int32
	_ = v400
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v423 int32
	_ = v423
	v16 = m.G0
	v18 = v16 - int32(144)
	m.G0 = v18
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = int32(0)
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_RemoveOldXlogFiles[0]))
	v24 = base.I64_extend_i32_s(v23)
	v25 = base.I64_div_u_s(l2, v24)
	*(*int64)(unsafe.Add(mBase, uint32(v18)+56)) = v25
	v28 = base.I64_div_u_s(int64(4294967296), v24)
	v29 = base.I64_div_u_s(l0, v28)
	*(*uint32)(unsafe.Add(mBase, uint32(v18)+36)) = uint32(v29)
	v32 = l0 - v28*v29
	*(*uint32)(unsafe.Add(mBase, uint32(v18)+40)) = uint32(v32)
	v35 = *(*int32)(unsafe.Add(mBase, _c_F_RemoveOldXlogFiles[1]))
	v37 = *(*float64)(unsafe.Add(mBase, _c_F_RemoveOldXlogFiles[2]))
	v39 = *(*float64)(unsafe.Add(mBase, _c_F_RemoveOldXlogFiles[3]))
	v41 = *(*int32)(unsafe.Add(mBase, _c_F_RemoveOldXlogFiles[4]))
	v43 = v18 - int32(-64)
	v48 = F_pg_snprintf(m, v43, int32(64), int32(_a_F_RemoveOldXlogFiles_0), v18+int32(32))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v50 = base.I64_div_u_s(l1, v24)
	v52 = base.I32_div_s(v23, int32(_a_F_RemoveOldXlogFiles_1))
	v53 = base.I32_div_s(v41, v52)
	v54 = base.I32_div_s(v35, v52)
	v57 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if v57 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v43
	F_errmsg_internal(m, int32(_a_F_RemoveOldXlogFiles_2), v18+int32(16))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v71 = F_AllocateDir(m, int32(_a_F_RemoveOldXlogFiles_3))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	F_errfinish(m, int32(_a_F_RemoveOldXlogFiles_4), int32(3884), int32(_a_F_RemoveOldXlogFiles_5))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	goto L6
L9:
	;
	v74 = F_ReadDir(m, v71, int32(_a_F_RemoveOldXlogFiles_3))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	if v74 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v77 = v50 - int64(1)
	v79 = v77 + base.I64_extend_i32_s(v53)
	v90 = base.I64_trunc_sat_f64_u(base.F64_ceil(base.F64_div(base.F64_add(base.F64_mul(base.F64_mul(base.F64_add(v39, float64(1)), v37), float64(1.1)), base.F64_convert_i64_u(l1)), base.F64_convert_i32_s(v23))))
	if base.Ui64(v90) < base.Ui64(v79) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	F_FreeDir(m, v71)
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L1
	} else {
		goto L95
	}
L14:
	;
	v92 = v79
	goto L16
L15:
	;
	v92 = v90
	goto L16
L16:
	;
	v94 = v77 + base.I64_extend_i32_s(v54)
	if base.Ui64(v92) < base.Ui64(v94) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v96 = v92
	goto L19
L18:
	;
	v96 = v94
	goto L19
L19:
	;
	v100 = v18 - int32(-64) | int32(8)
	v107 = v74
	goto L20
L20:
	;
	v117 = v107 + int32(19)
	v118 = F_strlen(m, v117)
	mBase = m.M
	switch v118 - int32(24) {
	case 0:
		goto L25
	default:
		goto L22
	case 8:
		goto L24
	}
L21:
	;
	goto L13
L22:
	;
	v405 = F_ReadDir(m, v71, int32(_a_F_RemoveOldXlogFiles_3))
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L1
	} else {
		goto L93
	}
L23:
	;
	v323 = v107 + int32(27)
	v326 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v323))))
	v329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100))))
	if base.B2i32(v326 == int32(0))|base.B2i32(v326 != v329) != 0 {
		v347 = v326
		v348 = v329
		goto L75
	} else {
		goto L76
	}
L24:
	;
	v207 = int32(_a_F_RemoveOldXlogFiles_6)
	v211 = m.G0
	v213 = v211 - int32(32)
	v214 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v213)+24)) = v214
	*(*int64)(unsafe.Add(mBase, uint32(v213)+16)) = v214
	*(*int64)(unsafe.Add(mBase, uint32(v213)+8)) = v214
	*(*int64)(unsafe.Add(mBase, uint32(v213))) = v214
	v222 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_RemoveOldXlogFiles[5])))
	if v222 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L25:
	;
	v121 = int32(_a_F_RemoveOldXlogFiles_6)
	v125 = m.G0
	v127 = v125 - int32(32)
	v128 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v127)+24)) = v128
	*(*int64)(unsafe.Add(mBase, uint32(v127)+16)) = v128
	*(*int64)(unsafe.Add(mBase, uint32(v127)+8)) = v128
	*(*int64)(unsafe.Add(mBase, uint32(v127))) = v128
	v136 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_RemoveOldXlogFiles[5])))
	if v136 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	if v204 == int32(24) {
		goto L23
	} else {
		goto L45
	}
L27:
	;
	v204 = int32(0)
	goto L26
L28:
	;
	goto L29
L29:
	;
	v140 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_RemoveOldXlogFiles[6])))
	if v140 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v144 = v117
	goto L33
L31:
	;
	goto L32
L32:
	;
	v154 = v121
	v155 = v136
	goto L36
L33:
	;
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144))))
	if v150 == v136 {
		v144 = v144 + int32(1)
		goto L33
	} else {
		goto L35
	}
L34:
	;
	v204 = v144 - v117
	goto L26
L35:
	;
	goto L34
L36:
	;
	v162 = v127 + int32(base.Ui32(v155)>>(uint(int32(3))%32))&int32(28)
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v162)))
	v164 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v162))) = v163 | v164<<(uint(v155)%32)
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154)+1)))
	if v168 != 0 {
		v154 = v154 + v164
		v155 = v168
		goto L36
	} else {
		goto L38
	}
L37:
	;
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117))))
	if v171 == int32(0) {
		v194 = v117
		goto L39
	} else {
		goto L40
	}
L38:
	;
	goto L37
L39:
	;
	v204 = v194 - v117
	goto L26
L40:
	;
	v175 = v117
	v176 = v171
	goto L41
L41:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v127+int32(base.Ui32(v176)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v184)>>(uint(v176)%32))&int32(1) == int32(0) {
		v194 = v175
		goto L39
	} else {
		goto L43
	}
L42:
	;
	v194 = v192
	goto L39
L43:
	;
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+1)))
	v192 = v175 + int32(1)
	if v190 != 0 {
		v175 = v192
		v176 = v190
		goto L41
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	goto L22
L46:
	;
	if v290 != int32(24) {
		goto L22
	} else {
		goto L65
	}
L47:
	;
	v290 = int32(0)
	goto L46
L48:
	;
	goto L49
L49:
	;
	v226 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_RemoveOldXlogFiles[6])))
	if v226 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v230 = v117
	goto L53
L51:
	;
	goto L52
L52:
	;
	v240 = v207
	v241 = v222
	goto L56
L53:
	;
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230))))
	if v236 == v222 {
		v230 = v230 + int32(1)
		goto L53
	} else {
		goto L55
	}
L54:
	;
	v290 = v230 - v117
	goto L46
L55:
	;
	goto L54
L56:
	;
	v248 = v213 + int32(base.Ui32(v241)>>(uint(int32(3))%32))&int32(28)
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v248)))
	v250 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v248))) = v249 | v250<<(uint(v241)%32)
	v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v240)+1)))
	if v254 != 0 {
		v240 = v240 + v250
		v241 = v254
		goto L56
	} else {
		goto L58
	}
L57:
	;
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117))))
	if v257 == int32(0) {
		v280 = v117
		goto L59
	} else {
		goto L60
	}
L58:
	;
	goto L57
L59:
	;
	v290 = v280 - v117
	goto L46
L60:
	;
	v261 = v117
	v262 = v257
	goto L61
L61:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v213+int32(base.Ui32(v262)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v270)>>(uint(v262)%32))&int32(1) == int32(0) {
		v280 = v261
		goto L59
	} else {
		goto L63
	}
L62:
	;
	v280 = v278
	goto L59
L63:
	;
	v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v261)+1)))
	v278 = v261 + int32(1)
	if v276 != 0 {
		v261 = v278
		v262 = v276
		goto L61
	} else {
		goto L64
	}
L64:
	;
	goto L62
L65:
	;
	v294 = v107 + int32(43)
	v295 = int32(_a_F_RemoveOldXlogFiles_7)
	v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v294))))
	v301 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_RemoveOldXlogFiles[7])))
	if base.B2i32(v298 == int32(0))|base.B2i32(v298 != v301) != 0 {
		v319 = v298
		v320 = v301
		goto L67
	} else {
		goto L68
	}
L66:
	;
	if v319-v320 != 0 {
		goto L22
	} else {
		goto L73
	}
L67:
	;
	goto L66
L68:
	;
	v304 = v294
	v305 = v295
	goto L69
L69:
	;
	v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v305)+1)))
	v309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v304)+1)))
	if v309 == int32(0) {
		v319 = v309
		v320 = v308
		goto L67
	} else {
		goto L71
	}
L70:
	;
	v319 = v309
	v320 = v308
	goto L67
L71:
	;
	v312 = int32(1)
	if v309 == v308 {
		v304 = v304 + v312
		v305 = v305 + v312
		goto L69
	} else {
		goto L72
	}
L72:
	;
	goto L70
L73:
	;
	goto L23
L74:
	;
	if int32(0) < v347-v348 {
		goto L22
	} else {
		goto L81
	}
L75:
	;
	goto L74
L76:
	;
	v332 = v323
	v333 = v100
	goto L77
L77:
	;
	v336 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v333)+1)))
	v337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v332)+1)))
	if v337 == int32(0) {
		v347 = v337
		v348 = v336
		goto L75
	} else {
		goto L79
	}
L78:
	;
	v347 = v337
	v348 = v336
	goto L75
L79:
	;
	v340 = int32(1)
	if v337 == v336 {
		v332 = v332 + v340
		v333 = v333 + v340
		goto L77
	} else {
		goto L80
	}
L80:
	;
	goto L78
L81:
	;
	v352 = F_XLogArchiveCheckDone(m, v117)
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	if v352 == int32(0) {
		goto L22
	} else {
		goto L83
	}
L83:
	;
	v357 = int64(*(*int32)(unsafe.Add(mBase, _c_F_RemoveOldXlogFiles[0])))
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v18 + int32(132)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = v18 + int32(140)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v18 + int32(136)
	v368 = F_sscanf(m, v117, int32(_a_F_RemoveOldXlogFiles_0), v18)
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	v370 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v18)+136)))
	v371 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v18)+140)))
	v373 = *(*int32)(unsafe.Add(mBase, _c_F_RemoveOldXlogFiles[8]))
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v373)+440))
	*(*int32)(unsafe.Add(mBase, uint32(v373)+440)) = int32(1)
	v378 = base.I64_div_u_s(int64(4294967296), v357)
	v380 = v370 + v371*v378
	if v374 != 0 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v382 = *(*int32)(unsafe.Add(mBase, _c_F_RemoveOldXlogFiles[8]))
	F_s_lock(m, v382+int32(440), int32(_a_F_RemoveOldXlogFiles_4), int32(3817), int32(_a_F_RemoveOldXlogFiles_8))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L1
	} else {
		goto L88
	}
L86:
	;
	goto L87
L87:
	;
	v391 = *(*int32)(unsafe.Add(mBase, _c_F_RemoveOldXlogFiles[8]))
	v392 = *(*int64)(unsafe.Add(mBase, uint32(v391)+232))
	if base.Ui64(v392) < base.Ui64(v380) {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	goto L87
L89:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v391)+232)) = v380
	goto L91
L90:
	;
	goto L91
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v391)+440)) = int32(0)
	F_RemoveXlogFile(m, v107, v96, v18+int32(56), l3)
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	goto L22
L93:
	;
	if v405 != 0 {
		v107 = v405
		goto L20
	} else {
		goto L94
	}
L94:
	;
	goto L21
L95:
	;
	m.G0 = v18 + int32(144)
	return
}
func F_ResetUsage(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v11 int32
	_ = v11
	v1 = int32(_a_F_ResetUsage_0)
	*(*int64)(unsafe.Add(mBase, _c_F_ResetUsage[0])) = int64(4)
	*(*int64)(unsafe.Add(mBase, _c_F_ResetUsage[1])) = int64(3)
	*(*int64)(unsafe.Add(mBase, _c_F_ResetUsage[2])) = int64(2)
	*(*int64)(unsafe.Add(mBase, _c_F_ResetUsage[3])) = int64(1)
	v11 = F___syscall_ret(m, int32(0))
	mBase = m.M
	F_gettimeofday(m, int32(_a_F_ResetUsage_1))
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
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int64
	_ = v76
	var v78 int64
	_ = v78
	var v79 int64
	_ = v79
	var v81 int64
	_ = v81
	var v82 int64
	_ = v82
	var v85 int64
	_ = v85
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v197 int32
	_ = v197
	var v202 int64
	_ = v202
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v353 int32
	_ = v353
	var v359 int32
	_ = v359
	var v365 int32
	_ = v365
	var v389 int32
	_ = v389
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
	var v415 int32
	_ = v415
	var v420 int32
	_ = v420
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v437 int32
	_ = v437
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v458 int32
	_ = v458
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	v6 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(2336)
	m.G0 = v15
	v18 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_RestoreArchivedFile[0])))
	if v18 != int32(1) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_proc_exit(m, int32(1))
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L8
	} else {
		goto L121
	}
L2:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L8
	} else {
		goto L117
	}
L3:
	;
	m.G0 = v15 + int32(2336)
	return v437
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = l1
	v431 = F_pg_snprintf(m, l0, int32(1024), int32(_a_F_RestoreArchivedFile_0), v15)
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L8
	} else {
		goto L116
	}
L5:
	;
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_RestoreArchivedFile[1]))
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
	v30 = v15 + int32(1312)
	v35 = F_pg_snprintf(m, v30, int32(1024), int32(_a_F_RestoreArchivedFile_0), v15+int32(176))
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
	v43 = F___fstatat(m, int32(-100), v30, v15+int32(192), int32(0))
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
	if v43 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v45 = *(*int32)(unsafe.Add(mBase, _c_F_RestoreArchivedFile[2]))
	if v45 == int32(44) {
		goto L10
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v67 = F_unlink(m, v15+int32(1312))
	mBase = m.M
	if v67 != 0 {
		goto L2
	} else {
		goto L20
	}
L15:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L8
	} else {
		goto L16
	}
L16:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L8
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+160)) = v30
	F_errmsg(m, int32(_a_F_RestoreArchivedFile_1), v15+int32(160))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L8
	} else {
		goto L18
	}
L18:
	;
	F_errfinish(m, int32(_a_F_RestoreArchivedFile_2), int32(112), int32(_a_F_RestoreArchivedFile_3))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
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
	v111 = *(*int32)(unsafe.Add(mBase, _c_F_RestoreArchivedFile[1]))
	v113 = v15 + int32(288)
	v114 = m.G0
	v116 = v114 - int32(32)
	m.G0 = v116
	v119 = v15 + int32(1312)
	if v119 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L22:
	;
	F_GetOldestRestartPoint(m, v15+int32(184), v15+int32(180))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
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
	v105 = F_pg_snprintf(m, v15+int32(288), int32(64), int32(_a_F_RestoreArchivedFile_4), v15+int32(128))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L8
	} else {
		goto L27
	}
L25:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v15)+180))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+112)) = v74
	v76 = *(*int64)(unsafe.Add(mBase, uint32(v15)+184))
	v78 = int64(*(*int32)(unsafe.Add(mBase, _c_F_RestoreArchivedFile[3])))
	v79 = base.I64_div_u_s(v76, v78)
	v81 = base.I64_div_u_s(int64(4294967296), v78)
	v82 = base.I64_div_u_s(v79, v81)
	*(*uint32)(unsafe.Add(mBase, uint32(v15)+116)) = uint32(v82)
	v85 = v79 - v81*v82
	*(*uint32)(unsafe.Add(mBase, uint32(v15)+120)) = uint32(v85)
	v93 = F_pg_snprintf(m, v15+int32(288), int32(64), int32(_a_F_RestoreArchivedFile_4), v15+int32(112))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
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
	m.G0 = v116 + int32(32)
	v152 = F_errstart(m, int32(12), int32(0))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L8
	} else {
		goto L37
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v116)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v116)+4)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v116))) = l1
	v128 = F_replace_percent_placeholders(m, v111, int32(_a_F_RestoreArchivedFile_5), int32(_a_F_RestoreArchivedFile_6), v116)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L8
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v130 = F_pstrdup(m, v119)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L8
	} else {
		goto L33
	}
L32:
	;
	v145 = v128
	goto L28
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v116)+24)) = v130
	*(*int32)(unsafe.Add(mBase, uint32(v116)+20)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v116)+16)) = l1
	v139 = F_replace_percent_placeholders(m, v111, int32(_a_F_RestoreArchivedFile_5), int32(_a_F_RestoreArchivedFile_6), v116+int32(16))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L8
	} else {
		goto L34
	}
L34:
	;
	if v130 == int32(0) {
		v145 = v139
		goto L28
	} else {
		goto L35
	}
L35:
	;
	F_pfree(m, v130)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L8
	} else {
		goto L36
	}
L36:
	;
	v145 = v139
	goto L28
L37:
	;
	if v152 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+96)) = v145
	F_errmsg_internal(m, int32(_a_F_RestoreArchivedFile_7), v15+int32(96))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L8
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v166 = F_fflush(m, int32(0))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L8
	} else {
		goto L43
	}
L41:
	;
	F_errfinish(m, int32(_a_F_RestoreArchivedFile_2), int32(159), int32(_a_F_RestoreArchivedFile_3))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L8
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	v169 = *(*int32)(unsafe.Add(mBase, _c_F_RestoreArchivedFile[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v169))) = int32(134217778)
	*(*int32)(unsafe.Add(mBase, _c_F_RestoreArchivedFile[5])) = int32(1)
	v176 = *(*int32)(unsafe.Add(mBase, _c_F_RestoreArchivedFile[6]))
	if v176 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	F_proc_exit(m, int32(1))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L8
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v180 = F_pgl_system(m, v145)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
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
	v183 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_RestoreArchivedFile[5])) = v183
	v186 = *(*int32)(unsafe.Add(mBase, _c_F_RestoreArchivedFile[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v186))) = v183
	F_pfree(m, v145)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L8
	} else {
		goto L49
	}
L49:
	;
	if v180 != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v359 = v180 & int32(127)
	v365 = int32(255)
	goto L105
L51:
	;
	v197 = F___fstatat(m, int32(-100), v15+int32(1312), v15+int32(192), int32(0))
	mBase = m.M
	goto L52
L52:
	;
	if v197 == int32(0) {
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
	v326 = *(*int32)(unsafe.Add(mBase, _c_F_RestoreArchivedFile[2]))
	if v326 == int32(44) {
		goto L96
	} else {
		goto L97
	}
L56:
	;
	v233 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L8
	} else {
		goto L69
	}
L57:
	;
	v202 = *(*int64)(unsafe.Add(mBase, uint32(v15)+216))
	if v202 == l3 {
		goto L56
	} else {
		goto L58
	}
L58:
	;
	v204 = int32(22)
	v208 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_RestoreArchivedFile[7])))
	if v208 != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v209 = int32(14)
	goto L61
L60:
	;
	v209 = v204
	goto L61
L61:
	;
	if l3 <= v202 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v211 = v204
	goto L64
L63:
	;
	v211 = v209
	goto L64
L64:
	;
	v213 = F_errstart(m, v211, int32(0))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L8
	} else {
		goto L65
	}
L65:
	;
	if v213 == int32(0) {
		v437 = v6
		goto L3
	} else {
		goto L66
	}
L66:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v15)+48)) = l3
	*(*int64)(unsafe.Add(mBase, uint32(v15)+40)) = v202
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = l1
	F_errmsg(m, int32(_a_F_RestoreArchivedFile_8), v15+int32(32))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L8
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(_a_F_RestoreArchivedFile_2), int32(216), int32(_a_F_RestoreArchivedFile_3))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L8
	} else {
		goto L68
	}
L68:
	;
	v437 = v6
	goto L3
L69:
	;
	if v233 != 0 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = l1
	F_errmsg(m, int32(_a_F_RestoreArchivedFile_9), v15+int32(16))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L8
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	v247 = v15 + int32(1312)
	if (v247^l0)&int32(3) != 0 {
		goto L78
	} else {
		goto L79
	}
L73:
	;
	F_errfinish(m, int32(_a_F_RestoreArchivedFile_2), int32(223), int32(_a_F_RestoreArchivedFile_3))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L8
	} else {
		goto L74
	}
L74:
	;
	goto L72
L75:
	;
	v437 = int32(1)
	goto L3
L76:
	;
	goto L75
L77:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v302))) = uint8(v301)
	if v301&int32(255) == int32(0) {
		goto L76
	} else {
		goto L92
	}
L78:
	;
	v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v247))))
	v300 = v247
	v301 = v253
	v302 = l0
	goto L77
L79:
	;
	goto L80
L80:
	;
	if v247&int32(3) != 0 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v257 = v247
	v259 = l0
	goto L84
L82:
	;
	v271 = v247
	v273 = l0
	goto L83
L83:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v271)))
	v278 = int32(-2139062144)
	if (int32(16843008)-v275|v275)&v278 != v278 {
		v300 = v271
		v301 = v275
		v302 = v273
		goto L77
	} else {
		goto L88
	}
L84:
	;
	v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v257))))
	*(*uint8)(unsafe.Add(mBase, uint32(v259))) = uint8(v260)
	if v260 == int32(0) {
		goto L76
	} else {
		goto L86
	}
L85:
	;
	v271 = v267
	v273 = v265
	goto L83
L86:
	;
	v264 = int32(1)
	v265 = v259 + v264
	v267 = v257 + v264
	if v267&int32(3) != 0 {
		v257 = v267
		v259 = v265
		goto L84
	} else {
		goto L87
	}
L87:
	;
	goto L85
L88:
	;
	v283 = v271
	v284 = v275
	v285 = v273
	goto L89
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v285))) = v284
	v287 = int32(4)
	v288 = v285 + v287
	v290 = v283 + v287
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v283)+4))
	v295 = int32(-2139062144)
	if (int32(16843008)-v292|v292)&v295 == v295 {
		v283 = v290
		v284 = v292
		v285 = v288
		goto L89
	} else {
		goto L91
	}
L90:
	;
	v300 = v290
	v301 = v292
	v302 = v288
	goto L77
L91:
	;
	goto L90
L92:
	;
	v309 = v300
	v311 = v302
	goto L93
L93:
	;
	v312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v309)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v311)+1)) = uint8(v312)
	v314 = int32(1)
	if v312 != 0 {
		v309 = v309 + v314
		v311 = v311 + v314
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
	v329 = int32(15)
	goto L98
L97:
	;
	v329 = int32(22)
	goto L98
L98:
	;
	v331 = F_errstart(m, v329, int32(0))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L8
	} else {
		goto L99
	}
L99:
	;
	if v331 == int32(0) {
		goto L50
	} else {
		goto L100
	}
L100:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L8
	} else {
		goto L101
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+80)) = v15 + int32(1312)
	F_errmsg(m, int32(_a_F_RestoreArchivedFile_1), v15+int32(80))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L8
	} else {
		goto L102
	}
L102:
	;
	F_errdetail(m, int32(_a_F_RestoreArchivedFile_10), int32(0))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L8
	} else {
		goto L103
	}
L103:
	;
	F_errfinish(m, int32(_a_F_RestoreArchivedFile_2), int32(236), int32(_a_F_RestoreArchivedFile_3))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L8
	} else {
		goto L104
	}
L104:
	;
	goto L50
L105:
	;
	if base.B2i32(int32(15) == v359)&base.B2i32(base.Ui32(v180&int32(_a_F_RestoreArchivedFile_11)-int32(1)) < base.Ui32(v365))|base.B2i32(v359 == int32(0))&base.B2i32(int32(143) == int32(base.Ui32(v180)>>(uint(int32(8))%32))&v365) != 0 {
		goto L1
	} else {
		goto L106
	}
L106:
	;
	v389 = int32(255)
	goto L107
L107:
	;
	if base.B2i32(v180&int32(127) == int32(0))&base.B2i32(base.Ui32(int32(125)) < base.Ui32(int32(base.Ui32(v180)>>(uint(int32(8))%32))&v389))|base.B2i32(base.Ui32(v180&int32(_a_F_RestoreArchivedFile_11)-int32(1)) < base.Ui32(v389)) != 0 {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v401 = int32(22)
	goto L110
L109:
	;
	v401 = int32(13)
	goto L110
L110:
	;
	v403 = F_errstart(m, v401, int32(0))
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L8
	} else {
		goto L111
	}
L111:
	;
	if v403 == int32(0) {
		goto L4
	} else {
		goto L112
	}
L112:
	;
	v407 = F_wait_result_to_str(m, v180)
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L8
	} else {
		goto L113
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v407
	*(*int32)(unsafe.Add(mBase, uint32(v15)+64)) = l1
	F_errmsg(m, int32(_a_F_RestoreArchivedFile_12), v15-int32(-64))
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L8
	} else {
		goto L114
	}
L114:
	;
	F_errfinish(m, int32(_a_F_RestoreArchivedFile_2), int32(269), int32(_a_F_RestoreArchivedFile_3))
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L8
	} else {
		goto L115
	}
L115:
	;
	goto L4
L116:
	;
	v437 = v6
	goto L3
L117:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L8
	} else {
		goto L118
	}
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+144)) = v15 + int32(1312)
	F_errmsg(m, int32(_a_F_RestoreArchivedFile_13), v15+int32(144))
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L8
	} else {
		goto L119
	}
L119:
	;
	F_errfinish(m, int32(_a_F_RestoreArchivedFile_2), int32(120), int32(_a_F_RestoreArchivedFile_3))
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L8
	} else {
		goto L120
	}
L120:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L121:
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
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
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
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
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
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	v4 = int32(0)
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+10)))
	if v8 != 0 {
		v90 = v4
		return v90
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l2)+208))
		if base.Ui32(v10) < base.Ui32(v9) {
			v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+13)))
			v14 = v12
		} else {
			v14 = int32(1)
		}
		if v14&int32(1) == int32(0) {
			v90 = v4
			return v90
		} else {
			v19 = F_IsBinaryTidClause(m, l1, l2)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				if v19 != 0 {
					v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
					if v24 != int32(387) {
						v31 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
						if v31 == int32(20) {
							v34 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
							if v34 != int32(387) {
								v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
								if v72 == int32(0) {
									v90 = v4
								} else {
									v75 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
									v77 = v72
									v80 = v75
									if v80 != int32(58) {
										v90 = v4
									} else {
										v83 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
										v84 = *(*int32)(unsafe.Add(mBase, uint32(l2)+68))
										v90 = base.B2i32(v83 == v84)
									}
								}
								return v90
							} else {
								v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+20)))
								if v37 != int32(1) {
									v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
									if v72 == int32(0) {
										v90 = v4
									} else {
										v75 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
										v77 = v72
										v80 = v75
										if v80 != int32(58) {
											v90 = v4
										} else {
											v83 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
											v84 = *(*int32)(unsafe.Add(mBase, uint32(l2)+68))
											v90 = base.B2i32(v83 == v84)
										}
									}
									return v90
								} else {
									v40 = *(*int32)(unsafe.Add(mBase, uint32(v23)+28))
									v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
									v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
									if v42 == int32(0) {
										v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
										if v72 == int32(0) {
											v90 = v4
										} else {
											v75 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
											v77 = v72
											v80 = v75
											if v80 != int32(58) {
												v90 = v4
											} else {
												v83 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
												v84 = *(*int32)(unsafe.Add(mBase, uint32(l2)+68))
												v90 = base.B2i32(v83 == v84)
											}
										}
										return v90
									} else {
										v45 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
										if v45 != int32(6) {
											v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
											if v72 == int32(0) {
												v90 = v4
											} else {
												v75 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
												v77 = v72
												v80 = v75
												if v80 != int32(58) {
													v90 = v4
												} else {
													v83 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
													v84 = *(*int32)(unsafe.Add(mBase, uint32(l2)+68))
													v90 = base.B2i32(v83 == v84)
												}
											}
											return v90
										} else {
											v48 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v42)+8)))
											if v48 != int32(_a_F_RestrictInfoIsTidQual_0) {
												v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
												if v72 == int32(0) {
													v90 = v4
												} else {
													v75 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
													v77 = v72
													v80 = v75
													if v80 != int32(58) {
														v90 = v4
													} else {
														v83 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
														v84 = *(*int32)(unsafe.Add(mBase, uint32(l2)+68))
														v90 = base.B2i32(v83 == v84)
													}
												}
												return v90
											} else {
												v51 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
												if v51 != int32(27) {
													v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
													if v72 == int32(0) {
														v90 = v4
													} else {
														v75 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
														v77 = v72
														v80 = v75
														if v80 != int32(58) {
															v90 = v4
														} else {
															v83 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
															v84 = *(*int32)(unsafe.Add(mBase, uint32(l2)+68))
															v90 = base.B2i32(v83 == v84)
														}
													}
													return v90
												} else {
													v54 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
													v55 = *(*int32)(unsafe.Add(mBase, uint32(l2)+68))
													if v54 != v55 {
														v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
														if v72 == int32(0) {
															v90 = v4
														} else {
															v75 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
															v77 = v72
															v80 = v75
															if v80 != int32(58) {
																v90 = v4
															} else {
																v83 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
																v84 = *(*int32)(unsafe.Add(mBase, uint32(l2)+68))
																v90 = base.B2i32(v83 == v84)
															}
														}
														return v90
													} else {
														v57 = *(*int32)(unsafe.Add(mBase, uint32(v42)+24))
														if v57 != 0 {
															v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
															if v72 == int32(0) {
																v90 = v4
															} else {
																v75 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
																v77 = v72
																v80 = v75
																if v80 != int32(58) {
																	v90 = v4
																} else {
																	v83 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
																	v84 = *(*int32)(unsafe.Add(mBase, uint32(l2)+68))
																	v90 = base.B2i32(v83 == v84)
																}
															}
															return v90
														} else {
															v58 = *(*int32)(unsafe.Add(mBase, uint32(v42)+28))
															if v58 != 0 {
																v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
																if v72 == int32(0) {
																	v90 = v4
																} else {
																	v75 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
																	v77 = v72
																	v80 = v75
																	if v80 != int32(58) {
																		v90 = v4
																	} else {
																		v83 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
																		v84 = *(*int32)(unsafe.Add(mBase, uint32(l2)+68))
																		v90 = base.B2i32(v83 == v84)
																	}
																}
																return v90
															} else {
																v59 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
																v60 = F_pull_varnos(m, l0, v59)
																mBase = m.M
																v61 = m.ExcPending
																if v61 != 0 {
																	return int32(0)
																} else {
																	v62 = F_bms_is_member(m, v54, v60)
																	mBase = m.M
																	v63 = m.ExcPending
																	if v63 != 0 {
																		return int32(0)
																	} else {
																		if v62 != 0 {
																			v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
																			if v72 == int32(0) {
																				v90 = v4
																			} else {
																				v75 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
																				v77 = v72
																				v80 = v75
																				if v80 != int32(58) {
																					v90 = v4
																				} else {
																					v83 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
																					v84 = *(*int32)(unsafe.Add(mBase, uint32(l2)+68))
																					v90 = base.B2i32(v83 == v84)
																				}
																			}
																			return v90
																		} else {
																			v64 = F_contain_volatile_functions(m, v59)
																			mBase = m.M
																			v65 = m.ExcPending
																			if v65 != 0 {
																				return int32(0)
																			} else {
																				if v64 != 0 {
																					v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
																					if v72 == int32(0) {
																						v90 = v4
																					} else {
																						v75 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
																						v77 = v72
																						v80 = v75
																						if v80 != int32(58) {
																							v90 = v4
																						} else {
																							v83 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
																							v84 = *(*int32)(unsafe.Add(mBase, uint32(l2)+68))
																							v90 = base.B2i32(v83 == v84)
																						}
																					}
																					return v90
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
							v77 = v23
							v80 = v31
							if v80 != int32(58) {
								v90 = v4
							} else {
								v83 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
								v84 = *(*int32)(unsafe.Add(mBase, uint32(l2)+68))
								v90 = base.B2i32(v83 == v84)
							}
							return v90
						}
					} else {
						return int32(1)
					}
				} else {
					if v23 == int32(0) {
						v90 = v4
						return v90
					} else {
						v31 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
						if v31 == int32(20) {
							v34 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
							if v34 != int32(387) {
								v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
								if v72 == int32(0) {
									v90 = v4
								} else {
									v75 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
									v77 = v72
									v80 = v75
									if v80 != int32(58) {
										v90 = v4
									} else {
										v83 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
										v84 = *(*int32)(unsafe.Add(mBase, uint32(l2)+68))
										v90 = base.B2i32(v83 == v84)
									}
								}
								return v90
							} else {
								v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+20)))
								if v37 != int32(1) {
									v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
									if v72 == int32(0) {
										v90 = v4
									} else {
										v75 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
										v77 = v72
										v80 = v75
										if v80 != int32(58) {
											v90 = v4
										} else {
											v83 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
											v84 = *(*int32)(unsafe.Add(mBase, uint32(l2)+68))
											v90 = base.B2i32(v83 == v84)
										}
									}
									return v90
								} else {
									v40 = *(*int32)(unsafe.Add(mBase, uint32(v23)+28))
									v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
									v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
									if v42 == int32(0) {
										v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
										if v72 == int32(0) {
											v90 = v4
										} else {
											v75 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
											v77 = v72
											v80 = v75
											if v80 != int32(58) {
												v90 = v4
											} else {
												v83 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
												v84 = *(*int32)(unsafe.Add(mBase, uint32(l2)+68))
												v90 = base.B2i32(v83 == v84)
											}
										}
										return v90
									} else {
										v45 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
										if v45 != int32(6) {
											v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
											if v72 == int32(0) {
												v90 = v4
											} else {
												v75 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
												v77 = v72
												v80 = v75
												if v80 != int32(58) {
													v90 = v4
												} else {
													v83 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
													v84 = *(*int32)(unsafe.Add(mBase, uint32(l2)+68))
													v90 = base.B2i32(v83 == v84)
												}
											}
											return v90
										} else {
											v48 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v42)+8)))
											if v48 != int32(_a_F_RestrictInfoIsTidQual_0) {
												v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
												if v72 == int32(0) {
													v90 = v4
												} else {
													v75 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
													v77 = v72
													v80 = v75
													if v80 != int32(58) {
														v90 = v4
													} else {
														v83 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
														v84 = *(*int32)(unsafe.Add(mBase, uint32(l2)+68))
														v90 = base.B2i32(v83 == v84)
													}
												}
												return v90
											} else {
												v51 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
												if v51 != int32(27) {
													v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
													if v72 == int32(0) {
														v90 = v4
													} else {
														v75 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
														v77 = v72
														v80 = v75
														if v80 != int32(58) {
															v90 = v4
														} else {
															v83 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
															v84 = *(*int32)(unsafe.Add(mBase, uint32(l2)+68))
															v90 = base.B2i32(v83 == v84)
														}
													}
													return v90
												} else {
													v54 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
													v55 = *(*int32)(unsafe.Add(mBase, uint32(l2)+68))
													if v54 != v55 {
														v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
														if v72 == int32(0) {
															v90 = v4
														} else {
															v75 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
															v77 = v72
															v80 = v75
															if v80 != int32(58) {
																v90 = v4
															} else {
																v83 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
																v84 = *(*int32)(unsafe.Add(mBase, uint32(l2)+68))
																v90 = base.B2i32(v83 == v84)
															}
														}
														return v90
													} else {
														v57 = *(*int32)(unsafe.Add(mBase, uint32(v42)+24))
														if v57 != 0 {
															v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
															if v72 == int32(0) {
																v90 = v4
															} else {
																v75 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
																v77 = v72
																v80 = v75
																if v80 != int32(58) {
																	v90 = v4
																} else {
																	v83 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
																	v84 = *(*int32)(unsafe.Add(mBase, uint32(l2)+68))
																	v90 = base.B2i32(v83 == v84)
																}
															}
															return v90
														} else {
															v58 = *(*int32)(unsafe.Add(mBase, uint32(v42)+28))
															if v58 != 0 {
																v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
																if v72 == int32(0) {
																	v90 = v4
																} else {
																	v75 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
																	v77 = v72
																	v80 = v75
																	if v80 != int32(58) {
																		v90 = v4
																	} else {
																		v83 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
																		v84 = *(*int32)(unsafe.Add(mBase, uint32(l2)+68))
																		v90 = base.B2i32(v83 == v84)
																	}
																}
																return v90
															} else {
																v59 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
																v60 = F_pull_varnos(m, l0, v59)
																mBase = m.M
																v61 = m.ExcPending
																if v61 != 0 {
																	return int32(0)
																} else {
																	v62 = F_bms_is_member(m, v54, v60)
																	mBase = m.M
																	v63 = m.ExcPending
																	if v63 != 0 {
																		return int32(0)
																	} else {
																		if v62 != 0 {
																			v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
																			if v72 == int32(0) {
																				v90 = v4
																			} else {
																				v75 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
																				v77 = v72
																				v80 = v75
																				if v80 != int32(58) {
																					v90 = v4
																				} else {
																					v83 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
																					v84 = *(*int32)(unsafe.Add(mBase, uint32(l2)+68))
																					v90 = base.B2i32(v83 == v84)
																				}
																			}
																			return v90
																		} else {
																			v64 = F_contain_volatile_functions(m, v59)
																			mBase = m.M
																			v65 = m.ExcPending
																			if v65 != 0 {
																				return int32(0)
																			} else {
																				if v64 != 0 {
																					v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
																					if v72 == int32(0) {
																						v90 = v4
																					} else {
																						v75 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
																						v77 = v72
																						v80 = v75
																						if v80 != int32(58) {
																							v90 = v4
																						} else {
																							v83 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
																							v84 = *(*int32)(unsafe.Add(mBase, uint32(l2)+68))
																							v90 = base.B2i32(v83 == v84)
																						}
																					}
																					return v90
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
							v77 = v23
							v80 = v31
							if v80 != int32(58) {
								v90 = v4
							} else {
								v83 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
								v84 = *(*int32)(unsafe.Add(mBase, uint32(l2)+68))
								v90 = base.B2i32(v83 == v84)
							}
							return v90
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
	*(*int32)(unsafe.Add(mBase, _c_F_RoleidCallback[0])) = int32(0)
	return
}
func F_r_CONSONANT(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v14 <= v15 {
		v123 = int32(-1)
		v130 = v123
	} else {
		v32 = int32(1)
		v33 = v14 - v32
		v35 = int32(*(*int8)(unsafe.Add(mBase, uint32(v16+v33))))
		v37 = v35 & int32(255)
		if base.B2i32(v33 == v15)|base.B2i32(int32(0) <= v35) != 0 {
			v95 = v37
			v99 = v32
		} else {
			v44 = v37 & int32(63)
			v46 = v14 - int32(2)
			v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16+v46))))
			v50 = v48 << (uint(int32(6)) % 32)
			if base.B2i32(v46 != v15)&base.B2i32(base.Ui32(v48) < base.Ui32(int32(192))) == int32(0) {
				v95 = v50&int32(1984) | v44
				v99 = int32(2)
			} else {
				v63 = v50&int32(4032) | v44
				v65 = v14 - int32(3)
				v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16+v65))))
				if base.B2i32(v65 != v15)&base.B2i32(base.Ui32(v67) < base.Ui32(int32(224))) == int32(0) {
					v95 = v67<<(uint(int32(12))%32)&int32(_a_F_r_CONSONANT_0) | v63
					v99 = int32(3)
				} else {
					v85 = int32(4)
					v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14+v16-v85))))
					v95 = v67<<(uint(int32(12))%32)&int32(_a_F_r_CONSONANT_1) | v87&int32(7)<<(uint(int32(18))%32) | v63
					v99 = v85
				}
			}
		}
		if int32(2399) < v95 {
			v130 = v99
		} else {
			v101 = v95 - int32(2325)
			if v101 < int32(0) {
				v130 = v99
			} else {
				v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v101)>>(uint(int32(3))%32)))+uint32(_c_F_r_CONSONANT[0]))))
				if int32(base.Ui32(v107)>>(uint(v101&int32(7))%32))&int32(1) == int32(0) {
					v130 = v99
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v14 - v99
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
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v74 int32
	_ = v74
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v111 int32
	_ = v111
	var v118 int32
	_ = v118
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
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v14 <= v13 {
		v118 = int32(-1)
	} else {
		v30 = int32(1)
		v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13+v15))))
		if base.Ui32(v32) < base.Ui32(int32(192)) {
			v89 = v32
			v90 = v30
		} else {
			v36 = v13 + int32(1)
			if v36 == v14 {
				v89 = v32
				v90 = v30
			} else {
				v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36+v15))))
				v41 = v39 & int32(63)
				if base.Ui32(int32(224)) <= base.Ui32(v32) {
					v45 = v13 + int32(2)
					if v45 != v14 {
						v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45+v15))))
						v57 = v55 & int32(63)
						if base.Ui32(int32(240)) <= base.Ui32(v32) {
							v61 = v13 + int32(3)
							if v61 != v14 {
								v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15+v61))))
								v89 = v74&int32(63) | (v32<<(uint(int32(18))%32)&int32(_a_F_r_KER_2_0) | v41<<(uint(int32(12))%32) | v57<<(uint(int32(6))%32))
								v90 = int32(4)
							} else {
								v89 = v32<<(uint(int32(12))%32)&int32(_a_F_r_KER_2_1) | v41<<(uint(int32(6))%32) | v57
								v90 = int32(3)
							}
						} else {
							v89 = v32<<(uint(int32(12))%32)&int32(_a_F_r_KER_2_1) | v41<<(uint(int32(6))%32) | v57
							v90 = int32(3)
						}
					} else {
						v89 = v32<<(uint(int32(6))%32)&int32(1984) | v41
						v90 = int32(2)
					}
				} else {
					v89 = v32<<(uint(int32(6))%32)&int32(1984) | v41
					v90 = int32(2)
				}
			}
		}
		if int32(117) < v89 {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v90 + v13
			v111 = int32(0)
		} else {
			v94 = v89 - int32(97)
			if v94 < int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v90 + v13
				v111 = int32(0)
			} else {
				v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v94)>>(uint(int32(3))%32)))+uint32(_c_F_r_KER_2[0]))))
				if int32(base.Ui32(v100)>>(uint(v94&int32(7))%32))&int32(1) != 0 {
					v111 = v90
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v90 + v13
					v111 = int32(0)
				}
			}
		}
		v118 = v111
	}
	if v118 != 0 {
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
			v130 = F_memcmp(m, v128+v125, int32(_a_F_r_KER_2_2), v120)
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
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	v2 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v4
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v4-int32(3) <= v6 {
		v119 = v2
		return v119
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10+v4-int32(1)))))
		if v14 != int32(170) {
			v119 = v2
			return v119
		} else {
			v19 = F_find_among_b(m, l0, int32(_a_F_r_Suffix_Noun_Step2b_0), int32(1))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				if v19 == int32(0) {
					v119 = v2
					return v119
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v25
					v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v28 = int32(0)
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v27-int32(4))))
					if v35 == v28 {
						v109 = int32(0)
					} else {
						v40 = v35 & int32(3)
						if base.Ui32(v35) < base.Ui32(int32(4)) {
							v76 = v27
							v77 = int32(0)
							v82 = v76
							v83 = v77
							v87 = v28
							for {
								v88 = int32(*(*int8)(unsafe.Add(mBase, uint32(v82))))
								v91 = v83 + base.B2i32(int32(-65) < v88)
								v92 = int32(1)
								v95 = v87 + v92
								if v95 != v40 {
									v82 = v82 + v92
									v83 = v91
									v87 = v95
									continue
								} else {
									break
								}
								break
							}
							v98 = v91
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
							if v40 == int32(0) {
								v98 = v68
							} else {
								v76 = v70
								v77 = v68
								v82 = v76
								v83 = v77
								v87 = v28
								for {
									v88 = int32(*(*int8)(unsafe.Add(mBase, uint32(v82))))
									v91 = v83 + base.B2i32(int32(-65) < v88)
									v92 = int32(1)
									v95 = v87 + v92
									if v95 != v40 {
										v82 = v82 + v92
										v83 = v91
										v87 = v95
										continue
									} else {
										break
									}
									break
								}
								v98 = v91
							}
						}
						v109 = v98
					}
					if v109 < int32(5) {
						v119 = v2
						return v119
					} else {
						v113 = F_slice_del(m, l0)
						mBase = m.M
						v114 = m.ExcPending
						if v114 != 0 {
							return int32(0)
						} else {
							if int32(0) <= v113 {
								v117 = int32(1)
							} else {
								v117 = v113
							}
							v119 = v117
							return v119
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
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
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
	v7 = v4 - int32(1)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v7 <= v8 {
		v117 = v2
		return v117
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10+v7))))
		if v12 != int32(170) {
			v117 = v2
			return v117
		} else {
			v17 = F_find_among_b(m, l0, int32(_a_F_r_Suffix_Noun_Step2c1_0), int32(1))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				if v17 == int32(0) {
					v117 = v2
					return v117
				} else {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v23
					v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v26 = int32(0)
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v25-int32(4))))
					if v33 == v26 {
						v107 = int32(0)
					} else {
						v38 = v33 & int32(3)
						if base.Ui32(v33) < base.Ui32(int32(4)) {
							v74 = v25
							v75 = int32(0)
							v80 = v74
							v81 = v75
							v85 = v26
							for {
								v86 = int32(*(*int8)(unsafe.Add(mBase, uint32(v80))))
								v89 = v81 + base.B2i32(int32(-65) < v86)
								v90 = int32(1)
								v93 = v85 + v90
								if v93 != v38 {
									v80 = v80 + v90
									v81 = v89
									v85 = v93
									continue
								} else {
									break
								}
								break
							}
							v96 = v89
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
							if v38 == int32(0) {
								v96 = v66
							} else {
								v74 = v68
								v75 = v66
								v80 = v74
								v81 = v75
								v85 = v26
								for {
									v86 = int32(*(*int8)(unsafe.Add(mBase, uint32(v80))))
									v89 = v81 + base.B2i32(int32(-65) < v86)
									v90 = int32(1)
									v93 = v85 + v90
									if v93 != v38 {
										v80 = v80 + v90
										v81 = v89
										v85 = v93
										continue
									} else {
										break
									}
									break
								}
								v96 = v89
							}
						}
						v107 = v96
					}
					if v107 < int32(4) {
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
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v293 int32
	_ = v293
	var v298 int32
	_ = v298
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v376 int32
	_ = v376
	v2 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v4
	v8 = F_find_among_b(m, l0, int32(_a_F_r_Suffix_Verb_Step2a_0), int32(11))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		if v8 == int32(0) {
			v376 = v2
			return v376
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v14
			switch v8 - int32(1) {
			case 0:
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v19 = int32(0)
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v18-int32(4))))
				if v26 == v19 {
					v100 = int32(0)
				} else {
					v31 = v26 & int32(3)
					if base.Ui32(v26) < base.Ui32(int32(4)) {
						v67 = v18
						v68 = int32(0)
						v73 = v67
						v74 = v68
						v78 = v19
						for {
							v79 = int32(*(*int8)(unsafe.Add(mBase, uint32(v73))))
							v82 = v74 + base.B2i32(int32(-65) < v79)
							v83 = int32(1)
							v86 = v78 + v83
							if v86 != v31 {
								v73 = v73 + v83
								v74 = v82
								v78 = v86
								continue
							} else {
								break
							}
							break
						}
						v89 = v82
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
						if v31 == int32(0) {
							v89 = v59
						} else {
							v67 = v61
							v68 = v59
							v73 = v67
							v74 = v68
							v78 = v19
							for {
								v79 = int32(*(*int8)(unsafe.Add(mBase, uint32(v73))))
								v82 = v74 + base.B2i32(int32(-65) < v79)
								v83 = int32(1)
								v86 = v78 + v83
								if v86 != v31 {
									v73 = v73 + v83
									v74 = v82
									v78 = v86
									continue
								} else {
									break
								}
								break
							}
							v89 = v82
						}
					}
					v100 = v89
				}
				if v100 < int32(4) {
					v376 = v2
					return v376
				} else {
					v103 = F_slice_del(m, l0)
					mBase = m.M
					v104 = m.ExcPending
					if v104 != 0 {
						return int32(0)
					} else {
						if int32(0) <= v103 {
							v376 = int32(1)
						} else {
							v376 = v103
						}
						return v376
					}
				}
			case 1:
				v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v108 = int32(0)
				v115 = *(*int32)(unsafe.Add(mBase, uint32(v107-int32(4))))
				if v115 == v108 {
					v189 = int32(0)
				} else {
					v120 = v115 & int32(3)
					if base.Ui32(v115) < base.Ui32(int32(4)) {
						v156 = v107
						v157 = int32(0)
						v162 = v156
						v163 = v157
						v167 = v108
						for {
							v168 = int32(*(*int8)(unsafe.Add(mBase, uint32(v162))))
							v171 = v163 + base.B2i32(int32(-65) < v168)
							v172 = int32(1)
							v175 = v167 + v172
							if v175 != v120 {
								v162 = v162 + v172
								v163 = v171
								v167 = v175
								continue
							} else {
								break
							}
							break
						}
						v178 = v171
					} else {
						v127 = v107
						v128 = int32(0)
						v131 = v108
						for {
							v133 = int32(*(*int8)(unsafe.Add(mBase, uint32(v127))))
							v134 = int32(-65)
							v137 = int32(*(*int8)(unsafe.Add(mBase, uint32(v127)+1)))
							v141 = int32(*(*int8)(unsafe.Add(mBase, uint32(v127)+2)))
							v145 = int32(*(*int8)(unsafe.Add(mBase, uint32(v127)+3)))
							v148 = v128 + base.B2i32(v134 < v133) + base.B2i32(v134 < v137) + base.B2i32(v134 < v141) + base.B2i32(v134 < v145)
							v149 = int32(4)
							v150 = v127 + v149
							v152 = v131 + v149
							if v152 != v115&int32(-4) {
								v127 = v150
								v128 = v148
								v131 = v152
								continue
							} else {
								break
							}
							break
						}
						if v120 == int32(0) {
							v178 = v148
						} else {
							v156 = v150
							v157 = v148
							v162 = v156
							v163 = v157
							v167 = v108
							for {
								v168 = int32(*(*int8)(unsafe.Add(mBase, uint32(v162))))
								v171 = v163 + base.B2i32(int32(-65) < v168)
								v172 = int32(1)
								v175 = v167 + v172
								if v175 != v120 {
									v162 = v162 + v172
									v163 = v171
									v167 = v175
									continue
								} else {
									break
								}
								break
							}
							v178 = v171
						}
					}
					v189 = v178
				}
				if v189 < int32(5) {
					v376 = v2
					return v376
				} else {
					v192 = F_slice_del(m, l0)
					mBase = m.M
					v193 = m.ExcPending
					if v193 != 0 {
						return int32(0)
					} else {
						if int32(0) <= v192 {
							v376 = int32(1)
						} else {
							v376 = v192
						}
						return v376
					}
				}
			case 2:
				v196 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v197 = int32(0)
				v204 = *(*int32)(unsafe.Add(mBase, uint32(v196-int32(4))))
				if v204 == v197 {
					v278 = int32(0)
				} else {
					v209 = v204 & int32(3)
					if base.Ui32(v204) < base.Ui32(int32(4)) {
						v245 = v196
						v246 = int32(0)
						v251 = v245
						v252 = v246
						v256 = v197
						for {
							v257 = int32(*(*int8)(unsafe.Add(mBase, uint32(v251))))
							v260 = v252 + base.B2i32(int32(-65) < v257)
							v261 = int32(1)
							v264 = v256 + v261
							if v264 != v209 {
								v251 = v251 + v261
								v252 = v260
								v256 = v264
								continue
							} else {
								break
							}
							break
						}
						v267 = v260
					} else {
						v216 = v196
						v217 = int32(0)
						v220 = v197
						for {
							v222 = int32(*(*int8)(unsafe.Add(mBase, uint32(v216))))
							v223 = int32(-65)
							v226 = int32(*(*int8)(unsafe.Add(mBase, uint32(v216)+1)))
							v230 = int32(*(*int8)(unsafe.Add(mBase, uint32(v216)+2)))
							v234 = int32(*(*int8)(unsafe.Add(mBase, uint32(v216)+3)))
							v237 = v217 + base.B2i32(v223 < v222) + base.B2i32(v223 < v226) + base.B2i32(v223 < v230) + base.B2i32(v223 < v234)
							v238 = int32(4)
							v239 = v216 + v238
							v241 = v220 + v238
							if v241 != v204&int32(-4) {
								v216 = v239
								v217 = v237
								v220 = v241
								continue
							} else {
								break
							}
							break
						}
						if v209 == int32(0) {
							v267 = v237
						} else {
							v245 = v239
							v246 = v237
							v251 = v245
							v252 = v246
							v256 = v197
							for {
								v257 = int32(*(*int8)(unsafe.Add(mBase, uint32(v251))))
								v260 = v252 + base.B2i32(int32(-65) < v257)
								v261 = int32(1)
								v264 = v256 + v261
								if v264 != v209 {
									v251 = v251 + v261
									v252 = v260
									v256 = v264
									continue
								} else {
									break
								}
								break
							}
							v267 = v260
						}
					}
					v278 = v267
				}
				if v278 < int32(6) {
					v376 = v2
					return v376
				} else {
					v281 = F_slice_del(m, l0)
					mBase = m.M
					v282 = m.ExcPending
					if v282 != 0 {
						return int32(0)
					} else {
						if int32(0) <= v281 {
							v376 = int32(1)
						} else {
							v376 = v281
						}
						return v376
					}
				}
			case 3:
				v285 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v286 = int32(0)
				v293 = *(*int32)(unsafe.Add(mBase, uint32(v285-int32(4))))
				if v293 == v286 {
					v367 = int32(0)
				} else {
					v298 = v293 & int32(3)
					if base.Ui32(v293) < base.Ui32(int32(4)) {
						v334 = v285
						v335 = int32(0)
						v340 = v334
						v341 = v335
						v345 = v286
						for {
							v346 = int32(*(*int8)(unsafe.Add(mBase, uint32(v340))))
							v349 = v341 + base.B2i32(int32(-65) < v346)
							v350 = int32(1)
							v353 = v345 + v350
							if v353 != v298 {
								v340 = v340 + v350
								v341 = v349
								v345 = v353
								continue
							} else {
								break
							}
							break
						}
						v356 = v349
					} else {
						v305 = v285
						v306 = int32(0)
						v309 = v286
						for {
							v311 = int32(*(*int8)(unsafe.Add(mBase, uint32(v305))))
							v312 = int32(-65)
							v315 = int32(*(*int8)(unsafe.Add(mBase, uint32(v305)+1)))
							v319 = int32(*(*int8)(unsafe.Add(mBase, uint32(v305)+2)))
							v323 = int32(*(*int8)(unsafe.Add(mBase, uint32(v305)+3)))
							v326 = v306 + base.B2i32(v312 < v311) + base.B2i32(v312 < v315) + base.B2i32(v312 < v319) + base.B2i32(v312 < v323)
							v327 = int32(4)
							v328 = v305 + v327
							v330 = v309 + v327
							if v330 != v293&int32(-4) {
								v305 = v328
								v306 = v326
								v309 = v330
								continue
							} else {
								break
							}
							break
						}
						if v298 == int32(0) {
							v356 = v326
						} else {
							v334 = v328
							v335 = v326
							v340 = v334
							v341 = v335
							v345 = v286
							for {
								v346 = int32(*(*int8)(unsafe.Add(mBase, uint32(v340))))
								v349 = v341 + base.B2i32(int32(-65) < v346)
								v350 = int32(1)
								v353 = v345 + v350
								if v353 != v298 {
									v340 = v340 + v350
									v341 = v349
									v345 = v353
									continue
								} else {
									break
								}
								break
							}
							v356 = v349
						}
					}
					v367 = v356
				}
				if v367 < int32(6) {
					v376 = v2
					return v376
				} else {
					v370 = F_slice_del(m, l0)
					mBase = m.M
					v371 = m.ExcPending
					if v371 != 0 {
						return int32(0)
					} else {
						if v370 < int32(0) {
							v376 = v370
						} else {
							v376 = int32(1)
						}
						return v376
					}
				}
			default:
				v376 = int32(1)
				return v376
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
	var v24 int32
	_ = v24
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
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
	var v170 int32
	_ = v170
	var v179 int32
	_ = v179
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
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
	var v320 int32
	_ = v320
	var v329 int32
	_ = v329
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v399 int32
	_ = v399
	var v403 int32
	_ = v403
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
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v481 int32
	_ = v481
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v523 int32
	_ = v523
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v551 int32
	_ = v551
	var v555 int32
	_ = v555
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
	var v619 int32
	_ = v619
	var v628 int32
	_ = v628
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v638 int32
	_ = v638
	var v640 int32
	_ = v640
	var v647 int32
	_ = v647
	var v649 int32
	_ = v649
	var v651 int32
	_ = v651
	var v653 int32
	_ = v653
	var v666 int32
	_ = v666
	var v668 int32
	_ = v668
	var v670 int32
	_ = v670
	var v688 int32
	_ = v688
	var v690 int32
	_ = v690
	var v698 int32
	_ = v698
	var v702 int32
	_ = v702
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
	var v738 int32
	_ = v738
	var v740 int32
	_ = v740
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v747 int32
	_ = v747
	var v751 int32
	_ = v751
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v776 int32
	_ = v776
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v786 int32
	_ = v786
	var v788 int32
	_ = v788
	var v795 int32
	_ = v795
	var v797 int32
	_ = v797
	var v799 int32
	_ = v799
	var v801 int32
	_ = v801
	var v814 int32
	_ = v814
	var v816 int32
	_ = v816
	var v818 int32
	_ = v818
	var v836 int32
	_ = v836
	var v838 int32
	_ = v838
	var v846 int32
	_ = v846
	var v850 int32
	_ = v850
	var v852 int32
	_ = v852
	var v858 int32
	_ = v858
	var v867 int32
	_ = v867
	var v882 int32
	_ = v882
	var v885 int32
	_ = v885
	var v887 int32
	_ = v887
	var v889 int32
	_ = v889
	var v891 int32
	_ = v891
	var v894 int32
	_ = v894
	var v897 int32
	_ = v897
	var v900 int32
	_ = v900
	var v904 int32
	_ = v904
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v928 int32
	_ = v928
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v938 int32
	_ = v938
	var v940 int32
	_ = v940
	var v947 int32
	_ = v947
	var v949 int32
	_ = v949
	var v951 int32
	_ = v951
	var v953 int32
	_ = v953
	var v966 int32
	_ = v966
	var v968 int32
	_ = v968
	var v970 int32
	_ = v970
	var v988 int32
	_ = v988
	var v990 int32
	_ = v990
	var v998 int32
	_ = v998
	var v1002 int32
	_ = v1002
	var v1004 int32
	_ = v1004
	var v1010 int32
	_ = v1010
	var v1019 int32
	_ = v1019
	var v1034 int32
	_ = v1034
	var v1037 int32
	_ = v1037
	var v1038 int32
	_ = v1038
	var v1040 int32
	_ = v1040
	var v1042 int32
	_ = v1042
	var v1046 int32
	_ = v1046
	var v1050 int32
	_ = v1050
	var v1065 int32
	_ = v1065
	var v1066 int32
	_ = v1066
	var v1075 int32
	_ = v1075
	var v1082 int32
	_ = v1082
	var v1083 int32
	_ = v1083
	var v1085 int32
	_ = v1085
	var v1087 int32
	_ = v1087
	var v1094 int32
	_ = v1094
	var v1096 int32
	_ = v1096
	var v1098 int32
	_ = v1098
	var v1100 int32
	_ = v1100
	var v1113 int32
	_ = v1113
	var v1115 int32
	_ = v1115
	var v1117 int32
	_ = v1117
	var v1135 int32
	_ = v1135
	var v1137 int32
	_ = v1137
	var v1145 int32
	_ = v1145
	var v1149 int32
	_ = v1149
	var v1151 int32
	_ = v1151
	var v1157 int32
	_ = v1157
	var v1166 int32
	_ = v1166
	var v1181 int32
	_ = v1181
	var v1184 int32
	_ = v1184
	var v1186 int32
	_ = v1186
	var v1188 int32
	_ = v1188
	var v1190 int32
	_ = v1190
	var v1193 int32
	_ = v1193
	var v1196 int32
	_ = v1196
	var v1199 int32
	_ = v1199
	var v1203 int32
	_ = v1203
	var v1218 int32
	_ = v1218
	var v1219 int32
	_ = v1219
	var v1220 int32
	_ = v1220
	var v1229 int32
	_ = v1229
	var v1236 int32
	_ = v1236
	var v1237 int32
	_ = v1237
	var v1239 int32
	_ = v1239
	var v1241 int32
	_ = v1241
	var v1248 int32
	_ = v1248
	var v1250 int32
	_ = v1250
	var v1252 int32
	_ = v1252
	var v1254 int32
	_ = v1254
	var v1267 int32
	_ = v1267
	var v1269 int32
	_ = v1269
	var v1271 int32
	_ = v1271
	var v1289 int32
	_ = v1289
	var v1291 int32
	_ = v1291
	var v1299 int32
	_ = v1299
	var v1303 int32
	_ = v1303
	var v1305 int32
	_ = v1305
	var v1311 int32
	_ = v1311
	var v1320 int32
	_ = v1320
	var v1335 int32
	_ = v1335
	var v1341 int32
	_ = v1341
	var v1349 int32
	_ = v1349
	v2 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v33 = v9
	goto L4
L1:
	;
	return v1349
L2:
	;
	if v139 < int32(0) {
		v1349 = v2
		goto L1
	} else {
		goto L20
	}
L3:
	;
	v139 = int32(-1)
	goto L2
L4:
	;
	if v33 <= v23 {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v40 = int32(1)
	v41 = v33 - v40
	v43 = int32(*(*int8)(unsafe.Add(mBase, uint32(v24+v41))))
	v45 = v43 & int32(255)
	if base.B2i32(v41 == v23)|base.B2i32(int32(0) <= v43) != 0 {
		v103 = v45
		v107 = v40
		goto L7
	} else {
		goto L8
	}
L7:
	;
	if int32(305) < v103 {
		goto L15
	} else {
		goto L16
	}
L8:
	;
	v52 = v45 & int32(63)
	v54 = v33 - int32(2)
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+v54))))
	v58 = v56 << (uint(int32(6)) % 32)
	if base.B2i32(v54 != v23)&base.B2i32(base.Ui32(v56) < base.Ui32(int32(192))) == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v103 = v58&int32(1984) | v52
	v107 = int32(2)
	goto L7
L10:
	;
	goto L11
L11:
	;
	v71 = v58&int32(4032) | v52
	v73 = v33 - int32(3)
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+v73))))
	if base.B2i32(v73 != v23)&base.B2i32(base.Ui32(v75) < base.Ui32(int32(224))) == int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v103 = v75<<(uint(int32(12))%32)&int32(_a_F_r_check_vowel_harmony_0) | v71
	v107 = int32(3)
	goto L7
L13:
	;
	goto L14
L14:
	;
	v93 = int32(4)
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33+v24-v93))))
	v103 = v75<<(uint(int32(12))%32)&int32(_a_F_r_check_vowel_harmony_1) | v95&int32(7)<<(uint(int32(18))%32) | v71
	v107 = v93
	goto L7
L15:
	;
	v124 = v33 - v107
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v124
	v33 = v124
	goto L4
L16:
	;
	v109 = v103 - int32(97)
	if v109 < int32(0) {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v109)>>(uint(int32(3))%32)))+uint32(_c_F_r_check_vowel_harmony[0]))))
	if int32(base.Ui32(v115)>>(uint(v109&int32(7))%32))&int32(1) == int32(0) {
		goto L15
	} else {
		goto L18
	}
L18:
	;
	v139 = v107
	goto L2
L20:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v143 <= v144 {
		v290 = v144
		v291 = v142
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v1341 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1341 + (v9 - v8)
	v1349 = int32(1)
	goto L1
L22:
	;
	v292 = v142 - v143
	v293 = v291 - v292
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v293
	if v293 <= v290 {
		v440 = v293
		goto L44
	} else {
		goto L45
	}
L23:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146+v143-int32(1)))))
	if v150 != int32(97) {
		v290 = v144
		v291 = v142
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v154 = v143 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v154
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v179 = v154
	goto L27
L25:
	;
	if int32(0) <= v285 {
		goto L21
	} else {
		goto L43
	}
L26:
	;
	v285 = int32(-1)
	goto L25
L27:
	;
	if v179 <= v169 {
		goto L26
	} else {
		goto L29
	}
L29:
	;
	v186 = int32(1)
	v187 = v179 - v186
	v189 = int32(*(*int8)(unsafe.Add(mBase, uint32(v170+v187))))
	v191 = v189 & int32(255)
	if base.B2i32(v187 == v169)|base.B2i32(int32(0) <= v189) != 0 {
		v249 = v191
		v253 = v186
		goto L30
	} else {
		goto L31
	}
L30:
	;
	if int32(305) < v249 {
		goto L38
	} else {
		goto L39
	}
L31:
	;
	v198 = v191 & int32(63)
	v200 = v179 - int32(2)
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170+v200))))
	v204 = v202 << (uint(int32(6)) % 32)
	if base.B2i32(v200 != v169)&base.B2i32(base.Ui32(v202) < base.Ui32(int32(192))) == int32(0) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v249 = v204&int32(1984) | v198
	v253 = int32(2)
	goto L30
L33:
	;
	goto L34
L34:
	;
	v217 = v204&int32(4032) | v198
	v219 = v179 - int32(3)
	v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170+v219))))
	if base.B2i32(v219 != v169)&base.B2i32(base.Ui32(v221) < base.Ui32(int32(224))) == int32(0) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v249 = v221<<(uint(int32(12))%32)&int32(_a_F_r_check_vowel_harmony_0) | v217
	v253 = int32(3)
	goto L30
L36:
	;
	goto L37
L37:
	;
	v239 = int32(4)
	v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179+v170-v239))))
	v249 = v221<<(uint(int32(12))%32)&int32(_a_F_r_check_vowel_harmony_1) | v241&int32(7)<<(uint(int32(18))%32) | v217
	v253 = v239
	goto L30
L38:
	;
	v270 = v179 - v253
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v270
	v179 = v270
	goto L27
L39:
	;
	v255 = v249 - int32(97)
	if v255 < int32(0) {
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v255)>>(uint(int32(3))%32)))+uint32(_c_F_r_check_vowel_harmony[1]))))
	if int32(base.Ui32(v261)>>(uint(v255&int32(7))%32))&int32(1) == int32(0) {
		goto L38
	} else {
		goto L41
	}
L41:
	;
	v285 = v253
	goto L25
L43:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v289 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v290 = v288
	v291 = v289
	goto L22
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v440
	v442 = int32(2)
	v444 = int32(0)
	v447 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v440-v447 < v442 {
		v457 = v444
		goto L67
	} else {
		goto L68
	}
L45:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v296+v293-int32(1)))))
	if v300 != int32(101) {
		v440 = v293
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v304 = v293 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v304
	v319 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v320 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v329 = v304
	goto L49
L47:
	;
	if int32(0) <= v435 {
		goto L21
	} else {
		goto L65
	}
L48:
	;
	v435 = int32(-1)
	goto L47
L49:
	;
	if v329 <= v319 {
		goto L48
	} else {
		goto L51
	}
L51:
	;
	v336 = int32(1)
	v337 = v329 - v336
	v339 = int32(*(*int8)(unsafe.Add(mBase, uint32(v320+v337))))
	v341 = v339 & int32(255)
	if base.B2i32(v337 == v319)|base.B2i32(int32(0) <= v339) != 0 {
		v399 = v341
		v403 = v336
		goto L52
	} else {
		goto L53
	}
L52:
	;
	if int32(252) < v399 {
		goto L60
	} else {
		goto L61
	}
L53:
	;
	v348 = v341 & int32(63)
	v350 = v329 - int32(2)
	v352 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v320+v350))))
	v354 = v352 << (uint(int32(6)) % 32)
	if base.B2i32(v350 != v319)&base.B2i32(base.Ui32(v352) < base.Ui32(int32(192))) == int32(0) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v399 = v354&int32(1984) | v348
	v403 = int32(2)
	goto L52
L55:
	;
	goto L56
L56:
	;
	v367 = v354&int32(4032) | v348
	v369 = v329 - int32(3)
	v371 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v320+v369))))
	if base.B2i32(v369 != v319)&base.B2i32(base.Ui32(v371) < base.Ui32(int32(224))) == int32(0) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v399 = v371<<(uint(int32(12))%32)&int32(_a_F_r_check_vowel_harmony_0) | v367
	v403 = int32(3)
	goto L52
L58:
	;
	goto L59
L59:
	;
	v389 = int32(4)
	v391 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v329+v320-v389))))
	v399 = v371<<(uint(int32(12))%32)&int32(_a_F_r_check_vowel_harmony_1) | v391&int32(7)<<(uint(int32(18))%32) | v367
	v403 = v389
	goto L52
L60:
	;
	v420 = v329 - v403
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v420
	v329 = v420
	goto L49
L61:
	;
	v405 = v399 - int32(101)
	if v405 < int32(0) {
		goto L60
	} else {
		goto L62
	}
L62:
	;
	v411 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v405)>>(uint(int32(3))%32)))+uint32(_c_F_r_check_vowel_harmony[2]))))
	if int32(base.Ui32(v411)>>(uint(v405&int32(7))%32))&int32(1) == int32(0) {
		goto L60
	} else {
		goto L63
	}
L63:
	;
	v435 = v403
	goto L47
L65:
	;
	v438 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v440 = v438 - v292
	goto L44
L66:
	;
	if v457 != 0 {
		goto L70
	} else {
		goto L71
	}
L67:
	;
	goto L66
L68:
	;
	v450 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v453 = F_memcmp(m, v450+v440-v442, int32(_a_F_r_check_vowel_harmony_2), v442)
	mBase = m.M
	if v453 != 0 {
		v457 = v444
		goto L67
	} else {
		goto L69
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v440 - v442
	v457 = int32(1)
	goto L67
L70:
	;
	v470 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v471 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v472 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v481 = v470
	goto L75
L71:
	;
	goto L72
L72:
	;
	v590 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v591 = v590 - v292
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v591
	v593 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v591 <= v593 {
		v887 = v591
		goto L92
	} else {
		goto L93
	}
L73:
	;
	if int32(0) <= v587 {
		goto L21
	} else {
		goto L91
	}
L74:
	;
	v587 = int32(-1)
	goto L73
L75:
	;
	if v481 <= v471 {
		goto L74
	} else {
		goto L77
	}
L77:
	;
	v488 = int32(1)
	v489 = v481 - v488
	v491 = int32(*(*int8)(unsafe.Add(mBase, uint32(v472+v489))))
	v493 = v491 & int32(255)
	if base.B2i32(v489 == v471)|base.B2i32(int32(0) <= v491) != 0 {
		v551 = v493
		v555 = v488
		goto L78
	} else {
		goto L79
	}
L78:
	;
	if int32(305) < v551 {
		goto L86
	} else {
		goto L87
	}
L79:
	;
	v500 = v493 & int32(63)
	v502 = v481 - int32(2)
	v504 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v472+v502))))
	v506 = v504 << (uint(int32(6)) % 32)
	if base.B2i32(v502 != v471)&base.B2i32(base.Ui32(v504) < base.Ui32(int32(192))) == int32(0) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v551 = v506&int32(1984) | v500
	v555 = int32(2)
	goto L78
L81:
	;
	goto L82
L82:
	;
	v519 = v506&int32(4032) | v500
	v521 = v481 - int32(3)
	v523 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v472+v521))))
	if base.B2i32(v521 != v471)&base.B2i32(base.Ui32(v523) < base.Ui32(int32(224))) == int32(0) {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v551 = v523<<(uint(int32(12))%32)&int32(_a_F_r_check_vowel_harmony_0) | v519
	v555 = int32(3)
	goto L78
L84:
	;
	goto L85
L85:
	;
	v541 = int32(4)
	v543 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v481+v472-v541))))
	v551 = v523<<(uint(int32(12))%32)&int32(_a_F_r_check_vowel_harmony_1) | v543&int32(7)<<(uint(int32(18))%32) | v519
	v555 = v541
	goto L78
L86:
	;
	v572 = v481 - v555
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v572
	v481 = v572
	goto L75
L87:
	;
	v557 = v551 - int32(97)
	if v557 < int32(0) {
		goto L86
	} else {
		goto L88
	}
L88:
	;
	v563 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v557)>>(uint(int32(3))%32)))+uint32(_c_F_r_check_vowel_harmony[3]))))
	if int32(base.Ui32(v563)>>(uint(v557&int32(7))%32))&int32(1) == int32(0) {
		goto L86
	} else {
		goto L89
	}
L89:
	;
	v587 = v555
	goto L73
L91:
	;
	goto L72
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v887
	v889 = int32(2)
	v891 = int32(0)
	v894 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v887-v894 < v889 {
		v904 = v891
		goto L138
	} else {
		goto L139
	}
L93:
	;
	v595 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v599 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v595+v591-int32(1)))))
	if v599 == int32(105) {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v603 = v591 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v603
	v618 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v619 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v628 = v603
	goto L99
L95:
	;
	v742 = v591
	goto L96
L96:
	;
	v743 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v747 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v743+v742-int32(1)))))
	if v747 != int32(111) {
		v887 = v742
		goto L92
	} else {
		goto L117
	}
L97:
	;
	if int32(0) <= v734 {
		goto L21
	} else {
		goto L115
	}
L98:
	;
	v734 = int32(-1)
	goto L97
L99:
	;
	if v628 <= v618 {
		goto L98
	} else {
		goto L101
	}
L101:
	;
	v635 = int32(1)
	v636 = v628 - v635
	v638 = int32(*(*int8)(unsafe.Add(mBase, uint32(v619+v636))))
	v640 = v638 & int32(255)
	if base.B2i32(v636 == v618)|base.B2i32(int32(0) <= v638) != 0 {
		v698 = v640
		v702 = v635
		goto L102
	} else {
		goto L103
	}
L102:
	;
	if int32(105) < v698 {
		goto L110
	} else {
		goto L111
	}
L103:
	;
	v647 = v640 & int32(63)
	v649 = v628 - int32(2)
	v651 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v619+v649))))
	v653 = v651 << (uint(int32(6)) % 32)
	if base.B2i32(v649 != v618)&base.B2i32(base.Ui32(v651) < base.Ui32(int32(192))) == int32(0) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v698 = v653&int32(1984) | v647
	v702 = int32(2)
	goto L102
L105:
	;
	goto L106
L106:
	;
	v666 = v653&int32(4032) | v647
	v668 = v628 - int32(3)
	v670 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v619+v668))))
	if base.B2i32(v668 != v618)&base.B2i32(base.Ui32(v670) < base.Ui32(int32(224))) == int32(0) {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v698 = v670<<(uint(int32(12))%32)&int32(_a_F_r_check_vowel_harmony_0) | v666
	v702 = int32(3)
	goto L102
L108:
	;
	goto L109
L109:
	;
	v688 = int32(4)
	v690 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v628+v619-v688))))
	v698 = v670<<(uint(int32(12))%32)&int32(_a_F_r_check_vowel_harmony_1) | v690&int32(7)<<(uint(int32(18))%32) | v666
	v702 = v688
	goto L102
L110:
	;
	v719 = v628 - v702
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v719
	v628 = v719
	goto L99
L111:
	;
	v704 = v698 - int32(101)
	if v704 < int32(0) {
		goto L110
	} else {
		goto L112
	}
L112:
	;
	v710 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v704)>>(uint(int32(3))%32)))+uint32(_c_F_r_check_vowel_harmony[4]))))
	if int32(base.Ui32(v710)>>(uint(v704&int32(7))%32))&int32(1) == int32(0) {
		goto L110
	} else {
		goto L113
	}
L113:
	;
	v734 = v702
	goto L97
L115:
	;
	v737 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v738 = v737 - v292
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v738
	v740 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v738 <= v740 {
		v887 = v738
		goto L92
	} else {
		goto L116
	}
L116:
	;
	v742 = v738
	goto L96
L117:
	;
	v751 = v742 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v751
	v766 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v767 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v776 = v751
	goto L120
L118:
	;
	if int32(0) <= v882 {
		goto L21
	} else {
		goto L136
	}
L119:
	;
	v882 = int32(-1)
	goto L118
L120:
	;
	if v776 <= v766 {
		goto L119
	} else {
		goto L122
	}
L122:
	;
	v783 = int32(1)
	v784 = v776 - v783
	v786 = int32(*(*int8)(unsafe.Add(mBase, uint32(v767+v784))))
	v788 = v786 & int32(255)
	if base.B2i32(v784 == v766)|base.B2i32(int32(0) <= v786) != 0 {
		v846 = v788
		v850 = v783
		goto L123
	} else {
		goto L124
	}
L123:
	;
	if int32(117) < v846 {
		goto L131
	} else {
		goto L132
	}
L124:
	;
	v795 = v788 & int32(63)
	v797 = v776 - int32(2)
	v799 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v767+v797))))
	v801 = v799 << (uint(int32(6)) % 32)
	if base.B2i32(v797 != v766)&base.B2i32(base.Ui32(v799) < base.Ui32(int32(192))) == int32(0) {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v846 = v801&int32(1984) | v795
	v850 = int32(2)
	goto L123
L126:
	;
	goto L127
L127:
	;
	v814 = v801&int32(4032) | v795
	v816 = v776 - int32(3)
	v818 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v767+v816))))
	if base.B2i32(v816 != v766)&base.B2i32(base.Ui32(v818) < base.Ui32(int32(224))) == int32(0) {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v846 = v818<<(uint(int32(12))%32)&int32(_a_F_r_check_vowel_harmony_0) | v814
	v850 = int32(3)
	goto L123
L129:
	;
	goto L130
L130:
	;
	v836 = int32(4)
	v838 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v776+v767-v836))))
	v846 = v818<<(uint(int32(12))%32)&int32(_a_F_r_check_vowel_harmony_1) | v838&int32(7)<<(uint(int32(18))%32) | v814
	v850 = v836
	goto L123
L131:
	;
	v867 = v776 - v850
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v867
	v776 = v867
	goto L120
L132:
	;
	v852 = v846 - int32(111)
	if v852 < int32(0) {
		goto L131
	} else {
		goto L133
	}
L133:
	;
	v858 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v852)>>(uint(int32(3))%32)))+uint32(_c_F_r_check_vowel_harmony[5]))))
	if int32(base.Ui32(v858)>>(uint(v852&int32(7))%32))&int32(1) == int32(0) {
		goto L131
	} else {
		goto L134
	}
L134:
	;
	v882 = v850
	goto L118
L136:
	;
	v885 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v887 = v885 - v292
	goto L92
L137:
	;
	if v904 != 0 {
		goto L141
	} else {
		goto L142
	}
L138:
	;
	goto L137
L139:
	;
	v897 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v900 = F_memcmp(m, v897+v887-v889, int32(_a_F_r_check_vowel_harmony_3), v889)
	mBase = m.M
	if v900 != 0 {
		v904 = v891
		goto L138
	} else {
		goto L140
	}
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v887 - v889
	v904 = int32(1)
	goto L138
L141:
	;
	v917 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v918 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v919 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v928 = v917
	goto L146
L142:
	;
	goto L143
L143:
	;
	v1037 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1038 = v1037 - v292
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1038
	v1040 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1038 <= v1040 {
		v1186 = v1038
		goto L163
	} else {
		goto L164
	}
L144:
	;
	if int32(0) <= v1034 {
		goto L21
	} else {
		goto L162
	}
L145:
	;
	v1034 = int32(-1)
	goto L144
L146:
	;
	if v928 <= v918 {
		goto L145
	} else {
		goto L148
	}
L148:
	;
	v935 = int32(1)
	v936 = v928 - v935
	v938 = int32(*(*int8)(unsafe.Add(mBase, uint32(v919+v936))))
	v940 = v938 & int32(255)
	if base.B2i32(v936 == v918)|base.B2i32(int32(0) <= v938) != 0 {
		v998 = v940
		v1002 = v935
		goto L149
	} else {
		goto L150
	}
L149:
	;
	if int32(252) < v998 {
		goto L157
	} else {
		goto L158
	}
L150:
	;
	v947 = v940 & int32(63)
	v949 = v928 - int32(2)
	v951 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v919+v949))))
	v953 = v951 << (uint(int32(6)) % 32)
	if base.B2i32(v949 != v918)&base.B2i32(base.Ui32(v951) < base.Ui32(int32(192))) == int32(0) {
		goto L151
	} else {
		goto L152
	}
L151:
	;
	v998 = v953&int32(1984) | v947
	v1002 = int32(2)
	goto L149
L152:
	;
	goto L153
L153:
	;
	v966 = v953&int32(4032) | v947
	v968 = v928 - int32(3)
	v970 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v919+v968))))
	if base.B2i32(v968 != v918)&base.B2i32(base.Ui32(v970) < base.Ui32(int32(224))) == int32(0) {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	v998 = v970<<(uint(int32(12))%32)&int32(_a_F_r_check_vowel_harmony_0) | v966
	v1002 = int32(3)
	goto L149
L155:
	;
	goto L156
L156:
	;
	v988 = int32(4)
	v990 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v928+v919-v988))))
	v998 = v970<<(uint(int32(12))%32)&int32(_a_F_r_check_vowel_harmony_1) | v990&int32(7)<<(uint(int32(18))%32) | v966
	v1002 = v988
	goto L149
L157:
	;
	v1019 = v928 - v1002
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1019
	v928 = v1019
	goto L146
L158:
	;
	v1004 = v998 - int32(246)
	if v1004 < int32(0) {
		goto L157
	} else {
		goto L159
	}
L159:
	;
	v1010 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1004)>>(uint(int32(3))%32)))+uint32(_c_F_r_check_vowel_harmony[6]))))
	if int32(base.Ui32(v1010)>>(uint(v1004&int32(7))%32))&int32(1) == int32(0) {
		goto L157
	} else {
		goto L160
	}
L160:
	;
	v1034 = v1002
	goto L144
L162:
	;
	goto L143
L163:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1186
	v1188 = int32(2)
	v1190 = int32(0)
	v1193 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1186-v1193 < v1188 {
		v1203 = v1190
		goto L186
	} else {
		goto L187
	}
L164:
	;
	v1042 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1046 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1042+v1038-int32(1)))))
	if v1046 != int32(117) {
		v1186 = v1038
		goto L163
	} else {
		goto L165
	}
L165:
	;
	v1050 = v1038 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1050
	v1065 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1066 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1075 = v1050
	goto L168
L166:
	;
	if int32(0) <= v1181 {
		goto L21
	} else {
		goto L184
	}
L167:
	;
	v1181 = int32(-1)
	goto L166
L168:
	;
	if v1075 <= v1065 {
		goto L167
	} else {
		goto L170
	}
L170:
	;
	v1082 = int32(1)
	v1083 = v1075 - v1082
	v1085 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1066+v1083))))
	v1087 = v1085 & int32(255)
	if base.B2i32(v1083 == v1065)|base.B2i32(int32(0) <= v1085) != 0 {
		v1145 = v1087
		v1149 = v1082
		goto L171
	} else {
		goto L172
	}
L171:
	;
	if int32(117) < v1145 {
		goto L179
	} else {
		goto L180
	}
L172:
	;
	v1094 = v1087 & int32(63)
	v1096 = v1075 - int32(2)
	v1098 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1066+v1096))))
	v1100 = v1098 << (uint(int32(6)) % 32)
	if base.B2i32(v1096 != v1065)&base.B2i32(base.Ui32(v1098) < base.Ui32(int32(192))) == int32(0) {
		goto L173
	} else {
		goto L174
	}
L173:
	;
	v1145 = v1100&int32(1984) | v1094
	v1149 = int32(2)
	goto L171
L174:
	;
	goto L175
L175:
	;
	v1113 = v1100&int32(4032) | v1094
	v1115 = v1075 - int32(3)
	v1117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1066+v1115))))
	if base.B2i32(v1115 != v1065)&base.B2i32(base.Ui32(v1117) < base.Ui32(int32(224))) == int32(0) {
		goto L176
	} else {
		goto L177
	}
L176:
	;
	v1145 = v1117<<(uint(int32(12))%32)&int32(_a_F_r_check_vowel_harmony_0) | v1113
	v1149 = int32(3)
	goto L171
L177:
	;
	goto L178
L178:
	;
	v1135 = int32(4)
	v1137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1075+v1066-v1135))))
	v1145 = v1117<<(uint(int32(12))%32)&int32(_a_F_r_check_vowel_harmony_1) | v1137&int32(7)<<(uint(int32(18))%32) | v1113
	v1149 = v1135
	goto L171
L179:
	;
	v1166 = v1075 - v1149
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1166
	v1075 = v1166
	goto L168
L180:
	;
	v1151 = v1145 - int32(111)
	if v1151 < int32(0) {
		goto L179
	} else {
		goto L181
	}
L181:
	;
	v1157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1151)>>(uint(int32(3))%32)))+uint32(_c_F_r_check_vowel_harmony[5]))))
	if int32(base.Ui32(v1157)>>(uint(v1151&int32(7))%32))&int32(1) == int32(0) {
		goto L179
	} else {
		goto L182
	}
L182:
	;
	v1181 = v1149
	goto L166
L184:
	;
	v1184 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1186 = v1184 - v292
	goto L163
L185:
	;
	if v1203 == int32(0) {
		v1349 = v2
		goto L1
	} else {
		goto L189
	}
L186:
	;
	goto L185
L187:
	;
	v1196 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1199 = F_memcmp(m, v1196+v1186-v1188, int32(_a_F_r_check_vowel_harmony_4), v1188)
	mBase = m.M
	if v1199 != 0 {
		v1203 = v1190
		goto L186
	} else {
		goto L188
	}
L188:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1186 - v1188
	v1203 = int32(1)
	goto L186
L189:
	;
	v1218 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1219 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1220 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1229 = v1218
	goto L192
L190:
	;
	if v1335 < int32(0) {
		v1349 = v2
		goto L1
	} else {
		goto L208
	}
L191:
	;
	v1335 = int32(-1)
	goto L190
L192:
	;
	if v1229 <= v1219 {
		goto L191
	} else {
		goto L194
	}
L194:
	;
	v1236 = int32(1)
	v1237 = v1229 - v1236
	v1239 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1220+v1237))))
	v1241 = v1239 & int32(255)
	if base.B2i32(v1237 == v1219)|base.B2i32(int32(0) <= v1239) != 0 {
		v1299 = v1241
		v1303 = v1236
		goto L195
	} else {
		goto L196
	}
L195:
	;
	if int32(252) < v1299 {
		goto L203
	} else {
		goto L204
	}
L196:
	;
	v1248 = v1241 & int32(63)
	v1250 = v1229 - int32(2)
	v1252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1220+v1250))))
	v1254 = v1252 << (uint(int32(6)) % 32)
	if base.B2i32(v1250 != v1219)&base.B2i32(base.Ui32(v1252) < base.Ui32(int32(192))) == int32(0) {
		goto L197
	} else {
		goto L198
	}
L197:
	;
	v1299 = v1254&int32(1984) | v1248
	v1303 = int32(2)
	goto L195
L198:
	;
	goto L199
L199:
	;
	v1267 = v1254&int32(4032) | v1248
	v1269 = v1229 - int32(3)
	v1271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1220+v1269))))
	if base.B2i32(v1269 != v1219)&base.B2i32(base.Ui32(v1271) < base.Ui32(int32(224))) == int32(0) {
		goto L200
	} else {
		goto L201
	}
L200:
	;
	v1299 = v1271<<(uint(int32(12))%32)&int32(_a_F_r_check_vowel_harmony_0) | v1267
	v1303 = int32(3)
	goto L195
L201:
	;
	goto L202
L202:
	;
	v1289 = int32(4)
	v1291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1229+v1220-v1289))))
	v1299 = v1271<<(uint(int32(12))%32)&int32(_a_F_r_check_vowel_harmony_1) | v1291&int32(7)<<(uint(int32(18))%32) | v1267
	v1303 = v1289
	goto L195
L203:
	;
	v1320 = v1229 - v1303
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1320
	v1229 = v1320
	goto L192
L204:
	;
	v1305 = v1299 - int32(246)
	if v1305 < int32(0) {
		goto L203
	} else {
		goto L205
	}
L205:
	;
	v1311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1305)>>(uint(int32(3))%32)))+uint32(_c_F_r_check_vowel_harmony[6]))))
	if int32(base.Ui32(v1311)>>(uint(v1305&int32(7))%32))&int32(1) == int32(0) {
		goto L203
	} else {
		goto L206
	}
L206:
	;
	v1335 = v1303
	goto L190
L208:
	;
	goto L21
}
func F_r_consonant_pair_2(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	if v6 < v8 {
		v120 = int32(0)
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v120
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v6
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v8
	v15 = v6 - int32(1)
	if v8 < v15 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v33 = F_find_among_b(m, l0, int32(_a_F_r_consonant_pair_2_0), int32(4))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L8
	} else {
		goto L9
	}
L4:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17+v15))))
	v21 = v19 - int32(100)
	if base.B2i32(v21 == int32(0))|base.B2i32(v21 == int32(16)) != 0 {
		goto L3
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v12
	return int32(0)
L7:
	;
	goto L6
L8:
	;
	return int32(0)
L9:
	;
	if v33 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v12
	return int32(0)
L11:
	;
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v12
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v43
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v47 = v45 + (v6 - v10)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v47
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L15
L13:
	;
	if v102 < int32(0) {
		v120 = int32(0)
		goto L1
	} else {
		goto L32
	}
L15:
	;
	goto L16
L16:
	;
	goto L17
L17:
	;
	v57 = v47
	v59 = int32(1)
	goto L20
L19:
	;
	v102 = v84
	goto L13
L20:
	;
	if v57 <= v12 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	goto L19
L22:
	;
	v102 = int32(-1)
	goto L13
L23:
	;
	goto L24
L24:
	;
	v64 = v57 - int32(1)
	v66 = int32(*(*int8)(unsafe.Add(mBase, uint32(v50+v64))))
	if base.B2i32(int32(0) <= v66)|base.B2i32(v64 <= v12) != 0 {
		v84 = v64
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v88 = int32(1)
	if v88 < v59 {
		v57 = v84
		v59 = v59 - v88
		goto L20
	} else {
		goto L31
	}
L26:
	;
	v72 = v64
	goto L27
L27:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50+v72))))
	if base.Ui32(int32(191)) < base.Ui32(v77) {
		v84 = v72
		goto L25
	} else {
		goto L29
	}
L28:
	;
	v84 = v12
	goto L25
L29:
	;
	v81 = v72 - int32(1)
	if v12 < v81 {
		v72 = v81
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	goto L21
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v102
	v108 = F_slice_del(m, l0)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L8
	} else {
		goto L33
	}
L33:
	;
	if int32(0) <= v108 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v115 = int32(1)
	goto L36
L35:
	;
	v115 = v108 >> (uint(int32(31)) % 32) & v108
	goto L36
L36:
	;
	v120 = v115
	goto L1
}
func F_r_remove_second_order_prefix_1(m *base.Module, l0 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn13983(m, l0, int32(_a_F_r_remove_second_order_prefix_1_0), int32(_a_F_r_remove_second_order_prefix_1_1), int32(_a_F_r_remove_second_order_prefix_1_2))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_r_undouble_2(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	v2 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v6 <= v7 {
		v27 = v2
		return v27
	} else {
		v10 = v6 - int32(2)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v10
		*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v6 - int32(1)
		if v10 < v7 {
			v27 = v2
			return v27
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v10
			v18 = F_slice_del(m, l0)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				if int32(0) <= v18 {
					v24 = int32(1)
				} else {
					v24 = v18
				}
				v27 = v24
				return v27
			}
		}
	}
}
func F_readstoplist(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
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
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v136 int32
	_ = v136
	v3 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v3
	if l0 == v3 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v10 + int32(48)
	return
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = int32(0)
	goto L1
L3:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v16 == int32(0) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v20 = v10 + int32(4)
	v22 = F_get_tsearch_config_filename(m, l0, int32(_a_F_readstoplist_0))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	F_tsearch_readline_end(m, v10+int32(4))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L7
	} else {
		goto L45
	}
L6:
	;
	v46 = v26
	v47 = v3
	v49 = v3
	goto L19
L7:
	;
	return
L8:
	;
	v24 = F_tsearch_readline_begin(m, v20, v22)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	if v24 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v26 = F_tsearch_readline(m, v20)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
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
	v31 = m.ExcPending
	if v31 != 0 {
		goto L7
	} else {
		goto L15
	}
L13:
	;
	if v26 != 0 {
		goto L6
	} else {
		goto L14
	}
L14:
	;
	v117 = v3
	goto L5
L15:
	;
	F_errcode(m, int32(22))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L7
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v22
	F_errmsg(m, int32(_a_F_readstoplist_1), v10)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L7
	} else {
		goto L17
	}
L17:
	;
	F_errfinish(m, int32(_a_F_readstoplist_2), int32(85), int32(_a_F_readstoplist_3))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
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
	v51 = v46
	goto L21
L20:
	;
	v117 = v108
	goto L5
L21:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	switch v58 {
	case 0, 9, 10, 11, 12, 13, 32:
		goto L23
	default:
		goto L24
	}
L22:
	;
	v62 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v51))) = uint8(v62)
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
	if v64 == v62 {
		goto L27
	} else {
		goto L28
	}
L23:
	;
	goto L22
L24:
	;
	v59 = F_pg_mblen_cstr(m, v51)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L7
	} else {
		goto L25
	}
L25:
	;
	v51 = v59 + v51
	goto L21
L26:
	;
	v112 = F_tsearch_readline(m, v10+int32(4))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L7
	} else {
		goto L43
	}
L27:
	;
	F_pfree(m, v46)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L7
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v69 < v49 {
		v83 = v47
		v84 = v49
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v108 = v47
	v109 = v49
	goto L26
L31:
	;
	v85 = F_strlen(m, v46)
	mBase = m.M
	v88 = m.T0[int32(1147)].(func(*base.Module, int32, int32, int32) int32)(m, v46, v85, int32(100))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L7
	} else {
		goto L38
	}
L32:
	;
	if v49 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v75 = F_palloc(m, int32(256))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L7
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v79 = F_repalloc(m, v47, v49<<(uint(int32(3))%32))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L7
	} else {
		goto L37
	}
L36:
	;
	v83 = v75
	v84 = int32(64)
	goto L31
L37:
	;
	v83 = v79
	v84 = v49 << (uint(int32(1)) % 32)
	goto L31
L38:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v91 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v83+v90<<(uint(v91)%32)))) = v88
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v83+v95<<(uint(v91)%32))))
	if v46 != v99 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	F_pfree(m, v46)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L7
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v103 + int32(1)
	v108 = v83
	v109 = v84
	goto L26
L42:
	;
	goto L41
L43:
	;
	if v112 != 0 {
		v46 = v112
		v47 = v108
		v49 = v109
		goto L19
	} else {
		goto L44
	}
L44:
	;
	goto L20
L45:
	;
	F_pfree(m, v22)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L7
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v117
	if v117 == int32(0) {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v130 <= int32(0) {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	F_pg_qsort(m, v117, v130, int32(4), int32(1170))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L7
	} else {
		goto L49
	}
L49:
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
								F_errmsg_internal(m, int32(_a_F_readtup_index_0), int32(0))
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_readtup_index_1), int32(1833), int32(_a_F_readtup_index_2))
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
					F_errmsg_internal(m, int32(_a_F_readtup_index_0), int32(0))
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_readtup_index_1), int32(1831), int32(_a_F_readtup_index_2))
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
								F_errmsg_internal(m, int32(_a_F_readtup_index_brin_0), int32(0))
								mBase = m.M
								v58 = m.ExcPending
								if v58 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_readtup_index_brin_1), int32(1909), int32(_a_F_readtup_index_brin_2))
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
					F_errmsg_internal(m, int32(_a_F_readtup_index_brin_0), int32(0))
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_readtup_index_brin_1), int32(1907), int32(_a_F_readtup_index_brin_2))
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
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
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
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
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
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+28))
	if v8 != 0 {
		v9 = int32(14)
		v12 = base.I32_clz(v8) ^ int32(31)
		if base.Ui32(v9) <= base.Ui32(v12) {
			v15 = v9
		} else {
			v15 = v12
		}
		v20 = v15 + int32(1)
	} else {
		v20 = int32(0)
	}
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	if v20 == v22 {
		return
	} else {
		v24 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
		if v24 != int32(-1) {
			v27 = F_get_segment_by_index(m, l0, v24)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return
			} else {
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
				v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+16))
				*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v31
				v39 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+16))
				if v40 != int32(-1) {
					v43 = F_get_segment_by_index(m, l0, v40)
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return
					} else {
						v45 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
						v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
						v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v45)+12)) = v47
						v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
						v50 = v49
						v51 = int32(-1)
						*(*int32)(unsafe.Add(mBase, uint32(v50)+12)) = v51
						v53 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
						v55 = v20 << (uint(int32(2)) % 32)
						v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v58 = *(*int32)(unsafe.Add(mBase, uint32(v55+v56)+160))
						*(*int32)(unsafe.Add(mBase, uint32(v53)+16)) = v58
						v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v60)+20)) = v20
						v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v68 = base.I32_div_s(l1-l0-int32(8), int32(20))
						*(*int32)(unsafe.Add(mBase, uint32(v62+v55)+160)) = v68
						v70 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
						v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+16))
						if v71 == v51 {
							return
						} else {
							v74 = F_get_segment_by_index(m, l0, v71)
							mBase = m.M
							v75 = m.ExcPending
							if v75 != 0 {
								return
							} else {
								v76 = *(*int32)(unsafe.Add(mBase, uint32(v74)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v76)+12)) = v68
								return
							}
						}
					}
				} else {
					v50 = v39
					v51 = int32(-1)
					*(*int32)(unsafe.Add(mBase, uint32(v50)+12)) = v51
					v53 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					v55 = v20 << (uint(int32(2)) % 32)
					v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v58 = *(*int32)(unsafe.Add(mBase, uint32(v55+v56)+160))
					*(*int32)(unsafe.Add(mBase, uint32(v53)+16)) = v58
					v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v60)+20)) = v20
					v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v68 = base.I32_div_s(l1-l0-int32(8), int32(20))
					*(*int32)(unsafe.Add(mBase, uint32(v62+v55)+160)) = v68
					v70 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+16))
					if v71 == v51 {
						return
					} else {
						v74 = F_get_segment_by_index(m, l0, v71)
						mBase = m.M
						v75 = m.ExcPending
						if v75 != 0 {
							return
						} else {
							v76 = *(*int32)(unsafe.Add(mBase, uint32(v74)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v76)+12)) = v68
							return
						}
					}
				}
			}
		} else {
			v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v37 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
			*(*int32)(unsafe.Add(mBase, uint32(v33+v22<<(uint(int32(2))%32))+160)) = v37
			v39 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+16))
			if v40 != int32(-1) {
				v43 = F_get_segment_by_index(m, l0, v40)
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return
				} else {
					v45 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
					v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v45)+12)) = v47
					v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					v50 = v49
					v51 = int32(-1)
					*(*int32)(unsafe.Add(mBase, uint32(v50)+12)) = v51
					v53 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					v55 = v20 << (uint(int32(2)) % 32)
					v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v58 = *(*int32)(unsafe.Add(mBase, uint32(v55+v56)+160))
					*(*int32)(unsafe.Add(mBase, uint32(v53)+16)) = v58
					v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v60)+20)) = v20
					v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v68 = base.I32_div_s(l1-l0-int32(8), int32(20))
					*(*int32)(unsafe.Add(mBase, uint32(v62+v55)+160)) = v68
					v70 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+16))
					if v71 == v51 {
						return
					} else {
						v74 = F_get_segment_by_index(m, l0, v71)
						mBase = m.M
						v75 = m.ExcPending
						if v75 != 0 {
							return
						} else {
							v76 = *(*int32)(unsafe.Add(mBase, uint32(v74)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v76)+12)) = v68
							return
						}
					}
				}
			} else {
				v50 = v39
				v51 = int32(-1)
				*(*int32)(unsafe.Add(mBase, uint32(v50)+12)) = v51
				v53 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				v55 = v20 << (uint(int32(2)) % 32)
				v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v58 = *(*int32)(unsafe.Add(mBase, uint32(v55+v56)+160))
				*(*int32)(unsafe.Add(mBase, uint32(v53)+16)) = v58
				v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v60)+20)) = v20
				v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v68 = base.I32_div_s(l1-l0-int32(8), int32(20))
				*(*int32)(unsafe.Add(mBase, uint32(v62+v55)+160)) = v68
				v70 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+16))
				if v71 == v51 {
					return
				} else {
					v74 = F_get_segment_by_index(m, l0, v71)
					mBase = m.M
					v75 = m.ExcPending
					if v75 != 0 {
						return
					} else {
						v76 = *(*int32)(unsafe.Add(mBase, uint32(v74)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v76)+12)) = v68
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
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
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
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int64
	_ = v88
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v128 int32
	_ = v128
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v175 int32
	_ = v175
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v230 int32
	_ = v230
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v281 int32
	_ = v281
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v293 int32
	_ = v293
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
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
	var v378 int32
	_ = v378
	var v383 int32
	_ = v383
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v397 int32
	_ = v397
	var v402 int32
	_ = v402
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v419 int32
	_ = v419
	var v424 int32
	_ = v424
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	v22 = m.G0
	v24 = v22 - int32(112)
	m.G0 = v24
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v27 = F_pg_detoast_datum(m, v26)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v32 = F_pg_detoast_datum(m, v31)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	F_check_stack_depth(m)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	v38 = F_lookup_rowtype_tupdesc(m, v36, v37)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	v43 = F_lookup_rowtype_tupdesc(m, v41, v42)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+108)) = v27
	v48 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+104)) = v48
	*(*uint16)(unsafe.Add(mBase, uint32(v24)+100)) = uint16(v48)
	v52 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+96)) = v52
	v54 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+92)) = int32(base.Ui32(v46) >> (uint(v54) % 32))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+88)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v24)+84)) = v48
	*(*uint16)(unsafe.Add(mBase, uint32(v24)+80)) = uint16(v48)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+76)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v24)+72)) = int32(base.Ui32(v57) >> (uint(v54) % 32))
	if v45 < v40 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v69 = v40
	goto L9
L8:
	;
	v69 = v45
	goto L9
L9:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+16))
	if v71 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	if v36 != v94 {
		goto L16
	} else {
		goto L17
	}
L11:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v70)+20))
	v82 = F_MemoryContextAlloc(m, v77, v69<<(uint(int32(2))%32)+int32(20))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L14
	}
L12:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
	if v74 < v69 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
	v93 = v71
	v94 = v76
	goto L10
L14:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v84)+16)) = v82
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)+16))
	v88 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v87)+4)) = v88
	*(*int32)(unsafe.Add(mBase, uint32(v87))) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v87)+12)) = v88
	v93 = v87
	v94 = int32(0)
	goto L10
L15:
	;
	v150 = F_palloc(m, v40<<(uint(int32(2))%32))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L1
	} else {
		goto L31
	}
L16:
	;
	v103 = v93 + int32(20)
	v107 = v69 << (uint(int32(2)) % 32)
	if v103&int32(3)|base.B2i32(base.Ui32(int32(1024)) < base.Ui32(v107)) == int32(0) {
		goto L22
	} else {
		goto L23
	}
L17:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v93)+8))
	if v96 != v37 {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
	if v98 != v41 {
		goto L16
	} else {
		goto L19
	}
L19:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v93)+16))
	if v100 == v42 {
		goto L15
	} else {
		goto L20
	}
L20:
	;
	goto L16
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v93)+16)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v93)+12)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v93)+8)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v93)+4)) = v36
	goto L15
L22:
	;
	if v107 == int32(0) {
		goto L21
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	if v107 == int32(0) {
		goto L21
	} else {
		goto L30
	}
L25:
	;
	v117 = v93 + v107 + int32(20)
	v119 = v93 + int32(24)
	if base.Ui32(v119) < base.Ui32(v117) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v121 = v117
	goto L28
L27:
	;
	v121 = v119
	goto L28
L28:
	;
	v128 = (v121-v93-int32(21))&int32(-4) + int32(4)
	if v128 == int32(0) {
		goto L21
	} else {
		goto L29
	}
L29:
	;
	base.MemoryFill(m, v103, int32(0), v128)
	goto L21
L30:
	;
	base.MemoryFill(m, v103, int32(0), v107)
	goto L21
L31:
	;
	v152 = F_palloc(m, v40)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	F_heap_deform_tuple(m, v24+int32(92), v38, v150, v152)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v160 = F_palloc(m, v45<<(uint(int32(2))%32))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v162 = F_palloc(m, v45)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	F_heap_deform_tuple(m, v24+int32(72), v43, v160, v162)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v166 = int32(0)
	v168 = base.B2i32(v166 < v40)
	if v168|base.B2i32(v166 < v45) == v166 {
		goto L43
	} else {
		goto L44
	}
L37:
	;
	F_pfree(m, v150)
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L1
	} else {
		goto L96
	}
L38:
	;
	v424 = int32(0)
	goto L37
L39:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L1
	} else {
		goto L92
	}
L40:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L1
	} else {
		goto L87
	}
L41:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L1
	} else {
		goto L81
	}
L42:
	;
	if v333 != v40 {
		goto L39
	} else {
		goto L79
	}
L43:
	;
	v333 = int32(0)
	v335 = v166
	goto L42
L44:
	;
	goto L45
L45:
	;
	v175 = int32(0)
	v182 = v175
	v184 = v166
	v185 = v168
	v186 = base.B2i32(v175 < v45)
	v189 = v175
	goto L46
L46:
	;
	v203 = v185 & int32(1)
	if v203 != 0 {
		goto L50
	} else {
		goto L51
	}
L47:
	;
	v333 = v321
	v335 = v322
	goto L42
L48:
	;
	v329 = base.B2i32(v322 < v45)
	v330 = base.B2i32(v321 < v40)
	if v329|v330 != 0 {
		v182 = v321
		v184 = v322
		v185 = v330
		v186 = v329
		v189 = v326
		goto L46
	} else {
		goto L78
	}
L49:
	;
	v321 = v182 + int32(1)
	v322 = v312
	v326 = v316
	goto L48
L50:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38+v204<<(uint(int32(4))%32)+v182*int32(100))+111)))
	if v211 == int32(1) {
		v312 = v184
		v316 = v189
		goto L49
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	if v186 == int32(0) {
		v333 = v182
		v335 = v184
		goto L42
	} else {
		goto L54
	}
L53:
	;
	goto L52
L54:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	v222 = v43 + v216<<(uint(int32(4))%32) + v184*int32(100)
	v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222)+111)))
	if v223 == int32(1) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v321 = v182
	v322 = v184 + int32(1)
	v326 = v189
	goto L48
L56:
	;
	goto L57
L57:
	;
	if v203 == int32(0) {
		v333 = v182
		v335 = v184
		goto L42
	} else {
		goto L58
	}
L58:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	v236 = v38 + v230<<(uint(int32(4))%32) + v182*int32(100)
	v237 = int32(20)
	v238 = v236 + v237
	v240 = v222 + v237
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v236)+88))
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v222)+88))
	if v241 != v242 {
		goto L41
	} else {
		goto L59
	}
L59:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v240)+96))
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v238)+96))
	v248 = v93 + int32(20) + v189<<(uint(int32(2))%32)
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v248)))
	if v249 != 0 {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184+v162))))
	v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182+v152))))
	if v263 == int32(1) {
		goto L68
	} else {
		goto L69
	}
L61:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v249)))
	if v250 == v241 {
		v259 = v249
		goto L60
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	v253 = F_lookup_type_cache(m, v241, int32(32))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L1
	} else {
		goto L65
	}
L64:
	;
	goto L63
L65:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v253)+80))
	if v255 == int32(0) {
		goto L40
	} else {
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v248))) = v253
	v259 = v253
	goto L60
L67:
	;
	v308 = int32(1)
	v312 = v184 + v308
	v316 = v189 + v308
	goto L49
L68:
	;
	if v261&int32(1) != 0 {
		goto L67
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	if v261&int32(1) != 0 {
		goto L38
	} else {
		goto L72
	}
L71:
	;
	goto L38
L72:
	;
	v270 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+52)) = uint8(v270)
	if v245 == v244 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v274 = v245
	goto L75
L74:
	;
	v274 = v270
	goto L75
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+48)) = v274
	*(*int64)(unsafe.Add(mBase, uint32(v24)+40)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+36)) = v259 + int32(76)
	v281 = int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v24)+54)) = uint16(v281)
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v150+v182<<(uint(v281)%32))))
	v287 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+60)) = uint8(v287)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+56)) = v286
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v160+v184<<(uint(v281)%32))))
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+68)) = uint8(v287)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+64)) = v293
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v259)+76))
	v300 = m.T0[v299].(func(*base.Module, int32) int32)(m, v24+int32(36))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+52)))
	if v302|base.B2i32(v300 == int32(0)) != 0 {
		goto L38
	} else {
		goto L77
	}
L77:
	;
	goto L67
L78:
	;
	goto L47
L79:
	;
	if v335 != v45 {
		goto L39
	} else {
		goto L80
	}
L80:
	;
	v424 = int32(1)
	goto L37
L81:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v238)+68))
	v364 = F_format_type_be(m, v363)
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v240)+68))
	v367 = F_format_type_be(m, v366)
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+24)) = v189 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+20)) = v367
	*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v364
	F_errmsg(m, int32(_a_F_record_eq_0), v24+int32(16))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	F_errfinish(m, int32(_a_F_record_eq_1), int32(1198), int32(_a_F_record_eq_2))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L1
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
	F_errcode(m, int32(52461700))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v253)))
	v392 = F_format_type_be(m, v391)
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v392
	F_errmsg(m, int32(_a_F_record_eq_3), v24)
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	F_errfinish(m, int32(_a_F_record_eq_1), int32(1221), int32(_a_F_record_eq_2))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L92:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	F_errmsg(m, int32(_a_F_record_eq_4), int32(0))
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	F_errfinish(m, int32(_a_F_record_eq_1), int32(1265), int32(_a_F_record_eq_2))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L1
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
	F_pfree(m, v152)
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	F_pfree(m, v160)
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	F_pfree(m, v162)
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v38)+12))
	if int32(0) <= v452 {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	F_DecrTupleDescRefCount(m, v38)
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L1
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
	if int32(0) <= v457 {
		goto L104
	} else {
		goto L105
	}
L103:
	;
	goto L102
L104:
	;
	F_DecrTupleDescRefCount(m, v43)
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L1
	} else {
		goto L107
	}
L105:
	;
	goto L106
L106:
	;
	v462 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v462 != v27 {
		goto L108
	} else {
		goto L109
	}
L107:
	;
	goto L106
L108:
	;
	F_pfree(m, v27)
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L1
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	v466 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v466 != v32 {
		goto L112
	} else {
		goto L113
	}
L111:
	;
	goto L110
L112:
	;
	F_pfree(m, v32)
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L1
	} else {
		goto L115
	}
L113:
	;
	goto L114
L114:
	;
	m.G0 = v24 + int32(112)
	return v424
L115:
	;
	goto L114
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
	var v45 int32
	_ = v45
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
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
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
	var v179 int32
	_ = v179
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v219 int32
	_ = v219
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
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
	v247 = m.ExcPending
	if v247 != 0 {
		goto L10
	} else {
		goto L84
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L10
	} else {
		goto L79
	}
L6:
	;
	m.G0 = v17 + int32(32)
	return v219
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
	v45 = v6
	goto L15
L10:
	;
	return int32(0)
L11:
	;
	if v21 == int32(0) {
		v219 = v6
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
		v219 = v6
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
		v219 = v6
		goto L6
	} else {
		goto L17
	}
L16:
	;
	if v200 == int32(3) {
		v219 = v6
		goto L6
	} else {
		goto L78
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
	if v179 != 0 {
		goto L64
	} else {
		goto L65
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
	v179 = int32(0)
	goto L18
L23:
	;
	goto L24
L24:
	;
	v53 = int32(0)
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	if v55 <= v53 {
		v179 = v53
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
	v179 = v95
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
	v179 = int32(0)
	goto L18
L42:
	;
	goto L43
L43:
	;
	v104 = int32(0)
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	if v106 <= v104 {
		v179 = v104
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
		v168 = v115
		v169 = v118
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v179 = v168
	goto L18
L47:
	;
	v171 = v110 + int32(1)
	if v171 < v169 {
		v110 = v171
		v115 = v168
		v118 = v169
		goto L45
	} else {
		goto L63
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
		v168 = v115
		v169 = v118
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
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138))))
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if base.B2i32(v141 == int32(0))|base.B2i32(v141 != v144) != 0 {
		v162 = v141
		v163 = v144
		goto L54
	} else {
		goto L55
	}
L52:
	;
	goto L51
L53:
	;
	if v162-v163 != 0 {
		v168 = v115
		v169 = v118
		goto L47
	} else {
		goto L60
	}
L54:
	;
	goto L53
L55:
	;
	v147 = v138
	v148 = l2
	goto L56
L56:
	;
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+1)))
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147)+1)))
	if v152 == int32(0) {
		v162 = v152
		v163 = v151
		goto L54
	} else {
		goto L58
	}
L57:
	;
	v162 = v152
	v163 = v151
	goto L54
L58:
	;
	v155 = int32(1)
	if v152 == v151 {
		v147 = v147 + v155
		v148 = v148 + v155
		goto L56
	} else {
		goto L59
	}
L59:
	;
	goto L57
L60:
	;
	if v115 != 0 {
		goto L4
	} else {
		goto L61
	}
L61:
	;
	F_check_lateral_ref_ok(m, v33, v127, l3)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L10
	} else {
		goto L62
	}
L62:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	v168 = v127
	v169 = v167
	goto L47
L63:
	;
	goto L46
L64:
	;
	v189 = int32(1)
	goto L66
L65:
	;
	v189 = int32(3)
	goto L66
L66:
	;
	if v179 != 0 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v190 = v179
	goto L69
L68:
	;
	v190 = v45
	goto L69
L69:
	;
	if l4 == int32(0) {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	if v200 == int32(0) {
		v33 = v199
		v45 = v201
		goto L15
	} else {
		goto L77
	}
L71:
	;
	v199 = v33
	v200 = v189
	v201 = v190
	goto L70
L72:
	;
	goto L73
L73:
	;
	if v179 != 0 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v199 = v33
	v200 = v189
	v201 = v190
	goto L70
L75:
	;
	goto L76
L76:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v193 + int32(1)
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v199 = v197
	v200 = int32(0)
	v201 = v45
	goto L70
L77:
	;
	goto L16
L78:
	;
	v219 = v201
	goto L6
L79:
	;
	F_errcode(m, int32(151126148))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L10
	} else {
		goto L80
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v32
	F_errmsg(m, int32(_a_F_refnameNamespaceItem_0), v17+int32(16))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L10
	} else {
		goto L81
	}
L81:
	;
	F_parser_errposition(m, v33, l3)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L10
	} else {
		goto L82
	}
L82:
	;
	F_errfinish(m, int32(_a_F_refnameNamespaceItem_1), int32(275), int32(_a_F_refnameNamespaceItem_2))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L10
	} else {
		goto L83
	}
L83:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L84:
	;
	F_errcode(m, int32(151126148))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L10
	} else {
		goto L85
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = l2
	F_errmsg(m, int32(_a_F_refnameNamespaceItem_3), v17)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L10
	} else {
		goto L86
	}
L86:
	;
	F_parser_errposition(m, v33, l3)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L10
	} else {
		goto L87
	}
L87:
	;
	F_errfinish(m, int32(_a_F_refnameNamespaceItem_1), int32(228), int32(_a_F_refnameNamespaceItem_4))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L10
	} else {
		goto L88
	}
L88:
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
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	v10 = m.G0
	v12 = v10 - int32(160)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	if v15 == int32(47) {
		v19 = F_palloc0(m, int32(32))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v19
			v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v27 = F_strlen(m, v24+int32(1))
			mBase = m.M
			v32 = F_palloc(m, v27<<(uint(int32(2))%32)+int32(4))
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v36 = v34 + int32(1)
				v37 = F_strlen(m, v36)
				mBase = m.M
				v38 = F_pg_mb2wchar_with_len(m, v36, v32, v37)
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int32(0)
				} else {
					v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v43 = F_pg_regcomp(m, v40, v32, v38, int32(3), int32(950))
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int32(0)
					} else {
						if v43 != 0 {
							v47 = v12 + int32(48)
							F_pg_regerror(m, v43, v47)
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return int32(0)
							} else {
								v51 = F_errstart(m, l4, int32(0))
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
									return int32(0)
								} else {
									if v51 != 0 {
										F_errcode(m, int32(302252162))
										mBase = m.M
										v55 = m.ExcPending
										if v55 != 0 {
											return int32(0)
										} else {
											v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v56 + int32(1)
											*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v47
											F_errmsg(m, int32(_a_F_regcomp_auth_token_0), v12+int32(32))
											mBase = m.M
											v65 = m.ExcPending
											if v65 != 0 {
												return int32(0)
											} else {
												F_set_errcontext_domain(m, int32(0))
												mBase = m.M
												v68 = m.ExcPending
												if v68 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = l1
													*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = l2
													F_errcontext_msg(m, int32(_a_F_regcomp_auth_token_1), v12+int32(16))
													mBase = m.M
													v75 = m.ExcPending
													if v75 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_regcomp_auth_token_2), int32(332), int32(_a_F_regcomp_auth_token_3))
														mBase = m.M
														v80 = m.ExcPending
														if v80 != 0 {
															return int32(0)
														} else {
															v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
															*(*int32)(unsafe.Add(mBase, uint32(v12))) = v81 + int32(1)
															*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v12 + int32(48)
															v89 = F_psprintf(m, int32(_a_F_regcomp_auth_token_0), v12)
															mBase = m.M
															v90 = m.ExcPending
															if v90 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(l3))) = v89
																F_pfree(m, v32)
																mBase = m.M
																v94 = m.ExcPending
																if v94 != 0 {
																	return int32(0)
																} else {
																	v95 = v43
																	m.G0 = v12 + int32(160)
																	return v95
																}
															}
														}
													}
												}
											}
										}
									} else {
										v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										*(*int32)(unsafe.Add(mBase, uint32(v12))) = v81 + int32(1)
										*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v12 + int32(48)
										v89 = F_psprintf(m, int32(_a_F_regcomp_auth_token_0), v12)
										mBase = m.M
										v90 = m.ExcPending
										if v90 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l3))) = v89
											F_pfree(m, v32)
											mBase = m.M
											v94 = m.ExcPending
											if v94 != 0 {
												return int32(0)
											} else {
												v95 = v43
												m.G0 = v12 + int32(160)
												return v95
											}
										}
									}
								}
							}
						} else {
							F_pfree(m, v32)
							mBase = m.M
							v94 = m.ExcPending
							if v94 != 0 {
								return int32(0)
							} else {
								v95 = v43
								m.G0 = v12 + int32(160)
								return v95
							}
						}
					}
				}
			}
		}
	} else {
		v95 = int32(0)
		m.G0 = v12 + int32(160)
		return v95
	}
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
		v14 = F_pstrdup(m, int32(_a_F_regdictionaryout_0))
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
					v43 = F_pg_snprintf(m, v38, int32(64), int32(_a_F_regdictionaryout_1), v8)
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
		v6 = F_pstrdup(m, int32(_a_F_regoperatorout_0))
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
		v12 = F_pstrdup(m, int32(_a_F_regroleout_0))
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
					v29 = F_pg_snprintf(m, v24, int32(64), int32(_a_F_regroleout_1), v6)
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
	var v30 int32
	_ = v30
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
	var v90 int32
	_ = v90
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v150 int32
	_ = v150
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v221 int32
	_ = v221
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v309 int32
	_ = v309
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
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v354 int32
	_ = v354
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
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
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v448 int32
	_ = v448
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
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v495 int32
	_ = v495
	var v500 int32
	_ = v500
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v509 int32
	_ = v509
	var v514 int32
	_ = v514
	v5 = int32(0)
	v12 = m.G0
	v14 = v12 + int32(-64)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v16 - int32(63) {
	case 0:
		v471 = l1
		goto L3
	case 1:
		goto L4
	case 2:
		goto L5
	default:
		goto L1
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L12
	} else {
		goto L134
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L12
	} else {
		goto L131
	}
L3:
	;
	m.G0 = v14 - int32(-64)
	return v471
L4:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v259 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v259 == int32(1) {
		goto L65
	} else {
		goto L66
	}
L5:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v19 == int32(0) {
		v471 = l1
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v28 = v19
	v30 = v5
	v32 = v5
	goto L7
L7:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	if v30 < v35 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	if v90 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L9:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
	v40 = v37 + v30<<(uint(int32(2))%32)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	v42 = F_remove_useless_results_recurse(m, l0, v41, l1+int32(8), l3)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v90 = v32
	goto L11
L11:
	;
	goto L8
L12:
	;
	return int32(0)
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40))) = v42
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v47 == int32(0) {
		v81 = v28
		v82 = v30
		v84 = v32
		goto L14
	} else {
		goto L15
	}
L14:
	;
	if v81 != 0 {
		v28 = v81
		v30 = v82 + int32(1)
		v32 = v84
		goto L7
	} else {
		goto L24
	}
L15:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	if v50 <= int32(1) {
		v81 = v28
		v82 = v30
		v84 = v32
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	if v53 != int32(63) {
		v81 = v28
		v82 = v30
		v84 = v32
		goto L14
	} else {
		goto L17
	}
L17:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	if v56 == int32(0) {
		v81 = v28
		v82 = v30
		v84 = v32
		goto L14
	} else {
		goto L18
	}
L18:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+52))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+12))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v61+v56<<(uint(int32(2))%32)-int32(4))))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+12))
	if v68 != int32(8) {
		v81 = v28
		v82 = v30
		v84 = v32
		goto L14
	} else {
		goto L19
	}
L19:
	;
	v71 = F_find_dependent_phvs_in_jointree(m, l0, l1, v56)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L12
	} else {
		goto L20
	}
L20:
	;
	if v71 != 0 {
		v81 = v28
		v82 = v30
		v84 = v32
		goto L14
	} else {
		goto L21
	}
L21:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v74 = F_list_delete_nth_cell(m, v73, v30)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L12
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v74
	v79 = F_bms_add_member(m, v32, v56)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L12
	} else {
		goto L23
	}
L23:
	;
	v81 = v74
	v82 = v30 - int32(1)
	v84 = v79
	goto L14
L24:
	;
	v90 = v84
	goto L11
L25:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v235 == int32(0) {
		v471 = l1
		goto L3
	} else {
		goto L54
	}
L26:
	;
	if v90 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	if v150 < int32(0) {
		goto L25
	} else {
		goto L38
	}
L28:
	;
	v150 = base.I32_ctz(v136) | v137<<(uint(int32(5))%32)
	goto L27
L29:
	;
	v150 = int32(-2)
	goto L27
L30:
	;
	v103 = base.I32_div_s(int32(0), int32(32))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
	if v104 <= v103 {
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v107 = v90 + int32(8)
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v107+v103<<(uint(int32(2))%32))))
	v114 = v111 & int32(-1)
	if v114 != 0 {
		v136 = v114
		v137 = v103
		goto L28
	} else {
		goto L32
	}
L32:
	;
	v116 = v103 + int32(1)
	if v116 == v104 {
		goto L29
	} else {
		goto L33
	}
L33:
	;
	v119 = v116
	goto L34
L34:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v107+v119<<(uint(int32(2))%32))))
	if v126 != 0 {
		v136 = v126
		v137 = v119
		goto L28
	} else {
		goto L36
	}
L35:
	;
	goto L29
L36:
	;
	v128 = v119 + int32(1)
	if v128 != v104 {
		v119 = v128
		goto L34
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	v159 = v150
	goto L39
L39:
	;
	F_remove_result_refs(m, l0, v159, l1)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L12
	} else {
		goto L41
	}
L40:
	;
	goto L25
L41:
	;
	if v90 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	if int32(0) <= v221 {
		v159 = v221
		goto L39
	} else {
		goto L53
	}
L43:
	;
	v221 = base.I32_ctz(v207) | v208<<(uint(int32(5))%32)
	goto L42
L44:
	;
	v221 = int32(-2)
	goto L42
L45:
	;
	v172 = v159 + int32(1)
	v174 = base.I32_div_s(v172, int32(32))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
	if v175 <= v174 {
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v178 = v90 + int32(8)
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v178+v174<<(uint(int32(2))%32))))
	v185 = v182 & (int32(-1) << (uint(v172) % 32))
	if v185 != 0 {
		v207 = v185
		v208 = v174
		goto L43
	} else {
		goto L47
	}
L47:
	;
	v187 = v174 + int32(1)
	if v187 == v175 {
		goto L44
	} else {
		goto L48
	}
L48:
	;
	v190 = v187
	goto L49
L49:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v178+v190<<(uint(int32(2))%32))))
	if v197 != 0 {
		v207 = v197
		v208 = v190
		goto L43
	} else {
		goto L51
	}
L50:
	;
	goto L44
L51:
	;
	v199 = v190 + int32(1)
	if v199 != v175 {
		v190 = v199
		goto L49
	} else {
		goto L52
	}
L52:
	;
	goto L50
L53:
	;
	goto L40
L54:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v235)+4))
	if v238 != int32(1) {
		v471 = l1
		goto L3
	} else {
		goto L55
	}
L55:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v241)+60))
	if l1 == v242 {
		v471 = l1
		goto L3
	} else {
		goto L56
	}
L56:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v245 != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v246 = l2
	goto L59
L58:
	;
	v246 = int32(1)
	goto L59
L59:
	;
	if v246 == int32(0) {
		v471 = l1
		goto L3
	} else {
		goto L60
	}
L60:
	;
	if v245 != 0 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v250 = F_list_concat(m, v245, v249)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L12
	} else {
		goto L64
	}
L62:
	;
	v254 = v235
	goto L63
L63:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v254)+12))
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v255)))
	v471 = v256
	goto L3
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v250
	v253 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v254 = v253
	goto L63
L65:
	;
	v262 = l2
	goto L67
L66:
	;
	v262 = int32(0)
	goto L67
L67:
	;
	v264 = l1 + int32(28)
	if v259 != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v265 = v262
	goto L70
L69:
	;
	v265 = v264
	goto L70
L70:
	;
	v266 = F_remove_useless_results_recurse(m, l0, v257, v265, l3)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L12
	} else {
		goto L71
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v266
	v269 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v271 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if base.Ui32(v271) < base.Ui32(int32(2)) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v274 = v264
	goto L74
L73:
	;
	v274 = int32(0)
	goto L74
L74:
	;
	v275 = F_remove_useless_results_recurse(m, l0, v269, v274, l3)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L12
	} else {
		goto L75
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v275
	v278 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	switch v278 {
	case 0:
		goto L79
	case 1:
		goto L78
	case 2, 5:
		v471 = l1
		goto L3
	default:
		goto L2
	case 4:
		goto L77
	}
L76:
	;
	v469 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v471 = v469
	goto L3
L77:
	;
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v275)))
	if v421 != int32(63) {
		v471 = l1
		goto L3
	} else {
		goto L120
	}
L78:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v275)))
	if v371 != int32(63) {
		v471 = l1
		goto L3
	} else {
		goto L107
	}
L79:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v279)))
	if v280 != int32(63) {
		v326 = v275
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v326)))
	if v327 != int32(63) {
		v471 = l1
		goto L3
	} else {
		goto L96
	}
L81:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v279)+4))
	if v283 == int32(0) {
		v326 = v275
		goto L80
	} else {
		goto L82
	}
L82:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v286)+52))
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v287)+12))
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v288+v283<<(uint(int32(2))%32)-int32(4))))
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v294)+12))
	if v295 != int32(8) {
		v326 = v275
		goto L80
	} else {
		goto L83
	}
L83:
	;
	v298 = F_find_dependent_phvs_in_jointree(m, l0, v275, v283)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L12
	} else {
		goto L84
	}
L84:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v298 != 0 {
		v326 = v300
		goto L80
	} else {
		goto L85
	}
L85:
	;
	F_remove_result_refs(m, l0, v283, v300)
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L12
	} else {
		goto L86
	}
L86:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v264)))
	v304 = int32(0)
	if l2|base.B2i32(v303 == v304) == v304 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v309
	*(*int32)(unsafe.Add(mBase, uint32(v14)+52)) = v309
	v315 = F_list_make1_impl(m, int32(1), v12+int32(-28))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L12
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	if v303 != 0 {
		goto L92
	} else {
		goto L93
	}
L90:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v318 = F_makeFromExpr(m, v315, v317)
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L12
	} else {
		goto L91
	}
L91:
	;
	v471 = v318
	goto L3
L92:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v321 = F_list_concat(m, v303, v320)
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L12
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v471 = v324
	goto L3
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v321
	goto L94
L96:
	;
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v326)+4))
	if v330 == int32(0) {
		v471 = l1
		goto L3
	} else {
		goto L97
	}
L97:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v333)+52))
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v334)+12))
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v335+v330<<(uint(int32(2))%32)-int32(4))))
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v341)+12))
	if v342 != int32(8) {
		v471 = l1
		goto L3
	} else {
		goto L98
	}
L98:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	F_remove_result_refs(m, l0, v330, v345)
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L12
	} else {
		goto L99
	}
L99:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v349 = int32(0)
	if l2|base.B2i32(v348 == v349) == v349 {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v354
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v354
	v360 = F_list_make1_impl(m, int32(1), v12+int32(-32))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L12
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	if v348 == int32(0) {
		goto L76
	} else {
		goto L105
	}
L103:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v363 = F_makeFromExpr(m, v360, v362)
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L12
	} else {
		goto L104
	}
L104:
	;
	v471 = v363
	goto L3
L105:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v368 = F_list_concat(m, v348, v367)
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L12
	} else {
		goto L106
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v368
	goto L76
L107:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v275)+4))
	if v374 == int32(0) {
		v471 = l1
		goto L3
	} else {
		goto L108
	}
L108:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v377)+52))
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v378)+12))
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v379+v374<<(uint(int32(2))%32)-int32(4))))
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v385)+12))
	if v386 != int32(8) {
		v471 = l1
		goto L3
	} else {
		goto L109
	}
L109:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v264)))
	if v389 == int32(0) {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v413 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	F_remove_result_refs(m, l0, v374, v413)
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L12
	} else {
		goto L118
	}
L111:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v392)+68))
	if v393 == int32(0) {
		goto L110
	} else {
		goto L112
	}
L112:
	;
	v396 = F_bms_make_singleton(m, v374)
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L12
	} else {
		goto L113
	}
L113:
	;
	v398 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+60)) = v398
	*(*int32)(unsafe.Add(mBase, uint32(v14)+56)) = v396
	v401 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v404 = v12 + int32(-8)
	v406 = F_query_tree_walker_impl(m, v401, int32(853), v404, v398)
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L12
	} else {
		goto L114
	}
L114:
	;
	if v406 != 0 {
		v471 = l1
		goto L3
	} else {
		goto L115
	}
L115:
	;
	v408 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v410 = F_expression_tree_walker_impl(m, v408, int32(853), v404)
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L12
	} else {
		goto L116
	}
L116:
	;
	if v410 != 0 {
		v471 = l1
		goto L3
	} else {
		goto L117
	}
L117:
	;
	goto L110
L118:
	;
	v416 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v417 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v418 = F_bms_add_member(m, v416, v417)
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L12
	} else {
		goto L119
	}
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v418
	goto L76
L120:
	;
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v275)+4))
	if v424 == int32(0) {
		v471 = l1
		goto L3
	} else {
		goto L121
	}
L121:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v427)+52))
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v428)+12))
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v429+v424<<(uint(int32(2))%32)-int32(4))))
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v435)+12))
	if v436 != int32(8) {
		v471 = l1
		goto L3
	} else {
		goto L122
	}
L122:
	;
	v439 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	F_remove_result_refs(m, l0, v424, v439)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L12
	} else {
		goto L123
	}
L123:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v443 = int32(0)
	if l2|base.B2i32(v442 == v443) == v443 {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+40)) = v448
	*(*int32)(unsafe.Add(mBase, uint32(v14)+44)) = v448
	v454 = F_list_make1_impl(m, int32(1), v12+int32(-24))
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L12
	} else {
		goto L127
	}
L125:
	;
	goto L126
L126:
	;
	if v442 == int32(0) {
		goto L76
	} else {
		goto L129
	}
L127:
	;
	v456 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v457 = F_makeFromExpr(m, v454, v456)
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L12
	} else {
		goto L128
	}
L128:
	;
	v471 = v457
	goto L3
L129:
	;
	v461 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v462 = F_list_concat(m, v442, v461)
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L12
	} else {
		goto L130
	}
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v462
	goto L76
L131:
	;
	v489 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v489
	F_errmsg_internal(m, int32(_a_F_remove_useless_results_recurse_0), v12+int32(-48))
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L12
	} else {
		goto L132
	}
L132:
	;
	F_errfinish(m, int32(_a_F_remove_useless_results_recurse_1), int32(3924), int32(_a_F_remove_useless_results_recurse_2))
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L12
	} else {
		goto L133
	}
L133:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L134:
	;
	v505 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v505
	F_errmsg_internal(m, int32(_a_F_remove_useless_results_recurse_3), v14)
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L12
	} else {
		goto L135
	}
L135:
	;
	F_errfinish(m, int32(_a_F_remove_useless_results_recurse_1), int32(3930), int32(_a_F_remove_useless_results_recurse_2))
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L12
	} else {
		goto L136
	}
L136:
	;
	base.Wasm_trap_unreachable()
	for {
	}
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
		F_errmsg_internal(m, int32(_a_F_removeabbrev_index_gin_0), int32(0))
		v11 = m.ExcPending
		if v11 != 0 {
			return
		} else {
			F_errfinish(m, int32(_a_F_removeabbrev_index_gin_1), int32(1924), int32(_a_F_removeabbrev_index_gin_2))
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
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
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
	var v68 int32
	_ = v68
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
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
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
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	if base.B2i32(v20 == int32(0))|base.B2i32(v20 != v23) != 0 {
		v41 = v20
		v42 = v23
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
		goto L12
	} else {
		goto L40
	}
L2:
	;
	if v41-v42 != 0 {
		goto L9
	} else {
		goto L10
	}
L3:
	;
	goto L2
L4:
	;
	v26 = v17
	v27 = l3
	goto L5
L5:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+1)))
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+1)))
	if v31 == int32(0) {
		v41 = v31
		v42 = v30
		goto L3
	} else {
		goto L7
	}
L6:
	;
	v41 = v31
	v42 = v30
	goto L3
L7:
	;
	v34 = int32(1)
	if v31 == v30 {
		v26 = v26 + v34
		v27 = v27 + v34
		goto L5
	} else {
		goto L8
	}
L8:
	;
	goto L6
L9:
	;
	v45 = v11 + int32(32)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	F_ScanKeyInit(m, v45, int32(2), int32(3), int32(184), v49)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	m.G0 = v11 + int32(128)
	return
L12:
	;
	return
L13:
	;
	F_ScanKeyInit(m, v11+int32(80), int32(4), int32(3), int32(62), l3)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v63 = F_systable_beginscan(m, l0, int32(2701), int32(1), int32(0), int32(2), v45)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	v65 = F_systable_getnext(m, v63)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L12
	} else {
		goto L16
	}
L16:
	;
	if v65 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	F_systable_endscan(m, v63)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L12
	} else {
		goto L18
	}
L18:
	;
	v69 = F_heap_copytuple(m, l2)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L12
	} else {
		goto L20
	}
L19:
	;
	v125 = F_strncpy(m, v75, l3, int32(64))
	mBase = m.M
	v126 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v125)+63)) = uint8(v126)
	goto L33
L20:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v69)+16))
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+22)))
	v73 = v71 + v72
	v75 = v73 + int32(12)
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75))))
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
	if base.B2i32(v78 == int32(0))|base.B2i32(v78 != v81) != 0 {
		v99 = v78
		v100 = v81
		goto L22
	} else {
		goto L23
	}
L21:
	;
	if v99-v100 == int32(0) {
		goto L19
	} else {
		goto L28
	}
L22:
	;
	goto L21
L23:
	;
	v84 = v75
	v85 = l4
	goto L24
L24:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+1)))
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+1)))
	if v89 == int32(0) {
		v99 = v89
		v100 = v88
		goto L22
	} else {
		goto L26
	}
L25:
	;
	v99 = v89
	v100 = v88
	goto L22
L26:
	;
	v92 = int32(1)
	if v89 == v88 {
		v84 = v84 + v92
		v85 = v85 + v92
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	v106 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L12
	} else {
		goto L29
	}
L29:
	;
	if v106 == int32(0) {
		goto L19
	} else {
		goto L30
	}
L30:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v110 + int32(4)
	F_errmsg(m, int32(_a_F_renametrig_internal_0), v11)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L12
	} else {
		goto L31
	}
L31:
	;
	F_errfinish(m, int32(_a_F_renametrig_internal_1), int32(1633), int32(_a_F_renametrig_internal_2))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L12
	} else {
		goto L32
	}
L32:
	;
	goto L19
L33:
	;
	F_CatalogTupleUpdate(m, l0, v69+int32(4), v69)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L12
	} else {
		goto L34
	}
L34:
	;
	v133 = *(*int32)(unsafe.Add(mBase, _c_F_renametrig_internal[0]))
	if v133 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	v136 = int32(0)
	F_RunObjectPostAlterHook(m, int32(2620), v135, v136, v136, v136)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L12
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	F_CacheInvalidateRelcache(m, l1)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L12
	} else {
		goto L39
	}
L38:
	;
	goto L37
L39:
	;
	goto L11
L40:
	;
	F_errcode(m, int32(_a_F_renametrig_internal_3))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L12
	} else {
		goto L41
	}
L41:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v157 + int32(4)
	F_errmsg(m, int32(_a_F_renametrig_internal_4), v11+int32(16))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L12
	} else {
		goto L42
	}
L42:
	;
	F_errfinish(m, int32(_a_F_renametrig_internal_1), int32(1615), int32(_a_F_renametrig_internal_2))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L12
	} else {
		goto L43
	}
L43:
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
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0-int32(8))))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v6&int32(15)*int32(36))+uint32(_c_F_repalloc[0])))
	v12 = m.T0[v11].(func(*base.Module, int32, int32, int32) int32)(m, l0, l1, int32(0))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		return v12
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
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v50 int64
	_ = v50
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
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
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
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
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L37
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
	v50 = base.I64_extend_i32_s(v17) * base.I64_extend_i32_s(v48)
	v54 = base.I32_wrap_i64(v50)
	if base.I32_wrap_i64(int64(base.Ui64(v50)>>(uint(int64(32))%64))) != v54>>(uint(int32(31))%32) {
		goto L3
	} else {
		goto L18
	}
L8:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+1)))
	if v25 == int32(18) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v36 = int32(1)
	if v19&v36 != 0 {
		v48 = int32(base.Ui32(v19)>>(uint(v36)%32)) - v36
		goto L7
	} else {
		goto L17
	}
L11:
	;
	v28 = int32(16)
	goto L13
L12:
	;
	v28 = int32(0)
	goto L13
L13:
	;
	if base.Ui32((v25-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v35 = int32(4)
	goto L16
L15:
	;
	v35 = v28
	goto L16
L16:
	;
	v48 = v35
	goto L7
L17:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v48 = int32(base.Ui32(v42)>>(uint(int32(2))%32)) - int32(4)
	goto L7
L18:
	;
	v59 = v54 + int32(4)
	if base.B2i32(v59 < v54)|base.B2i32(base.Ui32(int32(1073741824)) <= base.Ui32(v59)) != 0 {
		goto L3
	} else {
		goto L19
	}
L19:
	;
	v64 = F_palloc(m, v59)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v64))) = v59 << (uint(int32(2)) % 32)
	if int32(0) < v13 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v71 = int32(1)
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
	if v73&v71 != 0 {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L23
L23:
	;
	return v64
L24:
	;
	v76 = v71
	goto L26
L25:
	;
	v76 = int32(4)
	goto L26
L26:
	;
	v82 = v64 + int32(4)
	v83 = int32(0)
	goto L27
L27:
	;
	if v48 != 0 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	goto L23
L29:
	;
	base.MemoryCopy(m, v82, v9+v76, v48)
	goto L31
L30:
	;
	goto L31
L31:
	;
	v90 = *(*int32)(unsafe.Add(mBase, _c_F_repeat_2[0]))
	if v90 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v95 = v83 + int32(1)
	if v95 != v13 {
		v82 = v82 + v48
		v83 = v95
		goto L27
	} else {
		goto L36
	}
L35:
	;
	goto L34
L36:
	;
	goto L28
L37:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	F_errmsg(m, int32(_a_F_repeat_2_0), int32(0))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(_a_F_repeat_2_1), int32(1167), int32(_a_F_repeat_2_2))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
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
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_report_newlocale_failure[0]))
	if v9 == int32(0) {
		v13 = int32(44)
		*(*int32)(unsafe.Add(mBase, _c_F_report_newlocale_failure[0])) = v13
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
			F_errmsg(m, int32(_a_F_report_newlocale_failure_0), v6+int32(16))
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return
			} else {
				if v16 == int32(44) {
					*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
					F_errdetail(m, int32(_a_F_report_newlocale_failure_1), v6)
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_report_newlocale_failure_2), int32(829), int32(_a_F_report_newlocale_failure_3))
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
					F_errfinish(m, int32(_a_F_report_newlocale_failure_2), int32(829), int32(_a_F_report_newlocale_failure_3))
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
	v42 = F_pg_sprintf(m, v31, int32(_a_F_report_untranslatable_char_0), v11+int32(16))
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
	v48 = F_pg_sprintf(m, v44, int32(_a_F_report_untranslatable_char_1), int32(0))
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
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l1<<(uint(v70)%32))+uint32(_c_F_report_untranslatable_char[0])))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v74
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(v70)%32))+uint32(_c_F_report_untranslatable_char[0])))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v11 + int32(32)
	F_errmsg(m, int32(_a_F_report_untranslatable_char_2), v11)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(_a_F_report_untranslatable_char_3), int32(1901), int32(_a_F_report_untranslatable_char_4))
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
	var v135 int32
	_ = v135
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
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_rescanLatestTimeLine[0]))
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
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_rescanLatestTimeLine[0]))
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
	F_errfinish(m, int32(_a_F_rescanLatestTimeLine_0), v135, int32(_a_F_rescanLatestTimeLine_1))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L2
	} else {
		goto L32
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_rescanLatestTimeLine[0])) = v18
	v109 = *(*int32)(unsafe.Add(mBase, _c_F_rescanLatestTimeLine[1]))
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
	v37 = *(*int32)(unsafe.Add(mBase, _c_F_rescanLatestTimeLine[0]))
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
	F_errmsg(m, int32(_a_F_rescanLatestTimeLine_2), v14+int32(16))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L2
	} else {
		goto L23
	}
L23:
	;
	v135 = int32(_a_F_rescanLatestTimeLine_3)
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
	F_errmsg(m, int32(_a_F_rescanLatestTimeLine_4), v14)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L2
	} else {
		goto L26
	}
L26:
	;
	v135 = int32(_a_F_rescanLatestTimeLine_5)
	v146 = int32(0)
	goto L5
L27:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_rescanLatestTimeLine[1])) = v25
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
	v126 = *(*int32)(unsafe.Add(mBase, _c_F_rescanLatestTimeLine[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v126
	F_errmsg(m, int32(_a_F_rescanLatestTimeLine_6), v14+int32(32))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L2
	} else {
		goto L31
	}
L31:
	;
	v135 = int32(_a_F_rescanLatestTimeLine_7)
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
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)+208))
	if base.Ui32(v4) < base.Ui32(v3) {
		v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)))
		v8 = v6
	} else {
		v8 = int32(1)
	}
	return v8 & int32(1)
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
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v44 int32
	_ = v44
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v77 float64
	_ = v77
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v100 float64
	_ = v100
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
			v100 = float64(0.5)
			m.G0 = v12 + int32(16)
			return v100
		} else {
			v21 = m.G0
			v23 = v21 - int32(96)
			m.G0 = v23
			v26 = v23 + int32(16)
			v28 = *(*int32)(unsafe.Add(mBase, _c_F_restriction_selectivity[0]))
			F_fmgr_info_cxt_security(m, v14, v26, v28, int32(0))
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
				*(*int32)(unsafe.Add(mBase, uint32(v23)+44)) = v26
				v54 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
				v55 = m.T0[v54].(func(*base.Module, int32) int32)(m, v23+int32(44))
				mBase = m.M
				v56 = m.ExcPending
				if v56 != 0 {
					return float64(0)
				} else {
					v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+60)))
					if v57 == int32(1) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return float64(0)
						} else {
							v64 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
							*(*int32)(unsafe.Add(mBase, uint32(v23))) = v64
							F_errmsg_internal(m, int32(_a_F_restriction_selectivity_0), v23)
							mBase = m.M
							v68 = m.ExcPending
							if v68 != 0 {
								return float64(0)
							} else {
								F_errfinish(m, int32(_a_F_restriction_selectivity_1), int32(1217), int32(_a_F_restriction_selectivity_2))
								mBase = m.M
								v73 = m.ExcPending
								if v73 != 0 {
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
						v77 = *(*float64)(unsafe.Add(mBase, uint32(v55)))
						if base.F64_lt(v77, float64(0))|base.F64_gt(v77, float64(1)) == int32(0) {
							v100 = v77
							m.G0 = v12 + int32(16)
							return v100
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v88 = m.ExcPending
							if v88 != 0 {
								return float64(0)
							} else {
								*(*float64)(unsafe.Add(mBase, uint32(v12))) = v77
								F_errmsg_internal(m, int32(_a_F_restriction_selectivity_3), v12)
								mBase = m.M
								v92 = m.ExcPending
								if v92 != 0 {
									return float64(0)
								} else {
									F_errfinish(m, int32(_a_F_restriction_selectivity_4), int32(2007), int32(_a_F_restriction_selectivity_5))
									mBase = m.M
									v97 = m.ExcPending
									if v97 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = int32(_a_F_rollback_prepared_cb_wrapper_0)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v12
	v16 = *(*int64)(unsafe.Add(mBase, uint32(l1)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = int32(993)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = v16
	v20 = int32(_a_F_rollback_prepared_cb_wrapper_1)
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_rollback_prepared_cb_wrapper[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_rollback_prepared_cb_wrapper[0])) = v10 + int32(4)
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
				*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(_a_F_rollback_prepared_cb_wrapper_2)
				F_errmsg(m, int32(_a_F_rollback_prepared_cb_wrapper_3), v10)
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_rollback_prepared_cb_wrapper_4), int32(1077), int32(_a_F_rollback_prepared_cb_wrapper_5))
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
			*(*int32)(unsafe.Add(mBase, _c_F_rollback_prepared_cb_wrapper[0])) = v61
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
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
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
			if base.B2i32(v23 == int32(0))|base.B2i32(v27 != v28) != 0 {
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0))) = v27
				v42 = int32(1)
				v43 = v27 - v42
				v46 = v21 + v43<<(uint(v42)%32)
				v47 = int32(*(*int16)(unsafe.Add(mBase, uint32(v46))))
				v48 = int32(2)
				v50 = *(*int32)(unsafe.Add(mBase, uint32(v23<<(uint(v48)%32))+uint32(_c_F_round_var[0])))
				v51 = base.I32_rem_s(v47, v50)
				v52 = v47 - v51
				*(*uint16)(unsafe.Add(mBase, uint32(v46))) = uint16(v52)
				v55 = base.I32_div_s(v50, v48)
				if v51 < v55 {
					v93 = v43
				} else {
					v58 = v50 + base.I32_extend16_s(v52)
					if int32(_a_F_round_var_0) < v58 {
						v63 = v58 + int32(_a_F_round_var_1)
					} else {
						v63 = v58
					}
					*(*uint16)(unsafe.Add(mBase, uint32(v46))) = uint16(v63)
					if v58 < int32(_a_F_round_var_2) {
						v93 = v43
					} else {
						v67 = v43
						v73 = v67
						for {
							v79 = int32(1)
							v80 = v73 - v79
							v83 = v21 + v80<<(uint(v79)%32)
							v86 = int32(*(*int16)(unsafe.Add(mBase, uint32(v83))))
							v88 = base.B2i32(int32(_a_F_round_var_3) < v86)
							if int32(_a_F_round_var_3) < v86 {
								v89 = int32(-9999)
							} else {
								v89 = v79
							}
							v90 = v89 + v86
							*(*uint16)(unsafe.Add(mBase, uint32(v83))) = uint16(v90)
							if int32(_a_F_round_var_3) < v86 {
								v73 = v80
								continue
							} else {
								break
							}
							break
						}
						v93 = v80
					}
				}
				if int32(0) <= v93 {
				} else {
					v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v101 - int32(2)
					v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v106 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(l0))) = v105 + v106
					v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v109 + v106
				}
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = v27
			if v23 != 0 {
				v42 = int32(1)
				v43 = v27 - v42
				v46 = v21 + v43<<(uint(v42)%32)
				v47 = int32(*(*int16)(unsafe.Add(mBase, uint32(v46))))
				v48 = int32(2)
				v50 = *(*int32)(unsafe.Add(mBase, uint32(v23<<(uint(v48)%32))+uint32(_c_F_round_var[0])))
				v51 = base.I32_rem_s(v47, v50)
				v52 = v47 - v51
				*(*uint16)(unsafe.Add(mBase, uint32(v46))) = uint16(v52)
				v55 = base.I32_div_s(v50, v48)
				if v51 < v55 {
					v93 = v43
				} else {
					v58 = v50 + base.I32_extend16_s(v52)
					if int32(_a_F_round_var_0) < v58 {
						v63 = v58 + int32(_a_F_round_var_1)
					} else {
						v63 = v58
					}
					*(*uint16)(unsafe.Add(mBase, uint32(v46))) = uint16(v63)
					if v58 < int32(_a_F_round_var_2) {
						v93 = v43
					} else {
						v67 = v43
						v73 = v67
						for {
							v79 = int32(1)
							v80 = v73 - v79
							v83 = v21 + v80<<(uint(v79)%32)
							v86 = int32(*(*int16)(unsafe.Add(mBase, uint32(v83))))
							v88 = base.B2i32(int32(_a_F_round_var_3) < v86)
							if int32(_a_F_round_var_3) < v86 {
								v89 = int32(-9999)
							} else {
								v89 = v79
							}
							v90 = v89 + v86
							*(*uint16)(unsafe.Add(mBase, uint32(v83))) = uint16(v90)
							if int32(_a_F_round_var_3) < v86 {
								v73 = v80
								continue
							} else {
								break
							}
							break
						}
						v93 = v80
					}
				}
			} else {
				v39 = int32(*(*int16)(unsafe.Add(mBase, uint32(v21+v27<<(uint(int32(1))%32)))))
				if v39 <= int32(_a_F_round_var_4) {
					v93 = v27
				} else {
					v67 = v27
					v73 = v67
					for {
						v79 = int32(1)
						v80 = v73 - v79
						v83 = v21 + v80<<(uint(v79)%32)
						v86 = int32(*(*int16)(unsafe.Add(mBase, uint32(v83))))
						v88 = base.B2i32(int32(_a_F_round_var_3) < v86)
						if int32(_a_F_round_var_3) < v86 {
							v89 = int32(-9999)
						} else {
							v89 = v79
						}
						v90 = v89 + v86
						*(*uint16)(unsafe.Add(mBase, uint32(v83))) = uint16(v90)
						if int32(_a_F_round_var_3) < v86 {
							v73 = v80
							continue
						} else {
							break
						}
						break
					}
					v93 = v80
				}
			}
			if int32(0) <= v93 {
			} else {
				v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v101 - int32(2)
				v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v106 = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(l0))) = v105 + v106
				v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v109 + v106
			}
		}
		return
	}
}
func F_rt_cube_size(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 float64
	_ = v3
	var v4 int32
	_ = v4
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v28 float64
	_ = v28
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 float64
	_ = v42
	var v43 float64
	_ = v43
	var v50 float64
	_ = v50
	var v51 float64
	_ = v51
	var v54 float64
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v64 float64
	_ = v64
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v78 float64
	_ = v78
	var v79 float64
	_ = v79
	var v85 float64
	_ = v85
	v3 = float64(0)
	v4 = int32(0)
	if l0 == v4 {
		v85 = v3
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v13 <= int32(0) {
			v85 = v3
		} else {
			v17 = l0 + int32(8)
			if v13 == int32(1) {
				v64 = float64(1)
				v66 = v4
				v72 = int32(3)
				v74 = v17 + v66<<(uint(v72)%32)
				v78 = *(*float64)(unsafe.Add(mBase, uint32(v74+v13<<(uint(v72)%32))))
				v79 = *(*float64)(unsafe.Add(mBase, uint32(v74)))
				v85 = base.F64_mul(v64, base.F64_abs(base.F64_sub(v78, v79)))
			} else {
				v28 = float64(1)
				v30 = v4
				v35 = v4
				for {
					v36 = int32(3)
					v38 = v17 + v30<<(uint(v36)%32)
					v40 = v13 << (uint(v36) % 32)
					v42 = *(*float64)(unsafe.Add(mBase, uint32(v38+v40)))
					v43 = *(*float64)(unsafe.Add(mBase, uint32(v38)))
					v50 = *(*float64)(unsafe.Add(mBase, uint32(v38+int32(8)+v40)))
					v51 = *(*float64)(unsafe.Add(mBase, uint32(v38)+8))
					v54 = base.F64_mul(base.F64_mul(v28, base.F64_abs(base.F64_sub(v42, v43))), base.F64_abs(base.F64_sub(v50, v51)))
					v55 = int32(2)
					v56 = v30 + v55
					v58 = v35 + v55
					if v58 != v13&int32(2147483646) {
						v28 = v54
						v30 = v56
						v35 = v58
						continue
					} else {
						break
					}
					break
				}
				if v13&int32(1) == int32(0) {
					v85 = v54
				} else {
					v64 = v54
					v66 = v56
					v72 = int32(3)
					v74 = v17 + v66<<(uint(v72)%32)
					v78 = *(*float64)(unsafe.Add(mBase, uint32(v74+v13<<(uint(v72)%32))))
					v79 = *(*float64)(unsafe.Add(mBase, uint32(v74)))
					v85 = base.F64_mul(v64, base.F64_abs(base.F64_sub(v78, v79)))
				}
			}
		}
	}
	*(*float64)(unsafe.Add(mBase, uint32(l1))) = v85
	return
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
	var v126 int32
	_ = v126
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
	var v238 int32
	_ = v238
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
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v464 int32
	_ = v464
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v507 int32
	_ = v507
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v564 int32
	_ = v564
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
	var v576 int32
	_ = v576
	var v582 int32
	_ = v582
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v592 int32
	_ = v592
	var v595 int32
	_ = v595
	var v597 int32
	_ = v597
	var v601 int32
	_ = v601
	var v605 int32
	_ = v605
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v622 int32
	_ = v622
	var v626 int32
	_ = v626
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v645 int32
	_ = v645
	var v647 int32
	_ = v647
	var v650 int32
	_ = v650
	var v652 int32
	_ = v652
	var v654 int32
	_ = v654
	var v656 int32
	_ = v656
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v674 int32
	_ = v674
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v682 int32
	_ = v682
	var v684 int32
	_ = v684
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v690 int32
	_ = v690
	var v694 int32
	_ = v694
	var v700 int32
	_ = v700
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v710 int32
	_ = v710
	var v712 int32
	_ = v712
	var v716 int32
	_ = v716
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v737 int32
	_ = v737
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	goto L2
L1:
	;
	return v737
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
	v41 = F_slice_from_s(m, l0, v36, int32(_a_F_russian_KOI8_R_stem_0))
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
	v737 = v41
	goto L1
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v51
	v282 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v282
	v284 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v284)+4))
	if v282 < v285 {
		v737 = int32(0)
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
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v78)>>(uint(int32(3))%32)))+uint32(_c_F_russian_KOI8_R_stem[0]))))
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
	v126 = v117
	goto L41
L40:
	;
	v161 = int32(1)
	goto L36
L41:
	;
	if v126 == v120 {
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
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133+v126))))
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
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v137)>>(uint(int32(3))%32)))+uint32(_c_F_russian_KOI8_R_stem[0]))))
	if int32(base.Ui32(v143)>>(uint(v137&int32(7))%32))&int32(1) == int32(0) {
		goto L40
	} else {
		goto L48
	}
L48:
	;
	v152 = v126 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v152
	v126 = v152
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
	v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v192)>>(uint(int32(3))%32)))+uint32(_c_F_russian_KOI8_R_stem[0]))))
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
	v238 = v219
	goto L72
L71:
	;
	v273 = int32(1)
	goto L67
L72:
	;
	if v238 == v232 {
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
	v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v245+v238))))
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
	v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v249)>>(uint(int32(3))%32)))+uint32(_c_F_russian_KOI8_R_stem[0]))))
	if int32(base.Ui32(v255)>>(uint(v249&int32(7))%32))&int32(1) == int32(0) {
		goto L71
	} else {
		goto L79
	}
L79:
	;
	v264 = v238 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v264
	v238 = v264
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
		v345 = v285
		goto L84
	} else {
		goto L85
	}
L83:
	;
	v592 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v592
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v592
	v595 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v592 <= v595 {
		v615 = v595
		v616 = v592
		v617 = v592
		goto L164
	} else {
		goto L165
	}
L84:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v347
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v347
	v351 = v347 - int32(1)
	if v351 <= v345 {
		v363 = v347
		goto L101
	} else {
		goto L102
	}
L85:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v292 = int32(1)
	v294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v290+v282-v292))))
	if base.B2i32(v294&int32(224) != int32(192))|base.B2i32(v292<<(uint(v294)%32)&int32(25166336) == int32(0)) != 0 {
		v345 = v285
		goto L84
	} else {
		goto L86
	}
L86:
	;
	v308 = F_find_among_b(m, l0, int32(_a_F_russian_KOI8_R_stem_1), int32(9))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L16
	} else {
		goto L87
	}
L87:
	;
	if v308 == int32(0) {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v345 = v312
	goto L84
L89:
	;
	goto L90
L90:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v313
	switch v308 - int32(1) {
	case 0:
		goto L92
	case 1:
		goto L91
	default:
		goto L83
	}
L91:
	;
	v340 = F_slice_del(m, l0)
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L16
	} else {
		goto L97
	}
L92:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v313 <= v317 {
		v345 = v317
		goto L84
	} else {
		goto L93
	}
L93:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v319+v313-int32(1)))))
	v325 = v323 - int32(193)
	v326 = int32(0)
	if base.B2i32(v325 == v326)|base.B2i32(v325 == int32(16)) == v326 {
		v345 = v317
		goto L84
	} else {
		goto L94
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v313 - int32(1)
	v336 = F_slice_del(m, l0)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L16
	} else {
		goto L95
	}
L95:
	;
	if int32(0) <= v336 {
		goto L83
	} else {
		goto L96
	}
L96:
	;
	v737 = v336
	goto L1
L97:
	;
	if int32(0) <= v340 {
		goto L83
	} else {
		goto L98
	}
L98:
	;
	v737 = v340
	goto L1
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v373
	v377 = v373 - int32(1)
	v378 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v377 <= v378 {
		v399 = v374
		goto L111
	} else {
		goto L112
	}
L100:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v365
	v367 = F_slice_del(m, l0)
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L16
	} else {
		goto L106
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v363
	v373 = v363
	v374 = v363
	goto L99
L102:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v353+v351))))
	switch v355 - int32(209) {
	case 0, 7:
		goto L103
	default:
		v363 = v347
		goto L101
	}
L103:
	;
	v360 = F_find_among_b(m, l0, int32(_a_F_russian_KOI8_R_stem_2), int32(2))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L16
	} else {
		goto L104
	}
L104:
	;
	if v360 != 0 {
		goto L100
	} else {
		goto L105
	}
L105:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v363 = v362
	goto L101
L106:
	;
	if v367 < int32(0) {
		v737 = v367
		goto L1
	} else {
		goto L107
	}
L107:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v372 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v373 = v371
	v374 = v372
	goto L99
L108:
	;
	if v586 != 0 {
		v737 = v585
		goto L1
	} else {
		goto L163
	}
L109:
	;
	v538 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v539 = v538 - v400
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v539
	v541 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v539
	v544 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v539 <= v544 {
		v576 = v541
		goto L150
	} else {
		goto L151
	}
L110:
	;
	v475 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v475
	v477 = F_slice_del(m, l0)
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L16
	} else {
		goto L136
	}
L111:
	;
	v400 = v374 - v373
	v401 = v399 - v400
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v401
	v403 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v401
	v406 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v401 <= v406 {
		v464 = v403
		goto L117
	} else {
		goto L118
	}
L112:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v382 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v380+v377))))
	if v382&int32(224) != int32(192) {
		v399 = v374
		goto L111
	} else {
		goto L113
	}
L113:
	;
	if int32(1)<<(uint(v382)%32)&int32(_a_F_russian_KOI8_R_stem_3) == int32(0) {
		v399 = v374
		goto L111
	} else {
		goto L114
	}
L114:
	;
	v395 = F_find_among_b(m, l0, int32(_a_F_russian_KOI8_R_stem_4), int32(26))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L16
	} else {
		goto L115
	}
L115:
	;
	if v395 != 0 {
		goto L110
	} else {
		goto L116
	}
L116:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v399 = v397
	goto L111
L117:
	;
	v468 = int32(base.Ui32(v464) >> (uint(int32(31)) % 32))
	if v464 != 0 {
		goto L131
	} else {
		goto L132
	}
L118:
	;
	v408 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v410 = int32(1)
	v412 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v408+v401-v410))))
	if base.B2i32(v412&int32(224) != int32(192))|base.B2i32(v410<<(uint(v412)%32)&int32(51443235) == int32(0)) != 0 {
		v464 = v403
		goto L117
	} else {
		goto L119
	}
L119:
	;
	v426 = F_find_among_b(m, l0, int32(_a_F_russian_KOI8_R_stem_5), int32(46))
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L16
	} else {
		goto L120
	}
L120:
	;
	if v426 == int32(0) {
		v464 = v403
		goto L117
	} else {
		goto L121
	}
L121:
	;
	v430 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v430
	switch v426 - int32(1) {
	case 0:
		goto L124
	case 1:
		goto L123
	default:
		goto L122
	}
L122:
	;
	v464 = int32(1)
	goto L117
L123:
	;
	v457 = F_slice_del(m, l0)
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L16
	} else {
		goto L129
	}
L124:
	;
	v434 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v430 <= v434 {
		v464 = v403
		goto L117
	} else {
		goto L125
	}
L125:
	;
	v436 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v440 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v436+v430-int32(1)))))
	v442 = v440 - int32(193)
	v443 = int32(0)
	if base.B2i32(v442 == v443)|base.B2i32(v442 == int32(16)) == v443 {
		v464 = v403
		goto L117
	} else {
		goto L126
	}
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v430 - int32(1)
	v453 = F_slice_del(m, l0)
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L16
	} else {
		goto L127
	}
L127:
	;
	if int32(0) <= v453 {
		goto L122
	} else {
		goto L128
	}
L128:
	;
	v464 = v453
	goto L117
L129:
	;
	if v457 < int32(0) {
		v464 = v457
		goto L117
	} else {
		goto L130
	}
L130:
	;
	goto L122
L131:
	;
	v470 = v468
	goto L133
L132:
	;
	v470 = int32(13)
	goto L133
L133:
	;
	if v470 == int32(0) {
		goto L83
	} else {
		goto L134
	}
L134:
	;
	if v470 == int32(13) {
		goto L109
	} else {
		goto L135
	}
L135:
	;
	v585 = v464
	v586 = v468
	goto L108
L136:
	;
	if v477 < int32(0) {
		v737 = v477
		goto L1
	} else {
		goto L137
	}
L137:
	;
	v481 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v481
	v483 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v481 <= v483 {
		goto L83
	} else {
		goto L138
	}
L138:
	;
	v485 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v487 = int32(1)
	v489 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v485+v481-v487))))
	if base.B2i32(v489&int32(224) != int32(192))|base.B2i32(v487<<(uint(v489)%32)&int32(671113216) == int32(0)) != 0 {
		goto L83
	} else {
		goto L139
	}
L139:
	;
	v503 = F_find_among_b(m, l0, int32(_a_F_russian_KOI8_R_stem_6), int32(8))
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L16
	} else {
		goto L140
	}
L140:
	;
	if v503 == int32(0) {
		goto L83
	} else {
		goto L141
	}
L141:
	;
	v507 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v507
	switch v503 - int32(1) {
	case 0:
		goto L143
	case 1:
		goto L142
	default:
		goto L83
	}
L142:
	;
	v534 = F_slice_del(m, l0)
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L16
	} else {
		goto L148
	}
L143:
	;
	v511 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v507 <= v511 {
		goto L83
	} else {
		goto L144
	}
L144:
	;
	v513 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v517 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v513+v507-int32(1)))))
	v519 = v517 - int32(193)
	v520 = int32(0)
	if base.B2i32(v519 == v520)|base.B2i32(v519 == int32(16)) == v520 {
		goto L83
	} else {
		goto L145
	}
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v507 - int32(1)
	v530 = F_slice_del(m, l0)
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L16
	} else {
		goto L146
	}
L146:
	;
	if int32(0) <= v530 {
		goto L83
	} else {
		goto L147
	}
L147:
	;
	v737 = v530
	goto L1
L148:
	;
	if int32(0) <= v534 {
		goto L83
	} else {
		goto L149
	}
L149:
	;
	v737 = v534
	goto L1
L150:
	;
	if v576 == int32(0) {
		goto L83
	} else {
		goto L159
	}
L151:
	;
	v546 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v548 = int32(1)
	v550 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v546+v539-v548))))
	if base.B2i32(v550&int32(224) != int32(192))|base.B2i32(v548<<(uint(v550)%32)&int32(60991267) == int32(0)) != 0 {
		v576 = v541
		goto L150
	} else {
		goto L152
	}
L152:
	;
	v564 = F_find_among_b(m, l0, int32(_a_F_russian_KOI8_R_stem_7), int32(36))
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L16
	} else {
		goto L153
	}
L153:
	;
	if v564 == int32(0) {
		v576 = v541
		goto L150
	} else {
		goto L154
	}
L154:
	;
	v568 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v568
	v571 = F_slice_del(m, l0)
	mBase = m.M
	v572 = m.ExcPending
	if v572 != 0 {
		goto L16
	} else {
		goto L155
	}
L155:
	;
	if int32(0) <= v571 {
		goto L156
	} else {
		goto L157
	}
L156:
	;
	v575 = int32(1)
	goto L158
L157:
	;
	v575 = v571
	goto L158
L158:
	;
	v576 = v575
	goto L150
L159:
	;
	if v576 < int32(0) {
		goto L160
	} else {
		goto L161
	}
L160:
	;
	v582 = v576
	goto L162
L161:
	;
	v582 = v464
	goto L162
L162:
	;
	v585 = v582
	v586 = int32(base.Ui32(v576) >> (uint(int32(31)) % 32))
	goto L108
L163:
	;
	goto L83
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v616
	if v616-int32(2) <= v615 {
		goto L169
	} else {
		goto L170
	}
L165:
	;
	v597 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v601 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v597+v592-int32(1)))))
	if v601 != int32(201) {
		v615 = v595
		v616 = v592
		v617 = v592
		goto L164
	} else {
		goto L166
	}
L166:
	;
	v605 = v592 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v605
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v605
	v608 = F_slice_del(m, l0)
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L16
	} else {
		goto L167
	}
L167:
	;
	if v608 < int32(0) {
		v737 = v608
		goto L1
	} else {
		goto L168
	}
L168:
	;
	v612 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v613 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v614 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v615 = v612
	v616 = v613
	v617 = v614
	goto L164
L169:
	;
	v645 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v647 = v645 + (v616 - v617)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v647
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v647
	v650 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v647 <= v650 {
		goto L177
	} else {
		goto L178
	}
L170:
	;
	v622 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v626 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v622+v616-int32(1)))))
	switch v626 - int32(212) {
	case 0, 4:
		goto L171
	default:
		goto L169
	}
L171:
	;
	v631 = F_find_among_b(m, l0, int32(_a_F_russian_KOI8_R_stem_8), int32(2))
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L16
	} else {
		goto L172
	}
L172:
	;
	if v631 == int32(0) {
		goto L169
	} else {
		goto L173
	}
L173:
	;
	v635 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v635
	v637 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v638 = *(*int32)(unsafe.Add(mBase, uint32(v637)))
	if v635 < v638 {
		goto L169
	} else {
		goto L174
	}
L174:
	;
	v640 = F_slice_del(m, l0)
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
		goto L16
	} else {
		goto L175
	}
L175:
	;
	if v640 < int32(0) {
		v737 = v640
		goto L1
	} else {
		goto L176
	}
L176:
	;
	goto L169
L177:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v51
	v737 = int32(1)
	goto L1
L178:
	;
	v652 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v654 = int32(1)
	v656 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v652+v647-v654))))
	if base.B2i32(v656&int32(224) != int32(192))|base.B2i32(v654<<(uint(v656)%32)&int32(151011360) == int32(0)) != 0 {
		goto L177
	} else {
		goto L179
	}
L179:
	;
	v670 = F_find_among_b(m, l0, int32(_a_F_russian_KOI8_R_stem_9), int32(4))
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
		goto L16
	} else {
		goto L180
	}
L180:
	;
	if v670 == int32(0) {
		goto L177
	} else {
		goto L181
	}
L181:
	;
	v674 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v674
	switch v670 - int32(1) {
	case 0:
		goto L184
	case 1:
		goto L183
	case 2:
		goto L182
	default:
		goto L177
	}
L182:
	;
	v726 = F_slice_del(m, l0)
	mBase = m.M
	v727 = m.ExcPending
	if v727 != 0 {
		goto L16
	} else {
		goto L197
	}
L183:
	;
	v710 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v674 <= v710 {
		goto L177
	} else {
		goto L193
	}
L184:
	;
	v678 = F_slice_del(m, l0)
	mBase = m.M
	v679 = m.ExcPending
	if v679 != 0 {
		goto L16
	} else {
		goto L185
	}
L185:
	;
	if v678 < int32(0) {
		v737 = v678
		goto L1
	} else {
		goto L186
	}
L186:
	;
	v682 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v682
	v684 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v682 <= v684 {
		goto L177
	} else {
		goto L187
	}
L187:
	;
	v686 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v687 = v686 + v682
	v690 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v687-int32(1)))))
	if v690 != int32(206) {
		goto L177
	} else {
		goto L188
	}
L188:
	;
	v694 = v682 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v694
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v694
	if v694 <= v684 {
		goto L177
	} else {
		goto L189
	}
L189:
	;
	v700 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v687-int32(2)))))
	if v700 != int32(206) {
		goto L177
	} else {
		goto L190
	}
L190:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v682 - int32(2)
	v706 = F_slice_del(m, l0)
	mBase = m.M
	v707 = m.ExcPending
	if v707 != 0 {
		goto L16
	} else {
		goto L191
	}
L191:
	;
	if int32(0) <= v706 {
		goto L177
	} else {
		goto L192
	}
L192:
	;
	v737 = v706
	goto L1
L193:
	;
	v712 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v716 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v712+v674-int32(1)))))
	if v716 != int32(206) {
		goto L177
	} else {
		goto L194
	}
L194:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v674 - int32(1)
	v722 = F_slice_del(m, l0)
	mBase = m.M
	v723 = m.ExcPending
	if v723 != 0 {
		goto L16
	} else {
		goto L195
	}
L195:
	;
	if int32(0) <= v722 {
		goto L177
	} else {
		goto L196
	}
L196:
	;
	v737 = v722
	goto L1
L197:
	;
	if v726 < int32(0) {
		v737 = v726
		goto L1
	} else {
		goto L198
	}
L198:
	;
	goto L177
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
	var v35 int32
	_ = v35
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
	var v117 int32
	_ = v117
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
	var v874 int32
	_ = v874
	var v878 int32
	_ = v878
	var v880 int32
	_ = v880
	var v885 int32
	_ = v885
	var v886 int32
	_ = v886
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v897 int32
	_ = v897
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v904 int32
	_ = v904
	var v906 int32
	_ = v906
	var v911 int32
	_ = v911
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v920 int32
	_ = v920
	var v923 int32
	_ = v923
	var v925 int32
	_ = v925
	var v928 int32
	_ = v928
	var v931 int32
	_ = v931
	var v934 int32
	_ = v934
	var v938 int32
	_ = v938
	var v941 int32
	_ = v941
	var v943 int32
	_ = v943
	var v945 int32
	_ = v945
	var v946 int32
	_ = v946
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v951 int32
	_ = v951
	var v953 int32
	_ = v953
	var v955 int32
	_ = v955
	var v959 int32
	_ = v959
	var v963 int32
	_ = v963
	var v968 int32
	_ = v968
	var v969 int32
	_ = v969
	var v972 int32
	_ = v972
	var v974 int32
	_ = v974
	var v975 int32
	_ = v975
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v982 int32
	_ = v982
	var v984 int32
	_ = v984
	var v989 int32
	_ = v989
	var v990 int32
	_ = v990
	var v993 int32
	_ = v993
	var v997 int32
	_ = v997
	var v998 int32
	_ = v998
	var v1001 int32
	_ = v1001
	var v1003 int32
	_ = v1003
	var v1005 int32
	_ = v1005
	var v1008 int32
	_ = v1008
	var v1011 int32
	_ = v1011
	var v1014 int32
	_ = v1014
	var v1018 int32
	_ = v1018
	var v1021 int32
	_ = v1021
	var v1023 int32
	_ = v1023
	var v1025 int32
	_ = v1025
	var v1028 int32
	_ = v1028
	var v1031 int32
	_ = v1031
	var v1034 int32
	_ = v1034
	var v1038 int32
	_ = v1038
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1045 int32
	_ = v1045
	var v1047 int32
	_ = v1047
	var v1049 int32
	_ = v1049
	var v1050 int32
	_ = v1050
	var v1053 int32
	_ = v1053
	var v1056 int32
	_ = v1056
	var v1060 int32
	_ = v1060
	var v1063 int32
	_ = v1063
	var v1064 int32
	_ = v1064
	var v1067 int32
	_ = v1067
	var v1068 int32
	_ = v1068
	var v1077 int32
	_ = v1077
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	goto L2
L1:
	;
	return v1077
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
	v26 = F_memcmp(m, v24+v14, int32(_a_F_russian_UTF_8_stem_0), v16)
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
	v35 = v14
	goto L12
L10:
	;
	v117 = v14
	goto L11
L11:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v117
	v126 = F_slice_from_s(m, l0, int32(2), int32(_a_F_russian_UTF_8_stem_1))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L40
	} else {
		goto L41
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v35
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L16
L13:
	;
	v117 = v93
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
	v48 = v35
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
		v35 = v93
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
	v108 = F_memcmp(m, v106+v93, int32(_a_F_russian_UTF_8_stem_0), v98)
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
	v1077 = v126
	goto L1
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v137
	v634 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v634
	v636 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v637 = *(*int32)(unsafe.Add(mBase, uint32(v636)+4))
	if v634 < v637 {
		v1077 = int32(0)
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
	v226 = v211&int32(63) | (v169<<(uint(int32(18))%32)&int32(_a_F_russian_UTF_8_stem_2) | v178<<(uint(int32(12))%32) | v194<<(uint(int32(6))%32))
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
	v226 = v169<<(uint(int32(12))%32)&int32(_a_F_russian_UTF_8_stem_3) | v178<<(uint(int32(6))%32) | v194
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
	v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v231)>>(uint(int32(3))%32)))+uint32(_c_F_russian_UTF_8_stem[0]))))
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
	v350 = v335&int32(63) | (v293<<(uint(int32(18))%32)&int32(_a_F_russian_UTF_8_stem_2) | v302<<(uint(int32(12))%32) | v318<<(uint(int32(6))%32))
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
	v350 = v293<<(uint(int32(12))%32)&int32(_a_F_russian_UTF_8_stem_3) | v302<<(uint(int32(6))%32) | v318
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
	v361 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v355)>>(uint(int32(3))%32)))+uint32(_c_F_russian_UTF_8_stem[0]))))
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
	v473 = v458&int32(63) | (v416<<(uint(int32(18))%32)&int32(_a_F_russian_UTF_8_stem_2) | v425<<(uint(int32(12))%32) | v441<<(uint(int32(6))%32))
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
	v473 = v416<<(uint(int32(12))%32)&int32(_a_F_russian_UTF_8_stem_3) | v425<<(uint(int32(6))%32) | v441
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
	v484 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v478)>>(uint(int32(3))%32)))+uint32(_c_F_russian_UTF_8_stem[0]))))
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
	v595 = v580&int32(63) | (v538<<(uint(int32(18))%32)&int32(_a_F_russian_UTF_8_stem_2) | v547<<(uint(int32(12))%32) | v563<<(uint(int32(6))%32))
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
	v595 = v538<<(uint(int32(12))%32)&int32(_a_F_russian_UTF_8_stem_3) | v547<<(uint(int32(6))%32) | v563
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
	v606 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v600)>>(uint(int32(3))%32)))+uint32(_c_F_russian_UTF_8_stem[0]))))
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
	v643 = F_find_among_b(m, l0, int32(_a_F_russian_UTF_8_stem_4), int32(9))
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
		goto L40
	} else {
		goto L149
	}
L147:
	;
	v920 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v920
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v920
	v923 = int32(2)
	v925 = int32(0)
	v928 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v920-v928 < v923 {
		v938 = v925
		goto L247
	} else {
		goto L248
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
	v663 = F_memcmp(m, v660+v656-v652, int32(_a_F_russian_UTF_8_stem_5), v652)
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
	v685 = F_memcmp(m, v682+v672-v674, int32(_a_F_russian_UTF_8_stem_6), v674)
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
	v1077 = v692
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
	v1077 = v696
	goto L1
L169:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v731
	v736 = F_find_among_b(m, l0, int32(_a_F_russian_UTF_8_stem_7), int32(26))
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
	v718 = F_find_among_b(m, l0, int32(_a_F_russian_UTF_8_stem_8), int32(2))
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
		v1077 = v725
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
	v814 = F_find_among_b(m, l0, int32(_a_F_russian_UTF_8_stem_9), int32(46))
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
		v1077 = v740
		goto L1
	} else {
		goto L183
	}
L183:
	;
	v744 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v744
	v748 = F_find_among_b(m, l0, int32(_a_F_russian_UTF_8_stem_10), int32(8))
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
	v768 = F_memcmp(m, v765+v761-v757, int32(_a_F_russian_UTF_8_stem_11), v757)
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
	v790 = F_memcmp(m, v787+v777-v779, int32(_a_F_russian_UTF_8_stem_12), v779)
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
	v1077 = v797
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
	v1077 = v801
	goto L1
L204:
	;
	v878 = int32(base.Ui32(v874) >> (uint(int32(31)) % 32))
	if v874 != 0 {
		goto L226
	} else {
		goto L227
	}
L205:
	;
	if v814 == int32(0) {
		v874 = v809
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
	v874 = int32(1)
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
	v834 = F_memcmp(m, v831+v827-v823, int32(_a_F_russian_UTF_8_stem_13), v823)
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
		v874 = v809
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
	v856 = F_memcmp(m, v853+v843-v845, int32(_a_F_russian_UTF_8_stem_14), v845)
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
	v874 = v863
	goto L204
L224:
	;
	if v867 < int32(0) {
		v874 = v867
		goto L204
	} else {
		goto L225
	}
L225:
	;
	goto L207
L226:
	;
	v880 = v878
	goto L228
L227:
	;
	v880 = int32(13)
	goto L228
L228:
	;
	if v880 == int32(0) {
		goto L147
	} else {
		goto L229
	}
L229:
	;
	if v880 == int32(13) {
		goto L230
	} else {
		goto L231
	}
L230:
	;
	v885 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v886 = v885 - v806
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v886
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v886
	v893 = F_find_among_b(m, l0, int32(_a_F_russian_UTF_8_stem_15), int32(36))
	mBase = m.M
	v894 = m.ExcPending
	if v894 != 0 {
		goto L40
	} else {
		goto L234
	}
L231:
	;
	v914 = v878
	v915 = v874
	goto L232
L232:
	;
	if v914 != 0 {
		v1077 = v915
		goto L1
	} else {
		goto L244
	}
L233:
	;
	if v906 == int32(0) {
		goto L147
	} else {
		goto L240
	}
L234:
	;
	if v893 == int32(0) {
		v906 = int32(0)
		goto L233
	} else {
		goto L235
	}
L235:
	;
	v897 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v897
	v900 = F_slice_del(m, l0)
	mBase = m.M
	v901 = m.ExcPending
	if v901 != 0 {
		goto L40
	} else {
		goto L236
	}
L236:
	;
	if int32(0) <= v900 {
		goto L237
	} else {
		goto L238
	}
L237:
	;
	v904 = int32(1)
	goto L239
L238:
	;
	v904 = v900
	goto L239
L239:
	;
	v906 = v904
	goto L233
L240:
	;
	if v906 < int32(0) {
		goto L241
	} else {
		goto L242
	}
L241:
	;
	v911 = v906
	goto L243
L242:
	;
	v911 = v874
	goto L243
L243:
	;
	v914 = int32(base.Ui32(v906) >> (uint(int32(31)) % 32))
	v915 = v911
	goto L232
L244:
	;
	goto L147
L245:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v951
	v955 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v951-int32(5) <= v955 {
		goto L255
	} else {
		goto L256
	}
L246:
	;
	if v938 == int32(0) {
		goto L250
	} else {
		goto L251
	}
L247:
	;
	goto L246
L248:
	;
	v931 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v934 = F_memcmp(m, v931+v920-v923, int32(_a_F_russian_UTF_8_stem_16), v923)
	mBase = m.M
	if v934 != 0 {
		v938 = v925
		goto L247
	} else {
		goto L249
	}
L249:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v920 - v923
	v938 = int32(1)
	goto L247
L250:
	;
	v941 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v941
	v951 = v941
	v953 = v941
	goto L245
L251:
	;
	goto L252
L252:
	;
	v943 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v943
	v945 = F_slice_del(m, l0)
	mBase = m.M
	v946 = m.ExcPending
	if v946 != 0 {
		goto L40
	} else {
		goto L253
	}
L253:
	;
	if v945 < int32(0) {
		v1077 = v945
		goto L1
	} else {
		goto L254
	}
L254:
	;
	v949 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v950 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v951 = v949
	v953 = v950
	goto L245
L255:
	;
	v982 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v984 = v982 + (v951 - v953)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v984
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v984
	v989 = F_find_among_b(m, l0, int32(_a_F_russian_UTF_8_stem_17), int32(4))
	mBase = m.M
	v990 = m.ExcPending
	if v990 != 0 {
		goto L40
	} else {
		goto L264
	}
L256:
	;
	v959 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v963 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v959+v951-int32(1)))))
	switch v963 - int32(130) {
	case 0, 10:
		goto L257
	default:
		goto L255
	}
L257:
	;
	v968 = F_find_among_b(m, l0, int32(_a_F_russian_UTF_8_stem_18), int32(2))
	mBase = m.M
	v969 = m.ExcPending
	if v969 != 0 {
		goto L40
	} else {
		goto L258
	}
L258:
	;
	if v968 == int32(0) {
		goto L255
	} else {
		goto L259
	}
L259:
	;
	v972 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v972
	v974 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v975 = *(*int32)(unsafe.Add(mBase, uint32(v974)))
	if v972 < v975 {
		goto L255
	} else {
		goto L260
	}
L260:
	;
	v977 = F_slice_del(m, l0)
	mBase = m.M
	v978 = m.ExcPending
	if v978 != 0 {
		goto L40
	} else {
		goto L261
	}
L261:
	;
	if v977 < int32(0) {
		v1077 = v977
		goto L1
	} else {
		goto L262
	}
L262:
	;
	goto L255
L263:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v137
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v137
	v1077 = int32(1)
	goto L1
L264:
	;
	if v989 == int32(0) {
		goto L263
	} else {
		goto L265
	}
L265:
	;
	v993 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v993
	switch v989 - int32(1) {
	case 0:
		goto L268
	case 1:
		goto L267
	case 2:
		goto L266
	default:
		goto L263
	}
L266:
	;
	v1067 = F_slice_del(m, l0)
	mBase = m.M
	v1068 = m.ExcPending
	if v1068 != 0 {
		goto L40
	} else {
		goto L290
	}
L267:
	;
	v1045 = int32(2)
	v1047 = int32(0)
	v1049 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1050 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1049-v1050 < v1045 {
		v1060 = v1047
		goto L284
	} else {
		goto L285
	}
L268:
	;
	v997 = F_slice_del(m, l0)
	mBase = m.M
	v998 = m.ExcPending
	if v998 != 0 {
		goto L40
	} else {
		goto L269
	}
L269:
	;
	if v997 < int32(0) {
		v1077 = v997
		goto L1
	} else {
		goto L270
	}
L270:
	;
	v1001 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1001
	v1003 = int32(2)
	v1005 = int32(0)
	v1008 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1001-v1008 < v1003 {
		v1018 = v1005
		goto L272
	} else {
		goto L273
	}
L271:
	;
	if v1018 == int32(0) {
		goto L263
	} else {
		goto L275
	}
L272:
	;
	goto L271
L273:
	;
	v1011 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1014 = F_memcmp(m, v1011+v1001-v1003, int32(_a_F_russian_UTF_8_stem_19), v1003)
	mBase = m.M
	if v1014 != 0 {
		v1018 = v1005
		goto L272
	} else {
		goto L274
	}
L274:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1001 - v1003
	v1018 = int32(1)
	goto L272
L275:
	;
	v1021 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1021
	v1023 = int32(2)
	v1025 = int32(0)
	v1028 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1021-v1028 < v1023 {
		v1038 = v1025
		goto L277
	} else {
		goto L278
	}
L276:
	;
	if v1038 == int32(0) {
		goto L263
	} else {
		goto L280
	}
L277:
	;
	goto L276
L278:
	;
	v1031 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1034 = F_memcmp(m, v1031+v1021-v1023, int32(_a_F_russian_UTF_8_stem_20), v1023)
	mBase = m.M
	if v1034 != 0 {
		v1038 = v1025
		goto L277
	} else {
		goto L279
	}
L279:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1021 - v1023
	v1038 = int32(1)
	goto L277
L280:
	;
	v1041 = F_slice_del(m, l0)
	mBase = m.M
	v1042 = m.ExcPending
	if v1042 != 0 {
		goto L40
	} else {
		goto L281
	}
L281:
	;
	if int32(0) <= v1041 {
		goto L263
	} else {
		goto L282
	}
L282:
	;
	v1077 = v1041
	goto L1
L283:
	;
	if v1060 == int32(0) {
		goto L263
	} else {
		goto L287
	}
L284:
	;
	goto L283
L285:
	;
	v1053 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1056 = F_memcmp(m, v1053+v1049-v1045, int32(_a_F_russian_UTF_8_stem_21), v1045)
	mBase = m.M
	if v1056 != 0 {
		v1060 = v1047
		goto L284
	} else {
		goto L286
	}
L286:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1049 - v1045
	v1060 = int32(1)
	goto L284
L287:
	;
	v1063 = F_slice_del(m, l0)
	mBase = m.M
	v1064 = m.ExcPending
	if v1064 != 0 {
		goto L40
	} else {
		goto L288
	}
L288:
	;
	if int32(0) <= v1063 {
		goto L263
	} else {
		goto L289
	}
L289:
	;
	v1077 = v1063
	goto L1
L290:
	;
	if v1067 < int32(0) {
		v1077 = v1067
		goto L1
	} else {
		goto L291
	}
L291:
	;
	goto L263
}
