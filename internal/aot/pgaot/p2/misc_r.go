package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
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
				v39 = int32(2987)
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
						v39 = int32(2999)
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
	F_errfinish(m, int32(_a_F_ReleaseOneSerializableXact_5), int32(3820), int32(_a_F_ReleaseOneSerializableXact_6))
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
func F_ReleaseSemaphores(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	v3 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_ReleaseSemaphores[0]))
	if v3 < v6 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v9 = v3
	goto L4
L2:
	;
	goto L3
L3:
	;
	return
L4:
	;
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_ReleaseSemaphores[1]))
	v18 = int32(0)
	v19 = m.Env.Pgmem_sem(m, int32(1), v14+v9<<(uint(int32(7))%32), v18)
	mBase = m.M
	if v18 <= v19 {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	goto L3
L6:
	;
	v46 = v9 + int32(1)
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_ReleaseSemaphores[0]))
	if v46 < v48 {
		v9 = v46
		goto L4
	} else {
		goto L17
	}
L7:
	;
	if int32(0) <= v27 {
		goto L6
	} else {
		goto L11
	}
L8:
	;
	v27 = v19
	goto L7
L9:
	;
	goto L10
L10:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ReleaseSemaphores[2])) = int32(0) - v19
	v27 = int32(-1)
	goto L7
L11:
	;
	v32 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	return
L13:
	;
	if v32 == int32(0) {
		goto L6
	} else {
		goto L14
	}
L14:
	;
	F_errmsg_internal(m, int32(_a_F_ReleaseSemaphores_0), int32(0))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	F_errfinish(m, int32(_a_F_ReleaseSemaphores_1), int32(156), int32(_a_F_ReleaseSemaphores_2))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L12
	} else {
		goto L16
	}
L16:
	;
	goto L6
L17:
	;
	goto L5
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
	var v373 int64
	_ = v373
	var v375 int64
	_ = v375
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v388 int64
	_ = v388
	var v391 int32
	_ = v391
	var v397 int32
	_ = v397
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v419 int32
	_ = v419
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
	F_errfinish(m, int32(_a_F_RemoveOldXlogFiles_4), int32(3947), int32(_a_F_RemoveOldXlogFiles_5))
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
	v419 = m.ExcPending
	if v419 != 0 {
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
	v401 = F_ReadDir(m, v71, int32(_a_F_RemoveOldXlogFiles_3))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
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
	v373 = base.I64_div_u_s(int64(4294967296), v357)
	v375 = v370 + v371*v373
	v377 = *(*int32)(unsafe.Add(mBase, _c_F_RemoveOldXlogFiles[8]))
	v380 = base.AtomicRmwXchg32(m, v377, int32(440), int32(1))
	if v380 != 0 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	F_s_lock(m, v377+int32(440), int32(_a_F_RemoveOldXlogFiles_8))
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L1
	} else {
		goto L88
	}
L86:
	;
	goto L87
L87:
	;
	v387 = *(*int32)(unsafe.Add(mBase, _c_F_RemoveOldXlogFiles[8]))
	v388 = *(*int64)(unsafe.Add(mBase, uint32(v387)+224))
	if base.Ui64(v388) < base.Ui64(v375) {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	goto L87
L89:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v387)+224)) = v375
	goto L91
L90:
	;
	goto L91
L91:
	;
	v391 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v387)+440)), uint32(v391))
	F_RemoveXlogFile(m, v107, v96, v18+int32(56), l3)
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	goto L22
L93:
	;
	if v401 != 0 {
		v107 = v401
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
	var v347 int32
	_ = v347
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
	F_errfinish(m, int32(_a_F_RestoreArchivedFile_2), int32(113), int32(_a_F_RestoreArchivedFile_3))
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
	F_errfinish(m, int32(_a_F_RestoreArchivedFile_2), int32(160), int32(_a_F_RestoreArchivedFile_3))
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
	*(*int32)(unsafe.Add(mBase, uint32(v169))) = int32(134217779)
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
	F_errfinish(m, int32(_a_F_RestoreArchivedFile_2), int32(217), int32(_a_F_RestoreArchivedFile_3))
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
	F_errfinish(m, int32(_a_F_RestoreArchivedFile_2), int32(224), int32(_a_F_RestoreArchivedFile_3))
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
	v347 = F_errdetail(m, int32(_a_F_RestoreArchivedFile_10), int32(0))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L8
	} else {
		goto L103
	}
L103:
	;
	F_errfinish(m, int32(_a_F_RestoreArchivedFile_2), int32(237), int32(_a_F_RestoreArchivedFile_3))
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
	F_errfinish(m, int32(_a_F_RestoreArchivedFile_2), int32(270), int32(_a_F_RestoreArchivedFile_3))
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
	F_errfinish(m, int32(_a_F_RestoreArchivedFile_2), int32(121), int32(_a_F_RestoreArchivedFile_3))
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
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l2)+224))
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
										v84 = *(*int32)(unsafe.Add(mBase, uint32(l2)+76))
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
											v84 = *(*int32)(unsafe.Add(mBase, uint32(l2)+76))
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
												v84 = *(*int32)(unsafe.Add(mBase, uint32(l2)+76))
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
													v84 = *(*int32)(unsafe.Add(mBase, uint32(l2)+76))
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
														v84 = *(*int32)(unsafe.Add(mBase, uint32(l2)+76))
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
															v84 = *(*int32)(unsafe.Add(mBase, uint32(l2)+76))
															v90 = base.B2i32(v83 == v84)
														}
													}
													return v90
												} else {
													v54 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
													v55 = *(*int32)(unsafe.Add(mBase, uint32(l2)+76))
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
																v84 = *(*int32)(unsafe.Add(mBase, uint32(l2)+76))
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
																	v84 = *(*int32)(unsafe.Add(mBase, uint32(l2)+76))
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
																		v84 = *(*int32)(unsafe.Add(mBase, uint32(l2)+76))
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
																					v84 = *(*int32)(unsafe.Add(mBase, uint32(l2)+76))
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
																							v84 = *(*int32)(unsafe.Add(mBase, uint32(l2)+76))
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
								v84 = *(*int32)(unsafe.Add(mBase, uint32(l2)+76))
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
										v84 = *(*int32)(unsafe.Add(mBase, uint32(l2)+76))
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
											v84 = *(*int32)(unsafe.Add(mBase, uint32(l2)+76))
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
												v84 = *(*int32)(unsafe.Add(mBase, uint32(l2)+76))
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
													v84 = *(*int32)(unsafe.Add(mBase, uint32(l2)+76))
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
														v84 = *(*int32)(unsafe.Add(mBase, uint32(l2)+76))
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
															v84 = *(*int32)(unsafe.Add(mBase, uint32(l2)+76))
															v90 = base.B2i32(v83 == v84)
														}
													}
													return v90
												} else {
													v54 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
													v55 = *(*int32)(unsafe.Add(mBase, uint32(l2)+76))
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
																v84 = *(*int32)(unsafe.Add(mBase, uint32(l2)+76))
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
																	v84 = *(*int32)(unsafe.Add(mBase, uint32(l2)+76))
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
																		v84 = *(*int32)(unsafe.Add(mBase, uint32(l2)+76))
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
																					v84 = *(*int32)(unsafe.Add(mBase, uint32(l2)+76))
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
																							v84 = *(*int32)(unsafe.Add(mBase, uint32(l2)+76))
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
								v84 = *(*int32)(unsafe.Add(mBase, uint32(l2)+76))
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
func F_RoleidCallback(m *base.Module, l0 int64, l1 int32, l2 int32) {
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
func F_r_C_1(m *base.Module, l0 int32) int32 {
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
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v35 int32
	_ = v35
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	v2 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v7 = int32(2)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v5-v12 < v7 {
		v22 = v2
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v18 = F_memcmp(m, v15+v5-v7, int32(_a_F_r_C_1_0), v7)
		mBase = m.M
		if v18 != 0 {
			v22 = v2
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v5 - v7
			v22 = int32(1)
		}
	}
	if v22 != 0 {
		v81 = v2
	} else {
		v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v24 = v6 - v5
		v25 = v23 - v24
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v25
		v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v25 <= v35 {
			v75 = int32(-1)
		} else {
			v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46+v25-int32(1)))))
			if int32(252) < v50 {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v25 - int32(1)
				v72 = int32(0)
			} else {
				v52 = v50 - int32(97)
				if v52 < int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v25 - int32(1)
					v72 = int32(0)
				} else {
					v55 = int32(1)
					v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v52)>>(uint(int32(3))%32)))+uint32(_c_F_r_C_1[0]))))
					if int32(base.Ui32(v59)>>(uint(v52&int32(7))%32))&v55 != 0 {
						v72 = v55
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v25 - int32(1)
						v72 = int32(0)
					}
				}
			}
			v75 = v72
		}
		if v75 != 0 {
			v81 = v2
		} else {
			v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v76 - v24
			v81 = int32(1)
		}
	}
	return v81
}
func F_r_Suffix_Noun_Step2b(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
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
	var v22 int32
	_ = v22
	v5 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v6
	v8 = F_eq_s_b(m, l0, int32(4), int32(_a_F_r_Suffix_Noun_Step2b_0))
	mBase = m.M
	if v8 == v5 {
		v22 = v5
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v11
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v14 = F_len_utf8(m, v13)
		mBase = m.M
		if v14 < int32(5) {
			v22 = v5
		} else {
			v17 = F_slice_del(m, l0)
			mBase = m.M
			if int32(0) <= v17 {
				v20 = int32(1)
			} else {
				v20 = v17
			}
			v22 = v20
		}
	}
	return v22
}
func F_r_Suffix_Noun_Step2c1(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
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
	var v22 int32
	_ = v22
	v5 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v6
	v8 = F_eq_s_b(m, l0, int32(2), int32(_a_F_r_Suffix_Noun_Step2c1_0))
	mBase = m.M
	if v8 == v5 {
		v22 = v5
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v11
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v14 = F_len_utf8(m, v13)
		mBase = m.M
		if v14 < int32(4) {
			v22 = v5
		} else {
			v17 = F_slice_del(m, l0)
			mBase = m.M
			if int32(0) <= v17 {
				v20 = int32(1)
			} else {
				v20 = v17
			}
			v22 = v20
		}
	}
	return v22
}
func F_r_Suffix_Verb_Step2a(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v101 int32
	_ = v101
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
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v291 int32
	_ = v291
	var v296 int32
	_ = v296
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
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
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v373 int32
	_ = v373
	v2 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v4
	v9 = F_find_among_b(m, l0, int32(_a_F_r_Suffix_Verb_Step2a_0), int32(11), v2)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		if v9 == int32(0) {
			v373 = v2
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v15
			switch v9 - int32(1) {
			case 0:
				v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v20 = int32(0)
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v19-int32(4))))
				if v27 == v20 {
					v101 = int32(0)
				} else {
					v32 = v27 & int32(3)
					if base.Ui32(v27) < base.Ui32(int32(4)) {
						v68 = v19
						v69 = int32(0)
						v74 = v68
						v75 = v69
						v79 = v20
						for {
							v80 = int32(*(*int8)(unsafe.Add(mBase, uint32(v74))))
							v83 = v75 + base.B2i32(int32(-65) < v80)
							v84 = int32(1)
							v87 = v79 + v84
							if v87 != v32 {
								v74 = v74 + v84
								v75 = v83
								v79 = v87
								continue
							} else {
								break
							}
							break
						}
						v90 = v83
					} else {
						v39 = v19
						v40 = int32(0)
						v43 = v20
						for {
							v45 = int32(*(*int8)(unsafe.Add(mBase, uint32(v39))))
							v46 = int32(-65)
							v49 = int32(*(*int8)(unsafe.Add(mBase, uint32(v39)+1)))
							v53 = int32(*(*int8)(unsafe.Add(mBase, uint32(v39)+2)))
							v57 = int32(*(*int8)(unsafe.Add(mBase, uint32(v39)+3)))
							v60 = v40 + base.B2i32(v46 < v45) + base.B2i32(v46 < v49) + base.B2i32(v46 < v53) + base.B2i32(v46 < v57)
							v61 = int32(4)
							v62 = v39 + v61
							v64 = v43 + v61
							if v64 != v27&int32(-4) {
								v39 = v62
								v40 = v60
								v43 = v64
								continue
							} else {
								break
							}
							break
						}
						if v32 == int32(0) {
							v90 = v60
						} else {
							v68 = v62
							v69 = v60
							v74 = v68
							v75 = v69
							v79 = v20
							for {
								v80 = int32(*(*int8)(unsafe.Add(mBase, uint32(v74))))
								v83 = v75 + base.B2i32(int32(-65) < v80)
								v84 = int32(1)
								v87 = v79 + v84
								if v87 != v32 {
									v74 = v74 + v84
									v75 = v83
									v79 = v87
									continue
								} else {
									break
								}
								break
							}
							v90 = v83
						}
					}
					v101 = v90
				}
				if v101 < int32(4) {
					v373 = v2
				} else {
					v104 = F_slice_del(m, l0)
					mBase = m.M
					if int32(0) <= v104 {
						v373 = int32(1)
					} else {
						v373 = v104
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
					v373 = v2
				} else {
					v192 = F_slice_del(m, l0)
					mBase = m.M
					if int32(0) <= v192 {
						v373 = int32(1)
					} else {
						v373 = v192
					}
				}
			case 2:
				v195 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v196 = int32(0)
				v203 = *(*int32)(unsafe.Add(mBase, uint32(v195-int32(4))))
				if v203 == v196 {
					v277 = int32(0)
				} else {
					v208 = v203 & int32(3)
					if base.Ui32(v203) < base.Ui32(int32(4)) {
						v244 = v195
						v245 = int32(0)
						v250 = v244
						v251 = v245
						v255 = v196
						for {
							v256 = int32(*(*int8)(unsafe.Add(mBase, uint32(v250))))
							v259 = v251 + base.B2i32(int32(-65) < v256)
							v260 = int32(1)
							v263 = v255 + v260
							if v263 != v208 {
								v250 = v250 + v260
								v251 = v259
								v255 = v263
								continue
							} else {
								break
							}
							break
						}
						v266 = v259
					} else {
						v215 = v195
						v216 = int32(0)
						v219 = v196
						for {
							v221 = int32(*(*int8)(unsafe.Add(mBase, uint32(v215))))
							v222 = int32(-65)
							v225 = int32(*(*int8)(unsafe.Add(mBase, uint32(v215)+1)))
							v229 = int32(*(*int8)(unsafe.Add(mBase, uint32(v215)+2)))
							v233 = int32(*(*int8)(unsafe.Add(mBase, uint32(v215)+3)))
							v236 = v216 + base.B2i32(v222 < v221) + base.B2i32(v222 < v225) + base.B2i32(v222 < v229) + base.B2i32(v222 < v233)
							v237 = int32(4)
							v238 = v215 + v237
							v240 = v219 + v237
							if v240 != v203&int32(-4) {
								v215 = v238
								v216 = v236
								v219 = v240
								continue
							} else {
								break
							}
							break
						}
						if v208 == int32(0) {
							v266 = v236
						} else {
							v244 = v238
							v245 = v236
							v250 = v244
							v251 = v245
							v255 = v196
							for {
								v256 = int32(*(*int8)(unsafe.Add(mBase, uint32(v250))))
								v259 = v251 + base.B2i32(int32(-65) < v256)
								v260 = int32(1)
								v263 = v255 + v260
								if v263 != v208 {
									v250 = v250 + v260
									v251 = v259
									v255 = v263
									continue
								} else {
									break
								}
								break
							}
							v266 = v259
						}
					}
					v277 = v266
				}
				if v277 < int32(6) {
					v373 = v2
				} else {
					v280 = F_slice_del(m, l0)
					mBase = m.M
					if int32(0) <= v280 {
						v373 = int32(1)
					} else {
						v373 = v280
					}
				}
			case 3:
				v283 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v284 = int32(0)
				v291 = *(*int32)(unsafe.Add(mBase, uint32(v283-int32(4))))
				if v291 == v284 {
					v365 = int32(0)
				} else {
					v296 = v291 & int32(3)
					if base.Ui32(v291) < base.Ui32(int32(4)) {
						v332 = v283
						v333 = int32(0)
						v338 = v332
						v339 = v333
						v343 = v284
						for {
							v344 = int32(*(*int8)(unsafe.Add(mBase, uint32(v338))))
							v347 = v339 + base.B2i32(int32(-65) < v344)
							v348 = int32(1)
							v351 = v343 + v348
							if v351 != v296 {
								v338 = v338 + v348
								v339 = v347
								v343 = v351
								continue
							} else {
								break
							}
							break
						}
						v354 = v347
					} else {
						v303 = v283
						v304 = int32(0)
						v307 = v284
						for {
							v309 = int32(*(*int8)(unsafe.Add(mBase, uint32(v303))))
							v310 = int32(-65)
							v313 = int32(*(*int8)(unsafe.Add(mBase, uint32(v303)+1)))
							v317 = int32(*(*int8)(unsafe.Add(mBase, uint32(v303)+2)))
							v321 = int32(*(*int8)(unsafe.Add(mBase, uint32(v303)+3)))
							v324 = v304 + base.B2i32(v310 < v309) + base.B2i32(v310 < v313) + base.B2i32(v310 < v317) + base.B2i32(v310 < v321)
							v325 = int32(4)
							v326 = v303 + v325
							v328 = v307 + v325
							if v328 != v291&int32(-4) {
								v303 = v326
								v304 = v324
								v307 = v328
								continue
							} else {
								break
							}
							break
						}
						if v296 == int32(0) {
							v354 = v324
						} else {
							v332 = v326
							v333 = v324
							v338 = v332
							v339 = v333
							v343 = v284
							for {
								v344 = int32(*(*int8)(unsafe.Add(mBase, uint32(v338))))
								v347 = v339 + base.B2i32(int32(-65) < v344)
								v348 = int32(1)
								v351 = v343 + v348
								if v351 != v296 {
									v338 = v338 + v348
									v339 = v347
									v343 = v351
									continue
								} else {
									break
								}
								break
							}
							v354 = v347
						}
					}
					v365 = v354
				}
				if v365 < int32(6) {
					v373 = v2
				} else {
					v368 = F_slice_del(m, l0)
					mBase = m.M
					if v368 < int32(0) {
						v373 = v368
					} else {
						v373 = int32(1)
					}
				}
			default:
				v373 = int32(1)
			}
		}
		return v373
	}
}
func F_r_V_2(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
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
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v128 int32
	_ = v128
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v6 = v4 - v5
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v5 <= v20 {
		v128 = int32(-1)
		v135 = v128
	} else {
		v37 = int32(1)
		v38 = v5 - v37
		v40 = int32(*(*int8)(unsafe.Add(mBase, uint32(v21+v38))))
		v42 = v40 & int32(255)
		if base.B2i32(v38 == v20)|base.B2i32(int32(0) <= v40) != 0 {
			v100 = v42
			v104 = v37
		} else {
			v49 = v42 & int32(63)
			v51 = v5 - int32(2)
			v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21+v51))))
			v55 = v53 << (uint(int32(6)) % 32)
			if base.B2i32(v51 != v20)&base.B2i32(base.Ui32(v53) < base.Ui32(int32(192))) == int32(0) {
				v100 = v55&int32(1984) | v49
				v104 = int32(2)
			} else {
				v68 = v55&int32(4032) | v49
				v70 = v5 - int32(3)
				v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21+v70))))
				if base.B2i32(v70 != v20)&base.B2i32(base.Ui32(v72) < base.Ui32(int32(224))) == int32(0) {
					v100 = v72<<(uint(int32(12))%32)&int32(_a_F_r_V_2_0) | v68
					v104 = int32(3)
				} else {
					v90 = int32(4)
					v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5+v21-v90))))
					v100 = v72<<(uint(int32(12))%32)&int32(_a_F_r_V_2_1) | v92&int32(7)<<(uint(int32(18))%32) | v68
					v104 = v90
				}
			}
		}
		if int32(252) < v100 {
			v135 = v104
		} else {
			v106 = v100 - int32(97)
			if v106 < int32(0) {
				v135 = v104
			} else {
				v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v106)>>(uint(int32(3))%32)))+uint32(_c_F_r_V_2[0]))))
				if int32(base.Ui32(v112)>>(uint(v106&int32(7))%32))&int32(1) == int32(0) {
					v135 = v104
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v5 - v104
					v128 = int32(0)
					v135 = v128
				}
			}
		}
	}
	if v135 != 0 {
		v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v137 = v136 - v6
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v137
		v139 = int32(2)
		v141 = int32(0)
		v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v137-v144 < v139 {
			v154 = v141
		} else {
			v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v150 = F_memcmp(m, v147+v137-v139, int32(_a_F_r_V_2_2), v139)
			mBase = m.M
			if v150 != 0 {
				v154 = v141
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v137 - v139
				v154 = int32(1)
			}
		}
		if v154 == int32(0) {
			v161 = int32(0)
		} else {
			v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v157 - v6
			v161 = int32(1)
		}
	} else {
		v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v157 - v6
		v161 = int32(1)
	}
	return v161
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
	var v2 int32
	_ = v2
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
	var v34 int32
	_ = v34
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
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	v2 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v7 < v8 {
		v119 = v2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v119
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v7
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v8
	v15 = v7 - int32(1)
	if v8 < v15 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v34 = F_find_among_b(m, l0, int32(_a_F_r_consonant_pair_2_0), int32(4), int32(0))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
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
	if v34 == int32(0) {
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
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v44
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v48 = v46 + (v7 - v10)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v48
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L15
L13:
	;
	if v102 < int32(0) {
		v119 = v2
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
	v57 = v48
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
	if int32(0) <= v108 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v114 = int32(1)
	goto L35
L34:
	;
	v114 = v108 >> (uint(int32(31)) % 32) & v108
	goto L35
L35:
	;
	v119 = v114
	goto L1
}
func F_r_remove_second_order_prefix_1(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
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
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	v2 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v5
	v8 = v5 + int32(1)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v9 <= v8 {
		v177 = v2
		return v177
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11+v8))))
		if v13 != int32(101) {
			v177 = v2
			return v177
		} else {
			v16 = int32(2)
			v20 = F_find_among(m, l0, int32(_a_F_r_remove_second_order_prefix_1_0), v16, int32(0))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				switch v20 {
				case 0:
					v177 = v20
				case 1:
					v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					if v24 == v25 {
						v155 = v24
						v156 = v16
						*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v155
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v155
						v161 = v156
						*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v161
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27+v24))))
						switch v29 - int32(108) {
						case 0:
							v33 = v24 + int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v33
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v33
							v36 = int32(4)
							v38 = int32(0)
							v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							if v40-v33 < v36 {
								v50 = v38
							} else {
								v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v46 = F_memcmp(m, v44+v33, int32(_a_F_r_remove_second_order_prefix_1_1), v36)
								mBase = m.M
								if v46 != 0 {
									v50 = v38
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v24 + int32(5)
									v50 = int32(1)
								}
							}
							if v50 == int32(0) {
								v155 = v24
								v156 = v16
								*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v155
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v155
								v161 = v156
								*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v161
							} else {
							}
						default:
							v155 = v24
							v156 = v16
							*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v155
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v155
							v161 = v156
							*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v161
						case 6:
							v155 = v24 + int32(1)
							v156 = v16
							*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v155
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v155
							v161 = v156
							*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v161
						}
					}
					v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
					v168 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v167 - v168
					v172 = F_slice_del(m, l0)
					mBase = m.M
					if int32(0) <= v172 {
						v175 = v168
					} else {
						v175 = v172
					}
					v177 = v175
				case 2:
					v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					if v53 == v54 {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v53
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v53
						v87 = int32(0)
						v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						if v96 < v53 {
							v98 = v53
						} else {
							v98 = v96
						}
						if v53 == v98 {
							v136 = int32(-1)
						} else {
							v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109+v53))))
							if int32(117) < v111 {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v53 + int32(1)
								v133 = int32(0)
							} else {
								v113 = v111 - int32(97)
								if v113 < int32(0) {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v53 + int32(1)
									v133 = int32(0)
								} else {
									v116 = int32(1)
									v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v113)>>(uint(int32(3))%32)))+uint32(_c_F_r_remove_second_order_prefix_1[0]))))
									if int32(base.Ui32(v120)>>(uint(v113&int32(7))%32))&v116 != 0 {
										v133 = v116
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v53 + int32(1)
										v133 = int32(0)
									}
								}
							}
							v136 = v133
						}
						if v136 != 0 {
							v177 = v87
						} else {
							v138 = int32(2)
							v140 = int32(0)
							v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							if v142-v143 < v138 {
								v152 = v140
							} else {
								v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v148 = F_memcmp(m, v146+v143, int32(_a_F_r_remove_second_order_prefix_1_2), v138)
								mBase = m.M
								if v148 != 0 {
									v152 = v140
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v138 + v143
									v152 = int32(1)
								}
							}
							if v152 != 0 {
								v161 = int32(4)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v161
								v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
								v168 = int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v167 - v168
								v172 = F_slice_del(m, l0)
								mBase = m.M
								if int32(0) <= v172 {
									v175 = v168
								} else {
									v175 = v172
								}
								v177 = v175
							} else {
								v177 = v87
							}
						}
					} else {
						v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56+v53))))
						switch v58 - int32(108) {
						case 0:
							v65 = v53 + int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v65
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v65
							v68 = int32(4)
							v71 = int32(0)
							v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							if v73-v65 < v68 {
								v83 = v71
							} else {
								v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v79 = F_memcmp(m, v77+v65, int32(_a_F_r_remove_second_order_prefix_1_3), v68)
								mBase = m.M
								if v79 != 0 {
									v83 = v71
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v53 + int32(5)
									v83 = int32(1)
								}
							}
							if v83 != 0 {
								v161 = v68
								*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v161
								v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
								v168 = int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v167 - v168
								v172 = F_slice_del(m, l0)
								mBase = m.M
								if int32(0) <= v172 {
									v175 = v168
								} else {
									v175 = v172
								}
								v177 = v175
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v53
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v53
								v87 = int32(0)
								v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								if v96 < v53 {
									v98 = v53
								} else {
									v98 = v96
								}
								if v53 == v98 {
									v136 = int32(-1)
								} else {
									v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109+v53))))
									if int32(117) < v111 {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v53 + int32(1)
										v133 = int32(0)
									} else {
										v113 = v111 - int32(97)
										if v113 < int32(0) {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v53 + int32(1)
											v133 = int32(0)
										} else {
											v116 = int32(1)
											v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v113)>>(uint(int32(3))%32)))+uint32(_c_F_r_remove_second_order_prefix_1[0]))))
											if int32(base.Ui32(v120)>>(uint(v113&int32(7))%32))&v116 != 0 {
												v133 = v116
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v53 + int32(1)
												v133 = int32(0)
											}
										}
									}
									v136 = v133
								}
								if v136 != 0 {
									v177 = v87
								} else {
									v138 = int32(2)
									v140 = int32(0)
									v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									if v142-v143 < v138 {
										v152 = v140
									} else {
										v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										v148 = F_memcmp(m, v146+v143, int32(_a_F_r_remove_second_order_prefix_1_2), v138)
										mBase = m.M
										if v148 != 0 {
											v152 = v140
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v138 + v143
											v152 = int32(1)
										}
									}
									if v152 != 0 {
										v161 = int32(4)
										*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v161
										v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
										v168 = int32(1)
										*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v167 - v168
										v172 = F_slice_del(m, l0)
										mBase = m.M
										if int32(0) <= v172 {
											v175 = v168
										} else {
											v175 = v172
										}
										v177 = v175
									} else {
										v177 = v87
									}
								}
							}
						default:
							*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v53
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v53
							v87 = int32(0)
							v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							if v96 < v53 {
								v98 = v53
							} else {
								v98 = v96
							}
							if v53 == v98 {
								v136 = int32(-1)
							} else {
								v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109+v53))))
								if int32(117) < v111 {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v53 + int32(1)
									v133 = int32(0)
								} else {
									v113 = v111 - int32(97)
									if v113 < int32(0) {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v53 + int32(1)
										v133 = int32(0)
									} else {
										v116 = int32(1)
										v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v113)>>(uint(int32(3))%32)))+uint32(_c_F_r_remove_second_order_prefix_1[0]))))
										if int32(base.Ui32(v120)>>(uint(v113&int32(7))%32))&v116 != 0 {
											v133 = v116
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v53 + int32(1)
											v133 = int32(0)
										}
									}
								}
								v136 = v133
							}
							if v136 != 0 {
								v177 = v87
							} else {
								v138 = int32(2)
								v140 = int32(0)
								v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								if v142-v143 < v138 {
									v152 = v140
								} else {
									v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									v148 = F_memcmp(m, v146+v143, int32(_a_F_r_remove_second_order_prefix_1_2), v138)
									mBase = m.M
									if v148 != 0 {
										v152 = v140
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v138 + v143
										v152 = int32(1)
									}
								}
								if v152 != 0 {
									v161 = int32(4)
									*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v161
									v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
									v168 = int32(1)
									*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v167 - v168
									v172 = F_slice_del(m, l0)
									mBase = m.M
									if int32(0) <= v172 {
										v175 = v168
									} else {
										v175 = v172
									}
									v177 = v175
								} else {
									v177 = v87
								}
							}
						case 6:
							v155 = v53 + int32(1)
							v156 = int32(4)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v155
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v155
							v161 = v156
							*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v161
							v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
							v168 = int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v167 - v168
							v172 = F_slice_del(m, l0)
							mBase = m.M
							if int32(0) <= v172 {
								v175 = v168
							} else {
								v175 = v172
							}
							v177 = v175
						}
					}
				default:
					v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
					v168 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v167 - v168
					v172 = F_slice_del(m, l0)
					mBase = m.M
					if int32(0) <= v172 {
						v175 = v168
					} else {
						v175 = v172
					}
					v177 = v175
				}
				return v177
			}
		}
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
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	v2 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v6 <= v7 {
		v26 = v2
	} else {
		v10 = v6 - int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v10
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v10
		if v10 <= v7 {
			v26 = v2
		} else {
			v15 = v6 - int32(2)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v15
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v15
			v19 = F_slice_del(m, l0)
			mBase = m.M
			if int32(0) <= v19 {
				v22 = int32(1)
			} else {
				v22 = v19
			}
			v26 = v22
		}
	}
	return v26
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
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
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
	v88 = m.T0[int32(1262)].(func(*base.Module, int32, int32, int32) int32)(m, v46, v85, int32(100))
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
	v73 = int32(64)
	v76 = F_palloc_mul(m, int32(4), v73)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L7
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v80 = v49 << (uint(int32(1)) % 32)
	v81 = F_repalloc_mul(m, v47, int32(4), v80)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L7
	} else {
		goto L37
	}
L36:
	;
	v83 = v76
	v84 = v73
	goto L31
L37:
	;
	v83 = v81
	v84 = v80
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
	F_pg_qsort(m, v117, v130, int32(4), int32(1285))
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
	var v37 int64
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
							v37 = F_index_getattr_2(m, v16, int32(1), v34, l1+int32(16))
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(l1)+8)) = v37
								m.G0 = v10 + int32(16)
								return
							}
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v16
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
					v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+52))
					v37 = F_index_getattr_2(m, v16, int32(1), v34, l1+int32(16))
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(l1)+8)) = v37
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
	var v33 int64
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
							v33 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v14)+4)))
							*(*int64)(unsafe.Add(mBase, uint32(l1)+8)) = v33
							m.G0 = v9 + int32(16)
							return
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v14
					v33 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v14)+4)))
					*(*int64)(unsafe.Add(mBase, uint32(l1)+8)) = v33
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
func F_record_eq(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
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
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int64
	_ = v89
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
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
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v232 int32
	_ = v232
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v285 int32
	_ = v285
	var v288 int64
	_ = v288
	var v289 int32
	_ = v289
	var v295 int64
	_ = v295
	var v301 int32
	_ = v301
	var v302 int64
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v363 int32
	_ = v363
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
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v382 int32
	_ = v382
	var v387 int32
	_ = v387
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v401 int32
	_ = v401
	var v406 int32
	_ = v406
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v418 int32
	_ = v418
	var v423 int32
	_ = v423
	var v447 int64
	_ = v447
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v473 int32
	_ = v473
	v23 = m.G0
	v25 = v23 - int32(128)
	m.G0 = v25
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v28 = F_pg_detoast_datum(m, v27)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v33 = F_pg_detoast_datum(m, v32)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	F_check_stack_depth(m)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	v39 = F_lookup_rowtype_tupdesc(m, v37, v38)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	v44 = F_lookup_rowtype_tupdesc(m, v42, v43)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+124)) = v28
	v49 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+120)) = v49
	*(*uint16)(unsafe.Add(mBase, uint32(v25)+116)) = uint16(v49)
	v53 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+112)) = v53
	v55 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+108)) = int32(base.Ui32(v47) >> (uint(v55) % 32))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+104)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v25)+100)) = v49
	*(*uint16)(unsafe.Add(mBase, uint32(v25)+96)) = uint16(v49)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+92)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v25)+88)) = int32(base.Ui32(v58) >> (uint(v55) % 32))
	if v46 < v41 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v70 = v41
	goto L9
L8:
	;
	v70 = v46
	goto L9
L9:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+16))
	if v72 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	if v37 != v95 {
		goto L16
	} else {
		goto L17
	}
L11:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v71)+20))
	v83 = F_MemoryContextAlloc(m, v78, v70<<(uint(int32(2))%32)+int32(20))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L14
	}
L12:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	if v75 < v70 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v72)+4))
	v94 = v72
	v95 = v77
	goto L10
L14:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v85)+16)) = v83
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)+16))
	v89 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v88)+4)) = v89
	*(*int32)(unsafe.Add(mBase, uint32(v88))) = v70
	*(*int64)(unsafe.Add(mBase, uint32(v88)+12)) = v89
	v94 = v88
	v95 = int32(0)
	goto L10
L15:
	;
	v151 = F_palloc(m, v41<<(uint(int32(3))%32))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L31
	}
L16:
	;
	v104 = v94 + int32(20)
	v108 = v70 << (uint(int32(2)) % 32)
	if v104&int32(3)|base.B2i32(base.Ui32(int32(1024)) < base.Ui32(v108)) == int32(0) {
		goto L22
	} else {
		goto L23
	}
L17:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v94)+8))
	if v97 != v38 {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v94)+12))
	if v99 != v42 {
		goto L16
	} else {
		goto L19
	}
L19:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v94)+16))
	if v101 == v43 {
		goto L15
	} else {
		goto L20
	}
L20:
	;
	goto L16
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v94)+16)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v94)+12)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v94)+8)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v94)+4)) = v37
	goto L15
L22:
	;
	if v108 == int32(0) {
		goto L21
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	if v108 == int32(0) {
		goto L21
	} else {
		goto L30
	}
L25:
	;
	v118 = v94 + v108 + int32(20)
	v120 = v94 + int32(24)
	if base.Ui32(v120) < base.Ui32(v118) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v122 = v118
	goto L28
L27:
	;
	v122 = v120
	goto L28
L28:
	;
	v129 = (v122-v94-int32(21))&int32(-4) + int32(4)
	if v129 == int32(0) {
		goto L21
	} else {
		goto L29
	}
L29:
	;
	base.MemoryFill(m, v104, int32(0), v129)
	goto L21
L30:
	;
	base.MemoryFill(m, v104, int32(0), v108)
	goto L21
L31:
	;
	v153 = F_palloc(m, v41)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	F_heap_deform_tuple(m, v25+int32(108), v39, v151, v153)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v161 = F_palloc(m, v46<<(uint(int32(3))%32))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v163 = F_palloc(m, v46)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	F_heap_deform_tuple(m, v25+int32(88), v44, v161, v163)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v167 = int32(0)
	v169 = base.B2i32(v167 < v41)
	if v169|base.B2i32(v167 < v46) == v167 {
		goto L43
	} else {
		goto L44
	}
L37:
	;
	F_pfree(m, v151)
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L1
	} else {
		goto L96
	}
L38:
	;
	v447 = int64(0)
	goto L37
L39:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L1
	} else {
		goto L92
	}
L40:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L1
	} else {
		goto L87
	}
L41:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L1
	} else {
		goto L81
	}
L42:
	;
	if v336 != v41 {
		goto L39
	} else {
		goto L79
	}
L43:
	;
	v336 = int32(0)
	v338 = v167
	goto L42
L44:
	;
	goto L45
L45:
	;
	v176 = int32(0)
	v183 = v176
	v185 = v167
	v186 = v169
	v189 = base.B2i32(v176 < v46)
	v190 = v176
	goto L46
L46:
	;
	v205 = v186 & int32(1)
	if v205 != 0 {
		goto L50
	} else {
		goto L51
	}
L47:
	;
	v336 = v323
	v338 = v324
	goto L42
L48:
	;
	v332 = base.B2i32(v324 < v46)
	v333 = base.B2i32(v323 < v41)
	if v332|v333 != 0 {
		v183 = v323
		v185 = v324
		v186 = v333
		v189 = v332
		v190 = v328
		goto L46
	} else {
		goto L78
	}
L49:
	;
	v323 = v183 + int32(1)
	v324 = v313
	v328 = v317
	goto L48
L50:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39+v206<<(uint(int32(3))%32)+v183*int32(100))+119)))
	if v213 == int32(1) {
		v313 = v185
		v317 = v190
		goto L49
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	if v189 == int32(0) {
		v336 = v183
		v338 = v185
		goto L42
	} else {
		goto L54
	}
L53:
	;
	goto L52
L54:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	v224 = v44 + v218<<(uint(int32(3))%32) + v185*int32(100)
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v224)+119)))
	if v225 == int32(1) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v323 = v183
	v324 = v185 + int32(1)
	v328 = v190
	goto L48
L56:
	;
	goto L57
L57:
	;
	if v205 == int32(0) {
		v336 = v183
		v338 = v185
		goto L42
	} else {
		goto L58
	}
L58:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	v238 = v39 + v232<<(uint(int32(3))%32) + v183*int32(100)
	v239 = int32(28)
	v240 = v238 + v239
	v242 = v224 + v239
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v238)+96))
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v224)+96))
	if v243 != v244 {
		goto L41
	} else {
		goto L59
	}
L59:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v242)+96))
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v240)+96))
	v250 = v94 + int32(20) + v190<<(uint(int32(2))%32)
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v250)))
	if v251 != 0 {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v185+v163))))
	v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183+v153))))
	if v265 == int32(1) {
		goto L68
	} else {
		goto L69
	}
L61:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v251)))
	if v252 == v243 {
		v261 = v251
		goto L60
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	v255 = F_lookup_type_cache(m, v243, int32(32))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L1
	} else {
		goto L65
	}
L64:
	;
	goto L63
L65:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v255)+80))
	if v257 == int32(0) {
		goto L40
	} else {
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v250))) = v255
	v261 = v255
	goto L60
L67:
	;
	v309 = int32(1)
	v313 = v185 + v309
	v317 = v190 + v309
	goto L49
L68:
	;
	if v263&int32(1) != 0 {
		goto L67
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	if v263&int32(1) != 0 {
		goto L38
	} else {
		goto L72
	}
L71:
	;
	goto L38
L72:
	;
	v272 = int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v25)+50)) = uint16(v272)
	v274 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+48)) = uint8(v274)
	if v247 == v246 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v278 = v247
	goto L75
L74:
	;
	v278 = v274
	goto L75
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+44)) = v278
	*(*int64)(unsafe.Add(mBase, uint32(v25)+36)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+32)) = v261 + int32(76)
	v285 = int32(3)
	v288 = *(*int64)(unsafe.Add(mBase, uint32(v151+v183<<(uint(v285)%32))))
	v289 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+64)) = uint8(v289)
	*(*int64)(unsafe.Add(mBase, uint32(v25)+56)) = v288
	v295 = *(*int64)(unsafe.Add(mBase, uint32(v161+v185<<(uint(v285)%32))))
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+80)) = uint8(v289)
	*(*int64)(unsafe.Add(mBase, uint32(v25)+72)) = v295
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v261)+76))
	v302 = m.T0[v301].(func(*base.Module, int32) int64)(m, v25+int32(32))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+48)))
	if v304|base.B2i32(v302 == int64(0)) != 0 {
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
	if v338 != v46 {
		goto L39
	} else {
		goto L80
	}
L80:
	;
	v447 = int64(1)
	goto L37
L81:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v240)+68))
	v368 = F_format_type_be(m, v367)
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v242)+68))
	v371 = F_format_type_be(m, v370)
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+24)) = v190 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+20)) = v371
	*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = v368
	F_errmsg(m, int32(_a_F_record_eq_0), v25+int32(16))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	F_errfinish(m, int32(_a_F_record_eq_1), int32(1198), int32(_a_F_record_eq_2))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
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
	v394 = m.ExcPending
	if v394 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v255)))
	v396 = F_format_type_be(m, v395)
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = v396
	F_errmsg(m, int32(_a_F_record_eq_3), v25)
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	F_errfinish(m, int32(_a_F_record_eq_1), int32(1221), int32(_a_F_record_eq_2))
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
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
	v414 = m.ExcPending
	if v414 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	F_errmsg(m, int32(_a_F_record_eq_4), int32(0))
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	F_errfinish(m, int32(_a_F_record_eq_1), int32(1265), int32(_a_F_record_eq_2))
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
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
	F_pfree(m, v153)
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	F_pfree(m, v161)
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	F_pfree(m, v163)
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
	if int32(0) <= v456 {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	F_DecrTupleDescRefCount(m, v39)
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L1
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v44)+12))
	if int32(0) <= v461 {
		goto L104
	} else {
		goto L105
	}
L103:
	;
	goto L102
L104:
	;
	F_DecrTupleDescRefCount(m, v44)
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L1
	} else {
		goto L107
	}
L105:
	;
	goto L106
L106:
	;
	v466 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v466 != v28 {
		goto L108
	} else {
		goto L109
	}
L107:
	;
	goto L106
L108:
	;
	F_pfree(m, v28)
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L1
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	v470 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v470 != v33 {
		goto L112
	} else {
		goto L113
	}
L111:
	;
	goto L110
L112:
	;
	F_pfree(m, v33)
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L1
	} else {
		goto L115
	}
L113:
	;
	goto L114
L114:
	;
	m.G0 = v25 + int32(128)
	return v447
L115:
	;
	goto L114
}
func F_record_ge(m *base.Module, l0 int32) int64 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_record_cmp(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(base.B2i32(int32(0) <= v2))
	}
}
func F_record_image_lt(m *base.Module, l0 int32) int64 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_record_image_cmp(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(int32(base.Ui32(v2) >> (uint(int32(31)) % 32)))
	}
}
func F_reduce_semijoin_in_jointree(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
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
	var v121 int32
	_ = v121
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	v3 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	if l0 == v3 {
		v121 = v3
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L11
	} else {
		goto L34
	}
L2:
	;
	m.G0 = v8 + int32(16)
	return v121
L3:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v12 - int32(63) {
	case 0:
		goto L4
	case 1:
		goto L5
	case 2:
		goto L6
	default:
		goto L1
	}
L4:
	;
	v121 = int32(0)
	goto L2
L5:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v41 != int32(4) {
		goto L15
	} else {
		goto L16
	}
L6:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v15 == int32(0) {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v18 = int32(0)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v19 <= v18 {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v23 = v18
	goto L9
L9:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v28+v23<<(uint(int32(2))%32))))
	v33 = F_reduce_semijoin_in_jointree(m, v32, l1)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L4
L11:
	;
	return int32(0)
L12:
	;
	if v33 != 0 {
		v121 = int32(1)
		goto L2
	} else {
		goto L13
	}
L13:
	;
	v38 = v23 + int32(1)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v38 < v39 {
		v23 = v38
		goto L9
	} else {
		goto L14
	}
L14:
	;
	goto L10
L15:
	;
	v106 = int32(1)
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v108 = F_reduce_semijoin_in_jointree(m, v107, l1)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L11
	} else {
		goto L30
	}
L16:
	;
	v44 = int32(1)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v48 = F_get_relids_in_jointree(m, v45, v44, int32(0))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L11
	} else {
		goto L17
	}
L17:
	;
	v50 = int32(0)
	if base.B2i32(v48 == v50)|base.B2i32(l1 == v50) != 0 {
		v96 = base.B2i32(v48|l1 == v50)
		goto L19
	} else {
		goto L20
	}
L18:
	;
	if v96 == int32(0) {
		goto L15
	} else {
		goto L29
	}
L19:
	;
	goto L18
L20:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v64 != v65 {
		v96 = int32(0)
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v67 = int32(1)
	if v64 <= v67 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v70 = v67
	goto L24
L23:
	;
	v70 = v64
	goto L24
L24:
	;
	v71 = int32(8)
	v76 = int32(0)
	goto L25
L25:
	;
	v84 = v76 << (uint(int32(2)) % 32)
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v48+v71+v84)))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l1+v71+v84)))
	v89 = base.B2i32(v86 == v88)
	if v86 != v88 {
		v96 = v89
		goto L19
	} else {
		goto L27
	}
L26:
	;
	v96 = v89
	goto L19
L27:
	;
	v92 = v76 + int32(1)
	if v92 != v70 {
		v76 = v92
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(0)
	v121 = v44
	goto L2
L30:
	;
	if v108 != 0 {
		v121 = v106
		goto L2
	} else {
		goto L31
	}
L31:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v111 = F_reduce_semijoin_in_jointree(m, v110, l1)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L11
	} else {
		goto L32
	}
L32:
	;
	if v111 != 0 {
		v121 = v106
		goto L2
	} else {
		goto L33
	}
L33:
	;
	goto L4
L34:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v132
	F_errmsg_internal(m, int32(_a_F_reduce_semijoin_in_jointree_0), v8)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L11
	} else {
		goto L35
	}
L35:
	;
	F_errfinish(m, int32(_a_F_reduce_semijoin_in_jointree_1), int32(578), int32(_a_F_reduce_semijoin_in_jointree_2))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L11
	} else {
		goto L36
	}
L36:
	;
	base.Wasm_trap_unreachable()
	for {
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
	F_errfinish(m, int32(_a_F_refnameNamespaceItem_1), int32(271), int32(_a_F_refnameNamespaceItem_2))
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
	F_errfinish(m, int32(_a_F_refnameNamespaceItem_1), int32(224), int32(_a_F_refnameNamespaceItem_4))
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
func F_reform_tuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
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
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
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
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l2)+52))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v13 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+18)))
	if v13&int32(2047) < v11 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_heap_deform_tuple(m, l0, v9, l3, l4)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L11
	} else {
		goto L16
	}
L2:
	;
	if int32(0) < v11 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v21 = v11
	v25 = int32(0)
	goto L6
L4:
	;
	goto L5
L5:
	;
	v52 = F_heap_copytuple(m, l0)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L11
	} else {
		goto L15
	}
L6:
	;
	v28 = v25 + int32(1)
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10+v25<<(uint(int32(3))%32))+34)))
	if v32&int32(4) != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	goto L5
L8:
	;
	v35 = F_heap_attisnull(m, l0, v28, v10)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v42 = v21
	goto L10
L10:
	;
	if v28 < v42 {
		v21 = v42
		v25 = v28
		goto L6
	} else {
		goto L14
	}
L11:
	;
	return int32(0)
L12:
	;
	if v35 == int32(0) {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v42 = v41
	goto L10
L14:
	;
	goto L7
L15:
	;
	return v52
L16:
	;
	v65 = int32(0)
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	if v65 < v66 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v71 = v65
	v75 = v66
	goto L20
L18:
	;
	goto L19
L19:
	;
	v99 = F_heap_form_tuple(m, v10, l3, l4)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L11
	} else {
		goto L26
	}
L20:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10+v71<<(uint(int32(3))%32))+34)))
	if v80&int32(4) != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	goto L19
L22:
	;
	v84 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v71+l4))) = uint8(v84)
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v87 = v86
	goto L24
L23:
	;
	v87 = v75
	goto L24
L24:
	;
	v89 = v71 + int32(1)
	if v89 < v87 {
		v71 = v89
		v75 = v87
		goto L20
	} else {
		goto L25
	}
L25:
	;
	goto L21
L26:
	;
	return v99
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
														F_errfinish(m, int32(_a_F_regcomp_auth_token_2), int32(328), int32(_a_F_regcomp_auth_token_3))
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
func F_regdictionaryout(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int64
	_ = v11
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
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
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
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
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = base.I32_wrap_i64(v11)
	if v12 == int32(0) {
		v16 = F_pstrdup(m, int32(_a_F_regdictionaryout_0))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int64(0)
		} else {
			v49 = v16
			m.G0 = v9 + int32(16)
			return base.I64_extend_i32_u(v49)
		}
	} else {
		v23 = F_SearchSysCache1(m, int32(76), v11&int64(4294967295))
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int64(0)
		} else {
			if v23 != 0 {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
				v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+22)))
				v27 = v25 + v26
				v30 = F_TSDictionaryIsVisible(m, v12)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int64(0)
				} else {
					if v30 != 0 {
						v36 = int32(0)
						v37 = F_quote_qualified_identifier(m, v36, v27+int32(4))
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int64(0)
						} else {
							F_ReleaseCatCache(m, v23)
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return int64(0)
							} else {
								v49 = v37
								m.G0 = v9 + int32(16)
								return base.I64_extend_i32_u(v49)
							}
						}
					} else {
						v33 = *(*int32)(unsafe.Add(mBase, uint32(v27)+68))
						v34 = F_get_namespace_name(m, v33)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int64(0)
						} else {
							v36 = v34
							v37 = F_quote_qualified_identifier(m, v36, v27+int32(4))
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return int64(0)
							} else {
								F_ReleaseCatCache(m, v23)
								mBase = m.M
								v40 = m.ExcPending
								if v40 != 0 {
									return int64(0)
								} else {
									v49 = v37
									m.G0 = v9 + int32(16)
									return base.I64_extend_i32_u(v49)
								}
							}
						}
					}
				}
			} else {
				v42 = F_palloc(m, int32(64))
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return int64(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = v12
					v47 = F_pg_snprintf(m, v42, int32(64), int32(_a_F_regdictionaryout_1), v9)
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return int64(0)
					} else {
						v49 = v42
						m.G0 = v9 + int32(16)
						return base.I64_extend_i32_u(v49)
					}
				}
			}
		}
	}
}
func F_regexnejoinsel(m *base.Module, l0 int32) int64 {
	return int64(4607137382803743703)
}
func F_regoperatorout(m *base.Module, l0 int32) int64 {
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
	var v14 int32
	_ = v14
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v2 == int32(0) {
		v6 = F_pstrdup(m, int32(_a_F_regoperatorout_0))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int64(0)
		} else {
			return base.I64_extend_i32_u(v6)
		}
	} else {
		v13 = F_format_operator_extended(m, v2, int32(0))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int64(0)
		} else {
			return base.I64_extend_i32_u(v13)
		}
	}
}
func F_regroleout(m *base.Module, l0 int32) int64 {
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
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v8 == int32(0) {
		v12 = F_pstrdup(m, int32(_a_F_regroleout_0))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int64(0)
		} else {
			v31 = v12
			m.G0 = v6 + int32(16)
			return base.I64_extend_i32_u(v31)
		}
	} else {
		v17 = F_GetUserNameFromId(m, v8, int32(1))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int64(0)
		} else {
			if v17 != 0 {
				v19 = F_quote_identifier(m, v17)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int64(0)
				} else {
					v21 = F_pstrdup(m, v19)
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return int64(0)
					} else {
						v31 = v21
						m.G0 = v6 + int32(16)
						return base.I64_extend_i32_u(v31)
					}
				}
			} else {
				v24 = F_palloc(m, int32(64))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int64(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v6))) = v8
					v29 = F_pg_snprintf(m, v24, int32(64), int32(_a_F_regroleout_1), v6)
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int64(0)
					} else {
						v31 = v24
						m.G0 = v6 + int32(16)
						return base.I64_extend_i32_u(v31)
					}
				}
			}
		}
	}
}
func F_remove_useless_results_recurse(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v152 int32
	_ = v152
	var v162 int32
	_ = v162
	var v168 int32
	_ = v168
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v224 int32
	_ = v224
	var v239 int32
	_ = v239
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
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v313 int32
	_ = v313
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
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v358 int32
	_ = v358
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
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v454 int32
	_ = v454
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v503 int32
	_ = v503
	var v508 int32
	_ = v508
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v517 int32
	_ = v517
	var v522 int32
	_ = v522
	v6 = int32(0)
	v13 = m.G0
	v15 = v13 + int32(-64)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v17 - int32(63) {
	case 0:
		v478 = l1
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
	v512 = m.ExcPending
	if v512 != 0 {
		goto L12
	} else {
		goto L134
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L12
	} else {
		goto L131
	}
L3:
	;
	m.G0 = v15 - int32(-64)
	return v478
L4:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v263 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v263 == int32(1) {
		goto L65
	} else {
		goto L66
	}
L5:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v20 == int32(0) {
		v478 = l1
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v31 = v20
	v32 = v6
	v34 = v6
	goto L7
L7:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	if v32 < v37 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	if v92 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L9:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	v42 = v39 + v32<<(uint(int32(2))%32)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	v44 = F_remove_useless_results_recurse(m, l0, v43, l2, l1+int32(8), l4)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v92 = v34
	goto L11
L11:
	;
	goto L8
L12:
	;
	return int32(0)
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42))) = v44
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v49 == int32(0) {
		v83 = v31
		v84 = v32
		v86 = v34
		goto L14
	} else {
		goto L15
	}
L14:
	;
	if v83 != 0 {
		v31 = v83
		v32 = v84 + int32(1)
		v34 = v86
		goto L7
	} else {
		goto L24
	}
L15:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	if v52 <= int32(1) {
		v83 = v31
		v84 = v32
		v86 = v34
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	if v55 != int32(63) {
		v83 = v31
		v84 = v32
		v86 = v34
		goto L14
	} else {
		goto L17
	}
L17:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	if v58 == int32(0) {
		v83 = v31
		v84 = v32
		v86 = v34
		goto L14
	} else {
		goto L18
	}
L18:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+52))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+12))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v63+v58<<(uint(int32(2))%32)-int32(4))))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+12))
	if v70 != int32(8) {
		v83 = v31
		v84 = v32
		v86 = v34
		goto L14
	} else {
		goto L19
	}
L19:
	;
	v73 = F_find_dependent_phvs_in_jointree(m, l0, l1, v58, l2)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L12
	} else {
		goto L20
	}
L20:
	;
	if v73 != 0 {
		v83 = v31
		v84 = v32
		v86 = v34
		goto L14
	} else {
		goto L21
	}
L21:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v76 = F_list_delete_nth_cell(m, v75, v32)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L12
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v76
	v81 = F_bms_add_member(m, v34, v58)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L12
	} else {
		goto L23
	}
L23:
	;
	v83 = v76
	v84 = v32 - int32(1)
	v86 = v81
	goto L14
L24:
	;
	v92 = v86
	goto L11
L25:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v239 == int32(0) {
		v478 = l1
		goto L3
	} else {
		goto L54
	}
L26:
	;
	if v92 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	if v152 < int32(0) {
		goto L25
	} else {
		goto L38
	}
L28:
	;
	v152 = base.I32_ctz(v138) | v139<<(uint(int32(5))%32)
	goto L27
L29:
	;
	v152 = int32(-2)
	goto L27
L30:
	;
	v103 = int32(0)
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v92)+4))
	if v106 <= v103 {
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v109 = v92 + int32(8)
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	v116 = v113 & int32(-1)
	if v116 != 0 {
		v138 = v116
		v139 = v103
		goto L28
	} else {
		goto L32
	}
L32:
	;
	v117 = int32(1)
	if v117 == v106 {
		goto L29
	} else {
		goto L33
	}
L33:
	;
	v121 = v117
	goto L34
L34:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v109+v121<<(uint(int32(2))%32))))
	if v128 != 0 {
		v138 = v128
		v139 = v121
		goto L28
	} else {
		goto L36
	}
L35:
	;
	goto L29
L36:
	;
	v130 = v121 + int32(1)
	if v130 != v106 {
		v121 = v130
		goto L34
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	v162 = v152
	goto L39
L39:
	;
	F_remove_result_refs(m, l0, v162, l1)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L12
	} else {
		goto L41
	}
L40:
	;
	goto L25
L41:
	;
	if v92 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	if int32(0) <= v224 {
		v162 = v224
		goto L39
	} else {
		goto L53
	}
L43:
	;
	v224 = base.I32_ctz(v210) | v211<<(uint(int32(5))%32)
	goto L42
L44:
	;
	v224 = int32(-2)
	goto L42
L45:
	;
	v175 = v162 + int32(1)
	v177 = int32(base.Ui32(v175) >> (uint(int32(5)) % 32))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v92)+4))
	if v178 <= v177 {
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v181 = v92 + int32(8)
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v181+v177<<(uint(int32(2))%32))))
	v188 = v185 & (int32(-1) << (uint(v175) % 32))
	if v188 != 0 {
		v210 = v188
		v211 = v177
		goto L43
	} else {
		goto L47
	}
L47:
	;
	v190 = v177 + int32(1)
	if v190 == v178 {
		goto L44
	} else {
		goto L48
	}
L48:
	;
	v193 = v190
	goto L49
L49:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v181+v193<<(uint(int32(2))%32))))
	if v200 != 0 {
		v210 = v200
		v211 = v193
		goto L43
	} else {
		goto L51
	}
L50:
	;
	goto L44
L51:
	;
	v202 = v193 + int32(1)
	if v202 != v178 {
		v193 = v202
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
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v239)+4))
	if v242 != int32(1) {
		v478 = l1
		goto L3
	} else {
		goto L55
	}
L55:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v245)+60))
	if l1 == v246 {
		v478 = l1
		goto L3
	} else {
		goto L56
	}
L56:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v249 != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v250 = l3
	goto L59
L58:
	;
	v250 = int32(1)
	goto L59
L59:
	;
	if v250 == int32(0) {
		v478 = l1
		goto L3
	} else {
		goto L60
	}
L60:
	;
	if v249 != 0 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v254 = F_list_concat(m, v249, v253)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L12
	} else {
		goto L64
	}
L62:
	;
	v258 = v239
	goto L63
L63:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v258)+12))
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v259)))
	v478 = v260
	goto L3
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v254
	v257 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v258 = v257
	goto L63
L65:
	;
	v266 = l3
	goto L67
L66:
	;
	v266 = int32(0)
	goto L67
L67:
	;
	v268 = l1 + int32(28)
	if v263 != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v269 = v266
	goto L70
L69:
	;
	v269 = v268
	goto L70
L70:
	;
	v270 = F_remove_useless_results_recurse(m, l0, v261, l2, v269, l4)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L12
	} else {
		goto L71
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v270
	v273 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v275 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if base.Ui32(v275) < base.Ui32(int32(2)) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v278 = v268
	goto L74
L73:
	;
	v278 = int32(0)
	goto L74
L74:
	;
	v279 = F_remove_useless_results_recurse(m, l0, v273, l2, v278, l4)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L12
	} else {
		goto L75
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v279
	v282 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	switch v282 {
	case 0:
		goto L79
	case 1:
		goto L78
	case 2, 5:
		v478 = l1
		goto L3
	default:
		goto L2
	case 4:
		goto L77
	}
L76:
	;
	v476 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v478 = v476
	goto L3
L77:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v279)))
	if v427 != int32(63) {
		v478 = l1
		goto L3
	} else {
		goto L120
	}
L78:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v279)))
	if v375 != int32(63) {
		v478 = l1
		goto L3
	} else {
		goto L107
	}
L79:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v283)))
	if v284 != int32(63) {
		v330 = v279
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v330)))
	if v331 != int32(63) {
		v478 = l1
		goto L3
	} else {
		goto L96
	}
L81:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v283)+4))
	if v287 == int32(0) {
		v330 = v279
		goto L80
	} else {
		goto L82
	}
L82:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v290)+52))
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v291)+12))
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v292+v287<<(uint(int32(2))%32)-int32(4))))
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v298)+12))
	if v299 != int32(8) {
		v330 = v279
		goto L80
	} else {
		goto L83
	}
L83:
	;
	v302 = F_find_dependent_phvs_in_jointree(m, l0, v279, v287, l2)
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L12
	} else {
		goto L84
	}
L84:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v302 != 0 {
		v330 = v304
		goto L80
	} else {
		goto L85
	}
L85:
	;
	F_remove_result_refs(m, l0, v287, v304)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L12
	} else {
		goto L86
	}
L86:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v268)))
	v308 = int32(0)
	if l3|base.B2i32(v307 == v308) == v308 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v313
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v313
	v319 = F_list_make1_impl(m, int32(1), v13+int32(-32))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L12
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	if v307 != 0 {
		goto L92
	} else {
		goto L93
	}
L90:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v322 = F_makeFromExpr(m, v319, v321)
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L12
	} else {
		goto L91
	}
L91:
	;
	v478 = v322
	goto L3
L92:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v325 = F_list_concat(m, v307, v324)
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L12
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v478 = v328
	goto L3
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v325
	goto L94
L96:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v330)+4))
	if v334 == int32(0) {
		v478 = l1
		goto L3
	} else {
		goto L97
	}
L97:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v337)+52))
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v338)+12))
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v339+v334<<(uint(int32(2))%32)-int32(4))))
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v345)+12))
	if v346 != int32(8) {
		v478 = l1
		goto L3
	} else {
		goto L98
	}
L98:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	F_remove_result_refs(m, l0, v334, v349)
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L12
	} else {
		goto L99
	}
L99:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v353 = int32(0)
	if l3|base.B2i32(v352 == v353) == v353 {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = v358
	*(*int32)(unsafe.Add(mBase, uint32(v15)+44)) = v358
	v364 = F_list_make1_impl(m, int32(1), v13+int32(-36))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L12
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	if v352 == int32(0) {
		goto L76
	} else {
		goto L105
	}
L103:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v367 = F_makeFromExpr(m, v364, v366)
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L12
	} else {
		goto L104
	}
L104:
	;
	v478 = v367
	goto L3
L105:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v372 = F_list_concat(m, v352, v371)
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L12
	} else {
		goto L106
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v372
	goto L76
L107:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v279)+4))
	if v378 == int32(0) {
		v478 = l1
		goto L3
	} else {
		goto L108
	}
L108:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v381)+52))
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v382)+12))
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v383+v378<<(uint(int32(2))%32)-int32(4))))
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v389)+12))
	if v390 != int32(8) {
		v478 = l1
		goto L3
	} else {
		goto L109
	}
L109:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v268)))
	if v393 == int32(0) {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v419 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	F_remove_result_refs(m, l0, v378, v419)
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L12
	} else {
		goto L118
	}
L111:
	;
	v396 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v396)+80))
	if v397 == int32(0) {
		goto L110
	} else {
		goto L112
	}
L112:
	;
	v400 = F_bms_make_singleton(m, v378)
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L12
	} else {
		goto L113
	}
L113:
	;
	v402 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+60)) = v402
	*(*int32)(unsafe.Add(mBase, uint32(v15)+56)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v400
	v406 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v409 = v13 + int32(-12)
	v411 = F_query_tree_walker_impl(m, v406, int32(901), v409, v402)
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L12
	} else {
		goto L114
	}
L114:
	;
	if v411 != 0 {
		v478 = l1
		goto L3
	} else {
		goto L115
	}
L115:
	;
	v413 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v415 = F_expression_tree_walker_impl(m, v413, int32(901), v409)
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L12
	} else {
		goto L116
	}
L116:
	;
	if v415 != 0 {
		v478 = l1
		goto L3
	} else {
		goto L117
	}
L117:
	;
	goto L110
L118:
	;
	v422 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v423 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v424 = F_bms_add_member(m, v422, v423)
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L12
	} else {
		goto L119
	}
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v424
	goto L76
L120:
	;
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v279)+4))
	if v430 == int32(0) {
		v478 = l1
		goto L3
	} else {
		goto L121
	}
L121:
	;
	v433 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v433)+52))
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v434)+12))
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v435+v430<<(uint(int32(2))%32)-int32(4))))
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v441)+12))
	if v442 != int32(8) {
		v478 = l1
		goto L3
	} else {
		goto L122
	}
L122:
	;
	v445 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	F_remove_result_refs(m, l0, v430, v445)
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L12
	} else {
		goto L123
	}
L123:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v449 = int32(0)
	if l3|base.B2i32(v448 == v449) == v449 {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v454 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = v454
	*(*int32)(unsafe.Add(mBase, uint32(v15)+40)) = v454
	v460 = F_list_make1_impl(m, int32(1), v13+int32(-28))
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L12
	} else {
		goto L127
	}
L125:
	;
	goto L126
L126:
	;
	if v448 == int32(0) {
		goto L76
	} else {
		goto L129
	}
L127:
	;
	v462 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v463 = F_makeFromExpr(m, v460, v462)
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L12
	} else {
		goto L128
	}
L128:
	;
	v478 = v463
	goto L3
L129:
	;
	v467 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v468 = F_list_concat(m, v448, v467)
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L12
	} else {
		goto L130
	}
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v468
	goto L76
L131:
	;
	v497 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v497
	F_errmsg_internal(m, int32(_a_F_remove_useless_results_recurse_0), v13+int32(-48))
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L12
	} else {
		goto L132
	}
L132:
	;
	F_errfinish(m, int32(_a_F_remove_useless_results_recurse_1), int32(_a_F_remove_useless_results_recurse_2), int32(_a_F_remove_useless_results_recurse_3))
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
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
	v513 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v513
	F_errmsg_internal(m, int32(_a_F_remove_useless_results_recurse_4), v15)
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L12
	} else {
		goto L135
	}
L135:
	;
	F_errfinish(m, int32(_a_F_remove_useless_results_recurse_1), int32(_a_F_remove_useless_results_recurse_5), int32(_a_F_remove_useless_results_recurse_3))
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
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
	var v49 int64
	_ = v49
	var v51 int32
	_ = v51
	var v59 int32
	_ = v59
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
	var v82 int32
	_ = v82
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
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	v9 = m.G0
	v11 = v9 - int32(144)
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
	v154 = m.ExcPending
	if v154 != 0 {
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
	v49 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l1)+56)))
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
	m.G0 = v11 + int32(144)
	return
L12:
	;
	return
L13:
	;
	F_ScanKeyInit(m, v11+int32(88), int32(4), int32(3), int32(62), base.I64_extend_i32_u(l3))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v64 = F_systable_beginscan(m, l0, int32(2701), int32(1), int32(0), int32(2), v45)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	v66 = F_systable_getnext(m, v64)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L12
	} else {
		goto L16
	}
L16:
	;
	if v66 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	F_systable_endscan(m, v64)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L12
	} else {
		goto L18
	}
L18:
	;
	v70 = F_heap_copytuple(m, l2)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L12
	} else {
		goto L20
	}
L19:
	;
	v126 = F_strncpy(m, v76, l3, int32(64))
	mBase = m.M
	v127 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v126)+63)) = uint8(v127)
	goto L33
L20:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v70)+16))
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+22)))
	v74 = v72 + v73
	v76 = v74 + int32(12)
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76))))
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
	if base.B2i32(v79 == int32(0))|base.B2i32(v79 != v82) != 0 {
		v100 = v79
		v101 = v82
		goto L22
	} else {
		goto L23
	}
L21:
	;
	if v100-v101 == int32(0) {
		goto L19
	} else {
		goto L28
	}
L22:
	;
	goto L21
L23:
	;
	v85 = v76
	v86 = l4
	goto L24
L24:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+1)))
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+1)))
	if v90 == int32(0) {
		v100 = v90
		v101 = v89
		goto L22
	} else {
		goto L26
	}
L25:
	;
	v100 = v90
	v101 = v89
	goto L22
L26:
	;
	v93 = int32(1)
	if v90 == v89 {
		v85 = v85 + v93
		v86 = v86 + v93
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	v107 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L12
	} else {
		goto L29
	}
L29:
	;
	if v107 == int32(0) {
		goto L19
	} else {
		goto L30
	}
L30:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v76
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v111 + int32(4)
	F_errmsg(m, int32(_a_F_renametrig_internal_0), v11)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L12
	} else {
		goto L31
	}
L31:
	;
	F_errfinish(m, int32(_a_F_renametrig_internal_1), int32(1648), int32(_a_F_renametrig_internal_2))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L12
	} else {
		goto L32
	}
L32:
	;
	goto L19
L33:
	;
	F_CatalogTupleUpdate(m, l0, v70+int32(4), v70)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L12
	} else {
		goto L34
	}
L34:
	;
	v134 = *(*int32)(unsafe.Add(mBase, _c_F_renametrig_internal[0]))
	if v134 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	v137 = int32(0)
	F_RunObjectPostAlterHook(m, int32(2620), v136, v137, v137, v137)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
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
	v143 = m.ExcPending
	if v143 != 0 {
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
	v157 = m.ExcPending
	if v157 != 0 {
		goto L12
	} else {
		goto L41
	}
L41:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v158 + int32(4)
	F_errmsg(m, int32(_a_F_renametrig_internal_4), v11+int32(16))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L12
	} else {
		goto L42
	}
L42:
	;
	F_errfinish(m, int32(_a_F_renametrig_internal_1), int32(1630), int32(_a_F_renametrig_internal_2))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
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
func F_repeat_2(m *base.Module, l0 int32) int64 {
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
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
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
	return int64(0)
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
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
	return base.I64_extend_i32_u(v64)
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
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	F_errmsg(m, int32(_a_F_repeat_2_0), int32(0))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(_a_F_repeat_2_1), int32(1182), int32(_a_F_repeat_2_2))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
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
	var v34 int32
	_ = v34
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
					v34 = F_errdetail(m, int32(_a_F_report_newlocale_failure_1), v6)
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_report_newlocale_failure_2), int32(1195), int32(_a_F_report_newlocale_failure_3))
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
					F_errfinish(m, int32(_a_F_report_newlocale_failure_2), int32(1195), int32(_a_F_report_newlocale_failure_3))
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
	F_errfinish(m, int32(_a_F_report_untranslatable_char_3), int32(1902), int32(_a_F_report_untranslatable_char_4))
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
func F_reset_conflict_slot_xmin_to_safe_horizon(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_reset_conflict_slot_xmin_to_safe_horizon[0]))
	v14 = F_LWLockAcquire(m, v10+int32(_a_F_reset_conflict_slot_xmin_to_safe_horizon_0), int32(0))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, _c_F_reset_conflict_slot_xmin_to_safe_horizon[0]))
		v21 = F_LWLockAcquire(m, v17+int32(512), int32(0))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return
		} else {
			v24 = F_GetOldestSafeDecodingTransactionId(m, int32(0))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return
			} else {
				v27 = *(*int32)(unsafe.Add(mBase, _c_F_reset_conflict_slot_xmin_to_safe_horizon[1]))
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+96))
				if v28 == int32(0) {
					v56 = base.AtomicRmwXchg32(m, v27, int32(0), int32(1))
					if v56 != 0 {
						F_s_lock(m, v27, int32(_a_F_reset_conflict_slot_xmin_to_safe_horizon_1))
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return
						} else {
							v61 = *(*int32)(unsafe.Add(mBase, _c_F_reset_conflict_slot_xmin_to_safe_horizon[1]))
							*(*int32)(unsafe.Add(mBase, uint32(v61)+96)) = v24
							*(*int32)(unsafe.Add(mBase, uint32(v61)+16)) = v24
							v64 = int32(0)
							atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v61))), uint32(v64))
							v69 = F_errstart(m, int32(14), v64)
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return
							} else {
								if v69 != 0 {
									*(*int32)(unsafe.Add(mBase, uint32(v7))) = v24
									F_errmsg_internal(m, int32(_a_F_reset_conflict_slot_xmin_to_safe_horizon_2), v7)
									mBase = m.M
									v74 = m.ExcPending
									if v74 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_reset_conflict_slot_xmin_to_safe_horizon_3), int32(1614), int32(_a_F_reset_conflict_slot_xmin_to_safe_horizon_4))
										mBase = m.M
										v79 = m.ExcPending
										if v79 != 0 {
											return
										} else {
											F_ReplicationSlotsComputeRequiredXmin(m, int32(1))
											mBase = m.M
											v82 = m.ExcPending
											if v82 != 0 {
												return
											} else {
												v84 = *(*int32)(unsafe.Add(mBase, _c_F_reset_conflict_slot_xmin_to_safe_horizon[0]))
												F_LWLockRelease(m, v84+int32(512))
												mBase = m.M
												v88 = m.ExcPending
												if v88 != 0 {
													return
												} else {
													v90 = *(*int32)(unsafe.Add(mBase, _c_F_reset_conflict_slot_xmin_to_safe_horizon[0]))
													F_LWLockRelease(m, v90+int32(_a_F_reset_conflict_slot_xmin_to_safe_horizon_0))
													mBase = m.M
													v94 = m.ExcPending
													if v94 != 0 {
														return
													} else {
														F_ReplicationSlotMarkDirty(m)
														mBase = m.M
														v96 = m.ExcPending
														if v96 != 0 {
															return
														} else {
															F_ReplicationSlotSave(m)
															mBase = m.M
															v98 = m.ExcPending
															if v98 != 0 {
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
									F_ReplicationSlotsComputeRequiredXmin(m, int32(1))
									mBase = m.M
									v82 = m.ExcPending
									if v82 != 0 {
										return
									} else {
										v84 = *(*int32)(unsafe.Add(mBase, _c_F_reset_conflict_slot_xmin_to_safe_horizon[0]))
										F_LWLockRelease(m, v84+int32(512))
										mBase = m.M
										v88 = m.ExcPending
										if v88 != 0 {
											return
										} else {
											v90 = *(*int32)(unsafe.Add(mBase, _c_F_reset_conflict_slot_xmin_to_safe_horizon[0]))
											F_LWLockRelease(m, v90+int32(_a_F_reset_conflict_slot_xmin_to_safe_horizon_0))
											mBase = m.M
											v94 = m.ExcPending
											if v94 != 0 {
												return
											} else {
												F_ReplicationSlotMarkDirty(m)
												mBase = m.M
												v96 = m.ExcPending
												if v96 != 0 {
													return
												} else {
													F_ReplicationSlotSave(m)
													mBase = m.M
													v98 = m.ExcPending
													if v98 != 0 {
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
						}
					} else {
						v61 = *(*int32)(unsafe.Add(mBase, _c_F_reset_conflict_slot_xmin_to_safe_horizon[1]))
						*(*int32)(unsafe.Add(mBase, uint32(v61)+96)) = v24
						*(*int32)(unsafe.Add(mBase, uint32(v61)+16)) = v24
						v64 = int32(0)
						atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v61))), uint32(v64))
						v69 = F_errstart(m, int32(14), v64)
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
							return
						} else {
							if v69 != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(v7))) = v24
								F_errmsg_internal(m, int32(_a_F_reset_conflict_slot_xmin_to_safe_horizon_2), v7)
								mBase = m.M
								v74 = m.ExcPending
								if v74 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_reset_conflict_slot_xmin_to_safe_horizon_3), int32(1614), int32(_a_F_reset_conflict_slot_xmin_to_safe_horizon_4))
									mBase = m.M
									v79 = m.ExcPending
									if v79 != 0 {
										return
									} else {
										F_ReplicationSlotsComputeRequiredXmin(m, int32(1))
										mBase = m.M
										v82 = m.ExcPending
										if v82 != 0 {
											return
										} else {
											v84 = *(*int32)(unsafe.Add(mBase, _c_F_reset_conflict_slot_xmin_to_safe_horizon[0]))
											F_LWLockRelease(m, v84+int32(512))
											mBase = m.M
											v88 = m.ExcPending
											if v88 != 0 {
												return
											} else {
												v90 = *(*int32)(unsafe.Add(mBase, _c_F_reset_conflict_slot_xmin_to_safe_horizon[0]))
												F_LWLockRelease(m, v90+int32(_a_F_reset_conflict_slot_xmin_to_safe_horizon_0))
												mBase = m.M
												v94 = m.ExcPending
												if v94 != 0 {
													return
												} else {
													F_ReplicationSlotMarkDirty(m)
													mBase = m.M
													v96 = m.ExcPending
													if v96 != 0 {
														return
													} else {
														F_ReplicationSlotSave(m)
														mBase = m.M
														v98 = m.ExcPending
														if v98 != 0 {
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
								F_ReplicationSlotsComputeRequiredXmin(m, int32(1))
								mBase = m.M
								v82 = m.ExcPending
								if v82 != 0 {
									return
								} else {
									v84 = *(*int32)(unsafe.Add(mBase, _c_F_reset_conflict_slot_xmin_to_safe_horizon[0]))
									F_LWLockRelease(m, v84+int32(512))
									mBase = m.M
									v88 = m.ExcPending
									if v88 != 0 {
										return
									} else {
										v90 = *(*int32)(unsafe.Add(mBase, _c_F_reset_conflict_slot_xmin_to_safe_horizon[0]))
										F_LWLockRelease(m, v90+int32(_a_F_reset_conflict_slot_xmin_to_safe_horizon_0))
										mBase = m.M
										v94 = m.ExcPending
										if v94 != 0 {
											return
										} else {
											F_ReplicationSlotMarkDirty(m)
											mBase = m.M
											v96 = m.ExcPending
											if v96 != 0 {
												return
											} else {
												F_ReplicationSlotSave(m)
												mBase = m.M
												v98 = m.ExcPending
												if v98 != 0 {
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
					}
				} else {
					v31 = int32(3)
					if base.B2i32(base.Ui32(v24) < base.Ui32(v31))|base.B2i32(base.Ui32(v28) < base.Ui32(v31)) == int32(0) {
						if v28-v24 <= int32(0) {
							v43 = *(*int32)(unsafe.Add(mBase, _c_F_reset_conflict_slot_xmin_to_safe_horizon[0]))
							F_LWLockRelease(m, v43+int32(512))
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return
							} else {
								v49 = *(*int32)(unsafe.Add(mBase, _c_F_reset_conflict_slot_xmin_to_safe_horizon[0]))
								F_LWLockRelease(m, v49+int32(_a_F_reset_conflict_slot_xmin_to_safe_horizon_0))
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return
								} else {
									m.G0 = v7 + int32(16)
									return
								}
							}
						} else {
							v56 = base.AtomicRmwXchg32(m, v27, int32(0), int32(1))
							if v56 != 0 {
								F_s_lock(m, v27, int32(_a_F_reset_conflict_slot_xmin_to_safe_horizon_1))
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
									return
								} else {
									v61 = *(*int32)(unsafe.Add(mBase, _c_F_reset_conflict_slot_xmin_to_safe_horizon[1]))
									*(*int32)(unsafe.Add(mBase, uint32(v61)+96)) = v24
									*(*int32)(unsafe.Add(mBase, uint32(v61)+16)) = v24
									v64 = int32(0)
									atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v61))), uint32(v64))
									v69 = F_errstart(m, int32(14), v64)
									mBase = m.M
									v70 = m.ExcPending
									if v70 != 0 {
										return
									} else {
										if v69 != 0 {
											*(*int32)(unsafe.Add(mBase, uint32(v7))) = v24
											F_errmsg_internal(m, int32(_a_F_reset_conflict_slot_xmin_to_safe_horizon_2), v7)
											mBase = m.M
											v74 = m.ExcPending
											if v74 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_reset_conflict_slot_xmin_to_safe_horizon_3), int32(1614), int32(_a_F_reset_conflict_slot_xmin_to_safe_horizon_4))
												mBase = m.M
												v79 = m.ExcPending
												if v79 != 0 {
													return
												} else {
													F_ReplicationSlotsComputeRequiredXmin(m, int32(1))
													mBase = m.M
													v82 = m.ExcPending
													if v82 != 0 {
														return
													} else {
														v84 = *(*int32)(unsafe.Add(mBase, _c_F_reset_conflict_slot_xmin_to_safe_horizon[0]))
														F_LWLockRelease(m, v84+int32(512))
														mBase = m.M
														v88 = m.ExcPending
														if v88 != 0 {
															return
														} else {
															v90 = *(*int32)(unsafe.Add(mBase, _c_F_reset_conflict_slot_xmin_to_safe_horizon[0]))
															F_LWLockRelease(m, v90+int32(_a_F_reset_conflict_slot_xmin_to_safe_horizon_0))
															mBase = m.M
															v94 = m.ExcPending
															if v94 != 0 {
																return
															} else {
																F_ReplicationSlotMarkDirty(m)
																mBase = m.M
																v96 = m.ExcPending
																if v96 != 0 {
																	return
																} else {
																	F_ReplicationSlotSave(m)
																	mBase = m.M
																	v98 = m.ExcPending
																	if v98 != 0 {
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
											F_ReplicationSlotsComputeRequiredXmin(m, int32(1))
											mBase = m.M
											v82 = m.ExcPending
											if v82 != 0 {
												return
											} else {
												v84 = *(*int32)(unsafe.Add(mBase, _c_F_reset_conflict_slot_xmin_to_safe_horizon[0]))
												F_LWLockRelease(m, v84+int32(512))
												mBase = m.M
												v88 = m.ExcPending
												if v88 != 0 {
													return
												} else {
													v90 = *(*int32)(unsafe.Add(mBase, _c_F_reset_conflict_slot_xmin_to_safe_horizon[0]))
													F_LWLockRelease(m, v90+int32(_a_F_reset_conflict_slot_xmin_to_safe_horizon_0))
													mBase = m.M
													v94 = m.ExcPending
													if v94 != 0 {
														return
													} else {
														F_ReplicationSlotMarkDirty(m)
														mBase = m.M
														v96 = m.ExcPending
														if v96 != 0 {
															return
														} else {
															F_ReplicationSlotSave(m)
															mBase = m.M
															v98 = m.ExcPending
															if v98 != 0 {
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
								}
							} else {
								v61 = *(*int32)(unsafe.Add(mBase, _c_F_reset_conflict_slot_xmin_to_safe_horizon[1]))
								*(*int32)(unsafe.Add(mBase, uint32(v61)+96)) = v24
								*(*int32)(unsafe.Add(mBase, uint32(v61)+16)) = v24
								v64 = int32(0)
								atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v61))), uint32(v64))
								v69 = F_errstart(m, int32(14), v64)
								mBase = m.M
								v70 = m.ExcPending
								if v70 != 0 {
									return
								} else {
									if v69 != 0 {
										*(*int32)(unsafe.Add(mBase, uint32(v7))) = v24
										F_errmsg_internal(m, int32(_a_F_reset_conflict_slot_xmin_to_safe_horizon_2), v7)
										mBase = m.M
										v74 = m.ExcPending
										if v74 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_reset_conflict_slot_xmin_to_safe_horizon_3), int32(1614), int32(_a_F_reset_conflict_slot_xmin_to_safe_horizon_4))
											mBase = m.M
											v79 = m.ExcPending
											if v79 != 0 {
												return
											} else {
												F_ReplicationSlotsComputeRequiredXmin(m, int32(1))
												mBase = m.M
												v82 = m.ExcPending
												if v82 != 0 {
													return
												} else {
													v84 = *(*int32)(unsafe.Add(mBase, _c_F_reset_conflict_slot_xmin_to_safe_horizon[0]))
													F_LWLockRelease(m, v84+int32(512))
													mBase = m.M
													v88 = m.ExcPending
													if v88 != 0 {
														return
													} else {
														v90 = *(*int32)(unsafe.Add(mBase, _c_F_reset_conflict_slot_xmin_to_safe_horizon[0]))
														F_LWLockRelease(m, v90+int32(_a_F_reset_conflict_slot_xmin_to_safe_horizon_0))
														mBase = m.M
														v94 = m.ExcPending
														if v94 != 0 {
															return
														} else {
															F_ReplicationSlotMarkDirty(m)
															mBase = m.M
															v96 = m.ExcPending
															if v96 != 0 {
																return
															} else {
																F_ReplicationSlotSave(m)
																mBase = m.M
																v98 = m.ExcPending
																if v98 != 0 {
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
										F_ReplicationSlotsComputeRequiredXmin(m, int32(1))
										mBase = m.M
										v82 = m.ExcPending
										if v82 != 0 {
											return
										} else {
											v84 = *(*int32)(unsafe.Add(mBase, _c_F_reset_conflict_slot_xmin_to_safe_horizon[0]))
											F_LWLockRelease(m, v84+int32(512))
											mBase = m.M
											v88 = m.ExcPending
											if v88 != 0 {
												return
											} else {
												v90 = *(*int32)(unsafe.Add(mBase, _c_F_reset_conflict_slot_xmin_to_safe_horizon[0]))
												F_LWLockRelease(m, v90+int32(_a_F_reset_conflict_slot_xmin_to_safe_horizon_0))
												mBase = m.M
												v94 = m.ExcPending
												if v94 != 0 {
													return
												} else {
													F_ReplicationSlotMarkDirty(m)
													mBase = m.M
													v96 = m.ExcPending
													if v96 != 0 {
														return
													} else {
														F_ReplicationSlotSave(m)
														mBase = m.M
														v98 = m.ExcPending
														if v98 != 0 {
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
							}
						}
					} else {
						if base.Ui32(v24) < base.Ui32(v28) {
							v56 = base.AtomicRmwXchg32(m, v27, int32(0), int32(1))
							if v56 != 0 {
								F_s_lock(m, v27, int32(_a_F_reset_conflict_slot_xmin_to_safe_horizon_1))
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
									return
								} else {
									v61 = *(*int32)(unsafe.Add(mBase, _c_F_reset_conflict_slot_xmin_to_safe_horizon[1]))
									*(*int32)(unsafe.Add(mBase, uint32(v61)+96)) = v24
									*(*int32)(unsafe.Add(mBase, uint32(v61)+16)) = v24
									v64 = int32(0)
									atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v61))), uint32(v64))
									v69 = F_errstart(m, int32(14), v64)
									mBase = m.M
									v70 = m.ExcPending
									if v70 != 0 {
										return
									} else {
										if v69 != 0 {
											*(*int32)(unsafe.Add(mBase, uint32(v7))) = v24
											F_errmsg_internal(m, int32(_a_F_reset_conflict_slot_xmin_to_safe_horizon_2), v7)
											mBase = m.M
											v74 = m.ExcPending
											if v74 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_reset_conflict_slot_xmin_to_safe_horizon_3), int32(1614), int32(_a_F_reset_conflict_slot_xmin_to_safe_horizon_4))
												mBase = m.M
												v79 = m.ExcPending
												if v79 != 0 {
													return
												} else {
													F_ReplicationSlotsComputeRequiredXmin(m, int32(1))
													mBase = m.M
													v82 = m.ExcPending
													if v82 != 0 {
														return
													} else {
														v84 = *(*int32)(unsafe.Add(mBase, _c_F_reset_conflict_slot_xmin_to_safe_horizon[0]))
														F_LWLockRelease(m, v84+int32(512))
														mBase = m.M
														v88 = m.ExcPending
														if v88 != 0 {
															return
														} else {
															v90 = *(*int32)(unsafe.Add(mBase, _c_F_reset_conflict_slot_xmin_to_safe_horizon[0]))
															F_LWLockRelease(m, v90+int32(_a_F_reset_conflict_slot_xmin_to_safe_horizon_0))
															mBase = m.M
															v94 = m.ExcPending
															if v94 != 0 {
																return
															} else {
																F_ReplicationSlotMarkDirty(m)
																mBase = m.M
																v96 = m.ExcPending
																if v96 != 0 {
																	return
																} else {
																	F_ReplicationSlotSave(m)
																	mBase = m.M
																	v98 = m.ExcPending
																	if v98 != 0 {
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
											F_ReplicationSlotsComputeRequiredXmin(m, int32(1))
											mBase = m.M
											v82 = m.ExcPending
											if v82 != 0 {
												return
											} else {
												v84 = *(*int32)(unsafe.Add(mBase, _c_F_reset_conflict_slot_xmin_to_safe_horizon[0]))
												F_LWLockRelease(m, v84+int32(512))
												mBase = m.M
												v88 = m.ExcPending
												if v88 != 0 {
													return
												} else {
													v90 = *(*int32)(unsafe.Add(mBase, _c_F_reset_conflict_slot_xmin_to_safe_horizon[0]))
													F_LWLockRelease(m, v90+int32(_a_F_reset_conflict_slot_xmin_to_safe_horizon_0))
													mBase = m.M
													v94 = m.ExcPending
													if v94 != 0 {
														return
													} else {
														F_ReplicationSlotMarkDirty(m)
														mBase = m.M
														v96 = m.ExcPending
														if v96 != 0 {
															return
														} else {
															F_ReplicationSlotSave(m)
															mBase = m.M
															v98 = m.ExcPending
															if v98 != 0 {
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
								}
							} else {
								v61 = *(*int32)(unsafe.Add(mBase, _c_F_reset_conflict_slot_xmin_to_safe_horizon[1]))
								*(*int32)(unsafe.Add(mBase, uint32(v61)+96)) = v24
								*(*int32)(unsafe.Add(mBase, uint32(v61)+16)) = v24
								v64 = int32(0)
								atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v61))), uint32(v64))
								v69 = F_errstart(m, int32(14), v64)
								mBase = m.M
								v70 = m.ExcPending
								if v70 != 0 {
									return
								} else {
									if v69 != 0 {
										*(*int32)(unsafe.Add(mBase, uint32(v7))) = v24
										F_errmsg_internal(m, int32(_a_F_reset_conflict_slot_xmin_to_safe_horizon_2), v7)
										mBase = m.M
										v74 = m.ExcPending
										if v74 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_reset_conflict_slot_xmin_to_safe_horizon_3), int32(1614), int32(_a_F_reset_conflict_slot_xmin_to_safe_horizon_4))
											mBase = m.M
											v79 = m.ExcPending
											if v79 != 0 {
												return
											} else {
												F_ReplicationSlotsComputeRequiredXmin(m, int32(1))
												mBase = m.M
												v82 = m.ExcPending
												if v82 != 0 {
													return
												} else {
													v84 = *(*int32)(unsafe.Add(mBase, _c_F_reset_conflict_slot_xmin_to_safe_horizon[0]))
													F_LWLockRelease(m, v84+int32(512))
													mBase = m.M
													v88 = m.ExcPending
													if v88 != 0 {
														return
													} else {
														v90 = *(*int32)(unsafe.Add(mBase, _c_F_reset_conflict_slot_xmin_to_safe_horizon[0]))
														F_LWLockRelease(m, v90+int32(_a_F_reset_conflict_slot_xmin_to_safe_horizon_0))
														mBase = m.M
														v94 = m.ExcPending
														if v94 != 0 {
															return
														} else {
															F_ReplicationSlotMarkDirty(m)
															mBase = m.M
															v96 = m.ExcPending
															if v96 != 0 {
																return
															} else {
																F_ReplicationSlotSave(m)
																mBase = m.M
																v98 = m.ExcPending
																if v98 != 0 {
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
										F_ReplicationSlotsComputeRequiredXmin(m, int32(1))
										mBase = m.M
										v82 = m.ExcPending
										if v82 != 0 {
											return
										} else {
											v84 = *(*int32)(unsafe.Add(mBase, _c_F_reset_conflict_slot_xmin_to_safe_horizon[0]))
											F_LWLockRelease(m, v84+int32(512))
											mBase = m.M
											v88 = m.ExcPending
											if v88 != 0 {
												return
											} else {
												v90 = *(*int32)(unsafe.Add(mBase, _c_F_reset_conflict_slot_xmin_to_safe_horizon[0]))
												F_LWLockRelease(m, v90+int32(_a_F_reset_conflict_slot_xmin_to_safe_horizon_0))
												mBase = m.M
												v94 = m.ExcPending
												if v94 != 0 {
													return
												} else {
													F_ReplicationSlotMarkDirty(m)
													mBase = m.M
													v96 = m.ExcPending
													if v96 != 0 {
														return
													} else {
														F_ReplicationSlotSave(m)
														mBase = m.M
														v98 = m.ExcPending
														if v98 != 0 {
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
							}
						} else {
							v43 = *(*int32)(unsafe.Add(mBase, _c_F_reset_conflict_slot_xmin_to_safe_horizon[0]))
							F_LWLockRelease(m, v43+int32(512))
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return
							} else {
								v49 = *(*int32)(unsafe.Add(mBase, _c_F_reset_conflict_slot_xmin_to_safe_horizon[0]))
								F_LWLockRelease(m, v49+int32(_a_F_reset_conflict_slot_xmin_to_safe_horizon_0))
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
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
	}
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
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)+224))
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
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v51 int32
	_ = v51
	var v61 int32
	_ = v61
	var v62 int64
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v84 float64
	_ = v84
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v107 float64
	_ = v107
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v17 = F_get_oprrest(m, l1)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return float64(0)
	} else {
		if v17 == int32(0) {
			v107 = float64(0.5)
			m.G0 = v15 + int32(16)
			return v107
		} else {
			v28 = m.G0
			v30 = v28 - int32(128)
			m.G0 = v30
			v33 = v30 + int32(12)
			v35 = *(*int32)(unsafe.Add(mBase, _c_F_restriction_selectivity[0]))
			F_fmgr_info_cxt_security(m, v17, v33, v35, int32(0))
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return float64(0)
			} else {
				v39 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v30)+120)) = uint8(v39)
				*(*int64)(unsafe.Add(mBase, uint32(v30)+112)) = base.I64_extend_i32_s(l4)
				*(*uint8)(unsafe.Add(mBase, uint32(v30)+104)) = uint8(v39)
				*(*int64)(unsafe.Add(mBase, uint32(v30)+96)) = base.I64_extend_i32_u(l2)
				*(*uint8)(unsafe.Add(mBase, uint32(v30)+88)) = uint8(v39)
				*(*int64)(unsafe.Add(mBase, uint32(v30)+80)) = base.I64_extend_i32_u(l1)
				*(*uint8)(unsafe.Add(mBase, uint32(v30)+72)) = uint8(v39)
				*(*int64)(unsafe.Add(mBase, uint32(v30)+64)) = base.I64_extend_i32_u(l0)
				v51 = int32(4)
				*(*uint16)(unsafe.Add(mBase, uint32(v30)+58)) = uint16(v51)
				*(*uint8)(unsafe.Add(mBase, uint32(v30)+56)) = uint8(v39)
				*(*int32)(unsafe.Add(mBase, uint32(v30)+52)) = l3
				*(*int64)(unsafe.Add(mBase, uint32(v30)+44)) = int64(0)
				*(*int32)(unsafe.Add(mBase, uint32(v30)+40)) = v33
				v61 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
				v62 = m.T0[v61].(func(*base.Module, int32) int64)(m, v30+int32(40))
				mBase = m.M
				v63 = m.ExcPending
				if v63 != 0 {
					return float64(0)
				} else {
					v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+56)))
					if v64 == int32(1) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
							return float64(0)
						} else {
							v71 = *(*int32)(unsafe.Add(mBase, uint32(v30)+16))
							*(*int32)(unsafe.Add(mBase, uint32(v30))) = v71
							F_errmsg_internal(m, int32(_a_F_restriction_selectivity_0), v30)
							mBase = m.M
							v75 = m.ExcPending
							if v75 != 0 {
								return float64(0)
							} else {
								F_errfinish(m, int32(_a_F_restriction_selectivity_1), int32(1219), int32(_a_F_restriction_selectivity_2))
								mBase = m.M
								v80 = m.ExcPending
								if v80 != 0 {
									return float64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						m.G0 = v30 + int32(128)
						v84 = base.F64_reinterpret_i64(v62)
						if base.F64_lt(v84, float64(0))|base.F64_gt(v84, float64(1)) == int32(0) {
							v107 = v84
							m.G0 = v15 + int32(16)
							return v107
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v95 = m.ExcPending
							if v95 != 0 {
								return float64(0)
							} else {
								*(*float64)(unsafe.Add(mBase, uint32(v15))) = v84
								F_errmsg_internal(m, int32(_a_F_restriction_selectivity_3), v15)
								mBase = m.M
								v99 = m.ExcPending
								if v99 != 0 {
									return float64(0)
								} else {
									F_errfinish(m, int32(_a_F_restriction_selectivity_4), int32(2226), int32(_a_F_restriction_selectivity_5))
									mBase = m.M
									v104 = m.ExcPending
									if v104 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = int32(1058)
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
					F_errfinish(m, int32(_a_F_rollback_prepared_cb_wrapper_4), int32(1157), int32(_a_F_rollback_prepared_cb_wrapper_5))
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
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v91 int32
	_ = v91
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v140 int32
	_ = v140
	var v149 int32
	_ = v149
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v204 int32
	_ = v204
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v235 int32
	_ = v235
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v252 int32
	_ = v252
	var v261 int32
	_ = v261
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
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
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
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
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v458 int32
	_ = v458
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v501 int32
	_ = v501
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v542 int32
	_ = v542
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v561 int32
	_ = v561
	var v564 int32
	_ = v564
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v574 int32
	_ = v574
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v584 int32
	_ = v584
	var v587 int32
	_ = v587
	var v589 int32
	_ = v589
	var v593 int32
	_ = v593
	var v597 int32
	_ = v597
	var v600 int32
	_ = v600
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
	var v608 int32
	_ = v608
	var v613 int32
	_ = v613
	var v617 int32
	_ = v617
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v627 int32
	_ = v627
	var v629 int32
	_ = v629
	var v631 int32
	_ = v631
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v640 int32
	_ = v640
	var v642 int32
	_ = v642
	var v644 int32
	_ = v644
	var v646 int32
	_ = v646
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v665 int32
	_ = v665
	var v669 int32
	_ = v669
	var v672 int32
	_ = v672
	var v674 int32
	_ = v674
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v680 int32
	_ = v680
	var v684 int32
	_ = v684
	var v690 int32
	_ = v690
	var v696 int32
	_ = v696
	var v699 int32
	_ = v699
	var v701 int32
	_ = v701
	var v705 int32
	_ = v705
	var v711 int32
	_ = v711
	var v714 int32
	_ = v714
	var v724 int32
	_ = v724
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	goto L2
L1:
	;
	return v724
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v15
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v7
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v15
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v59 < v7 {
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
	v724 = v41
	goto L1
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v7
	v278 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v278
	v280 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v278 < v280 {
		v724 = int32(0)
		goto L1
	} else {
		goto L82
	}
L20:
	;
	if v99 < int32(0) {
		goto L19
	} else {
		goto L35
	}
L21:
	;
	v61 = v7
	goto L23
L22:
	;
	v61 = v59
	goto L23
L23:
	;
	v68 = v7
	goto L25
L24:
	;
	v99 = v79
	goto L20
L25:
	;
	if v68 == v61 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v99 = int32(-1)
	goto L20
L28:
	;
	goto L29
L29:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72+v68))))
	if int32(220) < v74 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v91 = v68 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v91
	v68 = v91
	goto L25
L31:
	;
	v76 = v74 - int32(192)
	if v76 < int32(0) {
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v79 = int32(1)
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v76)>>(uint(int32(3))%32)))+uint32(_c_F_russian_KOI8_R_stem[0]))))
	if int32(base.Ui32(v83)>>(uint(v76&int32(7))%32))&v79 != 0 {
		goto L24
	} else {
		goto L33
	}
L33:
	;
	goto L30
L35:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v103 = v102 + v99
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v103
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v103
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v115 < v103 {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	if v158 < int32(0) {
		goto L19
	} else {
		goto L50
	}
L37:
	;
	v117 = v103
	goto L39
L38:
	;
	v117 = v115
	goto L39
L39:
	;
	v123 = v103
	goto L41
L40:
	;
	v158 = int32(1)
	goto L36
L41:
	;
	if v123 == v117 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v158 = int32(-1)
	goto L36
L44:
	;
	goto L45
L45:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130+v123))))
	if int32(220) < v132 {
		goto L40
	} else {
		goto L46
	}
L46:
	;
	v134 = v132 - int32(192)
	if v134 < int32(0) {
		goto L40
	} else {
		goto L47
	}
L47:
	;
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v134)>>(uint(int32(3))%32)))+uint32(_c_F_russian_KOI8_R_stem[0]))))
	if int32(base.Ui32(v140)>>(uint(v134&int32(7))%32))&int32(1) == int32(0) {
		goto L40
	} else {
		goto L48
	}
L48:
	;
	v149 = v123 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v149
	v123 = v149
	goto L41
L50:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v162 = v161 + v158
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v162
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v172 < v162 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	if v212 < int32(0) {
		goto L19
	} else {
		goto L66
	}
L52:
	;
	v174 = v162
	goto L54
L53:
	;
	v174 = v172
	goto L54
L54:
	;
	v181 = v162
	goto L56
L55:
	;
	v212 = v192
	goto L51
L56:
	;
	if v181 == v174 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v212 = int32(-1)
	goto L51
L59:
	;
	goto L60
L60:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v185+v181))))
	if int32(220) < v187 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v204 = v181 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v204
	v181 = v204
	goto L56
L62:
	;
	v189 = v187 - int32(192)
	if v189 < int32(0) {
		goto L61
	} else {
		goto L63
	}
L63:
	;
	v192 = int32(1)
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v189)>>(uint(int32(3))%32)))+uint32(_c_F_russian_KOI8_R_stem[0]))))
	if int32(base.Ui32(v196)>>(uint(v189&int32(7))%32))&v192 != 0 {
		goto L55
	} else {
		goto L64
	}
L64:
	;
	goto L61
L66:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v216 = v215 + v212
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v216
	v227 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v227 < v216 {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	if v270 < int32(0) {
		goto L19
	} else {
		goto L81
	}
L68:
	;
	v229 = v216
	goto L70
L69:
	;
	v229 = v227
	goto L70
L70:
	;
	v235 = v216
	goto L72
L71:
	;
	v270 = int32(1)
	goto L67
L72:
	;
	if v235 == v229 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v270 = int32(-1)
	goto L67
L75:
	;
	goto L76
L76:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v242+v235))))
	if int32(220) < v244 {
		goto L71
	} else {
		goto L77
	}
L77:
	;
	v246 = v244 - int32(192)
	if v246 < int32(0) {
		goto L71
	} else {
		goto L78
	}
L78:
	;
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v246)>>(uint(int32(3))%32)))+uint32(_c_F_russian_KOI8_R_stem[0]))))
	if int32(base.Ui32(v252)>>(uint(v246&int32(7))%32))&int32(1) == int32(0) {
		goto L71
	} else {
		goto L79
	}
L79:
	;
	v261 = v235 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v261
	v235 = v261
	goto L72
L81:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v273 + v270
	goto L19
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v278
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v280
	if v278 <= v280 {
		v339 = v280
		goto L84
	} else {
		goto L85
	}
L83:
	;
	v584 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v584
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v584
	v587 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v584 <= v587 {
		v606 = v587
		v607 = v584
		v608 = v584
		goto L155
	} else {
		goto L156
	}
L84:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v341
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v341
	v345 = v341 - int32(1)
	if v345 <= v339 {
		v358 = v341
		goto L99
	} else {
		goto L100
	}
L85:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v287 = int32(1)
	v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v285+v278-v287))))
	if base.B2i32(v289&int32(224) != int32(192))|base.B2i32(v287<<(uint(v289)%32)&int32(25166336) == int32(0)) != 0 {
		v339 = v280
		goto L84
	} else {
		goto L86
	}
L86:
	;
	v304 = F_find_among_b(m, l0, int32(_a_F_russian_KOI8_R_stem_1), int32(9), int32(0))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L16
	} else {
		goto L87
	}
L87:
	;
	if v304 == int32(0) {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v339 = v308
	goto L84
L89:
	;
	goto L90
L90:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v309
	switch v304 - int32(1) {
	case 0:
		goto L92
	case 1:
		goto L91
	default:
		goto L83
	}
L91:
	;
	v335 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v335 {
		goto L83
	} else {
		goto L96
	}
L92:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v309 <= v313 {
		v339 = v313
		goto L84
	} else {
		goto L93
	}
L93:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v315+v309-int32(1)))))
	v321 = v319 - int32(193)
	v322 = int32(0)
	if base.B2i32(v321 == v322)|base.B2i32(v321 == int32(16)) == v322 {
		v339 = v313
		goto L84
	} else {
		goto L94
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v309 - int32(1)
	v332 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v332 {
		goto L83
	} else {
		goto L95
	}
L95:
	;
	v724 = v332
	goto L1
L96:
	;
	v724 = v335
	goto L1
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v367
	v371 = v367 - int32(1)
	v372 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v371 <= v372 {
		v394 = v368
		goto L108
	} else {
		goto L109
	}
L98:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v360
	v362 = F_slice_del(m, l0)
	mBase = m.M
	if v362 < int32(0) {
		v724 = v362
		goto L1
	} else {
		goto L104
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v358
	v367 = v358
	v368 = v358
	goto L97
L100:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v347+v345))))
	switch v349 - int32(209) {
	case 0, 7:
		goto L101
	default:
		v358 = v341
		goto L99
	}
L101:
	;
	v355 = F_find_among_b(m, l0, int32(_a_F_russian_KOI8_R_stem_2), int32(2), int32(0))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L16
	} else {
		goto L102
	}
L102:
	;
	if v355 != 0 {
		goto L98
	} else {
		goto L103
	}
L103:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v358 = v357
	goto L99
L104:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v366 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v367 = v365
	v368 = v366
	goto L97
L105:
	;
	if v578 != 0 {
		v724 = v577
		goto L1
	} else {
		goto L154
	}
L106:
	;
	v530 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v531 = v530 - v395
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v531
	v533 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v531
	v536 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v531 <= v536 {
		v568 = v533
		goto L142
	} else {
		goto L143
	}
L107:
	;
	v469 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v469
	v471 = F_slice_del(m, l0)
	mBase = m.M
	if v471 < int32(0) {
		v724 = v471
		goto L1
	} else {
		goto L131
	}
L108:
	;
	v395 = v368 - v367
	v396 = v394 - v395
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v396
	v398 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v396
	v401 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v396 <= v401 {
		v458 = v398
		goto L114
	} else {
		goto L115
	}
L109:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v374+v371))))
	if v376&int32(224) != int32(192) {
		v394 = v368
		goto L108
	} else {
		goto L110
	}
L110:
	;
	if int32(1)<<(uint(v376)%32)&int32(_a_F_russian_KOI8_R_stem_3) == int32(0) {
		v394 = v368
		goto L108
	} else {
		goto L111
	}
L111:
	;
	v390 = F_find_among_b(m, l0, int32(_a_F_russian_KOI8_R_stem_4), int32(26), int32(0))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L16
	} else {
		goto L112
	}
L112:
	;
	if v390 != 0 {
		goto L107
	} else {
		goto L113
	}
L113:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v394 = v392
	goto L108
L114:
	;
	v462 = int32(base.Ui32(v458) >> (uint(int32(31)) % 32))
	if v458 != 0 {
		goto L126
	} else {
		goto L127
	}
L115:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v405 = int32(1)
	v407 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v403+v396-v405))))
	if base.B2i32(v407&int32(224) != int32(192))|base.B2i32(v405<<(uint(v407)%32)&int32(51443235) == int32(0)) != 0 {
		v458 = v398
		goto L114
	} else {
		goto L116
	}
L116:
	;
	v422 = F_find_among_b(m, l0, int32(_a_F_russian_KOI8_R_stem_5), int32(46), int32(0))
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L16
	} else {
		goto L117
	}
L117:
	;
	if v422 == int32(0) {
		v458 = v398
		goto L114
	} else {
		goto L118
	}
L118:
	;
	v426 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v426
	switch v422 - int32(1) {
	case 0:
		goto L121
	case 1:
		goto L120
	default:
		goto L119
	}
L119:
	;
	v458 = int32(1)
	goto L114
L120:
	;
	v452 = F_slice_del(m, l0)
	mBase = m.M
	if v452 < int32(0) {
		v458 = v452
		goto L114
	} else {
		goto L125
	}
L121:
	;
	v430 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v426 <= v430 {
		v458 = v398
		goto L114
	} else {
		goto L122
	}
L122:
	;
	v432 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v436 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v432+v426-int32(1)))))
	v438 = v436 - int32(193)
	v439 = int32(0)
	if base.B2i32(v438 == v439)|base.B2i32(v438 == int32(16)) == v439 {
		v458 = v398
		goto L114
	} else {
		goto L123
	}
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v426 - int32(1)
	v449 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v449 {
		goto L119
	} else {
		goto L124
	}
L124:
	;
	v458 = v449
	goto L114
L125:
	;
	goto L119
L126:
	;
	v464 = v462
	goto L128
L127:
	;
	v464 = int32(15)
	goto L128
L128:
	;
	if v464 == int32(0) {
		goto L83
	} else {
		goto L129
	}
L129:
	;
	if v464 == int32(15) {
		goto L106
	} else {
		goto L130
	}
L130:
	;
	v577 = v458
	v578 = v462
	goto L105
L131:
	;
	v474 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v474
	v476 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v474 <= v476 {
		goto L83
	} else {
		goto L132
	}
L132:
	;
	v478 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v480 = int32(1)
	v482 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v478+v474-v480))))
	if base.B2i32(v482&int32(224) != int32(192))|base.B2i32(v480<<(uint(v482)%32)&int32(671113216) == int32(0)) != 0 {
		goto L83
	} else {
		goto L133
	}
L133:
	;
	v497 = F_find_among_b(m, l0, int32(_a_F_russian_KOI8_R_stem_6), int32(8), int32(0))
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L16
	} else {
		goto L134
	}
L134:
	;
	if v497 == int32(0) {
		goto L83
	} else {
		goto L135
	}
L135:
	;
	v501 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v501
	switch v497 - int32(1) {
	case 0:
		goto L137
	case 1:
		goto L136
	default:
		goto L83
	}
L136:
	;
	v527 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v527 {
		goto L83
	} else {
		goto L141
	}
L137:
	;
	v505 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v501 <= v505 {
		goto L83
	} else {
		goto L138
	}
L138:
	;
	v507 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v511 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v507+v501-int32(1)))))
	v513 = v511 - int32(193)
	v514 = int32(0)
	if base.B2i32(v513 == v514)|base.B2i32(v513 == int32(16)) == v514 {
		goto L83
	} else {
		goto L139
	}
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v501 - int32(1)
	v524 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v524 {
		goto L83
	} else {
		goto L140
	}
L140:
	;
	v724 = v524
	goto L1
L141:
	;
	v724 = v527
	goto L1
L142:
	;
	if v568 == int32(0) {
		goto L83
	} else {
		goto L150
	}
L143:
	;
	v538 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v540 = int32(1)
	v542 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v538+v531-v540))))
	if base.B2i32(v542&int32(224) != int32(192))|base.B2i32(v540<<(uint(v542)%32)&int32(60991267) == int32(0)) != 0 {
		v568 = v533
		goto L142
	} else {
		goto L144
	}
L144:
	;
	v557 = F_find_among_b(m, l0, int32(_a_F_russian_KOI8_R_stem_7), int32(36), int32(0))
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L16
	} else {
		goto L145
	}
L145:
	;
	if v557 == int32(0) {
		v568 = v533
		goto L142
	} else {
		goto L146
	}
L146:
	;
	v561 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v561
	v564 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v564 {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	v567 = int32(1)
	goto L149
L148:
	;
	v567 = v564
	goto L149
L149:
	;
	v568 = v567
	goto L142
L150:
	;
	if v568 < int32(0) {
		goto L151
	} else {
		goto L152
	}
L151:
	;
	v574 = v568
	goto L153
L152:
	;
	v574 = v458
	goto L153
L153:
	;
	v577 = v574
	v578 = int32(base.Ui32(v568) >> (uint(int32(31)) % 32))
	goto L105
L154:
	;
	goto L83
L155:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v607
	if v607-int32(2) <= v606 {
		goto L159
	} else {
		goto L160
	}
L156:
	;
	v589 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v593 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v589+v584-int32(1)))))
	if v593 != int32(201) {
		v606 = v587
		v607 = v584
		v608 = v584
		goto L155
	} else {
		goto L157
	}
L157:
	;
	v597 = v584 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v597
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v597
	v600 = F_slice_del(m, l0)
	mBase = m.M
	if v600 < int32(0) {
		v724 = v600
		goto L1
	} else {
		goto L158
	}
L158:
	;
	v603 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v604 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v605 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v606 = v603
	v607 = v604
	v608 = v605
	goto L155
L159:
	;
	v635 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v637 = v635 + (v607 - v608)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v637
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v637
	v640 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v637 <= v640 {
		goto L166
	} else {
		goto L167
	}
L160:
	;
	v613 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v617 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v613+v607-int32(1)))))
	switch v617 - int32(212) {
	case 0, 4:
		goto L161
	default:
		goto L159
	}
L161:
	;
	v623 = F_find_among_b(m, l0, int32(_a_F_russian_KOI8_R_stem_8), int32(2), int32(0))
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L16
	} else {
		goto L162
	}
L162:
	;
	if v623 == int32(0) {
		goto L159
	} else {
		goto L163
	}
L163:
	;
	v627 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v627
	v629 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v627 < v629 {
		goto L159
	} else {
		goto L164
	}
L164:
	;
	v631 = F_slice_del(m, l0)
	mBase = m.M
	if v631 < int32(0) {
		v724 = v631
		goto L1
	} else {
		goto L165
	}
L165:
	;
	goto L159
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v7
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v7
	v724 = int32(1)
	goto L1
L167:
	;
	v642 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v644 = int32(1)
	v646 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v642+v637-v644))))
	if base.B2i32(v646&int32(224) != int32(192))|base.B2i32(v644<<(uint(v646)%32)&int32(151011360) == int32(0)) != 0 {
		goto L166
	} else {
		goto L168
	}
L168:
	;
	v661 = F_find_among_b(m, l0, int32(_a_F_russian_KOI8_R_stem_9), int32(4), int32(0))
	mBase = m.M
	v662 = m.ExcPending
	if v662 != 0 {
		goto L16
	} else {
		goto L169
	}
L169:
	;
	if v661 == int32(0) {
		goto L166
	} else {
		goto L170
	}
L170:
	;
	v665 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v665
	switch v661 - int32(1) {
	case 0:
		goto L173
	case 1:
		goto L172
	case 2:
		goto L171
	default:
		goto L166
	}
L171:
	;
	v714 = F_slice_del(m, l0)
	mBase = m.M
	if v714 < int32(0) {
		v724 = v714
		goto L1
	} else {
		goto L183
	}
L172:
	;
	v699 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v665 <= v699 {
		goto L166
	} else {
		goto L180
	}
L173:
	;
	v669 = F_slice_del(m, l0)
	mBase = m.M
	if v669 < int32(0) {
		v724 = v669
		goto L1
	} else {
		goto L174
	}
L174:
	;
	v672 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v672
	v674 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v672 <= v674 {
		goto L166
	} else {
		goto L175
	}
L175:
	;
	v676 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v677 = v676 + v672
	v680 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v677-int32(1)))))
	if v680 != int32(206) {
		goto L166
	} else {
		goto L176
	}
L176:
	;
	v684 = v672 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v684
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v684
	if v684 <= v674 {
		goto L166
	} else {
		goto L177
	}
L177:
	;
	v690 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v677-int32(2)))))
	if v690 != int32(206) {
		goto L166
	} else {
		goto L178
	}
L178:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v672 - int32(2)
	v696 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v696 {
		goto L166
	} else {
		goto L179
	}
L179:
	;
	v724 = v696
	goto L1
L180:
	;
	v701 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v705 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v701+v665-int32(1)))))
	if v705 != int32(206) {
		goto L166
	} else {
		goto L181
	}
L181:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v665 - int32(1)
	v711 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v711 {
		goto L166
	} else {
		goto L182
	}
L182:
	;
	v724 = v711
	goto L1
L183:
	;
	goto L166
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
	var v150 int32
	_ = v150
	var v158 int32
	_ = v158
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v209 int32
	_ = v209
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v235 int32
	_ = v235
	var v242 int32
	_ = v242
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v281 int32
	_ = v281
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v332 int32
	_ = v332
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v358 int32
	_ = v358
	var v366 int32
	_ = v366
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v404 int32
	_ = v404
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v426 int32
	_ = v426
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v442 int32
	_ = v442
	var v455 int32
	_ = v455
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v475 int32
	_ = v475
	var v481 int32
	_ = v481
	var v488 int32
	_ = v488
	var v499 int32
	_ = v499
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v526 int32
	_ = v526
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v539 int32
	_ = v539
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v548 int32
	_ = v548
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v564 int32
	_ = v564
	var v577 int32
	_ = v577
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v597 int32
	_ = v597
	var v603 int32
	_ = v603
	var v611 int32
	_ = v611
	var v622 int32
	_ = v622
	var v625 int32
	_ = v625
	var v630 int32
	_ = v630
	var v632 int32
	_ = v632
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
	var v650 int32
	_ = v650
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v656 int32
	_ = v656
	var v659 int32
	_ = v659
	var v663 int32
	_ = v663
	var v666 int32
	_ = v666
	var v668 int32
	_ = v668
	var v670 int32
	_ = v670
	var v672 int32
	_ = v672
	var v675 int32
	_ = v675
	var v678 int32
	_ = v678
	var v681 int32
	_ = v681
	var v685 int32
	_ = v685
	var v688 int32
	_ = v688
	var v691 int32
	_ = v691
	var v696 int32
	_ = v696
	var v699 int32
	_ = v699
	var v703 int32
	_ = v703
	var v707 int32
	_ = v707
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v718 int32
	_ = v718
	var v720 int32
	_ = v720
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v735 int32
	_ = v735
	var v738 int32
	_ = v738
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v747 int32
	_ = v747
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v754 int32
	_ = v754
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v760 int32
	_ = v760
	var v763 int32
	_ = v763
	var v767 int32
	_ = v767
	var v770 int32
	_ = v770
	var v772 int32
	_ = v772
	var v774 int32
	_ = v774
	var v776 int32
	_ = v776
	var v779 int32
	_ = v779
	var v782 int32
	_ = v782
	var v785 int32
	_ = v785
	var v789 int32
	_ = v789
	var v792 int32
	_ = v792
	var v795 int32
	_ = v795
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v802 int32
	_ = v802
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v812 int32
	_ = v812
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v819 int32
	_ = v819
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v825 int32
	_ = v825
	var v828 int32
	_ = v828
	var v832 int32
	_ = v832
	var v835 int32
	_ = v835
	var v837 int32
	_ = v837
	var v839 int32
	_ = v839
	var v841 int32
	_ = v841
	var v844 int32
	_ = v844
	var v847 int32
	_ = v847
	var v850 int32
	_ = v850
	var v854 int32
	_ = v854
	var v857 int32
	_ = v857
	var v860 int32
	_ = v860
	var v866 int32
	_ = v866
	var v870 int32
	_ = v870
	var v872 int32
	_ = v872
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v885 int32
	_ = v885
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v890 int32
	_ = v890
	var v893 int32
	_ = v893
	var v896 int32
	_ = v896
	var v901 int32
	_ = v901
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v910 int32
	_ = v910
	var v913 int32
	_ = v913
	var v915 int32
	_ = v915
	var v918 int32
	_ = v918
	var v921 int32
	_ = v921
	var v924 int32
	_ = v924
	var v928 int32
	_ = v928
	var v931 int32
	_ = v931
	var v933 int32
	_ = v933
	var v935 int32
	_ = v935
	var v938 int32
	_ = v938
	var v939 int32
	_ = v939
	var v940 int32
	_ = v940
	var v942 int32
	_ = v942
	var v944 int32
	_ = v944
	var v948 int32
	_ = v948
	var v952 int32
	_ = v952
	var v958 int32
	_ = v958
	var v959 int32
	_ = v959
	var v962 int32
	_ = v962
	var v964 int32
	_ = v964
	var v966 int32
	_ = v966
	var v970 int32
	_ = v970
	var v972 int32
	_ = v972
	var v978 int32
	_ = v978
	var v979 int32
	_ = v979
	var v982 int32
	_ = v982
	var v986 int32
	_ = v986
	var v989 int32
	_ = v989
	var v991 int32
	_ = v991
	var v993 int32
	_ = v993
	var v996 int32
	_ = v996
	var v999 int32
	_ = v999
	var v1002 int32
	_ = v1002
	var v1006 int32
	_ = v1006
	var v1009 int32
	_ = v1009
	var v1011 int32
	_ = v1011
	var v1013 int32
	_ = v1013
	var v1016 int32
	_ = v1016
	var v1019 int32
	_ = v1019
	var v1022 int32
	_ = v1022
	var v1026 int32
	_ = v1026
	var v1029 int32
	_ = v1029
	var v1032 int32
	_ = v1032
	var v1034 int32
	_ = v1034
	var v1036 int32
	_ = v1036
	var v1037 int32
	_ = v1037
	var v1040 int32
	_ = v1040
	var v1043 int32
	_ = v1043
	var v1047 int32
	_ = v1047
	var v1050 int32
	_ = v1050
	var v1053 int32
	_ = v1053
	var v1062 int32
	_ = v1062
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	goto L2
L1:
	;
	return v1062
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
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v133
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v133
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v158 = v7
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
	v1062 = v126
	goto L1
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v7
	v630 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v630
	v632 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v630 < v632 {
		v1062 = int32(0)
		goto L1
	} else {
		goto L146
	}
L44:
	;
	if v253 < int32(0) {
		goto L43
	} else {
		goto L69
	}
L45:
	;
	v253 = v225
	goto L44
L46:
	;
	if v133 <= v158 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v253 = int32(-1)
	goto L44
L49:
	;
	goto L50
L50:
	;
	v165 = int32(1)
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158+v150))))
	if base.Ui32(v167) < base.Ui32(int32(192)) {
		v224 = v167
		v225 = v165
		goto L51
	} else {
		goto L52
	}
L51:
	;
	if int32(1103) < v224 {
		goto L64
	} else {
		goto L65
	}
L52:
	;
	v171 = v158 + int32(1)
	if v171 == v133 {
		v224 = v167
		v225 = v165
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v171+v150))))
	v176 = v174 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v167) {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180+v150))))
	v192 = v190 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v167) {
		goto L60
	} else {
		goto L61
	}
L55:
	;
	v180 = v158 + int32(2)
	if v180 != v133 {
		goto L54
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v224 = v167<<(uint(int32(6))%32)&int32(1984) | v176
	v225 = int32(2)
	goto L51
L58:
	;
	goto L57
L59:
	;
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150+v196))))
	v224 = v209&int32(63) | (v167<<(uint(int32(18))%32)&int32(_a_F_russian_UTF_8_stem_2) | v176<<(uint(int32(12))%32) | v192<<(uint(int32(6))%32))
	v225 = int32(4)
	goto L51
L60:
	;
	v196 = v158 + int32(3)
	if v196 != v133 {
		goto L59
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v224 = v167<<(uint(int32(12))%32)&int32(_a_F_russian_UTF_8_stem_3) | v176<<(uint(int32(6))%32) | v192
	v225 = int32(3)
	goto L51
L63:
	;
	goto L62
L64:
	;
	v242 = v225 + v158
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v242
	v158 = v242
	goto L46
L65:
	;
	v229 = v224 - int32(1072)
	if v229 < int32(0) {
		goto L64
	} else {
		goto L66
	}
L66:
	;
	v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v229)>>(uint(int32(3))%32)))+uint32(_c_F_russian_UTF_8_stem[0]))))
	if int32(base.Ui32(v235)>>(uint(v229&int32(7))%32))&int32(1) != 0 {
		goto L45
	} else {
		goto L67
	}
L67:
	;
	goto L64
L69:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v257 = v256 + v253
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v257
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v257
	v272 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v273 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v281 = v257
	goto L72
L70:
	;
	if v377 < int32(0) {
		goto L43
	} else {
		goto L94
	}
L71:
	;
	v377 = v348
	goto L70
L72:
	;
	if v272 <= v281 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v377 = int32(-1)
	goto L70
L75:
	;
	goto L76
L76:
	;
	v288 = int32(1)
	v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v281+v273))))
	if base.Ui32(v290) < base.Ui32(int32(192)) {
		v347 = v290
		v348 = v288
		goto L77
	} else {
		goto L78
	}
L77:
	;
	if int32(1103) < v347 {
		goto L71
	} else {
		goto L90
	}
L78:
	;
	v294 = v281 + int32(1)
	if v294 == v272 {
		v347 = v290
		v348 = v288
		goto L77
	} else {
		goto L79
	}
L79:
	;
	v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v294+v273))))
	v299 = v297 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v290) {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	v313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v303+v273))))
	v315 = v313 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v290) {
		goto L86
	} else {
		goto L87
	}
L81:
	;
	v303 = v281 + int32(2)
	if v303 != v272 {
		goto L80
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	v347 = v290<<(uint(int32(6))%32)&int32(1984) | v299
	v348 = int32(2)
	goto L77
L84:
	;
	goto L83
L85:
	;
	v332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v273+v319))))
	v347 = v332&int32(63) | (v290<<(uint(int32(18))%32)&int32(_a_F_russian_UTF_8_stem_2) | v299<<(uint(int32(12))%32) | v315<<(uint(int32(6))%32))
	v348 = int32(4)
	goto L77
L86:
	;
	v319 = v281 + int32(3)
	if v319 != v272 {
		goto L85
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	v347 = v290<<(uint(int32(12))%32)&int32(_a_F_russian_UTF_8_stem_3) | v299<<(uint(int32(6))%32) | v315
	v348 = int32(3)
	goto L77
L89:
	;
	goto L88
L90:
	;
	v352 = v347 - int32(1072)
	if v352 < int32(0) {
		goto L71
	} else {
		goto L91
	}
L91:
	;
	v358 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v352)>>(uint(int32(3))%32)))+uint32(_c_F_russian_UTF_8_stem[0]))))
	if int32(base.Ui32(v358)>>(uint(v352&int32(7))%32))&int32(1) == int32(0) {
		goto L71
	} else {
		goto L92
	}
L92:
	;
	v366 = v348 + v281
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v366
	v281 = v366
	goto L72
L94:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v381 = v380 + v377
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v381
	v395 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v396 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v404 = v381
	goto L97
L95:
	;
	if v499 < int32(0) {
		goto L43
	} else {
		goto L120
	}
L96:
	;
	v499 = v471
	goto L95
L97:
	;
	if v395 <= v404 {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v499 = int32(-1)
	goto L95
L100:
	;
	goto L101
L101:
	;
	v411 = int32(1)
	v413 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v404+v396))))
	if base.Ui32(v413) < base.Ui32(int32(192)) {
		v470 = v413
		v471 = v411
		goto L102
	} else {
		goto L103
	}
L102:
	;
	if int32(1103) < v470 {
		goto L115
	} else {
		goto L116
	}
L103:
	;
	v417 = v404 + int32(1)
	if v417 == v395 {
		v470 = v413
		v471 = v411
		goto L102
	} else {
		goto L104
	}
L104:
	;
	v420 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v417+v396))))
	v422 = v420 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v413) {
		goto L106
	} else {
		goto L107
	}
L105:
	;
	v436 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v426+v396))))
	v438 = v436 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v413) {
		goto L111
	} else {
		goto L112
	}
L106:
	;
	v426 = v404 + int32(2)
	if v426 != v395 {
		goto L105
	} else {
		goto L109
	}
L107:
	;
	goto L108
L108:
	;
	v470 = v413<<(uint(int32(6))%32)&int32(1984) | v422
	v471 = int32(2)
	goto L102
L109:
	;
	goto L108
L110:
	;
	v455 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v396+v442))))
	v470 = v455&int32(63) | (v413<<(uint(int32(18))%32)&int32(_a_F_russian_UTF_8_stem_2) | v422<<(uint(int32(12))%32) | v438<<(uint(int32(6))%32))
	v471 = int32(4)
	goto L102
L111:
	;
	v442 = v404 + int32(3)
	if v442 != v395 {
		goto L110
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	v470 = v413<<(uint(int32(12))%32)&int32(_a_F_russian_UTF_8_stem_3) | v422<<(uint(int32(6))%32) | v438
	v471 = int32(3)
	goto L102
L114:
	;
	goto L113
L115:
	;
	v488 = v471 + v404
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v488
	v404 = v488
	goto L97
L116:
	;
	v475 = v470 - int32(1072)
	if v475 < int32(0) {
		goto L115
	} else {
		goto L117
	}
L117:
	;
	v481 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v475)>>(uint(int32(3))%32)))+uint32(_c_F_russian_UTF_8_stem[0]))))
	if int32(base.Ui32(v481)>>(uint(v475&int32(7))%32))&int32(1) != 0 {
		goto L96
	} else {
		goto L118
	}
L118:
	;
	goto L115
L120:
	;
	v502 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v503 = v502 + v499
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v503
	v517 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v518 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v526 = v503
	goto L123
L121:
	;
	if v622 < int32(0) {
		goto L43
	} else {
		goto L145
	}
L122:
	;
	v622 = v593
	goto L121
L123:
	;
	if v517 <= v526 {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v622 = int32(-1)
	goto L121
L126:
	;
	goto L127
L127:
	;
	v533 = int32(1)
	v535 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v526+v518))))
	if base.Ui32(v535) < base.Ui32(int32(192)) {
		v592 = v535
		v593 = v533
		goto L128
	} else {
		goto L129
	}
L128:
	;
	if int32(1103) < v592 {
		goto L122
	} else {
		goto L141
	}
L129:
	;
	v539 = v526 + int32(1)
	if v539 == v517 {
		v592 = v535
		v593 = v533
		goto L128
	} else {
		goto L130
	}
L130:
	;
	v542 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v539+v518))))
	v544 = v542 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v535) {
		goto L132
	} else {
		goto L133
	}
L131:
	;
	v558 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v548+v518))))
	v560 = v558 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v535) {
		goto L137
	} else {
		goto L138
	}
L132:
	;
	v548 = v526 + int32(2)
	if v548 != v517 {
		goto L131
	} else {
		goto L135
	}
L133:
	;
	goto L134
L134:
	;
	v592 = v535<<(uint(int32(6))%32)&int32(1984) | v544
	v593 = int32(2)
	goto L128
L135:
	;
	goto L134
L136:
	;
	v577 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v518+v564))))
	v592 = v577&int32(63) | (v535<<(uint(int32(18))%32)&int32(_a_F_russian_UTF_8_stem_2) | v544<<(uint(int32(12))%32) | v560<<(uint(int32(6))%32))
	v593 = int32(4)
	goto L128
L137:
	;
	v564 = v526 + int32(3)
	if v564 != v517 {
		goto L136
	} else {
		goto L140
	}
L138:
	;
	goto L139
L139:
	;
	v592 = v535<<(uint(int32(12))%32)&int32(_a_F_russian_UTF_8_stem_3) | v544<<(uint(int32(6))%32) | v560
	v593 = int32(3)
	goto L128
L140:
	;
	goto L139
L141:
	;
	v597 = v592 - int32(1072)
	if v597 < int32(0) {
		goto L122
	} else {
		goto L142
	}
L142:
	;
	v603 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v597)>>(uint(int32(3))%32)))+uint32(_c_F_russian_UTF_8_stem[0]))))
	if int32(base.Ui32(v603)>>(uint(v597&int32(7))%32))&int32(1) == int32(0) {
		goto L122
	} else {
		goto L143
	}
L143:
	;
	v611 = v593 + v526
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v611
	v526 = v611
	goto L123
L145:
	;
	v625 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v625 + v622
	goto L43
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v630
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v632
	v639 = F_find_among_b(m, l0, int32(_a_F_russian_UTF_8_stem_4), int32(9), int32(0))
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		goto L40
	} else {
		goto L149
	}
L147:
	;
	v910 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v910
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v910
	v913 = int32(2)
	v915 = int32(0)
	v918 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v910-v918 < v913 {
		v928 = v915
		goto L239
	} else {
		goto L240
	}
L148:
	;
	v696 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v696
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v696
	v699 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v696-int32(3) <= v699 {
		v716 = v696
		goto L169
	} else {
		goto L170
	}
L149:
	;
	if v639 == int32(0) {
		goto L148
	} else {
		goto L150
	}
L150:
	;
	v643 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v643
	switch v639 - int32(1) {
	case 0:
		goto L152
	case 1:
		goto L151
	default:
		goto L147
	}
L151:
	;
	v691 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v691 {
		goto L147
	} else {
		goto L166
	}
L152:
	;
	v647 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v648 = int32(2)
	v650 = int32(0)
	v652 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v653 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v652-v653 < v648 {
		v663 = v650
		goto L154
	} else {
		goto L155
	}
L153:
	;
	if v663 == int32(0) {
		goto L157
	} else {
		goto L158
	}
L154:
	;
	goto L153
L155:
	;
	v656 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v659 = F_memcmp(m, v656+v652-v648, int32(_a_F_russian_UTF_8_stem_5), v648)
	mBase = m.M
	if v659 != 0 {
		v663 = v650
		goto L154
	} else {
		goto L156
	}
L156:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v652 - v648
	v663 = int32(1)
	goto L154
L157:
	;
	v666 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v668 = v666 + (v643 - v647)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v668
	v670 = int32(2)
	v672 = int32(0)
	v675 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v668-v675 < v670 {
		v685 = v672
		goto L161
	} else {
		goto L162
	}
L158:
	;
	goto L159
L159:
	;
	v688 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v688 {
		goto L147
	} else {
		goto L165
	}
L160:
	;
	if v685 == int32(0) {
		goto L148
	} else {
		goto L164
	}
L161:
	;
	goto L160
L162:
	;
	v678 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v681 = F_memcmp(m, v678+v668-v670, int32(_a_F_russian_UTF_8_stem_6), v670)
	mBase = m.M
	if v681 != 0 {
		v685 = v672
		goto L161
	} else {
		goto L163
	}
L163:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v668 - v670
	v685 = int32(1)
	goto L161
L164:
	;
	goto L159
L165:
	;
	v1062 = v688
	goto L1
L166:
	;
	v1062 = v691
	goto L1
L167:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v725
	v731 = F_find_among_b(m, l0, int32(_a_F_russian_UTF_8_stem_7), int32(26), int32(0))
	mBase = m.M
	v732 = m.ExcPending
	if v732 != 0 {
		goto L40
	} else {
		goto L175
	}
L168:
	;
	v718 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v718
	v720 = F_slice_del(m, l0)
	mBase = m.M
	if v720 < int32(0) {
		v1062 = v720
		goto L1
	} else {
		goto L174
	}
L169:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v716
	v725 = v716
	v726 = v716
	goto L167
L170:
	;
	v703 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v707 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v703+v696-int32(1)))))
	switch v707 - int32(140) {
	case 0, 3:
		goto L171
	default:
		v716 = v696
		goto L169
	}
L171:
	;
	v713 = F_find_among_b(m, l0, int32(_a_F_russian_UTF_8_stem_8), int32(2), int32(0))
	mBase = m.M
	v714 = m.ExcPending
	if v714 != 0 {
		goto L40
	} else {
		goto L172
	}
L172:
	;
	if v713 != 0 {
		goto L168
	} else {
		goto L173
	}
L173:
	;
	v715 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v716 = v715
	goto L169
L174:
	;
	v723 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v724 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v725 = v723
	v726 = v724
	goto L167
L175:
	;
	if v731 != 0 {
		goto L176
	} else {
		goto L177
	}
L176:
	;
	v733 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v733
	v735 = F_slice_del(m, l0)
	mBase = m.M
	if v735 < int32(0) {
		v1062 = v735
		goto L1
	} else {
		goto L179
	}
L177:
	;
	goto L178
L178:
	;
	v798 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v799 = v726 - v725
	v800 = v798 - v799
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v800
	v802 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v800
	v808 = F_find_among_b(m, l0, int32(_a_F_russian_UTF_8_stem_9), int32(46), v802)
	mBase = m.M
	v809 = m.ExcPending
	if v809 != 0 {
		goto L40
	} else {
		goto L199
	}
L179:
	;
	v738 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v738
	v743 = F_find_among_b(m, l0, int32(_a_F_russian_UTF_8_stem_10), int32(8), int32(0))
	mBase = m.M
	v744 = m.ExcPending
	if v744 != 0 {
		goto L40
	} else {
		goto L180
	}
L180:
	;
	if v743 == int32(0) {
		goto L147
	} else {
		goto L181
	}
L181:
	;
	v747 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v747
	switch v743 - int32(1) {
	case 0:
		goto L183
	case 1:
		goto L182
	default:
		goto L147
	}
L182:
	;
	v795 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v795 {
		goto L147
	} else {
		goto L197
	}
L183:
	;
	v751 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v752 = int32(2)
	v754 = int32(0)
	v756 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v757 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v756-v757 < v752 {
		v767 = v754
		goto L185
	} else {
		goto L186
	}
L184:
	;
	if v767 == int32(0) {
		goto L188
	} else {
		goto L189
	}
L185:
	;
	goto L184
L186:
	;
	v760 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v763 = F_memcmp(m, v760+v756-v752, int32(_a_F_russian_UTF_8_stem_11), v752)
	mBase = m.M
	if v763 != 0 {
		v767 = v754
		goto L185
	} else {
		goto L187
	}
L187:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v756 - v752
	v767 = int32(1)
	goto L185
L188:
	;
	v770 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v772 = v770 + (v747 - v751)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v772
	v774 = int32(2)
	v776 = int32(0)
	v779 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v772-v779 < v774 {
		v789 = v776
		goto L192
	} else {
		goto L193
	}
L189:
	;
	goto L190
L190:
	;
	v792 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v792 {
		goto L147
	} else {
		goto L196
	}
L191:
	;
	if v789 == int32(0) {
		goto L147
	} else {
		goto L195
	}
L192:
	;
	goto L191
L193:
	;
	v782 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v785 = F_memcmp(m, v782+v772-v774, int32(_a_F_russian_UTF_8_stem_12), v774)
	mBase = m.M
	if v785 != 0 {
		v789 = v776
		goto L192
	} else {
		goto L194
	}
L194:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v772 - v774
	v789 = int32(1)
	goto L192
L195:
	;
	goto L190
L196:
	;
	v1062 = v792
	goto L1
L197:
	;
	v1062 = v795
	goto L1
L198:
	;
	v870 = int32(base.Ui32(v866) >> (uint(int32(31)) % 32))
	if v866 != 0 {
		goto L218
	} else {
		goto L219
	}
L199:
	;
	if v808 == int32(0) {
		v866 = v802
		goto L198
	} else {
		goto L200
	}
L200:
	;
	v812 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v812
	switch v808 - int32(1) {
	case 0:
		goto L203
	case 1:
		goto L202
	default:
		goto L201
	}
L201:
	;
	v866 = int32(1)
	goto L198
L202:
	;
	v860 = F_slice_del(m, l0)
	mBase = m.M
	if v860 < int32(0) {
		v866 = v860
		goto L198
	} else {
		goto L217
	}
L203:
	;
	v816 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v817 = int32(2)
	v819 = int32(0)
	v821 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v822 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v821-v822 < v817 {
		v832 = v819
		goto L205
	} else {
		goto L206
	}
L204:
	;
	if v832 == int32(0) {
		goto L208
	} else {
		goto L209
	}
L205:
	;
	goto L204
L206:
	;
	v825 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v828 = F_memcmp(m, v825+v821-v817, int32(_a_F_russian_UTF_8_stem_13), v817)
	mBase = m.M
	if v828 != 0 {
		v832 = v819
		goto L205
	} else {
		goto L207
	}
L207:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v821 - v817
	v832 = int32(1)
	goto L205
L208:
	;
	v835 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v837 = v835 + (v812 - v816)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v837
	v839 = int32(2)
	v841 = int32(0)
	v844 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v837-v844 < v839 {
		v854 = v841
		goto L212
	} else {
		goto L213
	}
L209:
	;
	goto L210
L210:
	;
	v857 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v857 {
		goto L201
	} else {
		goto L216
	}
L211:
	;
	if v854 == int32(0) {
		v866 = v802
		goto L198
	} else {
		goto L215
	}
L212:
	;
	goto L211
L213:
	;
	v847 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v850 = F_memcmp(m, v847+v837-v839, int32(_a_F_russian_UTF_8_stem_14), v839)
	mBase = m.M
	if v850 != 0 {
		v854 = v841
		goto L212
	} else {
		goto L214
	}
L214:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v837 - v839
	v854 = int32(1)
	goto L212
L215:
	;
	goto L210
L216:
	;
	v866 = v857
	goto L198
L217:
	;
	goto L201
L218:
	;
	v872 = v870
	goto L220
L219:
	;
	v872 = int32(15)
	goto L220
L220:
	;
	if v872 == int32(0) {
		goto L147
	} else {
		goto L221
	}
L221:
	;
	if v872 == int32(15) {
		goto L222
	} else {
		goto L223
	}
L222:
	;
	v877 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v878 = v877 - v799
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v878
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v878
	v885 = F_find_among_b(m, l0, int32(_a_F_russian_UTF_8_stem_15), int32(36), int32(0))
	mBase = m.M
	v886 = m.ExcPending
	if v886 != 0 {
		goto L40
	} else {
		goto L225
	}
L223:
	;
	v904 = v870
	v905 = v866
	goto L224
L224:
	;
	if v904 != 0 {
		v1062 = v905
		goto L1
	} else {
		goto L236
	}
L225:
	;
	if v885 != 0 {
		goto L226
	} else {
		goto L227
	}
L226:
	;
	v887 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v887
	v890 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v890 {
		goto L229
	} else {
		goto L230
	}
L227:
	;
	v896 = int32(0)
	goto L228
L228:
	;
	if v896 == int32(0) {
		goto L147
	} else {
		goto L232
	}
L229:
	;
	v893 = int32(1)
	goto L231
L230:
	;
	v893 = v890
	goto L231
L231:
	;
	v896 = v893
	goto L228
L232:
	;
	if v896 < int32(0) {
		goto L233
	} else {
		goto L234
	}
L233:
	;
	v901 = v896
	goto L235
L234:
	;
	v901 = v866
	goto L235
L235:
	;
	v904 = int32(base.Ui32(v896) >> (uint(int32(31)) % 32))
	v905 = v901
	goto L224
L236:
	;
	goto L147
L237:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v940
	v944 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v940-int32(5) <= v944 {
		goto L246
	} else {
		goto L247
	}
L238:
	;
	if v928 == int32(0) {
		goto L242
	} else {
		goto L243
	}
L239:
	;
	goto L238
L240:
	;
	v921 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v924 = F_memcmp(m, v921+v910-v913, int32(_a_F_russian_UTF_8_stem_16), v913)
	mBase = m.M
	if v924 != 0 {
		v928 = v915
		goto L239
	} else {
		goto L241
	}
L241:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v910 - v913
	v928 = int32(1)
	goto L239
L242:
	;
	v931 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v931
	v940 = v931
	v942 = v931
	goto L237
L243:
	;
	goto L244
L244:
	;
	v933 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v933
	v935 = F_slice_del(m, l0)
	mBase = m.M
	if v935 < int32(0) {
		v1062 = v935
		goto L1
	} else {
		goto L245
	}
L245:
	;
	v938 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v939 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v940 = v938
	v942 = v939
	goto L237
L246:
	;
	v970 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v972 = v970 + (v940 - v942)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v972
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v972
	v978 = F_find_among_b(m, l0, int32(_a_F_russian_UTF_8_stem_17), int32(4), int32(0))
	mBase = m.M
	v979 = m.ExcPending
	if v979 != 0 {
		goto L40
	} else {
		goto L254
	}
L247:
	;
	v948 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v952 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v948+v940-int32(1)))))
	switch v952 - int32(130) {
	case 0, 10:
		goto L248
	default:
		goto L246
	}
L248:
	;
	v958 = F_find_among_b(m, l0, int32(_a_F_russian_UTF_8_stem_18), int32(2), int32(0))
	mBase = m.M
	v959 = m.ExcPending
	if v959 != 0 {
		goto L40
	} else {
		goto L249
	}
L249:
	;
	if v958 == int32(0) {
		goto L246
	} else {
		goto L250
	}
L250:
	;
	v962 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v962
	v964 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v962 < v964 {
		goto L246
	} else {
		goto L251
	}
L251:
	;
	v966 = F_slice_del(m, l0)
	mBase = m.M
	if v966 < int32(0) {
		v1062 = v966
		goto L1
	} else {
		goto L252
	}
L252:
	;
	goto L246
L253:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v7
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v7
	v1062 = int32(1)
	goto L1
L254:
	;
	if v978 == int32(0) {
		goto L253
	} else {
		goto L255
	}
L255:
	;
	v982 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v982
	switch v978 - int32(1) {
	case 0:
		goto L258
	case 1:
		goto L257
	case 2:
		goto L256
	default:
		goto L253
	}
L256:
	;
	v1053 = F_slice_del(m, l0)
	mBase = m.M
	if v1053 < int32(0) {
		v1062 = v1053
		goto L1
	} else {
		goto L277
	}
L257:
	;
	v1032 = int32(2)
	v1034 = int32(0)
	v1036 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1037 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1036-v1037 < v1032 {
		v1047 = v1034
		goto L272
	} else {
		goto L273
	}
L258:
	;
	v986 = F_slice_del(m, l0)
	mBase = m.M
	if v986 < int32(0) {
		v1062 = v986
		goto L1
	} else {
		goto L259
	}
L259:
	;
	v989 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v989
	v991 = int32(2)
	v993 = int32(0)
	v996 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v989-v996 < v991 {
		v1006 = v993
		goto L261
	} else {
		goto L262
	}
L260:
	;
	if v1006 == int32(0) {
		goto L253
	} else {
		goto L264
	}
L261:
	;
	goto L260
L262:
	;
	v999 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1002 = F_memcmp(m, v999+v989-v991, int32(_a_F_russian_UTF_8_stem_19), v991)
	mBase = m.M
	if v1002 != 0 {
		v1006 = v993
		goto L261
	} else {
		goto L263
	}
L263:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v989 - v991
	v1006 = int32(1)
	goto L261
L264:
	;
	v1009 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1009
	v1011 = int32(2)
	v1013 = int32(0)
	v1016 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1009-v1016 < v1011 {
		v1026 = v1013
		goto L266
	} else {
		goto L267
	}
L265:
	;
	if v1026 == int32(0) {
		goto L253
	} else {
		goto L269
	}
L266:
	;
	goto L265
L267:
	;
	v1019 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1022 = F_memcmp(m, v1019+v1009-v1011, int32(_a_F_russian_UTF_8_stem_20), v1011)
	mBase = m.M
	if v1022 != 0 {
		v1026 = v1013
		goto L266
	} else {
		goto L268
	}
L268:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1009 - v1011
	v1026 = int32(1)
	goto L266
L269:
	;
	v1029 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v1029 {
		goto L253
	} else {
		goto L270
	}
L270:
	;
	v1062 = v1029
	goto L1
L271:
	;
	if v1047 == int32(0) {
		goto L253
	} else {
		goto L275
	}
L272:
	;
	goto L271
L273:
	;
	v1040 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1043 = F_memcmp(m, v1040+v1036-v1032, int32(_a_F_russian_UTF_8_stem_21), v1032)
	mBase = m.M
	if v1043 != 0 {
		v1047 = v1034
		goto L272
	} else {
		goto L274
	}
L274:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1036 - v1032
	v1047 = int32(1)
	goto L272
L275:
	;
	v1050 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v1050 {
		goto L253
	} else {
		goto L276
	}
L276:
	;
	v1062 = v1050
	goto L1
L277:
	;
	goto L253
}
