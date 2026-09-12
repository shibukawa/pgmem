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
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = l0 - int32(8)
	v12 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
	if v12&int64(16) != int64(0) {
		v18 = l0 - int32(40)
		if v18 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = v11
				F_errmsg_internal(m, int32(236732), v8)
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return
				} else {
					F_errfinish(m, int32(490866), int32(711), int32(406201))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l0-int32(32))))
			if v23 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = v11
					F_errmsg_internal(m, int32(236732), v8)
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return
					} else {
						F_errfinish(m, int32(490866), int32(711), int32(406201))
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
				if v26 == int32(475) {
					v50 = v18
					v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+20))
					v53 = v51 + int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v50)+20)) = v53
					v55 = *(*int32)(unsafe.Add(mBase, uint32(v50)+16))
					if v53 < v55 {
					} else {
						v57 = *(*int32)(unsafe.Add(mBase, uint32(v50)+8))
						if v57+int32(80) != v50 {
							v61 = *(*int32)(unsafe.Add(mBase, uint32(v57)+60))
							if v61 != v50 {
								v68 = *(*int32)(unsafe.Add(mBase, uint32(v57)+64))
								if v68 == int32(0) {
									*(*int64)(unsafe.Add(mBase, uint32(v50)+16)) = int64(0)
									*(*int32)(unsafe.Add(mBase, uint32(v50)+24)) = v50 + int32(32)
									*(*int32)(unsafe.Add(mBase, uint32(v57)+64)) = v50
								} else {
									v77 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
									v78 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
									*(*int32)(unsafe.Add(mBase, uint32(v77)+4)) = v78
									v80 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
									*(*int32)(unsafe.Add(mBase, uint32(v78))) = v80
									v82 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
									v83 = *(*int32)(unsafe.Add(mBase, uint32(v50)+12))
									*(*int32)(unsafe.Add(mBase, uint32(v57)+8)) = v82 - v83
									F_emscripten_builtin_free(m, v50)
									mBase = m.M
								}
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(v50)+16)) = int64(0)
								*(*int32)(unsafe.Add(mBase, uint32(v50)+24)) = v50 + int32(32)
							}
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(v50)+16)) = int64(0)
							*(*int32)(unsafe.Add(mBase, uint32(v50)+24)) = v50 + int32(32)
						}
					}
					m.G0 = v8 + int32(16)
					return
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = v11
						F_errmsg_internal(m, int32(236732), v8)
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return
						} else {
							F_errfinish(m, int32(490866), int32(711), int32(406201))
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			}
		}
	} else {
		v50 = v11 - base.I32_wrap_i64(int64(base.Ui64(v12)>>(uint(int64(34))%64)))&int32(1073741822)
		v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+20))
		v53 = v51 + int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(v50)+20)) = v53
		v55 = *(*int32)(unsafe.Add(mBase, uint32(v50)+16))
		if v53 < v55 {
		} else {
			v57 = *(*int32)(unsafe.Add(mBase, uint32(v50)+8))
			if v57+int32(80) != v50 {
				v61 = *(*int32)(unsafe.Add(mBase, uint32(v57)+60))
				if v61 != v50 {
					v68 = *(*int32)(unsafe.Add(mBase, uint32(v57)+64))
					if v68 == int32(0) {
						*(*int64)(unsafe.Add(mBase, uint32(v50)+16)) = int64(0)
						*(*int32)(unsafe.Add(mBase, uint32(v50)+24)) = v50 + int32(32)
						*(*int32)(unsafe.Add(mBase, uint32(v57)+64)) = v50
					} else {
						v77 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
						v78 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v77)+4)) = v78
						v80 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
						*(*int32)(unsafe.Add(mBase, uint32(v78))) = v80
						v82 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
						v83 = *(*int32)(unsafe.Add(mBase, uint32(v50)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v57)+8)) = v82 - v83
						F_emscripten_builtin_free(m, v50)
						mBase = m.M
					}
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v50)+16)) = int64(0)
					*(*int32)(unsafe.Add(mBase, uint32(v50)+24)) = v50 + int32(32)
				}
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v50)+16)) = int64(0)
				*(*int32)(unsafe.Add(mBase, uint32(v50)+24)) = v50 + int32(32)
			}
		}
		m.G0 = v8 + int32(16)
		return
	}
}
