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
	var v19 int32
	_ = v19
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
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	v10 = m.G0
	v12 = v10 - int32(48)
	m.G0 = v12
	v16 = F_table_open(m, int32(2606), int32(3))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return
	} else {
		v19 = F_SearchSysCache1(m, int32(19), l0)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return
		} else {
			if v19 != 0 {
				v21 = F_heap_copytuple(m, v19)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return
				} else {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
					v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+22)))
					v25 = v23 + v24
					if l1 != 0 {
						v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+92))
						if v26 != 0 {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v114 = m.ExcPending
							if v114 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = l0
								F_errmsg_internal(m, int32(_a_F_ConstraintSetParentConstraint_0), v12+int32(16))
								mBase = m.M
								v120 = m.ExcPending
								if v120 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_ConstraintSetParentConstraint_1), int32(1147), int32(_a_F_ConstraintSetParentConstraint_2))
									mBase = m.M
									v125 = m.ExcPending
									if v125 != 0 {
										return
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v27 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v25)+103)) = uint8(v27)
							v29 = int32(*(*int16)(unsafe.Add(mBase, uint32(v25)+104)))
							v31 = v29 + int32(1)
							*(*uint16)(unsafe.Add(mBase, uint32(v25)+104)) = uint16(v31)
							if base.I32_extend16_s(v31) != v31 {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v129 = m.ExcPending
								if v129 != 0 {
									return
								} else {
									F_errcode(m, int32(261))
									mBase = m.M
									v132 = m.ExcPending
									if v132 != 0 {
										return
									} else {
										F_errmsg(m, int32(_a_F_ConstraintSetParentConstraint_3), int32(0))
										mBase = m.M
										v136 = m.ExcPending
										if v136 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_ConstraintSetParentConstraint_1), int32(1154), int32(_a_F_ConstraintSetParentConstraint_2))
											mBase = m.M
											v141 = m.ExcPending
											if v141 != 0 {
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
								*(*int32)(unsafe.Add(mBase, uint32(v25)+92)) = l1
								F_CatalogTupleUpdate(m, v16, v19+int32(4), v21)
								mBase = m.M
								v39 = m.ExcPending
								if v39 != 0 {
									return
								} else {
									v40 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = v40
									*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = l0
									v43 = int32(2606)
									*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v43
									*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v40
									*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = l1
									*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v43
									v51 = v12 + int32(36)
									v53 = v12 + int32(24)
									F_recordDependencyOn(m, v51, v53, int32(80))
									mBase = m.M
									v56 = m.ExcPending
									if v56 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = l2
										*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = int32(1259)
										F_recordDependencyOn(m, v51, v53, int32(83))
										mBase = m.M
										v64 = m.ExcPending
										if v64 != 0 {
											return
										} else {
											F_ReleaseCatCache(m, v19)
											mBase = m.M
											v91 = m.ExcPending
											if v91 != 0 {
												return
											} else {
												F_relation_close(m, v16, int32(3))
												mBase = m.M
												v94 = m.ExcPending
												if v94 != 0 {
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
						v65 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v25)+103)) = uint8(v65)
						*(*int32)(unsafe.Add(mBase, uint32(v25)+92)) = int32(0)
						v69 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+104)))
						v71 = v69 - v65
						*(*uint16)(unsafe.Add(mBase, uint32(v25)+104)) = uint16(v71)
						F_CatalogTupleUpdate(m, v16, v19+int32(4), v21)
						mBase = m.M
						v76 = m.ExcPending
						if v76 != 0 {
							return
						} else {
							v77 = int32(2606)
							v80 = F_deleteDependencyRecordsForClass(m, v77, l0, v77, int32(80))
							mBase = m.M
							v81 = m.ExcPending
							if v81 != 0 {
								return
							} else {
								v85 = F_deleteDependencyRecordsForClass(m, int32(2606), l0, int32(1259), int32(83))
								mBase = m.M
								v86 = m.ExcPending
								if v86 != 0 {
									return
								} else {
									F_ReleaseCatCache(m, v19)
									mBase = m.M
									v91 = m.ExcPending
									if v91 != 0 {
										return
									} else {
										F_relation_close(m, v16, int32(3))
										mBase = m.M
										v94 = m.ExcPending
										if v94 != 0 {
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
				v101 = m.ExcPending
				if v101 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v12))) = l0
					F_errmsg_internal(m, int32(_a_F_ConstraintSetParentConstraint_4), v12)
					mBase = m.M
					v105 = m.ExcPending
					if v105 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_ConstraintSetParentConstraint_1), int32(1138), int32(_a_F_ConstraintSetParentConstraint_2))
						mBase = m.M
						v110 = m.ExcPending
						if v110 != 0 {
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
