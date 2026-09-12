package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_MultiXactIdPrecedes(m *base.Module, l0 int32, l1 int32) int32 {
	return int32(base.Ui32(l0-l1) >> (uint(int32(31)) % 32))
}
func F_MultiXactMemberFreezeThreshold(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
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
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v45 float64
	_ = v45
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	v1 = int32(0)
	v9 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v13 = F_LWLockAcquire(m, v9+int32(1664), int32(1))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, _consts[78]))
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
		v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+24)))
		v25 = *(*int32)(unsafe.Add(mBase, _consts[7]))
		F_LWLockRelease(m, v25+int32(1664))
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return int32(0)
		} else {
			if v23 != int32(1) {
				v61 = v1
				return v61
			} else {
				v32 = v22 - v19
				if int32(0) <= v32 {
					v36 = *(*int32)(unsafe.Add(mBase, _consts[53]))
					return v36
				} else {
					v43 = v20 - v21
					v45 = base.F64_mul(base.F64_div(base.F64_convert_i32_u(v32-int32(2147483647)), float64(1.073741825e+09)), base.F64_convert_i32_u(v43))
					if base.F64_lt(v45, float64(4.294967296e+09))&base.F64_ge(v45, float64(0)) != 0 {
						v51 = base.I32_trunc_f64_u(v45)
						v53 = v51
					} else {
						v53 = int32(0)
					}
					if base.Ui32(v43) < base.Ui32(v53) {
						v61 = v1
					} else {
						v55 = v43 - v53
						v57 = *(*int32)(unsafe.Add(mBase, _consts[53]))
						if v55 < v57 {
							v59 = v55
						} else {
							v59 = v57
						}
						v61 = v59
					}
					return v61
				}
			}
		}
	}
}
func F_MultiXactMemberPagePrecedes(m *base.Module, l0 int64, l1 int64) int32 {
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	v5 = int32(1636)
	v10 = base.I32_wrap_i64(l0)*v5 - base.I32_wrap_i64(l1)*v5
	return int32(base.Ui32(v10&(v10-int32(1635))) >> (uint(int32(31)) % 32))
}
func F_MultiXactSetNextMXact(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int64
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	v6 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v10 = F_LWLockAcquire(m, v6+int32(1664), int32(0))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, _consts[78]))
		*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v13))) = l0
		v17 = *(*int32)(unsafe.Add(mBase, _consts[7]))
		F_LWLockRelease(m, v17+int32(1664))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return
		} else {
			v23 = int32(*(*uint8)(unsafe.Add(mBase, _consts[79])))
			if v23 == int32(1) {
				v27 = *(*int32)(unsafe.Add(mBase, _consts[80]))
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+28))
				v30 = *(*int32)(unsafe.Add(mBase, _consts[78]))
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
				v33 = int32(base.Ui32(v31) >> (uint(int32(11)) % 32))
				v35 = int32(*(*uint16)(unsafe.Add(mBase, _consts[81])))
				v36 = base.I32_rem_u_s(v33, v35)
				v39 = v28 + v36<<(uint(int32(7))%32)
				v41 = F_LWLockAcquire(m, v39, int32(0))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return
				} else {
					v44 = base.I64_extend_i32_u(v33)
					v45 = F_SimpleLruDoesPhysicalPageExist(m, int32(4410100), v44)
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return
					} else {
						if v45 == int32(0) {
							v49 = int32(4410100)
							v51 = F_SimpleLruZeroPage(m, v49, v44)
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return
							} else {
								F_SimpleLruWritePage(m, v49, v51)
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return
								} else {
									F_LWLockRelease(m, v39)
									mBase = m.M
									v56 = m.ExcPending
									if v56 != 0 {
										return
									} else {
										return
									}
								}
							}
						} else {
							F_LWLockRelease(m, v39)
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return
							} else {
								return
							}
						}
					}
				}
			} else {
				return
			}
		}
	}
}
func F_multi_sort_init(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = F_palloc0(m, l0*int32(36)+int32(4))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
		return v7
	}
}
