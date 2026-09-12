package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_VectorItemSize(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	v4 = F_mul_size(m, int32(4), l0)
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = F_add_size(m, int32(8), v4)
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			return v8
		}
	}
}
func F_VectorUpdateCenter(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 float32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 float32
	_ = v50
	var v53 int32
	_ = v53
	var v56 float32
	_ = v56
	var v59 int32
	_ = v59
	var v62 float32
	_ = v62
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v93 float32
	_ = v93
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	v2 = l1
	v4 = int32(0)
	v12 = F_mul_size(m, int32(4), v2)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		v14 = F_add_size(m, int32(8), v12)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return
		} else {
			*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v2)
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = v14 << (uint(int32(2)) % 32)
			if v2 <= int32(0) {
			} else {
				v23 = v2 & int32(3)
				v25 = l0 + int32(8)
				v26 = int32(0)
				if base.Ui32(int32(4)) <= base.Ui32(v2) {
					v31 = v26
					v38 = v4
					for {
						v41 = v31 << (uint(int32(2)) % 32)
						v44 = *(*float32)(unsafe.Add(mBase, uint32(v41+l2)))
						*(*float32)(unsafe.Add(mBase, uint32(v25+v41))) = v44
						v46 = int32(4)
						v47 = v41 | v46
						v50 = *(*float32)(unsafe.Add(mBase, uint32(l2+v47)))
						*(*float32)(unsafe.Add(mBase, uint32(v25+v47))) = v50
						v53 = v41 | int32(8)
						v56 = *(*float32)(unsafe.Add(mBase, uint32(l2+v53)))
						*(*float32)(unsafe.Add(mBase, uint32(v25+v53))) = v56
						v59 = v41 | int32(12)
						v62 = *(*float32)(unsafe.Add(mBase, uint32(v59+l2)))
						*(*float32)(unsafe.Add(mBase, uint32(v25+v59))) = v62
						v65 = v31 + v46
						v67 = v38 + v46
						if v67 != v2&int32(2147483644) {
							v31 = v65
							v38 = v67
							continue
						} else {
							break
						}
						break
					}
					v69 = v65
				} else {
					v69 = v26
				}
				if v23 == int32(0) {
				} else {
					v80 = v69
					v86 = v4
					for {
						v90 = v80 << (uint(int32(2)) % 32)
						v93 = *(*float32)(unsafe.Add(mBase, uint32(v90+l2)))
						*(*float32)(unsafe.Add(mBase, uint32(v25+v90))) = v93
						v95 = int32(1)
						v98 = v86 + v95
						if v98 != v23 {
							v80 = v80 + v95
							v86 = v98
							continue
						} else {
							break
						}
						break
					}
				}
			}
			return
		}
	}
}
func F_vector(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		if v14 == int32(-1) {
			m.G0 = v7 + int32(16)
			return v10
		} else {
			v17 = int32(*(*int16)(unsafe.Add(mBase, uint32(v10)+4)))
			if v14 == v17 {
				m.G0 = v7 + int32(16)
				return v10
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(130))
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v17
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = v14
						F_errmsg(m, int32(488627), v7)
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(517975), int32(88), int32(302025))
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return int32(0)
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
func F_vector_accum(m *base.Module, l0 int32) int32 {
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
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 float64
	_ = v43
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
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v71 int32
	_ = v71
	var v75 float64
	_ = v75
	var v77 float32
	_ = v77
	var v79 float64
	_ = v79
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v110 float32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
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
	return int32(0)
L2:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v20 = F_pg_detoast_datum(m, v19)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v22 != int32(1) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	F_float_overflow_error(m)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L1
	} else {
		goto L43
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L39
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L36
	}
L7:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
	if v25 <= int32(0) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	if v28 != 0 {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	if v29 != int32(701) {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	v32 = int32(*(*int16)(unsafe.Add(mBase, uint32(v20)+4)))
	v34 = v25 - int32(1)
	if v34 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v40 = v20 + int32(8)
	v42 = v15 + int32(24)
	v43 = *(*float64)(unsafe.Add(mBase, uint32(v42)))
	v46 = v38 + int32(1)
	v47 = F_mul_size(m, int32(4), v46)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L16
	}
L12:
	;
	v38 = v32
	goto L11
L13:
	;
	goto L14
L14:
	;
	if v34 != v32 {
		goto L5
	} else {
		goto L15
	}
L15:
	;
	v38 = v34
	goto L11
L16:
	;
	v49 = F_palloc(m, v47)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v53 = F_Float8GetDatum(m, base.F64_add(v43, float64(1)))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49))) = v53
	if v34 != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v129 = F_construct_array(m, v49, v46, int32(701), int32(8), int32(0), int32(100))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L34
	}
L20:
	;
	v56 = int32(0)
	if v38 <= v56 {
		goto L19
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	if v38 <= int32(0) {
		goto L19
	} else {
		goto L29
	}
L23:
	;
	v60 = v56
	goto L24
L24:
	;
	v71 = v60 + int32(1)
	v75 = *(*float64)(unsafe.Add(mBase, uint32(v42+v71<<(uint(int32(3))%32))))
	v77 = *(*float32)(unsafe.Add(mBase, uint32(v60<<(uint(int32(2))%32)+v40)))
	v79 = base.F64_add(v75, base.F64_promote_f32(v77))
	if base.F64_eq(base.F64_abs(v79), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L4
	} else {
		goto L26
	}
L25:
	;
	goto L19
L26:
	;
	v86 = F_Float8GetDatum(m, v79)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49+v71<<(uint(int32(2))%32)))) = v86
	if v71 != v38 {
		v60 = v71
		goto L24
	} else {
		goto L28
	}
L28:
	;
	goto L25
L29:
	;
	v94 = int32(0)
	goto L30
L30:
	;
	v103 = v94 + int32(1)
	v104 = int32(2)
	v110 = *(*float32)(unsafe.Add(mBase, uint32(v40+v94<<(uint(v104)%32))))
	v112 = F_Float8GetDatum(m, base.F64_promote_f32(v110))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L32
	}
L31:
	;
	goto L19
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49+v103<<(uint(v104)%32)))) = v112
	if v38 != v103 {
		v94 = v103
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	F_pfree(m, v49)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	m.G0 = v12 + int32(32)
	return v129
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(300546)
	F_errmsg_internal(m, int32(26326), v12)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	F_errfinish(m, int32(517975), int32(169), int32(26817))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
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
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v34
	F_errmsg(m, int32(488627), v12+int32(16))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(517975), int32(88), int32(302025))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L43:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_vector_combine(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 float64
	_ = v45
	var v47 int32
	_ = v47
	var v48 float64
	_ = v48
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v75 int32
	_ = v75
	var v82 float64
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v113 int32
	_ = v113
	var v120 float64
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 float64
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v156 float64
	_ = v156
	var v158 float64
	_ = v158
	var v159 float64
	_ = v159
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v179 float64
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v189 float64
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	v11 = m.G0
	v13 = v11 - int32(48)
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
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v23 != int32(1) {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	F_float_overflow_error(m)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L1
	} else {
		goto L65
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L1
	} else {
		goto L61
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L1
	} else {
		goto L58
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L1
	} else {
		goto L55
	}
L8:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
	if v26 <= int32(0) {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	if v29 != 0 {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	if v30 != int32(701) {
		goto L7
	} else {
		goto L11
	}
L11:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	if v33 != int32(1) {
		goto L6
	} else {
		goto L12
	}
L12:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
	if v36 <= int32(0) {
		goto L6
	} else {
		goto L13
	}
L13:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	if v39 != 0 {
		goto L6
	} else {
		goto L14
	}
L14:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	if v40 != int32(701) {
		goto L6
	} else {
		goto L15
	}
L15:
	;
	v43 = int32(24)
	v44 = v21 + v43
	v45 = *(*float64)(unsafe.Add(mBase, uint32(v21)+24))
	v47 = v16 + v43
	v48 = *(*float64)(unsafe.Add(mBase, uint32(v47)))
	if base.F64_eq(v48, float64(0)) != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v190 = F_Float8GetDatum(m, v189)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L52
	}
L17:
	;
	v52 = v36 - int32(1)
	F_CheckDim_3(m, v52)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	if base.F64_eq(v45, float64(0)) != 0 {
		goto L31
	} else {
		goto L32
	}
L20:
	;
	v57 = F_mul_size(m, int32(4), v36)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v59 = F_palloc(m, v57)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	if v36 == int32(1) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v180 = int32(1)
	v181 = v59
	v189 = v45
	goto L16
L24:
	;
	goto L25
L25:
	;
	v64 = int32(0)
	goto L26
L26:
	;
	v75 = v64 + int32(1)
	v82 = *(*float64)(unsafe.Add(mBase, uint32(v44+v75<<(uint(int32(3))%32))))
	v83 = F_Float8GetDatum(m, v82)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L28
	}
L27:
	;
	v180 = v36
	v181 = v59
	v189 = v45
	goto L16
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v59+v75<<(uint(int32(2))%32)))) = v83
	if v75 != v52 {
		v64 = v75
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	v180 = v26
	v181 = v171
	v189 = v179
	goto L16
L31:
	;
	v90 = v26 - int32(1)
	F_CheckDim_3(m, v90)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v126 = v26 - int32(1)
	F_CheckDim_3(m, v126)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L42
	}
L34:
	;
	v95 = F_mul_size(m, int32(4), v26)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v97 = F_palloc(m, v95)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	if v26 == int32(1) {
		v180 = int32(1)
		v181 = v97
		v189 = v48
		goto L16
	} else {
		goto L37
	}
L37:
	;
	v102 = int32(0)
	goto L38
L38:
	;
	v113 = v102 + int32(1)
	v120 = *(*float64)(unsafe.Add(mBase, uint32(v47+v113<<(uint(int32(3))%32))))
	v121 = F_Float8GetDatum(m, v120)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L40
	}
L39:
	;
	v171 = v97
	v179 = v48
	goto L30
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v97+v113<<(uint(int32(2))%32)))) = v121
	if v113 != v90 {
		v102 = v113
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
	if v26 != v129 {
		goto L5
	} else {
		goto L43
	}
L43:
	;
	v131 = base.F64_add(v45, v48)
	v134 = F_mul_size(m, int32(4), v26)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v136 = F_palloc(m, v134)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	if v26 == int32(1) {
		v180 = int32(1)
		v181 = v136
		v189 = v131
		goto L16
	} else {
		goto L46
	}
L46:
	;
	v141 = int32(0)
	goto L47
L47:
	;
	v152 = v141 + int32(1)
	v154 = v152 << (uint(int32(3)) % 32)
	v156 = *(*float64)(unsafe.Add(mBase, uint32(v44+v154)))
	v158 = *(*float64)(unsafe.Add(mBase, uint32(v154+v47)))
	v159 = base.F64_add(v156, v158)
	if base.F64_eq(base.F64_abs(v159), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L4
	} else {
		goto L49
	}
L48:
	;
	v171 = v136
	v179 = v131
	goto L30
L49:
	;
	v166 = F_Float8GetDatum(m, v159)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v136+v152<<(uint(int32(2))%32)))) = v166
	if v152 != v126 {
		v141 = v152
		goto L47
	} else {
		goto L51
	}
L51:
	;
	goto L48
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v181))) = v190
	v197 = F_construct_array(m, v181, v180, int32(701), int32(8), int32(0), int32(100))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	F_pfree(m, v181)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	m.G0 = v13 + int32(48)
	return v197
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(391298)
	F_errmsg_internal(m, int32(26326), v13)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(517975), int32(169), int32(26817))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = int32(391298)
	F_errmsg_internal(m, int32(26326), v13+int32(16))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	F_errfinish(m, int32(517975), int32(169), int32(26817))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L61:
	;
	F_errcode(m, int32(130))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+36)) = v129 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v126
	F_errmsg(m, int32(488627), v13+int32(32))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(517975), int32(88), int32(302025))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L65:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
