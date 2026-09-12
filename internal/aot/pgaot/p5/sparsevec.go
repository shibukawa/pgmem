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
				F_errmsg(m, int32(28626), v5)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return
				} else {
					F_errfinish(m, int32(505052), int32(1369), int32(355374))
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
	var v41 int32
	_ = v41
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v66 float32
	_ = v66
	var v69 int32
	_ = v69
	var v76 float32
	_ = v76
	var v79 int32
	_ = v79
	var v81 float32
	_ = v81
	var v83 float32
	_ = v83
	var v89 int32
	_ = v89
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v116 float32
	_ = v116
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v132 float32
	_ = v132
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v157 int32
	_ = v157
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
	return int32(base.Ui32(v157^int32(-1)) >> (uint(int32(31)) % 32))
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
	v41 = int32(0)
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
	v55 = v41 << (uint(int32(2)) % 32)
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
	v66 = *(*float32)(unsafe.Add(mBase, uint32(v34+v41<<(uint(int32(2))%32))))
	if base.F32_lt(v66, float32(0)) != 0 {
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
	v69 = int32(-1)
	goto L18
L17:
	;
	v69 = int32(1)
	goto L18
L18:
	;
	v157 = v69
	goto L4
L19:
	;
	v76 = *(*float32)(unsafe.Add(mBase, uint32(v28+v41<<(uint(int32(2))%32))))
	if base.F32_lt(v76, float32(0)) != 0 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L21
L21:
	;
	v81 = *(*float32)(unsafe.Add(mBase, uint32(v55+v34)))
	v83 = *(*float32)(unsafe.Add(mBase, uint32(v55+v28)))
	if base.F32_lt(v81, v83) != 0 {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	v79 = int32(1)
	goto L24
L23:
	;
	v79 = int32(-1)
	goto L24
L24:
	;
	v157 = v79
	goto L4
L25:
	;
	v157 = int32(-1)
	goto L4
L26:
	;
	goto L27
L27:
	;
	if base.F32_gt(v81, v83) != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v157 = int32(1)
	goto L4
L29:
	;
	goto L30
L30:
	;
	v89 = v41 + int32(1)
	if v89 != v36 {
		v41 = v89
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
	v108 = v31 << (uint(int32(2)) % 32)
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v24+v108)))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v3)+4))
	if v111 <= v110 {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v116 = *(*float32)(unsafe.Add(mBase, uint32(v108+v28)))
	if base.F32_lt(v116, float32(0)) != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v119 = int32(1)
	goto L37
L36:
	;
	v119 = int32(-1)
	goto L37
L37:
	;
	v157 = v119
	goto L4
L38:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v3)+4))
	if v138 < v137 {
		goto L46
	} else {
		goto L47
	}
L39:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v137 = v122
	goto L38
L40:
	;
	goto L41
L41:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v125 = v36 << (uint(int32(2)) % 32)
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v30+v125)))
	if v123 <= v127 {
		v137 = v123
		goto L38
	} else {
		goto L42
	}
L42:
	;
	v132 = *(*float32)(unsafe.Add(mBase, uint32(v125+v34)))
	if base.F32_lt(v132, float32(0)) != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v135 = int32(-1)
	goto L45
L44:
	;
	v135 = int32(1)
	goto L45
L45:
	;
	v157 = v135
	goto L4
L46:
	;
	v157 = int32(-1)
	goto L4
L47:
	;
	goto L48
L48:
	;
	v157 = base.B2i32(v137 < v138)
	goto L4
}
func F_sparsevec_l2_distance(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v16 float32
	_ = v16
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
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v63 float32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v87 float32
	_ = v87
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
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v128 float32
	_ = v128
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
	var v160 int32
	_ = v160
	var v174 float32
	_ = v174
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v198 float32
	_ = v198
	var v202 float32
	_ = v202
	var v204 float32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v226 float32
	_ = v226
	var v236 int32
	_ = v236
	var v252 float32
	_ = v252
	var v254 int32
	_ = v254
	var v256 float32
	_ = v256
	var v259 float32
	_ = v259
	var v262 float32
	_ = v262
	var v265 float32
	_ = v265
	var v270 float32
	_ = v270
	var v272 int32
	_ = v272
	var v290 float32
	_ = v290
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	v2 = int32(0)
	v16 = float32(0)
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
	if v35 <= v160 {
		v290 = v174
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
		v174 = v16
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
	v55 = v2
	v63 = v16
	goto L9
L9:
	;
	v65 = v55 << (uint(int32(2)) % 32)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v43+v65)))
	if v35 <= v49 {
		v114 = v49
		v115 = int32(-1)
		v128 = v63
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v160 = v114
	v174 = v135
	goto L4
L11:
	;
	if v67 != v115 {
		goto L25
	} else {
		goto L26
	}
L12:
	;
	v71 = v49
	v73 = v49
	v87 = v63
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
	v114 = v109
	v115 = v91
	v128 = v105
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
	v105 = base.F32_add(base.F32_mul(v96, v96), v87)
	goto L15
L17:
	;
	goto L18
L18:
	;
	if v67 <= v91 {
		v105 = v87
		goto L15
	} else {
		goto L19
	}
L19:
	;
	v101 = *(*float32)(unsafe.Add(mBase, uint32(v38+v89)))
	v105 = base.F32_add(base.F32_mul(v101, v101), v87)
	goto L15
L20:
	;
	v109 = v73
	goto L22
L21:
	;
	v109 = v107
	goto L22
L22:
	;
	if v67 <= v91 {
		v114 = v109
		v115 = v91
		v128 = v105
		goto L11
	} else {
		goto L23
	}
L23:
	;
	if v107 < v35 {
		v71 = v107
		v73 = v109
		v87 = v105
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
	v135 = base.F32_add(base.F32_mul(v131, v131), v128)
	goto L27
L26:
	;
	v135 = v128
	goto L27
L27:
	;
	v137 = v55 + int32(1)
	if v137 != v39 {
		v49 = v114
		v55 = v137
		v63 = v135
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
	F_errmsg(m, int32(488289), v20)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	F_errfinish(m, int32(512592), int32(50), int32(154322))
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
	v293 = F_Float8GetDatum(m, base.F64_sqrt(base.F64_promote_f32(v290)))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L1
	} else {
		goto L46
	}
L34:
	;
	v178 = (v35 - v160) & int32(3)
	if v178 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	if base.Ui32(int32(-4)) < base.Ui32(v160-v35) {
		v290 = v226
		goto L33
	} else {
		goto L42
	}
L36:
	;
	v210 = v160
	v226 = v174
	goto L35
L37:
	;
	goto L38
L38:
	;
	v183 = v160
	v185 = int32(0)
	v198 = v174
	goto L39
L39:
	;
	v202 = *(*float32)(unsafe.Add(mBase, uint32(v38+v183<<(uint(int32(2))%32))))
	v204 = base.F32_add(base.F32_mul(v202, v202), v198)
	v205 = int32(1)
	v206 = v183 + v205
	v208 = v185 + v205
	if v208 != v178 {
		v183 = v206
		v185 = v208
		v198 = v204
		goto L39
	} else {
		goto L41
	}
L40:
	;
	v210 = v206
	v226 = v204
	goto L35
L41:
	;
	goto L40
L42:
	;
	v236 = v210
	v252 = v226
	goto L43
L43:
	;
	v254 = v236 << (uint(int32(2)) % 32)
	v256 = *(*float32)(unsafe.Add(mBase, uint32(v38+int32(12)+v254)))
	v259 = *(*float32)(unsafe.Add(mBase, uint32(v254+(v38+int32(8)))))
	v262 = *(*float32)(unsafe.Add(mBase, uint32(v254+(v38+int32(4)))))
	v265 = *(*float32)(unsafe.Add(mBase, uint32(v254+v38)))
	v270 = base.F32_add(base.F32_mul(v256, v256), base.F32_add(base.F32_mul(v259, v259), base.F32_add(base.F32_mul(v262, v262), base.F32_add(base.F32_mul(v265, v265), v252))))
	v272 = v236 + int32(4)
	if v272 != v35 {
		v236 = v272
		v252 = v270
		goto L43
	} else {
		goto L45
	}
L44:
	;
	v290 = v270
	goto L33
L45:
	;
	goto L44
L46:
	;
	m.G0 = v20 + int32(16)
	return v293
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
	var v41 int32
	_ = v41
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v66 float32
	_ = v66
	var v69 int32
	_ = v69
	var v76 float32
	_ = v76
	var v79 int32
	_ = v79
	var v81 float32
	_ = v81
	var v83 float32
	_ = v83
	var v89 int32
	_ = v89
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v116 float32
	_ = v116
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v132 float32
	_ = v132
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v157 int32
	_ = v157
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
	return int32(base.Ui32(v157) >> (uint(int32(31)) % 32))
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
	v41 = int32(0)
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
	v55 = v41 << (uint(int32(2)) % 32)
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
	v66 = *(*float32)(unsafe.Add(mBase, uint32(v34+v41<<(uint(int32(2))%32))))
	if base.F32_lt(v66, float32(0)) != 0 {
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
	v69 = int32(-1)
	goto L18
L17:
	;
	v69 = int32(1)
	goto L18
L18:
	;
	v157 = v69
	goto L4
L19:
	;
	v76 = *(*float32)(unsafe.Add(mBase, uint32(v28+v41<<(uint(int32(2))%32))))
	if base.F32_lt(v76, float32(0)) != 0 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L21
L21:
	;
	v81 = *(*float32)(unsafe.Add(mBase, uint32(v55+v34)))
	v83 = *(*float32)(unsafe.Add(mBase, uint32(v55+v28)))
	if base.F32_lt(v81, v83) != 0 {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	v79 = int32(1)
	goto L24
L23:
	;
	v79 = int32(-1)
	goto L24
L24:
	;
	v157 = v79
	goto L4
L25:
	;
	v157 = int32(-1)
	goto L4
L26:
	;
	goto L27
L27:
	;
	if base.F32_gt(v81, v83) != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v157 = int32(1)
	goto L4
L29:
	;
	goto L30
L30:
	;
	v89 = v41 + int32(1)
	if v89 != v36 {
		v41 = v89
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
	v108 = v31 << (uint(int32(2)) % 32)
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v24+v108)))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v3)+4))
	if v111 <= v110 {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v116 = *(*float32)(unsafe.Add(mBase, uint32(v108+v28)))
	if base.F32_lt(v116, float32(0)) != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v119 = int32(1)
	goto L37
L36:
	;
	v119 = int32(-1)
	goto L37
L37:
	;
	v157 = v119
	goto L4
L38:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v3)+4))
	if v138 < v137 {
		goto L46
	} else {
		goto L47
	}
L39:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v137 = v122
	goto L38
L40:
	;
	goto L41
L41:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v125 = v36 << (uint(int32(2)) % 32)
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v30+v125)))
	if v123 <= v127 {
		v137 = v123
		goto L38
	} else {
		goto L42
	}
L42:
	;
	v132 = *(*float32)(unsafe.Add(mBase, uint32(v125+v34)))
	if base.F32_lt(v132, float32(0)) != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v135 = int32(-1)
	goto L45
L44:
	;
	v135 = int32(1)
	goto L45
L45:
	;
	v157 = v135
	goto L4
L46:
	;
	v157 = int32(-1)
	goto L4
L47:
	;
	goto L48
L48:
	;
	v157 = base.B2i32(v137 < v138)
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
	var v50 int32
	_ = v50
	var v61 float32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
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
	v50 = v2
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
	v65 = v50 << (uint(int32(2)) % 32)
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v65+v42)))
	v70 = v48
	v71 = v48
	v84 = v61
	goto L15
L15:
	;
	if v70 != v63 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v108 = v50 + int32(1)
	if v108 != v32 {
		v48 = v103
		v50 = v108
		v61 = v106
		goto L10
	} else {
		goto L27
	}
L17:
	;
	v87 = v70 << (uint(int32(2)) % 32)
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
	v98 = v70 + int32(1)
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
		v70 = v98
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
	F_errmsg(m, int32(488289), v19)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	F_errfinish(m, int32(512592), int32(50), int32(154322))
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
	F_errmsg(m, int32(161546), int32(0))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	F_errfinish(m, int32(512592), int32(551), int32(37713))
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
	F_errmsg(m, int32(477735), v12+int32(16))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	F_errfinish(m, int32(512592), int32(62), int32(294762))
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
	F_errmsg(m, int32(477766), v12)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(512592), int32(531), int32(37713))
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
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
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
	var v57 int32
	_ = v57
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v160 float32
	_ = v160
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
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
	v27 = int32(24)
	v29 = int32(65280)
	v31 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v24+v25))) = v20<<(uint(v27)%32) | v20&v29<<(uint(v31)%32) | (int32(base.Ui32(v20)>>(uint(v31)%32))&v29 | int32(base.Ui32(v20)>>(uint(v27)%32)))
	v43 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v24 + v43
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	F_enlargeStringInfo(m, v10, v43)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v53 = int32(24)
	v55 = int32(65280)
	v57 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v50+v51))) = v46<<(uint(v53)%32) | v46&v55<<(uint(v57)%32) | (int32(base.Ui32(v46)>>(uint(v57)%32))&v55 | int32(base.Ui32(v46)>>(uint(v53)%32)))
	v69 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v50 + v69
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	F_enlargeStringInfo(m, v10, v69)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v79 = int32(24)
	v81 = int32(65280)
	v83 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v76+v77))) = v72<<(uint(v79)%32) | v72&v81<<(uint(v83)%32) | (int32(base.Ui32(v72)>>(uint(v83)%32))&v81 | int32(base.Ui32(v72)>>(uint(v79)%32)))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v76 + int32(4)
	v98 = int32(0)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	if v99 <= v98 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v175))) = v176 << (uint(int32(2)) % 32)
	goto L18
L8:
	;
	v103 = v13 + int32(16)
	v109 = v98
	goto L9
L9:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v103+v109<<(uint(int32(2))%32))))
	F_enlargeStringInfo(m, v10, int32(4))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L11
	}
L10:
	;
	v147 = int32(0)
	if v145 <= v147 {
		goto L7
	} else {
		goto L13
	}
L11:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v124 = int32(24)
	v126 = int32(65280)
	v128 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v121+v122))) = v117<<(uint(v124)%32) | v117&v126<<(uint(v128)%32) | (int32(base.Ui32(v117)>>(uint(v128)%32))&v126 | int32(base.Ui32(v117)>>(uint(v124)%32)))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v121 + int32(4)
	v144 = v109 + int32(1)
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	if v144 < v145 {
		v109 = v144
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	v150 = v147
	goto L14
L14:
	;
	v160 = *(*float32)(unsafe.Add(mBase, uint32(v103+v17<<(uint(int32(2))%32)+v150<<(uint(int32(2))%32))))
	F_pq_sendfloat4(m, v10, v160)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L16
	}
L15:
	;
	goto L7
L16:
	;
	v164 = v150 + int32(1)
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	if v164 < v165 {
		v150 = v164
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
	return v175
}
