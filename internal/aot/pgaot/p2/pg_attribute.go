package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pg_clear_attribute_stats(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int64
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	v2 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = v2
	F_stats_check_required_arg(m, l0, int32(_a_F_pg_clear_attribute_stats_0), v2)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int64(0)
	} else {
		F_stats_check_required_arg(m, l0, int32(_a_F_pg_clear_attribute_stats_0), int32(1))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int64(0)
		} else {
			F_stats_check_required_arg(m, l0, int32(_a_F_pg_clear_attribute_stats_0), int32(2))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int64(0)
			} else {
				F_stats_check_required_arg(m, l0, int32(_a_F_pg_clear_attribute_stats_0), int32(3))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int64(0)
				} else {
					v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
					v32 = F_text_to_cstring(m, v31)
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int64(0)
					} else {
						v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						v35 = F_text_to_cstring(m, v34)
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return int64(0)
						} else {
							v39 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_clear_attribute_stats[0])))
							if v39 == int32(1) {
								v44 = *(*int32)(unsafe.Add(mBase, _c_F_pg_clear_attribute_stats[1]))
								v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+308))
								v47 = base.B2i32(v45 != int32(2))
								*(*uint8)(unsafe.Add(mBase, _c_F_pg_clear_attribute_stats[0])) = uint8(v47)
								v49 = v47
							} else {
								v49 = int32(0)
							}
							if v49 == int32(0) {
								v53 = F_makeRangeVar(m, v32, v35, int32(-1))
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return int64(0)
								} else {
									v60 = F_RangeVarGetRelidExtended(m, v53, int32(4), int32(0), int32(1138), v9+int32(28))
									mBase = m.M
									v61 = m.ExcPending
									if v61 != 0 {
										return int64(0)
									} else {
										v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
										v63 = F_text_to_cstring(m, v62)
										mBase = m.M
										v64 = m.ExcPending
										if v64 != 0 {
											return int64(0)
										} else {
											v65 = F_get_attnum(m, v60, v63)
											mBase = m.M
											v66 = m.ExcPending
											if v66 != 0 {
												return int64(0)
											} else {
												if v65 < int32(0) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v123 = m.ExcPending
													if v123 != 0 {
														return int64(0)
													} else {
														F_errcode(m, int32(1088))
														mBase = m.M
														v126 = m.ExcPending
														if v126 != 0 {
															return int64(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v9))) = v63
															F_errmsg(m, int32(_a_F_pg_clear_attribute_stats_1), v9)
															mBase = m.M
															v130 = m.ExcPending
															if v130 != 0 {
																return int64(0)
															} else {
																F_errfinish(m, int32(_a_F_pg_clear_attribute_stats_2), int32(658), int32(_a_F_pg_clear_attribute_stats_3))
																mBase = m.M
																v135 = m.ExcPending
																if v135 != 0 {
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
													if v65 == int32(0) {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v139 = m.ExcPending
														if v139 != 0 {
															return int64(0)
														} else {
															F_errcode(m, int32(50360452))
															mBase = m.M
															v142 = m.ExcPending
															if v142 != 0 {
																return int64(0)
															} else {
																v143 = F_get_rel_name(m, v60)
																mBase = m.M
																v144 = m.ExcPending
																if v144 != 0 {
																	return int64(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v143
																	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v63
																	F_errmsg(m, int32(_a_F_pg_clear_attribute_stats_4), v9+int32(16))
																	mBase = m.M
																	v151 = m.ExcPending
																	if v151 != 0 {
																		return int64(0)
																	} else {
																		F_errfinish(m, int32(_a_F_pg_clear_attribute_stats_2), int32(664), int32(_a_F_pg_clear_attribute_stats_3))
																		mBase = m.M
																		v156 = m.ExcPending
																		if v156 != 0 {
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
													} else {
														v71 = *(*int64)(unsafe.Add(mBase, uint32(l0)+72))
														v74 = F_table_open(m, int32(2619), int32(3))
														mBase = m.M
														v75 = m.ExcPending
														if v75 != 0 {
															return int64(0)
														} else {
															v82 = F_SearchSysCache3(m, int32(65), base.I64_extend_i32_u(v60), base.I64_extend_i32_s(v65), base.I64_extend_i32_u(base.B2i32(v71 != int64(0))))
															mBase = m.M
															v83 = m.ExcPending
															if v83 != 0 {
																return int64(0)
															} else {
																if v82 != 0 {
																	F_simple_heap_delete(m, v74, v82+int32(4))
																	mBase = m.M
																	v87 = m.ExcPending
																	if v87 != 0 {
																		return int64(0)
																	} else {
																		F_ReleaseCatCache(m, v82)
																		mBase = m.M
																		v89 = m.ExcPending
																		if v89 != 0 {
																			return int64(0)
																		} else {
																			F_relation_close(m, v74, int32(3))
																			mBase = m.M
																			v92 = m.ExcPending
																			if v92 != 0 {
																				return int64(0)
																			} else {
																				F_CommandCounterIncrement(m)
																				mBase = m.M
																				v94 = m.ExcPending
																				if v94 != 0 {
																					return int64(0)
																				} else {
																					m.G0 = v9 + int32(32)
																					return int64(0)
																				}
																			}
																		}
																	}
																} else {
																	F_relation_close(m, v74, int32(3))
																	mBase = m.M
																	v92 = m.ExcPending
																	if v92 != 0 {
																		return int64(0)
																	} else {
																		F_CommandCounterIncrement(m)
																		mBase = m.M
																		v94 = m.ExcPending
																		if v94 != 0 {
																			return int64(0)
																		} else {
																			m.G0 = v9 + int32(32)
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
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v103 = m.ExcPending
								if v103 != 0 {
									return int64(0)
								} else {
									F_errcode(m, int32(325))
									mBase = m.M
									v106 = m.ExcPending
									if v106 != 0 {
										return int64(0)
									} else {
										F_errmsg(m, int32(_a_F_pg_clear_attribute_stats_5), int32(0))
										mBase = m.M
										v110 = m.ExcPending
										if v110 != 0 {
											return int64(0)
										} else {
											F_errhint(m, int32(_a_F_pg_clear_attribute_stats_6), int32(0))
											mBase = m.M
											v114 = m.ExcPending
											if v114 != 0 {
												return int64(0)
											} else {
												F_errfinish(m, int32(_a_F_pg_clear_attribute_stats_2), int32(645), int32(_a_F_pg_clear_attribute_stats_3))
												mBase = m.M
												v119 = m.ExcPending
												if v119 != 0 {
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
						}
					}
				}
			}
		}
	}
}
