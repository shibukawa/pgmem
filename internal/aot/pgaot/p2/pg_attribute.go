package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pg_clear_attribute_stats(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
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
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
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
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	v2 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	*(*int32)(unsafe.Add(mBase, uint32(v8)+28)) = v2
	F_stats_check_required_arg(m, l0, int32(4126304), v2)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		F_stats_check_required_arg(m, l0, int32(4126304), int32(1))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			F_stats_check_required_arg(m, l0, int32(4126304), int32(2))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				F_stats_check_required_arg(m, l0, int32(4126304), int32(3))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v31 = F_text_to_cstring(m, v30)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						v34 = F_text_to_cstring(m, v33)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int32(0)
						} else {
							v38 = int32(*(*uint8)(unsafe.Add(mBase, _consts[185])))
							if v38 == int32(1) {
								v43 = *(*int32)(unsafe.Add(mBase, _consts[178]))
								v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+316))
								v46 = base.B2i32(v44 != int32(2))
								*(*uint8)(unsafe.Add(mBase, _consts[185])) = uint8(v46)
								v48 = v46
							} else {
								v48 = int32(0)
							}
							if v48 == int32(0) {
								v52 = F_makeRangeVar(m, v31, v34, int32(-1))
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return int32(0)
								} else {
									v59 = F_RangeVarGetRelidExtended(m, v52, int32(4), int32(0), int32(1060), v8+int32(28))
									mBase = m.M
									v60 = m.ExcPending
									if v60 != 0 {
										return int32(0)
									} else {
										v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
										v62 = F_text_to_cstring(m, v61)
										mBase = m.M
										v63 = m.ExcPending
										if v63 != 0 {
											return int32(0)
										} else {
											v64 = F_get_attnum(m, v59, v62)
											mBase = m.M
											v65 = m.ExcPending
											if v65 != 0 {
												return int32(0)
											} else {
												if v64 < int32(0) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v119 = m.ExcPending
													if v119 != 0 {
														return int32(0)
													} else {
														F_errcode(m, int32(1088))
														mBase = m.M
														v122 = m.ExcPending
														if v122 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v8))) = v62
															F_errmsg(m, int32(712795), v8)
															mBase = m.M
															v126 = m.ExcPending
															if v126 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(494823), int32(950), int32(126360))
																mBase = m.M
																v131 = m.ExcPending
																if v131 != 0 {
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
													if v64 == int32(0) {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v135 = m.ExcPending
														if v135 != 0 {
															return int32(0)
														} else {
															F_errcode(m, int32(50360452))
															mBase = m.M
															v138 = m.ExcPending
															if v138 != 0 {
																return int32(0)
															} else {
																v139 = F_get_rel_name(m, v59)
																mBase = m.M
																v140 = m.ExcPending
																if v140 != 0 {
																	return int32(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v139
																	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v62
																	F_errmsg(m, int32(71908), v8+int32(16))
																	mBase = m.M
																	v147 = m.ExcPending
																	if v147 != 0 {
																		return int32(0)
																	} else {
																		F_errfinish(m, int32(494823), int32(956), int32(126360))
																		mBase = m.M
																		v152 = m.ExcPending
																		if v152 != 0 {
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
													} else {
														v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
														v73 = F_table_open(m, int32(2619), int32(3))
														mBase = m.M
														v74 = m.ExcPending
														if v74 != 0 {
															return int32(0)
														} else {
															v78 = F_SearchSysCache3(m, int32(65), v59, v64, base.B2i32(v70 != int32(0)))
															mBase = m.M
															v79 = m.ExcPending
															if v79 != 0 {
																return int32(0)
															} else {
																if v78 != 0 {
																	F_CatalogTupleDelete(m, v73, v78+int32(4))
																	mBase = m.M
																	v83 = m.ExcPending
																	if v83 != 0 {
																		return int32(0)
																	} else {
																		F_ReleaseCatCache(m, v78)
																		mBase = m.M
																		v85 = m.ExcPending
																		if v85 != 0 {
																			return int32(0)
																		} else {
																			F_sequence_close(m, v73, int32(3))
																			mBase = m.M
																			v88 = m.ExcPending
																			if v88 != 0 {
																				return int32(0)
																			} else {
																				F_CommandCounterIncrement(m)
																				mBase = m.M
																				v90 = m.ExcPending
																				if v90 != 0 {
																					return int32(0)
																				} else {
																					m.G0 = v8 + int32(32)
																					return int32(0)
																				}
																			}
																		}
																	}
																} else {
																	F_sequence_close(m, v73, int32(3))
																	mBase = m.M
																	v88 = m.ExcPending
																	if v88 != 0 {
																		return int32(0)
																	} else {
																		F_CommandCounterIncrement(m)
																		mBase = m.M
																		v90 = m.ExcPending
																		if v90 != 0 {
																			return int32(0)
																		} else {
																			m.G0 = v8 + int32(32)
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
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v99 = m.ExcPending
								if v99 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(325))
									mBase = m.M
									v102 = m.ExcPending
									if v102 != 0 {
										return int32(0)
									} else {
										F_errmsg(m, int32(128450), int32(0))
										mBase = m.M
										v106 = m.ExcPending
										if v106 != 0 {
											return int32(0)
										} else {
											F_errhint(m, int32(574572), int32(0))
											mBase = m.M
											v110 = m.ExcPending
											if v110 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(494823), int32(937), int32(126360))
												mBase = m.M
												v115 = m.ExcPending
												if v115 != 0 {
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
						}
					}
				}
			}
		}
	}
}
