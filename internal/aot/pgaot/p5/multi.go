package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_MultiXactIdPrecedes(m *base.Module, l0 int32, l1 int32) int32 {
	return int32(base.Ui32(l0-l1) >> (uint(int32(31)) % 32))
}
func F_MultiXactIdPrecedesOrEquals(m *base.Module, l0 int32, l1 int32) int32 {
	return base.B2i32(l0-l1 <= int32(0))
}
func F_MultiXactMemberFreezeThreshold(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
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
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	v1 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactMemberFreezeThreshold[0]))
	v12 = F_LWLockAcquire(m, v8+int32(1664), int32(1))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactMemberFreezeThreshold[1]))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
		v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+24)))
		v24 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactMemberFreezeThreshold[0]))
		F_LWLockRelease(m, v24+int32(1664))
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return int32(0)
		} else {
			if v22 != int32(1) {
				v54 = v1
				return v54
			} else {
				v31 = v21 - v18
				if int32(0) <= v31 {
					v35 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactMemberFreezeThreshold[2]))
					return v35
				} else {
					v37 = v19 - v20
					v45 = base.I32_trunc_sat_f64_u(base.F64_mul(base.F64_div(base.F64_convert_i32_u(v31-int32(2147483647)), float64(1.073741825e+09)), base.F64_convert_i32_u(v37)))
					if base.Ui32(v37) < base.Ui32(v45) {
						v54 = v1
					} else {
						v47 = v37 - v45
						v49 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactMemberFreezeThreshold[2]))
						if v47 < v49 {
							v51 = v47
						} else {
							v51 = v49
						}
						v54 = v51
					}
					return v54
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
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactSetNextMXact[0]))
	v10 = F_LWLockAcquire(m, v6+int32(1664), int32(0))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactSetNextMXact[1]))
		*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v13))) = l0
		v17 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactSetNextMXact[0]))
		F_LWLockRelease(m, v17+int32(1664))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return
		} else {
			v23 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_MultiXactSetNextMXact[2])))
			if v23 == int32(1) {
				v27 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactSetNextMXact[3]))
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+28))
				v30 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactSetNextMXact[1]))
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
				v33 = int32(base.Ui32(v31) >> (uint(int32(11)) % 32))
				v35 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_MultiXactSetNextMXact[4])))
				v36 = base.I32_rem_u_s(v33, v35)
				v39 = v28 + v36<<(uint(int32(7))%32)
				v41 = F_LWLockAcquire(m, v39, int32(0))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return
				} else {
					v44 = base.I64_extend_i32_u(v33)
					v45 = F_SimpleLruDoesPhysicalPageExist(m, int32(_a_F_MultiXactSetNextMXact_0), v44)
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return
					} else {
						if v45 == int32(0) {
							v49 = int32(_a_F_MultiXactSetNextMXact_0)
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
func F_ReadMultiXactIdRange(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_ReadMultiXactIdRange[0]))
	v9 = F_LWLockAcquire(m, v5+int32(1664), int32(1))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, _c_F_ReadMultiXactIdRange[1]))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = v13
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v15
		v18 = *(*int32)(unsafe.Add(mBase, _c_F_ReadMultiXactIdRange[0]))
		F_LWLockRelease(m, v18+int32(1664))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return
		} else {
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			if v23 == int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1)
			} else {
			}
			v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			if v28 == int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(1)
			} else {
			}
			return
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
