package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_tbm_create(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v27 int64
	_ = v27
	v6 = F_palloc0(m, int32(116))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(478)
		v13 = *(*int32)(unsafe.Add(mBase, _c_F_tbm_create[0]))
		*(*int32)(unsafe.Add(mBase, uint32(v6)+112)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v6)+36)) = int32(0)
		v17 = int32(16)
		v19 = base.I32_div_u_s(l0, int32(56))
		if base.Ui32(v19) <= base.Ui32(v17) {
			v22 = v17
		} else {
			v22 = v19
		}
		*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = v22
		*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v13
		v27 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v6)+96)) = v27
		*(*int64)(unsafe.Add(mBase, uint32(v6)+104)) = v27
		return v6
	}
}
func F_tbm_extract_page_tuple(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	v3 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v12 = v3
	v13 = v3
	for {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v7+int32(8)+v13<<(uint(int32(2))%32))))
		if v19 != 0 {
			v24 = v19
			v26 = v12
			v28 = v13<<(uint(int32(5))%32) | int32(1)
			for {
				if v24&int32(1) != 0 {
					if base.Ui32(v26) < base.Ui32(int32(291)) {
						*(*uint16)(unsafe.Add(mBase, uint32(l1+v26<<(uint(int32(1))%32)))) = uint16(v28)
					} else {
					}
					v40 = v26 + int32(1)
				} else {
					v40 = v26
				}
				v41 = int32(1)
				v44 = int32(base.Ui32(v24) >> (uint(v41) % 32))
				if v44 != 0 {
					v24 = v44
					v26 = v40
					v28 = v28 + v41
					continue
				} else {
					break
				}
				break
			}
			v47 = v40
		} else {
			v47 = v12
		}
		v52 = v13 + int32(1)
		if v52 != int32(10) {
			v12 = v47
			v13 = v52
			continue
		} else {
			break
		}
		break
	}
	return v47
}
