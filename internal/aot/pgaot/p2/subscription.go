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
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int64
	_ = v25
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	if l4 != 0 {
		v19 = int32(0)
		v20 = F_table_open(m, int32(_a_F_UpdateSubscriptionRelState_0), v19)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return
		} else {
			v23 = F_SearchSysCacheCopy(m, int32(68), l1, l0)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return
			} else {
				if v23 != 0 {
					v25 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = v25
					*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = v25
					*(*int32)(unsafe.Add(mBase, uint32(v10)+44)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = int32(16842752)
					*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = l2
					if l3 != v25 {
						v36 = F_Int64GetDatum(m, l3)
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = v36
							v41 = *(*int32)(unsafe.Add(mBase, uint32(v20)+52))
							v48 = F_heap_modify_tuple(m, v23, v41, v10+int32(16), v10+int32(44), v10+int32(12))
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return
							} else {
								F_CatalogTupleUpdate(m, v20, v48+int32(4), v48)
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return
								} else {
									F_relation_close(m, v20, int32(0))
									mBase = m.M
									v56 = m.ExcPending
									if v56 != 0 {
										return
									} else {
										m.G0 = v10 + int32(48)
										return
									}
								}
							}
						}
					} else {
						v39 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v10)+47)) = uint8(v39)
						v41 = *(*int32)(unsafe.Add(mBase, uint32(v20)+52))
						v48 = F_heap_modify_tuple(m, v23, v41, v10+int32(16), v10+int32(44), v10+int32(12))
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return
						} else {
							F_CatalogTupleUpdate(m, v20, v48+int32(4), v48)
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return
							} else {
								F_relation_close(m, v20, int32(0))
								mBase = m.M
								v56 = m.ExcPending
								if v56 != 0 {
									return
								} else {
									m.G0 = v10 + int32(48)
									return
								}
							}
						}
					}
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = l0
						*(*int32)(unsafe.Add(mBase, uint32(v10))) = l1
						F_errmsg_internal(m, int32(_a_F_UpdateSubscriptionRelState_1), v10)
						mBase = m.M
						v68 = m.ExcPending
						if v68 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_UpdateSubscriptionRelState_2), int32(355), int32(_a_F_UpdateSubscriptionRelState_3))
							mBase = m.M
							v73 = m.ExcPending
							if v73 != 0 {
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
				v23 = F_SearchSysCacheCopy(m, int32(68), l1, l0)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return
				} else {
					if v23 != 0 {
						v25 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = v25
						*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = v25
						*(*int32)(unsafe.Add(mBase, uint32(v10)+44)) = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = int32(16842752)
						*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = l2
						if l3 != v25 {
							v36 = F_Int64GetDatum(m, l3)
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = v36
								v41 = *(*int32)(unsafe.Add(mBase, uint32(v20)+52))
								v48 = F_heap_modify_tuple(m, v23, v41, v10+int32(16), v10+int32(44), v10+int32(12))
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
									return
								} else {
									F_CatalogTupleUpdate(m, v20, v48+int32(4), v48)
									mBase = m.M
									v53 = m.ExcPending
									if v53 != 0 {
										return
									} else {
										F_relation_close(m, v20, int32(0))
										mBase = m.M
										v56 = m.ExcPending
										if v56 != 0 {
											return
										} else {
											m.G0 = v10 + int32(48)
											return
										}
									}
								}
							}
						} else {
							v39 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v10)+47)) = uint8(v39)
							v41 = *(*int32)(unsafe.Add(mBase, uint32(v20)+52))
							v48 = F_heap_modify_tuple(m, v23, v41, v10+int32(16), v10+int32(44), v10+int32(12))
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return
							} else {
								F_CatalogTupleUpdate(m, v20, v48+int32(4), v48)
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return
								} else {
									F_relation_close(m, v20, int32(0))
									mBase = m.M
									v56 = m.ExcPending
									if v56 != 0 {
										return
									} else {
										m.G0 = v10 + int32(48)
										return
									}
								}
							}
						}
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = l0
							*(*int32)(unsafe.Add(mBase, uint32(v10))) = l1
							F_errmsg_internal(m, int32(_a_F_UpdateSubscriptionRelState_1), v10)
							mBase = m.M
							v68 = m.ExcPending
							if v68 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_UpdateSubscriptionRelState_2), int32(355), int32(_a_F_UpdateSubscriptionRelState_3))
								mBase = m.M
								v73 = m.ExcPending
								if v73 != 0 {
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
