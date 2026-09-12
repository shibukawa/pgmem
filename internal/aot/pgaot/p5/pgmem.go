package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_gen_pgmem_iv_size(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)+92))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+12))
	return v4
}
func F_pgmem_aes_init(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+92))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	if base.Ui32(l2) < base.Ui32(int32(17)) {
		v21 = int32(16)
		*(*int32)(unsafe.Add(mBase, uint32(v8)+84)) = v21
		if l2 != 0 {
			v26 = F__emscripten_memcpy_bulkmem(m, v8+int32(4), l1, l2)
			mBase = m.M
		} else {
		}
		v29 = v8 + int32(68)
		if l3 != 0 {
			if v10 != 0 {
				v30 = F__emscripten_memcpy_bulkmem(m, v29, l3, v10)
				mBase = m.M
			} else {
			}
			return int32(0)
		} else {
			v34 = int32(0)
			v37 = F__emscripten_memset_bulkmem(m, v29, base.I32_extend8_s(v34), v10)
			mBase = m.M
			v40 = v34
			return v40
		}
	} else {
		if base.Ui32(l2) < base.Ui32(int32(25)) {
			v21 = int32(24)
			*(*int32)(unsafe.Add(mBase, uint32(v8)+84)) = v21
			if l2 != 0 {
				v26 = F__emscripten_memcpy_bulkmem(m, v8+int32(4), l1, l2)
				mBase = m.M
			} else {
			}
			v29 = v8 + int32(68)
			if l3 != 0 {
				if v10 != 0 {
					v30 = F__emscripten_memcpy_bulkmem(m, v29, l3, v10)
					mBase = m.M
				} else {
				}
				return int32(0)
			} else {
				v34 = int32(0)
				v37 = F__emscripten_memset_bulkmem(m, v29, base.I32_extend8_s(v34), v10)
				mBase = m.M
				v40 = v34
				return v40
			}
		} else {
			v17 = int32(32)
			if base.Ui32(v17) < base.Ui32(l2) {
				v40 = int32(-7)
				return v40
			} else {
				v21 = v17
				*(*int32)(unsafe.Add(mBase, uint32(v8)+84)) = v21
				if l2 != 0 {
					v26 = F__emscripten_memcpy_bulkmem(m, v8+int32(4), l1, l2)
					mBase = m.M
				} else {
				}
				v29 = v8 + int32(68)
				if l3 != 0 {
					if v10 != 0 {
						v30 = F__emscripten_memcpy_bulkmem(m, v29, l3, v10)
						mBase = m.M
					} else {
					}
					return int32(0)
				} else {
					v34 = int32(0)
					v37 = F__emscripten_memset_bulkmem(m, v29, base.I32_extend8_s(v34), v10)
					mBase = m.M
					v40 = v34
					return v40
				}
			}
		}
	}
}
func F_pgmem_call_sighandler(m *base.Module, l0 int32, l1 int32) {
	var v4 int32
	_ = v4
	m.T0[l0].(func(*base.Module, int32))(m, l1)
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		return
	}
}
func F_pgmem_des3_init(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int64
	_ = v11
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+92))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	v9 = int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+84)) = v9
	v11 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v6)+4)) = v11
	*(*int64)(unsafe.Add(mBase, uint32(v6)+12)) = v11
	*(*int64)(unsafe.Add(mBase, uint32(v6)+20)) = v11
	if base.Ui32(v9) <= base.Ui32(l2) {
		v22 = v9
	} else {
		v22 = l2
	}
	if v22 != 0 {
		v23 = F__emscripten_memcpy_bulkmem(m, v6+int32(4), l1, v22)
		mBase = m.M
	} else {
	}
	v26 = v6 + int32(68)
	if l3 != 0 {
		if v8 != 0 {
			v27 = F__emscripten_memcpy_bulkmem(m, v26, l3, v8)
			mBase = m.M
		} else {
		}
		return int32(0)
	} else {
		v33 = F__emscripten_memset_bulkmem(m, v26, base.I32_extend8_s(int32(0)), v8)
		mBase = m.M
		return int32(0)
	}
}
func F_pgmem_init(m *base.Module) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, _consts[1341])) = int32(4731)
	*(*int32)(unsafe.Add(mBase, _consts[1342])) = int32(4730)
	*(*int32)(unsafe.Add(mBase, _consts[1343])) = int32(4732)
	*(*int32)(unsafe.Add(mBase, _consts[1344])) = int32(4733)
	*(*int32)(unsafe.Add(mBase, _consts[1345])) = int32(4734)
	return
}
