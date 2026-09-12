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
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v292 int32
	_ = v292
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v307 int32
	_ = v307
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v327 int32
	_ = v327
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v384 int32
	_ = v384
	var v396 int32
	_ = v396
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v410 int32
	_ = v410
	var v415 int32
	_ = v415
	var v419 int32
	_ = v419
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v430 int32
	_ = v430
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v445 int32
	_ = v445
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v459 int32
	_ = v459
	var v463 int32
	_ = v463
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v478 int32
	_ = v478
	var v487 int32
	_ = v487
	var v494 int32
	_ = v494
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v514 int32
	_ = v514
	var v522 int32
	_ = v522
	var v525 int32
	_ = v525
	var v534 int32
	_ = v534
	var v539 int32
	_ = v539
	var v543 int32
	_ = v543
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v554 int32
	_ = v554
	var v560 int32
	_ = v560
	var v563 int32
	_ = v563
	var v569 int32
	_ = v569
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v583 int32
	_ = v583
	var v585 int32
	_ = v585
	var v590 int32
	_ = v590
	var v597 int32
	_ = v597
	var v600 int32
	_ = v600
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v618 int32
	_ = v618
	var v620 int32
	_ = v620
	var v627 int32
	_ = v627
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v636 int32
	_ = v636
	var v642 int32
	_ = v642
	var v649 int32
	_ = v649
	var v653 int32
	_ = v653
	var v656 int32
	_ = v656
	var v662 int32
	_ = v662
	var v669 int32
	_ = v669
	var v673 int32
	_ = v673
	var v676 int32
	_ = v676
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v690 int32
	_ = v690
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v695 int32
	_ = v695
	var v697 int32
	_ = v697
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v712 int32
	_ = v712
	var v714 int32
	_ = v714
	var v719 int32
	_ = v719
	var v731 int32
	_ = v731
	var v742 int32
	_ = v742
	var v747 int32
	_ = v747
	var v751 int32
	_ = v751
	var v756 int32
	_ = v756
	var v758 int32
	_ = v758
	var v762 int32
	_ = v762
	var v768 int32
	_ = v768
	var v771 int32
	_ = v771
	var v777 int32
	_ = v777
	var v781 int32
	_ = v781
	var v783 int32
	_ = v783
	var v791 int32
	_ = v791
	var v796 int32
	_ = v796
	var v803 int32
	_ = v803
	var v806 int32
	_ = v806
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v819 int32
	_ = v819
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v824 int32
	_ = v824
	var v826 int32
	_ = v826
	var v833 int32
	_ = v833
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v842 int32
	_ = v842
	var v848 int32
	_ = v848
	var v855 int32
	_ = v855
	var v859 int32
	_ = v859
	var v862 int32
	_ = v862
	var v868 int32
	_ = v868
	var v875 int32
	_ = v875
	var v879 int32
	_ = v879
	var v882 int32
	_ = v882
	var v885 int32
	_ = v885
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v896 int32
	_ = v896
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v901 int32
	_ = v901
	var v903 int32
	_ = v903
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v918 int32
	_ = v918
	var v920 int32
	_ = v920
	var v925 int32
	_ = v925
	var v943 int32
	_ = v943
	var v950 int32
	_ = v950
	var v953 int32
	_ = v953
	var v960 int32
	_ = v960
	var v961 int32
	_ = v961
	var v962 int32
	_ = v962
	var v966 int32
	_ = v966
	var v968 int32
	_ = v968
	var v969 int32
	_ = v969
	var v971 int32
	_ = v971
	var v973 int32
	_ = v973
	var v980 int32
	_ = v980
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v989 int32
	_ = v989
	var v995 int32
	_ = v995
	var v1002 int32
	_ = v1002
	var v1006 int32
	_ = v1006
	var v1009 int32
	_ = v1009
	var v1015 int32
	_ = v1015
	var v1022 int32
	_ = v1022
	var v1026 int32
	_ = v1026
	var v1029 int32
	_ = v1029
	var v1032 int32
	_ = v1032
	var v1033 int32
	_ = v1033
	var v1034 int32
	_ = v1034
	var v1039 int32
	_ = v1039
	var v1040 int32
	_ = v1040
	var v1041 int32
	_ = v1041
	var v1043 int32
	_ = v1043
	var v1045 int32
	_ = v1045
	var v1046 int32
	_ = v1046
	var v1048 int32
	_ = v1048
	var v1050 int32
	_ = v1050
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	var v1056 int32
	_ = v1056
	var v1061 int32
	_ = v1061
	var v1062 int32
	_ = v1062
	var v1063 int32
	_ = v1063
	var v1065 int32
	_ = v1065
	var v1067 int32
	_ = v1067
	var v1072 int32
	_ = v1072
	var v1087 int32
	_ = v1087
	var v1089 int32
	_ = v1089
	var v1097 int32
	_ = v1097
	var v1102 int32
	_ = v1102
	var v1106 int32
	_ = v1106
	var v1111 int32
	_ = v1111
	var v1113 int32
	_ = v1113
	var v1117 int32
	_ = v1117
	var v1123 int32
	_ = v1123
	var v1126 int32
	_ = v1126
	var v1132 int32
	_ = v1132
	var v1136 int32
	_ = v1136
	var v1138 int32
	_ = v1138
	var v1146 int32
	_ = v1146
	var v1150 int32
	_ = v1150
	var v1157 int32
	_ = v1157
	var v1158 int32
	_ = v1158
	var v1165 int32
	_ = v1165
	var v1174 int32
	_ = v1174
	var v1181 int32
	_ = v1181
	var v1188 int32
	_ = v1188
	var v1189 int32
	_ = v1189
	var v1197 int32
	_ = v1197
	var v1198 int32
	_ = v1198
	var v1201 int32
	_ = v1201
	var v1209 int32
	_ = v1209
	var v1213 int32
	_ = v1213
	var v1215 int32
	_ = v1215
	var v1224 int32
	_ = v1224
	var v1229 int32
	_ = v1229
	var v1233 int32
	_ = v1233
	var v1238 int32
	_ = v1238
	var v1240 int32
	_ = v1240
	var v1244 int32
	_ = v1244
	var v1250 int32
	_ = v1250
	var v1253 int32
	_ = v1253
	var v1259 int32
	_ = v1259
	var v1263 int32
	_ = v1263
	var v1265 int32
	_ = v1265
	var v1273 int32
	_ = v1273
	var v1275 int32
	_ = v1275
	var v1280 int32
	_ = v1280
	var v1287 int32
	_ = v1287
	var v1290 int32
	_ = v1290
	var v1297 int32
	_ = v1297
	var v1298 int32
	_ = v1298
	var v1299 int32
	_ = v1299
	var v1303 int32
	_ = v1303
	var v1305 int32
	_ = v1305
	var v1306 int32
	_ = v1306
	var v1308 int32
	_ = v1308
	var v1310 int32
	_ = v1310
	var v1317 int32
	_ = v1317
	var v1323 int32
	_ = v1323
	var v1324 int32
	_ = v1324
	var v1326 int32
	_ = v1326
	var v1332 int32
	_ = v1332
	var v1339 int32
	_ = v1339
	var v1343 int32
	_ = v1343
	var v1346 int32
	_ = v1346
	var v1352 int32
	_ = v1352
	var v1359 int32
	_ = v1359
	var v1363 int32
	_ = v1363
	var v1366 int32
	_ = v1366
	var v1369 int32
	_ = v1369
	var v1370 int32
	_ = v1370
	var v1371 int32
	_ = v1371
	var v1376 int32
	_ = v1376
	var v1377 int32
	_ = v1377
	var v1378 int32
	_ = v1378
	var v1380 int32
	_ = v1380
	var v1382 int32
	_ = v1382
	var v1383 int32
	_ = v1383
	var v1385 int32
	_ = v1385
	var v1387 int32
	_ = v1387
	var v1391 int32
	_ = v1391
	var v1392 int32
	_ = v1392
	var v1393 int32
	_ = v1393
	var v1398 int32
	_ = v1398
	var v1399 int32
	_ = v1399
	var v1400 int32
	_ = v1400
	var v1402 int32
	_ = v1402
	var v1404 int32
	_ = v1404
	var v1409 int32
	_ = v1409
	var v1424 int32
	_ = v1424
	var v1427 int32
	_ = v1427
	var v1435 int32
	_ = v1435
	var v1440 int32
	_ = v1440
	var v1444 int32
	_ = v1444
	var v1449 int32
	_ = v1449
	var v1451 int32
	_ = v1451
	var v1455 int32
	_ = v1455
	var v1461 int32
	_ = v1461
	var v1464 int32
	_ = v1464
	var v1470 int32
	_ = v1470
	var v1474 int32
	_ = v1474
	var v1476 int32
	_ = v1476
	var v1484 int32
	_ = v1484
	var v1489 int32
	_ = v1489
	var v1496 int32
	_ = v1496
	var v1499 int32
	_ = v1499
	var v1506 int32
	_ = v1506
	var v1507 int32
	_ = v1507
	var v1508 int32
	_ = v1508
	var v1512 int32
	_ = v1512
	var v1514 int32
	_ = v1514
	var v1515 int32
	_ = v1515
	var v1517 int32
	_ = v1517
	var v1519 int32
	_ = v1519
	var v1526 int32
	_ = v1526
	var v1532 int32
	_ = v1532
	var v1533 int32
	_ = v1533
	var v1535 int32
	_ = v1535
	var v1541 int32
	_ = v1541
	var v1548 int32
	_ = v1548
	var v1552 int32
	_ = v1552
	var v1555 int32
	_ = v1555
	var v1561 int32
	_ = v1561
	var v1568 int32
	_ = v1568
	var v1572 int32
	_ = v1572
	var v1575 int32
	_ = v1575
	var v1578 int32
	_ = v1578
	var v1579 int32
	_ = v1579
	var v1580 int32
	_ = v1580
	var v1585 int32
	_ = v1585
	var v1586 int32
	_ = v1586
	var v1587 int32
	_ = v1587
	var v1589 int32
	_ = v1589
	var v1591 int32
	_ = v1591
	var v1592 int32
	_ = v1592
	var v1594 int32
	_ = v1594
	var v1596 int32
	_ = v1596
	var v1600 int32
	_ = v1600
	var v1601 int32
	_ = v1601
	var v1602 int32
	_ = v1602
	var v1607 int32
	_ = v1607
	var v1608 int32
	_ = v1608
	var v1609 int32
	_ = v1609
	var v1611 int32
	_ = v1611
	var v1613 int32
	_ = v1613
	var v1618 int32
	_ = v1618
	var v1636 int32
	_ = v1636
	var v1643 int32
	_ = v1643
	var v1646 int32
	_ = v1646
	var v1653 int32
	_ = v1653
	var v1654 int32
	_ = v1654
	var v1655 int32
	_ = v1655
	var v1659 int32
	_ = v1659
	var v1661 int32
	_ = v1661
	var v1662 int32
	_ = v1662
	var v1664 int32
	_ = v1664
	var v1666 int32
	_ = v1666
	var v1673 int32
	_ = v1673
	var v1679 int32
	_ = v1679
	var v1680 int32
	_ = v1680
	var v1682 int32
	_ = v1682
	var v1688 int32
	_ = v1688
	var v1695 int32
	_ = v1695
	var v1699 int32
	_ = v1699
	var v1702 int32
	_ = v1702
	var v1708 int32
	_ = v1708
	var v1715 int32
	_ = v1715
	var v1719 int32
	_ = v1719
	var v1722 int32
	_ = v1722
	var v1725 int32
	_ = v1725
	var v1726 int32
	_ = v1726
	var v1727 int32
	_ = v1727
	var v1732 int32
	_ = v1732
	var v1733 int32
	_ = v1733
	var v1734 int32
	_ = v1734
	var v1736 int32
	_ = v1736
	var v1738 int32
	_ = v1738
	var v1739 int32
	_ = v1739
	var v1741 int32
	_ = v1741
	var v1743 int32
	_ = v1743
	var v1747 int32
	_ = v1747
	var v1748 int32
	_ = v1748
	var v1749 int32
	_ = v1749
	var v1754 int32
	_ = v1754
	var v1755 int32
	_ = v1755
	var v1756 int32
	_ = v1756
	var v1758 int32
	_ = v1758
	var v1760 int32
	_ = v1760
	var v1765 int32
	_ = v1765
	var v1780 int32
	_ = v1780
	var v1785 int32
	_ = v1785
	var v1789 int32
	_ = v1789
	var v1794 int32
	_ = v1794
	var v1795 int32
	_ = v1795
	var v1796 int32
	_ = v1796
	var v1798 int32
	_ = v1798
	var v1806 int32
	_ = v1806
	var v1810 int32
	_ = v1810
	var v1811 int32
	_ = v1811
	v2 = int32(0)
	if l0&int32(3) == v2 {
		v32 = l0
		goto L4
	} else {
		goto L5
	}
L1:
	;
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v95 != 0 {
		goto L25
	} else {
		goto L26
	}
L2:
	;
	v66 = v65 + l0
	if base.Ui32(v66) <= base.Ui32(l0) {
		goto L1
	} else {
		goto L19
	}
L3:
	;
	v65 = v57 - l0
	goto L2
L4:
	;
	v36 = v32
	goto L13
L5:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v16 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v65 = int32(0)
	goto L2
L7:
	;
	goto L8
L8:
	;
	v21 = l0
	goto L9
L9:
	;
	v25 = v21 + int32(1)
	if v25&int32(3) == int32(0) {
		v32 = v25
		goto L4
	} else {
		goto L11
	}
L10:
	;
	v57 = v25
	goto L3
L11:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	if v30 != 0 {
		v21 = v25
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	v45 = int32(-2139062144)
	if (int32(16843008)-v42|v42)&v45 == v45 {
		v36 = v36 + int32(4)
		goto L13
	} else {
		goto L15
	}
L14:
	;
	v51 = v36
	goto L16
L15:
	;
	goto L14
L16:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	if v55 != 0 {
		v51 = v51 + int32(1)
		goto L16
	} else {
		goto L18
	}
L17:
	;
	v57 = v51
	goto L3
L18:
	;
	goto L17
L19:
	;
	v69 = v66 - int32(1)
	if base.Ui32(v69) <= base.Ui32(l0) {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v76 = v69
	goto L21
L21:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76))))
	if v79 != int32(47) {
		goto L1
	} else {
		goto L23
	}
L22:
	;
	goto L1
L23:
	;
	v82 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v76))) = uint8(v82)
	v85 = v76 - int32(1)
	if base.Ui32(l0) < base.Ui32(v85) {
		v76 = v85
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	v97 = l0
	v98 = l0
	v100 = v2
	goto L28
L26:
	;
	v122 = l0
	goto L27
L27:
	;
	v129 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v122))) = uint8(v129)
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v131 != 0 {
		goto L35
	} else {
		goto L36
	}
L28:
	;
	v105 = v98 + int32(1)
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98))))
	if base.B2i32(v106 == int32(47))&v100 != 0 {
		v98 = v105
		goto L28
	} else {
		goto L30
	}
L29:
	;
	v122 = v115
	goto L27
L30:
	;
	if v97 != v98 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v97))) = uint8(v106)
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98))))
	v113 = v112
	goto L33
L32:
	;
	v113 = v106
	goto L33
L33:
	;
	v115 = v97 + int32(1)
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105))))
	if v120 != 0 {
		v97 = v115
		v98 = v105
		v100 = base.B2i32(v113&int32(255) == int32(47))
		goto L28
	} else {
		goto L34
	}
L34:
	;
	goto L29
L35:
	;
	if v131 == int32(47) {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	goto L37
L37:
	;
	return
L38:
	;
	v137 = l0 + int32(1)
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137))))
	v139 = v138
	v140 = v137
	v141 = int32(0)
	goto L40
L39:
	;
	v139 = v131
	v140 = l0
	v141 = int32(2)
	goto L40
L40:
	;
	if v139&int32(255) == int32(0) {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	if l0 == v1798 {
		goto L589
	} else {
		goto L590
	}
L42:
	;
	v1798 = v140
	goto L41
L43:
	;
	goto L44
L44:
	;
	v147 = v140
	v150 = v139
	v151 = v140
	v152 = v2
	v153 = v141
	goto L45
L45:
	;
	v156 = v150
	v159 = v151
	goto L48
L46:
	;
	v1798 = v1789
	goto L41
L47:
	;
	v178 = int32(0)
	if v176&int32(255) != int32(46) {
		v192 = v178
		goto L55
	} else {
		goto L56
	}
L48:
	;
	v163 = v156 & int32(255)
	if v163 == int32(0) {
		v176 = v150
		v177 = v159
		goto L47
	} else {
		goto L50
	}
L49:
	;
	v171 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v159))) = uint8(v171)
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151))))
	v176 = v175
	v177 = v159 + int32(1)
	goto L47
L50:
	;
	if v163 != int32(47) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v169 = v159 + int32(1)
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169))))
	v156 = v170
	v159 = v169
	goto L48
L52:
	;
	goto L53
L53:
	;
	goto L49
L54:
	;
	v1796 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177))))
	if v1796 != 0 {
		v147 = v1789
		v150 = v1796
		v151 = v177
		v152 = v1794
		v153 = v1795
		goto L45
	} else {
		goto L588
	}
L55:
	;
	switch v153 {
	case 0:
		goto L64
	case 1:
		goto L63
	case 2:
		goto L62
	case 3:
		goto L61
	case 4:
		goto L60
	default:
		v1789 = v147
		v1794 = v152
		v1795 = v153
		goto L54
	}
L56:
	;
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+1)))
	if v183 == int32(0) {
		v1789 = v147
		v1794 = v152
		v1795 = v153
		goto L54
	} else {
		goto L57
	}
L57:
	;
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+1)))
	if v186 != int32(46) {
		v192 = v178
		goto L55
	} else {
		goto L58
	}
L58:
	;
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+2)))
	v192 = base.B2i32(v189 == int32(0))
	goto L55
L59:
	;
	v1789 = v1780
	v1794 = v1785
	v1795 = int32(3)
	goto L54
L60:
	;
	v1424 = int32(47)
	*(*uint8)(unsafe.Add(mBase, uint32(v147))) = uint8(v1424)
	v1427 = v147 + int32(1)
	if v151&int32(3) == int32(0) {
		v1451 = v151
		goto L472
	} else {
		goto L473
	}
L61:
	;
	if v192 != 0 {
		goto L359
	} else {
		goto L360
	}
L62:
	;
	if v151&int32(3) == int32(0) {
		v758 = v151
		goto L243
	} else {
		goto L244
	}
L63:
	;
	if v192 != 0 {
		goto L134
	} else {
		goto L135
	}
L64:
	;
	if v192 != 0 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v1789 = v147
	v1794 = v152
	v1795 = int32(0)
	goto L54
L66:
	;
	goto L67
L67:
	;
	if v151&int32(3) == int32(0) {
		v217 = v151
		goto L70
	} else {
		goto L71
	}
L68:
	;
	if v147 != v151 {
		goto L85
	} else {
		goto L86
	}
L69:
	;
	v250 = v242 - v151
	goto L68
L70:
	;
	v221 = v217
	goto L79
L71:
	;
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151))))
	if v201 == int32(0) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v250 = int32(0)
	goto L68
L73:
	;
	goto L74
L74:
	;
	v206 = v151
	goto L75
L75:
	;
	v210 = v206 + int32(1)
	if v210&int32(3) == int32(0) {
		v217 = v210
		goto L70
	} else {
		goto L77
	}
L76:
	;
	v242 = v210
	goto L69
L77:
	;
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210))))
	if v215 != 0 {
		v206 = v210
		goto L75
	} else {
		goto L78
	}
L78:
	;
	goto L76
L79:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v221)))
	v230 = int32(-2139062144)
	if (int32(16843008)-v227|v227)&v230 == v230 {
		v221 = v221 + int32(4)
		goto L79
	} else {
		goto L81
	}
L80:
	;
	v236 = v221
	goto L82
L81:
	;
	goto L80
L82:
	;
	v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v236))))
	if v240 != 0 {
		v236 = v236 + int32(1)
		goto L82
	} else {
		goto L84
	}
L83:
	;
	v242 = v236
	goto L69
L84:
	;
	goto L83
L85:
	;
	if v147 == v151 {
		goto L89
	} else {
		goto L90
	}
L86:
	;
	goto L87
L87:
	;
	v396 = int32(1)
	v1789 = v147 + v250
	v1794 = v152 + v396
	v1795 = v396
	goto L54
L88:
	;
	goto L87
L89:
	;
	goto L88
L90:
	;
	v255 = v147 + v250
	if base.Ui32(v151-v255) <= base.Ui32(int32(0)-v250<<(uint(int32(1))%32)) {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v262 = F___memcpy(m, v147, v151, v250)
	mBase = m.M
	goto L88
L92:
	;
	goto L93
L93:
	;
	v265 = (v147 ^ v151) & int32(3)
	if base.Ui32(v147) < base.Ui32(v151) {
		goto L96
	} else {
		goto L97
	}
L94:
	;
	if v367 == int32(0) {
		goto L89
	} else {
		goto L130
	}
L95:
	;
	if base.Ui32(v345) <= base.Ui32(int32(3)) {
		v366 = v344
		v367 = v345
		v368 = v346
		goto L94
	} else {
		goto L126
	}
L96:
	;
	if v265 != 0 {
		goto L99
	} else {
		goto L100
	}
L97:
	;
	goto L98
L98:
	;
	if v265 != 0 {
		v327 = v250
		goto L109
	} else {
		goto L110
	}
L99:
	;
	v366 = v151
	v367 = v250
	v368 = v147
	goto L94
L100:
	;
	goto L101
L101:
	;
	if v147&int32(3) == int32(0) {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v344 = v151
	v345 = v250
	v346 = v147
	goto L95
L103:
	;
	goto L104
L104:
	;
	v272 = v151
	v273 = v250
	v274 = v147
	goto L105
L105:
	;
	if v273 == int32(0) {
		goto L89
	} else {
		goto L107
	}
L106:
	;
	v344 = v281
	v345 = v283
	v346 = v285
	goto L95
L107:
	;
	v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272))))
	*(*uint8)(unsafe.Add(mBase, uint32(v274))) = uint8(v278)
	v280 = int32(1)
	v281 = v272 + v280
	v283 = v273 - v280
	v285 = v274 + v280
	if v285&int32(3) != 0 {
		v272 = v281
		v273 = v283
		v274 = v285
		goto L105
	} else {
		goto L108
	}
L108:
	;
	goto L106
L109:
	;
	if v327 == int32(0) {
		goto L89
	} else {
		goto L122
	}
L110:
	;
	if v255&int32(3) != 0 {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v292 = v250
	goto L114
L112:
	;
	v307 = v250
	goto L113
L113:
	;
	if base.Ui32(v307) <= base.Ui32(int32(3)) {
		v327 = v307
		goto L109
	} else {
		goto L118
	}
L114:
	;
	if v292 == int32(0) {
		goto L89
	} else {
		goto L116
	}
L115:
	;
	v307 = v298
	goto L113
L116:
	;
	v298 = v292 - int32(1)
	v299 = v147 + v298
	v301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151+v298))))
	*(*uint8)(unsafe.Add(mBase, uint32(v299))) = uint8(v301)
	if v299&int32(3) != 0 {
		v292 = v298
		goto L114
	} else {
		goto L117
	}
L117:
	;
	goto L115
L118:
	;
	v314 = v307
	goto L119
L119:
	;
	v318 = v314 - int32(4)
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v151+v318)))
	*(*int32)(unsafe.Add(mBase, uint32(v147+v318))) = v321
	if base.Ui32(int32(3)) < base.Ui32(v318) {
		v314 = v318
		goto L119
	} else {
		goto L121
	}
L120:
	;
	v327 = v318
	goto L109
L121:
	;
	goto L120
L122:
	;
	v334 = v327
	goto L123
L123:
	;
	v338 = v334 - int32(1)
	v341 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151+v338))))
	*(*uint8)(unsafe.Add(mBase, uint32(v147+v338))) = uint8(v341)
	if v338 != 0 {
		v334 = v338
		goto L123
	} else {
		goto L125
	}
L124:
	;
	goto L89
L125:
	;
	goto L124
L126:
	;
	v351 = v344
	v352 = v345
	v353 = v346
	goto L127
L127:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v351)))
	*(*int32)(unsafe.Add(mBase, uint32(v353))) = v355
	v357 = int32(4)
	v358 = v351 + v357
	v360 = v353 + v357
	v362 = v352 - v357
	if base.Ui32(int32(3)) < base.Ui32(v362) {
		v351 = v358
		v352 = v362
		v353 = v360
		goto L127
	} else {
		goto L129
	}
L128:
	;
	v366 = v358
	v367 = v362
	v368 = v360
	goto L94
L129:
	;
	goto L128
L130:
	;
	v373 = v366
	v374 = v367
	v375 = v368
	goto L131
L131:
	;
	v377 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v373))))
	*(*uint8)(unsafe.Add(mBase, uint32(v375))) = uint8(v377)
	v379 = int32(1)
	v384 = v374 - v379
	if v384 != 0 {
		v373 = v373 + v379
		v374 = v384
		v375 = v375 + v379
		goto L131
	} else {
		goto L133
	}
L132:
	;
	goto L89
L133:
	;
	goto L132
L134:
	;
	v400 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v147))) = uint8(v400)
	v402 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v402 != 0 {
		goto L137
	} else {
		goto L138
	}
L135:
	;
	goto L136
L136:
	;
	v525 = int32(47)
	*(*uint8)(unsafe.Add(mBase, uint32(v147))) = uint8(v525)
	if v151&int32(3) == int32(0) {
		v550 = v151
		goto L177
	} else {
		goto L178
	}
L137:
	;
	if l0&int32(3) == int32(0) {
		v426 = l0
		goto L142
	} else {
		goto L143
	}
L138:
	;
	v514 = l0
	goto L139
L139:
	;
	v522 = v152 - int32(1)
	v1789 = v514
	v1794 = v522
	v1795 = base.B2i32(v522 != int32(0))
	goto L54
L140:
	;
	v463 = v459 + l0
	goto L157
L141:
	;
	v459 = v451 - l0
	goto L140
L142:
	;
	v430 = v426
	goto L151
L143:
	;
	v410 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v410 == int32(0) {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v459 = int32(0)
	goto L140
L145:
	;
	goto L146
L146:
	;
	v415 = l0
	goto L147
L147:
	;
	v419 = v415 + int32(1)
	if v419&int32(3) == int32(0) {
		v426 = v419
		goto L142
	} else {
		goto L149
	}
L148:
	;
	v451 = v419
	goto L141
L149:
	;
	v424 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v419))))
	if v424 != 0 {
		v415 = v419
		goto L147
	} else {
		goto L150
	}
L150:
	;
	goto L148
L151:
	;
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v430)))
	v439 = int32(-2139062144)
	if (int32(16843008)-v436|v436)&v439 == v439 {
		v430 = v430 + int32(4)
		goto L151
	} else {
		goto L153
	}
L152:
	;
	v445 = v430
	goto L154
L153:
	;
	goto L152
L154:
	;
	v449 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v445))))
	if v449 != 0 {
		v445 = v445 + int32(1)
		goto L154
	} else {
		goto L156
	}
L155:
	;
	v451 = v445
	goto L141
L156:
	;
	goto L155
L157:
	;
	v470 = v463 - int32(1)
	v471 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v470))))
	if base.B2i32(v471 == int32(47))&base.B2i32(base.Ui32(l0) < base.Ui32(v470)) != 0 {
		v463 = v470
		goto L157
	} else {
		goto L159
	}
L158:
	;
	v478 = v470
	goto L160
L159:
	;
	goto L158
L160:
	;
	if base.Ui32(l0) < base.Ui32(v478) {
		goto L162
	} else {
		goto L163
	}
L161:
	;
	v494 = v478
	goto L166
L162:
	;
	v487 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v478))))
	if v487 != int32(47) {
		v478 = v478 - int32(1)
		goto L160
	} else {
		goto L165
	}
L163:
	;
	goto L164
L164:
	;
	goto L161
L165:
	;
	goto L164
L166:
	;
	if base.Ui32(l0) < base.Ui32(v494) {
		goto L168
	} else {
		goto L169
	}
L167:
	;
	if l0 == v494 {
		goto L172
	} else {
		goto L173
	}
L168:
	;
	v501 = v494 - int32(1)
	v502 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v501))))
	if v502 == int32(47) {
		v494 = v501
		goto L166
	} else {
		goto L171
	}
L169:
	;
	goto L170
L170:
	;
	goto L167
L171:
	;
	goto L170
L172:
	;
	v510 = l0 + base.B2i32(v402 == int32(47))
	goto L174
L173:
	;
	v510 = v494
	goto L174
L174:
	;
	v511 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v510))) = uint8(v511)
	v514 = v510
	goto L139
L175:
	;
	v585 = v147 + int32(1)
	if v151 != v585 {
		goto L192
	} else {
		goto L193
	}
L176:
	;
	v583 = v575 - v151
	goto L175
L177:
	;
	v554 = v550
	goto L186
L178:
	;
	v534 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151))))
	if v534 == int32(0) {
		goto L179
	} else {
		goto L180
	}
L179:
	;
	v583 = int32(0)
	goto L175
L180:
	;
	goto L181
L181:
	;
	v539 = v151
	goto L182
L182:
	;
	v543 = v539 + int32(1)
	if v543&int32(3) == int32(0) {
		v550 = v543
		goto L177
	} else {
		goto L184
	}
L183:
	;
	v575 = v543
	goto L176
L184:
	;
	v548 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v543))))
	if v548 != 0 {
		v539 = v543
		goto L182
	} else {
		goto L185
	}
L185:
	;
	goto L183
L186:
	;
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v554)))
	v563 = int32(-2139062144)
	if (int32(16843008)-v560|v560)&v563 == v563 {
		v554 = v554 + int32(4)
		goto L186
	} else {
		goto L188
	}
L187:
	;
	v569 = v554
	goto L189
L188:
	;
	goto L187
L189:
	;
	v573 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v569))))
	if v573 != 0 {
		v569 = v569 + int32(1)
		goto L189
	} else {
		goto L191
	}
L190:
	;
	v575 = v569
	goto L176
L191:
	;
	goto L190
L192:
	;
	if v585 == v151 {
		goto L196
	} else {
		goto L197
	}
L193:
	;
	goto L194
L194:
	;
	v731 = int32(1)
	v1789 = v583 + v585
	v1794 = v152 + v731
	v1795 = v731
	goto L54
L195:
	;
	goto L194
L196:
	;
	goto L195
L197:
	;
	v590 = v585 + v583
	if base.Ui32(v151-v590) <= base.Ui32(int32(0)-v583<<(uint(int32(1))%32)) {
		goto L198
	} else {
		goto L199
	}
L198:
	;
	v597 = F___memcpy(m, v585, v151, v583)
	mBase = m.M
	goto L195
L199:
	;
	goto L200
L200:
	;
	v600 = (v585 ^ v151) & int32(3)
	if base.Ui32(v585) < base.Ui32(v151) {
		goto L203
	} else {
		goto L204
	}
L201:
	;
	if v702 == int32(0) {
		goto L196
	} else {
		goto L237
	}
L202:
	;
	if base.Ui32(v680) <= base.Ui32(int32(3)) {
		v701 = v679
		v702 = v680
		v703 = v681
		goto L201
	} else {
		goto L233
	}
L203:
	;
	if v600 != 0 {
		goto L206
	} else {
		goto L207
	}
L204:
	;
	goto L205
L205:
	;
	if v600 != 0 {
		v662 = v583
		goto L216
	} else {
		goto L217
	}
L206:
	;
	v701 = v151
	v702 = v583
	v703 = v585
	goto L201
L207:
	;
	goto L208
L208:
	;
	if v585&int32(3) == int32(0) {
		goto L209
	} else {
		goto L210
	}
L209:
	;
	v679 = v151
	v680 = v583
	v681 = v585
	goto L202
L210:
	;
	goto L211
L211:
	;
	v607 = v151
	v608 = v583
	v609 = v585
	goto L212
L212:
	;
	if v608 == int32(0) {
		goto L196
	} else {
		goto L214
	}
L213:
	;
	v679 = v616
	v680 = v618
	v681 = v620
	goto L202
L214:
	;
	v613 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v607))))
	*(*uint8)(unsafe.Add(mBase, uint32(v609))) = uint8(v613)
	v615 = int32(1)
	v616 = v607 + v615
	v618 = v608 - v615
	v620 = v609 + v615
	if v620&int32(3) != 0 {
		v607 = v616
		v608 = v618
		v609 = v620
		goto L212
	} else {
		goto L215
	}
L215:
	;
	goto L213
L216:
	;
	if v662 == int32(0) {
		goto L196
	} else {
		goto L229
	}
L217:
	;
	if v590&int32(3) != 0 {
		goto L218
	} else {
		goto L219
	}
L218:
	;
	v627 = v583
	goto L221
L219:
	;
	v642 = v583
	goto L220
L220:
	;
	if base.Ui32(v642) <= base.Ui32(int32(3)) {
		v662 = v642
		goto L216
	} else {
		goto L225
	}
L221:
	;
	if v627 == int32(0) {
		goto L196
	} else {
		goto L223
	}
L222:
	;
	v642 = v633
	goto L220
L223:
	;
	v633 = v627 - int32(1)
	v634 = v585 + v633
	v636 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151+v633))))
	*(*uint8)(unsafe.Add(mBase, uint32(v634))) = uint8(v636)
	if v634&int32(3) != 0 {
		v627 = v633
		goto L221
	} else {
		goto L224
	}
L224:
	;
	goto L222
L225:
	;
	v649 = v642
	goto L226
L226:
	;
	v653 = v649 - int32(4)
	v656 = *(*int32)(unsafe.Add(mBase, uint32(v151+v653)))
	*(*int32)(unsafe.Add(mBase, uint32(v585+v653))) = v656
	if base.Ui32(int32(3)) < base.Ui32(v653) {
		v649 = v653
		goto L226
	} else {
		goto L228
	}
L227:
	;
	v662 = v653
	goto L216
L228:
	;
	goto L227
L229:
	;
	v669 = v662
	goto L230
L230:
	;
	v673 = v669 - int32(1)
	v676 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151+v673))))
	*(*uint8)(unsafe.Add(mBase, uint32(v585+v673))) = uint8(v676)
	if v673 != 0 {
		v669 = v673
		goto L230
	} else {
		goto L232
	}
L231:
	;
	goto L196
L232:
	;
	goto L231
L233:
	;
	v686 = v679
	v687 = v680
	v688 = v681
	goto L234
L234:
	;
	v690 = *(*int32)(unsafe.Add(mBase, uint32(v686)))
	*(*int32)(unsafe.Add(mBase, uint32(v688))) = v690
	v692 = int32(4)
	v693 = v686 + v692
	v695 = v688 + v692
	v697 = v687 - v692
	if base.Ui32(int32(3)) < base.Ui32(v697) {
		v686 = v693
		v687 = v697
		v688 = v695
		goto L234
	} else {
		goto L236
	}
L235:
	;
	v701 = v693
	v702 = v697
	v703 = v695
	goto L201
L236:
	;
	goto L235
L237:
	;
	v708 = v701
	v709 = v702
	v710 = v703
	goto L238
L238:
	;
	v712 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v708))))
	*(*uint8)(unsafe.Add(mBase, uint32(v710))) = uint8(v712)
	v714 = int32(1)
	v719 = v709 - v714
	if v719 != 0 {
		v708 = v708 + v714
		v709 = v719
		v710 = v710 + v714
		goto L238
	} else {
		goto L240
	}
L239:
	;
	goto L196
L240:
	;
	goto L239
L241:
	;
	if v192 != 0 {
		goto L258
	} else {
		goto L259
	}
L242:
	;
	v791 = v783 - v151
	goto L241
L243:
	;
	v762 = v758
	goto L252
L244:
	;
	v742 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151))))
	if v742 == int32(0) {
		goto L245
	} else {
		goto L246
	}
L245:
	;
	v791 = int32(0)
	goto L241
L246:
	;
	goto L247
L247:
	;
	v747 = v151
	goto L248
L248:
	;
	v751 = v747 + int32(1)
	if v751&int32(3) == int32(0) {
		v758 = v751
		goto L243
	} else {
		goto L250
	}
L249:
	;
	v783 = v751
	goto L242
L250:
	;
	v756 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v751))))
	if v756 != 0 {
		v747 = v751
		goto L248
	} else {
		goto L251
	}
L251:
	;
	goto L249
L252:
	;
	v768 = *(*int32)(unsafe.Add(mBase, uint32(v762)))
	v771 = int32(-2139062144)
	if (int32(16843008)-v768|v768)&v771 == v771 {
		v762 = v762 + int32(4)
		goto L252
	} else {
		goto L254
	}
L253:
	;
	v777 = v762
	goto L255
L254:
	;
	goto L253
L255:
	;
	v781 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v777))))
	if v781 != 0 {
		v777 = v777 + int32(1)
		goto L255
	} else {
		goto L257
	}
L256:
	;
	v783 = v777
	goto L242
L257:
	;
	goto L256
L258:
	;
	if v147 != v151 {
		goto L261
	} else {
		goto L262
	}
L259:
	;
	goto L260
L260:
	;
	if v147 != v151 {
		goto L310
	} else {
		goto L311
	}
L261:
	;
	if v147 == v151 {
		goto L265
	} else {
		goto L266
	}
L262:
	;
	goto L263
L263:
	;
	v1789 = v147 + v791
	v1794 = v152
	v1795 = int32(4)
	goto L54
L264:
	;
	goto L263
L265:
	;
	goto L264
L266:
	;
	v796 = v147 + v791
	if base.Ui32(v151-v796) <= base.Ui32(int32(0)-v791<<(uint(int32(1))%32)) {
		goto L267
	} else {
		goto L268
	}
L267:
	;
	v803 = F___memcpy(m, v147, v151, v791)
	mBase = m.M
	goto L264
L268:
	;
	goto L269
L269:
	;
	v806 = (v147 ^ v151) & int32(3)
	if base.Ui32(v147) < base.Ui32(v151) {
		goto L272
	} else {
		goto L273
	}
L270:
	;
	if v908 == int32(0) {
		goto L265
	} else {
		goto L306
	}
L271:
	;
	if base.Ui32(v886) <= base.Ui32(int32(3)) {
		v907 = v885
		v908 = v886
		v909 = v887
		goto L270
	} else {
		goto L302
	}
L272:
	;
	if v806 != 0 {
		goto L275
	} else {
		goto L276
	}
L273:
	;
	goto L274
L274:
	;
	if v806 != 0 {
		v868 = v791
		goto L285
	} else {
		goto L286
	}
L275:
	;
	v907 = v151
	v908 = v791
	v909 = v147
	goto L270
L276:
	;
	goto L277
L277:
	;
	if v147&int32(3) == int32(0) {
		goto L278
	} else {
		goto L279
	}
L278:
	;
	v885 = v151
	v886 = v791
	v887 = v147
	goto L271
L279:
	;
	goto L280
L280:
	;
	v813 = v151
	v814 = v791
	v815 = v147
	goto L281
L281:
	;
	if v814 == int32(0) {
		goto L265
	} else {
		goto L283
	}
L282:
	;
	v885 = v822
	v886 = v824
	v887 = v826
	goto L271
L283:
	;
	v819 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v813))))
	*(*uint8)(unsafe.Add(mBase, uint32(v815))) = uint8(v819)
	v821 = int32(1)
	v822 = v813 + v821
	v824 = v814 - v821
	v826 = v815 + v821
	if v826&int32(3) != 0 {
		v813 = v822
		v814 = v824
		v815 = v826
		goto L281
	} else {
		goto L284
	}
L284:
	;
	goto L282
L285:
	;
	if v868 == int32(0) {
		goto L265
	} else {
		goto L298
	}
L286:
	;
	if v796&int32(3) != 0 {
		goto L287
	} else {
		goto L288
	}
L287:
	;
	v833 = v791
	goto L290
L288:
	;
	v848 = v791
	goto L289
L289:
	;
	if base.Ui32(v848) <= base.Ui32(int32(3)) {
		v868 = v848
		goto L285
	} else {
		goto L294
	}
L290:
	;
	if v833 == int32(0) {
		goto L265
	} else {
		goto L292
	}
L291:
	;
	v848 = v839
	goto L289
L292:
	;
	v839 = v833 - int32(1)
	v840 = v147 + v839
	v842 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151+v839))))
	*(*uint8)(unsafe.Add(mBase, uint32(v840))) = uint8(v842)
	if v840&int32(3) != 0 {
		v833 = v839
		goto L290
	} else {
		goto L293
	}
L293:
	;
	goto L291
L294:
	;
	v855 = v848
	goto L295
L295:
	;
	v859 = v855 - int32(4)
	v862 = *(*int32)(unsafe.Add(mBase, uint32(v151+v859)))
	*(*int32)(unsafe.Add(mBase, uint32(v147+v859))) = v862
	if base.Ui32(int32(3)) < base.Ui32(v859) {
		v855 = v859
		goto L295
	} else {
		goto L297
	}
L296:
	;
	v868 = v859
	goto L285
L297:
	;
	goto L296
L298:
	;
	v875 = v868
	goto L299
L299:
	;
	v879 = v875 - int32(1)
	v882 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151+v879))))
	*(*uint8)(unsafe.Add(mBase, uint32(v147+v879))) = uint8(v882)
	if v879 != 0 {
		v875 = v879
		goto L299
	} else {
		goto L301
	}
L300:
	;
	goto L265
L301:
	;
	goto L300
L302:
	;
	v892 = v885
	v893 = v886
	v894 = v887
	goto L303
L303:
	;
	v896 = *(*int32)(unsafe.Add(mBase, uint32(v892)))
	*(*int32)(unsafe.Add(mBase, uint32(v894))) = v896
	v898 = int32(4)
	v899 = v892 + v898
	v901 = v894 + v898
	v903 = v893 - v898
	if base.Ui32(int32(3)) < base.Ui32(v903) {
		v892 = v899
		v893 = v903
		v894 = v901
		goto L303
	} else {
		goto L305
	}
L304:
	;
	v907 = v899
	v908 = v903
	v909 = v901
	goto L270
L305:
	;
	goto L304
L306:
	;
	v914 = v907
	v915 = v908
	v916 = v909
	goto L307
L307:
	;
	v918 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v914))))
	*(*uint8)(unsafe.Add(mBase, uint32(v916))) = uint8(v918)
	v920 = int32(1)
	v925 = v915 - v920
	if v925 != 0 {
		v914 = v914 + v920
		v915 = v925
		v916 = v916 + v920
		goto L307
	} else {
		goto L309
	}
L308:
	;
	goto L265
L309:
	;
	goto L308
L310:
	;
	if v147 == v151 {
		goto L314
	} else {
		goto L315
	}
L311:
	;
	goto L312
L312:
	;
	v1780 = v147 + v791
	v1785 = v152 + int32(1)
	goto L59
L313:
	;
	goto L312
L314:
	;
	goto L313
L315:
	;
	v943 = v147 + v791
	if base.Ui32(v151-v943) <= base.Ui32(int32(0)-v791<<(uint(int32(1))%32)) {
		goto L316
	} else {
		goto L317
	}
L316:
	;
	v950 = F___memcpy(m, v147, v151, v791)
	mBase = m.M
	goto L313
L317:
	;
	goto L318
L318:
	;
	v953 = (v147 ^ v151) & int32(3)
	if base.Ui32(v147) < base.Ui32(v151) {
		goto L321
	} else {
		goto L322
	}
L319:
	;
	if v1055 == int32(0) {
		goto L314
	} else {
		goto L355
	}
L320:
	;
	if base.Ui32(v1033) <= base.Ui32(int32(3)) {
		v1054 = v1032
		v1055 = v1033
		v1056 = v1034
		goto L319
	} else {
		goto L351
	}
L321:
	;
	if v953 != 0 {
		goto L324
	} else {
		goto L325
	}
L322:
	;
	goto L323
L323:
	;
	if v953 != 0 {
		v1015 = v791
		goto L334
	} else {
		goto L335
	}
L324:
	;
	v1054 = v151
	v1055 = v791
	v1056 = v147
	goto L319
L325:
	;
	goto L326
L326:
	;
	if v147&int32(3) == int32(0) {
		goto L327
	} else {
		goto L328
	}
L327:
	;
	v1032 = v151
	v1033 = v791
	v1034 = v147
	goto L320
L328:
	;
	goto L329
L329:
	;
	v960 = v151
	v961 = v791
	v962 = v147
	goto L330
L330:
	;
	if v961 == int32(0) {
		goto L314
	} else {
		goto L332
	}
L331:
	;
	v1032 = v969
	v1033 = v971
	v1034 = v973
	goto L320
L332:
	;
	v966 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v960))))
	*(*uint8)(unsafe.Add(mBase, uint32(v962))) = uint8(v966)
	v968 = int32(1)
	v969 = v960 + v968
	v971 = v961 - v968
	v973 = v962 + v968
	if v973&int32(3) != 0 {
		v960 = v969
		v961 = v971
		v962 = v973
		goto L330
	} else {
		goto L333
	}
L333:
	;
	goto L331
L334:
	;
	if v1015 == int32(0) {
		goto L314
	} else {
		goto L347
	}
L335:
	;
	if v943&int32(3) != 0 {
		goto L336
	} else {
		goto L337
	}
L336:
	;
	v980 = v791
	goto L339
L337:
	;
	v995 = v791
	goto L338
L338:
	;
	if base.Ui32(v995) <= base.Ui32(int32(3)) {
		v1015 = v995
		goto L334
	} else {
		goto L343
	}
L339:
	;
	if v980 == int32(0) {
		goto L314
	} else {
		goto L341
	}
L340:
	;
	v995 = v986
	goto L338
L341:
	;
	v986 = v980 - int32(1)
	v987 = v147 + v986
	v989 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151+v986))))
	*(*uint8)(unsafe.Add(mBase, uint32(v987))) = uint8(v989)
	if v987&int32(3) != 0 {
		v980 = v986
		goto L339
	} else {
		goto L342
	}
L342:
	;
	goto L340
L343:
	;
	v1002 = v995
	goto L344
L344:
	;
	v1006 = v1002 - int32(4)
	v1009 = *(*int32)(unsafe.Add(mBase, uint32(v151+v1006)))
	*(*int32)(unsafe.Add(mBase, uint32(v147+v1006))) = v1009
	if base.Ui32(int32(3)) < base.Ui32(v1006) {
		v1002 = v1006
		goto L344
	} else {
		goto L346
	}
L345:
	;
	v1015 = v1006
	goto L334
L346:
	;
	goto L345
L347:
	;
	v1022 = v1015
	goto L348
L348:
	;
	v1026 = v1022 - int32(1)
	v1029 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151+v1026))))
	*(*uint8)(unsafe.Add(mBase, uint32(v147+v1026))) = uint8(v1029)
	if v1026 != 0 {
		v1022 = v1026
		goto L348
	} else {
		goto L350
	}
L349:
	;
	goto L314
L350:
	;
	goto L349
L351:
	;
	v1039 = v1032
	v1040 = v1033
	v1041 = v1034
	goto L352
L352:
	;
	v1043 = *(*int32)(unsafe.Add(mBase, uint32(v1039)))
	*(*int32)(unsafe.Add(mBase, uint32(v1041))) = v1043
	v1045 = int32(4)
	v1046 = v1039 + v1045
	v1048 = v1041 + v1045
	v1050 = v1040 - v1045
	if base.Ui32(int32(3)) < base.Ui32(v1050) {
		v1039 = v1046
		v1040 = v1050
		v1041 = v1048
		goto L352
	} else {
		goto L354
	}
L353:
	;
	v1054 = v1046
	v1055 = v1050
	v1056 = v1048
	goto L319
L354:
	;
	goto L353
L355:
	;
	v1061 = v1054
	v1062 = v1055
	v1063 = v1056
	goto L356
L356:
	;
	v1065 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1061))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1063))) = uint8(v1065)
	v1067 = int32(1)
	v1072 = v1062 - v1067
	if v1072 != 0 {
		v1061 = v1061 + v1067
		v1062 = v1072
		v1063 = v1063 + v1067
		goto L356
	} else {
		goto L358
	}
L357:
	;
	goto L314
L358:
	;
	goto L357
L359:
	;
	v1087 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v147))) = uint8(v1087)
	v1089 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v1089 != 0 {
		goto L362
	} else {
		goto L363
	}
L360:
	;
	goto L361
L361:
	;
	v1215 = int32(47)
	*(*uint8)(unsafe.Add(mBase, uint32(v147))) = uint8(v1215)
	if v151&int32(3) == int32(0) {
		v1240 = v151
		goto L406
	} else {
		goto L407
	}
L362:
	;
	if l0&int32(3) == int32(0) {
		v1113 = l0
		goto L367
	} else {
		goto L368
	}
L363:
	;
	v1201 = l0
	goto L364
L364:
	;
	v1209 = v152 - int32(1)
	if v1209 != 0 {
		v1780 = v1201
		v1785 = v1209
		goto L59
	} else {
		goto L400
	}
L365:
	;
	v1150 = v1146 + l0
	goto L382
L366:
	;
	v1146 = v1138 - l0
	goto L365
L367:
	;
	v1117 = v1113
	goto L376
L368:
	;
	v1097 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v1097 == int32(0) {
		goto L369
	} else {
		goto L370
	}
L369:
	;
	v1146 = int32(0)
	goto L365
L370:
	;
	goto L371
L371:
	;
	v1102 = l0
	goto L372
L372:
	;
	v1106 = v1102 + int32(1)
	if v1106&int32(3) == int32(0) {
		v1113 = v1106
		goto L367
	} else {
		goto L374
	}
L373:
	;
	v1138 = v1106
	goto L366
L374:
	;
	v1111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1106))))
	if v1111 != 0 {
		v1102 = v1106
		goto L372
	} else {
		goto L375
	}
L375:
	;
	goto L373
L376:
	;
	v1123 = *(*int32)(unsafe.Add(mBase, uint32(v1117)))
	v1126 = int32(-2139062144)
	if (int32(16843008)-v1123|v1123)&v1126 == v1126 {
		v1117 = v1117 + int32(4)
		goto L376
	} else {
		goto L378
	}
L377:
	;
	v1132 = v1117
	goto L379
L378:
	;
	goto L377
L379:
	;
	v1136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1132))))
	if v1136 != 0 {
		v1132 = v1132 + int32(1)
		goto L379
	} else {
		goto L381
	}
L380:
	;
	v1138 = v1132
	goto L366
L381:
	;
	goto L380
L382:
	;
	v1157 = v1150 - int32(1)
	v1158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1157))))
	if base.B2i32(v1158 == int32(47))&base.B2i32(base.Ui32(l0) < base.Ui32(v1157)) != 0 {
		v1150 = v1157
		goto L382
	} else {
		goto L384
	}
L383:
	;
	v1165 = v1157
	goto L385
L384:
	;
	goto L383
L385:
	;
	if base.Ui32(l0) < base.Ui32(v1165) {
		goto L387
	} else {
		goto L388
	}
L386:
	;
	v1181 = v1165
	goto L391
L387:
	;
	v1174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1165))))
	if v1174 != int32(47) {
		v1165 = v1165 - int32(1)
		goto L385
	} else {
		goto L390
	}
L388:
	;
	goto L389
L389:
	;
	goto L386
L390:
	;
	goto L389
L391:
	;
	if base.Ui32(l0) < base.Ui32(v1181) {
		goto L393
	} else {
		goto L394
	}
L392:
	;
	if l0 == v1181 {
		goto L397
	} else {
		goto L398
	}
L393:
	;
	v1188 = v1181 - int32(1)
	v1189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1188))))
	if v1189 == int32(47) {
		v1181 = v1188
		goto L391
	} else {
		goto L396
	}
L394:
	;
	goto L395
L395:
	;
	goto L392
L396:
	;
	goto L395
L397:
	;
	v1197 = l0 + base.B2i32(v1089 == int32(47))
	goto L399
L398:
	;
	v1197 = v1181
	goto L399
L399:
	;
	v1198 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1197))) = uint8(v1198)
	v1201 = v1197
	goto L364
L400:
	;
	if l0 == v1201 {
		goto L401
	} else {
		goto L402
	}
L401:
	;
	v1213 = int32(2)
	goto L403
L402:
	;
	v1213 = int32(4)
	goto L403
L403:
	;
	v1789 = v1201
	v1794 = int32(0)
	v1795 = v1213
	goto L54
L404:
	;
	v1275 = v147 + int32(1)
	if v151 != v1275 {
		goto L421
	} else {
		goto L422
	}
L405:
	;
	v1273 = v1265 - v151
	goto L404
L406:
	;
	v1244 = v1240
	goto L415
L407:
	;
	v1224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151))))
	if v1224 == int32(0) {
		goto L408
	} else {
		goto L409
	}
L408:
	;
	v1273 = int32(0)
	goto L404
L409:
	;
	goto L410
L410:
	;
	v1229 = v151
	goto L411
L411:
	;
	v1233 = v1229 + int32(1)
	if v1233&int32(3) == int32(0) {
		v1240 = v1233
		goto L406
	} else {
		goto L413
	}
L412:
	;
	v1265 = v1233
	goto L405
L413:
	;
	v1238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1233))))
	if v1238 != 0 {
		v1229 = v1233
		goto L411
	} else {
		goto L414
	}
L414:
	;
	goto L412
L415:
	;
	v1250 = *(*int32)(unsafe.Add(mBase, uint32(v1244)))
	v1253 = int32(-2139062144)
	if (int32(16843008)-v1250|v1250)&v1253 == v1253 {
		v1244 = v1244 + int32(4)
		goto L415
	} else {
		goto L417
	}
L416:
	;
	v1259 = v1244
	goto L418
L417:
	;
	goto L416
L418:
	;
	v1263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1259))))
	if v1263 != 0 {
		v1259 = v1259 + int32(1)
		goto L418
	} else {
		goto L420
	}
L419:
	;
	v1265 = v1259
	goto L405
L420:
	;
	goto L419
L421:
	;
	if v1275 == v151 {
		goto L425
	} else {
		goto L426
	}
L422:
	;
	goto L423
L423:
	;
	v1780 = v1273 + v1275
	v1785 = v152 + int32(1)
	goto L59
L424:
	;
	goto L423
L425:
	;
	goto L424
L426:
	;
	v1280 = v1275 + v1273
	if base.Ui32(v151-v1280) <= base.Ui32(int32(0)-v1273<<(uint(int32(1))%32)) {
		goto L427
	} else {
		goto L428
	}
L427:
	;
	v1287 = F___memcpy(m, v1275, v151, v1273)
	mBase = m.M
	goto L424
L428:
	;
	goto L429
L429:
	;
	v1290 = (v1275 ^ v151) & int32(3)
	if base.Ui32(v1275) < base.Ui32(v151) {
		goto L432
	} else {
		goto L433
	}
L430:
	;
	if v1392 == int32(0) {
		goto L425
	} else {
		goto L466
	}
L431:
	;
	if base.Ui32(v1370) <= base.Ui32(int32(3)) {
		v1391 = v1369
		v1392 = v1370
		v1393 = v1371
		goto L430
	} else {
		goto L462
	}
L432:
	;
	if v1290 != 0 {
		goto L435
	} else {
		goto L436
	}
L433:
	;
	goto L434
L434:
	;
	if v1290 != 0 {
		v1352 = v1273
		goto L445
	} else {
		goto L446
	}
L435:
	;
	v1391 = v151
	v1392 = v1273
	v1393 = v1275
	goto L430
L436:
	;
	goto L437
L437:
	;
	if v1275&int32(3) == int32(0) {
		goto L438
	} else {
		goto L439
	}
L438:
	;
	v1369 = v151
	v1370 = v1273
	v1371 = v1275
	goto L431
L439:
	;
	goto L440
L440:
	;
	v1297 = v151
	v1298 = v1273
	v1299 = v1275
	goto L441
L441:
	;
	if v1298 == int32(0) {
		goto L425
	} else {
		goto L443
	}
L442:
	;
	v1369 = v1306
	v1370 = v1308
	v1371 = v1310
	goto L431
L443:
	;
	v1303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1297))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1299))) = uint8(v1303)
	v1305 = int32(1)
	v1306 = v1297 + v1305
	v1308 = v1298 - v1305
	v1310 = v1299 + v1305
	if v1310&int32(3) != 0 {
		v1297 = v1306
		v1298 = v1308
		v1299 = v1310
		goto L441
	} else {
		goto L444
	}
L444:
	;
	goto L442
L445:
	;
	if v1352 == int32(0) {
		goto L425
	} else {
		goto L458
	}
L446:
	;
	if v1280&int32(3) != 0 {
		goto L447
	} else {
		goto L448
	}
L447:
	;
	v1317 = v1273
	goto L450
L448:
	;
	v1332 = v1273
	goto L449
L449:
	;
	if base.Ui32(v1332) <= base.Ui32(int32(3)) {
		v1352 = v1332
		goto L445
	} else {
		goto L454
	}
L450:
	;
	if v1317 == int32(0) {
		goto L425
	} else {
		goto L452
	}
L451:
	;
	v1332 = v1323
	goto L449
L452:
	;
	v1323 = v1317 - int32(1)
	v1324 = v1275 + v1323
	v1326 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151+v1323))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1324))) = uint8(v1326)
	if v1324&int32(3) != 0 {
		v1317 = v1323
		goto L450
	} else {
		goto L453
	}
L453:
	;
	goto L451
L454:
	;
	v1339 = v1332
	goto L455
L455:
	;
	v1343 = v1339 - int32(4)
	v1346 = *(*int32)(unsafe.Add(mBase, uint32(v151+v1343)))
	*(*int32)(unsafe.Add(mBase, uint32(v1275+v1343))) = v1346
	if base.Ui32(int32(3)) < base.Ui32(v1343) {
		v1339 = v1343
		goto L455
	} else {
		goto L457
	}
L456:
	;
	v1352 = v1343
	goto L445
L457:
	;
	goto L456
L458:
	;
	v1359 = v1352
	goto L459
L459:
	;
	v1363 = v1359 - int32(1)
	v1366 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151+v1363))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1275+v1363))) = uint8(v1366)
	if v1363 != 0 {
		v1359 = v1363
		goto L459
	} else {
		goto L461
	}
L460:
	;
	goto L425
L461:
	;
	goto L460
L462:
	;
	v1376 = v1369
	v1377 = v1370
	v1378 = v1371
	goto L463
L463:
	;
	v1380 = *(*int32)(unsafe.Add(mBase, uint32(v1376)))
	*(*int32)(unsafe.Add(mBase, uint32(v1378))) = v1380
	v1382 = int32(4)
	v1383 = v1376 + v1382
	v1385 = v1378 + v1382
	v1387 = v1377 - v1382
	if base.Ui32(int32(3)) < base.Ui32(v1387) {
		v1376 = v1383
		v1377 = v1387
		v1378 = v1385
		goto L463
	} else {
		goto L465
	}
L464:
	;
	v1391 = v1383
	v1392 = v1387
	v1393 = v1385
	goto L430
L465:
	;
	goto L464
L466:
	;
	v1398 = v1391
	v1399 = v1392
	v1400 = v1393
	goto L467
L467:
	;
	v1402 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1398))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1400))) = uint8(v1402)
	v1404 = int32(1)
	v1409 = v1399 - v1404
	if v1409 != 0 {
		v1398 = v1398 + v1404
		v1399 = v1409
		v1400 = v1400 + v1404
		goto L467
	} else {
		goto L469
	}
L468:
	;
	goto L425
L469:
	;
	goto L468
L470:
	;
	if v192 != 0 {
		goto L487
	} else {
		goto L488
	}
L471:
	;
	v1484 = v1476 - v151
	goto L470
L472:
	;
	v1455 = v1451
	goto L481
L473:
	;
	v1435 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151))))
	if v1435 == int32(0) {
		goto L474
	} else {
		goto L475
	}
L474:
	;
	v1484 = int32(0)
	goto L470
L475:
	;
	goto L476
L476:
	;
	v1440 = v151
	goto L477
L477:
	;
	v1444 = v1440 + int32(1)
	if v1444&int32(3) == int32(0) {
		v1451 = v1444
		goto L472
	} else {
		goto L479
	}
L478:
	;
	v1476 = v1444
	goto L471
L479:
	;
	v1449 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1444))))
	if v1449 != 0 {
		v1440 = v1444
		goto L477
	} else {
		goto L480
	}
L480:
	;
	goto L478
L481:
	;
	v1461 = *(*int32)(unsafe.Add(mBase, uint32(v1455)))
	v1464 = int32(-2139062144)
	if (int32(16843008)-v1461|v1461)&v1464 == v1464 {
		v1455 = v1455 + int32(4)
		goto L481
	} else {
		goto L483
	}
L482:
	;
	v1470 = v1455
	goto L484
L483:
	;
	goto L482
L484:
	;
	v1474 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1470))))
	if v1474 != 0 {
		v1470 = v1470 + int32(1)
		goto L484
	} else {
		goto L486
	}
L485:
	;
	v1476 = v1470
	goto L471
L486:
	;
	goto L485
L487:
	;
	if v151 != v1427 {
		goto L490
	} else {
		goto L491
	}
L488:
	;
	goto L489
L489:
	;
	if v151 != v1427 {
		goto L539
	} else {
		goto L540
	}
L490:
	;
	if v1427 == v151 {
		goto L494
	} else {
		goto L495
	}
L491:
	;
	goto L492
L492:
	;
	v1789 = v1484 + v1427
	v1794 = v152
	v1795 = int32(4)
	goto L54
L493:
	;
	goto L492
L494:
	;
	goto L493
L495:
	;
	v1489 = v1427 + v1484
	if base.Ui32(v151-v1489) <= base.Ui32(int32(0)-v1484<<(uint(int32(1))%32)) {
		goto L496
	} else {
		goto L497
	}
L496:
	;
	v1496 = F___memcpy(m, v1427, v151, v1484)
	mBase = m.M
	goto L493
L497:
	;
	goto L498
L498:
	;
	v1499 = (v1427 ^ v151) & int32(3)
	if base.Ui32(v1427) < base.Ui32(v151) {
		goto L501
	} else {
		goto L502
	}
L499:
	;
	if v1601 == int32(0) {
		goto L494
	} else {
		goto L535
	}
L500:
	;
	if base.Ui32(v1579) <= base.Ui32(int32(3)) {
		v1600 = v1578
		v1601 = v1579
		v1602 = v1580
		goto L499
	} else {
		goto L531
	}
L501:
	;
	if v1499 != 0 {
		goto L504
	} else {
		goto L505
	}
L502:
	;
	goto L503
L503:
	;
	if v1499 != 0 {
		v1561 = v1484
		goto L514
	} else {
		goto L515
	}
L504:
	;
	v1600 = v151
	v1601 = v1484
	v1602 = v1427
	goto L499
L505:
	;
	goto L506
L506:
	;
	if v1427&int32(3) == int32(0) {
		goto L507
	} else {
		goto L508
	}
L507:
	;
	v1578 = v151
	v1579 = v1484
	v1580 = v1427
	goto L500
L508:
	;
	goto L509
L509:
	;
	v1506 = v151
	v1507 = v1484
	v1508 = v1427
	goto L510
L510:
	;
	if v1507 == int32(0) {
		goto L494
	} else {
		goto L512
	}
L511:
	;
	v1578 = v1515
	v1579 = v1517
	v1580 = v1519
	goto L500
L512:
	;
	v1512 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1506))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1508))) = uint8(v1512)
	v1514 = int32(1)
	v1515 = v1506 + v1514
	v1517 = v1507 - v1514
	v1519 = v1508 + v1514
	if v1519&int32(3) != 0 {
		v1506 = v1515
		v1507 = v1517
		v1508 = v1519
		goto L510
	} else {
		goto L513
	}
L513:
	;
	goto L511
L514:
	;
	if v1561 == int32(0) {
		goto L494
	} else {
		goto L527
	}
L515:
	;
	if v1489&int32(3) != 0 {
		goto L516
	} else {
		goto L517
	}
L516:
	;
	v1526 = v1484
	goto L519
L517:
	;
	v1541 = v1484
	goto L518
L518:
	;
	if base.Ui32(v1541) <= base.Ui32(int32(3)) {
		v1561 = v1541
		goto L514
	} else {
		goto L523
	}
L519:
	;
	if v1526 == int32(0) {
		goto L494
	} else {
		goto L521
	}
L520:
	;
	v1541 = v1532
	goto L518
L521:
	;
	v1532 = v1526 - int32(1)
	v1533 = v1427 + v1532
	v1535 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151+v1532))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1533))) = uint8(v1535)
	if v1533&int32(3) != 0 {
		v1526 = v1532
		goto L519
	} else {
		goto L522
	}
L522:
	;
	goto L520
L523:
	;
	v1548 = v1541
	goto L524
L524:
	;
	v1552 = v1548 - int32(4)
	v1555 = *(*int32)(unsafe.Add(mBase, uint32(v151+v1552)))
	*(*int32)(unsafe.Add(mBase, uint32(v1427+v1552))) = v1555
	if base.Ui32(int32(3)) < base.Ui32(v1552) {
		v1548 = v1552
		goto L524
	} else {
		goto L526
	}
L525:
	;
	v1561 = v1552
	goto L514
L526:
	;
	goto L525
L527:
	;
	v1568 = v1561
	goto L528
L528:
	;
	v1572 = v1568 - int32(1)
	v1575 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151+v1572))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1427+v1572))) = uint8(v1575)
	if v1572 != 0 {
		v1568 = v1572
		goto L528
	} else {
		goto L530
	}
L529:
	;
	goto L494
L530:
	;
	goto L529
L531:
	;
	v1585 = v1578
	v1586 = v1579
	v1587 = v1580
	goto L532
L532:
	;
	v1589 = *(*int32)(unsafe.Add(mBase, uint32(v1585)))
	*(*int32)(unsafe.Add(mBase, uint32(v1587))) = v1589
	v1591 = int32(4)
	v1592 = v1585 + v1591
	v1594 = v1587 + v1591
	v1596 = v1586 - v1591
	if base.Ui32(int32(3)) < base.Ui32(v1596) {
		v1585 = v1592
		v1586 = v1596
		v1587 = v1594
		goto L532
	} else {
		goto L534
	}
L533:
	;
	v1600 = v1592
	v1601 = v1596
	v1602 = v1594
	goto L499
L534:
	;
	goto L533
L535:
	;
	v1607 = v1600
	v1608 = v1601
	v1609 = v1602
	goto L536
L536:
	;
	v1611 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1607))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1609))) = uint8(v1611)
	v1613 = int32(1)
	v1618 = v1608 - v1613
	if v1618 != 0 {
		v1607 = v1607 + v1613
		v1608 = v1618
		v1609 = v1609 + v1613
		goto L536
	} else {
		goto L538
	}
L537:
	;
	goto L494
L538:
	;
	goto L537
L539:
	;
	if v1427 == v151 {
		goto L543
	} else {
		goto L544
	}
L540:
	;
	goto L541
L541:
	;
	v1780 = v1484 + v1427
	v1785 = int32(1)
	goto L59
L542:
	;
	goto L541
L543:
	;
	goto L542
L544:
	;
	v1636 = v1427 + v1484
	if base.Ui32(v151-v1636) <= base.Ui32(int32(0)-v1484<<(uint(int32(1))%32)) {
		goto L545
	} else {
		goto L546
	}
L545:
	;
	v1643 = F___memcpy(m, v1427, v151, v1484)
	mBase = m.M
	goto L542
L546:
	;
	goto L547
L547:
	;
	v1646 = (v1427 ^ v151) & int32(3)
	if base.Ui32(v1427) < base.Ui32(v151) {
		goto L550
	} else {
		goto L551
	}
L548:
	;
	if v1748 == int32(0) {
		goto L543
	} else {
		goto L584
	}
L549:
	;
	if base.Ui32(v1726) <= base.Ui32(int32(3)) {
		v1747 = v1725
		v1748 = v1726
		v1749 = v1727
		goto L548
	} else {
		goto L580
	}
L550:
	;
	if v1646 != 0 {
		goto L553
	} else {
		goto L554
	}
L551:
	;
	goto L552
L552:
	;
	if v1646 != 0 {
		v1708 = v1484
		goto L563
	} else {
		goto L564
	}
L553:
	;
	v1747 = v151
	v1748 = v1484
	v1749 = v1427
	goto L548
L554:
	;
	goto L555
L555:
	;
	if v1427&int32(3) == int32(0) {
		goto L556
	} else {
		goto L557
	}
L556:
	;
	v1725 = v151
	v1726 = v1484
	v1727 = v1427
	goto L549
L557:
	;
	goto L558
L558:
	;
	v1653 = v151
	v1654 = v1484
	v1655 = v1427
	goto L559
L559:
	;
	if v1654 == int32(0) {
		goto L543
	} else {
		goto L561
	}
L560:
	;
	v1725 = v1662
	v1726 = v1664
	v1727 = v1666
	goto L549
L561:
	;
	v1659 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1653))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1655))) = uint8(v1659)
	v1661 = int32(1)
	v1662 = v1653 + v1661
	v1664 = v1654 - v1661
	v1666 = v1655 + v1661
	if v1666&int32(3) != 0 {
		v1653 = v1662
		v1654 = v1664
		v1655 = v1666
		goto L559
	} else {
		goto L562
	}
L562:
	;
	goto L560
L563:
	;
	if v1708 == int32(0) {
		goto L543
	} else {
		goto L576
	}
L564:
	;
	if v1636&int32(3) != 0 {
		goto L565
	} else {
		goto L566
	}
L565:
	;
	v1673 = v1484
	goto L568
L566:
	;
	v1688 = v1484
	goto L567
L567:
	;
	if base.Ui32(v1688) <= base.Ui32(int32(3)) {
		v1708 = v1688
		goto L563
	} else {
		goto L572
	}
L568:
	;
	if v1673 == int32(0) {
		goto L543
	} else {
		goto L570
	}
L569:
	;
	v1688 = v1679
	goto L567
L570:
	;
	v1679 = v1673 - int32(1)
	v1680 = v1427 + v1679
	v1682 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151+v1679))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1680))) = uint8(v1682)
	if v1680&int32(3) != 0 {
		v1673 = v1679
		goto L568
	} else {
		goto L571
	}
L571:
	;
	goto L569
L572:
	;
	v1695 = v1688
	goto L573
L573:
	;
	v1699 = v1695 - int32(4)
	v1702 = *(*int32)(unsafe.Add(mBase, uint32(v151+v1699)))
	*(*int32)(unsafe.Add(mBase, uint32(v1427+v1699))) = v1702
	if base.Ui32(int32(3)) < base.Ui32(v1699) {
		v1695 = v1699
		goto L573
	} else {
		goto L575
	}
L574:
	;
	v1708 = v1699
	goto L563
L575:
	;
	goto L574
L576:
	;
	v1715 = v1708
	goto L577
L577:
	;
	v1719 = v1715 - int32(1)
	v1722 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151+v1719))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1427+v1719))) = uint8(v1722)
	if v1719 != 0 {
		v1715 = v1719
		goto L577
	} else {
		goto L579
	}
L578:
	;
	goto L543
L579:
	;
	goto L578
L580:
	;
	v1732 = v1725
	v1733 = v1726
	v1734 = v1727
	goto L581
L581:
	;
	v1736 = *(*int32)(unsafe.Add(mBase, uint32(v1732)))
	*(*int32)(unsafe.Add(mBase, uint32(v1734))) = v1736
	v1738 = int32(4)
	v1739 = v1732 + v1738
	v1741 = v1734 + v1738
	v1743 = v1733 - v1738
	if base.Ui32(int32(3)) < base.Ui32(v1743) {
		v1732 = v1739
		v1733 = v1743
		v1734 = v1741
		goto L581
	} else {
		goto L583
	}
L582:
	;
	v1747 = v1739
	v1748 = v1743
	v1749 = v1741
	goto L548
L583:
	;
	goto L582
L584:
	;
	v1754 = v1747
	v1755 = v1748
	v1756 = v1749
	goto L585
L585:
	;
	v1758 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1754))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1756))) = uint8(v1758)
	v1760 = int32(1)
	v1765 = v1755 - v1760
	if v1765 != 0 {
		v1754 = v1754 + v1760
		v1755 = v1765
		v1756 = v1756 + v1760
		goto L585
	} else {
		goto L587
	}
L586:
	;
	goto L543
L587:
	;
	goto L586
L588:
	;
	goto L46
L589:
	;
	v1806 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v1798))) = uint8(v1806)
	v1810 = v1798 + int32(1)
	goto L591
L590:
	;
	v1810 = v1798
	goto L591
L591:
	;
	v1811 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1810))) = uint8(v1811)
	goto L37
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
