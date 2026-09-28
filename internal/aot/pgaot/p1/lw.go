package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_GetLWTrancheName(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	if base.Ui32(l0) <= base.Ui32(int32(99)) {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_c_F_GetLWTrancheName[0])))
		v49 = v14
		m.G0 = v8 + int32(16)
		return v49
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, _c_F_GetLWTrancheName[1]))
		v18 = l0 - int32(100)
		v20 = *(*int32)(unsafe.Add(mBase, _c_F_GetLWTrancheName[2]))
		if v20 <= v18 {
			v24 = base.AtomicRmwXchg32(m, v16, int32(_a_F_GetLWTrancheName_0), int32(1))
			if v24 != 0 {
				F_s_lock(m, v16+int32(_a_F_GetLWTrancheName_0), int32(_a_F_GetLWTrancheName_1))
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int32(0)
				} else {
					v34 = *(*int32)(unsafe.Add(mBase, _c_F_GetLWTrancheName[1]))
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_GetLWTrancheName[3])))
					*(*int32)(unsafe.Add(mBase, _c_F_GetLWTrancheName[2])) = v35
					v37 = int32(0)
					atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_GetLWTrancheName[4]))), uint32(v37))
					if v35 <= v18 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
							F_errmsg_internal(m, int32(_a_F_GetLWTrancheName_2), v8)
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_GetLWTrancheName_3), int32(737), int32(_a_F_GetLWTrancheName_4))
								mBase = m.M
								v66 = m.ExcPending
								if v66 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v41 = v34
						v49 = v41 + v18*int32(68)
						m.G0 = v8 + int32(16)
						return v49
					}
				}
			} else {
				v34 = *(*int32)(unsafe.Add(mBase, _c_F_GetLWTrancheName[1]))
				v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_GetLWTrancheName[3])))
				*(*int32)(unsafe.Add(mBase, _c_F_GetLWTrancheName[2])) = v35
				v37 = int32(0)
				atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_GetLWTrancheName[4]))), uint32(v37))
				if v35 <= v18 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
						F_errmsg_internal(m, int32(_a_F_GetLWTrancheName_2), v8)
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_GetLWTrancheName_3), int32(737), int32(_a_F_GetLWTrancheName_4))
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v41 = v34
					v49 = v41 + v18*int32(68)
					m.G0 = v8 + int32(16)
					return v49
				}
			}
		} else {
			v41 = v16
			v49 = v41 + v18*int32(68)
			m.G0 = v8 + int32(16)
			return v49
		}
	}
}
func F_LWLockReleaseClearVar(m *base.Module, l0 int32, l1 int32) {
	var v5 int64
	_ = v5
	var v7 int32
	_ = v7
	v5 = base.AtomicRmwXchg64(m, l1, int32(0), int64(0))
	F_LWLockRelease(m, l0)
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		return
	}
}
