package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_sparsevec_cmp(m *base.Module, l0 int32) int32 {
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
	return v155
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
func F_sparsevec_cmp_internal(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v57 float32
	_ = v57
	var v60 int32
	_ = v60
	var v68 float32
	_ = v68
	var v71 int32
	_ = v71
	var v74 float32
	_ = v74
	var v76 float32
	_ = v76
	var v84 int32
	_ = v84
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v111 float32
	_ = v111
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v128 float32
	_ = v128
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	v16 = int32(16)
	v17 = l1 + v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v19 = int32(2)
	v21 = v17 + v18<<(uint(v19)%32)
	v23 = l0 + v16
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v27 = v23 + v24<<(uint(v19)%32)
	if v24 < v18 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v29 = v24
	goto L3
L2:
	;
	v29 = v18
	goto L3
L3:
	;
	if int32(0) < v29 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v35 = int32(0)
	goto L7
L5:
	;
	goto L6
L6:
	;
	if v18 <= v24 {
		goto L28
	} else {
		goto L29
	}
L7:
	;
	v48 = v35 << (uint(int32(2)) % 32)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v23+v48)))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v48+v17)))
	if v50 < v52 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L6
L9:
	;
	v57 = *(*float32)(unsafe.Add(mBase, uint32(v48+v27)))
	if base.F32_lt(v57, float32(0)) != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	if v52 < v50 {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	v60 = int32(-1)
	goto L14
L13:
	;
	v60 = int32(1)
	goto L14
L14:
	;
	return v60
L15:
	;
	v68 = *(*float32)(unsafe.Add(mBase, uint32(v21+v35<<(uint(int32(2))%32))))
	if base.F32_lt(v68, float32(0)) != 0 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L17
L17:
	;
	v74 = *(*float32)(unsafe.Add(mBase, uint32(v48+v27)))
	v76 = *(*float32)(unsafe.Add(mBase, uint32(v48+v21)))
	if base.F32_lt(v74, v76) != 0 {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	v71 = int32(1)
	goto L20
L19:
	;
	v71 = int32(-1)
	goto L20
L20:
	;
	return v71
L21:
	;
	return int32(-1)
L22:
	;
	goto L23
L23:
	;
	if base.F32_gt(v74, v76) != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	return int32(1)
L25:
	;
	goto L26
L26:
	;
	v84 = v35 + int32(1)
	if v84 != v29 {
		v35 = v84
		goto L7
	} else {
		goto L27
	}
L27:
	;
	goto L8
L28:
	;
	if v24 <= v18 {
		goto L35
	} else {
		goto L36
	}
L29:
	;
	v103 = v24 << (uint(int32(2)) % 32)
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v17+v103)))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v106 <= v105 {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v111 = *(*float32)(unsafe.Add(mBase, uint32(v103+v21)))
	if base.F32_lt(v111, float32(0)) != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v114 = int32(1)
	goto L33
L32:
	;
	v114 = int32(-1)
	goto L33
L33:
	;
	return v114
L34:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v135 < v133 {
		goto L42
	} else {
		goto L43
	}
L35:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v133 = v118
	goto L34
L36:
	;
	goto L37
L37:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v121 = v29 << (uint(int32(2)) % 32)
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v23+v121)))
	if v119 <= v123 {
		v133 = v119
		goto L34
	} else {
		goto L38
	}
L38:
	;
	v128 = *(*float32)(unsafe.Add(mBase, uint32(v121+v27)))
	if base.F32_lt(v128, float32(0)) != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v131 = int32(-1)
	goto L41
L40:
	;
	v131 = int32(1)
	goto L41
L41:
	;
	return v131
L42:
	;
	return int32(-1)
L43:
	;
	goto L44
L44:
	;
	return base.B2i32(v133 < v135)
}
func F_sparsevec_l2_squared_distance(m *base.Module, l0 int32) int32 {
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
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
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
	F_errmsg(m, int32(_a_F_sparsevec_l2_squared_distance_0), v20)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	F_errfinish(m, int32(_a_F_sparsevec_l2_squared_distance_1), int32(50), int32(_a_F_sparsevec_l2_squared_distance_2))
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
	v283 = F_Float8GetDatum(m, base.F64_promote_f32(v267))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
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
	return v283
}
