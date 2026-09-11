package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_int4_numeric(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v28 int64
	_ = v28
	var v39 int64
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v48 int64
	_ = v48
	var v51 int32
	_ = v51
	var v53 int64
	_ = v53
	var v56 int64
	_ = v56
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = int64(0)
	v17 = F_palloc(m, int32(12))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v17
		v22 = int32(0)
		*(*uint16)(unsafe.Add(mBase, uint32(v17))) = uint16(v22)
		*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v17 + int32(2)
		v28 = base.I64_extend_i32_s(v13)
		if v13 < v22 {
			*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = int64(16384)
			v39 = int64(0) - v28
			v42 = v22
			v44 = v17 + int32(12)
			v48 = v39
			for {
				v51 = v44 - int32(2)
				v53 = base.I64_div_u_s(v48, int64(10000))
				v56 = v53*int64(55536) + v48
				*(*uint16)(unsafe.Add(mBase, uint32(v51))) = uint16(v56)
				v59 = v42 + int32(1)
				if base.Ui64(int64(9999)) < base.Ui64(v48) {
					v42 = v59
					v44 = v51
					v48 = v53
					continue
				} else {
					break
				}
				break
			}
			*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v51
			v63 = v59
			v67 = v42
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = int64(0)
			if v13 == int32(0) {
				v63 = v22
				v67 = int32(0)
			} else {
				v39 = v28
				v42 = v22
				v44 = v17 + int32(12)
				v48 = v39
				for {
					v51 = v44 - int32(2)
					v53 = base.I64_div_u_s(v48, int64(10000))
					v56 = v53*int64(55536) + v48
					*(*uint16)(unsafe.Add(mBase, uint32(v51))) = uint16(v56)
					v59 = v42 + int32(1)
					if base.Ui64(int64(9999)) < base.Ui64(v48) {
						v42 = v59
						v44 = v51
						v48 = v53
						continue
					} else {
						break
					}
					break
				}
				*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v51
				v63 = v59
				v67 = v42
			}
		}
		*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v67
		*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v63
		v76 = F_make_result_opt_error(m, v11+int32(8), int32(0))
		mBase = m.M
		v77 = m.ExcPending
		if v77 != 0 {
			return int32(0)
		} else {
			F_pfree(m, v17)
			mBase = m.M
			v79 = m.ExcPending
			if v79 != 0 {
				return int32(0)
			} else {
				m.G0 = v11 + int32(32)
				return v76
			}
		}
	}
}
