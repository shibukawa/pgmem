package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ResOwnerPrintFile(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13836(m, l0, int32(_a_F_ResOwnerPrintFile_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_ResOwnerReleaseCatCache(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	v4 = l0 - int32(8)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
	v6 = int32(1)
	v7 = v5 - v6
	*(*int32)(unsafe.Add(mBase, uint32(v4))) = v7
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0-int32(4)))))
	if base.B2i32(v11 != v6)|v7 != 0 {
		return
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v15 != 0 {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+32))
			if v16 != 0 {
				return
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				F_CatCacheRemoveCTup(m, v17, l0-int32(40))
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return
				} else {
					return
				}
			}
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			F_CatCacheRemoveCTup(m, v17, l0-int32(40))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_ResOwnerReleasePGMEMZHandle(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	m.Env.Pgmem_zstream_free(m, v5)
	mBase = m.M
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v7 != 0 {
		F_ResourceOwnerForget(m, v7, l0, int32(_a_F_ResOwnerReleasePGMEMZHandle_0))
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			F_pfree(m, l0)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return
			} else {
				return
			}
		}
	} else {
		F_pfree(m, l0)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			return
		}
	}
}
func F_ResOwnerReleaseRelation(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
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
	var v22 int32
	_ = v22
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v5 = v3 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v5
	if v5 != 0 {
		return
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
		if v7 == int32(0) {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
			if v15 == int32(0) {
				return
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
				if v18 == int32(0) {
					return
				} else {
					F_MemoryContextDeleteChildren(m, v15)
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return
					} else {
						return
					}
				}
			}
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
			if v10 == int32(0) {
				v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
				if v15 == int32(0) {
					return
				} else {
					v18 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
					if v18 == int32(0) {
						return
					} else {
						F_MemoryContextDeleteChildren(m, v15)
						mBase = m.M
						v22 = m.ExcPending
						if v22 != 0 {
							return
						} else {
							return
						}
					}
				}
			} else {
				F_MemoryContextDeleteChildren(m, v7)
				mBase = m.M
				v14 = m.ExcPending
				if v14 != 0 {
					return
				} else {
					v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
					if v15 == int32(0) {
						return
					} else {
						v18 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
						if v18 == int32(0) {
							return
						} else {
							F_MemoryContextDeleteChildren(m, v15)
							mBase = m.M
							v22 = m.ExcPending
							if v22 != 0 {
								return
							} else {
								return
							}
						}
					}
				}
			}
		}
	}
}
