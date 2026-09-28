package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_poly_circle(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = F_pg_detoast_datum(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		v10 = F_palloc(m, int32(24))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int64(0)
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			F_poly_to_circle(m, v10, v5, v12)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int64(0)
			} else {
				v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				if v15 == int32(0) {
					return base.I64_extend_i32_u(v10)
				} else {
					v18 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
					if v18 != int32(453) {
						return base.I64_extend_i32_u(v10)
					} else {
						v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)))
						if v21 != int32(1) {
							return base.I64_extend_i32_u(v10)
						} else {
							v24 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v24)
							return int64(0)
						}
					}
				}
			}
		}
	}
}
func F_poly_contain(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v9 int64
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 float64
	_ = v22
	var v23 float64
	_ = v23
	var v29 float64
	_ = v29
	var v30 float64
	_ = v30
	var v36 float64
	_ = v36
	var v37 float64
	_ = v37
	var v43 float64
	_ = v43
	var v44 float64
	_ = v44
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int64
	_ = v58
	var v60 int64
	_ = v60
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v79 int64
	_ = v79
	var v81 int64
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int64
	_ = v88
	var v90 int64
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v114 int64
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	v9 = int64(0)
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v15 = F_pg_detoast_datum(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v20 = F_pg_detoast_datum(m, v19)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L4
	}
L3:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v115 != v15 {
		goto L17
	} else {
		goto L18
	}
L4:
	;
	v22 = *(*float64)(unsafe.Add(mBase, uint32(v20)+8))
	v23 = *(*float64)(unsafe.Add(mBase, uint32(v15)+8))
	if base.F64_le(v22, base.F64_add(v23, float64(1e-06))) == int32(0) {
		v114 = v9
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v29 = *(*float64)(unsafe.Add(mBase, uint32(v15)+24))
	v30 = *(*float64)(unsafe.Add(mBase, uint32(v20)+24))
	if base.F64_le(v29, base.F64_add(v30, float64(1e-06))) == int32(0) {
		v114 = v9
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v36 = *(*float64)(unsafe.Add(mBase, uint32(v20)+16))
	v37 = *(*float64)(unsafe.Add(mBase, uint32(v15)+16))
	if base.F64_le(v36, base.F64_add(v37, float64(1e-06))) == int32(0) {
		v114 = v9
		goto L3
	} else {
		goto L7
	}
L7:
	;
	v43 = *(*float64)(unsafe.Add(mBase, uint32(v15)+32))
	v44 = *(*float64)(unsafe.Add(mBase, uint32(v20)+32))
	if base.F64_le(v43, base.F64_add(v44, float64(1e-06))) == int32(0) {
		v114 = v9
		goto L3
	} else {
		goto L8
	}
L8:
	;
	v51 = v20 + int32(40)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	v57 = v51 + v52<<(uint(int32(4))%32) - int32(16)
	v58 = *(*int64)(unsafe.Add(mBase, uint32(v57)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v58
	v60 = *(*int64)(unsafe.Add(mBase, uint32(v57)))
	*(*int64)(unsafe.Add(mBase, uint32(v12))) = v60
	if int32(0) < v52 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v65 = v12 + int32(16)
	v72 = int32(0)
	goto L12
L10:
	;
	goto L11
L11:
	;
	v114 = int64(1)
	goto L3
L12:
	;
	v78 = v51 + v72<<(uint(int32(4))%32)
	v79 = *(*int64)(unsafe.Add(mBase, uint32(v78)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v65)+8)) = v79
	v81 = *(*int64)(unsafe.Add(mBase, uint32(v78)))
	*(*int64)(unsafe.Add(mBase, uint32(v65))) = v81
	v84 = F_lseg_inside_poly(m, v12, v65, v15, int32(0))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L14
	}
L13:
	;
	goto L11
L14:
	;
	if v84 == int32(0) {
		v114 = v9
		goto L3
	} else {
		goto L15
	}
L15:
	;
	v88 = *(*int64)(unsafe.Add(mBase, uint32(v65)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v88
	v90 = *(*int64)(unsafe.Add(mBase, uint32(v65)))
	*(*int64)(unsafe.Add(mBase, uint32(v12))) = v90
	v93 = v72 + int32(1)
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if v93 < v94 {
		v72 = v93
		goto L12
	} else {
		goto L16
	}
L16:
	;
	goto L13
L17:
	;
	F_pfree(m, v15)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v119 != v20 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	goto L19
L21:
	;
	F_pfree(m, v20)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	m.G0 = v12 + int32(32)
	return v114
L24:
	;
	goto L23
}
func F_poly_out(m *base.Module, l0 int32) int64 {
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
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_pg_detoast_datum(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
		v11 = F_path_encode(m, int32(2), v8, v4+int32(40))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int64(0)
		} else {
			return base.I64_extend_i32_u(v11)
		}
	}
}
func F_poly_recv(m *base.Module, l0 int32) int64 {
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
	var v74 float64
	_ = v74
	var v75 float64
	_ = v75
	var v76 float64
	_ = v76
	var v77 float64
	_ = v77
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
	var v105 int32
	_ = v105
	var v109 float64
	_ = v109
	var v115 float64
	_ = v115
	var v116 float64
	_ = v116
	var v117 float64
	_ = v117
	var v119 float64
	_ = v119
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
	var v144 float64
	_ = v144
	var v145 float64
	_ = v145
	var v146 float64
	_ = v146
	var v147 float64
	_ = v147
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
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
	return int64(0)
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
	v159 = m.ExcPending
	if v159 != 0 {
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
	*(*float64)(unsafe.Add(mBase, uint32(v27)+32)) = v147
	*(*float64)(unsafe.Add(mBase, uint32(v27)+8)) = v145
	*(*float64)(unsafe.Add(mBase, uint32(v27)+24)) = v146
	*(*float64)(unsafe.Add(mBase, uint32(v27)+16)) = v144
	return base.I64_extend_i32_u(v27)
L13:
	;
	v144 = v60
	v145 = v61
	v146 = v61
	v147 = v60
	goto L12
L14:
	;
	goto L15
L15:
	;
	v68 = int32(1)
	v74 = v60
	v75 = v61
	v76 = v61
	v77 = v60
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
	v144 = v134
	v145 = v125
	v146 = v99
	v147 = v116
	goto L12
L18:
	;
	if base.F64_gt(v76, v83) != 0 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v99 = v76
	goto L20
L20:
	;
	v100 = *(*float64)(unsafe.Add(mBase, uint32(v82)+8))
	v105 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v100)&int64(9223372036854775807)))
	if v105 == int32(0) {
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
	v92 = v76
	goto L23
L23:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v76)&int64(9223372036854775807)) {
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
	if base.F64_gt(v77, v100) != 0 {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	v116 = v77
	goto L29
L29:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v83)&int64(9223372036854775807)) {
		goto L36
	} else {
		goto L37
	}
L30:
	;
	v109 = v100
	goto L32
L31:
	;
	v109 = v77
	goto L32
L32:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v77)&int64(9223372036854775807)) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v115 = v100
	goto L35
L34:
	;
	v115 = v109
	goto L35
L35:
	;
	v116 = v115
	goto L29
L36:
	;
	v117 = v83
	goto L38
L37:
	;
	v117 = v75
	goto L38
L38:
	;
	if base.F64_lt(v75, v83) != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v119 = v83
	goto L41
L40:
	;
	v119 = v117
	goto L41
L41:
	;
	if base.Ui64(base.I64_reinterpret_f64(v75)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v125 = v119
	goto L44
L43:
	;
	v125 = v75
	goto L44
L44:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v100)&int64(9223372036854775807)) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v126 = v100
	goto L47
L46:
	;
	v126 = v74
	goto L47
L47:
	;
	if base.F64_lt(v74, v100) != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v128 = v100
	goto L50
L49:
	;
	v128 = v126
	goto L50
L50:
	;
	if base.Ui64(base.I64_reinterpret_f64(v74)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
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
	v134 = v74
	goto L53
L53:
	;
	v136 = v68 + int32(1)
	if v136 != v62 {
		v68 = v136
		v74 = v134
		v75 = v125
		v76 = v99
		v77 = v116
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
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	F_errmsg(m, int32(_a_F_poly_recv_0), int32(0))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	F_errfinish(m, int32(_a_F_poly_recv_1), int32(3571), int32(_a_F_poly_recv_2))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
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
func F_poly_right(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 float64
	_ = v14
	var v15 float64
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = F_pg_detoast_datum(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v12 = F_pg_detoast_datum(m, v11)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int64(0)
		} else {
			v14 = *(*float64)(unsafe.Add(mBase, uint32(v12)+8))
			v15 = *(*float64)(unsafe.Add(mBase, uint32(v7)+24))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			if v16 != v7 {
				F_pfree(m, v7)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int64(0)
				} else {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					if v20 != v12 {
						F_pfree(m, v12)
						mBase = m.M
						v23 = m.ExcPending
						if v23 != 0 {
							return int64(0)
						} else {
							return base.I64_extend_i32_u(base.F64_lt(v14, v15))
						}
					} else {
						return base.I64_extend_i32_u(base.F64_lt(v14, v15))
					}
				}
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
				if v20 != v12 {
					F_pfree(m, v12)
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int64(0)
					} else {
						return base.I64_extend_i32_u(base.F64_lt(v14, v15))
					}
				} else {
					return base.I64_extend_i32_u(base.F64_lt(v14, v15))
				}
			}
		}
	}
}
