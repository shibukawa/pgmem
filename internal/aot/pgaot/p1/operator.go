package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GetOperatorFromCompareType(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
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
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v150 int32
	_ = v150
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	v9 = m.G0
	v11 = v9 - int32(80)
	m.G0 = v11
	v17 = F_get_opclass_opfamily_and_input_type(m, l0, v11+int32(76), v11+int32(72))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return
	} else {
		if v17 != 0 {
			v19 = m.G0
			v21 = v19 - int32(16)
			m.G0 = v21
			v25 = F_SearchSysCache1(m, int32(14), base.I64_extend_i32_u(l0))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return
			} else {
				if v25 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v21))) = l0
						F_errmsg_internal(m, int32(_a_F_GetOperatorFromCompareType_0), v21)
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_GetOperatorFromCompareType_1), int32(1518), int32(_a_F_GetOperatorFromCompareType_2))
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v42 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
					v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+22)))
					v45 = *(*int32)(unsafe.Add(mBase, uint32(v42+v43)+4))
					F_ReleaseCatCache(m, v25)
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return
					} else {
						m.G0 = v21 + int32(16)
						v51 = *(*int32)(unsafe.Add(mBase, uint32(v11)+76))
						v53 = F_IndexAmTranslateCompareType(m, l2, v45, v51, int32(1))
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return
						} else {
							*(*uint16)(unsafe.Add(mBase, uint32(l4))) = uint16(v53)
							if v53 == int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v61 = m.ExcPending
								if v61 != 0 {
									return
								} else {
									F_errcode(m, int32(67137668))
									mBase = m.M
									v64 = m.ExcPending
									if v64 != 0 {
										return
									} else {
										switch l2 - int32(3) {
										case 0:
											v70 = int32(_a_F_GetOperatorFromCompareType_3)
											v71 = *(*int32)(unsafe.Add(mBase, uint32(v11)+72))
											v72 = F_format_type_be(m, v71)
											mBase = m.M
											v73 = m.ExcPending
											if v73 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v72
												F_errmsg(m, v70, v11+int32(16))
												mBase = m.M
												v78 = m.ExcPending
												if v78 != 0 {
													return
												} else {
													v80 = *(*int32)(unsafe.Add(mBase, uint32(v11)+76))
													v81 = F_get_opfamily_name(m, v80)
													mBase = m.M
													v82 = m.ExcPending
													if v82 != 0 {
														return
													} else {
														v83 = F_get_am_name(m, v45)
														mBase = m.M
														v84 = m.ExcPending
														if v84 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v83
															*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v81
															*(*int32)(unsafe.Add(mBase, uint32(v11))) = l2
															v89 = F_errdetail(m, int32(_a_F_GetOperatorFromCompareType_4), v11)
															mBase = m.M
															v90 = m.ExcPending
															if v90 != 0 {
																return
															} else {
																F_errfinish(m, int32(_a_F_GetOperatorFromCompareType_5), int32(2518), int32(_a_F_GetOperatorFromCompareType_6))
																mBase = m.M
																v95 = m.ExcPending
																if v95 != 0 {
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
										default:
											v80 = *(*int32)(unsafe.Add(mBase, uint32(v11)+76))
											v81 = F_get_opfamily_name(m, v80)
											mBase = m.M
											v82 = m.ExcPending
											if v82 != 0 {
												return
											} else {
												v83 = F_get_am_name(m, v45)
												mBase = m.M
												v84 = m.ExcPending
												if v84 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v83
													*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v81
													*(*int32)(unsafe.Add(mBase, uint32(v11))) = l2
													v89 = F_errdetail(m, int32(_a_F_GetOperatorFromCompareType_4), v11)
													mBase = m.M
													v90 = m.ExcPending
													if v90 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_GetOperatorFromCompareType_5), int32(2518), int32(_a_F_GetOperatorFromCompareType_6))
														mBase = m.M
														v95 = m.ExcPending
														if v95 != 0 {
															return
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											}
										case 4:
											v70 = int32(_a_F_GetOperatorFromCompareType_7)
											v71 = *(*int32)(unsafe.Add(mBase, uint32(v11)+72))
											v72 = F_format_type_be(m, v71)
											mBase = m.M
											v73 = m.ExcPending
											if v73 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v72
												F_errmsg(m, v70, v11+int32(16))
												mBase = m.M
												v78 = m.ExcPending
												if v78 != 0 {
													return
												} else {
													v80 = *(*int32)(unsafe.Add(mBase, uint32(v11)+76))
													v81 = F_get_opfamily_name(m, v80)
													mBase = m.M
													v82 = m.ExcPending
													if v82 != 0 {
														return
													} else {
														v83 = F_get_am_name(m, v45)
														mBase = m.M
														v84 = m.ExcPending
														if v84 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v83
															*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v81
															*(*int32)(unsafe.Add(mBase, uint32(v11))) = l2
															v89 = F_errdetail(m, int32(_a_F_GetOperatorFromCompareType_4), v11)
															mBase = m.M
															v90 = m.ExcPending
															if v90 != 0 {
																return
															} else {
																F_errfinish(m, int32(_a_F_GetOperatorFromCompareType_5), int32(2518), int32(_a_F_GetOperatorFromCompareType_6))
																mBase = m.M
																v95 = m.ExcPending
																if v95 != 0 {
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
										case 5:
											v70 = int32(_a_F_GetOperatorFromCompareType_8)
											v71 = *(*int32)(unsafe.Add(mBase, uint32(v11)+72))
											v72 = F_format_type_be(m, v71)
											mBase = m.M
											v73 = m.ExcPending
											if v73 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v72
												F_errmsg(m, v70, v11+int32(16))
												mBase = m.M
												v78 = m.ExcPending
												if v78 != 0 {
													return
												} else {
													v80 = *(*int32)(unsafe.Add(mBase, uint32(v11)+76))
													v81 = F_get_opfamily_name(m, v80)
													mBase = m.M
													v82 = m.ExcPending
													if v82 != 0 {
														return
													} else {
														v83 = F_get_am_name(m, v45)
														mBase = m.M
														v84 = m.ExcPending
														if v84 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v83
															*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v81
															*(*int32)(unsafe.Add(mBase, uint32(v11))) = l2
															v89 = F_errdetail(m, int32(_a_F_GetOperatorFromCompareType_4), v11)
															mBase = m.M
															v90 = m.ExcPending
															if v90 != 0 {
																return
															} else {
																F_errfinish(m, int32(_a_F_GetOperatorFromCompareType_5), int32(2518), int32(_a_F_GetOperatorFromCompareType_6))
																mBase = m.M
																v95 = m.ExcPending
																if v95 != 0 {
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
									}
								}
							} else {
								v96 = *(*int32)(unsafe.Add(mBase, uint32(v11)+76))
								v97 = *(*int32)(unsafe.Add(mBase, uint32(v11)+72))
								if l1 != 0 {
									v98 = l1
								} else {
									v98 = v97
								}
								v100 = F_get_opfamily_member(m, v96, v97, v98, base.I32_extend16_s(v53))
								mBase = m.M
								v101 = m.ExcPending
								if v101 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l3))) = v100
									if v100 == int32(0) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v108 = m.ExcPending
										if v108 != 0 {
											return
										} else {
											F_errcode(m, int32(67137668))
											mBase = m.M
											v111 = m.ExcPending
											if v111 != 0 {
												return
											} else {
												switch l2 - int32(3) {
												case 0:
													v117 = int32(_a_F_GetOperatorFromCompareType_3)
													v118 = *(*int32)(unsafe.Add(mBase, uint32(v11)+72))
													v119 = F_format_type_be(m, v118)
													mBase = m.M
													v120 = m.ExcPending
													if v120 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v119
														F_errmsg(m, v117, v11+int32(48))
														mBase = m.M
														v125 = m.ExcPending
														if v125 != 0 {
															return
														} else {
															v127 = *(*int32)(unsafe.Add(mBase, uint32(v11)+76))
															v128 = F_get_opfamily_name(m, v127)
															mBase = m.M
															v129 = m.ExcPending
															if v129 != 0 {
																return
															} else {
																v130 = F_get_am_name(m, v45)
																mBase = m.M
																v131 = m.ExcPending
																if v131 != 0 {
																	return
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v130
																	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v128
																	v137 = F_errdetail(m, int32(_a_F_GetOperatorFromCompareType_9), v11+int32(32))
																	mBase = m.M
																	v138 = m.ExcPending
																	if v138 != 0 {
																		return
																	} else {
																		F_errfinish(m, int32(_a_F_GetOperatorFromCompareType_5), int32(2536), int32(_a_F_GetOperatorFromCompareType_6))
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
														}
													}
												default:
													v127 = *(*int32)(unsafe.Add(mBase, uint32(v11)+76))
													v128 = F_get_opfamily_name(m, v127)
													mBase = m.M
													v129 = m.ExcPending
													if v129 != 0 {
														return
													} else {
														v130 = F_get_am_name(m, v45)
														mBase = m.M
														v131 = m.ExcPending
														if v131 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v130
															*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v128
															v137 = F_errdetail(m, int32(_a_F_GetOperatorFromCompareType_9), v11+int32(32))
															mBase = m.M
															v138 = m.ExcPending
															if v138 != 0 {
																return
															} else {
																F_errfinish(m, int32(_a_F_GetOperatorFromCompareType_5), int32(2536), int32(_a_F_GetOperatorFromCompareType_6))
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
												case 4:
													v117 = int32(_a_F_GetOperatorFromCompareType_7)
													v118 = *(*int32)(unsafe.Add(mBase, uint32(v11)+72))
													v119 = F_format_type_be(m, v118)
													mBase = m.M
													v120 = m.ExcPending
													if v120 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v119
														F_errmsg(m, v117, v11+int32(48))
														mBase = m.M
														v125 = m.ExcPending
														if v125 != 0 {
															return
														} else {
															v127 = *(*int32)(unsafe.Add(mBase, uint32(v11)+76))
															v128 = F_get_opfamily_name(m, v127)
															mBase = m.M
															v129 = m.ExcPending
															if v129 != 0 {
																return
															} else {
																v130 = F_get_am_name(m, v45)
																mBase = m.M
																v131 = m.ExcPending
																if v131 != 0 {
																	return
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v130
																	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v128
																	v137 = F_errdetail(m, int32(_a_F_GetOperatorFromCompareType_9), v11+int32(32))
																	mBase = m.M
																	v138 = m.ExcPending
																	if v138 != 0 {
																		return
																	} else {
																		F_errfinish(m, int32(_a_F_GetOperatorFromCompareType_5), int32(2536), int32(_a_F_GetOperatorFromCompareType_6))
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
														}
													}
												case 5:
													v117 = int32(_a_F_GetOperatorFromCompareType_8)
													v118 = *(*int32)(unsafe.Add(mBase, uint32(v11)+72))
													v119 = F_format_type_be(m, v118)
													mBase = m.M
													v120 = m.ExcPending
													if v120 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v119
														F_errmsg(m, v117, v11+int32(48))
														mBase = m.M
														v125 = m.ExcPending
														if v125 != 0 {
															return
														} else {
															v127 = *(*int32)(unsafe.Add(mBase, uint32(v11)+76))
															v128 = F_get_opfamily_name(m, v127)
															mBase = m.M
															v129 = m.ExcPending
															if v129 != 0 {
																return
															} else {
																v130 = F_get_am_name(m, v45)
																mBase = m.M
																v131 = m.ExcPending
																if v131 != 0 {
																	return
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v130
																	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v128
																	v137 = F_errdetail(m, int32(_a_F_GetOperatorFromCompareType_9), v11+int32(32))
																	mBase = m.M
																	v138 = m.ExcPending
																	if v138 != 0 {
																		return
																	} else {
																		F_errfinish(m, int32(_a_F_GetOperatorFromCompareType_5), int32(2536), int32(_a_F_GetOperatorFromCompareType_6))
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
														}
													}
												}
											}
										}
									} else {
										m.G0 = v11 + int32(80)
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
			v150 = m.ExcPending
			if v150 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v11)+64)) = l0
				F_errmsg_internal(m, int32(_a_F_GetOperatorFromCompareType_0), v11-int32(-64))
				mBase = m.M
				v156 = m.ExcPending
				if v156 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_GetOperatorFromCompareType_5), int32(2503), int32(_a_F_GetOperatorFromCompareType_6))
					mBase = m.M
					v161 = m.ExcPending
					if v161 != 0 {
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
func F_operator_same_subexprs_proof(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	if l2 != 0 {
		v6 = F_get_negator(m, l0)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			if v6 != l1 {
				v14 = int32(11)
				v15 = F_lookup_proof_cache(m, l0, l1, l2)
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return int32(0)
				} else {
					v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15+v14))))
					v21 = v18
					return v21 & int32(1)
				}
			} else {
				v21 = int32(1)
				return v21 & int32(1)
			}
		}
	} else {
		if l0 == l1 {
			v21 = int32(1)
			return v21 & int32(1)
		} else {
			v14 = int32(10)
			v15 = F_lookup_proof_cache(m, l0, l1, l2)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15+v14))))
				v21 = v18
				return v21 & int32(1)
			}
		}
	}
}
