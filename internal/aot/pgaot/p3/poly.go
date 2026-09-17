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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 float64
	_ = v21
	var v22 float64
	_ = v22
	var v28 float64
	_ = v28
	var v29 float64
	_ = v29
	var v35 float64
	_ = v35
	var v36 float64
	_ = v36
	var v42 float64
	_ = v42
	var v43 float64
	_ = v43
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int64
	_ = v57
	var v59 int64
	_ = v59
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int64
	_ = v77
	var v79 int64
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int64
	_ = v87
	var v89 int64
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	v2 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = F_pg_detoast_datum(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v19 = F_pg_detoast_datum(m, v18)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L4
	}
L3:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v104 != v14 {
		goto L17
	} else {
		goto L18
	}
L4:
	;
	v21 = *(*float64)(unsafe.Add(mBase, uint32(v19)+8))
	v22 = *(*float64)(unsafe.Add(mBase, uint32(v14)+8))
	if base.F64_le(v21, base.F64_add(v22, float64(1e-06))) == int32(0) {
		v101 = v2
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v28 = *(*float64)(unsafe.Add(mBase, uint32(v14)+24))
	v29 = *(*float64)(unsafe.Add(mBase, uint32(v19)+24))
	if base.F64_le(v28, base.F64_add(v29, float64(1e-06))) == int32(0) {
		v101 = v2
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v35 = *(*float64)(unsafe.Add(mBase, uint32(v19)+16))
	v36 = *(*float64)(unsafe.Add(mBase, uint32(v14)+16))
	if base.F64_le(v35, base.F64_add(v36, float64(1e-06))) == int32(0) {
		v101 = v2
		goto L3
	} else {
		goto L7
	}
L7:
	;
	v42 = *(*float64)(unsafe.Add(mBase, uint32(v14)+32))
	v43 = *(*float64)(unsafe.Add(mBase, uint32(v19)+32))
	if base.F64_le(v42, base.F64_add(v43, float64(1e-06))) == int32(0) {
		v101 = v2
		goto L3
	} else {
		goto L8
	}
L8:
	;
	v50 = v19 + int32(40)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v56 = v50 + v51<<(uint(int32(4))%32) - int32(16)
	v57 = *(*int64)(unsafe.Add(mBase, uint32(v56)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = v57
	v59 = *(*int64)(unsafe.Add(mBase, uint32(v56)))
	*(*int64)(unsafe.Add(mBase, uint32(v11))) = v59
	if v51 <= int32(0) {
		v101 = int32(1)
		goto L3
	} else {
		goto L9
	}
L9:
	;
	v65 = v11 + int32(16)
	v72 = v2
	goto L10
L10:
	;
	v76 = v50 + v72<<(uint(int32(4))%32)
	v77 = *(*int64)(unsafe.Add(mBase, uint32(v76)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v65)+8)) = v77
	v79 = *(*int64)(unsafe.Add(mBase, uint32(v76)))
	*(*int64)(unsafe.Add(mBase, uint32(v65))) = v79
	v82 = F_lseg_inside_poly(m, v11, v65, v14, int32(0))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L12
	}
L11:
	;
	v101 = v91
	goto L3
L12:
	;
	if v82 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v101 = int32(0)
	goto L3
L14:
	;
	goto L15
L15:
	;
	v87 = *(*int64)(unsafe.Add(mBase, uint32(v65)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = v87
	v89 = *(*int64)(unsafe.Add(mBase, uint32(v65)))
	*(*int64)(unsafe.Add(mBase, uint32(v11))) = v89
	v91 = int32(1)
	v93 = v72 + v91
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v93 < v94 {
		v72 = v93
		goto L10
	} else {
		goto L16
	}
L16:
	;
	goto L11
L17:
	;
	F_pfree(m, v14)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v108 != v19 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	goto L19
L21:
	;
	F_pfree(m, v19)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	m.G0 = v11 + int32(32)
	return v101
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
	*(*float64)(unsafe.Add(mBase, uint32(v27)+32)) = v147
	*(*float64)(unsafe.Add(mBase, uint32(v27)+8)) = v145
	*(*float64)(unsafe.Add(mBase, uint32(v27)+24)) = v146
	*(*float64)(unsafe.Add(mBase, uint32(v27)+16)) = v144
	return v27
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
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	F_errmsg(m, int32(_a_F_poly_recv_0), int32(0))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	F_errfinish(m, int32(_a_F_poly_recv_1), int32(3487), int32(_a_F_poly_recv_2))
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
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v12 = F_pg_detoast_datum(m, v11)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			v14 = *(*float64)(unsafe.Add(mBase, uint32(v12)+8))
			v15 = *(*float64)(unsafe.Add(mBase, uint32(v7)+24))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v16 != v7 {
				F_pfree(m, v7)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v20 != v12 {
						F_pfree(m, v12)
						mBase = m.M
						v23 = m.ExcPending
						if v23 != 0 {
							return int32(0)
						} else {
							return base.F64_lt(v14, v15)
						}
					} else {
						return base.F64_lt(v14, v15)
					}
				}
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				if v20 != v12 {
					F_pfree(m, v12)
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int32(0)
					} else {
						return base.F64_lt(v14, v15)
					}
				} else {
					return base.F64_lt(v14, v15)
				}
			}
		}
	}
}
