package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ConstraintSetParentConstraint(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	v10 = m.G0
	v12 = v10 - int32(48)
	m.G0 = v12
	v16 = F_table_open(m, int32(2606), int32(3))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return
	} else {
		v20 = F_SearchSysCache1(m, int32(19), base.I64_extend_i32_u(l0))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return
		} else {
			if v20 != 0 {
				v22 = F_heap_copytuple(m, v20)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					v24 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
					v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+22)))
					v26 = v24 + v25
					if l1 != 0 {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+92))
						if v27 != 0 {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v115 = m.ExcPending
							if v115 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = l0
								F_errmsg_internal(m, int32(_a_F_ConstraintSetParentConstraint_0), v12+int32(16))
								mBase = m.M
								v121 = m.ExcPending
								if v121 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_ConstraintSetParentConstraint_1), int32(1152), int32(_a_F_ConstraintSetParentConstraint_2))
									mBase = m.M
									v126 = m.ExcPending
									if v126 != 0 {
										return
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v28 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v26)+103)) = uint8(v28)
							v30 = int32(*(*int16)(unsafe.Add(mBase, uint32(v26)+104)))
							v32 = v30 + int32(1)
							*(*uint16)(unsafe.Add(mBase, uint32(v26)+104)) = uint16(v32)
							if base.I32_extend16_s(v32) != v32 {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v130 = m.ExcPending
								if v130 != 0 {
									return
								} else {
									F_errcode(m, int32(261))
									mBase = m.M
									v133 = m.ExcPending
									if v133 != 0 {
										return
									} else {
										F_errmsg(m, int32(_a_F_ConstraintSetParentConstraint_3), int32(0))
										mBase = m.M
										v137 = m.ExcPending
										if v137 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_ConstraintSetParentConstraint_1), int32(1159), int32(_a_F_ConstraintSetParentConstraint_2))
											mBase = m.M
											v142 = m.ExcPending
											if v142 != 0 {
												return
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v26)+92)) = l1
								F_CatalogTupleUpdate(m, v16, v20+int32(4), v22)
								mBase = m.M
								v40 = m.ExcPending
								if v40 != 0 {
									return
								} else {
									v41 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = v41
									*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = l0
									v44 = int32(2606)
									*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v44
									*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v41
									*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = l1
									*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v44
									v52 = v12 + int32(36)
									v54 = v12 + int32(24)
									F_recordDependencyOn(m, v52, v54, int32(80))
									mBase = m.M
									v57 = m.ExcPending
									if v57 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = l2
										*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = int32(1259)
										F_recordDependencyOn(m, v52, v54, int32(83))
										mBase = m.M
										v65 = m.ExcPending
										if v65 != 0 {
											return
										} else {
											F_ReleaseCatCache(m, v20)
											mBase = m.M
											v92 = m.ExcPending
											if v92 != 0 {
												return
											} else {
												F_relation_close(m, v16, int32(3))
												mBase = m.M
												v95 = m.ExcPending
												if v95 != 0 {
													return
												} else {
													m.G0 = v12 + int32(48)
													return
												}
											}
										}
									}
								}
							}
						}
					} else {
						v66 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v26)+103)) = uint8(v66)
						*(*int32)(unsafe.Add(mBase, uint32(v26)+92)) = int32(0)
						v70 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26)+104)))
						v72 = v70 - v66
						*(*uint16)(unsafe.Add(mBase, uint32(v26)+104)) = uint16(v72)
						F_CatalogTupleUpdate(m, v16, v20+int32(4), v22)
						mBase = m.M
						v77 = m.ExcPending
						if v77 != 0 {
							return
						} else {
							v78 = int32(2606)
							v81 = F_deleteDependencyRecordsForClass(m, v78, l0, v78, int32(80))
							mBase = m.M
							v82 = m.ExcPending
							if v82 != 0 {
								return
							} else {
								v86 = F_deleteDependencyRecordsForClass(m, int32(2606), l0, int32(1259), int32(83))
								mBase = m.M
								v87 = m.ExcPending
								if v87 != 0 {
									return
								} else {
									F_ReleaseCatCache(m, v20)
									mBase = m.M
									v92 = m.ExcPending
									if v92 != 0 {
										return
									} else {
										F_relation_close(m, v16, int32(3))
										mBase = m.M
										v95 = m.ExcPending
										if v95 != 0 {
											return
										} else {
											m.G0 = v12 + int32(48)
											return
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
				v102 = m.ExcPending
				if v102 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v12))) = l0
					F_errmsg_internal(m, int32(_a_F_ConstraintSetParentConstraint_4), v12)
					mBase = m.M
					v106 = m.ExcPending
					if v106 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_ConstraintSetParentConstraint_1), int32(1143), int32(_a_F_ConstraintSetParentConstraint_2))
						mBase = m.M
						v111 = m.ExcPending
						if v111 != 0 {
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
