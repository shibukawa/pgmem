package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_on_ppath(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
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
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 float64
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v40 float64
	_ = v40
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 float64
	_ = v52
	var v53 int32
	_ = v53
	var v54 float64
	_ = v54
	var v64 int32
	_ = v64
	var v68 float64
	_ = v68
	var v69 int32
	_ = v69
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v96 int32
	_ = v96
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
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
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	if v19 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	F_float_overflow_error(m)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L26
	}
L4:
	;
	return v86
L5:
	;
	v23 = v18 - int32(1)
	v24 = int32(0)
	if v24 < v23 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v79 = F_point_inside(m, v12, v18, v14+int32(16))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L25
	}
L8:
	;
	v27 = v23
	goto L10
L9:
	;
	v27 = v24
	goto L10
L10:
	;
	v29 = v14 + int32(16)
	v30 = F_point_dt(m, v12, v29)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v33 = int32(0)
	v40 = v30
	goto L12
L12:
	;
	if v33 == v27 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v86 = v64
	goto L4
L14:
	;
	return int32(0)
L15:
	;
	goto L16
L16:
	;
	v48 = v33 + int32(1)
	v51 = v29 + v48<<(uint(int32(4))%32)
	v52 = F_point_dt(m, v12, v51)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L18
	}
L17:
	;
	v64 = int32(1)
	v68 = F_point_dt(m, v29+v33<<(uint(int32(4))%32), v51)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L22
	}
L18:
	;
	v54 = base.F64_add(v40, v52)
	if base.F64_ne(base.F64_abs(v54), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	if base.F64_eq(base.F64_abs(v40), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L17
	} else {
		goto L20
	}
L20:
	;
	if base.F64_ne(base.F64_abs(v52), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L3
	} else {
		goto L21
	}
L21:
	;
	goto L17
L22:
	;
	if base.F64_eq(v54, v68) != 0 {
		v86 = v64
		goto L4
	} else {
		goto L23
	}
L23:
	;
	if base.F64_le(base.F64_abs(base.F64_sub(v54, v68)), float64(1e-06)) == int32(0) {
		v33 = v48
		v40 = v52
		goto L12
	} else {
		goto L24
	}
L24:
	;
	goto L13
L25:
	;
	v86 = base.B2i32(v79 != int32(0))
	goto L4
L26:
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
	var v30 int32
	_ = v30
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	v5 = *(*int32)(unsafe.Add(mBase, _consts[769]))
	if v5 < int32(20) {
		v9 = v5 << (uint(int32(3)) % 32)
		v12 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v9)+uint32(_consts[770]))) = v12
		*(*int32)(unsafe.Add(mBase, uint32(v9)+uint32(_consts[771]))) = l0
		*(*int32)(unsafe.Add(mBase, _consts[769])) = v5 + int32(1)
		v22 = int32(*(*uint8)(unsafe.Add(mBase, _consts[772])))
		if v22 == v12 {
			v30 = *(*int32)(unsafe.Add(mBase, _consts[773]))
			if v30 <= int32(31) {
				*(*int32)(unsafe.Add(mBase, _consts[773])) = v30 + int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v30<<(uint(int32(2))%32))+uint32(_consts[774]))) = int32(1100)
			} else {
			}
			v45 = int32(1)
			*(*uint8)(unsafe.Add(mBase, _consts[772])) = uint8(v45)
		} else {
		}
		return
	} else {
		F_errstart_cold(m, int32(22), int32(0))
		mBase = m.M
		v50 = m.ExcPending
		if v50 != 0 {
			return
		} else {
			F_errcode(m, int32(261))
			mBase = m.M
			v53 = m.ExcPending
			if v53 != 0 {
				return
			} else {
				F_errmsg_internal(m, int32(125739), int32(0))
				mBase = m.M
				v57 = m.ExcPending
				if v57 != 0 {
					return
				} else {
					F_errfinish(m, int32(520543), int32(321), int32(105511))
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
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
	var v7 int32
	_ = v7
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
	var v36 float64
	_ = v36
	var v45 int32
	_ = v45
	v7 = int32(0)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v10 = *(*float64)(unsafe.Add(mBase, uint32(v9)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*float64)(unsafe.Add(mBase, uint32(v11)))
	if base.F64_ge(v10, v12) == v7 {
		v45 = v7
	} else {
		v16 = *(*float64)(unsafe.Add(mBase, uint32(v9)+16))
		if base.F64_le(v16, v12) == int32(0) {
			v45 = v7
		} else {
			v20 = *(*float64)(unsafe.Add(mBase, uint32(v9)+8))
			v21 = *(*float64)(unsafe.Add(mBase, uint32(v11)+8))
			if base.F64_ge(v20, v21) == int32(0) {
				v45 = v7
			} else {
				v25 = *(*float64)(unsafe.Add(mBase, uint32(v9)+24))
				if base.F64_le(v25, v21) == int32(0) {
					v45 = v7
				} else {
					v29 = *(*float64)(unsafe.Add(mBase, uint32(v11)+16))
					if base.F64_ge(v10, v29) == int32(0) {
						v45 = v7
					} else {
						if base.F64_ge(v29, v16) == int32(0) {
							v45 = v7
						} else {
							v36 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
							if base.F64_ge(v20, v36) == int32(0) {
								v45 = v7
							} else {
								v45 = base.F64_ge(v36, v25)
							}
						}
					}
				}
			}
		}
	}
	return v45
}
