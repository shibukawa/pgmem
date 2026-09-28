package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_BufferLockUnlock(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v46 int64
	_ = v46
	var v49 int64
	_ = v49
	var v51 int64
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_BufferLockUnlock[0]))
	if v11 != int32(-1) {
		v15 = v11 << (uint(int32(4)) % 32)
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v15)+uint32(_c_F_BufferLockUnlock[1])))
		if v18 == l0 {
			v37 = v15 + int32(_a_F_BufferLockUnlock_0)
			v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
			*(*int32)(unsafe.Add(mBase, uint32(v37)+12)) = int32(0)
			if v38 == int32(2) {
				v46 = int64(4503599627370496)
			} else {
				v46 = int64(17179869184)
			}
			if v38 == int32(3) {
				v49 = int64(9007199254740992)
			} else {
				v49 = v46
			}
			v51 = base.AtomicRmwSub64(m, l1, int32(24), v49)
			F_BufferLockProcessRelease(m, l1, v38, v51-v49)
			mBase = m.M
			v54 = m.ExcPending
			if v54 != 0 {
				return
			} else {
				v55 = int32(_a_F_BufferLockUnlock_1)
				v57 = *(*int32)(unsafe.Add(mBase, _c_F_BufferLockUnlock[2]))
				*(*int32)(unsafe.Add(mBase, _c_F_BufferLockUnlock[2])) = v57 - int32(1)
				m.G0 = v8 + int32(16)
				return
			}
		} else {
			v22 = F_GetPrivateRefCountEntrySlow(m, l0, int32(0))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return
			} else {
				if v22 != 0 {
					v37 = v22
					v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v37)+12)) = int32(0)
					if v38 == int32(2) {
						v46 = int64(4503599627370496)
					} else {
						v46 = int64(17179869184)
					}
					if v38 == int32(3) {
						v49 = int64(9007199254740992)
					} else {
						v49 = v46
					}
					v51 = base.AtomicRmwSub64(m, l1, int32(24), v49)
					F_BufferLockProcessRelease(m, l1, v38, v51-v49)
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return
					} else {
						v55 = int32(_a_F_BufferLockUnlock_1)
						v57 = *(*int32)(unsafe.Add(mBase, _c_F_BufferLockUnlock[2]))
						*(*int32)(unsafe.Add(mBase, _c_F_BufferLockUnlock[2])) = v57 - int32(1)
						m.G0 = v8 + int32(16)
						return
					}
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
						F_errmsg_internal(m, int32(_a_F_BufferLockUnlock_2), v8)
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_BufferLockUnlock_3), int32(_a_F_BufferLockUnlock_4), int32(_a_F_BufferLockUnlock_5))
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
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
	} else {
		v22 = F_GetPrivateRefCountEntrySlow(m, l0, int32(0))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return
		} else {
			if v22 != 0 {
				v37 = v22
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
				*(*int32)(unsafe.Add(mBase, uint32(v37)+12)) = int32(0)
				if v38 == int32(2) {
					v46 = int64(4503599627370496)
				} else {
					v46 = int64(17179869184)
				}
				if v38 == int32(3) {
					v49 = int64(9007199254740992)
				} else {
					v49 = v46
				}
				v51 = base.AtomicRmwSub64(m, l1, int32(24), v49)
				F_BufferLockProcessRelease(m, l1, v38, v51-v49)
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
					return
				} else {
					v55 = int32(_a_F_BufferLockUnlock_1)
					v57 = *(*int32)(unsafe.Add(mBase, _c_F_BufferLockUnlock[2]))
					*(*int32)(unsafe.Add(mBase, _c_F_BufferLockUnlock[2])) = v57 - int32(1)
					m.G0 = v8 + int32(16)
					return
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
					F_errmsg_internal(m, int32(_a_F_BufferLockUnlock_2), v8)
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_BufferLockUnlock_3), int32(_a_F_BufferLockUnlock_4), int32(_a_F_BufferLockUnlock_5))
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
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
}
