package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GenerationDelete(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	v2 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v2
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v8 == v2 {
	} else {
		v12 = l0 + int32(68)
		if v8 == v12 {
		} else {
			v17 = v8
			for {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
				if v17 == l0+int32(80) {
					*(*int64)(unsafe.Add(mBase, uint32(v17)+16)) = int64(0)
					*(*int32)(unsafe.Add(mBase, uint32(v17)+24)) = v17 + int32(32)
				} else {
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
					*(*int32)(unsafe.Add(mBase, uint32(v28)+4)) = v21
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
					*(*int32)(unsafe.Add(mBase, uint32(v21))) = v30
					v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v32 - v33
					F_emscripten_builtin_free(m, v17)
					mBase = m.M
				}
				if v21 != v12 {
					v17 = v21
					continue
				} else {
					break
				}
				break
			}
		}
	}
	F_emscripten_builtin_free(m, l0)
	mBase = m.M
	return
}
func F_GenerationGetChunkSpace(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int64
	_ = v5
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l0-int32(8))))
	if v5&int64(16) != int64(0) {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0-int32(12))))
		v19 = v12 - l0
	} else {
		v19 = base.I32_wrap_i64(int64(base.Ui64(v5)>>(uint(int64(5))%64))) & int32(1073741823)
	}
	return v19 + int32(8)
}
func F_GenerationReset(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v46 int32
	_ = v46
	v2 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v2
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v8 == v2 {
	} else {
		v12 = l0 + int32(68)
		if v8 == v12 {
		} else {
			v17 = v8
			for {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
				if v17 == l0+int32(80) {
					*(*int64)(unsafe.Add(mBase, uint32(v17)+16)) = int64(0)
					*(*int32)(unsafe.Add(mBase, uint32(v17)+24)) = v17 + int32(32)
				} else {
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
					*(*int32)(unsafe.Add(mBase, uint32(v28)+4)) = v21
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
					*(*int32)(unsafe.Add(mBase, uint32(v21))) = v30
					v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v32 - v33
					F_emscripten_builtin_free(m, v17)
					mBase = m.M
				}
				if v21 != v12 {
					v17 = v21
					continue
				} else {
					break
				}
				break
			}
		}
	}
	*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = l0 + int32(80)
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v46
	return
}
