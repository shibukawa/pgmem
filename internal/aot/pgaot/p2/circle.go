package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_circle_le(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 float64
	_ = v8
	var v9 float64
	_ = v9
	var v11 float64
	_ = v11
	var v20 float64
	_ = v20
	var v23 int32
	_ = v23
	var v24 float64
	_ = v24
	var v30 float64
	_ = v30
	var v31 int32
	_ = v31
	var v32 float64
	_ = v32
	var v34 float64
	_ = v34
	var v37 float64
	_ = v37
	var v45 float64
	_ = v45
	var v46 int32
	_ = v46
	var v47 float64
	_ = v47
	var v53 float64
	_ = v53
	var v54 int32
	_ = v54
	var v55 float64
	_ = v55
	var v56 float64
	_ = v56
	var v57 float64
	_ = v57
	var v59 float64
	_ = v59
	var v68 float64
	_ = v68
	var v69 int32
	_ = v69
	var v70 float64
	_ = v70
	var v76 float64
	_ = v76
	var v77 int32
	_ = v77
	var v78 float64
	_ = v78
	var v80 float64
	_ = v80
	var v83 float64
	_ = v83
	var v91 float64
	_ = v91
	var v92 int32
	_ = v92
	var v93 float64
	_ = v93
	var v99 float64
	_ = v99
	var v100 int32
	_ = v100
	var v101 float64
	_ = v101
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v8 = *(*float64)(unsafe.Add(mBase, uint32(v7)+16))
	v9 = base.F64_mul(v8, v8)
	v11 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v9), v11)|base.F64_eq(base.F64_abs(v8), v11) == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v34 = math.Float64frombits(uint64(0x7ff0000000000000))
	v37 = base.F64_mul(v32, float64(3.141592653589793))
	if base.F64_eq(base.F64_abs(v32), v34)|base.F64_ne(base.F64_abs(v37), v34) == int32(0) {
		goto L10
	} else {
		goto L11
	}
L2:
	;
	v20 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v24 = float64(0)
	if base.F64_eq(v8, v24)|base.F64_ne(v9, v24) != 0 {
		v32 = v9
		goto L1
	} else {
		goto L7
	}
L5:
	;
	return int64(0)
L6:
	;
	v32 = v20
	goto L1
L7:
	;
	v30 = F_float_underflow_error_ext(m, int32(0))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v32 = v30
	goto L1
L9:
	;
	v56 = *(*float64)(unsafe.Add(mBase, uint32(v6)+16))
	v57 = base.F64_mul(v56, v56)
	v59 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v57), v59)|base.F64_eq(base.F64_abs(v56), v59) == int32(0) {
		goto L17
	} else {
		goto L18
	}
L10:
	;
	v45 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L5
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v47 = float64(0)
	if base.F64_eq(v32, v47)|base.F64_ne(v37, v47) != 0 {
		v55 = v37
		goto L9
	} else {
		goto L14
	}
L13:
	;
	v55 = v45
	goto L9
L14:
	;
	v53 = F_float_underflow_error_ext(m, int32(0))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L5
	} else {
		goto L15
	}
L15:
	;
	v55 = v53
	goto L9
L16:
	;
	v80 = math.Float64frombits(uint64(0x7ff0000000000000))
	v83 = base.F64_mul(v78, float64(3.141592653589793))
	if base.F64_eq(base.F64_abs(v78), v80)|base.F64_ne(base.F64_abs(v83), v80) == int32(0) {
		goto L24
	} else {
		goto L25
	}
L17:
	;
	v68 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L5
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v70 = float64(0)
	if base.F64_eq(v56, v70)|base.F64_ne(v57, v70) != 0 {
		v78 = v57
		goto L16
	} else {
		goto L21
	}
L20:
	;
	v78 = v68
	goto L16
L21:
	;
	v76 = F_float_underflow_error_ext(m, int32(0))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L5
	} else {
		goto L22
	}
L22:
	;
	v78 = v76
	goto L16
L23:
	;
	return base.I64_extend_i32_u(base.F64_le(v55, base.F64_add(v101, float64(1e-06))))
L24:
	;
	v91 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L5
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v93 = float64(0)
	if base.F64_eq(v78, v93)|base.F64_ne(v83, v93) != 0 {
		v101 = v83
		goto L23
	} else {
		goto L28
	}
L27:
	;
	v101 = v91
	goto L23
L28:
	;
	v99 = F_float_underflow_error_ext(m, int32(0))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L5
	} else {
		goto L29
	}
L29:
	;
	v101 = v99
	goto L23
}
func F_circle_mul_pt(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 float64
	_ = v16
	var v17 float64
	_ = v17
	var v18 float64
	_ = v18
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 float64
	_ = v31
	var v32 float64
	_ = v32
	var v35 int32
	_ = v35
	var v36 float64
	_ = v36
	var v37 int64
	_ = v37
	var v39 int64
	_ = v39
	var v42 float64
	_ = v42
	var v45 int64
	_ = v45
	var v47 int64
	_ = v47
	var v58 float64
	_ = v58
	var v66 float64
	_ = v66
	var v71 float64
	_ = v71
	var v72 float64
	_ = v72
	var v73 float64
	_ = v73
	var v82 float64
	_ = v82
	var v83 float64
	_ = v83
	var v85 float64
	_ = v85
	var v87 float64
	_ = v87
	var v94 float64
	_ = v94
	var v100 float64
	_ = v100
	var v102 float64
	_ = v102
	var v115 float64
	_ = v115
	var v116 int32
	_ = v116
	var v117 float64
	_ = v117
	var v126 float64
	_ = v126
	var v127 int32
	_ = v127
	var v128 float64
	_ = v128
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = F_palloc(m, int32(24))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int64(0)
	} else {
		F_point_mul_point(m, v10, v8, v7)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int64(0)
		} else {
			v16 = *(*float64)(unsafe.Add(mBase, uint32(v8)+16))
			v17 = *(*float64)(unsafe.Add(mBase, uint32(v7)))
			v18 = *(*float64)(unsafe.Add(mBase, uint32(v7)+8))
			v27 = m.G0
			v29 = v27 - int32(32)
			m.G0 = v29
			v31 = base.F64_abs(v17)
			v32 = base.F64_abs(v18)
			v35 = base.B2i32(base.Ui64(base.I64_reinterpret_f64(v31)) < base.Ui64(base.I64_reinterpret_f64(v32)))
			if base.Ui64(base.I64_reinterpret_f64(v31)) < base.Ui64(base.I64_reinterpret_f64(v32)) {
				v36 = v31
			} else {
				v36 = v32
			}
			v37 = base.I64_reinterpret_f64(v36)
			v39 = int64(base.Ui64(v37) >> (uint(int64(52)) % 64))
			if v39 == int64(2047) {
				v94 = v36
			} else {
				if base.Ui64(base.I64_reinterpret_f64(v31)) < base.Ui64(base.I64_reinterpret_f64(v32)) {
					v42 = v32
				} else {
					v42 = v31
				}
				if v37 == int64(0) {
					v94 = v42
				} else {
					v45 = base.I64_reinterpret_f64(v42)
					v47 = int64(base.Ui64(v45) >> (uint(int64(52)) % 64))
					if v47 == int64(2047) {
						v94 = v42
					} else {
						if int32(65) <= base.I32_wrap_i64(v47)-base.I32_wrap_i64(v39) {
							v94 = base.F64_add(v31, v32)
						} else {
							if base.Ui64(int64(6908521828386340864)) <= base.Ui64(v45) {
								v58 = float64(1.90109156629516e-211)
								v71 = base.F64_mul(v42, v58)
								v72 = base.F64_mul(v36, v58)
								v73 = float64(5.260135901548374e+210)
							} else {
								if base.Ui64(int64(2580562586483294207)) < base.Ui64(v37) {
									v71 = v42
									v72 = v36
									v73 = float64(1)
								} else {
									v66 = float64(5.260135901548374e+210)
									v71 = base.F64_mul(v42, v66)
									v72 = base.F64_mul(v36, v66)
									v73 = float64(1.90109156629516e-211)
								}
							}
							F_sq(m, v29+int32(24), v29+int32(16), v71)
							mBase = m.M
							F_sq(m, v29+int32(8), v29, v72)
							mBase = m.M
							v82 = *(*float64)(unsafe.Add(mBase, uint32(v29)))
							v83 = *(*float64)(unsafe.Add(mBase, uint32(v29)+16))
							v85 = *(*float64)(unsafe.Add(mBase, uint32(v29)+8))
							v87 = *(*float64)(unsafe.Add(mBase, uint32(v29)+24))
							v94 = base.F64_mul(v73, base.F64_sqrt(base.F64_add(base.F64_add(base.F64_add(v82, v83), v85), v87)))
						}
					}
				}
			}
			m.G0 = v29 + int32(32)
			v100 = base.F64_mul(v16, v94)
			v102 = math.Float64frombits(uint64(0x7ff0000000000000))
			if base.F64_ne(base.F64_abs(v100), v102)|base.F64_eq(base.F64_abs(v16), v102)|base.F64_eq(base.F64_abs(v94), v102) == int32(0) {
				v115 = F_float_overflow_error_ext(m, int32(0))
				mBase = m.M
				v116 = m.ExcPending
				if v116 != 0 {
					return int64(0)
				} else {
					v128 = v115
					*(*float64)(unsafe.Add(mBase, uint32(v10)+16)) = v128
					return base.I64_extend_i32_u(v10)
				}
			} else {
				v117 = float64(0)
				if base.F64_eq(v16, v117)|base.F64_ne(v100, v117)|base.F64_eq(v94, v117) != 0 {
					v128 = v100
					*(*float64)(unsafe.Add(mBase, uint32(v10)+16)) = v128
					return base.I64_extend_i32_u(v10)
				} else {
					v126 = F_float_underflow_error_ext(m, int32(0))
					mBase = m.M
					v127 = m.ExcPending
					if v127 != 0 {
						return int64(0)
					} else {
						v128 = v126
						*(*float64)(unsafe.Add(mBase, uint32(v10)+16)) = v128
						return base.I64_extend_i32_u(v10)
					}
				}
			}
		}
	}
}
