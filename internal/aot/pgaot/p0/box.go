package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_box_contain(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v5 float64
	_ = v5
	var v8 int32
	_ = v8
	var v9 float64
	_ = v9
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
	v5 = *(*float64)(unsafe.Add(mBase, uint32(v4)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v9 = *(*float64)(unsafe.Add(mBase, uint32(v8)))
	if base.F64_ge(base.F64_add(v5, float64(1e-06)), v9) == v2 {
		v32 = v2
	} else {
		v13 = *(*float64)(unsafe.Add(mBase, uint32(v4)+16))
		v14 = *(*float64)(unsafe.Add(mBase, uint32(v8)+16))
		if base.F64_le(v13, base.F64_add(v14, float64(1e-06))) == int32(0) {
			v32 = v2
		} else {
			v20 = *(*float64)(unsafe.Add(mBase, uint32(v8)+8))
			v21 = *(*float64)(unsafe.Add(mBase, uint32(v4)+8))
			if base.F64_le(v20, base.F64_add(v21, float64(1e-06))) == int32(0) {
				v32 = v2
			} else {
				v27 = *(*float64)(unsafe.Add(mBase, uint32(v4)+24))
				v28 = *(*float64)(unsafe.Add(mBase, uint32(v8)+24))
				v32 = base.F64_le(v27, base.F64_add(v28, float64(1e-06)))
			}
		}
	}
	return v32
}
func F_box_interpt_lseg(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 float64
	_ = v18
	var v19 float64
	_ = v19
	var v22 int64
	_ = v22
	var v30 float64
	_ = v30
	var v32 float64
	_ = v32
	var v33 float64
	_ = v33
	var v34 float64
	_ = v34
	var v40 float64
	_ = v40
	var v43 float64
	_ = v43
	var v45 float64
	_ = v45
	var v51 float64
	_ = v51
	var v57 float64
	_ = v57
	var v60 float64
	_ = v60
	var v61 float64
	_ = v61
	var v66 int32
	_ = v66
	var v67 float64
	_ = v67
	var v69 float64
	_ = v69
	var v74 int32
	_ = v74
	var v75 float64
	_ = v75
	var v79 float64
	_ = v79
	var v80 float64
	_ = v80
	var v82 float64
	_ = v82
	var v83 float64
	_ = v83
	var v92 int32
	_ = v92
	var v93 float64
	_ = v93
	var v94 int32
	_ = v94
	var v95 float64
	_ = v95
	var v96 float64
	_ = v96
	var v97 float64
	_ = v97
	var v98 float64
	_ = v98
	var v99 float64
	_ = v99
	var v100 float64
	_ = v100
	var v107 float64
	_ = v107
	var v108 float64
	_ = v108
	var v112 float64
	_ = v112
	var v115 float64
	_ = v115
	var v119 float64
	_ = v119
	var v120 float64
	_ = v120
	var v124 float64
	_ = v124
	var v128 float64
	_ = v128
	var v135 float64
	_ = v135
	var v138 float64
	_ = v138
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 float64
	_ = v148
	var v150 float64
	_ = v150
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 float64
	_ = v159
	var v160 float64
	_ = v160
	var v161 float64
	_ = v161
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 float64
	_ = v171
	var v173 float64
	_ = v173
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v187 int32
	_ = v187
	v9 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(48)
	m.G0 = v16
	v18 = *(*float64)(unsafe.Add(mBase, uint32(l2)+16))
	v19 = *(*float64)(unsafe.Add(mBase, uint32(l2)))
	v22 = base.I64_reinterpret_f64(v19) & int64(9223372036854775807)
	if base.Ui64(v22) <= base.Ui64(int64(9218868437227405312)) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v16 + int32(48)
	return v187
L2:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v18)&int64(9223372036854775807)) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	v33 = v18
	goto L4
L4:
	;
	v34 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
	if base.F64_le(v33, base.F64_add(v34, float64(1e-06))) == int32(0) {
		v187 = v9
		goto L1
	} else {
		goto L11
	}
L5:
	;
	v30 = v19
	goto L7
L6:
	;
	v30 = v18
	goto L7
L7:
	;
	if base.F64_gt(v18, v19) != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v32 = v19
	goto L10
L9:
	;
	v32 = v30
	goto L10
L10:
	;
	v33 = v32
	goto L4
L11:
	;
	v40 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(v22) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v43 = v19
	goto L14
L13:
	;
	v43 = v18
	goto L14
L14:
	;
	if base.F64_lt(v18, v19) != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v45 = v19
	goto L17
L16:
	;
	v45 = v43
	goto L17
L17:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v18)&int64(9223372036854775807)) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v51 = v18
	goto L20
L19:
	;
	v51 = v45
	goto L20
L20:
	;
	if base.F64_le(v40, base.F64_add(v51, float64(1e-06))) == int32(0) {
		v187 = v9
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v57 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
	v60 = *(*float64)(unsafe.Add(mBase, uint32(l2)+24))
	v61 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
	v66 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v60)&int64(9223372036854775807)))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v60)&int64(9223372036854775807)) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v67 = v61
	goto L24
L23:
	;
	v67 = v60
	goto L24
L24:
	;
	if base.F64_gt(v60, v61) != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v69 = v61
	goto L27
L26:
	;
	v69 = v67
	goto L27
L27:
	;
	v74 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v61)&int64(9223372036854775807)))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v61)&int64(9223372036854775807)) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v75 = v60
	goto L30
L29:
	;
	v75 = v69
	goto L30
L30:
	;
	if base.F64_ge(base.F64_add(v57, float64(1e-06)), v75) == int32(0) {
		v187 = v9
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v79 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v61)&int64(9223372036854775807)) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v80 = v61
	goto L34
L33:
	;
	v80 = v60
	goto L34
L34:
	;
	if base.F64_lt(v60, v61) != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v82 = v61
	goto L37
L36:
	;
	v82 = v80
	goto L37
L37:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v60)&int64(9223372036854775807)) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v83 = v60
	goto L40
L39:
	;
	v83 = v82
	goto L40
L40:
	;
	if base.F64_le(v79, base.F64_add(v83, float64(1e-06))) == int32(0) {
		v187 = v9
		goto L1
	} else {
		goto L41
	}
L41:
	;
	if l0 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	F_box_cn(m, v16, l1)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	v98 = v19
	v99 = v40
	v100 = v34
	goto L44
L44:
	;
	if base.F64_le(v98, v100) == int32(0) {
		goto L50
	} else {
		goto L51
	}
L45:
	;
	return int32(0)
L46:
	;
	v93 = F_lseg_closept_point(m, l0, l2, v16)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L45
	} else {
		goto L47
	}
L47:
	;
	v95 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
	v96 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
	v97 = *(*float64)(unsafe.Add(mBase, uint32(l2)))
	v98 = v97
	v99 = v95
	v100 = v96
	goto L44
L48:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v99
	v138 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
	*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v135
	*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v99
	*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v138
	v142 = int32(1)
	v146 = F_lseg_interpt_lseg(m, int32(0), v16+int32(16), l2)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L45
	} else {
		goto L61
	}
L49:
	;
	v187 = int32(1)
	goto L1
L50:
	;
	v115 = *(*float64)(unsafe.Add(mBase, uint32(l2)+16))
	if base.F64_ge(v100, v115) == int32(0) {
		goto L55
	} else {
		goto L56
	}
L51:
	;
	if base.F64_ge(v98, v99) == int32(0) {
		goto L50
	} else {
		goto L52
	}
L52:
	;
	v107 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
	v108 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
	if base.F64_le(v107, v108) == int32(0) {
		goto L50
	} else {
		goto L53
	}
L53:
	;
	v112 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
	if base.F64_le(v112, v107) != 0 {
		goto L49
	} else {
		goto L54
	}
L54:
	;
	goto L50
L55:
	;
	v119 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
	v135 = v119
	goto L48
L56:
	;
	goto L57
L57:
	;
	v120 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
	if base.F64_ge(v115, v99) == int32(0) {
		v135 = v120
		goto L48
	} else {
		goto L58
	}
L58:
	;
	v124 = *(*float64)(unsafe.Add(mBase, uint32(l2)+24))
	if base.F64_ge(v120, v124) == int32(0) {
		v135 = v120
		goto L48
	} else {
		goto L59
	}
L59:
	;
	v128 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
	if base.F64_le(v128, v124) == int32(0) {
		v135 = v120
		goto L48
	} else {
		goto L60
	}
L60:
	;
	goto L49
L61:
	;
	if v146 != 0 {
		v187 = v142
		goto L1
	} else {
		goto L62
	}
L62:
	;
	v148 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
	*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v148
	v150 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v135
	*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v99
	*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v150
	v157 = F_lseg_interpt_lseg(m, int32(0), v16+int32(16), l2)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L45
	} else {
		goto L63
	}
L63:
	;
	if v157 != 0 {
		v187 = v142
		goto L1
	} else {
		goto L64
	}
L64:
	;
	v159 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
	v160 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
	v161 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
	*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v161
	*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v160
	*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v161
	*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v159
	v169 = F_lseg_interpt_lseg(m, int32(0), v16+int32(16), l2)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L45
	} else {
		goto L65
	}
L65:
	;
	if v169 != 0 {
		v187 = v142
		goto L1
	} else {
		goto L66
	}
L66:
	;
	v171 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
	*(*float64)(unsafe.Add(mBase, uint32(v16)+16)) = v171
	v173 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v161
	*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v160
	*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v173
	v180 = F_lseg_interpt_lseg(m, int32(0), v16+int32(16), l2)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L45
	} else {
		goto L67
	}
L67:
	;
	v187 = v180
	goto L1
}
func F_box_lt(m *base.Module, l0 int32) int32 {
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
	var v11 float64
	_ = v11
	var v12 int32
	_ = v12
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_box_ar(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v11 = F_box_ar(m, v3)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			return base.F64_lt(base.F64_add(v5, float64(1e-06)), v11)
		}
	}
}
func F_box_out(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_path_encode(m, int32(0), int32(2), v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_box_right(m *base.Module, l0 int32) int32 {
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
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*float64)(unsafe.Add(mBase, uint32(v2)+16))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*float64)(unsafe.Add(mBase, uint32(v4)))
	return base.F64_gt(v3, base.F64_add(v5, float64(1e-06)))
}
func F_box_send(m *base.Module, l0 int32) int32 {
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
	var v21 float64
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
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
					v21 = *(*float64)(unsafe.Add(mBase, uint32(v7)+24))
					F_pq_sendfloat8(m, v5, v21)
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int32(0)
					} else {
						v25 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
						v26 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v25))) = v26 << (uint(int32(2)) % 32)
						m.G0 = v5 + int32(16)
						return v25
					}
				}
			}
		}
	}
}
func F_box_sub(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 float64
	_ = v15
	var v16 float64
	_ = v16
	var v17 float64
	_ = v17
	var v27 float64
	_ = v27
	var v28 float64
	_ = v28
	var v29 float64
	_ = v29
	var v41 float64
	_ = v41
	var v42 float64
	_ = v42
	var v43 float64
	_ = v43
	var v53 float64
	_ = v53
	var v54 float64
	_ = v54
	var v55 float64
	_ = v55
	var v73 int32
	_ = v73
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_palloc(m, int32(32))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v15 = *(*float64)(unsafe.Add(mBase, uint32(v9)))
	v16 = *(*float64)(unsafe.Add(mBase, uint32(v8)))
	v17 = base.F64_sub(v15, v16)
	if base.F64_ne(base.F64_abs(v17), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	F_float_overflow_error(m)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L20
	}
L4:
	;
	v27 = *(*float64)(unsafe.Add(mBase, uint32(v9)+8))
	v28 = *(*float64)(unsafe.Add(mBase, uint32(v8)+8))
	v29 = base.F64_sub(v27, v28)
	if base.F64_ne(base.F64_abs(v29), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	if base.F64_eq(base.F64_abs(v15), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	if base.F64_ne(base.F64_abs(v16), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	goto L4
L8:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v29
	*(*float64)(unsafe.Add(mBase, uint32(v11))) = v17
	v41 = *(*float64)(unsafe.Add(mBase, uint32(v9)+16))
	v42 = *(*float64)(unsafe.Add(mBase, uint32(v8)))
	v43 = base.F64_sub(v41, v42)
	if base.F64_ne(base.F64_abs(v43), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	if base.F64_eq(base.F64_abs(v27), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	if base.F64_ne(base.F64_abs(v28), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L3
	} else {
		goto L11
	}
L11:
	;
	goto L8
L12:
	;
	v53 = *(*float64)(unsafe.Add(mBase, uint32(v9)+24))
	v54 = *(*float64)(unsafe.Add(mBase, uint32(v8)+8))
	v55 = base.F64_sub(v53, v54)
	if base.F64_ne(base.F64_abs(v55), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	if base.F64_eq(base.F64_abs(v41), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	if base.F64_ne(base.F64_abs(v42), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L3
	} else {
		goto L15
	}
L15:
	;
	goto L12
L16:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v55
	*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v43
	return v11
L17:
	;
	if base.F64_eq(base.F64_abs(v53), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	if base.F64_ne(base.F64_abs(v54), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L3
	} else {
		goto L19
	}
L19:
	;
	goto L16
L20:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_box_width(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 float64
	_ = v6
	var v7 float64
	_ = v7
	var v8 float64
	_ = v8
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*float64)(unsafe.Add(mBase, uint32(v5)))
	v7 = *(*float64)(unsafe.Add(mBase, uint32(v5)+16))
	v8 = base.F64_sub(v6, v7)
	if base.F64_ne(base.F64_abs(v8), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v22 = F_Float8GetDatum(m, v8)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			return v22
		}
	} else {
		if base.F64_eq(base.F64_abs(v6), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
			v22 = F_Float8GetDatum(m, v8)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				return v22
			}
		} else {
			if base.F64_eq(base.F64_abs(v7), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
				v22 = F_Float8GetDatum(m, v8)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					return v22
				}
			} else {
				F_float_overflow_error(m)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
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
