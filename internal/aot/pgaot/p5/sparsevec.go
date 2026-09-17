package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SparsevecCheckValue(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if int32(1001) <= v7 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			F_errcode(m, int32(261))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(1000)
				F_errmsg(m, int32(_a_F_SparsevecCheckValue_0), v5)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_SparsevecCheckValue_1), int32(1369), int32(_a_F_SparsevecCheckValue_2))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
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
		m.G0 = v5 + int32(16)
		return
	}
}
func F_sparsevec_ge(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
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
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v64 float32
	_ = v64
	var v67 int32
	_ = v67
	var v74 float32
	_ = v74
	var v77 int32
	_ = v77
	var v79 float32
	_ = v79
	var v81 float32
	_ = v81
	var v87 int32
	_ = v87
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v114 float32
	_ = v114
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v130 float32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v155 int32
	_ = v155
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v8 = F_pg_detoast_datum(m, v7)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v23 = int32(16)
	v24 = v8 + v23
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	v26 = int32(2)
	v28 = v24 + v25<<(uint(v26)%32)
	v30 = v3 + v23
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v3)+8))
	v34 = v30 + v31<<(uint(v26)%32)
	if v31 < v25 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	return int32(base.Ui32(v155^int32(-1)) >> (uint(int32(31)) % 32))
L5:
	;
	v36 = v31
	goto L7
L6:
	;
	v36 = v25
	goto L7
L7:
	;
	if int32(0) < v36 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v42 = int32(0)
	goto L11
L9:
	;
	goto L10
L10:
	;
	if v25 <= v31 {
		goto L32
	} else {
		goto L33
	}
L11:
	;
	v55 = v42 << (uint(int32(2)) % 32)
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v30+v55)))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v55+v24)))
	if v57 < v59 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	goto L10
L13:
	;
	v64 = *(*float32)(unsafe.Add(mBase, uint32(v55+v34)))
	if base.F32_lt(v64, float32(0)) != 0 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	if v59 < v57 {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	v67 = int32(-1)
	goto L18
L17:
	;
	v67 = int32(1)
	goto L18
L18:
	;
	v155 = v67
	goto L4
L19:
	;
	v74 = *(*float32)(unsafe.Add(mBase, uint32(v28+v42<<(uint(int32(2))%32))))
	if base.F32_lt(v74, float32(0)) != 0 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L21
L21:
	;
	v79 = *(*float32)(unsafe.Add(mBase, uint32(v55+v34)))
	v81 = *(*float32)(unsafe.Add(mBase, uint32(v55+v28)))
	if base.F32_lt(v79, v81) != 0 {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	v77 = int32(1)
	goto L24
L23:
	;
	v77 = int32(-1)
	goto L24
L24:
	;
	v155 = v77
	goto L4
L25:
	;
	v155 = int32(-1)
	goto L4
L26:
	;
	goto L27
L27:
	;
	if base.F32_gt(v79, v81) != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v155 = int32(1)
	goto L4
L29:
	;
	goto L30
L30:
	;
	v87 = v42 + int32(1)
	if v87 != v36 {
		v42 = v87
		goto L11
	} else {
		goto L31
	}
L31:
	;
	goto L12
L32:
	;
	if v31 <= v25 {
		goto L39
	} else {
		goto L40
	}
L33:
	;
	v106 = v31 << (uint(int32(2)) % 32)
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v24+v106)))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v3)+4))
	if v109 <= v108 {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v114 = *(*float32)(unsafe.Add(mBase, uint32(v106+v28)))
	if base.F32_lt(v114, float32(0)) != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v117 = int32(1)
	goto L37
L36:
	;
	v117 = int32(-1)
	goto L37
L37:
	;
	v155 = v117
	goto L4
L38:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v3)+4))
	if v136 < v134 {
		goto L46
	} else {
		goto L47
	}
L39:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v134 = v120
	goto L38
L40:
	;
	goto L41
L41:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v123 = v36 << (uint(int32(2)) % 32)
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v30+v123)))
	if v121 <= v125 {
		v134 = v121
		goto L38
	} else {
		goto L42
	}
L42:
	;
	v130 = *(*float32)(unsafe.Add(mBase, uint32(v123+v34)))
	if base.F32_lt(v130, float32(0)) != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v133 = int32(-1)
	goto L45
L44:
	;
	v133 = int32(1)
	goto L45
L45:
	;
	v155 = v133
	goto L4
L46:
	;
	v155 = int32(-1)
	goto L4
L47:
	;
	goto L48
L48:
	;
	v155 = base.B2i32(v134 < v136)
	goto L4
}
func F_sparsevec_l2_distance(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 float32
	_ = v2
	var v4 int32
	_ = v4
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
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
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 float32
	_ = v49
	var v50 int32
	_ = v50
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v73 float32
	_ = v73
	var v74 int32
	_ = v74
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 float32
	_ = v93
	var v95 float32
	_ = v95
	var v96 float32
	_ = v96
	var v101 float32
	_ = v101
	var v105 float32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v114 float32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v131 float32
	_ = v131
	var v135 float32
	_ = v135
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v160 float32
	_ = v160
	var v161 int32
	_ = v161
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v184 float32
	_ = v184
	var v187 int32
	_ = v187
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v204 float32
	_ = v204
	var v206 float32
	_ = v206
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v212 float32
	_ = v212
	var v230 int32
	_ = v230
	var v232 float32
	_ = v232
	var v249 int32
	_ = v249
	var v250 float32
	_ = v250
	var v252 float32
	_ = v252
	var v254 float32
	_ = v254
	var v256 float32
	_ = v256
	var v261 float32
	_ = v261
	var v263 int32
	_ = v263
	var v267 float32
	_ = v267
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	v2 = float32(0)
	v4 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(16)
	m.G0 = v20
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v23 = F_pg_detoast_datum(m, v22)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v28 = F_pg_detoast_datum(m, v27)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	if v30 == v31 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	if v35 <= v161 {
		v267 = v160
		goto L33
	} else {
		goto L34
	}
L5:
	;
	v34 = v28 + int32(16)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
	v38 = v34 + v35<<(uint(int32(2))%32)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	if v39 <= int32(0) {
		v160 = v2
		v161 = v4
		goto L4
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L1
	} else {
		goto L29
	}
L8:
	;
	v43 = v23 + int32(16)
	v46 = v43 + v39<<(uint(int32(2))%32)
	v49 = v2
	v50 = v4
	v58 = v4
	goto L9
L9:
	;
	v65 = v58 << (uint(int32(2)) % 32)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v43+v65)))
	if v35 <= v50 {
		v114 = v49
		v115 = v50
		v117 = int32(-1)
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v160 = v135
	v161 = v115
	goto L4
L11:
	;
	if v67 != v117 {
		goto L25
	} else {
		goto L26
	}
L12:
	;
	v71 = v50
	v73 = v49
	v74 = v50
	goto L13
L13:
	;
	v89 = v71 << (uint(int32(2)) % 32)
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v34+v89)))
	if v91 == v67 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v114 = v105
	v115 = v109
	v117 = v91
	goto L11
L15:
	;
	v107 = v71 + int32(1)
	if v67 < v91 {
		goto L20
	} else {
		goto L21
	}
L16:
	;
	v93 = *(*float32)(unsafe.Add(mBase, uint32(v46+v65)))
	v95 = *(*float32)(unsafe.Add(mBase, uint32(v38+v89)))
	v96 = base.F32_sub(v93, v95)
	v105 = base.F32_add(base.F32_mul(v96, v96), v73)
	goto L15
L17:
	;
	goto L18
L18:
	;
	if v67 <= v91 {
		v105 = v73
		goto L15
	} else {
		goto L19
	}
L19:
	;
	v101 = *(*float32)(unsafe.Add(mBase, uint32(v38+v89)))
	v105 = base.F32_add(base.F32_mul(v101, v101), v73)
	goto L15
L20:
	;
	v109 = v74
	goto L22
L21:
	;
	v109 = v107
	goto L22
L22:
	;
	if v67 <= v91 {
		v114 = v105
		v115 = v109
		v117 = v91
		goto L11
	} else {
		goto L23
	}
L23:
	;
	if v107 != v35 {
		v71 = v107
		v73 = v105
		v74 = v109
		goto L13
	} else {
		goto L24
	}
L24:
	;
	goto L14
L25:
	;
	v131 = *(*float32)(unsafe.Add(mBase, uint32(v46+v65)))
	v135 = base.F32_add(base.F32_mul(v131, v131), v114)
	goto L27
L26:
	;
	v135 = v114
	goto L27
L27:
	;
	v137 = v58 + int32(1)
	if v137 != v39 {
		v49 = v135
		v50 = v115
		v58 = v137
		goto L9
	} else {
		goto L28
	}
L28:
	;
	goto L10
L29:
	;
	F_errcode(m, int32(130))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v146
	F_errmsg(m, int32(_a_F_sparsevec_l2_distance_0), v20)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	F_errfinish(m, int32(_a_F_sparsevec_l2_distance_1), int32(50), int32(_a_F_sparsevec_l2_distance_2))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L33:
	;
	v284 = F_Float8GetDatum(m, base.F64_sqrt(base.F64_promote_f32(v267)))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L1
	} else {
		goto L46
	}
L34:
	;
	v178 = (v35 - v161) & int32(3)
	if v178 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	if base.Ui32(int32(-4)) < base.Ui32(v161-v35) {
		v267 = v212
		goto L33
	} else {
		goto L42
	}
L36:
	;
	v210 = v161
	v212 = v160
	goto L35
L37:
	;
	goto L38
L38:
	;
	v182 = v161
	v184 = v160
	v187 = int32(0)
	goto L39
L39:
	;
	v199 = int32(1)
	v200 = v182 + v199
	v204 = *(*float32)(unsafe.Add(mBase, uint32(v38+v182<<(uint(int32(2))%32))))
	v206 = base.F32_add(base.F32_mul(v204, v204), v184)
	v208 = v187 + v199
	if v208 != v178 {
		v182 = v200
		v184 = v206
		v187 = v208
		goto L39
	} else {
		goto L41
	}
L40:
	;
	v210 = v200
	v212 = v206
	goto L35
L41:
	;
	goto L40
L42:
	;
	v230 = v210
	v232 = v212
	goto L43
L43:
	;
	v249 = v38 + v230<<(uint(int32(2))%32)
	v250 = *(*float32)(unsafe.Add(mBase, uint32(v249)+12))
	v252 = *(*float32)(unsafe.Add(mBase, uint32(v249)+8))
	v254 = *(*float32)(unsafe.Add(mBase, uint32(v249)+4))
	v256 = *(*float32)(unsafe.Add(mBase, uint32(v249)))
	v261 = base.F32_add(base.F32_mul(v250, v250), base.F32_add(base.F32_mul(v252, v252), base.F32_add(base.F32_mul(v254, v254), base.F32_add(base.F32_mul(v256, v256), v232))))
	v263 = v230 + int32(4)
	if v263 != v35 {
		v230 = v263
		v232 = v261
		goto L43
	} else {
		goto L45
	}
L44:
	;
	v267 = v261
	goto L33
L45:
	;
	goto L44
L46:
	;
	m.G0 = v20 + int32(16)
	return v284
}
func F_sparsevec_lt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
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
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v64 float32
	_ = v64
	var v67 int32
	_ = v67
	var v74 float32
	_ = v74
	var v77 int32
	_ = v77
	var v79 float32
	_ = v79
	var v81 float32
	_ = v81
	var v87 int32
	_ = v87
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v114 float32
	_ = v114
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v130 float32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v155 int32
	_ = v155
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v8 = F_pg_detoast_datum(m, v7)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v23 = int32(16)
	v24 = v8 + v23
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	v26 = int32(2)
	v28 = v24 + v25<<(uint(v26)%32)
	v30 = v3 + v23
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v3)+8))
	v34 = v30 + v31<<(uint(v26)%32)
	if v31 < v25 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	return int32(base.Ui32(v155) >> (uint(int32(31)) % 32))
L5:
	;
	v36 = v31
	goto L7
L6:
	;
	v36 = v25
	goto L7
L7:
	;
	if int32(0) < v36 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v42 = int32(0)
	goto L11
L9:
	;
	goto L10
L10:
	;
	if v25 <= v31 {
		goto L32
	} else {
		goto L33
	}
L11:
	;
	v55 = v42 << (uint(int32(2)) % 32)
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v30+v55)))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v55+v24)))
	if v57 < v59 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	goto L10
L13:
	;
	v64 = *(*float32)(unsafe.Add(mBase, uint32(v55+v34)))
	if base.F32_lt(v64, float32(0)) != 0 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	if v59 < v57 {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	v67 = int32(-1)
	goto L18
L17:
	;
	v67 = int32(1)
	goto L18
L18:
	;
	v155 = v67
	goto L4
L19:
	;
	v74 = *(*float32)(unsafe.Add(mBase, uint32(v28+v42<<(uint(int32(2))%32))))
	if base.F32_lt(v74, float32(0)) != 0 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L21
L21:
	;
	v79 = *(*float32)(unsafe.Add(mBase, uint32(v55+v34)))
	v81 = *(*float32)(unsafe.Add(mBase, uint32(v55+v28)))
	if base.F32_lt(v79, v81) != 0 {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	v77 = int32(1)
	goto L24
L23:
	;
	v77 = int32(-1)
	goto L24
L24:
	;
	v155 = v77
	goto L4
L25:
	;
	v155 = int32(-1)
	goto L4
L26:
	;
	goto L27
L27:
	;
	if base.F32_gt(v79, v81) != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v155 = int32(1)
	goto L4
L29:
	;
	goto L30
L30:
	;
	v87 = v42 + int32(1)
	if v87 != v36 {
		v42 = v87
		goto L11
	} else {
		goto L31
	}
L31:
	;
	goto L12
L32:
	;
	if v31 <= v25 {
		goto L39
	} else {
		goto L40
	}
L33:
	;
	v106 = v31 << (uint(int32(2)) % 32)
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v24+v106)))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v3)+4))
	if v109 <= v108 {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v114 = *(*float32)(unsafe.Add(mBase, uint32(v106+v28)))
	if base.F32_lt(v114, float32(0)) != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v117 = int32(1)
	goto L37
L36:
	;
	v117 = int32(-1)
	goto L37
L37:
	;
	v155 = v117
	goto L4
L38:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v3)+4))
	if v136 < v134 {
		goto L46
	} else {
		goto L47
	}
L39:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v134 = v120
	goto L38
L40:
	;
	goto L41
L41:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v123 = v36 << (uint(int32(2)) % 32)
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v30+v123)))
	if v121 <= v125 {
		v134 = v121
		goto L38
	} else {
		goto L42
	}
L42:
	;
	v130 = *(*float32)(unsafe.Add(mBase, uint32(v123+v34)))
	if base.F32_lt(v130, float32(0)) != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v133 = int32(-1)
	goto L45
L44:
	;
	v133 = int32(1)
	goto L45
L45:
	;
	v155 = v133
	goto L4
L46:
	;
	v155 = int32(-1)
	goto L4
L47:
	;
	goto L48
L48:
	;
	v155 = base.B2i32(v134 < v136)
	goto L4
}
func F_sparsevec_negative_inner_product(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v16 float32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
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
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v61 float32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v84 float32
	_ = v84
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 float32
	_ = v91
	var v93 float32
	_ = v93
	var v96 float32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v106 float32
	_ = v106
	var v108 int32
	_ = v108
	var v125 float32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	v2 = int32(0)
	v16 = float32(0)
	v17 = m.G0
	v19 = v17 - int32(16)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v22 = F_pg_detoast_datum(m, v21)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v27 = F_pg_detoast_datum(m, v26)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	if v29 == v30 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	if int32(0) < v32 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L29
	}
L7:
	;
	v35 = int32(16)
	v36 = v27 + v35
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	v38 = int32(2)
	v42 = v22 + v35
	v48 = v2
	v52 = v2
	v61 = v16
	goto L10
L8:
	;
	v125 = v16
	goto L9
L9:
	;
	v128 = F_Float8GetDatum(m, base.F64_promote_f32(base.F32_neg(v125)))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L28
	}
L10:
	;
	if v37 < v48 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v125 = v106
	goto L9
L12:
	;
	v63 = v48
	goto L14
L13:
	;
	v63 = v37
	goto L14
L14:
	;
	v65 = v52 << (uint(int32(2)) % 32)
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v65+v42)))
	v69 = v48
	v71 = v48
	v84 = v61
	goto L15
L15:
	;
	if v69 != v63 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v108 = v52 + int32(1)
	if v108 != v32 {
		v48 = v103
		v52 = v108
		v61 = v106
		goto L10
	} else {
		goto L27
	}
L17:
	;
	v87 = v69 << (uint(int32(2)) % 32)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v36+v87)))
	if v89 == v68 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v103 = v71
	v106 = v84
	goto L19
L19:
	;
	goto L16
L20:
	;
	v91 = *(*float32)(unsafe.Add(mBase, uint32(v42+v32<<(uint(v38)%32)+v65)))
	v93 = *(*float32)(unsafe.Add(mBase, uint32(v36+v37<<(uint(v38)%32)+v87)))
	v96 = base.F32_add(base.F32_mul(v91, v93), v84)
	goto L22
L21:
	;
	v96 = v84
	goto L22
L22:
	;
	v98 = v69 + int32(1)
	if v68 < v89 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v100 = v71
	goto L25
L24:
	;
	v100 = v98
	goto L25
L25:
	;
	if v89 < v68 {
		v69 = v98
		v71 = v100
		v84 = v96
		goto L15
	} else {
		goto L26
	}
L26:
	;
	v103 = v100
	v106 = v96
	goto L19
L27:
	;
	goto L11
L28:
	;
	m.G0 = v19 + int32(16)
	return v128
L29:
	;
	F_errcode(m, int32(130))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v141
	F_errmsg(m, int32(_a_F_sparsevec_negative_inner_product_0), v19)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	F_errfinish(m, int32(_a_F_sparsevec_negative_inner_product_1), int32(50), int32(_a_F_sparsevec_negative_inner_product_2))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_sparsevec_recv(m *base.Module, l0 int32) int32 {
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
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
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
	var v30 int32
	_ = v30
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
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v96 int32
	_ = v96
	var v97 float32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 float32
	_ = v102
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v17 = F_pq_getmsgint(m, v15, int32(4))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v22 = F_pq_getmsgint(m, v15, int32(4))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v25 = F_pq_getmsgint(m, v15, int32(4))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	F_CheckDim_2(m, v17)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	F_CheckNnz(m, v22, v17)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	if base.B2i32(v14 != int32(-1))&base.B2i32(v14 != v17) == int32(0) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	m.G0 = v12 + int32(32)
	return v48
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L39
	}
L9:
	;
	if v25 != 0 {
		goto L8
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L35
	}
L12:
	;
	v39 = F_mul_size(m, int32(4), v22)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v41 = F_add_size(m, int32(16), v39)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v44 = F_mul_size(m, int32(4), v22)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v46 = F_add_size(m, v41, v44)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v48 = F_palloc0(m, v46)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+8)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v48)+4)) = v17
	*(*int32)(unsafe.Add(mBase, uint32(v48))) = v46 << (uint(int32(2)) % 32)
	if v22 <= int32(0) {
		goto L7
	} else {
		goto L18
	}
L18:
	;
	v58 = v48 + int32(16)
	v63 = int32(0)
	goto L19
L19:
	;
	v76 = F_pq_getmsgint(m, v15, int32(4))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L21
	}
L20:
	;
	v85 = int32(0)
	goto L24
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58+v63<<(uint(int32(2))%32)))) = v76
	F_CheckIndex(m, v58, v63, v17)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v82 = v63 + int32(1)
	if v82 != v22 {
		v63 = v82
		goto L19
	} else {
		goto L23
	}
L23:
	;
	goto L20
L24:
	;
	v96 = v58 + v22<<(uint(int32(2))%32) + v85<<(uint(int32(2))%32)
	v97 = F_pq_getmsgfloat4(m, v15)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L1
	} else {
		goto L27
	}
L25:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L31
	}
L26:
	;
	goto L25
L27:
	;
	*(*float32)(unsafe.Add(mBase, uint32(v96))) = v97
	F_CheckElement_2(m, v97)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v102 = *(*float32)(unsafe.Add(mBase, uint32(v96)))
	if base.F32_eq(v102, float32(0)) != 0 {
		goto L26
	} else {
		goto L29
	}
L29:
	;
	v106 = v85 + int32(1)
	if v22 != v106 {
		v85 = v106
		goto L24
	} else {
		goto L30
	}
L30:
	;
	goto L7
L31:
	;
	F_errcode(m, int32(130))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	F_errmsg(m, int32(_a_F_sparsevec_recv_0), int32(0))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	F_errfinish(m, int32(_a_F_sparsevec_recv_1), int32(551), int32(_a_F_sparsevec_recv_2))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L35:
	;
	F_errcode(m, int32(130))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v17
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v14
	F_errmsg(m, int32(_a_F_sparsevec_recv_3), v12+int32(16))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	F_errfinish(m, int32(_a_F_sparsevec_recv_1), int32(62), int32(_a_F_sparsevec_recv_4))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L39:
	;
	F_errcode(m, int32(130))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v25
	F_errmsg(m, int32(_a_F_sparsevec_recv_5), v12)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(_a_F_sparsevec_recv_1), int32(531), int32(_a_F_sparsevec_recv_2))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
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
func F_sparsevec_send(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
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
	var v29 int32
	_ = v29
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v136 float32
	_ = v136
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = F_pg_detoast_datum(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	F_pq_begintypsend(m, v10)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	F_enlargeStringInfo(m, v10, int32(4))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v29 = int32(16711935)
	*(*int32)(unsafe.Add(mBase, uint32(v24+v25))) = base.I32_rotr(v20, int32(24))&v29 | base.I32_rotr(v20&v29, int32(8))
	v37 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v24 + v37
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	F_enlargeStringInfo(m, v10, v37)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v49 = int32(16711935)
	*(*int32)(unsafe.Add(mBase, uint32(v44+v45))) = base.I32_rotr(v40, int32(24))&v49 | base.I32_rotr(v40&v49, int32(8))
	v57 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v44 + v57
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	F_enlargeStringInfo(m, v10, v57)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v69 = int32(16711935)
	*(*int32)(unsafe.Add(mBase, uint32(v64+v65))) = base.I32_rotr(v60, int32(24))&v69 | base.I32_rotr(v60&v69, int32(8))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v64 + int32(4)
	v80 = int32(0)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	if v81 <= v80 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v151))) = v152 << (uint(int32(2)) % 32)
	goto L18
L8:
	;
	v85 = v13 + int32(16)
	v89 = v80
	goto L9
L9:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v85+v89<<(uint(int32(2))%32))))
	F_enlargeStringInfo(m, v10, int32(4))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L11
	}
L10:
	;
	v123 = int32(0)
	if v121 <= v123 {
		goto L7
	} else {
		goto L13
	}
L11:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v108 = int32(16711935)
	*(*int32)(unsafe.Add(mBase, uint32(v103+v104))) = base.I32_rotr(v99, int32(24))&v108 | base.I32_rotr(v99&v108, int32(8))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v103 + int32(4)
	v120 = v89 + int32(1)
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	if v120 < v121 {
		v89 = v120
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	v126 = v123
	goto L14
L14:
	;
	v136 = *(*float32)(unsafe.Add(mBase, uint32(v85+v17<<(uint(int32(2))%32)+v126<<(uint(int32(2))%32))))
	F_pq_sendfloat4(m, v10, v136)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L16
	}
L15:
	;
	goto L7
L16:
	;
	v140 = v126 + int32(1)
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	if v140 < v141 {
		v126 = v140
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	m.G0 = v10 + int32(16)
	return v151
}
