package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_makeStringInfoExt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	v5 = F_palloc(m, int32(16))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = F_palloc(m, l0)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v5)+8)) = l0
			*(*int32)(unsafe.Add(mBase, uint32(v5))) = v9
			v13 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v9))) = uint8(v13)
			*(*int32)(unsafe.Add(mBase, uint32(v5)+12)) = v13
			*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = v13
			return v5
		}
	}
}
func F_resetStringInfo(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v3 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2))) = uint8(v3)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v3
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v3
	return
}
func F_string_agg_finalfn(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v6 != 0 {
		v30 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v30)
		return int32(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v7 == int32(0) {
			v30 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v30)
			return int32(0)
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
			v11 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
			v13 = v11 - v12
			v15 = v13 + int32(4)
			v16 = F_palloc(m, v15)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v16))) = v15 << (uint(int32(2)) % 32)
				if v13 != 0 {
					v26 = F__emscripten_memcpy_bulkmem(m, v16+int32(4), v12+v10, v13)
					mBase = m.M
				} else {
				}
				return v16
			}
		}
	}
}
func F_string_agg_serialize(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_pq_begintypsend(m, v7)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
		F_enlargeStringInfo(m, v7, int32(4))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
			v21 = int32(24)
			v23 = int32(65280)
			v25 = int32(8)
			*(*int32)(unsafe.Add(mBase, uint32(v18+v19))) = v14<<(uint(v21)%32) | v14&v23<<(uint(v25)%32) | (int32(base.Ui32(v14)>>(uint(v25)%32))&v23 | int32(base.Ui32(v14)>>(uint(v21)%32)))
			*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v18 + int32(4)
			v40 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
			v41 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
			F_pq_sendbytes(m, v7, v40, v41)
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return int32(0)
			} else {
				v45 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
				v46 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v45))) = v46 << (uint(int32(2)) % 32)
				m.G0 = v7 + int32(16)
				return v45
			}
		}
	}
}
