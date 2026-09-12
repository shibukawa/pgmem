package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_cidr_recv(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_network_recv(m, v2, int32(1))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
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
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	v2 = l1
	v8 = F_palloc0(m, int32(22))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v12 = int32(1)
		v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
		v16 = v14 & v12
		if v16 != 0 {
			v17 = v12
		} else {
			v17 = int32(4)
		}
		v19 = int32(1)
		v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
		if v21&v19 != 0 {
			v24 = v19
		} else {
			v24 = int32(4)
		}
		v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v24))))
		*(*uint8)(unsafe.Add(mBase, uint32(v8+v17))) = uint8(v26)
		v29 = v8 + int32(1)
		v31 = v8 + int32(4)
		if v16 != 0 {
			v32 = v29
		} else {
			v32 = v31
		}
		*(*uint8)(unsafe.Add(mBase, uint32(v32)+1)) = uint8(v2)
		if v2 <= int32(0) {
		} else {
			v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
			if v36&int32(1) != 0 {
				v39 = v29
			} else {
				v39 = v31
			}
			v42 = int32(1)
			v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
			if v46&v42 != 0 {
				v49 = l0 + v42
			} else {
				v49 = l0 + int32(4)
			}
			v55 = base.I32_div_s(v2+int32(7), int32(8))
			if v55 != 0 {
				v56 = F__emscripten_memcpy_bulkmem(m, v39+int32(2), v49+int32(2), v55)
				mBase = m.M
			} else {
			}
			v59 = v2 & int32(7)
			if v59 == int32(0) {
			} else {
				v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
				if v64&int32(1) != 0 {
					v67 = v29
				} else {
					v67 = v31
				}
				v70 = int32(base.Ui32(v2)>>(uint(int32(3))%32)) + v67 + int32(2)
				v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
				v74 = v71 & (int32(-256) >> (uint(v59) % 32))
				*(*uint8)(unsafe.Add(mBase, uint32(v70))) = uint8(v74)
			}
		}
		v80 = int32(1)
		v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
		if v82&v80 != 0 {
			v85 = v80
		} else {
			v85 = int32(4)
		}
		v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8+v85))))
		if v87 == int32(2) {
			v90 = int32(40)
		} else {
			v90 = int32(88)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v8))) = v90
		return v8
	}
}
