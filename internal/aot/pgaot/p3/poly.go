package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_poly_circle(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_pg_detoast_datum(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v9 = F_palloc(m, int32(24))
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			F_poly_to_circle(m, v9, v4)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return int32(0)
			} else {
				return v9
			}
		}
	}
}
func F_poly_contain(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 float64
	_ = v23
	var v24 float64
	_ = v24
	var v30 float64
	_ = v30
	var v31 float64
	_ = v31
	var v37 float64
	_ = v37
	var v38 float64
	_ = v38
	var v44 float64
	_ = v44
	var v45 float64
	_ = v45
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int64
	_ = v61
	var v63 int64
	_ = v63
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v83 int64
	_ = v83
	var v86 int32
	_ = v86
	var v87 int64
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int64
	_ = v95
	var v97 int64
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	v2 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = F_pg_detoast_datum(m, v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v21 = F_pg_detoast_datum(m, v20)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v23 = *(*float64)(unsafe.Add(mBase, uint32(v21)+8))
	v24 = *(*float64)(unsafe.Add(mBase, uint32(v16)+8))
	if base.F64_le(v23, base.F64_add(v24, float64(1e-06))) == int32(0) {
		v106 = v2
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v114 != v16 {
		goto L17
	} else {
		goto L18
	}
L5:
	;
	v30 = *(*float64)(unsafe.Add(mBase, uint32(v16)+24))
	v31 = *(*float64)(unsafe.Add(mBase, uint32(v21)+24))
	if base.F64_le(v30, base.F64_add(v31, float64(1e-06))) == int32(0) {
		v106 = v2
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v37 = *(*float64)(unsafe.Add(mBase, uint32(v21)+16))
	v38 = *(*float64)(unsafe.Add(mBase, uint32(v16)+16))
	if base.F64_le(v37, base.F64_add(v38, float64(1e-06))) == int32(0) {
		v106 = v2
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v44 = *(*float64)(unsafe.Add(mBase, uint32(v16)+32))
	v45 = *(*float64)(unsafe.Add(mBase, uint32(v21)+32))
	if base.F64_le(v44, base.F64_add(v45, float64(1e-06))) == int32(0) {
		v106 = v2
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v52 = v13 + int32(8)
	v54 = v21 + int32(40)
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	v60 = v54 + v55<<(uint(int32(4))%32) - int32(16)
	v61 = *(*int64)(unsafe.Add(mBase, uint32(v60)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v52))) = v61
	v63 = *(*int64)(unsafe.Add(mBase, uint32(v60)))
	*(*int64)(unsafe.Add(mBase, uint32(v13))) = v63
	if v55 <= int32(0) {
		v106 = int32(1)
		goto L4
	} else {
		goto L9
	}
L9:
	;
	v69 = v13 + int32(16)
	v76 = v2
	goto L10
L10:
	;
	v82 = v54 + v76<<(uint(int32(4))%32)
	v83 = *(*int64)(unsafe.Add(mBase, uint32(v82)))
	*(*int64)(unsafe.Add(mBase, uint32(v69))) = v83
	v86 = v13 + int32(24)
	v87 = *(*int64)(unsafe.Add(mBase, uint32(v82)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v86))) = v87
	v90 = F_lseg_inside_poly(m, v13, v69, v16, int32(0))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L12
	}
L11:
	;
	v106 = v99
	goto L4
L12:
	;
	if v90 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v106 = int32(0)
	goto L4
L14:
	;
	goto L15
L15:
	;
	v95 = *(*int64)(unsafe.Add(mBase, uint32(v86)))
	*(*int64)(unsafe.Add(mBase, uint32(v52))) = v95
	v97 = *(*int64)(unsafe.Add(mBase, uint32(v69)))
	*(*int64)(unsafe.Add(mBase, uint32(v13))) = v97
	v99 = int32(1)
	v101 = v76 + v99
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	if v101 < v102 {
		v76 = v101
		goto L10
	} else {
		goto L16
	}
L16:
	;
	goto L11
L17:
	;
	F_pfree(m, v16)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v118 != v21 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	goto L19
L21:
	;
	F_pfree(m, v21)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	m.G0 = v13 + int32(32)
	return v106
L24:
	;
	goto L23
}
func F_poly_out(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_pg_detoast_datum(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
		v11 = F_path_encode(m, int32(2), v8, v4+int32(40))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			return v11
		}
	}
}
func F_poly_recv(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v36 int32
	_ = v36
	var v50 int32
	_ = v50
	var v51 float64
	_ = v51
	var v52 int32
	_ = v52
	var v54 float64
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 float64
	_ = v60
	var v61 float64
	_ = v61
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v70 float64
	_ = v70
	var v71 float64
	_ = v71
	var v72 float64
	_ = v72
	var v73 float64
	_ = v73
	var v82 int32
	_ = v82
	var v83 float64
	_ = v83
	var v88 int32
	_ = v88
	var v92 float64
	_ = v92
	var v98 float64
	_ = v98
	var v99 float64
	_ = v99
	var v100 float64
	_ = v100
	var v102 float64
	_ = v102
	var v104 int64
	_ = v104
	var v108 float64
	_ = v108
	var v113 int32
	_ = v113
	var v117 float64
	_ = v117
	var v123 float64
	_ = v123
	var v124 float64
	_ = v124
	var v125 float64
	_ = v125
	var v126 float64
	_ = v126
	var v128 float64
	_ = v128
	var v134 float64
	_ = v134
	var v136 int32
	_ = v136
	var v140 float64
	_ = v140
	var v141 float64
	_ = v141
	var v142 float64
	_ = v142
	var v143 float64
	_ = v143
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = F_pq_getmsgint(m, v13, int32(4))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if base.Ui32(int32(-134217725)) < base.Ui32(v15-int32(134217725)) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v26 = v15<<(uint(int32(4))%32) + int32(40)
	v27 = F_palloc0(m, v26)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L55
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+4)) = v15
	*(*int32)(unsafe.Add(mBase, uint32(v27))) = v26 << (uint(int32(2)) % 32)
	v36 = int32(0)
	goto L7
L7:
	;
	v50 = v27 + int32(40) + v36<<(uint(int32(4))%32)
	v51 = F_pq_getmsgfloat8(m, v13)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	v60 = *(*float64)(unsafe.Add(mBase, uint32(v27)+48))
	v61 = *(*float64)(unsafe.Add(mBase, uint32(v27)+40))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	if v62 < int32(2) {
		goto L13
	} else {
		goto L14
	}
L9:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v50))) = v51
	v54 = F_pq_getmsgfloat8(m, v13)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v50)+8)) = v54
	v58 = v36 + int32(1)
	if v58 != v15 {
		v36 = v58
		goto L7
	} else {
		goto L11
	}
L11:
	;
	goto L8
L12:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v27)+32)) = v143
	*(*float64)(unsafe.Add(mBase, uint32(v27)+8)) = v141
	*(*float64)(unsafe.Add(mBase, uint32(v27)+24)) = v142
	*(*float64)(unsafe.Add(mBase, uint32(v27)+16)) = v140
	return v27
L13:
	;
	v140 = v60
	v141 = v61
	v142 = v61
	v143 = v60
	goto L12
L14:
	;
	goto L15
L15:
	;
	v68 = int32(1)
	v70 = v60
	v71 = v61
	v72 = v61
	v73 = v60
	goto L16
L16:
	;
	v82 = v27 + int32(40) + v68<<(uint(int32(4))%32)
	v83 = *(*float64)(unsafe.Add(mBase, uint32(v82)))
	v88 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v83)&int64(9223372036854775807)))
	if v88 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v140 = v134
	v141 = v125
	v142 = v99
	v143 = v124
	goto L12
L18:
	;
	if base.F64_lt(v83, v72) != 0 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v99 = v72
	goto L20
L20:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v83)&int64(9223372036854775807)) {
		goto L27
	} else {
		goto L28
	}
L21:
	;
	v92 = v83
	goto L23
L22:
	;
	v92 = v72
	goto L23
L23:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v72)&int64(9223372036854775807)) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v98 = v83
	goto L26
L25:
	;
	v98 = v92
	goto L26
L26:
	;
	v99 = v98
	goto L20
L27:
	;
	v100 = v83
	goto L29
L28:
	;
	v100 = v71
	goto L29
L29:
	;
	if base.F64_gt(v83, v71) != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v102 = v83
	goto L32
L31:
	;
	v102 = v100
	goto L32
L32:
	;
	v104 = int64(9223372036854775807)
	v108 = *(*float64)(unsafe.Add(mBase, uint32(v82)+8))
	v113 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v108)&v104))
	if v113 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	if base.F64_lt(v108, v73) != 0 {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	v124 = v73
	goto L35
L35:
	;
	if base.Ui64(base.I64_reinterpret_f64(v71)&v104) < base.Ui64(int64(9218868437227405313)) {
		goto L42
	} else {
		goto L43
	}
L36:
	;
	v117 = v108
	goto L38
L37:
	;
	v117 = v73
	goto L38
L38:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v73)&int64(9223372036854775807)) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v123 = v108
	goto L41
L40:
	;
	v123 = v117
	goto L41
L41:
	;
	v124 = v123
	goto L35
L42:
	;
	v125 = v102
	goto L44
L43:
	;
	v125 = v71
	goto L44
L44:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v108)&v104) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v126 = v108
	goto L47
L46:
	;
	v126 = v70
	goto L47
L47:
	;
	if base.F64_gt(v108, v70) != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v128 = v108
	goto L50
L49:
	;
	v128 = v126
	goto L50
L50:
	;
	if base.Ui64(base.I64_reinterpret_f64(v70)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v134 = v128
	goto L53
L52:
	;
	v134 = v70
	goto L53
L53:
	;
	v136 = v68 + int32(1)
	if v136 != v62 {
		v68 = v136
		v70 = v134
		v71 = v125
		v72 = v99
		v73 = v124
		goto L16
	} else {
		goto L54
	}
L54:
	;
	goto L17
L55:
	;
	F_errcode(m, int32(50462850))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	F_errmsg(m, int32(363774), int32(0))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	F_errfinish(m, int32(516252), int32(3487), int32(37605))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_poly_right(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 float64
	_ = v17
	var v18 float64
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_pg_detoast_datum(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v13 = l0 + int32(28)
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
		v15 = F_pg_detoast_datum(m, v14)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v17 = *(*float64)(unsafe.Add(mBase, uint32(v15)+8))
			v18 = *(*float64)(unsafe.Add(mBase, uint32(v8)+24))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v19 != v8 {
				F_pfree(m, v8)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
					if v23 != v15 {
						F_pfree(m, v15)
						mBase = m.M
						v26 = m.ExcPending
						if v26 != 0 {
							return int32(0)
						} else {
							return base.F64_lt(v17, v18)
						}
					} else {
						return base.F64_lt(v17, v18)
					}
				}
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
				if v23 != v15 {
					F_pfree(m, v15)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						return base.F64_lt(v17, v18)
					}
				} else {
					return base.F64_lt(v17, v18)
				}
			}
		}
	}
}
