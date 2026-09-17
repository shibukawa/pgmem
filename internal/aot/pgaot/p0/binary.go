package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_appendBinaryStringInfoNT(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	F_enlargeStringInfo(m, l0, l2)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		if l2 != 0 {
			v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			base.MemoryCopy(m, v6+v7, l1, l2)
		} else {
		}
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v10 + l2
		return
	}
}
func F_binary_upgrade_set_missing_value(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
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
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v54 int64
	_ = v54
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v109 int32
	_ = v109
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v150 int32
	_ = v150
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v176 int32
	_ = v176
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v17 = F_pg_detoast_datum(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v19 = F_text_to_cstring(m, v12)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				v21 = F_text_to_cstring(m, v17)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					v24 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_binary_upgrade_set_missing_value[0])))
					if v24 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(33685829))
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(_a_F_binary_upgrade_set_missing_value_0), int32(0))
								mBase = m.M
								v37 = m.ExcPending
								if v37 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_binary_upgrade_set_missing_value_1), int32(268), int32(_a_F_binary_upgrade_set_missing_value_2))
									mBase = m.M
									v42 = m.ExcPending
									if v42 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					} else {
						v43 = m.G0
						v45 = v43 - int32(192)
						m.G0 = v45
						v49 = int32(0)
						base.MemoryFill(m, v45+int32(80), v49, int32(96))
						*(*uint8)(unsafe.Add(mBase, uint32(v45)+72)) = uint8(v49)
						v54 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v45)+64)) = v54
						*(*int64)(unsafe.Add(mBase, uint32(v45)+56)) = v54
						*(*int64)(unsafe.Add(mBase, uint32(v45)+48)) = v54
						*(*uint8)(unsafe.Add(mBase, uint32(v45)+40)) = uint8(v49)
						*(*int64)(unsafe.Add(mBase, uint32(v45)+32)) = v54
						*(*int64)(unsafe.Add(mBase, uint32(v45)+24)) = v54
						*(*int64)(unsafe.Add(mBase, uint32(v45)+16)) = v54
						v69 = F_table_open(m, v10, int32(8))
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
							return int32(0)
						} else {
							v71 = *(*int32)(unsafe.Add(mBase, uint32(v69)+48))
							v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+119)))
							if v72 != int32(114) {
								F_relation_close(m, v69, int32(8))
								mBase = m.M
								v176 = m.ExcPending
								if v176 != 0 {
									return int32(0)
								} else {
									m.G0 = v45 + int32(192)
									return int32(0)
								}
							} else {
								v77 = F_table_open(m, int32(1249), int32(3))
								mBase = m.M
								v78 = m.ExcPending
								if v78 != 0 {
									return int32(0)
								} else {
									v79 = F_SearchSysCacheAttName(m, v10, v19)
									mBase = m.M
									v80 = m.ExcPending
									if v80 != 0 {
										return int32(0)
									} else {
										if v79 == int32(0) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v183 = m.ExcPending
											if v183 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v45)+4)) = v10
												*(*int32)(unsafe.Add(mBase, uint32(v45))) = v19
												F_errmsg_internal(m, int32(_a_F_binary_upgrade_set_missing_value_3), v45)
												mBase = m.M
												v188 = m.ExcPending
												if v188 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_binary_upgrade_set_missing_value_4), int32(2113), int32(_a_F_binary_upgrade_set_missing_value_5))
													mBase = m.M
													v193 = m.ExcPending
													if v193 != 0 {
														return int32(0)
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											}
										} else {
											v83 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
											v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+22)))
											v85 = v83 + v84
											v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+68))
											v87 = *(*int32)(unsafe.Add(mBase, uint32(v85)+76))
											v88 = m.G0
											v90 = v88 - int32(80)
											m.G0 = v90
											v94 = v90 + int32(8)
											v96 = *(*int32)(unsafe.Add(mBase, _c_F_binary_upgrade_set_missing_value[1]))
											F_fmgr_info_cxt_security(m, int32(750), v94, v96, int32(0))
											mBase = m.M
											v99 = m.ExcPending
											if v99 != 0 {
												return int32(0)
											} else {
												v100 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(v90)+76)) = uint8(v100)
												*(*int32)(unsafe.Add(mBase, uint32(v90)+72)) = v87
												*(*uint8)(unsafe.Add(mBase, uint32(v90)+68)) = uint8(v100)
												*(*int32)(unsafe.Add(mBase, uint32(v90)+64)) = v86
												*(*uint8)(unsafe.Add(mBase, uint32(v90)+60)) = uint8(v100)
												*(*int32)(unsafe.Add(mBase, uint32(v90)+56)) = v21
												v109 = int32(3)
												*(*uint16)(unsafe.Add(mBase, uint32(v90)+54)) = uint16(v109)
												*(*uint8)(unsafe.Add(mBase, uint32(v90)+52)) = uint8(v100)
												*(*int32)(unsafe.Add(mBase, uint32(v90)+48)) = v100
												*(*int64)(unsafe.Add(mBase, uint32(v90)+40)) = int64(0)
												*(*int32)(unsafe.Add(mBase, uint32(v90)+36)) = v94
												v120 = *(*int32)(unsafe.Add(mBase, uint32(v90)+8))
												v121 = m.T0[v120].(func(*base.Module, int32) int32)(m, v90+int32(36))
												mBase = m.M
												v122 = m.ExcPending
												if v122 != 0 {
													return int32(0)
												} else {
													v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90)+52)))
													if v123 == int32(1) {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v129 = m.ExcPending
														if v129 != 0 {
															return int32(0)
														} else {
															v130 = *(*int32)(unsafe.Add(mBase, uint32(v90)+12))
															*(*int32)(unsafe.Add(mBase, uint32(v90))) = v130
															F_errmsg_internal(m, int32(_a_F_binary_upgrade_set_missing_value_6), v90)
															mBase = m.M
															v134 = m.ExcPending
															if v134 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(_a_F_binary_upgrade_set_missing_value_7), int32(1190), int32(_a_F_binary_upgrade_set_missing_value_8))
																mBase = m.M
																v139 = m.ExcPending
																if v139 != 0 {
																	return int32(0)
																} else {
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															}
														}
													} else {
														v140 = int32(80)
														m.G0 = v90 + v140
														v143 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v45)+29)) = uint8(v143)
														*(*int32)(unsafe.Add(mBase, uint32(v45)+132)) = v143
														*(*int32)(unsafe.Add(mBase, uint32(v45)+176)) = v121
														*(*uint8)(unsafe.Add(mBase, uint32(v45)+40)) = uint8(v143)
														v150 = *(*int32)(unsafe.Add(mBase, uint32(v77)+52))
														v157 = F_heap_modify_tuple(m, v79, v150, v45+v140, v45+int32(48), v45+int32(16))
														mBase = m.M
														v158 = m.ExcPending
														if v158 != 0 {
															return int32(0)
														} else {
															F_CatalogTupleUpdate(m, v77, v157+int32(4), v157)
															mBase = m.M
															v162 = m.ExcPending
															if v162 != 0 {
																return int32(0)
															} else {
																F_ReleaseCatCache(m, v79)
																mBase = m.M
																v164 = m.ExcPending
																if v164 != 0 {
																	return int32(0)
																} else {
																	F_relation_close(m, v77, int32(3))
																	mBase = m.M
																	v167 = m.ExcPending
																	if v167 != 0 {
																		return int32(0)
																	} else {
																		F_relation_close(m, v69, int32(8))
																		mBase = m.M
																		v176 = m.ExcPending
																		if v176 != 0 {
																			return int32(0)
																		} else {
																			m.G0 = v45 + int32(192)
																			return int32(0)
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
							}
						}
					}
				}
			}
		}
	}
}
func F_binary_upgrade_set_next_index_relfilenode(m *base.Module, l0 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn13849(m, l0, int32(_a_F_binary_upgrade_set_next_index_relfilenode_0), int32(_a_F_binary_upgrade_set_next_index_relfilenode_1), int32(134))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_binary_upgrade_set_next_pg_enum_oid(m *base.Module, l0 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn13849(m, l0, int32(_a_F_binary_upgrade_set_next_pg_enum_oid_0), int32(_a_F_binary_upgrade_set_next_pg_enum_oid_1), int32(167))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
