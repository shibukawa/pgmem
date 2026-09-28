package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ResOwnerPrintFile(m *base.Module, l0 int64) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14221(m, l0, int32(_a_F_ResOwnerPrintFile_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_ResOwnerReleaseCatCache(m *base.Module, l0 int64) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	v4 = base.I32_wrap_i64(l0)
	v6 = v4 - int32(8)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	v8 = int32(1)
	v9 = v7 - v8
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = v9
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4-int32(4)))))
	if base.B2i32(v13 != v8)|v9 != 0 {
		return
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v4)+20))
		if v17 != 0 {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+48))
			if v18 != 0 {
				return
			} else {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v4)+24))
				F_CatCacheRemoveCTup(m, v19, v4-int32(56))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					return
				}
			}
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v4)+24))
			F_CatCacheRemoveCTup(m, v19, v4-int32(56))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_ResOwnerReleasePGMEMZHandle(m *base.Module, l0 int64) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	v4 = base.I32_wrap_i64(l0)
	*(*int32)(unsafe.Add(mBase, uint32(v4)+4)) = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
	m.Env.Pgmem_zstream_free(m, v7)
	mBase = m.M
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
	if v9 != 0 {
		F_ResourceOwnerForget(m, v9, l0&int64(4294967295), int32(_a_F_ResOwnerReleasePGMEMZHandle_0))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return
		} else {
			F_pfree(m, v4)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				return
			}
		}
	} else {
		F_pfree(m, v4)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			return
		}
	}
}
func F_ResOwnerReleaseRelation(m *base.Module, l0 int64) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	v4 = base.I32_wrap_i64(l0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+16))
	v7 = v5 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v4)+16)) = v7
	if v7 != 0 {
		return
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v4)+104))
		if v9 == int32(0) {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v4)+112))
			if v17 == int32(0) {
				return
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
				if v20 == int32(0) {
					return
				} else {
					F_MemoryContextDeleteChildren(m, v17)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return
					} else {
						return
					}
				}
			}
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
			if v12 == int32(0) {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(v4)+112))
				if v17 == int32(0) {
					return
				} else {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
					if v20 == int32(0) {
						return
					} else {
						F_MemoryContextDeleteChildren(m, v17)
						mBase = m.M
						v24 = m.ExcPending
						if v24 != 0 {
							return
						} else {
							return
						}
					}
				}
			} else {
				F_MemoryContextDeleteChildren(m, v9)
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return
				} else {
					v17 = *(*int32)(unsafe.Add(mBase, uint32(v4)+112))
					if v17 == int32(0) {
						return
					} else {
						v20 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
						if v20 == int32(0) {
							return
						} else {
							F_MemoryContextDeleteChildren(m, v17)
							mBase = m.M
							v24 = m.ExcPending
							if v24 != 0 {
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
