package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_MultiXactMemberFreezeThreshold(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int64
	_ = v17
	var v18 int64
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int64
	_ = v27
	var v31 int32
	_ = v31
	var v37 float64
	_ = v37
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactMemberFreezeThreshold[0]))
	v11 = F_LWLockAcquire(m, v7+int32(1664), int32(1))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactMemberFreezeThreshold[1]))
		v17 = *(*int64)(unsafe.Add(mBase, uint32(v16)+32))
		v18 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v16)+20))
		v22 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactMemberFreezeThreshold[0]))
		F_LWLockRelease(m, v22+int32(1664))
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return int32(0)
		} else {
			v27 = v18 - v17
			if base.Ui64(v27) <= base.Ui64(int64(2000000000)) {
				v31 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactMemberFreezeThreshold[2]))
				return v31
			} else {
				v37 = base.F64_div(base.F64_convert_i64_u(v27-int64(2000000000)), float64(2e+09))
				if base.F64_ge(v37, float64(1)) != 0 {
					v52 = int32(0)
				} else {
					v41 = v19 - v20
					v45 = v41 - base.I32_trunc_sat_f64_u(base.F64_mul(v37, base.F64_convert_i32_u(v41)))
					v47 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactMemberFreezeThreshold[2]))
					if v45 < v47 {
						v49 = v45
					} else {
						v49 = v47
					}
					v52 = v49
				}
				return v52
			}
		}
	}
}
func F_MultiXactMemberPagePrecedes(m *base.Module, l0 int64, l1 int64) int32 {
	return base.B2i32(l0 < l1)
}
func F_MultiXactOffsetIoErrorDetail(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14249(m, l0, int32(_a_F_MultiXactOffsetIoErrorDetail_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_MultiXactSetNextMXact(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactSetNextMXact[0]))
	v8 = F_LWLockAcquire(m, v4+int32(1664), int32(0))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		v10 = int32(_a_F_MultiXactSetNextMXact_0)
		v11 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactSetNextMXact[1]))
		*(*int32)(unsafe.Add(mBase, uint32(v11))) = l0
		v14 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactSetNextMXact[1]))
		*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = l1
		v17 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactSetNextMXact[0]))
		F_LWLockRelease(m, v17+int32(1664))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return
		} else {
			return
		}
	}
}
func F_ReadMultiXactIdRange(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_ReadMultiXactIdRange[0]))
	v8 = F_LWLockAcquire(m, v4+int32(1664), int32(1))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		v10 = int32(_a_F_ReadMultiXactIdRange_0)
		v11 = *(*int32)(unsafe.Add(mBase, _c_F_ReadMultiXactIdRange[1]))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = v12
		v15 = *(*int32)(unsafe.Add(mBase, _c_F_ReadMultiXactIdRange[1]))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v16
		v19 = *(*int32)(unsafe.Add(mBase, _c_F_ReadMultiXactIdRange[0]))
		F_LWLockRelease(m, v19+int32(1664))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return
		} else {
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
