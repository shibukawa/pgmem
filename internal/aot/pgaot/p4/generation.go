package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GenerationFree(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = l0 - int32(8)
	v12 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
	if v12&int64(16) != int64(0) {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0-int32(32))))
		if v19 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = v11
				F_errmsg_internal(m, int32(_a_F_GenerationFree_0), v8)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_GenerationFree_1), int32(711), int32(_a_F_GenerationFree_2))
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
			if v22 != int32(475) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = v11
					F_errmsg_internal(m, int32(_a_F_GenerationFree_0), v8)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_GenerationFree_1), int32(711), int32(_a_F_GenerationFree_2))
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v47 = l0 - int32(40)
				v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+20))
				v50 = v48 + int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v47)+20)) = v50
				v52 = *(*int32)(unsafe.Add(mBase, uint32(v47)+16))
				if v50 < v52 {
				} else {
					v54 = *(*int32)(unsafe.Add(mBase, uint32(v47)+8))
					if v54+int32(80) != v47 {
						v58 = *(*int32)(unsafe.Add(mBase, uint32(v54)+60))
						if v58 != v47 {
							v65 = *(*int32)(unsafe.Add(mBase, uint32(v54)+64))
							if v65 == int32(0) {
								*(*int64)(unsafe.Add(mBase, uint32(v47)+16)) = int64(0)
								*(*int32)(unsafe.Add(mBase, uint32(v47)+24)) = v47 + int32(32)
								*(*int32)(unsafe.Add(mBase, uint32(v54)+64)) = v47
							} else {
								v74 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
								v75 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
								*(*int32)(unsafe.Add(mBase, uint32(v74)+4)) = v75
								v77 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
								*(*int32)(unsafe.Add(mBase, uint32(v75))) = v77
								v79 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
								v80 = *(*int32)(unsafe.Add(mBase, uint32(v47)+12))
								*(*int32)(unsafe.Add(mBase, uint32(v54)+8)) = v79 - v80
								F_emscripten_builtin_free(m, v47)
								mBase = m.M
							}
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(v47)+16)) = int64(0)
							*(*int32)(unsafe.Add(mBase, uint32(v47)+24)) = v47 + int32(32)
						}
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v47)+16)) = int64(0)
						*(*int32)(unsafe.Add(mBase, uint32(v47)+24)) = v47 + int32(32)
					}
				}
				m.G0 = v8 + int32(16)
				return
			}
		}
	} else {
		v47 = v11 - base.I32_wrap_i64(int64(base.Ui64(v12)>>(uint(int64(34))%64)))&int32(1073741822)
		v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+20))
		v50 = v48 + int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(v47)+20)) = v50
		v52 = *(*int32)(unsafe.Add(mBase, uint32(v47)+16))
		if v50 < v52 {
		} else {
			v54 = *(*int32)(unsafe.Add(mBase, uint32(v47)+8))
			if v54+int32(80) != v47 {
				v58 = *(*int32)(unsafe.Add(mBase, uint32(v54)+60))
				if v58 != v47 {
					v65 = *(*int32)(unsafe.Add(mBase, uint32(v54)+64))
					if v65 == int32(0) {
						*(*int64)(unsafe.Add(mBase, uint32(v47)+16)) = int64(0)
						*(*int32)(unsafe.Add(mBase, uint32(v47)+24)) = v47 + int32(32)
						*(*int32)(unsafe.Add(mBase, uint32(v54)+64)) = v47
					} else {
						v74 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
						v75 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v74)+4)) = v75
						v77 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
						*(*int32)(unsafe.Add(mBase, uint32(v75))) = v77
						v79 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
						v80 = *(*int32)(unsafe.Add(mBase, uint32(v47)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v54)+8)) = v79 - v80
						F_emscripten_builtin_free(m, v47)
						mBase = m.M
					}
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v47)+16)) = int64(0)
					*(*int32)(unsafe.Add(mBase, uint32(v47)+24)) = v47 + int32(32)
				}
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v47)+16)) = int64(0)
				*(*int32)(unsafe.Add(mBase, uint32(v47)+24)) = v47 + int32(32)
			}
		}
		m.G0 = v8 + int32(16)
		return
	}
}
