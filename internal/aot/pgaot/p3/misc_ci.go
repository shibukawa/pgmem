package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_cidr_recv(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_network_recv(m, v2, int32(1))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(v4)
	}
}
func F_cidr_set_masklen_internal(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v67 int32
	_ = v67
	v2 = l1
	v8 = F_palloc0(m, int32(22))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v12 = int32(1)
		v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
		if v14&v12 != 0 {
			v17 = v12
		} else {
			v17 = int32(4)
		}
		v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v17))))
		v20 = int32(1)
		v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
		if v22&v20 != 0 {
			v25 = v20
		} else {
			v25 = int32(4)
		}
		v26 = v8 + v25
		*(*uint8)(unsafe.Add(mBase, uint32(v26)+1)) = uint8(v2)
		*(*uint8)(unsafe.Add(mBase, uint32(v26))) = uint8(v19)
		if v2 <= int32(0) {
		} else {
			v32 = v26 + int32(2)
			v36 = base.I32_div_s(v2+int32(7), int32(8))
			if v36 != 0 {
				v37 = int32(1)
				v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
				if v39&v37 != 0 {
					v42 = v37
				} else {
					v42 = int32(4)
				}
				base.MemoryCopy(m, v32, l0+v42+int32(2), v36)
			} else {
			}
			v48 = v2 & int32(7)
			if v48 == int32(0) {
			} else {
				v53 = v32 + int32(base.Ui32(v2)>>(uint(int32(3))%32))
				v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
				v57 = v54 & (int32(-256) >> (uint(v48) % 32))
				*(*uint8)(unsafe.Add(mBase, uint32(v53))) = uint8(v57)
			}
		}
		if v19 == int32(2) {
			v67 = int32(40)
		} else {
			v67 = int32(88)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v8))) = v67
		return v8
	}
}
