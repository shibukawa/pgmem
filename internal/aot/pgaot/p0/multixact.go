package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_find_multixact_start(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int64
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
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
	var v45 int64
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = l0
	F_SimpleLruWriteAll(m, int32(_a_F_find_multixact_start_0))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		F_SimpleLruWriteAll(m, int32(_a_F_find_multixact_start_1))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			v24 = int32(base.Ui32(l0) >> (uint(int32(10)) % 32))
			v25 = base.I64_extend_i32_u(v24)
			v26 = F_SimpleLruDoesPhysicalPageExist(m, int32(_a_F_find_multixact_start_0), v25)
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				if v26 != 0 {
					v31 = F_SimpleLruReadPage_ReadOnly(m, int32(_a_F_find_multixact_start_0), v25, v11+int32(12))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						v34 = *(*int32)(unsafe.Add(mBase, _c_F_find_multixact_start[0]))
						v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
						v39 = *(*int32)(unsafe.Add(mBase, uint32(v35+v31<<(uint(int32(2))%32))))
						v45 = *(*int64)(unsafe.Add(mBase, uint32(v39+l0&int32(1023)<<(uint(int32(3))%32))))
						v46 = *(*int32)(unsafe.Add(mBase, uint32(v34)+28))
						v48 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_find_multixact_start[1])))
						v49 = base.I32_rem_u_s(v24, v48)
						F_LWLockRelease(m, v46+v49<<(uint(int32(7))%32))
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return int32(0)
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(l1))) = v45
							m.G0 = v11 + int32(16)
							return v26
						}
					}
				} else {
					m.G0 = v11 + int32(16)
					return v26
				}
			}
		}
	}
}
