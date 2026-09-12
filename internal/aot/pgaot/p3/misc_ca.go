package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

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
	var v64 int32
	_ = v64
	var v73 int32
	_ = v73
	var v82 int32
	_ = v82
	var v91 int32
	_ = v91
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
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
								F_add_exact_object_address(m, v13+int32(-52), v49)
								mBase = m.M
								v64 = m.ExcPending
								if v64 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = l2
									*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = int32(1247)
									F_add_exact_object_address(m, v13+int32(-52), v49)
									mBase = m.M
									v73 = m.ExcPending
									if v73 != 0 {
										return
									} else {
										if l3 != 0 {
											*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = l3
											*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = int32(1255)
											F_add_exact_object_address(m, v13+int32(-52), v49)
											mBase = m.M
											v82 = m.ExcPending
											if v82 != 0 {
												return
											} else {
												if l4 != 0 {
													*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = int32(0)
													*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = l4
													*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = int32(2605)
													F_add_exact_object_address(m, v13+int32(-52), v49)
													mBase = m.M
													v91 = m.ExcPending
													if v91 != 0 {
														return
													} else {
														if l5 != 0 {
															*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = int32(0)
															*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = l5
															*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = int32(2605)
															F_add_exact_object_address(m, v13+int32(-52), v49)
															mBase = m.M
															v100 = m.ExcPending
															if v100 != 0 {
																return
															} else {
																F_record_object_address_dependencies(m, l0, v49, l8)
																mBase = m.M
																v102 = m.ExcPending
																if v102 != 0 {
																	return
																} else {
																	F_free_object_addresses(m, v49)
																	mBase = m.M
																	v104 = m.ExcPending
																	if v104 != 0 {
																		return
																	} else {
																		F_recordDependencyOnCurrentExtension(m, l0, int32(0))
																		mBase = m.M
																		v107 = m.ExcPending
																		if v107 != 0 {
																			return
																		} else {
																			v109 = *(*int32)(unsafe.Add(mBase, _c_F_CastCreate[0]))
																			if v109 != 0 {
																				v111 = int32(0)
																				F_RunObjectPostCreateHook(m, int32(2605), v32, v111, v111)
																				mBase = m.M
																				v114 = m.ExcPending
																				if v114 != 0 {
																					return
																				} else {
																					F_pfree(m, v45)
																					mBase = m.M
																					v116 = m.ExcPending
																					if v116 != 0 {
																						return
																					} else {
																						F_sequence_close(m, v23, int32(3))
																						mBase = m.M
																						v119 = m.ExcPending
																						if v119 != 0 {
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
																				v116 = m.ExcPending
																				if v116 != 0 {
																					return
																				} else {
																					F_sequence_close(m, v23, int32(3))
																					mBase = m.M
																					v119 = m.ExcPending
																					if v119 != 0 {
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
															v102 = m.ExcPending
															if v102 != 0 {
																return
															} else {
																F_free_object_addresses(m, v49)
																mBase = m.M
																v104 = m.ExcPending
																if v104 != 0 {
																	return
																} else {
																	F_recordDependencyOnCurrentExtension(m, l0, int32(0))
																	mBase = m.M
																	v107 = m.ExcPending
																	if v107 != 0 {
																		return
																	} else {
																		v109 = *(*int32)(unsafe.Add(mBase, _c_F_CastCreate[0]))
																		if v109 != 0 {
																			v111 = int32(0)
																			F_RunObjectPostCreateHook(m, int32(2605), v32, v111, v111)
																			mBase = m.M
																			v114 = m.ExcPending
																			if v114 != 0 {
																				return
																			} else {
																				F_pfree(m, v45)
																				mBase = m.M
																				v116 = m.ExcPending
																				if v116 != 0 {
																					return
																				} else {
																					F_sequence_close(m, v23, int32(3))
																					mBase = m.M
																					v119 = m.ExcPending
																					if v119 != 0 {
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
																			v116 = m.ExcPending
																			if v116 != 0 {
																				return
																			} else {
																				F_sequence_close(m, v23, int32(3))
																				mBase = m.M
																				v119 = m.ExcPending
																				if v119 != 0 {
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
														v100 = m.ExcPending
														if v100 != 0 {
															return
														} else {
															F_record_object_address_dependencies(m, l0, v49, l8)
															mBase = m.M
															v102 = m.ExcPending
															if v102 != 0 {
																return
															} else {
																F_free_object_addresses(m, v49)
																mBase = m.M
																v104 = m.ExcPending
																if v104 != 0 {
																	return
																} else {
																	F_recordDependencyOnCurrentExtension(m, l0, int32(0))
																	mBase = m.M
																	v107 = m.ExcPending
																	if v107 != 0 {
																		return
																	} else {
																		v109 = *(*int32)(unsafe.Add(mBase, _c_F_CastCreate[0]))
																		if v109 != 0 {
																			v111 = int32(0)
																			F_RunObjectPostCreateHook(m, int32(2605), v32, v111, v111)
																			mBase = m.M
																			v114 = m.ExcPending
																			if v114 != 0 {
																				return
																			} else {
																				F_pfree(m, v45)
																				mBase = m.M
																				v116 = m.ExcPending
																				if v116 != 0 {
																					return
																				} else {
																					F_sequence_close(m, v23, int32(3))
																					mBase = m.M
																					v119 = m.ExcPending
																					if v119 != 0 {
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
																			v116 = m.ExcPending
																			if v116 != 0 {
																				return
																			} else {
																				F_sequence_close(m, v23, int32(3))
																				mBase = m.M
																				v119 = m.ExcPending
																				if v119 != 0 {
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
														v102 = m.ExcPending
														if v102 != 0 {
															return
														} else {
															F_free_object_addresses(m, v49)
															mBase = m.M
															v104 = m.ExcPending
															if v104 != 0 {
																return
															} else {
																F_recordDependencyOnCurrentExtension(m, l0, int32(0))
																mBase = m.M
																v107 = m.ExcPending
																if v107 != 0 {
																	return
																} else {
																	v109 = *(*int32)(unsafe.Add(mBase, _c_F_CastCreate[0]))
																	if v109 != 0 {
																		v111 = int32(0)
																		F_RunObjectPostCreateHook(m, int32(2605), v32, v111, v111)
																		mBase = m.M
																		v114 = m.ExcPending
																		if v114 != 0 {
																			return
																		} else {
																			F_pfree(m, v45)
																			mBase = m.M
																			v116 = m.ExcPending
																			if v116 != 0 {
																				return
																			} else {
																				F_sequence_close(m, v23, int32(3))
																				mBase = m.M
																				v119 = m.ExcPending
																				if v119 != 0 {
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
																		v116 = m.ExcPending
																		if v116 != 0 {
																			return
																		} else {
																			F_sequence_close(m, v23, int32(3))
																			mBase = m.M
																			v119 = m.ExcPending
																			if v119 != 0 {
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
												v91 = m.ExcPending
												if v91 != 0 {
													return
												} else {
													if l5 != 0 {
														*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = int32(0)
														*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = l5
														*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = int32(2605)
														F_add_exact_object_address(m, v13+int32(-52), v49)
														mBase = m.M
														v100 = m.ExcPending
														if v100 != 0 {
															return
														} else {
															F_record_object_address_dependencies(m, l0, v49, l8)
															mBase = m.M
															v102 = m.ExcPending
															if v102 != 0 {
																return
															} else {
																F_free_object_addresses(m, v49)
																mBase = m.M
																v104 = m.ExcPending
																if v104 != 0 {
																	return
																} else {
																	F_recordDependencyOnCurrentExtension(m, l0, int32(0))
																	mBase = m.M
																	v107 = m.ExcPending
																	if v107 != 0 {
																		return
																	} else {
																		v109 = *(*int32)(unsafe.Add(mBase, _c_F_CastCreate[0]))
																		if v109 != 0 {
																			v111 = int32(0)
																			F_RunObjectPostCreateHook(m, int32(2605), v32, v111, v111)
																			mBase = m.M
																			v114 = m.ExcPending
																			if v114 != 0 {
																				return
																			} else {
																				F_pfree(m, v45)
																				mBase = m.M
																				v116 = m.ExcPending
																				if v116 != 0 {
																					return
																				} else {
																					F_sequence_close(m, v23, int32(3))
																					mBase = m.M
																					v119 = m.ExcPending
																					if v119 != 0 {
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
																			v116 = m.ExcPending
																			if v116 != 0 {
																				return
																			} else {
																				F_sequence_close(m, v23, int32(3))
																				mBase = m.M
																				v119 = m.ExcPending
																				if v119 != 0 {
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
														v102 = m.ExcPending
														if v102 != 0 {
															return
														} else {
															F_free_object_addresses(m, v49)
															mBase = m.M
															v104 = m.ExcPending
															if v104 != 0 {
																return
															} else {
																F_recordDependencyOnCurrentExtension(m, l0, int32(0))
																mBase = m.M
																v107 = m.ExcPending
																if v107 != 0 {
																	return
																} else {
																	v109 = *(*int32)(unsafe.Add(mBase, _c_F_CastCreate[0]))
																	if v109 != 0 {
																		v111 = int32(0)
																		F_RunObjectPostCreateHook(m, int32(2605), v32, v111, v111)
																		mBase = m.M
																		v114 = m.ExcPending
																		if v114 != 0 {
																			return
																		} else {
																			F_pfree(m, v45)
																			mBase = m.M
																			v116 = m.ExcPending
																			if v116 != 0 {
																				return
																			} else {
																				F_sequence_close(m, v23, int32(3))
																				mBase = m.M
																				v119 = m.ExcPending
																				if v119 != 0 {
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
																		v116 = m.ExcPending
																		if v116 != 0 {
																			return
																		} else {
																			F_sequence_close(m, v23, int32(3))
																			mBase = m.M
																			v119 = m.ExcPending
																			if v119 != 0 {
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
													v100 = m.ExcPending
													if v100 != 0 {
														return
													} else {
														F_record_object_address_dependencies(m, l0, v49, l8)
														mBase = m.M
														v102 = m.ExcPending
														if v102 != 0 {
															return
														} else {
															F_free_object_addresses(m, v49)
															mBase = m.M
															v104 = m.ExcPending
															if v104 != 0 {
																return
															} else {
																F_recordDependencyOnCurrentExtension(m, l0, int32(0))
																mBase = m.M
																v107 = m.ExcPending
																if v107 != 0 {
																	return
																} else {
																	v109 = *(*int32)(unsafe.Add(mBase, _c_F_CastCreate[0]))
																	if v109 != 0 {
																		v111 = int32(0)
																		F_RunObjectPostCreateHook(m, int32(2605), v32, v111, v111)
																		mBase = m.M
																		v114 = m.ExcPending
																		if v114 != 0 {
																			return
																		} else {
																			F_pfree(m, v45)
																			mBase = m.M
																			v116 = m.ExcPending
																			if v116 != 0 {
																				return
																			} else {
																				F_sequence_close(m, v23, int32(3))
																				mBase = m.M
																				v119 = m.ExcPending
																				if v119 != 0 {
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
																		v116 = m.ExcPending
																		if v116 != 0 {
																			return
																		} else {
																			F_sequence_close(m, v23, int32(3))
																			mBase = m.M
																			v119 = m.ExcPending
																			if v119 != 0 {
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
													v102 = m.ExcPending
													if v102 != 0 {
														return
													} else {
														F_free_object_addresses(m, v49)
														mBase = m.M
														v104 = m.ExcPending
														if v104 != 0 {
															return
														} else {
															F_recordDependencyOnCurrentExtension(m, l0, int32(0))
															mBase = m.M
															v107 = m.ExcPending
															if v107 != 0 {
																return
															} else {
																v109 = *(*int32)(unsafe.Add(mBase, _c_F_CastCreate[0]))
																if v109 != 0 {
																	v111 = int32(0)
																	F_RunObjectPostCreateHook(m, int32(2605), v32, v111, v111)
																	mBase = m.M
																	v114 = m.ExcPending
																	if v114 != 0 {
																		return
																	} else {
																		F_pfree(m, v45)
																		mBase = m.M
																		v116 = m.ExcPending
																		if v116 != 0 {
																			return
																		} else {
																			F_sequence_close(m, v23, int32(3))
																			mBase = m.M
																			v119 = m.ExcPending
																			if v119 != 0 {
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
																	v116 = m.ExcPending
																	if v116 != 0 {
																		return
																	} else {
																		F_sequence_close(m, v23, int32(3))
																		mBase = m.M
																		v119 = m.ExcPending
																		if v119 != 0 {
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
				v126 = m.ExcPending
				if v126 != 0 {
					return
				} else {
					F_errcode(m, int32(_a_F_CastCreate_0))
					mBase = m.M
					v129 = m.ExcPending
					if v129 != 0 {
						return
					} else {
						v130 = F_format_type_be(m, l1)
						mBase = m.M
						v131 = m.ExcPending
						if v131 != 0 {
							return
						} else {
							v132 = F_format_type_be(m, l2)
							mBase = m.M
							v133 = m.ExcPending
							if v133 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v132
								*(*int32)(unsafe.Add(mBase, uint32(v15))) = v130
								F_errmsg(m, int32(_a_F_CastCreate_1), v15)
								mBase = m.M
								v138 = m.ExcPending
								if v138 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_CastCreate_2), int32(77), int32(_a_F_CastCreate_3))
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
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v86 int64
	_ = v86
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
	v60 = F___fstatat(m, int32(-100), v8+int32(144), v8+int32(48), int32(0))
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
	v86 = *(*int64)(unsafe.Add(mBase, uint32(v8)+72))
	v20 = v20 + int32(1)
	v24 = v86 + v24
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
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v8 + int32(144)
	F_errmsg(m, int32(_a_F_calculate_relation_size_2), v8)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(_a_F_calculate_relation_size_3), int32(355), int32(_a_F_calculate_relation_size_4))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
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
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v144 int64
	_ = v144
	var v145 int32
	_ = v145
	var v147 int64
	_ = v147
	var v148 int64
	_ = v148
	var v150 int64
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v158 int64
	_ = v158
	var v160 int32
	_ = v160
	var v164 int64
	_ = v164
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
	return v164
L18:
	;
	if v64 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v164 = int64(-1)
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
	v158 = v4
	goto L25
L25:
	;
	F_FreeDir(m, v64)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
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
	v158 = v150
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
	v153 = F_ReadDir(m, v64, v7+int32(2192))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L3
	} else {
		goto L52
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v73 + int32(19)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v7 + int32(2192)
	v105 = F_pg_snprintf(m, v7+int32(144), int32(2048), int32(_a_F_calculate_tablespace_size_5), v7+int32(16))
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
		v150 = v76
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
		v150 = v76
		goto L32
	} else {
		goto L37
	}
L37:
	;
	goto L33
L38:
	;
	v113 = F___fstatat(m, int32(-100), v7+int32(144), v7+int32(48), int32(0))
	mBase = m.M
	goto L39
L39:
	;
	if v113 < int32(0) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v117 = *(*int32)(unsafe.Add(mBase, _c_F_calculate_tablespace_size[3]))
	if v117 == int32(44) {
		v150 = v76
		goto L32
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v7)+52))
	if v137&int32(_a_F_calculate_tablespace_size_6) == int32(_a_F_calculate_tablespace_size_7) {
		goto L48
	} else {
		goto L49
	}
L43:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L3
	} else {
		goto L44
	}
L44:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L3
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v7 + int32(144)
	F_errmsg(m, int32(_a_F_calculate_tablespace_size_8), v7)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L3
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(_a_F_calculate_tablespace_size_9), int32(266), int32(_a_F_calculate_tablespace_size_10))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
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
	v144 = F_db_dir_size(m, v7+int32(144))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L3
	} else {
		goto L51
	}
L49:
	;
	v147 = v76
	goto L50
L50:
	;
	v148 = *(*int64)(unsafe.Add(mBase, uint32(v7)+72))
	v150 = v147 + v148
	goto L32
L51:
	;
	v147 = v144 + v76
	goto L50
L52:
	;
	if v153 != 0 {
		v73 = v153
		v76 = v150
		goto L26
	} else {
		goto L53
	}
L53:
	;
	goto L27
L54:
	;
	v164 = v158
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
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
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
	var v42 int32
	_ = v42
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
	var v81 int32
	_ = v81
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
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
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
	var v143 int32
	_ = v143
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v180 int32
	_ = v180
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v195 int32
	_ = v195
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v215 int32
	_ = v215
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
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
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v272 int32
	_ = v272
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v310 int32
	_ = v310
	var v319 int32
	_ = v319
	var v326 int32
	_ = v326
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v366 int32
	_ = v366
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v403 int32
	_ = v403
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v418 int32
	_ = v418
	var v425 int32
	_ = v425
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v438 int32
	_ = v438
	var v445 int32
	_ = v445
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v495 int32
	_ = v495
	var v507 int32
	_ = v507
	var v511 int32
	_ = v511
	var v516 int32
	_ = v516
	var v523 int32
	_ = v523
	var v526 int32
	_ = v526
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v553 int32
	_ = v553
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v562 int32
	_ = v562
	var v568 int32
	_ = v568
	var v575 int32
	_ = v575
	var v579 int32
	_ = v579
	var v582 int32
	_ = v582
	var v588 int32
	_ = v588
	var v595 int32
	_ = v595
	var v599 int32
	_ = v599
	var v602 int32
	_ = v602
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v616 int32
	_ = v616
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v621 int32
	_ = v621
	var v623 int32
	_ = v623
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v638 int32
	_ = v638
	var v640 int32
	_ = v640
	var v645 int32
	_ = v645
	var v663 int32
	_ = v663
	var v670 int32
	_ = v670
	var v673 int32
	_ = v673
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v686 int32
	_ = v686
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v691 int32
	_ = v691
	var v693 int32
	_ = v693
	var v700 int32
	_ = v700
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v709 int32
	_ = v709
	var v715 int32
	_ = v715
	var v722 int32
	_ = v722
	var v726 int32
	_ = v726
	var v729 int32
	_ = v729
	var v735 int32
	_ = v735
	var v742 int32
	_ = v742
	var v746 int32
	_ = v746
	var v749 int32
	_ = v749
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v763 int32
	_ = v763
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v768 int32
	_ = v768
	var v770 int32
	_ = v770
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v785 int32
	_ = v785
	var v787 int32
	_ = v787
	var v792 int32
	_ = v792
	var v807 int32
	_ = v807
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v814 int32
	_ = v814
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v829 int32
	_ = v829
	var v838 int32
	_ = v838
	var v845 int32
	_ = v845
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v865 int32
	_ = v865
	var v873 int32
	_ = v873
	var v877 int32
	_ = v877
	var v879 int32
	_ = v879
	var v881 int32
	_ = v881
	var v883 int32
	_ = v883
	var v888 int32
	_ = v888
	var v895 int32
	_ = v895
	var v898 int32
	_ = v898
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v911 int32
	_ = v911
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v916 int32
	_ = v916
	var v918 int32
	_ = v918
	var v925 int32
	_ = v925
	var v931 int32
	_ = v931
	var v932 int32
	_ = v932
	var v934 int32
	_ = v934
	var v940 int32
	_ = v940
	var v947 int32
	_ = v947
	var v951 int32
	_ = v951
	var v954 int32
	_ = v954
	var v960 int32
	_ = v960
	var v967 int32
	_ = v967
	var v971 int32
	_ = v971
	var v974 int32
	_ = v974
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v979 int32
	_ = v979
	var v984 int32
	_ = v984
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v988 int32
	_ = v988
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v993 int32
	_ = v993
	var v995 int32
	_ = v995
	var v999 int32
	_ = v999
	var v1000 int32
	_ = v1000
	var v1001 int32
	_ = v1001
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1008 int32
	_ = v1008
	var v1010 int32
	_ = v1010
	var v1012 int32
	_ = v1012
	var v1017 int32
	_ = v1017
	var v1032 int32
	_ = v1032
	var v1035 int32
	_ = v1035
	var v1036 int32
	_ = v1036
	var v1041 int32
	_ = v1041
	var v1048 int32
	_ = v1048
	var v1051 int32
	_ = v1051
	var v1058 int32
	_ = v1058
	var v1059 int32
	_ = v1059
	var v1060 int32
	_ = v1060
	var v1064 int32
	_ = v1064
	var v1066 int32
	_ = v1066
	var v1067 int32
	_ = v1067
	var v1069 int32
	_ = v1069
	var v1071 int32
	_ = v1071
	var v1078 int32
	_ = v1078
	var v1084 int32
	_ = v1084
	var v1085 int32
	_ = v1085
	var v1087 int32
	_ = v1087
	var v1093 int32
	_ = v1093
	var v1100 int32
	_ = v1100
	var v1104 int32
	_ = v1104
	var v1107 int32
	_ = v1107
	var v1113 int32
	_ = v1113
	var v1120 int32
	_ = v1120
	var v1124 int32
	_ = v1124
	var v1127 int32
	_ = v1127
	var v1130 int32
	_ = v1130
	var v1131 int32
	_ = v1131
	var v1132 int32
	_ = v1132
	var v1137 int32
	_ = v1137
	var v1138 int32
	_ = v1138
	var v1139 int32
	_ = v1139
	var v1141 int32
	_ = v1141
	var v1143 int32
	_ = v1143
	var v1144 int32
	_ = v1144
	var v1146 int32
	_ = v1146
	var v1148 int32
	_ = v1148
	var v1152 int32
	_ = v1152
	var v1153 int32
	_ = v1153
	var v1154 int32
	_ = v1154
	var v1159 int32
	_ = v1159
	var v1160 int32
	_ = v1160
	var v1161 int32
	_ = v1161
	var v1163 int32
	_ = v1163
	var v1165 int32
	_ = v1165
	var v1170 int32
	_ = v1170
	var v1188 int32
	_ = v1188
	var v1195 int32
	_ = v1195
	var v1198 int32
	_ = v1198
	var v1205 int32
	_ = v1205
	var v1206 int32
	_ = v1206
	var v1207 int32
	_ = v1207
	var v1211 int32
	_ = v1211
	var v1213 int32
	_ = v1213
	var v1214 int32
	_ = v1214
	var v1216 int32
	_ = v1216
	var v1218 int32
	_ = v1218
	var v1225 int32
	_ = v1225
	var v1231 int32
	_ = v1231
	var v1232 int32
	_ = v1232
	var v1234 int32
	_ = v1234
	var v1240 int32
	_ = v1240
	var v1247 int32
	_ = v1247
	var v1251 int32
	_ = v1251
	var v1254 int32
	_ = v1254
	var v1260 int32
	_ = v1260
	var v1267 int32
	_ = v1267
	var v1271 int32
	_ = v1271
	var v1274 int32
	_ = v1274
	var v1277 int32
	_ = v1277
	var v1278 int32
	_ = v1278
	var v1279 int32
	_ = v1279
	var v1284 int32
	_ = v1284
	var v1285 int32
	_ = v1285
	var v1286 int32
	_ = v1286
	var v1288 int32
	_ = v1288
	var v1290 int32
	_ = v1290
	var v1291 int32
	_ = v1291
	var v1293 int32
	_ = v1293
	var v1295 int32
	_ = v1295
	var v1299 int32
	_ = v1299
	var v1300 int32
	_ = v1300
	var v1301 int32
	_ = v1301
	var v1306 int32
	_ = v1306
	var v1307 int32
	_ = v1307
	var v1308 int32
	_ = v1308
	var v1310 int32
	_ = v1310
	var v1312 int32
	_ = v1312
	var v1317 int32
	_ = v1317
	var v1332 int32
	_ = v1332
	var v1337 int32
	_ = v1337
	var v1341 int32
	_ = v1341
	var v1346 int32
	_ = v1346
	var v1347 int32
	_ = v1347
	var v1348 int32
	_ = v1348
	var v1350 int32
	_ = v1350
	var v1358 int32
	_ = v1358
	var v1362 int32
	_ = v1362
	var v1363 int32
	_ = v1363
	v2 = int32(0)
	v9 = F_strlen(m, l0)
	mBase = m.M
	v10 = v9 + l0
	if base.Ui32(v10) <= base.Ui32(l0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v39 != 0 {
		goto L8
	} else {
		goto L9
	}
L2:
	;
	v13 = v10 - int32(1)
	if base.Ui32(v13) <= base.Ui32(l0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v20 = v13
	goto L4
L4:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if v23 != int32(47) {
		goto L1
	} else {
		goto L6
	}
L5:
	;
	goto L1
L6:
	;
	v26 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v20))) = uint8(v26)
	v29 = v20 - int32(1)
	if base.Ui32(l0) < base.Ui32(v29) {
		v20 = v29
		goto L4
	} else {
		goto L7
	}
L7:
	;
	goto L5
L8:
	;
	v41 = l0
	v42 = l0
	v44 = v2
	goto L11
L9:
	;
	v66 = l0
	goto L10
L10:
	;
	v73 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v66))) = uint8(v73)
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v75 != 0 {
		goto L18
	} else {
		goto L19
	}
L11:
	;
	v49 = v42 + int32(1)
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
	if base.B2i32(v50 == int32(47))&v44 != 0 {
		v42 = v49
		goto L11
	} else {
		goto L13
	}
L12:
	;
	v66 = v59
	goto L10
L13:
	;
	if v41 != v42 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v41))) = uint8(v50)
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
	v57 = v56
	goto L16
L15:
	;
	v57 = v50
	goto L16
L16:
	;
	v59 = v41 + int32(1)
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
	if v64 != 0 {
		v41 = v59
		v42 = v49
		v44 = base.B2i32(v57&int32(255) == int32(47))
		goto L11
	} else {
		goto L17
	}
L17:
	;
	goto L12
L18:
	;
	if v75 == int32(47) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L20
L20:
	;
	return
L21:
	;
	v81 = l0 + int32(1)
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81))))
	v83 = v82
	v84 = v81
	v85 = int32(0)
	goto L23
L22:
	;
	v83 = v75
	v84 = l0
	v85 = int32(2)
	goto L23
L23:
	;
	if v83&int32(255) == int32(0) {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	if l0 == v1350 {
		goto L453
	} else {
		goto L454
	}
L25:
	;
	v1350 = v84
	goto L24
L26:
	;
	goto L27
L27:
	;
	v91 = v84
	v94 = v83
	v95 = v84
	v96 = v2
	v97 = v85
	goto L28
L28:
	;
	v100 = v94
	v103 = v95
	goto L31
L29:
	;
	v1350 = v1341
	goto L24
L30:
	;
	v122 = int32(0)
	if v120&int32(255) != int32(46) {
		v136 = v122
		goto L38
	} else {
		goto L39
	}
L31:
	;
	v107 = v100 & int32(255)
	if v107 == int32(0) {
		v120 = v94
		v121 = v103
		goto L30
	} else {
		goto L33
	}
L32:
	;
	v115 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v103))) = uint8(v115)
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95))))
	v120 = v119
	v121 = v103 + int32(1)
	goto L30
L33:
	;
	if v107 != int32(47) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v113 = v103 + int32(1)
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113))))
	v100 = v114
	v103 = v113
	goto L31
L35:
	;
	goto L36
L36:
	;
	goto L32
L37:
	;
	v1348 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121))))
	if v1348 != 0 {
		v91 = v1341
		v94 = v1348
		v95 = v121
		v96 = v1346
		v97 = v1347
		goto L28
	} else {
		goto L452
	}
L38:
	;
	switch v97 {
	case 0:
		goto L47
	case 1:
		goto L46
	case 2:
		goto L45
	case 3:
		goto L44
	case 4:
		goto L43
	default:
		v1341 = v91
		v1346 = v96
		v1347 = v97
		goto L37
	}
L39:
	;
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95)+1)))
	if v127 == int32(0) {
		v1341 = v91
		v1346 = v96
		v1347 = v97
		goto L37
	} else {
		goto L40
	}
L40:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95)+1)))
	if v130 != int32(46) {
		v136 = v122
		goto L38
	} else {
		goto L41
	}
L41:
	;
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95)+2)))
	v136 = base.B2i32(v133 == int32(0))
	goto L38
L42:
	;
	v1341 = v1332
	v1346 = v1337
	v1347 = int32(3)
	goto L37
L43:
	;
	v1032 = int32(47)
	*(*uint8)(unsafe.Add(mBase, uint32(v91))) = uint8(v1032)
	v1035 = v91 + int32(1)
	v1036 = F_strlen(m, v95)
	mBase = m.M
	if v136 != 0 {
		goto L351
	} else {
		goto L352
	}
L44:
	;
	if v136 != 0 {
		goto L274
	} else {
		goto L275
	}
L45:
	;
	v511 = F_strlen(m, v95)
	mBase = m.M
	if v136 != 0 {
		goto L173
	} else {
		goto L174
	}
L46:
	;
	if v136 != 0 {
		goto L100
	} else {
		goto L101
	}
L47:
	;
	if v136 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v1341 = v91
	v1346 = v96
	v1347 = int32(0)
	goto L37
L49:
	;
	goto L50
L50:
	;
	v138 = F_strlen(m, v95)
	mBase = m.M
	if v91 != v95 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	if v91 == v95 {
		goto L55
	} else {
		goto L56
	}
L52:
	;
	goto L53
L53:
	;
	v284 = int32(1)
	v1341 = v91 + v138
	v1346 = v96 + v284
	v1347 = v284
	goto L37
L54:
	;
	goto L53
L55:
	;
	goto L54
L56:
	;
	v143 = v91 + v138
	if base.Ui32(v95-v143) <= base.Ui32(int32(0)-v138<<(uint(int32(1))%32)) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v150 = F___memcpy(m, v91, v95, v138)
	mBase = m.M
	goto L54
L58:
	;
	goto L59
L59:
	;
	v153 = (v91 ^ v95) & int32(3)
	if base.Ui32(v91) < base.Ui32(v95) {
		goto L62
	} else {
		goto L63
	}
L60:
	;
	if v255 == int32(0) {
		goto L55
	} else {
		goto L96
	}
L61:
	;
	if base.Ui32(v233) <= base.Ui32(int32(3)) {
		v254 = v232
		v255 = v233
		v256 = v234
		goto L60
	} else {
		goto L92
	}
L62:
	;
	if v153 != 0 {
		goto L65
	} else {
		goto L66
	}
L63:
	;
	goto L64
L64:
	;
	if v153 != 0 {
		v215 = v138
		goto L75
	} else {
		goto L76
	}
L65:
	;
	v254 = v95
	v255 = v138
	v256 = v91
	goto L60
L66:
	;
	goto L67
L67:
	;
	if v91&int32(3) == int32(0) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v232 = v95
	v233 = v138
	v234 = v91
	goto L61
L69:
	;
	goto L70
L70:
	;
	v160 = v95
	v161 = v138
	v162 = v91
	goto L71
L71:
	;
	if v161 == int32(0) {
		goto L55
	} else {
		goto L73
	}
L72:
	;
	v232 = v169
	v233 = v171
	v234 = v173
	goto L61
L73:
	;
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160))))
	*(*uint8)(unsafe.Add(mBase, uint32(v162))) = uint8(v166)
	v168 = int32(1)
	v169 = v160 + v168
	v171 = v161 - v168
	v173 = v162 + v168
	if v173&int32(3) != 0 {
		v160 = v169
		v161 = v171
		v162 = v173
		goto L71
	} else {
		goto L74
	}
L74:
	;
	goto L72
L75:
	;
	if v215 == int32(0) {
		goto L55
	} else {
		goto L88
	}
L76:
	;
	if v143&int32(3) != 0 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v180 = v138
	goto L80
L78:
	;
	v195 = v138
	goto L79
L79:
	;
	if base.Ui32(v195) <= base.Ui32(int32(3)) {
		v215 = v195
		goto L75
	} else {
		goto L84
	}
L80:
	;
	if v180 == int32(0) {
		goto L55
	} else {
		goto L82
	}
L81:
	;
	v195 = v186
	goto L79
L82:
	;
	v186 = v180 - int32(1)
	v187 = v91 + v186
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95+v186))))
	*(*uint8)(unsafe.Add(mBase, uint32(v187))) = uint8(v189)
	if v187&int32(3) != 0 {
		v180 = v186
		goto L80
	} else {
		goto L83
	}
L83:
	;
	goto L81
L84:
	;
	v202 = v195
	goto L85
L85:
	;
	v206 = v202 - int32(4)
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v95+v206)))
	*(*int32)(unsafe.Add(mBase, uint32(v91+v206))) = v209
	if base.Ui32(int32(3)) < base.Ui32(v206) {
		v202 = v206
		goto L85
	} else {
		goto L87
	}
L86:
	;
	v215 = v206
	goto L75
L87:
	;
	goto L86
L88:
	;
	v222 = v215
	goto L89
L89:
	;
	v226 = v222 - int32(1)
	v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95+v226))))
	*(*uint8)(unsafe.Add(mBase, uint32(v91+v226))) = uint8(v229)
	if v226 != 0 {
		v222 = v226
		goto L89
	} else {
		goto L91
	}
L90:
	;
	goto L55
L91:
	;
	goto L90
L92:
	;
	v239 = v232
	v240 = v233
	v241 = v234
	goto L93
L93:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v239)))
	*(*int32)(unsafe.Add(mBase, uint32(v241))) = v243
	v245 = int32(4)
	v246 = v239 + v245
	v248 = v241 + v245
	v250 = v240 - v245
	if base.Ui32(int32(3)) < base.Ui32(v250) {
		v239 = v246
		v240 = v250
		v241 = v248
		goto L93
	} else {
		goto L95
	}
L94:
	;
	v254 = v246
	v255 = v250
	v256 = v248
	goto L60
L95:
	;
	goto L94
L96:
	;
	v261 = v254
	v262 = v255
	v263 = v256
	goto L97
L97:
	;
	v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v261))))
	*(*uint8)(unsafe.Add(mBase, uint32(v263))) = uint8(v265)
	v267 = int32(1)
	v272 = v262 - v267
	if v272 != 0 {
		v261 = v261 + v267
		v262 = v272
		v263 = v263 + v267
		goto L97
	} else {
		goto L99
	}
L98:
	;
	goto L55
L99:
	;
	goto L98
L100:
	;
	v288 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v91))) = uint8(v288)
	v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v290 != 0 {
		goto L103
	} else {
		goto L104
	}
L101:
	;
	goto L102
L102:
	;
	v357 = int32(47)
	*(*uint8)(unsafe.Add(mBase, uint32(v91))) = uint8(v357)
	v359 = F_strlen(m, v95)
	mBase = m.M
	v361 = v91 + int32(1)
	if v95 != v361 {
		goto L124
	} else {
		goto L125
	}
L103:
	;
	v291 = F_strlen(m, l0)
	mBase = m.M
	v295 = v291 + l0
	goto L106
L104:
	;
	v346 = l0
	goto L105
L105:
	;
	v354 = v96 - int32(1)
	v1341 = v346
	v1346 = v354
	v1347 = base.B2i32(v354 != int32(0))
	goto L37
L106:
	;
	v302 = v295 - int32(1)
	v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v302))))
	if base.B2i32(v303 == int32(47))&base.B2i32(base.Ui32(l0) < base.Ui32(v302)) != 0 {
		v295 = v302
		goto L106
	} else {
		goto L108
	}
L107:
	;
	v310 = v302
	goto L109
L108:
	;
	goto L107
L109:
	;
	if base.Ui32(l0) < base.Ui32(v310) {
		goto L111
	} else {
		goto L112
	}
L110:
	;
	v326 = v310
	goto L115
L111:
	;
	v319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310))))
	if v319 != int32(47) {
		v310 = v310 - int32(1)
		goto L109
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	goto L110
L114:
	;
	goto L113
L115:
	;
	if base.Ui32(l0) < base.Ui32(v326) {
		goto L117
	} else {
		goto L118
	}
L116:
	;
	if l0 == v326 {
		goto L121
	} else {
		goto L122
	}
L117:
	;
	v333 = v326 - int32(1)
	v334 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v333))))
	if v334 == int32(47) {
		v326 = v333
		goto L115
	} else {
		goto L120
	}
L118:
	;
	goto L119
L119:
	;
	goto L116
L120:
	;
	goto L119
L121:
	;
	v342 = l0 + base.B2i32(v290 == int32(47))
	goto L123
L122:
	;
	v342 = v326
	goto L123
L123:
	;
	v343 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v342))) = uint8(v343)
	v346 = v342
	goto L105
L124:
	;
	if v361 == v95 {
		goto L128
	} else {
		goto L129
	}
L125:
	;
	goto L126
L126:
	;
	v507 = int32(1)
	v1341 = v359 + v361
	v1346 = v96 + v507
	v1347 = v507
	goto L37
L127:
	;
	goto L126
L128:
	;
	goto L127
L129:
	;
	v366 = v361 + v359
	if base.Ui32(v95-v366) <= base.Ui32(int32(0)-v359<<(uint(int32(1))%32)) {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v373 = F___memcpy(m, v361, v95, v359)
	mBase = m.M
	goto L127
L131:
	;
	goto L132
L132:
	;
	v376 = (v361 ^ v95) & int32(3)
	if base.Ui32(v361) < base.Ui32(v95) {
		goto L135
	} else {
		goto L136
	}
L133:
	;
	if v478 == int32(0) {
		goto L128
	} else {
		goto L169
	}
L134:
	;
	if base.Ui32(v456) <= base.Ui32(int32(3)) {
		v477 = v455
		v478 = v456
		v479 = v457
		goto L133
	} else {
		goto L165
	}
L135:
	;
	if v376 != 0 {
		goto L138
	} else {
		goto L139
	}
L136:
	;
	goto L137
L137:
	;
	if v376 != 0 {
		v438 = v359
		goto L148
	} else {
		goto L149
	}
L138:
	;
	v477 = v95
	v478 = v359
	v479 = v361
	goto L133
L139:
	;
	goto L140
L140:
	;
	if v361&int32(3) == int32(0) {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	v455 = v95
	v456 = v359
	v457 = v361
	goto L134
L142:
	;
	goto L143
L143:
	;
	v383 = v95
	v384 = v359
	v385 = v361
	goto L144
L144:
	;
	if v384 == int32(0) {
		goto L128
	} else {
		goto L146
	}
L145:
	;
	v455 = v392
	v456 = v394
	v457 = v396
	goto L134
L146:
	;
	v389 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v383))))
	*(*uint8)(unsafe.Add(mBase, uint32(v385))) = uint8(v389)
	v391 = int32(1)
	v392 = v383 + v391
	v394 = v384 - v391
	v396 = v385 + v391
	if v396&int32(3) != 0 {
		v383 = v392
		v384 = v394
		v385 = v396
		goto L144
	} else {
		goto L147
	}
L147:
	;
	goto L145
L148:
	;
	if v438 == int32(0) {
		goto L128
	} else {
		goto L161
	}
L149:
	;
	if v366&int32(3) != 0 {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	v403 = v359
	goto L153
L151:
	;
	v418 = v359
	goto L152
L152:
	;
	if base.Ui32(v418) <= base.Ui32(int32(3)) {
		v438 = v418
		goto L148
	} else {
		goto L157
	}
L153:
	;
	if v403 == int32(0) {
		goto L128
	} else {
		goto L155
	}
L154:
	;
	v418 = v409
	goto L152
L155:
	;
	v409 = v403 - int32(1)
	v410 = v361 + v409
	v412 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95+v409))))
	*(*uint8)(unsafe.Add(mBase, uint32(v410))) = uint8(v412)
	if v410&int32(3) != 0 {
		v403 = v409
		goto L153
	} else {
		goto L156
	}
L156:
	;
	goto L154
L157:
	;
	v425 = v418
	goto L158
L158:
	;
	v429 = v425 - int32(4)
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v95+v429)))
	*(*int32)(unsafe.Add(mBase, uint32(v361+v429))) = v432
	if base.Ui32(int32(3)) < base.Ui32(v429) {
		v425 = v429
		goto L158
	} else {
		goto L160
	}
L159:
	;
	v438 = v429
	goto L148
L160:
	;
	goto L159
L161:
	;
	v445 = v438
	goto L162
L162:
	;
	v449 = v445 - int32(1)
	v452 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95+v449))))
	*(*uint8)(unsafe.Add(mBase, uint32(v361+v449))) = uint8(v452)
	if v449 != 0 {
		v445 = v449
		goto L162
	} else {
		goto L164
	}
L163:
	;
	goto L128
L164:
	;
	goto L163
L165:
	;
	v462 = v455
	v463 = v456
	v464 = v457
	goto L166
L166:
	;
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v462)))
	*(*int32)(unsafe.Add(mBase, uint32(v464))) = v466
	v468 = int32(4)
	v469 = v462 + v468
	v471 = v464 + v468
	v473 = v463 - v468
	if base.Ui32(int32(3)) < base.Ui32(v473) {
		v462 = v469
		v463 = v473
		v464 = v471
		goto L166
	} else {
		goto L168
	}
L167:
	;
	v477 = v469
	v478 = v473
	v479 = v471
	goto L133
L168:
	;
	goto L167
L169:
	;
	v484 = v477
	v485 = v478
	v486 = v479
	goto L170
L170:
	;
	v488 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v484))))
	*(*uint8)(unsafe.Add(mBase, uint32(v486))) = uint8(v488)
	v490 = int32(1)
	v495 = v485 - v490
	if v495 != 0 {
		v484 = v484 + v490
		v485 = v495
		v486 = v486 + v490
		goto L170
	} else {
		goto L172
	}
L171:
	;
	goto L128
L172:
	;
	goto L171
L173:
	;
	if v91 != v95 {
		goto L176
	} else {
		goto L177
	}
L174:
	;
	goto L175
L175:
	;
	if v91 != v95 {
		goto L225
	} else {
		goto L226
	}
L176:
	;
	if v91 == v95 {
		goto L180
	} else {
		goto L181
	}
L177:
	;
	goto L178
L178:
	;
	v1341 = v91 + v511
	v1346 = v96
	v1347 = int32(4)
	goto L37
L179:
	;
	goto L178
L180:
	;
	goto L179
L181:
	;
	v516 = v91 + v511
	if base.Ui32(v95-v516) <= base.Ui32(int32(0)-v511<<(uint(int32(1))%32)) {
		goto L182
	} else {
		goto L183
	}
L182:
	;
	v523 = F___memcpy(m, v91, v95, v511)
	mBase = m.M
	goto L179
L183:
	;
	goto L184
L184:
	;
	v526 = (v91 ^ v95) & int32(3)
	if base.Ui32(v91) < base.Ui32(v95) {
		goto L187
	} else {
		goto L188
	}
L185:
	;
	if v628 == int32(0) {
		goto L180
	} else {
		goto L221
	}
L186:
	;
	if base.Ui32(v606) <= base.Ui32(int32(3)) {
		v627 = v605
		v628 = v606
		v629 = v607
		goto L185
	} else {
		goto L217
	}
L187:
	;
	if v526 != 0 {
		goto L190
	} else {
		goto L191
	}
L188:
	;
	goto L189
L189:
	;
	if v526 != 0 {
		v588 = v511
		goto L200
	} else {
		goto L201
	}
L190:
	;
	v627 = v95
	v628 = v511
	v629 = v91
	goto L185
L191:
	;
	goto L192
L192:
	;
	if v91&int32(3) == int32(0) {
		goto L193
	} else {
		goto L194
	}
L193:
	;
	v605 = v95
	v606 = v511
	v607 = v91
	goto L186
L194:
	;
	goto L195
L195:
	;
	v533 = v95
	v534 = v511
	v535 = v91
	goto L196
L196:
	;
	if v534 == int32(0) {
		goto L180
	} else {
		goto L198
	}
L197:
	;
	v605 = v542
	v606 = v544
	v607 = v546
	goto L186
L198:
	;
	v539 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v533))))
	*(*uint8)(unsafe.Add(mBase, uint32(v535))) = uint8(v539)
	v541 = int32(1)
	v542 = v533 + v541
	v544 = v534 - v541
	v546 = v535 + v541
	if v546&int32(3) != 0 {
		v533 = v542
		v534 = v544
		v535 = v546
		goto L196
	} else {
		goto L199
	}
L199:
	;
	goto L197
L200:
	;
	if v588 == int32(0) {
		goto L180
	} else {
		goto L213
	}
L201:
	;
	if v516&int32(3) != 0 {
		goto L202
	} else {
		goto L203
	}
L202:
	;
	v553 = v511
	goto L205
L203:
	;
	v568 = v511
	goto L204
L204:
	;
	if base.Ui32(v568) <= base.Ui32(int32(3)) {
		v588 = v568
		goto L200
	} else {
		goto L209
	}
L205:
	;
	if v553 == int32(0) {
		goto L180
	} else {
		goto L207
	}
L206:
	;
	v568 = v559
	goto L204
L207:
	;
	v559 = v553 - int32(1)
	v560 = v91 + v559
	v562 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95+v559))))
	*(*uint8)(unsafe.Add(mBase, uint32(v560))) = uint8(v562)
	if v560&int32(3) != 0 {
		v553 = v559
		goto L205
	} else {
		goto L208
	}
L208:
	;
	goto L206
L209:
	;
	v575 = v568
	goto L210
L210:
	;
	v579 = v575 - int32(4)
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v95+v579)))
	*(*int32)(unsafe.Add(mBase, uint32(v91+v579))) = v582
	if base.Ui32(int32(3)) < base.Ui32(v579) {
		v575 = v579
		goto L210
	} else {
		goto L212
	}
L211:
	;
	v588 = v579
	goto L200
L212:
	;
	goto L211
L213:
	;
	v595 = v588
	goto L214
L214:
	;
	v599 = v595 - int32(1)
	v602 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95+v599))))
	*(*uint8)(unsafe.Add(mBase, uint32(v91+v599))) = uint8(v602)
	if v599 != 0 {
		v595 = v599
		goto L214
	} else {
		goto L216
	}
L215:
	;
	goto L180
L216:
	;
	goto L215
L217:
	;
	v612 = v605
	v613 = v606
	v614 = v607
	goto L218
L218:
	;
	v616 = *(*int32)(unsafe.Add(mBase, uint32(v612)))
	*(*int32)(unsafe.Add(mBase, uint32(v614))) = v616
	v618 = int32(4)
	v619 = v612 + v618
	v621 = v614 + v618
	v623 = v613 - v618
	if base.Ui32(int32(3)) < base.Ui32(v623) {
		v612 = v619
		v613 = v623
		v614 = v621
		goto L218
	} else {
		goto L220
	}
L219:
	;
	v627 = v619
	v628 = v623
	v629 = v621
	goto L185
L220:
	;
	goto L219
L221:
	;
	v634 = v627
	v635 = v628
	v636 = v629
	goto L222
L222:
	;
	v638 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v634))))
	*(*uint8)(unsafe.Add(mBase, uint32(v636))) = uint8(v638)
	v640 = int32(1)
	v645 = v635 - v640
	if v645 != 0 {
		v634 = v634 + v640
		v635 = v645
		v636 = v636 + v640
		goto L222
	} else {
		goto L224
	}
L223:
	;
	goto L180
L224:
	;
	goto L223
L225:
	;
	if v91 == v95 {
		goto L229
	} else {
		goto L230
	}
L226:
	;
	goto L227
L227:
	;
	v1332 = v91 + v511
	v1337 = v96 + int32(1)
	goto L42
L228:
	;
	goto L227
L229:
	;
	goto L228
L230:
	;
	v663 = v91 + v511
	if base.Ui32(v95-v663) <= base.Ui32(int32(0)-v511<<(uint(int32(1))%32)) {
		goto L231
	} else {
		goto L232
	}
L231:
	;
	v670 = F___memcpy(m, v91, v95, v511)
	mBase = m.M
	goto L228
L232:
	;
	goto L233
L233:
	;
	v673 = (v91 ^ v95) & int32(3)
	if base.Ui32(v91) < base.Ui32(v95) {
		goto L236
	} else {
		goto L237
	}
L234:
	;
	if v775 == int32(0) {
		goto L229
	} else {
		goto L270
	}
L235:
	;
	if base.Ui32(v753) <= base.Ui32(int32(3)) {
		v774 = v752
		v775 = v753
		v776 = v754
		goto L234
	} else {
		goto L266
	}
L236:
	;
	if v673 != 0 {
		goto L239
	} else {
		goto L240
	}
L237:
	;
	goto L238
L238:
	;
	if v673 != 0 {
		v735 = v511
		goto L249
	} else {
		goto L250
	}
L239:
	;
	v774 = v95
	v775 = v511
	v776 = v91
	goto L234
L240:
	;
	goto L241
L241:
	;
	if v91&int32(3) == int32(0) {
		goto L242
	} else {
		goto L243
	}
L242:
	;
	v752 = v95
	v753 = v511
	v754 = v91
	goto L235
L243:
	;
	goto L244
L244:
	;
	v680 = v95
	v681 = v511
	v682 = v91
	goto L245
L245:
	;
	if v681 == int32(0) {
		goto L229
	} else {
		goto L247
	}
L246:
	;
	v752 = v689
	v753 = v691
	v754 = v693
	goto L235
L247:
	;
	v686 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v680))))
	*(*uint8)(unsafe.Add(mBase, uint32(v682))) = uint8(v686)
	v688 = int32(1)
	v689 = v680 + v688
	v691 = v681 - v688
	v693 = v682 + v688
	if v693&int32(3) != 0 {
		v680 = v689
		v681 = v691
		v682 = v693
		goto L245
	} else {
		goto L248
	}
L248:
	;
	goto L246
L249:
	;
	if v735 == int32(0) {
		goto L229
	} else {
		goto L262
	}
L250:
	;
	if v663&int32(3) != 0 {
		goto L251
	} else {
		goto L252
	}
L251:
	;
	v700 = v511
	goto L254
L252:
	;
	v715 = v511
	goto L253
L253:
	;
	if base.Ui32(v715) <= base.Ui32(int32(3)) {
		v735 = v715
		goto L249
	} else {
		goto L258
	}
L254:
	;
	if v700 == int32(0) {
		goto L229
	} else {
		goto L256
	}
L255:
	;
	v715 = v706
	goto L253
L256:
	;
	v706 = v700 - int32(1)
	v707 = v91 + v706
	v709 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95+v706))))
	*(*uint8)(unsafe.Add(mBase, uint32(v707))) = uint8(v709)
	if v707&int32(3) != 0 {
		v700 = v706
		goto L254
	} else {
		goto L257
	}
L257:
	;
	goto L255
L258:
	;
	v722 = v715
	goto L259
L259:
	;
	v726 = v722 - int32(4)
	v729 = *(*int32)(unsafe.Add(mBase, uint32(v95+v726)))
	*(*int32)(unsafe.Add(mBase, uint32(v91+v726))) = v729
	if base.Ui32(int32(3)) < base.Ui32(v726) {
		v722 = v726
		goto L259
	} else {
		goto L261
	}
L260:
	;
	v735 = v726
	goto L249
L261:
	;
	goto L260
L262:
	;
	v742 = v735
	goto L263
L263:
	;
	v746 = v742 - int32(1)
	v749 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95+v746))))
	*(*uint8)(unsafe.Add(mBase, uint32(v91+v746))) = uint8(v749)
	if v746 != 0 {
		v742 = v746
		goto L263
	} else {
		goto L265
	}
L264:
	;
	goto L229
L265:
	;
	goto L264
L266:
	;
	v759 = v752
	v760 = v753
	v761 = v754
	goto L267
L267:
	;
	v763 = *(*int32)(unsafe.Add(mBase, uint32(v759)))
	*(*int32)(unsafe.Add(mBase, uint32(v761))) = v763
	v765 = int32(4)
	v766 = v759 + v765
	v768 = v761 + v765
	v770 = v760 - v765
	if base.Ui32(int32(3)) < base.Ui32(v770) {
		v759 = v766
		v760 = v770
		v761 = v768
		goto L267
	} else {
		goto L269
	}
L268:
	;
	v774 = v766
	v775 = v770
	v776 = v768
	goto L234
L269:
	;
	goto L268
L270:
	;
	v781 = v774
	v782 = v775
	v783 = v776
	goto L271
L271:
	;
	v785 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v781))))
	*(*uint8)(unsafe.Add(mBase, uint32(v783))) = uint8(v785)
	v787 = int32(1)
	v792 = v782 - v787
	if v792 != 0 {
		v781 = v781 + v787
		v782 = v792
		v783 = v783 + v787
		goto L271
	} else {
		goto L273
	}
L272:
	;
	goto L229
L273:
	;
	goto L272
L274:
	;
	v807 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v91))) = uint8(v807)
	v809 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v809 != 0 {
		goto L277
	} else {
		goto L278
	}
L275:
	;
	goto L276
L276:
	;
	v879 = int32(47)
	*(*uint8)(unsafe.Add(mBase, uint32(v91))) = uint8(v879)
	v881 = F_strlen(m, v95)
	mBase = m.M
	v883 = v91 + int32(1)
	if v95 != v883 {
		goto L302
	} else {
		goto L303
	}
L277:
	;
	v810 = F_strlen(m, l0)
	mBase = m.M
	v814 = v810 + l0
	goto L280
L278:
	;
	v865 = l0
	goto L279
L279:
	;
	v873 = v96 - int32(1)
	if v873 != 0 {
		v1332 = v865
		v1337 = v873
		goto L42
	} else {
		goto L298
	}
L280:
	;
	v821 = v814 - int32(1)
	v822 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v821))))
	if base.B2i32(v822 == int32(47))&base.B2i32(base.Ui32(l0) < base.Ui32(v821)) != 0 {
		v814 = v821
		goto L280
	} else {
		goto L282
	}
L281:
	;
	v829 = v821
	goto L283
L282:
	;
	goto L281
L283:
	;
	if base.Ui32(l0) < base.Ui32(v829) {
		goto L285
	} else {
		goto L286
	}
L284:
	;
	v845 = v829
	goto L289
L285:
	;
	v838 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v829))))
	if v838 != int32(47) {
		v829 = v829 - int32(1)
		goto L283
	} else {
		goto L288
	}
L286:
	;
	goto L287
L287:
	;
	goto L284
L288:
	;
	goto L287
L289:
	;
	if base.Ui32(l0) < base.Ui32(v845) {
		goto L291
	} else {
		goto L292
	}
L290:
	;
	if l0 == v845 {
		goto L295
	} else {
		goto L296
	}
L291:
	;
	v852 = v845 - int32(1)
	v853 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v852))))
	if v853 == int32(47) {
		v845 = v852
		goto L289
	} else {
		goto L294
	}
L292:
	;
	goto L293
L293:
	;
	goto L290
L294:
	;
	goto L293
L295:
	;
	v861 = l0 + base.B2i32(v809 == int32(47))
	goto L297
L296:
	;
	v861 = v845
	goto L297
L297:
	;
	v862 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v861))) = uint8(v862)
	v865 = v861
	goto L279
L298:
	;
	if l0 == v865 {
		goto L299
	} else {
		goto L300
	}
L299:
	;
	v877 = int32(2)
	goto L301
L300:
	;
	v877 = int32(4)
	goto L301
L301:
	;
	v1341 = v865
	v1346 = int32(0)
	v1347 = v877
	goto L37
L302:
	;
	if v883 == v95 {
		goto L306
	} else {
		goto L307
	}
L303:
	;
	goto L304
L304:
	;
	v1332 = v881 + v883
	v1337 = v96 + int32(1)
	goto L42
L305:
	;
	goto L304
L306:
	;
	goto L305
L307:
	;
	v888 = v883 + v881
	if base.Ui32(v95-v888) <= base.Ui32(int32(0)-v881<<(uint(int32(1))%32)) {
		goto L308
	} else {
		goto L309
	}
L308:
	;
	v895 = F___memcpy(m, v883, v95, v881)
	mBase = m.M
	goto L305
L309:
	;
	goto L310
L310:
	;
	v898 = (v883 ^ v95) & int32(3)
	if base.Ui32(v883) < base.Ui32(v95) {
		goto L313
	} else {
		goto L314
	}
L311:
	;
	if v1000 == int32(0) {
		goto L306
	} else {
		goto L347
	}
L312:
	;
	if base.Ui32(v978) <= base.Ui32(int32(3)) {
		v999 = v977
		v1000 = v978
		v1001 = v979
		goto L311
	} else {
		goto L343
	}
L313:
	;
	if v898 != 0 {
		goto L316
	} else {
		goto L317
	}
L314:
	;
	goto L315
L315:
	;
	if v898 != 0 {
		v960 = v881
		goto L326
	} else {
		goto L327
	}
L316:
	;
	v999 = v95
	v1000 = v881
	v1001 = v883
	goto L311
L317:
	;
	goto L318
L318:
	;
	if v883&int32(3) == int32(0) {
		goto L319
	} else {
		goto L320
	}
L319:
	;
	v977 = v95
	v978 = v881
	v979 = v883
	goto L312
L320:
	;
	goto L321
L321:
	;
	v905 = v95
	v906 = v881
	v907 = v883
	goto L322
L322:
	;
	if v906 == int32(0) {
		goto L306
	} else {
		goto L324
	}
L323:
	;
	v977 = v914
	v978 = v916
	v979 = v918
	goto L312
L324:
	;
	v911 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v905))))
	*(*uint8)(unsafe.Add(mBase, uint32(v907))) = uint8(v911)
	v913 = int32(1)
	v914 = v905 + v913
	v916 = v906 - v913
	v918 = v907 + v913
	if v918&int32(3) != 0 {
		v905 = v914
		v906 = v916
		v907 = v918
		goto L322
	} else {
		goto L325
	}
L325:
	;
	goto L323
L326:
	;
	if v960 == int32(0) {
		goto L306
	} else {
		goto L339
	}
L327:
	;
	if v888&int32(3) != 0 {
		goto L328
	} else {
		goto L329
	}
L328:
	;
	v925 = v881
	goto L331
L329:
	;
	v940 = v881
	goto L330
L330:
	;
	if base.Ui32(v940) <= base.Ui32(int32(3)) {
		v960 = v940
		goto L326
	} else {
		goto L335
	}
L331:
	;
	if v925 == int32(0) {
		goto L306
	} else {
		goto L333
	}
L332:
	;
	v940 = v931
	goto L330
L333:
	;
	v931 = v925 - int32(1)
	v932 = v883 + v931
	v934 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95+v931))))
	*(*uint8)(unsafe.Add(mBase, uint32(v932))) = uint8(v934)
	if v932&int32(3) != 0 {
		v925 = v931
		goto L331
	} else {
		goto L334
	}
L334:
	;
	goto L332
L335:
	;
	v947 = v940
	goto L336
L336:
	;
	v951 = v947 - int32(4)
	v954 = *(*int32)(unsafe.Add(mBase, uint32(v95+v951)))
	*(*int32)(unsafe.Add(mBase, uint32(v883+v951))) = v954
	if base.Ui32(int32(3)) < base.Ui32(v951) {
		v947 = v951
		goto L336
	} else {
		goto L338
	}
L337:
	;
	v960 = v951
	goto L326
L338:
	;
	goto L337
L339:
	;
	v967 = v960
	goto L340
L340:
	;
	v971 = v967 - int32(1)
	v974 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95+v971))))
	*(*uint8)(unsafe.Add(mBase, uint32(v883+v971))) = uint8(v974)
	if v971 != 0 {
		v967 = v971
		goto L340
	} else {
		goto L342
	}
L341:
	;
	goto L306
L342:
	;
	goto L341
L343:
	;
	v984 = v977
	v985 = v978
	v986 = v979
	goto L344
L344:
	;
	v988 = *(*int32)(unsafe.Add(mBase, uint32(v984)))
	*(*int32)(unsafe.Add(mBase, uint32(v986))) = v988
	v990 = int32(4)
	v991 = v984 + v990
	v993 = v986 + v990
	v995 = v985 - v990
	if base.Ui32(int32(3)) < base.Ui32(v995) {
		v984 = v991
		v985 = v995
		v986 = v993
		goto L344
	} else {
		goto L346
	}
L345:
	;
	v999 = v991
	v1000 = v995
	v1001 = v993
	goto L311
L346:
	;
	goto L345
L347:
	;
	v1006 = v999
	v1007 = v1000
	v1008 = v1001
	goto L348
L348:
	;
	v1010 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1006))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1008))) = uint8(v1010)
	v1012 = int32(1)
	v1017 = v1007 - v1012
	if v1017 != 0 {
		v1006 = v1006 + v1012
		v1007 = v1017
		v1008 = v1008 + v1012
		goto L348
	} else {
		goto L350
	}
L349:
	;
	goto L306
L350:
	;
	goto L349
L351:
	;
	if v95 != v1035 {
		goto L354
	} else {
		goto L355
	}
L352:
	;
	goto L353
L353:
	;
	if v95 != v1035 {
		goto L403
	} else {
		goto L404
	}
L354:
	;
	if v1035 == v95 {
		goto L358
	} else {
		goto L359
	}
L355:
	;
	goto L356
L356:
	;
	v1341 = v1036 + v1035
	v1346 = v96
	v1347 = int32(4)
	goto L37
L357:
	;
	goto L356
L358:
	;
	goto L357
L359:
	;
	v1041 = v1035 + v1036
	if base.Ui32(v95-v1041) <= base.Ui32(int32(0)-v1036<<(uint(int32(1))%32)) {
		goto L360
	} else {
		goto L361
	}
L360:
	;
	v1048 = F___memcpy(m, v1035, v95, v1036)
	mBase = m.M
	goto L357
L361:
	;
	goto L362
L362:
	;
	v1051 = (v1035 ^ v95) & int32(3)
	if base.Ui32(v1035) < base.Ui32(v95) {
		goto L365
	} else {
		goto L366
	}
L363:
	;
	if v1153 == int32(0) {
		goto L358
	} else {
		goto L399
	}
L364:
	;
	if base.Ui32(v1131) <= base.Ui32(int32(3)) {
		v1152 = v1130
		v1153 = v1131
		v1154 = v1132
		goto L363
	} else {
		goto L395
	}
L365:
	;
	if v1051 != 0 {
		goto L368
	} else {
		goto L369
	}
L366:
	;
	goto L367
L367:
	;
	if v1051 != 0 {
		v1113 = v1036
		goto L378
	} else {
		goto L379
	}
L368:
	;
	v1152 = v95
	v1153 = v1036
	v1154 = v1035
	goto L363
L369:
	;
	goto L370
L370:
	;
	if v1035&int32(3) == int32(0) {
		goto L371
	} else {
		goto L372
	}
L371:
	;
	v1130 = v95
	v1131 = v1036
	v1132 = v1035
	goto L364
L372:
	;
	goto L373
L373:
	;
	v1058 = v95
	v1059 = v1036
	v1060 = v1035
	goto L374
L374:
	;
	if v1059 == int32(0) {
		goto L358
	} else {
		goto L376
	}
L375:
	;
	v1130 = v1067
	v1131 = v1069
	v1132 = v1071
	goto L364
L376:
	;
	v1064 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1058))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1060))) = uint8(v1064)
	v1066 = int32(1)
	v1067 = v1058 + v1066
	v1069 = v1059 - v1066
	v1071 = v1060 + v1066
	if v1071&int32(3) != 0 {
		v1058 = v1067
		v1059 = v1069
		v1060 = v1071
		goto L374
	} else {
		goto L377
	}
L377:
	;
	goto L375
L378:
	;
	if v1113 == int32(0) {
		goto L358
	} else {
		goto L391
	}
L379:
	;
	if v1041&int32(3) != 0 {
		goto L380
	} else {
		goto L381
	}
L380:
	;
	v1078 = v1036
	goto L383
L381:
	;
	v1093 = v1036
	goto L382
L382:
	;
	if base.Ui32(v1093) <= base.Ui32(int32(3)) {
		v1113 = v1093
		goto L378
	} else {
		goto L387
	}
L383:
	;
	if v1078 == int32(0) {
		goto L358
	} else {
		goto L385
	}
L384:
	;
	v1093 = v1084
	goto L382
L385:
	;
	v1084 = v1078 - int32(1)
	v1085 = v1035 + v1084
	v1087 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95+v1084))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1085))) = uint8(v1087)
	if v1085&int32(3) != 0 {
		v1078 = v1084
		goto L383
	} else {
		goto L386
	}
L386:
	;
	goto L384
L387:
	;
	v1100 = v1093
	goto L388
L388:
	;
	v1104 = v1100 - int32(4)
	v1107 = *(*int32)(unsafe.Add(mBase, uint32(v95+v1104)))
	*(*int32)(unsafe.Add(mBase, uint32(v1035+v1104))) = v1107
	if base.Ui32(int32(3)) < base.Ui32(v1104) {
		v1100 = v1104
		goto L388
	} else {
		goto L390
	}
L389:
	;
	v1113 = v1104
	goto L378
L390:
	;
	goto L389
L391:
	;
	v1120 = v1113
	goto L392
L392:
	;
	v1124 = v1120 - int32(1)
	v1127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95+v1124))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1035+v1124))) = uint8(v1127)
	if v1124 != 0 {
		v1120 = v1124
		goto L392
	} else {
		goto L394
	}
L393:
	;
	goto L358
L394:
	;
	goto L393
L395:
	;
	v1137 = v1130
	v1138 = v1131
	v1139 = v1132
	goto L396
L396:
	;
	v1141 = *(*int32)(unsafe.Add(mBase, uint32(v1137)))
	*(*int32)(unsafe.Add(mBase, uint32(v1139))) = v1141
	v1143 = int32(4)
	v1144 = v1137 + v1143
	v1146 = v1139 + v1143
	v1148 = v1138 - v1143
	if base.Ui32(int32(3)) < base.Ui32(v1148) {
		v1137 = v1144
		v1138 = v1148
		v1139 = v1146
		goto L396
	} else {
		goto L398
	}
L397:
	;
	v1152 = v1144
	v1153 = v1148
	v1154 = v1146
	goto L363
L398:
	;
	goto L397
L399:
	;
	v1159 = v1152
	v1160 = v1153
	v1161 = v1154
	goto L400
L400:
	;
	v1163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1159))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1161))) = uint8(v1163)
	v1165 = int32(1)
	v1170 = v1160 - v1165
	if v1170 != 0 {
		v1159 = v1159 + v1165
		v1160 = v1170
		v1161 = v1161 + v1165
		goto L400
	} else {
		goto L402
	}
L401:
	;
	goto L358
L402:
	;
	goto L401
L403:
	;
	if v1035 == v95 {
		goto L407
	} else {
		goto L408
	}
L404:
	;
	goto L405
L405:
	;
	v1332 = v1036 + v1035
	v1337 = int32(1)
	goto L42
L406:
	;
	goto L405
L407:
	;
	goto L406
L408:
	;
	v1188 = v1035 + v1036
	if base.Ui32(v95-v1188) <= base.Ui32(int32(0)-v1036<<(uint(int32(1))%32)) {
		goto L409
	} else {
		goto L410
	}
L409:
	;
	v1195 = F___memcpy(m, v1035, v95, v1036)
	mBase = m.M
	goto L406
L410:
	;
	goto L411
L411:
	;
	v1198 = (v1035 ^ v95) & int32(3)
	if base.Ui32(v1035) < base.Ui32(v95) {
		goto L414
	} else {
		goto L415
	}
L412:
	;
	if v1300 == int32(0) {
		goto L407
	} else {
		goto L448
	}
L413:
	;
	if base.Ui32(v1278) <= base.Ui32(int32(3)) {
		v1299 = v1277
		v1300 = v1278
		v1301 = v1279
		goto L412
	} else {
		goto L444
	}
L414:
	;
	if v1198 != 0 {
		goto L417
	} else {
		goto L418
	}
L415:
	;
	goto L416
L416:
	;
	if v1198 != 0 {
		v1260 = v1036
		goto L427
	} else {
		goto L428
	}
L417:
	;
	v1299 = v95
	v1300 = v1036
	v1301 = v1035
	goto L412
L418:
	;
	goto L419
L419:
	;
	if v1035&int32(3) == int32(0) {
		goto L420
	} else {
		goto L421
	}
L420:
	;
	v1277 = v95
	v1278 = v1036
	v1279 = v1035
	goto L413
L421:
	;
	goto L422
L422:
	;
	v1205 = v95
	v1206 = v1036
	v1207 = v1035
	goto L423
L423:
	;
	if v1206 == int32(0) {
		goto L407
	} else {
		goto L425
	}
L424:
	;
	v1277 = v1214
	v1278 = v1216
	v1279 = v1218
	goto L413
L425:
	;
	v1211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1205))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1207))) = uint8(v1211)
	v1213 = int32(1)
	v1214 = v1205 + v1213
	v1216 = v1206 - v1213
	v1218 = v1207 + v1213
	if v1218&int32(3) != 0 {
		v1205 = v1214
		v1206 = v1216
		v1207 = v1218
		goto L423
	} else {
		goto L426
	}
L426:
	;
	goto L424
L427:
	;
	if v1260 == int32(0) {
		goto L407
	} else {
		goto L440
	}
L428:
	;
	if v1188&int32(3) != 0 {
		goto L429
	} else {
		goto L430
	}
L429:
	;
	v1225 = v1036
	goto L432
L430:
	;
	v1240 = v1036
	goto L431
L431:
	;
	if base.Ui32(v1240) <= base.Ui32(int32(3)) {
		v1260 = v1240
		goto L427
	} else {
		goto L436
	}
L432:
	;
	if v1225 == int32(0) {
		goto L407
	} else {
		goto L434
	}
L433:
	;
	v1240 = v1231
	goto L431
L434:
	;
	v1231 = v1225 - int32(1)
	v1232 = v1035 + v1231
	v1234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95+v1231))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1232))) = uint8(v1234)
	if v1232&int32(3) != 0 {
		v1225 = v1231
		goto L432
	} else {
		goto L435
	}
L435:
	;
	goto L433
L436:
	;
	v1247 = v1240
	goto L437
L437:
	;
	v1251 = v1247 - int32(4)
	v1254 = *(*int32)(unsafe.Add(mBase, uint32(v95+v1251)))
	*(*int32)(unsafe.Add(mBase, uint32(v1035+v1251))) = v1254
	if base.Ui32(int32(3)) < base.Ui32(v1251) {
		v1247 = v1251
		goto L437
	} else {
		goto L439
	}
L438:
	;
	v1260 = v1251
	goto L427
L439:
	;
	goto L438
L440:
	;
	v1267 = v1260
	goto L441
L441:
	;
	v1271 = v1267 - int32(1)
	v1274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95+v1271))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1035+v1271))) = uint8(v1274)
	if v1271 != 0 {
		v1267 = v1271
		goto L441
	} else {
		goto L443
	}
L442:
	;
	goto L407
L443:
	;
	goto L442
L444:
	;
	v1284 = v1277
	v1285 = v1278
	v1286 = v1279
	goto L445
L445:
	;
	v1288 = *(*int32)(unsafe.Add(mBase, uint32(v1284)))
	*(*int32)(unsafe.Add(mBase, uint32(v1286))) = v1288
	v1290 = int32(4)
	v1291 = v1284 + v1290
	v1293 = v1286 + v1290
	v1295 = v1285 - v1290
	if base.Ui32(int32(3)) < base.Ui32(v1295) {
		v1284 = v1291
		v1285 = v1295
		v1286 = v1293
		goto L445
	} else {
		goto L447
	}
L446:
	;
	v1299 = v1291
	v1300 = v1295
	v1301 = v1293
	goto L412
L447:
	;
	goto L446
L448:
	;
	v1306 = v1299
	v1307 = v1300
	v1308 = v1301
	goto L449
L449:
	;
	v1310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1306))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1308))) = uint8(v1310)
	v1312 = int32(1)
	v1317 = v1307 - v1312
	if v1317 != 0 {
		v1306 = v1306 + v1312
		v1307 = v1317
		v1308 = v1308 + v1312
		goto L449
	} else {
		goto L451
	}
L450:
	;
	goto L407
L451:
	;
	goto L450
L452:
	;
	goto L29
L453:
	;
	v1358 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v1350))) = uint8(v1358)
	v1362 = v1350 + int32(1)
	goto L455
L454:
	;
	v1362 = v1350
	goto L455
L455:
	;
	v1363 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1362))) = uint8(v1363)
	goto L20
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
