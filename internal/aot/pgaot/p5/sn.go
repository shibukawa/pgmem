package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_sn_array_element_start(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	if l1 == int32(0) {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
		if v15 <= int32(0) {
			return int32(0)
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
			v19 = v18 + v15
			v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19-int32(1)))))
			if v22 == int32(91) {
				return int32(0)
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
				if v25 <= v15+int32(1) {
					F_appendStringInfoChar(m, v14, int32(44))
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int32(0)
					} else {
						return int32(0)
					}
				} else {
					v36 = int32(44)
					*(*uint8)(unsafe.Add(mBase, uint32(v19))) = uint8(v36)
					v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
					v41 = v39 + int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v38)+4)) = v41
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
					v45 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v43+v41))) = uint8(v45)
					return int32(0)
				}
			}
		}
	} else {
		v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+9)))
		if v7 != int32(1) {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
			if v15 <= int32(0) {
				return int32(0)
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
				v19 = v18 + v15
				v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19-int32(1)))))
				if v22 == int32(91) {
					return int32(0)
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
					if v25 <= v15+int32(1) {
						F_appendStringInfoChar(m, v14, int32(44))
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int32(0)
						} else {
							return int32(0)
						}
					} else {
						v36 = int32(44)
						*(*uint8)(unsafe.Add(mBase, uint32(v19))) = uint8(v36)
						v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
						v41 = v39 + int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(v38)+4)) = v41
						v43 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
						v45 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v43+v41))) = uint8(v45)
						return int32(0)
					}
				}
			}
		} else {
			v10 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)) = uint8(v10)
			return int32(0)
		}
	}
}
func F_sn_object_end(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v4)+8))
	if v8 <= v5+int32(1) {
		F_appendStringInfoChar(m, v4, int32(125))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			return int32(0)
		}
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
		v19 = int32(125)
		*(*uint8)(unsafe.Add(mBase, uint32(v17+v5))) = uint8(v19)
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
		v24 = v22 + int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = v24
		v26 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
		v28 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v26+v24))) = uint8(v28)
		return v28
	}
}
func F_sn_write(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v14 = v12 - v13
	if base.Ui32(v11) < base.Ui32(v14) {
		v16 = v11
	} else {
		v16 = v14
	}
	if v16 != 0 {
		if v16 != 0 {
			v17 = F__emscripten_memcpy_bulkmem(m, v10, v13, v16)
			mBase = m.M
		} else {
		}
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
		v20 = v19 + v16
		*(*int32)(unsafe.Add(mBase, uint32(v9))) = v20
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
		v23 = v22 - v16
		*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v23
		v25 = v23
		v26 = v20
	} else {
		v25 = v11
		v26 = v10
	}
	if base.Ui32(v25) < base.Ui32(l2) {
		v28 = v25
	} else {
		v28 = l2
	}
	if v28 != 0 {
		if v28 != 0 {
			v29 = F__emscripten_memcpy_bulkmem(m, v26, l1, v28)
			mBase = m.M
		} else {
		}
		v31 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
		v32 = v31 + v28
		*(*int32)(unsafe.Add(mBase, uint32(v9))) = v32
		v34 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
		*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v34 - v28
		v37 = v32
	} else {
		v37 = v26
	}
	v38 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v37))) = uint8(v38)
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v40
	return l2
}
