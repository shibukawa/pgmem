package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_RI_FKey_check_upd(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	F_ri_CheckTrigger(m, l0, int32(_a_F_RI_FKey_check_upd_0), int32(2))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		F_RI_FKey_check(m, v8)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int64(0)
		} else {
			return int64(0)
		}
	}
}
func F_ri_HashCompareOp(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	v5 = m.G0
	v7 = v5 - int32(80)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_ri_HashCompareOp[0]))
	if v10 == int32(0) {
		*(*int64)(unsafe.Add(mBase, uint32(v7)+40)) = int64(3092376453124)
		v19 = v7 + int32(32)
		v21 = F_hash_create(m, int32(_a_F_ri_HashCompareOp_0), int64(64), v19, int32(40))
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_ri_HashCompareOp[1])) = v21
			F_CacheRegisterSyscacheCallback(m, int32(19), int32(1698), int64(0))
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				F_CacheRegisterSyscacheCallback(m, int32(3), int32(1698), int64(0))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return int32(0)
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v7)+40)) = int64(51539607560)
					v42 = F_hash_create(m, int32(_a_F_ri_HashCompareOp_1), int64(256), v19, int32(40))
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, _c_F_ri_HashCompareOp[2])) = v42
						*(*int64)(unsafe.Add(mBase, uint32(v7)+40)) = int64(292057776136)
						v51 = F_hash_create(m, int32(_a_F_ri_HashCompareOp_2), int64(256), v19, int32(40))
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, _c_F_ri_HashCompareOp[0])) = v51
							v54 = v51
							*(*int32)(unsafe.Add(mBase, uint32(v7)+36)) = l1
							*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = l0
							v62 = F_hash_search(m, v54, v7+int32(32), int32(1), v7+int32(31))
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return int32(0)
							} else {
								v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+31)))
								if v64 == int32(0) {
									v67 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v62)+8)) = uint8(v67)
									v70 = F_get_opcode(m, l0)
									mBase = m.M
									v71 = m.ExcPending
									if v71 != 0 {
										return int32(0)
									} else {
										v75 = *(*int32)(unsafe.Add(mBase, _c_F_ri_HashCompareOp[3]))
										F_fmgr_info_cxt(m, v70, v62+int32(12), v75)
										mBase = m.M
										v77 = m.ExcPending
										if v77 != 0 {
											return int32(0)
										} else {
											F_op_input_types(m, l0, v7+int32(24), v7+int32(20))
											mBase = m.M
											v83 = m.ExcPending
											if v83 != 0 {
												return int32(0)
											} else {
												v84 = F_getBaseType(m, l1)
												mBase = m.M
												v85 = m.ExcPending
												if v85 != 0 {
													return int32(0)
												} else {
													v86 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
													if v84 == v86 {
														*(*int32)(unsafe.Add(mBase, uint32(v62)+44)) = int32(0)
														v116 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v62)+8)) = uint8(v116)
														m.G0 = v7 + int32(80)
														return v62
													} else {
														v88 = *(*int32)(unsafe.Add(mBase, uint32(v7)+24))
														v92 = F_find_coercion_pathway(m, v88, l1, int32(0), v7+int32(16))
														mBase = m.M
														v93 = m.ExcPending
														if v93 != 0 {
															return int32(0)
														} else {
															if base.Ui32(v92-int32(3)) <= base.Ui32(int32(-3)) {
																v98 = *(*int32)(unsafe.Add(mBase, uint32(v7)+24))
																v99 = F_IsBinaryCoercible(m, l1, v98)
																mBase = m.M
																v100 = m.ExcPending
																if v100 != 0 {
																	return int32(0)
																} else {
																	if v99 == int32(0) {
																		F_errstart_cold(m, int32(21), int32(0))
																		mBase = m.M
																		v126 = m.ExcPending
																		if v126 != 0 {
																			return int32(0)
																		} else {
																			v127 = F_format_type_be(m, l1)
																			mBase = m.M
																			v128 = m.ExcPending
																			if v128 != 0 {
																				return int32(0)
																			} else {
																				v129 = *(*int32)(unsafe.Add(mBase, uint32(v7)+24))
																				v130 = F_format_type_be(m, v129)
																				mBase = m.M
																				v131 = m.ExcPending
																				if v131 != 0 {
																					return int32(0)
																				} else {
																					*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v130
																					*(*int32)(unsafe.Add(mBase, uint32(v7))) = v127
																					F_errmsg_internal(m, int32(_a_F_ri_HashCompareOp_3), v7)
																					mBase = m.M
																					v136 = m.ExcPending
																					if v136 != 0 {
																						return int32(0)
																					} else {
																						F_errfinish(m, int32(_a_F_ri_HashCompareOp_4), int32(3902), int32(_a_F_ri_HashCompareOp_5))
																						mBase = m.M
																						v141 = m.ExcPending
																						if v141 != 0 {
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
																		v103 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
																		if v103 == int32(0) {
																			*(*int32)(unsafe.Add(mBase, uint32(v62)+44)) = int32(0)
																			v116 = int32(1)
																			*(*uint8)(unsafe.Add(mBase, uint32(v62)+8)) = uint8(v116)
																			m.G0 = v7 + int32(80)
																			return v62
																		} else {
																			v109 = *(*int32)(unsafe.Add(mBase, _c_F_ri_HashCompareOp[3]))
																			F_fmgr_info_cxt(m, v103, v62+int32(40), v109)
																			mBase = m.M
																			v111 = m.ExcPending
																			if v111 != 0 {
																				return int32(0)
																			} else {
																				v116 = int32(1)
																				*(*uint8)(unsafe.Add(mBase, uint32(v62)+8)) = uint8(v116)
																				m.G0 = v7 + int32(80)
																				return v62
																			}
																		}
																	}
																}
															} else {
																v103 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
																if v103 == int32(0) {
																	*(*int32)(unsafe.Add(mBase, uint32(v62)+44)) = int32(0)
																	v116 = int32(1)
																	*(*uint8)(unsafe.Add(mBase, uint32(v62)+8)) = uint8(v116)
																	m.G0 = v7 + int32(80)
																	return v62
																} else {
																	v109 = *(*int32)(unsafe.Add(mBase, _c_F_ri_HashCompareOp[3]))
																	F_fmgr_info_cxt(m, v103, v62+int32(40), v109)
																	mBase = m.M
																	v111 = m.ExcPending
																	if v111 != 0 {
																		return int32(0)
																	} else {
																		v116 = int32(1)
																		*(*uint8)(unsafe.Add(mBase, uint32(v62)+8)) = uint8(v116)
																		m.G0 = v7 + int32(80)
																		return v62
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
									v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+8)))
									if v69 != 0 {
										m.G0 = v7 + int32(80)
										return v62
									} else {
										v70 = F_get_opcode(m, l0)
										mBase = m.M
										v71 = m.ExcPending
										if v71 != 0 {
											return int32(0)
										} else {
											v75 = *(*int32)(unsafe.Add(mBase, _c_F_ri_HashCompareOp[3]))
											F_fmgr_info_cxt(m, v70, v62+int32(12), v75)
											mBase = m.M
											v77 = m.ExcPending
											if v77 != 0 {
												return int32(0)
											} else {
												F_op_input_types(m, l0, v7+int32(24), v7+int32(20))
												mBase = m.M
												v83 = m.ExcPending
												if v83 != 0 {
													return int32(0)
												} else {
													v84 = F_getBaseType(m, l1)
													mBase = m.M
													v85 = m.ExcPending
													if v85 != 0 {
														return int32(0)
													} else {
														v86 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
														if v84 == v86 {
															*(*int32)(unsafe.Add(mBase, uint32(v62)+44)) = int32(0)
															v116 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(v62)+8)) = uint8(v116)
															m.G0 = v7 + int32(80)
															return v62
														} else {
															v88 = *(*int32)(unsafe.Add(mBase, uint32(v7)+24))
															v92 = F_find_coercion_pathway(m, v88, l1, int32(0), v7+int32(16))
															mBase = m.M
															v93 = m.ExcPending
															if v93 != 0 {
																return int32(0)
															} else {
																if base.Ui32(v92-int32(3)) <= base.Ui32(int32(-3)) {
																	v98 = *(*int32)(unsafe.Add(mBase, uint32(v7)+24))
																	v99 = F_IsBinaryCoercible(m, l1, v98)
																	mBase = m.M
																	v100 = m.ExcPending
																	if v100 != 0 {
																		return int32(0)
																	} else {
																		if v99 == int32(0) {
																			F_errstart_cold(m, int32(21), int32(0))
																			mBase = m.M
																			v126 = m.ExcPending
																			if v126 != 0 {
																				return int32(0)
																			} else {
																				v127 = F_format_type_be(m, l1)
																				mBase = m.M
																				v128 = m.ExcPending
																				if v128 != 0 {
																					return int32(0)
																				} else {
																					v129 = *(*int32)(unsafe.Add(mBase, uint32(v7)+24))
																					v130 = F_format_type_be(m, v129)
																					mBase = m.M
																					v131 = m.ExcPending
																					if v131 != 0 {
																						return int32(0)
																					} else {
																						*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v130
																						*(*int32)(unsafe.Add(mBase, uint32(v7))) = v127
																						F_errmsg_internal(m, int32(_a_F_ri_HashCompareOp_3), v7)
																						mBase = m.M
																						v136 = m.ExcPending
																						if v136 != 0 {
																							return int32(0)
																						} else {
																							F_errfinish(m, int32(_a_F_ri_HashCompareOp_4), int32(3902), int32(_a_F_ri_HashCompareOp_5))
																							mBase = m.M
																							v141 = m.ExcPending
																							if v141 != 0 {
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
																			v103 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
																			if v103 == int32(0) {
																				*(*int32)(unsafe.Add(mBase, uint32(v62)+44)) = int32(0)
																				v116 = int32(1)
																				*(*uint8)(unsafe.Add(mBase, uint32(v62)+8)) = uint8(v116)
																				m.G0 = v7 + int32(80)
																				return v62
																			} else {
																				v109 = *(*int32)(unsafe.Add(mBase, _c_F_ri_HashCompareOp[3]))
																				F_fmgr_info_cxt(m, v103, v62+int32(40), v109)
																				mBase = m.M
																				v111 = m.ExcPending
																				if v111 != 0 {
																					return int32(0)
																				} else {
																					v116 = int32(1)
																					*(*uint8)(unsafe.Add(mBase, uint32(v62)+8)) = uint8(v116)
																					m.G0 = v7 + int32(80)
																					return v62
																				}
																			}
																		}
																	}
																} else {
																	v103 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
																	if v103 == int32(0) {
																		*(*int32)(unsafe.Add(mBase, uint32(v62)+44)) = int32(0)
																		v116 = int32(1)
																		*(*uint8)(unsafe.Add(mBase, uint32(v62)+8)) = uint8(v116)
																		m.G0 = v7 + int32(80)
																		return v62
																	} else {
																		v109 = *(*int32)(unsafe.Add(mBase, _c_F_ri_HashCompareOp[3]))
																		F_fmgr_info_cxt(m, v103, v62+int32(40), v109)
																		mBase = m.M
																		v111 = m.ExcPending
																		if v111 != 0 {
																			return int32(0)
																		} else {
																			v116 = int32(1)
																			*(*uint8)(unsafe.Add(mBase, uint32(v62)+8)) = uint8(v116)
																			m.G0 = v7 + int32(80)
																			return v62
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
	} else {
		v54 = v10
		*(*int32)(unsafe.Add(mBase, uint32(v7)+36)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = l0
		v62 = F_hash_search(m, v54, v7+int32(32), int32(1), v7+int32(31))
		mBase = m.M
		v63 = m.ExcPending
		if v63 != 0 {
			return int32(0)
		} else {
			v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+31)))
			if v64 == int32(0) {
				v67 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v62)+8)) = uint8(v67)
				v70 = F_get_opcode(m, l0)
				mBase = m.M
				v71 = m.ExcPending
				if v71 != 0 {
					return int32(0)
				} else {
					v75 = *(*int32)(unsafe.Add(mBase, _c_F_ri_HashCompareOp[3]))
					F_fmgr_info_cxt(m, v70, v62+int32(12), v75)
					mBase = m.M
					v77 = m.ExcPending
					if v77 != 0 {
						return int32(0)
					} else {
						F_op_input_types(m, l0, v7+int32(24), v7+int32(20))
						mBase = m.M
						v83 = m.ExcPending
						if v83 != 0 {
							return int32(0)
						} else {
							v84 = F_getBaseType(m, l1)
							mBase = m.M
							v85 = m.ExcPending
							if v85 != 0 {
								return int32(0)
							} else {
								v86 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
								if v84 == v86 {
									*(*int32)(unsafe.Add(mBase, uint32(v62)+44)) = int32(0)
									v116 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v62)+8)) = uint8(v116)
									m.G0 = v7 + int32(80)
									return v62
								} else {
									v88 = *(*int32)(unsafe.Add(mBase, uint32(v7)+24))
									v92 = F_find_coercion_pathway(m, v88, l1, int32(0), v7+int32(16))
									mBase = m.M
									v93 = m.ExcPending
									if v93 != 0 {
										return int32(0)
									} else {
										if base.Ui32(v92-int32(3)) <= base.Ui32(int32(-3)) {
											v98 = *(*int32)(unsafe.Add(mBase, uint32(v7)+24))
											v99 = F_IsBinaryCoercible(m, l1, v98)
											mBase = m.M
											v100 = m.ExcPending
											if v100 != 0 {
												return int32(0)
											} else {
												if v99 == int32(0) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v126 = m.ExcPending
													if v126 != 0 {
														return int32(0)
													} else {
														v127 = F_format_type_be(m, l1)
														mBase = m.M
														v128 = m.ExcPending
														if v128 != 0 {
															return int32(0)
														} else {
															v129 = *(*int32)(unsafe.Add(mBase, uint32(v7)+24))
															v130 = F_format_type_be(m, v129)
															mBase = m.M
															v131 = m.ExcPending
															if v131 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v130
																*(*int32)(unsafe.Add(mBase, uint32(v7))) = v127
																F_errmsg_internal(m, int32(_a_F_ri_HashCompareOp_3), v7)
																mBase = m.M
																v136 = m.ExcPending
																if v136 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(_a_F_ri_HashCompareOp_4), int32(3902), int32(_a_F_ri_HashCompareOp_5))
																	mBase = m.M
																	v141 = m.ExcPending
																	if v141 != 0 {
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
													v103 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
													if v103 == int32(0) {
														*(*int32)(unsafe.Add(mBase, uint32(v62)+44)) = int32(0)
														v116 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v62)+8)) = uint8(v116)
														m.G0 = v7 + int32(80)
														return v62
													} else {
														v109 = *(*int32)(unsafe.Add(mBase, _c_F_ri_HashCompareOp[3]))
														F_fmgr_info_cxt(m, v103, v62+int32(40), v109)
														mBase = m.M
														v111 = m.ExcPending
														if v111 != 0 {
															return int32(0)
														} else {
															v116 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(v62)+8)) = uint8(v116)
															m.G0 = v7 + int32(80)
															return v62
														}
													}
												}
											}
										} else {
											v103 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
											if v103 == int32(0) {
												*(*int32)(unsafe.Add(mBase, uint32(v62)+44)) = int32(0)
												v116 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v62)+8)) = uint8(v116)
												m.G0 = v7 + int32(80)
												return v62
											} else {
												v109 = *(*int32)(unsafe.Add(mBase, _c_F_ri_HashCompareOp[3]))
												F_fmgr_info_cxt(m, v103, v62+int32(40), v109)
												mBase = m.M
												v111 = m.ExcPending
												if v111 != 0 {
													return int32(0)
												} else {
													v116 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v62)+8)) = uint8(v116)
													m.G0 = v7 + int32(80)
													return v62
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
				v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+8)))
				if v69 != 0 {
					m.G0 = v7 + int32(80)
					return v62
				} else {
					v70 = F_get_opcode(m, l0)
					mBase = m.M
					v71 = m.ExcPending
					if v71 != 0 {
						return int32(0)
					} else {
						v75 = *(*int32)(unsafe.Add(mBase, _c_F_ri_HashCompareOp[3]))
						F_fmgr_info_cxt(m, v70, v62+int32(12), v75)
						mBase = m.M
						v77 = m.ExcPending
						if v77 != 0 {
							return int32(0)
						} else {
							F_op_input_types(m, l0, v7+int32(24), v7+int32(20))
							mBase = m.M
							v83 = m.ExcPending
							if v83 != 0 {
								return int32(0)
							} else {
								v84 = F_getBaseType(m, l1)
								mBase = m.M
								v85 = m.ExcPending
								if v85 != 0 {
									return int32(0)
								} else {
									v86 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
									if v84 == v86 {
										*(*int32)(unsafe.Add(mBase, uint32(v62)+44)) = int32(0)
										v116 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v62)+8)) = uint8(v116)
										m.G0 = v7 + int32(80)
										return v62
									} else {
										v88 = *(*int32)(unsafe.Add(mBase, uint32(v7)+24))
										v92 = F_find_coercion_pathway(m, v88, l1, int32(0), v7+int32(16))
										mBase = m.M
										v93 = m.ExcPending
										if v93 != 0 {
											return int32(0)
										} else {
											if base.Ui32(v92-int32(3)) <= base.Ui32(int32(-3)) {
												v98 = *(*int32)(unsafe.Add(mBase, uint32(v7)+24))
												v99 = F_IsBinaryCoercible(m, l1, v98)
												mBase = m.M
												v100 = m.ExcPending
												if v100 != 0 {
													return int32(0)
												} else {
													if v99 == int32(0) {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v126 = m.ExcPending
														if v126 != 0 {
															return int32(0)
														} else {
															v127 = F_format_type_be(m, l1)
															mBase = m.M
															v128 = m.ExcPending
															if v128 != 0 {
																return int32(0)
															} else {
																v129 = *(*int32)(unsafe.Add(mBase, uint32(v7)+24))
																v130 = F_format_type_be(m, v129)
																mBase = m.M
																v131 = m.ExcPending
																if v131 != 0 {
																	return int32(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v130
																	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v127
																	F_errmsg_internal(m, int32(_a_F_ri_HashCompareOp_3), v7)
																	mBase = m.M
																	v136 = m.ExcPending
																	if v136 != 0 {
																		return int32(0)
																	} else {
																		F_errfinish(m, int32(_a_F_ri_HashCompareOp_4), int32(3902), int32(_a_F_ri_HashCompareOp_5))
																		mBase = m.M
																		v141 = m.ExcPending
																		if v141 != 0 {
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
														v103 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
														if v103 == int32(0) {
															*(*int32)(unsafe.Add(mBase, uint32(v62)+44)) = int32(0)
															v116 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(v62)+8)) = uint8(v116)
															m.G0 = v7 + int32(80)
															return v62
														} else {
															v109 = *(*int32)(unsafe.Add(mBase, _c_F_ri_HashCompareOp[3]))
															F_fmgr_info_cxt(m, v103, v62+int32(40), v109)
															mBase = m.M
															v111 = m.ExcPending
															if v111 != 0 {
																return int32(0)
															} else {
																v116 = int32(1)
																*(*uint8)(unsafe.Add(mBase, uint32(v62)+8)) = uint8(v116)
																m.G0 = v7 + int32(80)
																return v62
															}
														}
													}
												}
											} else {
												v103 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
												if v103 == int32(0) {
													*(*int32)(unsafe.Add(mBase, uint32(v62)+44)) = int32(0)
													v116 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v62)+8)) = uint8(v116)
													m.G0 = v7 + int32(80)
													return v62
												} else {
													v109 = *(*int32)(unsafe.Add(mBase, _c_F_ri_HashCompareOp[3]))
													F_fmgr_info_cxt(m, v103, v62+int32(40), v109)
													mBase = m.M
													v111 = m.ExcPending
													if v111 != 0 {
														return int32(0)
													} else {
														v116 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v62)+8)) = uint8(v116)
														m.G0 = v7 + int32(80)
														return v62
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
func F_ri_KeysEqual(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v35 int32
	_ = v35
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int64
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int64
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
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
	var v95 int32
	_ = v95
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v118 int64
	_ = v118
	var v119 int32
	_ = v119
	var v123 int64
	_ = v123
	var v124 int32
	_ = v124
	var v126 int64
	_ = v126
	var v127 int64
	_ = v127
	var v130 int64
	_ = v130
	var v131 int32
	_ = v131
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v153 int32
	_ = v153
	v6 = int32(0)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l3)+168))
	if v17 <= v6 {
		v153 = int32(1)
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v153
L2:
	;
	if l4 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v22 = int32(172)
	goto L5
L4:
	;
	v22 = int32(236)
	goto L5
L5:
	;
	v35 = v6
	goto L6
L6:
	;
	v45 = l3 + v22 + v35<<(uint(int32(1))%32)
	v46 = int32(*(*int16)(unsafe.Add(mBase, uint32(v45))))
	v47 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+6)))
	if v47 < v46 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v153 = v139
	goto L1
L8:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+16))
	m.T0[v50].(func(*base.Module, int32, int32))(m, l1, v46)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v55 = int32(0)
	v57 = v46 - int32(1)
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57+v58))))
	if v60 != 0 {
		v153 = v55
		goto L1
	} else {
		goto L13
	}
L11:
	;
	return int32(0)
L12:
	;
	goto L10
L13:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v65 = *(*int64)(unsafe.Add(mBase, uint32(v61+v57<<(uint(int32(3))%32))))
	v66 = int32(*(*int16)(unsafe.Add(mBase, uint32(v45))))
	v67 = int32(*(*int16)(unsafe.Add(mBase, uint32(l2)+6)))
	if v67 < v66 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+16))
	m.T0[v70].(func(*base.Module, int32, int32))(m, l2, v66)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L11
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v74 = v66 - int32(1)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74+v75))))
	if v77 != 0 {
		v153 = v55
		goto L1
	} else {
		goto L18
	}
L17:
	;
	goto L16
L18:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v82 = *(*int64)(unsafe.Add(mBase, uint32(v78+v74<<(uint(int32(3))%32))))
	if l4 != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v139 = int32(1)
	v141 = v35 + v139
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l3)+168))
	if v141 < v142 {
		v35 = v141
		goto L6
	} else {
		goto L40
	}
L20:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v84 = int32(*(*int16)(unsafe.Add(mBase, uint32(v45))))
	v87 = v83 + v84<<(uint(int32(3))%32)
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+24)))
	v89 = int32(*(*int16)(unsafe.Add(mBase, uint32(v87)+22)))
	v90 = F_datum_image_eq(m, v65, v82, v88, v89)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L11
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+165)))
	if v92 == int32(1) {
		goto L26
	} else {
		goto L27
	}
L23:
	;
	if v90 != 0 {
		goto L19
	} else {
		goto L24
	}
L24:
	;
	v153 = v55
	goto L1
L25:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)))
	v104 = int32(*(*int16)(unsafe.Add(mBase, uint32(v45))))
	v105 = F_attnumTypeId(m, l0, v104)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L11
	} else {
		goto L30
	}
L26:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l3)+168))
	if v35 == v95-int32(1) {
		v102 = l3 + int32(684)
		goto L25
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v102 = l3 + int32(556) + v35<<(uint(int32(2))%32)
	goto L25
L29:
	;
	goto L28
L30:
	;
	v107 = int32(*(*int16)(unsafe.Add(mBase, uint32(v45))))
	v108 = F_attnumCollationId(m, l0, v107)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L11
	} else {
		goto L31
	}
L31:
	;
	v110 = F_ri_HashCompareOp(m, v103, v105)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L11
	} else {
		goto L32
	}
L32:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v110)+44))
	if v112 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v114 = v110 + int32(40)
	v118 = F_FunctionCall3Coll(m, v114, int32(0), v82, int64(-1), int64(0))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L11
	} else {
		goto L36
	}
L34:
	;
	v126 = v65
	v127 = v82
	goto L35
L35:
	;
	v130 = F_FunctionCall2Coll(m, v110+int32(12), v108, v127, v126)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L11
	} else {
		goto L38
	}
L36:
	;
	v123 = F_FunctionCall3Coll(m, v114, int32(0), v65, int64(-1), int64(0))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L11
	} else {
		goto L37
	}
L37:
	;
	v126 = v123
	v127 = v118
	goto L35
L38:
	;
	if v130 == int64(0) {
		v153 = v55
		goto L1
	} else {
		goto L39
	}
L39:
	;
	goto L19
L40:
	;
	goto L7
}
func F_ri_set(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v140 int32
	_ = v140
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v188 int32
	_ = v188
	var v194 int32
	_ = v194
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v279 int32
	_ = v279
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v341 int32
	_ = v341
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v383 int32
	_ = v383
	var v389 int32
	_ = v389
	var v395 int32
	_ = v395
	var v408 int32
	_ = v408
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v445 int32
	_ = v445
	var v450 int32
	_ = v450
	var v457 int32
	_ = v457
	var v461 int32
	_ = v461
	var v466 int32
	_ = v466
	v18 = m.G0
	v20 = v18 - int32(640)
	m.G0 = v20
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v25 = F_ri_FetchConstraintInfo(m, v22, v23, int32(1))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+88))
	v29 = F_table_open(m, v27, int32(3))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_SPI_connect_ext(m, int32(0))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	if l1 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v39 = int32(8)
	goto L7
L6:
	;
	v39 = int32(10)
	goto L7
L7:
	;
	if l1 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v42 = int32(7)
	goto L10
L9:
	;
	v42 = int32(9)
	goto L10
L10:
	;
	v44 = base.B2i32(l2 == int32(2))
	if l2 == int32(2) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v45 = v39
	goto L13
L12:
	;
	v45 = v42
	goto L13
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+636)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v20)+632)) = v36
	v49 = v20 + int32(632)
	v50 = F_ri_FetchPreparedPlan(m, v49)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	if v50 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v432 = v50
	goto L17
L16:
	;
	if v44 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	v433 = int32(0)
	v437 = F_ri_PerformCheck(m, v25, v49, v432, v29, v32, v31, v433, v433, int32(1), int32(9))
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L1
	} else {
		goto L85
	}
L18:
	;
	F_initStringInfo(m, v20+int32(616))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L23
	}
L19:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v25)+96))
	if v55 != 0 {
		v59 = v55
		v60 = int32(100)
		goto L18
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v25)+168))
	v59 = v57
	v60 = int32(236)
	goto L18
L22:
	;
	goto L21
L23:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v29)+48))
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+119)))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v65)+68))
	v68 = F_get_namespace_name(m, v67)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v70 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+352)) = uint8(v70)
	v76 = v68
	v78 = v20 + int32(352)
	goto L25
L25:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76))))
	if v91 != int32(34) {
		goto L29
	} else {
		goto L30
	}
L26:
	;
	v108 = int32(34)
	*(*uint16)(unsafe.Add(mBase, uint32(v78)+1)) = uint16(v108)
	v111 = v20 + int32(352)
	v112 = F_strlen(m, v111)
	mBase = m.M
	v113 = v112 + v111
	v114 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v113))) = uint8(v114)
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v29)+48))
	*(*uint8)(unsafe.Add(mBase, uint32(v113)+1)) = uint8(v108)
	v125 = v113 + int32(1)
	v127 = v116 + int32(4)
	goto L33
L27:
	;
	goto L26
L28:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v104))) = uint8(v103)
	v76 = v76 + int32(1)
	v78 = v104
	goto L25
L29:
	;
	if v91 == int32(0) {
		goto L27
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v98 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v78)+1)) = uint8(v98)
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76))))
	v103 = v100
	v104 = v78 + int32(2)
	goto L28
L32:
	;
	v103 = v91
	v104 = v78 + int32(1)
	goto L28
L33:
	;
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
	if v140 != int32(34) {
		goto L37
	} else {
		goto L38
	}
L34:
	;
	v157 = int32(34)
	*(*uint16)(unsafe.Add(mBase, uint32(v125)+1)) = uint16(v157)
	v159 = int32(_a_F_ri_set_0)
	if v66 == int32(112) {
		goto L41
	} else {
		goto L42
	}
L35:
	;
	goto L34
L36:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v153))) = uint8(v152)
	v125 = v153
	v127 = v127 + int32(1)
	goto L33
L37:
	;
	if v140 == int32(0) {
		goto L35
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v147 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v125)+1)) = uint8(v147)
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
	v152 = v149
	v153 = v125 + int32(2)
	goto L36
L40:
	;
	v152 = v140
	v153 = v125 + int32(1)
	goto L36
L41:
	;
	v164 = v159
	goto L43
L42:
	;
	v164 = int32(_a_F_ri_set_1)
	goto L43
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+48)) = v164
	*(*int32)(unsafe.Add(mBase, uint32(v20)+52)) = v20 + int32(352)
	F_appendStringInfo(m, v20+int32(616), int32(_a_F_ri_set_2), v20+int32(48))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	if v59 <= int32(0) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v25)+168))
	if int32(0) < v279 {
		goto L64
	} else {
		goto L65
	}
L46:
	;
	if l1 != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v180 = int32(_a_F_ri_set_3)
	goto L49
L48:
	;
	v180 = int32(_a_F_ri_set_4)
	goto L49
L49:
	;
	v188 = int32(0)
	v194 = v159
	goto L50
L50:
	;
	v203 = int32(*(*int16)(unsafe.Add(mBase, uint32(v60+v25+v188<<(uint(int32(1))%32)))))
	v204 = F_attnumAttName(m, v29, v203)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v206 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+208)) = uint8(v206)
	v212 = v204
	v214 = v20 + int32(208)
	goto L53
L53:
	;
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212))))
	if v227 != int32(34) {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v258))) = uint8(v257)
	v212 = v212 + int32(1)
	v214 = v258
	goto L53
L56:
	;
	if v227 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	goto L58
L58:
	;
	v252 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v214)+1)) = uint8(v252)
	v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212))))
	v257 = v254
	v258 = v214 + int32(2)
	goto L55
L59:
	;
	v232 = int32(34)
	*(*uint16)(unsafe.Add(mBase, uint32(v214)+1)) = uint16(v232)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+40)) = v180
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v194
	*(*int32)(unsafe.Add(mBase, uint32(v20)+36)) = v20 + int32(208)
	F_appendStringInfo(m, v20+int32(616), int32(_a_F_ri_set_5), v20+int32(32))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L1
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	v257 = v227
	v258 = v214 + int32(1)
	goto L55
L62:
	;
	v248 = v188 + int32(1)
	if v248 != v59 {
		v188 = v248
		v194 = int32(_a_F_ri_set_6)
		goto L50
	} else {
		goto L63
	}
L63:
	;
	goto L45
L64:
	;
	v295 = int32(0)
	v299 = int32(_a_F_ri_set_7)
	goto L67
L65:
	;
	v395 = v279
	goto L66
L66:
	;
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v20)+616))
	v413 = F_ri_PlanCheck(m, v408, v395, v20-int32(-64), v20+int32(632), v29, v32)
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L1
	} else {
		goto L84
	}
L67:
	;
	v308 = v295 << (uint(int32(1)) % 32)
	v310 = int32(*(*int16)(unsafe.Add(mBase, uint32(v25+int32(172)+v308))))
	v311 = F_attnumTypeId(m, v32, v310)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L1
	} else {
		goto L69
	}
L68:
	;
	v395 = v389
	goto L66
L69:
	;
	v313 = v308 + (v25 + int32(236))
	v314 = int32(*(*int16)(unsafe.Add(mBase, uint32(v313))))
	v315 = F_attnumTypeId(m, v29, v314)
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	v317 = int32(*(*int16)(unsafe.Add(mBase, uint32(v313))))
	v318 = F_attnumAttName(m, v29, v317)
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	v320 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+208)) = uint8(v320)
	v326 = v318
	v328 = v20 + int32(208)
	goto L72
L72:
	;
	v341 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v326))))
	if v341 != int32(34) {
		goto L76
	} else {
		goto L77
	}
L73:
	;
	v358 = int32(34)
	*(*uint16)(unsafe.Add(mBase, uint32(v328)+1)) = uint16(v358)
	v361 = v295 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v361
	v364 = v20 + int32(192)
	v368 = F_pg_sprintf(m, v364, int32(_a_F_ri_set_8), v20+int32(16))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L1
	} else {
		goto L80
	}
L74:
	;
	goto L73
L75:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v354))) = uint8(v353)
	v326 = v326 + int32(1)
	v328 = v354
	goto L72
L76:
	;
	if v341 == int32(0) {
		goto L74
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	v348 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v328)+1)) = uint8(v348)
	v350 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v326))))
	v353 = v350
	v354 = v328 + int32(2)
	goto L75
L79:
	;
	v353 = v341
	v354 = v328 + int32(1)
	goto L75
L80:
	;
	v371 = v295 << (uint(int32(2)) % 32)
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v25+int32(300)+v371)))
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v299
	v376 = v20 + int32(616)
	F_appendStringInfo(m, v376, int32(_a_F_ri_set_9), v20)
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	F_generate_operator_clause(m, v376, v364, v311, v373, v20+int32(208), v315)
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20-int32(-64)+v371))) = v311
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v25)+168))
	if v361 < v389 {
		v295 = v361
		v299 = int32(_a_F_ri_set_10)
		goto L67
	} else {
		goto L83
	}
L83:
	;
	goto L68
L84:
	;
	v432 = v413
	goto L17
L85:
	;
	v439 = F_SPI_finish(m)
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	if v439 == int32(2) {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	F_relation_close(m, v29, int32(3))
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L1
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L1
	} else {
		goto L95
	}
L90:
	;
	if l1 == int32(0) {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	F_ri_restrict(m, l0, int32(1))
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L1
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	m.G0 = v20 + int32(640)
	return
L94:
	;
	goto L93
L95:
	;
	F_errmsg_internal(m, int32(_a_F_ri_set_11), int32(0))
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	F_errfinish(m, int32(_a_F_ri_set_12), int32(1439), int32(_a_F_ri_set_13))
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
