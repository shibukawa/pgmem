package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_binary_quantize(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
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
	var v29 int32
	_ = v29
	var v43 int32
	_ = v43
	var v44 float32
	_ = v44
	var v47 int32
	_ = v47
	var v50 float32
	_ = v50
	var v53 int32
	_ = v53
	var v57 float32
	_ = v57
	var v60 int32
	_ = v60
	var v64 float32
	_ = v64
	var v67 int32
	_ = v67
	var v71 float32
	_ = v71
	var v74 int32
	_ = v74
	var v78 float32
	_ = v78
	var v81 int32
	_ = v81
	var v85 float32
	_ = v85
	var v88 int32
	_ = v88
	var v90 float32
	_ = v90
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v108 int32
	_ = v108
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v122 float32
	_ = v122
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v9 = F_pg_detoast_datum(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int64(0)
	} else {
		v14 = v9 + int32(8)
		v15 = int32(*(*int16)(unsafe.Add(mBase, uint32(v9)+4)))
		v16 = F_InitBitVector(m, v15)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int64(0)
		} else {
			v18 = int32(8)
			v19 = v16 + v18
			v20 = int32(*(*int16)(unsafe.Add(mBase, uint32(v9)+4)))
			v22 = base.I32_div_s(v20, v18)
			v23 = int32(0)
			if v18 <= v20 {
				v29 = v23
				for {
					v43 = v14 + v29<<(uint(int32(2))%32)
					v44 = *(*float32)(unsafe.Add(mBase, uint32(v43)))
					if base.F32_gt(v44, float32(0)) != 0 {
						v47 = int32(-128)
					} else {
						v47 = int32(0)
					}
					v50 = *(*float32)(unsafe.Add(mBase, uint32(v43)+4))
					if base.F32_gt(v50, float32(0)) != 0 {
						v53 = int32(64)
					} else {
						v53 = int32(0)
					}
					v57 = *(*float32)(unsafe.Add(mBase, uint32(v43)+8))
					if base.F32_gt(v57, float32(0)) != 0 {
						v60 = int32(32)
					} else {
						v60 = int32(0)
					}
					v64 = *(*float32)(unsafe.Add(mBase, uint32(v43)+12))
					if base.F32_gt(v64, float32(0)) != 0 {
						v67 = int32(16)
					} else {
						v67 = int32(0)
					}
					v71 = *(*float32)(unsafe.Add(mBase, uint32(v43)+16))
					if base.F32_gt(v71, float32(0)) != 0 {
						v74 = int32(8)
					} else {
						v74 = int32(0)
					}
					v78 = *(*float32)(unsafe.Add(mBase, uint32(v43)+20))
					if base.F32_gt(v78, float32(0)) != 0 {
						v81 = int32(4)
					} else {
						v81 = int32(0)
					}
					v85 = *(*float32)(unsafe.Add(mBase, uint32(v43)+24))
					if base.F32_gt(v85, float32(0)) != 0 {
						v88 = int32(2)
					} else {
						v88 = int32(0)
					}
					v90 = *(*float32)(unsafe.Add(mBase, uint32(v43)+28))
					v93 = v47 | v53 | v60 | v67 | v74 | v81 | v88 | base.F32_gt(v90, float32(0))
					*(*uint8)(unsafe.Add(mBase, uint32(v19+int32(base.Ui32(v29)>>(uint(int32(3))%32))))) = uint8(v93)
					v96 = v29 + int32(8)
					if v96 < base.I32_extend16_s(v22)<<(uint(int32(3))%32) {
						v29 = v96
						continue
					} else {
						break
					}
					break
				}
				v98 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9)+4)))
				v99 = v96
				v100 = v98
			} else {
				v99 = v23
				v100 = v20
			}
			if v99 < base.I32_extend16_s(v100) {
				v108 = v99
				for {
					v117 = v19 + int32(base.Ui32(v108)>>(uint(int32(3))%32))
					v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117))))
					v122 = *(*float32)(unsafe.Add(mBase, uint32(v14+v108<<(uint(int32(2))%32))))
					v130 = v118 | base.F32_gt(v122, float32(0))<<(uint((v108^int32(-1))&int32(7))%32)
					*(*uint8)(unsafe.Add(mBase, uint32(v117))) = uint8(v130)
					v133 = v108 + int32(1)
					v134 = int32(*(*int16)(unsafe.Add(mBase, uint32(v9)+4)))
					if v133 < v134 {
						v108 = v133
						continue
					} else {
						break
					}
					break
				}
			} else {
			}
			return base.I64_extend_i32_u(v16)
		}
	}
}
func F_binary_upgrade_create_conflict_detection_slot(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	v3 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_binary_upgrade_create_conflict_detection_slot[0])))
	if v3 == int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int64(0)
		} else {
			F_errcode(m, int32(33685829))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int64(0)
			} else {
				F_errmsg(m, int32(_a_F_binary_upgrade_create_conflict_detection_slot_0), int32(0))
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_binary_upgrade_create_conflict_detection_slot_1), int32(430), int32(_a_F_binary_upgrade_create_conflict_detection_slot_2))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int64(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		v25 = *(*int32)(unsafe.Add(mBase, _c_F_binary_upgrade_create_conflict_detection_slot[1]))
		if v25 == int32(0) {
			v30 = F_errstart(m, int32(15), int32(0))
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int64(0)
			} else {
				if v30 != 0 {
					F_errmsg(m, int32(_a_F_binary_upgrade_create_conflict_detection_slot_3), int32(0))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_binary_upgrade_create_conflict_detection_slot_4), int32(1643), int32(_a_F_binary_upgrade_create_conflict_detection_slot_5))
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int64(0)
						} else {
							v42 = int32(0)
							F_ReplicationSlotCreate(m, int32(_a_F_binary_upgrade_create_conflict_detection_slot_6), v42, v42, v42, v42, v42, v42)
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return int64(0)
							} else {
								F_reset_conflict_slot_xmin_to_safe_horizon(m)
								mBase = m.M
								v51 = m.ExcPending
								if v51 != 0 {
									return int64(0)
								} else {
									F_ReplicationSlotRelease(m)
									mBase = m.M
									v53 = m.ExcPending
									if v53 != 0 {
										return int64(0)
									} else {
										return int64(0)
									}
								}
							}
						}
					}
				} else {
					v42 = int32(0)
					F_ReplicationSlotCreate(m, int32(_a_F_binary_upgrade_create_conflict_detection_slot_6), v42, v42, v42, v42, v42, v42)
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return int64(0)
					} else {
						F_reset_conflict_slot_xmin_to_safe_horizon(m)
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return int64(0)
						} else {
							F_ReplicationSlotRelease(m)
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return int64(0)
							} else {
								return int64(0)
							}
						}
					}
				}
			}
		} else {
			F_ReplicationSlotRelease(m)
			mBase = m.M
			v53 = m.ExcPending
			if v53 != 0 {
				return int64(0)
			} else {
				return int64(0)
			}
		}
	}
}
func F_binary_upgrade_set_next_pg_tablespace_oid(m *base.Module, l0 int32) int64 {
	var v5 int64
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn14232(m, l0, int32(_a_F_binary_upgrade_set_next_pg_tablespace_oid_0), int32(_a_F_binary_upgrade_set_next_pg_tablespace_oid_1), int32(47))
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		return v5
	}
}
func F_binary_upgrade_set_next_pg_type_oid(m *base.Module, l0 int32) int64 {
	var v5 int64
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn14232(m, l0, int32(_a_F_binary_upgrade_set_next_pg_type_oid_0), int32(_a_F_binary_upgrade_set_next_pg_type_oid_1), int32(58))
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		return v5
	}
}
func F_binary_upgrade_set_next_toast_relfilenode(m *base.Module, l0 int32) int64 {
	var v5 int64
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn14232(m, l0, int32(_a_F_binary_upgrade_set_next_toast_relfilenode_0), int32(_a_F_binary_upgrade_set_next_toast_relfilenode_1), int32(157))
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		return v5
	}
}
