package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_shm_mq_set_handle(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = l1
	return
}
func F_shm_toc_allocate(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
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
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(1)
	v12 = l0 + int32(8)
	v16 = (l1 + int32(31)) & int32(-32)
	if v8 != 0 {
		F_s_lock(m, v12, int32(500094), int32(104), int32(357559))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v30 = v24 + v25<<(uint(int32(4))%32) + int32(24)
			v31 = v30 + v16
			v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			if base.B2i32(base.Ui32(v31) <= base.Ui32(v32))&base.B2i32(base.Ui32(v30) <= base.Ui32(v31)) == int32(0) {
				v38 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v12))) = v38
				F_errstart_cold(m, int32(21), v38)
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(8389))
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(14090), int32(0))
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(500094), int32(118), int32(357559))
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
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
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
				v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v58 + v16
				return l0 + (v32 - (v16 + v24))
			}
		}
	} else {
		v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v30 = v24 + v25<<(uint(int32(4))%32) + int32(24)
		v31 = v30 + v16
		v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if base.B2i32(base.Ui32(v31) <= base.Ui32(v32))&base.B2i32(base.Ui32(v30) <= base.Ui32(v31)) == int32(0) {
			v38 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v12))) = v38
			F_errstart_cold(m, int32(21), v38)
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(8389))
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(14090), int32(0))
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(500094), int32(118), int32(357559))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
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
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
			v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v58 + v16
			return l0 + (v32 - (v16 + v24))
		}
	}
}
