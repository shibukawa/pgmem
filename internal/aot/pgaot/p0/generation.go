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
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v14 = l0 - int32(8)
	v15 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
	if v15&int64(16) != int64(0) {
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l0-int32(32))))
		if v22 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v69 = m.ExcPending
			if v69 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v11))) = v14
				F_errmsg_internal(m, int32(_a_F_GenerationRealloc_0), v11)
				mBase = m.M
				v73 = m.ExcPending
				if v73 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_GenerationRealloc_1), int32(820), int32(_a_F_GenerationRealloc_2))
					mBase = m.M
					v78 = m.ExcPending
					if v78 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
			if v25 != int32(475) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v69 = m.ExcPending
				if v69 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v11))) = v14
					F_errmsg_internal(m, int32(_a_F_GenerationRealloc_0), v11)
					mBase = m.M
					v73 = m.ExcPending
					if v73 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_GenerationRealloc_1), int32(820), int32(_a_F_GenerationRealloc_2))
						mBase = m.M
						v78 = m.ExcPending
						if v78 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v30 = *(*int32)(unsafe.Add(mBase, uint32(l0-int32(12))))
				v45 = v30 - l0
				v46 = l0 - int32(40)
				if base.Ui32(l1) <= base.Ui32(v45) {
					v60 = l0
					m.G0 = v11 + int32(16)
					return v60
				} else {
					v48 = *(*int32)(unsafe.Add(mBase, uint32(v46)+8))
					v49 = F_GenerationAlloc(m, v48, l1, l2)
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return int32(0)
					} else {
						if v49 == int32(0) {
							v55 = F_MemoryContextAllocationFailure(m, v48, l1, l2)
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return int32(0)
							} else {
								v60 = v55
								m.G0 = v11 + int32(16)
								return v60
							}
						} else {
							if v45 != 0 {
								base.MemoryCopy(m, v49, l0, v45)
							} else {
							}
							F_GenerationFree(m, l0)
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
								return int32(0)
							} else {
								v60 = v49
								m.G0 = v11 + int32(16)
								return v60
							}
						}
					}
				}
			}
		}
	} else {
		v45 = base.I32_wrap_i64(int64(base.Ui64(v15)>>(uint(int64(5))%64))) & int32(1073741823)
		v46 = v14 - base.I32_wrap_i64(int64(base.Ui64(v15)>>(uint(int64(34))%64)))&int32(1073741822)
		if base.Ui32(l1) <= base.Ui32(v45) {
			v60 = l0
			m.G0 = v11 + int32(16)
			return v60
		} else {
			v48 = *(*int32)(unsafe.Add(mBase, uint32(v46)+8))
			v49 = F_GenerationAlloc(m, v48, l1, l2)
			mBase = m.M
			v52 = m.ExcPending
			if v52 != 0 {
				return int32(0)
			} else {
				if v49 == int32(0) {
					v55 = F_MemoryContextAllocationFailure(m, v48, l1, l2)
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return int32(0)
					} else {
						v60 = v55
						m.G0 = v11 + int32(16)
						return v60
					}
				} else {
					if v45 != 0 {
						base.MemoryCopy(m, v49, l0, v45)
					} else {
					}
					F_GenerationFree(m, l0)
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
						return int32(0)
					} else {
						v60 = v49
						m.G0 = v11 + int32(16)
						return v60
					}
				}
			}
		}
	}
}
