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
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	v8 = int32(8)
	v9 = l0 + v8
	v12 = base.AtomicRmwXchg32(m, l0, v8, int32(1))
	if v12 != 0 {
		F_s_lock(m, v9, int32(_a_F_shm_toc_allocate_0))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v21 = (l1 + int32(31)) & int32(-32)
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v28 = v22 + v23<<(uint(int32(4))%32) + int32(24)
			v29 = v21 + v28
			v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			if base.B2i32(base.Ui32(v29) <= base.Ui32(v30))&base.B2i32(base.Ui32(v28) <= base.Ui32(v29)) == int32(0) {
				v36 = int32(0)
				atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v9))), uint32(v36))
				F_errstart_cold(m, int32(21), v36)
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(_a_F_shm_toc_allocate_1))
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_shm_toc_allocate_2), int32(0))
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_shm_toc_allocate_3), int32(118), int32(_a_F_shm_toc_allocate_4))
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
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
				v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v55 + v21
				v58 = int32(0)
				atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l0)+8)), uint32(v58))
				return l0 + (v30 - (v21 + v22))
			}
		}
	} else {
		v21 = (l1 + int32(31)) & int32(-32)
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v28 = v22 + v23<<(uint(int32(4))%32) + int32(24)
		v29 = v21 + v28
		v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if base.B2i32(base.Ui32(v29) <= base.Ui32(v30))&base.B2i32(base.Ui32(v28) <= base.Ui32(v29)) == int32(0) {
			v36 = int32(0)
			atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v9))), uint32(v36))
			F_errstart_cold(m, int32(21), v36)
			mBase = m.M
			v42 = m.ExcPending
			if v42 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(_a_F_shm_toc_allocate_1))
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_shm_toc_allocate_2), int32(0))
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_shm_toc_allocate_3), int32(118), int32(_a_F_shm_toc_allocate_4))
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
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
			v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v55 + v21
			v58 = int32(0)
			atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l0)+8)), uint32(v58))
			return l0 + (v30 - (v21 + v22))
		}
	}
}
