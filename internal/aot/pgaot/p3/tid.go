package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_TidStoreCreateShared(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	v8 = F_palloc0(m, int32(12))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v14 = int32(134217728)
		for {
			if base.Ui32(l0) < base.Ui32(v14<<(uint(int32(3))%32)) {
				v14 = int32(base.Ui32(v14) >> (uint(int32(1)) % 32))
				continue
			} else {
				break
			}
			break
		}
		v24 = int32(_a_F_TidStoreCreateShared_0)
		if base.Ui32(v14) <= base.Ui32(v24) {
			v27 = v24
		} else {
			v27 = v14
		}
		if base.Ui32(int32(_a_F_TidStoreCreateShared_1)) <= base.Ui32(v27) {
			v30 = int32(_a_F_TidStoreCreateShared_1)
		} else {
			v30 = v27
		}
		v31 = F_dsa_create_ext(m, int32(97), v30, v27)
		mBase = m.M
		v32 = m.ExcPending
		if v32 != 0 {
			return int32(0)
		} else {
			v34 = F_palloc0(m, int32(8))
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v34)+4)) = v31
				v39 = F_dsa_allocate_extended(m, v31, int32(56), int32(4))
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return int32(0)
				} else {
					v41 = F_dsa_get_address(m, v31, v39)
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v34))) = v41
						*(*int32)(unsafe.Add(mBase, uint32(v41))) = v39
						v45 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
						*(*int32)(unsafe.Add(mBase, uint32(v45)+4)) = int32(1420067175)
						v48 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
						F_LWLockInitialize(m, v48+int32(8), int32(97))
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return int32(0)
						} else {
							v54 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
							v57 = F_dsa_allocate_extended(m, v54, int32(24), int32(0))
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return int32(0)
							} else {
								v59 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
								v60 = F_dsa_get_address(m, v59, v57)
								mBase = m.M
								v61 = m.ExcPending
								if v61 != 0 {
									return int32(0)
								} else {
									*(*int64)(unsafe.Add(mBase, uint32(v60))) = int64(1024)
									v64 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
									*(*int32)(unsafe.Add(mBase, uint32(v64)+24)) = v57
									v66 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
									*(*int32)(unsafe.Add(mBase, uint32(v66)+48)) = int32(0)
									v69 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
									*(*int64)(unsafe.Add(mBase, uint32(v69)+32)) = int64(255)
									*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v31
									*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v34
									return v8
								}
							}
						}
					}
				}
			}
		}
	}
}
