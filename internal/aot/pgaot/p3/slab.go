package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SlabDelete(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v5 == int32(0) {
	} else {
		v9 = l0 + int32(68)
		if v5 == v9 {
		} else {
			v12 = v5
			for {
				v15 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
				v16 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v16
				v18 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
				*(*int32)(unsafe.Add(mBase, uint32(v16))) = v18
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+76)) = v20 - int32(1)
				F_emscripten_builtin_free(m, v12-int32(20))
				mBase = m.M
				v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v27 - v28
				if v16 != v9 {
					v12 = v16
					continue
				} else {
					break
				}
				break
			}
		}
	}
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	if v36 == int32(0) {
	} else {
		v40 = l0 + int32(80)
		if v36 == v40 {
		} else {
			v43 = v36
			for {
				v46 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
				v47 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v46)+4)) = v47
				v49 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
				*(*int32)(unsafe.Add(mBase, uint32(v47))) = v49
				F_emscripten_builtin_free(m, v43-int32(20))
				mBase = m.M
				v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v54 - v55
				if v47 != v40 {
					v43 = v47
					continue
				} else {
					break
				}
				break
			}
		}
	}
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	if v63 == int32(0) {
	} else {
		v67 = l0 + int32(88)
		if v63 == v67 {
		} else {
			v70 = v63
			for {
				v73 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
				v74 = *(*int32)(unsafe.Add(mBase, uint32(v70)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v73)+4)) = v74
				v76 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
				*(*int32)(unsafe.Add(mBase, uint32(v74))) = v76
				F_emscripten_builtin_free(m, v70-int32(20))
				mBase = m.M
				v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v81 - v82
				if v74 != v67 {
					v70 = v74
					continue
				} else {
					break
				}
				break
			}
		}
	}
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	if v90 == int32(0) {
	} else {
		v94 = l0 + int32(96)
		if v90 == v94 {
		} else {
			v97 = v90
			for {
				v100 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
				v101 = *(*int32)(unsafe.Add(mBase, uint32(v97)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v100)+4)) = v101
				v103 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
				*(*int32)(unsafe.Add(mBase, uint32(v101))) = v103
				F_emscripten_builtin_free(m, v97-int32(20))
				mBase = m.M
				v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v108 - v109
				if v101 != v94 {
					v97 = v101
					continue
				} else {
					break
				}
				break
			}
		}
	}
	*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = int32(0)
	F_emscripten_builtin_free(m, l0)
	mBase = m.M
	return
}
func F_SlabGetChunkSpace(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int64
	_ = v4
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	v3 = l0 - int32(8)
	v4 = *(*int64)(unsafe.Add(mBase, uint32(v3)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v3-base.I32_wrap_i64(int64(base.Ui64(v4)>>(uint(int64(34))%64)))&int32(1073741822))))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+48))
	return v12
}
func F_SlabIsEmpty(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	return base.B2i32(v2 == int32(0))
}
