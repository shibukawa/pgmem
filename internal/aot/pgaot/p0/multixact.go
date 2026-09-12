package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_find_multixact_start(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int64
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	F_SimpleLruWriteAll(m, int32(4443476))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		F_SimpleLruWriteAll(m, int32(4443556))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v18 = int32(base.Ui32(l0) >> (uint(int32(11)) % 32))
			v19 = base.I64_extend_i32_u(v18)
			v20 = F_SimpleLruDoesPhysicalPageExist(m, int32(4443476), v19)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				if v20 != 0 {
					v23 = F_SimpleLruReadPage_ReadOnly(m, int32(4443476), v19, l0)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						v26 = *(*int32)(unsafe.Add(mBase, _consts[65]))
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
						v28 = int32(2)
						v31 = *(*int32)(unsafe.Add(mBase, uint32(v27+v23<<(uint(v28)%32))))
						v37 = *(*int32)(unsafe.Add(mBase, uint32(v31+l0&int32(2047)<<(uint(v28)%32))))
						v38 = *(*int32)(unsafe.Add(mBase, uint32(v26)+28))
						v40 = int32(*(*uint16)(unsafe.Add(mBase, _consts[66])))
						v41 = base.I32_rem_u_s(v18, v40)
						F_LWLockRelease(m, v38+v41<<(uint(int32(7))%32))
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l1))) = v37
							return v20
						}
					}
				} else {
					return v20
				}
			}
		}
	}
}
