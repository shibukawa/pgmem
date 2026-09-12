package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_box_ar(m *base.Module, l0 int32) float64 {
	mBase := m.M
	_ = mBase
	var v8 float64
	_ = v8
	var v9 float64
	_ = v9
	var v10 float64
	_ = v10
	var v11 float64
	_ = v11
	var v20 float64
	_ = v20
	var v21 float64
	_ = v21
	var v22 float64
	_ = v22
	var v23 float64
	_ = v23
	var v32 float64
	_ = v32
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	v8 = *(*float64)(unsafe.Add(mBase, uint32(l0)))
	v9 = *(*float64)(unsafe.Add(mBase, uint32(l0)+16))
	v10 = base.F64_sub(v8, v9)
	v11 = base.F64_abs(v10)
	if base.F64_ne(v11, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_float_underflow_error(m)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L19
	} else {
		goto L21
	}
L2:
	;
	F_float_overflow_error(m)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L19
	} else {
		goto L20
	}
L3:
	;
	v20 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
	v21 = *(*float64)(unsafe.Add(mBase, uint32(l0)+24))
	v22 = base.F64_sub(v20, v21)
	v23 = base.F64_abs(v22)
	if base.F64_ne(v23, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	if base.F64_eq(base.F64_abs(v8), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	if base.F64_ne(base.F64_abs(v9), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L2
	} else {
		goto L6
	}
L6:
	;
	goto L3
L7:
	;
	v32 = base.F64_mul(v10, v22)
	if base.F64_ne(base.F64_abs(v32), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	if base.F64_eq(base.F64_abs(v20), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	if base.F64_ne(base.F64_abs(v21), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	goto L7
L11:
	;
	if base.F64_ne(v32, float64(0)) != 0 {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	if base.F64_eq(v11, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	if base.F64_ne(v23, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L2
	} else {
		goto L14
	}
L14:
	;
	goto L11
L15:
	;
	return v32
L16:
	;
	if base.F64_eq(v10, float64(0)) != 0 {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	if base.F64_ne(v22, float64(0)) != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	goto L15
L19:
	;
	return float64(0)
L20:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L21:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_box_center(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_palloc(m, int32(16))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		F_box_cn(m, v5, v3)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			return v5
		}
	}
}
func F_box_distance(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 float64
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_box_cn(m, v6+int32(16), v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		F_box_cn(m, v6, v8)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v20 = F_point_dt(m, v6+int32(16), v6)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				v22 = F_Float8GetDatum(m, v20)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					m.G0 = v6 + int32(32)
					return v22
				}
			}
		}
	}
}
func F_box_le(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 float64
	_ = v5
	var v8 int32
	_ = v8
	var v9 float64
	_ = v9
	var v10 int32
	_ = v10
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_box_ar(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = F_box_ar(m, v3)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			return base.F64_le(v5, base.F64_add(v9, float64(1e-06)))
		}
	}
}
func F_box_left(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 float64
	_ = v3
	var v4 int32
	_ = v4
	var v5 float64
	_ = v5
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v3 = *(*float64)(unsafe.Add(mBase, uint32(v2)+16))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*float64)(unsafe.Add(mBase, uint32(v4)))
	return base.F64_gt(v3, base.F64_add(v5, float64(1e-06)))
}
func F_box_overlap(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v5 float64
	_ = v5
	var v6 int32
	_ = v6
	var v7 float64
	_ = v7
	var v13 float64
	_ = v13
	var v14 float64
	_ = v14
	var v20 float64
	_ = v20
	var v21 float64
	_ = v21
	var v27 float64
	_ = v27
	var v28 float64
	_ = v28
	var v32 int32
	_ = v32
	v2 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*float64)(unsafe.Add(mBase, uint32(v4)+16))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = *(*float64)(unsafe.Add(mBase, uint32(v6)))
	if base.F64_le(v5, base.F64_add(v7, float64(1e-06))) == v2 {
		v32 = v2
	} else {
		v13 = *(*float64)(unsafe.Add(mBase, uint32(v6)+16))
		v14 = *(*float64)(unsafe.Add(mBase, uint32(v4)))
		if base.F64_le(v13, base.F64_add(v14, float64(1e-06))) == int32(0) {
			v32 = v2
		} else {
			v20 = *(*float64)(unsafe.Add(mBase, uint32(v4)+24))
			v21 = *(*float64)(unsafe.Add(mBase, uint32(v6)+8))
			if base.F64_le(v20, base.F64_add(v21, float64(1e-06))) == int32(0) {
				v32 = v2
			} else {
				v27 = *(*float64)(unsafe.Add(mBase, uint32(v6)+24))
				v28 = *(*float64)(unsafe.Add(mBase, uint32(v4)+8))
				v32 = base.F64_le(v27, base.F64_add(v28, float64(1e-06)))
			}
		}
	}
	return v32
}
func F_write_box(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v40 float64
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v79 int32
	_ = v79
	var v83 float64
	_ = v83
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v121 int32
	_ = v121
	var v128 float64
	_ = v128
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v142 float64
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v174 int32
	_ = v174
	var v178 float64
	_ = v178
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v196 float64
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v224 int32
	_ = v224
	var v242 int32
	_ = v242
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v20 = l0<<(uint(int32(4))%32) | int32(8)
	v21 = F_palloc0(m, v20)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v20 << (uint(int32(2)) % 32)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = v28&int32(-2147483648) | l0
	v33 = int32(1)
	if l0 <= int32(0) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	m.G0 = v15 + int32(16)
	return v242
L4:
	;
	v53 = int32(44)
	v54 = F___strchrnul(m, l1, v53)
	mBase = m.M
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
	if v56 == v53 {
		goto L13
	} else {
		goto L14
	}
L5:
	;
	v52 = int32(0)
	goto L4
L6:
	;
	goto L7
L7:
	;
	v40 = F_float8in_internal(m, l1, v15+int32(12), int32(429367), l1, l4)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v21)+8)) = v40
	if l4 == int32(0) {
		v52 = v33
		goto L4
	} else {
		goto L9
	}
L9:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v45 != int32(447) {
		v52 = v33
		goto L4
	} else {
		goto L10
	}
L10:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+4)))
	if v48 == int32(0) {
		v52 = v33
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v242 = int32(0)
	goto L3
L12:
	;
	if v60 != 0 {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	v60 = v54
	goto L15
L14:
	;
	v60 = int32(0)
	goto L15
L15:
	;
	goto L12
L16:
	;
	v68 = v52
	v70 = v60
	goto L19
L17:
	;
	v110 = v52
	goto L18
L18:
	;
	if int32(0) < l0 {
		goto L31
	} else {
		goto L32
	}
L19:
	;
	v79 = v70 + int32(1)
	v83 = F_float8in_internal(m, v79, v15+int32(12), int32(429367), l1, l4)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L21
	}
L20:
	;
	v110 = v96
	goto L18
L21:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v21+int32(8)+v68<<(uint(int32(3))%32)))) = v83
	if l4 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v96 = v68 + int32(1)
	v97 = int32(44)
	v98 = F___strchrnul(m, v79, v97)
	mBase = m.M
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98))))
	if v100 == v97 {
		goto L27
	} else {
		goto L28
	}
L23:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v88 != int32(447) {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+4)))
	if v91 == int32(0) {
		goto L22
	} else {
		goto L25
	}
L25:
	;
	v242 = int32(0)
	goto L3
L26:
	;
	if v104 != 0 {
		v68 = v96
		v70 = v104
		goto L19
	} else {
		goto L30
	}
L27:
	;
	v104 = v98
	goto L29
L28:
	;
	v104 = int32(0)
	goto L29
L29:
	;
	goto L26
L30:
	;
	goto L20
L31:
	;
	v121 = v21 + int32(8)
	v128 = F_float8in_internal(m, l2, v15+int32(12), int32(429367), l2, l4)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L34
	}
L32:
	;
	v144 = int32(1)
	v145 = v110
	goto L33
L33:
	;
	v148 = int32(44)
	v149 = F___strchrnul(m, l2, v148)
	mBase = m.M
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149))))
	if v151 == v148 {
		goto L40
	} else {
		goto L41
	}
L34:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v121+v110<<(uint(int32(3))%32)))) = v128
	if l4 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v142 = *(*float64)(unsafe.Add(mBase, uint32(v121)))
	v144 = base.F64_eq(v128, v142)
	v145 = v110 + int32(1)
	goto L33
L36:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v133 != int32(447) {
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+4)))
	if v136 == int32(0) {
		goto L35
	} else {
		goto L38
	}
L38:
	;
	v242 = int32(0)
	goto L3
L39:
	;
	if v155 != 0 {
		goto L43
	} else {
		goto L44
	}
L40:
	;
	v155 = v149
	goto L42
L41:
	;
	v155 = int32(0)
	goto L42
L42:
	;
	goto L39
L43:
	;
	v157 = v21 + int32(8)
	v159 = v144
	v163 = v145
	v166 = v155
	goto L46
L44:
	;
	v208 = v144
	goto L45
L45:
	;
	if v208 != 0 {
		goto L58
	} else {
		goto L59
	}
L46:
	;
	v174 = v166 + int32(1)
	v178 = F_float8in_internal(m, v174, v15+int32(12), int32(429367), l2, l4)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L1
	} else {
		goto L48
	}
L47:
	;
	v208 = v198
	goto L45
L48:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v157+v163<<(uint(int32(3))%32)))) = v178
	if l4 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v196 = *(*float64)(unsafe.Add(mBase, uint32(v157+(v163-l0)<<(uint(int32(3))%32))))
	v198 = v159 & base.F64_eq(v178, v196)
	v199 = int32(44)
	v200 = F___strchrnul(m, v174, v199)
	mBase = m.M
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v200))))
	if v202 == v199 {
		goto L54
	} else {
		goto L55
	}
L50:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v183 != int32(447) {
		goto L49
	} else {
		goto L51
	}
L51:
	;
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+4)))
	if v186 == int32(0) {
		goto L49
	} else {
		goto L52
	}
L52:
	;
	v242 = int32(0)
	goto L3
L53:
	;
	if v206 != 0 {
		v159 = v198
		v163 = v163 + int32(1)
		v166 = v206
		goto L46
	} else {
		goto L57
	}
L54:
	;
	v206 = v200
	goto L56
L55:
	;
	v206 = int32(0)
	goto L56
L56:
	;
	goto L53
L57:
	;
	goto L47
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = l0<<(uint(int32(5))%32) + int32(32)
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = v224 | int32(-2147483648)
	goto L60
L59:
	;
	goto L60
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v21
	v242 = int32(1)
	goto L3
}
