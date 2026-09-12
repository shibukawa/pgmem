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
	var v56 int32
	_ = v56
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
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
							v116 = m.ExcPending
							if v116 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = l0
								F_errmsg_internal(m, int32(96200), v12+int32(16))
								mBase = m.M
								v122 = m.ExcPending
								if v122 != 0 {
									return
								} else {
									F_errfinish(m, int32(518178), int32(1147), int32(97022))
									mBase = m.M
									v127 = m.ExcPending
									if v127 != 0 {
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
								v131 = m.ExcPending
								if v131 != 0 {
									return
								} else {
									F_errcode(m, int32(261))
									mBase = m.M
									v134 = m.ExcPending
									if v134 != 0 {
										return
									} else {
										F_errmsg(m, int32(128843), int32(0))
										mBase = m.M
										v138 = m.ExcPending
										if v138 != 0 {
											return
										} else {
											F_errfinish(m, int32(518178), int32(1154), int32(97022))
											mBase = m.M
											v143 = m.ExcPending
											if v143 != 0 {
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
									F_recordDependencyOn(m, v12+int32(36), v12+int32(24), int32(80))
									mBase = m.M
									v56 = m.ExcPending
									if v56 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = l2
										*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = int32(1259)
										F_recordDependencyOn(m, v12+int32(36), v12+int32(24), int32(83))
										mBase = m.M
										v68 = m.ExcPending
										if v68 != 0 {
											return
										} else {
											F_ReleaseCatCache(m, v19)
											mBase = m.M
											v93 = m.ExcPending
											if v93 != 0 {
												return
											} else {
												F_sequence_close(m, v16, int32(3))
												mBase = m.M
												v96 = m.ExcPending
												if v96 != 0 {
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
						v69 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v25)+103)) = uint8(v69)
						*(*int32)(unsafe.Add(mBase, uint32(v25)+92)) = int32(0)
						v73 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+104)))
						v75 = v73 - v69
						*(*uint16)(unsafe.Add(mBase, uint32(v25)+104)) = uint16(v75)
						F_CatalogTupleUpdate(m, v16, v19+int32(4), v21)
						mBase = m.M
						v80 = m.ExcPending
						if v80 != 0 {
							return
						} else {
							v81 = int32(2606)
							v84 = F_deleteDependencyRecordsForClass(m, v81, l0, v81, int32(80))
							mBase = m.M
							v85 = m.ExcPending
							if v85 != 0 {
								return
							} else {
								v89 = F_deleteDependencyRecordsForClass(m, int32(2606), l0, int32(1259), int32(83))
								mBase = m.M
								v90 = m.ExcPending
								if v90 != 0 {
									return
								} else {
									F_ReleaseCatCache(m, v19)
									mBase = m.M
									v93 = m.ExcPending
									if v93 != 0 {
										return
									} else {
										F_sequence_close(m, v16, int32(3))
										mBase = m.M
										v96 = m.ExcPending
										if v96 != 0 {
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
				v103 = m.ExcPending
				if v103 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v12))) = l0
					F_errmsg_internal(m, int32(43939), v12)
					mBase = m.M
					v107 = m.ExcPending
					if v107 != 0 {
						return
					} else {
						F_errfinish(m, int32(518178), int32(1138), int32(97022))
						mBase = m.M
						v112 = m.ExcPending
						if v112 != 0 {
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
