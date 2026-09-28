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
func F_binary_upgrade_set_missing_value(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v9 int64
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
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
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v53 int64
	_ = v53
	var v67 int32
	_ = v67
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
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int64
	_ = v88
	var v89 int64
	_ = v89
	var v90 int64
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	v9 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int64(0)
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
		v16 = F_pg_detoast_datum(m, v15)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int64(0)
		} else {
			v18 = F_text_to_cstring(m, v11)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int64(0)
			} else {
				v20 = F_text_to_cstring(m, v16)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int64(0)
				} else {
					v23 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_binary_upgrade_set_missing_value[0])))
					if v23 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return int64(0)
						} else {
							F_errcode(m, int32(33685829))
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
								return int64(0)
							} else {
								F_errmsg(m, int32(_a_F_binary_upgrade_set_missing_value_0), int32(0))
								mBase = m.M
								v36 = m.ExcPending
								if v36 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_binary_upgrade_set_missing_value_1), int32(269), int32(_a_F_binary_upgrade_set_missing_value_2))
									mBase = m.M
									v41 = m.ExcPending
									if v41 != 0 {
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
						v42 = m.G0
						v44 = v42 - int32(288)
						m.G0 = v44
						v48 = int32(0)
						base.MemoryFill(m, v44+int32(80), v48, int32(192))
						*(*uint8)(unsafe.Add(mBase, uint32(v44)+72)) = uint8(v48)
						v53 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v44)+64)) = v53
						*(*int64)(unsafe.Add(mBase, uint32(v44)+56)) = v53
						*(*int64)(unsafe.Add(mBase, uint32(v44)+48)) = v53
						*(*uint8)(unsafe.Add(mBase, uint32(v44)+40)) = uint8(v48)
						*(*int64)(unsafe.Add(mBase, uint32(v44)+32)) = v53
						*(*int64)(unsafe.Add(mBase, uint32(v44)+24)) = v53
						*(*int64)(unsafe.Add(mBase, uint32(v44)+16)) = v53
						v67 = base.I32_wrap_i64(v9)
						v69 = F_table_open(m, v67, int32(8))
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
							return int64(0)
						} else {
							v71 = *(*int32)(unsafe.Add(mBase, uint32(v69)+48))
							v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+119)))
							if v72 != int32(114) {
								F_relation_close(m, v69, int32(8))
								mBase = m.M
								v123 = m.ExcPending
								if v123 != 0 {
									return int64(0)
								} else {
									m.G0 = v44 + int32(288)
									return int64(0)
								}
							} else {
								v77 = F_table_open(m, int32(1249), int32(3))
								mBase = m.M
								v78 = m.ExcPending
								if v78 != 0 {
									return int64(0)
								} else {
									v79 = F_SearchSysCacheAttName(m, v67, v18)
									mBase = m.M
									v80 = m.ExcPending
									if v80 != 0 {
										return int64(0)
									} else {
										if v79 == int32(0) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v130 = m.ExcPending
											if v130 != 0 {
												return int64(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v44)+4)) = v67
												*(*int32)(unsafe.Add(mBase, uint32(v44))) = v18
												F_errmsg_internal(m, int32(_a_F_binary_upgrade_set_missing_value_3), v44)
												mBase = m.M
												v135 = m.ExcPending
												if v135 != 0 {
													return int64(0)
												} else {
													F_errfinish(m, int32(_a_F_binary_upgrade_set_missing_value_4), int32(2131), int32(_a_F_binary_upgrade_set_missing_value_5))
													mBase = m.M
													v140 = m.ExcPending
													if v140 != 0 {
														return int64(0)
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											}
										} else {
											v85 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
											v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+22)))
											v87 = v85 + v86
											v88 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v87)+68)))
											v89 = int64(*(*int32)(unsafe.Add(mBase, uint32(v87)+76)))
											v90 = F_OidFunctionCall3Coll(m, int32(750), base.I64_extend_i32_u(v20), v88, v89)
											mBase = m.M
											v91 = m.ExcPending
											if v91 != 0 {
												return int64(0)
											} else {
												v92 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v44)+29)) = uint8(v92)
												*(*int64)(unsafe.Add(mBase, uint32(v44)+184)) = int64(1)
												*(*int64)(unsafe.Add(mBase, uint32(v44)+272)) = v90
												*(*uint8)(unsafe.Add(mBase, uint32(v44)+40)) = uint8(v92)
												v99 = *(*int32)(unsafe.Add(mBase, uint32(v77)+52))
												v106 = F_heap_modify_tuple(m, v79, v99, v44+int32(80), v44+int32(48), v44+int32(16))
												mBase = m.M
												v107 = m.ExcPending
												if v107 != 0 {
													return int64(0)
												} else {
													F_CatalogTupleUpdate(m, v77, v106+int32(4), v106)
													mBase = m.M
													v111 = m.ExcPending
													if v111 != 0 {
														return int64(0)
													} else {
														F_ReleaseCatCache(m, v79)
														mBase = m.M
														v113 = m.ExcPending
														if v113 != 0 {
															return int64(0)
														} else {
															F_relation_close(m, v77, int32(3))
															mBase = m.M
															v116 = m.ExcPending
															if v116 != 0 {
																return int64(0)
															} else {
																F_relation_close(m, v69, int32(8))
																mBase = m.M
																v123 = m.ExcPending
																if v123 != 0 {
																	return int64(0)
																} else {
																	m.G0 = v44 + int32(288)
																	return int64(0)
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
func F_binary_upgrade_set_next_index_relfilenode(m *base.Module, l0 int32) int64 {
	var v5 int64
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn14232(m, l0, int32(_a_F_binary_upgrade_set_next_index_relfilenode_0), int32(_a_F_binary_upgrade_set_next_index_relfilenode_1), int32(135))
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		return v5
	}
}
func F_binary_upgrade_set_next_pg_enum_oid(m *base.Module, l0 int32) int64 {
	var v5 int64
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn14232(m, l0, int32(_a_F_binary_upgrade_set_next_pg_enum_oid_0), int32(_a_F_binary_upgrade_set_next_pg_enum_oid_1), int32(168))
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		return v5
	}
}
