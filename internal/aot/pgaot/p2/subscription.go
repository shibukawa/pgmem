package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_UpdateSubscriptionRelState(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int64
	_ = v27
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	v8 = m.G0
	v10 = v8 + int32(-64)
	m.G0 = v10
	if l4 != 0 {
		v19 = int32(0)
		v20 = F_table_open(m, int32(_a_F_UpdateSubscriptionRelState_0), v19)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return
		} else {
			v25 = F_SearchSysCacheCopy(m, int32(68), base.I64_extend_i32_u(l1), base.I64_extend_i32_u(l0))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return
			} else {
				if v25 != 0 {
					v27 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v10)+40)) = v27
					*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = v27
					*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = v27
					*(*int32)(unsafe.Add(mBase, uint32(v10)+60)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = int32(16842752)
					*(*int64)(unsafe.Add(mBase, uint32(v10)+32)) = base.I64_extend_i32_s(l2)
					if l3 != v27 {
						*(*int64)(unsafe.Add(mBase, uint32(v10)+40)) = l3
					} else {
						v42 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v10)+63)) = uint8(v42)
					}
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v20)+52))
					v51 = F_heap_modify_tuple(m, v25, v44, v8+int32(-48), v8+int32(-4), v8+int32(-52))
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return
					} else {
						F_CatalogTupleUpdate(m, v20, v51+int32(4), v51)
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return
						} else {
							F_relation_close(m, v20, int32(0))
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
								return
							} else {
								m.G0 = v10 - int32(-64)
								return
							}
						}
					}
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v66 = m.ExcPending
					if v66 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = l0
						*(*int32)(unsafe.Add(mBase, uint32(v10))) = l1
						F_errmsg_internal(m, int32(_a_F_UpdateSubscriptionRelState_1), v10)
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_UpdateSubscriptionRelState_2), int32(431), int32(_a_F_UpdateSubscriptionRelState_3))
							mBase = m.M
							v76 = m.ExcPending
							if v76 != 0 {
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
		F_LockSharedObject(m, int32(_a_F_UpdateSubscriptionRelState_4), l0, int32(1))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			v19 = int32(3)
			v20 = F_table_open(m, int32(_a_F_UpdateSubscriptionRelState_0), v19)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				v25 = F_SearchSysCacheCopy(m, int32(68), base.I64_extend_i32_u(l1), base.I64_extend_i32_u(l0))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return
				} else {
					if v25 != 0 {
						v27 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v10)+40)) = v27
						*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = v27
						*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = v27
						*(*int32)(unsafe.Add(mBase, uint32(v10)+60)) = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = int32(16842752)
						*(*int64)(unsafe.Add(mBase, uint32(v10)+32)) = base.I64_extend_i32_s(l2)
						if l3 != v27 {
							*(*int64)(unsafe.Add(mBase, uint32(v10)+40)) = l3
						} else {
							v42 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v10)+63)) = uint8(v42)
						}
						v44 = *(*int32)(unsafe.Add(mBase, uint32(v20)+52))
						v51 = F_heap_modify_tuple(m, v25, v44, v8+int32(-48), v8+int32(-4), v8+int32(-52))
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return
						} else {
							F_CatalogTupleUpdate(m, v20, v51+int32(4), v51)
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return
							} else {
								F_relation_close(m, v20, int32(0))
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
									return
								} else {
									m.G0 = v10 - int32(-64)
									return
								}
							}
						}
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = l0
							*(*int32)(unsafe.Add(mBase, uint32(v10))) = l1
							F_errmsg_internal(m, int32(_a_F_UpdateSubscriptionRelState_1), v10)
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_UpdateSubscriptionRelState_2), int32(431), int32(_a_F_UpdateSubscriptionRelState_3))
								mBase = m.M
								v76 = m.ExcPending
								if v76 != 0 {
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
		}
	}
}
