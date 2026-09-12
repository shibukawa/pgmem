package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_line_eq(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 float64
	_ = v16
	var v22 float64
	_ = v22
	var v28 float64
	_ = v28
	var v34 float64
	_ = v34
	var v35 float64
	_ = v35
	var v39 float64
	_ = v39
	var v40 float64
	_ = v40
	var v44 float64
	_ = v44
	var v45 float64
	_ = v45
	var v56 float64
	_ = v56
	var v63 int64
	_ = v63
	var v74 float64
	_ = v74
	var v77 int32
	_ = v77
	var v82 float64
	_ = v82
	var v83 int32
	_ = v83
	var v87 float64
	_ = v87
	var v88 int32
	_ = v88
	var v89 float64
	_ = v89
	var v90 float64
	_ = v90
	var v91 float64
	_ = v91
	var v92 float64
	_ = v92
	var v115 float64
	_ = v115
	var v116 float64
	_ = v116
	var v117 float64
	_ = v117
	var v118 int32
	_ = v118
	var v126 float64
	_ = v126
	var v127 float64
	_ = v127
	var v128 float64
	_ = v128
	var v129 int32
	_ = v129
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v148 float64
	_ = v148
	var v150 int64
	_ = v150
	var v151 int64
	_ = v151
	var v152 float64
	_ = v152
	var v165 float64
	_ = v165
	var v167 int64
	_ = v167
	var v168 int64
	_ = v168
	var v169 float64
	_ = v169
	var v192 int32
	_ = v192
	v11 = int32(0)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = *(*float64)(unsafe.Add(mBase, uint32(v15)))
	if base.Ui64(base.I64_reinterpret_f64(v16)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	return v192
L2:
	;
	v148 = *(*float64)(unsafe.Add(mBase, uint32(v14)+8))
	v150 = int64(9223372036854775807)
	v151 = base.I64_reinterpret_f64(v148) & v150
	v152 = *(*float64)(unsafe.Add(mBase, uint32(v15)+8))
	if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v152)&v150) {
		goto L51
	} else {
		goto L52
	}
L3:
	;
	if base.F64_le(v35, float64(1e-06)) == int32(0) {
		goto L17
	} else {
		goto L18
	}
L4:
	;
	v22 = *(*float64)(unsafe.Add(mBase, uint32(v15)+8))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v22)&int64(9223372036854775807)) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v63 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(v63&int64(9223372036854775807)) {
		goto L2
	} else {
		goto L15
	}
L7:
	;
	v56 = *(*float64)(unsafe.Add(mBase, uint32(v14)))
	if base.F64_ne(v16, v56) != 0 {
		v192 = v11
		goto L1
	} else {
		goto L13
	}
L8:
	;
	v28 = *(*float64)(unsafe.Add(mBase, uint32(v15)+16))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v28)&int64(9223372036854775807)) {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v34 = *(*float64)(unsafe.Add(mBase, uint32(v14)))
	v35 = base.F64_abs(v34)
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v35)) {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v39 = *(*float64)(unsafe.Add(mBase, uint32(v14)+8))
	v40 = base.F64_abs(v39)
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v40)) {
		goto L7
	} else {
		goto L11
	}
L11:
	;
	v44 = *(*float64)(unsafe.Add(mBase, uint32(v14)+16))
	v45 = base.F64_abs(v44)
	if base.Ui64(base.I64_reinterpret_f64(v45)) <= base.Ui64(int64(9218868437227405312)) {
		goto L3
	} else {
		goto L12
	}
L12:
	;
	goto L7
L13:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v56)&int64(9223372036854775807)) {
		v192 = v11
		goto L1
	} else {
		goto L14
	}
L14:
	;
	goto L2
L15:
	;
	return int32(0)
L16:
	;
	v90 = *(*float64)(unsafe.Add(mBase, uint32(v15)))
	v91 = *(*float64)(unsafe.Add(mBase, uint32(v14)))
	v92 = base.F64_mul(v89, v91)
	if base.F64_ne(base.F64_abs(v92), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L30
	} else {
		goto L31
	}
L17:
	;
	v74 = F_float8_div(m, v16, v34)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	goto L19
L19:
	;
	if base.F64_le(v40, float64(1e-06)) == int32(0) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	return int32(0)
L21:
	;
	v89 = v74
	goto L16
L22:
	;
	v82 = F_float8_div(m, v22, v39)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L20
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	if base.F64_le(v45, float64(1e-06)) != 0 {
		v89 = float64(1)
		goto L16
	} else {
		goto L26
	}
L25:
	;
	v89 = v82
	goto L16
L26:
	;
	v87 = F_float8_div(m, v28, v44)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L20
	} else {
		goto L27
	}
L27:
	;
	v89 = v87
	goto L16
L28:
	;
	F_float_underflow_error(m)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L20
	} else {
		goto L49
	}
L29:
	;
	F_float_overflow_error(m)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L20
	} else {
		goto L48
	}
L30:
	;
	if base.F64_ne(v92, float64(0)) != 0 {
		goto L34
	} else {
		goto L35
	}
L31:
	;
	if base.F64_eq(base.F64_abs(v89), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L30
	} else {
		goto L32
	}
L32:
	;
	if base.F64_ne(base.F64_abs(v91), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L29
	} else {
		goto L33
	}
L33:
	;
	goto L30
L34:
	;
	if base.F64_ne(v92, v90) != 0 {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	if base.F64_eq(v89, float64(0)) != 0 {
		goto L34
	} else {
		goto L36
	}
L36:
	;
	if base.F64_ne(v91, float64(0)) != 0 {
		goto L28
	} else {
		goto L37
	}
L37:
	;
	goto L34
L38:
	;
	if base.F64_le(base.F64_abs(base.F64_sub(v90, v92)), float64(1e-06)) == int32(0) {
		v192 = v11
		goto L1
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v115 = *(*float64)(unsafe.Add(mBase, uint32(v15)+8))
	v116 = *(*float64)(unsafe.Add(mBase, uint32(v14)+8))
	v117 = F_float8_mul(m, v89, v116)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L20
	} else {
		goto L42
	}
L41:
	;
	goto L40
L42:
	;
	if base.F64_ne(v115, v117) != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	if base.F64_le(base.F64_abs(base.F64_sub(v115, v117)), float64(1e-06)) == int32(0) {
		v192 = v11
		goto L1
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v126 = *(*float64)(unsafe.Add(mBase, uint32(v15)+16))
	v127 = *(*float64)(unsafe.Add(mBase, uint32(v14)+16))
	v128 = F_float8_mul(m, v89, v127)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L20
	} else {
		goto L47
	}
L46:
	;
	goto L45
L47:
	;
	v192 = base.F64_eq(v126, v128) | base.F64_le(base.F64_abs(base.F64_sub(v126, v128)), float64(1e-06))
	goto L1
L48:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L49:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L50:
	;
	v165 = *(*float64)(unsafe.Add(mBase, uint32(v14)+16))
	v167 = int64(9223372036854775807)
	v168 = base.I64_reinterpret_f64(v165) & v167
	v169 = *(*float64)(unsafe.Add(mBase, uint32(v15)+16))
	if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v169)&v167) {
		goto L57
	} else {
		goto L58
	}
L51:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(v151) {
		goto L50
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	if base.F64_ne(v152, v148) != 0 {
		v192 = v11
		goto L1
	} else {
		goto L55
	}
L54:
	;
	return int32(0)
L55:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(v151) {
		v192 = v11
		goto L1
	} else {
		goto L56
	}
L56:
	;
	goto L50
L57:
	;
	return base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v168))
L58:
	;
	goto L59
L59:
	;
	return base.B2i32(base.Ui64(v168) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v169, v165)
}
func F_line_recv(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 float64
	_ = v9
	var v10 int32
	_ = v10
	var v12 float64
	_ = v12
	var v13 int32
	_ = v13
	var v15 float64
	_ = v15
	var v16 int32
	_ = v16
	var v18 float64
	_ = v18
	var v24 float64
	_ = v24
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_palloc(m, int32(24))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = F_pq_getmsgfloat8(m, v3)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			*(*float64)(unsafe.Add(mBase, uint32(v5))) = v9
			v12 = F_pq_getmsgfloat8(m, v3)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				*(*float64)(unsafe.Add(mBase, uint32(v5)+8)) = v12
				v15 = F_pq_getmsgfloat8(m, v3)
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return int32(0)
				} else {
					*(*float64)(unsafe.Add(mBase, uint32(v5)+16)) = v15
					v18 = *(*float64)(unsafe.Add(mBase, uint32(v5)))
					if base.F64_le(base.F64_abs(v18), float64(1e-06)) == int32(0) {
						return v5
					} else {
						v24 = *(*float64)(unsafe.Add(mBase, uint32(v5)+8))
						if base.F64_le(base.F64_abs(v24), float64(1e-06)) == int32(0) {
							return v5
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(50462850))
								mBase = m.M
								v36 = m.ExcPending
								if v36 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(240119), int32(0))
									mBase = m.M
									v40 = m.ExcPending
									if v40 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(493522), int32(1052), int32(36573))
										mBase = m.M
										v45 = m.ExcPending
										if v45 != 0 {
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
		}
	}
}
func F_line_send(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 float64
	_ = v12
	var v14 int32
	_ = v14
	var v15 float64
	_ = v15
	var v17 int32
	_ = v17
	var v18 float64
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_pq_begintypsend(m, v5)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v12 = *(*float64)(unsafe.Add(mBase, uint32(v7)))
		F_pq_sendfloat8(m, v5, v12)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			v15 = *(*float64)(unsafe.Add(mBase, uint32(v7)+8))
			F_pq_sendfloat8(m, v5, v15)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				v18 = *(*float64)(unsafe.Add(mBase, uint32(v7)+16))
				F_pq_sendfloat8(m, v5, v18)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
					v23 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v22))) = v23 << (uint(int32(2)) % 32)
					m.G0 = v5 + int32(16)
					return v22
				}
			}
		}
	}
}
