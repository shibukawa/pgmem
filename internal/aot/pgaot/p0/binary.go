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
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	F_enlargeStringInfo(m, l0, l2)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if l2 != 0 {
			v9 = F__emscripten_memcpy_bulkmem(m, v6+v7, l1, l2)
			mBase = m.M
		} else {
		}
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v11 + l2
		return
	}
}
func F_binary_upgrade_set_missing_value(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
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
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int64
	_ = v56
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v114 int32
	_ = v114
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v157 int32
	_ = v157
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v16 = F_pg_detoast_datum(m, v15)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v18 = F_text_to_cstring(m, v11)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				v20 = F_text_to_cstring(m, v16)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					v23 = int32(*(*uint8)(unsafe.Add(mBase, _consts[200])))
					if v23 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(33685829))
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(392988), int32(0))
								mBase = m.M
								v36 = m.ExcPending
								if v36 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(470094), int32(268), int32(328376))
									mBase = m.M
									v41 = m.ExcPending
									if v41 != 0 {
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
						v42 = m.G0
						v44 = v42 - int32(192)
						m.G0 = v44
						v51 = F__emscripten_memset_bulkmem(m, v44+int32(80), base.I32_extend8_s(int32(0)), int32(96))
						mBase = m.M
						v52 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v44)+72)) = uint8(v52)
						v56 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v44-int32(-64)))) = v56
						*(*int64)(unsafe.Add(mBase, uint32(v44)+56)) = v56
						*(*int64)(unsafe.Add(mBase, uint32(v44)+48)) = v56
						*(*uint8)(unsafe.Add(mBase, uint32(v44)+40)) = uint8(v52)
						*(*int64)(unsafe.Add(mBase, uint32(v44)+32)) = v56
						*(*int64)(unsafe.Add(mBase, uint32(v44)+24)) = v56
						*(*int64)(unsafe.Add(mBase, uint32(v44)+16)) = v56
						v71 = F_table_open(m, v9, int32(8))
						mBase = m.M
						v72 = m.ExcPending
						if v72 != 0 {
							return int32(0)
						} else {
							v73 = *(*int32)(unsafe.Add(mBase, uint32(v71)+48))
							v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+119)))
							if v74 != int32(114) {
								F_sequence_close(m, v71, int32(8))
								mBase = m.M
								v79 = m.ExcPending
								if v79 != 0 {
									return int32(0)
								} else {
									m.G0 = v44 + int32(192)
									return int32(0)
								}
							} else {
								v82 = F_table_open(m, int32(1249), int32(3))
								mBase = m.M
								v83 = m.ExcPending
								if v83 != 0 {
									return int32(0)
								} else {
									v84 = F_SearchSysCacheAttName(m, v9, v18)
									mBase = m.M
									v85 = m.ExcPending
									if v85 != 0 {
										return int32(0)
									} else {
										if v84 == int32(0) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v190 = m.ExcPending
											if v190 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v44)+4)) = v9
												*(*int32)(unsafe.Add(mBase, uint32(v44))) = v18
												F_errmsg_internal(m, int32(43957), v44)
												mBase = m.M
												v195 = m.ExcPending
												if v195 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(473108), int32(2113), int32(313033))
													mBase = m.M
													v200 = m.ExcPending
													if v200 != 0 {
														return int32(0)
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											}
										} else {
											v88 = *(*int32)(unsafe.Add(mBase, uint32(v84)+16))
											v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+22)))
											v90 = v88 + v89
											v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)+68))
											v92 = *(*int32)(unsafe.Add(mBase, uint32(v90)+76))
											v93 = m.G0
											v95 = v93 - int32(80)
											m.G0 = v95
											v101 = *(*int32)(unsafe.Add(mBase, _consts[0]))
											F_fmgr_info_cxt_security(m, int32(750), v95+int32(8), v101, int32(0))
											mBase = m.M
											v104 = m.ExcPending
											if v104 != 0 {
												return int32(0)
											} else {
												v105 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(v95)+76)) = uint8(v105)
												*(*int32)(unsafe.Add(mBase, uint32(v95)+72)) = v92
												*(*uint8)(unsafe.Add(mBase, uint32(v95)+68)) = uint8(v105)
												*(*int32)(unsafe.Add(mBase, uint32(v95)+64)) = v91
												*(*uint8)(unsafe.Add(mBase, uint32(v95)+60)) = uint8(v105)
												*(*int32)(unsafe.Add(mBase, uint32(v95)+56)) = v20
												v114 = int32(3)
												*(*uint16)(unsafe.Add(mBase, uint32(v95)+54)) = uint16(v114)
												*(*uint8)(unsafe.Add(mBase, uint32(v95)+52)) = uint8(v105)
												*(*int32)(unsafe.Add(mBase, uint32(v95)+48)) = v105
												*(*int64)(unsafe.Add(mBase, uint32(v95)+40)) = int64(0)
												*(*int32)(unsafe.Add(mBase, uint32(v95)+36)) = v95 + int32(8)
												v127 = *(*int32)(unsafe.Add(mBase, uint32(v95)+8))
												v128 = m.T0[v127].(func(*base.Module, int32) int32)(m, v95+int32(36))
												mBase = m.M
												v129 = m.ExcPending
												if v129 != 0 {
													return int32(0)
												} else {
													v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95)+52)))
													if v130 == int32(1) {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v136 = m.ExcPending
														if v136 != 0 {
															return int32(0)
														} else {
															v137 = *(*int32)(unsafe.Add(mBase, uint32(v95)+12))
															*(*int32)(unsafe.Add(mBase, uint32(v95))) = v137
															F_errmsg_internal(m, int32(508732), v95)
															mBase = m.M
															v141 = m.ExcPending
															if v141 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(472279), int32(1190), int32(289223))
																mBase = m.M
																v146 = m.ExcPending
																if v146 != 0 {
																	return int32(0)
																} else {
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															}
														}
													} else {
														v147 = int32(80)
														m.G0 = v95 + v147
														v150 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v44)+29)) = uint8(v150)
														*(*int32)(unsafe.Add(mBase, uint32(v44)+132)) = v150
														*(*int32)(unsafe.Add(mBase, uint32(v44)+176)) = v128
														*(*uint8)(unsafe.Add(mBase, uint32(v44)+40)) = uint8(v150)
														v157 = *(*int32)(unsafe.Add(mBase, uint32(v82)+52))
														v164 = F_heap_modify_tuple(m, v84, v157, v44+v147, v44+int32(48), v44+int32(16))
														mBase = m.M
														v165 = m.ExcPending
														if v165 != 0 {
															return int32(0)
														} else {
															F_CatalogTupleUpdate(m, v82, v164+int32(4), v164)
															mBase = m.M
															v169 = m.ExcPending
															if v169 != 0 {
																return int32(0)
															} else {
																F_ReleaseCatCache(m, v84)
																mBase = m.M
																v171 = m.ExcPending
																if v171 != 0 {
																	return int32(0)
																} else {
																	F_sequence_close(m, v82, int32(3))
																	mBase = m.M
																	v174 = m.ExcPending
																	if v174 != 0 {
																		return int32(0)
																	} else {
																		F_sequence_close(m, v71, int32(8))
																		mBase = m.M
																		v177 = m.ExcPending
																		if v177 != 0 {
																			return int32(0)
																		} else {
																			m.G0 = v44 + int32(192)
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
	v3 = int32(*(*uint8)(unsafe.Add(mBase, _consts[200])))
	if v3 == int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(33685829))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(392988), int32(0))
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(470094), int32(134), int32(391419))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
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
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		*(*int32)(unsafe.Add(mBase, _consts[810])) = v25
		return int32(0)
	}
}
func F_binary_upgrade_set_next_pg_enum_oid(m *base.Module, l0 int32) int32 {
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
	v3 = int32(*(*uint8)(unsafe.Add(mBase, _consts[200])))
	if v3 == int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(33685829))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(392988), int32(0))
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(470094), int32(167), int32(413807))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
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
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		*(*int32)(unsafe.Add(mBase, _consts[811])) = v25
		return int32(0)
	}
}
