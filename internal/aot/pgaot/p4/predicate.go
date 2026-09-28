package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_PredicateLockShmemInit(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int64
	_ = v22
	var v33 int32
	_ = v33
	var v42 int64
	_ = v42
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
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
	var v75 int32
	_ = v75
	var v80 int64
	_ = v80
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v122 int64
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int64
	_ = v133
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
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
	v2 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_PredicateLockShmemInit[0]))
	v18 = F_hash_search(m, v13, int32(_a_F_PredicateLockShmemInit_0), int32(1), v10+int32(15))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_PredicateLockShmemInit[1]))
	v22 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v21)+40)) = v22
	*(*int64)(unsafe.Add(mBase, uint32(v21)+32)) = int64(1)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+24)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v21)+16)) = v22
	*(*int64)(unsafe.Add(mBase, uint32(v21)+48)) = v22
	v33 = v21 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+12)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+60)) = v21 - int32(-64)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = v21
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v21
	v42 = *(*int64)(unsafe.Add(mBase, _c_F_PredicateLockShmemInit[2]))
	if v42 <= v22 {
		v116 = v21
		v118 = v2
		v122 = v42
		goto L3
	} else {
		goto L4
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v116)+56)) = v118
	v124 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v118))) = v124
	v126 = int32(0)
	v128 = *(*int32)(unsafe.Add(mBase, _c_F_PredicateLockShmemInit[1]))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v128)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v129)+4)) = v126
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v128)+56))
	v133 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v132)+8)) = v133
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v128)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v135)+16)) = v133
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v128)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v138)+24)) = v133
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v128)+56))
	v143 = v141 + int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v141)+36)) = v143
	*(*int32)(unsafe.Add(mBase, uint32(v141)+32)) = v143
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v128)+56))
	v148 = v146 + int32(40)
	*(*int32)(unsafe.Add(mBase, uint32(v146)+44)) = v148
	*(*int32)(unsafe.Add(mBase, uint32(v146)+40)) = v148
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v128)+56))
	v153 = v151 + int32(48)
	*(*int32)(unsafe.Add(mBase, uint32(v151)+52)) = v153
	*(*int32)(unsafe.Add(mBase, uint32(v151)+48)) = v153
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v128)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v156)+56)) = v133
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v128)+56))
	v161 = v159 + int32(88)
	*(*int32)(unsafe.Add(mBase, uint32(v159)+92)) = v161
	*(*int32)(unsafe.Add(mBase, uint32(v159)+88)) = v161
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v128)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v164)+96)) = v126
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v128)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v167)+100)) = v126
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v128)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v170)+104)) = v126
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v128)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v173)+108)) = int32(1)
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v128)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v176)+112)) = v126
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v128)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v179)+116)) = v124
	v183 = *(*int32)(unsafe.Add(mBase, _c_F_PredicateLockShmemInit[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v183)+8)) = v183 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v183)+4)) = v183
	*(*int32)(unsafe.Add(mBase, uint32(v183))) = v183
	v191 = base.I32_wrap_i64(v122) * int32(5)
	if v126 < v191 {
		goto L16
	} else {
		goto L17
	}
L4:
	;
	v47 = v2
	goto L5
L5:
	;
	v53 = v47 * int32(120)
	v55 = *(*int32)(unsafe.Add(mBase, _c_F_PredicateLockShmemInit[1]))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+60))
	F_LWLockInitialize(m, v53+v56+int32(72), int32(82))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L7
	}
L6:
	;
	v85 = int32(0)
	v87 = *(*int32)(unsafe.Add(mBase, _c_F_PredicateLockShmemInit[1]))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
	if base.B2i32(v88 == v85)|base.B2i32(v87 == v88) != 0 {
		v116 = v87
		v118 = v85
		v122 = v80
		goto L3
	} else {
		goto L12
	}
L7:
	;
	v64 = *(*int32)(unsafe.Add(mBase, _c_F_PredicateLockShmemInit[1]))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+60))
	v66 = v53 + v65
	v68 = v66 - int32(-64)
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	if v69 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v64)+4)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v64))) = v64
	goto L10
L9:
	;
	goto L10
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v66)+68)) = v64
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	*(*int32)(unsafe.Add(mBase, uint32(v66)+64)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v75)+4)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v64))) = v68
	v80 = *(*int64)(unsafe.Add(mBase, _c_F_PredicateLockShmemInit[2]))
	v82 = v47 + int32(1)
	if base.I64_extend_i32_s(v82) < v80 {
		v47 = v82
		goto L5
	} else {
		goto L11
	}
L11:
	;
	goto L6
L12:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v88)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v93)+4)) = v94
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	*(*int32)(unsafe.Add(mBase, uint32(v94))) = v96
	v99 = *(*int32)(unsafe.Add(mBase, _c_F_PredicateLockShmemInit[1]))
	v101 = v99 + int32(8)
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v99)+12))
	if v102 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v99)+12)) = v101
	*(*int32)(unsafe.Add(mBase, uint32(v99)+8)) = v101
	goto L15
L14:
	;
	goto L15
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+4)) = v101
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	*(*int32)(unsafe.Add(mBase, uint32(v88))) = v110
	*(*int32)(unsafe.Add(mBase, uint32(v110)+4)) = v88
	*(*int32)(unsafe.Add(mBase, uint32(v101))) = v88
	v115 = *(*int32)(unsafe.Add(mBase, _c_F_PredicateLockShmemInit[1]))
	v116 = v115
	v118 = v88 + int32(-64)
	v122 = v80
	goto L3
L16:
	;
	v196 = v126
	goto L19
L17:
	;
	goto L18
L18:
	;
	v228 = *(*int32)(unsafe.Add(mBase, _c_F_PredicateLockShmemInit[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v228)+4)) = v228
	*(*int32)(unsafe.Add(mBase, uint32(v228))) = v228
	v232 = *(*int32)(unsafe.Add(mBase, _c_F_PredicateLockShmemInit[5]))
	v236 = F_LWLockAcquire(m, v232+int32(_a_F_PredicateLockShmemInit_1), int32(0))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L1
	} else {
		goto L25
	}
L19:
	;
	v202 = *(*int32)(unsafe.Add(mBase, _c_F_PredicateLockShmemInit[3]))
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v202)+8))
	v206 = v203 + v196*int32(24)
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v202)+4))
	if v207 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	goto L18
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v202)+4)) = v202
	*(*int32)(unsafe.Add(mBase, uint32(v202))) = v202
	goto L23
L22:
	;
	goto L23
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v206)+4)) = v202
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v202)))
	*(*int32)(unsafe.Add(mBase, uint32(v206))) = v213
	*(*int32)(unsafe.Add(mBase, uint32(v213)+4)) = v206
	*(*int32)(unsafe.Add(mBase, uint32(v202))) = v206
	v218 = v196 + int32(1)
	if v218 != v191 {
		v196 = v218
		goto L19
	} else {
		goto L24
	}
L24:
	;
	goto L20
L25:
	;
	v239 = *(*int32)(unsafe.Add(mBase, _c_F_PredicateLockShmemInit[6]))
	*(*int64)(unsafe.Add(mBase, uint32(v239)+8)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v239))) = int64(-1)
	v245 = *(*int32)(unsafe.Add(mBase, _c_F_PredicateLockShmemInit[5]))
	F_LWLockRelease(m, v245+int32(_a_F_PredicateLockShmemInit_1))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v252 = *(*int32)(unsafe.Add(mBase, _c_F_PredicateLockShmemInit[1]))
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v252)+56))
	*(*int32)(unsafe.Add(mBase, _c_F_PredicateLockShmemInit[7])) = v253
	v257 = *(*int32)(unsafe.Add(mBase, _c_F_PredicateLockShmemInit[0]))
	v259 = F_get_hash_value(m, v257, int32(_a_F_PredicateLockShmemInit_0))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PredicateLockShmemInit[8])) = v259
	v264 = *(*int32)(unsafe.Add(mBase, _c_F_PredicateLockShmemInit[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_PredicateLockShmemInit[9])) = v264 + v259&int32(15)<<(uint(int32(7))%32) + int32(_a_F_PredicateLockShmemInit_2)
	m.G0 = v10 + int32(16)
	return
}
func F_predicate_classify(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v17 int32
	_ = v17
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
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
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	v3 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v6 - int32(1) {
	case 0:
		*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = int32(933)
		*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = int32(934)
		*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(935)
		return int32(1)
	default:
		v85 = v3
		return v85
	case 19:
		v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
		v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
		if v36 == int32(0) {
			v85 = v3
			return v85
		} else {
			v39 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
			if v39 != int32(35) {
				if v39 != int32(7) {
					v85 = v3
					return v85
				} else {
					v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+32)))
					if v44 != 0 {
						v85 = v3
						return v85
					} else {
						v45 = *(*int32)(unsafe.Add(mBase, uint32(v36)+24))
						v46 = F_pg_detoast_datum(m, v45)
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return int32(0)
						} else {
							v50 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
							v53 = F_ArrayGetNItemsSafe(m, v50, v46+int32(16))
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return int32(0)
							} else {
								if int32(100) < v53 {
									v85 = v3
									return v85
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = int32(937)
									*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = int32(938)
									*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(939)
									v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
									if v65 != 0 {
										v66 = int32(2)
									} else {
										v66 = int32(1)
									}
									return v66
								}
							}
						}
					}
				}
			} else {
				v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+20)))
				if v68 != 0 {
					v85 = v3
				} else {
					v69 = *(*int32)(unsafe.Add(mBase, uint32(v36)+16))
					if v69 != 0 {
						v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
						if int32(100) < v70 {
							v85 = v3
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = int32(940)
							*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = int32(941)
							*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(942)
							v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
							if v81 != 0 {
								v82 = int32(2)
							} else {
								v82 = int32(1)
							}
							v85 = v82
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = int32(940)
						*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = int32(941)
						*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(942)
						v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
						if v81 != 0 {
							v82 = int32(2)
						} else {
							v82 = int32(1)
						}
						v85 = v82
					}
				}
				return v85
			}
		}
	case 20:
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		switch v17 {
		case 0:
			*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = int32(933)
			*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = int32(934)
			*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(936)
			return int32(1)
		case 1:
			*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = int32(933)
			*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = int32(934)
			*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(936)
			return int32(2)
		default:
			v85 = v3
			return v85
		}
	}
}
