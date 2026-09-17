package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_on_ppath(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 float64
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v39 float64
	_ = v39
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 float64
	_ = v53
	var v54 int32
	_ = v54
	var v55 float64
	_ = v55
	var v57 float64
	_ = v57
	var v66 int32
	_ = v66
	var v70 float64
	_ = v70
	var v71 int32
	_ = v71
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v89 int32
	_ = v89
	var v97 int32
	_ = v97
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v13 = F_pg_detoast_datum(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	if v18 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	F_float_overflow_error(m)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L23
	}
L4:
	;
	return v89
L5:
	;
	v22 = v17 - int32(1)
	v23 = int32(0)
	if v23 < v22 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v81 = F_point_inside(m, v11, v17, v13+int32(16))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L22
	}
L8:
	;
	v26 = v22
	goto L10
L9:
	;
	v26 = v23
	goto L10
L10:
	;
	v28 = v13 + int32(16)
	v29 = F_point_dt(m, v11, v28)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v32 = int32(0)
	v39 = v29
	goto L12
L12:
	;
	if v32 == v26 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v89 = v66
	goto L4
L14:
	;
	return int32(0)
L15:
	;
	goto L16
L16:
	;
	v49 = v32 + int32(1)
	v52 = v28 + v49<<(uint(int32(4))%32)
	v53 = F_point_dt(m, v11, v52)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v55 = base.F64_add(v39, v53)
	v57 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.B2i32(base.F64_eq(base.F64_abs(v39), math.Float64frombits(uint64(0x7ff0000000000000)))|base.F64_ne(base.F64_abs(v55), v57) == int32(0))&base.F64_ne(base.F64_abs(v53), v57) != 0 {
		goto L3
	} else {
		goto L18
	}
L18:
	;
	v66 = int32(1)
	v70 = F_point_dt(m, v28+v32<<(uint(int32(4))%32), v52)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	if base.F64_eq(v55, v70) != 0 {
		v89 = v66
		goto L4
	} else {
		goto L20
	}
L20:
	;
	if base.F64_le(base.F64_abs(base.F64_sub(v55, v70)), float64(1e-06)) == int32(0) {
		v32 = v49
		v39 = v53
		goto L12
	} else {
		goto L21
	}
L21:
	;
	goto L13
L22:
	;
	v89 = base.B2i32(v81 != int32(0))
	goto L4
L23:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_on_proc_exit(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_on_proc_exit[0]))
	if v5 < int32(20) {
		v9 = v5 << (uint(int32(3)) % 32)
		v12 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v9)+uint32(_c_F_on_proc_exit[1]))) = v12
		*(*int32)(unsafe.Add(mBase, uint32(v9)+uint32(_c_F_on_proc_exit[2]))) = l0
		*(*int32)(unsafe.Add(mBase, _c_F_on_proc_exit[0])) = v5 + int32(1)
		v22 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_on_proc_exit[3])))
		if v22 == v12 {
			v27 = *(*int32)(unsafe.Add(mBase, _c_F_on_proc_exit[4]))
			if v27 <= int32(31) {
				*(*int32)(unsafe.Add(mBase, _c_F_on_proc_exit[4])) = v27 + int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v27<<(uint(int32(2))%32))+uint32(_c_F_on_proc_exit[5]))) = int32(1103)
			} else {
			}
			v44 = int32(1)
			*(*uint8)(unsafe.Add(mBase, _c_F_on_proc_exit[3])) = uint8(v44)
		} else {
		}
		return
	} else {
		F_errstart_cold(m, int32(22), int32(0))
		mBase = m.M
		v49 = m.ExcPending
		if v49 != 0 {
			return
		} else {
			F_errcode(m, int32(261))
			mBase = m.M
			v52 = m.ExcPending
			if v52 != 0 {
				return
			} else {
				F_errmsg_internal(m, int32(_a_F_on_proc_exit_0), int32(0))
				mBase = m.M
				v56 = m.ExcPending
				if v56 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_on_proc_exit_1), int32(321), int32(_a_F_on_proc_exit_2))
					mBase = m.M
					v61 = m.ExcPending
					if v61 != 0 {
						return
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
func F_on_ps(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 float64
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 float64
	_ = v13
	var v14 int32
	_ = v14
	var v15 float64
	_ = v15
	var v16 float64
	_ = v16
	var v17 int32
	_ = v17
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = F_point_dt(m, v5, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v12 = v6 + int32(16)
		v13 = F_point_dt(m, v5, v12)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			v15 = base.F64_add(v7, v13)
			v16 = F_point_dt(m, v6, v12)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				return base.F64_eq(v15, v16) | base.F64_le(base.F64_abs(base.F64_sub(v15, v16)), float64(1e-06))
			}
		}
	}
}
func F_on_sb(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v10 float64
	_ = v10
	var v11 int32
	_ = v11
	var v12 float64
	_ = v12
	var v16 float64
	_ = v16
	var v20 float64
	_ = v20
	var v21 float64
	_ = v21
	var v25 float64
	_ = v25
	var v29 float64
	_ = v29
	var v31 int32
	_ = v31
	var v37 float64
	_ = v37
	var v42 int32
	_ = v42
	v2 = int32(0)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v10 = *(*float64)(unsafe.Add(mBase, uint32(v9)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*float64)(unsafe.Add(mBase, uint32(v11)))
	if base.F64_ge(v10, v12) == v2 {
		v42 = v2
	} else {
		v16 = *(*float64)(unsafe.Add(mBase, uint32(v9)+16))
		if base.F64_ge(v12, v16) == int32(0) {
			v42 = v2
		} else {
			v20 = *(*float64)(unsafe.Add(mBase, uint32(v9)+8))
			v21 = *(*float64)(unsafe.Add(mBase, uint32(v11)+8))
			if base.F64_ge(v20, v21) == int32(0) {
				v42 = v2
			} else {
				v25 = *(*float64)(unsafe.Add(mBase, uint32(v9)+24))
				if base.F64_ge(v21, v25) == int32(0) {
					v42 = v2
				} else {
					v29 = *(*float64)(unsafe.Add(mBase, uint32(v11)+16))
					v31 = int32(0)
					if base.B2i32(base.F64_ge(v10, v29) == v31)|base.B2i32(base.F64_le(v16, v29) == v31) != 0 {
						v42 = v2
					} else {
						v37 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
						if base.F64_ge(v20, v37) == int32(0) {
							v42 = v2
						} else {
							v42 = base.F64_ge(v37, v25)
						}
					}
				}
			}
		}
	}
	return v42
}
