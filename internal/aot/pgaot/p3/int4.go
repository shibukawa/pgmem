package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_int4_numeric(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int64
	_ = v13
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v28 int64
	_ = v28
	var v35 int64
	_ = v35
	var v41 int64
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int64
	_ = v50
	var v53 int32
	_ = v53
	var v55 int64
	_ = v55
	var v58 int64
	_ = v58
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v13 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = int64(0)
	v17 = F_palloc(m, int32(12))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int64(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v17
		v22 = int32(0)
		*(*uint16)(unsafe.Add(mBase, uint32(v17))) = uint16(v22)
		*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v17 + int32(2)
		v28 = base.I64_extend32_s(v13)
		if v28 < int64(0) {
			*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = int64(16384)
			v41 = int64(0) - v28
			v44 = v22
			v47 = v17 + int32(12)
			v50 = v41
			for {
				v53 = v47 - int32(2)
				v55 = base.I64_div_u_s(v50, int64(10000))
				v58 = v55*int64(55536) + v50
				*(*uint16)(unsafe.Add(mBase, uint32(v53))) = uint16(v58)
				v61 = v44 + int32(1)
				if base.Ui64(int64(9999)) < base.Ui64(v50) {
					v44 = v61
					v47 = v53
					v50 = v55
					continue
				} else {
					break
				}
				break
			}
			*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v53
			v65 = v61
			v69 = v44
		} else {
			v35 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = v35
			if v13<<(uint(int64(32))%64) == v35 {
				v65 = v22
				v69 = int32(0)
			} else {
				v41 = v28
				v44 = v22
				v47 = v17 + int32(12)
				v50 = v41
				for {
					v53 = v47 - int32(2)
					v55 = base.I64_div_u_s(v50, int64(10000))
					v58 = v55*int64(55536) + v50
					*(*uint16)(unsafe.Add(mBase, uint32(v53))) = uint16(v58)
					v61 = v44 + int32(1)
					if base.Ui64(int64(9999)) < base.Ui64(v50) {
						v44 = v61
						v47 = v53
						v50 = v55
						continue
					} else {
						break
					}
					break
				}
				*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v53
				v65 = v61
				v69 = v44
			}
		}
		*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v69
		*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v65
		v78 = F_make_result_safe(m, v11+int32(8), int32(0))
		mBase = m.M
		v79 = m.ExcPending
		if v79 != 0 {
			return int64(0)
		} else {
			F_pfree(m, v17)
			mBase = m.M
			v81 = m.ExcPending
			if v81 != 0 {
				return int64(0)
			} else {
				m.G0 = v11 + int32(32)
				return base.I64_extend_i32_u(v78)
			}
		}
	}
}
