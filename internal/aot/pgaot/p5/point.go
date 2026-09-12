package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_point_eq(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 float64
	_ = v12
	var v18 float64
	_ = v18
	var v20 int64
	_ = v20
	var v21 int64
	_ = v21
	var v22 float64
	_ = v22
	var v25 int64
	_ = v25
	var v32 float64
	_ = v32
	var v39 int64
	_ = v39
	var v44 float64
	_ = v44
	var v61 int32
	_ = v61
	var v67 float64
	_ = v67
	var v71 int64
	_ = v71
	var v72 float64
	_ = v72
	var v75 int64
	_ = v75
	var v92 int32
	_ = v92
	v8 = int32(0)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*float64)(unsafe.Add(mBase, uint32(v11)))
	if base.Ui64(base.I64_reinterpret_f64(v12)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		v18 = *(*float64)(unsafe.Add(mBase, uint32(v10)))
		v20 = int64(9223372036854775807)
		v21 = base.I64_reinterpret_f64(v18) & v20
		v22 = *(*float64)(unsafe.Add(mBase, uint32(v11)+8))
		v25 = base.I64_reinterpret_f64(v22) & v20
		if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v25) {
			v61 = base.B2i32(base.Ui64(v21) < base.Ui64(int64(9218868437227405313)))
			if base.F64_ne(v12, v18) != 0 {
				v92 = v8
				return v92
			} else {
				if v61 == int32(0) {
					v92 = v8
					return v92
				} else {
					v67 = v22
					v71 = v25
					v72 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
					v75 = base.I64_reinterpret_f64(v72) & int64(9223372036854775807)
					if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v71) {
						return base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v75))
					} else {
						return base.B2i32(base.Ui64(v75) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v72, v67)
					}
				}
			}
		} else {
			if base.Ui64(int64(9218868437227405312)) < base.Ui64(v21) {
				v92 = v8
				return v92
			} else {
				v32 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
				if base.Ui64(base.I64_reinterpret_f64(v32)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
					if base.F64_ne(v12, v18) != 0 {
						if base.F64_le(base.F64_abs(base.F64_sub(v12, v18)), float64(1e-06)) == int32(0) {
							v92 = v8
						} else {
							v92 = base.F64_eq(v22, v32) | base.F64_le(base.F64_abs(base.F64_sub(v22, v32)), float64(1e-06))
						}
					} else {
						v92 = base.F64_eq(v22, v32) | base.F64_le(base.F64_abs(base.F64_sub(v22, v32)), float64(1e-06))
					}
					return v92
				} else {
					v61 = int32(1)
					if base.F64_ne(v12, v18) != 0 {
						v92 = v8
						return v92
					} else {
						if v61 == int32(0) {
							v92 = v8
							return v92
						} else {
							v67 = v22
							v71 = v25
							v72 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
							v75 = base.I64_reinterpret_f64(v72) & int64(9223372036854775807)
							if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v71) {
								return base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v75))
							} else {
								return base.B2i32(base.Ui64(v75) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v72, v67)
							}
						}
					}
				}
			}
		}
	} else {
		v39 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
		if base.Ui64(v39&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
			v92 = v8
			return v92
		} else {
			v44 = *(*float64)(unsafe.Add(mBase, uint32(v11)+8))
			v67 = v44
			v71 = base.I64_reinterpret_f64(v44) & int64(9223372036854775807)
			v72 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
			v75 = base.I64_reinterpret_f64(v72) & int64(9223372036854775807)
			if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v71) {
				return base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v75))
			} else {
				return base.B2i32(base.Ui64(v75) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v72, v67)
			}
		}
	}
}
func F_point_mul(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_palloc(m, int32(16))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		F_point_mul_point(m, v7, v5, v4)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			return v7
		}
	}
}
func F_point_recv(m *base.Module, l0 int32) int32 {
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
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_palloc(m, int32(16))
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
				return v5
			}
		}
	}
}
func F_write_point_as_box(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v34 float64
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v69 int32
	_ = v69
	var v73 float64
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v113 int32
	_ = v113
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v16 = l0<<(uint(int32(3))%32) + int32(8)
	v17 = F_palloc0(m, v16)
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
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = l0 | int32(-2147483648)
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v16 << (uint(int32(2)) % 32)
	v27 = int32(1)
	if l0 <= int32(0) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	m.G0 = v11 + int32(16)
	return v113
L4:
	;
	v47 = int32(44)
	v48 = F___strchrnul(m, l1, v47)
	mBase = m.M
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48))))
	if v50 == v47 {
		goto L13
	} else {
		goto L14
	}
L5:
	;
	v46 = int32(0)
	goto L4
L6:
	;
	goto L7
L7:
	;
	v34 = F_float8in_internal(m, l1, v11+int32(12), int32(420039), l1, l3)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v17)+8)) = v34
	if l3 == int32(0) {
		v46 = v27
		goto L4
	} else {
		goto L9
	}
L9:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v39 != int32(447) {
		v46 = v27
		goto L4
	} else {
		goto L10
	}
L10:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)))
	if v42 == int32(0) {
		v46 = v27
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v113 = int32(0)
	goto L3
L12:
	;
	if v54 != 0 {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	v54 = v48
	goto L15
L14:
	;
	v54 = int32(0)
	goto L15
L15:
	;
	goto L12
L16:
	;
	v57 = v54
	v61 = v46
	goto L19
L17:
	;
	goto L18
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v17
	v113 = int32(1)
	goto L3
L19:
	;
	v69 = v57 + int32(1)
	v73 = F_float8in_internal(m, v69, v11+int32(12), int32(420039), l1, l3)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L21
	}
L20:
	;
	goto L18
L21:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v17+int32(8)+v61<<(uint(int32(3))%32)))) = v73
	if l3 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v87 = int32(44)
	v88 = F___strchrnul(m, v69, v87)
	mBase = m.M
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
	if v90 == v87 {
		goto L27
	} else {
		goto L28
	}
L23:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v78 != int32(447) {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)))
	if v81 == int32(0) {
		goto L22
	} else {
		goto L25
	}
L25:
	;
	v113 = int32(0)
	goto L3
L26:
	;
	if v94 != 0 {
		v57 = v94
		v61 = v61 + int32(1)
		goto L19
	} else {
		goto L30
	}
L27:
	;
	v94 = v88
	goto L29
L28:
	;
	v94 = int32(0)
	goto L29
L29:
	;
	goto L26
L30:
	;
	goto L20
}
