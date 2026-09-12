package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GetForeignDataWrapperExtended(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v12 = F_SearchSysCache1(m, int32(30), l0)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		if v12 == int32(0) {
			if l1&int32(1) != 0 {
				v67 = int32(0)
				m.G0 = v9 + int32(16)
				return v67
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
					F_errmsg_internal(m, int32(42780), v9)
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(485525), int32(63), int32(451405))
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			v33 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
			v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+22)))
			v36 = F_palloc(m, int32(24))
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v36))) = l0
				v39 = v33 + v34
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+68))
				*(*int32)(unsafe.Add(mBase, uint32(v36)+4)) = v40
				v44 = F_pstrdup(m, v39+int32(4))
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v36)+8)) = v44
					v47 = *(*int32)(unsafe.Add(mBase, uint32(v39)+72))
					*(*int32)(unsafe.Add(mBase, uint32(v36)+12)) = v47
					v49 = *(*int32)(unsafe.Add(mBase, uint32(v39)+76))
					*(*int32)(unsafe.Add(mBase, uint32(v36)+16)) = v49
					v55 = F_SysCacheGetAttr(m, int32(30), v12, int32(7), v9+int32(15))
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return int32(0)
					} else {
						v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
						if v57 != 0 {
							v61 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v36)+20)) = v61
							F_ReleaseCatCache(m, v12)
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
								return int32(0)
							} else {
								v67 = v36
								m.G0 = v9 + int32(16)
								return v67
							}
						} else {
							v59 = F_untransformRelOptions(m, v55)
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return int32(0)
							} else {
								v61 = v59
								*(*int32)(unsafe.Add(mBase, uint32(v36)+20)) = v61
								F_ReleaseCatCache(m, v12)
								mBase = m.M
								v64 = m.ExcPending
								if v64 != 0 {
									return int32(0)
								} else {
									v67 = v36
									m.G0 = v9 + int32(16)
									return v67
								}
							}
						}
					}
				}
			}
		}
	}
}
