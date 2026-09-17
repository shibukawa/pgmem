package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pg_create_logical_replication_slot(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v65 int64
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v20 = F_get_call_result_type(m, l0, int32(0), v10+int32(16))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		return int32(0)
	} else {
		if v20 == int32(1) {
			F_CheckSlotPermissions(m)
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				F_CheckLogicalDecodingRequirements(m)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					v30 = int32(1)
					if v14 != 0 {
						v33 = int32(2)
					} else {
						v33 = v30
					}
					v34 = int32(0)
					F_ReplicationSlotCreate(m, v16, v30, v33, base.B2i32(v13 != v34), base.B2i32(v12 != v34), v34)
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = int32(394)
						*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = int32(395)
						*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = int32(396)
						v47 = int32(0)
						v50 = v10 + int32(20)
						v54 = F_CreateInitDecodingContext(m, v15, v47, int64(0), v50, v47, v47, v47)
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int32(0)
						} else {
							F_DecodingContextFindStartpoint(m, v54)
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return int32(0)
							} else {
								F_FreeDecodingContext(m, v54)
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
									return int32(0)
								} else {
									v61 = *(*int32)(unsafe.Add(mBase, _c_F_pg_create_logical_replication_slot[0]))
									*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v61 + int32(24)
									v65 = *(*int64)(unsafe.Add(mBase, uint32(v61)+120))
									v66 = F_Int64GetDatum(m, v65)
									mBase = m.M
									v67 = m.ExcPending
									if v67 != 0 {
										return int32(0)
									} else {
										v68 = int32(0)
										*(*uint16)(unsafe.Add(mBase, uint32(v10)+20)) = uint16(v68)
										*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v66
										v71 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
										v74 = F_heap_form_tuple(m, v71, v10+int32(8), v50)
										mBase = m.M
										v75 = m.ExcPending
										if v75 != 0 {
											return int32(0)
										} else {
											v76 = *(*int32)(unsafe.Add(mBase, uint32(v74)+16))
											v77 = F_HeapTupleHeaderGetDatum(m, v76)
											mBase = m.M
											v78 = m.ExcPending
											if v78 != 0 {
												return int32(0)
											} else {
												if v14 == int32(0) {
													F_ReplicationSlotPersist(m)
													mBase = m.M
													v82 = m.ExcPending
													if v82 != 0 {
														return int32(0)
													} else {
														F_ReplicationSlotRelease(m)
														mBase = m.M
														v84 = m.ExcPending
														if v84 != 0 {
															return int32(0)
														} else {
															m.G0 = v10 + int32(32)
															return v77
														}
													}
												} else {
													F_ReplicationSlotRelease(m)
													mBase = m.M
													v84 = m.ExcPending
													if v84 != 0 {
														return int32(0)
													} else {
														m.G0 = v10 + int32(32)
														return v77
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
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v92 = m.ExcPending
			if v92 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(_a_F_pg_create_logical_replication_slot_0), int32(0))
				mBase = m.M
				v96 = m.ExcPending
				if v96 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_pg_create_logical_replication_slot_1), int32(183), int32(_a_F_pg_create_logical_replication_slot_2))
					mBase = m.M
					v101 = m.ExcPending
					if v101 != 0 {
						return int32(0)
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
