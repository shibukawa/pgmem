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
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int64
	_ = v24
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	if l4 != 0 {
		v18 = int32(0)
		v19 = F_table_open(m, int32(6102), v18)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return
		} else {
			v22 = F_SearchSysCacheCopy(m, int32(68), l1, l0)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return
			} else {
				if v22 != 0 {
					v24 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = v24
					*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = v24
					*(*int32)(unsafe.Add(mBase, uint32(v10)+44)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = int32(16842752)
					*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = l2
					if l3 != v24 {
						v35 = F_Int64GetDatum(m, l3)
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = v35
							v40 = *(*int32)(unsafe.Add(mBase, uint32(v19)+52))
							v47 = F_heap_modify_tuple(m, v22, v40, v10+int32(16), v10+int32(44), v10+int32(12))
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return
							} else {
								F_CatalogTupleUpdate(m, v19, v47+int32(4), v47)
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
									return
								} else {
									F_sequence_close(m, v19, int32(0))
									mBase = m.M
									v55 = m.ExcPending
									if v55 != 0 {
										return
									} else {
										m.G0 = v10 + int32(48)
										return
									}
								}
							}
						}
					} else {
						v38 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v10)+47)) = uint8(v38)
						v40 = *(*int32)(unsafe.Add(mBase, uint32(v19)+52))
						v47 = F_heap_modify_tuple(m, v22, v40, v10+int32(16), v10+int32(44), v10+int32(12))
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return
						} else {
							F_CatalogTupleUpdate(m, v19, v47+int32(4), v47)
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return
							} else {
								F_sequence_close(m, v19, int32(0))
								mBase = m.M
								v55 = m.ExcPending
								if v55 != 0 {
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
					v62 = m.ExcPending
					if v62 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = l0
						*(*int32)(unsafe.Add(mBase, uint32(v10))) = l1
						F_errmsg_internal(m, int32(68798), v10)
						mBase = m.M
						v67 = m.ExcPending
						if v67 != 0 {
							return
						} else {
							F_errfinish(m, int32(496196), int32(355), int32(353810))
							mBase = m.M
							v72 = m.ExcPending
							if v72 != 0 {
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
		F_LockSharedObject(m, int32(6100), l0, int32(1))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			v18 = int32(3)
			v19 = F_table_open(m, int32(6102), v18)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return
			} else {
				v22 = F_SearchSysCacheCopy(m, int32(68), l1, l0)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					if v22 != 0 {
						v24 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = v24
						*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = v24
						*(*int32)(unsafe.Add(mBase, uint32(v10)+44)) = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = int32(16842752)
						*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = l2
						if l3 != v24 {
							v35 = F_Int64GetDatum(m, l3)
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = v35
								v40 = *(*int32)(unsafe.Add(mBase, uint32(v19)+52))
								v47 = F_heap_modify_tuple(m, v22, v40, v10+int32(16), v10+int32(44), v10+int32(12))
								mBase = m.M
								v48 = m.ExcPending
								if v48 != 0 {
									return
								} else {
									F_CatalogTupleUpdate(m, v19, v47+int32(4), v47)
									mBase = m.M
									v52 = m.ExcPending
									if v52 != 0 {
										return
									} else {
										F_sequence_close(m, v19, int32(0))
										mBase = m.M
										v55 = m.ExcPending
										if v55 != 0 {
											return
										} else {
											m.G0 = v10 + int32(48)
											return
										}
									}
								}
							}
						} else {
							v38 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v10)+47)) = uint8(v38)
							v40 = *(*int32)(unsafe.Add(mBase, uint32(v19)+52))
							v47 = F_heap_modify_tuple(m, v22, v40, v10+int32(16), v10+int32(44), v10+int32(12))
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return
							} else {
								F_CatalogTupleUpdate(m, v19, v47+int32(4), v47)
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
									return
								} else {
									F_sequence_close(m, v19, int32(0))
									mBase = m.M
									v55 = m.ExcPending
									if v55 != 0 {
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
						v62 = m.ExcPending
						if v62 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = l0
							*(*int32)(unsafe.Add(mBase, uint32(v10))) = l1
							F_errmsg_internal(m, int32(68798), v10)
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return
							} else {
								F_errfinish(m, int32(496196), int32(355), int32(353810))
								mBase = m.M
								v72 = m.ExcPending
								if v72 != 0 {
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
