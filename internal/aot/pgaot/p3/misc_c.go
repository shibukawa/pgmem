package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_CLOGPagePrecedes(m *base.Module, l0 int64, l1 int64) int32 {
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	v4 = int32(15)
	v9 = int32(4)
	v10 = base.I32_wrap_i64(l0)<<(uint(v4)%32) | v9
	v12 = base.I32_wrap_i64(l1) << (uint(v4) % 32)
	v15 = F_TransactionIdPrecedes(m, v10, v12|v9)
	if v15 != 0 {
		v17 = F_TransactionIdPrecedes(m, v10, v12+int32(_a_F_CLOGPagePrecedes_0))
		v19 = v17
	} else {
		v19 = int32(0)
	}
	return v19
}
func F_CallerFInfoFunctionCall2(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	v6 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+44)) = uint8(v6)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+40)) = l4
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+36)) = uint8(v6)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = l3
	v17 = int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v9)+30)) = uint16(v17)
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+28)) = uint8(v6)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = l2
	*(*int64)(unsafe.Add(mBase, uint32(v9)+16)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = l1
	v27 = m.T0[l0].(func(*base.Module, int32) int32)(m, v9+int32(12))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		return int32(0)
	} else {
		v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+28)))
		if v31 == int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
				F_errmsg_internal(m, int32(_a_F_CallerFInfoFunctionCall2_0), v9)
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_CallerFInfoFunctionCall2_1), int32(1101), int32(_a_F_CallerFInfoFunctionCall2_2))
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			m.G0 = v9 + int32(48)
			return v27
		}
	}
}
func F_CastCreate(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v87 int32
	_ = v87
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	v10 = int32(0)
	v13 = m.G0
	v15 = v13 + int32(-64)
	m.G0 = v15
	*(*uint16)(unsafe.Add(mBase, uint32(v15)+28)) = uint16(v10)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = v10
	v23 = F_table_open(m, int32(2605), int32(3))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		return
	} else {
		v26 = F_SearchSysCache2(m, int32(12), l1, l2)
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return
		} else {
			if v26 == int32(0) {
				v32 = F_GetNewOidWithIndex(m, v23, int32(2660), int32(1))
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = l7
					*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = l6
					*(*int32)(unsafe.Add(mBase, uint32(v15)+44)) = l3
					*(*int32)(unsafe.Add(mBase, uint32(v15)+40)) = l2
					*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = l1
					*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v32
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v23)+52))
					v45 = F_heap_form_tuple(m, v40, v13+int32(-32), v13+int32(-40))
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return
					} else {
						F_CatalogTupleInsert(m, v23, v45)
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return
						} else {
							v49 = F_new_object_addresses(m)
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return
							} else {
								v51 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v51
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v32
								*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(2605)
								*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v51
								*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = l1
								*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = int32(1247)
								v62 = v13 + int32(-52)
								F_add_exact_object_address(m, v62, v49)
								mBase = m.M
								v64 = m.ExcPending
								if v64 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = l2
									*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = int32(1247)
									F_add_exact_object_address(m, v62, v49)
									mBase = m.M
									v71 = m.ExcPending
									if v71 != 0 {
										return
									} else {
										if l3 != 0 {
											*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = l3
											*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = int32(1255)
											F_add_exact_object_address(m, v62, v49)
											mBase = m.M
											v78 = m.ExcPending
											if v78 != 0 {
												return
											} else {
												if l4 != 0 {
													*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = int32(0)
													*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = l4
													*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = int32(2605)
													F_add_exact_object_address(m, v13+int32(-52), v49)
													mBase = m.M
													v87 = m.ExcPending
													if v87 != 0 {
														return
													} else {
														if l5 != 0 {
															*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = int32(0)
															*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = l5
															*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = int32(2605)
															F_add_exact_object_address(m, v13+int32(-52), v49)
															mBase = m.M
															v96 = m.ExcPending
															if v96 != 0 {
																return
															} else {
																F_record_object_address_dependencies(m, l0, v49, l8)
																mBase = m.M
																v98 = m.ExcPending
																if v98 != 0 {
																	return
																} else {
																	F_free_object_addresses(m, v49)
																	mBase = m.M
																	v100 = m.ExcPending
																	if v100 != 0 {
																		return
																	} else {
																		F_recordDependencyOnCurrentExtension(m, l0, int32(0))
																		mBase = m.M
																		v103 = m.ExcPending
																		if v103 != 0 {
																			return
																		} else {
																			v105 = *(*int32)(unsafe.Add(mBase, _c_F_CastCreate[0]))
																			if v105 != 0 {
																				v107 = int32(0)
																				F_RunObjectPostCreateHook(m, int32(2605), v32, v107, v107)
																				mBase = m.M
																				v110 = m.ExcPending
																				if v110 != 0 {
																					return
																				} else {
																					F_pfree(m, v45)
																					mBase = m.M
																					v112 = m.ExcPending
																					if v112 != 0 {
																						return
																					} else {
																						F_relation_close(m, v23, int32(3))
																						mBase = m.M
																						v115 = m.ExcPending
																						if v115 != 0 {
																							return
																						} else {
																							m.G0 = v15 - int32(-64)
																							return
																						}
																					}
																				}
																			} else {
																				F_pfree(m, v45)
																				mBase = m.M
																				v112 = m.ExcPending
																				if v112 != 0 {
																					return
																				} else {
																					F_relation_close(m, v23, int32(3))
																					mBase = m.M
																					v115 = m.ExcPending
																					if v115 != 0 {
																						return
																					} else {
																						m.G0 = v15 - int32(-64)
																						return
																					}
																				}
																			}
																		}
																	}
																}
															}
														} else {
															F_record_object_address_dependencies(m, l0, v49, l8)
															mBase = m.M
															v98 = m.ExcPending
															if v98 != 0 {
																return
															} else {
																F_free_object_addresses(m, v49)
																mBase = m.M
																v100 = m.ExcPending
																if v100 != 0 {
																	return
																} else {
																	F_recordDependencyOnCurrentExtension(m, l0, int32(0))
																	mBase = m.M
																	v103 = m.ExcPending
																	if v103 != 0 {
																		return
																	} else {
																		v105 = *(*int32)(unsafe.Add(mBase, _c_F_CastCreate[0]))
																		if v105 != 0 {
																			v107 = int32(0)
																			F_RunObjectPostCreateHook(m, int32(2605), v32, v107, v107)
																			mBase = m.M
																			v110 = m.ExcPending
																			if v110 != 0 {
																				return
																			} else {
																				F_pfree(m, v45)
																				mBase = m.M
																				v112 = m.ExcPending
																				if v112 != 0 {
																					return
																				} else {
																					F_relation_close(m, v23, int32(3))
																					mBase = m.M
																					v115 = m.ExcPending
																					if v115 != 0 {
																						return
																					} else {
																						m.G0 = v15 - int32(-64)
																						return
																					}
																				}
																			}
																		} else {
																			F_pfree(m, v45)
																			mBase = m.M
																			v112 = m.ExcPending
																			if v112 != 0 {
																				return
																			} else {
																				F_relation_close(m, v23, int32(3))
																				mBase = m.M
																				v115 = m.ExcPending
																				if v115 != 0 {
																					return
																				} else {
																					m.G0 = v15 - int32(-64)
																					return
																				}
																			}
																		}
																	}
																}
															}
														}
													}
												} else {
													if l5 != 0 {
														*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = int32(0)
														*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = l5
														*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = int32(2605)
														F_add_exact_object_address(m, v13+int32(-52), v49)
														mBase = m.M
														v96 = m.ExcPending
														if v96 != 0 {
															return
														} else {
															F_record_object_address_dependencies(m, l0, v49, l8)
															mBase = m.M
															v98 = m.ExcPending
															if v98 != 0 {
																return
															} else {
																F_free_object_addresses(m, v49)
																mBase = m.M
																v100 = m.ExcPending
																if v100 != 0 {
																	return
																} else {
																	F_recordDependencyOnCurrentExtension(m, l0, int32(0))
																	mBase = m.M
																	v103 = m.ExcPending
																	if v103 != 0 {
																		return
																	} else {
																		v105 = *(*int32)(unsafe.Add(mBase, _c_F_CastCreate[0]))
																		if v105 != 0 {
																			v107 = int32(0)
																			F_RunObjectPostCreateHook(m, int32(2605), v32, v107, v107)
																			mBase = m.M
																			v110 = m.ExcPending
																			if v110 != 0 {
																				return
																			} else {
																				F_pfree(m, v45)
																				mBase = m.M
																				v112 = m.ExcPending
																				if v112 != 0 {
																					return
																				} else {
																					F_relation_close(m, v23, int32(3))
																					mBase = m.M
																					v115 = m.ExcPending
																					if v115 != 0 {
																						return
																					} else {
																						m.G0 = v15 - int32(-64)
																						return
																					}
																				}
																			}
																		} else {
																			F_pfree(m, v45)
																			mBase = m.M
																			v112 = m.ExcPending
																			if v112 != 0 {
																				return
																			} else {
																				F_relation_close(m, v23, int32(3))
																				mBase = m.M
																				v115 = m.ExcPending
																				if v115 != 0 {
																					return
																				} else {
																					m.G0 = v15 - int32(-64)
																					return
																				}
																			}
																		}
																	}
																}
															}
														}
													} else {
														F_record_object_address_dependencies(m, l0, v49, l8)
														mBase = m.M
														v98 = m.ExcPending
														if v98 != 0 {
															return
														} else {
															F_free_object_addresses(m, v49)
															mBase = m.M
															v100 = m.ExcPending
															if v100 != 0 {
																return
															} else {
																F_recordDependencyOnCurrentExtension(m, l0, int32(0))
																mBase = m.M
																v103 = m.ExcPending
																if v103 != 0 {
																	return
																} else {
																	v105 = *(*int32)(unsafe.Add(mBase, _c_F_CastCreate[0]))
																	if v105 != 0 {
																		v107 = int32(0)
																		F_RunObjectPostCreateHook(m, int32(2605), v32, v107, v107)
																		mBase = m.M
																		v110 = m.ExcPending
																		if v110 != 0 {
																			return
																		} else {
																			F_pfree(m, v45)
																			mBase = m.M
																			v112 = m.ExcPending
																			if v112 != 0 {
																				return
																			} else {
																				F_relation_close(m, v23, int32(3))
																				mBase = m.M
																				v115 = m.ExcPending
																				if v115 != 0 {
																					return
																				} else {
																					m.G0 = v15 - int32(-64)
																					return
																				}
																			}
																		}
																	} else {
																		F_pfree(m, v45)
																		mBase = m.M
																		v112 = m.ExcPending
																		if v112 != 0 {
																			return
																		} else {
																			F_relation_close(m, v23, int32(3))
																			mBase = m.M
																			v115 = m.ExcPending
																			if v115 != 0 {
																				return
																			} else {
																				m.G0 = v15 - int32(-64)
																				return
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
											if l4 != 0 {
												*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = int32(0)
												*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = l4
												*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = int32(2605)
												F_add_exact_object_address(m, v13+int32(-52), v49)
												mBase = m.M
												v87 = m.ExcPending
												if v87 != 0 {
													return
												} else {
													if l5 != 0 {
														*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = int32(0)
														*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = l5
														*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = int32(2605)
														F_add_exact_object_address(m, v13+int32(-52), v49)
														mBase = m.M
														v96 = m.ExcPending
														if v96 != 0 {
															return
														} else {
															F_record_object_address_dependencies(m, l0, v49, l8)
															mBase = m.M
															v98 = m.ExcPending
															if v98 != 0 {
																return
															} else {
																F_free_object_addresses(m, v49)
																mBase = m.M
																v100 = m.ExcPending
																if v100 != 0 {
																	return
																} else {
																	F_recordDependencyOnCurrentExtension(m, l0, int32(0))
																	mBase = m.M
																	v103 = m.ExcPending
																	if v103 != 0 {
																		return
																	} else {
																		v105 = *(*int32)(unsafe.Add(mBase, _c_F_CastCreate[0]))
																		if v105 != 0 {
																			v107 = int32(0)
																			F_RunObjectPostCreateHook(m, int32(2605), v32, v107, v107)
																			mBase = m.M
																			v110 = m.ExcPending
																			if v110 != 0 {
																				return
																			} else {
																				F_pfree(m, v45)
																				mBase = m.M
																				v112 = m.ExcPending
																				if v112 != 0 {
																					return
																				} else {
																					F_relation_close(m, v23, int32(3))
																					mBase = m.M
																					v115 = m.ExcPending
																					if v115 != 0 {
																						return
																					} else {
																						m.G0 = v15 - int32(-64)
																						return
																					}
																				}
																			}
																		} else {
																			F_pfree(m, v45)
																			mBase = m.M
																			v112 = m.ExcPending
																			if v112 != 0 {
																				return
																			} else {
																				F_relation_close(m, v23, int32(3))
																				mBase = m.M
																				v115 = m.ExcPending
																				if v115 != 0 {
																					return
																				} else {
																					m.G0 = v15 - int32(-64)
																					return
																				}
																			}
																		}
																	}
																}
															}
														}
													} else {
														F_record_object_address_dependencies(m, l0, v49, l8)
														mBase = m.M
														v98 = m.ExcPending
														if v98 != 0 {
															return
														} else {
															F_free_object_addresses(m, v49)
															mBase = m.M
															v100 = m.ExcPending
															if v100 != 0 {
																return
															} else {
																F_recordDependencyOnCurrentExtension(m, l0, int32(0))
																mBase = m.M
																v103 = m.ExcPending
																if v103 != 0 {
																	return
																} else {
																	v105 = *(*int32)(unsafe.Add(mBase, _c_F_CastCreate[0]))
																	if v105 != 0 {
																		v107 = int32(0)
																		F_RunObjectPostCreateHook(m, int32(2605), v32, v107, v107)
																		mBase = m.M
																		v110 = m.ExcPending
																		if v110 != 0 {
																			return
																		} else {
																			F_pfree(m, v45)
																			mBase = m.M
																			v112 = m.ExcPending
																			if v112 != 0 {
																				return
																			} else {
																				F_relation_close(m, v23, int32(3))
																				mBase = m.M
																				v115 = m.ExcPending
																				if v115 != 0 {
																					return
																				} else {
																					m.G0 = v15 - int32(-64)
																					return
																				}
																			}
																		}
																	} else {
																		F_pfree(m, v45)
																		mBase = m.M
																		v112 = m.ExcPending
																		if v112 != 0 {
																			return
																		} else {
																			F_relation_close(m, v23, int32(3))
																			mBase = m.M
																			v115 = m.ExcPending
																			if v115 != 0 {
																				return
																			} else {
																				m.G0 = v15 - int32(-64)
																				return
																			}
																		}
																	}
																}
															}
														}
													}
												}
											} else {
												if l5 != 0 {
													*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = int32(0)
													*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = l5
													*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = int32(2605)
													F_add_exact_object_address(m, v13+int32(-52), v49)
													mBase = m.M
													v96 = m.ExcPending
													if v96 != 0 {
														return
													} else {
														F_record_object_address_dependencies(m, l0, v49, l8)
														mBase = m.M
														v98 = m.ExcPending
														if v98 != 0 {
															return
														} else {
															F_free_object_addresses(m, v49)
															mBase = m.M
															v100 = m.ExcPending
															if v100 != 0 {
																return
															} else {
																F_recordDependencyOnCurrentExtension(m, l0, int32(0))
																mBase = m.M
																v103 = m.ExcPending
																if v103 != 0 {
																	return
																} else {
																	v105 = *(*int32)(unsafe.Add(mBase, _c_F_CastCreate[0]))
																	if v105 != 0 {
																		v107 = int32(0)
																		F_RunObjectPostCreateHook(m, int32(2605), v32, v107, v107)
																		mBase = m.M
																		v110 = m.ExcPending
																		if v110 != 0 {
																			return
																		} else {
																			F_pfree(m, v45)
																			mBase = m.M
																			v112 = m.ExcPending
																			if v112 != 0 {
																				return
																			} else {
																				F_relation_close(m, v23, int32(3))
																				mBase = m.M
																				v115 = m.ExcPending
																				if v115 != 0 {
																					return
																				} else {
																					m.G0 = v15 - int32(-64)
																					return
																				}
																			}
																		}
																	} else {
																		F_pfree(m, v45)
																		mBase = m.M
																		v112 = m.ExcPending
																		if v112 != 0 {
																			return
																		} else {
																			F_relation_close(m, v23, int32(3))
																			mBase = m.M
																			v115 = m.ExcPending
																			if v115 != 0 {
																				return
																			} else {
																				m.G0 = v15 - int32(-64)
																				return
																			}
																		}
																	}
																}
															}
														}
													}
												} else {
													F_record_object_address_dependencies(m, l0, v49, l8)
													mBase = m.M
													v98 = m.ExcPending
													if v98 != 0 {
														return
													} else {
														F_free_object_addresses(m, v49)
														mBase = m.M
														v100 = m.ExcPending
														if v100 != 0 {
															return
														} else {
															F_recordDependencyOnCurrentExtension(m, l0, int32(0))
															mBase = m.M
															v103 = m.ExcPending
															if v103 != 0 {
																return
															} else {
																v105 = *(*int32)(unsafe.Add(mBase, _c_F_CastCreate[0]))
																if v105 != 0 {
																	v107 = int32(0)
																	F_RunObjectPostCreateHook(m, int32(2605), v32, v107, v107)
																	mBase = m.M
																	v110 = m.ExcPending
																	if v110 != 0 {
																		return
																	} else {
																		F_pfree(m, v45)
																		mBase = m.M
																		v112 = m.ExcPending
																		if v112 != 0 {
																			return
																		} else {
																			F_relation_close(m, v23, int32(3))
																			mBase = m.M
																			v115 = m.ExcPending
																			if v115 != 0 {
																				return
																			} else {
																				m.G0 = v15 - int32(-64)
																				return
																			}
																		}
																	}
																} else {
																	F_pfree(m, v45)
																	mBase = m.M
																	v112 = m.ExcPending
																	if v112 != 0 {
																		return
																	} else {
																		F_relation_close(m, v23, int32(3))
																		mBase = m.M
																		v115 = m.ExcPending
																		if v115 != 0 {
																			return
																		} else {
																			m.G0 = v15 - int32(-64)
																			return
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
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v122 = m.ExcPending
				if v122 != 0 {
					return
				} else {
					F_errcode(m, int32(_a_F_CastCreate_0))
					mBase = m.M
					v125 = m.ExcPending
					if v125 != 0 {
						return
					} else {
						v126 = F_format_type_be(m, l1)
						mBase = m.M
						v127 = m.ExcPending
						if v127 != 0 {
							return
						} else {
							v128 = F_format_type_be(m, l2)
							mBase = m.M
							v129 = m.ExcPending
							if v129 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v128
								*(*int32)(unsafe.Add(mBase, uint32(v15))) = v126
								F_errmsg(m, int32(_a_F_CastCreate_1), v15)
								mBase = m.M
								v134 = m.ExcPending
								if v134 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_CastCreate_2), int32(77), int32(_a_F_CastCreate_3))
									mBase = m.M
									v139 = m.ExcPending
									if v139 != 0 {
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
}
func F_CheckBuiltinCryptoMode(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	v2 = *(*int32)(unsafe.Add(mBase, _c_F_CheckBuiltinCryptoMode[0]))
	if v2 == int32(1) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			F_errmsg(m, int32(_a_F_CheckBuiltinCryptoMode_0), int32(0))
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_CheckBuiltinCryptoMode_1), int32(735), int32(_a_F_CheckBuiltinCryptoMode_2))
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		return
	}
}
func F_CheckElement_2(m *base.Module, l0 float32) {
	var v8 int32
	_ = v8
	Fn13824(m, l0, int32(147), int32(_a_F_CheckElement_2_0), int32(_a_F_CheckElement_2_1), int32(142), int32(_a_F_CheckElement_2_2))
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		return
	}
}
func F_CheckSASLAuth(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
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
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v185 int32
	_ = v185
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	v5 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(80)
	m.G0 = v13
	*(*int32)(unsafe.Add(mBase, uint32(v13)+44)) = v5
	*(*int32)(unsafe.Add(mBase, uint32(v13)+40)) = v5
	v20 = v13 - int32(-64)
	F_initStringInfo(m, v20)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	m.T0[v25].(func(*base.Module, int32, int32))(m, l1, v20)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	F_appendStringInfoChar(m, v20, int32(0))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v13)+64))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v13)+68))
	F_sendAuthRequest(m, int32(10), v32, v33)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v13)+64))
	F_pfree(m, v36)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v47 = int32(1)
	v50 = v5
	goto L10
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L1
	} else {
		goto L66
	}
L8:
	;
	m.G0 = v13 + int32(80)
	return v185
L9:
	;
	v185 = int32(-1)
	goto L8
L10:
	;
	F_pq_startmsgread(m)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L12
	}
L11:
	;
	if v135 == int32(1) {
		v185 = int32(0)
		goto L8
	} else {
		goto L65
	}
L12:
	;
	v53 = F_pq_getbyte(m)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	if v53 != int32(112) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	if v53 == int32(-1) {
		v185 = int32(-2)
		goto L8
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v76 = v13 + int32(48)
	F_initStringInfo(m, v76)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L22
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v53
	F_errmsg(m, int32(_a_F_CheckSASLAuth_0), v13)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	F_errfinish(m, int32(_a_F_CheckSASLAuth_1), int32(90), int32(_a_F_CheckSASLAuth_2))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L22:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v80 = F_pq_getmessage(m, v76, v79)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	if v80 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v13)+48))
	F_pfree(m, v82)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v87 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L28
	}
L27:
	;
	goto L9
L28:
	;
	if v87 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v13)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v89
	F_errmsg_internal(m, int32(_a_F_CheckSASLAuth_3), v13+int32(32))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v102 = v13 + int32(48)
	if v47 != 0 {
		goto L36
	} else {
		goto L37
	}
L32:
	;
	F_errfinish(m, int32(_a_F_CheckSASLAuth_1), int32(105), int32(_a_F_CheckSASLAuth_2))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	F_pq_getmsgend(m, v13+int32(48))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L46
	}
L35:
	;
	v120 = F_pq_getmsgbytes(m, v102, v119)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L45
	}
L36:
	;
	v104 = F_pq_getmsgrawstring(m, v102)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v13)+52))
	v118 = v50
	v119 = v115
	goto L35
L39:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v107 = m.T0[v106].(func(*base.Module, int32, int32, int32) int32)(m, l1, v104, l2)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	v110 = F_pq_getmsgint(m, v102, int32(4))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	if v110 != int32(-1) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v118 = v107
	v119 = v110
	goto L35
L43:
	;
	goto L44
L44:
	;
	v122 = int32(-1)
	v124 = v107
	v125 = int32(0)
	goto L34
L45:
	;
	v122 = v119
	v124 = v118
	v125 = v120
	goto L34
L46:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v135 = m.T0[v134].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v124, v125, v122, v13+int32(44), v13+int32(40), l3)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v13)+48))
	F_pfree(m, v137)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v13)+44))
	if v140 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	if v135 == int32(2) {
		goto L7
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	v171 = int32(0)
	if v135 == v171 {
		v47 = v171
		v50 = v124
		goto L10
	} else {
		goto L64
	}
L52:
	;
	v145 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	if v145 != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v13)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v147
	F_errmsg_internal(m, int32(_a_F_CheckSASLAuth_4), v13+int32(16))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	if v135 == int32(1) {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	F_errfinish(m, int32(_a_F_CheckSASLAuth_1), int32(176), int32(_a_F_CheckSASLAuth_2))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	goto L56
L59:
	;
	v163 = int32(12)
	goto L61
L60:
	;
	v163 = int32(11)
	goto L61
L61:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v13)+44))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v13)+40))
	F_sendAuthRequest(m, v163, v164, v165)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v13)+44))
	F_pfree(m, v168)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	goto L51
L64:
	;
	goto L11
L65:
	;
	goto L9
L66:
	;
	F_errmsg_internal(m, int32(_a_F_CheckSASLAuth_5), int32(0))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(_a_F_CheckSASLAuth_1), int32(171), int32(_a_F_CheckSASLAuth_2))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ClosePipeToProgram(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v9 = F_ClosePipeStream(m, v8)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		if v9 != 0 {
			if v9 == int32(-1) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return
				} else {
					F_errcode_for_file_access(m)
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return
					} else {
						F_errmsg(m, int32(_a_F_ClosePipeToProgram_0), int32(0))
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_ClosePipeToProgram_1), int32(572), int32(_a_F_ClosePipeToProgram_2))
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
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
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return
				} else {
					F_errcode(m, int32(515))
					mBase = m.M
					v19 = m.ExcPending
					if v19 != 0 {
						return
					} else {
						v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
						*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v20
						F_errmsg(m, int32(_a_F_ClosePipeToProgram_3), v6+int32(16))
						mBase = m.M
						v26 = m.ExcPending
						if v26 != 0 {
							return
						} else {
							v27 = F_wait_result_to_str(m, v9)
							mBase = m.M
							v28 = m.ExcPending
							if v28 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v6))) = v27
								F_errdetail_internal(m, int32(_a_F_ClosePipeToProgram_4), v6)
								mBase = m.M
								v32 = m.ExcPending
								if v32 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_ClosePipeToProgram_1), int32(579), int32(_a_F_ClosePipeToProgram_2))
									mBase = m.M
									v37 = m.ExcPending
									if v37 != 0 {
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
		} else {
			m.G0 = v6 + int32(32)
			return
		}
	}
}
func F_CombineRangeTables(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v9 == int32(0) {
		v56 = int32(0)
	} else {
		v13 = int32(0)
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
		if base.B2i32(l2 == v13)|base.B2i32(v15 <= v13) != 0 {
			v56 = v9
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
			if int32(0) < v19 {
				v27 = int32(0)
				for {
					v31 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v31+v27<<(uint(int32(2))%32))))
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+28))
					if v36 != 0 {
						*(*int32)(unsafe.Add(mBase, uint32(v35)+28)) = v15 + v36
					} else {
					}
					v40 = v27 + int32(1)
					v41 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
					if v40 < v41 {
						v27 = v40
						continue
					} else {
						break
					}
					break
				}
			} else {
			}
			v51 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v56 = v51
		}
	}
	v60 = F_list_concat(m, v56, l3)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v60
		v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v64 = F_list_concat(m, v63, l2)
		mBase = m.M
		v65 = m.ExcPending
		if v65 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = v64
			return
		}
	}
}
func F_CompareCandidateDistancesOffset(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 float32
	_ = v7
	var v8 int32
	_ = v8
	var v9 float32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = *(*float32)(unsafe.Add(mBase, uint32(v6)+4))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v9 = *(*float32)(unsafe.Add(mBase, uint32(v8)+4))
	if base.F32_lt(v7, v9) != 0 {
		v27 = int32(1)
	} else {
		if base.F32_gt(v7, v9) != 0 {
			v27 = int32(-1)
		} else {
			v13 = int32(1)
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
			v16 = v14 - v13
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
			v19 = v17 - v13
			if base.Ui32(v16) < base.Ui32(v19) {
				v27 = v13
			} else {
				if base.Ui32(v19) < base.Ui32(v16) {
					v24 = int32(-1)
				} else {
					v24 = int32(0)
				}
				v27 = v24
			}
		}
	}
	return v27
}
func F_CompareNearestCandidates(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 float64
	_ = v9
	var v10 float64
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	v9 = *(*float64)(unsafe.Add(mBase, uint32(l0)+32))
	v10 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	if base.F64_gt(v9, v10) != 0 {
		v12 = int32(-1)
	} else {
		v12 = int32(0)
	}
	if base.F64_lt(v9, v10) != 0 {
		v14 = int32(1)
	} else {
		v14 = v12
	}
	return v14
}
func F_ConditionVariableTimedSleep(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v29 int64
	_ = v29
	var v32 int64
	_ = v32
	var v35 int32
	_ = v35
	var v36 int64
	_ = v36
	var v37 int32
	_ = v37
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v111 int64
	_ = v111
	var v112 int64
	_ = v112
	var v121 int32
	_ = v121
	var v130 int32
	_ = v130
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableTimedSleep[0]))
	if l0 != v17 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v14 + int32(16)
	return v130
L2:
	;
	F_ConditionVariablePrepareToSleep(m, l0)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v24 = base.B2i32(l1 < int32(0))
	if l1 < int32(0) {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	return int32(0)
L6:
	;
	v130 = int32(0)
	goto L1
L7:
	;
	v46 = v37
	goto L11
L8:
	;
	v35 = int32(33)
	v36 = int64(0)
	v37 = int32(-1)
	goto L7
L9:
	;
	goto L10
L10:
	;
	F___clock_gettime(m, int32(1), v14)
	mBase = m.M
	v29 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
	v32 = int64(*(*int32)(unsafe.Add(mBase, uint32(v14)+8)))
	v35 = int32(41)
	v36 = v29*int64(-1000000000) - v32
	v37 = l1
	goto L7
L11:
	;
	v50 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableTimedSleep[1]))
	v51 = F_WaitLatch(m, v50, v35, v46, l2)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L5
	} else {
		goto L13
	}
L12:
	;
	v130 = int32(0)
	goto L1
L13:
	;
	v54 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableTimedSleep[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v54))) = int32(0)
	goto L14
L14:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1)
	if v57 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	F_s_lock(m, l0, int32(_a_F_ConditionVariableTimedSleep_0), int32(183), int32(_a_F_ConditionVariableTimedSleep_1))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L5
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v65 = int32(0)
	v67 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableTimedSleep[2]))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	v70 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableTimedSleep[3]))
	v73 = v68 + v70*int32(640)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+88))
	if v74 != 0 {
		v94 = v65
		goto L19
	} else {
		goto L20
	}
L18:
	;
	goto L17
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
	v98 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableTimedSleep[4]))
	if v98 != 0 {
		goto L26
	} else {
		goto L27
	}
L20:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v73)+84))
	if v75 != 0 {
		v94 = v65
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v76 == int32(-1) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v70
	v94 = int32(1)
	goto L19
L23:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v73)+84)) = int64(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v70
	goto L22
L24:
	;
	goto L25
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v73)+88)) = v76
	v84 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableTimedSleep[2]))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	*(*int32)(unsafe.Add(mBase, uint32(v85+v76*int32(640))+84)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v73)+84)) = int32(-1)
	goto L22
L26:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L5
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v102 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableTimedSleep[0]))
	v104 = v94 | base.B2i32(l0 != v102)
	if v104|v24 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	goto L28
L30:
	;
	v108 = int32(1)
	F___clock_gettime(m, v108, v14)
	mBase = m.M
	v111 = int64(*(*int32)(unsafe.Add(mBase, uint32(v14)+8)))
	v112 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
	v121 = l1 - base.I32_trunc_sat_f64_s(base.F64_div(base.F64_convert_i64_s(v111+(v112*int64(1000000000)+v36)), float64(1e+06)))
	if int32(0) < v121 {
		v46 = v121
		goto L11
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	if v104 == int32(0) {
		goto L11
	} else {
		goto L34
	}
L33:
	;
	v130 = v108
	goto L1
L34:
	;
	goto L12
}
func F_CopySendEndOfRow(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
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
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int64
	_ = v77
	var v78 int64
	_ = v78
	var v79 int64
	_ = v79
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	switch v5 {
	case 0:
		v6 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v10 = F_fwrite(m, v6, v7, int32(1), v9)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return
		} else {
			if v10 == int32(1) {
				v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
				if int32(base.Ui32(v15)>>(uint(int32(5))%32))&int32(1) == int32(0) {
					v77 = *(*int64)(unsafe.Add(mBase, uint32(l0)+176))
					v78 = int64(*(*int32)(unsafe.Add(mBase, uint32(v4)+4)))
					v79 = v77 + v78
					*(*int64)(unsafe.Add(mBase, uint32(l0)+176)) = v79
					v84 = *(*int32)(unsafe.Add(mBase, _c_F_CopySendEndOfRow[0]))
					if v84 == int32(0) {
					} else {
						v88 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CopySendEndOfRow[1])))
						if v88&int32(1) == int32(0) {
						} else {
							v93 = int32(_a_F_CopySendEndOfRow_0)
							v95 = *(*int32)(unsafe.Add(mBase, _c_F_CopySendEndOfRow[2]))
							v96 = int32(1)
							*(*int32)(unsafe.Add(mBase, _c_F_CopySendEndOfRow[2])) = v95 + v96
							v99 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
							*(*int32)(unsafe.Add(mBase, uint32(v84))) = v99 + v96
							*(*int64)(unsafe.Add(mBase, uint32(v84+int32(0))+232)) = v79
							v107 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
							*(*int32)(unsafe.Add(mBase, uint32(v84))) = v107 + v96
							v113 = *(*int32)(unsafe.Add(mBase, _c_F_CopySendEndOfRow[2]))
							*(*int32)(unsafe.Add(mBase, _c_F_CopySendEndOfRow[2])) = v113 - v96
						}
					}
					v117 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
					v118 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v117))) = uint8(v118)
					*(*int32)(unsafe.Add(mBase, uint32(v4)+12)) = v118
					*(*int32)(unsafe.Add(mBase, uint32(v4)+4)) = v118
					return
				} else {
					v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)))
					if v22 == int32(1) {
						v26 = *(*int32)(unsafe.Add(mBase, _c_F_CopySendEndOfRow[3]))
						if v26 == int32(64) {
							F_ClosePipeToProgram(m, l0)
							mBase = m.M
							v30 = m.ExcPending
							if v30 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, _c_F_CopySendEndOfRow[3])) = int32(64)
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v37 = m.ExcPending
								if v37 != 0 {
									return
								} else {
									F_errcode_for_file_access(m)
									mBase = m.M
									v39 = m.ExcPending
									if v39 != 0 {
										return
									} else {
										F_errmsg(m, int32(_a_F_CopySendEndOfRow_1), int32(0))
										mBase = m.M
										v43 = m.ExcPending
										if v43 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_CopySendEndOfRow_2), int32(477), int32(_a_F_CopySendEndOfRow_3))
											mBase = m.M
											v48 = m.ExcPending
											if v48 != 0 {
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
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return
							} else {
								F_errcode_for_file_access(m)
								mBase = m.M
								v39 = m.ExcPending
								if v39 != 0 {
									return
								} else {
									F_errmsg(m, int32(_a_F_CopySendEndOfRow_1), int32(0))
									mBase = m.M
									v43 = m.ExcPending
									if v43 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_CopySendEndOfRow_2), int32(477), int32(_a_F_CopySendEndOfRow_3))
										mBase = m.M
										v48 = m.ExcPending
										if v48 != 0 {
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
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return
						} else {
							F_errcode_for_file_access(m)
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return
							} else {
								F_errmsg(m, int32(_a_F_CopySendEndOfRow_4), int32(0))
								mBase = m.M
								v58 = m.ExcPending
								if v58 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_CopySendEndOfRow_2), int32(482), int32(_a_F_CopySendEndOfRow_3))
									mBase = m.M
									v63 = m.ExcPending
									if v63 != 0 {
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
			} else {
				v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)))
				if v22 == int32(1) {
					v26 = *(*int32)(unsafe.Add(mBase, _c_F_CopySendEndOfRow[3]))
					if v26 == int32(64) {
						F_ClosePipeToProgram(m, l0)
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, _c_F_CopySendEndOfRow[3])) = int32(64)
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return
							} else {
								F_errcode_for_file_access(m)
								mBase = m.M
								v39 = m.ExcPending
								if v39 != 0 {
									return
								} else {
									F_errmsg(m, int32(_a_F_CopySendEndOfRow_1), int32(0))
									mBase = m.M
									v43 = m.ExcPending
									if v43 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_CopySendEndOfRow_2), int32(477), int32(_a_F_CopySendEndOfRow_3))
										mBase = m.M
										v48 = m.ExcPending
										if v48 != 0 {
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
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return
						} else {
							F_errcode_for_file_access(m)
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
								return
							} else {
								F_errmsg(m, int32(_a_F_CopySendEndOfRow_1), int32(0))
								mBase = m.M
								v43 = m.ExcPending
								if v43 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_CopySendEndOfRow_2), int32(477), int32(_a_F_CopySendEndOfRow_3))
									mBase = m.M
									v48 = m.ExcPending
									if v48 != 0 {
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
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return
					} else {
						F_errcode_for_file_access(m)
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return
						} else {
							F_errmsg(m, int32(_a_F_CopySendEndOfRow_4), int32(0))
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_CopySendEndOfRow_2), int32(482), int32(_a_F_CopySendEndOfRow_3))
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
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
	case 1:
		v65 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
		v66 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
		v68 = *(*int32)(unsafe.Add(mBase, _c_F_CopySendEndOfRow[4]))
		v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+16))
		v70 = m.T0[v69].(func(*base.Module, int32, int32, int32) int32)(m, int32(100), v65, v66)
		mBase = m.M
		v71 = m.ExcPending
		if v71 != 0 {
			return
		} else {
			v77 = *(*int64)(unsafe.Add(mBase, uint32(l0)+176))
			v78 = int64(*(*int32)(unsafe.Add(mBase, uint32(v4)+4)))
			v79 = v77 + v78
			*(*int64)(unsafe.Add(mBase, uint32(l0)+176)) = v79
			v84 = *(*int32)(unsafe.Add(mBase, _c_F_CopySendEndOfRow[0]))
			if v84 == int32(0) {
			} else {
				v88 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CopySendEndOfRow[1])))
				if v88&int32(1) == int32(0) {
				} else {
					v93 = int32(_a_F_CopySendEndOfRow_0)
					v95 = *(*int32)(unsafe.Add(mBase, _c_F_CopySendEndOfRow[2]))
					v96 = int32(1)
					*(*int32)(unsafe.Add(mBase, _c_F_CopySendEndOfRow[2])) = v95 + v96
					v99 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
					*(*int32)(unsafe.Add(mBase, uint32(v84))) = v99 + v96
					*(*int64)(unsafe.Add(mBase, uint32(v84+int32(0))+232)) = v79
					v107 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
					*(*int32)(unsafe.Add(mBase, uint32(v84))) = v107 + v96
					v113 = *(*int32)(unsafe.Add(mBase, _c_F_CopySendEndOfRow[2]))
					*(*int32)(unsafe.Add(mBase, _c_F_CopySendEndOfRow[2])) = v113 - v96
				}
			}
			v117 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
			v118 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v117))) = uint8(v118)
			*(*int32)(unsafe.Add(mBase, uint32(v4)+12)) = v118
			*(*int32)(unsafe.Add(mBase, uint32(v4)+4)) = v118
			return
		}
	case 2:
		v72 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
		v73 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
		v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
		m.T0[v74].(func(*base.Module, int32, int32))(m, v72, v73)
		mBase = m.M
		v76 = m.ExcPending
		if v76 != 0 {
			return
		} else {
			v77 = *(*int64)(unsafe.Add(mBase, uint32(l0)+176))
			v78 = int64(*(*int32)(unsafe.Add(mBase, uint32(v4)+4)))
			v79 = v77 + v78
			*(*int64)(unsafe.Add(mBase, uint32(l0)+176)) = v79
			v84 = *(*int32)(unsafe.Add(mBase, _c_F_CopySendEndOfRow[0]))
			if v84 == int32(0) {
			} else {
				v88 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CopySendEndOfRow[1])))
				if v88&int32(1) == int32(0) {
				} else {
					v93 = int32(_a_F_CopySendEndOfRow_0)
					v95 = *(*int32)(unsafe.Add(mBase, _c_F_CopySendEndOfRow[2]))
					v96 = int32(1)
					*(*int32)(unsafe.Add(mBase, _c_F_CopySendEndOfRow[2])) = v95 + v96
					v99 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
					*(*int32)(unsafe.Add(mBase, uint32(v84))) = v99 + v96
					*(*int64)(unsafe.Add(mBase, uint32(v84+int32(0))+232)) = v79
					v107 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
					*(*int32)(unsafe.Add(mBase, uint32(v84))) = v107 + v96
					v113 = *(*int32)(unsafe.Add(mBase, _c_F_CopySendEndOfRow[2]))
					*(*int32)(unsafe.Add(mBase, _c_F_CopySendEndOfRow[2])) = v113 - v96
				}
			}
			v117 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
			v118 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v117))) = uint8(v118)
			*(*int32)(unsafe.Add(mBase, uint32(v4)+12)) = v118
			*(*int32)(unsafe.Add(mBase, uint32(v4)+4)) = v118
			return
		}
	default:
		v77 = *(*int64)(unsafe.Add(mBase, uint32(l0)+176))
		v78 = int64(*(*int32)(unsafe.Add(mBase, uint32(v4)+4)))
		v79 = v77 + v78
		*(*int64)(unsafe.Add(mBase, uint32(l0)+176)) = v79
		v84 = *(*int32)(unsafe.Add(mBase, _c_F_CopySendEndOfRow[0]))
		if v84 == int32(0) {
		} else {
			v88 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CopySendEndOfRow[1])))
			if v88&int32(1) == int32(0) {
			} else {
				v93 = int32(_a_F_CopySendEndOfRow_0)
				v95 = *(*int32)(unsafe.Add(mBase, _c_F_CopySendEndOfRow[2]))
				v96 = int32(1)
				*(*int32)(unsafe.Add(mBase, _c_F_CopySendEndOfRow[2])) = v95 + v96
				v99 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
				*(*int32)(unsafe.Add(mBase, uint32(v84))) = v99 + v96
				*(*int64)(unsafe.Add(mBase, uint32(v84+int32(0))+232)) = v79
				v107 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
				*(*int32)(unsafe.Add(mBase, uint32(v84))) = v107 + v96
				v113 = *(*int32)(unsafe.Add(mBase, _c_F_CopySendEndOfRow[2]))
				*(*int32)(unsafe.Add(mBase, _c_F_CopySendEndOfRow[2])) = v113 - v96
			}
		}
		v117 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
		v118 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v117))) = uint8(v118)
		*(*int32)(unsafe.Add(mBase, uint32(v4)+12)) = v118
		*(*int32)(unsafe.Add(mBase, uint32(v4)+4)) = v118
		return
	}
}
func F_CreateStandaloneExprContext(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int64
	_ = v21
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	v4 = F_palloc0(m, int32(72))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		*(*int64)(unsafe.Add(mBase, uint32(v4)+8)) = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v4))) = int64(382)
		v13 = *(*int32)(unsafe.Add(mBase, _c_F_CreateStandaloneExprContext[0]))
		*(*int32)(unsafe.Add(mBase, uint32(v4)+16)) = v13
		v19 = F_AllocSetContextCreateInternal(m, v13, int32(_a_F_CreateStandaloneExprContext_0), int32(0), int32(_a_F_CreateStandaloneExprContext_1), int32(_a_F_CreateStandaloneExprContext_2))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			v21 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v4)+24)) = v21
			*(*int32)(unsafe.Add(mBase, uint32(v4)+20)) = v19
			*(*int64)(unsafe.Add(mBase, uint32(v4)+32)) = v21
			v26 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v4)+40)) = v26
			*(*int64)(unsafe.Add(mBase, uint32(v4)+64)) = v21
			v30 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(v4)+52)) = uint8(v30)
			*(*int32)(unsafe.Add(mBase, uint32(v4)+48)) = v26
			*(*uint8)(unsafe.Add(mBase, uint32(v4)+44)) = uint8(v30)
			return v4
		}
	}
}
func F_CreateTemplateTupleDesc(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = F_palloc(m, l0*int32(116)+int32(20))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
		*(*int64)(unsafe.Add(mBase, uint32(v7)+12)) = int64(4294967295)
		*(*int64)(unsafe.Add(mBase, uint32(v7)+4)) = int64(-4294965047)
		return v7
	}
}
func F_calculate_relation_size(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int64
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v84 int64
	_ = v84
	v6 = m.G0
	v8 = v6 - int32(1248)
	m.G0 = v8
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_GetRelationPath(m, v8+int32(1176), v12, v13, v14, l1, l2)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v20 = int32(0)
	v24 = int64(0)
	goto L3
L3:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_calculate_relation_size[0]))
	if v26 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	m.G0 = v8 + int32(1248)
	return v24
L5:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	if v20 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L7
L9:
	;
	v55 = v8 + int32(144)
	v60 = F___fstatat(m, int32(-100), v55, v8+int32(48), int32(0))
	mBase = m.M
	goto L16
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v8 + int32(1176)
	v40 = F_pg_snprintf(m, v8+int32(144), int32(1024), int32(_a_F_calculate_relation_size_0), v8+int32(16))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v20
	*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v8 + int32(1176)
	v52 = F_pg_snprintf(m, v8+int32(144), int32(1024), int32(_a_F_calculate_relation_size_1), v8+int32(32))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L14
	}
L13:
	;
	goto L9
L14:
	;
	goto L9
L15:
	;
	goto L4
L16:
	;
	if v60 < int32(0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v64 = *(*int32)(unsafe.Add(mBase, _c_F_calculate_relation_size[1]))
	if v64 == int32(44) {
		goto L15
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v84 = *(*int64)(unsafe.Add(mBase, uint32(v8)+72))
	v20 = v20 + int32(1)
	v24 = v84 + v24
	goto L3
L20:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v55
	F_errmsg(m, int32(_a_F_calculate_relation_size_2), v8)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(_a_F_calculate_relation_size_3), int32(355), int32(_a_F_calculate_relation_size_4))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_calculate_table_size(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v11 int32
	_ = v11
	var v13 int64
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int64
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int64
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int64
	_ = v27
	var v28 int32
	_ = v28
	var v31 int64
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int64
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int64
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int64
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int64
	_ = v53
	var v54 int32
	_ = v54
	var v55 int64
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v67 int64
	_ = v67
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int64
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int64
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int64
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int64
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v103 int64
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v112 int64
	_ = v112
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v134 int64
	_ = v134
	v2 = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = F_calculate_relation_size(m, l0, v11, v2)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v19 = F_calculate_relation_size(m, l0, v17, int32(1))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v23 = F_calculate_relation_size(m, l0, v21, int32(2))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v27 = F_calculate_relation_size(m, l0, v25, int32(3))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v31 = v27 + (v23 + (v13 + v19))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+112))
	if v33 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v35 = F_relation_open(m, v33, int32(1))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	v134 = v31
	goto L8
L8:
	;
	return v134
L9:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v35)+20))
	v39 = F_calculate_relation_size(m, v35, v37, int32(0))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v35)+20))
	v43 = F_calculate_relation_size(m, v35, v41, int32(1))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v35)+20))
	v48 = F_calculate_relation_size(m, v35, v46, int32(2))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v35)+20))
	v53 = F_calculate_relation_size(m, v35, v51, int32(3))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v55 = v39 + v43 + v48 + v53
	v56 = F_RelationGetIndexList(m, v35)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L15
	}
L14:
	;
	F_list_free(m, v56)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L1
	} else {
		goto L27
	}
L15:
	;
	if v56 == int32(0) {
		v112 = v55
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	if v60 <= int32(0) {
		v112 = v55
		goto L14
	} else {
		goto L17
	}
L17:
	;
	v66 = v2
	v67 = v55
	goto L18
L18:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v56)+12))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v73+v66<<(uint(int32(2))%32))))
	v79 = F_relation_open(m, v77, int32(1))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L20
	}
L19:
	;
	v112 = v103
	goto L14
L20:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v79)+20))
	v83 = F_calculate_relation_size(m, v79, v81, int32(0))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v79)+20))
	v87 = F_calculate_relation_size(m, v79, v85, int32(1))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v79)+20))
	v91 = F_calculate_relation_size(m, v79, v89, int32(2))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v79)+20))
	v95 = F_calculate_relation_size(m, v79, v93, int32(3))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	F_relation_close(m, v79, int32(1))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v103 = v95 + (v91 + (v87 + (v67 + v83)))
	v105 = v66 + int32(1)
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	if v105 < v106 {
		v66 = v105
		v67 = v103
		goto L18
	} else {
		goto L26
	}
L26:
	;
	goto L19
L27:
	;
	F_relation_close(m, v35, int32(1))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v134 = v112 + v31
	goto L8
}
func F_calculate_tablespace_size(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int64
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v140 int64
	_ = v140
	var v141 int32
	_ = v141
	var v143 int64
	_ = v143
	var v144 int64
	_ = v144
	var v147 int64
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v155 int64
	_ = v155
	var v157 int32
	_ = v157
	var v161 int64
	_ = v161
	v4 = int64(0)
	v5 = m.G0
	v7 = v5 - int32(3216)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_calculate_tablespace_size[0]))
	if l0 == v10 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	switch l0 - int32(1663) {
	case 0:
		goto L13
	case 1:
		goto L12
	default:
		goto L11
	}
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_calculate_tablespace_size[1]))
	v15 = F_has_privs_of_role(m, v13, int32(3375))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int64(0)
L4:
	;
	if v15 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_calculate_tablespace_size[1]))
	v23 = F_object_aclcheck(m, int32(1213), l0, v21, int64(512))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	if v23 == int32(0) {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v28 = F_get_tablespace_name(m, l0)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L3
	} else {
		goto L8
	}
L8:
	;
	F_aclcheck_error(m, v23, int32(42), v28)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L3
	} else {
		goto L9
	}
L9:
	;
	goto L1
L10:
	;
	v64 = F_AllocateDir(m, v7+int32(2192))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L3
	} else {
		goto L18
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+40)) = int32(_a_F_calculate_tablespace_size_0)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+36)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = int32(_a_F_calculate_tablespace_size_1)
	v60 = F_pg_snprintf(m, v7+int32(2192), int32(1024), int32(_a_F_calculate_tablespace_size_2), v7+int32(32))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L3
	} else {
		goto L16
	}
L12:
	;
	v47 = F_pg_snprintf(m, v7+int32(2192), int32(1024), int32(_a_F_calculate_tablespace_size_3), int32(0))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L3
	} else {
		goto L15
	}
L13:
	;
	v40 = F_pg_snprintf(m, v7+int32(2192), int32(1024), int32(_a_F_calculate_tablespace_size_4), int32(0))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L3
	} else {
		goto L14
	}
L14:
	;
	goto L10
L15:
	;
	goto L10
L16:
	;
	goto L10
L17:
	;
	m.G0 = v7 + int32(3216)
	return v161
L18:
	;
	if v64 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v161 = int64(-1)
	goto L17
L20:
	;
	goto L21
L21:
	;
	v71 = F_ReadDir(m, v64, v7+int32(2192))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L3
	} else {
		goto L22
	}
L22:
	;
	if v71 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v73 = v71
	v76 = v4
	goto L26
L24:
	;
	v155 = v4
	goto L25
L25:
	;
	F_FreeDir(m, v64)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L3
	} else {
		goto L54
	}
L26:
	;
	v78 = *(*int32)(unsafe.Add(mBase, _c_F_calculate_tablespace_size[2]))
	if v78 != 0 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v155 = v147
	goto L25
L28:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L3
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+19)))
	if v81 != int32(46) {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	goto L30
L32:
	;
	v150 = F_ReadDir(m, v64, v7+int32(2192))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L3
	} else {
		goto L52
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v73 + int32(19)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v7 + int32(2192)
	v100 = v7 + int32(144)
	v105 = F_pg_snprintf(m, v100, int32(2048), int32(_a_F_calculate_tablespace_size_5), v7+int32(16))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L3
	} else {
		goto L38
	}
L34:
	;
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+20)))
	if v84 == int32(0) {
		v147 = v76
		goto L32
	} else {
		goto L35
	}
L35:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+20)))
	if v87 != int32(46) {
		goto L33
	} else {
		goto L36
	}
L36:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+21)))
	if v90 == int32(0) {
		v147 = v76
		goto L32
	} else {
		goto L37
	}
L37:
	;
	goto L33
L38:
	;
	v111 = F___fstatat(m, int32(-100), v100, v7+int32(48), int32(0))
	mBase = m.M
	goto L39
L39:
	;
	if v111 < int32(0) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v115 = *(*int32)(unsafe.Add(mBase, _c_F_calculate_tablespace_size[3]))
	if v115 == int32(44) {
		v147 = v76
		goto L32
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v7)+52))
	if v133&int32(_a_F_calculate_tablespace_size_6) == int32(_a_F_calculate_tablespace_size_7) {
		goto L48
	} else {
		goto L49
	}
L43:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L3
	} else {
		goto L44
	}
L44:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L3
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v100
	F_errmsg(m, int32(_a_F_calculate_tablespace_size_8), v7)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L3
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(_a_F_calculate_tablespace_size_9), int32(266), int32(_a_F_calculate_tablespace_size_10))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L3
	} else {
		goto L47
	}
L47:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L48:
	;
	v140 = F_db_dir_size(m, v7+int32(144))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L3
	} else {
		goto L51
	}
L49:
	;
	v143 = v76
	goto L50
L50:
	;
	v144 = *(*int64)(unsafe.Add(mBase, uint32(v7)+72))
	v147 = v143 + v144
	goto L32
L51:
	;
	v143 = v140 + v76
	goto L50
L52:
	;
	if v150 != 0 {
		v73 = v150
		v76 = v147
		goto L26
	} else {
		goto L53
	}
L53:
	;
	goto L27
L54:
	;
	v161 = v155
	goto L17
}
func F_cancel_parser_errposition_callback(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, _c_F_cancel_parser_errposition_callback[0])) = v3
	return
}
func F_cannotCastJsonbValue(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	switch l0 {
	case 0:
		v29 = int32(_a_F_cannotCastJsonbValue_0)
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return
		} else {
			F_errcode(m, int32(50856066))
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l1
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
				F_errmsg(m, v38, v7+int32(16))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_cannotCastJsonbValue_1), int32(2031), int32(_a_F_cannotCastJsonbValue_2))
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	case 1:
		v29 = int32(_a_F_cannotCastJsonbValue_3)
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return
		} else {
			F_errcode(m, int32(50856066))
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l1
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
				F_errmsg(m, v38, v7+int32(16))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_cannotCastJsonbValue_1), int32(2031), int32(_a_F_cannotCastJsonbValue_2))
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	case 2:
		v29 = int32(_a_F_cannotCastJsonbValue_4)
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return
		} else {
			F_errcode(m, int32(50856066))
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l1
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
				F_errmsg(m, v38, v7+int32(16))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_cannotCastJsonbValue_1), int32(2031), int32(_a_F_cannotCastJsonbValue_2))
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	case 3:
		v29 = int32(_a_F_cannotCastJsonbValue_5)
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return
		} else {
			F_errcode(m, int32(50856066))
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l1
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
				F_errmsg(m, v38, v7+int32(16))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_cannotCastJsonbValue_1), int32(2031), int32(_a_F_cannotCastJsonbValue_2))
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
			F_errmsg_internal(m, int32(_a_F_cannotCastJsonbValue_6), v7)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_cannotCastJsonbValue_1), int32(2034), int32(_a_F_cannotCastJsonbValue_2))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	case 16:
		v29 = int32(_a_F_cannotCastJsonbValue_7)
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return
		} else {
			F_errcode(m, int32(50856066))
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l1
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
				F_errmsg(m, v38, v7+int32(16))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_cannotCastJsonbValue_1), int32(2031), int32(_a_F_cannotCastJsonbValue_2))
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	case 17:
		v29 = int32(_a_F_cannotCastJsonbValue_8)
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return
		} else {
			F_errcode(m, int32(50856066))
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l1
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
				F_errmsg(m, v38, v7+int32(16))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_cannotCastJsonbValue_1), int32(2031), int32(_a_F_cannotCastJsonbValue_2))
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	case 18:
		v29 = int32(_a_F_cannotCastJsonbValue_9)
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return
		} else {
			F_errcode(m, int32(50856066))
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l1
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
				F_errmsg(m, v38, v7+int32(16))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_cannotCastJsonbValue_1), int32(2031), int32(_a_F_cannotCastJsonbValue_2))
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
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
func F_canonicalize_path_enc(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v174 int32
	_ = v174
	var v181 int32
	_ = v181
	var v187 int32
	_ = v187
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v245 int32
	_ = v245
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v264 int32
	_ = v264
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v279 int32
	_ = v279
	var v286 int32
	_ = v286
	var v292 int32
	_ = v292
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v356 int32
	_ = v356
	var v366 int32
	_ = v366
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v392 int32
	_ = v392
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	v2 = int32(0)
	v9 = F_strlen(m, l0)
	mBase = m.M
	if v9 < int32(2) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v39 != 0 {
		goto L7
	} else {
		goto L8
	}
L2:
	;
	v20 = l0 + v9 - int32(1)
	goto L3
L3:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if v23 != int32(47) {
		goto L1
	} else {
		goto L5
	}
L4:
	;
	goto L1
L5:
	;
	v26 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v20))) = uint8(v26)
	v29 = v20 - int32(1)
	if base.Ui32(l0) < base.Ui32(v29) {
		v20 = v29
		goto L3
	} else {
		goto L6
	}
L6:
	;
	goto L4
L7:
	;
	v41 = l0
	v43 = v2
	v44 = l0
	goto L10
L8:
	;
	v66 = l0
	goto L9
L9:
	;
	v73 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v66))) = uint8(v73)
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v75 != 0 {
		goto L17
	} else {
		goto L18
	}
L10:
	;
	v49 = v44 + int32(1)
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
	if base.B2i32(v50 == int32(47))&v43 != 0 {
		v44 = v49
		goto L10
	} else {
		goto L12
	}
L11:
	;
	v66 = v59
	goto L9
L12:
	;
	if v41 != v44 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v41))) = uint8(v50)
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
	v57 = v56
	goto L15
L14:
	;
	v57 = v50
	goto L15
L15:
	;
	v59 = v41 + int32(1)
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
	if v64 != 0 {
		v41 = v59
		v43 = base.B2i32(v57&int32(255) == int32(47))
		v44 = v49
		goto L10
	} else {
		goto L16
	}
L16:
	;
	goto L11
L17:
	;
	if v75 == int32(47) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	goto L19
L19:
	;
	return
L20:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	v83 = v82
	v84 = l0 + int32(1)
	v85 = int32(0)
	goto L22
L21:
	;
	v83 = v75
	v84 = l0
	v85 = int32(2)
	goto L22
L22:
	;
	if v83&int32(255) == int32(0) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	if l0 == v384 {
		goto L130
	} else {
		goto L131
	}
L24:
	;
	v384 = v84
	goto L23
L25:
	;
	goto L26
L26:
	;
	v91 = v84
	v92 = v83
	v95 = v84
	v96 = v2
	v97 = v85
	goto L27
L27:
	;
	v102 = v92
	v103 = v95
	goto L30
L28:
	;
	v384 = v375
	goto L23
L29:
	;
	v122 = int32(0)
	if v120&int32(255) != int32(46) {
		v136 = v122
		goto L37
	} else {
		goto L38
	}
L30:
	;
	v107 = v102 & int32(255)
	if v107 == int32(0) {
		v120 = v92
		v121 = v103
		goto L29
	} else {
		goto L32
	}
L31:
	;
	v115 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v103))) = uint8(v115)
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95))))
	v120 = v119
	v121 = v103 + int32(1)
	goto L29
L32:
	;
	if v107 != int32(47) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+1)))
	v102 = v112
	v103 = v103 + int32(1)
	goto L30
L34:
	;
	goto L35
L35:
	;
	goto L31
L36:
	;
	v382 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121))))
	if v382 != 0 {
		v91 = v375
		v92 = v382
		v95 = v121
		v96 = v380
		v97 = v381
		goto L27
	} else {
		goto L129
	}
L37:
	;
	switch v97 {
	case 0:
		goto L46
	case 1:
		goto L45
	case 2:
		goto L44
	case 3:
		goto L43
	case 4:
		goto L42
	default:
		v375 = v91
		v380 = v96
		v381 = v97
		goto L36
	}
L38:
	;
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95)+1)))
	if v127 == int32(0) {
		v375 = v91
		v380 = v96
		v381 = v97
		goto L36
	} else {
		goto L39
	}
L39:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95)+1)))
	if v130 != int32(46) {
		v136 = v122
		goto L37
	} else {
		goto L40
	}
L40:
	;
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95)+2)))
	v136 = base.B2i32(v133 == int32(0))
	goto L37
L41:
	;
	v375 = v366
	v380 = v371
	v381 = int32(3)
	goto L36
L42:
	;
	v342 = int32(47)
	*(*uint8)(unsafe.Add(mBase, uint32(v91))) = uint8(v342)
	v345 = v91 + int32(1)
	v346 = F_strlen(m, v95)
	mBase = m.M
	if v136 != 0 {
		goto L120
	} else {
		goto L121
	}
L43:
	;
	if v136 != 0 {
		goto L89
	} else {
		goto L90
	}
L44:
	;
	v235 = F_strlen(m, v95)
	mBase = m.M
	if v136 != 0 {
		goto L80
	} else {
		goto L81
	}
L45:
	;
	if v136 != 0 {
		goto L53
	} else {
		goto L54
	}
L46:
	;
	if v136 != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v375 = v91
	v380 = v96
	v381 = int32(0)
	goto L36
L48:
	;
	goto L49
L49:
	;
	v138 = F_strlen(m, v95)
	mBase = m.M
	v139 = int32(0)
	if base.B2i32(v138 == v139)|base.B2i32(v91 == v95) == v139 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	base.MemoryCopy(m, v91, v95, v138)
	goto L52
L51:
	;
	goto L52
L52:
	;
	v146 = int32(1)
	v375 = v91 + v138
	v380 = v96 + v146
	v381 = v146
	goto L36
L53:
	;
	v150 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v91))) = uint8(v150)
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v152 != 0 {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	goto L55
L55:
	;
	v219 = int32(47)
	*(*uint8)(unsafe.Add(mBase, uint32(v91))) = uint8(v219)
	v221 = F_strlen(m, v95)
	mBase = m.M
	v222 = int32(0)
	v225 = v91 + int32(1)
	if base.B2i32(v221 == v222)|base.B2i32(v225 == v95) == v222 {
		goto L77
	} else {
		goto L78
	}
L56:
	;
	v153 = F_strlen(m, l0)
	mBase = m.M
	v159 = v153 + l0
	goto L59
L57:
	;
	v208 = l0
	goto L58
L58:
	;
	v216 = v96 - int32(1)
	v375 = v208
	v380 = v216
	v381 = base.B2i32(v216 != int32(0))
	goto L36
L59:
	;
	v164 = v159 - int32(1)
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164))))
	if base.B2i32(v165 == int32(47))&base.B2i32(base.Ui32(l0) < base.Ui32(v164)) != 0 {
		v159 = v164
		goto L59
	} else {
		goto L61
	}
L60:
	;
	v174 = v164
	goto L62
L61:
	;
	goto L60
L62:
	;
	if base.Ui32(l0) < base.Ui32(v174) {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	v187 = v174
	goto L68
L64:
	;
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174))))
	if v181 != int32(47) {
		v174 = v174 - int32(1)
		goto L62
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	goto L63
L67:
	;
	goto L66
L68:
	;
	if base.Ui32(l0) < base.Ui32(v187) {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	if l0 == v187 {
		goto L74
	} else {
		goto L75
	}
L70:
	;
	v195 = v187 - int32(1)
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v195))))
	if v196 == int32(47) {
		v187 = v195
		goto L68
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	goto L69
L73:
	;
	goto L72
L74:
	;
	v204 = l0 + base.B2i32(v152 == int32(47))
	goto L76
L75:
	;
	v204 = v187
	goto L76
L76:
	;
	v205 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v204))) = uint8(v205)
	v208 = v204
	goto L58
L77:
	;
	base.MemoryCopy(m, v225, v95, v221)
	goto L79
L78:
	;
	goto L79
L79:
	;
	v231 = int32(1)
	v375 = v225 + v221
	v380 = v96 + v231
	v381 = v231
	goto L36
L80:
	;
	v236 = int32(0)
	if base.B2i32(v235 == v236)|base.B2i32(v91 == v95) == v236 {
		goto L83
	} else {
		goto L84
	}
L81:
	;
	goto L82
L82:
	;
	v245 = int32(0)
	if base.B2i32(v235 == v245)|base.B2i32(v91 == v95) == v245 {
		goto L86
	} else {
		goto L87
	}
L83:
	;
	base.MemoryCopy(m, v91, v95, v235)
	goto L85
L84:
	;
	goto L85
L85:
	;
	v375 = v91 + v235
	v380 = v96
	v381 = int32(4)
	goto L36
L86:
	;
	base.MemoryCopy(m, v91, v95, v235)
	goto L88
L87:
	;
	goto L88
L88:
	;
	v366 = v91 + v235
	v371 = v96 + int32(1)
	goto L41
L89:
	;
	v255 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v91))) = uint8(v255)
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v257 != 0 {
		goto L92
	} else {
		goto L93
	}
L90:
	;
	goto L91
L91:
	;
	v327 = int32(47)
	*(*uint8)(unsafe.Add(mBase, uint32(v91))) = uint8(v327)
	v329 = F_strlen(m, v95)
	mBase = m.M
	v330 = int32(0)
	v333 = v91 + int32(1)
	if base.B2i32(v329 == v330)|base.B2i32(v333 == v95) == v330 {
		goto L117
	} else {
		goto L118
	}
L92:
	;
	v258 = F_strlen(m, l0)
	mBase = m.M
	v264 = v258 + l0
	goto L95
L93:
	;
	v313 = l0
	goto L94
L94:
	;
	v321 = v96 - int32(1)
	if v321 != 0 {
		v366 = v313
		v371 = v321
		goto L41
	} else {
		goto L113
	}
L95:
	;
	v269 = v264 - int32(1)
	v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v269))))
	if base.B2i32(v270 == int32(47))&base.B2i32(base.Ui32(l0) < base.Ui32(v269)) != 0 {
		v264 = v269
		goto L95
	} else {
		goto L97
	}
L96:
	;
	v279 = v269
	goto L98
L97:
	;
	goto L96
L98:
	;
	if base.Ui32(l0) < base.Ui32(v279) {
		goto L100
	} else {
		goto L101
	}
L99:
	;
	v292 = v279
	goto L104
L100:
	;
	v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v279))))
	if v286 != int32(47) {
		v279 = v279 - int32(1)
		goto L98
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	goto L99
L103:
	;
	goto L102
L104:
	;
	if base.Ui32(l0) < base.Ui32(v292) {
		goto L106
	} else {
		goto L107
	}
L105:
	;
	if l0 == v292 {
		goto L110
	} else {
		goto L111
	}
L106:
	;
	v300 = v292 - int32(1)
	v301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v300))))
	if v301 == int32(47) {
		v292 = v300
		goto L104
	} else {
		goto L109
	}
L107:
	;
	goto L108
L108:
	;
	goto L105
L109:
	;
	goto L108
L110:
	;
	v309 = l0 + base.B2i32(v257 == int32(47))
	goto L112
L111:
	;
	v309 = v292
	goto L112
L112:
	;
	v310 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v309))) = uint8(v310)
	v313 = v309
	goto L94
L113:
	;
	if l0 == v313 {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v325 = int32(2)
	goto L116
L115:
	;
	v325 = int32(4)
	goto L116
L116:
	;
	v375 = v313
	v380 = int32(0)
	v381 = v325
	goto L36
L117:
	;
	base.MemoryCopy(m, v333, v95, v329)
	goto L119
L118:
	;
	goto L119
L119:
	;
	v366 = v333 + v329
	v371 = v96 + int32(1)
	goto L41
L120:
	;
	v347 = int32(0)
	if base.B2i32(v346 == v347)|base.B2i32(v345 == v95) == v347 {
		goto L123
	} else {
		goto L124
	}
L121:
	;
	goto L122
L122:
	;
	v356 = int32(0)
	if base.B2i32(v346 == v356)|base.B2i32(v345 == v95) == v356 {
		goto L126
	} else {
		goto L127
	}
L123:
	;
	base.MemoryCopy(m, v345, v95, v346)
	goto L125
L124:
	;
	goto L125
L125:
	;
	v375 = v346 + v345
	v380 = v96
	v381 = int32(4)
	goto L36
L126:
	;
	base.MemoryCopy(m, v345, v95, v346)
	goto L128
L127:
	;
	goto L128
L128:
	;
	v366 = v346 + v345
	v371 = int32(1)
	goto L41
L129:
	;
	goto L28
L130:
	;
	v392 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v384))) = uint8(v392)
	v396 = v384 + int32(1)
	goto L132
L131:
	;
	v396 = v384
	goto L132
L132:
	;
	v397 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v396))) = uint8(v397)
	goto L19
}
func F_cashsmaller(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int64
	_ = v5
	var v6 int32
	_ = v6
	var v7 int64
	_ = v7
	var v9 int64
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(v4)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
	if v5 < v7 {
		v9 = v5
	} else {
		v9 = v7
	}
	v10 = F_Int64GetDatum(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		return v10
	}
}
func F_cfunc_match(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v5 int64
	_ = v5
	var v7 int64
	_ = v7
	var v8 int64
	_ = v8
	var v11 int64
	_ = v11
	var v12 int64
	_ = v12
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	v7 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	v8 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	v12 = *(*int64)(unsafe.Add(mBase, uint32(l1)+16))
	if v4^v5|(v7^v8)|(v11^v12) == int64(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return int32(0)
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if int32(0) < v17 {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	goto L4
L4:
	;
	return int32(1)
L5:
	;
	goto L4
L6:
	;
	v20 = int32(28)
	v21 = l0 + v20
	v23 = l1 + v20
	v25 = v17 << (uint(int32(2)) % 32)
	if base.Ui32(int32(4)) <= base.Ui32(v25) {
		goto L12
	} else {
		goto L13
	}
L7:
	;
	goto L8
L8:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v89 != 0 {
		goto L28
	} else {
		goto L29
	}
L9:
	;
	if v87 != 0 {
		goto L5
	} else {
		goto L27
	}
L10:
	;
	v87 = int32(0)
	goto L9
L11:
	;
	v61 = v56
	v62 = v57
	v63 = v58
	goto L21
L12:
	;
	if (v21|v23)&int32(3) != 0 {
		v56 = v21
		v57 = v23
		v58 = v25
		goto L11
	} else {
		goto L15
	}
L13:
	;
	v49 = v21
	v50 = v23
	v51 = v25
	goto L14
L14:
	;
	if v51 == int32(0) {
		goto L10
	} else {
		goto L20
	}
L15:
	;
	v33 = v21
	v34 = v23
	v35 = v25
	goto L16
L16:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	if v38 != v39 {
		v56 = v33
		v57 = v34
		v58 = v35
		goto L11
	} else {
		goto L18
	}
L17:
	;
	v49 = v44
	v50 = v42
	v51 = v46
	goto L14
L18:
	;
	v41 = int32(4)
	v42 = v34 + v41
	v44 = v33 + v41
	v46 = v35 - v41
	if base.Ui32(int32(3)) < base.Ui32(v46) {
		v33 = v44
		v34 = v42
		v35 = v46
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	v56 = v49
	v57 = v50
	v58 = v51
	goto L11
L21:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62))))
	if v66 == v67 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v87 = v66 - v67
	goto L9
L23:
	;
	v69 = int32(1)
	v74 = v63 - v69
	if v74 != 0 {
		v61 = v61 + v69
		v62 = v62 + v69
		v63 = v74
		goto L21
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	goto L22
L26:
	;
	goto L10
L27:
	;
	goto L8
L28:
	;
	if v88 == int32(0) {
		goto L5
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	if v88 == int32(0) {
		goto L1
	} else {
		goto L47
	}
L31:
	;
	v92 = int32(0)
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	if v96 != v97 {
		v148 = v92
		goto L33
	} else {
		goto L34
	}
L32:
	;
	if v148 == int32(0) {
		goto L5
	} else {
		goto L46
	}
L33:
	;
	goto L32
L34:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v88)+4))
	if v99 != v100 {
		v148 = v92
		goto L33
	} else {
		goto L35
	}
L35:
	;
	if v96 <= int32(0) {
		v148 = int32(1)
		goto L33
	} else {
		goto L36
	}
L36:
	;
	v106 = v96 << (uint(int32(4)) % 32)
	v108 = int32(20)
	v114 = int32(0)
	goto L37
L37:
	;
	v121 = v114 * int32(100)
	v122 = v89 + v106 + v108 + v121
	v123 = int32(4)
	v125 = v121 + (v88 + v106 + v108)
	v128 = F_strcmp(m, v122+v123, v125+v123)
	mBase = m.M
	if v128 != 0 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v148 = int32(0)
	goto L33
L39:
	;
	goto L38
L40:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v122)+68))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v125)+68))
	if v129 != v130 {
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v122)+76))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v125)+76))
	if v132 != v133 {
		goto L39
	} else {
		goto L42
	}
L42:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v122)+96))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v125)+96))
	if v135 != v136 {
		goto L39
	} else {
		goto L43
	}
L43:
	;
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+91)))
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125)+91)))
	if v138 != v139 {
		goto L39
	} else {
		goto L44
	}
L44:
	;
	v141 = int32(1)
	v143 = v114 + v141
	if v96 != v143 {
		v114 = v143
		goto L37
	} else {
		goto L45
	}
L45:
	;
	v148 = v141
	goto L33
L46:
	;
	goto L1
L47:
	;
	goto L5
}
func F_cfunc_resolve_polymorphic_argtypes(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v54 int32
	_ = v54
	var v63 int32
	_ = v63
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v213 int64
	_ = v213
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v279 int32
	_ = v279
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v409 int32
	_ = v409
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v445 int32
	_ = v445
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v543 int32
	_ = v543
	var v546 int32
	_ = v546
	var v562 int32
	_ = v562
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v568 int32
	_ = v568
	var v584 int32
	_ = v584
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v627 int32
	_ = v627
	var v639 int32
	_ = v639
	var v674 int32
	_ = v674
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v717 int32
	_ = v717
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v737 int32
	_ = v737
	var v741 int32
	_ = v741
	var v787 int32
	_ = v787
	var v790 int32
	_ = v790
	var v794 int32
	_ = v794
	var v799 int32
	_ = v799
	v7 = int32(0)
	v39 = m.G0
	v41 = v39 - int32(16)
	m.G0 = v41
	if l4 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v787 = m.ExcPending
	if v787 != 0 {
		goto L80
	} else {
		goto L225
	}
L2:
	;
	m.G0 = v41 + int32(16)
	return
L3:
	;
	if l0 <= int32(0) {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v209 = m.G0
	v211 = v209 - int32(32)
	m.G0 = v211
	v213 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v211)+16)) = v213
	*(*int64)(unsafe.Add(mBase, uint32(v211)+24)) = v213
	*(*int64)(unsafe.Add(mBase, uint32(v211))) = v213
	*(*int64)(unsafe.Add(mBase, uint32(v211)+8)) = v213
	if l0 <= int32(0) {
		v639 = int32(1)
		goto L49
	} else {
		goto L50
	}
L6:
	;
	v45 = int32(0)
	if l0 != int32(1) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v54 = v45
	v63 = v7
	goto L10
L8:
	;
	v148 = v45
	goto L9
L9:
	;
	v184 = int32(23)
	v187 = l1 + v148<<(uint(int32(2))%32)
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v187)))
	if v188 <= int32(3830) {
		goto L40
	} else {
		goto L41
	}
L10:
	;
	v90 = int32(23)
	v93 = l1 + v54<<(uint(int32(2))%32)
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
	if v94 <= int32(3830) {
		goto L16
	} else {
		goto L17
	}
L11:
	;
	if l0&int32(1) == int32(0) {
		goto L2
	} else {
		goto L37
	}
L12:
	;
	v116 = int32(23)
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
	if v117 <= int32(3830) {
		goto L27
	} else {
		goto L28
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v93))) = v113
	goto L12
L14:
	;
	v113 = int32(3904)
	goto L13
L15:
	;
	v113 = int32(1007)
	goto L13
L16:
	;
	switch v94 - int32(2277) {
	case 0:
		goto L15
	case 1, 2, 3, 4, 5:
		goto L12
	case 6:
		v113 = v90
		goto L13
	default:
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	switch v94 - int32(_a_F_cfunc_resolve_polymorphic_argtypes_0) {
	case 0, 2:
		v113 = v90
		goto L13
	case 1:
		goto L15
	case 3:
		goto L14
	default:
		goto L21
	}
L19:
	;
	if base.B2i32(v94 == int32(2776))|base.B2i32(v94 == int32(3500)) != 0 {
		v113 = v90
		goto L13
	} else {
		goto L20
	}
L20:
	;
	goto L12
L21:
	;
	if v94 == int32(3831) {
		goto L14
	} else {
		goto L22
	}
L22:
	;
	if v94 != int32(_a_F_cfunc_resolve_polymorphic_argtypes_1) {
		goto L12
	} else {
		goto L23
	}
L23:
	;
	v113 = int32(_a_F_cfunc_resolve_polymorphic_argtypes_2)
	goto L13
L24:
	;
	v139 = int32(2)
	v140 = v54 + v139
	v142 = v63 + v139
	if v142 != l0&int32(2147483646) {
		v54 = v140
		v63 = v142
		goto L10
	} else {
		goto L36
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v93)+4)) = v136
	goto L24
L26:
	;
	v136 = int32(1007)
	goto L25
L27:
	;
	switch v117 - int32(2277) {
	case 0:
		goto L26
	case 1, 2, 3, 4, 5:
		goto L24
	case 6:
		v136 = v116
		goto L25
	default:
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	switch v117 - int32(_a_F_cfunc_resolve_polymorphic_argtypes_0) {
	case 0, 2:
		v136 = v116
		goto L25
	case 1:
		goto L26
	case 3:
		goto L32
	default:
		goto L33
	}
L30:
	;
	if base.B2i32(v117 == int32(2776))|base.B2i32(v117 == int32(3500)) != 0 {
		v136 = v116
		goto L25
	} else {
		goto L31
	}
L31:
	;
	goto L24
L32:
	;
	v136 = int32(3904)
	goto L25
L33:
	;
	if v117 == int32(3831) {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	if v117 != int32(_a_F_cfunc_resolve_polymorphic_argtypes_1) {
		goto L24
	} else {
		goto L35
	}
L35:
	;
	v136 = int32(_a_F_cfunc_resolve_polymorphic_argtypes_2)
	goto L25
L36:
	;
	goto L11
L37:
	;
	v148 = v140
	goto L9
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v187))) = v207
	goto L2
L39:
	;
	v207 = int32(1007)
	goto L38
L40:
	;
	switch v188 - int32(2277) {
	case 0:
		goto L39
	case 1, 2, 3, 4, 5:
		goto L2
	case 6:
		v207 = v184
		goto L38
	default:
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	switch v188 - int32(_a_F_cfunc_resolve_polymorphic_argtypes_0) {
	case 0, 2:
		v207 = v184
		goto L38
	case 1:
		goto L39
	case 3:
		goto L45
	default:
		goto L46
	}
L43:
	;
	if base.B2i32(v188 == int32(2776))|base.B2i32(v188 == int32(3500)) != 0 {
		v207 = v184
		goto L38
	} else {
		goto L44
	}
L44:
	;
	goto L2
L45:
	;
	v207 = int32(3904)
	goto L38
L46:
	;
	if v188 == int32(3831) {
		goto L45
	} else {
		goto L47
	}
L47:
	;
	if v188 != int32(_a_F_cfunc_resolve_polymorphic_argtypes_1) {
		goto L2
	} else {
		goto L48
	}
L48:
	;
	v207 = int32(_a_F_cfunc_resolve_polymorphic_argtypes_2)
	goto L38
L49:
	;
	m.G0 = v211 + int32(32)
	if v639 == int32(0) {
		goto L1
	} else {
		goto L213
	}
L50:
	;
	v235 = int32(0)
	v236 = v7
	v238 = v7
	v239 = v7
	v241 = v7
	v245 = v7
	v246 = v7
	v248 = v7
	v249 = v7
	v251 = v7
	v252 = v7
	v253 = v7
	v254 = v7
	v256 = v7
	v257 = v7
	v258 = v7
	v259 = v7
	v260 = v7
	v261 = v7
	goto L51
L51:
	;
	if l2 != 0 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v211)+8)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v211)+12)) = v392
	*(*int32)(unsafe.Add(mBase, uint32(v211)+4)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v211)+28)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v211)+24)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v211)+20)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v211))) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v211)+16)) = v385
	v409 = int32(1)
	if v377&v409 == int32(0) {
		v639 = v409
		goto L49
	} else {
		goto L127
	}
L53:
	;
	v264 = int32(*(*int8)(unsafe.Add(mBase, uint32(l2+v256))))
	v266 = v264
	goto L55
L54:
	;
	v266 = int32(105)
	goto L55
L55:
	;
	v269 = l1 + v256<<(uint(int32(2))%32)
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v269)))
	if v270 <= int32(3830) {
		goto L67
	} else {
		goto L68
	}
L56:
	;
	switch v266 - int32(111) {
	case 0, 5:
		v397 = v249
		goto L124
	default:
		goto L125
	}
L57:
	;
	v376 = v239
	v377 = v261
	v378 = v235
	v379 = v248
	v380 = v238
	v381 = v236
	v382 = v241
	v383 = v245
	v384 = v246
	v385 = v368
	v386 = v369
	v387 = v370
	v388 = v371
	v389 = v372
	v390 = v373
	v391 = v374
	v392 = v375
	goto L56
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v269))) = v342
	v368 = v350
	v369 = v351
	v370 = v352
	v371 = v353
	v372 = v354
	v373 = v355
	v374 = v356
	v375 = v357
	goto L57
L59:
	;
	v334 = int32(1)
	switch v266 - int32(111) {
	case 0, 5:
		v376 = v239
		v377 = v334
		v378 = v235
		v379 = v334
		v380 = v238
		v381 = v236
		v382 = v241
		v383 = v245
		v384 = v246
		v385 = v251
		v386 = v252
		v387 = v253
		v388 = v254
		v389 = v257
		v390 = v258
		v391 = v259
		v392 = v260
		goto L56
	default:
		goto L118
	}
L60:
	;
	v327 = int32(1)
	switch v266 - int32(111) {
	case 0, 5:
		v376 = v239
		v377 = v327
		v378 = v235
		v379 = v248
		v380 = v238
		v381 = v236
		v382 = v241
		v383 = v245
		v384 = v327
		v385 = v251
		v386 = v252
		v387 = v253
		v388 = v254
		v389 = v257
		v390 = v258
		v391 = v259
		v392 = v260
		goto L56
	default:
		goto L112
	}
L61:
	;
	v320 = int32(1)
	switch v266 - int32(111) {
	case 0, 5:
		v376 = v239
		v377 = v320
		v378 = v235
		v379 = v248
		v380 = v238
		v381 = v236
		v382 = v241
		v383 = v320
		v384 = v246
		v385 = v251
		v386 = v252
		v387 = v253
		v388 = v254
		v389 = v257
		v390 = v258
		v391 = v259
		v392 = v260
		goto L56
	default:
		goto L106
	}
L62:
	;
	v313 = int32(1)
	switch v266 - int32(111) {
	case 0, 5:
		v376 = v239
		v377 = v313
		v378 = v235
		v379 = v248
		v380 = v238
		v381 = v236
		v382 = v313
		v383 = v245
		v384 = v246
		v385 = v251
		v386 = v252
		v387 = v253
		v388 = v254
		v389 = v257
		v390 = v258
		v391 = v259
		v392 = v260
		goto L56
	default:
		goto L100
	}
L63:
	;
	v306 = int32(1)
	switch v266 - int32(111) {
	case 0, 5:
		v376 = v239
		v377 = v306
		v378 = v235
		v379 = v248
		v380 = v238
		v381 = v306
		v382 = v241
		v383 = v245
		v384 = v246
		v385 = v251
		v386 = v252
		v387 = v253
		v388 = v254
		v389 = v257
		v390 = v258
		v391 = v259
		v392 = v260
		goto L56
	default:
		goto L94
	}
L64:
	;
	if v257 != 0 {
		goto L89
	} else {
		goto L90
	}
L65:
	;
	v296 = int32(1)
	switch v266 - int32(111) {
	case 0, 5:
		v376 = v239
		v377 = v296
		v378 = v235
		v379 = v248
		v380 = v296
		v381 = v236
		v382 = v241
		v383 = v245
		v384 = v246
		v385 = v251
		v386 = v252
		v387 = v253
		v388 = v254
		v389 = v257
		v390 = v258
		v391 = v259
		v392 = v260
		goto L56
	default:
		goto L83
	}
L66:
	;
	if v251 != 0 {
		goto L77
	} else {
		goto L78
	}
L67:
	;
	switch v270 - int32(2277) {
	case 0:
		goto L65
	case 1, 2, 3, 4, 5:
		v376 = v239
		v377 = v261
		v378 = v235
		v379 = v248
		v380 = v238
		v381 = v236
		v382 = v241
		v383 = v245
		v384 = v246
		v385 = v251
		v386 = v252
		v387 = v253
		v388 = v254
		v389 = v257
		v390 = v258
		v391 = v259
		v392 = v260
		goto L56
	case 6:
		goto L70
	default:
		goto L71
	}
L68:
	;
	goto L69
L69:
	;
	switch v270 - int32(_a_F_cfunc_resolve_polymorphic_argtypes_0) {
	case 0, 2:
		goto L62
	case 1:
		goto L61
	case 3:
		goto L60
	default:
		goto L74
	}
L70:
	;
	v279 = int32(1)
	switch v266 - int32(111) {
	case 0, 5:
		v376 = v279
		v377 = v279
		v378 = v235
		v379 = v248
		v380 = v238
		v381 = v236
		v382 = v241
		v383 = v245
		v384 = v246
		v385 = v251
		v386 = v252
		v387 = v253
		v388 = v254
		v389 = v257
		v390 = v258
		v391 = v259
		v392 = v260
		goto L56
	default:
		goto L66
	}
L71:
	;
	if v270 == int32(2776) {
		goto L70
	} else {
		goto L72
	}
L72:
	;
	if v270 != int32(3500) {
		v368 = v251
		v369 = v252
		v370 = v253
		v371 = v254
		v372 = v257
		v373 = v258
		v374 = v259
		v375 = v260
		goto L57
	} else {
		goto L73
	}
L73:
	;
	goto L70
L74:
	;
	switch v270 - int32(_a_F_cfunc_resolve_polymorphic_argtypes_1) {
	case 0:
		goto L63
	case 1:
		goto L59
	default:
		goto L75
	}
L75:
	;
	if v270 != int32(3831) {
		v368 = v251
		v369 = v252
		v370 = v253
		v371 = v254
		v372 = v257
		v373 = v258
		v374 = v259
		v375 = v260
		goto L57
	} else {
		goto L76
	}
L76:
	;
	v289 = int32(1)
	switch v266 - int32(111) {
	case 0, 5:
		v376 = v239
		v377 = v289
		v378 = v289
		v379 = v248
		v380 = v238
		v381 = v236
		v382 = v241
		v383 = v245
		v384 = v246
		v385 = v251
		v386 = v252
		v387 = v253
		v388 = v254
		v389 = v257
		v390 = v258
		v391 = v259
		v392 = v260
		goto L56
	default:
		goto L64
	}
L77:
	;
	v342 = v251
	v350 = v251
	v351 = v252
	v352 = v253
	v353 = v254
	v354 = v257
	v355 = v258
	v356 = v259
	v357 = v260
	goto L58
L78:
	;
	goto L79
L79:
	;
	v293 = F_get_call_expr_argtype(m, l3, v249)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	return
L81:
	;
	if v293 != 0 {
		v342 = v293
		v350 = v293
		v351 = v252
		v352 = v253
		v353 = v254
		v354 = v257
		v355 = v258
		v356 = v259
		v357 = v260
		goto L58
	} else {
		goto L82
	}
L82:
	;
	v639 = int32(0)
	goto L49
L83:
	;
	if v252 != 0 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v342 = v252
	v350 = v251
	v351 = v252
	v352 = v253
	v353 = v254
	v354 = v257
	v355 = v258
	v356 = v259
	v357 = v260
	goto L58
L85:
	;
	goto L86
L86:
	;
	v300 = F_get_call_expr_argtype(m, l3, v249)
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L80
	} else {
		goto L87
	}
L87:
	;
	if v300 != 0 {
		v342 = v300
		v350 = v251
		v351 = v300
		v352 = v253
		v353 = v254
		v354 = v257
		v355 = v258
		v356 = v259
		v357 = v260
		goto L58
	} else {
		goto L88
	}
L88:
	;
	v639 = int32(0)
	goto L49
L89:
	;
	v342 = v257
	v350 = v251
	v351 = v252
	v352 = v253
	v353 = v254
	v354 = v257
	v355 = v258
	v356 = v259
	v357 = v260
	goto L58
L90:
	;
	goto L91
L91:
	;
	v303 = F_get_call_expr_argtype(m, l3, v249)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L80
	} else {
		goto L92
	}
L92:
	;
	if v303 != 0 {
		v342 = v303
		v350 = v251
		v351 = v252
		v352 = v253
		v353 = v254
		v354 = v303
		v355 = v258
		v356 = v259
		v357 = v260
		goto L58
	} else {
		goto L93
	}
L93:
	;
	v639 = int32(0)
	goto L49
L94:
	;
	if v258 != 0 {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v342 = v258
	v350 = v251
	v351 = v252
	v352 = v253
	v353 = v254
	v354 = v257
	v355 = v258
	v356 = v259
	v357 = v260
	goto L58
L96:
	;
	goto L97
L97:
	;
	v310 = F_get_call_expr_argtype(m, l3, v249)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L80
	} else {
		goto L98
	}
L98:
	;
	if v310 != 0 {
		v342 = v310
		v350 = v251
		v351 = v252
		v352 = v253
		v353 = v254
		v354 = v257
		v355 = v310
		v356 = v259
		v357 = v260
		goto L58
	} else {
		goto L99
	}
L99:
	;
	v639 = int32(0)
	goto L49
L100:
	;
	if v253 != 0 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v342 = v253
	v350 = v251
	v351 = v252
	v352 = v253
	v353 = v254
	v354 = v257
	v355 = v258
	v356 = v259
	v357 = v260
	goto L58
L102:
	;
	goto L103
L103:
	;
	v317 = F_get_call_expr_argtype(m, l3, v249)
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L80
	} else {
		goto L104
	}
L104:
	;
	if v317 != 0 {
		v342 = v317
		v350 = v251
		v351 = v252
		v352 = v317
		v353 = v254
		v354 = v257
		v355 = v258
		v356 = v259
		v357 = v260
		goto L58
	} else {
		goto L105
	}
L105:
	;
	v639 = int32(0)
	goto L49
L106:
	;
	if v254 != 0 {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v342 = v254
	v350 = v251
	v351 = v252
	v352 = v253
	v353 = v254
	v354 = v257
	v355 = v258
	v356 = v259
	v357 = v260
	goto L58
L108:
	;
	goto L109
L109:
	;
	v324 = F_get_call_expr_argtype(m, l3, v249)
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L80
	} else {
		goto L110
	}
L110:
	;
	if v324 != 0 {
		v342 = v324
		v350 = v251
		v351 = v252
		v352 = v253
		v353 = v324
		v354 = v257
		v355 = v258
		v356 = v259
		v357 = v260
		goto L58
	} else {
		goto L111
	}
L111:
	;
	v639 = int32(0)
	goto L49
L112:
	;
	if v259 != 0 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v342 = v259
	v350 = v251
	v351 = v252
	v352 = v253
	v353 = v254
	v354 = v257
	v355 = v258
	v356 = v259
	v357 = v260
	goto L58
L114:
	;
	goto L115
L115:
	;
	v331 = F_get_call_expr_argtype(m, l3, v249)
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L80
	} else {
		goto L116
	}
L116:
	;
	if v331 != 0 {
		v342 = v331
		v350 = v251
		v351 = v252
		v352 = v253
		v353 = v254
		v354 = v257
		v355 = v258
		v356 = v331
		v357 = v260
		goto L58
	} else {
		goto L117
	}
L117:
	;
	v639 = int32(0)
	goto L49
L118:
	;
	if v260 != 0 {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v342 = v260
	v350 = v251
	v351 = v252
	v352 = v253
	v353 = v254
	v354 = v257
	v355 = v258
	v356 = v259
	v357 = v260
	goto L58
L120:
	;
	goto L121
L121:
	;
	v338 = F_get_call_expr_argtype(m, l3, v249)
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L80
	} else {
		goto L122
	}
L122:
	;
	if v338 != 0 {
		v342 = v338
		v350 = v251
		v351 = v252
		v352 = v253
		v353 = v254
		v354 = v257
		v355 = v258
		v356 = v259
		v357 = v338
		goto L58
	} else {
		goto L123
	}
L123:
	;
	v639 = int32(0)
	goto L49
L124:
	;
	v399 = v256 + int32(1)
	if v399 != l0 {
		v235 = v378
		v236 = v381
		v238 = v380
		v239 = v376
		v241 = v382
		v245 = v383
		v246 = v384
		v248 = v379
		v249 = v397
		v251 = v385
		v252 = v386
		v253 = v387
		v254 = v388
		v256 = v399
		v257 = v389
		v258 = v390
		v259 = v391
		v260 = v392
		v261 = v377
		goto L51
	} else {
		goto L126
	}
L125:
	;
	v397 = v249 + int32(1)
	goto L124
L126:
	;
	goto L52
L127:
	;
	if base.B2i32(v385 == int32(0))&v376 != 0 {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	F_resolve_anyelement_from_others(m, v211+int32(16))
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L80
	} else {
		goto L131
	}
L129:
	;
	v422 = v386
	goto L130
L130:
	;
	if base.B2i32(v422 == int32(0))&v380 != 0 {
		goto L132
	} else {
		goto L133
	}
L131:
	;
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v211)+20))
	v422 = v421
	goto L130
L132:
	;
	F_resolve_anyarray_from_others(m, v211+int32(16))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L80
	} else {
		goto L135
	}
L133:
	;
	goto L134
L134:
	;
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v211)+24))
	if base.B2i32(v430 == int32(0))&v378 != 0 {
		goto L136
	} else {
		goto L137
	}
L135:
	;
	goto L134
L136:
	;
	F_resolve_anyrange_from_others(m, v211+int32(16))
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L80
	} else {
		goto L139
	}
L137:
	;
	goto L138
L138:
	;
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v211)+28))
	if base.B2i32(v438 == int32(0))&v381 != 0 {
		goto L140
	} else {
		goto L141
	}
L139:
	;
	goto L138
L140:
	;
	F_resolve_anymultirange_from_others(m, v211+int32(16))
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L80
	} else {
		goto L143
	}
L141:
	;
	goto L142
L142:
	;
	if base.B2i32(v387 == int32(0))&v382 != 0 {
		goto L144
	} else {
		goto L145
	}
L143:
	;
	goto L142
L144:
	;
	F_resolve_anyelement_from_others(m, v211)
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L80
	} else {
		goto L147
	}
L145:
	;
	v452 = v388
	goto L146
L146:
	;
	if base.B2i32(v452 == int32(0))&v383 != 0 {
		goto L148
	} else {
		goto L149
	}
L147:
	;
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v211)+4))
	v452 = v451
	goto L146
L148:
	;
	F_resolve_anyarray_from_others(m, v211)
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L80
	} else {
		goto L151
	}
L149:
	;
	goto L150
L150:
	;
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v211)+8))
	if base.B2i32(v458 == int32(0))&v384 != 0 {
		goto L152
	} else {
		goto L153
	}
L151:
	;
	goto L150
L152:
	;
	F_resolve_anyrange_from_others(m, v211)
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L80
	} else {
		goto L155
	}
L153:
	;
	goto L154
L154:
	;
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v211)+12))
	if base.B2i32(v464 == int32(0))&v379 != 0 {
		goto L156
	} else {
		goto L157
	}
L155:
	;
	goto L154
L156:
	;
	F_resolve_anymultirange_from_others(m, v211)
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L80
	} else {
		goto L159
	}
L157:
	;
	goto L158
L158:
	;
	v470 = int32(0)
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v211)+16))
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v211)+20))
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v211)+24))
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v211)+28))
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v211)))
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v211)+4))
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v211)+8))
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v211)+12))
	if l0 == int32(1) {
		v584 = v470
		goto L160
	} else {
		goto L161
	}
L159:
	;
	goto L158
L160:
	;
	v610 = l1 + v584<<(uint(int32(2))%32)
	v611 = *(*int32)(unsafe.Add(mBase, uint32(v610)))
	if v611 <= int32(3830) {
		goto L200
	} else {
		goto L201
	}
L161:
	;
	v498 = int32(0)
	v500 = v470
	goto L162
L162:
	;
	v526 = l1 + v500<<(uint(int32(2))%32)
	v527 = *(*int32)(unsafe.Add(mBase, uint32(v526)))
	if v527 <= int32(3830) {
		goto L172
	} else {
		goto L173
	}
L163:
	;
	if l0&int32(1) != 0 {
		v584 = v566
		goto L160
	} else {
		goto L197
	}
L164:
	;
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v526)+4))
	if v546 <= int32(3830) {
		goto L183
	} else {
		goto L184
	}
L165:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v526))) = v543
	goto L164
L166:
	;
	v543 = v478
	goto L165
L167:
	;
	v543 = v477
	goto L165
L168:
	;
	v543 = v476
	goto L165
L169:
	;
	v543 = v475
	goto L165
L170:
	;
	v543 = v474
	goto L165
L171:
	;
	v543 = v472
	goto L165
L172:
	;
	switch v527 - int32(2277) {
	case 0:
		goto L171
	case 1, 2, 3, 4, 5:
		goto L164
	case 6:
		v543 = v471
		goto L165
	default:
		goto L175
	}
L173:
	;
	goto L174
L174:
	;
	switch v527 - int32(_a_F_cfunc_resolve_polymorphic_argtypes_0) {
	case 0, 2:
		goto L169
	case 1:
		goto L168
	case 3:
		goto L167
	default:
		goto L177
	}
L175:
	;
	if base.B2i32(v527 == int32(2776))|base.B2i32(v527 == int32(3500)) != 0 {
		v543 = v471
		goto L165
	} else {
		goto L176
	}
L176:
	;
	goto L164
L177:
	;
	switch v527 - int32(_a_F_cfunc_resolve_polymorphic_argtypes_1) {
	case 0:
		goto L170
	case 1:
		goto L166
	default:
		goto L178
	}
L178:
	;
	if v527 == int32(3831) {
		v543 = v473
		goto L165
	} else {
		goto L179
	}
L179:
	;
	goto L164
L180:
	;
	v565 = int32(2)
	v566 = v500 + v565
	v568 = v498 + v565
	if v568 != l0&int32(2147483646) {
		v498 = v568
		v500 = v566
		goto L162
	} else {
		goto L196
	}
L181:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v526)+4)) = v562
	goto L180
L182:
	;
	v562 = v472
	goto L181
L183:
	;
	switch v546 - int32(2277) {
	case 0:
		goto L182
	case 1, 2, 3, 4, 5:
		goto L180
	case 6:
		v562 = v471
		goto L181
	default:
		goto L186
	}
L184:
	;
	goto L185
L185:
	;
	switch v546 - int32(_a_F_cfunc_resolve_polymorphic_argtypes_0) {
	case 0, 2:
		goto L189
	case 1:
		goto L190
	case 3:
		goto L191
	default:
		goto L192
	}
L186:
	;
	if base.B2i32(v546 == int32(2776))|base.B2i32(v546 == int32(3500)) != 0 {
		v562 = v471
		goto L181
	} else {
		goto L187
	}
L187:
	;
	goto L180
L188:
	;
	v562 = v474
	goto L181
L189:
	;
	v562 = v475
	goto L181
L190:
	;
	v562 = v476
	goto L181
L191:
	;
	v562 = v477
	goto L181
L192:
	;
	switch v546 - int32(_a_F_cfunc_resolve_polymorphic_argtypes_1) {
	case 0:
		goto L188
	case 1:
		goto L193
	default:
		goto L194
	}
L193:
	;
	v562 = v478
	goto L181
L194:
	;
	if v546 == int32(3831) {
		v562 = v473
		goto L181
	} else {
		goto L195
	}
L195:
	;
	goto L180
L196:
	;
	goto L163
L197:
	;
	v639 = v409
	goto L49
L198:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v610))) = v627
	v639 = v409
	goto L49
L199:
	;
	v627 = v472
	goto L198
L200:
	;
	switch v611 - int32(2277) {
	case 0:
		goto L199
	case 1, 2, 3, 4, 5:
		v639 = v409
		goto L49
	case 6:
		v627 = v471
		goto L198
	default:
		goto L203
	}
L201:
	;
	goto L202
L202:
	;
	switch v611 - int32(_a_F_cfunc_resolve_polymorphic_argtypes_0) {
	case 0, 2:
		goto L206
	case 1:
		goto L207
	case 3:
		goto L208
	default:
		goto L209
	}
L203:
	;
	if base.B2i32(v611 == int32(2776))|base.B2i32(v611 == int32(3500)) != 0 {
		v627 = v471
		goto L198
	} else {
		goto L204
	}
L204:
	;
	v639 = v409
	goto L49
L205:
	;
	v627 = v474
	goto L198
L206:
	;
	v627 = v475
	goto L198
L207:
	;
	v627 = v476
	goto L198
L208:
	;
	v627 = v477
	goto L198
L209:
	;
	switch v611 - int32(_a_F_cfunc_resolve_polymorphic_argtypes_1) {
	case 0:
		goto L205
	case 1:
		goto L210
	default:
		goto L211
	}
L210:
	;
	v627 = v478
	goto L198
L211:
	;
	if v611 == int32(3831) {
		v627 = v473
		goto L198
	} else {
		goto L212
	}
L212:
	;
	v639 = v409
	goto L49
L213:
	;
	if l0 <= int32(0) {
		goto L2
	} else {
		goto L214
	}
L214:
	;
	v674 = int32(0)
	v680 = v674
	v681 = v674
	goto L215
L215:
	;
	if l2 == int32(0) {
		goto L218
	} else {
		goto L219
	}
L216:
	;
	goto L2
L217:
	;
	v741 = v680 + int32(1)
	if v741 != l0 {
		v680 = v741
		v681 = v737
		goto L215
	} else {
		goto L224
	}
L218:
	;
	v722 = l1 + v680<<(uint(int32(2))%32)
	v723 = *(*int32)(unsafe.Add(mBase, uint32(v722)))
	if base.B2i32(v723 != int32(2287))&base.B2i32(v723 != int32(2249)) != 0 {
		goto L220
	} else {
		goto L221
	}
L219:
	;
	v717 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v680))))
	switch v717 - int32(111) {
	case 0, 5:
		v737 = v681
		goto L217
	default:
		goto L218
	}
L220:
	;
	v737 = v681 + int32(1)
	goto L217
L221:
	;
	v729 = F_get_call_expr_argtype(m, l3, v681)
	mBase = m.M
	v730 = m.ExcPending
	if v730 != 0 {
		goto L80
	} else {
		goto L222
	}
L222:
	;
	if v729 == int32(0) {
		goto L220
	} else {
		goto L223
	}
L223:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v722))) = v729
	goto L220
L224:
	;
	goto L216
L225:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v790 = m.ExcPending
	if v790 != 0 {
		goto L80
	} else {
		goto L226
	}
L226:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41))) = l5
	F_errmsg(m, int32(_a_F_cfunc_resolve_polymorphic_argtypes_3), v41)
	mBase = m.M
	v794 = m.ExcPending
	if v794 != 0 {
		goto L80
	} else {
		goto L227
	}
L227:
	;
	F_errfinish(m, int32(_a_F_cfunc_resolve_polymorphic_argtypes_4), int32(366), int32(_a_F_cfunc_resolve_polymorphic_argtypes_5))
	mBase = m.M
	v799 = m.ExcPending
	if v799 != 0 {
		goto L80
	} else {
		goto L228
	}
L228:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_char2wchar(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	if l1 == int32(0) {
		return
	} else {
		v9 = F_pnstrdup(m, l2, l3)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			if l4 == int32(0) {
				v13 = F_mbstowcs(m, l0, v9, l1)
				mBase = m.M
				v42 = v13
			} else {
				v14 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
				v17 = *(*int32)(unsafe.Add(mBase, _c_F_char2wchar[0]))
				if v14 != 0 {
					if v14 == int32(-1) {
						v22 = int32(_a_F_char2wchar_0)
					} else {
						v22 = v14
					}
					*(*int32)(unsafe.Add(mBase, _c_F_char2wchar[0])) = v22
				} else {
				}
				if v17 == int32(_a_F_char2wchar_0) {
					v27 = int32(-1)
				} else {
					v27 = v17
				}
				v28 = F_mbstowcs(m, l0, v9, l1)
				mBase = m.M
				if v27 != 0 {
					if v27 == int32(-1) {
						v36 = int32(_a_F_char2wchar_0)
					} else {
						v36 = v27
					}
					*(*int32)(unsafe.Add(mBase, _c_F_char2wchar[0])) = v36
				} else {
				}
				v42 = v28
			}
			F_pfree(m, v9)
			mBase = m.M
			v44 = m.ExcPending
			if v44 != 0 {
				return
			} else {
				if v42 != int32(-1) {
					return
				} else {
					F_pg_verifymbstr(m, l2, l3)
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return
						} else {
							F_errcode(m, int32(17301634))
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
								return
							} else {
								F_errmsg(m, int32(_a_F_char2wchar_1), int32(0))
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
									return
								} else {
									F_errhint(m, int32(_a_F_char2wchar_2), int32(0))
									mBase = m.M
									v63 = m.ExcPending
									if v63 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_char2wchar_3), int32(1001), int32(_a_F_char2wchar_4))
										mBase = m.M
										v68 = m.ExcPending
										if v68 != 0 {
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
	}
}
func F_charne(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	v3 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	return base.B2i32(v2 != v3)
}
func F_checkWellFormedRecursionWalker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v204 int32
	_ = v204
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v230 int32
	_ = v230
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v287 int32
	_ = v287
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v369 int32
	_ = v369
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v394 int32
	_ = v394
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v439 int32
	_ = v439
	var v455 int32
	_ = v455
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v494 int32
	_ = v494
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v530 int32
	_ = v530
	var v538 int32
	_ = v538
	var v547 int32
	_ = v547
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v556 int32
	_ = v556
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v577 int32
	_ = v577
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	v3 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(48)
	m.G0 = v16
	if l0 == v3 {
		v538 = v3
		goto L4
	} else {
		goto L5
	}
L1:
	;
	m.G0 = v16 + int32(48)
	return base.B2i32(v593 == int32(0)) & v594
L2:
	;
	v584 = F_raw_expression_tree_walker_impl(m, v570, int32(485), l1)
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L19
	} else {
		goto L169
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L19
	} else {
		goto L164
	}
L4:
	;
	v593 = v538
	v594 = v3
	goto L1
L5:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v20 <= int32(109) {
		goto L12
	} else {
		goto L13
	}
L6:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v338)+8))
	if v351 != 0 {
		v593 = v345
		v594 = v3
		goto L1
	} else {
		goto L122
	}
L7:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v197)+64))
	if v210 != 0 {
		goto L91
	} else {
		goto L92
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L19
	} else {
		goto L88
	}
L9:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v82+l0)))
	if v84 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L10:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	switch v41 {
	case 0:
		goto L22
	case 1:
		goto L23
	case 2:
		goto L24
	case 3:
		goto L25
	default:
		v168 = l0
		goto L8
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = int32(2)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v35 = F_checkWellFormedRecursionWalker(m, v34, l1)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L19
	} else {
		goto L20
	}
L12:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	switch v20 - int32(3) {
	case 0:
		v338 = l0
		v341 = v23
		v345 = v3
		goto L6
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18:
		v570 = l0
		v577 = v3
		goto L2
	case 19:
		goto L11
	default:
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	if v20 == int32(110) {
		v593 = v3
		v594 = v3
		goto L1
	} else {
		goto L17
	}
L15:
	;
	if v20 == int32(64) {
		goto L10
	} else {
		goto L16
	}
L16:
	;
	v570 = l0
	v577 = v3
	goto L2
L17:
	;
	if v20 == int32(141) {
		v197 = l0
		v204 = v3
		goto L7
	} else {
		goto L18
	}
L18:
	;
	v570 = l0
	v577 = v3
	goto L2
L19:
	;
	return int32(0)
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v23
	v82 = int32(12)
	goto L9
L21:
	;
	v82 = int32(28)
	goto L9
L22:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v76 = F_checkWellFormedRecursionWalker(m, v75, l1)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L19
	} else {
		goto L41
	}
L23:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v65 = F_checkWellFormedRecursionWalker(m, v64, l1)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L19
	} else {
		goto L36
	}
L24:
	;
	if v23 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L25:
	;
	if v23 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = int32(3)
	goto L28
L27:
	;
	goto L28
L28:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v47 = F_checkWellFormedRecursionWalker(m, v46, l1)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L19
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v23
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v51 = F_checkWellFormedRecursionWalker(m, v50, l1)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L19
	} else {
		goto L30
	}
L30:
	;
	goto L21
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = int32(3)
	goto L33
L32:
	;
	goto L33
L33:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v58 = F_checkWellFormedRecursionWalker(m, v57, l1)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L19
	} else {
		goto L34
	}
L34:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v61 = F_checkWellFormedRecursionWalker(m, v60, l1)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L19
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v23
	goto L21
L36:
	;
	if v23 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = int32(3)
	goto L39
L38:
	;
	goto L39
L39:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v72 = F_checkWellFormedRecursionWalker(m, v71, l1)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L19
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v23
	goto L21
L41:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v79 = F_checkWellFormedRecursionWalker(m, v78, l1)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L19
	} else {
		goto L42
	}
L42:
	;
	goto L21
L43:
	;
	v593 = int32(1)
	v594 = v3
	goto L1
L44:
	;
	goto L45
L45:
	;
	v88 = v84
	goto L46
L46:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v102 = int32(1)
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	if v103 <= int32(63) {
		goto L51
	} else {
		goto L52
	}
L47:
	;
	v593 = v102
	v594 = v3
	goto L1
L48:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v165+v88)))
	if v167 != 0 {
		v88 = v167
		goto L46
	} else {
		goto L87
	}
L49:
	;
	v165 = int32(28)
	goto L48
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = int32(2)
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v88)+20))
	v159 = F_checkWellFormedRecursionWalker(m, v158, l1)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L19
	} else {
		goto L86
	}
L51:
	;
	v107 = v103 - int32(3)
	if v107 != 0 {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	goto L53
L53:
	;
	if v103 != int32(64) {
		goto L60
	} else {
		goto L61
	}
L54:
	;
	if v107 == int32(19) {
		goto L57
	} else {
		goto L58
	}
L55:
	;
	v338 = v88
	v341 = v101
	v345 = v102
	goto L6
L57:
	;
	goto L50
L58:
	;
	v570 = v88
	v577 = v102
	goto L2
L60:
	;
	if v103 == int32(110) {
		v593 = v102
		v594 = v3
		goto L1
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v88)+4))
	switch v116 {
	case 0:
		goto L68
	case 1:
		goto L67
	case 2:
		goto L66
	case 3:
		goto L65
	default:
		v168 = v88
		goto L8
	}
L63:
	;
	if v103 == int32(141) {
		v197 = v88
		v204 = v102
		goto L7
	} else {
		goto L64
	}
L64:
	;
	v570 = v88
	v577 = v102
	goto L2
L65:
	;
	if v101 == int32(0) {
		goto L81
	} else {
		goto L82
	}
L66:
	;
	if v101 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L67:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v88)+12))
	v124 = F_checkWellFormedRecursionWalker(m, v123, l1)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L19
	} else {
		goto L71
	}
L68:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v88)+12))
	v118 = F_checkWellFormedRecursionWalker(m, v117, l1)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L19
	} else {
		goto L69
	}
L69:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v88)+16))
	v121 = F_checkWellFormedRecursionWalker(m, v120, l1)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L19
	} else {
		goto L70
	}
L70:
	;
	goto L49
L71:
	;
	if v101 == int32(0) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = int32(3)
	goto L74
L73:
	;
	goto L74
L74:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v88)+16))
	v131 = F_checkWellFormedRecursionWalker(m, v130, l1)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L19
	} else {
		goto L75
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v101
	goto L49
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = int32(3)
	goto L78
L77:
	;
	goto L78
L78:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v88)+12))
	v139 = F_checkWellFormedRecursionWalker(m, v138, l1)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L19
	} else {
		goto L79
	}
L79:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v88)+16))
	v142 = F_checkWellFormedRecursionWalker(m, v141, l1)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L19
	} else {
		goto L80
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v101
	goto L49
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = int32(3)
	goto L83
L82:
	;
	goto L83
L83:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v88)+12))
	v150 = F_checkWellFormedRecursionWalker(m, v149, l1)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L19
	} else {
		goto L84
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v101
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v88)+16))
	v154 = F_checkWellFormedRecursionWalker(m, v153, l1)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L19
	} else {
		goto L85
	}
L85:
	;
	goto L49
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v101
	v165 = int32(12)
	goto L48
L87:
	;
	goto L47
L88:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v168)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v185
	F_errmsg_internal(m, int32(_a_F_checkWellFormedRecursionWalker_0), v16+int32(32))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L19
	} else {
		goto L89
	}
L89:
	;
	F_errfinish(m, int32(_a_F_checkWellFormedRecursionWalker_1), int32(1179), int32(_a_F_checkWellFormedRecursionWalker_2))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L19
	} else {
		goto L90
	}
L90:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L91:
	;
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+8)))
	if v211 == int32(1) {
		goto L94
	} else {
		goto L95
	}
L92:
	;
	goto L93
L93:
	;
	F_checkWellFormedSelectStmt(m, v197, l1)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L19
	} else {
		goto L121
	}
L94:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v210)+4))
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v216 = F_lcons(m, v214, v215)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L19
	} else {
		goto L97
	}
L95:
	;
	goto L96
L96:
	;
	v271 = int32(0)
	v273 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v274 = F_lcons(m, v271, v273)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L19
	} else {
		goto L107
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v216
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v197)+64))
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v219)+4))
	if v220 == int32(0) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	F_checkWellFormedSelectStmt(m, v197, l1)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L19
	} else {
		goto L105
	}
L99:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v220)+4))
	if v223 <= int32(0) {
		goto L98
	} else {
		goto L100
	}
L100:
	;
	v230 = int32(0)
	goto L101
L101:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v220)+12))
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v240+v230<<(uint(int32(2))%32))))
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v244)+16))
	v246 = F_checkWellFormedRecursionWalker(m, v245, l1)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L19
	} else {
		goto L103
	}
L102:
	;
	goto L98
L103:
	;
	v249 = v230 + int32(1)
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v220)+4))
	if v249 < v250 {
		v230 = v249
		goto L101
	} else {
		goto L104
	}
L104:
	;
	goto L102
L105:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v268 = F_list_delete_first(m, v267)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L19
	} else {
		goto L106
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v268
	v593 = v204
	v594 = v3
	goto L1
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v274
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v197)+64))
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v277)+4))
	if v278 == int32(0) {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	F_checkWellFormedSelectStmt(m, v197, l1)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L19
	} else {
		goto L119
	}
L109:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v278)+4))
	if v281 <= int32(0) {
		goto L108
	} else {
		goto L110
	}
L110:
	;
	v287 = v271
	goto L111
L111:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v278)+12))
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v297+v287<<(uint(int32(2))%32))))
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v301)+16))
	v303 = F_checkWellFormedRecursionWalker(m, v302, l1)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L19
	} else {
		goto L113
	}
L112:
	;
	goto L108
L113:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v305 != 0 {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v305)+12))
	v308 = v306
	goto L116
L115:
	;
	v308 = int32(0)
	goto L116
L116:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v308)))
	v310 = F_lappend(m, v309, v301)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L19
	} else {
		goto L117
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v308))) = v310
	v314 = v287 + int32(1)
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v278)+4))
	if v314 < v315 {
		v287 = v314
		goto L111
	} else {
		goto L118
	}
L118:
	;
	goto L112
L119:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v333 = F_list_delete_first(m, v332)
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L19
	} else {
		goto L120
	}
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v333
	v538 = v204
	goto L4
L121:
	;
	v593 = v204
	v594 = v3
	goto L1
L122:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v352 == int32(0) {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v338)+12))
	v471 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v472 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v471+v472*int32(12))))
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v476)+4))
	v480 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v470))))
	v483 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v477))))
	if base.B2i32(v480 == int32(0))|base.B2i32(v480 != v483) != 0 {
		v501 = v480
		v502 = v483
		goto L150
	} else {
		goto L151
	}
L124:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v352)+4))
	if v355 <= int32(0) {
		goto L123
	} else {
		goto L125
	}
L125:
	;
	v358 = int32(0)
	if v358 < v355 {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v361 = v355
	goto L128
L127:
	;
	v361 = v358
	goto L128
L128:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v352)+12))
	v369 = v3
	goto L129
L129:
	;
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v362+v369<<(uint(int32(2))%32))))
	if v379 == int32(0) {
		goto L131
	} else {
		goto L132
	}
L130:
	;
	goto L123
L131:
	;
	v455 = v369 + int32(1)
	if v455 != v361 {
		v369 = v455
		goto L129
	} else {
		goto L148
	}
L132:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v379)+4))
	if v382 <= int32(0) {
		goto L131
	} else {
		goto L133
	}
L133:
	;
	v385 = int32(0)
	if v385 < v382 {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v388 = v382
	goto L136
L135:
	;
	v388 = v385
	goto L136
L136:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v338)+12))
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v379)+12))
	v394 = int32(0)
	goto L137
L137:
	;
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v390+v394<<(uint(int32(2))%32))))
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v408)+4))
	v412 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v389))))
	v415 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v409))))
	if base.B2i32(v412 == int32(0))|base.B2i32(v412 != v415) != 0 {
		v433 = v412
		v434 = v415
		goto L140
	} else {
		goto L141
	}
L138:
	;
	goto L131
L139:
	;
	if v433-v434 == int32(0) {
		v538 = v345
		goto L4
	} else {
		goto L146
	}
L140:
	;
	goto L139
L141:
	;
	v418 = v389
	v419 = v409
	goto L142
L142:
	;
	v422 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v419)+1)))
	v423 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v418)+1)))
	if v423 == int32(0) {
		v433 = v423
		v434 = v422
		goto L140
	} else {
		goto L144
	}
L143:
	;
	v433 = v423
	v434 = v422
	goto L140
L144:
	;
	v426 = int32(1)
	if v423 == v422 {
		v418 = v418 + v426
		v419 = v419 + v426
		goto L142
	} else {
		goto L145
	}
L145:
	;
	goto L143
L146:
	;
	v439 = v394 + int32(1)
	if v439 != v388 {
		v394 = v439
		goto L137
	} else {
		goto L147
	}
L147:
	;
	goto L138
L148:
	;
	goto L130
L149:
	;
	if v501-v502 != 0 {
		v593 = v345
		v594 = v3
		goto L1
	} else {
		goto L156
	}
L150:
	;
	goto L149
L151:
	;
	v486 = v470
	v487 = v477
	goto L152
L152:
	;
	v490 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v487)+1)))
	v491 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v486)+1)))
	if v491 == int32(0) {
		v501 = v491
		v502 = v490
		goto L150
	} else {
		goto L154
	}
L153:
	;
	v501 = v491
	v502 = v490
	goto L150
L154:
	;
	v494 = int32(1)
	if v491 == v490 {
		v486 = v486 + v494
		v487 = v487 + v494
		goto L152
	} else {
		goto L155
	}
L155:
	;
	goto L153
L156:
	;
	if v341 != 0 {
		goto L3
	} else {
		goto L157
	}
L157:
	;
	v504 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v506 = v504 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v506
	if v506 < int32(2) {
		v593 = v345
		v594 = v3
		goto L1
	} else {
		goto L158
	}
L158:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L19
	} else {
		goto L159
	}
L159:
	;
	F_errcode(m, int32(151388292))
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L19
	} else {
		goto L160
	}
L160:
	;
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v476)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v517
	F_errmsg(m, int32(_a_F_checkWellFormedRecursionWalker_3), v16)
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L19
	} else {
		goto L161
	}
L161:
	;
	v522 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v338)+24))
	F_parser_errposition(m, v522, v523)
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L19
	} else {
		goto L162
	}
L162:
	;
	F_errfinish(m, int32(_a_F_checkWellFormedRecursionWalker_1), int32(1077), int32(_a_F_checkWellFormedRecursionWalker_2))
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L19
	} else {
		goto L163
	}
L163:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L164:
	;
	F_errcode(m, int32(151388292))
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L19
	} else {
		goto L165
	}
L165:
	;
	v551 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v476)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v552
	v556 = *(*int32)(unsafe.Add(mBase, uint32(v551<<(uint(int32(2))%32))+uint32(_c_F_checkWellFormedRecursionWalker[0])))
	F_errmsg(m, v556, v16+int32(16))
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L19
	} else {
		goto L166
	}
L166:
	;
	v561 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v338)+24))
	F_parser_errposition(m, v561, v562)
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L19
	} else {
		goto L167
	}
L167:
	;
	F_errfinish(m, int32(_a_F_checkWellFormedRecursionWalker_1), int32(1069), int32(_a_F_checkWellFormedRecursionWalker_2))
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L19
	} else {
		goto L168
	}
L168:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L169:
	;
	v593 = v577
	v594 = v584
	goto L1
}
func F_checkWellFormedSelectStmt(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
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
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
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
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
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
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v8 != 0 {
		v10 = F_raw_expression_tree_walker_impl(m, l0, int32(485), l1)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return
		} else {
			m.G0 = v6 + int32(16)
			return
		}
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
		switch v12 {
		case 0, 1:
			v80 = F_raw_expression_tree_walker_impl(m, l0, int32(485), l1)
			mBase = m.M
			v81 = m.ExcPending
			if v81 != 0 {
				return
			} else {
				m.G0 = v6 + int32(16)
				return
			}
		case 2:
			v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+72)))
			if v13 == int32(1) {
				*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = int32(4)
			} else {
			}
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
			v19 = F_checkWellFormedRecursionWalker(m, v18, l1)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
				v22 = F_checkWellFormedRecursionWalker(m, v21, l1)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = int32(0)
					v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
					v27 = F_checkWellFormedRecursionWalker(m, v26, l1)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return
					} else {
						v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
						v30 = F_checkWellFormedRecursionWalker(m, v29, l1)
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return
						} else {
							v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
							v33 = F_checkWellFormedRecursionWalker(m, v32, l1)
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return
							} else {
								v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
								v36 = F_checkWellFormedRecursionWalker(m, v35, l1)
								mBase = m.M
								v37 = m.ExcPending
								if v37 != 0 {
									return
								} else {
									m.G0 = v6 + int32(16)
									return
								}
							}
						}
					}
				}
			}
		case 3:
			v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+72)))
			if v38 == int32(1) {
				*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = int32(5)
			} else {
			}
			v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
			v44 = F_checkWellFormedRecursionWalker(m, v43, l1)
			mBase = m.M
			v45 = m.ExcPending
			if v45 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = int32(5)
				v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
				v49 = F_checkWellFormedRecursionWalker(m, v48, l1)
				mBase = m.M
				v50 = m.ExcPending
				if v50 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = int32(0)
					v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
					v54 = F_checkWellFormedRecursionWalker(m, v53, l1)
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return
					} else {
						v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
						v57 = F_checkWellFormedRecursionWalker(m, v56, l1)
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return
						} else {
							v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
							v60 = F_checkWellFormedRecursionWalker(m, v59, l1)
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return
							} else {
								v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
								v63 = F_checkWellFormedRecursionWalker(m, v62, l1)
								mBase = m.M
								v64 = m.ExcPending
								if v64 != 0 {
									return
								} else {
									m.G0 = v6 + int32(16)
									return
								}
							}
						}
					}
				}
			}
		default:
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v68 = m.ExcPending
			if v68 != 0 {
				return
			} else {
				v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = v69
				F_errmsg_internal(m, int32(_a_F_checkWellFormedSelectStmt_0), v6)
				mBase = m.M
				v73 = m.ExcPending
				if v73 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_checkWellFormedSelectStmt_1), int32(1267), int32(_a_F_checkWellFormedSelectStmt_2))
					mBase = m.M
					v78 = m.ExcPending
					if v78 != 0 {
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
func F_check_createrole_self_grant(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v78 int32
	_ = v78
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v120 int32
	_ = v120
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v174 int32
	_ = v174
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v195 int32
	_ = v195
	v4 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v14 = F_pstrdup(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v11 + int32(16)
	return v195
L2:
	;
	return int32(0)
L3:
	;
	v21 = F_SplitIdentifierString(m, v14, int32(44), v11+int32(12))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	if v21 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_check_createrole_self_grant[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_check_createrole_self_grant[1])) = v26
	goto L8
L6:
	;
	goto L7
L7:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	if v40 == int32(0) {
		v158 = v4
		goto L13
	} else {
		goto L14
	}
L8:
	;
	v32 = F_format_elog_string(m, int32(_a_F_check_createrole_self_grant_0), int32(0))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_createrole_self_grant[2])) = v32
	F_pfree(m, v14)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	F_list_free(m, v37)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	v195 = v4
	goto L1
L12:
	;
	v174 = *(*int32)(unsafe.Add(mBase, _c_F_check_createrole_self_grant[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_check_createrole_self_grant[1])) = v174
	goto L53
L13:
	;
	F_pfree(m, v14)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L2
	} else {
		goto L49
	}
L14:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	if v43 <= int32(0) {
		v158 = v4
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v47 = int32(0)
	v53 = v4
	goto L16
L16:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v55+v47<<(uint(int32(2))%32))))
	v63 = v59
	v64 = int32(_a_F_check_createrole_self_grant_1)
	goto L19
L17:
	;
	v158 = v147
	goto L13
L18:
	;
	if v101 != 0 {
		goto L31
	} else {
		goto L32
	}
L19:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64))))
	if v67 == v68 {
		v90 = v67
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v101 = int32(0)
	goto L18
L21:
	;
	v92 = int32(1)
	if v90 != 0 {
		v63 = v63 + v92
		v64 = v64 + v92
		goto L19
	} else {
		goto L30
	}
L22:
	;
	if base.Ui32((v67-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v78 = v67 | int32(32)
	goto L25
L24:
	;
	v78 = v67
	goto L25
L25:
	;
	if base.Ui32((v68-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v87 = v68 | int32(32)
	goto L28
L27:
	;
	v87 = v68
	goto L28
L28:
	;
	if v78 == v87 {
		v90 = v78
		goto L21
	} else {
		goto L29
	}
L29:
	;
	v101 = v78 - v87
	goto L18
L30:
	;
	goto L20
L31:
	;
	v105 = v59
	v106 = int32(_a_F_check_createrole_self_grant_2)
	goto L35
L32:
	;
	v146 = int32(4)
	goto L33
L33:
	;
	v147 = v146 | v53
	v149 = v47 + int32(1)
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	if v149 < v150 {
		v47 = v149
		v53 = v147
		goto L16
	} else {
		goto L48
	}
L34:
	;
	if v143 != 0 {
		goto L12
	} else {
		goto L47
	}
L35:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105))))
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106))))
	if v109 == v110 {
		v132 = v109
		goto L37
	} else {
		goto L38
	}
L36:
	;
	v143 = int32(0)
	goto L34
L37:
	;
	v134 = int32(1)
	if v132 != 0 {
		v105 = v105 + v134
		v106 = v106 + v134
		goto L35
	} else {
		goto L46
	}
L38:
	;
	if base.Ui32((v109-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v120 = v109 | int32(32)
	goto L41
L40:
	;
	v120 = v109
	goto L41
L41:
	;
	if base.Ui32((v110-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v129 = v110 | int32(32)
	goto L44
L43:
	;
	v129 = v110
	goto L44
L44:
	;
	if v120 == v129 {
		v132 = v120
		goto L37
	} else {
		goto L45
	}
L45:
	;
	v143 = v120 - v129
	goto L34
L46:
	;
	goto L36
L47:
	;
	v146 = int32(2)
	goto L33
L48:
	;
	goto L17
L49:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	F_list_free(m, v162)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L2
	} else {
		goto L50
	}
L50:
	;
	v166 = F_guc_malloc(m, int32(4))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L2
	} else {
		goto L51
	}
L51:
	;
	if v166 == int32(0) {
		v195 = v4
		goto L1
	} else {
		goto L52
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v166))) = v158
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v166
	v195 = int32(1)
	goto L1
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v59
	v180 = F_format_elog_string(m, int32(_a_F_check_createrole_self_grant_3), v11)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L2
	} else {
		goto L54
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_createrole_self_grant[2])) = v180
	F_pfree(m, v14)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L2
	} else {
		goto L55
	}
L55:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	F_list_free(m, v185)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L2
	} else {
		goto L56
	}
L56:
	;
	v195 = v4
	goto L1
}
func F_check_debug_io_direct(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v81 int32
	_ = v81
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v104 int32
	_ = v104
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v126 int32
	_ = v126
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v149 int32
	_ = v149
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v170 int32
	_ = v170
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v223 int32
	_ = v223
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v244 int32
	_ = v244
	v4 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v14 = F_pstrdup(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v11 + int32(32)
	return v244
L2:
	;
	return int32(0)
L3:
	;
	v20 = F_SplitGUCList(m, v14, v11+int32(28))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	if v20 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_check_debug_io_direct[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_check_debug_io_direct[1])) = v25
	goto L8
L6:
	;
	goto L7
L7:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
	if v42 == int32(0) {
		v207 = v4
		goto L13
	} else {
		goto L14
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = int32(_a_F_check_debug_io_direct_0)
	v34 = F_format_elog_string(m, int32(_a_F_check_debug_io_direct_1), v11+int32(16))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_debug_io_direct[2])) = v34
	F_pfree(m, v14)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
	F_list_free(m, v39)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	v244 = v4
	goto L1
L12:
	;
	v223 = *(*int32)(unsafe.Add(mBase, _c_F_check_debug_io_direct[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_check_debug_io_direct[1])) = v223
	goto L66
L13:
	;
	F_pfree(m, v14)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L2
	} else {
		goto L62
	}
L14:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	if v45 <= int32(0) {
		v207 = v4
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v49 = int32(0)
	v55 = v4
	goto L16
L16:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v58+v49<<(uint(int32(2))%32))))
	v66 = v62
	v67 = int32(_a_F_check_debug_io_direct_2)
	goto L20
L17:
	;
	v207 = v196
	goto L13
L18:
	;
	v196 = v195 | v55
	v198 = v49 + int32(1)
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	if v198 < v199 {
		v49 = v198
		v55 = v196
		goto L16
	} else {
		goto L61
	}
L19:
	;
	if v104 == int32(0) {
		v195 = int32(1)
		goto L18
	} else {
		goto L32
	}
L20:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67))))
	if v70 == v71 {
		v93 = v70
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v104 = int32(0)
	goto L19
L22:
	;
	v95 = int32(1)
	if v93 != 0 {
		v66 = v66 + v95
		v67 = v67 + v95
		goto L20
	} else {
		goto L31
	}
L23:
	;
	if base.Ui32((v70-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v81 = v70 | int32(32)
	goto L26
L25:
	;
	v81 = v70
	goto L26
L26:
	;
	if base.Ui32((v71-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v90 = v71 | int32(32)
	goto L29
L28:
	;
	v90 = v71
	goto L29
L29:
	;
	if v81 == v90 {
		v93 = v81
		goto L22
	} else {
		goto L30
	}
L30:
	;
	v104 = v81 - v90
	goto L19
L31:
	;
	goto L21
L32:
	;
	v111 = v62
	v112 = int32(_a_F_check_debug_io_direct_3)
	goto L34
L33:
	;
	if v149 == int32(0) {
		v195 = int32(2)
		goto L18
	} else {
		goto L46
	}
L34:
	;
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111))))
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112))))
	if v115 == v116 {
		v138 = v115
		goto L36
	} else {
		goto L37
	}
L35:
	;
	v149 = int32(0)
	goto L33
L36:
	;
	v140 = int32(1)
	if v138 != 0 {
		v111 = v111 + v140
		v112 = v112 + v140
		goto L34
	} else {
		goto L45
	}
L37:
	;
	if base.Ui32((v115-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v126 = v115 | int32(32)
	goto L40
L39:
	;
	v126 = v115
	goto L40
L40:
	;
	if base.Ui32((v116-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v135 = v116 | int32(32)
	goto L43
L42:
	;
	v135 = v116
	goto L43
L43:
	;
	if v126 == v135 {
		v138 = v126
		goto L36
	} else {
		goto L44
	}
L44:
	;
	v149 = v126 - v135
	goto L33
L45:
	;
	goto L35
L46:
	;
	v155 = v62
	v156 = int32(_a_F_check_debug_io_direct_4)
	goto L48
L47:
	;
	if v193 != 0 {
		goto L12
	} else {
		goto L60
	}
L48:
	;
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155))))
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156))))
	if v159 == v160 {
		v182 = v159
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v193 = int32(0)
	goto L47
L50:
	;
	v184 = int32(1)
	if v182 != 0 {
		v155 = v155 + v184
		v156 = v156 + v184
		goto L48
	} else {
		goto L59
	}
L51:
	;
	if base.Ui32((v159-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v170 = v159 | int32(32)
	goto L54
L53:
	;
	v170 = v159
	goto L54
L54:
	;
	if base.Ui32((v160-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v179 = v160 | int32(32)
	goto L57
L56:
	;
	v179 = v160
	goto L57
L57:
	;
	if v170 == v179 {
		v182 = v170
		goto L50
	} else {
		goto L58
	}
L58:
	;
	v193 = v170 - v179
	goto L47
L59:
	;
	goto L49
L60:
	;
	v195 = int32(4)
	goto L18
L61:
	;
	goto L17
L62:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
	F_list_free(m, v211)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L2
	} else {
		goto L63
	}
L63:
	;
	v215 = F_guc_malloc(m, int32(4))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L2
	} else {
		goto L64
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v215
	if v215 == int32(0) {
		v244 = v4
		goto L1
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v215))) = v207
	v244 = int32(1)
	goto L1
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v62
	v229 = F_format_elog_string(m, int32(_a_F_check_debug_io_direct_5), v11)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L2
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_debug_io_direct[2])) = v229
	F_pfree(m, v14)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L2
	} else {
		goto L68
	}
L68:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
	F_list_free(m, v234)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L2
	} else {
		goto L69
	}
L69:
	;
	v244 = v4
	goto L1
}
func F_check_exclusion_constraint(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) {
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	v9 = int32(0)
	v12 = F_check_exclusion_or_unique_constraint(m, l0, l1, l2, l3, l4, l5, l6, l7, v9, v9, v9)
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		return
	}
}
func F_check_indirection(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	v3 = int32(0)
	if l0 == v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return l0
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v10 <= int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v13 = int32(0)
	if v13 < v10 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v16 = v10
	goto L6
L5:
	;
	v16 = v13
	goto L6
L6:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v23 = v3
	goto L7
L7:
	;
	v28 = v23 << (uint(int32(2)) % 32)
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v19+v28)))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	if base.B2i32(v31 == int32(77))&base.B2i32(v28+int32(4) < v10<<(uint(int32(2))%32)) == int32(0) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	F_scanner_yyerror(m, int32(_a_F_check_indirection_0), l1)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L13
	} else {
		goto L14
	}
L9:
	;
	v41 = v23 + int32(1)
	if v16 != v41 {
		v23 = v41
		goto L7
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	goto L8
L12:
	;
	goto L1
L13:
	;
	return int32(0)
L14:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_check_lateral_ref_ok(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
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
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+22)))
	if v11 != int32(1) {
		m.G0 = v9 + int32(32)
		return
	} else {
		v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+23)))
		if v14 != 0 {
			m.G0 = v9 + int32(32)
			return
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				F_errcode(m, int32(_a_F_check_lateral_ref_ok_0))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v17
					F_errmsg(m, int32(_a_F_check_lateral_ref_ok_1), v9+int32(16))
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return
					} else {
						v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
						if v31 == int32(0) {
							F_errdetail(m, int32(_a_F_check_lateral_ref_ok_2), int32(0))
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
								return
							} else {
								F_parser_errposition(m, l0, l2)
								mBase = m.M
								v45 = m.ExcPending
								if v45 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_check_lateral_ref_ok_3), int32(509), int32(_a_F_check_lateral_ref_ok_4))
									mBase = m.M
									v50 = m.ExcPending
									if v50 != 0 {
										return
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v34 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
							if v15 != v34 {
								F_errdetail(m, int32(_a_F_check_lateral_ref_ok_2), int32(0))
								mBase = m.M
								v43 = m.ExcPending
								if v43 != 0 {
									return
								} else {
									F_parser_errposition(m, l0, l2)
									mBase = m.M
									v45 = m.ExcPending
									if v45 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_check_lateral_ref_ok_3), int32(509), int32(_a_F_check_lateral_ref_ok_4))
										mBase = m.M
										v50 = m.ExcPending
										if v50 != 0 {
											return
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v17
								F_errhint(m, int32(_a_F_check_lateral_ref_ok_5), v9)
								mBase = m.M
								v39 = m.ExcPending
								if v39 != 0 {
									return
								} else {
									F_parser_errposition(m, l0, l2)
									mBase = m.M
									v45 = m.ExcPending
									if v45 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_check_lateral_ref_ok_3), int32(509), int32(_a_F_check_lateral_ref_ok_4))
										mBase = m.M
										v50 = m.ExcPending
										if v50 != 0 {
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
	}
}
func F_check_publications_origin(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v38 int32
	_ = v38
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
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
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v257 int32
	_ = v257
	var v262 int32
	_ = v262
	v8 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(80)
	m.G0 = v11
	*(*int32)(unsafe.Add(mBase, uint32(v11)+60)) = int32(25)
	if base.B2i32(l2 == v8)|base.B2i32(l3 == v8) != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L18
	} else {
		goto L76
	}
L2:
	;
	m.G0 = v11 + int32(80)
	return
L3:
	;
	v23 = l3
	v24 = int32(_a_F_check_publications_origin_0)
	goto L5
L4:
	;
	if v61 != 0 {
		goto L2
	} else {
		goto L17
	}
L5:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	if v27 == v28 {
		v50 = v27
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v61 = int32(0)
	goto L4
L7:
	;
	v52 = int32(1)
	if v50 != 0 {
		v23 = v23 + v52
		v24 = v24 + v52
		goto L5
	} else {
		goto L16
	}
L8:
	;
	if base.Ui32((v27-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v38 = v27 | int32(32)
	goto L11
L10:
	;
	v38 = v27
	goto L11
L11:
	;
	if base.Ui32((v28-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v47 = v28 | int32(32)
	goto L14
L13:
	;
	v47 = v28
	goto L14
L14:
	;
	if v38 == v47 {
		v50 = v38
		goto L7
	} else {
		goto L15
	}
L15:
	;
	v61 = v38 - v47
	goto L4
L16:
	;
	goto L6
L17:
	;
	v63 = v11 - int32(-64)
	F_initStringInfo(m, v63)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	return
L19:
	;
	F_appendStringInfoString(m, v63, int32(_a_F_check_publications_origin_1))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	F_GetPublicationsStr(m, l1, v63, int32(1))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L18
	} else {
		goto L21
	}
L21:
	;
	F_appendStringInfoString(m, v63, int32(_a_F_check_publications_origin_2))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L18
	} else {
		goto L22
	}
L22:
	;
	if int32(0) < l5 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v81 = int32(0)
	goto L26
L24:
	;
	goto L25
L25:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
	v121 = *(*int32)(unsafe.Add(mBase, _c_F_check_publications_origin[0]))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)+60))
	v123 = m.T0[v122].(func(*base.Module, int32, int32, int32, int32) int32)(m, l0, v116, int32(1), v11+int32(60))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L18
	} else {
		goto L33
	}
L26:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l4+v81<<(uint(int32(2))%32))))
	v90 = F_get_rel_namespace(m, v89)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L18
	} else {
		goto L28
	}
L27:
	;
	goto L25
L28:
	;
	v92 = F_get_namespace_name(m, v90)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L18
	} else {
		goto L29
	}
L29:
	;
	v94 = F_get_rel_name(m, v89)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L18
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = v94
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v92
	F_appendStringInfo(m, v11-int32(-64), int32(_a_F_check_publications_origin_3), v11+int32(48))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L18
	} else {
		goto L31
	}
L31:
	;
	v106 = v81 + int32(1)
	if v106 != l5 {
		v81 = v106
		goto L26
	} else {
		goto L32
	}
L32:
	;
	goto L27
L33:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
	F_pfree(m, v125)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L18
	} else {
		goto L34
	}
L34:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v123)))
	if v128 != int32(2) {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v123)+16))
	v133 = F_MakeTupleTableSlot(m, v131, int32(_a_F_check_publications_origin_4))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L18
	} else {
		goto L36
	}
L36:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v123)+12))
	v138 = F_tuplestore_gettupleslot(m, v135, int32(1), int32(0), v133)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L18
	} else {
		goto L38
	}
L37:
	;
	F_ExecDropSingleTupleTableSlot(m, v133)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L18
	} else {
		goto L62
	}
L38:
	;
	if v138 == int32(0) {
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v145 = int32(0)
	goto L40
L40:
	;
	v151 = int32(*(*int16)(unsafe.Add(mBase, uint32(v133)+6)))
	if v151 <= int32(0) {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	if v167 == int32(0) {
		goto L37
	} else {
		goto L52
	}
L42:
	;
	F_slot_getsomeattrs_int(m, v133, int32(1))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L18
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v133)+16))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)))
	v159 = F_text_to_cstring(m, v158)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L18
	} else {
		goto L46
	}
L45:
	;
	goto L44
L46:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v133)+8))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v161)+12))
	m.T0[v162].(func(*base.Module, int32))(m, v133)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L18
	} else {
		goto L47
	}
L47:
	;
	v165 = F_makeString(m, v159)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L18
	} else {
		goto L48
	}
L48:
	;
	v167 = F_list_append_unique(m, v145, v165)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L18
	} else {
		goto L49
	}
L49:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v123)+12))
	v172 = F_tuplestore_gettupleslot(m, v169, int32(1), int32(0), v133)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L18
	} else {
		goto L50
	}
L50:
	;
	if v172 != 0 {
		v145 = v167
		goto L40
	} else {
		goto L51
	}
L51:
	;
	goto L41
L52:
	;
	v176 = F_makeStringInfo(m)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L18
	} else {
		goto L53
	}
L53:
	;
	F_GetPublicationsStr(m, v167, v176, int32(0))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L18
	} else {
		goto L54
	}
L54:
	;
	v183 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L18
	} else {
		goto L55
	}
L55:
	;
	if v183 == int32(0) {
		goto L37
	} else {
		goto L56
	}
L56:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L18
	} else {
		goto L57
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = l6
	F_errmsg(m, int32(_a_F_check_publications_origin_5), v11+int32(16))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L18
	} else {
		goto L58
	}
L58:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v176)))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v197
	F_errdetail_plural(m, int32(_a_F_check_publications_origin_6), int32(_a_F_check_publications_origin_7), v196, v11)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L18
	} else {
		goto L59
	}
L59:
	;
	F_errhint(m, int32(_a_F_check_publications_origin_8), int32(0))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L18
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(_a_F_check_publications_origin_9), int32(2193), int32(_a_F_check_publications_origin_10))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L18
	} else {
		goto L61
	}
L61:
	;
	goto L37
L62:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v123)+8))
	if v222 != 0 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	F_pfree(m, v222)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L18
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v123)+12))
	if v225 != 0 {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	goto L65
L67:
	;
	F_tuplestore_end(m, v225)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L18
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v123)+16))
	if v228 != 0 {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	goto L69
L71:
	;
	F_FreeTupleDesc(m, v228)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L18
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	F_pfree(m, v123)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L18
	} else {
		goto L75
	}
L74:
	;
	goto L73
L75:
	;
	goto L2
L76:
	;
	F_errcode(m, int32(100663808))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L18
	} else {
		goto L77
	}
L77:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v123)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v251
	F_errmsg(m, int32(_a_F_check_publications_origin_11), v11+int32(32))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L18
	} else {
		goto L78
	}
L78:
	;
	F_errfinish(m, int32(_a_F_check_publications_origin_9), int32(2153), int32(_a_F_check_publications_origin_10))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L18
	} else {
		goto L79
	}
L79:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_check_stack_depth(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_check_stack_depth[0]))
	if v8 == int32(0) {
		m.G0 = v5 + int32(16)
		return
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, _c_F_check_stack_depth[1]))
		v15 = v8 - (v5 + int32(15))
		v17 = v15 >> (uint(int32(31)) % 32)
		if v15^v17-v17 <= v12 {
			m.G0 = v5 + int32(16)
			return
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return
			} else {
				F_errcode(m, int32(16777477))
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return
				} else {
					F_errmsg(m, int32(_a_F_check_stack_depth_0), int32(0))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return
					} else {
						v33 = *(*int32)(unsafe.Add(mBase, _c_F_check_stack_depth[2]))
						*(*int32)(unsafe.Add(mBase, uint32(v5))) = v33
						F_errhint(m, int32(_a_F_check_stack_depth_1), v5)
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_check_stack_depth_2), int32(104), int32(_a_F_check_stack_depth_3))
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
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
func F_check_synchronized_standby_slots(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v300 int32
	_ = v300
	v4 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
	if v14 == v4 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v11 + int32(48)
	return v300
L2:
	;
	v300 = int32(1)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v18 = F_pstrdup(m, v13)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	F_pfree(m, v18)
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L6
	} else {
		goto L72
	}
L6:
	;
	return int32(0)
L7:
	;
	v25 = F_SplitIdentifierString(m, v18, int32(44), v11+int32(32))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	if v25 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v30 = *(*int32)(unsafe.Add(mBase, _c_F_check_synchronized_standby_slots[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_check_synchronized_standby_slots[1])) = v30
	goto L12
L10:
	;
	goto L11
L11:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v11)+32))
	if v39 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v36 = F_format_elog_string(m, int32(_a_F_check_synchronized_standby_slots_0), int32(0))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L6
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_synchronized_standby_slots[2])) = v36
	v287 = v4
	goto L5
L14:
	;
	v287 = int32(1)
	goto L5
L15:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	if int32(0) < v42 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v46 = int32(0)
	goto L19
L17:
	;
	v112 = v39
	goto L18
L18:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v112)+4))
	if v117 <= int32(0) {
		goto L34
	} else {
		goto L35
	}
L19:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v54+v46<<(uint(int32(2))%32))))
	v59 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v59
	v69 = F_ReplicationSlotValidateNameInternal(m, v58, v11+int32(44), v11+int32(40), v11+int32(36))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L6
	} else {
		goto L21
	}
L20:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v11)+32))
	if v106 == int32(0) {
		goto L14
	} else {
		goto L32
	}
L21:
	;
	if v69 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v11)+44))
	*(*int32)(unsafe.Add(mBase, _c_F_check_synchronized_standby_slots[3])) = v73
	goto L25
L23:
	;
	goto L24
L24:
	;
	v103 = v46 + int32(1)
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	if v103 < v104 {
		v46 = v103
		goto L19
	} else {
		goto L31
	}
L25:
	;
	v77 = *(*int32)(unsafe.Add(mBase, _c_F_check_synchronized_standby_slots[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_check_synchronized_standby_slots[1])) = v77
	goto L26
L26:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v11)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v80
	v86 = F_format_elog_string(m, int32(_a_F_check_synchronized_standby_slots_1), v11+int32(16))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L6
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_synchronized_standby_slots[2])) = v86
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v11)+36))
	if v89 == int32(0) {
		v287 = v4
		goto L5
	} else {
		goto L28
	}
L28:
	;
	v93 = *(*int32)(unsafe.Add(mBase, _c_F_check_synchronized_standby_slots[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_check_synchronized_standby_slots[1])) = v93
	goto L29
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v89
	v99 = F_format_elog_string(m, int32(_a_F_check_synchronized_standby_slots_1), v11)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L6
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_synchronized_standby_slots[4])) = v99
	v287 = v4
	goto L5
L31:
	;
	goto L20
L32:
	;
	v112 = v106
	goto L18
L33:
	;
	v151 = F_guc_malloc(m, v146)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L6
	} else {
		goto L40
	}
L34:
	;
	v146 = int32(4)
	goto L33
L35:
	;
	goto L36
L36:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v112)+12))
	v124 = int32(0)
	v127 = int32(4)
	goto L37
L37:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v121+v124<<(uint(int32(2))%32))))
	v136 = F_strlen(m, v135)
	mBase = m.M
	v138 = int32(1)
	v139 = v136 + v127 + v138
	v141 = v124 + v138
	if v141 != v117 {
		v124 = v141
		v127 = v139
		goto L37
	} else {
		goto L39
	}
L38:
	;
	v146 = v139
	goto L33
L39:
	;
	goto L38
L40:
	;
	if v151 == int32(0) {
		v300 = v4
		goto L1
	} else {
		goto L41
	}
L41:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v11)+32))
	if v155 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v155)+4))
	v158 = v156
	goto L44
L43:
	;
	v158 = int32(0)
	goto L44
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v151))) = v158
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v11)+32))
	if v160 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v151
	goto L14
L46:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v160)+4))
	if v163 <= int32(0) {
		goto L45
	} else {
		goto L47
	}
L47:
	;
	v169 = int32(0)
	v173 = v151 + int32(4)
	goto L48
L48:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v160)+12))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v177+v169<<(uint(int32(2))%32))))
	if (v181^v173)&int32(3) != 0 {
		goto L53
	} else {
		goto L54
	}
L49:
	;
	goto L45
L50:
	;
	v256 = F_strlen(m, v181)
	mBase = m.M
	v258 = int32(1)
	v261 = v169 + v258
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v160)+4))
	if v261 < v262 {
		v169 = v261
		v173 = v173 + v256 + v258
		goto L48
	} else {
		goto L71
	}
L51:
	;
	goto L50
L52:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v236))) = uint8(v235)
	if v235&int32(255) == int32(0) {
		goto L51
	} else {
		goto L67
	}
L53:
	;
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v181))))
	v234 = v181
	v235 = v187
	v236 = v173
	goto L52
L54:
	;
	goto L55
L55:
	;
	if v181&int32(3) != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v191 = v181
	v193 = v173
	goto L59
L57:
	;
	v205 = v181
	v207 = v173
	goto L58
L58:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v205)))
	v212 = int32(-2139062144)
	if (int32(16843008)-v209|v209)&v212 != v212 {
		v234 = v205
		v235 = v209
		v236 = v207
		goto L52
	} else {
		goto L63
	}
L59:
	;
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v191))))
	*(*uint8)(unsafe.Add(mBase, uint32(v193))) = uint8(v194)
	if v194 == int32(0) {
		goto L51
	} else {
		goto L61
	}
L60:
	;
	v205 = v201
	v207 = v199
	goto L58
L61:
	;
	v198 = int32(1)
	v199 = v193 + v198
	v201 = v191 + v198
	if v201&int32(3) != 0 {
		v191 = v201
		v193 = v199
		goto L59
	} else {
		goto L62
	}
L62:
	;
	goto L60
L63:
	;
	v217 = v205
	v218 = v209
	v219 = v207
	goto L64
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v219))) = v218
	v221 = int32(4)
	v222 = v219 + v221
	v224 = v217 + v221
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v217)+4))
	v229 = int32(-2139062144)
	if (int32(16843008)-v226|v226)&v229 == v229 {
		v217 = v224
		v218 = v226
		v219 = v222
		goto L64
	} else {
		goto L66
	}
L65:
	;
	v234 = v224
	v235 = v226
	v236 = v222
	goto L52
L66:
	;
	goto L65
L67:
	;
	v243 = v234
	v245 = v236
	goto L68
L68:
	;
	v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v245)+1)) = uint8(v246)
	v248 = int32(1)
	if v246 != 0 {
		v243 = v243 + v248
		v245 = v245 + v248
		goto L68
	} else {
		goto L70
	}
L69:
	;
	goto L51
L70:
	;
	goto L69
L71:
	;
	goto L49
L72:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v11)+32))
	F_list_free(m, v292)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L6
	} else {
		goto L73
	}
L73:
	;
	v300 = v287
	goto L1
}
func F_check_synchronous_standby_names(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v77 int64
	_ = v77
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v166 int32
	_ = v166
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v247 int32
	_ = v247
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v433 int32
	_ = v433
	var v441 int32
	_ = v441
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v454 int32
	_ = v454
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v469 int32
	_ = v469
	var v474 int32
	_ = v474
	var v479 int32
	_ = v479
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v496 int32
	_ = v496
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v551 int32
	_ = v551
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v577 int32
	_ = v577
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v603 int32
	_ = v603
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v613 int32
	_ = v613
	var v619 int32
	_ = v619
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v645 int32
	_ = v645
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v667 int32
	_ = v667
	var v677 int32
	_ = v677
	var v697 int32
	_ = v697
	var v703 int32
	_ = v703
	var v718 int32
	_ = v718
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v743 int32
	_ = v743
	var v746 int32
	_ = v746
	var v750 int32
	_ = v750
	var v768 int32
	_ = v768
	var v770 int32
	_ = v770
	var v777 int32
	_ = v777
	var v784 int32
	_ = v784
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v807 int32
	_ = v807
	var v809 int32
	_ = v809
	var v812 int32
	_ = v812
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v824 int32
	_ = v824
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v856 int32
	_ = v856
	var v857 int32
	_ = v857
	var v860 int32
	_ = v860
	var v866 int32
	_ = v866
	var v868 int32
	_ = v868
	var v869 int32
	_ = v869
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v880 int32
	_ = v880
	var v882 int32
	_ = v882
	var v883 int32
	_ = v883
	var v885 int32
	_ = v885
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v890 int32
	_ = v890
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v907 int32
	_ = v907
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v913 int32
	_ = v913
	var v920 int32
	_ = v920
	var v924 int32
	_ = v924
	var v942 int32
	_ = v942
	var v943 int32
	_ = v943
	var v945 int32
	_ = v945
	var v953 int32
	_ = v953
	var v954 int32
	_ = v954
	var v958 int32
	_ = v958
	var v959 int32
	_ = v959
	var v964 int32
	_ = v964
	var v970 int32
	_ = v970
	var v975 int32
	_ = v975
	var v976 int32
	_ = v976
	var v996 int32
	_ = v996
	var v1004 int32
	_ = v1004
	var v1005 int32
	_ = v1005
	var v1007 int32
	_ = v1007
	var v1008 int32
	_ = v1008
	var v1012 int32
	_ = v1012
	var v1013 int32
	_ = v1013
	var v1018 int32
	_ = v1018
	var v1028 int32
	_ = v1028
	var v1048 int32
	_ = v1048
	var v1052 int32
	_ = v1052
	var v1054 int32
	_ = v1054
	var v1060 int32
	_ = v1060
	var v1088 int32
	_ = v1088
	var v1092 int32
	_ = v1092
	var v1094 int32
	_ = v1094
	var v1099 int32
	_ = v1099
	var v1105 int32
	_ = v1105
	var v1127 int32
	_ = v1127
	var v1131 int32
	_ = v1131
	var v1132 int32
	_ = v1132
	var v1137 int32
	_ = v1137
	var v1139 int32
	_ = v1139
	var v1144 int32
	_ = v1144
	var v1152 int32
	_ = v1152
	var v1182 int32
	_ = v1182
	var v1183 int32
	_ = v1183
	var v1189 int32
	_ = v1189
	var v1193 int32
	_ = v1193
	var v1194 int32
	_ = v1194
	var v1202 int32
	_ = v1202
	var v1205 int32
	_ = v1205
	var v1206 int32
	_ = v1206
	var v1219 int32
	_ = v1219
	var v1221 int32
	_ = v1221
	var v1224 int32
	_ = v1224
	var v1241 int32
	_ = v1241
	var v1243 int32
	_ = v1243
	var v1245 int32
	_ = v1245
	var v1247 int32
	_ = v1247
	var v1249 int32
	_ = v1249
	var v1251 int32
	_ = v1251
	var v1253 int32
	_ = v1253
	var v1255 int32
	_ = v1255
	var v1257 int32
	_ = v1257
	var v1258 int32
	_ = v1258
	var v1260 int32
	_ = v1260
	var v1262 int32
	_ = v1262
	var v1270 int32
	_ = v1270
	var v1272 int32
	_ = v1272
	var v1297 int32
	_ = v1297
	var v1299 int32
	_ = v1299
	var v1302 int32
	_ = v1302
	var v1319 int32
	_ = v1319
	var v1321 int32
	_ = v1321
	var v1326 int32
	_ = v1326
	var v1354 int32
	_ = v1354
	var v1355 int32
	_ = v1355
	var v1359 int32
	_ = v1359
	var v1364 int32
	_ = v1364
	var v1369 int32
	_ = v1369
	var v1370 int32
	_ = v1370
	var v1386 int32
	_ = v1386
	var v1389 int32
	_ = v1389
	var v1395 int32
	_ = v1395
	var v1396 int32
	_ = v1396
	var v1397 int32
	_ = v1397
	var v1400 int32
	_ = v1400
	var v1405 int32
	_ = v1405
	var v1409 int32
	_ = v1409
	var v1411 int32
	_ = v1411
	var v1427 int32
	_ = v1427
	var v1432 int32
	_ = v1432
	var v1434 int32
	_ = v1434
	var v1438 int32
	_ = v1438
	var v1440 int32
	_ = v1440
	var v1443 int32
	_ = v1443
	var v1444 int32
	_ = v1444
	var v1445 int32
	_ = v1445
	var v1446 int32
	_ = v1446
	var v1447 int32
	_ = v1447
	var v1448 int32
	_ = v1448
	var v1453 int32
	_ = v1453
	var v1455 int32
	_ = v1455
	var v1456 int32
	_ = v1456
	var v1460 int32
	_ = v1460
	var v1461 int32
	_ = v1461
	var v1462 int32
	_ = v1462
	var v1469 int32
	_ = v1469
	var v1471 int32
	_ = v1471
	var v1491 int32
	_ = v1491
	var v1494 int32
	_ = v1494
	var v1496 int32
	_ = v1496
	var v1503 int32
	_ = v1503
	var v1523 int32
	_ = v1523
	var v1524 int32
	_ = v1524
	var v1525 int32
	_ = v1525
	var v1527 int32
	_ = v1527
	var v1528 int32
	_ = v1528
	var v1529 int32
	_ = v1529
	var v1533 int32
	_ = v1533
	var v1534 int32
	_ = v1534
	var v1539 int32
	_ = v1539
	var v1541 int32
	_ = v1541
	var v1542 int32
	_ = v1542
	var v1543 int32
	_ = v1543
	var v1552 int32
	_ = v1552
	var v1553 int32
	_ = v1553
	var v1554 int32
	_ = v1554
	var v1558 int32
	_ = v1558
	var v1559 int32
	_ = v1559
	var v1562 int32
	_ = v1562
	var v1566 int32
	_ = v1566
	var v1571 int32
	_ = v1571
	var v1572 int32
	_ = v1572
	var v1576 int32
	_ = v1576
	var v1577 int32
	_ = v1577
	var v1580 int32
	_ = v1580
	var v1581 int32
	_ = v1581
	var v1582 int32
	_ = v1582
	var v1589 int32
	_ = v1589
	var v1610 int32
	_ = v1610
	var v1611 int32
	_ = v1611
	var v1620 int32
	_ = v1620
	var v1626 int32
	_ = v1626
	var v1627 int32
	_ = v1627
	var v1631 int32
	_ = v1631
	var v1632 int32
	_ = v1632
	var v1636 int32
	_ = v1636
	var v1637 int32
	_ = v1637
	var v1640 int32
	_ = v1640
	var v1641 int32
	_ = v1641
	var v1642 int32
	_ = v1642
	var v1645 int32
	_ = v1645
	var v1651 int32
	_ = v1651
	var v1658 int32
	_ = v1658
	var v1661 int32
	_ = v1661
	var v1668 int32
	_ = v1668
	var v1688 int32
	_ = v1688
	var v1689 int32
	_ = v1689
	var v1699 int32
	_ = v1699
	var v1719 int32
	_ = v1719
	var v1720 int32
	_ = v1720
	var v1723 int32
	_ = v1723
	var v1726 int32
	_ = v1726
	var v1727 int32
	_ = v1727
	var v1728 int32
	_ = v1728
	var v1732 int32
	_ = v1732
	var v1734 int32
	_ = v1734
	var v1735 int32
	_ = v1735
	var v1736 int32
	_ = v1736
	var v1737 int32
	_ = v1737
	var v1738 int32
	_ = v1738
	var v1739 int32
	_ = v1739
	var v1744 int32
	_ = v1744
	var v1746 int32
	_ = v1746
	var v1749 int32
	_ = v1749
	var v1751 int32
	_ = v1751
	var v1752 int32
	_ = v1752
	var v1756 int32
	_ = v1756
	var v1760 int32
	_ = v1760
	var v1762 int32
	_ = v1762
	var v1765 int32
	_ = v1765
	var v1772 int32
	_ = v1772
	var v1774 int32
	_ = v1774
	var v1777 int32
	_ = v1777
	var v1780 int32
	_ = v1780
	var v1781 int32
	_ = v1781
	var v1783 int32
	_ = v1783
	var v1785 int32
	_ = v1785
	var v1786 int32
	_ = v1786
	var v1789 int32
	_ = v1789
	var v1790 int32
	_ = v1790
	var v1792 int32
	_ = v1792
	var v1796 int32
	_ = v1796
	var v1797 int32
	_ = v1797
	var v1801 int32
	_ = v1801
	var v1802 int32
	_ = v1802
	var v1806 int32
	_ = v1806
	var v1814 int32
	_ = v1814
	var v1815 int32
	_ = v1815
	var v1818 int32
	_ = v1818
	var v1819 int32
	_ = v1819
	var v1820 int32
	_ = v1820
	var v1822 int32
	_ = v1822
	var v1823 int32
	_ = v1823
	var v1826 int32
	_ = v1826
	var v1827 int32
	_ = v1827
	var v1829 int32
	_ = v1829
	var v1832 int32
	_ = v1832
	var v1833 int32
	_ = v1833
	var v1834 int32
	_ = v1834
	var v1837 int32
	_ = v1837
	var v1846 int32
	_ = v1846
	var v1847 int32
	_ = v1847
	var v1848 int32
	_ = v1848
	var v1849 int32
	_ = v1849
	var v1850 int32
	_ = v1850
	var v1854 int32
	_ = v1854
	var v1855 int32
	_ = v1855
	var v1859 int32
	_ = v1859
	var v1860 int32
	_ = v1860
	var v1861 int32
	_ = v1861
	var v1862 int32
	_ = v1862
	var v1863 int32
	_ = v1863
	var v1864 int32
	_ = v1864
	var v1865 int32
	_ = v1865
	var v1866 int32
	_ = v1866
	var v1867 int32
	_ = v1867
	var v1868 int32
	_ = v1868
	var v1871 int32
	_ = v1871
	var v1873 int32
	_ = v1873
	var v1874 int32
	_ = v1874
	var v1878 int32
	_ = v1878
	var v1879 int32
	_ = v1879
	var v1885 int32
	_ = v1885
	var v1886 int32
	_ = v1886
	var v1888 int32
	_ = v1888
	var v1889 int32
	_ = v1889
	var v1891 int32
	_ = v1891
	var v1892 int32
	_ = v1892
	var v1894 int32
	_ = v1894
	var v1897 int32
	_ = v1897
	var v1898 int32
	_ = v1898
	var v1900 int32
	_ = v1900
	var v1902 int32
	_ = v1902
	var v1903 int32
	_ = v1903
	var v1907 int32
	_ = v1907
	var v1908 int32
	_ = v1908
	var v1909 int32
	_ = v1909
	var v1913 int32
	_ = v1913
	var v1914 int32
	_ = v1914
	var v1917 int32
	_ = v1917
	var v1918 int32
	_ = v1918
	var v1919 int32
	_ = v1919
	var v1927 int32
	_ = v1927
	var v1928 int32
	_ = v1928
	var v1930 int32
	_ = v1930
	var v1937 int32
	_ = v1937
	var v1944 int32
	_ = v1944
	var v1959 int32
	_ = v1959
	var v1960 int32
	_ = v1960
	var v1962 int32
	_ = v1962
	var v1970 int32
	_ = v1970
	var v1971 int32
	_ = v1971
	var v1975 int32
	_ = v1975
	var v1976 int32
	_ = v1976
	var v1981 int32
	_ = v1981
	var v1987 int32
	_ = v1987
	var v1992 int32
	_ = v1992
	var v1993 int32
	_ = v1993
	var v2013 int32
	_ = v2013
	var v2021 int32
	_ = v2021
	var v2022 int32
	_ = v2022
	var v2024 int32
	_ = v2024
	var v2025 int32
	_ = v2025
	var v2029 int32
	_ = v2029
	var v2030 int32
	_ = v2030
	var v2035 int32
	_ = v2035
	var v2045 int32
	_ = v2045
	var v2065 int32
	_ = v2065
	var v2069 int32
	_ = v2069
	var v2071 int32
	_ = v2071
	var v2073 int32
	_ = v2073
	var v2074 int32
	_ = v2074
	var v2075 int32
	_ = v2075
	var v2081 int32
	_ = v2081
	var v2082 int32
	_ = v2082
	var v2089 int32
	_ = v2089
	var v2102 int32
	_ = v2102
	var v2104 int32
	_ = v2104
	var v2110 int32
	_ = v2110
	var v2114 int32
	_ = v2114
	var v2132 int32
	_ = v2132
	var v2133 int32
	_ = v2133
	var v2135 int32
	_ = v2135
	var v2143 int32
	_ = v2143
	var v2144 int32
	_ = v2144
	var v2148 int32
	_ = v2148
	var v2149 int32
	_ = v2149
	var v2154 int32
	_ = v2154
	var v2160 int32
	_ = v2160
	var v2165 int32
	_ = v2165
	var v2166 int32
	_ = v2166
	var v2186 int32
	_ = v2186
	var v2194 int32
	_ = v2194
	var v2195 int32
	_ = v2195
	var v2197 int32
	_ = v2197
	var v2198 int32
	_ = v2198
	var v2202 int32
	_ = v2202
	var v2203 int32
	_ = v2203
	var v2208 int32
	_ = v2208
	var v2218 int32
	_ = v2218
	var v2238 int32
	_ = v2238
	var v2242 int32
	_ = v2242
	var v2244 int32
	_ = v2244
	var v2248 int32
	_ = v2248
	var v2254 int32
	_ = v2254
	var v2260 int32
	_ = v2260
	var v2284 int32
	_ = v2284
	var v2288 int32
	_ = v2288
	var v2291 int32
	_ = v2291
	var v2292 int32
	_ = v2292
	var v2293 int32
	_ = v2293
	var v2294 int32
	_ = v2294
	var v2295 int32
	_ = v2295
	var v2303 int32
	_ = v2303
	var v2305 int32
	_ = v2305
	var v2306 int32
	_ = v2306
	var v2307 int32
	_ = v2307
	var v2309 int32
	_ = v2309
	var v2310 int32
	_ = v2310
	var v2311 int32
	_ = v2311
	var v2313 int32
	_ = v2313
	var v2314 int32
	_ = v2314
	var v2319 int32
	_ = v2319
	var v2327 int32
	_ = v2327
	var v2328 int32
	_ = v2328
	var v2329 int32
	_ = v2329
	var v2330 int32
	_ = v2330
	var v2333 int32
	_ = v2333
	var v2335 int32
	_ = v2335
	var v2340 int32
	_ = v2340
	var v2341 int32
	_ = v2341
	var v2342 int32
	_ = v2342
	var v2343 int32
	_ = v2343
	var v2344 int32
	_ = v2344
	var v2345 int32
	_ = v2345
	var v2353 int32
	_ = v2353
	var v2355 int32
	_ = v2355
	var v2356 int32
	_ = v2356
	var v2357 int32
	_ = v2357
	var v2359 int32
	_ = v2359
	var v2360 int32
	_ = v2360
	var v2361 int32
	_ = v2361
	var v2363 int32
	_ = v2363
	var v2364 int32
	_ = v2364
	var v2373 int32
	_ = v2373
	var v2374 int32
	_ = v2374
	var v2375 int32
	_ = v2375
	var v2376 int32
	_ = v2376
	var v2378 int32
	_ = v2378
	var v2381 int32
	_ = v2381
	var v2387 int32
	_ = v2387
	var v2391 int32
	_ = v2391
	var v2399 int32
	_ = v2399
	var v2401 int32
	_ = v2401
	var v2403 int32
	_ = v2403
	var v2406 int32
	_ = v2406
	var v2407 int32
	_ = v2407
	var v2413 int32
	_ = v2413
	var v2414 int32
	_ = v2414
	var v2417 int32
	_ = v2417
	var v2418 int32
	_ = v2418
	var v2419 int32
	_ = v2419
	var v2428 int32
	_ = v2428
	var v2431 int32
	_ = v2431
	var v2433 int32
	_ = v2433
	var v2436 int32
	_ = v2436
	var v2441 int32
	_ = v2441
	var v2444 int32
	_ = v2444
	var v2446 int32
	_ = v2446
	var v2447 int32
	_ = v2447
	var v2450 int32
	_ = v2450
	var v2453 int32
	_ = v2453
	var v2455 int32
	_ = v2455
	var v2456 int32
	_ = v2456
	var v2459 int32
	_ = v2459
	var v2462 int32
	_ = v2462
	var v2464 int32
	_ = v2464
	var v2465 int32
	_ = v2465
	var v2468 int32
	_ = v2468
	var v2471 int32
	_ = v2471
	var v2473 int32
	_ = v2473
	var v2474 int32
	_ = v2474
	var v2475 int32
	_ = v2475
	var v2481 int32
	_ = v2481
	var v2482 int32
	_ = v2482
	var v2485 int32
	_ = v2485
	var v2486 int32
	_ = v2486
	var v2487 int32
	_ = v2487
	var v2488 int32
	_ = v2488
	var v2489 int32
	_ = v2489
	var v2490 int32
	_ = v2490
	var v2493 int32
	_ = v2493
	var v2496 int32
	_ = v2496
	var v2497 int32
	_ = v2497
	var v2498 int32
	_ = v2498
	var v2501 int32
	_ = v2501
	var v2504 int32
	_ = v2504
	var v2505 int32
	_ = v2505
	var v2508 int32
	_ = v2508
	var v2514 int32
	_ = v2514
	var v2517 int32
	_ = v2517
	var v2518 int32
	_ = v2518
	var v2519 int32
	_ = v2519
	var v2520 int32
	_ = v2520
	var v2521 int32
	_ = v2521
	var v2522 int32
	_ = v2522
	var v2530 int32
	_ = v2530
	var v2532 int32
	_ = v2532
	var v2533 int32
	_ = v2533
	var v2534 int32
	_ = v2534
	var v2536 int32
	_ = v2536
	var v2538 int32
	_ = v2538
	var v2540 int32
	_ = v2540
	var v2544 int32
	_ = v2544
	var v2550 int32
	_ = v2550
	var v2551 int32
	_ = v2551
	var v2554 int32
	_ = v2554
	var v2556 int32
	_ = v2556
	var v2563 int32
	_ = v2563
	var v2579 int32
	_ = v2579
	var v2583 int32
	_ = v2583
	var v2585 int32
	_ = v2585
	var v2588 int32
	_ = v2588
	var v2593 int32
	_ = v2593
	var v2595 int32
	_ = v2595
	var v2599 int32
	_ = v2599
	var v2604 int32
	_ = v2604
	var v2605 int32
	_ = v2605
	var v2609 int32
	_ = v2609
	var v2610 int32
	_ = v2610
	var v2611 int32
	_ = v2611
	var v2614 int32
	_ = v2614
	var v2618 int32
	_ = v2618
	var v2621 int32
	_ = v2621
	var v2622 int32
	_ = v2622
	var v2628 int32
	_ = v2628
	var v2629 int32
	_ = v2629
	var v2633 int32
	_ = v2633
	var v2634 int32
	_ = v2634
	var v2635 int32
	_ = v2635
	var v2638 int32
	_ = v2638
	var v2639 int32
	_ = v2639
	var v2646 int32
	_ = v2646
	var v2673 int32
	_ = v2673
	var v2697 int32
	_ = v2697
	v4 = int32(0)
	v27 = m.G0
	v29 = v27 - int32(48)
	m.G0 = v29
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v31 == v4 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v2673 + int32(48)
	return v2697
L2:
	;
	v2673 = v2646
	v2697 = int32(1)
	goto L1
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
	v2646 = v29
	goto L2
L4:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
	if v34 == int32(0) {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v37 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+40)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v29)+36)) = v37
	v42 = v29 + int32(44)
	v44 = F_palloc0(m, int32(16))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return int32(0)
L7:
	;
	if v42 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v71 = int32(0)
	base.MemoryFill(m, v50, v71, int32(96))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	*(*int32)(unsafe.Add(mBase, uint32(v74)+60)) = v71
	v77 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v74)+52)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v74)+44)) = v71
	*(*int64)(unsafe.Add(mBase, uint32(v74)+36)) = v77
	*(*int64)(unsafe.Add(mBase, uint32(v74)+4)) = v77
	*(*int64)(unsafe.Add(mBase, uint32(v74)+12)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v74)+20)) = v71
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	*(*int32)(unsafe.Add(mBase, uint32(v89))) = v44
	v91 = F_strlen(m, v31)
	mBase = m.M
	v93 = v91 + int32(2)
	v94 = F_palloc(m, v93)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L6
	} else {
		goto L19
	}
L9:
	;
	v50 = F_palloc(m, int32(96))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L6
	} else {
		goto L12
	}
L10:
	;
	v56 = int32(28)
	goto L11
L11:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_synchronous_standby_names[0])) = v56
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L6
	} else {
		goto L14
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42))) = v50
	if v50 != 0 {
		goto L8
	} else {
		goto L13
	}
L13:
	;
	v56 = int32(48)
	goto L11
L14:
	;
	F_errmsg_internal(m, int32(_a_F_check_synchronous_standby_names_0), int32(0))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L6
	} else {
		goto L15
	}
L15:
	;
	F_errfinish(m, int32(_a_F_check_synchronous_standby_names_1), int32(179), int32(_a_F_check_synchronous_standby_names_2))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L6
	} else {
		goto L16
	}
L16:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L17:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v29)+44))
	v354 = m.G0
	v356 = v354 - int32(1024)
	m.G0 = v356
	*(*int32)(unsafe.Add(mBase, uint32(v356)+1020)) = int32(0)
	v362 = v356 + int32(816)
	v364 = v356 + int32(16)
	v366 = v362
	v367 = l1
	v368 = v29
	v369 = v353
	v370 = int32(-2)
	v378 = v364
	v380 = v356
	v381 = v362
	v382 = int32(200)
	v384 = v29 + int32(36)
	v385 = v4
	v386 = v364
	v388 = v29 + int32(40)
	goto L53
L18:
	;
	F_yy_fatal_error_4(m, int32(_a_F_check_synchronous_standby_names_3))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L6
	} else {
		goto L51
	}
L19:
	;
	if v94 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	if v91 <= int32(0) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	goto L22
L22:
	;
	F_yy_fatal_error_4(m, int32(_a_F_check_synchronous_standby_names_4))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L6
	} else {
		goto L50
	}
L23:
	;
	v247 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v94+v91))) = uint16(v247)
	if base.Ui32(v93) < base.Ui32(int32(2)) {
		v334 = v247
		goto L37
	} else {
		goto L38
	}
L24:
	;
	v99 = v91 & int32(3)
	if base.Ui32(int32(4)) <= base.Ui32(v91) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v111 = v4
	v114 = v4
	goto L28
L26:
	;
	v166 = v4
	goto L27
L27:
	;
	v192 = v166
	v196 = v4
	goto L32
L28:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31+v111))))
	*(*uint8)(unsafe.Add(mBase, uint32(v94+v111))) = uint8(v132)
	v135 = v111 | int32(1)
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+v31))))
	*(*uint8)(unsafe.Add(mBase, uint32(v94+v135))) = uint8(v138)
	v141 = v111 | int32(2)
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141+v31))))
	*(*uint8)(unsafe.Add(mBase, uint32(v94+v141))) = uint8(v144)
	v147 = v111 | int32(3)
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147+v31))))
	*(*uint8)(unsafe.Add(mBase, uint32(v94+v147))) = uint8(v150)
	v152 = int32(4)
	v153 = v111 + v152
	v155 = v114 + v152
	if v155 != v91&int32(2147483644) {
		v111 = v153
		v114 = v155
		goto L28
	} else {
		goto L30
	}
L29:
	;
	if v99 == int32(0) {
		goto L23
	} else {
		goto L31
	}
L30:
	;
	goto L29
L31:
	;
	v166 = v153
	goto L27
L32:
	;
	v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31+v192))))
	*(*uint8)(unsafe.Add(mBase, uint32(v94+v192))) = uint8(v213)
	v215 = int32(1)
	v218 = v196 + v215
	if v218 != v99 {
		v192 = v192 + v215
		v196 = v218
		goto L32
	} else {
		goto L34
	}
L33:
	;
	goto L23
L34:
	;
	goto L33
L35:
	;
	if v334 == int32(0) {
		goto L18
	} else {
		goto L49
	}
L36:
	;
	F_yy_fatal_error_4(m, int32(_a_F_check_synchronous_standby_names_5))
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L6
	} else {
		goto L48
	}
L37:
	;
	goto L35
L38:
	;
	v253 = v93 - int32(2)
	v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94+v253))))
	if v255 != 0 {
		v334 = v247
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94+v93-int32(1)))))
	if v259 != 0 {
		v334 = v247
		goto L37
	} else {
		goto L40
	}
L40:
	;
	v261 = F_palloc(m, int32(48))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L6
	} else {
		goto L41
	}
L41:
	;
	if v261 == int32(0) {
		goto L36
	} else {
		goto L42
	}
L42:
	;
	v265 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v261)+20)) = v265
	*(*int32)(unsafe.Add(mBase, uint32(v261)+8)) = v94
	*(*int32)(unsafe.Add(mBase, uint32(v261)+4)) = v94
	*(*int32)(unsafe.Add(mBase, uint32(v261)+12)) = v253
	*(*int64)(unsafe.Add(mBase, uint32(v261)+40)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v261)+24)) = int64(4294967296)
	*(*int32)(unsafe.Add(mBase, uint32(v261)+16)) = v253
	*(*int32)(unsafe.Add(mBase, uint32(v261))) = v265
	F_syncrep_yyensure_buffer_stack(m, v89)
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L6
	} else {
		goto L43
	}
L43:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v89)+20))
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v279+v280<<(uint(int32(2))%32))))
	if v284 == v261 {
		v334 = v261
		goto L37
	} else {
		goto L44
	}
L44:
	;
	if v284 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v89)+36))
	v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v286))) = uint8(v287)
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v89)+20))
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	v291 = int32(2)
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v289+v290<<(uint(v291)%32))))
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v89)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v294)+8)) = v295
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v89)+20))
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v297+v298<<(uint(v291)%32))))
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v89)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v302)+16)) = v303
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v89)+20))
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	v307 = v305
	v308 = v306
	goto L47
L46:
	;
	v307 = v279
	v308 = v280
	goto L47
L47:
	;
	v309 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v308<<(uint(v309)%32)+v307))) = v261
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v89)+20))
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	v317 = v313 + v314<<(uint(v309)%32)
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v317)))
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v318)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v89)+28)) = v319
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v317)))
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v321)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v89)+36)) = v322
	*(*int32)(unsafe.Add(mBase, uint32(v89)+80)) = v322
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v317)))
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v325)))
	*(*int32)(unsafe.Add(mBase, uint32(v89)+4)) = v326
	v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v322))))
	*(*uint8)(unsafe.Add(mBase, uint32(v89)+24)) = uint8(v328)
	*(*int32)(unsafe.Add(mBase, uint32(v89)+48)) = int32(1)
	v334 = v261
	goto L37
L48:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v334)+20)) = int32(1)
	goto L17
L50:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L51:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L52:
	;
	if v2563+int32(816) != v2556 {
		goto L398
	} else {
		goto L399
	}
L53:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v381))) = uint8(v385)
	if base.Ui32(v366+v382-int32(1)) <= base.Ui32(v381) {
		goto L59
	} else {
		goto L60
	}
L54:
	;
	v2550 = v367
	v2551 = v368
	v2554 = int32(1)
	v2556 = v411
	v2563 = v380
	goto L52
L55:
	;
	goto L54
L56:
	;
	v366 = v2518
	v367 = v2519
	v368 = v2520
	v369 = v2521
	v370 = v2522
	v378 = v2530
	v380 = v2532
	v381 = v2533 + int32(1)
	v382 = v2534
	v384 = v2536
	v385 = base.I32_extend8_s(v2544)
	v386 = v2538
	v388 = v2540
	goto L53
L57:
	;
	v2428 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2360)+uint32(_c_F_check_synchronous_standby_names[1]))))
	v2431 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2428)+uint32(_c_F_check_synchronous_standby_names[2]))))
	v2433 = int32(2)
	v2436 = *(*int32)(unsafe.Add(mBase, uint32(v2353+(int32(1)-v2431)<<(uint(v2433)%32))))
	switch v2428&int32(255) - v2433 {
	case 0:
		goto L386
	case 1:
		goto L385
	case 2:
		goto L384
	case 3:
		goto L383
	case 4:
		goto L382
	case 5:
		goto L381
	case 6:
		goto L380
	case 7, 8:
		goto L379
	default:
		v2490 = v2436
		goto L378
	}
L58:
	;
	v2399 = m.G0
	v2401 = v2399 - int32(32)
	m.G0 = v2401
	v2403 = *(*int32)(unsafe.Add(mBase, uint32(v2391)))
	if v2403 == int32(0) {
		goto L369
	} else {
		goto L370
	}
L59:
	;
	v397 = int32(2)
	v398 = int32(_a_F_check_synchronous_standby_names_6)
	if int32(_a_F_check_synchronous_standby_names_7) < v382 {
		v2373 = v366
		v2374 = v367
		v2375 = v368
		v2376 = v369
		v2378 = v397
		v2381 = v398
		v2387 = v380
		v2391 = v384
		goto L58
	} else {
		goto L62
	}
L60:
	;
	v441 = v366
	v446 = v378
	v447 = v381
	v448 = v382
	v449 = v386
	goto L61
L61:
	;
	if v385 == int32(12) {
		goto L79
	} else {
		goto L80
	}
L62:
	;
	v401 = int32(_a_F_check_synchronous_standby_names_8)
	v403 = v382 << (uint(int32(1)) % 32)
	if v401 <= v403 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v406 = v401
	goto L65
L64:
	;
	v406 = v403
	goto L65
L65:
	;
	v411 = F_palloc(m, v406*int32(5)+int32(3))
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L6
	} else {
		goto L66
	}
L66:
	;
	if v411 == int32(0) {
		v2373 = v366
		v2374 = v367
		v2375 = v368
		v2376 = v369
		v2378 = v397
		v2381 = v398
		v2387 = v380
		v2391 = v384
		goto L58
	} else {
		goto L67
	}
L67:
	;
	v415 = v381 - v366
	v417 = v415 + int32(1)
	if v417 != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	base.MemoryCopy(m, v411, v366, v417)
	goto L70
L69:
	;
	goto L70
L70:
	;
	v422 = base.I32_div_s(v406+int32(3), int32(4))
	v423 = int32(2)
	v425 = v411 + v422<<(uint(v423)%32)
	v427 = v417 << (uint(v423) % 32)
	if v427 != 0 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	base.MemoryCopy(m, v425, v386, v427)
	goto L73
L72:
	;
	goto L73
L73:
	;
	if v380+int32(816) != v366 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	F_pfree(m, v366)
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L6
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	if v406-int32(1) <= v415 {
		goto L55
	} else {
		goto L78
	}
L77:
	;
	goto L76
L78:
	;
	v441 = v411
	v446 = v425 + v427 - int32(4)
	v447 = v411 + v415
	v448 = v406
	v449 = v425
	goto L61
L79:
	;
	v2550 = v367
	v2551 = v368
	v2554 = int32(0)
	v2556 = v441
	v2563 = v380
	goto L52
L80:
	;
	goto L81
L81:
	;
	v454 = int32(1) << (uint(v385) % 32)
	if v454&int32(13390146) != 0 {
		v2341 = v441
		v2342 = v367
		v2343 = v368
		v2344 = v369
		v2345 = v370
		v2353 = v446
		v2355 = v380
		v2356 = v447
		v2357 = v448
		v2359 = v384
		v2360 = v385
		v2361 = v449
		v2363 = v388
		v2364 = v454
		goto L82
	} else {
		goto L83
	}
L82:
	;
	if v2364&int32(_a_F_check_synchronous_standby_names_9) == int32(0) {
		goto L57
	} else {
		goto L368
	}
L83:
	;
	v459 = int32(*(*int8)(unsafe.Add(mBase, uint32(v385)+uint32(_c_F_check_synchronous_standby_names[3]))))
	if v370 == int32(-2) {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	v2330 = v459 + v2329
	if base.Ui32(int32(22)) < base.Ui32(v2330) {
		v2341 = v2291
		v2342 = v2292
		v2343 = v2293
		v2344 = v2294
		v2345 = v2328
		v2353 = v2303
		v2355 = v2305
		v2356 = v2306
		v2357 = v2307
		v2359 = v2309
		v2360 = v2310
		v2361 = v2311
		v2363 = v2313
		v2364 = v2314
		goto L82
	} else {
		goto L366
	}
L85:
	;
	v462 = m.G0
	v464 = v462 - int32(32)
	m.G0 = v464
	*(*int32)(unsafe.Add(mBase, uint32(v369)+92)) = v380 + int32(1020)
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v369)+40))
	if v469 == int32(0) {
		goto L89
	} else {
		goto L90
	}
L86:
	;
	v2291 = v441
	v2292 = v367
	v2293 = v368
	v2294 = v369
	v2295 = v370
	v2303 = v446
	v2305 = v380
	v2306 = v447
	v2307 = v448
	v2309 = v384
	v2310 = v385
	v2311 = v449
	v2313 = v388
	v2314 = v454
	goto L87
L87:
	;
	if v2295 <= int32(0) {
		goto L359
	} else {
		goto L360
	}
L88:
	;
	v2291 = v569
	v2292 = v570
	v2293 = v571
	v2294 = v572
	v2295 = v1651
	v2303 = v581
	v2305 = v583
	v2306 = v584
	v2307 = v585
	v2309 = v587
	v2310 = v588
	v2311 = v589
	v2313 = v591
	v2314 = v592
	goto L87
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v369)+40)) = int32(1)
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v369)+44))
	if v474 == int32(0) {
		goto L92
	} else {
		goto L93
	}
L90:
	;
	goto L91
L91:
	;
	v539 = v441
	v540 = v367
	v541 = v368
	v542 = v369
	v551 = v446
	v553 = v380
	v554 = v447
	v555 = v448
	v556 = v464
	v557 = v384
	v558 = v385
	v559 = v449
	v561 = v388
	v562 = v454
	goto L108
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v369)+44)) = int32(1)
	goto L94
L93:
	;
	goto L94
L94:
	;
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v369)+4))
	if v479 == int32(0) {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v483 = *(*int32)(unsafe.Add(mBase, _c_F_check_synchronous_standby_names[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v369)+4)) = v483
	goto L97
L96:
	;
	goto L97
L97:
	;
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v369)+8))
	if v485 == int32(0) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v489 = *(*int32)(unsafe.Add(mBase, _c_F_check_synchronous_standby_names[5]))
	*(*int32)(unsafe.Add(mBase, uint32(v369)+8)) = v489
	goto L100
L99:
	;
	goto L100
L100:
	;
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v369)+20))
	if v491 != 0 {
		goto L102
	} else {
		goto L103
	}
L101:
	;
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v519)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v369)+28)) = v520
	v524 = v517 + v516<<(uint(int32(2))%32)
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v524)))
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v525)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v369)+80)) = v526
	*(*int32)(unsafe.Add(mBase, uint32(v369)+36)) = v526
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v524)))
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v529)))
	*(*int32)(unsafe.Add(mBase, uint32(v369)+4)) = v530
	v532 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v526))))
	*(*uint8)(unsafe.Add(mBase, uint32(v369)+24)) = uint8(v532)
	goto L91
L102:
	;
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v369)+12))
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v491+v492<<(uint(int32(2))%32))))
	if v496 != 0 {
		v516 = v492
		v517 = v491
		v519 = v496
		goto L101
	} else {
		goto L105
	}
L103:
	;
	goto L104
L104:
	;
	F_syncrep_yyensure_buffer_stack(m, v369)
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L6
	} else {
		goto L106
	}
L105:
	;
	goto L104
L106:
	;
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v369)+4))
	v502 = F_syncrep_yy_create_buffer(m, v501, v369)
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L6
	} else {
		goto L107
	}
L107:
	;
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v369)+20))
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v369)+12))
	v506 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v504+v505<<(uint(v506)%32)))) = v502
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v369)+20))
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v369)+12))
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v510+v511<<(uint(v506)%32))))
	v516 = v511
	v517 = v510
	v519 = v515
	goto L101
L108:
	;
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v542)+36))
	v566 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v542)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v565))) = uint8(v566)
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v542)+44))
	v569 = v539
	v570 = v540
	v571 = v541
	v572 = v542
	v573 = v568
	v574 = v565
	v577 = v565
	v581 = v551
	v583 = v553
	v584 = v554
	v585 = v555
	v586 = v556
	v587 = v557
	v588 = v558
	v589 = v559
	v591 = v561
	v592 = v562
	goto L110
L110:
	;
	v595 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v577))))
	v596 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v595)+uint32(_c_F_check_synchronous_standby_names[6]))))
	if base.Ui32(int32(-26)) <= base.Ui32(v573-int32(31)) {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v572)+68)) = v577
	*(*int32)(unsafe.Add(mBase, uint32(v572)+64)) = v573
	goto L114
L113:
	;
	goto L114
L114:
	;
	v603 = int32(1)
	v607 = int32(*(*int16)(unsafe.Add(mBase, uint32(v573<<(uint(v603)%32))+uint32(_c_F_check_synchronous_standby_names[7]))))
	v608 = v607 + v596
	v613 = int32(*(*int16)(unsafe.Add(mBase, uint32(v608<<(uint(v603)%32))+uint32(_c_F_check_synchronous_standby_names[8]))))
	if v613 != v573 {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v619 = v573
	v624 = v596
	v625 = v596
	goto L118
L116:
	;
	v677 = v608
	goto L117
L117:
	;
	v697 = int32(1)
	v703 = int32(*(*int16)(unsafe.Add(mBase, uint32(v677<<(uint(v697)%32))+uint32(_c_F_check_synchronous_standby_names[9]))))
	if int64(1)<<(uint(base.I64_extend_i32_u(v677))%64)&int64(-32985348833280) == int64(0) {
		v573 = v703
		v577 = v577 + v697
		goto L110
	} else {
		goto L124
	}
L118:
	;
	v645 = int32(*(*int16)(unsafe.Add(mBase, uint32(v619<<(uint(int32(1))%32))+uint32(_c_F_check_synchronous_standby_names[10]))))
	if int64(1)<<(uint(base.I64_extend_i32_u(v619))%64)&int64(2076672024) != int64(0) {
		goto L120
	} else {
		goto L121
	}
L119:
	;
	v677 = v662
	goto L117
L120:
	;
	v653 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v625)+uint32(_c_F_check_synchronous_standby_names[11]))))
	v654 = v653
	goto L122
L121:
	;
	v654 = v624
	goto L122
L122:
	;
	v656 = v654 & int32(255)
	v657 = int32(1)
	v661 = int32(*(*int16)(unsafe.Add(mBase, uint32(v645<<(uint(v657)%32))+uint32(_c_F_check_synchronous_standby_names[7]))))
	v662 = v656 + v661
	v667 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v662<<(uint(v657)%32))+uint32(_c_F_check_synchronous_standby_names[8]))))
	if v667 != v645&int32(_a_F_check_synchronous_standby_names_10) {
		v619 = v645
		v624 = v654
		v625 = v656
		goto L118
	} else {
		goto L123
	}
L123:
	;
	goto L119
L124:
	;
	v718 = v574
	goto L125
L125:
	;
	v737 = *(*int32)(unsafe.Add(mBase, uint32(v572)+64))
	v738 = *(*int32)(unsafe.Add(mBase, uint32(v572)+68))
	v743 = v737
	v746 = v718
	v750 = v738
	goto L127
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v572)+80)) = v746
	*(*int32)(unsafe.Add(mBase, uint32(v572)+32)) = v750 - v746
	v768 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v750))))
	*(*uint8)(unsafe.Add(mBase, uint32(v572)+24)) = uint8(v768)
	v770 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v750))) = uint8(v770)
	*(*int32)(unsafe.Add(mBase, uint32(v572)+36)) = v750
	v777 = int32(*(*int16)(unsafe.Add(mBase, uint32(v743<<(uint(int32(1))%32))+uint32(_c_F_check_synchronous_standby_names[12]))))
	v784 = v777
	goto L129
L129:
	;
	v804 = int32(260)
	switch v784 {
	case 0:
		goto L158
	case 1:
		v539 = v569
		v540 = v570
		v541 = v571
		v542 = v572
		v551 = v581
		v553 = v583
		v554 = v584
		v555 = v585
		v556 = v586
		v557 = v587
		v558 = v588
		v559 = v589
		v561 = v591
		v562 = v592
		goto L108
	case 2:
		goto L145
	case 3:
		goto L144
	case 4:
		goto L157
	case 5:
		goto L156
	case 6:
		goto L155
	case 7:
		goto L154
	case 8:
		goto L152
	case 9:
		goto L151
	case 10:
		goto L150
	case 11:
		goto L143
	case 12:
		goto L142
	case 13:
		goto L141
	case 14:
		v1651 = v804
		goto L140
	case 15:
		goto L149
	case 16:
		goto L147
	case 17:
		goto L148
	case 18:
		goto L153
	default:
		goto L146
	}
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v572)+36)) = v2260
	*(*int32)(unsafe.Add(mBase, uint32(v572)+48)) = int32(0)
	v2284 = *(*int32)(unsafe.Add(mBase, uint32(v572)+44))
	v2288 = base.I32_div_s(v2284-int32(1), int32(2))
	v784 = v2288 + int32(17)
	goto L129
L132:
	;
	F_yy_fatal_error_4(m, int32(_a_F_check_synchronous_standby_names_11))
	mBase = m.M
	v2254 = m.ExcPending
	if v2254 != 0 {
		goto L6
	} else {
		goto L358
	}
L133:
	;
	F_yy_fatal_error_4(m, int32(_a_F_check_synchronous_standby_names_12))
	mBase = m.M
	v2248 = m.ExcPending
	if v2248 != 0 {
		goto L6
	} else {
		goto L357
	}
L134:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L135:
	;
	v2102 = v2082 + v2089
	*(*int32)(unsafe.Add(mBase, uint32(v572)+36)) = v2102
	v2104 = *(*int32)(unsafe.Add(mBase, uint32(v572)+44))
	if base.Ui32(v2102) <= base.Ui32(v2081) {
		v743 = v2104
		v746 = v2081
		v750 = v2102
		goto L127
	} else {
		goto L338
	}
L136:
	;
	v1720 = *(*int32)(unsafe.Add(mBase, uint32(v1719)))
	*(*int32)(unsafe.Add(mBase, uint32(v1720)+16)) = v1699
	v1723 = *(*int32)(unsafe.Add(mBase, uint32(v572)+28))
	if v1723 != 0 {
		v1846 = int32(0)
		goto L282
	} else {
		goto L283
	}
L137:
	;
	v1688 = *(*int32)(unsafe.Add(mBase, uint32(v572)+20))
	v1689 = *(*int32)(unsafe.Add(mBase, uint32(v572)+12))
	v1699 = v1668
	v1719 = v1688 + v1689<<(uint(int32(2))%32)
	goto L136
L138:
	;
	F_yy_fatal_error_4(m, int32(_a_F_check_synchronous_standby_names_13))
	mBase = m.M
	v1661 = m.ExcPending
	if v1661 != 0 {
		goto L6
	} else {
		goto L281
	}
L139:
	;
	F_yy_fatal_error_4(m, int32(_a_F_check_synchronous_standby_names_14))
	mBase = m.M
	v1658 = m.ExcPending
	if v1658 != 0 {
		goto L6
	} else {
		goto L280
	}
L140:
	;
	m.G0 = v586 + int32(32)
	goto L88
L141:
	;
	v1651 = int32(41)
	goto L140
L142:
	;
	v1651 = int32(40)
	goto L140
L143:
	;
	v1651 = int32(44)
	goto L140
L144:
	;
	v1651 = int32(262)
	goto L140
L145:
	;
	v1651 = int32(261)
	goto L140
L146:
	;
	F_yy_fatal_error_4(m, int32(_a_F_check_synchronous_standby_names_15))
	mBase = m.M
	v1645 = m.ExcPending
	if v1645 != 0 {
		goto L6
	} else {
		goto L279
	}
L147:
	;
	v868 = *(*int32)(unsafe.Add(mBase, uint32(v572)+80))
	v869 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v572)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v750))) = uint8(v869)
	v871 = *(*int32)(unsafe.Add(mBase, uint32(v572)+20))
	v872 = *(*int32)(unsafe.Add(mBase, uint32(v572)+12))
	v875 = v871 + v872<<(uint(int32(2))%32)
	v876 = *(*int32)(unsafe.Add(mBase, uint32(v875)))
	v877 = *(*int32)(unsafe.Add(mBase, uint32(v876)+44))
	if v877 == int32(0) {
		goto L171
	} else {
		goto L172
	}
L148:
	;
	v1651 = int32(0)
	goto L140
L149:
	;
	F_yy_fatal_error_4(m, int32(_a_F_check_synchronous_standby_names_16))
	mBase = m.M
	v866 = m.ExcPending
	if v866 != 0 {
		goto L6
	} else {
		goto L170
	}
L150:
	;
	v860 = *(*int32)(unsafe.Add(mBase, uint32(v572)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v860))) = int32(_a_F_check_synchronous_standby_names_17)
	v1651 = int32(258)
	goto L140
L151:
	;
	v854 = *(*int32)(unsafe.Add(mBase, uint32(v572)+80))
	v855 = F_pstrdup(m, v854)
	mBase = m.M
	v856 = m.ExcPending
	if v856 != 0 {
		goto L6
	} else {
		goto L169
	}
L152:
	;
	v848 = *(*int32)(unsafe.Add(mBase, uint32(v572)+80))
	v849 = F_pstrdup(m, v848)
	mBase = m.M
	v850 = m.ExcPending
	if v850 != 0 {
		goto L6
	} else {
		goto L168
	}
L153:
	;
	v830 = *(*int32)(unsafe.Add(mBase, uint32(v587)))
	if v830 != 0 {
		v1651 = v804
		goto L140
	} else {
		goto L162
	}
L154:
	;
	v820 = *(*int32)(unsafe.Add(mBase, uint32(v572)+92))
	v821 = *(*int32)(unsafe.Add(mBase, uint32(v572)))
	v822 = *(*int32)(unsafe.Add(mBase, uint32(v821)))
	*(*int32)(unsafe.Add(mBase, uint32(v820))) = v822
	v824 = *(*int32)(unsafe.Add(mBase, uint32(v572)))
	*(*int32)(unsafe.Add(mBase, uint32(v824))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v572)+44)) = int32(1)
	v1651 = int32(258)
	goto L140
L155:
	;
	v816 = *(*int32)(unsafe.Add(mBase, uint32(v572)))
	v817 = *(*int32)(unsafe.Add(mBase, uint32(v572)+80))
	F_appendStringInfoString(m, v816, v817)
	mBase = m.M
	v819 = m.ExcPending
	if v819 != 0 {
		goto L6
	} else {
		goto L161
	}
L156:
	;
	v812 = *(*int32)(unsafe.Add(mBase, uint32(v572)))
	F_appendStringInfoChar(m, v812, int32(34))
	mBase = m.M
	v815 = m.ExcPending
	if v815 != 0 {
		goto L6
	} else {
		goto L160
	}
L157:
	;
	v807 = *(*int32)(unsafe.Add(mBase, uint32(v572)))
	F_initStringInfo(m, v807)
	mBase = m.M
	v809 = m.ExcPending
	if v809 != 0 {
		goto L6
	} else {
		goto L159
	}
L158:
	;
	v805 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v572)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v750))) = uint8(v805)
	v718 = v746
	goto L125
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v572)+44)) = int32(3)
	v539 = v569
	v540 = v570
	v541 = v571
	v542 = v572
	v551 = v581
	v553 = v583
	v554 = v584
	v555 = v585
	v556 = v586
	v557 = v587
	v558 = v588
	v559 = v589
	v561 = v591
	v562 = v592
	goto L108
L160:
	;
	v539 = v569
	v540 = v570
	v541 = v571
	v542 = v572
	v551 = v581
	v553 = v583
	v554 = v584
	v555 = v585
	v556 = v586
	v557 = v587
	v558 = v588
	v559 = v589
	v561 = v591
	v562 = v592
	goto L108
L161:
	;
	v539 = v569
	v540 = v570
	v541 = v571
	v542 = v572
	v551 = v581
	v553 = v583
	v554 = v584
	v555 = v585
	v556 = v586
	v557 = v587
	v558 = v588
	v559 = v589
	v561 = v591
	v562 = v592
	goto L108
L162:
	;
	v831 = *(*int32)(unsafe.Add(mBase, uint32(v572)+80))
	v832 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v831))))
	if v832 != 0 {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v586)+20)) = v831
	*(*int32)(unsafe.Add(mBase, uint32(v586)+16)) = int32(_a_F_check_synchronous_standby_names_18)
	v839 = F_psprintf(m, int32(_a_F_check_synchronous_standby_names_19), v586+int32(16))
	mBase = m.M
	v840 = m.ExcPending
	if v840 != 0 {
		goto L6
	} else {
		goto L166
	}
L164:
	;
	goto L165
L165:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v586))) = int32(_a_F_check_synchronous_standby_names_18)
	v845 = F_psprintf(m, int32(_a_F_check_synchronous_standby_names_20), v586)
	mBase = m.M
	v846 = m.ExcPending
	if v846 != 0 {
		goto L6
	} else {
		goto L167
	}
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v587))) = v839
	v1651 = v804
	goto L140
L167:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v587))) = v845
	v1651 = v804
	goto L140
L168:
	;
	v851 = *(*int32)(unsafe.Add(mBase, uint32(v572)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v851))) = v849
	v1651 = int32(258)
	goto L140
L169:
	;
	v857 = *(*int32)(unsafe.Add(mBase, uint32(v572)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v857))) = v855
	v1651 = int32(259)
	goto L140
L170:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L171:
	;
	v880 = *(*int32)(unsafe.Add(mBase, uint32(v876)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v572)+28)) = v880
	v882 = *(*int32)(unsafe.Add(mBase, uint32(v875)))
	v883 = *(*int32)(unsafe.Add(mBase, uint32(v572)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v882))) = v883
	v885 = *(*int32)(unsafe.Add(mBase, uint32(v572)+20))
	v886 = *(*int32)(unsafe.Add(mBase, uint32(v572)+12))
	v887 = int32(2)
	v890 = *(*int32)(unsafe.Add(mBase, uint32(v885+v886<<(uint(v887)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v890)+44)) = int32(1)
	v893 = *(*int32)(unsafe.Add(mBase, uint32(v572)+20))
	v894 = *(*int32)(unsafe.Add(mBase, uint32(v572)+12))
	v898 = *(*int32)(unsafe.Add(mBase, uint32(v893+v894<<(uint(v887)%32))))
	v899 = v898
	v900 = v893
	v901 = v894
	goto L173
L172:
	;
	v899 = v876
	v900 = v871
	v901 = v872
	goto L173
L173:
	;
	v902 = *(*int32)(unsafe.Add(mBase, uint32(v572)+36))
	v903 = *(*int32)(unsafe.Add(mBase, uint32(v899)+4))
	v904 = *(*int32)(unsafe.Add(mBase, uint32(v572)+28))
	v905 = v903 + v904
	if base.Ui32(v902) <= base.Ui32(v905) {
		goto L174
	} else {
		goto L175
	}
L174:
	;
	v907 = *(*int32)(unsafe.Add(mBase, uint32(v572)+80))
	v910 = v868 ^ int32(-1) + v750
	v911 = v907 + v910
	*(*int32)(unsafe.Add(mBase, uint32(v572)+36)) = v911
	v913 = *(*int32)(unsafe.Add(mBase, uint32(v572)+44))
	if int32(0) < v910 {
		goto L177
	} else {
		goto L178
	}
L175:
	;
	goto L176
L176:
	;
	if base.Ui32(v905+int32(1)) < base.Ui32(v902) {
		goto L139
	} else {
		goto L208
	}
L177:
	;
	v920 = v913
	v924 = v907
	goto L180
L178:
	;
	v1060 = v913
	goto L179
L179:
	;
	if base.Ui32(int32(-26)) <= base.Ui32(v1060-int32(31)) {
		goto L198
	} else {
		goto L199
	}
L180:
	;
	v942 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v924))))
	if v942 != 0 {
		goto L182
	} else {
		goto L183
	}
L181:
	;
	v1060 = v1052
	goto L179
L182:
	;
	v943 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v942)+uint32(_c_F_check_synchronous_standby_names[6]))))
	v945 = v943
	goto L184
L183:
	;
	v945 = int32(1)
	goto L184
L184:
	;
	if base.Ui32(int32(-26)) <= base.Ui32(v920-int32(31)) {
		goto L185
	} else {
		goto L186
	}
L185:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v572)+68)) = v924
	*(*int32)(unsafe.Add(mBase, uint32(v572)+64)) = v920
	goto L187
L186:
	;
	goto L187
L187:
	;
	v953 = v945 & int32(255)
	v954 = int32(1)
	v958 = int32(*(*int16)(unsafe.Add(mBase, uint32(v920<<(uint(v954)%32))+uint32(_c_F_check_synchronous_standby_names[7]))))
	v959 = v953 + v958
	v964 = int32(*(*int16)(unsafe.Add(mBase, uint32(v959<<(uint(v954)%32))+uint32(_c_F_check_synchronous_standby_names[8]))))
	if v964 != v920 {
		goto L188
	} else {
		goto L189
	}
L188:
	;
	v970 = v920
	v975 = v945
	v976 = v953
	goto L191
L189:
	;
	v1028 = v959
	goto L190
L190:
	;
	v1048 = int32(1)
	v1052 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1028<<(uint(v1048)%32))+uint32(_c_F_check_synchronous_standby_names[9]))))
	v1054 = v924 + v1048
	if v1054 != v911 {
		v920 = v1052
		v924 = v1054
		goto L180
	} else {
		goto L197
	}
L191:
	;
	v996 = int32(*(*int16)(unsafe.Add(mBase, uint32(v970<<(uint(int32(1))%32))+uint32(_c_F_check_synchronous_standby_names[10]))))
	if int64(1)<<(uint(base.I64_extend_i32_u(v970))%64)&int64(2076672024) != int64(0) {
		goto L193
	} else {
		goto L194
	}
L192:
	;
	v1028 = v1013
	goto L190
L193:
	;
	v1004 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v976)+uint32(_c_F_check_synchronous_standby_names[11]))))
	v1005 = v1004
	goto L195
L194:
	;
	v1005 = v975
	goto L195
L195:
	;
	v1007 = v1005 & int32(255)
	v1008 = int32(1)
	v1012 = int32(*(*int16)(unsafe.Add(mBase, uint32(v996<<(uint(v1008)%32))+uint32(_c_F_check_synchronous_standby_names[7]))))
	v1013 = v1007 + v1012
	v1018 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1013<<(uint(v1008)%32))+uint32(_c_F_check_synchronous_standby_names[8]))))
	if v1018 != v996&int32(_a_F_check_synchronous_standby_names_10) {
		v970 = v996
		v975 = v1005
		v976 = v1007
		goto L191
	} else {
		goto L196
	}
L196:
	;
	goto L192
L197:
	;
	goto L181
L198:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v572)+68)) = v911
	*(*int32)(unsafe.Add(mBase, uint32(v572)+64)) = v1060
	goto L200
L199:
	;
	goto L200
L200:
	;
	v1088 = int32(1)
	v1092 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1060<<(uint(v1088)%32))+uint32(_c_F_check_synchronous_standby_names[7]))))
	v1094 = v1092 + v1088
	v1099 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1094<<(uint(v1088)%32))+uint32(_c_F_check_synchronous_standby_names[8]))))
	if v1099 != v1060 {
		goto L201
	} else {
		goto L202
	}
L201:
	;
	v1105 = v1060
	goto L204
L202:
	;
	v1152 = v1094
	goto L203
L203:
	;
	if base.B2i32(v1152 == int32(0))|base.B2i32(int64(1)<<(uint(base.I64_extend_i32_u(v1152))%64)&int64(-32985348833280) != int64(0)) != 0 {
		v718 = v907
		goto L125
	} else {
		goto L207
	}
L204:
	;
	v1127 = int32(1)
	v1131 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1105<<(uint(v1127)%32))+uint32(_c_F_check_synchronous_standby_names[10]))))
	v1132 = base.I32_extend16_s(v1131)
	v1137 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1132<<(uint(v1127)%32))+uint32(_c_F_check_synchronous_standby_names[7]))))
	v1139 = v1137 + v1127
	v1144 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1139<<(uint(v1127)%32))+uint32(_c_F_check_synchronous_standby_names[8]))))
	if v1131 != v1144 {
		v1105 = v1132
		goto L204
	} else {
		goto L206
	}
L205:
	;
	v1152 = v1139
	goto L203
L206:
	;
	goto L205
L207:
	;
	v1182 = int32(1)
	v1183 = v911 + v1182
	*(*int32)(unsafe.Add(mBase, uint32(v572)+36)) = v1183
	v1189 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1152<<(uint(v1182)%32))+uint32(_c_F_check_synchronous_standby_names[9]))))
	v573 = v1189
	v574 = v907
	v577 = v1183
	goto L110
L208:
	;
	v1193 = *(*int32)(unsafe.Add(mBase, uint32(v572)+80))
	v1194 = *(*int32)(unsafe.Add(mBase, uint32(v899)+40))
	if v1194 == int32(0) {
		goto L209
	} else {
		goto L210
	}
L209:
	;
	if v902-v1193 != int32(1) {
		v2081 = v1193
		v2082 = v903
		v2089 = v904
		goto L135
	} else {
		goto L212
	}
L210:
	;
	goto L211
L211:
	;
	v1202 = v1193 ^ int32(-1) + v902
	if int32(0) < v1202 {
		goto L213
	} else {
		goto L214
	}
L212:
	;
	v2260 = v1193
	goto L131
L213:
	;
	v1205 = int32(7)
	v1206 = v1202 & v1205
	if base.Ui32(v902-v1193-int32(2)) < base.Ui32(v1205) {
		goto L218
	} else {
		goto L219
	}
L214:
	;
	v1364 = v899
	v1369 = v900
	v1370 = v901
	goto L215
L215:
	;
	v1386 = *(*int32)(unsafe.Add(mBase, uint32(v1364)+44))
	if v1386 == int32(2) {
		goto L228
	} else {
		goto L229
	}
L216:
	;
	v1354 = *(*int32)(unsafe.Add(mBase, uint32(v572)+20))
	v1355 = *(*int32)(unsafe.Add(mBase, uint32(v572)+12))
	v1359 = *(*int32)(unsafe.Add(mBase, uint32(v1354+v1355<<(uint(int32(2))%32))))
	v1364 = v1359
	v1369 = v1354
	v1370 = v1355
	goto L215
L217:
	;
	v1297 = v1270
	v1299 = v1272
	v1302 = int32(0)
	goto L225
L218:
	;
	v1270 = v1193
	v1272 = v903
	goto L217
L219:
	;
	goto L220
L220:
	;
	v1219 = v1193
	v1221 = v903
	v1224 = int32(0)
	goto L221
L221:
	;
	v1241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1219))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1221))) = uint8(v1241)
	v1243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1219)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1221)+1)) = uint8(v1243)
	v1245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1219)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1221)+2)) = uint8(v1245)
	v1247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1219)+3)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1221)+3)) = uint8(v1247)
	v1249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1219)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1221)+4)) = uint8(v1249)
	v1251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1219)+5)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1221)+5)) = uint8(v1251)
	v1253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1219)+6)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1221)+6)) = uint8(v1253)
	v1255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1219)+7)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1221)+7)) = uint8(v1255)
	v1257 = int32(8)
	v1258 = v1221 + v1257
	v1260 = v1219 + v1257
	v1262 = v1224 + v1257
	if v1262 != v1202&int32(2147483640) {
		v1219 = v1260
		v1221 = v1258
		v1224 = v1262
		goto L221
	} else {
		goto L223
	}
L222:
	;
	if v1206 == int32(0) {
		goto L216
	} else {
		goto L224
	}
L223:
	;
	goto L222
L224:
	;
	v1270 = v1260
	v1272 = v1258
	goto L217
L225:
	;
	v1319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1297))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1299))) = uint8(v1319)
	v1321 = int32(1)
	v1326 = v1302 + v1321
	if v1326 != v1206 {
		v1297 = v1297 + v1321
		v1299 = v1299 + v1321
		v1302 = v1326
		goto L225
	} else {
		goto L227
	}
L226:
	;
	goto L216
L227:
	;
	goto L226
L228:
	;
	v1389 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v572)+28)) = v1389
	v1699 = v1389
	v1719 = v1369 + v1370<<(uint(int32(2))%32)
	goto L136
L229:
	;
	goto L230
L230:
	;
	v1395 = *(*int32)(unsafe.Add(mBase, uint32(v1364)+12))
	v1396 = v1193 - v902
	v1397 = v1395 + v1396
	if v1397 <= int32(0) {
		goto L231
	} else {
		goto L232
	}
L231:
	;
	v1400 = *(*int32)(unsafe.Add(mBase, uint32(v572)+36))
	v1405 = v1364
	v1409 = v1400
	v1411 = v1395
	goto L234
L232:
	;
	v1469 = v1364
	v1471 = v1397
	goto L233
L233:
	;
	v1491 = int32(_a_F_check_synchronous_standby_names_21)
	if base.Ui32(v1491) <= base.Ui32(v1471) {
		goto L250
	} else {
		goto L251
	}
L234:
	;
	v1427 = *(*int32)(unsafe.Add(mBase, uint32(v1405)+20))
	if v1427 == int32(0) {
		goto L236
	} else {
		goto L237
	}
L235:
	;
	v1469 = v1460
	v1471 = v1462
	goto L233
L236:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1405)+4)) = int32(0)
	goto L132
L237:
	;
	goto L238
L238:
	;
	v1432 = *(*int32)(unsafe.Add(mBase, uint32(v1405)+4))
	v1434 = v1411 << (uint(int32(1)) % 32)
	if v1434 <= int32(0) {
		goto L239
	} else {
		goto L240
	}
L239:
	;
	v1438 = base.I32_div_s(v1411, int32(8))
	v1440 = v1438 + v1411
	goto L241
L240:
	;
	v1440 = v1434
	goto L241
L241:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1405)+12)) = v1440
	v1443 = v1440 + int32(2)
	if v1432 != 0 {
		goto L243
	} else {
		goto L244
	}
L242:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1405)+4)) = v1448
	if v1448 == int32(0) {
		goto L132
	} else {
		goto L248
	}
L243:
	;
	v1444 = F_repalloc(m, v1432, v1443)
	mBase = m.M
	v1445 = m.ExcPending
	if v1445 != 0 {
		goto L6
	} else {
		goto L246
	}
L244:
	;
	goto L245
L245:
	;
	v1446 = F_palloc(m, v1443)
	mBase = m.M
	v1447 = m.ExcPending
	if v1447 != 0 {
		goto L6
	} else {
		goto L247
	}
L246:
	;
	v1448 = v1444
	goto L242
L247:
	;
	v1448 = v1446
	goto L242
L248:
	;
	v1453 = v1448 + (v1409 - v1432)
	*(*int32)(unsafe.Add(mBase, uint32(v572)+36)) = v1453
	v1455 = *(*int32)(unsafe.Add(mBase, uint32(v572)+20))
	v1456 = *(*int32)(unsafe.Add(mBase, uint32(v572)+12))
	v1460 = *(*int32)(unsafe.Add(mBase, uint32(v1455+v1456<<(uint(int32(2))%32))))
	v1461 = *(*int32)(unsafe.Add(mBase, uint32(v1460)+12))
	v1462 = v1461 + v1396
	if v1462 <= int32(0) {
		v1405 = v1460
		v1409 = v1453
		v1411 = v1461
		goto L234
	} else {
		goto L249
	}
L249:
	;
	goto L235
L250:
	;
	v1494 = v1491
	goto L252
L251:
	;
	v1494 = v1471
	goto L252
L252:
	;
	v1496 = *(*int32)(unsafe.Add(mBase, uint32(v1469)+24))
	if v1496 != 0 {
		goto L253
	} else {
		goto L254
	}
L253:
	;
	v1503 = int32(0)
	goto L257
L254:
	;
	goto L255
L255:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_synchronous_standby_names[0])) = int32(0)
	v1571 = *(*int32)(unsafe.Add(mBase, uint32(v572)+20))
	v1572 = *(*int32)(unsafe.Add(mBase, uint32(v572)+12))
	v1576 = *(*int32)(unsafe.Add(mBase, uint32(v1571+v1572<<(uint(int32(2))%32))))
	v1577 = *(*int32)(unsafe.Add(mBase, uint32(v1576)+4))
	v1580 = *(*int32)(unsafe.Add(mBase, uint32(v572)+4))
	v1581 = F_fread(m, v1577+v1202, int32(1), v1494, v1580)
	mBase = m.M
	v1582 = m.ExcPending
	if v1582 != 0 {
		goto L6
	} else {
		goto L268
	}
L256:
	;
	switch v1527 {
	case 0:
		goto L264
	default:
		v1566 = v1541
		goto L262
	case 11:
		goto L263
	}
L257:
	;
	v1523 = *(*int32)(unsafe.Add(mBase, uint32(v572)+4))
	v1524 = F_do_getc(m, v1523)
	mBase = m.M
	v1525 = m.ExcPending
	if v1525 != 0 {
		goto L6
	} else {
		goto L260
	}
L258:
	;
	v1541 = v1494
	goto L256
L259:
	;
	v1528 = *(*int32)(unsafe.Add(mBase, uint32(v572)+20))
	v1529 = *(*int32)(unsafe.Add(mBase, uint32(v572)+12))
	v1533 = *(*int32)(unsafe.Add(mBase, uint32(v1528+v1529<<(uint(int32(2))%32))))
	v1534 = *(*int32)(unsafe.Add(mBase, uint32(v1533)+4))
	*(*uint8)(unsafe.Add(mBase, uint32(v1534+v1202+v1503))) = uint8(v1524)
	v1539 = v1503 + int32(1)
	if v1539 != v1494 {
		v1503 = v1539
		goto L257
	} else {
		goto L261
	}
L260:
	;
	v1527 = v1524 + int32(1)
	switch v1527 {
	case 0, 11:
		v1541 = v1503
		goto L256
	default:
		goto L259
	}
L261:
	;
	goto L258
L262:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v572)+28)) = v1566
	v1668 = v1566
	goto L137
L263:
	;
	v1553 = *(*int32)(unsafe.Add(mBase, uint32(v572)+20))
	v1554 = *(*int32)(unsafe.Add(mBase, uint32(v572)+12))
	v1558 = *(*int32)(unsafe.Add(mBase, uint32(v1553+v1554<<(uint(int32(2))%32))))
	v1559 = *(*int32)(unsafe.Add(mBase, uint32(v1558)+4))
	v1562 = int32(10)
	*(*uint8)(unsafe.Add(mBase, uint32(v1559+v1202+v1541))) = uint8(v1562)
	v1566 = v1541 + int32(1)
	goto L262
L264:
	;
	v1542 = *(*int32)(unsafe.Add(mBase, uint32(v572)+4))
	v1543 = *(*int32)(unsafe.Add(mBase, uint32(v1542)))
	goto L265
L265:
	;
	if int32(base.Ui32(v1543)>>(uint(int32(5))%32))&int32(1) == int32(0) {
		v1566 = v1541
		goto L262
	} else {
		goto L266
	}
L266:
	;
	F_yy_fatal_error_4(m, int32(_a_F_check_synchronous_standby_names_13))
	mBase = m.M
	v1552 = m.ExcPending
	if v1552 != 0 {
		goto L6
	} else {
		goto L267
	}
L267:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L268:
	;
	v1589 = v1581
	goto L269
L269:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v572)+28)) = v1589
	if v1589 != 0 {
		v1668 = v1589
		goto L137
	} else {
		goto L271
	}
L271:
	;
	v1610 = *(*int32)(unsafe.Add(mBase, uint32(v572)+4))
	v1611 = *(*int32)(unsafe.Add(mBase, uint32(v1610)))
	goto L272
L272:
	;
	if int32(base.Ui32(v1611)>>(uint(int32(5))%32))&int32(1) == int32(0) {
		goto L273
	} else {
		goto L274
	}
L273:
	;
	v1668 = int32(0)
	goto L137
L274:
	;
	goto L275
L275:
	;
	v1620 = *(*int32)(unsafe.Add(mBase, _c_F_check_synchronous_standby_names[0]))
	if v1620 != int32(27) {
		goto L138
	} else {
		goto L276
	}
L276:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_synchronous_standby_names[0])) = int32(0)
	v1626 = *(*int32)(unsafe.Add(mBase, uint32(v572)+4))
	v1627 = *(*int32)(unsafe.Add(mBase, uint32(v1626)))
	*(*int32)(unsafe.Add(mBase, uint32(v1626))) = v1627 & int32(-49)
	goto L277
L277:
	;
	v1631 = *(*int32)(unsafe.Add(mBase, uint32(v572)+20))
	v1632 = *(*int32)(unsafe.Add(mBase, uint32(v572)+12))
	v1636 = *(*int32)(unsafe.Add(mBase, uint32(v1631+v1632<<(uint(int32(2))%32))))
	v1637 = *(*int32)(unsafe.Add(mBase, uint32(v1636)+4))
	v1640 = *(*int32)(unsafe.Add(mBase, uint32(v572)+4))
	v1641 = F_fread(m, v1637+v1202, int32(1), v1494, v1640)
	mBase = m.M
	v1642 = m.ExcPending
	if v1642 != 0 {
		goto L6
	} else {
		goto L278
	}
L278:
	;
	v1589 = v1641
	goto L269
L279:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L280:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L281:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L282:
	;
	v1847 = *(*int32)(unsafe.Add(mBase, uint32(v572)+28))
	v1848 = v1847 + v1202
	v1849 = *(*int32)(unsafe.Add(mBase, uint32(v572)+20))
	v1850 = *(*int32)(unsafe.Add(mBase, uint32(v572)+12))
	v1854 = *(*int32)(unsafe.Add(mBase, uint32(v1849+v1850<<(uint(int32(2))%32))))
	v1855 = *(*int32)(unsafe.Add(mBase, uint32(v1854)+12))
	if v1855 < v1848 {
		goto L306
	} else {
		goto L307
	}
L283:
	;
	if v1202 == int32(0) {
		goto L284
	} else {
		goto L285
	}
L284:
	;
	v1726 = *(*int32)(unsafe.Add(mBase, uint32(v572)+4))
	v1727 = *(*int32)(unsafe.Add(mBase, uint32(v572)+20))
	if v1727 != 0 {
		goto L289
	} else {
		goto L290
	}
L285:
	;
	goto L286
L286:
	;
	v1832 = *(*int32)(unsafe.Add(mBase, uint32(v572)+20))
	v1833 = *(*int32)(unsafe.Add(mBase, uint32(v572)+12))
	v1834 = int32(2)
	v1837 = *(*int32)(unsafe.Add(mBase, uint32(v1832+v1833<<(uint(v1834)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v1837)+44)) = v1834
	v1846 = v1834
	goto L282
L287:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1796)+40)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1796))) = v1726
	v1801 = *(*int32)(unsafe.Add(mBase, uint32(v572)+20))
	if v1801 != 0 {
		goto L302
	} else {
		goto L303
	}
L288:
	;
	v1751 = *(*int32)(unsafe.Add(mBase, _c_F_check_synchronous_standby_names[0]))
	v1752 = *(*int32)(unsafe.Add(mBase, uint32(v572)+12))
	v1756 = *(*int32)(unsafe.Add(mBase, uint32(v1749+v1752<<(uint(int32(2))%32))))
	if v1756 == int32(0) {
		goto L296
	} else {
		goto L297
	}
L289:
	;
	v1728 = *(*int32)(unsafe.Add(mBase, uint32(v572)+12))
	v1732 = *(*int32)(unsafe.Add(mBase, uint32(v1727+v1728<<(uint(int32(2))%32))))
	if v1732 != 0 {
		v1749 = v1727
		goto L288
	} else {
		goto L292
	}
L290:
	;
	goto L291
L291:
	;
	F_syncrep_yyensure_buffer_stack(m, v572)
	mBase = m.M
	v1734 = m.ExcPending
	if v1734 != 0 {
		goto L6
	} else {
		goto L293
	}
L292:
	;
	goto L291
L293:
	;
	v1735 = *(*int32)(unsafe.Add(mBase, uint32(v572)+4))
	v1736 = F_syncrep_yy_create_buffer(m, v1735, v572)
	mBase = m.M
	v1737 = m.ExcPending
	if v1737 != 0 {
		goto L6
	} else {
		goto L294
	}
L294:
	;
	v1738 = *(*int32)(unsafe.Add(mBase, uint32(v572)+20))
	v1739 = *(*int32)(unsafe.Add(mBase, uint32(v572)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1738+v1739<<(uint(int32(2))%32)))) = v1736
	v1744 = *(*int32)(unsafe.Add(mBase, uint32(v572)+20))
	if v1744 != 0 {
		v1749 = v1744
		goto L288
	} else {
		goto L295
	}
L295:
	;
	v1746 = *(*int32)(unsafe.Add(mBase, _c_F_check_synchronous_standby_names[0]))
	v1796 = int32(0)
	v1797 = v1746
	goto L287
L296:
	;
	v1796 = int32(0)
	v1797 = v1751
	goto L287
L297:
	;
	goto L298
L298:
	;
	v1760 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1756)+16)) = v1760
	v1762 = *(*int32)(unsafe.Add(mBase, uint32(v1756)+4))
	*(*uint8)(unsafe.Add(mBase, uint32(v1762))) = uint8(v1760)
	v1765 = *(*int32)(unsafe.Add(mBase, uint32(v1756)+4))
	*(*uint8)(unsafe.Add(mBase, uint32(v1765)+1)) = uint8(v1760)
	*(*int32)(unsafe.Add(mBase, uint32(v1756)+44)) = v1760
	*(*int32)(unsafe.Add(mBase, uint32(v1756)+28)) = int32(1)
	v1772 = *(*int32)(unsafe.Add(mBase, uint32(v1756)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1756)+8)) = v1772
	v1774 = *(*int32)(unsafe.Add(mBase, uint32(v572)+20))
	if v1774 == v1760 {
		v1796 = v1756
		v1797 = v1751
		goto L287
	} else {
		goto L299
	}
L299:
	;
	v1777 = *(*int32)(unsafe.Add(mBase, uint32(v572)+12))
	v1780 = v1774 + v1777<<(uint(int32(2))%32)
	v1781 = *(*int32)(unsafe.Add(mBase, uint32(v1780)))
	if v1756 != v1781 {
		v1796 = v1756
		v1797 = v1751
		goto L287
	} else {
		goto L300
	}
L300:
	;
	v1783 = *(*int32)(unsafe.Add(mBase, uint32(v1781)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v572)+28)) = v1783
	v1785 = *(*int32)(unsafe.Add(mBase, uint32(v1780)))
	v1786 = *(*int32)(unsafe.Add(mBase, uint32(v1785)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v572)+80)) = v1786
	*(*int32)(unsafe.Add(mBase, uint32(v572)+36)) = v1786
	v1789 = *(*int32)(unsafe.Add(mBase, uint32(v1780)))
	v1790 = *(*int32)(unsafe.Add(mBase, uint32(v1789)))
	*(*int32)(unsafe.Add(mBase, uint32(v572)+4)) = v1790
	v1792 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1786))))
	*(*uint8)(unsafe.Add(mBase, uint32(v572)+24)) = uint8(v1792)
	v1796 = v1756
	v1797 = v1751
	goto L287
L301:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1796)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_check_synchronous_standby_names[0])) = v1797
	v1814 = *(*int32)(unsafe.Add(mBase, uint32(v572)+20))
	v1815 = *(*int32)(unsafe.Add(mBase, uint32(v572)+12))
	v1818 = v1814 + v1815<<(uint(int32(2))%32)
	v1819 = *(*int32)(unsafe.Add(mBase, uint32(v1818)))
	v1820 = *(*int32)(unsafe.Add(mBase, uint32(v1819)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v572)+28)) = v1820
	v1822 = *(*int32)(unsafe.Add(mBase, uint32(v1818)))
	v1823 = *(*int32)(unsafe.Add(mBase, uint32(v1822)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v572)+36)) = v1823
	*(*int32)(unsafe.Add(mBase, uint32(v572)+80)) = v1823
	v1826 = *(*int32)(unsafe.Add(mBase, uint32(v1818)))
	v1827 = *(*int32)(unsafe.Add(mBase, uint32(v1826)))
	*(*int32)(unsafe.Add(mBase, uint32(v572)+4)) = v1827
	v1829 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1823))))
	*(*uint8)(unsafe.Add(mBase, uint32(v572)+24)) = uint8(v1829)
	v1846 = int32(1)
	goto L282
L302:
	;
	v1802 = *(*int32)(unsafe.Add(mBase, uint32(v572)+12))
	v1806 = *(*int32)(unsafe.Add(mBase, uint32(v1801+v1802<<(uint(int32(2))%32))))
	if v1796 == v1806 {
		goto L301
	} else {
		goto L305
	}
L303:
	;
	goto L304
L304:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1796)+32)) = int64(1)
	goto L301
L305:
	;
	goto L304
L306:
	;
	v1859 = v1848 + v1847>>(uint(int32(1))%32)
	v1860 = *(*int32)(unsafe.Add(mBase, uint32(v1854)+4))
	if v1860 != 0 {
		goto L310
	} else {
		goto L311
	}
L307:
	;
	v1889 = v1849
	v1891 = v1848
	v1892 = v1850
	goto L308
L308:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v572)+28)) = v1891
	v1894 = int32(2)
	v1897 = *(*int32)(unsafe.Add(mBase, uint32(v1889+v1892<<(uint(v1894)%32))))
	v1898 = *(*int32)(unsafe.Add(mBase, uint32(v1897)+4))
	v1900 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1898+v1891))) = uint8(v1900)
	v1902 = *(*int32)(unsafe.Add(mBase, uint32(v572)+20))
	v1903 = *(*int32)(unsafe.Add(mBase, uint32(v572)+12))
	v1907 = *(*int32)(unsafe.Add(mBase, uint32(v1902+v1903<<(uint(v1894)%32))))
	v1908 = *(*int32)(unsafe.Add(mBase, uint32(v1907)+4))
	v1909 = *(*int32)(unsafe.Add(mBase, uint32(v572)+28))
	*(*uint8)(unsafe.Add(mBase, uint32(v1908+v1909)+1)) = uint8(v1900)
	v1913 = *(*int32)(unsafe.Add(mBase, uint32(v572)+20))
	v1914 = *(*int32)(unsafe.Add(mBase, uint32(v572)+12))
	v1917 = v1913 + v1914<<(uint(v1894)%32)
	v1918 = *(*int32)(unsafe.Add(mBase, uint32(v1917)))
	v1919 = *(*int32)(unsafe.Add(mBase, uint32(v1918)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v572)+80)) = v1919
	if v1846 == int32(1) {
		v2260 = v1919
		goto L131
	} else {
		goto L316
	}
L309:
	;
	v1866 = *(*int32)(unsafe.Add(mBase, uint32(v572)+20))
	v1867 = *(*int32)(unsafe.Add(mBase, uint32(v572)+12))
	v1868 = int32(2)
	v1871 = *(*int32)(unsafe.Add(mBase, uint32(v1866+v1867<<(uint(v1868)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v1871)+4)) = v1865
	v1873 = *(*int32)(unsafe.Add(mBase, uint32(v572)+20))
	v1874 = *(*int32)(unsafe.Add(mBase, uint32(v572)+12))
	v1878 = *(*int32)(unsafe.Add(mBase, uint32(v1873+v1874<<(uint(v1868)%32))))
	v1879 = *(*int32)(unsafe.Add(mBase, uint32(v1878)+4))
	if v1879 == int32(0) {
		goto L133
	} else {
		goto L315
	}
L310:
	;
	v1861 = F_repalloc(m, v1860, v1859)
	mBase = m.M
	v1862 = m.ExcPending
	if v1862 != 0 {
		goto L6
	} else {
		goto L313
	}
L311:
	;
	goto L312
L312:
	;
	v1863 = F_palloc(m, v1859)
	mBase = m.M
	v1864 = m.ExcPending
	if v1864 != 0 {
		goto L6
	} else {
		goto L314
	}
L313:
	;
	v1865 = v1861
	goto L309
L314:
	;
	v1865 = v1863
	goto L309
L315:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1878)+12)) = v1859 - int32(2)
	v1885 = *(*int32)(unsafe.Add(mBase, uint32(v572)+12))
	v1886 = *(*int32)(unsafe.Add(mBase, uint32(v572)+28))
	v1888 = *(*int32)(unsafe.Add(mBase, uint32(v572)+20))
	v1889 = v1888
	v1891 = v1886 + v1202
	v1892 = v1885
	goto L308
L316:
	;
	switch v1846 - int32(1) {
	case 0:
		goto L134
	case 1:
		goto L317
	default:
		goto L318
	}
L317:
	;
	v2073 = *(*int32)(unsafe.Add(mBase, uint32(v572)+28))
	v2074 = *(*int32)(unsafe.Add(mBase, uint32(v1917)))
	v2075 = *(*int32)(unsafe.Add(mBase, uint32(v2074)+4))
	v2081 = v1919
	v2082 = v2075
	v2089 = v2073
	goto L135
L318:
	;
	v1927 = v868 ^ int32(-1) + v750
	v1928 = v1919 + v1927
	*(*int32)(unsafe.Add(mBase, uint32(v572)+36)) = v1928
	v1930 = *(*int32)(unsafe.Add(mBase, uint32(v572)+44))
	if v1927 <= int32(0) {
		v573 = v1930
		v574 = v1919
		v577 = v1928
		goto L110
	} else {
		goto L319
	}
L319:
	;
	v1937 = v1930
	v1944 = v1919
	goto L320
L320:
	;
	v1959 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1944))))
	if v1959 != 0 {
		goto L322
	} else {
		goto L323
	}
L321:
	;
	v573 = v2069
	v574 = v1919
	v577 = v1928
	goto L110
L322:
	;
	v1960 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1959)+uint32(_c_F_check_synchronous_standby_names[6]))))
	v1962 = v1960
	goto L324
L323:
	;
	v1962 = int32(1)
	goto L324
L324:
	;
	if base.Ui32(int32(-26)) <= base.Ui32(v1937-int32(31)) {
		goto L325
	} else {
		goto L326
	}
L325:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v572)+68)) = v1944
	*(*int32)(unsafe.Add(mBase, uint32(v572)+64)) = v1937
	goto L327
L326:
	;
	goto L327
L327:
	;
	v1970 = v1962 & int32(255)
	v1971 = int32(1)
	v1975 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1937<<(uint(v1971)%32))+uint32(_c_F_check_synchronous_standby_names[7]))))
	v1976 = v1970 + v1975
	v1981 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1976<<(uint(v1971)%32))+uint32(_c_F_check_synchronous_standby_names[8]))))
	if v1981 != v1937 {
		goto L328
	} else {
		goto L329
	}
L328:
	;
	v1987 = v1937
	v1992 = v1962
	v1993 = v1970
	goto L331
L329:
	;
	v2045 = v1976
	goto L330
L330:
	;
	v2065 = int32(1)
	v2069 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2045<<(uint(v2065)%32))+uint32(_c_F_check_synchronous_standby_names[9]))))
	v2071 = v1944 + v2065
	if v1928 != v2071 {
		v1937 = v2069
		v1944 = v2071
		goto L320
	} else {
		goto L337
	}
L331:
	;
	v2013 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1987<<(uint(int32(1))%32))+uint32(_c_F_check_synchronous_standby_names[10]))))
	if int64(1)<<(uint(base.I64_extend_i32_u(v1987))%64)&int64(2076672024) != int64(0) {
		goto L333
	} else {
		goto L334
	}
L332:
	;
	v2045 = v2030
	goto L330
L333:
	;
	v2021 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1993)+uint32(_c_F_check_synchronous_standby_names[11]))))
	v2022 = v2021
	goto L335
L334:
	;
	v2022 = v1992
	goto L335
L335:
	;
	v2024 = v2022 & int32(255)
	v2025 = int32(1)
	v2029 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2013<<(uint(v2025)%32))+uint32(_c_F_check_synchronous_standby_names[7]))))
	v2030 = v2024 + v2029
	v2035 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2030<<(uint(v2025)%32))+uint32(_c_F_check_synchronous_standby_names[8]))))
	if v2035 != v2013&int32(_a_F_check_synchronous_standby_names_10) {
		v1987 = v2013
		v1992 = v2022
		v1993 = v2024
		goto L331
	} else {
		goto L336
	}
L336:
	;
	goto L332
L337:
	;
	goto L321
L338:
	;
	v2110 = v2104
	v2114 = v2081
	goto L339
L339:
	;
	v2132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2114))))
	if v2132 != 0 {
		goto L341
	} else {
		goto L342
	}
L340:
	;
	v743 = v2242
	v746 = v2081
	v750 = v2102
	goto L127
L341:
	;
	v2133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2132)+uint32(_c_F_check_synchronous_standby_names[6]))))
	v2135 = v2133
	goto L343
L342:
	;
	v2135 = int32(1)
	goto L343
L343:
	;
	if base.Ui32(int32(-26)) <= base.Ui32(v2110-int32(31)) {
		goto L344
	} else {
		goto L345
	}
L344:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v572)+68)) = v2114
	*(*int32)(unsafe.Add(mBase, uint32(v572)+64)) = v2110
	goto L346
L345:
	;
	goto L346
L346:
	;
	v2143 = v2135 & int32(255)
	v2144 = int32(1)
	v2148 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2110<<(uint(v2144)%32))+uint32(_c_F_check_synchronous_standby_names[7]))))
	v2149 = v2143 + v2148
	v2154 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2149<<(uint(v2144)%32))+uint32(_c_F_check_synchronous_standby_names[8]))))
	if v2154 != v2110 {
		goto L347
	} else {
		goto L348
	}
L347:
	;
	v2160 = v2110
	v2165 = v2135
	v2166 = v2143
	goto L350
L348:
	;
	v2218 = v2149
	goto L349
L349:
	;
	v2238 = int32(1)
	v2242 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2218<<(uint(v2238)%32))+uint32(_c_F_check_synchronous_standby_names[9]))))
	v2244 = v2114 + v2238
	if v2244 != v2102 {
		v2110 = v2242
		v2114 = v2244
		goto L339
	} else {
		goto L356
	}
L350:
	;
	v2186 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2160<<(uint(int32(1))%32))+uint32(_c_F_check_synchronous_standby_names[10]))))
	if int64(1)<<(uint(base.I64_extend_i32_u(v2160))%64)&int64(2076672024) != int64(0) {
		goto L352
	} else {
		goto L353
	}
L351:
	;
	v2218 = v2203
	goto L349
L352:
	;
	v2194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2166)+uint32(_c_F_check_synchronous_standby_names[11]))))
	v2195 = v2194
	goto L354
L353:
	;
	v2195 = v2165
	goto L354
L354:
	;
	v2197 = v2195 & int32(255)
	v2198 = int32(1)
	v2202 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2186<<(uint(v2198)%32))+uint32(_c_F_check_synchronous_standby_names[7]))))
	v2203 = v2197 + v2202
	v2208 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2203<<(uint(v2198)%32))+uint32(_c_F_check_synchronous_standby_names[8]))))
	if v2208 != v2186&int32(_a_F_check_synchronous_standby_names_10) {
		v2160 = v2186
		v2165 = v2195
		v2166 = v2197
		goto L350
	} else {
		goto L355
	}
L355:
	;
	goto L351
L356:
	;
	goto L340
L357:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L358:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L359:
	;
	v2319 = int32(0)
	v2328 = v2319
	v2329 = v2319
	goto L84
L360:
	;
	goto L361
L361:
	;
	if v2295 == int32(256) {
		goto L362
	} else {
		goto L363
	}
L362:
	;
	v2550 = v2292
	v2551 = v2293
	v2554 = int32(1)
	v2556 = v2291
	v2563 = v2305
	goto L52
L363:
	;
	goto L364
L364:
	;
	if base.Ui32(int32(262)) < base.Ui32(v2295) {
		v2328 = v2295
		v2329 = int32(2)
		goto L84
	} else {
		goto L365
	}
L365:
	;
	v2327 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2295)+uint32(_c_F_check_synchronous_standby_names[13]))))
	v2328 = v2295
	v2329 = v2327
	goto L84
L366:
	;
	v2333 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2330)+uint32(_c_F_check_synchronous_standby_names[14]))))
	if v2329 != v2333 {
		v2341 = v2291
		v2342 = v2292
		v2343 = v2293
		v2344 = v2294
		v2345 = v2328
		v2353 = v2303
		v2355 = v2305
		v2356 = v2306
		v2357 = v2307
		v2359 = v2309
		v2360 = v2310
		v2361 = v2311
		v2363 = v2313
		v2364 = v2314
		goto L82
	} else {
		goto L367
	}
L367:
	;
	v2335 = *(*int32)(unsafe.Add(mBase, uint32(v2305)+1020))
	*(*int32)(unsafe.Add(mBase, uint32(v2303)+4)) = v2335
	v2340 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2330)+uint32(_c_F_check_synchronous_standby_names[15]))))
	v2518 = v2291
	v2519 = v2292
	v2520 = v2293
	v2521 = v2294
	v2522 = int32(-2)
	v2530 = v2303 + int32(4)
	v2532 = v2305
	v2533 = v2306
	v2534 = v2307
	v2536 = v2309
	v2538 = v2311
	v2540 = v2313
	v2544 = v2340
	goto L56
L368:
	;
	v2373 = v2341
	v2374 = v2342
	v2375 = v2343
	v2376 = v2344
	v2378 = int32(1)
	v2381 = int32(_a_F_check_synchronous_standby_names_22)
	v2387 = v2355
	v2391 = v2359
	goto L58
L369:
	;
	v2406 = *(*int32)(unsafe.Add(mBase, uint32(v2376)+80))
	v2407 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2406))))
	if v2407 != 0 {
		goto L373
	} else {
		goto L374
	}
L370:
	;
	goto L371
L371:
	;
	m.G0 = v2401 + int32(32)
	v2550 = v2374
	v2551 = v2375
	v2554 = v2378
	v2556 = v2373
	v2563 = v2387
	goto L52
L372:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2391))) = v2419
	goto L371
L373:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2401)+20)) = v2406
	*(*int32)(unsafe.Add(mBase, uint32(v2401)+16)) = v2381
	v2413 = F_psprintf(m, int32(_a_F_check_synchronous_standby_names_19), v2401+int32(16))
	mBase = m.M
	v2414 = m.ExcPending
	if v2414 != 0 {
		goto L6
	} else {
		goto L376
	}
L374:
	;
	goto L375
L375:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2401))) = v2381
	v2417 = F_psprintf(m, int32(_a_F_check_synchronous_standby_names_20), v2401)
	mBase = m.M
	v2418 = m.ExcPending
	if v2418 != 0 {
		goto L6
	} else {
		goto L377
	}
L376:
	;
	v2419 = v2413
	goto L372
L377:
	;
	v2419 = v2417
	goto L372
L378:
	;
	v2493 = v2353 - v2431<<(uint(int32(2))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v2493)+4)) = v2490
	v2496 = v2493 + int32(4)
	v2497 = v2356 - v2431
	v2498 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2497))))
	v2501 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2428)+uint32(_c_F_check_synchronous_standby_names[16]))))
	v2504 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2501)+uint32(_c_F_check_synchronous_standby_names[16]))))
	v2505 = v2498 + v2504
	if base.Ui32(v2505) <= base.Ui32(int32(22)) {
		goto L394
	} else {
		goto L395
	}
L379:
	;
	v2489 = *(*int32)(unsafe.Add(mBase, uint32(v2353)))
	v2490 = v2489
	goto L378
L380:
	;
	v2485 = *(*int32)(unsafe.Add(mBase, uint32(v2353-int32(8))))
	v2486 = *(*int32)(unsafe.Add(mBase, uint32(v2353)))
	v2487 = F_lappend(m, v2485, v2486)
	mBase = m.M
	v2488 = m.ExcPending
	if v2488 != 0 {
		goto L6
	} else {
		goto L392
	}
L381:
	;
	v2475 = *(*int32)(unsafe.Add(mBase, uint32(v2353)))
	*(*int32)(unsafe.Add(mBase, uint32(v2355)+8)) = v2475
	*(*int32)(unsafe.Add(mBase, uint32(v2355)+12)) = v2475
	v2481 = F_list_make1_impl(m, int32(1), v2355+int32(8))
	mBase = m.M
	v2482 = m.ExcPending
	if v2482 != 0 {
		goto L6
	} else {
		goto L391
	}
L382:
	;
	v2468 = *(*int32)(unsafe.Add(mBase, uint32(v2353-int32(12))))
	v2471 = *(*int32)(unsafe.Add(mBase, uint32(v2353-int32(4))))
	v2473 = F_create_syncrep_config(m, v2468, v2471, int32(0))
	mBase = m.M
	v2474 = m.ExcPending
	if v2474 != 0 {
		goto L6
	} else {
		goto L390
	}
L383:
	;
	v2459 = *(*int32)(unsafe.Add(mBase, uint32(v2353-int32(12))))
	v2462 = *(*int32)(unsafe.Add(mBase, uint32(v2353-int32(4))))
	v2464 = F_create_syncrep_config(m, v2459, v2462, int32(1))
	mBase = m.M
	v2465 = m.ExcPending
	if v2465 != 0 {
		goto L6
	} else {
		goto L389
	}
L384:
	;
	v2450 = *(*int32)(unsafe.Add(mBase, uint32(v2353-int32(12))))
	v2453 = *(*int32)(unsafe.Add(mBase, uint32(v2353-int32(4))))
	v2455 = F_create_syncrep_config(m, v2450, v2453, int32(0))
	mBase = m.M
	v2456 = m.ExcPending
	if v2456 != 0 {
		goto L6
	} else {
		goto L388
	}
L385:
	;
	v2444 = *(*int32)(unsafe.Add(mBase, uint32(v2353)))
	v2446 = F_create_syncrep_config(m, int32(_a_F_check_synchronous_standby_names_23), v2444, int32(0))
	mBase = m.M
	v2447 = m.ExcPending
	if v2447 != 0 {
		goto L6
	} else {
		goto L387
	}
L386:
	;
	v2441 = *(*int32)(unsafe.Add(mBase, uint32(v2353)))
	*(*int32)(unsafe.Add(mBase, uint32(v2363))) = v2441
	v2490 = v2436
	goto L378
L387:
	;
	v2490 = v2446
	goto L378
L388:
	;
	v2490 = v2455
	goto L378
L389:
	;
	v2490 = v2464
	goto L378
L390:
	;
	v2490 = v2473
	goto L378
L391:
	;
	v2490 = v2481
	goto L378
L392:
	;
	v2490 = v2487
	goto L378
L393:
	;
	v2517 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2505)+uint32(_c_F_check_synchronous_standby_names[15]))))
	v2518 = v2341
	v2519 = v2342
	v2520 = v2343
	v2521 = v2344
	v2522 = v2345
	v2530 = v2496
	v2532 = v2355
	v2533 = v2497
	v2534 = v2357
	v2536 = v2359
	v2538 = v2361
	v2540 = v2363
	v2544 = v2517
	goto L56
L394:
	;
	v2508 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2505)+uint32(_c_F_check_synchronous_standby_names[14]))))
	if v2508 == v2498&int32(255) {
		goto L393
	} else {
		goto L397
	}
L395:
	;
	goto L396
L396:
	;
	v2514 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2501)+uint32(_c_F_check_synchronous_standby_names[17]))))
	v2518 = v2341
	v2519 = v2342
	v2520 = v2343
	v2521 = v2344
	v2522 = v2345
	v2530 = v2496
	v2532 = v2355
	v2533 = v2497
	v2534 = v2357
	v2536 = v2359
	v2538 = v2361
	v2540 = v2363
	v2544 = v2514
	goto L56
L397:
	;
	goto L396
L398:
	;
	F_pfree(m, v2556)
	mBase = m.M
	v2579 = m.ExcPending
	if v2579 != 0 {
		goto L6
	} else {
		goto L401
	}
L399:
	;
	goto L400
L400:
	;
	m.G0 = v2563 + int32(1024)
	v2583 = *(*int32)(unsafe.Add(mBase, uint32(v2551)+44))
	F_replication_scanner_finish(m, v2583)
	mBase = m.M
	v2585 = m.ExcPending
	if v2585 != 0 {
		goto L6
	} else {
		goto L402
	}
L401:
	;
	goto L400
L402:
	;
	if v2554 == int32(0) {
		goto L404
	} else {
		goto L405
	}
L403:
	;
	v2614 = *(*int32)(unsafe.Add(mBase, uint32(v2588)+4))
	if v2614 <= int32(0) {
		goto L416
	} else {
		goto L417
	}
L404:
	;
	v2588 = *(*int32)(unsafe.Add(mBase, uint32(v2551)+40))
	if v2588 != 0 {
		goto L403
	} else {
		goto L407
	}
L405:
	;
	goto L406
L406:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_synchronous_standby_names[18])) = int32(16801924)
	goto L408
L407:
	;
	goto L406
L408:
	;
	v2593 = *(*int32)(unsafe.Add(mBase, uint32(v2551)+36))
	v2595 = *(*int32)(unsafe.Add(mBase, _c_F_check_synchronous_standby_names[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_check_synchronous_standby_names[19])) = v2595
	goto L409
L409:
	;
	if v2593 != 0 {
		goto L411
	} else {
		goto L412
	}
L410:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_synchronous_standby_names[20])) = v2611
	v2673 = v2551
	v2697 = int32(0)
	goto L1
L411:
	;
	v2599 = *(*int32)(unsafe.Add(mBase, uint32(v2551)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v2551)+16)) = v2599
	v2604 = F_format_elog_string(m, int32(_a_F_check_synchronous_standby_names_24), v2551+int32(16))
	mBase = m.M
	v2605 = m.ExcPending
	if v2605 != 0 {
		goto L6
	} else {
		goto L414
	}
L412:
	;
	goto L413
L413:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2551))) = int32(_a_F_check_synchronous_standby_names_25)
	v2609 = F_format_elog_string(m, int32(_a_F_check_synchronous_standby_names_26), v2551)
	mBase = m.M
	v2610 = m.ExcPending
	if v2610 != 0 {
		goto L6
	} else {
		goto L415
	}
L414:
	;
	v2611 = v2604
	goto L410
L415:
	;
	v2611 = v2609
	goto L410
L416:
	;
	v2618 = *(*int32)(unsafe.Add(mBase, _c_F_check_synchronous_standby_names[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_check_synchronous_standby_names[19])) = v2618
	goto L419
L417:
	;
	goto L418
L418:
	;
	v2633 = *(*int32)(unsafe.Add(mBase, uint32(v2588)))
	v2634 = F_guc_malloc(m, v2633)
	mBase = m.M
	v2635 = m.ExcPending
	if v2635 != 0 {
		goto L6
	} else {
		goto L421
	}
L419:
	;
	v2621 = *(*int32)(unsafe.Add(mBase, uint32(v2551)+40))
	v2622 = *(*int32)(unsafe.Add(mBase, uint32(v2621)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v2551)+32)) = v2622
	v2628 = F_format_elog_string(m, int32(_a_F_check_synchronous_standby_names_27), v2551+int32(32))
	mBase = m.M
	v2629 = m.ExcPending
	if v2629 != 0 {
		goto L6
	} else {
		goto L420
	}
L420:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_synchronous_standby_names[21])) = v2628
	v2673 = v2551
	v2697 = int32(0)
	goto L1
L421:
	;
	if v2634 == int32(0) {
		v2673 = v2551
		v2697 = int32(0)
		goto L1
	} else {
		goto L422
	}
L422:
	;
	v2638 = *(*int32)(unsafe.Add(mBase, uint32(v2551)+40))
	v2639 = *(*int32)(unsafe.Add(mBase, uint32(v2638)))
	if v2639 != 0 {
		goto L423
	} else {
		goto L424
	}
L423:
	;
	base.MemoryCopy(m, v2634, v2638, v2639)
	goto L425
L424:
	;
	goto L425
L425:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2550))) = v2634
	v2646 = v2551
	goto L2
}
func F_chmod(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	v3 = m.Env.X__syscall_chmod(m, l0, l1)
	mBase = m.M
	if base.Ui32(int32(-4095)) <= base.Ui32(v3) {
		*(*int32)(unsafe.Add(mBase, _c_F_chmod[0])) = int32(0) - v3
		v11 = int32(-1)
	} else {
		v11 = v3
	}
	return v11
}
func F_chr(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_chr[0]))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if int32(0) <= v11 {
		if v11 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v215 = m.ExcPending
			if v215 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(261))
				mBase = m.M
				v218 = m.ExcPending
				if v218 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_chr_0), int32(0))
					mBase = m.M
					v222 = m.ExcPending
					if v222 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_chr_1), int32(1049), int32(_a_F_chr_2))
						mBase = m.M
						v227 = m.ExcPending
						if v227 != 0 {
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
			if base.B2i32(v14 != int32(6))|base.B2i32(base.Ui32(v11) < base.Ui32(int32(128))) == int32(0) {
				if base.Ui32(int32(_a_F_chr_3)) <= base.Ui32(v11) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v231 = m.ExcPending
					if v231 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(261))
						mBase = m.M
						v234 = m.ExcPending
						if v234 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v11
							F_errmsg(m, int32(_a_F_chr_4), v9)
							mBase = m.M
							v238 = m.ExcPending
							if v238 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_chr_1), int32(1068), int32(_a_F_chr_2))
								mBase = m.M
								v243 = m.ExcPending
								if v243 != 0 {
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
					v32 = base.B2i32(base.Ui32(int32(2047)) < base.Ui32(v11))
					if base.Ui32(int32(2047)) < base.Ui32(v11) {
						v33 = int32(3)
					} else {
						v33 = int32(2)
					}
					if base.Ui32(int32(_a_F_chr_5)) < base.Ui32(v11) {
						v36 = int32(4)
					} else {
						v36 = v33
					}
					v38 = v36 + int32(4)
					v39 = F_palloc(m, v38)
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v39))) = v38 << (uint(int32(2)) % 32)
						v47 = v39 + int32(4)
						if v32 == int32(0) {
							v53 = int32(base.Ui32(v11)>>(uint(int32(6))%32)) | int32(192)
							*(*uint8)(unsafe.Add(mBase, uint32(v47))) = uint8(v53)
							v91 = int32(5)
						} else {
							if v11 <= int32(_a_F_chr_5) {
								v61 = int32(base.Ui32(v11)>>(uint(int32(12))%32)) | int32(224)
								*(*uint8)(unsafe.Add(mBase, uint32(v39)+4)) = uint8(v61)
								v63 = int32(6)
								v68 = int32(base.Ui32(v11)>>(uint(v63)%32))&int32(63) | int32(128)
								*(*uint8)(unsafe.Add(mBase, uint32(v39)+5)) = uint8(v68)
								v91 = v63
							} else {
								v74 = int32(base.Ui32(v11)>>(uint(int32(18))%32)) | int32(240)
								*(*uint8)(unsafe.Add(mBase, uint32(v39)+4)) = uint8(v74)
								v78 = int32(63)
								v80 = int32(128)
								v81 = int32(base.Ui32(v11)>>(uint(int32(6))%32))&v78 | v80
								*(*uint8)(unsafe.Add(mBase, uint32(v39)+6)) = uint8(v81)
								v88 = int32(base.Ui32(v11)>>(uint(int32(12))%32))&v78 | v80
								*(*uint8)(unsafe.Add(mBase, uint32(v39)+5)) = uint8(v88)
								v91 = int32(7)
							}
						}
						v96 = v11&int32(63) | int32(128)
						*(*uint8)(unsafe.Add(mBase, uint32(v91+v39))) = uint8(v96)
						v98 = int32(0)
						switch v36 - int32(1) {
						case 0:
							v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47))))
							v134 = v133
							if base.I32_extend8_s(v134) < int32(-62) {
								v147 = v98
							} else {
								v139 = v134
								v147 = base.B2i32(base.Ui32(v139&int32(255)) < base.Ui32(int32(245)))
							}
						case 1:
							v107 = int32(*(*int8)(unsafe.Add(mBase, uint32(v47)+1)))
							v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47))))
							switch v108 - int32(224) {
							case 0:
								v111 = int32(224)
								if base.Ui32(v111) <= base.Ui32((v107-int32(-64))&int32(255)) {
									v139 = v111
									v147 = base.B2i32(base.Ui32(v139&int32(255)) < base.Ui32(int32(245)))
								} else {
									v147 = v98
								}
							default:
								if v107 <= int32(-65) {
									v134 = v108
									if base.I32_extend8_s(v134) < int32(-62) {
										v147 = v98
									} else {
										v139 = v134
										v147 = base.B2i32(base.Ui32(v139&int32(255)) < base.Ui32(int32(245)))
									}
								} else {
									v147 = v98
								}
							case 13:
								if int32(-97) < v107 {
									v147 = v98
								} else {
									v139 = int32(237)
									v147 = base.B2i32(base.Ui32(v139&int32(255)) < base.Ui32(int32(245)))
								}
							case 16:
								if base.Ui32((v107-int32(-64))&int32(255)) < base.Ui32(int32(208)) {
									v147 = v98
								} else {
									v139 = int32(240)
									v147 = base.B2i32(base.Ui32(v139&int32(255)) < base.Ui32(int32(245)))
								}
							case 20:
								if int32(-113) < v107 {
									v147 = v98
								} else {
									v139 = int32(244)
									v147 = base.B2i32(base.Ui32(v139&int32(255)) < base.Ui32(int32(245)))
								}
							}
						case 2:
							v104 = int32(*(*int8)(unsafe.Add(mBase, uint32(v47)+2)))
							if int32(-65) < v104 {
								v147 = v98
							} else {
								v107 = int32(*(*int8)(unsafe.Add(mBase, uint32(v47)+1)))
								v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47))))
								switch v108 - int32(224) {
								case 0:
									v111 = int32(224)
									if base.Ui32(v111) <= base.Ui32((v107-int32(-64))&int32(255)) {
										v139 = v111
										v147 = base.B2i32(base.Ui32(v139&int32(255)) < base.Ui32(int32(245)))
									} else {
										v147 = v98
									}
								default:
									if v107 <= int32(-65) {
										v134 = v108
										if base.I32_extend8_s(v134) < int32(-62) {
											v147 = v98
										} else {
											v139 = v134
											v147 = base.B2i32(base.Ui32(v139&int32(255)) < base.Ui32(int32(245)))
										}
									} else {
										v147 = v98
									}
								case 13:
									if int32(-97) < v107 {
										v147 = v98
									} else {
										v139 = int32(237)
										v147 = base.B2i32(base.Ui32(v139&int32(255)) < base.Ui32(int32(245)))
									}
								case 16:
									if base.Ui32((v107-int32(-64))&int32(255)) < base.Ui32(int32(208)) {
										v147 = v98
									} else {
										v139 = int32(240)
										v147 = base.B2i32(base.Ui32(v139&int32(255)) < base.Ui32(int32(245)))
									}
								case 20:
									if int32(-113) < v107 {
										v147 = v98
									} else {
										v139 = int32(244)
										v147 = base.B2i32(base.Ui32(v139&int32(255)) < base.Ui32(int32(245)))
									}
								}
							}
						case 3:
							v101 = int32(*(*int8)(unsafe.Add(mBase, uint32(v47)+3)))
							if int32(-65) < v101 {
								v147 = v98
							} else {
								v104 = int32(*(*int8)(unsafe.Add(mBase, uint32(v47)+2)))
								if int32(-65) < v104 {
									v147 = v98
								} else {
									v107 = int32(*(*int8)(unsafe.Add(mBase, uint32(v47)+1)))
									v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47))))
									switch v108 - int32(224) {
									case 0:
										v111 = int32(224)
										if base.Ui32(v111) <= base.Ui32((v107-int32(-64))&int32(255)) {
											v139 = v111
											v147 = base.B2i32(base.Ui32(v139&int32(255)) < base.Ui32(int32(245)))
										} else {
											v147 = v98
										}
									default:
										if v107 <= int32(-65) {
											v134 = v108
											if base.I32_extend8_s(v134) < int32(-62) {
												v147 = v98
											} else {
												v139 = v134
												v147 = base.B2i32(base.Ui32(v139&int32(255)) < base.Ui32(int32(245)))
											}
										} else {
											v147 = v98
										}
									case 13:
										if int32(-97) < v107 {
											v147 = v98
										} else {
											v139 = int32(237)
											v147 = base.B2i32(base.Ui32(v139&int32(255)) < base.Ui32(int32(245)))
										}
									case 16:
										if base.Ui32((v107-int32(-64))&int32(255)) < base.Ui32(int32(208)) {
											v147 = v98
										} else {
											v139 = int32(240)
											v147 = base.B2i32(base.Ui32(v139&int32(255)) < base.Ui32(int32(245)))
										}
									case 20:
										if int32(-113) < v107 {
											v147 = v98
										} else {
											v139 = int32(244)
											v147 = base.B2i32(base.Ui32(v139&int32(255)) < base.Ui32(int32(245)))
										}
									}
								}
							}
						default:
							v147 = v98
						}
						if v147 != 0 {
							v188 = v39
							m.G0 = v9 + int32(48)
							return v188
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v151 = m.ExcPending
							if v151 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(261))
								mBase = m.M
								v154 = m.ExcPending
								if v154 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v11
									F_errmsg(m, int32(_a_F_chr_6), v9+int32(16))
									mBase = m.M
									v160 = m.ExcPending
									if v160 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_chr_1), int32(1109), int32(_a_F_chr_2))
										mBase = m.M
										v165 = m.ExcPending
										if v165 != 0 {
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
			} else {
				if base.Ui32(v14) <= base.Ui32(int32(41)) {
					v174 = *(*int32)(unsafe.Add(mBase, uint32(v14*int32(28))+uint32(_c_F_chr[1])))
					v176 = v174
				} else {
					v176 = int32(1)
				}
				if int32(1) < v176 {
					v179 = base.B2i32(base.Ui32(v11) < base.Ui32(int32(128)))
				} else {
					v179 = base.B2i32(base.Ui32(v11) < base.Ui32(int32(256)))
				}
				if v179 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v247 = m.ExcPending
					if v247 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(261))
						mBase = m.M
						v250 = m.ExcPending
						if v250 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v11
							F_errmsg(m, int32(_a_F_chr_4), v9+int32(32))
							mBase = m.M
							v256 = m.ExcPending
							if v256 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_chr_1), int32(1121), int32(_a_F_chr_2))
								mBase = m.M
								v261 = m.ExcPending
								if v261 != 0 {
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
					v183 = F_palloc(m, int32(5))
					mBase = m.M
					v184 = m.ExcPending
					if v184 != 0 {
						return int32(0)
					} else {
						*(*uint8)(unsafe.Add(mBase, uint32(v183)+4)) = uint8(v11)
						*(*int32)(unsafe.Add(mBase, uint32(v183))) = int32(20)
						v188 = v183
						m.G0 = v9 + int32(48)
						return v188
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v199 = m.ExcPending
		if v199 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50856066))
			mBase = m.M
			v202 = m.ExcPending
			if v202 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_chr_7), int32(0))
				mBase = m.M
				v206 = m.ExcPending
				if v206 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_chr_1), int32(1045), int32(_a_F_chr_2))
					mBase = m.M
					v211 = m.ExcPending
					if v211 != 0 {
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
func F_cidr_recv(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_network_recv(m, v2, int32(1))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_cidr_set_masklen_internal(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	v2 = l1
	v8 = F_palloc0(m, int32(22))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v12 = int32(1)
		v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
		v16 = v14 & v12
		if v16 != 0 {
			v17 = v12
		} else {
			v17 = int32(4)
		}
		v19 = int32(1)
		v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
		if v21&v19 != 0 {
			v24 = v19
		} else {
			v24 = int32(4)
		}
		v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v24))))
		*(*uint8)(unsafe.Add(mBase, uint32(v8+v17))) = uint8(v26)
		v29 = v8 + int32(1)
		v31 = v8 + int32(4)
		if v16 != 0 {
			v32 = v29
		} else {
			v32 = v31
		}
		*(*uint8)(unsafe.Add(mBase, uint32(v32)+1)) = uint8(v2)
		if v2 <= int32(0) {
		} else {
			v39 = base.I32_div_s(v2+int32(7), int32(8))
			if v39 != 0 {
				v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
				if v40&int32(1) != 0 {
					v43 = v29
				} else {
					v43 = v31
				}
				v46 = int32(1)
				v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
				if v50&v46 != 0 {
					v53 = l0 + v46
				} else {
					v53 = l0 + int32(4)
				}
				base.MemoryCopy(m, v43+int32(2), v53+int32(2), v39)
			} else {
			}
			v58 = v2 & int32(7)
			if v58 == int32(0) {
			} else {
				v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
				if v63&int32(1) != 0 {
					v66 = v29
				} else {
					v66 = v31
				}
				v67 = int32(base.Ui32(v2)>>(uint(int32(3))%32)) + v66
				v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+2)))
				v71 = v68 & (int32(-256) >> (uint(v58) % 32))
				*(*uint8)(unsafe.Add(mBase, uint32(v67)+2)) = uint8(v71)
			}
		}
		v78 = int32(1)
		v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
		if v80&v78 != 0 {
			v83 = v78
		} else {
			v83 = int32(4)
		}
		v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8+v83))))
		if v85 == int32(2) {
			v88 = int32(40)
		} else {
			v88 = int32(88)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v8))) = v88
		return v8
	}
}
func F_clause_selectivity_ext(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) float64 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v45 float64
	_ = v45
	var v50 float64
	_ = v50
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 float64
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 float64
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 float64
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 float64
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 float64
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 float64
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v128 float64
	_ = v128
	var v129 int32
	_ = v129
	var v130 float64
	_ = v130
	var v133 int32
	_ = v133
	var v139 int32
	_ = v139
	var v144 float64
	_ = v144
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v166 float64
	_ = v166
	var v167 int32
	_ = v167
	var v171 float64
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v183 float64
	_ = v183
	var v196 int32
	_ = v196
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 float64
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v223 float64
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v230 float64
	_ = v230
	var v231 int32
	_ = v231
	var v237 int32
	_ = v237
	var v240 float64
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v247 float64
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v280 float64
	_ = v280
	var v281 int32
	_ = v281
	var v282 float64
	_ = v282
	var v283 int32
	_ = v283
	var v284 float64
	_ = v284
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 float64
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v306 float32
	_ = v306
	var v307 float64
	_ = v307
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v320 float64
	_ = v320
	var v321 int32
	_ = v321
	var v322 float32
	_ = v322
	var v323 float64
	_ = v323
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 float64
	_ = v329
	var v330 float64
	_ = v330
	var v331 float64
	_ = v331
	var v337 int32
	_ = v337
	var v343 int32
	_ = v343
	var v348 int32
	_ = v348
	var v351 float64
	_ = v351
	var v355 int32
	_ = v355
	var v367 int32
	_ = v367
	var v373 int32
	_ = v373
	var v378 int32
	_ = v378
	var v384 float64
	_ = v384
	var v385 int32
	_ = v385
	var v390 int32
	_ = v390
	var v394 int32
	_ = v394
	var v399 int32
	_ = v399
	var v400 float64
	_ = v400
	var v401 int32
	_ = v401
	var v403 float64
	_ = v403
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v411 float64
	_ = v411
	var v421 float64
	_ = v421
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 float64
	_ = v431
	var v436 float64
	_ = v436
	var v437 int32
	_ = v437
	var v438 float64
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v441 float64
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v445 float64
	_ = v445
	var v446 int32
	_ = v446
	var v447 float64
	_ = v447
	var v449 int32
	_ = v449
	var v452 float64
	_ = v452
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v470 float64
	_ = v470
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v479 float64
	_ = v479
	var v489 float64
	_ = v489
	var v503 int32
	_ = v503
	var v512 float64
	_ = v512
	v11 = int32(0)
	if l1 == v11 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return float64(0.5)
L2:
	;
	goto L3
L3:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v21 != int32(318) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return v512
L5:
	;
	switch v65 - int32(6) {
	case 0:
		goto L44
	case 1:
		goto L43
	case 2:
		goto L42
	default:
		goto L29
	case 9:
		goto L39
	case 11, 12:
		goto L40
	case 14:
		goto L38
	case 15:
		goto L41
	case 21:
		goto L33
	case 31:
		goto L37
	case 46:
		goto L36
	case 47:
		goto L35
	case 49:
		goto L32
	case 52:
		goto L34
	}
L6:
	;
	v64 = l1
	v65 = v21
	v66 = v11
	v67 = v11
	goto L5
L7:
	;
	goto L8
L8:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+10)))
	if v24 != int32(1) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	if l2 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	if v28 == int32(7) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	return float64(1)
L12:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	if v57 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L13:
	;
	if l3 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L14:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	switch v35 {
	case 0:
		goto L13
	case 1:
		goto L15
	default:
		v56 = v11
		goto L12
	}
L15:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v37 = F_bms_is_member(m, l2, v36)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	return float64(0)
L17:
	;
	if v37 == int32(0) {
		v56 = v11
		goto L12
	} else {
		goto L18
	}
L18:
	;
	goto L13
L19:
	;
	v56 = int32(1)
	goto L12
L20:
	;
	v45 = *(*float64)(unsafe.Add(mBase, uint32(l1)+80))
	if base.F64_ge(v45, float64(0)) == int32(0) {
		goto L19
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v50 = *(*float64)(unsafe.Add(mBase, uint32(l1)+88))
	if base.F64_ge(v50, float64(0)) != 0 {
		v512 = v50
		goto L4
	} else {
		goto L24
	}
L23:
	;
	v512 = v45
	goto L4
L24:
	;
	goto L19
L25:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v61 = v60
	goto L27
L26:
	;
	v61 = v57
	goto L27
L27:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	v64 = v61
	v65 = v62
	v66 = l1
	v67 = v56
	goto L5
L28:
	;
	if v67 == int32(0) {
		v512 = v489
		goto L4
	} else {
		goto L186
	}
L29:
	;
	v454 = m.G0
	v456 = v454 - int32(32)
	m.G0 = v456
	F_examine_variable(m, l0, v64, l2, v456)
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L16
	} else {
		goto L178
	}
L30:
	;
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	if v449 == int32(18) {
		goto L175
	} else {
		goto L176
	}
L31:
	;
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v64)+28))
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v64)+24))
	v445 = F_restriction_selectivity(m, l0, v196, v443, v444, l2)
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L16
	} else {
		goto L174
	}
L32:
	;
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	v441 = F_clause_selectivity_ext(m, l0, v440, l2, l3, l4, l5)
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L16
	} else {
		goto L173
	}
L33:
	;
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	v438 = F_clause_selectivity_ext(m, l0, v437, l2, l3, l4, l5)
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L16
	} else {
		goto L172
	}
L34:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	v429 = F_find_base_rel(m, l0, v428)
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L16
	} else {
		goto L168
	}
L35:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v64)+8))
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	v294 = m.G0
	v296 = v294 - int32(112)
	m.G0 = v296
	F_examine_variable(m, l0, v293, l2, v296+int32(80))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L16
	} else {
		goto L122
	}
L36:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v64)+8))
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	v290 = F_nulltestsel(m, l0, v288, v289, l2)
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L16
	} else {
		goto L121
	}
L37:
	;
	v249 = m.G0
	v251 = v249 - int32(16)
	m.G0 = v251
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v64)+16))
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v253)+12))
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v254)))
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v64)+8))
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v256)+12))
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v257)))
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v64)+20))
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v259)+12))
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v260)))
	*(*int32)(unsafe.Add(mBase, uint32(v251)+12)) = v261
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v64)+24))
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v263)+12))
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v264)))
	*(*int32)(unsafe.Add(mBase, uint32(v251)+8)) = v265
	*(*int32)(unsafe.Add(mBase, uint32(v251)+4)) = v261
	*(*int32)(unsafe.Add(mBase, uint32(v251))) = v265
	v271 = F_list_make2_impl(m, v251+int32(4), v251)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L16
	} else {
		goto L113
	}
L38:
	;
	if l2|base.B2i32(l4 == int32(0)) != 0 {
		goto L104
	} else {
		goto L105
	}
L39:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v64)+28))
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v64)+24))
	if l2|base.B2i32(l4 == int32(0)) != 0 {
		goto L95
	} else {
		goto L96
	}
L40:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	if l2|base.B2i32(l4 == int32(0)) != 0 {
		goto L31
	} else {
		goto L86
	}
L41:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	switch v95 {
	case 0:
		goto L65
	case 1:
		goto L64
	case 2:
		goto L66
	default:
		goto L29
	}
L42:
	;
	v83 = F_estimate_expression_value(m, l0, v64)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L16
	} else {
		goto L54
	}
L43:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+24)))
	if v77 != 0 {
		goto L48
	} else {
		goto L49
	}
L44:
	;
	v71 = float64(0.5)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v64)+28))
	if v72 != 0 {
		v489 = v71
		goto L28
	} else {
		goto L45
	}
L45:
	;
	if l2 == int32(0) {
		goto L29
	} else {
		goto L46
	}
L46:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	if l2 != v75 {
		v489 = v71
		goto L28
	} else {
		goto L47
	}
L47:
	;
	goto L29
L48:
	;
	v489 = float64(0)
	goto L28
L49:
	;
	goto L50
L50:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v64)+20))
	if v81 != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v82 = float64(1)
	goto L53
L52:
	;
	v82 = float64(0)
	goto L53
L53:
	;
	v489 = v82
	goto L28
L54:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
	if v85 != int32(7) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v489 = float64(0.5)
	goto L28
L56:
	;
	goto L57
L57:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+24)))
	if v89 != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v489 = float64(0)
	goto L28
L59:
	;
	goto L60
L60:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v83)+20))
	if v93 != 0 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v94 = float64(1)
	goto L63
L62:
	;
	v94 = float64(0)
	goto L63
L63:
	;
	v489 = v94
	goto L28
L64:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v64)+8))
	v107 = float64(0)
	v108 = m.G0
	v110 = v108 - int32(16)
	m.G0 = v110
	v112 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v110)+12)) = v112
	v116 = F_find_single_rel_for_clauses(m, l0, v106)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L16
	} else {
		goto L70
	}
L65:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v64)+8))
	v104 = F_clauselist_selectivity_ext(m, l0, v103, l2, l3, l4, l5)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L16
	} else {
		goto L68
	}
L66:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v64)+8))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)+12))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	v100 = F_clause_selectivity_ext(m, l0, v99, l2, l3, l4, l5)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L16
	} else {
		goto L67
	}
L67:
	;
	v489 = base.F64_sub(float64(1), v100)
	goto L28
L68:
	;
	v489 = v104
	goto L28
L69:
	;
	if v106 == int32(0) {
		v183 = v130
		goto L75
	} else {
		goto L76
	}
L70:
	;
	if base.B2i32(l5 == v112)|base.B2i32(v116 == int32(0)) != 0 {
		v130 = v107
		goto L69
	} else {
		goto L71
	}
L71:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v116)+76))
	if v121 != 0 {
		v130 = v107
		goto L69
	} else {
		goto L72
	}
L72:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v116)+112))
	if v122 == int32(0) {
		v130 = v107
		goto L69
	} else {
		goto L73
	}
L73:
	;
	v128 = F_statext_clauselist_selectivity(m, l0, v106, l2, l3, l4, v116, v110+int32(12), int32(1))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L16
	} else {
		goto L74
	}
L74:
	;
	v130 = v128
	goto L69
L75:
	;
	m.G0 = v110 + int32(16)
	v489 = v183
	goto L28
L76:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v106)+4))
	if v133 <= int32(0) {
		v183 = v130
		goto L75
	} else {
		goto L77
	}
L77:
	;
	v139 = int32(0)
	v144 = v130
	v151 = int32(-1)
	goto L78
L78:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v106)+12))
	v156 = v151 + int32(1)
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v110)+12))
	v158 = F_bms_is_member(m, v156, v157)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L16
	} else {
		goto L80
	}
L79:
	;
	v183 = v171
	goto L75
L80:
	;
	if v158 == int32(0) {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v154+v139<<(uint(int32(2))%32))))
	v166 = F_clause_selectivity_ext(m, l0, v165, l2, l3, l4, l5)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L16
	} else {
		goto L84
	}
L82:
	;
	v171 = v144
	goto L83
L83:
	;
	v174 = v139 + int32(1)
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v106)+4))
	if v174 < v175 {
		v139 = v174
		v144 = v171
		v151 = v156
		goto L78
	} else {
		goto L85
	}
L84:
	;
	v171 = base.F64_sub(base.F64_add(v144, v166), base.F64_mul(v144, v166))
	goto L83
L85:
	;
	goto L79
L86:
	;
	if v66 == int32(0) {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v64)+28))
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v64)+24))
	v211 = F_join_selectivity(m, l0, v196, v209, v210, l3, l4)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L16
	} else {
		goto L94
	}
L88:
	;
	v202 = F_NumRelids(m, l0, v64)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L16
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v66)+24))
	if v206 < int32(2) {
		goto L31
	} else {
		goto L93
	}
L91:
	;
	if int32(1) < v202 {
		goto L87
	} else {
		goto L92
	}
L92:
	;
	goto L31
L93:
	;
	goto L87
L94:
	;
	v447 = v211
	goto L30
L95:
	;
	v229 = int32(0)
	goto L97
L96:
	;
	if v66 != 0 {
		goto L98
	} else {
		goto L99
	}
L97:
	;
	v230 = F_function_selectivity(m, l0, v213, v214, v215, v229, l2, l3, l4)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L16
	} else {
		goto L103
	}
L98:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v66)+24))
	v223 = F_function_selectivity(m, l0, v213, v214, v215, base.B2i32(int32(1) < v220), l2, l3, l4)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L16
	} else {
		goto L101
	}
L99:
	;
	goto L100
L100:
	;
	v225 = F_NumRelids(m, l0, v64)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L16
	} else {
		goto L102
	}
L101:
	;
	v489 = v223
	goto L28
L102:
	;
	v229 = base.B2i32(int32(1) < v225)
	goto L97
L103:
	;
	v489 = v230
	goto L28
L104:
	;
	v246 = int32(0)
	goto L106
L105:
	;
	if v66 != 0 {
		goto L107
	} else {
		goto L108
	}
L106:
	;
	v247 = F_scalararraysel(m, l0, v64, v246, l2, l3, l4)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L16
	} else {
		goto L112
	}
L107:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v66)+24))
	v240 = F_scalararraysel(m, l0, v64, base.B2i32(int32(1) < v237), l2, l3, l4)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L16
	} else {
		goto L110
	}
L108:
	;
	goto L109
L109:
	;
	v242 = F_NumRelids(m, l0, v64)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L16
	} else {
		goto L111
	}
L110:
	;
	v489 = v240
	goto L28
L111:
	;
	v246 = base.B2i32(int32(1) < v242)
	goto L106
L112:
	;
	v489 = v247
	goto L28
L113:
	;
	if l2|base.B2i32(l4 == int32(0)) != 0 {
		goto L115
	} else {
		goto L116
	}
L114:
	;
	m.G0 = v251 + int32(16)
	v489 = v284
	goto L28
L115:
	;
	v282 = F_restriction_selectivity(m, l0, v258, v271, v255, l2)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L16
	} else {
		goto L120
	}
L116:
	;
	v276 = F_NumRelids(m, l0, v271)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L16
	} else {
		goto L117
	}
L117:
	;
	if v276 < int32(2) {
		goto L115
	} else {
		goto L118
	}
L118:
	;
	v280 = F_join_selectivity(m, l0, v258, v271, v255, l3, l4)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L16
	} else {
		goto L119
	}
L119:
	;
	v284 = v280
	goto L114
L120:
	;
	v284 = v282
	goto L114
L121:
	;
	v489 = v290
	goto L28
L122:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v296)+88))
	if v302 != 0 {
		goto L125
	} else {
		goto L126
	}
L123:
	;
	m.G0 = v296 + int32(112)
	v489 = v421
	goto L28
L124:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v296)+88))
	if v407 != 0 {
		goto L162
	} else {
		goto L163
	}
L125:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v302)+16))
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v303)+22)))
	v306 = *(*float32)(unsafe.Add(mBase, uint32(v303+v304)+8))
	v307 = base.F64_promote_f32(v306)
	v313 = F_get_attstatsslot(m, v296+int32(44), v302, int32(1), int32(0), int32(3))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L16
	} else {
		goto L129
	}
L126:
	;
	goto L127
L127:
	;
	switch v292 {
	case 0, 3:
		goto L153
	case 1, 2:
		goto L155
	case 4:
		v421 = float64(0.005)
		goto L123
	case 5:
		goto L156
	default:
		goto L154
	}
L128:
	;
	switch v292 {
	case 0, 2:
		goto L149
	case 1, 3:
		goto L148
	case 4:
		v403 = v307
		goto L124
	case 5:
		goto L146
	default:
		goto L147
	}
L129:
	;
	if v313 == int32(0) {
		goto L128
	} else {
		goto L130
	}
L130:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v296)+68))
	if v317 <= int32(0) {
		goto L128
	} else {
		goto L131
	}
L131:
	;
	v320 = float64(1)
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v296)+64))
	v322 = *(*float32)(unsafe.Add(mBase, uint32(v321)))
	v323 = base.F64_promote_f32(v322)
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v296)+56))
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v327)))
	if v328 != 0 {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v329 = v323
	goto L134
L133:
	;
	v329 = base.F64_sub(base.F64_sub(v320, v323), v307)
	goto L134
L134:
	;
	v330 = base.F64_sub(v320, v329)
	v331 = base.F64_sub(v330, v307)
	switch v292 {
	case 0:
		goto L141
	case 1:
		goto L140
	case 2:
		goto L139
	case 3:
		goto L138
	case 4:
		v351 = v307
		goto L135
	case 5:
		goto L136
	default:
		goto L137
	}
L135:
	;
	F_free_attstatsslot(m, v296+int32(44))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L16
	} else {
		goto L145
	}
L136:
	;
	v351 = base.F64_sub(float64(1), v307)
	goto L135
L137:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L16
	} else {
		goto L142
	}
L138:
	;
	v351 = base.F64_sub(float64(1), v331)
	goto L135
L139:
	;
	v351 = v331
	goto L135
L140:
	;
	v351 = v330
	goto L135
L141:
	;
	v351 = v329
	goto L135
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v296)+16)) = v292
	F_errmsg_internal(m, int32(_a_F_clause_selectivity_ext_0), v296+int32(16))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L16
	} else {
		goto L143
	}
L143:
	;
	F_errfinish(m, int32(_a_F_clause_selectivity_ext_1), int32(1615), int32(_a_F_clause_selectivity_ext_2))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L16
	} else {
		goto L144
	}
L144:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L145:
	;
	v403 = v351
	goto L124
L146:
	;
	v403 = base.F64_sub(float64(1), v307)
	goto L124
L147:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L16
	} else {
		goto L150
	}
L148:
	;
	v403 = base.F64_mul(base.F64_add(v307, float64(1)), float64(0.5))
	goto L124
L149:
	;
	v403 = base.F64_mul(base.F64_sub(float64(1), v307), float64(0.5))
	goto L124
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v296)+32)) = v292
	F_errmsg_internal(m, int32(_a_F_clause_selectivity_ext_0), v296+int32(32))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L16
	} else {
		goto L151
	}
L151:
	;
	F_errfinish(m, int32(_a_F_clause_selectivity_ext_1), int32(1652), int32(_a_F_clause_selectivity_ext_2))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L16
	} else {
		goto L152
	}
L152:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L153:
	;
	v400 = F_clause_selectivity(m, l0, v293, l2, l3, l4)
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L16
	} else {
		goto L161
	}
L154:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L16
	} else {
		goto L158
	}
L155:
	;
	v384 = F_clause_selectivity(m, l0, v293, l2, l3, l4)
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L16
	} else {
		goto L157
	}
L156:
	;
	v421 = float64(0.995)
	goto L123
L157:
	;
	v403 = base.F64_sub(float64(1), v384)
	goto L124
L158:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v296))) = v292
	F_errmsg_internal(m, int32(_a_F_clause_selectivity_ext_0), v296)
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L16
	} else {
		goto L159
	}
L159:
	;
	F_errfinish(m, int32(_a_F_clause_selectivity_ext_1), int32(1688), int32(_a_F_clause_selectivity_ext_2))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L16
	} else {
		goto L160
	}
L160:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L161:
	;
	v403 = v400
	goto L124
L162:
	;
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v296)+92))
	m.T0[v408].(func(*base.Module, int32))(m, v407)
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L16
	} else {
		goto L165
	}
L163:
	;
	goto L164
L164:
	;
	v411 = float64(0)
	if base.F64_lt(v403, v411) != 0 {
		v421 = v411
		goto L123
	} else {
		goto L166
	}
L165:
	;
	goto L164
L166:
	;
	if base.F64_gt(v403, float64(1)) == int32(0) {
		v421 = v403
		goto L123
	} else {
		goto L167
	}
L167:
	;
	v421 = float64(1)
	goto L123
L168:
	;
	v431 = *(*float64)(unsafe.Add(mBase, uint32(v429)+120))
	if base.F64_gt(v431, float64(0)) != 0 {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	v436 = base.F64_div(float64(1), v431)
	goto L171
L170:
	;
	v436 = float64(0.5)
	goto L171
L171:
	;
	v489 = v436
	goto L28
L172:
	;
	v489 = v438
	goto L28
L173:
	;
	v489 = v441
	goto L28
L174:
	;
	v447 = v445
	goto L30
L175:
	;
	v452 = base.F64_sub(float64(1), v447)
	goto L177
L176:
	;
	v452 = v447
	goto L177
L177:
	;
	v489 = v452
	goto L28
L178:
	;
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v456)+8))
	if v460 == int32(0) {
		goto L180
	} else {
		goto L181
	}
L179:
	;
	m.G0 = v456 + int32(32)
	v489 = v479
	goto L28
L180:
	;
	v479 = float64(0.5)
	goto L179
L181:
	;
	goto L182
L182:
	;
	v465 = int32(0)
	v466 = int32(1)
	v470 = F_var_eq_const(m, v456, int32(91), v465, v466, v465, v466, v465)
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L16
	} else {
		goto L183
	}
L183:
	;
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v456)+8))
	if v472 == int32(0) {
		v479 = v470
		goto L179
	} else {
		goto L184
	}
L184:
	;
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v456)+12))
	m.T0[v475].(func(*base.Module, int32))(m, v472)
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L16
	} else {
		goto L185
	}
L185:
	;
	v479 = v470
	goto L179
L186:
	;
	if l3 != 0 {
		goto L187
	} else {
		goto L188
	}
L187:
	;
	v503 = int32(88)
	goto L189
L188:
	;
	v503 = int32(80)
	goto L189
L189:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v66+v503))) = v489
	v512 = v489
	goto L4
}
func F_clean_NOT_intree(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	F_check_stack_depth(m)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
		if v9 == int32(1) {
			return l0
		} else {
			v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+1)))
			switch v13 - int32(1) {
			case 0:
				F_freetree(m, l0)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int32(0)
				} else {
					return int32(0)
				}
			default:
				v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v35 = F_clean_NOT_intree(m, v34)
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0))) = v35
					v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v39 = F_clean_NOT_intree(m, v38)
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v39
						v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						if v39|v42 == int32(0) {
							F_pfree(m, l0)
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return int32(0)
							} else {
								return int32(0)
							}
						} else {
							if v42 == int32(0) {
								F_pfree(m, l0)
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return int32(0)
								} else {
									return v39
								}
							} else {
								if v39 != 0 {
									return l0
								} else {
									F_pfree(m, l0)
									mBase = m.M
									v57 = m.ExcPending
									if v57 != 0 {
										return int32(0)
									} else {
										return v42
									}
								}
							}
						}
					}
				}
			case 2:
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v17 = F_clean_NOT_intree(m, v16)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0))) = v17
					if v17 == int32(0) {
						F_freetree(m, l0)
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return int32(0)
						} else {
							return int32(0)
						}
					} else {
						v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v23 = F_clean_NOT_intree(m, v22)
						mBase = m.M
						v24 = m.ExcPending
						if v24 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v23
							if v23 == int32(0) {
								F_freetree(m, l0)
								mBase = m.M
								v31 = m.ExcPending
								if v31 != 0 {
									return int32(0)
								} else {
									return int32(0)
								}
							} else {
								return l0
							}
						}
					}
				}
			}
		}
	}
}
func F_clock_timestamp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int64
	_ = v8
	var v9 int64
	_ = v9
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	F_gettimeofday(m, v5)
	mBase = m.M
	v8 = int64(*(*int32)(unsafe.Add(mBase, uint32(v5)+8)))
	v9 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
	v15 = F_Int64GetDatum(m, v8+v9*int64(1000000)-int64(946684800000000))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		m.G0 = v5 + int32(16)
		return v15
	}
}
func F_clog_desc(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int64
	_ = v17
	var v21 int32
	_ = v21
	var v22 int64
	_ = v22
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+64))
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+48)))
	v14 = v12 & int32(240)
	if v14 != 0 {
		if v14 == int32(16) {
			v22 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = v23
			*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = v22
			F_appendStringInfo(m, l0, int32(_a_F_clog_desc_0), v8+int32(16))
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return
			} else {
				m.G0 = v8 + int32(32)
				return
			}
		} else {
			m.G0 = v8 + int32(32)
			return
		}
	} else {
		v17 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
		*(*int64)(unsafe.Add(mBase, uint32(v8))) = v17
		F_appendStringInfo(m, l0, int32(_a_F_clog_desc_1), v8)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return
		} else {
			m.G0 = v8 + int32(32)
			return
		}
	}
}
func F_closedir(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v4 = F_close(m, v3)
	mBase = m.M
	F_emscripten_builtin_free(m, l0)
	mBase = m.M
	return v4
}
func F_closerel(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	v4 = m.G0
	v6 = v4 - int32(48)
	m.G0 = v6
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_closerel[0]))
	if l0 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L17
	} else {
		goto L32
	}
L2:
	;
	v81 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L17
	} else {
		goto L25
	}
L3:
	;
	if v9 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	if v9 == int32(0) {
		goto L1
	} else {
		goto L24
	}
L6:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+48))
	v12 = v10 + int32(4)
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.B2i32(v15 == int32(0))|base.B2i32(v15 != v18) != 0 {
		v36 = v15
		v37 = v18
		goto L10
	} else {
		goto L11
	}
L7:
	;
	goto L8
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L17
	} else {
		goto L21
	}
L9:
	;
	if v36-v37 == int32(0) {
		goto L2
	} else {
		goto L16
	}
L10:
	;
	goto L9
L11:
	;
	v21 = v12
	v22 = l0
	goto L12
L12:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+1)))
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)))
	if v26 == int32(0) {
		v36 = v26
		v37 = v25
		goto L10
	} else {
		goto L14
	}
L13:
	;
	v36 = v26
	v37 = v25
	goto L10
L14:
	;
	v29 = int32(1)
	if v26 == v25 {
		v21 = v21 + v29
		v22 = v22 + v29
		goto L12
	} else {
		goto L15
	}
L15:
	;
	goto L13
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	return
L18:
	;
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_closerel[0]))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+32)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v6)+36)) = v47 + int32(4)
	F_errmsg_internal(m, int32(_a_F_closerel_0), v6+int32(32))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	F_errfinish(m, int32(_a_F_closerel_1), int32(493), int32(_a_F_closerel_2))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L17
	} else {
		goto L20
	}
L20:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = l0
	F_errmsg_internal(m, int32(_a_F_closerel_3), v6+int32(16))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L17
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(_a_F_closerel_1), int32(497), int32(_a_F_closerel_2))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L17
	} else {
		goto L23
	}
L23:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L24:
	;
	goto L2
L25:
	;
	if v81 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v84 = *(*int32)(unsafe.Add(mBase, _c_F_closerel[0]))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = v85 + int32(4)
	F_errmsg_internal(m, int32(_a_F_closerel_4), v6)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L17
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v98 = *(*int32)(unsafe.Add(mBase, _c_F_closerel[0]))
	F_relation_close(m, v98, int32(0))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L17
	} else {
		goto L31
	}
L29:
	;
	F_errfinish(m, int32(_a_F_closerel_1), int32(505), int32(_a_F_closerel_2))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L17
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_closerel[0])) = int32(0)
	m.G0 = v6 + int32(48)
	return
L32:
	;
	F_errmsg_internal(m, int32(_a_F_closerel_5), int32(0))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L17
	} else {
		goto L33
	}
L33:
	;
	F_errfinish(m, int32(_a_F_closerel_1), int32(501), int32(_a_F_closerel_2))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L17
	} else {
		goto L34
	}
L34:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_cmpLexemeQ(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v4 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return base.B2i32(v3 != int32(0))
L2:
	;
	goto L3
L3:
	;
	if v3 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(-1)
L5:
	;
	goto L6
L6:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4))))
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3))))
	if base.B2i32(v16 == int32(0))|base.B2i32(v16 != v19) != 0 {
		v37 = v16
		v38 = v19
		goto L8
	} else {
		goto L9
	}
L7:
	;
	return v37 - v38
L8:
	;
	goto L7
L9:
	;
	v22 = v4
	v23 = v3
	goto L10
L10:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+1)))
	if v27 == int32(0) {
		v37 = v27
		v38 = v26
		goto L8
	} else {
		goto L12
	}
L11:
	;
	v37 = v27
	v38 = v26
	goto L8
L12:
	;
	v30 = int32(1)
	if v27 == v26 {
		v22 = v22 + v30
		v23 = v23 + v30
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
}
func F_collect_corrupt_items(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	var v392 int32
	_ = v392
	var v398 int32
	_ = v398
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v445 int32
	_ = v445
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v505 int32
	_ = v505
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v524 int32
	_ = v524
	v2 = l1
	v3 = l2
	v4 = int32(0)
	v21 = m.G0
	v23 = v21 - int32(48)
	m.G0 = v23
	*(*int32)(unsafe.Add(mBase, uint32(v23)+44)) = v4
	v28 = F_GetAccessStrategy(m, int32(1))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v33 = F_relation_open(m, l0, int32(1))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v33)+48))
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+119)))
	v38 = v36 - int32(109)
	v45 = int32(0)
	if base.B2i32(base.Ui32(int32(7)) < base.Ui32(v38))|base.B2i32(int32(1)<<(uint(v38)%32)&int32(161) == v45) == v45 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	if v2 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L1
	} else {
		goto L132
	}
L7:
	;
	v50 = F_GetStrictOldestNonRemovableTransactionId(m, v33)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	v52 = v4
	goto L9
L9:
	;
	v54 = F_palloc0(m, int32(12))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L11
	}
L10:
	;
	v52 = v50
	goto L9
L11:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v54))) = int64(274877906944)
	v59 = F_palloc(m, int32(384))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+8)) = v59
	v62 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = v62
	v65 = F_RelationGetNumberOfBlocksInFork(m, v33, v62)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v67 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+40)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v23)+36)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v65
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+25)) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+24)) = uint8(v3)
	v79 = F_read_stream_begin_relation(m, int32(4), v28, v33, v67, int32(_a_F_collect_corrupt_items_0), v23+int32(24), v67)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v82 = F_read_stream_next_buffer(m, v79, int32(0))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	if v82 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v85 = v23 + int32(8)
	v93 = v52
	v95 = v82
	goto L19
L17:
	;
	goto L18
L18:
	;
	F_read_stream_end(m, v79)
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L1
	} else {
		goto L122
	}
L19:
	;
	v107 = *(*int32)(unsafe.Add(mBase, _c_F_collect_corrupt_items[0]))
	if v107 != 0 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	goto L18
L21:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	F_LockBuffer(m, v95, int32(1))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L25
	}
L24:
	;
	goto L23
L25:
	;
	if v95 < int32(0) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v131 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v130)+12)))
	if v95 < int32(0) {
		goto L31
	} else {
		goto L32
	}
L27:
	;
	v116 = *(*int32)(unsafe.Add(mBase, _c_F_collect_corrupt_items[1]))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v116+(v95^int32(-1))<<(uint(int32(2))%32))))
	v130 = v122
	goto L26
L28:
	;
	goto L29
L29:
	;
	v124 = *(*int32)(unsafe.Add(mBase, _c_F_collect_corrupt_items[2]))
	v130 = v124 + v95<<(uint(int32(13))%32) + int32(-8192)
	goto L26
L30:
	;
	if v3 != 0 {
		goto L34
	} else {
		goto L35
	}
L31:
	;
	v135 = *(*int32)(unsafe.Add(mBase, _c_F_collect_corrupt_items[3]))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v135+(v95^int32(-1))<<(uint(int32(6))%32))+16))
	v150 = v141
	goto L30
L32:
	;
	goto L33
L33:
	;
	v143 = *(*int32)(unsafe.Add(mBase, _c_F_collect_corrupt_items[4]))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v143+v95<<(uint(int32(6))%32)+int32(-64))+16))
	v150 = v149
	goto L30
L34:
	;
	v154 = F_visibilitymap_get_status(m, v33, v150, v23+int32(44))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L37
	}
L35:
	;
	v160 = int32(0)
	goto L36
L36:
	;
	if v2 != 0 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v160 = int32(base.Ui32(v154&int32(2)) >> (uint(int32(1)) % 32))
	goto L36
L38:
	;
	v164 = F_visibilitymap_get_status(m, v33, v150, v23+int32(44))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L1
	} else {
		goto L41
	}
L39:
	;
	v168 = int32(0)
	goto L40
L40:
	;
	if v160|v168 == int32(0) {
		v445 = v93
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v168 = v164 & int32(1)
	goto L40
L42:
	;
	F_UnlockReleaseBuffer(m, v95)
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L1
	} else {
		goto L119
	}
L43:
	;
	if base.Ui32(int32(25)) <= base.Ui32(v131) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v179 = int32(base.Ui32(v131+int32(_a_F_collect_corrupt_items_1)) >> (uint(int32(2)) % 32))
	goto L46
L45:
	;
	v179 = int32(0)
	goto L46
L46:
	;
	v181 = v179 & int32(_a_F_collect_corrupt_items_2)
	if v181 == int32(0) {
		v445 = v93
		goto L42
	} else {
		goto L47
	}
L47:
	;
	v185 = int32(base.Ui32(v150) >> (uint(int32(16)) % 32))
	v196 = v93
	v200 = int32(1)
	goto L48
L48:
	;
	v213 = v130 + int32(20) + v200&int32(_a_F_collect_corrupt_items_2)<<(uint(int32(2))%32)
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v213)))
	switch int32(base.Ui32(v214)>>(uint(int32(15))%32))&int32(3) - int32(1) {
	case 0:
		goto L52
	default:
		v432 = v196
		goto L50
	case 2:
		goto L53
	}
L49:
	;
	v445 = v432
	goto L42
L50:
	;
	v434 = v200 + int32(1)
	if base.Ui32(v434&int32(_a_F_collect_corrupt_items_2)) <= base.Ui32(v181) {
		v196 = v432
		v200 = v434
		goto L48
	} else {
		goto L118
	}
L51:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	if base.Ui32(v404) < base.Ui32(v405) {
		goto L114
	} else {
		goto L115
	}
L52:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+12)) = uint16(v200)
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+10)) = uint16(v150)
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+8)) = uint16(v185)
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v213)))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v130 + v227&int32(_a_F_collect_corrupt_items_3)
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v213)))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = int32(base.Ui32(v232) >> (uint(int32(17)) % 32))
	if v168 == int32(0) {
		v366 = v196
		goto L54
	} else {
		goto L55
	}
L53:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+12)) = uint16(v200)
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+10)) = uint16(v150)
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+8)) = uint16(v185)
	v403 = v196
	goto L51
L54:
	;
	if v160 == int32(0) {
		v432 = v366
		goto L50
	} else {
		goto L98
	}
L55:
	;
	v241 = F_HeapTupleSatisfiesVacuum(m, v23+int32(4), v196, v95)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	if v241 == int32(1) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
	v246 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v245)+20)))
	v247 = int32(768)
	if v246&v247 != v247 {
		goto L60
	} else {
		goto L61
	}
L58:
	;
	goto L59
L59:
	;
	v267 = F_GetStrictOldestNonRemovableTransactionId(m, v33)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L1
	} else {
		goto L69
	}
L60:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v245)))
	v253 = v251
	goto L62
L61:
	;
	v253 = int32(2)
	goto L62
L62:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v196))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v253)) == int32(0) {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	if v265 != 0 {
		v366 = v196
		goto L54
	} else {
		goto L67
	}
L64:
	;
	v265 = base.B2i32(base.Ui32(v253) < base.Ui32(v196))
	goto L63
L65:
	;
	goto L66
L66:
	;
	v265 = int32(base.Ui32(v253-v196) >> (uint(int32(31)) % 32))
	goto L63
L67:
	;
	goto L59
L68:
	;
	v360 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v85)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v358)+4)) = uint16(v360)
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	*(*int32)(unsafe.Add(mBase, uint32(v358))) = v362
	v366 = v359
	goto L54
L69:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v267))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v196)) == int32(0) {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	if v280 == int32(0) {
		goto L74
	} else {
		goto L75
	}
L71:
	;
	v280 = base.B2i32(base.Ui32(v196) < base.Ui32(v267))
	goto L70
L72:
	;
	goto L73
L73:
	;
	v280 = int32(base.Ui32(v196-v267) >> (uint(int32(31)) % 32))
	goto L70
L74:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	if base.Ui32(v283) < base.Ui32(v284) {
		goto L78
	} else {
		goto L79
	}
L75:
	;
	goto L76
L76:
	;
	v307 = F_HeapTupleSatisfiesVacuum(m, v23+int32(4), v267, v95)
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L1
	} else {
		goto L83
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54))) = v298 + int32(1)
	v358 = v297 + v298*int32(6)
	v359 = v196
	goto L68
L78:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
	v297 = v286
	v298 = v283
	goto L77
L79:
	;
	goto L80
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+4)) = v284 << (uint(int32(1)) % 32)
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
	v293 = F_repalloc(m, v290, v284*int32(12))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+8)) = v293
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	v297 = v293
	v298 = v296
	goto L77
L82:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	if base.Ui32(v335) < base.Ui32(v336) {
		goto L94
	} else {
		goto L95
	}
L83:
	;
	if v307 != int32(1) {
		goto L82
	} else {
		goto L84
	}
L84:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
	v312 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v311)+20)))
	v313 = int32(768)
	if v312&v313 != v313 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v311)))
	v319 = v317
	goto L87
L86:
	;
	v319 = int32(2)
	goto L87
L87:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v267))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v319)) == int32(0) {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	if v331 == int32(0) {
		goto L82
	} else {
		goto L92
	}
L89:
	;
	v331 = base.B2i32(base.Ui32(v319) < base.Ui32(v267))
	goto L88
L90:
	;
	goto L91
L91:
	;
	v331 = int32(base.Ui32(v319-v267) >> (uint(int32(31)) % 32))
	goto L88
L92:
	;
	v366 = v267
	goto L54
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54))) = v349 + int32(1)
	v358 = v350 + v349*int32(6)
	v359 = v267
	goto L68
L94:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
	v349 = v335
	v350 = v338
	goto L93
L95:
	;
	goto L96
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+4)) = v336 << (uint(int32(1)) % 32)
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
	v345 = F_repalloc(m, v342, v336*int32(12))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+8)) = v345
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	v349 = v348
	v350 = v345
	goto L93
L98:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
	v372 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v369)+20)))
	v373 = int32(768)
	if v372&v373 == v373 {
		goto L100
	} else {
		goto L101
	}
L99:
	;
	if v398 == int32(0) {
		v432 = v366
		goto L50
	} else {
		goto L112
	}
L100:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v369)+4))
	if v372&int32(_a_F_collect_corrupt_items_4) != 0 {
		goto L104
	} else {
		goto L105
	}
L101:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v369)))
	if base.Ui32(v377) <= base.Ui32(int32(2)) {
		goto L100
	} else {
		goto L102
	}
L102:
	;
	v398 = int32(1)
	goto L99
L103:
	;
	if base.Ui32(v372) < base.Ui32(int32(_a_F_collect_corrupt_items_5)) {
		goto L109
	} else {
		goto L110
	}
L104:
	;
	if v381 == int32(0) {
		goto L103
	} else {
		goto L107
	}
L105:
	;
	goto L106
L106:
	;
	if base.Ui32(v381) <= base.Ui32(int32(2)) {
		goto L103
	} else {
		goto L108
	}
L107:
	;
	v398 = int32(1)
	goto L99
L108:
	;
	v398 = int32(1)
	goto L99
L109:
	;
	v398 = int32(0)
	goto L99
L110:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v369)+8))
	if base.Ui32(v392) <= base.Ui32(int32(2)) {
		goto L109
	} else {
		goto L111
	}
L111:
	;
	v398 = int32(1)
	goto L99
L112:
	;
	v403 = v366
	goto L51
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54))) = v419 + int32(1)
	v425 = v418 + v419*int32(6)
	v426 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v85)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v425)+4)) = uint16(v426)
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	*(*int32)(unsafe.Add(mBase, uint32(v425))) = v428
	v432 = v403
	goto L50
L114:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
	v418 = v407
	v419 = v404
	goto L113
L115:
	;
	goto L116
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+4)) = v405 << (uint(int32(1)) % 32)
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
	v414 = F_repalloc(m, v411, v405*int32(12))
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L1
	} else {
		goto L117
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+8)) = v414
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	v418 = v414
	v419 = v417
	goto L113
L118:
	;
	goto L49
L119:
	;
	v461 = F_read_stream_next_buffer(m, v79, int32(0))
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L1
	} else {
		goto L120
	}
L120:
	;
	if v461 != 0 {
		v93 = v445
		v95 = v461
		goto L19
	} else {
		goto L121
	}
L121:
	;
	goto L20
L122:
	;
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v23)+44))
	if v485 != 0 {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	F_ReleaseBuffer(m, v485)
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L1
	} else {
		goto L126
	}
L124:
	;
	goto L125
L125:
	;
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v23)+40))
	if v488 != 0 {
		goto L127
	} else {
		goto L128
	}
L126:
	;
	goto L125
L127:
	;
	F_ReleaseBuffer(m, v488)
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L1
	} else {
		goto L130
	}
L128:
	;
	goto L129
L129:
	;
	F_relation_close(m, v33, int32(1))
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L1
	} else {
		goto L131
	}
L130:
	;
	goto L129
L131:
	;
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+4)) = v494
	*(*int32)(unsafe.Add(mBase, uint32(v54))) = int32(0)
	m.G0 = v23 + int32(48)
	return v54
L132:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L1
	} else {
		goto L133
	}
L133:
	;
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v33)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v509 + int32(4)
	F_errmsg(m, int32(_a_F_collect_corrupt_items_6), v23)
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L1
	} else {
		goto L134
	}
L134:
	;
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v33)+48))
	v517 = int32(*(*int8)(unsafe.Add(mBase, uint32(v516)+119)))
	F_errdetail_relkind_not_supported(m, v517)
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L1
	} else {
		goto L135
	}
L135:
	;
	F_errfinish(m, int32(_a_F_collect_corrupt_items_7), int32(951), int32(_a_F_collect_corrupt_items_8))
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_collect_corrupt_items_read_stream_next_block(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
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
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	v5 = int32(-1)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if base.Ui32(v7) <= base.Ui32(v6) {
		v63 = v5
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v63
L2:
	;
	v10 = l1 + int32(16)
	goto L3
L3:
	;
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_collect_corrupt_items_read_stream_next_block[0]))
	if v16 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v63 = v5
	goto L1
L5:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v21 == int32(1) {
		goto L14
	} else {
		goto L15
	}
L8:
	;
	return int32(0)
L9:
	;
	goto L7
L10:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v57 = v55 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v57
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if base.Ui32(v57) < base.Ui32(v59) {
		goto L3
	} else {
		goto L23
	}
L11:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v50 + int32(1)
	v63 = v50
	goto L1
L12:
	;
	if v29 == int32(0) {
		goto L10
	} else {
		goto L22
	}
L13:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v42 = F_visibilitymap_get_status(m, v40, v41, v10)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L8
	} else {
		goto L20
	}
L14:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v26 = F_visibilitymap_get_status(m, v24, v25, v10)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L8
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
	if v36 != int32(1) {
		goto L10
	} else {
		goto L19
	}
L17:
	;
	v29 = v26 & int32(2)
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
	if v30 == int32(0) {
		goto L12
	} else {
		goto L18
	}
L18:
	;
	v39 = int32(base.Ui32(v29) >> (uint(int32(1)) % 32))
	goto L13
L19:
	;
	v39 = int32(0)
	goto L13
L20:
	;
	if v42&int32(1)|v39 != 0 {
		goto L11
	} else {
		goto L21
	}
L21:
	;
	goto L10
L22:
	;
	goto L11
L23:
	;
	goto L4
}
func F_colorcomplement(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v89 int32
	_ = v89
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
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v148 int32
	_ = v148
	var v161 int32
	_ = v161
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v185 int32
	_ = v185
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v204 int32
	_ = v204
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v229 int32
	_ = v229
	var v239 int32
	_ = v239
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	if v12 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v173 = *(*int32)(unsafe.Add(mBase, _c_F_colorcomplement[0]))
	if v173 != 0 {
		goto L55
	} else {
		goto L56
	}
L2:
	;
	v19 = v12
	goto L5
L3:
	;
	v59 = v11
	goto L4
L4:
	;
	v62 = int32(24)
	v66 = v11 + v10*v62 + v62
	if base.Ui32(v66) <= base.Ui32(v59) {
		goto L18
	} else {
		goto L19
	}
L5:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	if v22 == int32(112) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v32 = v12
	goto L12
L7:
	;
	v25 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v19)+4)))
	if v25 == int32(_a_F_colorcomplement_0) {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
	if v28 != 0 {
		v19 = v28
		goto L5
	} else {
		goto L11
	}
L10:
	;
	goto L9
L11:
	;
	goto L6
L12:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	if v38 == int32(112) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v59 = v52
	goto L4
L14:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v42 = int32(*(*int16)(unsafe.Add(mBase, uint32(v32)+4)))
	v45 = v41 + v42*int32(24)
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v45)+20)) = v46 | int32(4)
	goto L16
L15:
	;
	goto L16
L16:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	if v51 != 0 {
		v32 = v51
		goto L12
	} else {
		goto L17
	}
L17:
	;
	goto L13
L18:
	;
	return
L19:
	;
	v75 = v59
	v76 = int32(0)
	goto L20
L20:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+12))
	if v79 != 0 {
		goto L18
	} else {
		goto L22
	}
L21:
	;
	goto L18
L22:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v75)+20))
	if v80&int32(4) != 0 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v161 = v75 + int32(24)
	if base.Ui32(v161) < base.Ui32(v66) {
		v75 = v161
		v76 = v76 + int32(1)
		goto L20
	} else {
		goto L54
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v75)+20)) = v80 & int32(-5)
	goto L23
L25:
	;
	goto L26
L26:
	;
	if v80&int32(3) != 0 {
		goto L23
	} else {
		goto L27
	}
L27:
	;
	v89 = *(*int32)(unsafe.Add(mBase, _c_F_colorcomplement[0]))
	if v89 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	goto L30
L30:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l5)+8))
	if v92 <= v93 {
		goto L34
	} else {
		goto L35
	}
L31:
	;
	return
L32:
	;
	goto L30
L33:
	;
	F_createarc(m, l0, l2, base.I32_extend16_s(v76), l4, l5)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L31
	} else {
		goto L53
	}
L34:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
	if v95 == int32(0) {
		goto L33
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l5)+16))
	if v116 == int32(0) {
		goto L33
	} else {
		goto L45
	}
L37:
	;
	v101 = v95
	goto L38
L38:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v101)+12))
	if v107 != l5 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	goto L33
L40:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v101)+16))
	if v115 != 0 {
		v101 = v115
		goto L38
	} else {
		goto L44
	}
L41:
	;
	v109 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v101)+4)))
	if v109 != v76&int32(_a_F_colorcomplement_1) {
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	if v113 == l2 {
		goto L23
	} else {
		goto L43
	}
L43:
	;
	goto L40
L44:
	;
	goto L39
L45:
	;
	v122 = v116
	goto L46
L46:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v122)+8))
	if v128 != l4 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	goto L33
L48:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v122)+24))
	if v136 != 0 {
		v122 = v136
		goto L46
	} else {
		goto L52
	}
L49:
	;
	v130 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v122)+4)))
	if v130 != v76&int32(_a_F_colorcomplement_1) {
		goto L48
	} else {
		goto L50
	}
L50:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
	if v134 == l2 {
		goto L23
	} else {
		goto L51
	}
L51:
	;
	goto L48
L52:
	;
	goto L47
L53:
	;
	goto L23
L54:
	;
	goto L21
L55:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L31
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l5)+8))
	if v176 <= v177 {
		goto L61
	} else {
		goto L62
	}
L58:
	;
	goto L57
L59:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v239 | int32(4)
	return
L60:
	;
	F_createarc(m, l0, int32(120), int32(0), l4, l5)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L31
	} else {
		goto L80
	}
L61:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
	if v179 == int32(0) {
		goto L60
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(l5)+16))
	if v198 == int32(0) {
		goto L60
	} else {
		goto L72
	}
L64:
	;
	v185 = v179
	goto L65
L65:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v185)+12))
	if v191 != l5 {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	goto L60
L67:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v185)+16))
	if v197 != 0 {
		v185 = v197
		goto L65
	} else {
		goto L71
	}
L68:
	;
	v193 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v185)+4)))
	if v193 != 0 {
		goto L67
	} else {
		goto L69
	}
L69:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v185)))
	if v194 == int32(120) {
		goto L59
	} else {
		goto L70
	}
L70:
	;
	goto L67
L71:
	;
	goto L66
L72:
	;
	v204 = v198
	goto L73
L73:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v204)+8))
	if v210 != l4 {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	goto L60
L75:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v204)+24))
	if v216 != 0 {
		v204 = v216
		goto L73
	} else {
		goto L79
	}
L76:
	;
	v212 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v204)+4)))
	if v212 != 0 {
		goto L75
	} else {
		goto L77
	}
L77:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v204)))
	if v213 == int32(120) {
		goto L59
	} else {
		goto L78
	}
L78:
	;
	goto L75
L79:
	;
	goto L74
L80:
	;
	goto L59
}
func F_combine(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	v6 = int32(3)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v11 = v7 | v8<<(uint(int32(8))%32)
	if v11 <= int32(_a_F_combine_0) {
		if v11 <= int32(_a_F_combine_1) {
			switch v11 - int32(_a_F_combine_2) {
			case 0, 18:
				v104 = v6
				return v104
			case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17:
				return int32(1)
			default:
				if v11 == int32(_a_F_combine_3) {
					v66 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
					v67 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)))
					if v66 == v67 {
						v69 = int32(2)
					} else {
						v69 = int32(1)
					}
					return v69
				} else {
					if v11 != int32(_a_F_combine_4) {
						return int32(1)
					} else {
						v104 = v6
						return v104
					}
				}
			}
		} else {
			switch v11 - int32(_a_F_combine_5) {
			case 0, 21:
				v104 = v6
				return v104
			case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 19, 20:
				return int32(1)
			case 18:
				v66 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
				v67 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)))
				if v66 == v67 {
					v69 = int32(2)
				} else {
					v69 = int32(1)
				}
				return v69
			default:
				if v11 != int32(_a_F_combine_6) {
					return int32(1)
				} else {
					v104 = v6
					return v104
				}
			}
		}
	} else {
		switch v11 - int32(_a_F_combine_7) {
		case 0, 18, 38:
			v104 = v6
			return v104
		case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 19, 20, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 37:
			return int32(1)
		case 21:
			v71 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
			v72 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)))
			if v71 == v72 {
				return int32(2)
			} else {
				if v71 == int32(_a_F_combine_8) {
					v78 = int32(2)
					v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
					v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+20))
					v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80+base.I32_extend16_s(v72)*int32(24))+20)))
					if v85&v78 != 0 {
						return int32(1)
					} else {
						v104 = v78
						return v104
					}
				} else {
					if v72 != int32(_a_F_combine_8) {
						return int32(1)
					} else {
						v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
						v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)+20))
						v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93+base.I32_extend16_s(v71)*int32(24))+20)))
						if v98&int32(2) != 0 {
							v101 = int32(1)
						} else {
							v101 = int32(4)
						}
						v104 = v101
						return v104
					}
				}
			}
		case 36:
			v32 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
			v33 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)))
			if v32 == v33 {
				return int32(2)
			} else {
				if v32 == int32(_a_F_combine_8) {
					v39 = int32(2)
					v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
					v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+20))
					v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41+base.I32_extend16_s(v33)*int32(24))+20)))
					if v46&v39 != 0 {
						return int32(1)
					} else {
						v104 = v39
						return v104
					}
				} else {
					if v33 != int32(_a_F_combine_8) {
						return int32(1)
					} else {
						v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
						v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+20))
						v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54+base.I32_extend16_s(v32)*int32(24))+20)))
						if v59&int32(2) != 0 {
							v62 = int32(1)
						} else {
							v62 = int32(4)
						}
						return v62
					}
				}
			}
		default:
			switch v11 - int32(_a_F_combine_9) {
			case 0, 21:
				v104 = v6
				return v104
			case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 37:
				return int32(1)
			case 36:
				v32 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
				v33 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)))
				if v32 == v33 {
					return int32(2)
				} else {
					if v32 == int32(_a_F_combine_8) {
						v39 = int32(2)
						v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
						v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+20))
						v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41+base.I32_extend16_s(v33)*int32(24))+20)))
						if v46&v39 != 0 {
							return int32(1)
						} else {
							v104 = v39
							return v104
						}
					} else {
						if v33 != int32(_a_F_combine_8) {
							return int32(1)
						} else {
							v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
							v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+20))
							v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54+base.I32_extend16_s(v32)*int32(24))+20)))
							if v59&int32(2) != 0 {
								v62 = int32(1)
							} else {
								v62 = int32(4)
							}
							return v62
						}
					}
				}
			case 38:
				v71 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
				v72 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)))
				if v71 == v72 {
					return int32(2)
				} else {
					if v71 == int32(_a_F_combine_8) {
						v78 = int32(2)
						v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
						v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+20))
						v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80+base.I32_extend16_s(v72)*int32(24))+20)))
						if v85&v78 != 0 {
							return int32(1)
						} else {
							v104 = v78
							return v104
						}
					} else {
						if v72 != int32(_a_F_combine_8) {
							return int32(1)
						} else {
							v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
							v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)+20))
							v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93+base.I32_extend16_s(v71)*int32(24))+20)))
							if v98&int32(2) != 0 {
								v101 = int32(1)
							} else {
								v101 = int32(4)
							}
							v104 = v101
							return v104
						}
					}
				}
			default:
				if v11 == int32(_a_F_combine_10) {
					v104 = v6
					return v104
				} else {
					return int32(1)
				}
			}
		}
	}
}
func F_compactify_tuples(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v180 int32
	_ = v180
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v319 int32
	_ = v319
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v334 int32
	_ = v334
	var v343 int32
	_ = v343
	v5 = int32(0)
	v14 = m.G0
	v16 = v14 + int32(-8192)
	m.G0 = v16
	if l3 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l2)+14)) = uint16(v343)
	m.G0 = v16 - int32(-8192)
	return
L2:
	;
	v18 = int32(1)
	if l1 <= v18 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v118 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v118) {
		goto L30
	} else {
		goto L31
	}
L5:
	;
	v21 = v18
	goto L7
L6:
	;
	v21 = l1
	goto L7
L7:
	;
	v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+16)))
	v26 = v22
	v29 = v5
	goto L9
L8:
	;
	v112 = v106 - v103
	if v112 == int32(0) {
		v343 = v102
		goto L1
	} else {
		goto L28
	}
L9:
	;
	v38 = l0 + v29*int32(6)
	v39 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38)+4)))
	v40 = int32(*(*int16)(unsafe.Add(mBase, uint32(v38)+2)))
	v41 = v39 + v40
	if v26 == v41 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	if l1 <= v29 {
		goto L15
	} else {
		goto L16
	}
L11:
	;
	v43 = v26 - v39
	v45 = v29 + int32(1)
	if v45 != v21 {
		v26 = v43
		v29 = v45
		goto L9
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	goto L10
L14:
	;
	v102 = v43
	v103 = v26
	v106 = v26
	goto L8
L15:
	;
	v102 = v26
	v103 = v41
	v106 = v41
	goto L8
L16:
	;
	goto L17
L17:
	;
	v53 = v26
	v54 = v41
	v56 = v29
	v57 = v41
	goto L18
L18:
	;
	v65 = l0 + v56*int32(6)
	v66 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v65))))
	v73 = l2 + int32(20) + (v66+int32(1))&int32(_a_F_compactify_tuples_0)<<(uint(int32(2))%32)
	v74 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v65)+4)))
	v75 = int32(*(*int16)(unsafe.Add(mBase, uint32(v65)+2)))
	if v74+v75 == v54 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v102 = v91
	v103 = v85
	v106 = v87
	goto L8
L20:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	v91 = v53 - v86
	*(*int32)(unsafe.Add(mBase, uint32(v73))) = v88&int32(-32768) | v91&int32(_a_F_compactify_tuples_1)
	v97 = v56 + int32(1)
	if v97 != l1 {
		v53 = v91
		v54 = v85
		v56 = v97
		v57 = v87
		goto L18
	} else {
		goto L27
	}
L21:
	;
	v85 = v75
	v86 = v74
	v87 = v57
	goto L20
L22:
	;
	goto L23
L23:
	;
	v78 = v57 - v54
	if v78 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	base.MemoryCopy(m, l2+v53, l2+v54, v78)
	goto L26
L25:
	;
	goto L26
L26:
	;
	v82 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v65)+4)))
	v83 = int32(*(*int16)(unsafe.Add(mBase, uint32(v65)+2)))
	v85 = v83
	v86 = v82
	v87 = v82 + v83
	goto L20
L27:
	;
	goto L19
L28:
	;
	base.MemoryCopy(m, l2+v102, l2+v103, v112)
	v343 = v102
	goto L1
L29:
	;
	if l1 <= v262 {
		goto L64
	} else {
		goto L65
	}
L30:
	;
	v128 = int32(base.Ui32(v118+int32(_a_F_compactify_tuples_2))>>(uint(int32(4))%32)) & int32(_a_F_compactify_tuples_3)
	goto L32
L31:
	;
	v128 = int32(0)
	goto L32
L32:
	;
	if l1 < v128 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	if int32(2) <= l1 {
		goto L37
	} else {
		goto L38
	}
L34:
	;
	goto L35
L35:
	;
	v218 = int32(1)
	if base.Ui32(l1) <= base.Ui32(v218) {
		goto L54
	} else {
		goto L55
	}
L36:
	;
	v213 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	v214 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+2)))
	v216 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+16)))
	v259 = v216
	v261 = v213 + v214
	v262 = int32(0)
	goto L29
L37:
	;
	v133 = int32(1)
	if l1 <= v133 {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	v180 = int32(0)
	goto L39
L39:
	;
	v192 = l0 + v180*int32(6)
	v193 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v192)+4)))
	if v193 == int32(0) {
		goto L36
	} else {
		goto L53
	}
L40:
	;
	v136 = v133
	goto L42
L41:
	;
	v136 = l1
	goto L42
L42:
	;
	v145 = int32(0)
	v151 = v5
	goto L43
L43:
	;
	v157 = l0 + v145*int32(6)
	v158 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v157)+4)))
	if v158 != 0 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	if v136&int32(1) == int32(0) {
		goto L36
	} else {
		goto L52
	}
L45:
	;
	v159 = int32(*(*int16)(unsafe.Add(mBase, uint32(v157)+2)))
	base.MemoryCopy(m, v16+v159, l2+v159, v158)
	goto L47
L46:
	;
	goto L47
L47:
	;
	v164 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v157)+10)))
	if v164 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v165 = int32(*(*int16)(unsafe.Add(mBase, uint32(v157)+8)))
	base.MemoryCopy(m, v16+v165, l2+v165, v164)
	goto L50
L49:
	;
	goto L50
L50:
	;
	v170 = int32(2)
	v171 = v145 + v170
	v173 = v151 + v170
	if v173 != v136&int32(2147483646) {
		v145 = v171
		v151 = v173
		goto L43
	} else {
		goto L51
	}
L51:
	;
	goto L44
L52:
	;
	v180 = v171
	goto L39
L53:
	;
	v196 = int32(*(*int16)(unsafe.Add(mBase, uint32(v192)+2)))
	base.MemoryCopy(m, v16+v196, l2+v196, v193)
	goto L36
L54:
	;
	v221 = v218
	goto L56
L55:
	;
	v221 = l1
	goto L56
L56:
	;
	v222 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+16)))
	v226 = v222
	v229 = v5
	goto L58
L57:
	;
	v249 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+14)))
	v250 = v247 - v249
	if v250 == int32(0) {
		v259 = v247
		v261 = v241
		v262 = v248
		goto L29
	} else {
		goto L62
	}
L58:
	;
	v238 = l0 + v229*int32(6)
	v239 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v238)+4)))
	v240 = int32(*(*int16)(unsafe.Add(mBase, uint32(v238)+2)))
	v241 = v239 + v240
	if v226 != v241 {
		v247 = v226
		v248 = v229
		goto L57
	} else {
		goto L60
	}
L59:
	;
	v247 = v243
	v248 = v221
	goto L57
L60:
	;
	v243 = v226 - v239
	v245 = v229 + int32(1)
	if v245 != v221 {
		v226 = v243
		v229 = v245
		goto L58
	} else {
		goto L61
	}
L61:
	;
	goto L59
L62:
	;
	base.MemoryCopy(m, v249+v16, l2+v249, v250)
	v259 = v247
	v261 = v241
	v262 = v248
	goto L29
L63:
	;
	v334 = v326 - v325
	if v334 == int32(0) {
		v343 = v324
		goto L1
	} else {
		goto L77
	}
L64:
	;
	v324 = v259
	v325 = v261
	v326 = v261
	goto L63
L65:
	;
	goto L66
L66:
	;
	v275 = v259
	v276 = v261
	v277 = v261
	v278 = v262
	goto L67
L67:
	;
	v287 = l0 + v278*int32(6)
	v288 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v287))))
	v295 = l2 + int32(20) + (v288+int32(1))&int32(_a_F_compactify_tuples_0)<<(uint(int32(2))%32)
	v296 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v287)+4)))
	v297 = int32(*(*int16)(unsafe.Add(mBase, uint32(v287)+2)))
	if v296+v297 == v276 {
		goto L70
	} else {
		goto L71
	}
L68:
	;
	v324 = v313
	v325 = v307
	v326 = v308
	goto L63
L69:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v295)))
	v313 = v275 - v309
	*(*int32)(unsafe.Add(mBase, uint32(v295))) = v310&int32(-32768) | v313&int32(_a_F_compactify_tuples_1)
	v319 = v278 + int32(1)
	if v319 != l1 {
		v275 = v313
		v276 = v307
		v277 = v308
		v278 = v319
		goto L67
	} else {
		goto L76
	}
L70:
	;
	v307 = v297
	v308 = v277
	v309 = v296
	goto L69
L71:
	;
	goto L72
L72:
	;
	v300 = v277 - v276
	if v300 != 0 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	base.MemoryCopy(m, l2+v275, v276+v16, v300)
	goto L75
L74:
	;
	goto L75
L75:
	;
	v304 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v287)+4)))
	v305 = int32(*(*int16)(unsafe.Add(mBase, uint32(v287)+2)))
	v307 = v305
	v308 = v304 + v305
	v309 = v304
	goto L69
L76:
	;
	goto L68
L77:
	;
	base.MemoryCopy(m, l2+v324, v325+v16, v334)
	v343 = v324
	goto L1
}
func F_compareDoubles(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 float64
	_ = v7
	var v8 float64
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	v7 = *(*float64)(unsafe.Add(mBase, uint32(l0)))
	v8 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
	if base.F64_gt(v7, v8) != 0 {
		v10 = int32(1)
	} else {
		v10 = int32(-1)
	}
	if base.F64_ne(v7, v8) != 0 {
		v13 = v10
	} else {
		v13 = int32(0)
	}
	return v13
}
func F_compareentry(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v5 = int32(12)
	v8 = int32(1)
	v10 = int32(2047)
	v11 = int32(base.Ui32(v4)>>(uint(v8)%32)) & v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v19 = int32(base.Ui32(v12)>>(uint(v8)%32)) & v10
	if v11 == int32(0) {
		v25 = int32(0)
		if v25 < v19 {
			v28 = int32(-1)
		} else {
			v28 = v25
		}
		v45 = v28
	} else {
		if v19 == int32(0) {
			v45 = base.B2i32(int32(0) < v11)
		} else {
			if base.Ui32(v11) < base.Ui32(v19) {
				v34 = v11
			} else {
				v34 = v19
			}
			v35 = F_memcmp(m, l2+int32(base.Ui32(v4)>>(uint(v5)%32)), l2+int32(base.Ui32(v12)>>(uint(v5)%32)), v34)
			mBase = m.M
			if v35 != 0 {
				v43 = v35
				v45 = v43
			} else {
				if v11 == v19 {
					v45 = int32(0)
				} else {
					if v11 < v19 {
						v42 = int32(-1)
					} else {
						v42 = int32(1)
					}
					v43 = v42
					v45 = v43
				}
			}
		}
	}
	return v45
}
func F_comparison_shim(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	v4 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = l0
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+44)) = uint8(v4)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v18 = m.T0[v17].(func(*base.Module, int32) int32)(m, v9+int32(28))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return int32(0)
	} else {
		v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+44)))
		if v22 == int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = v29
				F_errmsg_internal(m, int32(_a_F_comparison_shim_0), v7)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_comparison_shim_1), int32(58), int32(_a_F_comparison_shim_2))
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			m.G0 = v7 + int32(16)
			return v18
		}
	}
}
func F_compress_init(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+64)))
	if base.Ui32(v5-int32(3)) < base.Ui32(int32(-2)) {
		return int32(-102)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
		v15 = m.Env.Pgmem_deflate_create(m, base.B2i32(v5 == int32(1)), v14)
		mBase = m.M
		if v15 <= int32(0) {
			return int32(-105)
		} else {
			v21 = F_palloc0(m, int32(_a_F_compress_init_0))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = int32(_a_F_compress_init_1)
				v28 = *(*int32)(unsafe.Add(mBase, _c_F_compress_init[0]))
				F_ResourceOwnerEnlarge(m, v28)
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int32(0)
				} else {
					v32 = *(*int32)(unsafe.Add(mBase, _c_F_compress_init[1]))
					v34 = F_MemoryContextAlloc(m, v32, int32(8))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v34))) = v15
						v38 = *(*int32)(unsafe.Add(mBase, _c_F_compress_init[0]))
						*(*int32)(unsafe.Add(mBase, uint32(v34)+4)) = v38
						F_ResourceOwnerRemember(m, v38, v34, int32(_a_F_compress_init_2))
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v21))) = v34
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v21
							return int32(_a_F_compress_init_1)
						}
					}
				}
			}
		}
	}
}
func F_compute_gather_rows(m *base.Module, l0 int32) float64 {
	mBase := m.M
	_ = mBase
	var v5 float64
	_ = v5
	var v6 int32
	_ = v6
	var v7 float64
	_ = v7
	var v9 int32
	_ = v9
	var v15 float64
	_ = v15
	var v19 float64
	_ = v19
	var v21 float64
	_ = v21
	var v23 float64
	_ = v23
	var v24 float64
	_ = v24
	var v33 float64
	_ = v33
	var v37 float64
	_ = v37
	v5 = *(*float64)(unsafe.Add(mBase, uint32(l0)+32))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = base.F64_convert_i32_s(v6)
	v9 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_compute_gather_rows[0])))
	if v9 == int32(1) {
		v15 = base.F64_add(base.F64_mul(v7, float64(-0.3)), float64(1))
		if base.F64_gt(v15, float64(0)) != 0 {
			v19 = v15
		} else {
			v19 = math.Float64frombits(uint64(0x8000000000000000))
		}
		v21 = base.F64_add(v19, v7)
	} else {
		v21 = v7
	}
	v23 = float64(1e+100)
	v24 = base.F64_mul(v5, v21)
	if base.F64_gt(v24, v23)|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v24)&int64(9223372036854775807))) != 0 {
		v37 = v23
	} else {
		v33 = float64(1)
		if base.F64_le(v24, v33) != 0 {
			v37 = v33
		} else {
			v37 = base.F64_nearest(v24)
		}
	}
	return v37
}
func F_connectby_text(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
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
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
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
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	v16 = m.G0
	v18 = v16 - int32(16)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v21 = F_pg_detoast_datum_packed(m, v20)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		return int32(0)
	} else {
		v25 = F_text_to_cstring(m, v21)
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return int32(0)
		} else {
			v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			v28 = F_pg_detoast_datum_packed(m, v27)
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				v30 = F_text_to_cstring(m, v28)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int32(0)
				} else {
					v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
					v33 = F_pg_detoast_datum_packed(m, v32)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int32(0)
					} else {
						v35 = F_text_to_cstring(m, v33)
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return int32(0)
						} else {
							v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
							v38 = F_pg_detoast_datum_packed(m, v37)
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
								return int32(0)
							} else {
								v40 = F_text_to_cstring(m, v38)
								mBase = m.M
								v41 = m.ExcPending
								if v41 != 0 {
									return int32(0)
								} else {
									v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									if v42 == int32(0) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v129 = m.ExcPending
										if v129 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(1088))
											mBase = m.M
											v132 = m.ExcPending
											if v132 != 0 {
												return int32(0)
											} else {
												F_errmsg(m, int32(_a_F_connectby_text_0), int32(0))
												mBase = m.M
												v136 = m.ExcPending
												if v136 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_connectby_text_1), int32(997), int32(_a_F_connectby_text_2))
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
									} else {
										v45 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
										if v45 != int32(383) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v129 = m.ExcPending
											if v129 != 0 {
												return int32(0)
											} else {
												F_errcode(m, int32(1088))
												mBase = m.M
												v132 = m.ExcPending
												if v132 != 0 {
													return int32(0)
												} else {
													F_errmsg(m, int32(_a_F_connectby_text_0), int32(0))
													mBase = m.M
													v136 = m.ExcPending
													if v136 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_connectby_text_1), int32(997), int32(_a_F_connectby_text_2))
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
										} else {
											v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+12)))
											if v48&int32(2) == int32(0) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v145 = m.ExcPending
												if v145 != 0 {
													return int32(0)
												} else {
													F_errcode(m, int32(1088))
													mBase = m.M
													v148 = m.ExcPending
													if v148 != 0 {
														return int32(0)
													} else {
														F_errmsg(m, int32(_a_F_connectby_text_3), int32(0))
														mBase = m.M
														v152 = m.ExcPending
														if v152 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_connectby_text_1), int32(1002), int32(_a_F_connectby_text_2))
															mBase = m.M
															v157 = m.ExcPending
															if v157 != 0 {
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
												v53 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
												if v53 == int32(0) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v145 = m.ExcPending
													if v145 != 0 {
														return int32(0)
													} else {
														F_errcode(m, int32(1088))
														mBase = m.M
														v148 = m.ExcPending
														if v148 != 0 {
															return int32(0)
														} else {
															F_errmsg(m, int32(_a_F_connectby_text_3), int32(0))
															mBase = m.M
															v152 = m.ExcPending
															if v152 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(_a_F_connectby_text_1), int32(1002), int32(_a_F_connectby_text_2))
																mBase = m.M
																v157 = m.ExcPending
																if v157 != 0 {
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
													v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
													v57 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
													if v57 == int32(6) {
														v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
														v61 = F_pg_detoast_datum_packed(m, v60)
														mBase = m.M
														v62 = m.ExcPending
														if v62 != 0 {
															return int32(0)
														} else {
															v63 = F_text_to_cstring(m, v61)
															mBase = m.M
															v64 = m.ExcPending
															if v64 != 0 {
																return int32(0)
															} else {
																v68 = v63
																v69 = int32(_a_F_connectby_text_4)
																v70 = *(*int32)(unsafe.Add(mBase, _c_F_connectby_text[0]))
																v72 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
																v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)+16))
																*(*int32)(unsafe.Add(mBase, _c_F_connectby_text[0])) = v73
																v75 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
																v76 = F_CreateTupleDescCopy(m, v75)
																mBase = m.M
																v77 = m.ExcPending
																if v77 != 0 {
																	return int32(0)
																} else {
																	v79 = base.B2i32(v57 == int32(6))
																	F_validateConnectbyTupleDesc(m, v76, v79, int32(0))
																	mBase = m.M
																	v82 = m.ExcPending
																	if v82 != 0 {
																		return int32(0)
																	} else {
																		v83 = F_TupleDescGetAttInMetadata(m, v76)
																		mBase = m.M
																		v84 = m.ExcPending
																		if v84 != 0 {
																			return int32(0)
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v42)+16)) = int32(2)
																			v87 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
																			*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = int32(1)
																			F_SPI_connect_ext(m, int32(0))
																			mBase = m.M
																			v92 = m.ExcPending
																			if v92 != 0 {
																				return int32(0)
																			} else {
																				v93 = int32(_a_F_connectby_text_4)
																				v94 = *(*int32)(unsafe.Add(mBase, _c_F_connectby_text[0]))
																				*(*int32)(unsafe.Add(mBase, _c_F_connectby_text[0])) = v73
																				v103 = *(*int32)(unsafe.Add(mBase, _c_F_connectby_text[1]))
																				v104 = F_tuplestore_begin_heap(m, int32(base.Ui32(v87&int32(4))>>(uint(int32(2))%32)), int32(0), v103)
																				mBase = m.M
																				v105 = m.ExcPending
																				if v105 != 0 {
																					return int32(0)
																				} else {
																					*(*int32)(unsafe.Add(mBase, _c_F_connectby_text[0])) = v94
																					v108 = int32(0)
																					F_build_tuplestore_recursively(m, v30, v35, v25, v108, v68, v40, v40, v108, v18+int32(12), v56, v79, v108, v83, v104)
																					mBase = m.M
																					v114 = m.ExcPending
																					if v114 != 0 {
																						return int32(0)
																					} else {
																						v115 = F_SPI_finish(m)
																						mBase = m.M
																						v116 = m.ExcPending
																						if v116 != 0 {
																							return int32(0)
																						} else {
																							*(*int32)(unsafe.Add(mBase, uint32(v42)+28)) = v76
																							*(*int32)(unsafe.Add(mBase, uint32(v42)+24)) = v104
																							*(*int32)(unsafe.Add(mBase, _c_F_connectby_text[0])) = v70
																							m.G0 = v18 + int32(16)
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
													} else {
														v66 = F_pstrdup(m, int32(_a_F_connectby_text_5))
														mBase = m.M
														v67 = m.ExcPending
														if v67 != 0 {
															return int32(0)
														} else {
															v68 = v66
															v69 = int32(_a_F_connectby_text_4)
															v70 = *(*int32)(unsafe.Add(mBase, _c_F_connectby_text[0]))
															v72 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
															v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)+16))
															*(*int32)(unsafe.Add(mBase, _c_F_connectby_text[0])) = v73
															v75 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
															v76 = F_CreateTupleDescCopy(m, v75)
															mBase = m.M
															v77 = m.ExcPending
															if v77 != 0 {
																return int32(0)
															} else {
																v79 = base.B2i32(v57 == int32(6))
																F_validateConnectbyTupleDesc(m, v76, v79, int32(0))
																mBase = m.M
																v82 = m.ExcPending
																if v82 != 0 {
																	return int32(0)
																} else {
																	v83 = F_TupleDescGetAttInMetadata(m, v76)
																	mBase = m.M
																	v84 = m.ExcPending
																	if v84 != 0 {
																		return int32(0)
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v42)+16)) = int32(2)
																		v87 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
																		*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = int32(1)
																		F_SPI_connect_ext(m, int32(0))
																		mBase = m.M
																		v92 = m.ExcPending
																		if v92 != 0 {
																			return int32(0)
																		} else {
																			v93 = int32(_a_F_connectby_text_4)
																			v94 = *(*int32)(unsafe.Add(mBase, _c_F_connectby_text[0]))
																			*(*int32)(unsafe.Add(mBase, _c_F_connectby_text[0])) = v73
																			v103 = *(*int32)(unsafe.Add(mBase, _c_F_connectby_text[1]))
																			v104 = F_tuplestore_begin_heap(m, int32(base.Ui32(v87&int32(4))>>(uint(int32(2))%32)), int32(0), v103)
																			mBase = m.M
																			v105 = m.ExcPending
																			if v105 != 0 {
																				return int32(0)
																			} else {
																				*(*int32)(unsafe.Add(mBase, _c_F_connectby_text[0])) = v94
																				v108 = int32(0)
																				F_build_tuplestore_recursively(m, v30, v35, v25, v108, v68, v40, v40, v108, v18+int32(12), v56, v79, v108, v83, v104)
																				mBase = m.M
																				v114 = m.ExcPending
																				if v114 != 0 {
																					return int32(0)
																				} else {
																					v115 = F_SPI_finish(m)
																					mBase = m.M
																					v116 = m.ExcPending
																					if v116 != 0 {
																						return int32(0)
																					} else {
																						*(*int32)(unsafe.Add(mBase, uint32(v42)+28)) = v76
																						*(*int32)(unsafe.Add(mBase, uint32(v42)+24)) = v104
																						*(*int32)(unsafe.Add(mBase, _c_F_connectby_text[0])) = v70
																						m.G0 = v18 + int32(16)
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
		}
	}
}
func F_consider_groupingsets_paths(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 float64) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v31 float64
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 float64
	_ = v40
	var v42 int32
	_ = v42
	var v46 float64
	_ = v46
	var v47 float64
	_ = v47
	var v50 float64
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
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
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 float64
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 float64
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v225 int32
	_ = v225
	var v246 int32
	_ = v246
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v338 int32
	_ = v338
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v361 int32
	_ = v361
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v409 int32
	_ = v409
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v491 int32
	_ = v491
	var v497 int32
	_ = v497
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v535 int32
	_ = v535
	var v540 int32
	_ = v540
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v582 int32
	_ = v582
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v621 int32
	_ = v621
	var v650 float64
	_ = v650
	var v651 int32
	_ = v651
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v671 int32
	_ = v671
	var v674 int32
	_ = v674
	var v676 int32
	_ = v676
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v707 int32
	_ = v707
	var v710 int32
	_ = v710
	var v712 int32
	_ = v712
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v744 int32
	_ = v744
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v749 int32
	_ = v749
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v759 int32
	_ = v759
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v766 float64
	_ = v766
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v777 float64
	_ = v777
	var v782 int32
	_ = v782
	var v785 int32
	_ = v785
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v798 float64
	_ = v798
	var v799 float64
	_ = v799
	var v802 float64
	_ = v802
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v811 float64
	_ = v811
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v845 int32
	_ = v845
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v854 float64
	_ = v854
	var v855 int32
	_ = v855
	var v856 int32
	_ = v856
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v860 int32
	_ = v860
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v866 float64
	_ = v866
	var v868 float64
	_ = v868
	var v873 int32
	_ = v873
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v888 int32
	_ = v888
	var v913 int32
	_ = v913
	var v915 int32
	_ = v915
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v923 int32
	_ = v923
	var v927 int32
	_ = v927
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v934 int32
	_ = v934
	var v935 int32
	_ = v935
	var v942 int32
	_ = v942
	var v978 int32
	_ = v978
	var v979 int32
	_ = v979
	var v982 int32
	_ = v982
	var v1030 int32
	_ = v1030
	var v1053 int32
	_ = v1053
	var v1059 int32
	_ = v1059
	var v1087 int32
	_ = v1087
	var v1089 int32
	_ = v1089
	var v1090 float64
	_ = v1090
	var v1091 int32
	_ = v1091
	var v1094 int32
	_ = v1094
	var v1095 float64
	_ = v1095
	var v1101 int32
	_ = v1101
	var v1102 int32
	_ = v1102
	var v1106 int32
	_ = v1106
	var v1109 int32
	_ = v1109
	var v1112 int32
	_ = v1112
	var v1116 int32
	_ = v1116
	var v1117 int32
	_ = v1117
	var v1118 int32
	_ = v1118
	var v1125 int32
	_ = v1125
	var v1127 int32
	_ = v1127
	var v1128 int32
	_ = v1128
	var v1134 int32
	_ = v1134
	var v1135 int32
	_ = v1135
	var v1136 int32
	_ = v1136
	var v1137 int32
	_ = v1137
	var v1150 int32
	_ = v1150
	var v1175 int32
	_ = v1175
	var v1178 int32
	_ = v1178
	var v1181 int32
	_ = v1181
	var v1182 int32
	_ = v1182
	var v1195 int32
	_ = v1195
	var v1249 int32
	_ = v1249
	var v1261 int32
	_ = v1261
	var v1283 int32
	_ = v1283
	var v1284 int32
	_ = v1284
	var v1286 float64
	_ = v1286
	var v1323 int32
	_ = v1323
	var v1358 int32
	_ = v1358
	var v1360 int32
	_ = v1360
	var v1361 int32
	_ = v1361
	var v1362 int32
	_ = v1362
	var v1365 int32
	_ = v1365
	var v1366 int32
	_ = v1366
	var v1374 int32
	_ = v1374
	var v1375 int32
	_ = v1375
	var v1377 int32
	_ = v1377
	var v1382 int32
	_ = v1382
	var v1384 int32
	_ = v1384
	var v1397 int32
	_ = v1397
	var v1398 int32
	_ = v1398
	var v1399 int32
	_ = v1399
	var v1400 int32
	_ = v1400
	var v1401 int32
	_ = v1401
	var v1403 int32
	_ = v1403
	var v1406 int32
	_ = v1406
	var v1407 int32
	_ = v1407
	var v1408 int32
	_ = v1408
	var v1411 int32
	_ = v1411
	var v1415 int32
	_ = v1415
	var v1416 int32
	_ = v1416
	var v1417 int32
	_ = v1417
	var v1420 int32
	_ = v1420
	var v1428 int32
	_ = v1428
	var v1432 int32
	_ = v1432
	var v1438 int32
	_ = v1438
	var v1439 int32
	_ = v1439
	var v1456 int32
	_ = v1456
	var v1460 int32
	_ = v1460
	var v1461 int32
	_ = v1461
	var v1464 int32
	_ = v1464
	var v1465 int32
	_ = v1465
	var v1466 int32
	_ = v1466
	var v1467 int32
	_ = v1467
	var v1468 int32
	_ = v1468
	var v1471 int32
	_ = v1471
	var v1472 int32
	_ = v1472
	var v1475 int32
	_ = v1475
	var v1476 int32
	_ = v1476
	var v1477 int32
	_ = v1477
	var v1478 int32
	_ = v1478
	var v1479 int32
	_ = v1479
	var v1481 int32
	_ = v1481
	var v1482 int32
	_ = v1482
	var v1484 int32
	_ = v1484
	var v1485 int32
	_ = v1485
	var v1486 int32
	_ = v1486
	var v1489 int32
	_ = v1489
	var v1490 int32
	_ = v1490
	var v1498 int32
	_ = v1498
	var v1499 int32
	_ = v1499
	var v1501 int32
	_ = v1501
	var v1506 int32
	_ = v1506
	var v1508 int32
	_ = v1508
	var v1516 int32
	_ = v1516
	var v1521 int32
	_ = v1521
	var v1522 int32
	_ = v1522
	var v1523 int32
	_ = v1523
	var v1526 int32
	_ = v1526
	var v1527 int32
	_ = v1527
	var v1544 int32
	_ = v1544
	var v1554 int32
	_ = v1554
	var v1562 int32
	_ = v1562
	var v1566 int32
	_ = v1566
	var v1568 int32
	_ = v1568
	var v1569 int32
	_ = v1569
	var v1572 int32
	_ = v1572
	var v1573 int32
	_ = v1573
	var v1574 int32
	_ = v1574
	var v1581 int32
	_ = v1581
	var v1582 int32
	_ = v1582
	var v1584 int32
	_ = v1584
	var v1585 int32
	_ = v1585
	var v1588 int32
	_ = v1588
	var v1589 int32
	_ = v1589
	var v1600 int32
	_ = v1600
	var v1624 int32
	_ = v1624
	var v1625 int32
	_ = v1625
	var v1628 int32
	_ = v1628
	var v1629 int32
	_ = v1629
	var v1635 int32
	_ = v1635
	var v1636 int32
	_ = v1636
	var v1673 int32
	_ = v1673
	var v1675 int32
	_ = v1675
	var v1682 int32
	_ = v1682
	var v1688 int32
	_ = v1688
	var v1710 int32
	_ = v1710
	var v1711 int32
	_ = v1711
	var v1715 int32
	_ = v1715
	var v1716 int32
	_ = v1716
	var v1719 int32
	_ = v1719
	var v1720 int32
	_ = v1720
	var v1726 int32
	_ = v1726
	var v1731 int32
	_ = v1731
	var v1755 int32
	_ = v1755
	var v1756 int32
	_ = v1756
	var v1759 int32
	_ = v1759
	var v1763 int32
	_ = v1763
	var v1764 int32
	_ = v1764
	var v1765 int32
	_ = v1765
	var v1767 int32
	_ = v1767
	var v1768 int32
	_ = v1768
	var v1773 int32
	_ = v1773
	var v1802 int32
	_ = v1802
	var v1803 int32
	_ = v1803
	var v1805 int32
	_ = v1805
	var v1806 int32
	_ = v1806
	var v1812 int32
	_ = v1812
	var v1841 float64
	_ = v1841
	var v1842 int32
	_ = v1842
	var v1845 int32
	_ = v1845
	var v1846 int32
	_ = v1846
	var v1848 int32
	_ = v1848
	var v1849 int32
	_ = v1849
	var v1865 int32
	_ = v1865
	var v1885 int32
	_ = v1885
	var v1887 int32
	_ = v1887
	var v1888 int32
	_ = v1888
	var v1890 int32
	_ = v1890
	var v1891 int32
	_ = v1891
	var v1892 int32
	_ = v1892
	var v1893 int32
	_ = v1893
	var v1896 int32
	_ = v1896
	var v1897 int32
	_ = v1897
	var v1908 int32
	_ = v1908
	var v1913 int32
	_ = v1913
	var v1923 int32
	_ = v1923
	var v1924 int32
	_ = v1924
	var v1926 int32
	_ = v1926
	var v1927 int32
	_ = v1927
	var v1928 int32
	_ = v1928
	var v1946 int32
	_ = v1946
	var v1961 int32
	_ = v1961
	var v1963 int32
	_ = v1963
	var v1981 int32
	_ = v1981
	v9 = int32(0)
	v31 = float64(0)
	v33 = m.G0
	v35 = v33 - int32(32)
	m.G0 = v35
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v40 = *(*float64)(unsafe.Add(mBase, _c_F_consider_groupingsets_paths[0]))
	v42 = *(*int32)(unsafe.Add(mBase, _c_F_consider_groupingsets_paths[1]))
	v46 = base.F64_mul(base.F64_mul(v40, base.F64_convert_i32_s(v42)), float64(1024))
	v47 = float64(4.294967295e+09)
	if base.F64_lt(v46, v47) != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	if l3 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L2:
	;
	v50 = v46
	goto L4
L3:
	;
	v50 = v47
	goto L4
L4:
	;
	v51 = base.I32_trunc_sat_f64_u(v50)
	goto L1
L5:
	;
	m.G0 = v1981 + int32(32)
	return
L6:
	;
	F_add_path(m, l1, v1961)
	mBase = m.M
	v1963 = m.ExcPending
	if v1963 != 0 {
		goto L43
	} else {
		goto L261
	}
L7:
	;
	v55 = int32(0)
	if v52 == v55 {
		v134 = v9
		v135 = v31
		v136 = v55
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	if v52 == int32(0) {
		v1981 = v35
		goto L5
	} else {
		goto L115
	}
L10:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+332))
	if v138 != 0 {
		goto L36
	} else {
		goto L37
	}
L11:
	;
	v58 = int32(0)
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v52)+12))
	if v59 == v58 {
		v134 = v9
		v135 = v31
		v136 = v58
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l2)+64))
	if v62 == v63 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	if v116 == int32(0) {
		v134 = v9
		v135 = v31
		v136 = v59
		goto L10
	} else {
		goto L31
	}
L14:
	;
	v116 = int32(1)
	goto L13
L15:
	;
	goto L16
L16:
	;
	v72 = int32(0)
	goto L18
L17:
	;
	v116 = v108
	goto L13
L18:
	;
	v76 = int32(0)
	if v62 == v76 {
		v86 = v76
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v108 = int32(0)
	goto L17
L20:
	;
	if v63 != 0 {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
	if v80 <= v72 {
		v86 = int32(0)
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v62)+12))
	v86 = v82 + v72<<(uint(int32(2))%32)
	goto L20
L23:
	;
	v92 = base.B2i32(v86 == int32(0))
	if v86 == int32(0) {
		v108 = v92
		goto L17
	} else {
		goto L28
	}
L24:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	if v72 < v87 {
		goto L23
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v116 = base.B2i32(v86 == int32(0))
	goto L13
L27:
	;
	goto L26
L28:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v63)+12))
	if v95 == int32(0) {
		v108 = v92
		goto L17
	} else {
		goto L29
	}
L29:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v72<<(uint(int32(2))%32)+v95)))
	if v102 == v104 {
		v72 = v72 + int32(1)
		goto L18
	} else {
		goto L30
	}
L30:
	;
	goto L19
L31:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	v120 = *(*float64)(unsafe.Add(mBase, uint32(v119)+16))
	v122 = v59 + int32(4)
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v124)+12))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v124)+4))
	if base.Ui32(v122) < base.Ui32(v125+v126<<(uint(int32(2))%32)) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v131 = v122
	goto L34
L33:
	;
	v131 = int32(0)
	goto L34
L34:
	;
	v134 = v119
	v135 = v120
	v136 = v131
	goto L10
L35:
	;
	if base.F64_gt(base.F64_mul(base.F64_sub(l7, v135), base.F64_convert_i32_u(v145)), base.F64_convert_i32_u(v51)) != 0 {
		goto L39
	} else {
		goto L40
	}
L36:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)+4))
	v141 = v139
	goto L38
L37:
	;
	v141 = int32(0)
	goto L38
L38:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)+32))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l6)+32))
	v145 = F_hash_agg_entry_size(m, v141, v143, v144)
	mBase = m.M
	goto L35
L39:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	if v150 != 0 {
		v1981 = v35
		goto L5
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l5)+28))
	v152 = F_list_copy(m, v151)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	goto L41
L43:
	;
	return
L44:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	if v136 != 0 {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	if v225 == int32(0) {
		v1981 = v35
		goto L5
	} else {
		goto L57
	}
L46:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v154)+4))
	if v163 <= v162 {
		v225 = v152
		goto L45
	} else {
		goto L51
	}
L47:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v154)+12))
	v162 = (v136 - v155) >> (uint(int32(2)) % 32)
	goto L46
L48:
	;
	goto L49
L49:
	;
	if v154 == int32(0) {
		v225 = v152
		goto L45
	} else {
		goto L50
	}
L50:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v154)+4))
	v162 = v161
	goto L46
L51:
	;
	v173 = v162
	v178 = v152
	goto L52
L52:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v154)+12))
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v197+v173<<(uint(int32(2))%32))))
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201)+24)))
	if v202 != int32(1) {
		v1981 = v35
		goto L5
	} else {
		goto L54
	}
L53:
	;
	v225 = v206
	goto L45
L54:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v201)+12))
	v206 = F_list_concat(m, v178, v205)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L43
	} else {
		goto L55
	}
L55:
	;
	v209 = v173 + int32(1)
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v154)+4))
	if v209 < v210 {
		v173 = v209
		v178 = v206
		goto L52
	} else {
		goto L56
	}
L56:
	;
	goto L53
L57:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v225)+4))
	if int32(0) < v246 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v263 = v9
	v264 = v9
	v267 = v9
	v269 = v9
	goto L61
L59:
	;
	v707 = v9
	v710 = v9
	v712 = v9
	goto L60
L60:
	;
	if v707 == int32(0) {
		v1981 = v35
		goto L5
	} else {
		goto L104
	}
L61:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v225)+12))
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v281+v263<<(uint(int32(2))%32))))
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v285)+4))
	if v286 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L62:
	;
	v707 = v671
	v710 = v674
	v712 = v676
	goto L60
L63:
	;
	v689 = v263 + int32(1)
	v690 = *(*int32)(unsafe.Add(mBase, uint32(v225)+4))
	if v689 < v690 {
		v263 = v689
		v264 = v671
		v267 = v674
		v269 = v676
		goto L61
	} else {
		goto L103
	}
L64:
	;
	v289 = F_lappend(m, v269, v285)
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L43
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v295 = F_palloc0(m, int32(32))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L43
	} else {
		goto L69
	}
L67:
	;
	v292 = F_lappend(m, v267, int32(0))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L43
	} else {
		goto L68
	}
L68:
	;
	v671 = v264
	v674 = v292
	v676 = v289
	goto L63
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v295))) = int32(309)
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v286)+4))
	if v299 <= int32(0) {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v295)+4)) = v361
	*(*int32)(unsafe.Add(mBase, uint32(v35)+16)) = v285
	*(*int32)(unsafe.Add(mBase, uint32(v35)+28)) = v285
	v390 = F_list_make1_impl(m, int32(1), v35+int32(16))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L43
	} else {
		goto L79
	}
L71:
	;
	v361 = int32(0)
	goto L70
L72:
	;
	goto L73
L73:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v304 = int32(0)
	v314 = v304
	v315 = v304
	goto L74
L74:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v286)+12))
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v338+v314<<(uint(int32(2))%32))))
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v303)+100))
	v344 = F_get_sortgroupref_clause(m, v342, v343)
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L43
	} else {
		goto L76
	}
L75:
	;
	v361 = v346
	goto L70
L76:
	;
	v346 = F_lappend(m, v315, v344)
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L43
	} else {
		goto L77
	}
L77:
	;
	v349 = v314 + int32(1)
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v286)+4))
	if v349 < v350 {
		v314 = v349
		v315 = v346
		goto L74
	} else {
		goto L78
	}
L78:
	;
	goto L75
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v295)+12)) = v390
	v393 = *(*int32)(unsafe.Add(mBase, uint32(l5)+32))
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v295)+4))
	if v394 == int32(0) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	if v390 == int32(0) {
		goto L87
	} else {
		goto L88
	}
L81:
	;
	v397 = int32(0)
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v394)+4))
	if v398 <= v397 {
		goto L80
	} else {
		goto L82
	}
L82:
	;
	v409 = v397
	goto L83
L83:
	;
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v394)+12))
	v434 = int32(2)
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v433+v409<<(uint(v434)%32))))
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v437)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v393+v438<<(uint(v434)%32)))) = v409
	v444 = v409 + int32(1)
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v394)+4))
	if v444 < v445 {
		v409 = v444
		goto L83
	} else {
		goto L85
	}
L84:
	;
	goto L80
L85:
	;
	goto L84
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v295)+8)) = v621
	v650 = *(*float64)(unsafe.Add(mBase, uint32(v285)+8))
	v651 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v295)+24)) = uint16(v651)
	*(*float64)(unsafe.Add(mBase, uint32(v295)+16)) = v650
	v654 = F_lappend(m, v264, v295)
	mBase = m.M
	v655 = m.ExcPending
	if v655 != 0 {
		goto L43
	} else {
		goto L102
	}
L87:
	;
	v621 = int32(0)
	goto L86
L88:
	;
	goto L89
L89:
	;
	v482 = int32(0)
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v390)+4))
	if v484 <= v482 {
		v621 = v482
		goto L86
	} else {
		goto L90
	}
L90:
	;
	v491 = v482
	v497 = v482
	goto L91
L91:
	;
	v519 = int32(0)
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v390)+12))
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v520+v497<<(uint(int32(2))%32))))
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v524)+4))
	if v525 == v519 {
		v582 = v519
		goto L93
	} else {
		goto L94
	}
L92:
	;
	v621 = v611
	goto L86
L93:
	;
	v611 = F_lappend(m, v491, v582)
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
		goto L43
	} else {
		goto L100
	}
L94:
	;
	v528 = int32(0)
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v525)+4))
	if v529 <= v528 {
		v582 = v519
		goto L93
	} else {
		goto L95
	}
L95:
	;
	v535 = v519
	v540 = v528
	goto L96
L96:
	;
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v525)+12))
	v565 = int32(2)
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v564+v540<<(uint(v565)%32))))
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v393+v568<<(uint(v565)%32))))
	v573 = F_lappend_int(m, v535, v572)
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L43
	} else {
		goto L98
	}
L97:
	;
	v582 = v573
	goto L93
L98:
	;
	v576 = v540 + int32(1)
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v525)+4))
	if v576 < v577 {
		v535 = v573
		v540 = v576
		goto L96
	} else {
		goto L99
	}
L99:
	;
	goto L97
L100:
	;
	v614 = v497 + int32(1)
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v390)+4))
	if v614 < v615 {
		v491 = v611
		v497 = v614
		goto L91
	} else {
		goto L101
	}
L101:
	;
	goto L92
L102:
	;
	v671 = v654
	v674 = v267
	v676 = v269
	goto L63
L103:
	;
	goto L62
L104:
	;
	if v134 == int32(0) {
		goto L106
	} else {
		goto L107
	}
L105:
	;
	v752 = *(*int32)(unsafe.Add(mBase, uint32(v37)+112))
	v753 = F_create_groupingsets_path(m, l0, l1, l2, v752, v749, v751, l6)
	mBase = m.M
	v754 = m.ExcPending
	if v754 != 0 {
		goto L43
	} else {
		goto L114
	}
L106:
	;
	if v710 == int32(0) {
		goto L109
	} else {
		goto L110
	}
L107:
	;
	v744 = v134
	goto L108
L108:
	;
	v746 = F_lappend(m, v707, v744)
	mBase = m.M
	v747 = m.ExcPending
	if v747 != 0 {
		goto L43
	} else {
		goto L113
	}
L109:
	;
	v749 = int32(2)
	v751 = v707
	goto L105
L110:
	;
	goto L111
L111:
	;
	v732 = F_palloc0(m, int32(32))
	mBase = m.M
	v733 = m.ExcPending
	if v733 != 0 {
		goto L43
	} else {
		goto L112
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v732)+12)) = v712
	*(*int64)(unsafe.Add(mBase, uint32(v732))) = int64(309)
	*(*int32)(unsafe.Add(mBase, uint32(v732)+8)) = v710
	v738 = *(*int32)(unsafe.Add(mBase, uint32(v710)+4))
	v739 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v732)+24)) = uint16(v739)
	*(*float64)(unsafe.Add(mBase, uint32(v732)+16)) = base.F64_convert_i32_s(v738)
	v744 = v732
	goto L108
L113:
	;
	v749 = int32(3)
	v751 = v746
	goto L105
L114:
	;
	v1946 = v35
	v1961 = v753
	goto L6
L115:
	;
	if l4 == int32(0) {
		v1891 = l0
		v1892 = l1
		v1893 = l2
		v1896 = l5
		v1897 = l6
		v1908 = v35
		v1913 = v37
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v1923 = *(*int32)(unsafe.Add(mBase, uint32(v1896)+28))
	if v1923 != 0 {
		v1981 = v1908
		goto L5
	} else {
		goto L259
	}
L117:
	;
	v759 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5)+16)))
	if v759 != int32(1) {
		v1891 = l0
		v1892 = l1
		v1893 = l2
		v1896 = l5
		v1897 = l6
		v1908 = v35
		v1913 = v37
		goto L116
	} else {
		goto L118
	}
L118:
	;
	v762 = *(*int32)(unsafe.Add(mBase, uint32(l5)+28))
	v763 = F_list_copy(m, v762)
	mBase = m.M
	v764 = m.ExcPending
	if v764 != 0 {
		goto L43
	} else {
		goto L119
	}
L119:
	;
	v766 = *(*float64)(unsafe.Add(mBase, uint32(l5)+8))
	v767 = *(*int32)(unsafe.Add(mBase, uint32(l0)+332))
	if v767 != 0 {
		goto L122
	} else {
		goto L123
	}
L120:
	;
	v1516 = int32(0)
	if v1498|base.B2i32(v1499 == v1516) == v1516 {
		goto L221
	} else {
		goto L222
	}
L121:
	;
	v777 = base.F64_sub(base.F64_convert_i32_u(v51), base.F64_mul(v766, base.F64_convert_i32_u(v774)))
	if base.F64_gt(v777, float64(0)) == int32(0) {
		v1484 = l0
		v1485 = l1
		v1486 = l2
		v1489 = l5
		v1490 = l6
		v1498 = v9
		v1499 = v763
		v1501 = v35
		v1506 = v37
		v1508 = v9
		goto L120
	} else {
		goto L125
	}
L122:
	;
	v768 = *(*int32)(unsafe.Add(mBase, uint32(v767)+4))
	v770 = v768
	goto L124
L123:
	;
	v770 = int32(0)
	goto L124
L124:
	;
	v771 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v772 = *(*int32)(unsafe.Add(mBase, uint32(v771)+32))
	v773 = *(*int32)(unsafe.Add(mBase, uint32(l6)+32))
	v774 = F_hash_agg_entry_size(m, v770, v772, v773)
	mBase = m.M
	goto L121
L125:
	;
	v782 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	if v782 == int32(0) {
		v1484 = l0
		v1485 = l1
		v1486 = l2
		v1489 = l5
		v1490 = l6
		v1498 = v9
		v1499 = v763
		v1501 = v35
		v1506 = v37
		v1508 = v9
		goto L120
	} else {
		goto L126
	}
L126:
	;
	v785 = *(*int32)(unsafe.Add(mBase, uint32(v782)+4))
	if v785 < int32(2) {
		v1484 = l0
		v1485 = l1
		v1486 = l2
		v1489 = l5
		v1490 = l6
		v1498 = v9
		v1499 = v763
		v1501 = v35
		v1506 = v37
		v1508 = v9
		goto L120
	} else {
		goto L127
	}
L127:
	;
	v790 = F_palloc(m, v785<<(uint(int32(2))%32))
	mBase = m.M
	v791 = m.ExcPending
	if v791 != 0 {
		goto L43
	} else {
		goto L128
	}
L128:
	;
	v792 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	if v792 == int32(0) {
		v1484 = l0
		v1485 = l1
		v1486 = l2
		v1489 = l5
		v1490 = l6
		v1498 = v9
		v1499 = v763
		v1501 = v35
		v1506 = v37
		v1508 = v9
		goto L120
	} else {
		goto L129
	}
L129:
	;
	v798 = base.F64_div(v777, base.F64_mul(base.F64_convert_i32_u(v785), float64(20)))
	v799 = float64(1)
	if base.F64_gt(v798, v799) != 0 {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v802 = v798
	goto L132
L131:
	;
	v802 = v799
	goto L132
L132:
	;
	v805 = base.I32_trunc_sat_f64_s(base.F64_floor(base.F64_div(v777, v802)))
	v806 = *(*int32)(unsafe.Add(mBase, uint32(v792)+4))
	if int32(2) <= v806 {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v811 = base.F64_add(base.F64_convert_i32_s(v805), float64(1))
	v821 = int32(1)
	v822 = v9
	goto L136
L134:
	;
	v888 = v9
	goto L135
L135:
	;
	if v888 <= int32(0) {
		v1484 = l0
		v1485 = l1
		v1486 = l2
		v1489 = l5
		v1490 = l6
		v1498 = v9
		v1499 = v763
		v1501 = v35
		v1506 = v37
		v1508 = v9
		goto L120
	} else {
		goto L149
	}
L136:
	;
	v845 = *(*int32)(unsafe.Add(mBase, uint32(v792)+12))
	v849 = *(*int32)(unsafe.Add(mBase, uint32(v845+v821<<(uint(int32(2))%32))))
	v850 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v849)+24)))
	if v850 != 0 {
		goto L138
	} else {
		goto L139
	}
L137:
	;
	v888 = v873
	goto L135
L138:
	;
	v854 = *(*float64)(unsafe.Add(mBase, uint32(v849)+16))
	v855 = *(*int32)(unsafe.Add(mBase, uint32(l0)+332))
	if v855 != 0 {
		goto L142
	} else {
		goto L143
	}
L139:
	;
	v873 = v822
	goto L140
L140:
	;
	v876 = v821 + int32(1)
	v877 = *(*int32)(unsafe.Add(mBase, uint32(v792)+4))
	if v876 < v877 {
		v821 = v876
		v822 = v873
		goto L136
	} else {
		goto L148
	}
L141:
	;
	v866 = base.F64_floor(base.F64_div(base.F64_mul(v854, base.F64_convert_i32_u(v862)), v802))
	if base.F64_gt(v811, v866) != 0 {
		goto L145
	} else {
		goto L146
	}
L142:
	;
	v856 = *(*int32)(unsafe.Add(mBase, uint32(v855)+4))
	v858 = v856
	goto L144
L143:
	;
	v858 = int32(0)
	goto L144
L144:
	;
	v859 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v860 = *(*int32)(unsafe.Add(mBase, uint32(v859)+32))
	v861 = *(*int32)(unsafe.Add(mBase, uint32(l6)+32))
	v862 = F_hash_agg_entry_size(m, v858, v860, v861)
	mBase = m.M
	goto L141
L145:
	;
	v868 = v866
	goto L147
L146:
	;
	v868 = v811
	goto L147
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v790+v822<<(uint(int32(2))%32)))) = base.I32_trunc_sat_f64_s(v868)
	v873 = v822 + int32(1)
	goto L140
L148:
	;
	goto L137
L149:
	;
	v913 = int32(0)
	v915 = *(*int32)(unsafe.Add(mBase, _c_F_consider_groupingsets_paths[2]))
	v920 = F_AllocSetContextCreateInternal(m, v915, int32(_a_F_consider_groupingsets_paths_0), v913, int32(1024), int32(_a_F_consider_groupingsets_paths_1))
	mBase = m.M
	v921 = m.ExcPending
	if v921 != 0 {
		goto L43
	} else {
		goto L150
	}
L150:
	;
	v922 = int32(_a_F_consider_groupingsets_paths_2)
	v923 = *(*int32)(unsafe.Add(mBase, _c_F_consider_groupingsets_paths[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_consider_groupingsets_paths[2])) = v920
	v927 = v805 + int32(1)
	v930 = F_palloc(m, v927<<(uint(int32(3))%32))
	mBase = m.M
	v931 = m.ExcPending
	if v931 != 0 {
		goto L43
	} else {
		goto L151
	}
L151:
	;
	v934 = F_palloc(m, v927<<(uint(int32(2))%32))
	mBase = m.M
	v935 = m.ExcPending
	if v935 != 0 {
		goto L43
	} else {
		goto L152
	}
L152:
	;
	if int32(0) <= v805 {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	v942 = v913
	goto L156
L154:
	;
	goto L155
L155:
	;
	if int32(0) < v888 {
		goto L160
	} else {
		goto L161
	}
L156:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v930+v942<<(uint(int32(3))%32)))) = int64(0)
	v978 = F_bms_make_singleton(m, v888)
	mBase = m.M
	v979 = m.ExcPending
	if v979 != 0 {
		goto L43
	} else {
		goto L158
	}
L157:
	;
	goto L155
L158:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v934+v942<<(uint(int32(2))%32)))) = v978
	v982 = v942 + int32(1)
	if v982 <= v805 {
		v942 = v982
		goto L156
	} else {
		goto L159
	}
L159:
	;
	goto L157
L160:
	;
	v1030 = v9
	goto L163
L161:
	;
	v1360 = l0
	v1361 = l1
	v1362 = l2
	v1365 = l5
	v1366 = l6
	v1374 = v9
	v1375 = v763
	v1377 = v35
	v1382 = v37
	v1384 = v9
	goto L162
L162:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_consider_groupingsets_paths[2])) = v923
	v1397 = *(*int32)(unsafe.Add(mBase, uint32(v934+v805<<(uint(int32(2))%32))))
	v1398 = F_bms_copy(m, v1397)
	mBase = m.M
	v1399 = m.ExcPending
	if v1399 != 0 {
		goto L43
	} else {
		goto L198
	}
L163:
	;
	v1053 = *(*int32)(unsafe.Add(mBase, uint32(v790+v1030<<(uint(int32(2))%32))))
	if v1053 <= v805 {
		goto L165
	} else {
		goto L166
	}
L164:
	;
	v1360 = l0
	v1361 = l1
	v1362 = l2
	v1365 = l5
	v1366 = l6
	v1374 = v9
	v1375 = v763
	v1377 = v35
	v1382 = v37
	v1384 = v9
	goto L162
L165:
	;
	v1059 = v805
	goto L168
L166:
	;
	goto L167
L167:
	;
	v1358 = v1030 + int32(1)
	if v1358 != v888 {
		v1030 = v1358
		goto L163
	} else {
		goto L197
	}
L168:
	;
	v1087 = int32(3)
	v1089 = v930 + v1059<<(uint(v1087)%32)
	v1090 = *(*float64)(unsafe.Add(mBase, uint32(v1089)))
	v1091 = v1059 - v1053
	v1094 = v930 + v1091<<(uint(v1087)%32)
	v1095 = *(*float64)(unsafe.Add(mBase, uint32(v1094)))
	if base.F64_le(v1090, base.F64_add(v1095, float64(1))) != 0 {
		goto L170
	} else {
		goto L171
	}
L169:
	;
	goto L167
L170:
	;
	v1101 = v934 + v1059<<(uint(int32(2))%32)
	v1102 = *(*int32)(unsafe.Add(mBase, uint32(v1101)))
	if v1053 != 0 {
		goto L173
	} else {
		goto L174
	}
L171:
	;
	goto L172
L172:
	;
	v1323 = v1059 - int32(1)
	if v1053 <= v1323 {
		v1059 = v1323
		goto L168
	} else {
		goto L196
	}
L173:
	;
	v1106 = *(*int32)(unsafe.Add(mBase, uint32(v934+v1091<<(uint(int32(2))%32))))
	if v1102 == int32(0) {
		goto L178
	} else {
		goto L179
	}
L174:
	;
	v1261 = v1102
	goto L175
L175:
	;
	v1283 = F_bms_add_member(m, v1261, v1030)
	mBase = m.M
	v1284 = m.ExcPending
	if v1284 != 0 {
		goto L43
	} else {
		goto L195
	}
L176:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1101))) = v1249
	v1261 = v1249
	goto L175
L177:
	;
	v1249 = v1195
	goto L176
L178:
	;
	v1109 = int32(0)
	if v1106 == v1109 {
		v1249 = v1109
		goto L176
	} else {
		goto L181
	}
L179:
	;
	goto L180
L180:
	;
	if v1106 == int32(0) {
		goto L184
	} else {
		goto L185
	}
L181:
	;
	v1112 = *(*int32)(unsafe.Add(mBase, uint32(v1106)+4))
	v1116 = v1112<<(uint(int32(2))%32) + int32(8)
	v1117 = F_palloc(m, v1116)
	mBase = m.M
	v1118 = m.ExcPending
	if v1118 != 0 {
		goto L43
	} else {
		goto L182
	}
L182:
	;
	if v1116 == int32(0) {
		v1195 = v1117
		goto L177
	} else {
		goto L183
	}
L183:
	;
	base.MemoryCopy(m, v1117, v1106, v1116)
	v1249 = v1117
	goto L176
L184:
	;
	F_pfree(m, v1102)
	mBase = m.M
	v1125 = m.ExcPending
	if v1125 != 0 {
		goto L43
	} else {
		goto L187
	}
L185:
	;
	goto L186
L186:
	;
	v1127 = *(*int32)(unsafe.Add(mBase, uint32(v1106)+4))
	v1128 = *(*int32)(unsafe.Add(mBase, uint32(v1102)+4))
	if v1128 < v1127 {
		goto L188
	} else {
		goto L189
	}
L187:
	;
	v1249 = int32(0)
	goto L176
L188:
	;
	v1134 = F_repalloc(m, v1102, v1127<<(uint(int32(2))%32)+int32(8))
	mBase = m.M
	v1135 = m.ExcPending
	if v1135 != 0 {
		goto L43
	} else {
		goto L191
	}
L189:
	;
	v1136 = v1102
	goto L190
L190:
	;
	v1137 = int32(8)
	v1150 = int32(0)
	goto L192
L191:
	;
	v1136 = v1134
	goto L190
L192:
	;
	v1175 = v1150 << (uint(int32(2)) % 32)
	v1178 = *(*int32)(unsafe.Add(mBase, uint32(v1175+(v1106+v1137))))
	*(*int32)(unsafe.Add(mBase, uint32(v1136+v1137+v1175))) = v1178
	v1181 = v1150 + int32(1)
	v1182 = *(*int32)(unsafe.Add(mBase, uint32(v1106)+4))
	if v1181 < v1182 {
		v1150 = v1181
		goto L192
	} else {
		goto L194
	}
L193:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1136)+4)) = v1182
	v1195 = v1136
	goto L177
L194:
	;
	goto L193
L195:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1101))) = v1283
	v1286 = *(*float64)(unsafe.Add(mBase, uint32(v1094)))
	*(*float64)(unsafe.Add(mBase, uint32(v1089))) = base.F64_add(v1286, float64(1))
	goto L172
L196:
	;
	goto L169
L197:
	;
	goto L164
L198:
	;
	v1400 = F_bms_del_member(m, v1398, v888)
	mBase = m.M
	v1401 = m.ExcPending
	if v1401 != 0 {
		goto L43
	} else {
		goto L199
	}
L199:
	;
	F_MemoryContextDelete(m, v920)
	mBase = m.M
	v1403 = m.ExcPending
	if v1403 != 0 {
		goto L43
	} else {
		goto L200
	}
L200:
	;
	if v1400 == int32(0) {
		v1484 = v1360
		v1485 = v1361
		v1486 = v1362
		v1489 = v1365
		v1490 = v1366
		v1498 = v1374
		v1499 = v1375
		v1501 = v1377
		v1506 = v1382
		v1508 = v1384
		goto L120
	} else {
		goto L201
	}
L201:
	;
	v1406 = *(*int32)(unsafe.Add(mBase, uint32(v1365)))
	v1407 = *(*int32)(unsafe.Add(mBase, uint32(v1406)+12))
	v1408 = *(*int32)(unsafe.Add(mBase, uint32(v1407)))
	*(*int32)(unsafe.Add(mBase, uint32(v1377)+12)) = v1408
	*(*int32)(unsafe.Add(mBase, uint32(v1377)+24)) = v1408
	v1411 = int32(1)
	v1415 = F_list_make1_impl(m, v1411, v1377+int32(12))
	mBase = m.M
	v1416 = m.ExcPending
	if v1416 != 0 {
		goto L43
	} else {
		goto L202
	}
L202:
	;
	v1417 = *(*int32)(unsafe.Add(mBase, uint32(v1365)))
	if v1417 == int32(0) {
		v1484 = v1360
		v1485 = v1361
		v1486 = v1362
		v1489 = v1365
		v1490 = v1366
		v1498 = v1415
		v1499 = v1375
		v1501 = v1377
		v1506 = v1382
		v1508 = v1384
		goto L120
	} else {
		goto L203
	}
L203:
	;
	v1420 = *(*int32)(unsafe.Add(mBase, uint32(v1417)+4))
	if v1420 < int32(2) {
		v1484 = v1360
		v1485 = v1361
		v1486 = v1362
		v1489 = v1365
		v1490 = v1366
		v1498 = v1415
		v1499 = v1375
		v1501 = v1377
		v1506 = v1382
		v1508 = v1384
		goto L120
	} else {
		goto L204
	}
L204:
	;
	v1428 = int32(0)
	v1432 = v1411
	v1438 = v1415
	v1439 = v1375
	goto L205
L205:
	;
	v1456 = *(*int32)(unsafe.Add(mBase, uint32(v1417)+12))
	v1460 = *(*int32)(unsafe.Add(mBase, uint32(v1456+v1432<<(uint(int32(2))%32))))
	v1461 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1460)+24)))
	if v1461 == int32(1) {
		goto L208
	} else {
		goto L209
	}
L206:
	;
	v1484 = v1360
	v1485 = v1361
	v1486 = v1362
	v1489 = v1365
	v1490 = v1366
	v1498 = v1478
	v1499 = v1479
	v1501 = v1377
	v1506 = v1382
	v1508 = v1384
	goto L120
L207:
	;
	v1481 = v1432 + int32(1)
	v1482 = *(*int32)(unsafe.Add(mBase, uint32(v1417)+4))
	if v1481 < v1482 {
		v1428 = v1477
		v1432 = v1481
		v1438 = v1478
		v1439 = v1479
		goto L205
	} else {
		goto L218
	}
L208:
	;
	v1464 = F_bms_is_member(m, v1428, v1400)
	mBase = m.M
	v1465 = m.ExcPending
	if v1465 != 0 {
		goto L43
	} else {
		goto L211
	}
L209:
	;
	goto L210
L210:
	;
	v1475 = F_lappend(m, v1438, v1460)
	mBase = m.M
	v1476 = m.ExcPending
	if v1476 != 0 {
		goto L43
	} else {
		goto L217
	}
L211:
	;
	if v1464 != 0 {
		goto L212
	} else {
		goto L213
	}
L212:
	;
	v1466 = *(*int32)(unsafe.Add(mBase, uint32(v1460)+12))
	v1467 = F_list_concat(m, v1439, v1466)
	mBase = m.M
	v1468 = m.ExcPending
	if v1468 != 0 {
		goto L43
	} else {
		goto L215
	}
L213:
	;
	goto L214
L214:
	;
	v1471 = F_lappend(m, v1438, v1460)
	mBase = m.M
	v1472 = m.ExcPending
	if v1472 != 0 {
		goto L43
	} else {
		goto L216
	}
L215:
	;
	v1477 = v1428 + int32(1)
	v1478 = v1438
	v1479 = v1467
	goto L207
L216:
	;
	v1477 = v1428 + int32(1)
	v1478 = v1471
	v1479 = v1439
	goto L207
L217:
	;
	v1477 = v1428
	v1478 = v1475
	v1479 = v1439
	goto L207
L218:
	;
	goto L206
L219:
	;
	if v1865 == int32(0) {
		v1891 = v1484
		v1892 = v1485
		v1893 = v1486
		v1896 = v1489
		v1897 = v1490
		v1908 = v1501
		v1913 = v1506
		goto L116
	} else {
		goto L256
	}
L220:
	;
	v1527 = *(*int32)(unsafe.Add(mBase, uint32(v1499)+4))
	if v1527 <= int32(0) {
		v1865 = v1526
		goto L219
	} else {
		goto L226
	}
L221:
	;
	v1521 = *(*int32)(unsafe.Add(mBase, uint32(v1489)))
	v1522 = F_list_copy(m, v1521)
	mBase = m.M
	v1523 = m.ExcPending
	if v1523 != 0 {
		goto L43
	} else {
		goto L224
	}
L222:
	;
	goto L223
L223:
	;
	if v1499 == int32(0) {
		v1865 = v1498
		goto L219
	} else {
		goto L225
	}
L224:
	;
	v1526 = v1522
	goto L220
L225:
	;
	v1526 = v1498
	goto L220
L226:
	;
	v1544 = v1526
	v1554 = v1508
	goto L227
L227:
	;
	v1562 = *(*int32)(unsafe.Add(mBase, uint32(v1499)+12))
	v1566 = *(*int32)(unsafe.Add(mBase, uint32(v1562+v1554<<(uint(int32(2))%32))))
	v1568 = F_palloc0(m, int32(32))
	mBase = m.M
	v1569 = m.ExcPending
	if v1569 != 0 {
		goto L43
	} else {
		goto L229
	}
L228:
	;
	v1865 = v1845
	goto L219
L229:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1568))) = int32(309)
	v1572 = *(*int32)(unsafe.Add(mBase, uint32(v1566)+4))
	v1573 = F_preprocess_groupclause(m, v1484, v1572)
	mBase = m.M
	v1574 = m.ExcPending
	if v1574 != 0 {
		goto L43
	} else {
		goto L230
	}
L230:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1568)+4)) = v1573
	*(*int32)(unsafe.Add(mBase, uint32(v1501)+8)) = v1566
	*(*int32)(unsafe.Add(mBase, uint32(v1501)+20)) = v1566
	v1581 = F_list_make1_impl(m, int32(1), v1501+int32(8))
	mBase = m.M
	v1582 = m.ExcPending
	if v1582 != 0 {
		goto L43
	} else {
		goto L231
	}
L231:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1568)+12)) = v1581
	v1584 = *(*int32)(unsafe.Add(mBase, uint32(v1489)+32))
	v1585 = *(*int32)(unsafe.Add(mBase, uint32(v1568)+4))
	if v1585 == int32(0) {
		goto L232
	} else {
		goto L233
	}
L232:
	;
	if v1581 == int32(0) {
		goto L239
	} else {
		goto L240
	}
L233:
	;
	v1588 = int32(0)
	v1589 = *(*int32)(unsafe.Add(mBase, uint32(v1585)+4))
	if v1589 <= v1588 {
		goto L232
	} else {
		goto L234
	}
L234:
	;
	v1600 = v1588
	goto L235
L235:
	;
	v1624 = *(*int32)(unsafe.Add(mBase, uint32(v1585)+12))
	v1625 = int32(2)
	v1628 = *(*int32)(unsafe.Add(mBase, uint32(v1624+v1600<<(uint(v1625)%32))))
	v1629 = *(*int32)(unsafe.Add(mBase, uint32(v1628)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1584+v1629<<(uint(v1625)%32)))) = v1600
	v1635 = v1600 + int32(1)
	v1636 = *(*int32)(unsafe.Add(mBase, uint32(v1585)+4))
	if v1635 < v1636 {
		v1600 = v1635
		goto L235
	} else {
		goto L237
	}
L236:
	;
	goto L232
L237:
	;
	goto L236
L238:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1568)+8)) = v1812
	v1841 = *(*float64)(unsafe.Add(mBase, uint32(v1566)+8))
	v1842 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v1568)+24)) = uint16(v1842)
	*(*float64)(unsafe.Add(mBase, uint32(v1568)+16)) = v1841
	v1845 = F_lcons(m, v1568, v1544)
	mBase = m.M
	v1846 = m.ExcPending
	if v1846 != 0 {
		goto L43
	} else {
		goto L254
	}
L239:
	;
	v1812 = int32(0)
	goto L238
L240:
	;
	goto L241
L241:
	;
	v1673 = int32(0)
	v1675 = *(*int32)(unsafe.Add(mBase, uint32(v1581)+4))
	if v1675 <= v1673 {
		v1812 = v1673
		goto L238
	} else {
		goto L242
	}
L242:
	;
	v1682 = v1673
	v1688 = v1673
	goto L243
L243:
	;
	v1710 = int32(0)
	v1711 = *(*int32)(unsafe.Add(mBase, uint32(v1581)+12))
	v1715 = *(*int32)(unsafe.Add(mBase, uint32(v1711+v1688<<(uint(int32(2))%32))))
	v1716 = *(*int32)(unsafe.Add(mBase, uint32(v1715)+4))
	if v1716 == v1710 {
		v1773 = v1710
		goto L245
	} else {
		goto L246
	}
L244:
	;
	v1812 = v1802
	goto L238
L245:
	;
	v1802 = F_lappend(m, v1682, v1773)
	mBase = m.M
	v1803 = m.ExcPending
	if v1803 != 0 {
		goto L43
	} else {
		goto L252
	}
L246:
	;
	v1719 = int32(0)
	v1720 = *(*int32)(unsafe.Add(mBase, uint32(v1716)+4))
	if v1720 <= v1719 {
		v1773 = v1710
		goto L245
	} else {
		goto L247
	}
L247:
	;
	v1726 = v1710
	v1731 = v1719
	goto L248
L248:
	;
	v1755 = *(*int32)(unsafe.Add(mBase, uint32(v1716)+12))
	v1756 = int32(2)
	v1759 = *(*int32)(unsafe.Add(mBase, uint32(v1755+v1731<<(uint(v1756)%32))))
	v1763 = *(*int32)(unsafe.Add(mBase, uint32(v1584+v1759<<(uint(v1756)%32))))
	v1764 = F_lappend_int(m, v1726, v1763)
	mBase = m.M
	v1765 = m.ExcPending
	if v1765 != 0 {
		goto L43
	} else {
		goto L250
	}
L249:
	;
	v1773 = v1764
	goto L245
L250:
	;
	v1767 = v1731 + int32(1)
	v1768 = *(*int32)(unsafe.Add(mBase, uint32(v1716)+4))
	if v1767 < v1768 {
		v1726 = v1764
		v1731 = v1767
		goto L248
	} else {
		goto L251
	}
L251:
	;
	goto L249
L252:
	;
	v1805 = v1688 + int32(1)
	v1806 = *(*int32)(unsafe.Add(mBase, uint32(v1581)+4))
	if v1805 < v1806 {
		v1682 = v1802
		v1688 = v1805
		goto L243
	} else {
		goto L253
	}
L253:
	;
	goto L244
L254:
	;
	v1848 = v1554 + int32(1)
	v1849 = *(*int32)(unsafe.Add(mBase, uint32(v1499)+4))
	if v1848 < v1849 {
		v1544 = v1845
		v1554 = v1848
		goto L227
	} else {
		goto L255
	}
L255:
	;
	goto L228
L256:
	;
	v1885 = *(*int32)(unsafe.Add(mBase, uint32(v1506)+112))
	v1887 = F_create_groupingsets_path(m, v1484, v1485, v1486, v1885, int32(3), v1865, v1490)
	mBase = m.M
	v1888 = m.ExcPending
	if v1888 != 0 {
		goto L43
	} else {
		goto L257
	}
L257:
	;
	F_add_path(m, v1485, v1887)
	mBase = m.M
	v1890 = m.ExcPending
	if v1890 != 0 {
		goto L43
	} else {
		goto L258
	}
L258:
	;
	v1891 = v1484
	v1892 = v1485
	v1893 = v1486
	v1896 = v1489
	v1897 = v1490
	v1908 = v1501
	v1913 = v1506
	goto L116
L259:
	;
	v1924 = *(*int32)(unsafe.Add(mBase, uint32(v1913)+112))
	v1926 = *(*int32)(unsafe.Add(mBase, uint32(v1896)))
	v1927 = F_create_groupingsets_path(m, v1891, v1892, v1893, v1924, int32(1), v1926, v1897)
	mBase = m.M
	v1928 = m.ExcPending
	if v1928 != 0 {
		goto L43
	} else {
		goto L260
	}
L260:
	;
	v1946 = v1908
	v1961 = v1927
	goto L6
L261:
	;
	v1981 = v1946
	goto L5
}
func F_conv_18030_to_utf8(m *base.Module, l0 int32) int32 {
	var v15 int32
	_ = v15
	var v26 int32
	_ = v26
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v61 int32
	_ = v61
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v105 int32
	_ = v105
	var v113 int32
	_ = v113
	var v139 int32
	_ = v139
	var v150 int32
	_ = v150
	var v172 int32
	_ = v172
	var v180 int32
	_ = v180
	var v202 int32
	_ = v202
	var v210 int32
	_ = v210
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v272 int32
	_ = v272
	var v283 int32
	_ = v283
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v341 int32
	_ = v341
	var v349 int32
	_ = v349
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v409 int32
	_ = v409
	var v427 int32
	_ = v427
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v487 int32
	_ = v487
	var v505 int32
	_ = v505
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v567 int32
	_ = v567
	var v578 int32
	_ = v578
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v644 int32
	_ = v644
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v704 int32
	_ = v704
	var v722 int32
	_ = v722
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v774 int32
	_ = v774
	var v776 int32
	_ = v776
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v787 int32
	_ = v787
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	if base.Ui32(l0+int32(2127506640)) <= base.Ui32(int32(_a_F_conv_18030_to_utf8_0)) {
		v15 = int32(255)
		v26 = int32(base.Ui32(l0)>>(uint(int32(16))%32))&int32(55)*int32(1260) + l0&v15 + int32(base.Ui32(l0)>>(uint(int32(8))%32))&v15*int32(10) - int32(_a_F_conv_18030_to_utf8_1)
		if base.Ui32(v26) < base.Ui32(int32(128)) {
			v774 = v26
			return v774
		} else {
			if base.Ui32(v26) <= base.Ui32(int32(2047)) {
				return v26&int32(63) | v26<<(uint(int32(2))%32)&int32(_a_F_conv_18030_to_utf8_2) | int32(_a_F_conv_18030_to_utf8_3)
			} else {
				v42 = v26 & int32(63)
				v44 = v26 << (uint(int32(4)) % 32)
				v48 = v26 << (uint(int32(2)) % 32) & int32(_a_F_conv_18030_to_utf8_4)
				if base.Ui32(v26) <= base.Ui32(int32(_a_F_conv_18030_to_utf8_5)) {
					v776 = v42
					v778 = v44
					v779 = v48
					return v779 | v778&int32(_a_F_conv_18030_to_utf8_6) | v776 | int32(14712960)
				} else {
					v787 = v42
					v789 = v44
					v790 = v48
					return v787 | v789&int32(_a_F_conv_18030_to_utf8_7) | v790 | int32(-142573440)
				}
			}
		}
	} else {
		if base.Ui32(l0+int32(2127058887)) <= base.Ui32(int32(_a_F_conv_18030_to_utf8_8)) {
			v61 = int32(255)
			v70 = int32(base.Ui32(l0)>>(uint(int32(16))%32))&int32(63)*int32(1260) + l0&v61 + int32(base.Ui32(l0)>>(uint(int32(8))%32))&v61*int32(10)
			v72 = v70 - int32(_a_F_conv_18030_to_utf8_9)
			if base.Ui32(v72) < base.Ui32(int32(128)) {
				v774 = v72
				return v774
			} else {
				v76 = v70 - int32(_a_F_conv_18030_to_utf8_10)
				if base.Ui32(v72) <= base.Ui32(int32(2047)) {
					return v76&int32(63) | v72<<(uint(int32(2))%32)&int32(_a_F_conv_18030_to_utf8_2) | int32(_a_F_conv_18030_to_utf8_3)
				} else {
					v90 = v76 & int32(63)
					v92 = v72 << (uint(int32(4)) % 32)
					v96 = v72 << (uint(int32(2)) % 32) & int32(_a_F_conv_18030_to_utf8_4)
					if base.Ui32(v72) <= base.Ui32(int32(_a_F_conv_18030_to_utf8_5)) {
						v776 = v90
						v778 = v92
						v779 = v96
						return v779 | v778&int32(_a_F_conv_18030_to_utf8_6) | v776 | int32(14712960)
					} else {
						v787 = v90
						v789 = v92
						v790 = v96
						return v787 | v789&int32(_a_F_conv_18030_to_utf8_7) | v790 | int32(-142573440)
					}
				}
			}
		} else {
			if base.Ui32(l0+int32(2110740941)) <= base.Ui32(int32(_a_F_conv_18030_to_utf8_11)) {
				v105 = int32(255)
				v113 = int32(base.Ui32(l0)>>(uint(int32(8))%32))&v105*int32(10) + l0&v105 + int32(_a_F_conv_18030_to_utf8_12)
				return v113&int32(63) | v113<<(uint(int32(2))%32)&int32(_a_F_conv_18030_to_utf8_4) | v113<<(uint(int32(4))%32)&int32(_a_F_conv_18030_to_utf8_13) | int32(14712960)
			} else {
				if base.Ui32(l0+int32(2110663624)) <= base.Ui32(int32(_a_F_conv_18030_to_utf8_14)) {
					v139 = int32(255)
					v150 = int32(base.Ui32(l0)>>(uint(int32(16))%32))&int32(51)*int32(1260) + l0&v139 + int32(base.Ui32(l0)>>(uint(int32(8))%32))&v139*int32(10) - int32(_a_F_conv_18030_to_utf8_15)
					return v150&int32(63) | v150<<(uint(int32(2))%32)&int32(_a_F_conv_18030_to_utf8_4) | v150<<(uint(int32(4))%32)&int32(_a_F_conv_18030_to_utf8_6) | int32(14712960)
				} else {
					if base.Ui32(l0+int32(2110600905)) <= base.Ui32(int32(_a_F_conv_18030_to_utf8_16)) {
						v172 = int32(255)
						v180 = int32(base.Ui32(l0)>>(uint(int32(8))%32))&v172*int32(10) + l0&v172 + int32(_a_F_conv_18030_to_utf8_17)
						return v180&int32(63) | v180<<(uint(int32(2))%32)&int32(_a_F_conv_18030_to_utf8_4) | v180<<(uint(int32(4))%32)&int32(_a_F_conv_18030_to_utf8_13) | int32(14712960)
					} else {
						if base.Ui32(l0+int32(2110545095)) <= base.Ui32(int32(_a_F_conv_18030_to_utf8_18)) {
							v202 = int32(255)
							v210 = int32(base.Ui32(l0)>>(uint(int32(8))%32))&v202*int32(10) + l0&v202 + int32(_a_F_conv_18030_to_utf8_19)
							if base.Ui32(int32(128)) <= base.Ui32(v210) {
								if base.Ui32(v210) <= base.Ui32(int32(2047)) {
									v260 = v210&int32(63) | v210<<(uint(int32(2))%32)&int32(_a_F_conv_18030_to_utf8_2) | int32(_a_F_conv_18030_to_utf8_3)
								} else {
									if base.Ui32(v210) <= base.Ui32(int32(_a_F_conv_18030_to_utf8_5)) {
										v260 = v210<<(uint(int32(4))%32)&int32(_a_F_conv_18030_to_utf8_6) | (v210&int32(63) | v210<<(uint(int32(2))%32)&int32(_a_F_conv_18030_to_utf8_4)) | int32(14712960)
									} else {
										v259 = v210<<(uint(int32(2))%32)&int32(_a_F_conv_18030_to_utf8_4) | (v210<<(uint(int32(6))%32)&int32(117440512) | (v210&int32(63) | v210<<(uint(int32(4))%32)&int32(_a_F_conv_18030_to_utf8_7))) | int32(-260013952)
										v260 = v259
									}
								}
							} else {
								v259 = v210
								v260 = v259
							}
							return v260
						} else {
							if base.Ui32(l0+int32(2110527432)) <= base.Ui32(int32(_a_F_conv_18030_to_utf8_20)) {
								v272 = int32(255)
								v283 = int32(base.Ui32(l0)>>(uint(int32(16))%32))&int32(55)*int32(1260) + l0&v272 + int32(base.Ui32(l0)>>(uint(int32(8))%32))&v272*int32(10) - int32(_a_F_conv_18030_to_utf8_21)
								if base.Ui32(int32(128)) <= base.Ui32(v283) {
									if base.Ui32(v283) <= base.Ui32(int32(2047)) {
										v333 = v283&int32(63) | v283<<(uint(int32(2))%32)&int32(_a_F_conv_18030_to_utf8_2) | int32(_a_F_conv_18030_to_utf8_3)
									} else {
										if base.Ui32(v283) <= base.Ui32(int32(_a_F_conv_18030_to_utf8_5)) {
											v333 = v283<<(uint(int32(4))%32)&int32(_a_F_conv_18030_to_utf8_6) | (v283&int32(63) | v283<<(uint(int32(2))%32)&int32(_a_F_conv_18030_to_utf8_4)) | int32(14712960)
										} else {
											v332 = v283<<(uint(int32(2))%32)&int32(_a_F_conv_18030_to_utf8_4) | (v283<<(uint(int32(6))%32)&int32(117440512) | (v283&int32(63) | v283<<(uint(int32(4))%32)&int32(_a_F_conv_18030_to_utf8_7))) | int32(-260013952)
											v333 = v332
										}
									}
								} else {
									v332 = v283
									v333 = v332
								}
								return v333
							} else {
								if base.Ui32(l0+int32(2110480079)) <= base.Ui32(int32(_a_F_conv_18030_to_utf8_22)) {
									v341 = int32(255)
									v349 = int32(base.Ui32(l0)>>(uint(int32(8))%32))&v341*int32(10) + l0&v341 + int32(_a_F_conv_18030_to_utf8_23)
									if base.Ui32(int32(128)) <= base.Ui32(v349) {
										if base.Ui32(v349) <= base.Ui32(int32(2047)) {
											v399 = v349&int32(63) | v349<<(uint(int32(2))%32)&int32(_a_F_conv_18030_to_utf8_2) | int32(_a_F_conv_18030_to_utf8_3)
										} else {
											if base.Ui32(v349) <= base.Ui32(int32(_a_F_conv_18030_to_utf8_5)) {
												v399 = v349<<(uint(int32(4))%32)&int32(_a_F_conv_18030_to_utf8_6) | (v349&int32(63) | v349<<(uint(int32(2))%32)&int32(_a_F_conv_18030_to_utf8_4)) | int32(14712960)
											} else {
												v398 = v349<<(uint(int32(2))%32)&int32(_a_F_conv_18030_to_utf8_4) | (v349<<(uint(int32(6))%32)&int32(117440512) | (v349&int32(63) | v349<<(uint(int32(4))%32)&int32(_a_F_conv_18030_to_utf8_7))) | int32(-260013952)
												v399 = v398
											}
										}
									} else {
										v398 = v349
										v399 = v398
									}
									return v399
								} else {
									if base.Ui32(l0+int32(2110419149)) <= base.Ui32(int32(16857093)) {
										v409 = int32(255)
										v427 = int32(base.Ui32(l0)>>(uint(int32(24))%32))*int32(_a_F_conv_18030_to_utf8_24) + l0&v409 + int32(base.Ui32(l0)>>(uint(int32(16))%32))&v409*int32(1260) + int32(base.Ui32(l0)>>(uint(int32(8))%32))&v409*int32(10) - int32(_a_F_conv_18030_to_utf8_25)
										if base.Ui32(int32(128)) <= base.Ui32(v427) {
											if base.Ui32(v427) <= base.Ui32(int32(2047)) {
												v477 = v427&int32(63) | v427<<(uint(int32(2))%32)&int32(_a_F_conv_18030_to_utf8_2) | int32(_a_F_conv_18030_to_utf8_3)
											} else {
												if base.Ui32(v427) <= base.Ui32(int32(_a_F_conv_18030_to_utf8_5)) {
													v477 = v427<<(uint(int32(4))%32)&int32(_a_F_conv_18030_to_utf8_6) | (v427&int32(63) | v427<<(uint(int32(2))%32)&int32(_a_F_conv_18030_to_utf8_4)) | int32(14712960)
												} else {
													v476 = v427<<(uint(int32(2))%32)&int32(_a_F_conv_18030_to_utf8_4) | (v427<<(uint(int32(6))%32)&int32(117440512) | (v427&int32(63) | v427<<(uint(int32(4))%32)&int32(_a_F_conv_18030_to_utf8_7))) | int32(-260013952)
													v477 = v476
												}
											}
										} else {
											v476 = v427
											v477 = v476
										}
										return v477
									} else {
										if base.Ui32(l0+int32(2093559760)) <= base.Ui32(int32(16364804)) {
											v487 = int32(255)
											v505 = int32(base.Ui32(l0)>>(uint(int32(24))%32))*int32(_a_F_conv_18030_to_utf8_24) + l0&v487 + int32(base.Ui32(l0)>>(uint(int32(16))%32))&v487*int32(1260) + int32(base.Ui32(l0)>>(uint(int32(8))%32))&v487*int32(10) - int32(_a_F_conv_18030_to_utf8_26)
											if base.Ui32(int32(128)) <= base.Ui32(v505) {
												if base.Ui32(v505) <= base.Ui32(int32(2047)) {
													v555 = v505&int32(63) | v505<<(uint(int32(2))%32)&int32(_a_F_conv_18030_to_utf8_2) | int32(_a_F_conv_18030_to_utf8_3)
												} else {
													if base.Ui32(v505) <= base.Ui32(int32(_a_F_conv_18030_to_utf8_5)) {
														v555 = v505<<(uint(int32(4))%32)&int32(_a_F_conv_18030_to_utf8_6) | (v505&int32(63) | v505<<(uint(int32(2))%32)&int32(_a_F_conv_18030_to_utf8_4)) | int32(14712960)
													} else {
														v554 = v505<<(uint(int32(2))%32)&int32(_a_F_conv_18030_to_utf8_4) | (v505<<(uint(int32(6))%32)&int32(117440512) | (v505&int32(63) | v505<<(uint(int32(4))%32)&int32(_a_F_conv_18030_to_utf8_7))) | int32(-260013952)
														v555 = v554
													}
												}
											} else {
												v554 = v505
												v555 = v554
											}
											return v555
										} else {
											if base.Ui32(l0+int32(2077189064)) <= base.Ui32(int32(_a_F_conv_18030_to_utf8_27)) {
												v567 = int32(255)
												v578 = int32(base.Ui32(l0)>>(uint(int32(16))%32))&int32(49)*int32(1260) + l0&v567 + int32(base.Ui32(l0)>>(uint(int32(8))%32))&v567*int32(10) + int32(1946)
												if base.Ui32(int32(128)) <= base.Ui32(v578) {
													if base.Ui32(v578) <= base.Ui32(int32(2047)) {
														v628 = v578&int32(63) | v578<<(uint(int32(2))%32)&int32(_a_F_conv_18030_to_utf8_2) | int32(_a_F_conv_18030_to_utf8_3)
													} else {
														if base.Ui32(v578) <= base.Ui32(int32(_a_F_conv_18030_to_utf8_5)) {
															v628 = v578<<(uint(int32(4))%32)&int32(_a_F_conv_18030_to_utf8_6) | (v578&int32(63) | v578<<(uint(int32(2))%32)&int32(_a_F_conv_18030_to_utf8_4)) | int32(14712960)
														} else {
															v627 = v578<<(uint(int32(2))%32)&int32(_a_F_conv_18030_to_utf8_4) | (v578<<(uint(int32(6))%32)&int32(117440512) | (v578&int32(63) | v578<<(uint(int32(4))%32)&int32(_a_F_conv_18030_to_utf8_7))) | int32(-260013952)
															v628 = v627
														}
													}
												} else {
													v627 = v578
													v628 = v627
												}
												return v628
											} else {
												if base.Ui32(l0+int32(2077121996)) <= base.Ui32(int32(517)) {
													v644 = int32(base.Ui32(l0)>>(uint(int32(8))%32))&int32(167)*int32(10) + l0&int32(255) + int32(_a_F_conv_18030_to_utf8_28)
													if base.Ui32(int32(128)) <= base.Ui32(v644) {
														if base.Ui32(v644) <= base.Ui32(int32(2047)) {
															v694 = v644&int32(63) | v644<<(uint(int32(2))%32)&int32(_a_F_conv_18030_to_utf8_2) | int32(_a_F_conv_18030_to_utf8_3)
														} else {
															if base.Ui32(v644) <= base.Ui32(int32(_a_F_conv_18030_to_utf8_5)) {
																v694 = v644<<(uint(int32(4))%32)&int32(_a_F_conv_18030_to_utf8_6) | (v644&int32(63) | v644<<(uint(int32(2))%32)&int32(_a_F_conv_18030_to_utf8_4)) | int32(14712960)
															} else {
																v693 = v644<<(uint(int32(2))%32)&int32(_a_F_conv_18030_to_utf8_4) | (v644<<(uint(int32(6))%32)&int32(117440512) | (v644&int32(63) | v644<<(uint(int32(4))%32)&int32(_a_F_conv_18030_to_utf8_7))) | int32(-260013952)
																v694 = v693
															}
														}
													} else {
														v693 = v644
														v694 = v693
													}
													return v694
												} else {
													if base.Ui32(int32(1392646405)) < base.Ui32(l0+int32(1875869392)) {
														v774 = int32(0)
													} else {
														v704 = int32(255)
														v722 = int32(base.Ui32(l0)>>(uint(int32(24))%32))*int32(_a_F_conv_18030_to_utf8_24) + l0&v704 + int32(base.Ui32(l0)>>(uint(int32(16))%32))&v704*int32(1260) + int32(base.Ui32(l0)>>(uint(int32(8))%32))&v704*int32(10) - int32(_a_F_conv_18030_to_utf8_29)
														if base.Ui32(int32(128)) <= base.Ui32(v722) {
															if base.Ui32(v722) <= base.Ui32(int32(2047)) {
																v772 = v722&int32(63) | v722<<(uint(int32(2))%32)&int32(_a_F_conv_18030_to_utf8_2) | int32(_a_F_conv_18030_to_utf8_3)
															} else {
																if base.Ui32(v722) <= base.Ui32(int32(_a_F_conv_18030_to_utf8_5)) {
																	v772 = v722<<(uint(int32(4))%32)&int32(_a_F_conv_18030_to_utf8_6) | (v722&int32(63) | v722<<(uint(int32(2))%32)&int32(_a_F_conv_18030_to_utf8_4)) | int32(14712960)
																} else {
																	v771 = v722<<(uint(int32(2))%32)&int32(_a_F_conv_18030_to_utf8_4) | (v722<<(uint(int32(6))%32)&int32(117440512) | (v722&int32(63) | v722<<(uint(int32(4))%32)&int32(_a_F_conv_18030_to_utf8_7))) | int32(-260013952)
																	v772 = v771
																}
															}
														} else {
															v771 = v722
															v772 = v771
														}
														v774 = v772
													}
													return v774
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
func F_convert_saop_to_hashed_saop(m *base.Module, l0 int32) {
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = F_convert_saop_to_hashed_saop_walker(m, l0, l0)
	v3 = m.ExcPending
	if v3 != 0 {
		return
	} else {
		return
	}
}
func F_convert_tuples_by_name_attrmap(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v8 = F_palloc(m, int32(28))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = l2
		*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
		v17 = F_palloc(m, v6<<(uint(int32(2))%32))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v17
			v20 = F_palloc(m, v6)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = v20
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v25 = v23 + int32(1)
				v28 = F_palloc(m, v25<<(uint(int32(2))%32))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v28
					v31 = F_palloc(m, v25)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v31
						v34 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v34))) = int32(0)
						v37 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
						v38 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v37))) = uint8(v38)
						return v8
					}
				}
			}
		}
	}
}
func F_cost_ctescan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v7 float64
	_ = v7
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 float64
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 float64
	_ = v24
	var v25 int64
	_ = v25
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 float64
	_ = v62
	var v63 float64
	_ = v63
	var v64 float64
	_ = v64
	var v66 float64
	_ = v66
	var v73 float64
	_ = v73
	var v75 float64
	_ = v75
	var v76 float64
	_ = v76
	var v81 float64
	_ = v81
	var v82 float64
	_ = v82
	var v84 float64
	_ = v84
	var v86 float64
	_ = v86
	var v88 float64
	_ = v88
	var v90 float64
	_ = v90
	var v91 float64
	_ = v91
	var v99 float64
	_ = v99
	var v100 float64
	_ = v100
	var v101 float64
	_ = v101
	var v102 float64
	_ = v102
	var v106 float64
	_ = v106
	var v107 float64
	_ = v107
	var v108 int32
	_ = v108
	var v109 float64
	_ = v109
	var v110 float64
	_ = v110
	var v113 float64
	_ = v113
	var v115 float64
	_ = v115
	v7 = float64(0)
	v15 = m.G0
	v17 = v15 - int32(32)
	m.G0 = v17
	if l3 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v107 = *(*float64)(unsafe.Add(mBase, uint32(l2)+120))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v109 = *(*float64)(unsafe.Add(mBase, uint32(v108)+24))
	v110 = *(*float64)(unsafe.Add(mBase, uint32(v108)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = int32(0)
	v113 = float64(0)
	v115 = base.F64_add(v110, base.F64_add(v106, v113))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+48)) = v115
	*(*float64)(unsafe.Add(mBase, uint32(l0)+56)) = base.F64_add(v115, base.F64_add(base.F64_mul(v109, v101), base.F64_add(base.F64_mul(v107, base.F64_add(v99, base.F64_add(v100, v102))), v113)))
	m.G0 = v17 + int32(32)
	return
L2:
	;
	v19 = *(*float64)(unsafe.Add(mBase, uint32(l3)+8))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+32)) = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	v22 = int32(0)
	v24 = *(*float64)(unsafe.Add(mBase, _c_F_cost_ctescan[0]))
	v25 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v17)+16)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v17)+24)) = v25
	if v21 == v22 {
		v73 = v7
		v75 = v7
		v76 = v19
		v81 = v24
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v86 = *(*float64)(unsafe.Add(mBase, uint32(l2)+16))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+32)) = v86
	v88 = *(*float64)(unsafe.Add(mBase, uint32(l2)+200))
	v90 = *(*float64)(unsafe.Add(mBase, _c_F_cost_ctescan[0]))
	v91 = *(*float64)(unsafe.Add(mBase, uint32(l2)+192))
	v99 = v90
	v100 = v88
	v101 = v86
	v102 = v90
	v106 = v91
	goto L1
L5:
	;
	v82 = *(*float64)(unsafe.Add(mBase, uint32(l2)+200))
	v84 = *(*float64)(unsafe.Add(mBase, uint32(l2)+192))
	v99 = v24
	v100 = base.F64_add(v75, v82)
	v101 = v76
	v102 = v81
	v106 = base.F64_add(v73, v84)
	goto L1
L6:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	if v32 <= int32(0) {
		v73 = v7
		v75 = v7
		v76 = v19
		v81 = v24
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v38 = v22
	goto L8
L8:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v49+v38<<(uint(int32(2))%32))))
	v56 = F_cost_qual_eval_walker(m, v53, v17+int32(8))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v62 = *(*float64)(unsafe.Add(mBase, uint32(l0)+32))
	v63 = *(*float64)(unsafe.Add(mBase, uint32(v17)+24))
	v64 = *(*float64)(unsafe.Add(mBase, uint32(v17)+16))
	v66 = *(*float64)(unsafe.Add(mBase, _c_F_cost_ctescan[0]))
	v73 = v64
	v75 = v63
	v76 = v62
	v81 = v66
	goto L5
L10:
	;
	return
L11:
	;
	v59 = v38 + int32(1)
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	if v59 < v60 {
		v38 = v59
		goto L8
	} else {
		goto L12
	}
L12:
	;
	goto L9
}
func F_cost_incremental_sort(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 float64, l6 float64, l7 float64, l8 int32, l9 int32, l10 float64) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 float64
	_ = v23
	var v26 float64
	_ = v26
	var v29 int32
	_ = v29
	var v32 float64
	_ = v32
	var v35 float64
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
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
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v104 float64
	_ = v104
	var v105 int32
	_ = v105
	var v121 float64
	_ = v121
	var v125 int32
	_ = v125
	var v126 float64
	_ = v126
	var v132 float64
	_ = v132
	var v135 float64
	_ = v135
	var v137 float64
	_ = v137
	var v140 float64
	_ = v140
	var v143 int64
	_ = v143
	var v144 float64
	_ = v144
	var v151 float64
	_ = v151
	var v153 float64
	_ = v153
	var v157 int32
	_ = v157
	var v158 float64
	_ = v158
	var v160 float64
	_ = v160
	var v161 int32
	_ = v161
	var v164 float64
	_ = v164
	var v168 float64
	_ = v168
	var v170 float64
	_ = v170
	var v171 float64
	_ = v171
	var v173 float64
	_ = v173
	var v174 float64
	_ = v174
	var v178 float64
	_ = v178
	var v181 float64
	_ = v181
	var v185 float64
	_ = v185
	var v191 float64
	_ = v191
	var v192 float64
	_ = v192
	var v196 float64
	_ = v196
	var v200 float64
	_ = v200
	var v208 float64
	_ = v208
	var v211 float64
	_ = v211
	var v215 float64
	_ = v215
	var v216 float64
	_ = v216
	var v217 float64
	_ = v217
	var v221 float64
	_ = v221
	var v223 float64
	_ = v223
	var v231 float64
	_ = v231
	v12 = int32(0)
	v19 = m.G0
	v21 = v19 - int32(16)
	m.G0 = v21
	v23 = float64(2)
	if base.F64_lt(l7, v23) != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v26 = v23
	goto L3
L2:
	;
	v26 = l7
	goto L3
L3:
	;
	if l2 == int32(0) {
		v96 = v12
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v125 = v21 + int32(8)
	v126 = base.F64_div(v26, v121)
	v132 = float64(2)
	if base.F64_lt(v126, v132) != 0 {
		goto L25
	} else {
		goto L26
	}
L5:
	;
	v102 = int32(0)
	v104 = F_estimate_num_groups(m, l1, v96, v26, v102, v102)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L16
	} else {
		goto L23
	}
L6:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v29 <= int32(0) {
		v96 = v12
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v32 = float64(200)
	if base.F64_lt(v26, v32) != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v35 = v26
	goto L10
L9:
	;
	v35 = v32
	goto L10
L10:
	;
	v36 = int32(1)
	if l3 <= v36 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v39 = v36
	goto L13
L12:
	;
	v39 = l3
	goto L13
L13:
	;
	v46 = int32(0)
	v55 = v12
	goto L14
L14:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v62+v46<<(uint(int32(2))%32))))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+4))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+16))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+12))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+4))
	v72 = F_pull_varnos(m, l1, v71)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v96 = v77
	goto L5
L16:
	;
	return
L17:
	;
	v74 = F_bms_is_member(m, int32(0), v72)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	if v74 != 0 {
		v121 = v35
		goto L4
	} else {
		goto L19
	}
L19:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v70)+4))
	v77 = F_lappend(m, v55, v76)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L16
	} else {
		goto L20
	}
L20:
	;
	if v46 == v39-int32(1) {
		v96 = v77
		goto L5
	} else {
		goto L21
	}
L21:
	;
	v81 = v46 + int32(1)
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v81 < v82 {
		v46 = v81
		v55 = v77
		goto L14
	} else {
		goto L22
	}
L22:
	;
	goto L15
L23:
	;
	v121 = v104
	goto L4
L24:
	;
	v215 = *(*float64)(unsafe.Add(mBase, _c_F_cost_incremental_sort[0]))
	v216 = *(*float64)(unsafe.Add(mBase, uint32(v21)))
	v217 = *(*float64)(unsafe.Add(mBase, uint32(v21)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = l4
	*(*float64)(unsafe.Add(mBase, uint32(l0)+32)) = v26
	v221 = base.F64_div(base.F64_sub(l6, l5), v121)
	v223 = base.F64_add(v221, base.F64_add(l5, v217))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+48)) = v223
	v231 = base.F64_add(v121, float64(-1))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+56)) = base.F64_add(v223, base.F64_add(base.F64_mul(base.F64_add(v215, v215), v121), base.F64_add(base.F64_mul(base.F64_add(v215, float64(0)), v26), base.F64_add(base.F64_mul(v221, v231), base.F64_add(v216, base.F64_mul(base.F64_add(v217, v216), v231))))))
	m.G0 = v21 + int32(16)
	return
L25:
	;
	v135 = v132
	goto L27
L26:
	;
	v135 = v126
	goto L27
L27:
	;
	v137 = *(*float64)(unsafe.Add(mBase, _c_F_cost_incremental_sort[1]))
	v140 = base.F64_mul(v135, base.F64_add(base.F64_add(v137, v137), float64(0)))
	v143 = base.I64_extend_i32_s(l9) << (uint(int64(10)) % 64)
	v144 = base.F64_convert_i64_s(v143)
	v151 = base.F64_convert_i32_u((l8+int32(7))&int32(-8) + int32(24))
	v153 = base.F64_mul(v126, v151)
	v157 = base.F64_lt(l10, v135) & base.F64_gt(l10, float64(0))
	if v157 != 0 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v125))) = v208
	v211 = *(*float64)(unsafe.Add(mBase, _c_F_cost_incremental_sort[1]))
	*(*float64)(unsafe.Add(mBase, uint32(v21))) = base.F64_mul(v135, v211)
	goto L24
L29:
	;
	v158 = base.F64_mul(l10, v151)
	goto L31
L30:
	;
	v158 = v153
	goto L31
L31:
	;
	if base.F64_lt(v144, v158) != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v160 = F_log(m, v135)
	mBase = m.M
	v161 = F_tuplesort_merge_order(m, v143)
	mBase = m.M
	v164 = base.F64_mul(base.F64_div(v160, float64(0.693147180559945)), v140)
	*(*float64)(unsafe.Add(mBase, uint32(v125))) = v164
	v168 = base.F64_ceil(base.F64_mul(v153, float64(0.0001220703125)))
	v170 = base.F64_div(v153, v144)
	v171 = base.F64_convert_i32_s(v161)
	if base.F64_gt(v170, v171) != 0 {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	goto L34
L34:
	;
	if v157 != 0 {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	v173 = F_log(m, v170)
	mBase = m.M
	v174 = F_log(m, v171)
	mBase = m.M
	v178 = base.F64_ceil(base.F64_div(v173, v174))
	goto L37
L36:
	;
	v178 = float64(1)
	goto L37
L37:
	;
	v181 = *(*float64)(unsafe.Add(mBase, _c_F_cost_incremental_sort[2]))
	v185 = *(*float64)(unsafe.Add(mBase, _c_F_cost_incremental_sort[3]))
	v208 = base.F64_add(base.F64_mul(base.F64_mul(base.F64_add(v168, v168), v178), base.F64_add(base.F64_mul(v181, float64(0.75)), base.F64_mul(v185, float64(0.25)))), v164)
	goto L28
L38:
	;
	v191 = l10
	goto L40
L39:
	;
	v191 = v135
	goto L40
L40:
	;
	v192 = base.F64_add(v191, v191)
	if base.F64_gt(v135, v192)|base.F64_gt(v153, v144) != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v196 = F_log(m, v192)
	mBase = m.M
	v208 = base.F64_mul(base.F64_div(v196, float64(0.693147180559945)), v140)
	goto L28
L42:
	;
	goto L43
L43:
	;
	v200 = F_log(m, v135)
	mBase = m.M
	v208 = base.F64_mul(base.F64_div(v200, float64(0.693147180559945)), v140)
	goto L28
}
func F_create_hashjoin_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 float64
	_ = v12
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 float64
	_ = v90
	var v91 float64
	_ = v91
	var v93 float64
	_ = v93
	var v95 float64
	_ = v95
	var v97 float64
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 float64
	_ = v107
	var v109 int32
	_ = v109
	var v112 float64
	_ = v112
	var v114 int32
	_ = v114
	var v120 float64
	_ = v120
	var v124 float64
	_ = v124
	var v126 float64
	_ = v126
	var v128 float64
	_ = v128
	var v129 float64
	_ = v129
	var v138 float64
	_ = v138
	var v142 float64
	_ = v142
	var v150 float64
	_ = v150
	var v151 int32
	_ = v151
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v179 int32
	_ = v179
	var v190 float64
	_ = v190
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v313 int32
	_ = v313
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v332 int32
	_ = v332
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v395 int32
	_ = v395
	var v427 int32
	_ = v427
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v522 int32
	_ = v522
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v536 int32
	_ = v536
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v565 int32
	_ = v565
	var v567 int32
	_ = v567
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v607 int32
	_ = v607
	var v609 int32
	_ = v609
	var v621 float64
	_ = v621
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v648 float64
	_ = v648
	var v650 float64
	_ = v650
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v687 int32
	_ = v687
	var v689 int32
	_ = v689
	var v693 int32
	_ = v693
	var v698 int32
	_ = v698
	var v701 int32
	_ = v701
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v712 int32
	_ = v712
	var v715 int32
	_ = v715
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v721 int32
	_ = v721
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v732 int32
	_ = v732
	var v738 int32
	_ = v738
	var v744 int32
	_ = v744
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v755 int32
	_ = v755
	var v766 float64
	_ = v766
	var v789 int32
	_ = v789
	var v800 float64
	_ = v800
	var v826 int32
	_ = v826
	var v866 float64
	_ = v866
	var v867 int32
	_ = v867
	var v876 int32
	_ = v876
	var v882 float64
	_ = v882
	var v905 int32
	_ = v905
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v913 int32
	_ = v913
	var v922 int32
	_ = v922
	var v923 int32
	_ = v923
	var v925 int32
	_ = v925
	var v928 int32
	_ = v928
	var v929 int32
	_ = v929
	var v934 int32
	_ = v934
	var v941 int32
	_ = v941
	var v943 int32
	_ = v943
	var v945 int32
	_ = v945
	var v948 int32
	_ = v948
	var v950 int32
	_ = v950
	var v952 int32
	_ = v952
	var v959 int32
	_ = v959
	var v966 int32
	_ = v966
	var v967 float64
	_ = v967
	var v975 int32
	_ = v975
	var v976 int32
	_ = v976
	var v977 int32
	_ = v977
	var v980 int32
	_ = v980
	var v983 int32
	_ = v983
	var v984 int32
	_ = v984
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v990 int32
	_ = v990
	var v991 float64
	_ = v991
	var v992 float64
	_ = v992
	var v1001 int32
	_ = v1001
	var v1002 int32
	_ = v1002
	var v1003 int32
	_ = v1003
	var v1004 int32
	_ = v1004
	var v1006 int32
	_ = v1006
	var v1010 int32
	_ = v1010
	var v1011 float64
	_ = v1011
	var v1012 int32
	_ = v1012
	var v1015 float64
	_ = v1015
	var v1017 float64
	_ = v1017
	var v1018 float64
	_ = v1018
	var v1022 float64
	_ = v1022
	var v1024 int32
	_ = v1024
	var v1025 int32
	_ = v1025
	var v1038 float64
	_ = v1038
	var v1061 float64
	_ = v1061
	var v1062 float64
	_ = v1062
	var v1071 float64
	_ = v1071
	var v1075 float64
	_ = v1075
	var v1076 int32
	_ = v1076
	var v1077 int32
	_ = v1077
	var v1081 float64
	_ = v1081
	var v1083 int32
	_ = v1083
	var v1087 float64
	_ = v1087
	var v1088 float64
	_ = v1088
	var v1091 float64
	_ = v1091
	var v1094 float64
	_ = v1094
	var v1095 int64
	_ = v1095
	var v1104 int32
	_ = v1104
	var v1110 int32
	_ = v1110
	var v1141 int32
	_ = v1141
	var v1145 int32
	_ = v1145
	var v1148 int32
	_ = v1148
	var v1149 int32
	_ = v1149
	var v1151 int32
	_ = v1151
	var v1152 int32
	_ = v1152
	var v1154 float64
	_ = v1154
	var v1155 float64
	_ = v1155
	var v1174 float64
	_ = v1174
	var v1190 float64
	_ = v1190
	var v1191 int32
	_ = v1191
	var v1192 int64
	_ = v1192
	var v1201 int32
	_ = v1201
	var v1208 int32
	_ = v1208
	var v1239 int32
	_ = v1239
	var v1243 int32
	_ = v1243
	var v1246 int32
	_ = v1246
	var v1247 int32
	_ = v1247
	var v1249 int32
	_ = v1249
	var v1250 int32
	_ = v1250
	var v1252 float64
	_ = v1252
	var v1253 float64
	_ = v1253
	var v1277 float64
	_ = v1277
	var v1288 float64
	_ = v1288
	var v1289 int32
	_ = v1289
	var v1294 int32
	_ = v1294
	var v1297 float64
	_ = v1297
	var v1298 float64
	_ = v1298
	var v1300 float64
	_ = v1300
	var v1302 float64
	_ = v1302
	var v1305 float64
	_ = v1305
	var v1308 float64
	_ = v1308
	var v1312 float64
	_ = v1312
	var v1321 float64
	_ = v1321
	var v1325 float64
	_ = v1325
	var v1330 float64
	_ = v1330
	var v1339 float64
	_ = v1339
	var v1343 float64
	_ = v1343
	var v1346 float64
	_ = v1346
	var v1352 float64
	_ = v1352
	var v1353 float64
	_ = v1353
	var v1354 float64
	_ = v1354
	var v1363 float64
	_ = v1363
	var v1367 float64
	_ = v1367
	var v1368 int32
	_ = v1368
	var v1369 float64
	_ = v1369
	var v1370 int32
	_ = v1370
	var v1371 float64
	_ = v1371
	var v1373 int32
	_ = v1373
	var v1374 int32
	_ = v1374
	var v1375 int32
	_ = v1375
	var v1376 int32
	_ = v1376
	var v1377 int32
	_ = v1377
	var v1378 int64
	_ = v1378
	var v1398 float64
	_ = v1398
	var v1399 int32
	_ = v1399
	var v1406 int32
	_ = v1406
	var v1414 float64
	_ = v1414
	var v1437 int32
	_ = v1437
	var v1441 int32
	_ = v1441
	var v1442 int32
	_ = v1442
	var v1446 float64
	_ = v1446
	var v1447 int32
	_ = v1447
	var v1448 float64
	_ = v1448
	var v1450 int32
	_ = v1450
	var v1451 int32
	_ = v1451
	var v1464 float64
	_ = v1464
	var v1489 float64
	_ = v1489
	var v1491 float64
	_ = v1491
	var v1500 float64
	_ = v1500
	var v1504 float64
	_ = v1504
	var v1518 float64
	_ = v1518
	var v1540 float64
	_ = v1540
	var v1542 float64
	_ = v1542
	var v1543 int32
	_ = v1543
	var v1544 float64
	_ = v1544
	var v1556 float64
	_ = v1556
	var v1560 float64
	_ = v1560
	var v1561 float64
	_ = v1561
	var v1563 float64
	_ = v1563
	v12 = float64(0)
	v35 = m.G0
	v37 = v35 - int32(16)
	m.G0 = v37
	*(*int32)(unsafe.Add(mBase, uint32(v37)+12)) = l8
	v41 = F_palloc0(m, int32(112))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+8)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v41))) = int64(1541893259564)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+12)) = v48
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v53 = F_get_joinrel_parampathinfo(m, l0, l1, l5, l6, v50, l9, v37+int32(12))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+16)) = v53
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	v57 = l7 & v56
	*(*uint8)(unsafe.Add(mBase, uint32(v41)+20)) = uint8(v57)
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	if v60 != int32(1) {
		v68 = int32(0)
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v70 = v68 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v41)+21)) = uint8(v70)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l5)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+72)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v41)+64)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v41)+24)) = v72
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+8)))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+84)) = l6
	*(*int32)(unsafe.Add(mBase, uint32(v41)+80)) = l5
	*(*uint8)(unsafe.Add(mBase, uint32(v41)+76)) = uint8(v77)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+96)) = l10
	*(*int32)(unsafe.Add(mBase, uint32(v41)+88)) = v81
	v84 = m.G0
	v86 = v84 + int32(-64)
	m.G0 = v86
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l3)+84))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l3)+80))
	v90 = *(*float64)(unsafe.Add(mBase, uint32(l3)+24))
	v91 = *(*float64)(unsafe.Add(mBase, uint32(l3)+8))
	v93 = *(*float64)(unsafe.Add(mBase, uint32(l3)+88))
	v95 = *(*float64)(unsafe.Add(mBase, uint32(l6)+32))
	v97 = *(*float64)(unsafe.Add(mBase, uint32(l5)+32))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+40)) = v98
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v41)+16))
	if v100 != 0 {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5)+21)))
	if v64 != int32(1) {
		v68 = int32(0)
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6)+21)))
	v68 = v67
	goto L4
L7:
	;
	v107 = *(*float64)(unsafe.Add(mBase, uint32(v106)))
	*(*float64)(unsafe.Add(mBase, uint32(v41)+32)) = v107
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v41)+24))
	if int32(0) < v109 {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	v106 = v100 + int32(8)
	goto L7
L9:
	;
	goto L10
L10:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v41)+8))
	v106 = v103 + int32(16)
	goto L7
L11:
	;
	v112 = base.F64_convert_i32_u(v109)
	v114 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_hashjoin_path[0])))
	if v114 == int32(1) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v41)+104)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v41)+100)) = v88
	v150 = base.F64_mul(base.F64_convert_i32_s(v89), base.F64_convert_i32_s(v88))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	if v151 == int32(295) {
		goto L24
	} else {
		goto L25
	}
L14:
	;
	v120 = base.F64_add(base.F64_mul(v112, float64(-0.3)), float64(1))
	if base.F64_gt(v120, float64(0)) != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v126 = v112
	goto L16
L16:
	;
	v128 = float64(1e+100)
	v129 = base.F64_div(v107, v126)
	if base.F64_gt(v129, v128)|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v129)&int64(9223372036854775807))) != 0 {
		v142 = v128
		goto L20
	} else {
		goto L21
	}
L17:
	;
	v124 = v120
	goto L19
L18:
	;
	v124 = math.Float64frombits(uint64(0x8000000000000000))
	goto L19
L19:
	;
	v126 = base.F64_add(v124, v112)
	goto L16
L20:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v41)+32)) = v142
	goto L13
L21:
	;
	v138 = float64(1)
	if base.F64_le(v129, v138) != 0 {
		v142 = v138
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v142 = base.F64_nearest(v129)
	goto L20
L23:
	;
	v1061 = float64(1e+100)
	v1062 = base.F64_mul(v95, v1038)
	if base.F64_gt(v1062, v1061)|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1062)&int64(9223372036854775807))) != 0 {
		v1075 = v1061
		goto L190
	} else {
		goto L191
	}
L24:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v86))) = base.F64_div(float64(1), v150)
	v1038 = float64(0)
	goto L23
L25:
	;
	goto L26
L26:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v86))) = int64(4607182418800017408)
	v161 = m.G0
	v163 = v161 - int32(16)
	m.G0 = v163
	v165 = F_list_copy(m, l10)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	if l10 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	m.G0 = v163 + int32(16)
	if v826 == int32(0) {
		goto L145
	} else {
		goto L146
	}
L29:
	;
	v826 = int32(0)
	goto L28
L30:
	;
	goto L31
L31:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l10)+4))
	if v170 < int32(2) {
		v826 = l10
		goto L28
	} else {
		goto L32
	}
L32:
	;
	if v165 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v86))) = base.F64_div(float64(1), v800)
	v826 = v789
	goto L28
L34:
	;
	v789 = int32(0)
	v800 = float64(1)
	goto L33
L35:
	;
	goto L36
L36:
	;
	v179 = int32(0)
	v190 = float64(1)
	v209 = v165
	goto L37
L37:
	;
	v213 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v163)+12)) = v213
	v215 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v163)+8)) = v215
	v221 = v179
	v222 = v213
	v223 = v215
	v227 = v215
	v251 = v209
	v252 = v215
	goto L39
L38:
	;
	v789 = v755
	v800 = v766
	goto L33
L39:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v251)+4))
	if v227 < v255 {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v599 = *(*int32)(unsafe.Add(mBase, uint32(v163)+8))
	if v599 != 0 {
		goto L102
	} else {
		goto L103
	}
L41:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v251)+12))
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v257+v227<<(uint(int32(2))%32))))
	v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v261)+120)))
	if v264 != 0 {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	v565 = v221
	v567 = v223
	v595 = v251
	v596 = v252
	goto L43
L43:
	;
	goto L40
L44:
	;
	v265 = int32(48)
	goto L46
L45:
	;
	v265 = int32(44)
	goto L46
L46:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v261+v265)))
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v261)+4))
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v268)+28))
	if v264 != 0 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v286 = int32(0)
	if v267 == v286 {
		goto L61
	} else {
		goto L62
	}
L48:
	;
	v270 = int32(0)
	if v269 == v270 {
		v283 = v270
		goto L47
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	if v269 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v269)+4))
	if v273 < int32(2) {
		v283 = v270
		goto L47
	} else {
		goto L52
	}
L52:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v269)+12))
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v276)+4))
	v283 = v277
	goto L47
L53:
	;
	v283 = int32(0)
	goto L47
L54:
	;
	goto L55
L55:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v269)+12))
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v281)))
	v283 = v282
	goto L47
L56:
	;
	if v559 != 0 {
		v221 = v529
		v222 = v530
		v223 = v531
		v227 = v536 + int32(1)
		v251 = v559
		v252 = v560
		goto L39
	} else {
		goto L99
	}
L57:
	;
	v527 = F_list_delete_nth_cell(m, v251, v227)
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L1
	} else {
		goto L98
	}
L58:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v384 = F_remove_nulling_relids(m, v283, v382, int32(0))
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L1
	} else {
		goto L84
	}
L59:
	;
	v377 = F_lappend(m, v221, v261)
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L1
	} else {
		goto L83
	}
L60:
	;
	if v340 == int32(0) {
		goto L59
	} else {
		goto L75
	}
L61:
	;
	v340 = int32(0)
	goto L60
L62:
	;
	goto L63
L63:
	;
	v294 = int32(1)
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v267)+4))
	if v295 <= v294 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v298 = v294
	goto L66
L65:
	;
	v298 = v295
	goto L66
L66:
	;
	v303 = int32(0)
	v305 = int32(-1)
	goto L68
L67:
	;
	v340 = v332
	goto L60
L68:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v267+int32(8)+v303<<(uint(int32(2))%32))))
	if v313 != 0 {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v163+int32(12)))) = v324
	v332 = int32(1)
	goto L67
L70:
	;
	if base.B2i32(base.Ui32(int32(1)) < base.Ui32(base.I32_popcnt(v313)))|base.B2i32(int32(0) <= v305) != 0 {
		v332 = v286
		goto L67
	} else {
		goto L73
	}
L71:
	;
	v324 = v305
	goto L72
L72:
	;
	v326 = v303 + int32(1)
	if v326 != v298 {
		v303 = v326
		v305 = v324
		goto L68
	} else {
		goto L74
	}
L73:
	;
	v324 = base.I32_ctz(v313) | v303<<(uint(int32(5))%32)
	goto L72
L74:
	;
	goto L69
L75:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v163)+12))
	v345 = v343 << (uint(int32(2)) % 32)
	v346 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v345+v346)))
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v348)+112))
	if v349 == int32(0) {
		goto L59
	} else {
		goto L76
	}
L76:
	;
	if v222 < int32(0) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v354+v345)))
	if v356 == int32(0) {
		goto L59
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	if base.B2i32(v222 != v343) == int32(0) {
		v379 = v223
		v380 = v222
		goto L58
	} else {
		goto L82
	}
L80:
	;
	v359 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v356)+21)))
	v361 = v359 - int32(102)
	if base.B2i32(base.Ui32(int32(12)) < base.Ui32(v361))|base.B2i32(int32(1)<<(uint(v361)%32)&int32(_a_F_create_hashjoin_path_0) == int32(0)) != 0 {
		goto L59
	} else {
		goto L81
	}
L81:
	;
	v379 = v348
	v380 = v343
	goto L58
L82:
	;
	v529 = v221
	v530 = v222
	v531 = v223
	v536 = v227
	v559 = v251
	v560 = v252
	goto L56
L83:
	;
	v491 = v377
	v492 = v222
	v493 = v223
	v522 = v252
	goto L57
L84:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v163)+8))
	if v386 == int32(0) {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v476 = F_palloc0(m, int32(24))
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L1
	} else {
		goto L95
	}
L86:
	;
	v389 = int32(0)
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v386)+4))
	if v390 <= v389 {
		goto L85
	} else {
		goto L87
	}
L87:
	;
	v395 = v389
	goto L88
L88:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v386)+12))
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v427+v395<<(uint(int32(2))%32))))
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v431)))
	v433 = F_equal(m, v384, v432)
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L1
	} else {
		goto L90
	}
L89:
	;
	v529 = v221
	v530 = v380
	v531 = v379
	v536 = v227
	v559 = v251
	v560 = v252
	goto L56
L90:
	;
	if v433 == int32(0) {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v438 = v395 + int32(1)
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v386)+4))
	if v438 < v439 {
		v395 = v438
		goto L88
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	goto L89
L94:
	;
	goto L85
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v476))) = v384
	v479 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v163)+12))
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v479+v480<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v476)+4)) = v484
	v486 = F_lappend(m, v386, v476)
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v163)+8)) = v486
	v489 = F_lappend(m, v252, v261)
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	v491 = v221
	v492 = v380
	v493 = v379
	v522 = v489
	goto L57
L98:
	;
	v529 = v491
	v530 = v492
	v531 = v493
	v536 = v227 - int32(1)
	v559 = v527
	v560 = v522
	goto L56
L99:
	;
	v565 = v529
	v567 = v531
	v595 = v559
	v596 = v560
	goto L43
L100:
	;
	if v595 != 0 {
		v179 = v755
		v190 = v766
		v209 = v595
		goto L37
	} else {
		goto L144
	}
L101:
	;
	v621 = v190
	goto L109
L102:
	;
	v600 = *(*int32)(unsafe.Add(mBase, uint32(v599)+4))
	if int32(2) <= v600 {
		goto L101
	} else {
		goto L105
	}
L103:
	;
	goto L104
L104:
	;
	v603 = F_list_concat(m, v565, v596)
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L1
	} else {
		goto L106
	}
L105:
	;
	goto L104
L106:
	;
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v163)+8))
	F_list_free_deep(m, v605)
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	F_list_free(m, v596)
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	v755 = v603
	v766 = v190
	goto L100
L109:
	;
	v646 = F_estimate_multivariate_ndistinct(m, l0, v567, v163+int32(8), v163)
	mBase = m.M
	v647 = m.ExcPending
	if v647 != 0 {
		goto L1
	} else {
		goto L111
	}
L110:
	;
	v657 = v565
	v658 = int32(0)
	goto L116
L111:
	;
	v648 = *(*float64)(unsafe.Add(mBase, uint32(v163)))
	if base.F64_lt(v621, v648) != 0 {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v650 = v648
	goto L114
L113:
	;
	v650 = v621
	goto L114
L114:
	;
	if v646 != 0 {
		v621 = v650
		goto L109
	} else {
		goto L115
	}
L115:
	;
	goto L110
L116:
	;
	v687 = *(*int32)(unsafe.Add(mBase, uint32(v599)+4))
	if v658 < v687 {
		goto L118
	} else {
		goto L119
	}
L117:
	;
	v755 = v657
	v766 = v621
	goto L100
L118:
	;
	v689 = *(*int32)(unsafe.Add(mBase, uint32(v599)+12))
	v693 = v689 + v658<<(uint(int32(2))%32)
	goto L120
L119:
	;
	v693 = int32(0)
	goto L120
L120:
	;
	if v596 == int32(0) {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v755 = v565
	v766 = v621
	goto L100
L122:
	;
	goto L123
L123:
	;
	v698 = *(*int32)(unsafe.Add(mBase, uint32(v596)+4))
	if base.B2i32(v693 == int32(0))|base.B2i32(v698 <= v658) != 0 {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	goto L117
L125:
	;
	v701 = *(*int32)(unsafe.Add(mBase, uint32(v596)+12))
	if v701 == int32(0) {
		goto L124
	} else {
		goto L126
	}
L126:
	;
	v704 = *(*int32)(unsafe.Add(mBase, uint32(v163)+8))
	v705 = *(*int32)(unsafe.Add(mBase, uint32(v693)))
	v706 = int32(0)
	if v704 == v706 {
		goto L128
	} else {
		goto L129
	}
L127:
	;
	if v744 != 0 {
		goto L140
	} else {
		goto L141
	}
L128:
	;
	v744 = int32(0)
	goto L127
L129:
	;
	goto L130
L130:
	;
	v712 = *(*int32)(unsafe.Add(mBase, uint32(v704)+4))
	if v712 <= int32(0) {
		v738 = v706
		goto L131
	} else {
		goto L132
	}
L131:
	;
	v744 = v738
	goto L127
L132:
	;
	v715 = int32(0)
	if v715 < v712 {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v718 = v712
	goto L135
L134:
	;
	v718 = v715
	goto L135
L135:
	;
	v719 = *(*int32)(unsafe.Add(mBase, uint32(v704)+12))
	v721 = int32(0)
	goto L136
L136:
	;
	v729 = *(*int32)(unsafe.Add(mBase, uint32(v719+v721<<(uint(int32(2))%32))))
	v730 = base.B2i32(v729 == v705)
	if v729 == v705 {
		v738 = v730
		goto L131
	} else {
		goto L138
	}
L137:
	;
	v738 = v730
	goto L131
L138:
	;
	v732 = v721 + int32(1)
	if v732 != v718 {
		v721 = v732
		goto L136
	} else {
		goto L139
	}
L139:
	;
	goto L137
L140:
	;
	v748 = *(*int32)(unsafe.Add(mBase, uint32(v701+v658<<(uint(int32(2))%32))))
	v749 = F_lappend(m, v657, v748)
	mBase = m.M
	v750 = m.ExcPending
	if v750 != 0 {
		goto L1
	} else {
		goto L143
	}
L141:
	;
	v751 = v657
	goto L142
L142:
	;
	v657 = v751
	v658 = v658 + int32(1)
	goto L116
L143:
	;
	v751 = v749
	goto L142
L144:
	;
	goto L38
L145:
	;
	v1038 = float64(1)
	goto L23
L146:
	;
	goto L147
L147:
	;
	v866 = float64(1)
	v867 = *(*int32)(unsafe.Add(mBase, uint32(v826)+4))
	if v867 <= int32(0) {
		v1038 = v866
		goto L23
	} else {
		goto L148
	}
L148:
	;
	v876 = int32(0)
	v882 = v866
	goto L149
L149:
	;
	v905 = *(*int32)(unsafe.Add(mBase, uint32(v826)+12))
	v909 = *(*int32)(unsafe.Add(mBase, uint32(v905+v876<<(uint(int32(2))%32))))
	v910 = *(*int32)(unsafe.Add(mBase, uint32(v909)+48))
	v911 = *(*int32)(unsafe.Add(mBase, uint32(l6)+8))
	v912 = *(*int32)(unsafe.Add(mBase, uint32(v911)+8))
	v913 = int32(0)
	if v910 == v913 {
		goto L153
	} else {
		goto L154
	}
L150:
	;
	v1038 = v1022
	goto L23
L151:
	;
	v1017 = *(*float64)(unsafe.Add(mBase, uint32(v1012+v909)))
	v1018 = *(*float64)(unsafe.Add(mBase, uint32(v86)))
	if base.F64_lt(v1015, v1018) != 0 {
		goto L183
	} else {
		goto L184
	}
L152:
	;
	if v966 != 0 {
		goto L166
	} else {
		goto L167
	}
L153:
	;
	v966 = int32(1)
	goto L152
L154:
	;
	goto L155
L155:
	;
	if v912 == int32(0) {
		v959 = v913
		goto L156
	} else {
		goto L157
	}
L156:
	;
	v966 = v959
	goto L152
L157:
	;
	v922 = *(*int32)(unsafe.Add(mBase, uint32(v910)+4))
	v923 = *(*int32)(unsafe.Add(mBase, uint32(v912)+4))
	if v923 < v922 {
		v959 = v913
		goto L156
	} else {
		goto L158
	}
L158:
	;
	v925 = int32(1)
	if v922 <= v925 {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	v928 = v925
	goto L161
L160:
	;
	v928 = v922
	goto L161
L161:
	;
	v929 = int32(8)
	v934 = int32(0)
	goto L162
L162:
	;
	v941 = v934 << (uint(int32(2)) % 32)
	v943 = *(*int32)(unsafe.Add(mBase, uint32(v910+v929+v941)))
	v945 = *(*int32)(unsafe.Add(mBase, uint32(v912+v929+v941)))
	v948 = v943 & (v945 ^ int32(-1))
	v950 = base.B2i32(v948 == int32(0))
	if v948 != 0 {
		v959 = v950
		goto L156
	} else {
		goto L164
	}
L163:
	;
	v959 = v950
	goto L156
L164:
	;
	v952 = v934 + int32(1)
	if v952 != v928 {
		v934 = v952
		goto L162
	} else {
		goto L165
	}
L165:
	;
	goto L163
L166:
	;
	v967 = *(*float64)(unsafe.Add(mBase, uint32(v909)+136))
	if base.F64_lt(v967, float64(0)) == int32(0) {
		goto L169
	} else {
		goto L170
	}
L167:
	;
	goto L168
L168:
	;
	v992 = *(*float64)(unsafe.Add(mBase, uint32(v909)+128))
	if base.F64_lt(v992, float64(0)) == int32(0) {
		goto L176
	} else {
		goto L177
	}
L169:
	;
	v1012 = int32(152)
	v1015 = v967
	goto L151
L170:
	;
	goto L171
L171:
	;
	v975 = int32(0)
	v976 = *(*int32)(unsafe.Add(mBase, uint32(v909)+4))
	v977 = *(*int32)(unsafe.Add(mBase, uint32(v976)+28))
	if v977 == v975 {
		v985 = v975
		goto L172
	} else {
		goto L173
	}
L172:
	;
	v986 = int32(152)
	F_estimate_hash_bucket_stats(m, l0, v985, v150, v909+v986, v909+int32(136))
	mBase = m.M
	v990 = m.ExcPending
	if v990 != 0 {
		goto L1
	} else {
		goto L175
	}
L173:
	;
	v980 = *(*int32)(unsafe.Add(mBase, uint32(v977)+4))
	if v980 < int32(2) {
		v985 = v975
		goto L172
	} else {
		goto L174
	}
L174:
	;
	v983 = *(*int32)(unsafe.Add(mBase, uint32(v977)+12))
	v984 = *(*int32)(unsafe.Add(mBase, uint32(v983)+4))
	v985 = v984
	goto L172
L175:
	;
	v991 = *(*float64)(unsafe.Add(mBase, uint32(v909)+136))
	v1012 = v986
	v1015 = v991
	goto L151
L176:
	;
	v1012 = int32(144)
	v1015 = v992
	goto L151
L177:
	;
	goto L178
L178:
	;
	v1001 = *(*int32)(unsafe.Add(mBase, uint32(v909)+4))
	v1002 = *(*int32)(unsafe.Add(mBase, uint32(v1001)+28))
	if v1002 != 0 {
		goto L179
	} else {
		goto L180
	}
L179:
	;
	v1003 = *(*int32)(unsafe.Add(mBase, uint32(v1002)+12))
	v1004 = *(*int32)(unsafe.Add(mBase, uint32(v1003)))
	v1006 = v1004
	goto L181
L180:
	;
	v1006 = int32(0)
	goto L181
L181:
	;
	F_estimate_hash_bucket_stats(m, l0, v1006, v150, v909+int32(144), v909+int32(128))
	mBase = m.M
	v1010 = m.ExcPending
	if v1010 != 0 {
		goto L1
	} else {
		goto L182
	}
L182:
	;
	v1011 = *(*float64)(unsafe.Add(mBase, uint32(v909)+128))
	v1012 = int32(144)
	v1015 = v1011
	goto L151
L183:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v86))) = v1015
	goto L185
L184:
	;
	goto L185
L185:
	;
	if base.F64_gt(v882, v1017) != 0 {
		goto L186
	} else {
		goto L187
	}
L186:
	;
	v1022 = v1017
	goto L188
L187:
	;
	v1022 = v882
	goto L188
L188:
	;
	v1024 = v876 + int32(1)
	v1025 = *(*int32)(unsafe.Add(mBase, uint32(v826)+4))
	if v1024 < v1025 {
		v876 = v1024
		v882 = v1022
		goto L149
	} else {
		goto L189
	}
L189:
	;
	goto L150
L190:
	;
	v1076 = *(*int32)(unsafe.Add(mBase, uint32(l6)+12))
	v1077 = *(*int32)(unsafe.Add(mBase, uint32(v1076)+32))
	v1081 = *(*float64)(unsafe.Add(mBase, _c_F_create_hashjoin_path[1]))
	v1083 = *(*int32)(unsafe.Add(mBase, _c_F_create_hashjoin_path[2]))
	v1087 = base.F64_mul(base.F64_mul(v1081, base.F64_convert_i32_s(v1083)), float64(1024))
	v1088 = float64(4.294967295e+09)
	if base.F64_lt(v1087, v1088) != 0 {
		goto L194
	} else {
		goto L195
	}
L191:
	;
	v1071 = float64(1)
	if base.F64_le(v1062, v1071) != 0 {
		v1075 = v1071
		goto L190
	} else {
		goto L192
	}
L192:
	;
	v1075 = base.F64_nearest(v1062)
	goto L190
L193:
	;
	v1094 = *(*float64)(unsafe.Add(mBase, _c_F_create_hashjoin_path[3]))
	v1095 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v86)+16)) = v1095
	*(*int32)(unsafe.Add(mBase, uint32(v86)+8)) = l0
	*(*int64)(unsafe.Add(mBase, uint32(v86)+24)) = v1095
	if l10 == int32(0) {
		v1174 = v12
		v1190 = float64(0)
		goto L197
	} else {
		goto L198
	}
L194:
	;
	v1091 = v1087
	goto L196
L195:
	;
	v1091 = v1088
	goto L196
L196:
	;
	goto L193
L197:
	;
	v1191 = *(*int32)(unsafe.Add(mBase, uint32(v41)+88))
	v1192 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v86)+16)) = v1192
	*(*int32)(unsafe.Add(mBase, uint32(v86)+8)) = l0
	*(*int64)(unsafe.Add(mBase, uint32(v86)+24)) = v1192
	if v1191 == int32(0) {
		v1277 = v12
		v1288 = float64(0)
		goto L204
	} else {
		goto L205
	}
L198:
	;
	v1104 = *(*int32)(unsafe.Add(mBase, uint32(l10)+4))
	if v1104 <= int32(0) {
		v1174 = v12
		v1190 = float64(0)
		goto L197
	} else {
		goto L199
	}
L199:
	;
	v1110 = int32(0)
	goto L200
L200:
	;
	v1141 = *(*int32)(unsafe.Add(mBase, uint32(l10)+12))
	v1145 = *(*int32)(unsafe.Add(mBase, uint32(v1141+v1110<<(uint(int32(2))%32))))
	v1148 = F_cost_qual_eval_walker(m, v1145, v86+int32(8))
	mBase = m.M
	v1149 = m.ExcPending
	if v1149 != 0 {
		goto L1
	} else {
		goto L202
	}
L201:
	;
	v1154 = *(*float64)(unsafe.Add(mBase, uint32(v86)+16))
	v1155 = *(*float64)(unsafe.Add(mBase, uint32(v86)+24))
	v1174 = v1154
	v1190 = v1155
	goto L197
L202:
	;
	v1151 = v1110 + int32(1)
	v1152 = *(*int32)(unsafe.Add(mBase, uint32(l10)+4))
	if v1151 < v1152 {
		v1110 = v1151
		goto L200
	} else {
		goto L203
	}
L203:
	;
	goto L201
L204:
	;
	v1289 = *(*int32)(unsafe.Add(mBase, uint32(v41)+72))
	if v1289&int32(-2) != int32(4) {
		goto L213
	} else {
		goto L214
	}
L205:
	;
	v1201 = *(*int32)(unsafe.Add(mBase, uint32(v1191)+4))
	if v1201 <= int32(0) {
		v1277 = v12
		v1288 = float64(0)
		goto L204
	} else {
		goto L206
	}
L206:
	;
	v1208 = int32(0)
	goto L207
L207:
	;
	v1239 = *(*int32)(unsafe.Add(mBase, uint32(v1191)+12))
	v1243 = *(*int32)(unsafe.Add(mBase, uint32(v1239+v1208<<(uint(int32(2))%32))))
	v1246 = F_cost_qual_eval_walker(m, v1243, v86+int32(8))
	mBase = m.M
	v1247 = m.ExcPending
	if v1247 != 0 {
		goto L1
	} else {
		goto L209
	}
L208:
	;
	v1252 = *(*float64)(unsafe.Add(mBase, uint32(v86)+24))
	v1253 = *(*float64)(unsafe.Add(mBase, uint32(v86)+16))
	v1277 = v1252
	v1288 = v1253
	goto L204
L209:
	;
	v1249 = v1208 + int32(1)
	v1250 = *(*int32)(unsafe.Add(mBase, uint32(v1191)+4))
	if v1249 < v1250 {
		v1208 = v1249
		goto L207
	} else {
		goto L210
	}
L210:
	;
	goto L208
L211:
	;
	v1542 = *(*float64)(unsafe.Add(mBase, _c_F_create_hashjoin_path[4]))
	v1543 = *(*int32)(unsafe.Add(mBase, uint32(v41)+12))
	v1544 = *(*float64)(unsafe.Add(mBase, uint32(v1543)+24))
	if base.F64_lt(base.F64_convert_i32_u(base.I32_trunc_sat_f64_u(v1091)), base.F64_mul(v1075, base.F64_convert_i32_u((v1077+int32(7))&int32(-8)+int32(24)))) != 0 {
		goto L243
	} else {
		goto L244
	}
L212:
	;
	v1352 = float64(1e+100)
	v1353 = *(*float64)(unsafe.Add(mBase, uint32(v86)))
	v1354 = base.F64_mul(v95, v1353)
	if base.F64_gt(v1354, v1352)|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1354)&int64(9223372036854775807))) != 0 {
		v1367 = v1352
		goto L227
	} else {
		goto L228
	}
L213:
	;
	v1294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+8)))
	if v1294 != int32(1) {
		goto L212
	} else {
		goto L216
	}
L214:
	;
	goto L215
L215:
	;
	v1297 = float64(1e+100)
	v1298 = *(*float64)(unsafe.Add(mBase, uint32(l4)+16))
	v1300 = base.F64_nearest(base.F64_mul(v97, v1298))
	v1302 = base.F64_sub(v97, v1300)
	v1305 = *(*float64)(unsafe.Add(mBase, uint32(v86)))
	v1308 = *(*float64)(unsafe.Add(mBase, uint32(l4)+24))
	v1312 = base.F64_mul(base.F64_mul(v95, v1305), base.F64_div(float64(2), base.F64_add(v1308, float64(1))))
	if base.F64_gt(v1312, v1297) != 0 {
		v1325 = v1297
		goto L217
	} else {
		goto L218
	}
L216:
	;
	goto L215
L217:
	;
	v1330 = base.F64_div(v95, v150)
	if base.F64_gt(v1330, float64(1e+100))|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1330)&int64(9223372036854775807))) != 0 {
		v1343 = v1297
		goto L221
	} else {
		goto L222
	}
L218:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1312)&int64(9223372036854775807)) {
		v1325 = float64(1e+100)
		goto L217
	} else {
		goto L219
	}
L219:
	;
	v1321 = float64(1)
	if base.F64_le(v1312, v1321) != 0 {
		v1325 = v1321
		goto L217
	} else {
		goto L220
	}
L220:
	;
	v1325 = base.F64_nearest(v1312)
	goto L217
L221:
	;
	if v1289 == int32(5) {
		goto L224
	} else {
		goto L225
	}
L222:
	;
	v1339 = float64(1)
	if base.F64_le(v1330, v1339) != 0 {
		v1343 = v1339
		goto L221
	} else {
		goto L223
	}
L223:
	;
	v1343 = base.F64_nearest(v1330)
	goto L221
L224:
	;
	v1346 = v1302
	goto L226
L225:
	;
	v1346 = v1300
	goto L226
L226:
	;
	v1518 = v1346
	v1540 = base.F64_add(base.F64_mul(base.F64_mul(base.F64_mul(v1190, v1302), v1343), float64(0.05)), base.F64_add(base.F64_mul(base.F64_mul(base.F64_mul(v1190, v1300), v1325), float64(0.5)), v90))
	goto L211
L227:
	;
	v1368 = *(*int32)(unsafe.Add(mBase, uint32(v41)+84))
	v1369 = *(*float64)(unsafe.Add(mBase, uint32(v1368)+32))
	v1370 = *(*int32)(unsafe.Add(mBase, uint32(v41)+80))
	v1371 = *(*float64)(unsafe.Add(mBase, uint32(v1370)+32))
	v1373 = v86 + int32(8)
	v1374 = *(*int32)(unsafe.Add(mBase, uint32(v1370)+8))
	v1375 = *(*int32)(unsafe.Add(mBase, uint32(v1374)+8))
	v1376 = *(*int32)(unsafe.Add(mBase, uint32(v1368)+8))
	v1377 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+8))
	v1378 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1373)+48)) = v1378
	*(*int32)(unsafe.Add(mBase, uint32(v1373)+16)) = v1377
	*(*int32)(unsafe.Add(mBase, uint32(v1373)+12)) = v1375
	*(*int32)(unsafe.Add(mBase, uint32(v1373)+8)) = v1377
	*(*int32)(unsafe.Add(mBase, uint32(v1373)+4)) = v1375
	*(*int32)(unsafe.Add(mBase, uint32(v1373))) = int32(320)
	*(*int64)(unsafe.Add(mBase, uint32(v1373)+20)) = v1378
	*(*int64)(unsafe.Add(mBase, uint32(v1373)+28)) = v1378
	*(*int64)(unsafe.Add(mBase, uint32(v1373)+36)) = v1378
	*(*int32)(unsafe.Add(mBase, uint32(v1373)+43)) = int32(0)
	goto L230
L228:
	;
	v1363 = float64(1)
	if base.F64_le(v1354, v1363) != 0 {
		v1367 = v1363
		goto L227
	} else {
		goto L229
	}
L229:
	;
	v1367 = base.F64_nearest(v1354)
	goto L227
L230:
	;
	if l10 == int32(0) {
		goto L232
	} else {
		goto L233
	}
L231:
	;
	v1489 = float64(1e+100)
	v1491 = base.F64_mul(v1369, base.F64_mul(v1371, v1464))
	if base.F64_gt(v1491, v1489)|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1491)&int64(9223372036854775807))) != 0 {
		v1504 = v1489
		goto L240
	} else {
		goto L241
	}
L232:
	;
	v1464 = float64(1)
	goto L231
L233:
	;
	goto L234
L234:
	;
	v1398 = float64(1)
	v1399 = *(*int32)(unsafe.Add(mBase, uint32(l10)+4))
	if v1399 <= int32(0) {
		v1464 = v1398
		goto L231
	} else {
		goto L235
	}
L235:
	;
	v1406 = int32(0)
	v1414 = v1398
	goto L236
L236:
	;
	v1437 = *(*int32)(unsafe.Add(mBase, uint32(l10)+12))
	v1441 = *(*int32)(unsafe.Add(mBase, uint32(v1437+v1406<<(uint(int32(2))%32))))
	v1442 = int32(0)
	v1446 = F_clause_selectivity(m, l0, v1441, v1442, v1442, v86+int32(8))
	mBase = m.M
	v1447 = m.ExcPending
	if v1447 != 0 {
		goto L1
	} else {
		goto L238
	}
L237:
	;
	v1464 = v1448
	goto L231
L238:
	;
	v1448 = base.F64_mul(v1414, v1446)
	v1450 = v1406 + int32(1)
	v1451 = *(*int32)(unsafe.Add(mBase, uint32(l10)+4))
	if v1450 < v1451 {
		v1406 = v1450
		v1414 = v1448
		goto L236
	} else {
		goto L239
	}
L239:
	;
	goto L237
L240:
	;
	v1518 = v1504
	v1540 = base.F64_add(base.F64_mul(base.F64_mul(base.F64_mul(v97, v1190), v1367), float64(0.5)), v90)
	goto L211
L241:
	;
	v1500 = float64(1)
	if base.F64_le(v1491, v1500) != 0 {
		v1504 = v1500
		goto L240
	} else {
		goto L242
	}
L242:
	;
	v1504 = base.F64_nearest(v1491)
	goto L240
L243:
	;
	v1556 = base.F64_add(v91, v1094)
	goto L245
L244:
	;
	v1556 = v91
	goto L245
L245:
	;
	v1560 = *(*float64)(unsafe.Add(mBase, uint32(v1543)+16))
	v1561 = base.F64_add(base.F64_add(base.F64_add(v1556, v1174), base.F64_sub(v1288, v1174)), v1560)
	*(*float64)(unsafe.Add(mBase, uint32(v41)+48)) = v1561
	v1563 = *(*float64)(unsafe.Add(mBase, uint32(v41)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v41)+56)) = base.F64_add(v1561, base.F64_add(base.F64_mul(v1544, v1563), base.F64_add(base.F64_mul(base.F64_add(v1542, base.F64_sub(v1277, v1190)), v1518), v1540)))
	m.G0 = v86 - int32(-64)
	m.G0 = v37 + int32(16)
	return v41
}
func F_create_mergejoin_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 float64
	_ = v15
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 float64
	_ = v85
	var v86 float64
	_ = v86
	var v87 float64
	_ = v87
	var v88 float64
	_ = v88
	var v89 float64
	_ = v89
	var v90 float64
	_ = v90
	var v91 float64
	_ = v91
	var v93 float64
	_ = v93
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 float64
	_ = v106
	var v108 int32
	_ = v108
	var v111 float64
	_ = v111
	var v113 int32
	_ = v113
	var v119 float64
	_ = v119
	var v123 float64
	_ = v123
	var v125 float64
	_ = v125
	var v127 float64
	_ = v127
	var v128 float64
	_ = v128
	var v137 float64
	_ = v137
	var v141 float64
	_ = v141
	var v145 int64
	_ = v145
	var v150 float64
	_ = v150
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 float64
	_ = v204
	var v205 float64
	_ = v205
	var v223 float64
	_ = v223
	var v238 float64
	_ = v238
	var v239 int32
	_ = v239
	var v240 int64
	_ = v240
	var v249 int32
	_ = v249
	var v254 int32
	_ = v254
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v298 float64
	_ = v298
	var v299 float64
	_ = v299
	var v328 float64
	_ = v328
	var v332 float64
	_ = v332
	var v333 int32
	_ = v333
	var v338 int32
	_ = v338
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v357 float64
	_ = v357
	var v358 int32
	_ = v358
	var v359 float64
	_ = v359
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int64
	_ = v366
	var v385 float64
	_ = v385
	var v386 int32
	_ = v386
	var v391 int32
	_ = v391
	var v404 float64
	_ = v404
	var v422 int32
	_ = v422
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v431 float64
	_ = v431
	var v432 int32
	_ = v432
	var v433 float64
	_ = v433
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v452 float64
	_ = v452
	var v472 float64
	_ = v472
	var v474 float64
	_ = v474
	var v483 float64
	_ = v483
	var v487 float64
	_ = v487
	var v489 float64
	_ = v489
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v495 float64
	_ = v495
	var v501 int32
	_ = v501
	var v502 float64
	_ = v502
	var v505 int32
	_ = v505
	var v506 float64
	_ = v506
	var v510 float64
	_ = v510
	var v511 float64
	_ = v511
	var v512 int32
	_ = v512
	var v516 int32
	_ = v516
	var v520 float64
	_ = v520
	var v523 float64
	_ = v523
	var v528 int32
	_ = v528
	var v531 int32
	_ = v531
	var v561 int32
	_ = v561
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v571 int32
	_ = v571
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v578 int32
	_ = v578
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v586 int32
	_ = v586
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v592 int32
	_ = v592
	var v599 int32
	_ = v599
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v650 int32
	_ = v650
	var v662 float64
	_ = v662
	var v680 int32
	_ = v680
	var v681 float64
	_ = v681
	var v683 float64
	_ = v683
	var v691 float64
	_ = v691
	var v692 float64
	_ = v692
	var v694 float64
	_ = v694
	v15 = float64(0)
	v31 = int32(0)
	v33 = m.G0
	v35 = v33 - int32(16)
	m.G0 = v35
	*(*int32)(unsafe.Add(mBase, uint32(v35)+12)) = l7
	v39 = F_palloc0(m, int32(120))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+8)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v39))) = int64(1537598292267)
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+12)) = v46
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v51 = F_get_joinrel_parampathinfo(m, l0, l1, l5, l6, v48, l9, v35+int32(12))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v53 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v39)+20)) = uint8(v53)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+16)) = v51
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	if v56 != int32(1) {
		v63 = v31
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v65 = v63 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v39)+21)) = uint8(v65)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l5)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+72)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v39)+64)) = l8
	*(*int32)(unsafe.Add(mBase, uint32(v39)+24)) = v67
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+8)))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+84)) = l6
	*(*int32)(unsafe.Add(mBase, uint32(v39)+80)) = l5
	*(*uint8)(unsafe.Add(mBase, uint32(v39)+76)) = uint8(v71)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v35)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+108)) = l13
	*(*int32)(unsafe.Add(mBase, uint32(v39)+104)) = l12
	*(*int32)(unsafe.Add(mBase, uint32(v39)+100)) = l11
	*(*int32)(unsafe.Add(mBase, uint32(v39)+96)) = l10
	*(*int32)(unsafe.Add(mBase, uint32(v39)+88)) = v75
	v81 = m.G0
	v83 = v81 + int32(-64)
	m.G0 = v83
	v85 = *(*float64)(unsafe.Add(mBase, uint32(l3)+72))
	v86 = *(*float64)(unsafe.Add(mBase, uint32(l3)+64))
	v87 = *(*float64)(unsafe.Add(mBase, uint32(l3)+56))
	v88 = *(*float64)(unsafe.Add(mBase, uint32(l3)+48))
	v89 = *(*float64)(unsafe.Add(mBase, uint32(l3)+32))
	v90 = *(*float64)(unsafe.Add(mBase, uint32(l3)+24))
	v91 = *(*float64)(unsafe.Add(mBase, uint32(l3)+8))
	v93 = *(*float64)(unsafe.Add(mBase, uint32(l6)+32))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+40)) = v94
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v39)+16))
	if v99 != 0 {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5)+21)))
	if v59 != int32(1) {
		v63 = v31
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6)+21)))
	v63 = v62
	goto L4
L7:
	;
	v106 = *(*float64)(unsafe.Add(mBase, uint32(v105)))
	*(*float64)(unsafe.Add(mBase, uint32(v39)+32)) = v106
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v39)+24))
	if int32(0) < v108 {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	v105 = v99 + int32(8)
	goto L7
L9:
	;
	goto L10
L10:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v39)+8))
	v105 = v102 + int32(16)
	goto L7
L11:
	;
	v111 = base.F64_convert_i32_u(v108)
	v113 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_mergejoin_path[0])))
	if v113 == int32(1) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	v145 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v83)+16)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v83)+8)) = l0
	*(*int64)(unsafe.Add(mBase, uint32(v83)+24)) = v145
	v150 = float64(0)
	if l10 == int32(0) {
		v223 = v150
		v238 = v150
		goto L23
	} else {
		goto L24
	}
L14:
	;
	v119 = base.F64_add(base.F64_mul(v111, float64(-0.3)), float64(1))
	if base.F64_gt(v119, float64(0)) != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v125 = v111
	goto L16
L16:
	;
	v127 = float64(1e+100)
	v128 = base.F64_div(v106, v125)
	if base.F64_gt(v128, v127)|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v128)&int64(9223372036854775807))) != 0 {
		v141 = v127
		goto L20
	} else {
		goto L21
	}
L17:
	;
	v123 = v119
	goto L19
L18:
	;
	v123 = math.Float64frombits(uint64(0x8000000000000000))
	goto L19
L19:
	;
	v125 = base.F64_add(v123, v111)
	goto L16
L20:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v39)+32)) = v141
	goto L13
L21:
	;
	v137 = float64(1)
	if base.F64_le(v128, v137) != 0 {
		v141 = v137
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v141 = base.F64_nearest(v128)
	goto L20
L23:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v39)+88))
	v240 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v83)+16)) = v240
	*(*int32)(unsafe.Add(mBase, uint32(v83)+8)) = l0
	*(*int64)(unsafe.Add(mBase, uint32(v83)+24)) = v240
	if v239 == int32(0) {
		v328 = v15
		v332 = float64(0)
		goto L30
	} else {
		goto L31
	}
L24:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l10)+4))
	if v155 <= int32(0) {
		v223 = v150
		v238 = float64(0)
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v160 = int32(0)
	goto L26
L26:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(l10)+12))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v191+v160<<(uint(int32(2))%32))))
	v198 = F_cost_qual_eval_walker(m, v195, v81+int32(-56))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L1
	} else {
		goto L28
	}
L27:
	;
	v204 = *(*float64)(unsafe.Add(mBase, uint32(v83)+24))
	v205 = *(*float64)(unsafe.Add(mBase, uint32(v83)+16))
	v223 = v204
	v238 = v205
	goto L23
L28:
	;
	v201 = v160 + int32(1)
	v202 = *(*int32)(unsafe.Add(mBase, uint32(l10)+4))
	if v201 < v202 {
		v160 = v201
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v39)+72))
	if v333&int32(-2) != int32(4) {
		goto L39
	} else {
		goto L40
	}
L31:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v239)+4))
	if v249 <= int32(0) {
		v328 = v15
		v332 = float64(0)
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v254 = int32(0)
	goto L33
L33:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v239)+12))
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v285+v254<<(uint(int32(2))%32))))
	v292 = F_cost_qual_eval_walker(m, v289, v81+int32(-56))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L1
	} else {
		goto L35
	}
L34:
	;
	v298 = *(*float64)(unsafe.Add(mBase, uint32(v83)+24))
	v299 = *(*float64)(unsafe.Add(mBase, uint32(v83)+16))
	v328 = v298
	v332 = v299
	goto L30
L35:
	;
	v295 = v254 + int32(1)
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v239)+4))
	if v295 < v296 {
		v254 = v295
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v39)+112)) = uint8(v354)
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v39)+84))
	v357 = *(*float64)(unsafe.Add(mBase, uint32(v356)+32))
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v39)+80))
	v359 = *(*float64)(unsafe.Add(mBase, uint32(v358)+32))
	v361 = v81 + int32(-56)
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v358)+8))
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v362)+8))
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v356)+8))
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v364)+8))
	v366 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v361)+48)) = v366
	*(*int32)(unsafe.Add(mBase, uint32(v361)+16)) = v365
	*(*int32)(unsafe.Add(mBase, uint32(v361)+12)) = v363
	*(*int32)(unsafe.Add(mBase, uint32(v361)+8)) = v365
	*(*int32)(unsafe.Add(mBase, uint32(v361)+4)) = v363
	*(*int32)(unsafe.Add(mBase, uint32(v361))) = int32(320)
	*(*int64)(unsafe.Add(mBase, uint32(v361)+20)) = v366
	*(*int64)(unsafe.Add(mBase, uint32(v361)+28)) = v366
	*(*int64)(unsafe.Add(mBase, uint32(v361)+36)) = v366
	*(*int32)(unsafe.Add(mBase, uint32(v361)+43)) = int32(0)
	goto L50
L38:
	;
	v354 = int32(0)
	goto L37
L39:
	;
	v338 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+8)))
	if v338 != int32(1) {
		goto L38
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v39)+88))
	if v342 != 0 {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	goto L41
L43:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v342)+4))
	v345 = v343
	goto L45
L44:
	;
	v345 = int32(0)
	goto L45
L45:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v39)+96))
	if v346 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v346)+4))
	v349 = v347
	goto L48
L47:
	;
	v349 = int32(0)
	goto L48
L48:
	;
	if v345 == v349 {
		v354 = int32(1)
		goto L37
	} else {
		goto L49
	}
L49:
	;
	goto L38
L50:
	;
	if l10 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v472 = float64(1e+100)
	v474 = base.F64_mul(v357, base.F64_mul(v359, v452))
	if base.F64_gt(v474, v472)|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v474)&int64(9223372036854775807))) != 0 {
		v487 = v472
		goto L60
	} else {
		goto L61
	}
L52:
	;
	v452 = float64(1)
	goto L51
L53:
	;
	goto L54
L54:
	;
	v385 = float64(1)
	v386 = *(*int32)(unsafe.Add(mBase, uint32(l10)+4))
	if v386 <= int32(0) {
		v452 = v385
		goto L51
	} else {
		goto L55
	}
L55:
	;
	v391 = int32(0)
	v404 = v385
	goto L56
L56:
	;
	v422 = *(*int32)(unsafe.Add(mBase, uint32(l10)+12))
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v422+v391<<(uint(int32(2))%32))))
	v427 = int32(0)
	v431 = F_clause_selectivity(m, l0, v426, v427, v427, v81+int32(-56))
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L1
	} else {
		goto L58
	}
L57:
	;
	v452 = v433
	goto L51
L58:
	;
	v433 = base.F64_mul(v404, v431)
	v435 = v391 + int32(1)
	v436 = *(*int32)(unsafe.Add(mBase, uint32(l10)+4))
	if v435 < v436 {
		v391 = v435
		v404 = v433
		goto L56
	} else {
		goto L59
	}
L59:
	;
	goto L57
L60:
	;
	if base.F64_le(v93, float64(0)) != 0 {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	v483 = float64(1)
	if base.F64_le(v474, v483) != 0 {
		v487 = v483
		goto L60
	} else {
		goto L62
	}
L62:
	;
	v487 = base.F64_nearest(v474)
	goto L60
L63:
	;
	v489 = float64(1)
	goto L65
L64:
	;
	v489 = v93
	goto L65
L65:
	;
	v490 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+112)))
	v491 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	v492 = int32(295)
	v495 = float64(0)
	if v490&int32(1)|base.B2i32(v491 == v492) != 0 {
		v505 = v490 | base.B2i32(v491 != v492)
		v506 = v495
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v510 = base.F64_add(base.F64_div(v506, v87), float64(1))
	v511 = base.F64_mul(v89, v510)
	v512 = int32(1)
	if v505&v512 != 0 {
		goto L70
	} else {
		goto L71
	}
L67:
	;
	v501 = int32(0)
	v502 = base.F64_sub(v487, v489)
	if base.F64_lt(v502, float64(0)) != 0 {
		v505 = v501
		v506 = v495
		goto L66
	} else {
		goto L68
	}
L68:
	;
	v505 = v501
	v506 = v502
	goto L66
L69:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v39)+113)) = uint8(v650)
	v680 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
	v681 = *(*float64)(unsafe.Add(mBase, uint32(v680)+24))
	v683 = *(*float64)(unsafe.Add(mBase, _c_F_create_mergejoin_path[1]))
	v691 = *(*float64)(unsafe.Add(mBase, uint32(v680)+16))
	v692 = base.F64_add(base.F64_add(base.F64_sub(v332, v238), base.F64_add(base.F64_mul(v223, base.F64_add(base.F64_mul(v85, v510), v86)), base.F64_add(v91, v238))), v691)
	*(*float64)(unsafe.Add(mBase, uint32(v39)+48)) = v692
	v694 = *(*float64)(unsafe.Add(mBase, uint32(v39)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v39)+56)) = base.F64_add(v692, base.F64_add(base.F64_mul(v681, v694), base.F64_add(base.F64_mul(base.F64_add(v683, base.F64_sub(v328, v223)), v487), base.F64_add(base.F64_mul(v223, base.F64_add(base.F64_mul(base.F64_sub(v87, v85), v510), base.F64_sub(v88, v86))), base.F64_add(v90, v662)))))
	m.G0 = v83 - int32(-64)
	m.G0 = v35 + int32(16)
	return v39
L70:
	;
	v650 = int32(0)
	v662 = v511
	goto L69
L71:
	;
	v516 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_mergejoin_path[2])))
	v520 = *(*float64)(unsafe.Add(mBase, _c_F_create_mergejoin_path[3]))
	v523 = base.F64_add(base.F64_mul(base.F64_mul(v87, v520), v510), v89)
	if base.B2i32(v516 == int32(1))&base.F64_gt(v511, v523) != 0 {
		v650 = v512
		v662 = v523
		goto L69
	} else {
		goto L72
	}
L72:
	;
	if l12 == int32(0) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v528 = int32(0)
	v531 = l6
	goto L77
L74:
	;
	goto L75
L75:
	;
	if v516 == int32(0) {
		goto L70
	} else {
		goto L91
	}
L76:
	;
	if v592&int32(1) != 0 {
		goto L70
	} else {
		goto L90
	}
L77:
	;
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v531)+4))
	switch v561 - int32(331) {
	case 0:
		goto L82
	default:
		v592 = v528
		goto L76
	case 3:
		goto L81
	case 4:
		goto L80
	case 10, 11:
		goto L84
	case 24:
		goto L83
	case 29, 31:
		goto L79
	}
L78:
	;
	v592 = int32(1)
	goto L76
L79:
	;
	goto L78
L80:
	;
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v531)+72))
	if v583 == int32(0) {
		v592 = v528
		goto L76
	} else {
		goto L88
	}
L81:
	;
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v531)+72))
	if v575 == int32(0) {
		v592 = v528
		goto L76
	} else {
		goto L86
	}
L82:
	;
	v571 = *(*int32)(unsafe.Add(mBase, uint32(v531)))
	if v571 != int32(301) {
		v592 = v528
		goto L76
	} else {
		goto L85
	}
L83:
	;
	v566 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v531)+72)))
	v592 = int32(base.Ui32(v566&int32(2)) >> (uint(int32(1)) % 32))
	goto L76
L84:
	;
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v531)+72))
	v565 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v564)+112)))
	v592 = v565
	goto L76
L85:
	;
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v531)+72))
	v531 = v574
	goto L77
L86:
	;
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v575)+4))
	if v578 != int32(1) {
		v592 = v528
		goto L76
	} else {
		goto L87
	}
L87:
	;
	v581 = *(*int32)(unsafe.Add(mBase, uint32(v575)+12))
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v581)))
	v531 = v582
	goto L77
L88:
	;
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v583)+4))
	if v586 != int32(1) {
		v592 = v528
		goto L76
	} else {
		goto L89
	}
L89:
	;
	v589 = *(*int32)(unsafe.Add(mBase, uint32(v583)+12))
	v590 = *(*int32)(unsafe.Add(mBase, uint32(v589)))
	v531 = v590
	goto L77
L90:
	;
	v650 = v512
	v662 = v523
	goto L69
L91:
	;
	v599 = *(*int32)(unsafe.Add(mBase, _c_F_create_mergejoin_path[4]))
	v603 = *(*int32)(unsafe.Add(mBase, uint32(l6)+12))
	v604 = *(*int32)(unsafe.Add(mBase, uint32(v603)+32))
	if base.F64_lt(base.F64_convert_i32_u(v599<<(uint(int32(10))%32)), base.F64_mul(v489, base.F64_convert_i32_u((v604+int32(7))&int32(-8)+int32(24)))) != 0 {
		v650 = v512
		v662 = v523
		goto L69
	} else {
		goto L92
	}
L92:
	;
	goto L70
}
func F_crlf_process(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	if l3 <= int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L3
L3:
	;
	v10 = l2 + l3
	v13 = l2
	goto L4
L4:
	;
	v17 = v10 - v13
	v18 = int32(0)
	if base.B2i32(v13&int32(3) == v18)|base.B2i32(v17 == v18) != 0 {
		v48 = v13
		v50 = v17
		v51 = base.B2i32(v17 != v18)
		goto L11
	} else {
		goto L12
	}
L5:
	;
	return v165
L6:
	;
	goto L5
L7:
	;
	if base.Ui32(v10) <= base.Ui32(v136) {
		v159 = v135
		v160 = v136
		goto L42
	} else {
		goto L43
	}
L8:
	;
	if v122 != 0 {
		goto L33
	} else {
		goto L34
	}
L9:
	;
	v122 = int32(0)
	goto L8
L10:
	;
	v100 = v93
	v102 = v95
	goto L27
L11:
	;
	if v51 == int32(0) {
		goto L9
	} else {
		goto L18
	}
L12:
	;
	v31 = v13
	v33 = v17
	goto L13
L13:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
	if v36 == int32(10) {
		v93 = v31
		v95 = v33
		goto L10
	} else {
		goto L15
	}
L14:
	;
	v48 = v43
	v50 = v39
	v51 = v41
	goto L11
L15:
	;
	v38 = int32(1)
	v39 = v33 - v38
	v40 = int32(0)
	v41 = base.B2i32(v39 != v40)
	v43 = v31 + v38
	if v43&int32(3) == v40 {
		v48 = v43
		v50 = v39
		v51 = v41
		goto L11
	} else {
		goto L16
	}
L16:
	;
	if v39 != 0 {
		v31 = v43
		v33 = v39
		goto L13
	} else {
		goto L17
	}
L17:
	;
	goto L14
L18:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48))))
	if base.B2i32(int32(10) == v57)|base.B2i32(base.Ui32(v50) < base.Ui32(int32(4))) == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v66 = v48
	v68 = v50
	goto L22
L20:
	;
	v86 = v48
	v88 = v50
	goto L21
L21:
	;
	if v88 == int32(0) {
		goto L9
	} else {
		goto L26
	}
L22:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	v73 = v72 ^ int32(168430090)
	v76 = int32(-2139062144)
	if (int32(16843008)-v73|v73)&v76 != v76 {
		v93 = v66
		v95 = v68
		goto L10
	} else {
		goto L24
	}
L23:
	;
	v86 = v81
	v88 = v83
	goto L21
L24:
	;
	v80 = int32(4)
	v81 = v66 + v80
	v83 = v68 - v80
	if base.Ui32(int32(3)) < base.Ui32(v83) {
		v66 = v81
		v68 = v83
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	v93 = v86
	v95 = v88
	goto L10
L27:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100))))
	if int32(10) == v105 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	goto L9
L29:
	;
	v122 = v100
	goto L8
L30:
	;
	goto L31
L31:
	;
	v107 = int32(1)
	v110 = v102 - v107
	if v110 != 0 {
		v100 = v100 + v107
		v102 = v110
		goto L27
	} else {
		goto L32
	}
L32:
	;
	goto L28
L33:
	;
	v123 = v122
	goto L35
L34:
	;
	v123 = v10
	goto L35
L35:
	;
	v124 = v123 - v13
	if v124 <= int32(0) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v135 = int32(0)
	v136 = v13
	goto L7
L37:
	;
	goto L38
L38:
	;
	v128 = F_pushf_write(m, l0, v13, v124)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	return int32(0)
L40:
	;
	if v128 < int32(0) {
		v165 = v128
		goto L6
	} else {
		goto L41
	}
L41:
	;
	v135 = v128
	v136 = v13 + v124
	goto L7
L42:
	;
	if base.Ui32(v160) < base.Ui32(v10) {
		v13 = v160
		goto L4
	} else {
		goto L53
	}
L43:
	;
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136))))
	if v138 != int32(10) {
		v159 = v135
		v160 = v136
		goto L42
	} else {
		goto L44
	}
L44:
	;
	v143 = v136
	goto L45
L45:
	;
	v148 = F_pushf_write(m, l0, int32(_a_F_crlf_process_0), int32(2))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L39
	} else {
		goto L47
	}
L46:
	;
	v159 = v148
	v160 = v10
	goto L42
L47:
	;
	if v148 < int32(0) {
		v159 = v148
		v160 = v143
		goto L42
	} else {
		goto L48
	}
L48:
	;
	v153 = v143 + int32(1)
	if base.Ui32(v153) < base.Ui32(v10) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153))))
	if v155 != int32(10) {
		v159 = v148
		v160 = v153
		goto L42
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	goto L46
L52:
	;
	v143 = v153
	goto L45
L53:
	;
	v165 = v159
	goto L6
}
func F_crosstab(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int64
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v231 int64
	_ = v231
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v328 int32
	_ = v328
	var v329 int64
	_ = v329
	var v335 int32
	_ = v335
	var v350 int64
	_ = v350
	var v353 int32
	_ = v353
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v406 int64
	_ = v406
	var v408 int32
	_ = v408
	var v411 int32
	_ = v411
	var v427 int64
	_ = v427
	var v431 int32
	_ = v431
	var v450 int64
	_ = v450
	var v471 int64
	_ = v471
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v498 int64
	_ = v498
	var v502 int32
	_ = v502
	var v519 int64
	_ = v519
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v539 int64
	_ = v539
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v549 int32
	_ = v549
	var v571 int32
	_ = v571
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v597 int32
	_ = v597
	var v600 int64
	_ = v600
	var v602 int32
	_ = v602
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v636 int32
	_ = v636
	var v639 int32
	_ = v639
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v654 int32
	_ = v654
	var v659 int32
	_ = v659
	var v663 int32
	_ = v663
	var v666 int32
	_ = v666
	var v670 int32
	_ = v670
	var v674 int32
	_ = v674
	var v679 int32
	_ = v679
	var v683 int32
	_ = v683
	var v686 int32
	_ = v686
	var v690 int32
	_ = v690
	var v695 int32
	_ = v695
	var v699 int32
	_ = v699
	var v702 int32
	_ = v702
	var v706 int32
	_ = v706
	var v710 int32
	_ = v710
	var v715 int32
	_ = v715
	var v719 int32
	_ = v719
	var v722 int32
	_ = v722
	var v726 int32
	_ = v726
	var v731 int32
	_ = v731
	var v735 int32
	_ = v735
	var v738 int32
	_ = v738
	var v742 int32
	_ = v742
	var v747 int32
	_ = v747
	v20 = m.G0
	v22 = v20 - int32(32)
	m.G0 = v22
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v25 = F_pg_detoast_datum_packed(m, v24)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v29 = F_text_to_cstring(m, v25)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v31 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v735 = m.ExcPending
	if v735 != 0 {
		goto L1
	} else {
		goto L174
	}
L5:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	if v34 != int32(383) {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+12)))
	if v37&int32(2) != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+16))
	F_SPI_connect_ext(m, int32(0))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v719 = m.ExcPending
	if v719 != 0 {
		goto L1
	} else {
		goto L170
	}
L10:
	;
	v47 = F_SPI_execute(m, v29, int32(1), int32(0))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L17
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v699 = m.ExcPending
	if v699 != 0 {
		goto L1
	} else {
		goto L165
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v683 = m.ExcPending
	if v683 != 0 {
		goto L1
	} else {
		goto L161
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v663 = m.ExcPending
	if v663 != 0 {
		goto L1
	} else {
		goto L156
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v636 = m.ExcPending
	if v636 != 0 {
		goto L1
	} else {
		goto L149
	}
L15:
	;
	m.G0 = v22 + int32(32)
	return int32(0)
L16:
	;
	v63 = *(*int32)(unsafe.Add(mBase, _c_F_crosstab[0]))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	if v65 != int32(3) {
		goto L11
	} else {
		goto L23
	}
L17:
	;
	if v47 == int32(5) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v52 = *(*int64)(unsafe.Add(mBase, _c_F_crosstab[1]))
	if v52 != int64(0) {
		goto L16
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v56 = F_SPI_finish(m)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L22
	}
L21:
	;
	goto L20
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+20)) = int32(2)
	v60 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v60)
	goto L15
L23:
	;
	v71 = F_get_call_result_type(m, l0, int32(0), v22+int32(28))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	if v71 != int32(1) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	if v71 == int32(3) {
		goto L12
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
	if v94 <= int32(1) {
		goto L13
	} else {
		goto L33
	}
L28:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	F_errmsg(m, int32(_a_F_crosstab_0), int32(0))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	F_errfinish(m, int32(_a_F_crosstab_1), int32(442), int32(_a_F_crosstab_2))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L33:
	;
	v97 = int32(4)
	v99 = v93 + v94<<(uint(v97)%32)
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)+96))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	v104 = v64 + v101<<(uint(v97)%32)
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)+96))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v99)+88))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v104)+88))
	if base.B2i32(v100 != v105)&base.B2i32(int32(0) <= v100)|base.B2i32(v110 != v111) != 0 {
		goto L14
	} else {
		goto L34
	}
L34:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v104)+296))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v104)+288))
	v119 = int32(1)
	goto L35
L35:
	;
	v140 = v99 + int32(20) + v119*int32(100)
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v140)+76))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v140)+68))
	if v117 == v142 {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	v180 = int32(_a_F_crosstab_3)
	v181 = *(*int32)(unsafe.Add(mBase, _c_F_crosstab[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_crosstab[2])) = v41
	v184 = F_CreateTupleDescCopy(m, v93)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L50
	}
L37:
	;
	v178 = v119 + int32(1)
	if v178 != v94 {
		v119 = v178
		goto L35
	} else {
		goto L49
	}
L38:
	;
	if base.B2i32(v141 == v116)|base.B2i32(v141 < int32(0)) != 0 {
		goto L37
	} else {
		goto L41
	}
L39:
	;
	v148 = v142
	goto L40
L40:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L42
	}
L41:
	;
	v148 = v117
	goto L40
L42:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	F_errmsg(m, int32(_a_F_crosstab_4), int32(0))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v160 = F_format_type_with_typemod(m, v117, v116)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v162 = F_format_type_with_typemod(m, v148, v141)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = v119 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+4)) = v162
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v160
	F_errdetail(m, int32(_a_F_crosstab_5), v22)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(_a_F_crosstab_1), int32(1569), int32(_a_F_crosstab_6))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L49:
	;
	goto L36
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+28)) = v184
	v187 = int32(0)
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	v195 = *(*int32)(unsafe.Add(mBase, _c_F_crosstab[3]))
	v196 = F_tuplestore_begin_heap(m, int32(base.Ui32(v188&int32(4))>>(uint(int32(2))%32)), v187, v195)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_crosstab[2])) = v181
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	v201 = F_TupleDescGetAttInMetadata(m, v200)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v203)))
	v205 = int32(2)
	v209 = int32(1)
	v219 = v187
	v221 = v209
	v231 = int64(0)
	goto L53
L53:
	;
	v234 = F_palloc0(m, v204<<(uint(v205)%32))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L1
	} else {
		goto L55
	}
L54:
	;
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+28)) = v602
	*(*int32)(unsafe.Add(mBase, uint32(v31)+24)) = v196
	*(*int32)(unsafe.Add(mBase, uint32(v31)+16)) = int32(2)
	v607 = F_SPI_finish(m)
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		goto L1
	} else {
		goto L148
	}
L55:
	;
	if v204 < int32(2) {
		v471 = v231
		goto L59
	} else {
		goto L60
	}
L56:
	;
	v542 = int32(0)
	if v524 != 0 {
		goto L132
	} else {
		goto L133
	}
L57:
	;
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v234)))
	v524 = v522
	v539 = v519
	goto L56
L58:
	;
	F_pfree(m, v219)
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L1
	} else {
		goto L131
	}
L59:
	;
	v474 = F_BuildTupleFromCStrings(m, v201, v234)
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L1
	} else {
		goto L127
	}
L60:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v238+base.I32_wrap_i64(v231)<<(uint(int32(2))%32))))
	v245 = F_SPI_getvalue(m, v243, v64, int32(1))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	if v245 != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v247 = F_pstrdup(m, v245)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L1
	} else {
		goto L65
	}
L63:
	;
	v250 = int32(0)
	goto L64
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v234))) = v250
	if v221 != 0 {
		goto L71
	} else {
		goto L72
	}
L65:
	;
	v250 = v247
	goto L64
L66:
	;
	v471 = v450 - int64(1)
	goto L59
L67:
	;
	F_pfree(m, v411)
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L1
	} else {
		goto L126
	}
L68:
	;
	v322 = F_SPI_getvalue(m, v243, v64, int32(3))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L1
	} else {
		goto L95
	}
L69:
	;
	if v250 != 0 {
		v450 = v231
		goto L66
	} else {
		goto L94
	}
L70:
	;
	if v250 == int32(0) {
		v411 = v245
		v427 = v231
		goto L67
	} else {
		goto L85
	}
L71:
	;
	if v245 == int32(0) {
		goto L69
	} else {
		goto L84
	}
L72:
	;
	if v245|v219 == int32(0) {
		v524 = v250
		v539 = v231
		goto L56
	} else {
		goto L73
	}
L73:
	;
	v255 = int32(0)
	if base.B2i32(v219 == v255)|base.B2i32(v245 == v255) != 0 {
		goto L71
	} else {
		goto L74
	}
L74:
	;
	v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219))))
	v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v245))))
	if base.B2i32(v262 == int32(0))|base.B2i32(v262 != v265) != 0 {
		v283 = v262
		v284 = v265
		goto L76
	} else {
		goto L77
	}
L75:
	;
	if v283-v284 != 0 {
		goto L70
	} else {
		goto L82
	}
L76:
	;
	goto L75
L77:
	;
	v268 = v219
	v269 = v245
	goto L78
L78:
	;
	v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v269)+1)))
	v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v268)+1)))
	if v273 == int32(0) {
		v283 = v273
		v284 = v272
		goto L76
	} else {
		goto L80
	}
L79:
	;
	v283 = v273
	v284 = v272
	goto L76
L80:
	;
	v276 = int32(1)
	if v273 == v272 {
		v268 = v268 + v276
		v269 = v269 + v276
		goto L78
	} else {
		goto L81
	}
L81:
	;
	goto L79
L82:
	;
	F_pfree(m, v245)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	v498 = v231
	goto L58
L84:
	;
	goto L70
L85:
	;
	v294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v245))))
	v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v250))))
	if base.B2i32(v294 == int32(0))|base.B2i32(v294 != v297) != 0 {
		v315 = v294
		v316 = v297
		goto L87
	} else {
		goto L88
	}
L86:
	;
	if v315-v316 != 0 {
		v411 = v245
		v427 = v231
		goto L67
	} else {
		goto L93
	}
L87:
	;
	goto L86
L88:
	;
	v300 = v245
	v301 = v250
	goto L89
L89:
	;
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v301)+1)))
	v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v300)+1)))
	if v305 == int32(0) {
		v315 = v305
		v316 = v304
		goto L87
	} else {
		goto L91
	}
L90:
	;
	v315 = v305
	v316 = v304
	goto L87
L91:
	;
	v308 = int32(1)
	if v305 == v304 {
		v300 = v300 + v308
		v301 = v301 + v308
		goto L89
	} else {
		goto L92
	}
L92:
	;
	goto L90
L93:
	;
	v320 = int32(0)
	goto L68
L94:
	;
	v320 = int32(1)
	goto L68
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v234)+4)) = v322
	if v320 == int32(0) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	F_pfree(m, v245)
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L1
	} else {
		goto L99
	}
L97:
	;
	goto L98
L98:
	;
	v329 = v231 + base.I64_extend_i32_u(base.B2i32(v204 != v205))
	if v204 == int32(2) {
		v471 = v329
		goto L59
	} else {
		goto L100
	}
L99:
	;
	goto L98
L100:
	;
	if base.Ui64(v52) <= base.Ui64(v329) {
		v471 = v329
		goto L59
	} else {
		goto L101
	}
L101:
	;
	v335 = int32(1)
	v350 = v329
	goto L102
L102:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v353+base.I32_wrap_i64(v350)<<(uint(int32(2))%32))))
	v360 = F_SPI_getvalue(m, v358, v64, int32(1))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L1
	} else {
		goto L104
	}
L103:
	;
	v471 = v406
	goto L59
L104:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v234)))
	if v360 == int32(0) {
		goto L106
	} else {
		goto L107
	}
L105:
	;
	v399 = F_SPI_getvalue(m, v358, v64, int32(3))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L1
	} else {
		goto L119
	}
L106:
	;
	if v362 == int32(0) {
		goto L105
	} else {
		goto L109
	}
L107:
	;
	goto L108
L108:
	;
	if v362 == int32(0) {
		v411 = v360
		v427 = v350
		goto L67
	} else {
		goto L110
	}
L109:
	;
	v450 = v350
	goto L66
L110:
	;
	v371 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v360))))
	v374 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v362))))
	if base.B2i32(v371 == int32(0))|base.B2i32(v371 != v374) != 0 {
		v392 = v371
		v393 = v374
		goto L112
	} else {
		goto L113
	}
L111:
	;
	if v392-v393 != 0 {
		v411 = v360
		v427 = v350
		goto L67
	} else {
		goto L118
	}
L112:
	;
	goto L111
L113:
	;
	v377 = v360
	v378 = v362
	goto L114
L114:
	;
	v381 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v378)+1)))
	v382 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v377)+1)))
	if v382 == int32(0) {
		v392 = v382
		v393 = v381
		goto L112
	} else {
		goto L116
	}
L115:
	;
	v392 = v382
	v393 = v381
	goto L112
L116:
	;
	v385 = int32(1)
	if v382 == v381 {
		v377 = v377 + v385
		v378 = v378 + v385
		goto L114
	} else {
		goto L117
	}
L117:
	;
	goto L115
L118:
	;
	goto L105
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v234+v335<<(uint(int32(2))%32))+4)) = v399
	if v360 != 0 {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	F_pfree(m, v360)
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L1
	} else {
		goto L123
	}
L121:
	;
	goto L122
L122:
	;
	v406 = base.I64_extend_i32_u(base.B2i32(v335 < v204-v205)) + v350
	v408 = v335 + int32(1)
	if v204-v209 <= v408 {
		v471 = v406
		goto L59
	} else {
		goto L124
	}
L123:
	;
	goto L122
L124:
	;
	if base.Ui64(v406) < base.Ui64(v52) {
		v335 = v408
		v350 = v406
		goto L102
	} else {
		goto L125
	}
L125:
	;
	goto L103
L126:
	;
	v471 = v427 - int64(1)
	goto L59
L127:
	;
	F_tuplestore_puttuple(m, v196, v474)
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L1
	} else {
		goto L128
	}
L128:
	;
	F_pfree(m, v474)
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L1
	} else {
		goto L129
	}
L129:
	;
	if v219 == int32(0) {
		v519 = v471
		goto L57
	} else {
		goto L130
	}
L130:
	;
	v498 = v471
	goto L58
L131:
	;
	v519 = v498
	goto L57
L132:
	;
	v544 = F_pstrdup(m, v524)
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L1
	} else {
		goto L135
	}
L133:
	;
	v546 = v542
	goto L134
L134:
	;
	if int32(0) < v204 {
		goto L136
	} else {
		goto L137
	}
L135:
	;
	v546 = v544
	goto L134
L136:
	;
	v549 = v542
	goto L139
L137:
	;
	goto L138
L138:
	;
	F_pfree(m, v234)
	mBase = m.M
	v597 = m.ExcPending
	if v597 != 0 {
		goto L1
	} else {
		goto L146
	}
L139:
	;
	v571 = *(*int32)(unsafe.Add(mBase, uint32(v234+v549<<(uint(int32(2))%32))))
	if v571 != 0 {
		goto L141
	} else {
		goto L142
	}
L140:
	;
	goto L138
L141:
	;
	F_pfree(m, v571)
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L1
	} else {
		goto L144
	}
L142:
	;
	goto L143
L143:
	;
	v575 = v549 + int32(1)
	if v575 != v204 {
		v549 = v575
		goto L139
	} else {
		goto L145
	}
L144:
	;
	goto L143
L145:
	;
	goto L140
L146:
	;
	v600 = v539 + int64(1)
	if base.Ui64(v600) < base.Ui64(v52) {
		v219 = v546
		v221 = int32(0)
		v231 = v600
		goto L53
	} else {
		goto L147
	}
L147:
	;
	goto L54
L148:
	;
	goto L15
L149:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v639 = m.ExcPending
	if v639 != 0 {
		goto L1
	} else {
		goto L150
	}
L150:
	;
	F_errmsg(m, int32(_a_F_crosstab_4), int32(0))
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
		goto L1
	} else {
		goto L151
	}
L151:
	;
	v644 = F_format_type_with_typemod(m, v111, v105)
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L1
	} else {
		goto L152
	}
L152:
	;
	v646 = F_format_type_with_typemod(m, v110, v100)
	mBase = m.M
	v647 = m.ExcPending
	if v647 != 0 {
		goto L1
	} else {
		goto L153
	}
L153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = v646
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v644
	F_errdetail(m, int32(_a_F_crosstab_7), v22+int32(16))
	mBase = m.M
	v654 = m.ExcPending
	if v654 != 0 {
		goto L1
	} else {
		goto L154
	}
L154:
	;
	F_errfinish(m, int32(_a_F_crosstab_1), int32(1547), int32(_a_F_crosstab_6))
	mBase = m.M
	v659 = m.ExcPending
	if v659 != 0 {
		goto L1
	} else {
		goto L155
	}
L155:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L156:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v666 = m.ExcPending
	if v666 != 0 {
		goto L1
	} else {
		goto L157
	}
L157:
	;
	F_errmsg(m, int32(_a_F_crosstab_4), int32(0))
	mBase = m.M
	v670 = m.ExcPending
	if v670 != 0 {
		goto L1
	} else {
		goto L158
	}
L158:
	;
	F_errdetail(m, int32(_a_F_crosstab_8), int32(0))
	mBase = m.M
	v674 = m.ExcPending
	if v674 != 0 {
		goto L1
	} else {
		goto L159
	}
L159:
	;
	F_errfinish(m, int32(_a_F_crosstab_1), int32(1532), int32(_a_F_crosstab_6))
	mBase = m.M
	v679 = m.ExcPending
	if v679 != 0 {
		goto L1
	} else {
		goto L160
	}
L160:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L161:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v686 = m.ExcPending
	if v686 != 0 {
		goto L1
	} else {
		goto L162
	}
L162:
	;
	F_errmsg(m, int32(_a_F_crosstab_9), int32(0))
	mBase = m.M
	v690 = m.ExcPending
	if v690 != 0 {
		goto L1
	} else {
		goto L163
	}
L163:
	;
	F_errfinish(m, int32(_a_F_crosstab_1), int32(436), int32(_a_F_crosstab_2))
	mBase = m.M
	v695 = m.ExcPending
	if v695 != 0 {
		goto L1
	} else {
		goto L164
	}
L164:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L165:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v702 = m.ExcPending
	if v702 != 0 {
		goto L1
	} else {
		goto L166
	}
L166:
	;
	F_errmsg(m, int32(_a_F_crosstab_10), int32(0))
	mBase = m.M
	v706 = m.ExcPending
	if v706 != 0 {
		goto L1
	} else {
		goto L167
	}
L167:
	;
	F_errdetail(m, int32(_a_F_crosstab_11), int32(0))
	mBase = m.M
	v710 = m.ExcPending
	if v710 != 0 {
		goto L1
	} else {
		goto L168
	}
L168:
	;
	F_errfinish(m, int32(_a_F_crosstab_1), int32(423), int32(_a_F_crosstab_2))
	mBase = m.M
	v715 = m.ExcPending
	if v715 != 0 {
		goto L1
	} else {
		goto L169
	}
L169:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L170:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v722 = m.ExcPending
	if v722 != 0 {
		goto L1
	} else {
		goto L171
	}
L171:
	;
	F_errmsg(m, int32(_a_F_crosstab_12), int32(0))
	mBase = m.M
	v726 = m.ExcPending
	if v726 != 0 {
		goto L1
	} else {
		goto L172
	}
L172:
	;
	F_errfinish(m, int32(_a_F_crosstab_1), int32(386), int32(_a_F_crosstab_2))
	mBase = m.M
	v731 = m.ExcPending
	if v731 != 0 {
		goto L1
	} else {
		goto L173
	}
L173:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L174:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v738 = m.ExcPending
	if v738 != 0 {
		goto L1
	} else {
		goto L175
	}
L175:
	;
	F_errmsg(m, int32(_a_F_crosstab_13), int32(0))
	mBase = m.M
	v742 = m.ExcPending
	if v742 != 0 {
		goto L1
	} else {
		goto L176
	}
L176:
	;
	F_errfinish(m, int32(_a_F_crosstab_1), int32(382), int32(_a_F_crosstab_2))
	mBase = m.M
	v747 = m.ExcPending
	if v747 != 0 {
		goto L1
	} else {
		goto L177
	}
L177:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_currtid_internal(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v283 int32
	_ = v283
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v340 int32
	_ = v340
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v360 int32
	_ = v360
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v13 = F_palloc(m, int32(6))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_currtid_internal[0]))
	v21 = F_pg_class_aclcheck(m, v17, v19, int64(2))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if v21 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v24 = int32(*(*int8)(unsafe.Add(mBase, uint32(v23)+119)))
	switch v24 - int32(73) {
	case 0, 32:
		v34 = int32(20)
		goto L8
	default:
		goto L9
	case 10:
		goto L13
	case 29:
		goto L10
	case 36:
		goto L11
	case 45:
		goto L12
	}
L5:
	;
	goto L6
L6:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+119)))
	switch v43 - int32(83) {
	case 0, 22, 26, 31, 33:
		goto L18
	default:
		goto L19
	case 35:
		goto L20
	}
L7:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	F_aclcheck_error(m, v21, v36, v37+int32(4))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L14
	}
L8:
	;
	v36 = v34
	goto L7
L9:
	;
	v34 = int32(41)
	goto L8
L10:
	;
	v36 = int32(18)
	goto L7
L11:
	;
	v36 = int32(23)
	goto L7
L12:
	;
	v36 = int32(51)
	goto L7
L13:
	;
	v36 = int32(37)
	goto L7
L14:
	;
	goto L6
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L1
	} else {
		goto L101
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L1
	} else {
		goto L97
	}
L17:
	;
	m.G0 = v10 + int32(16)
	return v316
L18:
	;
	v289 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v13)+4)) = uint16(v289)
	v291 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v291
	v293 = F_GetLatestSnapshot(m)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L1
	} else {
		goto L91
	}
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L1
	} else {
		goto L86
	}
L20:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	if int32(0) < v47 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if v142 != 0 {
		goto L48
	} else {
		goto L49
	}
L22:
	;
	v58 = int32(0)
	goto L25
L23:
	;
	goto L24
L24:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L43
	}
L25:
	;
	v65 = v46 + v47<<(uint(int32(4))%32) + int32(20) + v58*int32(100)
	v67 = v65 + int32(4)
	v68 = int32(_a_F_currtid_internal_0)
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67))))
	v74 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_currtid_internal[1])))
	if base.B2i32(v71 == int32(0))|base.B2i32(v71 != v74) != 0 {
		v92 = v71
		v93 = v74
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L24
L27:
	;
	if v92-v93 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L28:
	;
	goto L27
L29:
	;
	v77 = v67
	v78 = v68
	goto L30
L30:
	;
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+1)))
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+1)))
	if v82 == int32(0) {
		v92 = v82
		v93 = v81
		goto L28
	} else {
		goto L32
	}
L31:
	;
	v92 = v82
	v93 = v81
	goto L28
L32:
	;
	v85 = int32(1)
	if v82 == v81 {
		v77 = v77 + v85
		v78 = v78 + v85
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v65)+68))
	if v97 == int32(27) {
		goto L21
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v117 = v58 + int32(1)
	if v117 != v47 {
		v58 = v117
		goto L25
	} else {
		goto L42
	}
L37:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	F_errmsg(m, int32(_a_F_currtid_internal_1), int32(0))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	F_errfinish(m, int32(_a_F_currtid_internal_2), int32(356), int32(_a_F_currtid_internal_3))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L42:
	;
	goto L26
L43:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	F_errmsg(m, int32(_a_F_currtid_internal_4), int32(0))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(_a_F_currtid_internal_2), int32(364), int32(_a_F_currtid_internal_3))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L47:
	;
	v164 = int32(0)
	goto L56
L48:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	if v143 <= int32(0) {
		goto L15
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L1
	} else {
		goto L52
	}
L51:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v142)+4))
	goto L47
L52:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	F_errmsg(m, int32(_a_F_currtid_internal_5), int32(0))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(_a_F_currtid_internal_2), int32(369), int32(_a_F_currtid_internal_3))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L56:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v146+v164<<(uint(int32(2))%32))))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v174)+4))
	if v175 != int32(1) {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v174)+12))
	if v181 == int32(0) {
		goto L16
	} else {
		goto L62
	}
L58:
	;
	v179 = v164 + int32(1)
	if v143 != v179 {
		v164 = v179
		goto L56
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	goto L57
L61:
	;
	goto L15
L62:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v181)+4))
	if v184 != int32(1) {
		goto L16
	} else {
		goto L63
	}
L63:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v181)+12))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v187)))
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v188)+76))
	if v189 != 0 {
		goto L66
	} else {
		goto L67
	}
L64:
	;
	if v230 == int32(0) {
		goto L15
	} else {
		goto L77
	}
L65:
	;
	goto L64
L66:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v189)+4))
	if v196 <= int32(0) {
		v230 = int32(0)
		goto L65
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v230 = int32(0)
	goto L65
L69:
	;
	v199 = int32(0)
	if v199 < v196 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v202 = v196
	goto L72
L71:
	;
	v202 = v199
	goto L72
L72:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v189)+12))
	v207 = int32(0)
	goto L73
L73:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v203+v207<<(uint(int32(2))%32))))
	v216 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v215)+8)))
	if v216 == base.I32_extend16_s(v58+int32(1))&int32(_a_F_currtid_internal_6) {
		v230 = v215
		goto L65
	} else {
		goto L75
	}
L74:
	;
	goto L68
L75:
	;
	v219 = v207 + int32(1)
	if v219 != v202 {
		v207 = v219
		goto L73
	} else {
		goto L76
	}
L76:
	;
	goto L74
L77:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v230)+4))
	if v234 == int32(0) {
		goto L15
	} else {
		goto L78
	}
L78:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v234)))
	if v237 != int32(6) {
		goto L15
	} else {
		goto L79
	}
L79:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v234)+4))
	if v240 < int32(0) {
		goto L15
	} else {
		goto L80
	}
L80:
	;
	v243 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v234)+8)))
	if v243 != int32(_a_F_currtid_internal_6) {
		goto L15
	} else {
		goto L81
	}
L81:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v188)+52))
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v246)+12))
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v247+v240<<(uint(int32(2))%32)-int32(4))))
	if v253 == int32(0) {
		goto L15
	} else {
		goto L82
	}
L82:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v253)+16))
	v258 = F_table_open(m, v256, int32(1))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	v260 = F_currtid_internal(m, v258, l1)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	F_relation_close(m, v258, int32(1))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	v316 = v260
	goto L17
L86:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v272)+68))
	v274 = F_get_namespace_name(m, v273)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v274
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v276 + int32(4)
	F_errmsg(m, int32(_a_F_currtid_internal_7), v10)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	F_errfinish(m, int32(_a_F_currtid_internal_2), int32(319), int32(_a_F_currtid_internal_8))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L91:
	;
	v295 = F_RegisterSnapshot(m, v293)
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	v297 = int32(0)
	v301 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v301)+8))
	v303 = m.T0[v302].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, l0, v295, v297, v297, v297, int32(8))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	F_table_tuple_get_latest_tid(m, v303, v13)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v303)))
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v307)+188))
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v308)+12))
	m.T0[v309].(func(*base.Module, int32))(m, v303)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	F_UnregisterSnapshot(m, v295)
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	v316 = v13
	goto L17
L97:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	F_errmsg(m, int32(_a_F_currtid_internal_9), int32(0))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	F_errfinish(m, int32(_a_F_currtid_internal_2), int32(381), int32(_a_F_currtid_internal_3))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L101:
	;
	F_errmsg_internal(m, int32(_a_F_currtid_internal_10), int32(0))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	F_errfinish(m, int32(_a_F_currtid_internal_2), int32(408), int32(_a_F_currtid_internal_3))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
