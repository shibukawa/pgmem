package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_InstrAlloc(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v91 int32
	_ = v91
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	v3 = l2
	v4 = int32(0)
	v14 = F_palloc0(m, l0*int32(416))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v20 = int32(0)
		if base.B2i32(l1&int32(11) == v20)|base.B2i32(l0 <= v20) != 0 {
		} else {
			v25 = int32(1)
			v26 = l1 & v25
			v27 = int32(3)
			v30 = int32(base.Ui32(l1)>>(uint(v27)%32)) & v25
			v34 = int32(base.Ui32(l1)>>(uint(v25)%32)) & v25
			v36 = l0 & v27
			v37 = int32(0)
			if base.Ui32(int32(4)) <= base.Ui32(l0) {
				v43 = v37
				v51 = v4
				for {
					v55 = v14 + v43*int32(416)
					*(*uint8)(unsafe.Add(mBase, uint32(v55)+2)) = uint8(v30)
					*(*uint8)(unsafe.Add(mBase, uint32(v55)+1)) = uint8(v34)
					*(*uint8)(unsafe.Add(mBase, uint32(v55)+418)) = uint8(v30)
					*(*uint8)(unsafe.Add(mBase, uint32(v55)+417)) = uint8(v34)
					*(*uint8)(unsafe.Add(mBase, uint32(v55)+3)) = uint8(v3)
					*(*uint8)(unsafe.Add(mBase, uint32(v55))) = uint8(v26)
					*(*uint8)(unsafe.Add(mBase, uint32(v55)+834)) = uint8(v30)
					*(*uint8)(unsafe.Add(mBase, uint32(v55)+833)) = uint8(v34)
					*(*uint8)(unsafe.Add(mBase, uint32(v55)+419)) = uint8(v3)
					*(*uint8)(unsafe.Add(mBase, uint32(v55)+416)) = uint8(v26)
					*(*uint8)(unsafe.Add(mBase, uint32(v55)+1250)) = uint8(v30)
					*(*uint8)(unsafe.Add(mBase, uint32(v55)+1249)) = uint8(v34)
					*(*uint8)(unsafe.Add(mBase, uint32(v55)+835)) = uint8(v3)
					*(*uint8)(unsafe.Add(mBase, uint32(v55)+832)) = uint8(v26)
					*(*uint8)(unsafe.Add(mBase, uint32(v55)+1251)) = uint8(v3)
					*(*uint8)(unsafe.Add(mBase, uint32(v55)+1248)) = uint8(v26)
					v72 = int32(4)
					v73 = v43 + v72
					v75 = v51 + v72
					if v75 != l0&int32(2147483644) {
						v43 = v73
						v51 = v75
						continue
					} else {
						break
					}
					break
				}
				if v36 == int32(0) {
				} else {
					v80 = v73
					v91 = v80
					v100 = v4
					for {
						v103 = v14 + v91*int32(416)
						*(*uint8)(unsafe.Add(mBase, uint32(v103)+2)) = uint8(v30)
						*(*uint8)(unsafe.Add(mBase, uint32(v103)+1)) = uint8(v34)
						*(*uint8)(unsafe.Add(mBase, uint32(v103)+3)) = uint8(v3)
						*(*uint8)(unsafe.Add(mBase, uint32(v103))) = uint8(v26)
						v108 = int32(1)
						v111 = v100 + v108
						if v111 != v36 {
							v91 = v91 + v108
							v100 = v111
							continue
						} else {
							break
						}
						break
					}
				}
			} else {
				v80 = v37
				v91 = v80
				v100 = v4
				for {
					v103 = v14 + v91*int32(416)
					*(*uint8)(unsafe.Add(mBase, uint32(v103)+2)) = uint8(v30)
					*(*uint8)(unsafe.Add(mBase, uint32(v103)+1)) = uint8(v34)
					*(*uint8)(unsafe.Add(mBase, uint32(v103)+3)) = uint8(v3)
					*(*uint8)(unsafe.Add(mBase, uint32(v103))) = uint8(v26)
					v108 = int32(1)
					v111 = v100 + v108
					if v111 != v36 {
						v91 = v91 + v108
						v100 = v111
						continue
					} else {
						break
					}
					break
				}
			}
		}
		return v14
	}
}
func F_InstrEndLoop(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int64
	_ = v5
	var v8 int32
	_ = v8
	var v10 float64
	_ = v10
	var v11 float64
	_ = v11
	var v14 float64
	_ = v14
	var v15 float64
	_ = v15
	var v18 float64
	_ = v18
	var v22 int64
	_ = v22
	var v26 float64
	_ = v26
	var v30 int32
	_ = v30
	var v31 int64
	_ = v31
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	v2 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
	if v2 == int32(1) {
		v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
		if v5 != int64(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return
			} else {
				F_errmsg_internal(m, int32(_a_F_InstrEndLoop_0), int32(0))
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_InstrEndLoop_1), int32(149), int32(_a_F_InstrEndLoop_2))
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v8 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v8)
			v10 = *(*float64)(unsafe.Add(mBase, uint32(l0)+24))
			v11 = *(*float64)(unsafe.Add(mBase, uint32(l0)+200))
			*(*float64)(unsafe.Add(mBase, uint32(l0)+200)) = base.F64_add(v10, v11)
			v14 = *(*float64)(unsafe.Add(mBase, uint32(l0)+32))
			v15 = *(*float64)(unsafe.Add(mBase, uint32(l0)+216))
			*(*float64)(unsafe.Add(mBase, uint32(l0)+216)) = base.F64_add(v14, v15)
			v18 = *(*float64)(unsafe.Add(mBase, uint32(l0)+232))
			*(*float64)(unsafe.Add(mBase, uint32(l0)+232)) = base.F64_add(v18, float64(1))
			v22 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
			v26 = *(*float64)(unsafe.Add(mBase, uint32(l0)+208))
			*(*float64)(unsafe.Add(mBase, uint32(l0)+208)) = base.F64_add(base.F64_div(base.F64_convert_i64_s(v22), float64(1e+09)), v26)
			v30 = l0 + int32(8)
			v31 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v30)+24)) = v31
			*(*int64)(unsafe.Add(mBase, uint32(v30)+16)) = v31
			*(*int64)(unsafe.Add(mBase, uint32(v30)+8)) = v31
			*(*int64)(unsafe.Add(mBase, uint32(v30))) = v31
			return
		}
	} else {
		return
	}
}
