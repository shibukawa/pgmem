package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pg_create_logical_replication_slot(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int64
	_ = v13
	var v14 int64
	_ = v14
	var v15 int64
	_ = v15
	var v16 int64
	_ = v16
	var v17 int64
	_ = v17
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int64
	_ = v39
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v76 int64
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int64
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v13 = *(*int64)(unsafe.Add(mBase, uint32(l0)+88))
	v14 = *(*int64)(unsafe.Add(mBase, uint32(l0)+72))
	v15 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	v16 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v17 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v21 = F_get_call_result_type(m, l0, int32(0), v11+int32(16))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		return int64(0)
	} else {
		if v21 == int32(1) {
			F_CheckSlotPermissions(m)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int64(0)
			} else {
				F_CheckLogicalDecodingRequirements(m, int32(0))
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int64(0)
				} else {
					v33 = int32(1)
					if v15 == int64(0) {
						v38 = v33
					} else {
						v38 = int32(2)
					}
					v39 = int64(0)
					v41 = int32(0)
					F_ReplicationSlotCreate(m, base.I32_wrap_i64(v17), v33, v38, base.B2i32(v14 != v39), v41, base.B2i32(v13 != v39), v41)
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return int64(0)
					} else {
						F_EnsureLogicalDecodingEnabled(m)
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = int32(414)
							*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = int32(415)
							*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = int32(416)
							v56 = int32(0)
							v60 = v11 + int32(20)
							v64 = F_CreateInitDecodingContext(m, base.I32_wrap_i64(v16), v56, v56, int64(0), v60, v56, v56, v56)
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return int64(0)
							} else {
								F_DecodingContextFindStartpoint(m, v64)
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return int64(0)
								} else {
									F_FreeDecodingContext(m, v64)
									mBase = m.M
									v69 = m.ExcPending
									if v69 != 0 {
										return int64(0)
									} else {
										v71 = *(*int32)(unsafe.Add(mBase, _c_F_pg_create_logical_replication_slot[0]))
										*(*int64)(unsafe.Add(mBase, uint32(v11))) = base.I64_extend_i32_u(v71 + int32(24))
										v76 = *(*int64)(unsafe.Add(mBase, uint32(v71)+120))
										*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = v76
										v78 = int32(0)
										*(*uint16)(unsafe.Add(mBase, uint32(v11)+20)) = uint16(v78)
										v80 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
										v81 = F_heap_form_tuple(m, v80, v11, v60)
										mBase = m.M
										v82 = m.ExcPending
										if v82 != 0 {
											return int64(0)
										} else {
											v83 = *(*int32)(unsafe.Add(mBase, uint32(v81)+16))
											v84 = F_HeapTupleHeaderGetDatum(m, v83)
											mBase = m.M
											v85 = m.ExcPending
											if v85 != 0 {
												return int64(0)
											} else {
												if v15 == int64(0) {
													F_ReplicationSlotPersist(m)
													mBase = m.M
													v89 = m.ExcPending
													if v89 != 0 {
														return int64(0)
													} else {
														F_ReplicationSlotRelease(m)
														mBase = m.M
														v91 = m.ExcPending
														if v91 != 0 {
															return int64(0)
														} else {
															m.G0 = v11 + int32(32)
															return v84
														}
													}
												} else {
													F_ReplicationSlotRelease(m)
													mBase = m.M
													v91 = m.ExcPending
													if v91 != 0 {
														return int64(0)
													} else {
														m.G0 = v11 + int32(32)
														return v84
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
			v99 = m.ExcPending
			if v99 != 0 {
				return int64(0)
			} else {
				F_errmsg_internal(m, int32(_a_F_pg_create_logical_replication_slot_0), int32(0))
				mBase = m.M
				v103 = m.ExcPending
				if v103 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_pg_create_logical_replication_slot_1), int32(212), int32(_a_F_pg_create_logical_replication_slot_2))
					mBase = m.M
					v108 = m.ExcPending
					if v108 != 0 {
						return int64(0)
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
