package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_int2_sum(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v15 int64
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int64
	_ = v54
	var v55 int64
	_ = v55
	var v59 int64
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int64
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v5 == int32(1) {
		v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
		if v8 == int32(1) {
			v11 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v11)
			return int32(0)
		} else {
			v15 = int64(*(*int16)(unsafe.Add(mBase, uint32(l0)+28)))
			v16 = F_Int64GetDatum(m, v15)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				return v16
			}
		}
	} else {
		v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v23 == int32(0) {
			v51 = int32(0)
		} else {
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
			switch v26 - int32(429) {
			case 0:
				v51 = int32(1)
			case 1:
				v51 = int32(2)
			default:
				v51 = int32(0)
			}
		}
		v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v51 != 0 {
			v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
			if v53 != 0 {
				v70 = v52
				return v70
			} else {
				v54 = int64(*(*int16)(unsafe.Add(mBase, uint32(l0)+28)))
				v55 = *(*int64)(unsafe.Add(mBase, uint32(v52)))
				*(*int64)(unsafe.Add(mBase, uint32(v52))) = v54 + v55
				return v52
			}
		} else {
			v59 = *(*int64)(unsafe.Add(mBase, uint32(v52)))
			v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
			if v60 == int32(1) {
				v63 = F_Int64GetDatum(m, v59)
				mBase = m.M
				v64 = m.ExcPending
				if v64 != 0 {
					return int32(0)
				} else {
					return v63
				}
			} else {
				v66 = int64(*(*int16)(unsafe.Add(mBase, uint32(l0)+28)))
				v68 = F_Int64GetDatum(m, v66+v59)
				mBase = m.M
				v69 = m.ExcPending
				if v69 != 0 {
					return int32(0)
				} else {
					v70 = v68
					return v70
				}
			}
		}
	}
}
