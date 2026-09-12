package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GenerationRealloc(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int64
	_ = v15
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v14 = l0 - int32(8)
	v15 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
	if v15&int64(16) != int64(0) {
		v21 = l0 - int32(40)
		if v21 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v74 = m.ExcPending
			if v74 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v11))) = v14
				F_errmsg_internal(m, int32(238564), v11)
				mBase = m.M
				v78 = m.ExcPending
				if v78 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(494519), int32(820), int32(487221))
					mBase = m.M
					v83 = m.ExcPending
					if v83 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v26 = *(*int32)(unsafe.Add(mBase, uint32(l0-int32(32))))
			if v26 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v74 = m.ExcPending
				if v74 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v11))) = v14
					F_errmsg_internal(m, int32(238564), v11)
					mBase = m.M
					v78 = m.ExcPending
					if v78 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(494519), int32(820), int32(487221))
						mBase = m.M
						v83 = m.ExcPending
						if v83 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
				if v29 != int32(475) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v74 = m.ExcPending
					if v74 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v11))) = v14
						F_errmsg_internal(m, int32(238564), v11)
						mBase = m.M
						v78 = m.ExcPending
						if v78 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(494519), int32(820), int32(487221))
							mBase = m.M
							v83 = m.ExcPending
							if v83 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v34 = *(*int32)(unsafe.Add(mBase, uint32(l0-int32(12))))
					v47 = v21
					v49 = v34 - l0
					if base.Ui32(l1) <= base.Ui32(v49) {
						v64 = l0
						m.G0 = v11 + int32(16)
						return v64
					} else {
						v51 = *(*int32)(unsafe.Add(mBase, uint32(v47)+8))
						v52 = F_GenerationAlloc(m, v51, l1, l2)
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int32(0)
						} else {
							if v52 == int32(0) {
								v58 = F_MemoryContextAllocationFailure(m, v51, l1, l2)
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
									return int32(0)
								} else {
									v64 = v58
									m.G0 = v11 + int32(16)
									return v64
								}
							} else {
								if v49 != 0 {
									v60 = F__emscripten_memcpy_bulkmem(m, v52, l0, v49)
									mBase = m.M
								} else {
								}
								F_GenerationFree(m, l0)
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return int32(0)
								} else {
									v64 = v52
									m.G0 = v11 + int32(16)
									return v64
								}
							}
						}
					}
				}
			}
		}
	} else {
		v47 = v14 - base.I32_wrap_i64(int64(base.Ui64(v15)>>(uint(int64(34))%64)))&int32(1073741822)
		v49 = base.I32_wrap_i64(int64(base.Ui64(v15)>>(uint(int64(5))%64))) & int32(1073741823)
		if base.Ui32(l1) <= base.Ui32(v49) {
			v64 = l0
			m.G0 = v11 + int32(16)
			return v64
		} else {
			v51 = *(*int32)(unsafe.Add(mBase, uint32(v47)+8))
			v52 = F_GenerationAlloc(m, v51, l1, l2)
			mBase = m.M
			v55 = m.ExcPending
			if v55 != 0 {
				return int32(0)
			} else {
				if v52 == int32(0) {
					v58 = F_MemoryContextAllocationFailure(m, v51, l1, l2)
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
						return int32(0)
					} else {
						v64 = v58
						m.G0 = v11 + int32(16)
						return v64
					}
				} else {
					if v49 != 0 {
						v60 = F__emscripten_memcpy_bulkmem(m, v52, l0, v49)
						mBase = m.M
					} else {
					}
					F_GenerationFree(m, l0)
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return int32(0)
					} else {
						v64 = v52
						m.G0 = v11 + int32(16)
						return v64
					}
				}
			}
		}
	}
}
