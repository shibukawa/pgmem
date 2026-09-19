package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_shm_toc_allocate(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	v8 = int32(8)
	v9 = l0 + v8
	v12 = base.AtomicRmwXchg32(m, l0, v8, int32(1))
	if v12 != 0 {
		F_s_lock(m, v9, int32(_a_F_shm_toc_allocate_0), int32(104), int32(_a_F_shm_toc_allocate_1))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			v23 = (l1 + int32(31)) & int32(-32)
			v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v30 = v24 + v25<<(uint(int32(4))%32) + int32(24)
			v31 = v23 + v30
			v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			if base.B2i32(base.Ui32(v31) <= base.Ui32(v32))&base.B2i32(base.Ui32(v30) <= base.Ui32(v31)) == int32(0) {
				v38 = int32(0)
				atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v9))), uint32(v38))
				F_errstart_cold(m, int32(21), v38)
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(_a_F_shm_toc_allocate_2))
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_shm_toc_allocate_3), int32(0))
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_shm_toc_allocate_0), int32(118), int32(_a_F_shm_toc_allocate_1))
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v57 + v23
				v60 = int32(0)
				atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l0)+8)), uint32(v60))
				return l0 + (v32 - (v23 + v24))
			}
		}
	} else {
		v23 = (l1 + int32(31)) & int32(-32)
		v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v30 = v24 + v25<<(uint(int32(4))%32) + int32(24)
		v31 = v23 + v30
		v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if base.B2i32(base.Ui32(v31) <= base.Ui32(v32))&base.B2i32(base.Ui32(v30) <= base.Ui32(v31)) == int32(0) {
			v38 = int32(0)
			atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v9))), uint32(v38))
			F_errstart_cold(m, int32(21), v38)
			mBase = m.M
			v44 = m.ExcPending
			if v44 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(_a_F_shm_toc_allocate_2))
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_shm_toc_allocate_3), int32(0))
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_shm_toc_allocate_0), int32(118), int32(_a_F_shm_toc_allocate_1))
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v57 + v23
			v60 = int32(0)
			atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l0)+8)), uint32(v60))
			return l0 + (v32 - (v23 + v24))
		}
	}
}
