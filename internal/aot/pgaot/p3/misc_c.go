package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_CLOGPagePrecedes(m *base.Module, l0 int64, l1 int64) int32 {
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	v3 = int32(0)
	v7 = int32(15)
	v9 = int32(4)
	v10 = base.I32_wrap_i64(l0)<<(uint(v7)%32) | v9
	v13 = base.I32_wrap_i64(l1) << (uint(v7) % 32)
	v15 = v13 | v9
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v15))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v10)) == v3 {
		v27 = base.B2i32(base.Ui32(v10) < base.Ui32(v15))
	} else {
		v27 = int32(base.Ui32(v10-v15) >> (uint(int32(31)) % 32))
	}
	if v27 != 0 {
		v29 = v13 + int32(32771)
		if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v29))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v10)) == int32(0) {
			v41 = base.B2i32(base.Ui32(v10) < base.Ui32(v29))
		} else {
			v41 = int32(base.Ui32(v10-v29) >> (uint(int32(31)) % 32))
		}
		v42 = v41
	} else {
		v42 = v3
	}
	return v42
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
																			v109 = *(*int32)(unsafe.Add(mBase, _consts[433]))
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
																		v109 = *(*int32)(unsafe.Add(mBase, _consts[433]))
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
																		v109 = *(*int32)(unsafe.Add(mBase, _consts[433]))
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
																	v109 = *(*int32)(unsafe.Add(mBase, _consts[433]))
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
																		v109 = *(*int32)(unsafe.Add(mBase, _consts[433]))
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
																	v109 = *(*int32)(unsafe.Add(mBase, _consts[433]))
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
																	v109 = *(*int32)(unsafe.Add(mBase, _consts[433]))
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
																v109 = *(*int32)(unsafe.Add(mBase, _consts[433]))
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
					F_errcode(m, int32(290948))
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
								F_errmsg(m, int32(108053), v15)
								mBase = m.M
								v138 = m.ExcPending
								if v138 != 0 {
									return
								} else {
									F_errfinish(m, int32(469765), int32(77), int32(338366))
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
func F_CheckBuiltinCryptoMode(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	v3 = *(*int32)(unsafe.Add(mBase, _consts[1297]))
	if v3 == int32(1) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			F_errmsg(m, int32(433626), int32(0))
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return
			} else {
				F_errfinish(m, int32(467498), int32(735), int32(394463))
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
	} else {
		return
	}
}
func F_CheckSASLAuth(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
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
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v194 int32
	_ = v194
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	v5 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(80)
	m.G0 = v12
	*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = v5
	*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v5
	F_initStringInfo(m, v12-int32(-64))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	m.T0[v26].(func(*base.Module, int32, int32))(m, l1, v12-int32(-64))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	F_appendStringInfoChar(m, v12-int32(-64), int32(0))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v12)+64))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v12)+68))
	F_sendAuthRequest(m, int32(10), v35, v36)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v12)+64))
	F_pfree(m, v39)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v50 = int32(1)
	v52 = v5
	goto L10
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L1
	} else {
		goto L66
	}
L8:
	;
	m.G0 = v12 + int32(80)
	return v194
L9:
	;
	v194 = int32(-1)
	goto L8
L10:
	;
	F_pq_startmsgread(m)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L12
	}
L11:
	;
	if v145 == int32(1) {
		v194 = int32(0)
		goto L8
	} else {
		goto L65
	}
L12:
	;
	v55 = F_pq_getbyte(m)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	if v55 != int32(112) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	if v55 == int32(-1) {
		v194 = int32(-2)
		goto L8
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	F_initStringInfo(m, v12+int32(48))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L22
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v55
	F_errmsg(m, int32(454097), v12)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	F_errfinish(m, int32(473951), int32(90), int32(304764))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
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
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v84 = F_pq_getmessage(m, v12+int32(48), v83)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	if v84 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
	F_pfree(m, v86)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v91 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L28
	}
L27:
	;
	goto L9
L28:
	;
	if v91 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v12)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v93
	F_errmsg_internal(m, int32(452099), v12+int32(32))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	if v50&int32(1) != 0 {
		goto L36
	} else {
		goto L37
	}
L32:
	;
	F_errfinish(m, int32(473951), int32(105), int32(304764))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	F_pq_getmsgend(m, v12+int32(48))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L46
	}
L35:
	;
	v130 = F_pq_getmsgbytes(m, v12+int32(48), v129)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L45
	}
L36:
	;
	v112 = F_pq_getmsgrawstring(m, v12+int32(48))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v12)+52))
	v128 = v52
	v129 = v125
	goto L35
L39:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v115 = m.T0[v114].(func(*base.Module, int32, int32, int32) int32)(m, l1, v112, l2)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	v120 = F_pq_getmsgint(m, v12+int32(48), int32(4))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	if v120 != int32(-1) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v128 = v115
	v129 = v120
	goto L35
L43:
	;
	goto L44
L44:
	;
	v132 = int32(-1)
	v134 = v115
	v135 = int32(0)
	goto L34
L45:
	;
	v132 = v129
	v134 = v128
	v135 = v130
	goto L34
L46:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v145 = m.T0[v144].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v134, v135, v132, v12+int32(44), v12+int32(40), l3)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
	F_pfree(m, v147)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v12)+44))
	if v150 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	if v145 == int32(2) {
		goto L7
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	v181 = int32(0)
	if v145 == v181 {
		v50 = v181
		v52 = v134
		goto L10
	} else {
		goto L64
	}
L52:
	;
	v155 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	if v155 != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v12)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v157
	F_errmsg_internal(m, int32(452146), v12+int32(16))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	if v145 == int32(1) {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	F_errfinish(m, int32(473951), int32(176), int32(304764))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	goto L56
L59:
	;
	v173 = int32(12)
	goto L61
L60:
	;
	v173 = int32(11)
	goto L61
L61:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v12)+44))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v12)+40))
	F_sendAuthRequest(m, v173, v174, v175)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v12)+44))
	F_pfree(m, v178)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
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
	F_errmsg_internal(m, int32(346026), int32(0))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(473951), int32(171), int32(304764))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
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
						F_errmsg(m, int32(280460), int32(0))
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return
						} else {
							F_errfinish(m, int32(472859), int32(572), int32(277151))
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
						F_errmsg(m, int32(433435), v6+int32(16))
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
								F_errdetail_internal(m, int32(195784), v6)
								mBase = m.M
								v32 = m.ExcPending
								if v32 != 0 {
									return
								} else {
									F_errfinish(m, int32(472859), int32(579), int32(277151))
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
	var v18 int32
	_ = v18
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v9 == int32(0) {
		v55 = int32(0)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
		if v13 <= int32(0) {
			v55 = v9
		} else {
			if l2 == int32(0) {
				v55 = v9
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
				if int32(0) < v18 {
					v26 = int32(0)
					for {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
						v34 = *(*int32)(unsafe.Add(mBase, uint32(v30+v26<<(uint(int32(2))%32))))
						v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+28))
						if v35 != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(v34)+28)) = v13 + v35
						} else {
						}
						v39 = v26 + int32(1)
						v40 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
						if v39 < v40 {
							v26 = v39
							continue
						} else {
							break
						}
						break
					}
				} else {
				}
				v50 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v55 = v50
			}
		}
	}
	v59 = F_list_concat(m, v55, l3)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v59
		v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v63 = F_list_concat(m, v62, l2)
		mBase = m.M
		v64 = m.ExcPending
		if v64 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = v63
			return
		}
	}
}
func F_ConditionVariableTimedSleep(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v30 int64
	_ = v30
	var v33 int64
	_ = v33
	var v36 int32
	_ = v36
	var v37 int64
	_ = v37
	var v38 int32
	_ = v38
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
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
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int64
	_ = v113
	var v114 int64
	_ = v114
	var v121 float64
	_ = v121
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v137 int32
	_ = v137
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v18 = *(*int32)(unsafe.Add(mBase, _consts[846]))
	if l0 != v18 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v15 + int32(16)
	return v137
L2:
	;
	F_ConditionVariablePrepareToSleep(m, l0)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v25 = base.B2i32(l1 < int32(0))
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
	v137 = int32(0)
	goto L1
L7:
	;
	v47 = v38
	goto L11
L8:
	;
	v36 = int32(33)
	v37 = int64(0)
	v38 = int32(-1)
	goto L7
L9:
	;
	goto L10
L10:
	;
	F___clock_gettime(m, int32(1), v15)
	mBase = m.M
	v30 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
	v33 = int64(*(*int32)(unsafe.Add(mBase, uint32(v15)+8)))
	v36 = int32(41)
	v37 = v30*int64(-1000000000) - v33
	v38 = l1
	goto L7
L11:
	;
	v52 = *(*int32)(unsafe.Add(mBase, _consts[302]))
	v53 = F_WaitLatch(m, v52, v36, v47, l2)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L5
	} else {
		goto L13
	}
L12:
	;
	v137 = int32(0)
	goto L1
L13:
	;
	v56 = *(*int32)(unsafe.Add(mBase, _consts[302]))
	*(*int32)(unsafe.Add(mBase, uint32(v56))) = int32(0)
	goto L14
L14:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1)
	if v59 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	F_s_lock(m, l0, int32(475679), int32(183), int32(225585))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L5
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v67 = int32(0)
	v69 = *(*int32)(unsafe.Add(mBase, _consts[141]))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	v72 = *(*int32)(unsafe.Add(mBase, _consts[100]))
	v75 = v70 + v72*int32(640)
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+88))
	if v76 != 0 {
		v98 = v67
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
	v103 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	if v103 != 0 {
		goto L26
	} else {
		goto L27
	}
L20:
	;
	v78 = v75 + int32(84)
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	if v79 != 0 {
		v98 = v67
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v80 == int32(-1) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v72
	v98 = int32(1)
	goto L19
L23:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v78))) = int64(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v72
	goto L22
L24:
	;
	goto L25
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v78)+4)) = v80
	v88 = *(*int32)(unsafe.Add(mBase, _consts[141]))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	*(*int32)(unsafe.Add(mBase, uint32(v89+v80*int32(640))+84)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v78))) = int32(-1)
	goto L22
L26:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L5
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v107 = *(*int32)(unsafe.Add(mBase, _consts[846]))
	v109 = v98 | base.B2i32(l0 != v107)
	if l1 < int32(0) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	goto L28
L30:
	;
	if v109 == int32(0) {
		goto L11
	} else {
		goto L38
	}
L31:
	;
	if v109 != 0 {
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v110 = int32(1)
	F___clock_gettime(m, v110, v15)
	mBase = m.M
	v113 = int64(*(*int32)(unsafe.Add(mBase, uint32(v15)+8)))
	v114 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
	v121 = base.F64_div(base.F64_convert_i64_s(v113+(v114*int64(1000000000)+v37)), float64(1e+06))
	if base.F64_lt(base.F64_abs(v121), float64(2.147483648e+09)) != 0 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v128 = l1 - v127
	if int32(0) < v128 {
		v47 = v128
		goto L11
	} else {
		goto L37
	}
L34:
	;
	v125 = base.I32_trunc_f64_s(v121)
	v127 = v125
	goto L33
L35:
	;
	goto L36
L36:
	;
	v127 = int32(-2147483648)
	goto L33
L37:
	;
	v137 = v110
	goto L1
L38:
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
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
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int64
	_ = v82
	var v83 int64
	_ = v83
	var v84 int64
	_ = v84
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
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
				v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+76))
				if v15 < int32(0) {
					v18 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
					v20 = v18
				} else {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
					v20 = v19
				}
				if int32(base.Ui32(v20)>>(uint(int32(5))%32))&int32(1) == int32(0) {
					v82 = *(*int64)(unsafe.Add(mBase, uint32(l0)+176))
					v83 = int64(*(*int32)(unsafe.Add(mBase, uint32(v4)+4)))
					v84 = v82 + v83
					*(*int64)(unsafe.Add(mBase, uint32(l0)+176)) = v84
					v89 = *(*int32)(unsafe.Add(mBase, _consts[454]))
					if v89 == int32(0) {
					} else {
						v93 = int32(*(*uint8)(unsafe.Add(mBase, _consts[455])))
						if v93 != int32(1) {
						} else {
							v96 = int32(4437236)
							v98 = *(*int32)(unsafe.Add(mBase, _consts[14]))
							v99 = int32(1)
							*(*int32)(unsafe.Add(mBase, _consts[14])) = v98 + v99
							v102 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
							*(*int32)(unsafe.Add(mBase, uint32(v89))) = v102 + v99
							*(*int64)(unsafe.Add(mBase, uint32(v89+int32(0))+232)) = v84
							v110 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
							*(*int32)(unsafe.Add(mBase, uint32(v89))) = v110 + v99
							v116 = *(*int32)(unsafe.Add(mBase, _consts[14]))
							*(*int32)(unsafe.Add(mBase, _consts[14])) = v116 - v99
						}
					}
					v120 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
					v121 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v120))) = uint8(v121)
					*(*int32)(unsafe.Add(mBase, uint32(v4)+12)) = v121
					*(*int32)(unsafe.Add(mBase, uint32(v4)+4)) = v121
					return
				} else {
					v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)))
					if v27 == int32(1) {
						v31 = *(*int32)(unsafe.Add(mBase, _consts[40]))
						if v31 == int32(64) {
							F_ClosePipeToProgram(m, l0)
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, _consts[40])) = int32(64)
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return
								} else {
									F_errcode_for_file_access(m)
									mBase = m.M
									v44 = m.ExcPending
									if v44 != 0 {
										return
									} else {
										F_errmsg(m, int32(279291), int32(0))
										mBase = m.M
										v48 = m.ExcPending
										if v48 != 0 {
											return
										} else {
											F_errfinish(m, int32(472859), int32(477), int32(29661))
											mBase = m.M
											v53 = m.ExcPending
											if v53 != 0 {
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
							v42 = m.ExcPending
							if v42 != 0 {
								return
							} else {
								F_errcode_for_file_access(m)
								mBase = m.M
								v44 = m.ExcPending
								if v44 != 0 {
									return
								} else {
									F_errmsg(m, int32(279291), int32(0))
									mBase = m.M
									v48 = m.ExcPending
									if v48 != 0 {
										return
									} else {
										F_errfinish(m, int32(472859), int32(477), int32(29661))
										mBase = m.M
										v53 = m.ExcPending
										if v53 != 0 {
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
						v57 = m.ExcPending
						if v57 != 0 {
							return
						} else {
							F_errcode_for_file_access(m)
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
								return
							} else {
								F_errmsg(m, int32(280207), int32(0))
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return
								} else {
									F_errfinish(m, int32(472859), int32(482), int32(29661))
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
			} else {
				v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)))
				if v27 == int32(1) {
					v31 = *(*int32)(unsafe.Add(mBase, _consts[40]))
					if v31 == int32(64) {
						F_ClosePipeToProgram(m, l0)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, _consts[40])) = int32(64)
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return
							} else {
								F_errcode_for_file_access(m)
								mBase = m.M
								v44 = m.ExcPending
								if v44 != 0 {
									return
								} else {
									F_errmsg(m, int32(279291), int32(0))
									mBase = m.M
									v48 = m.ExcPending
									if v48 != 0 {
										return
									} else {
										F_errfinish(m, int32(472859), int32(477), int32(29661))
										mBase = m.M
										v53 = m.ExcPending
										if v53 != 0 {
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
						v42 = m.ExcPending
						if v42 != 0 {
							return
						} else {
							F_errcode_for_file_access(m)
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return
							} else {
								F_errmsg(m, int32(279291), int32(0))
								mBase = m.M
								v48 = m.ExcPending
								if v48 != 0 {
									return
								} else {
									F_errfinish(m, int32(472859), int32(477), int32(29661))
									mBase = m.M
									v53 = m.ExcPending
									if v53 != 0 {
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
					v57 = m.ExcPending
					if v57 != 0 {
						return
					} else {
						F_errcode_for_file_access(m)
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return
						} else {
							F_errmsg(m, int32(280207), int32(0))
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return
							} else {
								F_errfinish(m, int32(472859), int32(482), int32(29661))
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
	case 1:
		v70 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
		v71 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
		v73 = *(*int32)(unsafe.Add(mBase, _consts[456]))
		v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+16))
		v75 = m.T0[v74].(func(*base.Module, int32, int32, int32) int32)(m, int32(100), v70, v71)
		mBase = m.M
		v76 = m.ExcPending
		if v76 != 0 {
			return
		} else {
			v82 = *(*int64)(unsafe.Add(mBase, uint32(l0)+176))
			v83 = int64(*(*int32)(unsafe.Add(mBase, uint32(v4)+4)))
			v84 = v82 + v83
			*(*int64)(unsafe.Add(mBase, uint32(l0)+176)) = v84
			v89 = *(*int32)(unsafe.Add(mBase, _consts[454]))
			if v89 == int32(0) {
			} else {
				v93 = int32(*(*uint8)(unsafe.Add(mBase, _consts[455])))
				if v93 != int32(1) {
				} else {
					v96 = int32(4437236)
					v98 = *(*int32)(unsafe.Add(mBase, _consts[14]))
					v99 = int32(1)
					*(*int32)(unsafe.Add(mBase, _consts[14])) = v98 + v99
					v102 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
					*(*int32)(unsafe.Add(mBase, uint32(v89))) = v102 + v99
					*(*int64)(unsafe.Add(mBase, uint32(v89+int32(0))+232)) = v84
					v110 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
					*(*int32)(unsafe.Add(mBase, uint32(v89))) = v110 + v99
					v116 = *(*int32)(unsafe.Add(mBase, _consts[14]))
					*(*int32)(unsafe.Add(mBase, _consts[14])) = v116 - v99
				}
			}
			v120 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
			v121 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v120))) = uint8(v121)
			*(*int32)(unsafe.Add(mBase, uint32(v4)+12)) = v121
			*(*int32)(unsafe.Add(mBase, uint32(v4)+4)) = v121
			return
		}
	case 2:
		v77 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
		v78 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
		v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
		m.T0[v79].(func(*base.Module, int32, int32))(m, v77, v78)
		mBase = m.M
		v81 = m.ExcPending
		if v81 != 0 {
			return
		} else {
			v82 = *(*int64)(unsafe.Add(mBase, uint32(l0)+176))
			v83 = int64(*(*int32)(unsafe.Add(mBase, uint32(v4)+4)))
			v84 = v82 + v83
			*(*int64)(unsafe.Add(mBase, uint32(l0)+176)) = v84
			v89 = *(*int32)(unsafe.Add(mBase, _consts[454]))
			if v89 == int32(0) {
			} else {
				v93 = int32(*(*uint8)(unsafe.Add(mBase, _consts[455])))
				if v93 != int32(1) {
				} else {
					v96 = int32(4437236)
					v98 = *(*int32)(unsafe.Add(mBase, _consts[14]))
					v99 = int32(1)
					*(*int32)(unsafe.Add(mBase, _consts[14])) = v98 + v99
					v102 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
					*(*int32)(unsafe.Add(mBase, uint32(v89))) = v102 + v99
					*(*int64)(unsafe.Add(mBase, uint32(v89+int32(0))+232)) = v84
					v110 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
					*(*int32)(unsafe.Add(mBase, uint32(v89))) = v110 + v99
					v116 = *(*int32)(unsafe.Add(mBase, _consts[14]))
					*(*int32)(unsafe.Add(mBase, _consts[14])) = v116 - v99
				}
			}
			v120 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
			v121 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v120))) = uint8(v121)
			*(*int32)(unsafe.Add(mBase, uint32(v4)+12)) = v121
			*(*int32)(unsafe.Add(mBase, uint32(v4)+4)) = v121
			return
		}
	default:
		v82 = *(*int64)(unsafe.Add(mBase, uint32(l0)+176))
		v83 = int64(*(*int32)(unsafe.Add(mBase, uint32(v4)+4)))
		v84 = v82 + v83
		*(*int64)(unsafe.Add(mBase, uint32(l0)+176)) = v84
		v89 = *(*int32)(unsafe.Add(mBase, _consts[454]))
		if v89 == int32(0) {
		} else {
			v93 = int32(*(*uint8)(unsafe.Add(mBase, _consts[455])))
			if v93 != int32(1) {
			} else {
				v96 = int32(4437236)
				v98 = *(*int32)(unsafe.Add(mBase, _consts[14]))
				v99 = int32(1)
				*(*int32)(unsafe.Add(mBase, _consts[14])) = v98 + v99
				v102 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
				*(*int32)(unsafe.Add(mBase, uint32(v89))) = v102 + v99
				*(*int64)(unsafe.Add(mBase, uint32(v89+int32(0))+232)) = v84
				v110 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
				*(*int32)(unsafe.Add(mBase, uint32(v89))) = v110 + v99
				v116 = *(*int32)(unsafe.Add(mBase, _consts[14]))
				*(*int32)(unsafe.Add(mBase, _consts[14])) = v116 - v99
			}
		}
		v120 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
		v121 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v120))) = uint8(v121)
		*(*int32)(unsafe.Add(mBase, uint32(v4)+12)) = v121
		*(*int32)(unsafe.Add(mBase, uint32(v4)+4)) = v121
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
		v13 = *(*int32)(unsafe.Add(mBase, _consts[0]))
		*(*int32)(unsafe.Add(mBase, uint32(v4)+16)) = v13
		v19 = F_AllocSetContextCreateInternal(m, v13, int32(59432), int32(0), int32(8192), int32(8388608))
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
	v26 = *(*int32)(unsafe.Add(mBase, _consts[8]))
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
	v40 = F_pg_snprintf(m, v8+int32(144), int32(1024), int32(195784), v8+int32(16))
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
	v52 = F_pg_snprintf(m, v8+int32(144), int32(1024), int32(36834), v8+int32(32))
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
	v64 = *(*int32)(unsafe.Add(mBase, _consts[40]))
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
	F_errmsg(m, int32(282964), v8)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(475011), int32(355), int32(323881))
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
	v10 = *(*int32)(unsafe.Add(mBase, _consts[804]))
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
	v13 = *(*int32)(unsafe.Add(mBase, _consts[3]))
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
	v21 = *(*int32)(unsafe.Add(mBase, _consts[3]))
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
	*(*int32)(unsafe.Add(mBase, uint32(v7)+40)) = int32(529025)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+36)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = int32(466108)
	v60 = F_pg_snprintf(m, v7+int32(2192), int32(1024), int32(166798), v7+int32(32))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L3
	} else {
		goto L16
	}
L12:
	;
	v47 = F_pg_snprintf(m, v7+int32(2192), int32(1024), int32(298826), int32(0))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L3
	} else {
		goto L15
	}
L13:
	;
	v40 = F_pg_snprintf(m, v7+int32(2192), int32(1024), int32(345140), int32(0))
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
	v78 = *(*int32)(unsafe.Add(mBase, _consts[8]))
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
	v105 = F_pg_snprintf(m, v7+int32(144), int32(2048), int32(166924), v7+int32(16))
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
	v117 = *(*int32)(unsafe.Add(mBase, _consts[40]))
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
	if v137&int32(61440) == int32(16384) {
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
	F_errmsg(m, int32(282964), v7)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L3
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(475011), int32(266), int32(324454))
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
	*(*int32)(unsafe.Add(mBase, _consts[49])) = v3
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
		v29 = int32(1625344)
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
					F_errfinish(m, int32(476670), int32(2031), int32(330558))
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
		v29 = int32(1625352)
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
					F_errfinish(m, int32(476670), int32(2031), int32(330558))
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
		v29 = int32(1625360)
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
					F_errfinish(m, int32(476670), int32(2031), int32(330558))
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
		v29 = int32(1625368)
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
					F_errfinish(m, int32(476670), int32(2031), int32(330558))
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
			F_errmsg_internal(m, int32(463494), v7)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return
			} else {
				F_errfinish(m, int32(476670), int32(2034), int32(330558))
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
		v29 = int32(1625376)
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
					F_errfinish(m, int32(476670), int32(2031), int32(330558))
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
		v29 = int32(1625384)
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
					F_errfinish(m, int32(476670), int32(2031), int32(330558))
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
		v29 = int32(1625392)
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
					F_errfinish(m, int32(476670), int32(2031), int32(330558))
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
func F_cfunc_match(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
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
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v168 int32
	_ = v168
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v217 int32
	_ = v217
	v5 = int32(24)
	goto L6
L1:
	;
	return int32(0)
L2:
	;
	return v217
L3:
	;
	if v67 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L4:
	;
	v67 = int32(0)
	goto L3
L5:
	;
	v41 = v36
	v42 = v37
	v43 = v38
	goto L15
L6:
	;
	if (l0|l1)&int32(3) != 0 {
		v36 = l0
		v37 = l1
		v38 = v5
		goto L5
	} else {
		goto L9
	}
L8:
	;
	if v26 == int32(0) {
		goto L4
	} else {
		goto L14
	}
L9:
	;
	v13 = l0
	v14 = l1
	v15 = v5
	goto L10
L10:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	if v18 != v19 {
		v36 = v13
		v37 = v14
		v38 = v15
		goto L5
	} else {
		goto L12
	}
L11:
	;
	goto L8
L12:
	;
	v21 = int32(4)
	v22 = v14 + v21
	v24 = v13 + v21
	v26 = v15 - v21
	if base.Ui32(int32(3)) < base.Ui32(v26) {
		v13 = v24
		v14 = v22
		v15 = v26
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	v36 = v24
	v37 = v22
	v38 = v26
	goto L5
L15:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
	if v46 == v47 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v67 = v46 - v47
	goto L3
L17:
	;
	v49 = int32(1)
	v54 = v43 - v49
	if v54 != 0 {
		v41 = v41 + v49
		v42 = v42 + v49
		v43 = v54
		goto L15
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	goto L16
L20:
	;
	goto L4
L21:
	;
	v70 = int32(1)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if int32(0) < v71 {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L23
L23:
	;
	v217 = int32(1)
	goto L2
L24:
	;
	v74 = int32(28)
	v75 = l0 + v74
	v77 = l1 + v74
	v79 = v71 << (uint(int32(2)) % 32)
	if base.Ui32(int32(4)) <= base.Ui32(v79) {
		goto L30
	} else {
		goto L31
	}
L25:
	;
	goto L26
L26:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v143 != 0 {
		goto L46
	} else {
		goto L47
	}
L27:
	;
	if v141 != 0 {
		v217 = v70
		goto L2
	} else {
		goto L45
	}
L28:
	;
	v141 = int32(0)
	goto L27
L29:
	;
	v115 = v110
	v116 = v111
	v117 = v112
	goto L39
L30:
	;
	if (v75|v77)&int32(3) != 0 {
		v110 = v75
		v111 = v77
		v112 = v79
		goto L29
	} else {
		goto L33
	}
L31:
	;
	v103 = v75
	v104 = v77
	v105 = v79
	goto L32
L32:
	;
	if v105 == int32(0) {
		goto L28
	} else {
		goto L38
	}
L33:
	;
	v87 = v75
	v88 = v77
	v89 = v79
	goto L34
L34:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	if v92 != v93 {
		v110 = v87
		v111 = v88
		v112 = v89
		goto L29
	} else {
		goto L36
	}
L35:
	;
	v103 = v98
	v104 = v96
	v105 = v100
	goto L32
L36:
	;
	v95 = int32(4)
	v96 = v88 + v95
	v98 = v87 + v95
	v100 = v89 - v95
	if base.Ui32(int32(3)) < base.Ui32(v100) {
		v87 = v98
		v88 = v96
		v89 = v100
		goto L34
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	v110 = v103
	v111 = v104
	v112 = v105
	goto L29
L39:
	;
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115))))
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116))))
	if v120 == v121 {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v141 = v120 - v121
	goto L27
L41:
	;
	v123 = int32(1)
	v128 = v117 - v123
	if v128 != 0 {
		v115 = v115 + v123
		v116 = v116 + v123
		v117 = v128
		goto L39
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	goto L40
L44:
	;
	goto L28
L45:
	;
	goto L26
L46:
	;
	if v142 == int32(0) {
		v217 = v70
		goto L2
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	if v142 == int32(0) {
		goto L1
	} else {
		goto L65
	}
L49:
	;
	v146 = int32(0)
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	if v150 != v151 {
		v202 = v146
		goto L51
	} else {
		goto L52
	}
L50:
	;
	if v202 == int32(0) {
		v217 = v70
		goto L2
	} else {
		goto L64
	}
L51:
	;
	goto L50
L52:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v143)+4))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v142)+4))
	if v153 != v154 {
		v202 = v146
		goto L51
	} else {
		goto L53
	}
L53:
	;
	if v150 <= int32(0) {
		v202 = int32(1)
		goto L51
	} else {
		goto L54
	}
L54:
	;
	v160 = v150 << (uint(int32(4)) % 32)
	v162 = int32(20)
	v168 = int32(0)
	goto L55
L55:
	;
	v175 = v168 * int32(100)
	v176 = v143 + v160 + v162 + v175
	v177 = int32(4)
	v179 = v175 + (v142 + v160 + v162)
	v182 = F_strcmp(m, v176+v177, v179+v177)
	mBase = m.M
	if v182 != 0 {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	v202 = int32(0)
	goto L51
L57:
	;
	goto L56
L58:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v176)+68))
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v179)+68))
	if v183 != v184 {
		goto L57
	} else {
		goto L59
	}
L59:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v176)+76))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v179)+76))
	if v186 != v187 {
		goto L57
	} else {
		goto L60
	}
L60:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v176)+96))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v179)+96))
	if v189 != v190 {
		goto L57
	} else {
		goto L61
	}
L61:
	;
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v176)+91)))
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179)+91)))
	if v192 != v193 {
		goto L57
	} else {
		goto L62
	}
L62:
	;
	v195 = int32(1)
	v197 = v168 + v195
	if v150 != v197 {
		v168 = v197
		goto L55
	} else {
		goto L63
	}
L63:
	;
	v202 = v195
	goto L51
L64:
	;
	goto L1
L65:
	;
	goto L23
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
	var v47 int32
	_ = v47
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
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
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v212 int64
	_ = v212
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
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
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v278 int32
	_ = v278
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v349 int32
	_ = v349
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
	var v367 int32
	_ = v367
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
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v408 int32
	_ = v408
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v444 int32
	_ = v444
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
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
	var v479 int32
	_ = v479
	var v496 int32
	_ = v496
	var v504 int32
	_ = v504
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v541 int32
	_ = v541
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v561 int32
	_ = v561
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v567 int32
	_ = v567
	var v580 int32
	_ = v580
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
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
		goto L83
	} else {
		goto L232
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
	v208 = m.G0
	v210 = v208 - int32(32)
	m.G0 = v210
	v212 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v210)+24)) = v212
	*(*int64)(unsafe.Add(mBase, uint32(v210)+8)) = v212
	*(*int64)(unsafe.Add(mBase, uint32(v210)+16)) = v212
	*(*int64)(unsafe.Add(mBase, uint32(v210))) = v212
	if l0 <= int32(0) {
		v639 = int32(1)
		goto L52
	} else {
		goto L53
	}
L6:
	;
	v45 = int32(1)
	v47 = int32(0)
	if l0 != v45 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v56 = v47
	v59 = v7
	goto L10
L8:
	;
	v148 = v47
	goto L9
L9:
	;
	if l0&v45 == int32(0) {
		goto L2
	} else {
		goto L39
	}
L10:
	;
	v90 = int32(23)
	v93 = l1 + v56<<(uint(int32(2))%32)
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
	if v94 <= int32(3830) {
		goto L16
	} else {
		goto L17
	}
L11:
	;
	v148 = v140
	goto L9
L12:
	;
	v115 = int32(23)
	v117 = v93 + int32(4)
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
	if v118 <= int32(3830) {
		goto L28
	} else {
		goto L29
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v93))) = v112
	goto L12
L14:
	;
	v112 = int32(3904)
	goto L13
L15:
	;
	v112 = int32(1007)
	goto L13
L16:
	;
	switch v94 - int32(2277) {
	case 0:
		goto L15
	case 1, 2, 3, 4, 5:
		goto L12
	case 6:
		v112 = v90
		goto L13
	default:
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	switch v94 - int32(5077) {
	case 0, 2:
		v112 = v90
		goto L13
	case 1:
		goto L15
	case 3:
		goto L14
	default:
		goto L22
	}
L19:
	;
	if v94 == int32(2776) {
		v112 = v90
		goto L13
	} else {
		goto L20
	}
L20:
	;
	if v94 == int32(3500) {
		v112 = v90
		goto L13
	} else {
		goto L21
	}
L21:
	;
	goto L12
L22:
	;
	if v94 == int32(3831) {
		goto L14
	} else {
		goto L23
	}
L23:
	;
	if v94 != int32(4537) {
		goto L12
	} else {
		goto L24
	}
L24:
	;
	v112 = int32(4451)
	goto L13
L25:
	;
	v139 = int32(2)
	v140 = v56 + v139
	v142 = v59 + v139
	if v142 != l0&int32(2147483646) {
		v56 = v140
		v59 = v142
		goto L10
	} else {
		goto L38
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v117))) = v136
	goto L25
L27:
	;
	v136 = int32(1007)
	goto L26
L28:
	;
	switch v118 - int32(2277) {
	case 0:
		goto L27
	case 1, 2, 3, 4, 5:
		goto L25
	case 6:
		v136 = v115
		goto L26
	default:
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	switch v118 - int32(5077) {
	case 0, 2:
		v136 = v115
		goto L26
	case 1:
		goto L27
	case 3:
		goto L34
	default:
		goto L35
	}
L31:
	;
	if v118 == int32(2776) {
		v136 = v115
		goto L26
	} else {
		goto L32
	}
L32:
	;
	if v118 == int32(3500) {
		v136 = v115
		goto L26
	} else {
		goto L33
	}
L33:
	;
	goto L25
L34:
	;
	v136 = int32(3904)
	goto L26
L35:
	;
	if v118 == int32(3831) {
		goto L34
	} else {
		goto L36
	}
L36:
	;
	if v118 != int32(4537) {
		goto L25
	} else {
		goto L37
	}
L37:
	;
	v136 = int32(4451)
	goto L26
L38:
	;
	goto L11
L39:
	;
	v184 = int32(23)
	v187 = l1 + v148<<(uint(int32(2))%32)
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v187)))
	if v188 <= int32(3830) {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v187))) = v206
	goto L2
L41:
	;
	v206 = int32(1007)
	goto L40
L42:
	;
	switch v188 - int32(2277) {
	case 0:
		goto L41
	case 1, 2, 3, 4, 5:
		goto L2
	case 6:
		v206 = v184
		goto L40
	default:
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	switch v188 - int32(5077) {
	case 0, 2:
		v206 = v184
		goto L40
	case 1:
		goto L41
	case 3:
		goto L48
	default:
		goto L49
	}
L45:
	;
	if v188 == int32(2776) {
		v206 = v184
		goto L40
	} else {
		goto L46
	}
L46:
	;
	if v188 == int32(3500) {
		v206 = v184
		goto L40
	} else {
		goto L47
	}
L47:
	;
	goto L2
L48:
	;
	v206 = int32(3904)
	goto L40
L49:
	;
	if v188 == int32(3831) {
		goto L48
	} else {
		goto L50
	}
L50:
	;
	if v188 != int32(4537) {
		goto L2
	} else {
		goto L51
	}
L51:
	;
	v206 = int32(4451)
	goto L40
L52:
	;
	m.G0 = v210 + int32(32)
	if v639 == int32(0) {
		goto L1
	} else {
		goto L220
	}
L53:
	;
	v234 = int32(0)
	v235 = v7
	v240 = v7
	v241 = v7
	v242 = v7
	v245 = v7
	v246 = v7
	v247 = v7
	v248 = v7
	v249 = v7
	v251 = v7
	v252 = v7
	v253 = v7
	v254 = v7
	v255 = v7
	v256 = v7
	v257 = v7
	v258 = v7
	v260 = v7
	goto L54
L54:
	;
	if l2 != 0 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v210)+8)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v210)+12)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v210)+4)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v210)+28)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v210)+24)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v210)+20)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v210))) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v210)+16)) = v387
	v408 = int32(1)
	if v375&v408 == int32(0) {
		v639 = v408
		goto L52
	} else {
		goto L130
	}
L56:
	;
	v263 = int32(*(*int8)(unsafe.Add(mBase, uint32(l2+v235))))
	v265 = v263
	goto L58
L57:
	;
	v265 = int32(105)
	goto L58
L58:
	;
	v268 = l1 + v235<<(uint(int32(2))%32)
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v268)))
	if v269 <= int32(3830) {
		goto L70
	} else {
		goto L71
	}
L59:
	;
	switch v265 - int32(111) {
	case 0, 5:
		v396 = v248
		goto L127
	default:
		goto L128
	}
L60:
	;
	v375 = v260
	v376 = v247
	v377 = v234
	v378 = v249
	v379 = v245
	v380 = v241
	v381 = v240
	v382 = v246
	v383 = v242
	v384 = v367
	v385 = v368
	v386 = v369
	v387 = v370
	v388 = v371
	v389 = v372
	v390 = v373
	v391 = v374
	goto L59
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v268))) = v340
	v367 = v349
	v368 = v350
	v369 = v351
	v370 = v352
	v371 = v353
	v372 = v354
	v373 = v355
	v374 = v356
	goto L60
L62:
	;
	v333 = int32(1)
	switch v265 - int32(111) {
	case 0, 5:
		v375 = v333
		v376 = v247
		v377 = v234
		v378 = v249
		v379 = v245
		v380 = v241
		v381 = v333
		v382 = v246
		v383 = v242
		v384 = v251
		v385 = v252
		v386 = v253
		v387 = v254
		v388 = v255
		v389 = v256
		v390 = v257
		v391 = v258
		goto L59
	default:
		goto L121
	}
L63:
	;
	v326 = int32(1)
	switch v265 - int32(111) {
	case 0, 5:
		v375 = v326
		v376 = v247
		v377 = v234
		v378 = v249
		v379 = v245
		v380 = v326
		v381 = v240
		v382 = v246
		v383 = v242
		v384 = v251
		v385 = v252
		v386 = v253
		v387 = v254
		v388 = v255
		v389 = v256
		v390 = v257
		v391 = v258
		goto L59
	default:
		goto L115
	}
L64:
	;
	v319 = int32(1)
	switch v265 - int32(111) {
	case 0, 5:
		v375 = v319
		v376 = v247
		v377 = v234
		v378 = v249
		v379 = v319
		v380 = v241
		v381 = v240
		v382 = v246
		v383 = v242
		v384 = v251
		v385 = v252
		v386 = v253
		v387 = v254
		v388 = v255
		v389 = v256
		v390 = v257
		v391 = v258
		goto L59
	default:
		goto L109
	}
L65:
	;
	v312 = int32(1)
	switch v265 - int32(111) {
	case 0, 5:
		v375 = v312
		v376 = v247
		v377 = v234
		v378 = v312
		v379 = v245
		v380 = v241
		v381 = v240
		v382 = v246
		v383 = v242
		v384 = v251
		v385 = v252
		v386 = v253
		v387 = v254
		v388 = v255
		v389 = v256
		v390 = v257
		v391 = v258
		goto L59
	default:
		goto L103
	}
L66:
	;
	v305 = int32(1)
	switch v265 - int32(111) {
	case 0, 5:
		v375 = v305
		v376 = v247
		v377 = v234
		v378 = v249
		v379 = v245
		v380 = v241
		v381 = v240
		v382 = v246
		v383 = v305
		v384 = v251
		v385 = v252
		v386 = v253
		v387 = v254
		v388 = v255
		v389 = v256
		v390 = v257
		v391 = v258
		goto L59
	default:
		goto L97
	}
L67:
	;
	if v255 != 0 {
		goto L92
	} else {
		goto L93
	}
L68:
	;
	v295 = int32(1)
	switch v265 - int32(111) {
	case 0, 5:
		v375 = v295
		v376 = v247
		v377 = v234
		v378 = v249
		v379 = v245
		v380 = v241
		v381 = v240
		v382 = v295
		v383 = v242
		v384 = v251
		v385 = v252
		v386 = v253
		v387 = v254
		v388 = v255
		v389 = v256
		v390 = v257
		v391 = v258
		goto L59
	default:
		goto L86
	}
L69:
	;
	if v254 != 0 {
		goto L80
	} else {
		goto L81
	}
L70:
	;
	switch v269 - int32(2277) {
	case 0:
		goto L68
	case 1, 2, 3, 4, 5:
		v375 = v260
		v376 = v247
		v377 = v234
		v378 = v249
		v379 = v245
		v380 = v241
		v381 = v240
		v382 = v246
		v383 = v242
		v384 = v251
		v385 = v252
		v386 = v253
		v387 = v254
		v388 = v255
		v389 = v256
		v390 = v257
		v391 = v258
		goto L59
	case 6:
		goto L73
	default:
		goto L74
	}
L71:
	;
	goto L72
L72:
	;
	switch v269 - int32(5077) {
	case 0, 2:
		goto L65
	case 1:
		goto L64
	case 3:
		goto L63
	default:
		goto L77
	}
L73:
	;
	v278 = int32(1)
	switch v265 - int32(111) {
	case 0, 5:
		v375 = v278
		v376 = v278
		v377 = v234
		v378 = v249
		v379 = v245
		v380 = v241
		v381 = v240
		v382 = v246
		v383 = v242
		v384 = v251
		v385 = v252
		v386 = v253
		v387 = v254
		v388 = v255
		v389 = v256
		v390 = v257
		v391 = v258
		goto L59
	default:
		goto L69
	}
L74:
	;
	if v269 == int32(2776) {
		goto L73
	} else {
		goto L75
	}
L75:
	;
	if v269 != int32(3500) {
		v367 = v251
		v368 = v252
		v369 = v253
		v370 = v254
		v371 = v255
		v372 = v256
		v373 = v257
		v374 = v258
		goto L60
	} else {
		goto L76
	}
L76:
	;
	goto L73
L77:
	;
	switch v269 - int32(4537) {
	case 0:
		goto L66
	case 1:
		goto L62
	default:
		goto L78
	}
L78:
	;
	if v269 != int32(3831) {
		v367 = v251
		v368 = v252
		v369 = v253
		v370 = v254
		v371 = v255
		v372 = v256
		v373 = v257
		v374 = v258
		goto L60
	} else {
		goto L79
	}
L79:
	;
	v288 = int32(1)
	switch v265 - int32(111) {
	case 0, 5:
		v375 = v288
		v376 = v247
		v377 = v288
		v378 = v249
		v379 = v245
		v380 = v241
		v381 = v240
		v382 = v246
		v383 = v242
		v384 = v251
		v385 = v252
		v386 = v253
		v387 = v254
		v388 = v255
		v389 = v256
		v390 = v257
		v391 = v258
		goto L59
	default:
		goto L67
	}
L80:
	;
	v340 = v254
	v349 = v251
	v350 = v252
	v351 = v253
	v352 = v254
	v353 = v255
	v354 = v256
	v355 = v257
	v356 = v258
	goto L61
L81:
	;
	goto L82
L82:
	;
	v292 = F_get_call_expr_argtype(m, l3, v248)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	return
L84:
	;
	if v292 != 0 {
		v340 = v292
		v349 = v251
		v350 = v252
		v351 = v253
		v352 = v292
		v353 = v255
		v354 = v256
		v355 = v257
		v356 = v258
		goto L61
	} else {
		goto L85
	}
L85:
	;
	v639 = int32(0)
	goto L52
L86:
	;
	if v251 != 0 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v340 = v251
	v349 = v251
	v350 = v252
	v351 = v253
	v352 = v254
	v353 = v255
	v354 = v256
	v355 = v257
	v356 = v258
	goto L61
L88:
	;
	goto L89
L89:
	;
	v299 = F_get_call_expr_argtype(m, l3, v248)
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L83
	} else {
		goto L90
	}
L90:
	;
	if v299 != 0 {
		v340 = v299
		v349 = v299
		v350 = v252
		v351 = v253
		v352 = v254
		v353 = v255
		v354 = v256
		v355 = v257
		v356 = v258
		goto L61
	} else {
		goto L91
	}
L91:
	;
	v639 = int32(0)
	goto L52
L92:
	;
	v340 = v255
	v349 = v251
	v350 = v252
	v351 = v253
	v352 = v254
	v353 = v255
	v354 = v256
	v355 = v257
	v356 = v258
	goto L61
L93:
	;
	goto L94
L94:
	;
	v302 = F_get_call_expr_argtype(m, l3, v248)
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L83
	} else {
		goto L95
	}
L95:
	;
	if v302 != 0 {
		v340 = v302
		v349 = v251
		v350 = v252
		v351 = v253
		v352 = v254
		v353 = v302
		v354 = v256
		v355 = v257
		v356 = v258
		goto L61
	} else {
		goto L96
	}
L96:
	;
	v639 = int32(0)
	goto L52
L97:
	;
	if v256 != 0 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v340 = v256
	v349 = v251
	v350 = v252
	v351 = v253
	v352 = v254
	v353 = v255
	v354 = v256
	v355 = v257
	v356 = v258
	goto L61
L99:
	;
	goto L100
L100:
	;
	v309 = F_get_call_expr_argtype(m, l3, v248)
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L83
	} else {
		goto L101
	}
L101:
	;
	if v309 != 0 {
		v340 = v309
		v349 = v251
		v350 = v252
		v351 = v253
		v352 = v254
		v353 = v255
		v354 = v309
		v355 = v257
		v356 = v258
		goto L61
	} else {
		goto L102
	}
L102:
	;
	v639 = int32(0)
	goto L52
L103:
	;
	if v253 != 0 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v340 = v253
	v349 = v251
	v350 = v252
	v351 = v253
	v352 = v254
	v353 = v255
	v354 = v256
	v355 = v257
	v356 = v258
	goto L61
L105:
	;
	goto L106
L106:
	;
	v316 = F_get_call_expr_argtype(m, l3, v248)
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L83
	} else {
		goto L107
	}
L107:
	;
	if v316 != 0 {
		v340 = v316
		v349 = v251
		v350 = v252
		v351 = v316
		v352 = v254
		v353 = v255
		v354 = v256
		v355 = v257
		v356 = v258
		goto L61
	} else {
		goto L108
	}
L108:
	;
	v639 = int32(0)
	goto L52
L109:
	;
	if v252 != 0 {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v340 = v252
	v349 = v251
	v350 = v252
	v351 = v253
	v352 = v254
	v353 = v255
	v354 = v256
	v355 = v257
	v356 = v258
	goto L61
L111:
	;
	goto L112
L112:
	;
	v323 = F_get_call_expr_argtype(m, l3, v248)
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L83
	} else {
		goto L113
	}
L113:
	;
	if v323 != 0 {
		v340 = v323
		v349 = v251
		v350 = v323
		v351 = v253
		v352 = v254
		v353 = v255
		v354 = v256
		v355 = v257
		v356 = v258
		goto L61
	} else {
		goto L114
	}
L114:
	;
	v639 = int32(0)
	goto L52
L115:
	;
	if v257 != 0 {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v340 = v257
	v349 = v251
	v350 = v252
	v351 = v253
	v352 = v254
	v353 = v255
	v354 = v256
	v355 = v257
	v356 = v258
	goto L61
L117:
	;
	goto L118
L118:
	;
	v330 = F_get_call_expr_argtype(m, l3, v248)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L83
	} else {
		goto L119
	}
L119:
	;
	if v330 != 0 {
		v340 = v330
		v349 = v251
		v350 = v252
		v351 = v253
		v352 = v254
		v353 = v255
		v354 = v256
		v355 = v330
		v356 = v258
		goto L61
	} else {
		goto L120
	}
L120:
	;
	v639 = int32(0)
	goto L52
L121:
	;
	if v258 != 0 {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v340 = v258
	v349 = v251
	v350 = v252
	v351 = v253
	v352 = v254
	v353 = v255
	v354 = v256
	v355 = v257
	v356 = v258
	goto L61
L123:
	;
	goto L124
L124:
	;
	v337 = F_get_call_expr_argtype(m, l3, v248)
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L83
	} else {
		goto L125
	}
L125:
	;
	if v337 != 0 {
		v340 = v337
		v349 = v251
		v350 = v252
		v351 = v253
		v352 = v254
		v353 = v255
		v354 = v256
		v355 = v257
		v356 = v337
		goto L61
	} else {
		goto L126
	}
L126:
	;
	v639 = int32(0)
	goto L52
L127:
	;
	v398 = v235 + int32(1)
	if v398 != l0 {
		v234 = v377
		v235 = v398
		v240 = v381
		v241 = v380
		v242 = v383
		v245 = v379
		v246 = v382
		v247 = v376
		v248 = v396
		v249 = v378
		v251 = v384
		v252 = v385
		v253 = v386
		v254 = v387
		v255 = v388
		v256 = v389
		v257 = v390
		v258 = v391
		v260 = v375
		goto L54
	} else {
		goto L129
	}
L128:
	;
	v396 = v248 + int32(1)
	goto L127
L129:
	;
	goto L55
L130:
	;
	if base.B2i32(v387 == int32(0))&v376 != 0 {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	F_resolve_anyelement_from_others(m, v210+int32(16))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L83
	} else {
		goto L134
	}
L132:
	;
	v421 = v384
	goto L133
L133:
	;
	if base.B2i32(v421 == int32(0))&v382 != 0 {
		goto L135
	} else {
		goto L136
	}
L134:
	;
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v210)+20))
	v421 = v420
	goto L133
L135:
	;
	F_resolve_anyarray_from_others(m, v210+int32(16))
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L83
	} else {
		goto L138
	}
L136:
	;
	goto L137
L137:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v210)+24))
	if base.B2i32(v429 == int32(0))&v377 != 0 {
		goto L139
	} else {
		goto L140
	}
L138:
	;
	goto L137
L139:
	;
	F_resolve_anyrange_from_others(m, v210+int32(16))
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L83
	} else {
		goto L142
	}
L140:
	;
	goto L141
L141:
	;
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v210)+28))
	if base.B2i32(v437 == int32(0))&v383 != 0 {
		goto L143
	} else {
		goto L144
	}
L142:
	;
	goto L141
L143:
	;
	F_resolve_anymultirange_from_others(m, v210+int32(16))
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L83
	} else {
		goto L146
	}
L144:
	;
	goto L145
L145:
	;
	if base.B2i32(v386 == int32(0))&v378 != 0 {
		goto L147
	} else {
		goto L148
	}
L146:
	;
	goto L145
L147:
	;
	F_resolve_anyelement_from_others(m, v210)
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L83
	} else {
		goto L150
	}
L148:
	;
	v451 = v385
	goto L149
L149:
	;
	if base.B2i32(v451 == int32(0))&v379 != 0 {
		goto L151
	} else {
		goto L152
	}
L150:
	;
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v210)+4))
	v451 = v450
	goto L149
L151:
	;
	F_resolve_anyarray_from_others(m, v210)
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L83
	} else {
		goto L154
	}
L152:
	;
	goto L153
L153:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v210)+8))
	if base.B2i32(v457 == int32(0))&v380 != 0 {
		goto L155
	} else {
		goto L156
	}
L154:
	;
	goto L153
L155:
	;
	F_resolve_anyrange_from_others(m, v210)
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L83
	} else {
		goto L158
	}
L156:
	;
	goto L157
L157:
	;
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v210)+12))
	if base.B2i32(v463 == int32(0))&v381 != 0 {
		goto L159
	} else {
		goto L160
	}
L158:
	;
	goto L157
L159:
	;
	F_resolve_anymultirange_from_others(m, v210)
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L83
	} else {
		goto L162
	}
L160:
	;
	goto L161
L161:
	;
	v469 = int32(1)
	v471 = int32(0)
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v210)+16))
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v210)+20))
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v210)+24))
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v210)+28))
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v210)+4))
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v210)+8))
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v210)+12))
	if l0 != v469 {
		goto L163
	} else {
		goto L164
	}
L162:
	;
	goto L161
L163:
	;
	v496 = v471
	v504 = int32(0)
	goto L166
L164:
	;
	v580 = v471
	goto L165
L165:
	;
	if l0&v469 == int32(0) {
		v639 = v408
		goto L52
	} else {
		goto L203
	}
L166:
	;
	v525 = l1 + v496<<(uint(int32(2))%32)
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v525)))
	if v526 <= int32(3830) {
		goto L176
	} else {
		goto L177
	}
L167:
	;
	v580 = v565
	goto L165
L168:
	;
	v545 = v525 + int32(4)
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v545)))
	if v546 <= int32(3830) {
		goto L188
	} else {
		goto L189
	}
L169:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v525))) = v541
	goto L168
L170:
	;
	v541 = v479
	goto L169
L171:
	;
	v541 = v478
	goto L169
L172:
	;
	v541 = v477
	goto L169
L173:
	;
	v541 = v476
	goto L169
L174:
	;
	v541 = v475
	goto L169
L175:
	;
	v541 = v473
	goto L169
L176:
	;
	switch v526 - int32(2277) {
	case 0:
		goto L175
	case 1, 2, 3, 4, 5:
		goto L168
	case 6:
		v541 = v472
		goto L169
	default:
		goto L179
	}
L177:
	;
	goto L178
L178:
	;
	switch v526 - int32(5077) {
	case 0, 2:
		goto L173
	case 1:
		goto L172
	case 3:
		goto L171
	default:
		goto L182
	}
L179:
	;
	if v526 == int32(2776) {
		v541 = v472
		goto L169
	} else {
		goto L180
	}
L180:
	;
	if v526 == int32(3500) {
		v541 = v472
		goto L169
	} else {
		goto L181
	}
L181:
	;
	goto L168
L182:
	;
	switch v526 - int32(4537) {
	case 0:
		goto L174
	case 1:
		goto L170
	default:
		goto L183
	}
L183:
	;
	if v526 == int32(3831) {
		v541 = v474
		goto L169
	} else {
		goto L184
	}
L184:
	;
	goto L168
L185:
	;
	v564 = int32(2)
	v565 = v496 + v564
	v567 = v504 + v564
	if v567 != l0&int32(2147483646) {
		v496 = v565
		v504 = v567
		goto L166
	} else {
		goto L202
	}
L186:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v545))) = v561
	goto L185
L187:
	;
	v561 = v473
	goto L186
L188:
	;
	switch v546 - int32(2277) {
	case 0:
		goto L187
	case 1, 2, 3, 4, 5:
		goto L185
	case 6:
		v561 = v472
		goto L186
	default:
		goto L191
	}
L189:
	;
	goto L190
L190:
	;
	switch v546 - int32(5077) {
	case 0, 2:
		goto L195
	case 1:
		goto L196
	case 3:
		goto L197
	default:
		goto L198
	}
L191:
	;
	if v546 == int32(2776) {
		v561 = v472
		goto L186
	} else {
		goto L192
	}
L192:
	;
	if v546 == int32(3500) {
		v561 = v472
		goto L186
	} else {
		goto L193
	}
L193:
	;
	goto L185
L194:
	;
	v561 = v475
	goto L186
L195:
	;
	v561 = v476
	goto L186
L196:
	;
	v561 = v477
	goto L186
L197:
	;
	v561 = v478
	goto L186
L198:
	;
	switch v546 - int32(4537) {
	case 0:
		goto L194
	case 1:
		goto L199
	default:
		goto L200
	}
L199:
	;
	v561 = v479
	goto L186
L200:
	;
	if v546 == int32(3831) {
		v561 = v474
		goto L186
	} else {
		goto L201
	}
L201:
	;
	goto L185
L202:
	;
	goto L167
L203:
	;
	v611 = l1 + v580<<(uint(int32(2))%32)
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v611)))
	if v612 <= int32(3830) {
		goto L206
	} else {
		goto L207
	}
L204:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v611))) = v627
	v639 = v408
	goto L52
L205:
	;
	v627 = v473
	goto L204
L206:
	;
	switch v612 - int32(2277) {
	case 0:
		goto L205
	case 1, 2, 3, 4, 5:
		v639 = v408
		goto L52
	case 6:
		v627 = v472
		goto L204
	default:
		goto L209
	}
L207:
	;
	goto L208
L208:
	;
	switch v612 - int32(5077) {
	case 0, 2:
		goto L213
	case 1:
		goto L214
	case 3:
		goto L215
	default:
		goto L216
	}
L209:
	;
	if v612 == int32(2776) {
		v627 = v472
		goto L204
	} else {
		goto L210
	}
L210:
	;
	if v612 == int32(3500) {
		v627 = v472
		goto L204
	} else {
		goto L211
	}
L211:
	;
	v639 = v408
	goto L52
L212:
	;
	v627 = v475
	goto L204
L213:
	;
	v627 = v476
	goto L204
L214:
	;
	v627 = v477
	goto L204
L215:
	;
	v627 = v478
	goto L204
L216:
	;
	switch v612 - int32(4537) {
	case 0:
		goto L212
	case 1:
		goto L217
	default:
		goto L218
	}
L217:
	;
	v627 = v479
	goto L204
L218:
	;
	if v612 == int32(3831) {
		v627 = v474
		goto L204
	} else {
		goto L219
	}
L219:
	;
	v639 = v408
	goto L52
L220:
	;
	if l0 <= int32(0) {
		goto L2
	} else {
		goto L221
	}
L221:
	;
	v674 = int32(0)
	v680 = v674
	v681 = v674
	goto L222
L222:
	;
	if l2 == int32(0) {
		goto L225
	} else {
		goto L226
	}
L223:
	;
	goto L2
L224:
	;
	v741 = v680 + int32(1)
	if v741 != l0 {
		v680 = v741
		v681 = v737
		goto L222
	} else {
		goto L231
	}
L225:
	;
	v722 = l1 + v680<<(uint(int32(2))%32)
	v723 = *(*int32)(unsafe.Add(mBase, uint32(v722)))
	if base.B2i32(v723 != int32(2287))&base.B2i32(v723 != int32(2249)) != 0 {
		goto L227
	} else {
		goto L228
	}
L226:
	;
	v717 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v680))))
	switch v717 - int32(111) {
	case 0, 5:
		v737 = v681
		goto L224
	default:
		goto L225
	}
L227:
	;
	v737 = v681 + int32(1)
	goto L224
L228:
	;
	v729 = F_get_call_expr_argtype(m, l3, v681)
	mBase = m.M
	v730 = m.ExcPending
	if v730 != 0 {
		goto L83
	} else {
		goto L229
	}
L229:
	;
	if v729 == int32(0) {
		goto L227
	} else {
		goto L230
	}
L230:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v722))) = v729
	goto L227
L231:
	;
	goto L223
L232:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v790 = m.ExcPending
	if v790 != 0 {
		goto L83
	} else {
		goto L233
	}
L233:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41))) = l5
	F_errmsg(m, int32(661880), v41)
	mBase = m.M
	v794 = m.ExcPending
	if v794 != 0 {
		goto L83
	} else {
		goto L234
	}
L234:
	;
	F_errfinish(m, int32(475861), int32(366), int32(151847))
	mBase = m.M
	v799 = m.ExcPending
	if v799 != 0 {
		goto L83
	} else {
		goto L235
	}
L235:
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
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
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
				v17 = *(*int32)(unsafe.Add(mBase, _consts[1003]))
				if v14 != 0 {
					if v14 == int32(-1) {
						v22 = int32(4606656)
					} else {
						v22 = v14
					}
					*(*int32)(unsafe.Add(mBase, _consts[1003])) = v22
				} else {
				}
				if v17 == int32(4606656) {
					v27 = int32(-1)
				} else {
					v27 = v17
				}
				v28 = F_mbstowcs(m, l0, v9, l1)
				mBase = m.M
				if v27 != 0 {
					if v27 == int32(-1) {
						v36 = int32(4606656)
					} else {
						v36 = v27
					}
					*(*int32)(unsafe.Add(mBase, _consts[1003])) = v36
				} else {
				}
				v42 = v28
			}
			F_pfree(m, v9)
			mBase = m.M
			v45 = m.ExcPending
			if v45 != 0 {
				return
			} else {
				if v42 != int32(-1) {
					return
				} else {
					F_pg_verifymbstr(m, l2, l3)
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return
						} else {
							F_errcode(m, int32(17301634))
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return
							} else {
								F_errmsg(m, int32(378831), int32(0))
								mBase = m.M
								v60 = m.ExcPending
								if v60 != 0 {
									return
								} else {
									F_errhint(m, int32(586682), int32(0))
									mBase = m.M
									v64 = m.ExcPending
									if v64 != 0 {
										return
									} else {
										F_errfinish(m, int32(476647), int32(1001), int32(218582))
										mBase = m.M
										v69 = m.ExcPending
										if v69 != 0 {
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
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
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
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v283 int32
	_ = v283
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
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
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v368 int32
	_ = v368
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v392 int32
	_ = v392
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v435 int32
	_ = v435
	var v451 int32
	_ = v451
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v525 int32
	_ = v525
	var v531 int32
	_ = v531
	var v543 int32
	_ = v543
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v554 int32
	_ = v554
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v562 int32
	_ = v562
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v573 int32
	_ = v573
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v589 int32
	_ = v589
	var v591 int32
	_ = v591
	v3 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(48)
	m.G0 = v16
	if l0 == v3 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	m.G0 = v16 + int32(48)
	return base.B2i32(v589 == int32(0)) & v591
L2:
	;
	v582 = F_raw_expression_tree_walker_impl(m, v568, int32(485), l1)
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L23
	} else {
		goto L169
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L23
	} else {
		goto L164
	}
L4:
	;
	v589 = v531
	v591 = int32(0)
	goto L1
L5:
	;
	v531 = v3
	goto L4
L6:
	;
	goto L7
L7:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v20 <= int32(109) {
		goto L14
	} else {
		goto L15
	}
L8:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v335)+8))
	if v348 != 0 {
		v589 = v340
		v591 = v3
		goto L1
	} else {
		goto L120
	}
L9:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v194)+64))
	if v207 != 0 {
		goto L89
	} else {
		goto L90
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L23
	} else {
		goto L86
	}
L11:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v82+l0)))
	if v84 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L12:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	switch v41 {
	case 0:
		goto L26
	case 1:
		goto L27
	case 2:
		goto L28
	case 3:
		goto L29
	default:
		v165 = l0
		goto L10
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = int32(2)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v35 = F_checkWellFormedRecursionWalker(m, v34, l1)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L23
	} else {
		goto L24
	}
L14:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	switch v20 - int32(3) {
	case 0:
		v335 = l0
		v337 = v23
		v340 = v3
		goto L8
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18:
		v568 = l0
		v573 = v3
		goto L2
	case 19:
		goto L13
	default:
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	if v20 == int32(110) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	if v20 == int32(64) {
		goto L12
	} else {
		goto L18
	}
L18:
	;
	v568 = l0
	v573 = v3
	goto L2
L19:
	;
	v589 = v3
	v591 = v3
	goto L1
L20:
	;
	goto L21
L21:
	;
	if v20 == int32(141) {
		v194 = l0
		v199 = v3
		goto L9
	} else {
		goto L22
	}
L22:
	;
	v568 = l0
	v573 = v3
	goto L2
L23:
	;
	return int32(0)
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v23
	v82 = int32(12)
	goto L11
L25:
	;
	v82 = int32(28)
	goto L11
L26:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v76 = F_checkWellFormedRecursionWalker(m, v75, l1)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L23
	} else {
		goto L45
	}
L27:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v65 = F_checkWellFormedRecursionWalker(m, v64, l1)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L23
	} else {
		goto L40
	}
L28:
	;
	if v23 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L29:
	;
	if v23 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = int32(3)
	goto L32
L31:
	;
	goto L32
L32:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v47 = F_checkWellFormedRecursionWalker(m, v46, l1)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L23
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v23
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v51 = F_checkWellFormedRecursionWalker(m, v50, l1)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L23
	} else {
		goto L34
	}
L34:
	;
	goto L25
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = int32(3)
	goto L37
L36:
	;
	goto L37
L37:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v58 = F_checkWellFormedRecursionWalker(m, v57, l1)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L23
	} else {
		goto L38
	}
L38:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v61 = F_checkWellFormedRecursionWalker(m, v60, l1)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L23
	} else {
		goto L39
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v23
	goto L25
L40:
	;
	if v23 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = int32(3)
	goto L43
L42:
	;
	goto L43
L43:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v72 = F_checkWellFormedRecursionWalker(m, v71, l1)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L23
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v23
	goto L25
L45:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v79 = F_checkWellFormedRecursionWalker(m, v78, l1)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L23
	} else {
		goto L46
	}
L46:
	;
	goto L25
L47:
	;
	v589 = int32(1)
	v591 = v3
	goto L1
L48:
	;
	goto L49
L49:
	;
	v88 = v84
	goto L50
L50:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v102 = int32(1)
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	if v103 <= int32(63) {
		goto L55
	} else {
		goto L56
	}
L51:
	;
	v589 = v102
	v591 = v3
	goto L1
L52:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v162+v88)))
	if v164 != 0 {
		v88 = v164
		goto L50
	} else {
		goto L85
	}
L53:
	;
	v162 = int32(28)
	goto L52
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = int32(2)
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v88)+20))
	v157 = F_checkWellFormedRecursionWalker(m, v156, l1)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L23
	} else {
		goto L84
	}
L55:
	;
	switch v103 - int32(3) {
	case 0:
		v335 = v88
		v337 = v101
		v340 = v102
		goto L8
	default:
		v568 = v88
		v573 = v102
		goto L2
	case 19:
		goto L54
	}
L56:
	;
	goto L57
L57:
	;
	if v103 != int32(64) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	if v103 == int32(110) {
		v589 = v102
		v591 = v3
		goto L1
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v88)+4))
	switch v114 {
	case 0:
		goto L66
	case 1:
		goto L65
	case 2:
		goto L64
	case 3:
		goto L63
	default:
		v165 = v88
		goto L10
	}
L61:
	;
	if v103 == int32(141) {
		v194 = v88
		v199 = v102
		goto L9
	} else {
		goto L62
	}
L62:
	;
	v568 = v88
	v573 = v102
	goto L2
L63:
	;
	if v101 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L64:
	;
	if v101 == int32(0) {
		goto L74
	} else {
		goto L75
	}
L65:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v88)+12))
	v122 = F_checkWellFormedRecursionWalker(m, v121, l1)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L23
	} else {
		goto L69
	}
L66:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v88)+12))
	v116 = F_checkWellFormedRecursionWalker(m, v115, l1)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L23
	} else {
		goto L67
	}
L67:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v88)+16))
	v119 = F_checkWellFormedRecursionWalker(m, v118, l1)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L23
	} else {
		goto L68
	}
L68:
	;
	goto L53
L69:
	;
	if v101 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = int32(3)
	goto L72
L71:
	;
	goto L72
L72:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v88)+16))
	v129 = F_checkWellFormedRecursionWalker(m, v128, l1)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L23
	} else {
		goto L73
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v101
	goto L53
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = int32(3)
	goto L76
L75:
	;
	goto L76
L76:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v88)+12))
	v137 = F_checkWellFormedRecursionWalker(m, v136, l1)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L23
	} else {
		goto L77
	}
L77:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v88)+16))
	v140 = F_checkWellFormedRecursionWalker(m, v139, l1)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L23
	} else {
		goto L78
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v101
	goto L53
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = int32(3)
	goto L81
L80:
	;
	goto L81
L81:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v88)+12))
	v148 = F_checkWellFormedRecursionWalker(m, v147, l1)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L23
	} else {
		goto L82
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v101
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v88)+16))
	v152 = F_checkWellFormedRecursionWalker(m, v151, l1)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L23
	} else {
		goto L83
	}
L83:
	;
	goto L53
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v101
	v162 = int32(12)
	goto L52
L85:
	;
	goto L51
L86:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v165)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v182
	F_errmsg_internal(m, int32(462600), v16+int32(32))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L23
	} else {
		goto L87
	}
L87:
	;
	F_errfinish(m, int32(475120), int32(1179), int32(211014))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L23
	} else {
		goto L88
	}
L88:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L89:
	;
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+8)))
	if v208 == int32(1) {
		goto L92
	} else {
		goto L93
	}
L90:
	;
	goto L91
L91:
	;
	F_checkWellFormedSelectStmt(m, v194, l1)
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L23
	} else {
		goto L119
	}
L92:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v207)+4))
	v212 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v213 = F_lcons(m, v211, v212)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L23
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	v268 = int32(0)
	v270 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v271 = F_lcons(m, v268, v270)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L23
	} else {
		goto L105
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v213
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v194)+64))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v216)+4))
	if v217 == int32(0) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	F_checkWellFormedSelectStmt(m, v194, l1)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L23
	} else {
		goto L103
	}
L97:
	;
	v220 = int32(0)
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v217)+4))
	if v221 <= v220 {
		goto L96
	} else {
		goto L98
	}
L98:
	;
	v226 = v220
	goto L99
L99:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v217)+12))
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v237+v226<<(uint(int32(2))%32))))
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v241)+16))
	v243 = F_checkWellFormedRecursionWalker(m, v242, l1)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L23
	} else {
		goto L101
	}
L100:
	;
	goto L96
L101:
	;
	v246 = v226 + int32(1)
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v217)+4))
	if v246 < v247 {
		v226 = v246
		goto L99
	} else {
		goto L102
	}
L102:
	;
	goto L100
L103:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v265 = F_list_delete_first(m, v264)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L23
	} else {
		goto L104
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v265
	v589 = v199
	v591 = v3
	goto L1
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v271
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v194)+64))
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v274)+4))
	if v275 == int32(0) {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	F_checkWellFormedSelectStmt(m, v194, l1)
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L23
	} else {
		goto L117
	}
L107:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v275)+4))
	if v278 <= int32(0) {
		goto L106
	} else {
		goto L108
	}
L108:
	;
	v283 = v268
	goto L109
L109:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v275)+12))
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v294+v283<<(uint(int32(2))%32))))
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v298)+16))
	v300 = F_checkWellFormedRecursionWalker(m, v299, l1)
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L23
	} else {
		goto L111
	}
L110:
	;
	goto L106
L111:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v302 != 0 {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v302)+12))
	v305 = v303
	goto L114
L113:
	;
	v305 = int32(0)
	goto L114
L114:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v305)))
	v307 = F_lappend(m, v306, v298)
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L23
	} else {
		goto L115
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v305))) = v307
	v311 = v283 + int32(1)
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v275)+4))
	if v311 < v312 {
		v283 = v311
		goto L109
	} else {
		goto L116
	}
L116:
	;
	goto L110
L117:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v330 = F_list_delete_first(m, v329)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L23
	} else {
		goto L118
	}
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v330
	v531 = v199
	goto L4
L119:
	;
	v589 = v199
	v591 = v3
	goto L1
L120:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v349 == int32(0) {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v335)+12))
	v467 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v468 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v467+v468*int32(12))))
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v472)+4))
	v476 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v473))))
	v477 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v466))))
	if v477 == int32(0) {
		v496 = v476
		v497 = v477
		goto L149
	} else {
		goto L150
	}
L122:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v349)+4))
	if v352 <= int32(0) {
		goto L121
	} else {
		goto L123
	}
L123:
	;
	v355 = int32(0)
	if v355 < v352 {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v358 = v352
	goto L126
L125:
	;
	v358 = v355
	goto L126
L126:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v349)+12))
	v368 = v3
	goto L127
L127:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v359+v368<<(uint(int32(2))%32))))
	if v376 == int32(0) {
		goto L129
	} else {
		goto L130
	}
L128:
	;
	goto L121
L129:
	;
	v451 = v368 + int32(1)
	if v451 != v358 {
		v368 = v451
		goto L127
	} else {
		goto L147
	}
L130:
	;
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v376)+4))
	if v379 <= int32(0) {
		goto L129
	} else {
		goto L131
	}
L131:
	;
	v382 = int32(0)
	if v382 < v379 {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v385 = v379
	goto L134
L133:
	;
	v385 = v382
	goto L134
L134:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v335)+12))
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v376)+12))
	v392 = int32(0)
	goto L135
L135:
	;
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v387+v392<<(uint(int32(2))%32))))
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v405)+4))
	v409 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v406))))
	v410 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v386))))
	if v410 == int32(0) {
		v429 = v409
		v430 = v410
		goto L138
	} else {
		goto L139
	}
L136:
	;
	goto L129
L137:
	;
	if v430-v429 == int32(0) {
		v531 = v340
		goto L4
	} else {
		goto L145
	}
L138:
	;
	goto L137
L139:
	;
	if v409 != v410 {
		v429 = v409
		v430 = v410
		goto L138
	} else {
		goto L140
	}
L140:
	;
	v414 = v386
	v415 = v406
	goto L141
L141:
	;
	v418 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v415)+1)))
	v419 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v414)+1)))
	if v419 == int32(0) {
		v429 = v418
		v430 = v419
		goto L138
	} else {
		goto L143
	}
L142:
	;
	v429 = v418
	v430 = v419
	goto L138
L143:
	;
	v422 = int32(1)
	if v418 == v419 {
		v414 = v414 + v422
		v415 = v415 + v422
		goto L141
	} else {
		goto L144
	}
L144:
	;
	goto L142
L145:
	;
	v435 = v392 + int32(1)
	if v435 != v385 {
		v392 = v435
		goto L135
	} else {
		goto L146
	}
L146:
	;
	goto L136
L147:
	;
	goto L128
L148:
	;
	if v497-v496 != 0 {
		v589 = v340
		v591 = v3
		goto L1
	} else {
		goto L156
	}
L149:
	;
	goto L148
L150:
	;
	if v476 != v477 {
		v496 = v476
		v497 = v477
		goto L149
	} else {
		goto L151
	}
L151:
	;
	v481 = v466
	v482 = v473
	goto L152
L152:
	;
	v485 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v482)+1)))
	v486 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v481)+1)))
	if v486 == int32(0) {
		v496 = v485
		v497 = v486
		goto L149
	} else {
		goto L154
	}
L153:
	;
	v496 = v485
	v497 = v486
	goto L149
L154:
	;
	v489 = int32(1)
	if v485 == v486 {
		v481 = v481 + v489
		v482 = v482 + v489
		goto L152
	} else {
		goto L155
	}
L155:
	;
	goto L153
L156:
	;
	if v337 != 0 {
		goto L3
	} else {
		goto L157
	}
L157:
	;
	v499 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v501 = v499 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v501
	if v501 < int32(2) {
		v589 = v340
		v591 = v3
		goto L1
	} else {
		goto L158
	}
L158:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L23
	} else {
		goto L159
	}
L159:
	;
	F_errcode(m, int32(151388292))
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L23
	} else {
		goto L160
	}
L160:
	;
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v472)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v512
	F_errmsg(m, int32(395256), v16)
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L23
	} else {
		goto L161
	}
L161:
	;
	v517 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v335)+24))
	F_parser_errposition(m, v517, v518)
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L23
	} else {
		goto L162
	}
L162:
	;
	F_errfinish(m, int32(475120), int32(1077), int32(211014))
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L23
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
	v546 = m.ExcPending
	if v546 != 0 {
		goto L23
	} else {
		goto L165
	}
L165:
	;
	v547 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v472)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v548
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v547<<(uint(int32(2))%32))+uint32(_consts[444])))
	F_errmsg(m, v554, v16+int32(16))
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L23
	} else {
		goto L166
	}
L166:
	;
	v559 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v335)+24))
	F_parser_errposition(m, v559, v560)
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L23
	} else {
		goto L167
	}
L167:
	;
	F_errfinish(m, int32(475120), int32(1069), int32(211014))
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L23
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
	v589 = v573
	v591 = v582
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
				F_errmsg_internal(m, int32(460164), v6)
				mBase = m.M
				v73 = m.ExcPending
				if v73 != 0 {
					return
				} else {
					F_errfinish(m, int32(475120), int32(1267), int32(91225))
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
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v76 int32
	_ = v76
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v118 int32
	_ = v118
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v171 int32
	_ = v171
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v191 int32
	_ = v191
	v4 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v13 = F_pstrdup(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v10 + int32(16)
	return v191
L2:
	;
	return int32(0)
L3:
	;
	v20 = F_SplitIdentifierString(m, v13, int32(44), v10+int32(12))
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
	v25 = *(*int32)(unsafe.Add(mBase, _consts[40]))
	*(*int32)(unsafe.Add(mBase, _consts[497])) = v25
	goto L8
L6:
	;
	goto L7
L7:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	if v39 == int32(0) {
		v153 = v4
		goto L13
	} else {
		goto L14
	}
L8:
	;
	v31 = F_format_elog_string(m, int32(605174), int32(0))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, _consts[498])) = v31
	F_pfree(m, v13)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	F_list_free(m, v36)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	v191 = v4
	goto L1
L12:
	;
	v171 = *(*int32)(unsafe.Add(mBase, _consts[40]))
	*(*int32)(unsafe.Add(mBase, _consts[497])) = v171
	goto L53
L13:
	;
	F_pfree(m, v13)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L2
	} else {
		goto L49
	}
L14:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	if v42 <= int32(0) {
		v153 = v4
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v46 = int32(0)
	v49 = v4
	goto L16
L16:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v53+v46<<(uint(int32(2))%32))))
	v61 = v57
	v62 = int32(497268)
	goto L19
L17:
	;
	v153 = v145
	goto L13
L18:
	;
	if v99 != 0 {
		goto L31
	} else {
		goto L32
	}
L19:
	;
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62))))
	if v65 == v66 {
		v88 = v65
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v99 = int32(0)
	goto L18
L21:
	;
	v90 = int32(1)
	if v88 != 0 {
		v61 = v61 + v90
		v62 = v62 + v90
		goto L19
	} else {
		goto L30
	}
L22:
	;
	if base.Ui32((v65-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v76 = v65 | int32(32)
	goto L25
L24:
	;
	v76 = v65
	goto L25
L25:
	;
	if base.Ui32((v66-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v85 = v66 | int32(32)
	goto L28
L27:
	;
	v85 = v66
	goto L28
L28:
	;
	if v76 == v85 {
		v88 = v76
		goto L21
	} else {
		goto L29
	}
L29:
	;
	v99 = v76 - v85
	goto L18
L30:
	;
	goto L20
L31:
	;
	v103 = v57
	v104 = int32(496999)
	goto L35
L32:
	;
	v144 = int32(4)
	goto L33
L33:
	;
	v145 = v49 | v144
	v147 = v46 + int32(1)
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	if v147 < v148 {
		v46 = v147
		v49 = v145
		goto L16
	} else {
		goto L48
	}
L34:
	;
	if v141 != 0 {
		goto L12
	} else {
		goto L47
	}
L35:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103))))
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104))))
	if v107 == v108 {
		v130 = v107
		goto L37
	} else {
		goto L38
	}
L36:
	;
	v141 = int32(0)
	goto L34
L37:
	;
	v132 = int32(1)
	if v130 != 0 {
		v103 = v103 + v132
		v104 = v104 + v132
		goto L35
	} else {
		goto L46
	}
L38:
	;
	if base.Ui32((v107-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v118 = v107 | int32(32)
	goto L41
L40:
	;
	v118 = v107
	goto L41
L41:
	;
	if base.Ui32((v108-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v127 = v108 | int32(32)
	goto L44
L43:
	;
	v127 = v108
	goto L44
L44:
	;
	if v118 == v127 {
		v130 = v118
		goto L37
	} else {
		goto L45
	}
L45:
	;
	v141 = v118 - v127
	goto L34
L46:
	;
	goto L36
L47:
	;
	v144 = int32(2)
	goto L33
L48:
	;
	goto L17
L49:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	F_list_free(m, v159)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L2
	} else {
		goto L50
	}
L50:
	;
	v163 = F_guc_malloc(m, int32(4))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L2
	} else {
		goto L51
	}
L51:
	;
	if v163 == int32(0) {
		v191 = v4
		goto L1
	} else {
		goto L52
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v163))) = v153
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v163
	v191 = int32(1)
	goto L1
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v57
	v177 = F_format_elog_string(m, int32(626402), v10)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L2
	} else {
		goto L54
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, _consts[498])) = v177
	F_pfree(m, v13)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L2
	} else {
		goto L55
	}
L55:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	F_list_free(m, v182)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L2
	} else {
		goto L56
	}
L56:
	;
	v191 = v4
	goto L1
}
func F_check_debug_io_direct(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v79 int32
	_ = v79
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v102 int32
	_ = v102
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v124 int32
	_ = v124
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v168 int32
	_ = v168
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v220 int32
	_ = v220
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v240 int32
	_ = v240
	v4 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v13 = F_pstrdup(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v10 + int32(32)
	return v240
L2:
	;
	return int32(0)
L3:
	;
	v19 = F_SplitGUCList(m, v13, v10+int32(28))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	if v19 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v24 = *(*int32)(unsafe.Add(mBase, _consts[40]))
	*(*int32)(unsafe.Add(mBase, _consts[497])) = v24
	goto L8
L6:
	;
	goto L7
L7:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
	if v41 == int32(0) {
		v202 = v4
		goto L13
	} else {
		goto L14
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = int32(102310)
	v33 = F_format_elog_string(m, int32(623778), v10+int32(16))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, _consts[498])) = v33
	F_pfree(m, v13)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
	F_list_free(m, v38)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	v240 = v4
	goto L1
L12:
	;
	v220 = *(*int32)(unsafe.Add(mBase, _consts[40]))
	*(*int32)(unsafe.Add(mBase, _consts[497])) = v220
	goto L66
L13:
	;
	F_pfree(m, v13)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L2
	} else {
		goto L62
	}
L14:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	if v44 <= int32(0) {
		v202 = v4
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v48 = int32(0)
	v51 = v4
	goto L16
L16:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v41)+12))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v56+v48<<(uint(int32(2))%32))))
	v64 = v60
	v65 = int32(481619)
	goto L20
L17:
	;
	v202 = v194
	goto L13
L18:
	;
	v194 = v51 | v193
	v196 = v48 + int32(1)
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	if v196 < v197 {
		v48 = v196
		v51 = v194
		goto L16
	} else {
		goto L61
	}
L19:
	;
	if v102 == int32(0) {
		v193 = int32(1)
		goto L18
	} else {
		goto L32
	}
L20:
	;
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64))))
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
	if v68 == v69 {
		v91 = v68
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v102 = int32(0)
	goto L19
L22:
	;
	v93 = int32(1)
	if v91 != 0 {
		v64 = v64 + v93
		v65 = v65 + v93
		goto L20
	} else {
		goto L31
	}
L23:
	;
	if base.Ui32((v68-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v79 = v68 | int32(32)
	goto L26
L25:
	;
	v79 = v68
	goto L26
L26:
	;
	if base.Ui32((v69-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v88 = v69 | int32(32)
	goto L29
L28:
	;
	v88 = v69
	goto L29
L29:
	;
	if v79 == v88 {
		v91 = v79
		goto L22
	} else {
		goto L30
	}
L30:
	;
	v102 = v79 - v88
	goto L19
L31:
	;
	goto L21
L32:
	;
	v109 = v60
	v110 = int32(293395)
	goto L34
L33:
	;
	if v147 == int32(0) {
		v193 = int32(2)
		goto L18
	} else {
		goto L46
	}
L34:
	;
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109))))
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110))))
	if v113 == v114 {
		v136 = v113
		goto L36
	} else {
		goto L37
	}
L35:
	;
	v147 = int32(0)
	goto L33
L36:
	;
	v138 = int32(1)
	if v136 != 0 {
		v109 = v109 + v138
		v110 = v110 + v138
		goto L34
	} else {
		goto L45
	}
L37:
	;
	if base.Ui32((v113-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v124 = v113 | int32(32)
	goto L40
L39:
	;
	v124 = v113
	goto L40
L40:
	;
	if base.Ui32((v114-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v133 = v114 | int32(32)
	goto L43
L42:
	;
	v133 = v114
	goto L43
L43:
	;
	if v124 == v133 {
		v136 = v124
		goto L36
	} else {
		goto L44
	}
L44:
	;
	v147 = v124 - v133
	goto L33
L45:
	;
	goto L35
L46:
	;
	v153 = v60
	v154 = int32(93959)
	goto L48
L47:
	;
	if v191 != 0 {
		goto L12
	} else {
		goto L60
	}
L48:
	;
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153))))
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154))))
	if v157 == v158 {
		v180 = v157
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v191 = int32(0)
	goto L47
L50:
	;
	v182 = int32(1)
	if v180 != 0 {
		v153 = v153 + v182
		v154 = v154 + v182
		goto L48
	} else {
		goto L59
	}
L51:
	;
	if base.Ui32((v157-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v168 = v157 | int32(32)
	goto L54
L53:
	;
	v168 = v157
	goto L54
L54:
	;
	if base.Ui32((v158-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v177 = v158 | int32(32)
	goto L57
L56:
	;
	v177 = v158
	goto L57
L57:
	;
	if v168 == v177 {
		v180 = v168
		goto L50
	} else {
		goto L58
	}
L58:
	;
	v191 = v168 - v177
	goto L47
L59:
	;
	goto L49
L60:
	;
	v193 = int32(4)
	goto L18
L61:
	;
	goto L17
L62:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
	F_list_free(m, v208)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L2
	} else {
		goto L63
	}
L63:
	;
	v212 = F_guc_malloc(m, int32(4))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L2
	} else {
		goto L64
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v212
	if v212 == int32(0) {
		v240 = v4
		goto L1
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v212))) = v202
	v240 = int32(1)
	goto L1
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v60
	v226 = F_format_elog_string(m, int32(624451), v10)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L2
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, _consts[498])) = v226
	F_pfree(m, v13)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L2
	} else {
		goto L68
	}
L68:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
	F_list_free(m, v231)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L2
	} else {
		goto L69
	}
L69:
	;
	v240 = v4
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
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v56 int32
	_ = v56
	v3 = int32(0)
	if l0 == v3 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_scanner_yyerror(m, int32(688381), l1)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L15
	} else {
		goto L16
	}
L2:
	;
	return l0
L3:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v10 <= int32(0) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v13 = int32(0)
	if v13 < v10 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v16 = v10
	goto L7
L6:
	;
	v16 = v13
	goto L7
L7:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v24 = v3
	goto L8
L8:
	;
	v30 = v17 + v24<<(uint(int32(2))%32)
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	if v32 != int32(77) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	goto L2
L10:
	;
	v42 = v24 + int32(1)
	if v42 != v16 {
		v24 = v42
		goto L8
	} else {
		goto L14
	}
L11:
	;
	v36 = v30 + int32(4)
	if v36 == int32(0) {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	if base.Ui32(v36) < base.Ui32(v17+v10<<(uint(int32(2))%32)) {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	goto L10
L14:
	;
	goto L9
L15:
	;
	return int32(0)
L16:
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
				F_errcode(m, int32(393348))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v17
					F_errmsg(m, int32(675551), v9+int32(16))
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return
					} else {
						v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
						if v31 == int32(0) {
							F_errdetail(m, int32(603044), int32(0))
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
									F_errfinish(m, int32(473085), int32(509), int32(299841))
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
								F_errdetail(m, int32(603044), int32(0))
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
										F_errfinish(m, int32(473085), int32(509), int32(299841))
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
								F_errhint(m, int32(533629), v9)
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
										F_errfinish(m, int32(473085), int32(509), int32(299841))
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
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v37 int32
	_ = v37
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v121 int32
	_ = v121
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
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	v9 = m.G0
	v11 = v9 - int32(80)
	m.G0 = v11
	*(*int32)(unsafe.Add(mBase, uint32(v11)+60)) = int32(25)
	if l2 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L19
	} else {
		goto L77
	}
L2:
	;
	m.G0 = v11 + int32(80)
	return
L3:
	;
	if l3 == int32(0) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v22 = l3
	v23 = int32(354802)
	goto L6
L5:
	;
	if v60 != 0 {
		goto L2
	} else {
		goto L18
	}
L6:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	if v26 == v27 {
		v49 = v26
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v60 = int32(0)
	goto L5
L8:
	;
	v51 = int32(1)
	if v49 != 0 {
		v22 = v22 + v51
		v23 = v23 + v51
		goto L6
	} else {
		goto L17
	}
L9:
	;
	if base.Ui32((v26-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v37 = v26 | int32(32)
	goto L12
L11:
	;
	v37 = v26
	goto L12
L12:
	;
	if base.Ui32((v27-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v46 = v27 | int32(32)
	goto L15
L14:
	;
	v46 = v27
	goto L15
L15:
	;
	if v37 == v46 {
		v49 = v37
		goto L8
	} else {
		goto L16
	}
L16:
	;
	v60 = v37 - v46
	goto L5
L17:
	;
	goto L7
L18:
	;
	F_initStringInfo(m, v11-int32(-64))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	return
L20:
	;
	F_appendStringInfoString(m, v11-int32(-64), int32(644446))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	F_GetPublicationsStr(m, l1, v11-int32(-64), int32(1))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L19
	} else {
		goto L22
	}
L22:
	;
	F_appendStringInfoString(m, v11-int32(-64), int32(714475))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L19
	} else {
		goto L23
	}
L23:
	;
	if int32(0) < l5 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v86 = int32(0)
	goto L27
L25:
	;
	goto L26
L26:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
	v126 = *(*int32)(unsafe.Add(mBase, _consts[484]))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)+60))
	v128 = m.T0[v127].(func(*base.Module, int32, int32, int32, int32) int32)(m, l0, v121, int32(1), v11+int32(60))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L19
	} else {
		goto L34
	}
L27:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l4+v86<<(uint(int32(2))%32))))
	v95 = F_get_rel_namespace(m, v94)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L19
	} else {
		goto L29
	}
L28:
	;
	goto L26
L29:
	;
	v97 = F_get_namespace_name(m, v95)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L19
	} else {
		goto L30
	}
L30:
	;
	v99 = F_get_rel_name(m, v94)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L19
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v97
	F_appendStringInfo(m, v11-int32(-64), int32(714429), v11+int32(48))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L19
	} else {
		goto L32
	}
L32:
	;
	v111 = v86 + int32(1)
	if v111 != l5 {
		v86 = v111
		goto L27
	} else {
		goto L33
	}
L33:
	;
	goto L28
L34:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
	F_pfree(m, v130)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L19
	} else {
		goto L35
	}
L35:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v128)))
	if v133 != int32(2) {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v128)+16))
	v138 = F_MakeSingleTupleTableSlot(m, v136, int32(1575788))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L19
	} else {
		goto L37
	}
L37:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v128)+12))
	v143 = F_tuplestore_gettupleslot(m, v140, int32(1), int32(0), v138)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L19
	} else {
		goto L39
	}
L38:
	;
	F_ExecDropSingleTupleTableSlot(m, v138)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L19
	} else {
		goto L63
	}
L39:
	;
	if v143 == int32(0) {
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v150 = int32(0)
	goto L41
L41:
	;
	v156 = int32(*(*int16)(unsafe.Add(mBase, uint32(v138)+6)))
	if v156 <= int32(0) {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	if v172 == int32(0) {
		goto L38
	} else {
		goto L53
	}
L43:
	;
	F_slot_getsomeattrs_int(m, v138, int32(1))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L19
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v138)+16))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v162)))
	v164 = F_text_to_cstring(m, v163)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L19
	} else {
		goto L47
	}
L46:
	;
	goto L45
L47:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v138)+8))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v166)+12))
	m.T0[v167].(func(*base.Module, int32))(m, v138)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L19
	} else {
		goto L48
	}
L48:
	;
	v170 = F_makeString(m, v164)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L19
	} else {
		goto L49
	}
L49:
	;
	v172 = F_list_append_unique(m, v150, v170)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L19
	} else {
		goto L50
	}
L50:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v128)+12))
	v177 = F_tuplestore_gettupleslot(m, v174, int32(1), int32(0), v138)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L19
	} else {
		goto L51
	}
L51:
	;
	if v177 != 0 {
		v150 = v172
		goto L41
	} else {
		goto L52
	}
L52:
	;
	goto L42
L53:
	;
	v181 = F_makeStringInfo(m)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L19
	} else {
		goto L54
	}
L54:
	;
	F_GetPublicationsStr(m, v172, v181, int32(0))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L19
	} else {
		goto L55
	}
L55:
	;
	v188 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L19
	} else {
		goto L56
	}
L56:
	;
	if v188 == int32(0) {
		goto L38
	} else {
		goto L57
	}
L57:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L19
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = l6
	F_errmsg(m, int32(262954), v11+int32(16))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L19
	} else {
		goto L59
	}
L59:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v172)+4))
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v181)))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v202
	F_errdetail_plural(m, int32(549686), int32(549815), v201, v11)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L19
	} else {
		goto L60
	}
L60:
	;
	F_errhint(m, int32(552659), int32(0))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L19
	} else {
		goto L61
	}
L61:
	;
	F_errfinish(m, int32(471205), int32(2193), int32(262876))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L19
	} else {
		goto L62
	}
L62:
	;
	goto L38
L63:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v128)+8))
	if v227 != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	F_pfree(m, v227)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L19
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v128)+12))
	if v230 != 0 {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	goto L66
L68:
	;
	F_tuplestore_end(m, v230)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L19
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v128)+16))
	if v233 != 0 {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	goto L70
L72:
	;
	F_FreeTupleDesc(m, v233)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L19
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	F_pfree(m, v128)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L19
	} else {
		goto L76
	}
L75:
	;
	goto L74
L76:
	;
	goto L2
L77:
	;
	F_errcode(m, int32(100663808))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L19
	} else {
		goto L78
	}
L78:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v128)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v256
	F_errmsg(m, int32(191399), v11+int32(32))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L19
	} else {
		goto L79
	}
L79:
	;
	F_errfinish(m, int32(471205), int32(2153), int32(262876))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L19
	} else {
		goto L80
	}
L80:
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
	v8 = *(*int32)(unsafe.Add(mBase, _consts[585]))
	if v8 == int32(0) {
		m.G0 = v5 + int32(16)
		return
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, _consts[584]))
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
					F_errmsg(m, int32(440644), int32(0))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return
					} else {
						v33 = *(*int32)(unsafe.Add(mBase, _consts[1138]))
						*(*int32)(unsafe.Add(mBase, uint32(v5))) = v33
						F_errhint(m, int32(589361), v5)
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return
						} else {
							F_errfinish(m, int32(474468), int32(104), int32(304824))
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
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v139 int32
	_ = v139
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v325 int32
	_ = v325
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v360 int32
	_ = v360
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v423 int32
	_ = v423
	v4 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(48)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	if v15 == v4 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v12 + int32(48)
	return v423
L2:
	;
	v423 = int32(1)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v19 = F_pstrdup(m, v14)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	F_pfree(m, v19)
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L6
	} else {
		goto L106
	}
L6:
	;
	return int32(0)
L7:
	;
	v26 = F_SplitIdentifierString(m, v19, int32(44), v12+int32(32))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	if v26 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _consts[40]))
	*(*int32)(unsafe.Add(mBase, _consts[497])) = v31
	goto L12
L10:
	;
	goto L11
L11:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
	if v40 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v37 = F_format_elog_string(m, int32(605174), int32(0))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L6
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, _consts[498])) = v37
	v409 = v4
	goto L5
L14:
	;
	v409 = int32(1)
	goto L5
L15:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	if int32(0) < v43 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v47 = int32(0)
	goto L19
L17:
	;
	v114 = v40
	goto L18
L18:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v114)+4))
	if v120 <= int32(0) {
		goto L34
	} else {
		goto L35
	}
L19:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v56+v47<<(uint(int32(2))%32))))
	v61 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v61
	v71 = F_ReplicationSlotValidateNameInternal(m, v60, v12+int32(44), v12+int32(40), v12+int32(36))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L6
	} else {
		goto L21
	}
L20:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
	if v108 == int32(0) {
		goto L14
	} else {
		goto L32
	}
L21:
	;
	if v71 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v12)+44))
	*(*int32)(unsafe.Add(mBase, _consts[518])) = v75
	goto L25
L23:
	;
	goto L24
L24:
	;
	v105 = v47 + int32(1)
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	if v105 < v106 {
		v47 = v105
		goto L19
	} else {
		goto L31
	}
L25:
	;
	v79 = *(*int32)(unsafe.Add(mBase, _consts[40]))
	*(*int32)(unsafe.Add(mBase, _consts[497])) = v79
	goto L26
L26:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v12)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v82
	v88 = F_format_elog_string(m, int32(195784), v12+int32(16))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L6
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, _consts[498])) = v88
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v12)+36))
	if v91 == int32(0) {
		v409 = v4
		goto L5
	} else {
		goto L28
	}
L28:
	;
	v95 = *(*int32)(unsafe.Add(mBase, _consts[40]))
	*(*int32)(unsafe.Add(mBase, _consts[497])) = v95
	goto L29
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v91
	v101 = F_format_elog_string(m, int32(195784), v12)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L6
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, _consts[702])) = v101
	v409 = v4
	goto L5
L31:
	;
	goto L20
L32:
	;
	v114 = v108
	goto L18
L33:
	;
	v212 = F_guc_malloc(m, v206)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L6
	} else {
		goto L57
	}
L34:
	;
	v206 = int32(4)
	goto L33
L35:
	;
	goto L36
L36:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v114)+12))
	v127 = int32(0)
	v130 = int32(4)
	goto L37
L37:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v124+v127<<(uint(int32(2))%32))))
	if v139&int32(3) == int32(0) {
		v163 = v139
		goto L41
	} else {
		goto L42
	}
L38:
	;
	v206 = v199
	goto L33
L39:
	;
	v198 = int32(1)
	v199 = v196 + v130 + v198
	v201 = v127 + v198
	if v201 != v120 {
		v127 = v201
		v130 = v199
		goto L37
	} else {
		goto L56
	}
L40:
	;
	v196 = v188 - v139
	goto L39
L41:
	;
	v167 = v163
	goto L50
L42:
	;
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v139))))
	if v147 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v196 = int32(0)
	goto L39
L44:
	;
	goto L45
L45:
	;
	v152 = v139
	goto L46
L46:
	;
	v156 = v152 + int32(1)
	if v156&int32(3) == int32(0) {
		v163 = v156
		goto L41
	} else {
		goto L48
	}
L47:
	;
	v188 = v156
	goto L40
L48:
	;
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156))))
	if v161 != 0 {
		v152 = v156
		goto L46
	} else {
		goto L49
	}
L49:
	;
	goto L47
L50:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	v176 = int32(-2139062144)
	if (int32(16843008)-v173|v173)&v176 == v176 {
		v167 = v167 + int32(4)
		goto L50
	} else {
		goto L52
	}
L51:
	;
	v182 = v167
	goto L53
L52:
	;
	goto L51
L53:
	;
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182))))
	if v186 != 0 {
		v182 = v182 + int32(1)
		goto L53
	} else {
		goto L55
	}
L54:
	;
	v188 = v182
	goto L40
L55:
	;
	goto L54
L56:
	;
	goto L38
L57:
	;
	if v212 == int32(0) {
		v423 = v4
		goto L1
	} else {
		goto L58
	}
L58:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
	if v216 != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v216)+4))
	v219 = v217
	goto L61
L60:
	;
	v219 = int32(0)
	goto L61
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v212))) = v219
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
	if v221 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v212
	goto L14
L63:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v221)+4))
	if v224 <= int32(0) {
		goto L62
	} else {
		goto L64
	}
L64:
	;
	v230 = int32(0)
	v234 = v212 + int32(4)
	goto L65
L65:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v221)+12))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v239+v230<<(uint(int32(2))%32))))
	if (v243^v234)&int32(3) != 0 {
		goto L70
	} else {
		goto L71
	}
L66:
	;
	goto L62
L67:
	;
	if v243&int32(3) == int32(0) {
		v341 = v243
		goto L90
	} else {
		goto L91
	}
L68:
	;
	goto L67
L69:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v298))) = uint8(v297)
	if v297&int32(255) == int32(0) {
		goto L68
	} else {
		goto L84
	}
L70:
	;
	v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243))))
	v296 = v243
	v297 = v249
	v298 = v234
	goto L69
L71:
	;
	goto L72
L72:
	;
	if v243&int32(3) != 0 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v253 = v243
	v255 = v234
	goto L76
L74:
	;
	v267 = v243
	v269 = v234
	goto L75
L75:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v267)))
	v274 = int32(-2139062144)
	if (int32(16843008)-v271|v271)&v274 != v274 {
		v296 = v267
		v297 = v271
		v298 = v269
		goto L69
	} else {
		goto L80
	}
L76:
	;
	v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v253))))
	*(*uint8)(unsafe.Add(mBase, uint32(v255))) = uint8(v256)
	if v256 == int32(0) {
		goto L68
	} else {
		goto L78
	}
L77:
	;
	v267 = v263
	v269 = v261
	goto L75
L78:
	;
	v260 = int32(1)
	v261 = v255 + v260
	v263 = v253 + v260
	if v263&int32(3) != 0 {
		v253 = v263
		v255 = v261
		goto L76
	} else {
		goto L79
	}
L79:
	;
	goto L77
L80:
	;
	v279 = v267
	v280 = v271
	v281 = v269
	goto L81
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v281))) = v280
	v283 = int32(4)
	v284 = v281 + v283
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v279)+4))
	v287 = v279 + v283
	v291 = int32(-2139062144)
	if (v285|(int32(16843008)-v285))&v291 == v291 {
		v279 = v287
		v280 = v285
		v281 = v284
		goto L81
	} else {
		goto L83
	}
L82:
	;
	v296 = v287
	v297 = v285
	v298 = v284
	goto L69
L83:
	;
	goto L82
L84:
	;
	v305 = v296
	v307 = v298
	goto L85
L85:
	;
	v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v305)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v307)+1)) = uint8(v308)
	v310 = int32(1)
	if v308 != 0 {
		v305 = v305 + v310
		v307 = v307 + v310
		goto L85
	} else {
		goto L87
	}
L86:
	;
	goto L68
L87:
	;
	goto L86
L88:
	;
	v376 = int32(1)
	v379 = v230 + v376
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v221)+4))
	if v379 < v380 {
		v230 = v379
		v234 = v234 + v374 + v376
		goto L65
	} else {
		goto L105
	}
L89:
	;
	v374 = v366 - v243
	goto L88
L90:
	;
	v345 = v341
	goto L99
L91:
	;
	v325 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243))))
	if v325 == int32(0) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v374 = int32(0)
	goto L88
L93:
	;
	goto L94
L94:
	;
	v330 = v243
	goto L95
L95:
	;
	v334 = v330 + int32(1)
	if v334&int32(3) == int32(0) {
		v341 = v334
		goto L90
	} else {
		goto L97
	}
L96:
	;
	v366 = v334
	goto L89
L97:
	;
	v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v334))))
	if v339 != 0 {
		v330 = v334
		goto L95
	} else {
		goto L98
	}
L98:
	;
	goto L96
L99:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v345)))
	v354 = int32(-2139062144)
	if (int32(16843008)-v351|v351)&v354 == v354 {
		v345 = v345 + int32(4)
		goto L99
	} else {
		goto L101
	}
L100:
	;
	v360 = v345
	goto L102
L101:
	;
	goto L100
L102:
	;
	v364 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v360))))
	if v364 != 0 {
		v360 = v360 + int32(1)
		goto L102
	} else {
		goto L104
	}
L103:
	;
	v366 = v360
	goto L89
L104:
	;
	goto L103
L105:
	;
	goto L66
L106:
	;
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
	F_list_free(m, v413)
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L6
	} else {
		goto L107
	}
L107:
	;
	v423 = v409
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
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int64
	_ = v78
	var v90 int32
	_ = v90
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v172 int32
	_ = v172
	var v179 int32
	_ = v179
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v225 int32
	_ = v225
	var v244 int32
	_ = v244
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v309 int32
	_ = v309
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v477 int32
	_ = v477
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v500 int32
	_ = v500
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v523 int32
	_ = v523
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v533 int32
	_ = v533
	var v538 int32
	_ = v538
	var v543 int32
	_ = v543
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v560 int32
	_ = v560
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v595 int32
	_ = v595
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v629 int32
	_ = v629
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v638 int32
	_ = v638
	var v656 int32
	_ = v656
	var v659 int32
	_ = v659
	var v668 int32
	_ = v668
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v678 int32
	_ = v678
	var v685 int32
	_ = v685
	var v687 int32
	_ = v687
	var v689 int32
	_ = v689
	var v710 int32
	_ = v710
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v738 int32
	_ = v738
	var v744 int32
	_ = v744
	var v766 int32
	_ = v766
	var v772 int32
	_ = v772
	var v792 int32
	_ = v792
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v815 int32
	_ = v815
	var v820 int32
	_ = v820
	var v824 int32
	_ = v824
	var v839 int32
	_ = v839
	var v841 int32
	_ = v841
	var v848 int32
	_ = v848
	var v853 int32
	_ = v853
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v878 int32
	_ = v878
	var v880 int32
	_ = v880
	var v883 int32
	_ = v883
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v890 int32
	_ = v890
	var v891 int32
	_ = v891
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v895 int32
	_ = v895
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v916 int32
	_ = v916
	var v917 int32
	_ = v917
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v925 int32
	_ = v925
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v931 int32
	_ = v931
	var v937 int32
	_ = v937
	var v939 int32
	_ = v939
	var v940 int32
	_ = v940
	var v942 int32
	_ = v942
	var v943 int32
	_ = v943
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v948 int32
	_ = v948
	var v951 int32
	_ = v951
	var v953 int32
	_ = v953
	var v954 int32
	_ = v954
	var v956 int32
	_ = v956
	var v957 int32
	_ = v957
	var v958 int32
	_ = v958
	var v961 int32
	_ = v961
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v969 int32
	_ = v969
	var v970 int32
	_ = v970
	var v971 int32
	_ = v971
	var v972 int32
	_ = v972
	var v973 int32
	_ = v973
	var v974 int32
	_ = v974
	var v975 int32
	_ = v975
	var v976 int32
	_ = v976
	var v978 int32
	_ = v978
	var v982 int32
	_ = v982
	var v984 int32
	_ = v984
	var v991 int32
	_ = v991
	var v994 int32
	_ = v994
	var v1012 int32
	_ = v1012
	var v1015 int32
	_ = v1015
	var v1017 int32
	_ = v1017
	var v1027 int32
	_ = v1027
	var v1028 int32
	_ = v1028
	var v1032 int32
	_ = v1032
	var v1033 int32
	_ = v1033
	var v1038 int32
	_ = v1038
	var v1045 int32
	_ = v1045
	var v1047 int32
	_ = v1047
	var v1049 int32
	_ = v1049
	var v1070 int32
	_ = v1070
	var v1082 int32
	_ = v1082
	var v1083 int32
	_ = v1083
	var v1087 int32
	_ = v1087
	var v1088 int32
	_ = v1088
	var v1092 int32
	_ = v1092
	var v1093 int32
	_ = v1093
	var v1098 int32
	_ = v1098
	var v1104 int32
	_ = v1104
	var v1126 int32
	_ = v1126
	var v1130 int32
	_ = v1130
	var v1132 int32
	_ = v1132
	var v1139 int32
	_ = v1139
	var v1168 int32
	_ = v1168
	var v1172 int32
	_ = v1172
	var v1174 int32
	_ = v1174
	var v1179 int32
	_ = v1179
	var v1186 int32
	_ = v1186
	var v1207 int32
	_ = v1207
	var v1211 int32
	_ = v1211
	var v1212 int32
	_ = v1212
	var v1217 int32
	_ = v1217
	var v1219 int32
	_ = v1219
	var v1224 int32
	_ = v1224
	var v1230 int32
	_ = v1230
	var v1254 int32
	_ = v1254
	var v1263 int32
	_ = v1263
	var v1264 int32
	_ = v1264
	var v1270 int32
	_ = v1270
	var v1274 int32
	_ = v1274
	var v1275 int32
	_ = v1275
	var v1283 int32
	_ = v1283
	var v1284 int32
	_ = v1284
	var v1285 int32
	_ = v1285
	var v1298 int32
	_ = v1298
	var v1299 int32
	_ = v1299
	var v1301 int32
	_ = v1301
	var v1320 int32
	_ = v1320
	var v1322 int32
	_ = v1322
	var v1324 int32
	_ = v1324
	var v1326 int32
	_ = v1326
	var v1328 int32
	_ = v1328
	var v1330 int32
	_ = v1330
	var v1332 int32
	_ = v1332
	var v1334 int32
	_ = v1334
	var v1336 int32
	_ = v1336
	var v1337 int32
	_ = v1337
	var v1339 int32
	_ = v1339
	var v1341 int32
	_ = v1341
	var v1347 int32
	_ = v1347
	var v1348 int32
	_ = v1348
	var v1374 int32
	_ = v1374
	var v1375 int32
	_ = v1375
	var v1377 int32
	_ = v1377
	var v1396 int32
	_ = v1396
	var v1398 int32
	_ = v1398
	var v1403 int32
	_ = v1403
	var v1431 int32
	_ = v1431
	var v1432 int32
	_ = v1432
	var v1436 int32
	_ = v1436
	var v1442 int32
	_ = v1442
	var v1444 int32
	_ = v1444
	var v1446 int32
	_ = v1446
	var v1463 int32
	_ = v1463
	var v1466 int32
	_ = v1466
	var v1472 int32
	_ = v1472
	var v1473 int32
	_ = v1473
	var v1474 int32
	_ = v1474
	var v1477 int32
	_ = v1477
	var v1482 int32
	_ = v1482
	var v1483 int32
	_ = v1483
	var v1487 int32
	_ = v1487
	var v1504 int32
	_ = v1504
	var v1509 int32
	_ = v1509
	var v1515 int32
	_ = v1515
	var v1518 int32
	_ = v1518
	var v1519 int32
	_ = v1519
	var v1520 int32
	_ = v1520
	var v1521 int32
	_ = v1521
	var v1522 int32
	_ = v1522
	var v1523 int32
	_ = v1523
	var v1528 int32
	_ = v1528
	var v1530 int32
	_ = v1530
	var v1531 int32
	_ = v1531
	var v1535 int32
	_ = v1535
	var v1536 int32
	_ = v1536
	var v1537 int32
	_ = v1537
	var v1545 int32
	_ = v1545
	var v1547 int32
	_ = v1547
	var v1566 int32
	_ = v1566
	var v1569 int32
	_ = v1569
	var v1571 int32
	_ = v1571
	var v1576 int32
	_ = v1576
	var v1598 int32
	_ = v1598
	var v1599 int32
	_ = v1599
	var v1600 int32
	_ = v1600
	var v1602 int32
	_ = v1602
	var v1603 int32
	_ = v1603
	var v1604 int32
	_ = v1604
	var v1608 int32
	_ = v1608
	var v1609 int32
	_ = v1609
	var v1614 int32
	_ = v1614
	var v1616 int32
	_ = v1616
	var v1617 int32
	_ = v1617
	var v1618 int32
	_ = v1618
	var v1621 int32
	_ = v1621
	var v1622 int32
	_ = v1622
	var v1623 int32
	_ = v1623
	var v1632 int32
	_ = v1632
	var v1633 int32
	_ = v1633
	var v1634 int32
	_ = v1634
	var v1638 int32
	_ = v1638
	var v1639 int32
	_ = v1639
	var v1642 int32
	_ = v1642
	var v1646 int32
	_ = v1646
	var v1651 int32
	_ = v1651
	var v1652 int32
	_ = v1652
	var v1656 int32
	_ = v1656
	var v1657 int32
	_ = v1657
	var v1660 int32
	_ = v1660
	var v1661 int32
	_ = v1661
	var v1662 int32
	_ = v1662
	var v1667 int32
	_ = v1667
	var v1690 int32
	_ = v1690
	var v1691 int32
	_ = v1691
	var v1694 int32
	_ = v1694
	var v1695 int32
	_ = v1695
	var v1696 int32
	_ = v1696
	var v1705 int32
	_ = v1705
	var v1709 int32
	_ = v1709
	var v1711 int32
	_ = v1711
	var v1712 int32
	_ = v1712
	var v1715 int32
	_ = v1715
	var v1719 int32
	_ = v1719
	var v1723 int32
	_ = v1723
	var v1724 int32
	_ = v1724
	var v1728 int32
	_ = v1728
	var v1729 int32
	_ = v1729
	var v1732 int32
	_ = v1732
	var v1733 int32
	_ = v1733
	var v1734 int32
	_ = v1734
	var v1737 int32
	_ = v1737
	var v1744 int32
	_ = v1744
	var v1750 int32
	_ = v1750
	var v1753 int32
	_ = v1753
	var v1758 int32
	_ = v1758
	var v1780 int32
	_ = v1780
	var v1781 int32
	_ = v1781
	var v1789 int32
	_ = v1789
	var v1811 int32
	_ = v1811
	var v1812 int32
	_ = v1812
	var v1815 int32
	_ = v1815
	var v1818 int32
	_ = v1818
	var v1819 int32
	_ = v1819
	var v1820 int32
	_ = v1820
	var v1824 int32
	_ = v1824
	var v1826 int32
	_ = v1826
	var v1827 int32
	_ = v1827
	var v1828 int32
	_ = v1828
	var v1829 int32
	_ = v1829
	var v1830 int32
	_ = v1830
	var v1831 int32
	_ = v1831
	var v1836 int32
	_ = v1836
	var v1838 int32
	_ = v1838
	var v1840 int32
	_ = v1840
	var v1842 int32
	_ = v1842
	var v1843 int32
	_ = v1843
	var v1847 int32
	_ = v1847
	var v1851 int32
	_ = v1851
	var v1853 int32
	_ = v1853
	var v1856 int32
	_ = v1856
	var v1863 int32
	_ = v1863
	var v1865 int32
	_ = v1865
	var v1868 int32
	_ = v1868
	var v1871 int32
	_ = v1871
	var v1872 int32
	_ = v1872
	var v1874 int32
	_ = v1874
	var v1876 int32
	_ = v1876
	var v1877 int32
	_ = v1877
	var v1880 int32
	_ = v1880
	var v1881 int32
	_ = v1881
	var v1883 int32
	_ = v1883
	var v1885 int32
	_ = v1885
	var v1888 int32
	_ = v1888
	var v1892 int32
	_ = v1892
	var v1893 int32
	_ = v1893
	var v1897 int32
	_ = v1897
	var v1905 int32
	_ = v1905
	var v1906 int32
	_ = v1906
	var v1909 int32
	_ = v1909
	var v1910 int32
	_ = v1910
	var v1911 int32
	_ = v1911
	var v1913 int32
	_ = v1913
	var v1914 int32
	_ = v1914
	var v1917 int32
	_ = v1917
	var v1918 int32
	_ = v1918
	var v1920 int32
	_ = v1920
	var v1923 int32
	_ = v1923
	var v1924 int32
	_ = v1924
	var v1925 int32
	_ = v1925
	var v1928 int32
	_ = v1928
	var v1937 int32
	_ = v1937
	var v1938 int32
	_ = v1938
	var v1939 int32
	_ = v1939
	var v1940 int32
	_ = v1940
	var v1941 int32
	_ = v1941
	var v1945 int32
	_ = v1945
	var v1946 int32
	_ = v1946
	var v1950 int32
	_ = v1950
	var v1951 int32
	_ = v1951
	var v1952 int32
	_ = v1952
	var v1953 int32
	_ = v1953
	var v1954 int32
	_ = v1954
	var v1955 int32
	_ = v1955
	var v1956 int32
	_ = v1956
	var v1957 int32
	_ = v1957
	var v1958 int32
	_ = v1958
	var v1959 int32
	_ = v1959
	var v1962 int32
	_ = v1962
	var v1964 int32
	_ = v1964
	var v1965 int32
	_ = v1965
	var v1969 int32
	_ = v1969
	var v1970 int32
	_ = v1970
	var v1976 int32
	_ = v1976
	var v1977 int32
	_ = v1977
	var v1979 int32
	_ = v1979
	var v1980 int32
	_ = v1980
	var v1981 int32
	_ = v1981
	var v1982 int32
	_ = v1982
	var v1984 int32
	_ = v1984
	var v1987 int32
	_ = v1987
	var v1988 int32
	_ = v1988
	var v1990 int32
	_ = v1990
	var v1992 int32
	_ = v1992
	var v1993 int32
	_ = v1993
	var v1997 int32
	_ = v1997
	var v1998 int32
	_ = v1998
	var v1999 int32
	_ = v1999
	var v2003 int32
	_ = v2003
	var v2004 int32
	_ = v2004
	var v2007 int32
	_ = v2007
	var v2008 int32
	_ = v2008
	var v2009 int32
	_ = v2009
	var v2018 int32
	_ = v2018
	var v2020 int32
	_ = v2020
	var v2027 int32
	_ = v2027
	var v2036 int32
	_ = v2036
	var v2048 int32
	_ = v2048
	var v2051 int32
	_ = v2051
	var v2053 int32
	_ = v2053
	var v2063 int32
	_ = v2063
	var v2064 int32
	_ = v2064
	var v2068 int32
	_ = v2068
	var v2069 int32
	_ = v2069
	var v2074 int32
	_ = v2074
	var v2081 int32
	_ = v2081
	var v2083 int32
	_ = v2083
	var v2085 int32
	_ = v2085
	var v2106 int32
	_ = v2106
	var v2118 int32
	_ = v2118
	var v2119 int32
	_ = v2119
	var v2123 int32
	_ = v2123
	var v2124 int32
	_ = v2124
	var v2128 int32
	_ = v2128
	var v2129 int32
	_ = v2129
	var v2134 int32
	_ = v2134
	var v2140 int32
	_ = v2140
	var v2162 int32
	_ = v2162
	var v2166 int32
	_ = v2166
	var v2168 int32
	_ = v2168
	var v2170 int32
	_ = v2170
	var v2171 int32
	_ = v2171
	var v2172 int32
	_ = v2172
	var v2177 int32
	_ = v2177
	var v2179 int32
	_ = v2179
	var v2189 int32
	_ = v2189
	var v2199 int32
	_ = v2199
	var v2201 int32
	_ = v2201
	var v2208 int32
	_ = v2208
	var v2211 int32
	_ = v2211
	var v2229 int32
	_ = v2229
	var v2232 int32
	_ = v2232
	var v2234 int32
	_ = v2234
	var v2244 int32
	_ = v2244
	var v2245 int32
	_ = v2245
	var v2249 int32
	_ = v2249
	var v2250 int32
	_ = v2250
	var v2255 int32
	_ = v2255
	var v2262 int32
	_ = v2262
	var v2264 int32
	_ = v2264
	var v2266 int32
	_ = v2266
	var v2287 int32
	_ = v2287
	var v2299 int32
	_ = v2299
	var v2300 int32
	_ = v2300
	var v2304 int32
	_ = v2304
	var v2305 int32
	_ = v2305
	var v2309 int32
	_ = v2309
	var v2310 int32
	_ = v2310
	var v2315 int32
	_ = v2315
	var v2321 int32
	_ = v2321
	var v2343 int32
	_ = v2343
	var v2347 int32
	_ = v2347
	var v2349 int32
	_ = v2349
	var v2353 int32
	_ = v2353
	var v2358 int32
	_ = v2358
	var v2365 int32
	_ = v2365
	var v2388 int32
	_ = v2388
	var v2392 int32
	_ = v2392
	var v2395 int32
	_ = v2395
	var v2396 int32
	_ = v2396
	var v2397 int32
	_ = v2397
	var v2398 int32
	_ = v2398
	var v2400 int32
	_ = v2400
	var v2406 int32
	_ = v2406
	var v2407 int32
	_ = v2407
	var v2408 int32
	_ = v2408
	var v2410 int32
	_ = v2410
	var v2412 int32
	_ = v2412
	var v2413 int32
	_ = v2413
	var v2415 int32
	_ = v2415
	var v2416 int32
	_ = v2416
	var v2420 int32
	_ = v2420
	var v2423 int32
	_ = v2423
	var v2430 int32
	_ = v2430
	var v2431 int32
	_ = v2431
	var v2432 int32
	_ = v2432
	var v2433 int32
	_ = v2433
	var v2438 int32
	_ = v2438
	var v2443 int32
	_ = v2443
	var v2444 int32
	_ = v2444
	var v2448 int32
	_ = v2448
	var v2451 int32
	_ = v2451
	var v2453 int32
	_ = v2453
	var v2454 int32
	_ = v2454
	var v2455 int32
	_ = v2455
	var v2456 int32
	_ = v2456
	var v2458 int32
	_ = v2458
	var v2464 int32
	_ = v2464
	var v2465 int32
	_ = v2465
	var v2466 int32
	_ = v2466
	var v2468 int32
	_ = v2468
	var v2470 int32
	_ = v2470
	var v2471 int32
	_ = v2471
	var v2473 int32
	_ = v2473
	var v2474 int32
	_ = v2474
	var v2478 int32
	_ = v2478
	var v2484 int32
	_ = v2484
	var v2487 int32
	_ = v2487
	var v2489 int32
	_ = v2489
	var v2492 int32
	_ = v2492
	var v2495 int32
	_ = v2495
	var v2498 int32
	_ = v2498
	var v2500 int32
	_ = v2500
	var v2501 int32
	_ = v2501
	var v2504 int32
	_ = v2504
	var v2507 int32
	_ = v2507
	var v2509 int32
	_ = v2509
	var v2510 int32
	_ = v2510
	var v2513 int32
	_ = v2513
	var v2516 int32
	_ = v2516
	var v2518 int32
	_ = v2518
	var v2519 int32
	_ = v2519
	var v2522 int32
	_ = v2522
	var v2525 int32
	_ = v2525
	var v2527 int32
	_ = v2527
	var v2528 int32
	_ = v2528
	var v2529 int32
	_ = v2529
	var v2535 int32
	_ = v2535
	var v2536 int32
	_ = v2536
	var v2539 int32
	_ = v2539
	var v2540 int32
	_ = v2540
	var v2541 int32
	_ = v2541
	var v2542 int32
	_ = v2542
	var v2543 int32
	_ = v2543
	var v2544 int32
	_ = v2544
	var v2549 int32
	_ = v2549
	var v2553 int32
	_ = v2553
	var v2554 int32
	_ = v2554
	var v2557 int32
	_ = v2557
	var v2560 int32
	_ = v2560
	var v2561 int32
	_ = v2561
	var v2566 int32
	_ = v2566
	var v2570 int32
	_ = v2570
	var v2573 int32
	_ = v2573
	var v2574 int32
	_ = v2574
	var v2575 int32
	_ = v2575
	var v2576 int32
	_ = v2576
	var v2577 int32
	_ = v2577
	var v2579 int32
	_ = v2579
	var v2585 int32
	_ = v2585
	var v2587 int32
	_ = v2587
	var v2589 int32
	_ = v2589
	var v2591 int32
	_ = v2591
	var v2594 int32
	_ = v2594
	var v2595 int32
	_ = v2595
	var v2599 int32
	_ = v2599
	var v2600 int32
	_ = v2600
	var v2605 int32
	_ = v2605
	var v2619 int32
	_ = v2619
	var v2642 int32
	_ = v2642
	var v2644 int32
	_ = v2644
	var v2645 int32
	_ = v2645
	var v2646 int32
	_ = v2646
	var v2659 int32
	_ = v2659
	var v2661 int32
	_ = v2661
	var v2674 int32
	_ = v2674
	var v2678 int32
	_ = v2678
	var v2680 int32
	_ = v2680
	var v2683 int32
	_ = v2683
	var v2688 int32
	_ = v2688
	var v2690 int32
	_ = v2690
	var v2694 int32
	_ = v2694
	var v2699 int32
	_ = v2699
	var v2700 int32
	_ = v2700
	var v2704 int32
	_ = v2704
	var v2705 int32
	_ = v2705
	var v2706 int32
	_ = v2706
	var v2709 int32
	_ = v2709
	var v2713 int32
	_ = v2713
	var v2716 int32
	_ = v2716
	var v2717 int32
	_ = v2717
	var v2723 int32
	_ = v2723
	var v2724 int32
	_ = v2724
	var v2728 int32
	_ = v2728
	var v2729 int32
	_ = v2729
	var v2730 int32
	_ = v2730
	var v2733 int32
	_ = v2733
	var v2734 int32
	_ = v2734
	var v2735 int32
	_ = v2735
	var v2736 int32
	_ = v2736
	var v2742 int32
	_ = v2742
	var v2769 int32
	_ = v2769
	var v2793 int32
	_ = v2793
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
	m.G0 = v2769 + int32(48)
	return v2793
L2:
	;
	v2769 = v2742
	v2793 = int32(1)
	goto L1
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
	v2742 = v29
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
	v74 = F__emscripten_memset_bulkmem(m, v50, base.I32_extend8_s(int32(0)), int32(96))
	mBase = m.M
	goto L17
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
	*(*int32)(unsafe.Add(mBase, _consts[40])) = v56
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
	F_errmsg_internal(m, int32(280896), int32(0))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L6
	} else {
		goto L15
	}
L15:
	;
	F_errfinish(m, int32(298966), int32(179), int32(93741))
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
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	v76 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v75)+60)) = v76
	v78 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v75)+52)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v75)+44)) = v76
	*(*int64)(unsafe.Add(mBase, uint32(v75)+36)) = v78
	*(*int64)(unsafe.Add(mBase, uint32(v75)+4)) = v78
	*(*int64)(unsafe.Add(mBase, uint32(v75)+12)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v75)+20)) = v76
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	*(*int32)(unsafe.Add(mBase, uint32(v90))) = v44
	if v31&int32(3) == v76 {
		v115 = v31
		goto L20
	} else {
		goto L21
	}
L18:
	;
	if base.Ui32(v148) < base.Ui32(int32(-2)) {
		goto L38
	} else {
		goto L39
	}
L19:
	;
	v148 = v140 - v31
	goto L18
L20:
	;
	v119 = v115
	goto L29
L21:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
	if v99 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v148 = int32(0)
	goto L18
L23:
	;
	goto L24
L24:
	;
	v104 = v31
	goto L25
L25:
	;
	v108 = v104 + int32(1)
	if v108&int32(3) == int32(0) {
		v115 = v108
		goto L20
	} else {
		goto L27
	}
L26:
	;
	v140 = v108
	goto L19
L27:
	;
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108))))
	if v113 != 0 {
		v104 = v108
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
	v128 = int32(-2139062144)
	if (int32(16843008)-v125|v125)&v128 == v128 {
		v119 = v119 + int32(4)
		goto L29
	} else {
		goto L31
	}
L30:
	;
	v134 = v119
	goto L32
L31:
	;
	goto L30
L32:
	;
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134))))
	if v138 != 0 {
		v134 = v134 + int32(1)
		goto L32
	} else {
		goto L34
	}
L33:
	;
	v140 = v134
	goto L19
L34:
	;
	goto L33
L35:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v29)+44))
	v419 = m.G0
	v421 = v419 - int32(1232)
	m.G0 = v421
	v425 = v421 + int32(816)
	v427 = v421 + int32(16)
	v429 = v425
	v430 = l1
	v431 = v29
	v432 = v418
	v434 = int32(-2)
	v440 = v427
	v442 = v425
	v444 = v427
	v446 = v421
	v447 = v4
	v449 = v29 + int32(36)
	v450 = int32(200)
	v454 = v29 + int32(40)
	goto L75
L36:
	;
	F_yy_fatal_error_4(m, int32(641595))
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L6
	} else {
		goto L72
	}
L37:
	;
	F_yy_fatal_error_4(m, int32(641554))
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L6
	} else {
		goto L71
	}
L38:
	;
	v153 = v148 + int32(2)
	v154 = F_palloc(m, v153)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L6
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	F_yy_fatal_error_4(m, int32(641625))
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L6
	} else {
		goto L70
	}
L41:
	;
	if v154 == int32(0) {
		goto L37
	} else {
		goto L42
	}
L42:
	;
	if v148 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v309 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v154+v148))) = uint16(v309)
	if base.Ui32(v153) < base.Ui32(int32(2)) {
		v397 = v309
		goto L57
	} else {
		goto L58
	}
L44:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v148) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v172 = v4
	v179 = int32(0)
	goto L48
L46:
	;
	v225 = v4
	goto L47
L47:
	;
	v244 = v148 & int32(3)
	if v244 == int32(0) {
		goto L43
	} else {
		goto L51
	}
L48:
	;
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31+v172))))
	*(*uint8)(unsafe.Add(mBase, uint32(v154+v172))) = uint8(v192)
	v195 = v172 | int32(1)
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31+v195))))
	*(*uint8)(unsafe.Add(mBase, uint32(v154+v195))) = uint8(v198)
	v201 = v172 | int32(2)
	v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31+v201))))
	*(*uint8)(unsafe.Add(mBase, uint32(v154+v201))) = uint8(v204)
	v207 = v172 | int32(3)
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31+v207))))
	*(*uint8)(unsafe.Add(mBase, uint32(v154+v207))) = uint8(v210)
	v212 = int32(4)
	v213 = v172 + v212
	v215 = v179 + v212
	if v215 != v148&int32(-4) {
		v172 = v213
		v179 = v215
		goto L48
	} else {
		goto L50
	}
L49:
	;
	v225 = v213
	goto L47
L50:
	;
	goto L49
L51:
	;
	v250 = v4
	v255 = v225
	goto L52
L52:
	;
	v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31+v255))))
	*(*uint8)(unsafe.Add(mBase, uint32(v154+v255))) = uint8(v275)
	v277 = int32(1)
	v280 = v250 + v277
	if v280 != v244 {
		v250 = v280
		v255 = v255 + v277
		goto L52
	} else {
		goto L54
	}
L53:
	;
	goto L43
L54:
	;
	goto L53
L55:
	;
	if v397 == int32(0) {
		goto L36
	} else {
		goto L69
	}
L56:
	;
	F_yy_fatal_error_4(m, int32(641955))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L6
	} else {
		goto L68
	}
L57:
	;
	goto L55
L58:
	;
	v315 = v153 - int32(2)
	v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154+v315))))
	if v317 != 0 {
		v397 = v309
		goto L57
	} else {
		goto L59
	}
L59:
	;
	v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153+v154-int32(1)))))
	if v321 != 0 {
		v397 = v309
		goto L57
	} else {
		goto L60
	}
L60:
	;
	v323 = F_palloc(m, int32(48))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L6
	} else {
		goto L61
	}
L61:
	;
	if v323 == int32(0) {
		goto L56
	} else {
		goto L62
	}
L62:
	;
	v327 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v323)+20)) = v327
	*(*int32)(unsafe.Add(mBase, uint32(v323)+8)) = v154
	*(*int32)(unsafe.Add(mBase, uint32(v323)+4)) = v154
	*(*int32)(unsafe.Add(mBase, uint32(v323)+12)) = v315
	*(*int64)(unsafe.Add(mBase, uint32(v323)+40)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v323)+24)) = int64(4294967296)
	*(*int32)(unsafe.Add(mBase, uint32(v323)+16)) = v315
	*(*int32)(unsafe.Add(mBase, uint32(v323))) = v327
	F_syncrep_yyensure_buffer_stack(m, v90)
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L6
	} else {
		goto L63
	}
L63:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v90)+20))
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v90)+12))
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v341+v342<<(uint(int32(2))%32))))
	if v346 == v323 {
		v397 = v323
		goto L57
	} else {
		goto L64
	}
L64:
	;
	if v346 != 0 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v90)+36))
	v349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v348))) = uint8(v349)
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v90)+20))
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v90)+12))
	v353 = int32(2)
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v351+v352<<(uint(v353)%32))))
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v90)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v356)+8)) = v357
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v90)+20))
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v90)+12))
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v359+v360<<(uint(v353)%32))))
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v90)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v364)+16)) = v365
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v90)+20))
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v90)+12))
	v369 = v367
	v370 = v368
	goto L67
L66:
	;
	v369 = v341
	v370 = v342
	goto L67
L67:
	;
	v371 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v370<<(uint(v371)%32)+v369))) = v323
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v90)+20))
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v90)+12))
	v379 = v375 + v376<<(uint(v371)%32)
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v379)))
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v380)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v90)+28)) = v381
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v379)))
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v383)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v90)+36)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v90)+80)) = v384
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v379)))
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v387)))
	*(*int32)(unsafe.Add(mBase, uint32(v90)+4)) = v388
	v390 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v384))))
	*(*uint8)(unsafe.Add(mBase, uint32(v90)+24)) = uint8(v390)
	*(*int32)(unsafe.Add(mBase, uint32(v90)+48)) = int32(1)
	v397 = v323
	goto L57
L68:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v397)+20)) = int32(1)
	goto L35
L70:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L71:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L72:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L73:
	;
	if v2661+int32(816) != v2644 {
		goto L429
	} else {
		goto L430
	}
L74:
	;
	F_syncrep_yyerror(m, v449, v432, int32(420004))
	mBase = m.M
	v2642 = m.ExcPending
	if v2642 != 0 {
		goto L6
	} else {
		goto L428
	}
L75:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v442))) = uint16(v447)
	v457 = v450 << (uint(int32(1)) % 32)
	if base.Ui32(v429+v457-int32(2)) <= base.Ui32(v442) {
		goto L77
	} else {
		goto L78
	}
L76:
	;
	F_syncrep_yyerror(m, v2473, v2456, int32(201609))
	mBase = m.M
	v2605 = m.ExcPending
	if v2605 != 0 {
		goto L6
	} else {
		goto L424
	}
L77:
	;
	if base.Ui32(int32(9999)) < base.Ui32(v450) {
		goto L74
	} else {
		goto L80
	}
L78:
	;
	v510 = v429
	v512 = v440
	v514 = v442
	v515 = v444
	v516 = v450
	goto L79
L79:
	;
	v518 = int32(1) << (uint(v447) % 32)
	if v518&int32(13390146) != 0 {
		v2453 = v510
		v2454 = v430
		v2455 = v431
		v2456 = v432
		v2458 = v434
		v2464 = v512
		v2465 = v518
		v2466 = v514
		v2468 = v515
		v2470 = v446
		v2471 = v447
		v2473 = v449
		v2474 = v516
		v2478 = v454
		goto L103
	} else {
		goto L104
	}
L80:
	;
	v464 = int32(10000)
	if base.Ui32(v464) <= base.Ui32(v457) {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v467 = v464
	goto L83
L82:
	;
	v467 = v457
	goto L83
L83:
	;
	v472 = F_palloc(m, v467*int32(6)|int32(3))
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L6
	} else {
		goto L84
	}
L84:
	;
	if v472 == int32(0) {
		goto L74
	} else {
		goto L85
	}
L85:
	;
	v477 = int32(1)
	v480 = (v442-v429)>>(uint(v477)%32) + v477
	v482 = v480 << (uint(v477) % 32)
	if v482 != 0 {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	v487 = v484 + v467<<(uint(int32(1))%32)
	v489 = v480 << (uint(int32(2)) % 32)
	if v489 != 0 {
		goto L91
	} else {
		goto L92
	}
L87:
	;
	v483 = F__emscripten_memcpy_bulkmem(m, v472, v429, v482)
	mBase = m.M
	v484 = v483
	goto L89
L88:
	;
	v484 = v472
	goto L89
L89:
	;
	goto L86
L90:
	;
	if v446+int32(816) != v429 {
		goto L94
	} else {
		goto L95
	}
L91:
	;
	v490 = F__emscripten_memcpy_bulkmem(m, v487, v444, v489)
	mBase = m.M
	v491 = v490
	goto L93
L92:
	;
	v491 = v487
	goto L93
L93:
	;
	goto L90
L94:
	;
	F_pfree(m, v429)
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L6
	} else {
		goto L97
	}
L95:
	;
	goto L96
L96:
	;
	v497 = int32(1)
	v500 = v484 + v480<<(uint(v497)%32)
	if base.Ui32(v484+v467<<(uint(v497)%32)) <= base.Ui32(v500) {
		goto L98
	} else {
		goto L99
	}
L97:
	;
	goto L96
L98:
	;
	v2644 = v484
	v2645 = v430
	v2646 = v431
	v2659 = v497
	v2661 = v446
	goto L73
L99:
	;
	goto L100
L100:
	;
	v510 = v484
	v512 = v491 + v489 - int32(4)
	v514 = v500 - int32(2)
	v515 = v491
	v516 = v467
	goto L79
L101:
	;
	goto L76
L102:
	;
	v429 = v2574
	v430 = v2575
	v431 = v2576
	v432 = v2577
	v434 = v2579
	v440 = v2585
	v442 = v2587 + int32(2)
	v444 = v2589
	v446 = v2591
	v447 = v2600
	v449 = v2594
	v450 = v2595
	v454 = v2599
	goto L75
L103:
	;
	if v2465&int32(3386937) != 0 {
		goto L101
	} else {
		goto L405
	}
L104:
	;
	v523 = int32(*(*int8)(unsafe.Add(mBase, uint32(v447)+uint32(_consts[707]))))
	if v434 == int32(-2) {
		goto L106
	} else {
		goto L107
	}
L105:
	;
	v2433 = v523 + v2432
	if base.Ui32(int32(22)) < base.Ui32(v2433) {
		v2453 = v2395
		v2454 = v2396
		v2455 = v2397
		v2456 = v2398
		v2458 = v2431
		v2464 = v2406
		v2465 = v2407
		v2466 = v2408
		v2468 = v2410
		v2470 = v2412
		v2471 = v2413
		v2473 = v2415
		v2474 = v2416
		v2478 = v2420
		goto L103
	} else {
		goto L397
	}
L106:
	;
	v526 = m.G0
	v528 = v526 - int32(32)
	m.G0 = v528
	*(*int32)(unsafe.Add(mBase, uint32(v432)+92)) = v446 + int32(1228)
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v432)+40))
	if v533 == int32(0) {
		goto L110
	} else {
		goto L111
	}
L107:
	;
	v2395 = v510
	v2396 = v430
	v2397 = v431
	v2398 = v432
	v2400 = v434
	v2406 = v512
	v2407 = v518
	v2408 = v514
	v2410 = v515
	v2412 = v446
	v2413 = v447
	v2415 = v449
	v2416 = v516
	v2420 = v454
	goto L108
L108:
	;
	if v2400 <= int32(0) {
		goto L393
	} else {
		goto L394
	}
L109:
	;
	v2395 = v600
	v2396 = v601
	v2397 = v602
	v2398 = v603
	v2400 = v1744
	v2406 = v611
	v2407 = v612
	v2408 = v613
	v2410 = v615
	v2412 = v617
	v2413 = v618
	v2415 = v620
	v2416 = v621
	v2420 = v625
	goto L108
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v432)+40)) = int32(1)
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v432)+44))
	if v538 == int32(0) {
		goto L113
	} else {
		goto L114
	}
L111:
	;
	goto L112
L112:
	;
	v600 = v510
	v601 = v430
	v602 = v431
	v603 = v432
	v611 = v512
	v612 = v518
	v613 = v514
	v615 = v515
	v617 = v446
	v618 = v447
	v619 = v528
	v620 = v449
	v621 = v516
	v625 = v454
	goto L129
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v432)+44)) = int32(1)
	goto L115
L114:
	;
	goto L115
L115:
	;
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v432)+4))
	if v543 == int32(0) {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v547 = *(*int32)(unsafe.Add(mBase, _consts[708]))
	*(*int32)(unsafe.Add(mBase, uint32(v432)+4)) = v547
	goto L118
L117:
	;
	goto L118
L118:
	;
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v432)+8))
	if v549 == int32(0) {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v553 = *(*int32)(unsafe.Add(mBase, _consts[709]))
	*(*int32)(unsafe.Add(mBase, uint32(v432)+8)) = v553
	goto L121
L120:
	;
	goto L121
L121:
	;
	v555 = *(*int32)(unsafe.Add(mBase, uint32(v432)+20))
	if v555 != 0 {
		goto L123
	} else {
		goto L124
	}
L122:
	;
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v582)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v432)+28)) = v583
	v587 = v580 + v581<<(uint(int32(2))%32)
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v587)))
	v589 = *(*int32)(unsafe.Add(mBase, uint32(v588)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v432)+80)) = v589
	*(*int32)(unsafe.Add(mBase, uint32(v432)+36)) = v589
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v587)))
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v592)))
	*(*int32)(unsafe.Add(mBase, uint32(v432)+4)) = v593
	v595 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v589))))
	*(*uint8)(unsafe.Add(mBase, uint32(v432)+24)) = uint8(v595)
	goto L112
L123:
	;
	v556 = *(*int32)(unsafe.Add(mBase, uint32(v432)+12))
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v555+v556<<(uint(int32(2))%32))))
	if v560 != 0 {
		v580 = v555
		v581 = v556
		v582 = v560
		goto L122
	} else {
		goto L126
	}
L124:
	;
	goto L125
L125:
	;
	F_syncrep_yyensure_buffer_stack(m, v432)
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L6
	} else {
		goto L127
	}
L126:
	;
	goto L125
L127:
	;
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v432)+4))
	v566 = F_syncrep_yy_create_buffer(m, v565, v432)
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L6
	} else {
		goto L128
	}
L128:
	;
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v432)+20))
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v432)+12))
	v570 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v568+v569<<(uint(v570)%32)))) = v566
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v432)+20))
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v432)+12))
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v574+v575<<(uint(v570)%32))))
	v580 = v574
	v581 = v575
	v582 = v579
	goto L122
L129:
	;
	v626 = *(*int32)(unsafe.Add(mBase, uint32(v603)+36))
	v627 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v603)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v626))) = uint8(v627)
	v629 = *(*int32)(unsafe.Add(mBase, uint32(v603)+44))
	v635 = v629
	v636 = v626
	v638 = v626
	goto L131
L131:
	;
	v656 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v638))))
	v659 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v656)+uint32(_consts[710]))))
	if base.Ui32(int32(-26)) <= base.Ui32(v635&int32(2147483647)-int32(31)) {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v603)+68)) = v638
	*(*int32)(unsafe.Add(mBase, uint32(v603)+64)) = v635
	goto L135
L134:
	;
	goto L135
L135:
	;
	v668 = int32(1)
	v672 = int32(*(*int16)(unsafe.Add(mBase, uint32(v635<<(uint(v668)%32))+uint32(_consts[711]))))
	v673 = v672 + v659
	v678 = int32(*(*int16)(unsafe.Add(mBase, uint32(v673<<(uint(v668)%32))+uint32(_consts[712]))))
	if v678 != v635 {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v685 = v635
	v687 = v659
	v689 = v659
	goto L139
L137:
	;
	v744 = v673
	goto L138
L138:
	;
	v766 = int32(1)
	v772 = int32(*(*int16)(unsafe.Add(mBase, uint32(v744<<(uint(v766)%32))+uint32(_consts[713]))))
	if int64(1)<<(uint(base.I64_extend_i32_u(v744&int32(2147483647)))%64)&int64(-32985348833280) == int64(0) {
		v635 = v772
		v638 = v638 + v766
		goto L131
	} else {
		goto L145
	}
L139:
	;
	v710 = int32(*(*int16)(unsafe.Add(mBase, uint32(v685<<(uint(int32(1))%32))+uint32(_consts[714]))))
	if int64(1)<<(uint(base.I64_extend_i32_u(v685&int32(2147483647)))%64)&int64(2076672024) != int64(0) {
		goto L141
	} else {
		goto L142
	}
L140:
	;
	v744 = v733
	goto L138
L141:
	;
	v722 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v687)+uint32(_consts[715]))))
	v723 = v722
	goto L143
L142:
	;
	v723 = v689
	goto L143
L143:
	;
	v727 = v723 & int32(255)
	v728 = int32(1)
	v732 = int32(*(*int16)(unsafe.Add(mBase, uint32(v710<<(uint(v728)%32))+uint32(_consts[711]))))
	v733 = v727 + v732
	v738 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v733<<(uint(v728)%32))+uint32(_consts[712]))))
	if v710&int32(65535) != v738 {
		v685 = v710
		v687 = v727
		v689 = v723
		goto L139
	} else {
		goto L144
	}
L144:
	;
	goto L140
L145:
	;
	v792 = v636
	goto L146
L146:
	;
	v808 = *(*int32)(unsafe.Add(mBase, uint32(v603)+64))
	v809 = *(*int32)(unsafe.Add(mBase, uint32(v603)+68))
	v815 = v808
	v820 = v792
	v824 = v809
	goto L148
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v603)+80)) = v820
	*(*int32)(unsafe.Add(mBase, uint32(v603)+32)) = v824 - v820
	v839 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v824))))
	*(*uint8)(unsafe.Add(mBase, uint32(v603)+24)) = uint8(v839)
	v841 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v824))) = uint8(v841)
	*(*int32)(unsafe.Add(mBase, uint32(v603)+36)) = v824
	v848 = int32(*(*int16)(unsafe.Add(mBase, uint32(v815<<(uint(int32(1))%32))+uint32(_consts[716]))))
	v853 = v848
	goto L150
L150:
	;
	v875 = int32(260)
	switch v853 {
	case 0:
		goto L179
	case 1:
		goto L129
	case 2:
		goto L166
	case 3:
		goto L165
	case 4:
		goto L178
	case 5:
		goto L177
	case 6:
		goto L176
	case 7:
		goto L175
	case 8:
		goto L173
	case 9:
		goto L172
	case 10:
		goto L171
	case 11:
		goto L164
	case 12:
		goto L163
	case 13:
		goto L162
	case 14:
		v1744 = v875
		goto L161
	case 15:
		goto L170
	case 16:
		goto L168
	case 17:
		goto L169
	case 18:
		goto L174
	default:
		goto L167
	}
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v603)+36)) = v2365
	*(*int32)(unsafe.Add(mBase, uint32(v603)+48)) = int32(0)
	v2388 = *(*int32)(unsafe.Add(mBase, uint32(v603)+44))
	v2392 = base.I32_div_s(v2388-int32(1), int32(2))
	v853 = v2392 + int32(17)
	goto L150
L153:
	;
	F_yy_fatal_error_4(m, int32(29049))
	mBase = m.M
	v2358 = m.ExcPending
	if v2358 != 0 {
		goto L6
	} else {
		goto L392
	}
L154:
	;
	F_yy_fatal_error_4(m, int32(641909))
	mBase = m.M
	v2353 = m.ExcPending
	if v2353 != 0 {
		goto L6
	} else {
		goto L391
	}
L155:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L156:
	;
	v2199 = v2177 + v2189
	*(*int32)(unsafe.Add(mBase, uint32(v603)+36)) = v2199
	v2201 = *(*int32)(unsafe.Add(mBase, uint32(v603)+44))
	if base.Ui32(v2199) <= base.Ui32(v2179) {
		v815 = v2201
		v820 = v2179
		v824 = v2199
		goto L148
	} else {
		goto L372
	}
L157:
	;
	v1812 = *(*int32)(unsafe.Add(mBase, uint32(v1811)))
	*(*int32)(unsafe.Add(mBase, uint32(v1812)+16)) = v1789
	v1815 = *(*int32)(unsafe.Add(mBase, uint32(v603)+28))
	if v1815 != 0 {
		v1937 = int32(0)
		goto L316
	} else {
		goto L317
	}
L158:
	;
	v1780 = *(*int32)(unsafe.Add(mBase, uint32(v603)+20))
	v1781 = *(*int32)(unsafe.Add(mBase, uint32(v603)+12))
	v1789 = v1758
	v1811 = v1780 + v1781<<(uint(int32(2))%32)
	goto L157
L159:
	;
	F_yy_fatal_error_4(m, int32(432918))
	mBase = m.M
	v1753 = m.ExcPending
	if v1753 != 0 {
		goto L6
	} else {
		goto L315
	}
L160:
	;
	F_yy_fatal_error_4(m, int32(428872))
	mBase = m.M
	v1750 = m.ExcPending
	if v1750 != 0 {
		goto L6
	} else {
		goto L314
	}
L161:
	;
	m.G0 = v619 + int32(32)
	goto L109
L162:
	;
	v1744 = int32(41)
	goto L161
L163:
	;
	v1744 = int32(40)
	goto L161
L164:
	;
	v1744 = int32(44)
	goto L161
L165:
	;
	v1744 = int32(262)
	goto L161
L166:
	;
	v1744 = int32(261)
	goto L161
L167:
	;
	F_yy_fatal_error_4(m, int32(403701))
	mBase = m.M
	v1737 = m.ExcPending
	if v1737 != 0 {
		goto L6
	} else {
		goto L313
	}
L168:
	;
	v939 = *(*int32)(unsafe.Add(mBase, uint32(v603)+80))
	v940 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v603)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v824))) = uint8(v940)
	v942 = *(*int32)(unsafe.Add(mBase, uint32(v603)+20))
	v943 = *(*int32)(unsafe.Add(mBase, uint32(v603)+12))
	v946 = v942 + v943<<(uint(int32(2))%32)
	v947 = *(*int32)(unsafe.Add(mBase, uint32(v946)))
	v948 = *(*int32)(unsafe.Add(mBase, uint32(v947)+44))
	if v948 == int32(0) {
		goto L192
	} else {
		goto L193
	}
L169:
	;
	v1744 = int32(0)
	goto L161
L170:
	;
	F_yy_fatal_error_4(m, int32(432270))
	mBase = m.M
	v937 = m.ExcPending
	if v937 != 0 {
		goto L6
	} else {
		goto L191
	}
L171:
	;
	v931 = *(*int32)(unsafe.Add(mBase, uint32(v603)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v931))) = int32(628501)
	v1744 = int32(258)
	goto L161
L172:
	;
	v925 = *(*int32)(unsafe.Add(mBase, uint32(v603)+80))
	v926 = F_pstrdup(m, v925)
	mBase = m.M
	v927 = m.ExcPending
	if v927 != 0 {
		goto L6
	} else {
		goto L190
	}
L173:
	;
	v919 = *(*int32)(unsafe.Add(mBase, uint32(v603)+80))
	v920 = F_pstrdup(m, v919)
	mBase = m.M
	v921 = m.ExcPending
	if v921 != 0 {
		goto L6
	} else {
		goto L189
	}
L174:
	;
	v901 = *(*int32)(unsafe.Add(mBase, uint32(v620)))
	if v901 != 0 {
		v1744 = v875
		goto L161
	} else {
		goto L183
	}
L175:
	;
	v891 = *(*int32)(unsafe.Add(mBase, uint32(v603)+92))
	v892 = *(*int32)(unsafe.Add(mBase, uint32(v603)))
	v893 = *(*int32)(unsafe.Add(mBase, uint32(v892)))
	*(*int32)(unsafe.Add(mBase, uint32(v891))) = v893
	v895 = *(*int32)(unsafe.Add(mBase, uint32(v603)))
	*(*int32)(unsafe.Add(mBase, uint32(v895))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v603)+44)) = int32(1)
	v1744 = int32(258)
	goto L161
L176:
	;
	v887 = *(*int32)(unsafe.Add(mBase, uint32(v603)))
	v888 = *(*int32)(unsafe.Add(mBase, uint32(v603)+80))
	F_appendStringInfoString(m, v887, v888)
	mBase = m.M
	v890 = m.ExcPending
	if v890 != 0 {
		goto L6
	} else {
		goto L182
	}
L177:
	;
	v883 = *(*int32)(unsafe.Add(mBase, uint32(v603)))
	F_appendStringInfoChar(m, v883, int32(34))
	mBase = m.M
	v886 = m.ExcPending
	if v886 != 0 {
		goto L6
	} else {
		goto L181
	}
L178:
	;
	v878 = *(*int32)(unsafe.Add(mBase, uint32(v603)))
	F_initStringInfo(m, v878)
	mBase = m.M
	v880 = m.ExcPending
	if v880 != 0 {
		goto L6
	} else {
		goto L180
	}
L179:
	;
	v876 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v603)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v824))) = uint8(v876)
	v792 = v820
	goto L146
L180:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v603)+44)) = int32(3)
	goto L129
L181:
	;
	goto L129
L182:
	;
	goto L129
L183:
	;
	v902 = *(*int32)(unsafe.Add(mBase, uint32(v603)+80))
	v903 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v902))))
	if v903 != 0 {
		goto L184
	} else {
		goto L185
	}
L184:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v619)+20)) = v902
	*(*int32)(unsafe.Add(mBase, uint32(v619)+16)) = int32(211476)
	v910 = F_psprintf(m, int32(660175), v619+int32(16))
	mBase = m.M
	v911 = m.ExcPending
	if v911 != 0 {
		goto L6
	} else {
		goto L187
	}
L185:
	;
	goto L186
L186:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v619))) = int32(211476)
	v916 = F_psprintf(m, int32(62096), v619)
	mBase = m.M
	v917 = m.ExcPending
	if v917 != 0 {
		goto L6
	} else {
		goto L188
	}
L187:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v620))) = v910
	v1744 = v875
	goto L161
L188:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v620))) = v916
	v1744 = v875
	goto L161
L189:
	;
	v922 = *(*int32)(unsafe.Add(mBase, uint32(v603)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v922))) = v920
	v1744 = int32(258)
	goto L161
L190:
	;
	v928 = *(*int32)(unsafe.Add(mBase, uint32(v603)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v928))) = v926
	v1744 = int32(259)
	goto L161
L191:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L192:
	;
	v951 = *(*int32)(unsafe.Add(mBase, uint32(v947)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v603)+28)) = v951
	v953 = *(*int32)(unsafe.Add(mBase, uint32(v946)))
	v954 = *(*int32)(unsafe.Add(mBase, uint32(v603)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v953))) = v954
	v956 = *(*int32)(unsafe.Add(mBase, uint32(v603)+20))
	v957 = *(*int32)(unsafe.Add(mBase, uint32(v603)+12))
	v958 = int32(2)
	v961 = *(*int32)(unsafe.Add(mBase, uint32(v956+v957<<(uint(v958)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v961)+44)) = int32(1)
	v964 = *(*int32)(unsafe.Add(mBase, uint32(v603)+20))
	v965 = *(*int32)(unsafe.Add(mBase, uint32(v603)+12))
	v969 = *(*int32)(unsafe.Add(mBase, uint32(v964+v965<<(uint(v958)%32))))
	v970 = v969
	v971 = v964
	v972 = v965
	goto L194
L193:
	;
	v970 = v947
	v971 = v942
	v972 = v943
	goto L194
L194:
	;
	v973 = *(*int32)(unsafe.Add(mBase, uint32(v603)+36))
	v974 = *(*int32)(unsafe.Add(mBase, uint32(v970)+4))
	v975 = *(*int32)(unsafe.Add(mBase, uint32(v603)+28))
	v976 = v974 + v975
	if base.Ui32(v973) <= base.Ui32(v976) {
		goto L195
	} else {
		goto L196
	}
L195:
	;
	v978 = *(*int32)(unsafe.Add(mBase, uint32(v603)+80))
	v982 = v978 + (v939 ^ int32(-1)) + v824
	*(*int32)(unsafe.Add(mBase, uint32(v603)+36)) = v982
	v984 = *(*int32)(unsafe.Add(mBase, uint32(v603)+44))
	if base.Ui32(v978) < base.Ui32(v982) {
		goto L198
	} else {
		goto L199
	}
L196:
	;
	goto L197
L197:
	;
	if base.Ui32(v976+int32(1)) < base.Ui32(v973) {
		goto L160
	} else {
		goto L230
	}
L198:
	;
	v991 = v984
	v994 = v978
	goto L201
L199:
	;
	v1139 = v984
	goto L200
L200:
	;
	if base.Ui32(int32(-26)) <= base.Ui32(v1139&int32(2147483647)-int32(31)) {
		goto L219
	} else {
		goto L220
	}
L201:
	;
	v1012 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v994))))
	if v1012 != 0 {
		goto L203
	} else {
		goto L204
	}
L202:
	;
	v1139 = v1130
	goto L200
L203:
	;
	v1015 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1012)+uint32(_consts[710]))))
	v1017 = v1015
	goto L205
L204:
	;
	v1017 = int32(1)
	goto L205
L205:
	;
	if base.Ui32(int32(-26)) <= base.Ui32(v991&int32(2147483647)-int32(31)) {
		goto L206
	} else {
		goto L207
	}
L206:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v603)+68)) = v994
	*(*int32)(unsafe.Add(mBase, uint32(v603)+64)) = v991
	goto L208
L207:
	;
	goto L208
L208:
	;
	v1027 = v1017 & int32(255)
	v1028 = int32(1)
	v1032 = int32(*(*int16)(unsafe.Add(mBase, uint32(v991<<(uint(v1028)%32))+uint32(_consts[711]))))
	v1033 = v1027 + v1032
	v1038 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1033<<(uint(v1028)%32))+uint32(_consts[712]))))
	if v1038 != v991 {
		goto L209
	} else {
		goto L210
	}
L209:
	;
	v1045 = v991
	v1047 = v1017
	v1049 = v1027
	goto L212
L210:
	;
	v1104 = v1033
	goto L211
L211:
	;
	v1126 = int32(1)
	v1130 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1104<<(uint(v1126)%32))+uint32(_consts[713]))))
	v1132 = v994 + v1126
	if v1132 != v982 {
		v991 = v1130
		v994 = v1132
		goto L201
	} else {
		goto L218
	}
L212:
	;
	v1070 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1045<<(uint(int32(1))%32))+uint32(_consts[714]))))
	if int64(1)<<(uint(base.I64_extend_i32_u(v1045&int32(2147483647)))%64)&int64(2076672024) != int64(0) {
		goto L214
	} else {
		goto L215
	}
L213:
	;
	v1104 = v1093
	goto L211
L214:
	;
	v1082 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1049)+uint32(_consts[715]))))
	v1083 = v1082
	goto L216
L215:
	;
	v1083 = v1047
	goto L216
L216:
	;
	v1087 = v1083 & int32(255)
	v1088 = int32(1)
	v1092 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1070<<(uint(v1088)%32))+uint32(_consts[711]))))
	v1093 = v1087 + v1092
	v1098 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1093<<(uint(v1088)%32))+uint32(_consts[712]))))
	if v1070&int32(65535) != v1098 {
		v1045 = v1070
		v1047 = v1083
		v1049 = v1087
		goto L212
	} else {
		goto L217
	}
L217:
	;
	goto L213
L218:
	;
	goto L202
L219:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v603)+68)) = v982
	*(*int32)(unsafe.Add(mBase, uint32(v603)+64)) = v1139
	goto L221
L220:
	;
	goto L221
L221:
	;
	v1168 = int32(1)
	v1172 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1139<<(uint(v1168)%32))+uint32(_consts[711]))))
	v1174 = v1172 + v1168
	v1179 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1174<<(uint(v1168)%32))+uint32(_consts[712]))))
	if v1179 != v1139 {
		goto L222
	} else {
		goto L223
	}
L222:
	;
	v1186 = v1139
	goto L225
L223:
	;
	v1230 = v1174
	goto L224
L224:
	;
	v1254 = v1230 & int32(2147483647)
	if int64(1)<<(uint(base.I64_extend_i32_u(v1254))%64)&int64(-32985348833280) != int64(0) {
		v792 = v978
		goto L146
	} else {
		goto L228
	}
L225:
	;
	v1207 = int32(1)
	v1211 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1186<<(uint(v1207)%32))+uint32(_consts[714]))))
	v1212 = base.I32_extend16_s(v1211)
	v1217 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1212<<(uint(v1207)%32))+uint32(_consts[711]))))
	v1219 = v1217 + v1207
	v1224 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1219<<(uint(v1207)%32))+uint32(_consts[712]))))
	if v1211 != v1224 {
		v1186 = v1212
		goto L225
	} else {
		goto L227
	}
L226:
	;
	v1230 = v1219
	goto L224
L227:
	;
	goto L226
L228:
	;
	if v1254 == int32(0) {
		v792 = v978
		goto L146
	} else {
		goto L229
	}
L229:
	;
	v1263 = int32(1)
	v1264 = v982 + v1263
	*(*int32)(unsafe.Add(mBase, uint32(v603)+36)) = v1264
	v1270 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1230<<(uint(v1263)%32))+uint32(_consts[713]))))
	v635 = v1270
	v636 = v978
	v638 = v1264
	goto L131
L230:
	;
	v1274 = *(*int32)(unsafe.Add(mBase, uint32(v603)+80))
	v1275 = *(*int32)(unsafe.Add(mBase, uint32(v970)+40))
	if v1275 == int32(0) {
		goto L231
	} else {
		goto L232
	}
L231:
	;
	if v973-v1274 != int32(1) {
		v2177 = v974
		v2179 = v1274
		v2189 = v975
		goto L156
	} else {
		goto L234
	}
L232:
	;
	goto L233
L233:
	;
	v1283 = v1274 ^ int32(-1) + v973
	if v1283 != 0 {
		goto L235
	} else {
		goto L236
	}
L234:
	;
	v2365 = v1274
	goto L152
L235:
	;
	v1284 = int32(7)
	v1285 = v1283 & v1284
	if base.Ui32(v973-v1274-int32(2)) < base.Ui32(v1284) {
		goto L239
	} else {
		goto L240
	}
L236:
	;
	v1442 = v970
	v1444 = v971
	v1446 = v972
	goto L237
L237:
	;
	v1463 = *(*int32)(unsafe.Add(mBase, uint32(v1442)+44))
	if v1463 == int32(2) {
		goto L251
	} else {
		goto L252
	}
L238:
	;
	if v1285 != 0 {
		goto L245
	} else {
		goto L246
	}
L239:
	;
	v1347 = v974
	v1348 = v1274
	goto L238
L240:
	;
	goto L241
L241:
	;
	v1298 = v974
	v1299 = v1274
	v1301 = int32(0)
	goto L242
L242:
	;
	v1320 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1299))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1298))) = uint8(v1320)
	v1322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1299)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1298)+1)) = uint8(v1322)
	v1324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1299)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1298)+2)) = uint8(v1324)
	v1326 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1299)+3)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1298)+3)) = uint8(v1326)
	v1328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1299)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1298)+4)) = uint8(v1328)
	v1330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1299)+5)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1298)+5)) = uint8(v1330)
	v1332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1299)+6)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1298)+6)) = uint8(v1332)
	v1334 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1299)+7)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1298)+7)) = uint8(v1334)
	v1336 = int32(8)
	v1337 = v1298 + v1336
	v1339 = v1299 + v1336
	v1341 = v1301 + v1336
	if v1341 != v1283&int32(-8) {
		v1298 = v1337
		v1299 = v1339
		v1301 = v1341
		goto L242
	} else {
		goto L244
	}
L243:
	;
	v1347 = v1337
	v1348 = v1339
	goto L238
L244:
	;
	goto L243
L245:
	;
	v1374 = v1347
	v1375 = v1348
	v1377 = int32(0)
	goto L248
L246:
	;
	goto L247
L247:
	;
	v1431 = *(*int32)(unsafe.Add(mBase, uint32(v603)+20))
	v1432 = *(*int32)(unsafe.Add(mBase, uint32(v603)+12))
	v1436 = *(*int32)(unsafe.Add(mBase, uint32(v1431+v1432<<(uint(int32(2))%32))))
	v1442 = v1436
	v1444 = v1431
	v1446 = v1432
	goto L237
L248:
	;
	v1396 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1375))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1374))) = uint8(v1396)
	v1398 = int32(1)
	v1403 = v1377 + v1398
	if v1403 != v1285 {
		v1374 = v1374 + v1398
		v1375 = v1375 + v1398
		v1377 = v1403
		goto L248
	} else {
		goto L250
	}
L249:
	;
	goto L247
L250:
	;
	goto L249
L251:
	;
	v1466 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v603)+28)) = v1466
	v1789 = v1466
	v1811 = v1444 + v1446<<(uint(int32(2))%32)
	goto L157
L252:
	;
	goto L253
L253:
	;
	v1472 = *(*int32)(unsafe.Add(mBase, uint32(v1442)+12))
	v1473 = v1274 - v973
	v1474 = v1472 + v1473
	if v1474 == int32(0) {
		goto L254
	} else {
		goto L255
	}
L254:
	;
	v1477 = *(*int32)(unsafe.Add(mBase, uint32(v603)+36))
	v1482 = v1472
	v1483 = v1442
	v1487 = v1477
	goto L257
L255:
	;
	v1545 = v1442
	v1547 = v1474
	goto L256
L256:
	;
	v1566 = int32(8192)
	if base.Ui32(v1566) <= base.Ui32(v1547) {
		goto L273
	} else {
		goto L274
	}
L257:
	;
	v1504 = *(*int32)(unsafe.Add(mBase, uint32(v1483)+20))
	if v1504 == int32(0) {
		goto L259
	} else {
		goto L260
	}
L258:
	;
	v1545 = v1535
	v1547 = v1537
	goto L256
L259:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1483)+4)) = int32(0)
	goto L153
L260:
	;
	goto L261
L261:
	;
	v1509 = *(*int32)(unsafe.Add(mBase, uint32(v1483)+4))
	if base.Ui32(int32(2147483646)) <= base.Ui32(v1482) {
		goto L262
	} else {
		goto L263
	}
L262:
	;
	v1515 = int32(-3)
	goto L264
L263:
	;
	v1515 = v1482 << (uint(int32(1)) % 32)
	goto L264
L264:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1483)+12)) = v1515
	v1518 = v1515 + int32(2)
	if v1509 != 0 {
		goto L266
	} else {
		goto L267
	}
L265:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1483)+4)) = v1523
	if v1523 == int32(0) {
		goto L153
	} else {
		goto L271
	}
L266:
	;
	v1519 = F_repalloc(m, v1509, v1518)
	mBase = m.M
	v1520 = m.ExcPending
	if v1520 != 0 {
		goto L6
	} else {
		goto L269
	}
L267:
	;
	goto L268
L268:
	;
	v1521 = F_palloc(m, v1518)
	mBase = m.M
	v1522 = m.ExcPending
	if v1522 != 0 {
		goto L6
	} else {
		goto L270
	}
L269:
	;
	v1523 = v1519
	goto L265
L270:
	;
	v1523 = v1521
	goto L265
L271:
	;
	v1528 = v1523 + (v1487 - v1509)
	*(*int32)(unsafe.Add(mBase, uint32(v603)+36)) = v1528
	v1530 = *(*int32)(unsafe.Add(mBase, uint32(v603)+20))
	v1531 = *(*int32)(unsafe.Add(mBase, uint32(v603)+12))
	v1535 = *(*int32)(unsafe.Add(mBase, uint32(v1530+v1531<<(uint(int32(2))%32))))
	v1536 = *(*int32)(unsafe.Add(mBase, uint32(v1535)+12))
	v1537 = v1536 + v1473
	if v1537 == int32(0) {
		v1482 = v1536
		v1483 = v1535
		v1487 = v1528
		goto L257
	} else {
		goto L272
	}
L272:
	;
	goto L258
L273:
	;
	v1569 = v1566
	goto L275
L274:
	;
	v1569 = v1547
	goto L275
L275:
	;
	v1571 = *(*int32)(unsafe.Add(mBase, uint32(v1545)+24))
	if v1571 != 0 {
		goto L276
	} else {
		goto L277
	}
L276:
	;
	v1576 = int32(0)
	goto L280
L277:
	;
	goto L278
L278:
	;
	*(*int32)(unsafe.Add(mBase, _consts[40])) = int32(0)
	v1651 = *(*int32)(unsafe.Add(mBase, uint32(v603)+20))
	v1652 = *(*int32)(unsafe.Add(mBase, uint32(v603)+12))
	v1656 = *(*int32)(unsafe.Add(mBase, uint32(v1651+v1652<<(uint(int32(2))%32))))
	v1657 = *(*int32)(unsafe.Add(mBase, uint32(v1656)+4))
	v1660 = *(*int32)(unsafe.Add(mBase, uint32(v603)+4))
	v1661 = F_fread(m, v1657+v1283, int32(1), v1569, v1660)
	mBase = m.M
	v1662 = m.ExcPending
	if v1662 != 0 {
		goto L6
	} else {
		goto L295
	}
L279:
	;
	switch v1602 {
	case 0:
		goto L287
	default:
		v1646 = v1616
		goto L285
	case 11:
		goto L286
	}
L280:
	;
	v1598 = *(*int32)(unsafe.Add(mBase, uint32(v603)+4))
	v1599 = F_do_getc(m, v1598)
	mBase = m.M
	v1600 = m.ExcPending
	if v1600 != 0 {
		goto L6
	} else {
		goto L283
	}
L281:
	;
	v1616 = v1569
	goto L279
L282:
	;
	v1603 = *(*int32)(unsafe.Add(mBase, uint32(v603)+20))
	v1604 = *(*int32)(unsafe.Add(mBase, uint32(v603)+12))
	v1608 = *(*int32)(unsafe.Add(mBase, uint32(v1603+v1604<<(uint(int32(2))%32))))
	v1609 = *(*int32)(unsafe.Add(mBase, uint32(v1608)+4))
	*(*uint8)(unsafe.Add(mBase, uint32(v1609+v1283+v1576))) = uint8(v1599)
	v1614 = v1576 + int32(1)
	if v1614 != v1569 {
		v1576 = v1614
		goto L280
	} else {
		goto L284
	}
L283:
	;
	v1602 = v1599 + int32(1)
	switch v1602 {
	case 0, 11:
		v1616 = v1576
		goto L279
	default:
		goto L282
	}
L284:
	;
	goto L281
L285:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v603)+28)) = v1646
	v1758 = v1646
	goto L158
L286:
	;
	v1633 = *(*int32)(unsafe.Add(mBase, uint32(v603)+20))
	v1634 = *(*int32)(unsafe.Add(mBase, uint32(v603)+12))
	v1638 = *(*int32)(unsafe.Add(mBase, uint32(v1633+v1634<<(uint(int32(2))%32))))
	v1639 = *(*int32)(unsafe.Add(mBase, uint32(v1638)+4))
	v1642 = int32(10)
	*(*uint8)(unsafe.Add(mBase, uint32(v1639+v1283+v1616))) = uint8(v1642)
	v1646 = v1616 + int32(1)
	goto L285
L287:
	;
	v1617 = *(*int32)(unsafe.Add(mBase, uint32(v603)+4))
	v1618 = *(*int32)(unsafe.Add(mBase, uint32(v1617)+76))
	if v1618 < int32(0) {
		goto L290
	} else {
		goto L291
	}
L288:
	;
	if int32(base.Ui32(v1623)>>(uint(int32(5))%32))&int32(1) == int32(0) {
		v1646 = v1616
		goto L285
	} else {
		goto L293
	}
L289:
	;
	goto L288
L290:
	;
	v1621 = *(*int32)(unsafe.Add(mBase, uint32(v1617)))
	v1623 = v1621
	goto L289
L291:
	;
	goto L292
L292:
	;
	v1622 = *(*int32)(unsafe.Add(mBase, uint32(v1617)))
	v1623 = v1622
	goto L289
L293:
	;
	F_yy_fatal_error_4(m, int32(432918))
	mBase = m.M
	v1632 = m.ExcPending
	if v1632 != 0 {
		goto L6
	} else {
		goto L294
	}
L294:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L295:
	;
	v1667 = v1661
	goto L296
L296:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v603)+28)) = v1667
	if v1667 != 0 {
		v1758 = v1667
		goto L158
	} else {
		goto L298
	}
L298:
	;
	v1690 = *(*int32)(unsafe.Add(mBase, uint32(v603)+4))
	v1691 = *(*int32)(unsafe.Add(mBase, uint32(v1690)+76))
	if v1691 < int32(0) {
		goto L301
	} else {
		goto L302
	}
L299:
	;
	if int32(base.Ui32(v1696)>>(uint(int32(5))%32))&int32(1) == int32(0) {
		goto L304
	} else {
		goto L305
	}
L300:
	;
	goto L299
L301:
	;
	v1694 = *(*int32)(unsafe.Add(mBase, uint32(v1690)))
	v1696 = v1694
	goto L300
L302:
	;
	goto L303
L303:
	;
	v1695 = *(*int32)(unsafe.Add(mBase, uint32(v1690)))
	v1696 = v1695
	goto L300
L304:
	;
	v1758 = int32(0)
	goto L158
L305:
	;
	goto L306
L306:
	;
	v1705 = *(*int32)(unsafe.Add(mBase, _consts[40]))
	if v1705 != int32(27) {
		goto L159
	} else {
		goto L307
	}
L307:
	;
	v1709 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[40])) = v1709
	v1711 = *(*int32)(unsafe.Add(mBase, uint32(v603)+4))
	v1712 = *(*int32)(unsafe.Add(mBase, uint32(v1711)+76))
	if v1709 <= v1712 {
		goto L309
	} else {
		goto L310
	}
L308:
	;
	v1723 = *(*int32)(unsafe.Add(mBase, uint32(v603)+20))
	v1724 = *(*int32)(unsafe.Add(mBase, uint32(v603)+12))
	v1728 = *(*int32)(unsafe.Add(mBase, uint32(v1723+v1724<<(uint(int32(2))%32))))
	v1729 = *(*int32)(unsafe.Add(mBase, uint32(v1728)+4))
	v1732 = *(*int32)(unsafe.Add(mBase, uint32(v603)+4))
	v1733 = F_fread(m, v1729+v1283, int32(1), v1569, v1732)
	mBase = m.M
	v1734 = m.ExcPending
	if v1734 != 0 {
		goto L6
	} else {
		goto L312
	}
L309:
	;
	v1715 = *(*int32)(unsafe.Add(mBase, uint32(v1711)))
	*(*int32)(unsafe.Add(mBase, uint32(v1711))) = v1715 & int32(-49)
	goto L308
L310:
	;
	goto L311
L311:
	;
	v1719 = *(*int32)(unsafe.Add(mBase, uint32(v1711)))
	*(*int32)(unsafe.Add(mBase, uint32(v1711))) = v1719 & int32(-49)
	goto L308
L312:
	;
	v1667 = v1733
	goto L296
L313:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L314:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L315:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L316:
	;
	v1938 = *(*int32)(unsafe.Add(mBase, uint32(v603)+28))
	v1939 = v1938 + v1283
	v1940 = *(*int32)(unsafe.Add(mBase, uint32(v603)+20))
	v1941 = *(*int32)(unsafe.Add(mBase, uint32(v603)+12))
	v1945 = *(*int32)(unsafe.Add(mBase, uint32(v1940+v1941<<(uint(int32(2))%32))))
	v1946 = *(*int32)(unsafe.Add(mBase, uint32(v1945)+12))
	if base.Ui32(v1946) < base.Ui32(v1939) {
		goto L340
	} else {
		goto L341
	}
L317:
	;
	if v1283 == int32(0) {
		goto L318
	} else {
		goto L319
	}
L318:
	;
	v1818 = *(*int32)(unsafe.Add(mBase, uint32(v603)+4))
	v1819 = *(*int32)(unsafe.Add(mBase, uint32(v603)+20))
	if v1819 != 0 {
		goto L323
	} else {
		goto L324
	}
L319:
	;
	goto L320
L320:
	;
	v1923 = *(*int32)(unsafe.Add(mBase, uint32(v603)+20))
	v1924 = *(*int32)(unsafe.Add(mBase, uint32(v603)+12))
	v1925 = int32(2)
	v1928 = *(*int32)(unsafe.Add(mBase, uint32(v1923+v1924<<(uint(v1925)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v1928)+44)) = v1925
	v1937 = v1925
	goto L316
L321:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1885)+40)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1885))) = v1818
	v1892 = *(*int32)(unsafe.Add(mBase, uint32(v603)+20))
	if v1892 != 0 {
		goto L336
	} else {
		goto L337
	}
L322:
	;
	v1842 = *(*int32)(unsafe.Add(mBase, _consts[40]))
	v1843 = *(*int32)(unsafe.Add(mBase, uint32(v603)+12))
	v1847 = *(*int32)(unsafe.Add(mBase, uint32(v1840+v1843<<(uint(int32(2))%32))))
	if v1847 == int32(0) {
		goto L330
	} else {
		goto L331
	}
L323:
	;
	v1820 = *(*int32)(unsafe.Add(mBase, uint32(v603)+12))
	v1824 = *(*int32)(unsafe.Add(mBase, uint32(v1819+v1820<<(uint(int32(2))%32))))
	if v1824 != 0 {
		v1840 = v1819
		goto L322
	} else {
		goto L326
	}
L324:
	;
	goto L325
L325:
	;
	F_syncrep_yyensure_buffer_stack(m, v603)
	mBase = m.M
	v1826 = m.ExcPending
	if v1826 != 0 {
		goto L6
	} else {
		goto L327
	}
L326:
	;
	goto L325
L327:
	;
	v1827 = *(*int32)(unsafe.Add(mBase, uint32(v603)+4))
	v1828 = F_syncrep_yy_create_buffer(m, v1827, v603)
	mBase = m.M
	v1829 = m.ExcPending
	if v1829 != 0 {
		goto L6
	} else {
		goto L328
	}
L328:
	;
	v1830 = *(*int32)(unsafe.Add(mBase, uint32(v603)+20))
	v1831 = *(*int32)(unsafe.Add(mBase, uint32(v603)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1830+v1831<<(uint(int32(2))%32)))) = v1828
	v1836 = *(*int32)(unsafe.Add(mBase, uint32(v603)+20))
	if v1836 != 0 {
		v1840 = v1836
		goto L322
	} else {
		goto L329
	}
L329:
	;
	v1838 = *(*int32)(unsafe.Add(mBase, _consts[40]))
	v1885 = int32(0)
	v1888 = v1838
	goto L321
L330:
	;
	v1885 = int32(0)
	v1888 = v1842
	goto L321
L331:
	;
	goto L332
L332:
	;
	v1851 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1847)+16)) = v1851
	v1853 = *(*int32)(unsafe.Add(mBase, uint32(v1847)+4))
	*(*uint8)(unsafe.Add(mBase, uint32(v1853))) = uint8(v1851)
	v1856 = *(*int32)(unsafe.Add(mBase, uint32(v1847)+4))
	*(*uint8)(unsafe.Add(mBase, uint32(v1856)+1)) = uint8(v1851)
	*(*int32)(unsafe.Add(mBase, uint32(v1847)+44)) = v1851
	*(*int32)(unsafe.Add(mBase, uint32(v1847)+28)) = int32(1)
	v1863 = *(*int32)(unsafe.Add(mBase, uint32(v1847)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1847)+8)) = v1863
	v1865 = *(*int32)(unsafe.Add(mBase, uint32(v603)+20))
	if v1865 == v1851 {
		v1885 = v1847
		v1888 = v1842
		goto L321
	} else {
		goto L333
	}
L333:
	;
	v1868 = *(*int32)(unsafe.Add(mBase, uint32(v603)+12))
	v1871 = v1865 + v1868<<(uint(int32(2))%32)
	v1872 = *(*int32)(unsafe.Add(mBase, uint32(v1871)))
	if v1847 != v1872 {
		v1885 = v1847
		v1888 = v1842
		goto L321
	} else {
		goto L334
	}
L334:
	;
	v1874 = *(*int32)(unsafe.Add(mBase, uint32(v1872)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v603)+28)) = v1874
	v1876 = *(*int32)(unsafe.Add(mBase, uint32(v1871)))
	v1877 = *(*int32)(unsafe.Add(mBase, uint32(v1876)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v603)+80)) = v1877
	*(*int32)(unsafe.Add(mBase, uint32(v603)+36)) = v1877
	v1880 = *(*int32)(unsafe.Add(mBase, uint32(v1871)))
	v1881 = *(*int32)(unsafe.Add(mBase, uint32(v1880)))
	*(*int32)(unsafe.Add(mBase, uint32(v603)+4)) = v1881
	v1883 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1877))))
	*(*uint8)(unsafe.Add(mBase, uint32(v603)+24)) = uint8(v1883)
	v1885 = v1847
	v1888 = v1842
	goto L321
L335:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1885)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[40])) = v1888
	v1905 = *(*int32)(unsafe.Add(mBase, uint32(v603)+20))
	v1906 = *(*int32)(unsafe.Add(mBase, uint32(v603)+12))
	v1909 = v1905 + v1906<<(uint(int32(2))%32)
	v1910 = *(*int32)(unsafe.Add(mBase, uint32(v1909)))
	v1911 = *(*int32)(unsafe.Add(mBase, uint32(v1910)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v603)+28)) = v1911
	v1913 = *(*int32)(unsafe.Add(mBase, uint32(v1909)))
	v1914 = *(*int32)(unsafe.Add(mBase, uint32(v1913)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v603)+36)) = v1914
	*(*int32)(unsafe.Add(mBase, uint32(v603)+80)) = v1914
	v1917 = *(*int32)(unsafe.Add(mBase, uint32(v1909)))
	v1918 = *(*int32)(unsafe.Add(mBase, uint32(v1917)))
	*(*int32)(unsafe.Add(mBase, uint32(v603)+4)) = v1918
	v1920 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1914))))
	*(*uint8)(unsafe.Add(mBase, uint32(v603)+24)) = uint8(v1920)
	v1937 = int32(1)
	goto L316
L336:
	;
	v1893 = *(*int32)(unsafe.Add(mBase, uint32(v603)+12))
	v1897 = *(*int32)(unsafe.Add(mBase, uint32(v1892+v1893<<(uint(int32(2))%32))))
	if v1885 == v1897 {
		goto L335
	} else {
		goto L339
	}
L337:
	;
	goto L338
L338:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1885)+32)) = int64(1)
	goto L335
L339:
	;
	goto L338
L340:
	;
	v1950 = v1939 + int32(base.Ui32(v1938)>>(uint(int32(1))%32))
	v1951 = *(*int32)(unsafe.Add(mBase, uint32(v1945)+4))
	if v1951 != 0 {
		goto L344
	} else {
		goto L345
	}
L341:
	;
	v1980 = v1939
	v1981 = v1940
	v1982 = v1941
	goto L342
L342:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v603)+28)) = v1980
	v1984 = int32(2)
	v1987 = *(*int32)(unsafe.Add(mBase, uint32(v1981+v1982<<(uint(v1984)%32))))
	v1988 = *(*int32)(unsafe.Add(mBase, uint32(v1987)+4))
	v1990 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1988+v1980))) = uint8(v1990)
	v1992 = *(*int32)(unsafe.Add(mBase, uint32(v603)+20))
	v1993 = *(*int32)(unsafe.Add(mBase, uint32(v603)+12))
	v1997 = *(*int32)(unsafe.Add(mBase, uint32(v1992+v1993<<(uint(v1984)%32))))
	v1998 = *(*int32)(unsafe.Add(mBase, uint32(v1997)+4))
	v1999 = *(*int32)(unsafe.Add(mBase, uint32(v603)+28))
	*(*uint8)(unsafe.Add(mBase, uint32(v1998+v1999)+1)) = uint8(v1990)
	v2003 = *(*int32)(unsafe.Add(mBase, uint32(v603)+20))
	v2004 = *(*int32)(unsafe.Add(mBase, uint32(v603)+12))
	v2007 = v2003 + v2004<<(uint(v1984)%32)
	v2008 = *(*int32)(unsafe.Add(mBase, uint32(v2007)))
	v2009 = *(*int32)(unsafe.Add(mBase, uint32(v2008)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v603)+80)) = v2009
	if v1937 == int32(1) {
		v2365 = v2009
		goto L152
	} else {
		goto L350
	}
L343:
	;
	v1957 = *(*int32)(unsafe.Add(mBase, uint32(v603)+20))
	v1958 = *(*int32)(unsafe.Add(mBase, uint32(v603)+12))
	v1959 = int32(2)
	v1962 = *(*int32)(unsafe.Add(mBase, uint32(v1957+v1958<<(uint(v1959)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v1962)+4)) = v1956
	v1964 = *(*int32)(unsafe.Add(mBase, uint32(v603)+20))
	v1965 = *(*int32)(unsafe.Add(mBase, uint32(v603)+12))
	v1969 = *(*int32)(unsafe.Add(mBase, uint32(v1964+v1965<<(uint(v1959)%32))))
	v1970 = *(*int32)(unsafe.Add(mBase, uint32(v1969)+4))
	if v1970 == int32(0) {
		goto L154
	} else {
		goto L349
	}
L344:
	;
	v1952 = F_repalloc(m, v1951, v1950)
	mBase = m.M
	v1953 = m.ExcPending
	if v1953 != 0 {
		goto L6
	} else {
		goto L347
	}
L345:
	;
	goto L346
L346:
	;
	v1954 = F_palloc(m, v1950)
	mBase = m.M
	v1955 = m.ExcPending
	if v1955 != 0 {
		goto L6
	} else {
		goto L348
	}
L347:
	;
	v1956 = v1952
	goto L343
L348:
	;
	v1956 = v1954
	goto L343
L349:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1969)+12)) = v1950 - int32(2)
	v1976 = *(*int32)(unsafe.Add(mBase, uint32(v603)+12))
	v1977 = *(*int32)(unsafe.Add(mBase, uint32(v603)+28))
	v1979 = *(*int32)(unsafe.Add(mBase, uint32(v603)+20))
	v1980 = v1977 + v1283
	v1981 = v1979
	v1982 = v1976
	goto L342
L350:
	;
	switch v1937 - int32(1) {
	case 0:
		goto L155
	case 1:
		goto L351
	default:
		goto L352
	}
L351:
	;
	v2170 = *(*int32)(unsafe.Add(mBase, uint32(v603)+28))
	v2171 = *(*int32)(unsafe.Add(mBase, uint32(v2007)))
	v2172 = *(*int32)(unsafe.Add(mBase, uint32(v2171)+4))
	v2177 = v2172
	v2179 = v2009
	v2189 = v2170
	goto L156
L352:
	;
	v2018 = v2009 + (v939 ^ int32(-1)) + v824
	*(*int32)(unsafe.Add(mBase, uint32(v603)+36)) = v2018
	v2020 = *(*int32)(unsafe.Add(mBase, uint32(v603)+44))
	if base.Ui32(v2018) <= base.Ui32(v2009) {
		v635 = v2020
		v636 = v2009
		v638 = v2018
		goto L131
	} else {
		goto L353
	}
L353:
	;
	v2027 = v2020
	v2036 = v2009
	goto L354
L354:
	;
	v2048 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2036))))
	if v2048 != 0 {
		goto L356
	} else {
		goto L357
	}
L355:
	;
	v635 = v2166
	v636 = v2009
	v638 = v2018
	goto L131
L356:
	;
	v2051 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2048)+uint32(_consts[710]))))
	v2053 = v2051
	goto L358
L357:
	;
	v2053 = int32(1)
	goto L358
L358:
	;
	if base.Ui32(int32(-26)) <= base.Ui32(v2027&int32(2147483647)-int32(31)) {
		goto L359
	} else {
		goto L360
	}
L359:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v603)+68)) = v2036
	*(*int32)(unsafe.Add(mBase, uint32(v603)+64)) = v2027
	goto L361
L360:
	;
	goto L361
L361:
	;
	v2063 = v2053 & int32(255)
	v2064 = int32(1)
	v2068 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2027<<(uint(v2064)%32))+uint32(_consts[711]))))
	v2069 = v2063 + v2068
	v2074 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2069<<(uint(v2064)%32))+uint32(_consts[712]))))
	if v2074 != v2027 {
		goto L362
	} else {
		goto L363
	}
L362:
	;
	v2081 = v2027
	v2083 = v2053
	v2085 = v2063
	goto L365
L363:
	;
	v2140 = v2069
	goto L364
L364:
	;
	v2162 = int32(1)
	v2166 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2140<<(uint(v2162)%32))+uint32(_consts[713]))))
	v2168 = v2036 + v2162
	if v2018 != v2168 {
		v2027 = v2166
		v2036 = v2168
		goto L354
	} else {
		goto L371
	}
L365:
	;
	v2106 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2081<<(uint(int32(1))%32))+uint32(_consts[714]))))
	if int64(1)<<(uint(base.I64_extend_i32_u(v2081&int32(2147483647)))%64)&int64(2076672024) != int64(0) {
		goto L367
	} else {
		goto L368
	}
L366:
	;
	v2140 = v2129
	goto L364
L367:
	;
	v2118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2085)+uint32(_consts[715]))))
	v2119 = v2118
	goto L369
L368:
	;
	v2119 = v2083
	goto L369
L369:
	;
	v2123 = v2119 & int32(255)
	v2124 = int32(1)
	v2128 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2106<<(uint(v2124)%32))+uint32(_consts[711]))))
	v2129 = v2123 + v2128
	v2134 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2129<<(uint(v2124)%32))+uint32(_consts[712]))))
	if v2106&int32(65535) != v2134 {
		v2081 = v2106
		v2083 = v2119
		v2085 = v2123
		goto L365
	} else {
		goto L370
	}
L370:
	;
	goto L366
L371:
	;
	goto L355
L372:
	;
	v2208 = v2201
	v2211 = v2179
	goto L373
L373:
	;
	v2229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2211))))
	if v2229 != 0 {
		goto L375
	} else {
		goto L376
	}
L374:
	;
	v815 = v2347
	v820 = v2179
	v824 = v2199
	goto L148
L375:
	;
	v2232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2229)+uint32(_consts[710]))))
	v2234 = v2232
	goto L377
L376:
	;
	v2234 = int32(1)
	goto L377
L377:
	;
	if base.Ui32(int32(-26)) <= base.Ui32(v2208&int32(2147483647)-int32(31)) {
		goto L378
	} else {
		goto L379
	}
L378:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v603)+68)) = v2211
	*(*int32)(unsafe.Add(mBase, uint32(v603)+64)) = v2208
	goto L380
L379:
	;
	goto L380
L380:
	;
	v2244 = v2234 & int32(255)
	v2245 = int32(1)
	v2249 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2208<<(uint(v2245)%32))+uint32(_consts[711]))))
	v2250 = v2244 + v2249
	v2255 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2250<<(uint(v2245)%32))+uint32(_consts[712]))))
	if v2255 != v2208 {
		goto L381
	} else {
		goto L382
	}
L381:
	;
	v2262 = v2208
	v2264 = v2234
	v2266 = v2244
	goto L384
L382:
	;
	v2321 = v2250
	goto L383
L383:
	;
	v2343 = int32(1)
	v2347 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2321<<(uint(v2343)%32))+uint32(_consts[713]))))
	v2349 = v2211 + v2343
	if v2349 != v2199 {
		v2208 = v2347
		v2211 = v2349
		goto L373
	} else {
		goto L390
	}
L384:
	;
	v2287 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2262<<(uint(int32(1))%32))+uint32(_consts[714]))))
	if int64(1)<<(uint(base.I64_extend_i32_u(v2262&int32(2147483647)))%64)&int64(2076672024) != int64(0) {
		goto L386
	} else {
		goto L387
	}
L385:
	;
	v2321 = v2310
	goto L383
L386:
	;
	v2299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2266)+uint32(_consts[715]))))
	v2300 = v2299
	goto L388
L387:
	;
	v2300 = v2264
	goto L388
L388:
	;
	v2304 = v2300 & int32(255)
	v2305 = int32(1)
	v2309 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2287<<(uint(v2305)%32))+uint32(_consts[711]))))
	v2310 = v2304 + v2309
	v2315 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2310<<(uint(v2305)%32))+uint32(_consts[712]))))
	if v2287&int32(65535) != v2315 {
		v2262 = v2287
		v2264 = v2300
		v2266 = v2304
		goto L384
	} else {
		goto L389
	}
L389:
	;
	goto L385
L390:
	;
	goto L374
L391:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L392:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L393:
	;
	v2423 = int32(0)
	v2431 = v2423
	v2432 = v2423
	goto L105
L394:
	;
	goto L395
L395:
	;
	if base.Ui32(int32(262)) < base.Ui32(v2400) {
		v2431 = v2400
		v2432 = int32(2)
		goto L105
	} else {
		goto L396
	}
L396:
	;
	v2430 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2400)+uint32(_consts[717]))))
	v2431 = v2400
	v2432 = v2430
	goto L105
L397:
	;
	v2438 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2433)+uint32(_consts[718]))))
	if v2432 != v2438 {
		v2453 = v2395
		v2454 = v2396
		v2455 = v2397
		v2456 = v2398
		v2458 = v2431
		v2464 = v2406
		v2465 = v2407
		v2466 = v2408
		v2468 = v2410
		v2470 = v2412
		v2471 = v2413
		v2473 = v2415
		v2474 = v2416
		v2478 = v2420
		goto L103
	} else {
		goto L398
	}
L398:
	;
	if v2433 != int32(19) {
		goto L399
	} else {
		goto L400
	}
L399:
	;
	v2443 = v2406 + int32(4)
	v2444 = *(*int32)(unsafe.Add(mBase, uint32(v2412)+1228))
	*(*int32)(unsafe.Add(mBase, uint32(v2443))) = v2444
	if v2431 != 0 {
		goto L402
	} else {
		goto L403
	}
L400:
	;
	goto L401
L401:
	;
	v2644 = v2395
	v2645 = v2396
	v2646 = v2397
	v2659 = int32(0)
	v2661 = v2412
	goto L73
L402:
	;
	v2448 = int32(-2)
	goto L404
L403:
	;
	v2448 = int32(0)
	goto L404
L404:
	;
	v2451 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2433)+uint32(_consts[719]))))
	v2574 = v2395
	v2575 = v2396
	v2576 = v2397
	v2577 = v2398
	v2579 = v2448
	v2585 = v2443
	v2587 = v2408
	v2589 = v2410
	v2591 = v2412
	v2594 = v2415
	v2595 = v2416
	v2599 = v2420
	v2600 = v2451
	goto L102
L405:
	;
	v2484 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2471)+uint32(_consts[720]))))
	v2487 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2484)+uint32(_consts[721]))))
	v2489 = int32(2)
	v2492 = *(*int32)(unsafe.Add(mBase, uint32(v2464+(int32(1)-v2487)<<(uint(v2489)%32))))
	switch v2484 - v2489 {
	case 0:
		goto L414
	case 1:
		goto L413
	case 2:
		goto L412
	case 3:
		goto L411
	case 4:
		goto L410
	case 5:
		goto L409
	case 6:
		goto L408
	case 7, 8:
		goto L407
	default:
		v2544 = v2492
		goto L406
	}
L406:
	;
	v2549 = v2464 - v2487<<(uint(int32(2))%32) + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v2549))) = v2544
	v2553 = v2466 - v2487<<(uint(int32(1))%32)
	v2554 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2553))))
	v2557 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2484)+uint32(_consts[722]))))
	v2560 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2557)+uint32(_consts[722]))))
	v2561 = v2554 + v2560
	if base.Ui32(int32(22)) < base.Ui32(v2561) {
		goto L421
	} else {
		goto L422
	}
L407:
	;
	v2543 = *(*int32)(unsafe.Add(mBase, uint32(v2464)))
	v2544 = v2543
	goto L406
L408:
	;
	v2539 = *(*int32)(unsafe.Add(mBase, uint32(v2464-int32(8))))
	v2540 = *(*int32)(unsafe.Add(mBase, uint32(v2464)))
	v2541 = F_lappend(m, v2539, v2540)
	mBase = m.M
	v2542 = m.ExcPending
	if v2542 != 0 {
		goto L6
	} else {
		goto L420
	}
L409:
	;
	v2529 = *(*int32)(unsafe.Add(mBase, uint32(v2464)))
	*(*int32)(unsafe.Add(mBase, uint32(v2470)+8)) = v2529
	*(*int32)(unsafe.Add(mBase, uint32(v2470)+12)) = v2529
	v2535 = F_list_make1_impl(m, int32(1), v2470+int32(8))
	mBase = m.M
	v2536 = m.ExcPending
	if v2536 != 0 {
		goto L6
	} else {
		goto L419
	}
L410:
	;
	v2522 = *(*int32)(unsafe.Add(mBase, uint32(v2464-int32(12))))
	v2525 = *(*int32)(unsafe.Add(mBase, uint32(v2464-int32(4))))
	v2527 = F_create_syncrep_config(m, v2522, v2525, int32(0))
	mBase = m.M
	v2528 = m.ExcPending
	if v2528 != 0 {
		goto L6
	} else {
		goto L418
	}
L411:
	;
	v2513 = *(*int32)(unsafe.Add(mBase, uint32(v2464-int32(12))))
	v2516 = *(*int32)(unsafe.Add(mBase, uint32(v2464-int32(4))))
	v2518 = F_create_syncrep_config(m, v2513, v2516, int32(1))
	mBase = m.M
	v2519 = m.ExcPending
	if v2519 != 0 {
		goto L6
	} else {
		goto L417
	}
L412:
	;
	v2504 = *(*int32)(unsafe.Add(mBase, uint32(v2464-int32(12))))
	v2507 = *(*int32)(unsafe.Add(mBase, uint32(v2464-int32(4))))
	v2509 = F_create_syncrep_config(m, v2504, v2507, int32(0))
	mBase = m.M
	v2510 = m.ExcPending
	if v2510 != 0 {
		goto L6
	} else {
		goto L416
	}
L413:
	;
	v2498 = *(*int32)(unsafe.Add(mBase, uint32(v2464)))
	v2500 = F_create_syncrep_config(m, int32(529597), v2498, int32(0))
	mBase = m.M
	v2501 = m.ExcPending
	if v2501 != 0 {
		goto L6
	} else {
		goto L415
	}
L414:
	;
	v2495 = *(*int32)(unsafe.Add(mBase, uint32(v2464)))
	*(*int32)(unsafe.Add(mBase, uint32(v2478))) = v2495
	v2544 = v2492
	goto L406
L415:
	;
	v2544 = v2500
	goto L406
L416:
	;
	v2544 = v2509
	goto L406
L417:
	;
	v2544 = v2518
	goto L406
L418:
	;
	v2544 = v2527
	goto L406
L419:
	;
	v2544 = v2535
	goto L406
L420:
	;
	v2544 = v2541
	goto L406
L421:
	;
	v2573 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2557)+uint32(_consts[723]))))
	v2574 = v2453
	v2575 = v2454
	v2576 = v2455
	v2577 = v2456
	v2579 = v2458
	v2585 = v2549
	v2587 = v2553
	v2589 = v2468
	v2591 = v2470
	v2594 = v2473
	v2595 = v2474
	v2599 = v2478
	v2600 = v2573
	goto L102
L422:
	;
	v2566 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2561)+uint32(_consts[718]))))
	if v2566 != v2554 {
		goto L421
	} else {
		goto L423
	}
L423:
	;
	v2570 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2561)+uint32(_consts[719]))))
	v2574 = v2453
	v2575 = v2454
	v2576 = v2455
	v2577 = v2456
	v2579 = v2458
	v2585 = v2549
	v2587 = v2553
	v2589 = v2468
	v2591 = v2470
	v2594 = v2473
	v2595 = v2474
	v2599 = v2478
	v2600 = v2570
	goto L102
L424:
	;
	v2619 = v2466
	goto L425
L425:
	;
	if base.B2i32(v2453 == v2619) == int32(0) {
		v2619 = v2619 - int32(2)
		goto L425
	} else {
		goto L427
	}
L426:
	;
	v2644 = v2453
	v2645 = v2454
	v2646 = v2455
	v2659 = int32(1)
	v2661 = v2470
	goto L73
L427:
	;
	goto L426
L428:
	;
	v2644 = v429
	v2645 = v430
	v2646 = v431
	v2659 = int32(2)
	v2661 = v446
	goto L73
L429:
	;
	F_pfree(m, v2644)
	mBase = m.M
	v2674 = m.ExcPending
	if v2674 != 0 {
		goto L6
	} else {
		goto L432
	}
L430:
	;
	goto L431
L431:
	;
	m.G0 = v2661 + int32(1232)
	v2678 = *(*int32)(unsafe.Add(mBase, uint32(v2646)+44))
	F_replication_scanner_finish(m, v2678)
	mBase = m.M
	v2680 = m.ExcPending
	if v2680 != 0 {
		goto L6
	} else {
		goto L433
	}
L432:
	;
	goto L431
L433:
	;
	if v2659 == int32(0) {
		goto L435
	} else {
		goto L436
	}
L434:
	;
	v2709 = *(*int32)(unsafe.Add(mBase, uint32(v2683)+4))
	if v2709 <= int32(0) {
		goto L447
	} else {
		goto L448
	}
L435:
	;
	v2683 = *(*int32)(unsafe.Add(mBase, uint32(v2646)+40))
	if v2683 != 0 {
		goto L434
	} else {
		goto L438
	}
L436:
	;
	goto L437
L437:
	;
	*(*int32)(unsafe.Add(mBase, _consts[518])) = int32(16801924)
	goto L439
L438:
	;
	goto L437
L439:
	;
	v2688 = *(*int32)(unsafe.Add(mBase, uint32(v2646)+36))
	v2690 = *(*int32)(unsafe.Add(mBase, _consts[40]))
	*(*int32)(unsafe.Add(mBase, _consts[497])) = v2690
	goto L440
L440:
	;
	if v2688 != 0 {
		goto L442
	} else {
		goto L443
	}
L441:
	;
	*(*int32)(unsafe.Add(mBase, _consts[498])) = v2706
	v2769 = v2646
	v2793 = int32(0)
	goto L1
L442:
	;
	v2694 = *(*int32)(unsafe.Add(mBase, uint32(v2646)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v2646)+16)) = v2694
	v2699 = F_format_elog_string(m, int32(195784), v2646+int32(16))
	mBase = m.M
	v2700 = m.ExcPending
	if v2700 != 0 {
		goto L6
	} else {
		goto L445
	}
L443:
	;
	goto L444
L444:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2646))) = int32(153312)
	v2704 = F_format_elog_string(m, int32(608714), v2646)
	mBase = m.M
	v2705 = m.ExcPending
	if v2705 != 0 {
		goto L6
	} else {
		goto L446
	}
L445:
	;
	v2706 = v2699
	goto L441
L446:
	;
	v2706 = v2704
	goto L441
L447:
	;
	v2713 = *(*int32)(unsafe.Add(mBase, _consts[40]))
	*(*int32)(unsafe.Add(mBase, _consts[497])) = v2713
	goto L450
L448:
	;
	goto L449
L449:
	;
	v2728 = *(*int32)(unsafe.Add(mBase, uint32(v2683)))
	v2729 = F_guc_malloc(m, v2728)
	mBase = m.M
	v2730 = m.ExcPending
	if v2730 != 0 {
		goto L6
	} else {
		goto L452
	}
L450:
	;
	v2716 = *(*int32)(unsafe.Add(mBase, uint32(v2646)+40))
	v2717 = *(*int32)(unsafe.Add(mBase, uint32(v2716)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v2646)+32)) = v2717
	v2723 = F_format_elog_string(m, int32(227705), v2646+int32(32))
	mBase = m.M
	v2724 = m.ExcPending
	if v2724 != 0 {
		goto L6
	} else {
		goto L451
	}
L451:
	;
	*(*int32)(unsafe.Add(mBase, _consts[516])) = v2723
	v2769 = v2646
	v2793 = int32(0)
	goto L1
L452:
	;
	if v2729 == int32(0) {
		v2769 = v2646
		v2793 = int32(0)
		goto L1
	} else {
		goto L453
	}
L453:
	;
	v2733 = *(*int32)(unsafe.Add(mBase, uint32(v2646)+40))
	v2734 = *(*int32)(unsafe.Add(mBase, uint32(v2733)))
	if v2734 != 0 {
		goto L455
	} else {
		goto L456
	}
L454:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2645))) = v2736
	v2742 = v2646
	goto L2
L455:
	;
	v2735 = F__emscripten_memcpy_bulkmem(m, v2729, v2733, v2734)
	mBase = m.M
	v2736 = v2735
	goto L457
L456:
	;
	v2736 = v2729
	goto L457
L457:
	;
	goto L454
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
		*(*int32)(unsafe.Add(mBase, _consts[40])) = int32(0) - v3
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
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = *(*int32)(unsafe.Add(mBase, _consts[451]))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if int32(0) <= v11 {
		if v11 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v217 = m.ExcPending
			if v217 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(261))
				mBase = m.M
				v220 = m.ExcPending
				if v220 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(419831), int32(0))
					mBase = m.M
					v224 = m.ExcPending
					if v224 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(470404), int32(1049), int32(203143))
						mBase = m.M
						v229 = m.ExcPending
						if v229 != 0 {
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
			if v14 != int32(6) {
				if base.Ui32(v14) <= base.Ui32(int32(41)) {
					v177 = *(*int32)(unsafe.Add(mBase, uint32(v14*int32(28))+uint32(_consts[970])))
					v178 = v177
				} else {
					v178 = int32(1)
				}
				if int32(1) < v178 {
					v181 = base.B2i32(base.Ui32(v11) < base.Ui32(int32(128)))
				} else {
					v181 = base.B2i32(base.Ui32(v11) < base.Ui32(int32(256)))
				}
				if v181 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v249 = m.ExcPending
					if v249 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(261))
						mBase = m.M
						v252 = m.ExcPending
						if v252 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v11
							F_errmsg(m, int32(55837), v9+int32(32))
							mBase = m.M
							v258 = m.ExcPending
							if v258 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(470404), int32(1121), int32(203143))
								mBase = m.M
								v263 = m.ExcPending
								if v263 != 0 {
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
					v185 = F_palloc(m, int32(5))
					mBase = m.M
					v186 = m.ExcPending
					if v186 != 0 {
						return int32(0)
					} else {
						*(*uint8)(unsafe.Add(mBase, uint32(v185)+4)) = uint8(v11)
						*(*int32)(unsafe.Add(mBase, uint32(v185))) = int32(20)
						v190 = v185
						m.G0 = v9 + int32(48)
						return v190
					}
				}
			} else {
				if base.Ui32(v11) < base.Ui32(int32(128)) {
					if base.Ui32(v14) <= base.Ui32(int32(41)) {
						v177 = *(*int32)(unsafe.Add(mBase, uint32(v14*int32(28))+uint32(_consts[970])))
						v178 = v177
					} else {
						v178 = int32(1)
					}
					if int32(1) < v178 {
						v181 = base.B2i32(base.Ui32(v11) < base.Ui32(int32(128)))
					} else {
						v181 = base.B2i32(base.Ui32(v11) < base.Ui32(int32(256)))
					}
					if v181 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v249 = m.ExcPending
						if v249 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(261))
							mBase = m.M
							v252 = m.ExcPending
							if v252 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v11
								F_errmsg(m, int32(55837), v9+int32(32))
								mBase = m.M
								v258 = m.ExcPending
								if v258 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(470404), int32(1121), int32(203143))
									mBase = m.M
									v263 = m.ExcPending
									if v263 != 0 {
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
						v185 = F_palloc(m, int32(5))
						mBase = m.M
						v186 = m.ExcPending
						if v186 != 0 {
							return int32(0)
						} else {
							*(*uint8)(unsafe.Add(mBase, uint32(v185)+4)) = uint8(v11)
							*(*int32)(unsafe.Add(mBase, uint32(v185))) = int32(20)
							v190 = v185
							m.G0 = v9 + int32(48)
							return v190
						}
					}
				} else {
					if base.Ui32(int32(1114112)) <= base.Ui32(v11) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v233 = m.ExcPending
						if v233 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(261))
							mBase = m.M
							v236 = m.ExcPending
							if v236 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v11
								F_errmsg(m, int32(55837), v9)
								mBase = m.M
								v240 = m.ExcPending
								if v240 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(470404), int32(1068), int32(203143))
									mBase = m.M
									v245 = m.ExcPending
									if v245 != 0 {
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
						v29 = base.B2i32(base.Ui32(int32(2047)) < base.Ui32(v11))
						if base.Ui32(int32(2047)) < base.Ui32(v11) {
							v30 = int32(3)
						} else {
							v30 = int32(2)
						}
						if base.Ui32(int32(65535)) < base.Ui32(v11) {
							v33 = int32(4)
						} else {
							v33 = v30
						}
						v35 = v33 + int32(4)
						v36 = F_palloc(m, v35)
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v36))) = v35 << (uint(int32(2)) % 32)
							v44 = v36 + int32(4)
							if v29 == int32(0) {
								v50 = int32(base.Ui32(v11)>>(uint(int32(6))%32)) | int32(192)
								*(*uint8)(unsafe.Add(mBase, uint32(v44))) = uint8(v50)
								v90 = int32(5)
							} else {
								if base.Ui32(v11-int32(2048)) <= base.Ui32(int32(63487)) {
									v60 = int32(base.Ui32(v11)>>(uint(int32(12))%32)) | int32(224)
									*(*uint8)(unsafe.Add(mBase, uint32(v36)+4)) = uint8(v60)
									v62 = int32(6)
									v67 = int32(base.Ui32(v11)>>(uint(v62)%32))&int32(63) | int32(128)
									*(*uint8)(unsafe.Add(mBase, uint32(v36)+5)) = uint8(v67)
									v90 = v62
								} else {
									v73 = int32(base.Ui32(v11)>>(uint(int32(18))%32)) | int32(240)
									*(*uint8)(unsafe.Add(mBase, uint32(v36)+4)) = uint8(v73)
									v77 = int32(63)
									v79 = int32(128)
									v80 = int32(base.Ui32(v11)>>(uint(int32(6))%32))&v77 | v79
									*(*uint8)(unsafe.Add(mBase, uint32(v36)+6)) = uint8(v80)
									v87 = int32(base.Ui32(v11)>>(uint(int32(12))%32))&v77 | v79
									*(*uint8)(unsafe.Add(mBase, uint32(v36)+5)) = uint8(v87)
									v90 = int32(7)
								}
							}
							v95 = v11&int32(63) | int32(128)
							*(*uint8)(unsafe.Add(mBase, uint32(v90+v36))) = uint8(v95)
							v97 = int32(0)
							switch v33 - int32(1) {
							case 0:
								v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
								v133 = v132
								if base.I32_extend8_s(v133) < int32(-62) {
									v146 = v97
								} else {
									v138 = v133
									v146 = base.B2i32(base.Ui32(v138&int32(255)) < base.Ui32(int32(245)))
								}
							case 1:
								v106 = int32(*(*int8)(unsafe.Add(mBase, uint32(v44)+1)))
								v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
								switch v107 - int32(224) {
								case 0:
									v110 = int32(224)
									if base.Ui32(v110) <= base.Ui32((v106-int32(-64))&int32(255)) {
										v138 = v110
										v146 = base.B2i32(base.Ui32(v138&int32(255)) < base.Ui32(int32(245)))
									} else {
										v146 = v97
									}
								default:
									if v106 <= int32(-65) {
										v133 = v107
										if base.I32_extend8_s(v133) < int32(-62) {
											v146 = v97
										} else {
											v138 = v133
											v146 = base.B2i32(base.Ui32(v138&int32(255)) < base.Ui32(int32(245)))
										}
									} else {
										v146 = v97
									}
								case 13:
									if int32(-97) < v106 {
										v146 = v97
									} else {
										v138 = int32(237)
										v146 = base.B2i32(base.Ui32(v138&int32(255)) < base.Ui32(int32(245)))
									}
								case 16:
									if base.Ui32((v106-int32(-64))&int32(255)) < base.Ui32(int32(208)) {
										v146 = v97
									} else {
										v138 = int32(240)
										v146 = base.B2i32(base.Ui32(v138&int32(255)) < base.Ui32(int32(245)))
									}
								case 20:
									if int32(-113) < v106 {
										v146 = v97
									} else {
										v138 = int32(244)
										v146 = base.B2i32(base.Ui32(v138&int32(255)) < base.Ui32(int32(245)))
									}
								}
							case 2:
								v103 = int32(*(*int8)(unsafe.Add(mBase, uint32(v44)+2)))
								if int32(-65) < v103 {
									v146 = v97
								} else {
									v106 = int32(*(*int8)(unsafe.Add(mBase, uint32(v44)+1)))
									v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
									switch v107 - int32(224) {
									case 0:
										v110 = int32(224)
										if base.Ui32(v110) <= base.Ui32((v106-int32(-64))&int32(255)) {
											v138 = v110
											v146 = base.B2i32(base.Ui32(v138&int32(255)) < base.Ui32(int32(245)))
										} else {
											v146 = v97
										}
									default:
										if v106 <= int32(-65) {
											v133 = v107
											if base.I32_extend8_s(v133) < int32(-62) {
												v146 = v97
											} else {
												v138 = v133
												v146 = base.B2i32(base.Ui32(v138&int32(255)) < base.Ui32(int32(245)))
											}
										} else {
											v146 = v97
										}
									case 13:
										if int32(-97) < v106 {
											v146 = v97
										} else {
											v138 = int32(237)
											v146 = base.B2i32(base.Ui32(v138&int32(255)) < base.Ui32(int32(245)))
										}
									case 16:
										if base.Ui32((v106-int32(-64))&int32(255)) < base.Ui32(int32(208)) {
											v146 = v97
										} else {
											v138 = int32(240)
											v146 = base.B2i32(base.Ui32(v138&int32(255)) < base.Ui32(int32(245)))
										}
									case 20:
										if int32(-113) < v106 {
											v146 = v97
										} else {
											v138 = int32(244)
											v146 = base.B2i32(base.Ui32(v138&int32(255)) < base.Ui32(int32(245)))
										}
									}
								}
							case 3:
								v100 = int32(*(*int8)(unsafe.Add(mBase, uint32(v44)+3)))
								if int32(-65) < v100 {
									v146 = v97
								} else {
									v103 = int32(*(*int8)(unsafe.Add(mBase, uint32(v44)+2)))
									if int32(-65) < v103 {
										v146 = v97
									} else {
										v106 = int32(*(*int8)(unsafe.Add(mBase, uint32(v44)+1)))
										v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
										switch v107 - int32(224) {
										case 0:
											v110 = int32(224)
											if base.Ui32(v110) <= base.Ui32((v106-int32(-64))&int32(255)) {
												v138 = v110
												v146 = base.B2i32(base.Ui32(v138&int32(255)) < base.Ui32(int32(245)))
											} else {
												v146 = v97
											}
										default:
											if v106 <= int32(-65) {
												v133 = v107
												if base.I32_extend8_s(v133) < int32(-62) {
													v146 = v97
												} else {
													v138 = v133
													v146 = base.B2i32(base.Ui32(v138&int32(255)) < base.Ui32(int32(245)))
												}
											} else {
												v146 = v97
											}
										case 13:
											if int32(-97) < v106 {
												v146 = v97
											} else {
												v138 = int32(237)
												v146 = base.B2i32(base.Ui32(v138&int32(255)) < base.Ui32(int32(245)))
											}
										case 16:
											if base.Ui32((v106-int32(-64))&int32(255)) < base.Ui32(int32(208)) {
												v146 = v97
											} else {
												v138 = int32(240)
												v146 = base.B2i32(base.Ui32(v138&int32(255)) < base.Ui32(int32(245)))
											}
										case 20:
											if int32(-113) < v106 {
												v146 = v97
											} else {
												v138 = int32(244)
												v146 = base.B2i32(base.Ui32(v138&int32(255)) < base.Ui32(int32(245)))
											}
										}
									}
								}
							default:
								v146 = v97
							}
							if v146 != 0 {
								v190 = v36
								m.G0 = v9 + int32(48)
								return v190
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v150 = m.ExcPending
								if v150 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(261))
									mBase = m.M
									v153 = m.ExcPending
									if v153 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v11
										F_errmsg(m, int32(55884), v9+int32(16))
										mBase = m.M
										v159 = m.ExcPending
										if v159 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(470404), int32(1109), int32(203143))
											mBase = m.M
											v164 = m.ExcPending
											if v164 != 0 {
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
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v201 = m.ExcPending
		if v201 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50856066))
			mBase = m.M
			v204 = m.ExcPending
			if v204 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(325745), int32(0))
				mBase = m.M
				v208 = m.ExcPending
				if v208 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(470404), int32(1045), int32(203143))
					mBase = m.M
					v213 = m.ExcPending
					if v213 != 0 {
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
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
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
			v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
			if v36&int32(1) != 0 {
				v39 = v29
			} else {
				v39 = v31
			}
			v42 = int32(1)
			v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
			if v46&v42 != 0 {
				v49 = l0 + v42
			} else {
				v49 = l0 + int32(4)
			}
			v55 = base.I32_div_s(v2+int32(7), int32(8))
			if v55 != 0 {
				v56 = F__emscripten_memcpy_bulkmem(m, v39+int32(2), v49+int32(2), v55)
				mBase = m.M
			} else {
			}
			v59 = v2 & int32(7)
			if v59 == int32(0) {
			} else {
				v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
				if v64&int32(1) != 0 {
					v67 = v29
				} else {
					v67 = v31
				}
				v70 = int32(base.Ui32(v2)>>(uint(int32(3))%32)) + v67 + int32(2)
				v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
				v74 = v71 & (int32(-256) >> (uint(v59) % 32))
				*(*uint8)(unsafe.Add(mBase, uint32(v70))) = uint8(v74)
			}
		}
		v80 = int32(1)
		v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
		if v82&v80 != 0 {
			v85 = v80
		} else {
			v85 = int32(4)
		}
		v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8+v85))))
		if v87 == int32(2) {
			v90 = int32(40)
		} else {
			v90 = int32(88)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v8))) = v90
		return v8
	}
}
func F_clause_selectivity_ext(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) float64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v46 float64
	_ = v46
	var v51 float64
	_ = v51
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
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
	var v68 int32
	_ = v68
	var v72 float64
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 float64
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 float64
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 float64
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 float64
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 float64
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
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
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v147 int32
	_ = v147
	var v150 float64
	_ = v150
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
	var v189 float64
	_ = v189
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 float64
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v218 float64
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v225 float64
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v233 float64
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v240 float64
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
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
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v272 float64
	_ = v272
	var v273 int32
	_ = v273
	var v274 float64
	_ = v274
	var v275 int32
	_ = v275
	var v276 float64
	_ = v276
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 float64
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v298 float32
	_ = v298
	var v299 float64
	_ = v299
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v312 float64
	_ = v312
	var v313 int32
	_ = v313
	var v314 float32
	_ = v314
	var v315 float64
	_ = v315
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 float64
	_ = v321
	var v322 float64
	_ = v322
	var v323 float64
	_ = v323
	var v329 int32
	_ = v329
	var v335 int32
	_ = v335
	var v340 int32
	_ = v340
	var v343 float64
	_ = v343
	var v347 int32
	_ = v347
	var v359 int32
	_ = v359
	var v365 int32
	_ = v365
	var v370 int32
	_ = v370
	var v376 float64
	_ = v376
	var v377 int32
	_ = v377
	var v382 int32
	_ = v382
	var v386 int32
	_ = v386
	var v391 int32
	_ = v391
	var v392 float64
	_ = v392
	var v393 int32
	_ = v393
	var v395 float64
	_ = v395
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v403 float64
	_ = v403
	var v413 float64
	_ = v413
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v423 float64
	_ = v423
	var v428 float64
	_ = v428
	var v429 int32
	_ = v429
	var v430 float64
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 float64
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v437 float64
	_ = v437
	var v438 int32
	_ = v438
	var v439 float64
	_ = v439
	var v441 int32
	_ = v441
	var v444 float64
	_ = v444
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v462 float64
	_ = v462
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v471 float64
	_ = v471
	var v487 float64
	_ = v487
	var v495 int32
	_ = v495
	var v510 float64
	_ = v510
	v7 = int32(0)
	if l1 == v7 {
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
	return v510
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
	v66 = v7
	v68 = int32(0)
	goto L5
L7:
	;
	goto L8
L8:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+10)))
	if v25 != int32(1) {
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
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	if v29 == int32(7) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	return float64(1)
L12:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	if v58 == int32(0) {
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
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	switch v36 {
	case 0:
		goto L13
	case 1:
		goto L15
	default:
		v56 = v7
		goto L12
	}
L15:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v38 = F_bms_is_member(m, l2, v37)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	return float64(0)
L17:
	;
	if v38 == int32(0) {
		v56 = v7
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
	v46 = *(*float64)(unsafe.Add(mBase, uint32(l1)+80))
	if base.F64_ge(v46, float64(0)) == int32(0) {
		goto L19
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v51 = *(*float64)(unsafe.Add(mBase, uint32(l1)+88))
	if base.F64_ge(v51, float64(0)) != 0 {
		v510 = v51
		goto L4
	} else {
		goto L24
	}
L23:
	;
	v510 = v46
	goto L4
L24:
	;
	goto L19
L25:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v62 = v61
	goto L27
L26:
	;
	v62 = v58
	goto L27
L27:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	v64 = v62
	v65 = v63
	v66 = v56
	v68 = l1
	goto L5
L28:
	;
	if v66 == int32(0) {
		v510 = v487
		goto L4
	} else {
		goto L189
	}
L29:
	;
	v446 = m.G0
	v448 = v446 - int32(32)
	m.G0 = v448
	F_examine_variable(m, l0, v64, l2, v448)
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L16
	} else {
		goto L181
	}
L30:
	;
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	if v441 == int32(18) {
		goto L178
	} else {
		goto L179
	}
L31:
	;
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v64)+28))
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v64)+24))
	v437 = F_restriction_selectivity(m, l0, v196, v435, v436, l2)
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L16
	} else {
		goto L177
	}
L32:
	;
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	v433 = F_clause_selectivity_ext(m, l0, v432, l2, l3, l4, l5)
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L16
	} else {
		goto L176
	}
L33:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	v430 = F_clause_selectivity_ext(m, l0, v429, l2, l3, l4, l5)
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L16
	} else {
		goto L175
	}
L34:
	;
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	v421 = F_find_base_rel(m, l0, v420)
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L16
	} else {
		goto L171
	}
L35:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v64)+8))
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	v286 = m.G0
	v288 = v286 - int32(112)
	m.G0 = v288
	F_examine_variable(m, l0, v285, l2, v288+int32(80))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L16
	} else {
		goto L125
	}
L36:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v64)+8))
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	v282 = F_nulltestsel(m, l0, v280, v281, l2)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L16
	} else {
		goto L124
	}
L37:
	;
	v242 = m.G0
	v244 = v242 - int32(16)
	m.G0 = v244
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v64)+16))
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v246)+12))
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v247)))
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v64)+8))
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v249)+12))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v250)))
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v64)+20))
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v252)+12))
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v253)))
	*(*int32)(unsafe.Add(mBase, uint32(v244)+12)) = v254
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v64)+24))
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v256)+12))
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v257)))
	*(*int32)(unsafe.Add(mBase, uint32(v244)+8)) = v258
	*(*int32)(unsafe.Add(mBase, uint32(v244)+4)) = v254
	*(*int32)(unsafe.Add(mBase, uint32(v244))) = v258
	v264 = F_list_make2_impl(m, v244+int32(4), v244)
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L16
	} else {
		goto L115
	}
L38:
	;
	v227 = int32(0)
	if l2 != 0 {
		v239 = v227
		goto L106
	} else {
		goto L107
	}
L39:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v64)+24))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v64)+28))
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	if l2 != 0 {
		v224 = v7
		goto L97
	} else {
		goto L98
	}
L40:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	if l2 != 0 {
		goto L31
	} else {
		goto L87
	}
L41:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	switch v96 {
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
	v84 = F_estimate_expression_value(m, l0, v64)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L16
	} else {
		goto L54
	}
L43:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+24)))
	if v78 != 0 {
		goto L48
	} else {
		goto L49
	}
L44:
	;
	v72 = float64(0.5)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v64)+28))
	if v73 != 0 {
		v487 = v72
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
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	if l2 != v76 {
		v487 = v72
		goto L28
	} else {
		goto L47
	}
L47:
	;
	goto L29
L48:
	;
	v487 = float64(0)
	goto L28
L49:
	;
	goto L50
L50:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v64)+20))
	if v82 != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v83 = float64(1)
	goto L53
L52:
	;
	v83 = float64(0)
	goto L53
L53:
	;
	v487 = v83
	goto L28
L54:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	if v86 != int32(7) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v487 = float64(0.5)
	goto L28
L56:
	;
	goto L57
L57:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+24)))
	if v90 != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v487 = float64(0)
	goto L28
L59:
	;
	goto L60
L60:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v84)+20))
	if v94 != 0 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v95 = float64(1)
	goto L63
L62:
	;
	v95 = float64(0)
	goto L63
L63:
	;
	v487 = v95
	goto L28
L64:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v64)+8))
	v108 = float64(0)
	v109 = m.G0
	v111 = v109 - int32(16)
	m.G0 = v111
	*(*int32)(unsafe.Add(mBase, uint32(v111)+12)) = int32(0)
	v115 = F_find_single_rel_for_clauses(m, l0, v107)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L16
	} else {
		goto L69
	}
L65:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v64)+8))
	v105 = F_clauselist_selectivity_ext(m, l0, v104, l2, l3, l4, l5)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L16
	} else {
		goto L68
	}
L66:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v64)+8))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v98)+12))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	v101 = F_clause_selectivity_ext(m, l0, v100, l2, l3, l4, l5)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L16
	} else {
		goto L67
	}
L67:
	;
	v487 = base.F64_sub(float64(1), v101)
	goto L28
L68:
	;
	v487 = v105
	goto L28
L69:
	;
	if l5 == int32(0) {
		v130 = v108
		goto L70
	} else {
		goto L71
	}
L70:
	;
	if v107 == int32(0) {
		v189 = v130
		goto L76
	} else {
		goto L77
	}
L71:
	;
	if v115 == int32(0) {
		v130 = v108
		goto L70
	} else {
		goto L72
	}
L72:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v115)+76))
	if v121 != 0 {
		v130 = v108
		goto L70
	} else {
		goto L73
	}
L73:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v115)+112))
	if v122 == int32(0) {
		v130 = v108
		goto L70
	} else {
		goto L74
	}
L74:
	;
	v128 = F_statext_clauselist_selectivity(m, l0, v107, l2, l3, l4, v115, v111+int32(12), int32(1))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L16
	} else {
		goto L75
	}
L75:
	;
	v130 = v128
	goto L70
L76:
	;
	m.G0 = v111 + int32(16)
	v487 = v189
	goto L28
L77:
	;
	v133 = int32(0)
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
	if v134 <= v133 {
		v189 = v130
		goto L76
	} else {
		goto L78
	}
L78:
	;
	v139 = v133
	v147 = int32(-1)
	v150 = v130
	goto L79
L79:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v107)+12))
	v156 = v147 + int32(1)
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v111)+12))
	v158 = F_bms_is_member(m, v156, v157)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L16
	} else {
		goto L81
	}
L80:
	;
	v189 = v171
	goto L76
L81:
	;
	if v158 == int32(0) {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v154+v139<<(uint(int32(2))%32))))
	v166 = F_clause_selectivity_ext(m, l0, v165, l2, l3, l4, l5)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L16
	} else {
		goto L85
	}
L83:
	;
	v171 = v150
	goto L84
L84:
	;
	v174 = v139 + int32(1)
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
	if v174 < v175 {
		v139 = v174
		v147 = v156
		v150 = v171
		goto L79
	} else {
		goto L86
	}
L85:
	;
	v171 = base.F64_sub(base.F64_add(v150, v166), base.F64_mul(v150, v166))
	goto L84
L86:
	;
	goto L80
L87:
	;
	if l4 == int32(0) {
		goto L31
	} else {
		goto L88
	}
L88:
	;
	if v68 != 0 {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v64)+28))
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v64)+24))
	v208 = F_join_selectivity(m, l0, v196, v206, v207, l3, l4)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L16
	} else {
		goto L96
	}
L90:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v68)+24))
	if int32(1) < v199 {
		goto L89
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	v202 = F_NumRelids(m, l0, v64)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L16
	} else {
		goto L94
	}
L93:
	;
	goto L31
L94:
	;
	if v202 < int32(2) {
		goto L31
	} else {
		goto L95
	}
L95:
	;
	goto L89
L96:
	;
	v439 = v208
	goto L30
L97:
	;
	v225 = F_function_selectivity(m, l0, v212, v211, v210, v224, l2, l3, l4)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L16
	} else {
		goto L105
	}
L98:
	;
	if l4 == int32(0) {
		v224 = v7
		goto L97
	} else {
		goto L99
	}
L99:
	;
	if v68 != 0 {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v68)+24))
	v218 = F_function_selectivity(m, l0, v212, v211, v210, base.B2i32(int32(1) < v215), l2, l3, l4)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L16
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	v220 = F_NumRelids(m, l0, v64)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L16
	} else {
		goto L104
	}
L103:
	;
	v487 = v218
	goto L28
L104:
	;
	v224 = base.B2i32(int32(1) < v220)
	goto L97
L105:
	;
	v487 = v225
	goto L28
L106:
	;
	v240 = F_scalararraysel(m, l0, v64, v239, l2, l3, l4)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L16
	} else {
		goto L114
	}
L107:
	;
	if l4 == int32(0) {
		v239 = v227
		goto L106
	} else {
		goto L108
	}
L108:
	;
	if v68 != 0 {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v68)+24))
	v233 = F_scalararraysel(m, l0, v64, base.B2i32(int32(1) < v230), l2, l3, l4)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L16
	} else {
		goto L112
	}
L110:
	;
	goto L111
L111:
	;
	v235 = F_NumRelids(m, l0, v64)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L16
	} else {
		goto L113
	}
L112:
	;
	v487 = v233
	goto L28
L113:
	;
	v239 = base.B2i32(int32(1) < v235)
	goto L106
L114:
	;
	v487 = v240
	goto L28
L115:
	;
	if l2 != 0 {
		goto L117
	} else {
		goto L118
	}
L116:
	;
	m.G0 = v244 + int32(16)
	v487 = v276
	goto L28
L117:
	;
	v274 = F_restriction_selectivity(m, l0, v251, v264, v248, l2)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L16
	} else {
		goto L123
	}
L118:
	;
	if l4 == int32(0) {
		goto L117
	} else {
		goto L119
	}
L119:
	;
	v268 = F_NumRelids(m, l0, v264)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L16
	} else {
		goto L120
	}
L120:
	;
	if v268 < int32(2) {
		goto L117
	} else {
		goto L121
	}
L121:
	;
	v272 = F_join_selectivity(m, l0, v251, v264, v248, l3, l4)
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L16
	} else {
		goto L122
	}
L122:
	;
	v276 = v272
	goto L116
L123:
	;
	v276 = v274
	goto L116
L124:
	;
	v487 = v282
	goto L28
L125:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v288)+88))
	if v294 != 0 {
		goto L128
	} else {
		goto L129
	}
L126:
	;
	m.G0 = v288 + int32(112)
	v487 = v413
	goto L28
L127:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v288)+88))
	if v399 != 0 {
		goto L165
	} else {
		goto L166
	}
L128:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v294)+16))
	v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v295)+22)))
	v298 = *(*float32)(unsafe.Add(mBase, uint32(v295+v296)+8))
	v299 = base.F64_promote_f32(v298)
	v305 = F_get_attstatsslot(m, v288+int32(44), v294, int32(1), int32(0), int32(3))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L16
	} else {
		goto L132
	}
L129:
	;
	goto L130
L130:
	;
	switch v284 {
	case 0, 3:
		goto L156
	case 1, 2:
		goto L158
	case 4:
		v413 = float64(0.005)
		goto L126
	case 5:
		goto L159
	default:
		goto L157
	}
L131:
	;
	switch v284 {
	case 0, 2:
		goto L152
	case 1, 3:
		goto L151
	case 4:
		v395 = v299
		goto L127
	case 5:
		goto L149
	default:
		goto L150
	}
L132:
	;
	if v305 == int32(0) {
		goto L131
	} else {
		goto L133
	}
L133:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v288)+68))
	if v309 <= int32(0) {
		goto L131
	} else {
		goto L134
	}
L134:
	;
	v312 = float64(1)
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v288)+64))
	v314 = *(*float32)(unsafe.Add(mBase, uint32(v313)))
	v315 = base.F64_promote_f32(v314)
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v288)+56))
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v319)))
	if v320 != 0 {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v321 = v315
	goto L137
L136:
	;
	v321 = base.F64_sub(base.F64_sub(v312, v315), v299)
	goto L137
L137:
	;
	v322 = base.F64_sub(v312, v321)
	v323 = base.F64_sub(v322, v299)
	switch v284 {
	case 0:
		goto L144
	case 1:
		goto L143
	case 2:
		goto L142
	case 3:
		goto L141
	case 4:
		v343 = v299
		goto L138
	case 5:
		goto L139
	default:
		goto L140
	}
L138:
	;
	F_free_attstatsslot(m, v288+int32(44))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L16
	} else {
		goto L148
	}
L139:
	;
	v343 = base.F64_sub(float64(1), v299)
	goto L138
L140:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L16
	} else {
		goto L145
	}
L141:
	;
	v343 = base.F64_sub(float64(1), v323)
	goto L138
L142:
	;
	v343 = v323
	goto L138
L143:
	;
	v343 = v322
	goto L138
L144:
	;
	v343 = v321
	goto L138
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v288)+16)) = v284
	F_errmsg_internal(m, int32(461367), v288+int32(16))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L16
	} else {
		goto L146
	}
L146:
	;
	F_errfinish(m, int32(471673), int32(1615), int32(291653))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L16
	} else {
		goto L147
	}
L147:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L148:
	;
	v395 = v343
	goto L127
L149:
	;
	v395 = base.F64_sub(float64(1), v299)
	goto L127
L150:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L16
	} else {
		goto L153
	}
L151:
	;
	v395 = base.F64_mul(base.F64_add(v299, float64(1)), float64(0.5))
	goto L127
L152:
	;
	v395 = base.F64_mul(base.F64_sub(float64(1), v299), float64(0.5))
	goto L127
L153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v288)+32)) = v284
	F_errmsg_internal(m, int32(461367), v288+int32(32))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L16
	} else {
		goto L154
	}
L154:
	;
	F_errfinish(m, int32(471673), int32(1652), int32(291653))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L16
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
	v392 = F_clause_selectivity(m, l0, v285, l2, l3, l4)
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L16
	} else {
		goto L164
	}
L157:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L16
	} else {
		goto L161
	}
L158:
	;
	v376 = F_clause_selectivity(m, l0, v285, l2, l3, l4)
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L16
	} else {
		goto L160
	}
L159:
	;
	v413 = float64(0.995)
	goto L126
L160:
	;
	v395 = base.F64_sub(float64(1), v376)
	goto L127
L161:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v288))) = v284
	F_errmsg_internal(m, int32(461367), v288)
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L16
	} else {
		goto L162
	}
L162:
	;
	F_errfinish(m, int32(471673), int32(1688), int32(291653))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L16
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
	v395 = v392
	goto L127
L165:
	;
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v288)+92))
	m.T0[v400].(func(*base.Module, int32))(m, v399)
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L16
	} else {
		goto L168
	}
L166:
	;
	goto L167
L167:
	;
	v403 = float64(0)
	if base.F64_lt(v395, v403) != 0 {
		v413 = v403
		goto L126
	} else {
		goto L169
	}
L168:
	;
	goto L167
L169:
	;
	if base.F64_gt(v395, float64(1)) == int32(0) {
		v413 = v395
		goto L126
	} else {
		goto L170
	}
L170:
	;
	v413 = float64(1)
	goto L126
L171:
	;
	v423 = *(*float64)(unsafe.Add(mBase, uint32(v421)+120))
	if base.F64_gt(v423, float64(0)) != 0 {
		goto L172
	} else {
		goto L173
	}
L172:
	;
	v428 = base.F64_div(float64(1), v423)
	goto L174
L173:
	;
	v428 = float64(0.5)
	goto L174
L174:
	;
	v487 = v428
	goto L28
L175:
	;
	v487 = v430
	goto L28
L176:
	;
	v487 = v433
	goto L28
L177:
	;
	v439 = v437
	goto L30
L178:
	;
	v444 = base.F64_sub(float64(1), v439)
	goto L180
L179:
	;
	v444 = v439
	goto L180
L180:
	;
	v487 = v444
	goto L28
L181:
	;
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v448)+8))
	if v452 == int32(0) {
		goto L183
	} else {
		goto L184
	}
L182:
	;
	m.G0 = v448 + int32(32)
	v487 = v471
	goto L28
L183:
	;
	v471 = float64(0.5)
	goto L182
L184:
	;
	goto L185
L185:
	;
	v457 = int32(0)
	v458 = int32(1)
	v462 = F_var_eq_const(m, v448, int32(91), v457, v458, v457, v458, v457)
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L16
	} else {
		goto L186
	}
L186:
	;
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v448)+8))
	if v464 == int32(0) {
		v471 = v462
		goto L182
	} else {
		goto L187
	}
L187:
	;
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v448)+12))
	m.T0[v467].(func(*base.Module, int32))(m, v464)
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L16
	} else {
		goto L188
	}
L188:
	;
	v471 = v462
	goto L182
L189:
	;
	if l3 != 0 {
		goto L190
	} else {
		goto L191
	}
L190:
	;
	v495 = int32(88)
	goto L192
L191:
	;
	v495 = int32(80)
	goto L192
L192:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v68+v495))) = v487
	v510 = v487
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
	F___gettimeofday(m, v5)
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
	var v15 int64
	_ = v15
	var v19 int32
	_ = v19
	var v20 int64
	_ = v20
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+64))
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+48)))
	switch v12 & int32(240) {
	case 0:
		v15 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
		*(*int64)(unsafe.Add(mBase, uint32(v8))) = v15
		F_appendStringInfo(m, l0, int32(409484), v8)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return
		} else {
			m.G0 = v8 + int32(32)
			return
		}
	default:
		m.G0 = v8 + int32(32)
		return
	case 16:
		v20 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = v21
		*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = v20
		F_appendStringInfo(m, l0, int32(39658), v8+int32(16))
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
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
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	v4 = m.G0
	v6 = v4 - int32(48)
	m.G0 = v6
	v9 = *(*int32)(unsafe.Add(mBase, _consts[415]))
	if l0 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L18
	} else {
		goto L33
	}
L2:
	;
	v80 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L18
	} else {
		goto L26
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
		goto L25
	}
L6:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+48))
	v12 = v10 + int32(4)
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v16 == int32(0) {
		v35 = v15
		v36 = v16
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
	v64 = m.ExcPending
	if v64 != 0 {
		goto L18
	} else {
		goto L22
	}
L9:
	;
	if v36-v35 == int32(0) {
		goto L2
	} else {
		goto L17
	}
L10:
	;
	goto L9
L11:
	;
	if v15 != v16 {
		v35 = v15
		v36 = v16
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v20 = v12
	v21 = l0
	goto L13
L13:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)))
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+1)))
	if v25 == int32(0) {
		v35 = v24
		v36 = v25
		goto L10
	} else {
		goto L15
	}
L14:
	;
	v35 = v24
	v36 = v25
	goto L10
L15:
	;
	v28 = int32(1)
	if v24 == v25 {
		v20 = v20 + v28
		v21 = v21 + v28
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	return
L19:
	;
	v45 = *(*int32)(unsafe.Add(mBase, _consts[415]))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+32)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v6)+36)) = v46 + int32(4)
	F_errmsg_internal(m, int32(426249), v6+int32(32))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	F_errfinish(m, int32(472726), int32(493), int32(292183))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L18
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
	*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = l0
	F_errmsg_internal(m, int32(431985), v6+int32(16))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L18
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(472726), int32(497), int32(292183))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L18
	} else {
		goto L24
	}
L24:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L25:
	;
	goto L2
L26:
	;
	if v80 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v83 = *(*int32)(unsafe.Add(mBase, _consts[415]))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = v84 + int32(4)
	F_errmsg_internal(m, int32(173725), v6)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L18
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v97 = *(*int32)(unsafe.Add(mBase, _consts[415]))
	F_sequence_close(m, v97, int32(0))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L18
	} else {
		goto L32
	}
L30:
	;
	F_errfinish(m, int32(472726), int32(505), int32(292183))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L18
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	*(*int32)(unsafe.Add(mBase, _consts[415])) = int32(0)
	m.G0 = v6 + int32(48)
	return
L33:
	;
	F_errmsg_internal(m, int32(343729), int32(0))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L18
	} else {
		goto L34
	}
L34:
	;
	F_errfinish(m, int32(472726), int32(501), int32(292183))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L18
	} else {
		goto L35
	}
L35:
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
	var v17 int32
	_ = v17
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
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3))))
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4))))
	if v17 == int32(0) {
		v36 = v16
		v37 = v17
		goto L8
	} else {
		goto L9
	}
L7:
	;
	return v37 - v36
L8:
	;
	goto L7
L9:
	;
	if v16 != v17 {
		v36 = v16
		v37 = v17
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v21 = v4
	v22 = v3
	goto L11
L11:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+1)))
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)))
	if v26 == int32(0) {
		v36 = v25
		v37 = v26
		goto L8
	} else {
		goto L13
	}
L12:
	;
	v36 = v25
	v37 = v26
	goto L8
L13:
	;
	v29 = int32(1)
	if v25 == v26 {
		v21 = v21 + v29
		v22 = v22 + v29
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
}
func F_colorcomplement(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
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
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v158 int32
	_ = v158
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v197 int32
	_ = v197
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v217 int32
	_ = v217
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v244 int32
	_ = v244
	var v255 int32
	_ = v255
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	if v13 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v185 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	if v185 != 0 {
		goto L55
	} else {
		goto L56
	}
L2:
	;
	v20 = v13
	goto L5
L3:
	;
	v66 = v12
	goto L4
L4:
	;
	v68 = int32(24)
	v70 = v12 + v11*v68
	if base.Ui32(v70+v68) <= base.Ui32(v66) {
		goto L18
	} else {
		goto L19
	}
L5:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	if v24 == int32(112) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v34 = v13
	goto L12
L7:
	;
	v27 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20)+4)))
	if v27 == int32(65534) {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
	if v30 != 0 {
		v20 = v30
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
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	if v41 == int32(112) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v66 = v57
	goto L4
L14:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v45 = int32(*(*int16)(unsafe.Add(mBase, uint32(v34)+4)))
	v50 = v44 + v45*int32(24) + int32(20)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	*(*int32)(unsafe.Add(mBase, uint32(v50))) = v51 | int32(4)
	goto L16
L15:
	;
	goto L16
L16:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v34)+16))
	if v56 != 0 {
		v34 = v56
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
	v82 = int32(0)
	v83 = v66
	goto L20
L20:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+12))
	if v86 != 0 {
		goto L18
	} else {
		goto L22
	}
L21:
	;
	goto L18
L22:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v83)+20))
	if v87&int32(4) != 0 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	if base.Ui32(v83) < base.Ui32(v70) {
		v82 = v82 + int32(1)
		v83 = v83 + int32(24)
		goto L20
	} else {
		goto L54
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v83)+20)) = v87 & int32(-5)
	goto L23
L25:
	;
	goto L26
L26:
	;
	if v87&int32(3) != 0 {
		goto L23
	} else {
		goto L27
	}
L27:
	;
	v96 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	if v96 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	goto L30
L30:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l5)+8))
	if v99 <= v100 {
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
	F_createarc(m, l0, l2, base.I32_extend16_s(v82), l4, l5)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L31
	} else {
		goto L53
	}
L34:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
	if v102 == int32(0) {
		goto L33
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l5)+16))
	if v124 == int32(0) {
		goto L33
	} else {
		goto L45
	}
L37:
	;
	v108 = v102
	goto L38
L38:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v108)+12))
	if v115 != l5 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	goto L33
L40:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v108)+16))
	if v123 != 0 {
		v108 = v123
		goto L38
	} else {
		goto L44
	}
L41:
	;
	v117 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v108)+4)))
	if v117 != v82&int32(65535) {
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
	if v121 == l2 {
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
	v130 = v124
	goto L46
L46:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v130)+8))
	if v137 != l4 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	goto L33
L48:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v130)+24))
	if v145 != 0 {
		v130 = v145
		goto L46
	} else {
		goto L52
	}
L49:
	;
	v139 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v130)+4)))
	if v139 != v82&int32(65535) {
		goto L48
	} else {
		goto L50
	}
L50:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
	if v143 == l2 {
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
	v187 = m.ExcPending
	if v187 != 0 {
		goto L31
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v189 = *(*int32)(unsafe.Add(mBase, uint32(l5)+8))
	if v188 <= v189 {
		goto L61
	} else {
		goto L62
	}
L58:
	;
	goto L57
L59:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v255 | int32(4)
	return
L60:
	;
	F_createarc(m, l0, int32(120), int32(0), l4, l5)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L31
	} else {
		goto L80
	}
L61:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
	if v191 == int32(0) {
		goto L60
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l5)+16))
	if v211 == int32(0) {
		goto L60
	} else {
		goto L72
	}
L64:
	;
	v197 = v191
	goto L65
L65:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v197)+12))
	if v204 != l5 {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	goto L60
L67:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v197)+16))
	if v210 != 0 {
		v197 = v210
		goto L65
	} else {
		goto L71
	}
L68:
	;
	v206 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v197)+4)))
	if v206 != 0 {
		goto L67
	} else {
		goto L69
	}
L69:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v197)))
	if v207 == int32(120) {
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
	v217 = v211
	goto L73
L73:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v217)+8))
	if v224 != l4 {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	goto L60
L75:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v217)+24))
	if v230 != 0 {
		v217 = v230
		goto L73
	} else {
		goto L79
	}
L76:
	;
	v226 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v217)+4)))
	if v226 != 0 {
		goto L75
	} else {
		goto L77
	}
L77:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v217)))
	if v227 == int32(120) {
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
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
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
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
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
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v114 int32
	_ = v114
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v10 = v6 | v7<<(uint(int32(8))%32)
	if v10 <= int32(24907) {
		if v10 <= int32(24099) {
			v15 = int32(1)
			switch v10 - int32(9292) {
			case 0, 18:
				v114 = int32(3)
				return v114
			case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17:
				v114 = v15
				return v114
			default:
				if v10 == int32(9252) {
					v69 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
					v70 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)))
					if v69 == v70 {
						v72 = int32(2)
					} else {
						v72 = int32(1)
					}
					return v72
				} else {
					if v10 == int32(9330) {
						v114 = int32(3)
					} else {
						v114 = v15
					}
					return v114
				}
			}
		} else {
			v22 = int32(1)
			switch v10 - int32(24140) {
			case 0, 21:
				v114 = int32(3)
				return v114
			case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 19, 20:
				v114 = v22
				return v114
			case 18:
				v69 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
				v70 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)))
				if v69 == v70 {
					v72 = int32(2)
				} else {
					v72 = int32(1)
				}
				return v72
			default:
				if v10 != int32(24100) {
					v114 = v22
				} else {
					v114 = int32(3)
				}
				return v114
			}
		}
	} else {
		v25 = int32(1)
		switch v10 - int32(24908) {
		case 0, 18, 38:
			v114 = int32(3)
			return v114
		case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 19, 20, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 37:
			v114 = v25
			return v114
		case 21:
			v74 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
			v75 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)))
			if v74 == v75 {
				return int32(2)
			} else {
				if v74 == int32(65534) {
					v81 = int32(2)
					v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
					v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+20))
					v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83+base.I32_extend16_s(v75)*int32(24))+20)))
					if v88&v81 == int32(0) {
						v114 = v81
						return v114
					} else {
						return int32(1)
					}
				} else {
					if v75 != int32(65534) {
						return int32(1)
					} else {
						v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
						v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)+20))
						v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98+base.I32_extend16_s(v74)*int32(24))+20)))
						if v103&int32(2) != 0 {
							v106 = int32(1)
						} else {
							v106 = int32(4)
						}
						return v106
					}
				}
			}
		case 36:
			v32 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
			v33 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)))
			if v32 == v33 {
				return int32(2)
			} else {
				if v32 == int32(65534) {
					v39 = int32(2)
					v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
					v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+20))
					v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41+base.I32_extend16_s(v33)*int32(24))+20)))
					if v46&v39 == int32(0) {
						v114 = v39
						return v114
					} else {
						return int32(1)
					}
				} else {
					if v33 != int32(65534) {
						return int32(1)
					} else {
						v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
						v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+20))
						v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56+base.I32_extend16_s(v32)*int32(24))+20)))
						if v61&int32(2) != 0 {
							v64 = int32(1)
						} else {
							v64 = int32(4)
						}
						return v64
					}
				}
			}
		default:
			switch v10 - int32(29260) {
			case 0, 21:
				v114 = int32(3)
				return v114
			case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 37:
				v114 = v25
				return v114
			case 36:
				v32 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
				v33 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)))
				if v32 == v33 {
					return int32(2)
				} else {
					if v32 == int32(65534) {
						v39 = int32(2)
						v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
						v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+20))
						v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41+base.I32_extend16_s(v33)*int32(24))+20)))
						if v46&v39 == int32(0) {
							v114 = v39
							return v114
						} else {
							return int32(1)
						}
					} else {
						if v33 != int32(65534) {
							return int32(1)
						} else {
							v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
							v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+20))
							v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56+base.I32_extend16_s(v32)*int32(24))+20)))
							if v61&int32(2) != 0 {
								v64 = int32(1)
							} else {
								v64 = int32(4)
							}
							return v64
						}
					}
				}
			case 38:
				v74 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
				v75 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)))
				if v74 == v75 {
					return int32(2)
				} else {
					if v74 == int32(65534) {
						v81 = int32(2)
						v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
						v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+20))
						v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83+base.I32_extend16_s(v75)*int32(24))+20)))
						if v88&v81 == int32(0) {
							v114 = v81
							return v114
						} else {
							return int32(1)
						}
					} else {
						if v75 != int32(65534) {
							return int32(1)
						} else {
							v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
							v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)+20))
							v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98+base.I32_extend16_s(v74)*int32(24))+20)))
							if v103&int32(2) != 0 {
								v106 = int32(1)
							} else {
								v106 = int32(4)
							}
							return v106
						}
					}
				}
			default:
				if v10 == int32(29220) {
					v114 = int32(3)
				} else {
					v114 = v25
				}
				return v114
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
	var v28 int32
	_ = v28
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
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v138 int32
	_ = v138
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v158 int32
	_ = v158
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v300 int32
	_ = v300
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v315 int32
	_ = v315
	var v322 int32
	_ = v322
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v335 int32
	_ = v335
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v392 int32
	_ = v392
	var v404 int32
	_ = v404
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v469 int32
	_ = v469
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
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
	var v535 int32
	_ = v535
	var v537 int32
	_ = v537
	var v541 int32
	_ = v541
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v557 int32
	_ = v557
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v592 int32
	_ = v592
	var v598 int32
	_ = v598
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v609 int32
	_ = v609
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v621 int32
	_ = v621
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
	*(*uint16)(unsafe.Add(mBase, uint32(l2)+14)) = uint16(v621)
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
	v404 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v404) {
		goto L118
	} else {
		goto L119
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
	v28 = v5
	goto L9
L8:
	;
	v257 = l2 + v247
	v258 = l2 + v248
	v259 = v251 - v248
	if v257 == v258 {
		goto L72
	} else {
		goto L73
	}
L9:
	;
	v38 = l0 + v28*int32(6)
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
	if l1 <= v28 {
		goto L15
	} else {
		goto L16
	}
L11:
	;
	v43 = v26 - v39
	v45 = v28 + int32(1)
	if v45 != v21 {
		v26 = v43
		v28 = v45
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
	v247 = v43
	v248 = v26
	v251 = v26
	goto L8
L15:
	;
	v247 = v26
	v248 = v41
	v251 = v41
	goto L8
L16:
	;
	goto L17
L17:
	;
	v53 = v26
	v54 = v41
	v55 = v28
	v57 = v41
	goto L18
L18:
	;
	v65 = l0 + v55*int32(6)
	v66 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v65))))
	v75 = (v66+int32(1))&int32(65535)<<(uint(int32(2))%32) + (l2 + int32(24)) - int32(4)
	v76 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v65)+4)))
	v77 = int32(*(*int16)(unsafe.Add(mBase, uint32(v65)+2)))
	if v76+v77 == v54 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v247 = v236
	v248 = v230
	v251 = v231
	goto L8
L20:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	v236 = v53 - v232
	*(*int32)(unsafe.Add(mBase, uint32(v75))) = v233&int32(-32768) | v236&int32(32767)
	v242 = v55 + int32(1)
	if v242 != l1 {
		v53 = v236
		v54 = v230
		v55 = v242
		v57 = v231
		goto L18
	} else {
		goto L70
	}
L21:
	;
	v230 = v77
	v231 = v57
	v232 = v76
	goto L20
L22:
	;
	goto L23
L23:
	;
	v80 = l2 + v53
	v81 = l2 + v54
	v82 = v57 - v54
	if v80 == v81 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v227 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v65)+4)))
	v228 = int32(*(*int16)(unsafe.Add(mBase, uint32(v65)+2)))
	v230 = v228
	v231 = v227 + v228
	v232 = v227
	goto L20
L25:
	;
	goto L24
L26:
	;
	v86 = v80 + v82
	if base.Ui32(v81-v86) <= base.Ui32(int32(0)-v82<<(uint(int32(1))%32)) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v93 = F___memcpy(m, v80, v81, v82)
	mBase = m.M
	goto L24
L28:
	;
	goto L29
L29:
	;
	v96 = (v80 ^ v81) & int32(3)
	if base.Ui32(v80) < base.Ui32(v81) {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	if v198 == int32(0) {
		goto L25
	} else {
		goto L66
	}
L31:
	;
	if base.Ui32(v176) <= base.Ui32(int32(3)) {
		v197 = v175
		v198 = v176
		v199 = v177
		goto L30
	} else {
		goto L62
	}
L32:
	;
	if v96 != 0 {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	goto L34
L34:
	;
	if v96 != 0 {
		v158 = v82
		goto L45
	} else {
		goto L46
	}
L35:
	;
	v197 = v81
	v198 = v82
	v199 = v80
	goto L30
L36:
	;
	goto L37
L37:
	;
	if v80&int32(3) == int32(0) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v175 = v81
	v176 = v82
	v177 = v80
	goto L31
L39:
	;
	goto L40
L40:
	;
	v103 = v81
	v104 = v82
	v105 = v80
	goto L41
L41:
	;
	if v104 == int32(0) {
		goto L25
	} else {
		goto L43
	}
L42:
	;
	v175 = v112
	v176 = v114
	v177 = v116
	goto L31
L43:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103))))
	*(*uint8)(unsafe.Add(mBase, uint32(v105))) = uint8(v109)
	v111 = int32(1)
	v112 = v103 + v111
	v114 = v104 - v111
	v116 = v105 + v111
	if v116&int32(3) != 0 {
		v103 = v112
		v104 = v114
		v105 = v116
		goto L41
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	if v158 == int32(0) {
		goto L25
	} else {
		goto L58
	}
L46:
	;
	if v86&int32(3) != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v123 = v82
	goto L50
L48:
	;
	v138 = v82
	goto L49
L49:
	;
	if base.Ui32(v138) <= base.Ui32(int32(3)) {
		v158 = v138
		goto L45
	} else {
		goto L54
	}
L50:
	;
	if v123 == int32(0) {
		goto L25
	} else {
		goto L52
	}
L51:
	;
	v138 = v129
	goto L49
L52:
	;
	v129 = v123 - int32(1)
	v130 = v80 + v129
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81+v129))))
	*(*uint8)(unsafe.Add(mBase, uint32(v130))) = uint8(v132)
	if v130&int32(3) != 0 {
		v123 = v129
		goto L50
	} else {
		goto L53
	}
L53:
	;
	goto L51
L54:
	;
	v145 = v138
	goto L55
L55:
	;
	v149 = v145 - int32(4)
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v81+v149)))
	*(*int32)(unsafe.Add(mBase, uint32(v80+v149))) = v152
	if base.Ui32(int32(3)) < base.Ui32(v149) {
		v145 = v149
		goto L55
	} else {
		goto L57
	}
L56:
	;
	v158 = v149
	goto L45
L57:
	;
	goto L56
L58:
	;
	v165 = v158
	goto L59
L59:
	;
	v169 = v165 - int32(1)
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81+v169))))
	*(*uint8)(unsafe.Add(mBase, uint32(v80+v169))) = uint8(v172)
	if v169 != 0 {
		v165 = v169
		goto L59
	} else {
		goto L61
	}
L60:
	;
	goto L25
L61:
	;
	goto L60
L62:
	;
	v182 = v175
	v183 = v176
	v184 = v177
	goto L63
L63:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v182)))
	*(*int32)(unsafe.Add(mBase, uint32(v184))) = v186
	v188 = int32(4)
	v189 = v182 + v188
	v191 = v184 + v188
	v193 = v183 - v188
	if base.Ui32(int32(3)) < base.Ui32(v193) {
		v182 = v189
		v183 = v193
		v184 = v191
		goto L63
	} else {
		goto L65
	}
L64:
	;
	v197 = v189
	v198 = v193
	v199 = v191
	goto L30
L65:
	;
	goto L64
L66:
	;
	v204 = v197
	v205 = v198
	v206 = v199
	goto L67
L67:
	;
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v204))))
	*(*uint8)(unsafe.Add(mBase, uint32(v206))) = uint8(v208)
	v210 = int32(1)
	v215 = v205 - v210
	if v215 != 0 {
		v204 = v204 + v210
		v205 = v215
		v206 = v206 + v210
		goto L67
	} else {
		goto L69
	}
L68:
	;
	goto L25
L69:
	;
	goto L68
L70:
	;
	goto L19
L71:
	;
	v621 = v247
	goto L1
L72:
	;
	goto L71
L73:
	;
	v263 = v257 + v259
	if base.Ui32(v258-v263) <= base.Ui32(int32(0)-v259<<(uint(int32(1))%32)) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v270 = F___memcpy(m, v257, v258, v259)
	mBase = m.M
	goto L71
L75:
	;
	goto L76
L76:
	;
	v273 = (v257 ^ v258) & int32(3)
	if base.Ui32(v257) < base.Ui32(v258) {
		goto L79
	} else {
		goto L80
	}
L77:
	;
	if v375 == int32(0) {
		goto L72
	} else {
		goto L113
	}
L78:
	;
	if base.Ui32(v353) <= base.Ui32(int32(3)) {
		v374 = v352
		v375 = v353
		v376 = v354
		goto L77
	} else {
		goto L109
	}
L79:
	;
	if v273 != 0 {
		goto L82
	} else {
		goto L83
	}
L80:
	;
	goto L81
L81:
	;
	if v273 != 0 {
		v335 = v259
		goto L92
	} else {
		goto L93
	}
L82:
	;
	v374 = v258
	v375 = v259
	v376 = v257
	goto L77
L83:
	;
	goto L84
L84:
	;
	if v257&int32(3) == int32(0) {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v352 = v258
	v353 = v259
	v354 = v257
	goto L78
L86:
	;
	goto L87
L87:
	;
	v280 = v258
	v281 = v259
	v282 = v257
	goto L88
L88:
	;
	if v281 == int32(0) {
		goto L72
	} else {
		goto L90
	}
L89:
	;
	v352 = v289
	v353 = v291
	v354 = v293
	goto L78
L90:
	;
	v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v280))))
	*(*uint8)(unsafe.Add(mBase, uint32(v282))) = uint8(v286)
	v288 = int32(1)
	v289 = v280 + v288
	v291 = v281 - v288
	v293 = v282 + v288
	if v293&int32(3) != 0 {
		v280 = v289
		v281 = v291
		v282 = v293
		goto L88
	} else {
		goto L91
	}
L91:
	;
	goto L89
L92:
	;
	if v335 == int32(0) {
		goto L72
	} else {
		goto L105
	}
L93:
	;
	if v263&int32(3) != 0 {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v300 = v259
	goto L97
L95:
	;
	v315 = v259
	goto L96
L96:
	;
	if base.Ui32(v315) <= base.Ui32(int32(3)) {
		v335 = v315
		goto L92
	} else {
		goto L101
	}
L97:
	;
	if v300 == int32(0) {
		goto L72
	} else {
		goto L99
	}
L98:
	;
	v315 = v306
	goto L96
L99:
	;
	v306 = v300 - int32(1)
	v307 = v257 + v306
	v309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258+v306))))
	*(*uint8)(unsafe.Add(mBase, uint32(v307))) = uint8(v309)
	if v307&int32(3) != 0 {
		v300 = v306
		goto L97
	} else {
		goto L100
	}
L100:
	;
	goto L98
L101:
	;
	v322 = v315
	goto L102
L102:
	;
	v326 = v322 - int32(4)
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v258+v326)))
	*(*int32)(unsafe.Add(mBase, uint32(v257+v326))) = v329
	if base.Ui32(int32(3)) < base.Ui32(v326) {
		v322 = v326
		goto L102
	} else {
		goto L104
	}
L103:
	;
	v335 = v326
	goto L92
L104:
	;
	goto L103
L105:
	;
	v342 = v335
	goto L106
L106:
	;
	v346 = v342 - int32(1)
	v349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258+v346))))
	*(*uint8)(unsafe.Add(mBase, uint32(v257+v346))) = uint8(v349)
	if v346 != 0 {
		v342 = v346
		goto L106
	} else {
		goto L108
	}
L107:
	;
	goto L72
L108:
	;
	goto L107
L109:
	;
	v359 = v352
	v360 = v353
	v361 = v354
	goto L110
L110:
	;
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v359)))
	*(*int32)(unsafe.Add(mBase, uint32(v361))) = v363
	v365 = int32(4)
	v366 = v359 + v365
	v368 = v361 + v365
	v370 = v360 - v365
	if base.Ui32(int32(3)) < base.Ui32(v370) {
		v359 = v366
		v360 = v370
		v361 = v368
		goto L110
	} else {
		goto L112
	}
L111:
	;
	v374 = v366
	v375 = v370
	v376 = v368
	goto L77
L112:
	;
	goto L111
L113:
	;
	v381 = v374
	v382 = v375
	v383 = v376
	goto L114
L114:
	;
	v385 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v381))))
	*(*uint8)(unsafe.Add(mBase, uint32(v383))) = uint8(v385)
	v387 = int32(1)
	v392 = v382 - v387
	if v392 != 0 {
		v381 = v381 + v387
		v382 = v392
		v383 = v383 + v387
		goto L114
	} else {
		goto L116
	}
L115:
	;
	goto L72
L116:
	;
	goto L115
L117:
	;
	if l1 <= v537 {
		goto L161
	} else {
		goto L162
	}
L118:
	;
	v414 = int32(base.Ui32(v404+int32(262120))>>(uint(int32(4))%32)) & int32(16383)
	goto L120
L119:
	;
	v414 = int32(0)
	goto L120
L120:
	;
	if l1 < v414 {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v416 = int32(1)
	if l1 <= v416 {
		goto L124
	} else {
		goto L125
	}
L122:
	;
	goto L123
L123:
	;
	v495 = int32(1)
	if base.Ui32(l1) <= base.Ui32(v495) {
		goto L148
	} else {
		goto L149
	}
L124:
	;
	v419 = v416
	goto L126
L125:
	;
	v419 = l1
	goto L126
L126:
	;
	v422 = int32(0)
	if int32(2) <= l1 {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v431 = v422
	v433 = int32(0)
	goto L130
L128:
	;
	v469 = v422
	goto L129
L129:
	;
	if v419&int32(1) != 0 {
		goto L141
	} else {
		goto L142
	}
L130:
	;
	v443 = l0 + v431*int32(6)
	v444 = int32(*(*int16)(unsafe.Add(mBase, uint32(v443)+2)))
	v447 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v443)+4)))
	if v447 != 0 {
		goto L133
	} else {
		goto L134
	}
L131:
	;
	v469 = v462
	goto L129
L132:
	;
	v454 = l0 + (v431|int32(1))*int32(6)
	v455 = int32(*(*int16)(unsafe.Add(mBase, uint32(v454)+2)))
	v458 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v454)+4)))
	if v458 != 0 {
		goto L137
	} else {
		goto L138
	}
L133:
	;
	v448 = F__emscripten_memcpy_bulkmem(m, v16+v444, l2+v444, v447)
	mBase = m.M
	goto L135
L134:
	;
	goto L135
L135:
	;
	goto L132
L136:
	;
	v461 = int32(2)
	v462 = v431 + v461
	v464 = v433 + v461
	if v464 != v419&int32(2147483646) {
		v431 = v462
		v433 = v464
		goto L130
	} else {
		goto L140
	}
L137:
	;
	v459 = F__emscripten_memcpy_bulkmem(m, v16+v455, l2+v455, v458)
	mBase = m.M
	goto L139
L138:
	;
	goto L139
L139:
	;
	goto L136
L140:
	;
	goto L131
L141:
	;
	v481 = l0 + v469*int32(6)
	v482 = int32(*(*int16)(unsafe.Add(mBase, uint32(v481)+2)))
	v485 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v481)+4)))
	if v485 != 0 {
		goto L145
	} else {
		goto L146
	}
L142:
	;
	goto L143
L143:
	;
	v490 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	v491 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+2)))
	v493 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+16)))
	v535 = v493
	v537 = int32(0)
	v541 = v490 + v491
	goto L117
L144:
	;
	goto L143
L145:
	;
	v486 = F__emscripten_memcpy_bulkmem(m, v16+v482, l2+v482, v485)
	mBase = m.M
	goto L147
L146:
	;
	goto L147
L147:
	;
	goto L144
L148:
	;
	v498 = v495
	goto L150
L149:
	;
	v498 = l1
	goto L150
L150:
	;
	v499 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+16)))
	v503 = v499
	v505 = v5
	goto L152
L151:
	;
	v526 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+14)))
	v529 = v524 - v526
	if v529 != 0 {
		goto L157
	} else {
		goto L158
	}
L152:
	;
	v515 = l0 + v505*int32(6)
	v516 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v515)+4)))
	v517 = int32(*(*int16)(unsafe.Add(mBase, uint32(v515)+2)))
	v518 = v516 + v517
	if v503 != v518 {
		v524 = v503
		v525 = v505
		goto L151
	} else {
		goto L154
	}
L153:
	;
	v524 = v520
	v525 = v498
	goto L151
L154:
	;
	v520 = v503 - v516
	v522 = v505 + int32(1)
	if v522 != v498 {
		v503 = v520
		v505 = v522
		goto L152
	} else {
		goto L155
	}
L155:
	;
	goto L153
L156:
	;
	v535 = v524
	v537 = v525
	v541 = v518
	goto L117
L157:
	;
	v530 = F__emscripten_memcpy_bulkmem(m, v16+v526, l2+v526, v529)
	mBase = m.M
	goto L159
L158:
	;
	goto L159
L159:
	;
	goto L156
L160:
	;
	v615 = v609 - v604
	if v615 != 0 {
		goto L176
	} else {
		goto L177
	}
L161:
	;
	v603 = v535
	v604 = v541
	v609 = v541
	goto L160
L162:
	;
	goto L163
L163:
	;
	v551 = v535
	v552 = v541
	v553 = v537
	v557 = v541
	goto L164
L164:
	;
	v563 = l0 + v553*int32(6)
	v564 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v563))))
	v573 = (v564+int32(1))&int32(65535)<<(uint(int32(2))%32) + (l2 + int32(24)) - int32(4)
	v574 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v563)+4)))
	v575 = int32(*(*int16)(unsafe.Add(mBase, uint32(v563)+2)))
	if v574+v575 == v552 {
		goto L167
	} else {
		goto L168
	}
L165:
	;
	v603 = v592
	v604 = v586
	v609 = v588
	goto L160
L166:
	;
	v589 = *(*int32)(unsafe.Add(mBase, uint32(v573)))
	v592 = v551 - v587
	*(*int32)(unsafe.Add(mBase, uint32(v573))) = v589&int32(-32768) | v592&int32(32767)
	v598 = v553 + int32(1)
	if v598 != l1 {
		v551 = v592
		v552 = v586
		v553 = v598
		v557 = v588
		goto L164
	} else {
		goto L174
	}
L167:
	;
	v586 = v575
	v587 = v574
	v588 = v557
	goto L166
L168:
	;
	goto L169
L169:
	;
	v580 = v557 - v552
	if v580 != 0 {
		goto L171
	} else {
		goto L172
	}
L170:
	;
	v583 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v563)+4)))
	v584 = int32(*(*int16)(unsafe.Add(mBase, uint32(v563)+2)))
	v586 = v584
	v587 = v583
	v588 = v583 + v584
	goto L166
L171:
	;
	v581 = F__emscripten_memcpy_bulkmem(m, l2+v551, v552+v16, v580)
	mBase = m.M
	goto L173
L172:
	;
	goto L173
L173:
	;
	goto L170
L174:
	;
	goto L165
L175:
	;
	v621 = v603
	goto L1
L176:
	;
	v616 = F__emscripten_memcpy_bulkmem(m, l2+v603, v604+v16, v615)
	mBase = m.M
	goto L178
L177:
	;
	goto L178
L178:
	;
	goto L175
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
	*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = l0
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+44)) = uint8(v4)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l1
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
				F_errmsg_internal(m, int32(508464), v7)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(469797), int32(58), int32(274291))
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
	var v6 int32
	_ = v6
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+64)))
	if base.Ui32(v6-int32(3)) < base.Ui32(int32(-2)) {
		return int32(-102)
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
		v16 = m.Env.Pgmem_deflate_create(m, base.B2i32(v6 == int32(1)), v15)
		mBase = m.M
		if v16 <= int32(0) {
			return int32(-105)
		} else {
			v22 = F_palloc0(m, int32(8200))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v22)+4)) = int32(8192)
				v28 = int32(4442628)
				v29 = *(*int32)(unsafe.Add(mBase, _consts[170]))
				F_ResourceOwnerEnlarge(m, v29)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int32(0)
				} else {
					v33 = *(*int32)(unsafe.Add(mBase, _consts[12]))
					v35 = F_MemoryContextAlloc(m, v33, int32(8))
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v35))) = v16
						v38 = *(*int32)(unsafe.Add(mBase, _consts[170]))
						*(*int32)(unsafe.Add(mBase, uint32(v35)+4)) = v38
						F_ResourceOwnerRemember(m, v38, v35, int32(4335212))
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v22))) = v35
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v22
							return int32(8192)
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
	var v32 float64
	_ = v32
	var v36 float64
	_ = v36
	v5 = *(*float64)(unsafe.Add(mBase, uint32(l0)+32))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = base.F64_convert_i32_s(v6)
	v9 = int32(*(*uint8)(unsafe.Add(mBase, _consts[586])))
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
	if base.F64_gt(v24, v23) != 0 {
		v36 = v23
	} else {
		if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v24)&int64(9223372036854775807)) {
			v36 = v23
		} else {
			v32 = float64(1)
			if base.F64_le(v24, v32) != 0 {
				v36 = v32
			} else {
				v36 = base.F64_nearest(v24)
			}
		}
	}
	return v36
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
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 float64
	_ = v127
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v143 float64
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
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
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v232 int32
	_ = v232
	var v254 int32
	_ = v254
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v346 int32
	_ = v346
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
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v369 int32
	_ = v369
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v417 int32
	_ = v417
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v506 int32
	_ = v506
	var v510 int32
	_ = v510
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v543 int32
	_ = v543
	var v548 int32
	_ = v548
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v576 int32
	_ = v576
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v590 int32
	_ = v590
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v640 int32
	_ = v640
	var v658 float64
	_ = v658
	var v659 int32
	_ = v659
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v680 int32
	_ = v680
	var v684 int32
	_ = v684
	var v686 int32
	_ = v686
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v704 int32
	_ = v704
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v711 float64
	_ = v711
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v722 float64
	_ = v722
	var v727 int32
	_ = v727
	var v730 int32
	_ = v730
	var v736 float64
	_ = v736
	var v737 float64
	_ = v737
	var v740 float64
	_ = v740
	var v742 float64
	_ = v742
	var v746 int32
	_ = v746
	var v748 int32
	_ = v748
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v763 float64
	_ = v763
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v796 int32
	_ = v796
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v807 float64
	_ = v807
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v819 float64
	_ = v819
	var v821 float64
	_ = v821
	var v825 int32
	_ = v825
	var v827 int32
	_ = v827
	var v831 int32
	_ = v831
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v846 int32
	_ = v846
	var v871 int32
	_ = v871
	var v873 int32
	_ = v873
	var v878 int32
	_ = v878
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v885 int32
	_ = v885
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v900 int32
	_ = v900
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v940 int32
	_ = v940
	var v993 int32
	_ = v993
	var v1011 int32
	_ = v1011
	var v1017 int32
	_ = v1017
	var v1045 int32
	_ = v1045
	var v1047 int32
	_ = v1047
	var v1048 float64
	_ = v1048
	var v1049 int32
	_ = v1049
	var v1052 int32
	_ = v1052
	var v1053 float64
	_ = v1053
	var v1059 int32
	_ = v1059
	var v1060 int32
	_ = v1060
	var v1064 int32
	_ = v1064
	var v1067 int32
	_ = v1067
	var v1070 int32
	_ = v1070
	var v1074 int32
	_ = v1074
	var v1075 int32
	_ = v1075
	var v1076 int32
	_ = v1076
	var v1077 int32
	_ = v1077
	var v1078 int32
	_ = v1078
	var v1082 int32
	_ = v1082
	var v1084 int32
	_ = v1084
	var v1085 int32
	_ = v1085
	var v1091 int32
	_ = v1091
	var v1092 int32
	_ = v1092
	var v1093 int32
	_ = v1093
	var v1094 int32
	_ = v1094
	var v1107 int32
	_ = v1107
	var v1132 int32
	_ = v1132
	var v1135 int32
	_ = v1135
	var v1138 int32
	_ = v1138
	var v1139 int32
	_ = v1139
	var v1174 int32
	_ = v1174
	var v1187 int32
	_ = v1187
	var v1208 int32
	_ = v1208
	var v1209 int32
	_ = v1209
	var v1211 float64
	_ = v1211
	var v1248 int32
	_ = v1248
	var v1283 int32
	_ = v1283
	var v1322 int32
	_ = v1322
	var v1323 int32
	_ = v1323
	var v1324 int32
	_ = v1324
	var v1325 int32
	_ = v1325
	var v1326 int32
	_ = v1326
	var v1328 int32
	_ = v1328
	var v1331 int32
	_ = v1331
	var v1332 int32
	_ = v1332
	var v1333 int32
	_ = v1333
	var v1336 int32
	_ = v1336
	var v1340 int32
	_ = v1340
	var v1341 int32
	_ = v1341
	var v1342 int32
	_ = v1342
	var v1345 int32
	_ = v1345
	var v1353 int32
	_ = v1353
	var v1357 int32
	_ = v1357
	var v1363 int32
	_ = v1363
	var v1365 int32
	_ = v1365
	var v1381 int32
	_ = v1381
	var v1385 int32
	_ = v1385
	var v1386 int32
	_ = v1386
	var v1389 int32
	_ = v1389
	var v1390 int32
	_ = v1390
	var v1391 int32
	_ = v1391
	var v1392 int32
	_ = v1392
	var v1393 int32
	_ = v1393
	var v1396 int32
	_ = v1396
	var v1397 int32
	_ = v1397
	var v1400 int32
	_ = v1400
	var v1401 int32
	_ = v1401
	var v1402 int32
	_ = v1402
	var v1403 int32
	_ = v1403
	var v1404 int32
	_ = v1404
	var v1406 int32
	_ = v1406
	var v1407 int32
	_ = v1407
	var v1409 int32
	_ = v1409
	var v1410 int32
	_ = v1410
	var v1411 int32
	_ = v1411
	var v1414 int32
	_ = v1414
	var v1415 int32
	_ = v1415
	var v1423 int32
	_ = v1423
	var v1425 int32
	_ = v1425
	var v1428 int32
	_ = v1428
	var v1432 int32
	_ = v1432
	var v1443 int32
	_ = v1443
	var v1444 int32
	_ = v1444
	var v1445 int32
	_ = v1445
	var v1448 int32
	_ = v1448
	var v1449 int32
	_ = v1449
	var v1466 int32
	_ = v1466
	var v1473 int32
	_ = v1473
	var v1484 int32
	_ = v1484
	var v1488 int32
	_ = v1488
	var v1490 int32
	_ = v1490
	var v1491 int32
	_ = v1491
	var v1494 int32
	_ = v1494
	var v1495 int32
	_ = v1495
	var v1496 int32
	_ = v1496
	var v1503 int32
	_ = v1503
	var v1504 int32
	_ = v1504
	var v1506 int32
	_ = v1506
	var v1507 int32
	_ = v1507
	var v1510 int32
	_ = v1510
	var v1511 int32
	_ = v1511
	var v1522 int32
	_ = v1522
	var v1546 int32
	_ = v1546
	var v1547 int32
	_ = v1547
	var v1550 int32
	_ = v1550
	var v1551 int32
	_ = v1551
	var v1557 int32
	_ = v1557
	var v1558 int32
	_ = v1558
	var v1595 int32
	_ = v1595
	var v1597 int32
	_ = v1597
	var v1611 int32
	_ = v1611
	var v1615 int32
	_ = v1615
	var v1632 int32
	_ = v1632
	var v1633 int32
	_ = v1633
	var v1637 int32
	_ = v1637
	var v1638 int32
	_ = v1638
	var v1641 int32
	_ = v1641
	var v1642 int32
	_ = v1642
	var v1648 int32
	_ = v1648
	var v1653 int32
	_ = v1653
	var v1677 int32
	_ = v1677
	var v1678 int32
	_ = v1678
	var v1681 int32
	_ = v1681
	var v1685 int32
	_ = v1685
	var v1686 int32
	_ = v1686
	var v1687 int32
	_ = v1687
	var v1689 int32
	_ = v1689
	var v1690 int32
	_ = v1690
	var v1695 int32
	_ = v1695
	var v1724 int32
	_ = v1724
	var v1725 int32
	_ = v1725
	var v1727 int32
	_ = v1727
	var v1728 int32
	_ = v1728
	var v1745 int32
	_ = v1745
	var v1763 float64
	_ = v1763
	var v1764 int32
	_ = v1764
	var v1767 int32
	_ = v1767
	var v1768 int32
	_ = v1768
	var v1770 int32
	_ = v1770
	var v1771 int32
	_ = v1771
	var v1787 int32
	_ = v1787
	var v1807 int32
	_ = v1807
	var v1809 int32
	_ = v1809
	var v1810 int32
	_ = v1810
	var v1812 int32
	_ = v1812
	var v1813 int32
	_ = v1813
	var v1814 int32
	_ = v1814
	var v1815 int32
	_ = v1815
	var v1818 int32
	_ = v1818
	var v1819 int32
	_ = v1819
	var v1832 int32
	_ = v1832
	var v1836 int32
	_ = v1836
	var v1845 int32
	_ = v1845
	var v1846 int32
	_ = v1846
	var v1848 int32
	_ = v1848
	var v1849 int32
	_ = v1849
	var v1850 int32
	_ = v1850
	var v1867 int32
	_ = v1867
	var v1871 int32
	_ = v1871
	var v1873 int32
	_ = v1873
	var v1891 int32
	_ = v1891
	var v1892 int32
	_ = v1892
	var v1897 int32
	_ = v1897
	var v1898 int32
	_ = v1898
	var v1903 int32
	_ = v1903
	var v1905 int32
	_ = v1905
	var v1906 int32
	_ = v1906
	var v1907 int32
	_ = v1907
	var v1908 int32
	_ = v1908
	var v1910 int32
	_ = v1910
	var v1911 int32
	_ = v1911
	var v1912 int32
	_ = v1912
	var v1932 int32
	_ = v1932
	var v1945 int32
	_ = v1945
	var v1947 int32
	_ = v1947
	var v1967 int32
	_ = v1967
	v9 = int32(0)
	v31 = float64(0)
	v33 = m.G0
	v35 = v33 - int32(32)
	m.G0 = v35
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v40 = *(*float64)(unsafe.Add(mBase, _consts[528]))
	v42 = *(*int32)(unsafe.Add(mBase, _consts[523]))
	v46 = base.F64_mul(base.F64_mul(v40, base.F64_convert_i32_s(v42)), float64(1024))
	v47 = float64(4.294967295e+09)
	if base.F64_lt(v46, v47) != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	if l3 == int32(0) {
		goto L11
	} else {
		goto L12
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
	if base.F64_lt(v50, float64(4.294967296e+09))&base.F64_ge(v50, float64(0)) != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v56 = base.I32_trunc_f64_u(v50)
	v58 = v56
	goto L1
L6:
	;
	goto L7
L7:
	;
	v58 = int32(0)
	goto L1
L8:
	;
	m.G0 = v1967 + int32(32)
	return
L9:
	;
	F_add_path(m, l1, v1945)
	mBase = m.M
	v1947 = m.ExcPending
	if v1947 != 0 {
		goto L47
	} else {
		goto L273
	}
L10:
	;
	if v1867 == int32(0) {
		v1967 = v35
		goto L8
	} else {
		goto L262
	}
L11:
	;
	v62 = int32(0)
	if v59 == v62 {
		v142 = v9
		v143 = v31
		v144 = v62
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	if v59 == int32(0) {
		v1967 = v35
		goto L8
	} else {
		goto L106
	}
L14:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)+332))
	if v146 != 0 {
		goto L40
	} else {
		goto L41
	}
L15:
	;
	v65 = int32(0)
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v59)+12))
	if v66 == v65 {
		v142 = v9
		v143 = v31
		v144 = v65
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l2)+64))
	if v69 == v70 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	if v123 == int32(0) {
		v142 = v9
		v143 = v31
		v144 = v66
		goto L14
	} else {
		goto L35
	}
L18:
	;
	v123 = int32(1)
	goto L17
L19:
	;
	goto L20
L20:
	;
	v79 = int32(0)
	goto L22
L21:
	;
	v123 = v115
	goto L17
L22:
	;
	v83 = int32(0)
	if v69 == v83 {
		v93 = v83
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v115 = int32(0)
	goto L21
L24:
	;
	if v70 != 0 {
		goto L28
	} else {
		goto L29
	}
L25:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	if v87 <= v79 {
		v93 = int32(0)
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v69)+12))
	v93 = v89 + v79<<(uint(int32(2))%32)
	goto L24
L27:
	;
	v99 = base.B2i32(v93 == int32(0))
	if v93 == int32(0) {
		v115 = v99
		goto L21
	} else {
		goto L32
	}
L28:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v70)+4))
	if v79 < v94 {
		goto L27
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v123 = base.B2i32(v93 == int32(0))
	goto L17
L31:
	;
	goto L30
L32:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v70)+12))
	v105 = v102 + v79<<(uint(int32(2))%32)
	if v105 == int32(0) {
		v115 = v99
		goto L21
	} else {
		goto L33
	}
L33:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
	if v110 == v111 {
		v79 = v79 + int32(1)
		goto L22
	} else {
		goto L34
	}
L34:
	;
	goto L23
L35:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	v127 = *(*float64)(unsafe.Add(mBase, uint32(v126)+16))
	v129 = v66 + int32(4)
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)+12))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v131)+4))
	if base.Ui32(v129) < base.Ui32(v132+v133<<(uint(int32(2))%32)) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v138 = v129
	goto L38
L37:
	;
	v138 = int32(0)
	goto L38
L38:
	;
	v142 = v126
	v143 = v127
	v144 = v138
	goto L14
L39:
	;
	if base.F64_gt(base.F64_mul(base.F64_sub(l7, v143), base.F64_convert_i32_u(v153)), base.F64_convert_i32_u(v58)) != 0 {
		goto L43
	} else {
		goto L44
	}
L40:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v146)+4))
	v149 = v147
	goto L42
L41:
	;
	v149 = int32(0)
	goto L42
L42:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v150)+32))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l6)+32))
	v153 = F_hash_agg_entry_size(m, v149, v151, v152)
	mBase = m.M
	goto L39
L43:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	if v158 != 0 {
		v1967 = v35
		goto L8
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l5)+28))
	v160 = F_list_copy(m, v159)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	goto L45
L47:
	;
	return
L48:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	if v144 != 0 {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	if v232 == int32(0) {
		v1967 = v35
		goto L8
	} else {
		goto L61
	}
L50:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v162)+4))
	if v171 <= v170 {
		v232 = v160
		goto L49
	} else {
		goto L55
	}
L51:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v162)+12))
	v170 = (v144 - v163) >> (uint(int32(2)) % 32)
	goto L50
L52:
	;
	goto L53
L53:
	;
	if v162 == int32(0) {
		v232 = v160
		goto L49
	} else {
		goto L54
	}
L54:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v162)+4))
	v170 = v169
	goto L50
L55:
	;
	v181 = v170
	v185 = v160
	goto L56
L56:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v162)+12))
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v205+v181<<(uint(int32(2))%32))))
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+24)))
	if v210 != int32(1) {
		v1967 = v35
		goto L8
	} else {
		goto L58
	}
L57:
	;
	v232 = v214
	goto L49
L58:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v209)+12))
	v214 = F_list_concat(m, v185, v213)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L47
	} else {
		goto L59
	}
L59:
	;
	v217 = v181 + int32(1)
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v162)+4))
	if v217 < v218 {
		v181 = v217
		v185 = v214
		goto L56
	} else {
		goto L60
	}
L60:
	;
	goto L57
L61:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v232)+4))
	if v254 <= int32(0) {
		v1867 = v9
		v1871 = v9
		v1873 = v9
		goto L10
	} else {
		goto L62
	}
L62:
	;
	v271 = v9
	v273 = v9
	v277 = v9
	v279 = v9
	goto L63
L63:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v232)+12))
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v289+v271<<(uint(int32(2))%32))))
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v293)+4))
	if v294 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L64:
	;
	v1867 = v680
	v1871 = v684
	v1873 = v686
	goto L10
L65:
	;
	v697 = v271 + int32(1)
	v698 = *(*int32)(unsafe.Add(mBase, uint32(v232)+4))
	if v697 < v698 {
		v271 = v697
		v273 = v680
		v277 = v684
		v279 = v686
		goto L63
	} else {
		goto L105
	}
L66:
	;
	v297 = F_lappend(m, v279, v293)
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L47
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v303 = F_palloc0(m, int32(32))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L47
	} else {
		goto L71
	}
L69:
	;
	v300 = F_lappend(m, v277, int32(0))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L47
	} else {
		goto L70
	}
L70:
	;
	v680 = v273
	v684 = v300
	v686 = v297
	goto L65
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v303))) = int32(309)
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v294)+4))
	if v307 <= int32(0) {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v303)+4)) = v369
	*(*int32)(unsafe.Add(mBase, uint32(v35)+16)) = v293
	*(*int32)(unsafe.Add(mBase, uint32(v35)+28)) = v293
	v398 = F_list_make1_impl(m, int32(1), v35+int32(16))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L47
	} else {
		goto L81
	}
L73:
	;
	v369 = int32(0)
	goto L72
L74:
	;
	goto L75
L75:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v312 = int32(0)
	v322 = v312
	v323 = v312
	goto L76
L76:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v294)+12))
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v346+v322<<(uint(int32(2))%32))))
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v311)+100))
	v352 = F_get_sortgroupref_clause(m, v350, v351)
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L47
	} else {
		goto L78
	}
L77:
	;
	v369 = v354
	goto L72
L78:
	;
	v354 = F_lappend(m, v323, v352)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L47
	} else {
		goto L79
	}
L79:
	;
	v357 = v322 + int32(1)
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v294)+4))
	if v357 < v358 {
		v322 = v357
		v323 = v354
		goto L76
	} else {
		goto L80
	}
L80:
	;
	goto L77
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v303)+12)) = v398
	v401 = *(*int32)(unsafe.Add(mBase, uint32(l5)+32))
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v303)+4))
	if v402 == int32(0) {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	if v398 == int32(0) {
		goto L89
	} else {
		goto L90
	}
L83:
	;
	v405 = int32(0)
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v402)+4))
	if v406 <= v405 {
		goto L82
	} else {
		goto L84
	}
L84:
	;
	v417 = v405
	goto L85
L85:
	;
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v402)+12))
	v442 = int32(2)
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v441+v417<<(uint(v442)%32))))
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v445)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v401+v446<<(uint(v442)%32)))) = v417
	v452 = v417 + int32(1)
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v402)+4))
	if v452 < v453 {
		v417 = v452
		goto L85
	} else {
		goto L87
	}
L86:
	;
	goto L82
L87:
	;
	goto L86
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v303)+8)) = v640
	v658 = *(*float64)(unsafe.Add(mBase, uint32(v293)+8))
	v659 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v303)+24)) = uint16(v659)
	*(*float64)(unsafe.Add(mBase, uint32(v303)+16)) = v658
	v662 = F_lappend(m, v273, v303)
	mBase = m.M
	v663 = m.ExcPending
	if v663 != 0 {
		goto L47
	} else {
		goto L104
	}
L89:
	;
	v640 = int32(0)
	goto L88
L90:
	;
	goto L91
L91:
	;
	v490 = int32(0)
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v398)+4))
	if v492 <= v490 {
		v640 = v490
		goto L88
	} else {
		goto L92
	}
L92:
	;
	v506 = v490
	v510 = v490
	goto L93
L93:
	;
	v527 = int32(0)
	v528 = *(*int32)(unsafe.Add(mBase, uint32(v398)+12))
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v528+v506<<(uint(int32(2))%32))))
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v532)+4))
	if v533 == v527 {
		v590 = v527
		goto L95
	} else {
		goto L96
	}
L94:
	;
	v640 = v619
	goto L88
L95:
	;
	v619 = F_lappend(m, v510, v590)
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L47
	} else {
		goto L102
	}
L96:
	;
	v536 = int32(0)
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v533)+4))
	if v537 <= v536 {
		v590 = v527
		goto L95
	} else {
		goto L97
	}
L97:
	;
	v543 = v527
	v548 = v536
	goto L98
L98:
	;
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v533)+12))
	v573 = int32(2)
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v572+v548<<(uint(v573)%32))))
	v580 = *(*int32)(unsafe.Add(mBase, uint32(v401+v576<<(uint(v573)%32))))
	v581 = F_lappend_int(m, v543, v580)
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L47
	} else {
		goto L100
	}
L99:
	;
	v590 = v581
	goto L95
L100:
	;
	v584 = v548 + int32(1)
	v585 = *(*int32)(unsafe.Add(mBase, uint32(v533)+4))
	if v584 < v585 {
		v543 = v581
		v548 = v584
		goto L98
	} else {
		goto L101
	}
L101:
	;
	goto L99
L102:
	;
	v622 = v506 + int32(1)
	v623 = *(*int32)(unsafe.Add(mBase, uint32(v398)+4))
	if v622 < v623 {
		v506 = v622
		v510 = v619
		goto L93
	} else {
		goto L103
	}
L103:
	;
	goto L94
L104:
	;
	v680 = v662
	v684 = v277
	v686 = v279
	goto L65
L105:
	;
	goto L64
L106:
	;
	if l4 == int32(0) {
		v1813 = l0
		v1814 = l1
		v1815 = l2
		v1818 = l5
		v1819 = l6
		v1832 = v35
		v1836 = v37
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v1845 = *(*int32)(unsafe.Add(mBase, uint32(v1818)+28))
	if v1845 != 0 {
		v1967 = v1832
		goto L8
	} else {
		goto L260
	}
L108:
	;
	v704 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5)+16)))
	if v704 != int32(1) {
		v1813 = l0
		v1814 = l1
		v1815 = l2
		v1818 = l5
		v1819 = l6
		v1832 = v35
		v1836 = v37
		goto L107
	} else {
		goto L109
	}
L109:
	;
	v707 = *(*int32)(unsafe.Add(mBase, uint32(l5)+28))
	v708 = F_list_copy(m, v707)
	mBase = m.M
	v709 = m.ExcPending
	if v709 != 0 {
		goto L47
	} else {
		goto L110
	}
L110:
	;
	v711 = *(*float64)(unsafe.Add(mBase, uint32(l5)+8))
	v712 = *(*int32)(unsafe.Add(mBase, uint32(l0)+332))
	if v712 != 0 {
		goto L113
	} else {
		goto L114
	}
L111:
	;
	if v1423 != 0 {
		goto L222
	} else {
		goto L223
	}
L112:
	;
	v722 = base.F64_sub(base.F64_convert_i32_u(v58), base.F64_mul(v711, base.F64_convert_i32_u(v719)))
	if base.F64_gt(v722, float64(0)) == int32(0) {
		v1409 = l0
		v1410 = l1
		v1411 = l2
		v1414 = l5
		v1415 = l6
		v1423 = v9
		v1425 = v708
		v1428 = v35
		v1432 = v37
		goto L111
	} else {
		goto L116
	}
L113:
	;
	v713 = *(*int32)(unsafe.Add(mBase, uint32(v712)+4))
	v715 = v713
	goto L115
L114:
	;
	v715 = int32(0)
	goto L115
L115:
	;
	v716 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v717 = *(*int32)(unsafe.Add(mBase, uint32(v716)+32))
	v718 = *(*int32)(unsafe.Add(mBase, uint32(l6)+32))
	v719 = F_hash_agg_entry_size(m, v715, v717, v718)
	mBase = m.M
	goto L112
L116:
	;
	v727 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	if v727 == int32(0) {
		v1409 = l0
		v1410 = l1
		v1411 = l2
		v1414 = l5
		v1415 = l6
		v1423 = v9
		v1425 = v708
		v1428 = v35
		v1432 = v37
		goto L111
	} else {
		goto L117
	}
L117:
	;
	v730 = *(*int32)(unsafe.Add(mBase, uint32(v727)+4))
	if v730 < int32(2) {
		v1409 = l0
		v1410 = l1
		v1411 = l2
		v1414 = l5
		v1415 = l6
		v1423 = v9
		v1425 = v708
		v1428 = v35
		v1432 = v37
		goto L111
	} else {
		goto L118
	}
L118:
	;
	v736 = base.F64_div(v722, base.F64_mul(base.F64_convert_i32_u(v730), float64(20)))
	v737 = float64(1)
	if base.F64_gt(v736, v737) != 0 {
		goto L120
	} else {
		goto L121
	}
L119:
	;
	v751 = F_palloc(m, v730<<(uint(int32(2))%32))
	mBase = m.M
	v752 = m.ExcPending
	if v752 != 0 {
		goto L47
	} else {
		goto L126
	}
L120:
	;
	v740 = v736
	goto L122
L121:
	;
	v740 = v737
	goto L122
L122:
	;
	v742 = base.F64_floor(base.F64_div(v722, v740))
	if base.F64_lt(base.F64_abs(v742), float64(2.147483648e+09)) != 0 {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v746 = base.I32_trunc_f64_s(v742)
	v748 = v746
	goto L119
L124:
	;
	goto L125
L125:
	;
	v748 = int32(-2147483648)
	goto L119
L126:
	;
	v753 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	if v753 == int32(0) {
		v1409 = l0
		v1410 = l1
		v1411 = l2
		v1414 = l5
		v1415 = l6
		v1423 = v9
		v1425 = v708
		v1428 = v35
		v1432 = v37
		goto L111
	} else {
		goto L127
	}
L127:
	;
	v756 = int32(1)
	v757 = int32(0)
	v758 = *(*int32)(unsafe.Add(mBase, uint32(v753)+4))
	if v756 < v758 {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v763 = base.F64_add(base.F64_convert_i32_s(v748), float64(1))
	v772 = v756
	v773 = v757
	goto L131
L129:
	;
	v846 = v757
	goto L130
L130:
	;
	if v846 <= int32(0) {
		v1409 = l0
		v1410 = l1
		v1411 = l2
		v1414 = l5
		v1415 = l6
		v1423 = v9
		v1425 = v708
		v1428 = v35
		v1432 = v37
		goto L111
	} else {
		goto L148
	}
L131:
	;
	v796 = *(*int32)(unsafe.Add(mBase, uint32(v753)+12))
	v800 = *(*int32)(unsafe.Add(mBase, uint32(v796+v772<<(uint(int32(2))%32))))
	v801 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v800)+24)))
	if v801 == int32(1) {
		goto L133
	} else {
		goto L134
	}
L132:
	;
	v846 = v831
	goto L130
L133:
	;
	v807 = *(*float64)(unsafe.Add(mBase, uint32(v800)+16))
	v808 = *(*int32)(unsafe.Add(mBase, uint32(l0)+332))
	if v808 != 0 {
		goto L138
	} else {
		goto L139
	}
L134:
	;
	v831 = v773
	goto L135
L135:
	;
	v834 = v772 + int32(1)
	v835 = *(*int32)(unsafe.Add(mBase, uint32(v753)+4))
	if v834 < v835 {
		v772 = v834
		v773 = v831
		goto L131
	} else {
		goto L147
	}
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v751+v773<<(uint(int32(2))%32)))) = v827
	v831 = v773 + int32(1)
	goto L135
L137:
	;
	v819 = base.F64_floor(base.F64_div(base.F64_mul(v807, base.F64_convert_i32_u(v815)), v740))
	if base.F64_gt(v763, v819) != 0 {
		goto L141
	} else {
		goto L142
	}
L138:
	;
	v809 = *(*int32)(unsafe.Add(mBase, uint32(v808)+4))
	v811 = v809
	goto L140
L139:
	;
	v811 = int32(0)
	goto L140
L140:
	;
	v812 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v813 = *(*int32)(unsafe.Add(mBase, uint32(v812)+32))
	v814 = *(*int32)(unsafe.Add(mBase, uint32(l6)+32))
	v815 = F_hash_agg_entry_size(m, v811, v813, v814)
	mBase = m.M
	goto L137
L141:
	;
	v821 = v819
	goto L143
L142:
	;
	v821 = v763
	goto L143
L143:
	;
	if base.F64_lt(base.F64_abs(v821), float64(2.147483648e+09)) != 0 {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v825 = base.I32_trunc_f64_s(v821)
	v827 = v825
	goto L136
L145:
	;
	goto L146
L146:
	;
	v827 = int32(-2147483648)
	goto L136
L147:
	;
	goto L132
L148:
	;
	v871 = int32(0)
	v873 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v878 = F_AllocSetContextCreateInternal(m, v873, int32(302922), v871, int32(1024), int32(8192))
	mBase = m.M
	v879 = m.ExcPending
	if v879 != 0 {
		goto L47
	} else {
		goto L149
	}
L149:
	;
	v880 = int32(4442576)
	v881 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v878
	v885 = v748 + int32(1)
	v888 = F_palloc(m, v885<<(uint(int32(3))%32))
	mBase = m.M
	v889 = m.ExcPending
	if v889 != 0 {
		goto L47
	} else {
		goto L150
	}
L150:
	;
	v892 = F_palloc(m, v885<<(uint(int32(2))%32))
	mBase = m.M
	v893 = m.ExcPending
	if v893 != 0 {
		goto L47
	} else {
		goto L151
	}
L151:
	;
	if int32(0) <= v748 {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	v900 = v871
	goto L155
L153:
	;
	goto L154
L154:
	;
	if int32(0) < v846 {
		goto L159
	} else {
		goto L160
	}
L155:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v888+v900<<(uint(int32(3))%32)))) = int64(0)
	v936 = F_bms_make_singleton(m, v846)
	mBase = m.M
	v937 = m.ExcPending
	if v937 != 0 {
		goto L47
	} else {
		goto L157
	}
L156:
	;
	goto L154
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v892+v900<<(uint(int32(2))%32)))) = v936
	v940 = v900 + int32(1)
	if v940 <= v748 {
		v900 = v940
		goto L155
	} else {
		goto L158
	}
L158:
	;
	goto L156
L159:
	;
	v993 = v9
	goto L162
L160:
	;
	goto L161
L161:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v881
	v1322 = *(*int32)(unsafe.Add(mBase, uint32(v892+v748<<(uint(int32(2))%32))))
	v1323 = F_bms_copy(m, v1322)
	mBase = m.M
	v1324 = m.ExcPending
	if v1324 != 0 {
		goto L47
	} else {
		goto L199
	}
L162:
	;
	v1011 = *(*int32)(unsafe.Add(mBase, uint32(v751+v993<<(uint(int32(2))%32))))
	if v1011 <= v748 {
		goto L164
	} else {
		goto L165
	}
L163:
	;
	goto L161
L164:
	;
	v1017 = v748
	goto L167
L165:
	;
	goto L166
L166:
	;
	v1283 = v993 + int32(1)
	if v1283 != v846 {
		v993 = v1283
		goto L162
	} else {
		goto L198
	}
L167:
	;
	v1045 = int32(3)
	v1047 = v888 + v1017<<(uint(v1045)%32)
	v1048 = *(*float64)(unsafe.Add(mBase, uint32(v1047)))
	v1049 = v1017 - v1011
	v1052 = v888 + v1049<<(uint(v1045)%32)
	v1053 = *(*float64)(unsafe.Add(mBase, uint32(v1052)))
	if base.F64_le(v1048, base.F64_add(v1053, float64(1))) != 0 {
		goto L169
	} else {
		goto L170
	}
L168:
	;
	goto L166
L169:
	;
	v1059 = v892 + v1017<<(uint(int32(2))%32)
	v1060 = *(*int32)(unsafe.Add(mBase, uint32(v1059)))
	if v1011 != 0 {
		goto L172
	} else {
		goto L173
	}
L170:
	;
	goto L171
L171:
	;
	v1248 = v1017 - int32(1)
	if v1011 <= v1248 {
		v1017 = v1248
		goto L167
	} else {
		goto L197
	}
L172:
	;
	v1064 = *(*int32)(unsafe.Add(mBase, uint32(v892+v1049<<(uint(int32(2))%32))))
	if v1060 == int32(0) {
		goto L176
	} else {
		goto L177
	}
L173:
	;
	v1187 = v1060
	goto L174
L174:
	;
	v1208 = F_bms_add_member(m, v1187, v993)
	mBase = m.M
	v1209 = m.ExcPending
	if v1209 != 0 {
		goto L47
	} else {
		goto L196
	}
L175:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1059))) = v1174
	v1187 = v1174
	goto L174
L176:
	;
	v1067 = int32(0)
	if v1064 == v1067 {
		v1174 = v1067
		goto L175
	} else {
		goto L179
	}
L177:
	;
	goto L178
L178:
	;
	if v1064 == int32(0) {
		goto L185
	} else {
		goto L186
	}
L179:
	;
	v1070 = *(*int32)(unsafe.Add(mBase, uint32(v1064)+4))
	v1074 = v1070<<(uint(int32(2))%32) + int32(8)
	v1075 = F_palloc(m, v1074)
	mBase = m.M
	v1076 = m.ExcPending
	if v1076 != 0 {
		goto L47
	} else {
		goto L180
	}
L180:
	;
	if v1074 != 0 {
		goto L182
	} else {
		goto L183
	}
L181:
	;
	v1174 = v1078
	goto L175
L182:
	;
	v1077 = F__emscripten_memcpy_bulkmem(m, v1075, v1064, v1074)
	mBase = m.M
	v1078 = v1077
	goto L184
L183:
	;
	v1078 = v1075
	goto L184
L184:
	;
	goto L181
L185:
	;
	F_pfree(m, v1060)
	mBase = m.M
	v1082 = m.ExcPending
	if v1082 != 0 {
		goto L47
	} else {
		goto L188
	}
L186:
	;
	goto L187
L187:
	;
	v1084 = *(*int32)(unsafe.Add(mBase, uint32(v1064)+4))
	v1085 = *(*int32)(unsafe.Add(mBase, uint32(v1060)+4))
	if v1085 < v1084 {
		goto L189
	} else {
		goto L190
	}
L188:
	;
	v1174 = int32(0)
	goto L175
L189:
	;
	v1091 = F_repalloc(m, v1060, v1084<<(uint(int32(2))%32)+int32(8))
	mBase = m.M
	v1092 = m.ExcPending
	if v1092 != 0 {
		goto L47
	} else {
		goto L192
	}
L190:
	;
	v1093 = v1060
	goto L191
L191:
	;
	v1094 = int32(8)
	v1107 = int32(0)
	goto L193
L192:
	;
	v1093 = v1091
	goto L191
L193:
	;
	v1132 = v1107 << (uint(int32(2)) % 32)
	v1135 = *(*int32)(unsafe.Add(mBase, uint32(v1132+(v1064+v1094))))
	*(*int32)(unsafe.Add(mBase, uint32(v1093+v1094+v1132))) = v1135
	v1138 = v1107 + int32(1)
	v1139 = *(*int32)(unsafe.Add(mBase, uint32(v1064)+4))
	if v1138 < v1139 {
		v1107 = v1138
		goto L193
	} else {
		goto L195
	}
L194:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1093)+4)) = v1139
	v1174 = v1093
	goto L175
L195:
	;
	goto L194
L196:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1059))) = v1208
	v1211 = *(*float64)(unsafe.Add(mBase, uint32(v1052)))
	*(*float64)(unsafe.Add(mBase, uint32(v1047))) = base.F64_add(v1211, float64(1))
	goto L171
L197:
	;
	goto L168
L198:
	;
	goto L163
L199:
	;
	v1325 = F_bms_del_member(m, v1323, v846)
	mBase = m.M
	v1326 = m.ExcPending
	if v1326 != 0 {
		goto L47
	} else {
		goto L200
	}
L200:
	;
	F_MemoryContextDelete(m, v878)
	mBase = m.M
	v1328 = m.ExcPending
	if v1328 != 0 {
		goto L47
	} else {
		goto L201
	}
L201:
	;
	if v1325 == int32(0) {
		v1409 = l0
		v1410 = l1
		v1411 = l2
		v1414 = l5
		v1415 = l6
		v1423 = v9
		v1425 = v708
		v1428 = v35
		v1432 = v37
		goto L111
	} else {
		goto L202
	}
L202:
	;
	v1331 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	v1332 = *(*int32)(unsafe.Add(mBase, uint32(v1331)+12))
	v1333 = *(*int32)(unsafe.Add(mBase, uint32(v1332)))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+12)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v35)+24)) = v1333
	v1336 = int32(1)
	v1340 = F_list_make1_impl(m, v1336, v35+int32(12))
	mBase = m.M
	v1341 = m.ExcPending
	if v1341 != 0 {
		goto L47
	} else {
		goto L203
	}
L203:
	;
	v1342 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	if v1342 == int32(0) {
		v1409 = l0
		v1410 = l1
		v1411 = l2
		v1414 = l5
		v1415 = l6
		v1423 = v1340
		v1425 = v708
		v1428 = v35
		v1432 = v37
		goto L111
	} else {
		goto L204
	}
L204:
	;
	v1345 = *(*int32)(unsafe.Add(mBase, uint32(v1342)+4))
	if v1345 < int32(2) {
		v1409 = l0
		v1410 = l1
		v1411 = l2
		v1414 = l5
		v1415 = l6
		v1423 = v1340
		v1425 = v708
		v1428 = v35
		v1432 = v37
		goto L111
	} else {
		goto L205
	}
L205:
	;
	v1353 = int32(0)
	v1357 = v1336
	v1363 = v1340
	v1365 = v708
	goto L206
L206:
	;
	v1381 = *(*int32)(unsafe.Add(mBase, uint32(v1342)+12))
	v1385 = *(*int32)(unsafe.Add(mBase, uint32(v1381+v1357<<(uint(int32(2))%32))))
	v1386 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1385)+24)))
	if v1386 == int32(1) {
		goto L209
	} else {
		goto L210
	}
L207:
	;
	v1409 = l0
	v1410 = l1
	v1411 = l2
	v1414 = l5
	v1415 = l6
	v1423 = v1403
	v1425 = v1404
	v1428 = v35
	v1432 = v37
	goto L111
L208:
	;
	v1406 = v1357 + int32(1)
	v1407 = *(*int32)(unsafe.Add(mBase, uint32(v1342)+4))
	if v1406 < v1407 {
		v1353 = v1402
		v1357 = v1406
		v1363 = v1403
		v1365 = v1404
		goto L206
	} else {
		goto L219
	}
L209:
	;
	v1389 = F_bms_is_member(m, v1353, v1325)
	mBase = m.M
	v1390 = m.ExcPending
	if v1390 != 0 {
		goto L47
	} else {
		goto L212
	}
L210:
	;
	goto L211
L211:
	;
	v1400 = F_lappend(m, v1363, v1385)
	mBase = m.M
	v1401 = m.ExcPending
	if v1401 != 0 {
		goto L47
	} else {
		goto L218
	}
L212:
	;
	if v1389 != 0 {
		goto L213
	} else {
		goto L214
	}
L213:
	;
	v1391 = *(*int32)(unsafe.Add(mBase, uint32(v1385)+12))
	v1392 = F_list_concat(m, v1365, v1391)
	mBase = m.M
	v1393 = m.ExcPending
	if v1393 != 0 {
		goto L47
	} else {
		goto L216
	}
L214:
	;
	goto L215
L215:
	;
	v1396 = F_lappend(m, v1363, v1385)
	mBase = m.M
	v1397 = m.ExcPending
	if v1397 != 0 {
		goto L47
	} else {
		goto L217
	}
L216:
	;
	v1402 = v1353 + int32(1)
	v1403 = v1363
	v1404 = v1392
	goto L208
L217:
	;
	v1402 = v1353 + int32(1)
	v1403 = v1396
	v1404 = v1365
	goto L208
L218:
	;
	v1402 = v1353
	v1403 = v1400
	v1404 = v1365
	goto L208
L219:
	;
	goto L207
L220:
	;
	if v1787 == int32(0) {
		v1813 = v1409
		v1814 = v1410
		v1815 = v1411
		v1818 = v1414
		v1819 = v1415
		v1832 = v1428
		v1836 = v1432
		goto L107
	} else {
		goto L257
	}
L221:
	;
	v1449 = *(*int32)(unsafe.Add(mBase, uint32(v1425)+4))
	if v1449 <= int32(0) {
		v1787 = v1448
		goto L220
	} else {
		goto L227
	}
L222:
	;
	if v1425 == int32(0) {
		v1787 = v1423
		goto L220
	} else {
		goto L226
	}
L223:
	;
	if v1425 == int32(0) {
		goto L222
	} else {
		goto L224
	}
L224:
	;
	v1443 = *(*int32)(unsafe.Add(mBase, uint32(v1414)))
	v1444 = F_list_copy(m, v1443)
	mBase = m.M
	v1445 = m.ExcPending
	if v1445 != 0 {
		goto L47
	} else {
		goto L225
	}
L225:
	;
	v1448 = v1444
	goto L221
L226:
	;
	v1448 = v1423
	goto L221
L227:
	;
	v1466 = v1448
	v1473 = v9
	goto L228
L228:
	;
	v1484 = *(*int32)(unsafe.Add(mBase, uint32(v1425)+12))
	v1488 = *(*int32)(unsafe.Add(mBase, uint32(v1484+v1473<<(uint(int32(2))%32))))
	v1490 = F_palloc0(m, int32(32))
	mBase = m.M
	v1491 = m.ExcPending
	if v1491 != 0 {
		goto L47
	} else {
		goto L230
	}
L229:
	;
	v1787 = v1767
	goto L220
L230:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1490))) = int32(309)
	v1494 = *(*int32)(unsafe.Add(mBase, uint32(v1488)+4))
	v1495 = F_preprocess_groupclause(m, v1409, v1494)
	mBase = m.M
	v1496 = m.ExcPending
	if v1496 != 0 {
		goto L47
	} else {
		goto L231
	}
L231:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1490)+4)) = v1495
	*(*int32)(unsafe.Add(mBase, uint32(v1428)+8)) = v1488
	*(*int32)(unsafe.Add(mBase, uint32(v1428)+20)) = v1488
	v1503 = F_list_make1_impl(m, int32(1), v1428+int32(8))
	mBase = m.M
	v1504 = m.ExcPending
	if v1504 != 0 {
		goto L47
	} else {
		goto L232
	}
L232:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1490)+12)) = v1503
	v1506 = *(*int32)(unsafe.Add(mBase, uint32(v1414)+32))
	v1507 = *(*int32)(unsafe.Add(mBase, uint32(v1490)+4))
	if v1507 == int32(0) {
		goto L233
	} else {
		goto L234
	}
L233:
	;
	if v1503 == int32(0) {
		goto L240
	} else {
		goto L241
	}
L234:
	;
	v1510 = int32(0)
	v1511 = *(*int32)(unsafe.Add(mBase, uint32(v1507)+4))
	if v1511 <= v1510 {
		goto L233
	} else {
		goto L235
	}
L235:
	;
	v1522 = v1510
	goto L236
L236:
	;
	v1546 = *(*int32)(unsafe.Add(mBase, uint32(v1507)+12))
	v1547 = int32(2)
	v1550 = *(*int32)(unsafe.Add(mBase, uint32(v1546+v1522<<(uint(v1547)%32))))
	v1551 = *(*int32)(unsafe.Add(mBase, uint32(v1550)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1506+v1551<<(uint(v1547)%32)))) = v1522
	v1557 = v1522 + int32(1)
	v1558 = *(*int32)(unsafe.Add(mBase, uint32(v1507)+4))
	if v1557 < v1558 {
		v1522 = v1557
		goto L236
	} else {
		goto L238
	}
L237:
	;
	goto L233
L238:
	;
	goto L237
L239:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1490)+8)) = v1745
	v1763 = *(*float64)(unsafe.Add(mBase, uint32(v1488)+8))
	v1764 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v1490)+24)) = uint16(v1764)
	*(*float64)(unsafe.Add(mBase, uint32(v1490)+16)) = v1763
	v1767 = F_lcons(m, v1490, v1466)
	mBase = m.M
	v1768 = m.ExcPending
	if v1768 != 0 {
		goto L47
	} else {
		goto L255
	}
L240:
	;
	v1745 = int32(0)
	goto L239
L241:
	;
	goto L242
L242:
	;
	v1595 = int32(0)
	v1597 = *(*int32)(unsafe.Add(mBase, uint32(v1503)+4))
	if v1597 <= v1595 {
		v1745 = v1595
		goto L239
	} else {
		goto L243
	}
L243:
	;
	v1611 = v1595
	v1615 = v1595
	goto L244
L244:
	;
	v1632 = int32(0)
	v1633 = *(*int32)(unsafe.Add(mBase, uint32(v1503)+12))
	v1637 = *(*int32)(unsafe.Add(mBase, uint32(v1633+v1611<<(uint(int32(2))%32))))
	v1638 = *(*int32)(unsafe.Add(mBase, uint32(v1637)+4))
	if v1638 == v1632 {
		v1695 = v1632
		goto L246
	} else {
		goto L247
	}
L245:
	;
	v1745 = v1724
	goto L239
L246:
	;
	v1724 = F_lappend(m, v1615, v1695)
	mBase = m.M
	v1725 = m.ExcPending
	if v1725 != 0 {
		goto L47
	} else {
		goto L253
	}
L247:
	;
	v1641 = int32(0)
	v1642 = *(*int32)(unsafe.Add(mBase, uint32(v1638)+4))
	if v1642 <= v1641 {
		v1695 = v1632
		goto L246
	} else {
		goto L248
	}
L248:
	;
	v1648 = v1632
	v1653 = v1641
	goto L249
L249:
	;
	v1677 = *(*int32)(unsafe.Add(mBase, uint32(v1638)+12))
	v1678 = int32(2)
	v1681 = *(*int32)(unsafe.Add(mBase, uint32(v1677+v1653<<(uint(v1678)%32))))
	v1685 = *(*int32)(unsafe.Add(mBase, uint32(v1506+v1681<<(uint(v1678)%32))))
	v1686 = F_lappend_int(m, v1648, v1685)
	mBase = m.M
	v1687 = m.ExcPending
	if v1687 != 0 {
		goto L47
	} else {
		goto L251
	}
L250:
	;
	v1695 = v1686
	goto L246
L251:
	;
	v1689 = v1653 + int32(1)
	v1690 = *(*int32)(unsafe.Add(mBase, uint32(v1638)+4))
	if v1689 < v1690 {
		v1648 = v1686
		v1653 = v1689
		goto L249
	} else {
		goto L252
	}
L252:
	;
	goto L250
L253:
	;
	v1727 = v1611 + int32(1)
	v1728 = *(*int32)(unsafe.Add(mBase, uint32(v1503)+4))
	if v1727 < v1728 {
		v1611 = v1727
		v1615 = v1724
		goto L244
	} else {
		goto L254
	}
L254:
	;
	goto L245
L255:
	;
	v1770 = v1473 + int32(1)
	v1771 = *(*int32)(unsafe.Add(mBase, uint32(v1425)+4))
	if v1770 < v1771 {
		v1466 = v1767
		v1473 = v1770
		goto L228
	} else {
		goto L256
	}
L256:
	;
	goto L229
L257:
	;
	v1807 = *(*int32)(unsafe.Add(mBase, uint32(v1432)+112))
	v1809 = F_create_groupingsets_path(m, v1409, v1410, v1411, v1807, int32(3), v1787, v1415)
	mBase = m.M
	v1810 = m.ExcPending
	if v1810 != 0 {
		goto L47
	} else {
		goto L258
	}
L258:
	;
	F_add_path(m, v1410, v1809)
	mBase = m.M
	v1812 = m.ExcPending
	if v1812 != 0 {
		goto L47
	} else {
		goto L259
	}
L259:
	;
	v1813 = v1409
	v1814 = v1410
	v1815 = v1411
	v1818 = v1414
	v1819 = v1415
	v1832 = v1428
	v1836 = v1432
	goto L107
L260:
	;
	v1846 = *(*int32)(unsafe.Add(mBase, uint32(v1836)+112))
	v1848 = *(*int32)(unsafe.Add(mBase, uint32(v1818)))
	v1849 = F_create_groupingsets_path(m, v1813, v1814, v1815, v1846, int32(1), v1848, v1819)
	mBase = m.M
	v1850 = m.ExcPending
	if v1850 != 0 {
		goto L47
	} else {
		goto L261
	}
L261:
	;
	v1932 = v1832
	v1945 = v1849
	goto L9
L262:
	;
	if v142 == int32(0) {
		goto L264
	} else {
		goto L265
	}
L263:
	;
	v1910 = *(*int32)(unsafe.Add(mBase, uint32(v37)+112))
	v1911 = F_create_groupingsets_path(m, l0, l1, l2, v1910, v1907, v1908, l6)
	mBase = m.M
	v1912 = m.ExcPending
	if v1912 != 0 {
		goto L47
	} else {
		goto L272
	}
L264:
	;
	if v1871 == int32(0) {
		goto L267
	} else {
		goto L268
	}
L265:
	;
	v1903 = v142
	goto L266
L266:
	;
	v1905 = F_lappend(m, v1867, v1903)
	mBase = m.M
	v1906 = m.ExcPending
	if v1906 != 0 {
		goto L47
	} else {
		goto L271
	}
L267:
	;
	v1907 = int32(2)
	v1908 = v1867
	goto L263
L268:
	;
	goto L269
L269:
	;
	v1891 = F_palloc0(m, int32(32))
	mBase = m.M
	v1892 = m.ExcPending
	if v1892 != 0 {
		goto L47
	} else {
		goto L270
	}
L270:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1891)+12)) = v1873
	*(*int64)(unsafe.Add(mBase, uint32(v1891))) = int64(309)
	*(*int32)(unsafe.Add(mBase, uint32(v1891)+8)) = v1871
	v1897 = *(*int32)(unsafe.Add(mBase, uint32(v1871)+4))
	v1898 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v1891)+24)) = uint16(v1898)
	*(*float64)(unsafe.Add(mBase, uint32(v1891)+16)) = base.F64_convert_i32_s(v1897)
	v1903 = v1891
	goto L266
L271:
	;
	v1907 = int32(3)
	v1908 = v1905
	goto L263
L272:
	;
	v1932 = v35
	v1945 = v1911
	goto L9
L273:
	;
	v1967 = v1932
	goto L8
}
func F_conv_18030_to_utf8(m *base.Module, l0 int32) int32 {
	var v13 int32
	_ = v13
	var v24 int32
	_ = v24
	var v86 int32
	_ = v86
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v157 int32
	_ = v157
	var v165 int32
	_ = v165
	var v191 int32
	_ = v191
	var v202 int32
	_ = v202
	var v224 int32
	_ = v224
	var v232 int32
	_ = v232
	var v254 int32
	_ = v254
	var v262 int32
	_ = v262
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v324 int32
	_ = v324
	var v335 int32
	_ = v335
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v393 int32
	_ = v393
	var v401 int32
	_ = v401
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v461 int32
	_ = v461
	var v479 int32
	_ = v479
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v539 int32
	_ = v539
	var v557 int32
	_ = v557
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v619 int32
	_ = v619
	var v630 int32
	_ = v630
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v696 int32
	_ = v696
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v756 int32
	_ = v756
	var v774 int32
	_ = v774
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v826 int32
	_ = v826
	if base.Ui32(l0+int32(2127506640)) <= base.Ui32(int32(381441)) {
		v13 = int32(255)
		v24 = int32(base.Ui32(l0)>>(uint(int32(16))%32))&int32(55)*int32(1260) + l0&v13 + int32(base.Ui32(l0)>>(uint(int32(8))%32))&v13*int32(10) - int32(61532)
		if base.Ui32(v24) < base.Ui32(int32(128)) {
			v826 = v24
			return v826
		} else {
			if base.Ui32(v24) <= base.Ui32(int32(2047)) {
				return v24&int32(63) | v24<<(uint(int32(2))%32)&int32(7936) | int32(49280)
			} else {
				if base.Ui32(v24) <= base.Ui32(int32(65535)) {
					return v24<<(uint(int32(4))%32)&int32(983040) | (v24&int32(63) | v24<<(uint(int32(2))%32)&int32(16128)) | int32(14712960)
				} else {
					return v24<<(uint(int32(2))%32)&int32(16128) | (v24<<(uint(int32(6))%32)&int32(117440512) | (v24&int32(63) | v24<<(uint(int32(4))%32)&int32(4128768))) | int32(-260013952)
				}
			}
		}
	} else {
		if base.Ui32(l0+int32(2127058887)) <= base.Ui32(int32(87295)) {
			v86 = int32(255)
			v95 = int32(base.Ui32(l0)>>(uint(int32(16))%32))&int32(63)*int32(1260) + l0&v86 + int32(base.Ui32(l0)>>(uint(int32(8))%32))&v86*int32(10)
			v97 = v95 - int32(61242)
			if base.Ui32(v97) < base.Ui32(int32(128)) {
				v826 = v97
				return v826
			} else {
				v101 = v95 - int32(61818)
				if base.Ui32(v97) <= base.Ui32(int32(2047)) {
					return v101&int32(63) | v97<<(uint(int32(2))%32)&int32(7936) | int32(49280)
				} else {
					if base.Ui32(v97) <= base.Ui32(int32(65535)) {
						return v97<<(uint(int32(4))%32)&int32(983040) | (v101&int32(63) | v97<<(uint(int32(2))%32)&int32(16128)) | int32(14712960)
					} else {
						return v97<<(uint(int32(2))%32)&int32(16128) | (v97<<(uint(int32(6))%32)&int32(117440512) | (v101&int32(63) | v97<<(uint(int32(4))%32)&int32(4128768))) | int32(-260013952)
					}
				}
			}
		} else {
			if base.Ui32(l0+int32(2110740941)) <= base.Ui32(int32(19460)) {
				v157 = int32(255)
				v165 = int32(base.Ui32(l0)>>(uint(int32(8))%32))&v157*int32(10) + l0&v157 + int32(12140)
				return v165&int32(63) | v165<<(uint(int32(2))%32)&int32(16128) | v165<<(uint(int32(4))%32)&int32(458752) | int32(14712960)
			} else {
				if base.Ui32(l0+int32(2110663624)) <= base.Ui32(int32(56058)) {
					v191 = int32(255)
					v202 = int32(base.Ui32(l0)>>(uint(int32(16))%32))&int32(51)*int32(1260) + l0&v191 + int32(base.Ui32(l0)>>(uint(int32(8))%32))&v191*int32(10) - int32(48331)
					return v202&int32(63) | v202<<(uint(int32(2))%32)&int32(16128) | v202<<(uint(int32(4))%32)&int32(983040) | int32(14712960)
				} else {
					if base.Ui32(l0+int32(2110600905)) <= base.Ui32(int32(12032)) {
						v224 = int32(255)
						v232 = int32(base.Ui32(l0)>>(uint(int32(8))%32))&v224*int32(10) + l0&v224 + int32(14671)
						return v232&int32(63) | v232<<(uint(int32(2))%32)&int32(16128) | v232<<(uint(int32(4))%32)&int32(458752) | int32(14712960)
					} else {
						if base.Ui32(l0+int32(2110545095)) <= base.Ui32(int32(9720)) {
							v254 = int32(255)
							v262 = int32(base.Ui32(l0)>>(uint(int32(8))%32))&v254*int32(10) + l0&v254 + int32(15936)
							if base.Ui32(int32(128)) <= base.Ui32(v262) {
								if base.Ui32(v262) <= base.Ui32(int32(2047)) {
									v312 = v262&int32(63) | v262<<(uint(int32(2))%32)&int32(7936) | int32(49280)
								} else {
									if base.Ui32(v262) <= base.Ui32(int32(65535)) {
										v312 = v262<<(uint(int32(4))%32)&int32(983040) | (v262&int32(63) | v262<<(uint(int32(2))%32)&int32(16128)) | int32(14712960)
									} else {
										v311 = v262<<(uint(int32(2))%32)&int32(16128) | (v262<<(uint(int32(6))%32)&int32(117440512) | (v262&int32(63) | v262<<(uint(int32(4))%32)&int32(4128768))) | int32(-260013952)
										v312 = v311
									}
								}
							} else {
								v311 = v262
								v312 = v311
							}
							return v312
						} else {
							if base.Ui32(l0+int32(2110527432)) <= base.Ui32(int32(44544)) {
								v324 = int32(255)
								v335 = int32(base.Ui32(l0)>>(uint(int32(16))%32))&int32(55)*int32(1260) + l0&v324 + int32(base.Ui32(l0)>>(uint(int32(8))%32))&v324*int32(10) - int32(48318)
								if base.Ui32(int32(128)) <= base.Ui32(v335) {
									if base.Ui32(v335) <= base.Ui32(int32(2047)) {
										v385 = v335&int32(63) | v335<<(uint(int32(2))%32)&int32(7936) | int32(49280)
									} else {
										if base.Ui32(v335) <= base.Ui32(int32(65535)) {
											v385 = v335<<(uint(int32(4))%32)&int32(983040) | (v335&int32(63) | v335<<(uint(int32(2))%32)&int32(16128)) | int32(14712960)
										} else {
											v384 = v335<<(uint(int32(2))%32)&int32(16128) | (v335<<(uint(int32(6))%32)&int32(117440512) | (v335&int32(63) | v335<<(uint(int32(4))%32)&int32(4128768))) | int32(-260013952)
											v385 = v384
										}
									}
								} else {
									v384 = v335
									v385 = v384
								}
								return v385
							} else {
								if base.Ui32(l0+int32(2110480079)) <= base.Ui32(int32(17922)) {
									v393 = int32(255)
									v401 = int32(base.Ui32(l0)>>(uint(int32(8))%32))&v393*int32(10) + l0&v393 + int32(17213)
									if base.Ui32(int32(128)) <= base.Ui32(v401) {
										if base.Ui32(v401) <= base.Ui32(int32(2047)) {
											v451 = v401&int32(63) | v401<<(uint(int32(2))%32)&int32(7936) | int32(49280)
										} else {
											if base.Ui32(v401) <= base.Ui32(int32(65535)) {
												v451 = v401<<(uint(int32(4))%32)&int32(983040) | (v401&int32(63) | v401<<(uint(int32(2))%32)&int32(16128)) | int32(14712960)
											} else {
												v450 = v401<<(uint(int32(2))%32)&int32(16128) | (v401<<(uint(int32(6))%32)&int32(117440512) | (v401&int32(63) | v401<<(uint(int32(4))%32)&int32(4128768))) | int32(-260013952)
												v451 = v450
											}
										}
									} else {
										v450 = v401
										v451 = v450
									}
									return v451
								} else {
									if base.Ui32(l0+int32(2110419149)) <= base.Ui32(int32(16857093)) {
										v461 = int32(255)
										v479 = int32(base.Ui32(l0)>>(uint(int32(24))%32))*int32(12600) + l0&v461 + int32(base.Ui32(l0)>>(uint(int32(16))%32))&v461*int32(1260) + int32(base.Ui32(l0)>>(uint(int32(8))%32))&v461*int32(10) - int32(1665391)
										if base.Ui32(int32(128)) <= base.Ui32(v479) {
											if base.Ui32(v479) <= base.Ui32(int32(2047)) {
												v529 = v479&int32(63) | v479<<(uint(int32(2))%32)&int32(7936) | int32(49280)
											} else {
												if base.Ui32(v479) <= base.Ui32(int32(65535)) {
													v529 = v479<<(uint(int32(4))%32)&int32(983040) | (v479&int32(63) | v479<<(uint(int32(2))%32)&int32(16128)) | int32(14712960)
												} else {
													v528 = v479<<(uint(int32(2))%32)&int32(16128) | (v479<<(uint(int32(6))%32)&int32(117440512) | (v479&int32(63) | v479<<(uint(int32(4))%32)&int32(4128768))) | int32(-260013952)
													v529 = v528
												}
											}
										} else {
											v528 = v479
											v529 = v528
										}
										return v529
									} else {
										if base.Ui32(l0+int32(2093559760)) <= base.Ui32(int32(16364804)) {
											v539 = int32(255)
											v557 = int32(base.Ui32(l0)>>(uint(int32(24))%32))*int32(12600) + l0&v539 + int32(base.Ui32(l0)>>(uint(int32(16))%32))&v539*int32(1260) + int32(base.Ui32(l0)>>(uint(int32(8))%32))&v539*int32(10) - int32(1661275)
											if base.Ui32(int32(128)) <= base.Ui32(v557) {
												if base.Ui32(v557) <= base.Ui32(int32(2047)) {
													v607 = v557&int32(63) | v557<<(uint(int32(2))%32)&int32(7936) | int32(49280)
												} else {
													if base.Ui32(v557) <= base.Ui32(int32(65535)) {
														v607 = v557<<(uint(int32(4))%32)&int32(983040) | (v557&int32(63) | v557<<(uint(int32(2))%32)&int32(16128)) | int32(14712960)
													} else {
														v606 = v557<<(uint(int32(2))%32)&int32(16128) | (v557<<(uint(int32(6))%32)&int32(117440512) | (v557&int32(63) | v557<<(uint(int32(4))%32)&int32(4128768))) | int32(-260013952)
														v607 = v606
													}
												}
											} else {
												v606 = v557
												v607 = v606
											}
											return v607
										} else {
											if base.Ui32(l0+int32(2077189064)) <= base.Ui32(int32(59647)) {
												v619 = int32(255)
												v630 = int32(base.Ui32(l0)>>(uint(int32(16))%32))&int32(49)*int32(1260) + l0&v619 + int32(base.Ui32(l0)>>(uint(int32(8))%32))&v619*int32(10) + int32(1946)
												if base.Ui32(int32(128)) <= base.Ui32(v630) {
													if base.Ui32(v630) <= base.Ui32(int32(2047)) {
														v680 = v630&int32(63) | v630<<(uint(int32(2))%32)&int32(7936) | int32(49280)
													} else {
														if base.Ui32(v630) <= base.Ui32(int32(65535)) {
															v680 = v630<<(uint(int32(4))%32)&int32(983040) | (v630&int32(63) | v630<<(uint(int32(2))%32)&int32(16128)) | int32(14712960)
														} else {
															v679 = v630<<(uint(int32(2))%32)&int32(16128) | (v630<<(uint(int32(6))%32)&int32(117440512) | (v630&int32(63) | v630<<(uint(int32(4))%32)&int32(4128768))) | int32(-260013952)
															v680 = v679
														}
													}
												} else {
													v679 = v630
													v680 = v679
												}
												return v680
											} else {
												if base.Ui32(l0+int32(2077121996)) <= base.Ui32(int32(517)) {
													v696 = int32(base.Ui32(l0)>>(uint(int32(8))%32))&int32(167)*int32(10) + l0&int32(255) + int32(63838)
													if base.Ui32(int32(128)) <= base.Ui32(v696) {
														if base.Ui32(v696) <= base.Ui32(int32(2047)) {
															v746 = v696&int32(63) | v696<<(uint(int32(2))%32)&int32(7936) | int32(49280)
														} else {
															if base.Ui32(v696) <= base.Ui32(int32(65535)) {
																v746 = v696<<(uint(int32(4))%32)&int32(983040) | (v696&int32(63) | v696<<(uint(int32(2))%32)&int32(16128)) | int32(14712960)
															} else {
																v745 = v696<<(uint(int32(2))%32)&int32(16128) | (v696<<(uint(int32(6))%32)&int32(117440512) | (v696&int32(63) | v696<<(uint(int32(4))%32)&int32(4128768))) | int32(-260013952)
																v746 = v745
															}
														}
													} else {
														v745 = v696
														v746 = v745
													}
													return v746
												} else {
													if base.Ui32(int32(1392646405)) < base.Ui32(l0+int32(1875869392)) {
														v826 = int32(0)
													} else {
														v756 = int32(255)
														v774 = int32(base.Ui32(l0)>>(uint(int32(24))%32))*int32(12600) + l0&v756 + int32(base.Ui32(l0)>>(uint(int32(16))%32))&v756*int32(1260) + int32(base.Ui32(l0)>>(uint(int32(8))%32))&v756*int32(10) - int32(1810682)
														if base.Ui32(int32(128)) <= base.Ui32(v774) {
															if base.Ui32(v774) <= base.Ui32(int32(2047)) {
																v824 = v774&int32(63) | v774<<(uint(int32(2))%32)&int32(7936) | int32(49280)
															} else {
																if base.Ui32(v774) <= base.Ui32(int32(65535)) {
																	v824 = v774<<(uint(int32(4))%32)&int32(983040) | (v774&int32(63) | v774<<(uint(int32(2))%32)&int32(16128)) | int32(14712960)
																} else {
																	v823 = v774<<(uint(int32(2))%32)&int32(16128) | (v774<<(uint(int32(6))%32)&int32(117440512) | (v774&int32(63) | v774<<(uint(int32(4))%32)&int32(4128768))) | int32(-260013952)
																	v824 = v823
																}
															}
														} else {
															v823 = v774
															v824 = v823
														}
														v826 = v824
													}
													return v826
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
	var v74 float64
	_ = v74
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
	*(*float64)(unsafe.Add(mBase, uint32(l0)+56)) = base.F64_add(v115, base.F64_add(base.F64_mul(v109, v101), base.F64_add(base.F64_mul(v107, base.F64_add(v100, base.F64_add(v99, v102))), v113)))
	m.G0 = v17 + int32(32)
	return
L2:
	;
	v19 = *(*float64)(unsafe.Add(mBase, uint32(l3)+8))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+32)) = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	v22 = int32(0)
	v24 = *(*float64)(unsafe.Add(mBase, _consts[589]))
	v25 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v17)+24)) = v25
	*(*int64)(unsafe.Add(mBase, uint32(v17)+16)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = l1
	if v21 == v22 {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	goto L4
L4:
	;
	v86 = *(*float64)(unsafe.Add(mBase, uint32(l2)+16))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+32)) = v86
	v88 = *(*float64)(unsafe.Add(mBase, uint32(l2)+200))
	v90 = *(*float64)(unsafe.Add(mBase, _consts[589]))
	v91 = *(*float64)(unsafe.Add(mBase, uint32(l2)+192))
	v99 = v88
	v100 = v90
	v101 = v86
	v102 = v90
	v106 = v91
	goto L1
L5:
	;
	v82 = *(*float64)(unsafe.Add(mBase, uint32(l2)+200))
	v84 = *(*float64)(unsafe.Add(mBase, uint32(l2)+192))
	v99 = base.F64_add(v74, v82)
	v100 = v24
	v101 = v76
	v102 = v81
	v106 = base.F64_add(v73, v84)
	goto L1
L6:
	;
	v73 = v7
	v74 = v7
	v76 = v19
	v81 = v24
	goto L5
L7:
	;
	goto L8
L8:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	if v32 <= int32(0) {
		v73 = v7
		v74 = v7
		v76 = v19
		v81 = v24
		goto L5
	} else {
		goto L9
	}
L9:
	;
	v38 = v22
	goto L10
L10:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v49+v38<<(uint(int32(2))%32))))
	v56 = F_cost_qual_eval_walker(m, v53, v17+int32(8))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v62 = *(*float64)(unsafe.Add(mBase, uint32(l0)+32))
	v63 = *(*float64)(unsafe.Add(mBase, uint32(v17)+24))
	v64 = *(*float64)(unsafe.Add(mBase, uint32(v17)+16))
	v66 = *(*float64)(unsafe.Add(mBase, _consts[589]))
	v73 = v64
	v74 = v63
	v76 = v62
	v81 = v66
	goto L5
L12:
	;
	return
L13:
	;
	v59 = v38 + int32(1)
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	if v59 < v60 {
		v38 = v59
		goto L10
	} else {
		goto L14
	}
L14:
	;
	goto L11
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
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
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
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v103 float64
	_ = v103
	var v104 int32
	_ = v104
	var v120 float64
	_ = v120
	var v124 int32
	_ = v124
	var v125 float64
	_ = v125
	var v131 float64
	_ = v131
	var v134 float64
	_ = v134
	var v136 float64
	_ = v136
	var v139 float64
	_ = v139
	var v146 float64
	_ = v146
	var v148 float64
	_ = v148
	var v152 int32
	_ = v152
	var v153 float64
	_ = v153
	var v156 int64
	_ = v156
	var v157 float64
	_ = v157
	var v159 float64
	_ = v159
	var v160 int32
	_ = v160
	var v163 float64
	_ = v163
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
	var v177 float64
	_ = v177
	var v180 float64
	_ = v180
	var v184 float64
	_ = v184
	var v191 float64
	_ = v191
	var v192 float64
	_ = v192
	var v195 float64
	_ = v195
	var v199 float64
	_ = v199
	var v209 float64
	_ = v209
	var v212 float64
	_ = v212
	var v216 float64
	_ = v216
	var v217 float64
	_ = v217
	var v218 float64
	_ = v218
	var v222 float64
	_ = v222
	var v224 float64
	_ = v224
	var v232 float64
	_ = v232
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
	v124 = v21 + int32(8)
	v125 = base.F64_div(v26, v120)
	v131 = float64(2)
	if base.F64_lt(v125, v131) != 0 {
		goto L25
	} else {
		goto L26
	}
L5:
	;
	v101 = int32(0)
	v103 = F_estimate_num_groups(m, l1, v96, v26, v101, v101)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
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
	v54 = v12
	v55 = v12
	goto L14
L14:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v61+v54<<(uint(int32(2))%32))))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+16))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+12))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	v71 = F_pull_varnos(m, l1, v70)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v96 = v76
	goto L5
L16:
	;
	return
L17:
	;
	v73 = F_bms_is_member(m, int32(0), v71)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	if v73 != 0 {
		v120 = v35
		goto L4
	} else {
		goto L19
	}
L19:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	v76 = F_lappend(m, v55, v75)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L16
	} else {
		goto L20
	}
L20:
	;
	if v54 == v39-int32(1) {
		v96 = v76
		goto L5
	} else {
		goto L21
	}
L21:
	;
	v80 = v54 + int32(1)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v80 < v81 {
		v54 = v80
		v55 = v76
		goto L14
	} else {
		goto L22
	}
L22:
	;
	goto L15
L23:
	;
	v120 = v103
	goto L4
L24:
	;
	v216 = *(*float64)(unsafe.Add(mBase, _consts[589]))
	v217 = *(*float64)(unsafe.Add(mBase, uint32(v21)))
	v218 = *(*float64)(unsafe.Add(mBase, uint32(v21)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = l4
	*(*float64)(unsafe.Add(mBase, uint32(l0)+32)) = v26
	v222 = base.F64_div(base.F64_sub(l6, l5), v120)
	v224 = base.F64_add(v222, base.F64_add(l5, v218))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+48)) = v224
	v232 = base.F64_add(v120, float64(-1))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+56)) = base.F64_add(v224, base.F64_add(base.F64_mul(base.F64_add(v216, v216), v120), base.F64_add(base.F64_mul(base.F64_add(v216, float64(0)), v26), base.F64_add(base.F64_mul(v222, v232), base.F64_add(v217, base.F64_mul(base.F64_add(v218, v217), v232))))))
	m.G0 = v21 + int32(16)
	return
L25:
	;
	v134 = v131
	goto L27
L26:
	;
	v134 = v125
	goto L27
L27:
	;
	v136 = *(*float64)(unsafe.Add(mBase, _consts[588]))
	v139 = base.F64_mul(v134, base.F64_add(base.F64_add(v136, v136), float64(0)))
	v146 = base.F64_convert_i32_u((l8+int32(7))&int32(-8) + int32(24))
	v148 = base.F64_mul(v125, v146)
	v152 = base.F64_lt(l10, v134) & base.F64_gt(l10, float64(0))
	if v152 != 0 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v124))) = v209
	v212 = *(*float64)(unsafe.Add(mBase, _consts[588]))
	*(*float64)(unsafe.Add(mBase, uint32(v21))) = base.F64_mul(v134, v212)
	goto L24
L29:
	;
	v153 = base.F64_mul(l10, v146)
	goto L31
L30:
	;
	v153 = v148
	goto L31
L31:
	;
	v156 = base.I64_extend_i32_s(l9) << (uint(int64(10)) % 64)
	v157 = base.F64_convert_i64_s(v156)
	if base.F64_gt(v153, v157) != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v159 = F_log(m, v134)
	mBase = m.M
	v160 = F_tuplesort_merge_order(m, v156)
	mBase = m.M
	v163 = base.F64_mul(base.F64_div(v159, float64(0.693147180559945)), v139)
	*(*float64)(unsafe.Add(mBase, uint32(v124))) = v163
	v168 = base.F64_ceil(base.F64_mul(v148, float64(0.0001220703125)))
	v170 = base.F64_div(v148, v157)
	v171 = base.F64_convert_i32_s(v160)
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
	if v152 != 0 {
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
	v177 = base.F64_ceil(base.F64_div(v173, v174))
	goto L37
L36:
	;
	v177 = float64(1)
	goto L37
L37:
	;
	v180 = *(*float64)(unsafe.Add(mBase, _consts[590]))
	v184 = *(*float64)(unsafe.Add(mBase, _consts[591]))
	v209 = base.F64_add(base.F64_mul(base.F64_mul(base.F64_add(v168, v168), v177), base.F64_add(base.F64_mul(v180, float64(0.75)), base.F64_mul(v184, float64(0.25)))), v163)
	goto L28
L38:
	;
	v191 = l10
	goto L40
L39:
	;
	v191 = v134
	goto L40
L40:
	;
	v192 = base.F64_add(v191, v191)
	if base.F64_gt(v148, v157)|base.F64_gt(v134, v192) != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v195 = F_log(m, v192)
	mBase = m.M
	v209 = base.F64_mul(base.F64_div(v195, float64(0.693147180559945)), v139)
	goto L28
L42:
	;
	goto L43
L43:
	;
	v199 = F_log(m, v134)
	mBase = m.M
	v209 = base.F64_mul(base.F64_div(v199, float64(0.693147180559945)), v139)
	goto L28
}
func F_create_hashjoin_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 float64
	_ = v12
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 float64
	_ = v92
	var v93 float64
	_ = v93
	var v95 float64
	_ = v95
	var v97 float64
	_ = v97
	var v99 float64
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 float64
	_ = v109
	var v111 int32
	_ = v111
	var v114 float64
	_ = v114
	var v116 int32
	_ = v116
	var v122 float64
	_ = v122
	var v126 float64
	_ = v126
	var v128 float64
	_ = v128
	var v130 float64
	_ = v130
	var v131 float64
	_ = v131
	var v139 float64
	_ = v139
	var v143 float64
	_ = v143
	var v151 float64
	_ = v151
	var v152 int32
	_ = v152
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v180 int32
	_ = v180
	var v197 float64
	_ = v197
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v318 int32
	_ = v318
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v336 int32
	_ = v336
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v423 int32
	_ = v423
	var v432 int32
	_ = v432
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v530 int32
	_ = v530
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v545 int32
	_ = v545
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v576 int32
	_ = v576
	var v578 int32
	_ = v578
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v620 int32
	_ = v620
	var v622 int32
	_ = v622
	var v640 float64
	_ = v640
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v663 float64
	_ = v663
	var v665 float64
	_ = v665
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v704 int32
	_ = v704
	var v706 int32
	_ = v706
	var v710 int32
	_ = v710
	var v713 int32
	_ = v713
	var v717 int32
	_ = v717
	var v720 int32
	_ = v720
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v731 int32
	_ = v731
	var v734 int32
	_ = v734
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v740 int32
	_ = v740
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v751 int32
	_ = v751
	var v756 int32
	_ = v756
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v771 int32
	_ = v771
	var v788 float64
	_ = v788
	var v807 int32
	_ = v807
	var v824 float64
	_ = v824
	var v846 int32
	_ = v846
	var v888 float64
	_ = v888
	var v889 int32
	_ = v889
	var v898 int32
	_ = v898
	var v904 float64
	_ = v904
	var v929 int32
	_ = v929
	var v933 int32
	_ = v933
	var v934 int32
	_ = v934
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v949 int32
	_ = v949
	var v952 int32
	_ = v952
	var v953 int32
	_ = v953
	var v958 int32
	_ = v958
	var v965 int32
	_ = v965
	var v967 int32
	_ = v967
	var v969 int32
	_ = v969
	var v972 int32
	_ = v972
	var v974 int32
	_ = v974
	var v976 int32
	_ = v976
	var v981 int32
	_ = v981
	var v990 int32
	_ = v990
	var v991 float64
	_ = v991
	var v999 int32
	_ = v999
	var v1000 int32
	_ = v1000
	var v1001 int32
	_ = v1001
	var v1004 int32
	_ = v1004
	var v1007 int32
	_ = v1007
	var v1008 int32
	_ = v1008
	var v1009 int32
	_ = v1009
	var v1010 int32
	_ = v1010
	var v1014 int32
	_ = v1014
	var v1015 float64
	_ = v1015
	var v1016 float64
	_ = v1016
	var v1024 int32
	_ = v1024
	var v1025 int32
	_ = v1025
	var v1026 int32
	_ = v1026
	var v1027 int32
	_ = v1027
	var v1029 int32
	_ = v1029
	var v1030 int32
	_ = v1030
	var v1034 int32
	_ = v1034
	var v1035 float64
	_ = v1035
	var v1036 int32
	_ = v1036
	var v1039 float64
	_ = v1039
	var v1041 float64
	_ = v1041
	var v1042 float64
	_ = v1042
	var v1046 float64
	_ = v1046
	var v1048 int32
	_ = v1048
	var v1049 int32
	_ = v1049
	var v1062 float64
	_ = v1062
	var v1087 float64
	_ = v1087
	var v1088 float64
	_ = v1088
	var v1096 float64
	_ = v1096
	var v1100 float64
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1102 int32
	_ = v1102
	var v1106 float64
	_ = v1106
	var v1108 int32
	_ = v1108
	var v1112 float64
	_ = v1112
	var v1113 float64
	_ = v1113
	var v1116 float64
	_ = v1116
	var v1122 int32
	_ = v1122
	var v1124 int32
	_ = v1124
	var v1126 float64
	_ = v1126
	var v1128 int32
	_ = v1128
	var v1129 int64
	_ = v1129
	var v1134 float64
	_ = v1134
	var v1139 int32
	_ = v1139
	var v1145 int32
	_ = v1145
	var v1178 int32
	_ = v1178
	var v1182 int32
	_ = v1182
	var v1185 int32
	_ = v1185
	var v1186 int32
	_ = v1186
	var v1188 int32
	_ = v1188
	var v1189 int32
	_ = v1189
	var v1191 float64
	_ = v1191
	var v1192 float64
	_ = v1192
	var v1208 float64
	_ = v1208
	var v1229 float64
	_ = v1229
	var v1230 int32
	_ = v1230
	var v1231 int64
	_ = v1231
	var v1240 int32
	_ = v1240
	var v1247 int32
	_ = v1247
	var v1280 int32
	_ = v1280
	var v1284 int32
	_ = v1284
	var v1287 int32
	_ = v1287
	var v1288 int32
	_ = v1288
	var v1290 int32
	_ = v1290
	var v1291 int32
	_ = v1291
	var v1293 float64
	_ = v1293
	var v1294 float64
	_ = v1294
	var v1319 float64
	_ = v1319
	var v1331 float64
	_ = v1331
	var v1332 int32
	_ = v1332
	var v1337 int32
	_ = v1337
	var v1340 float64
	_ = v1340
	var v1341 float64
	_ = v1341
	var v1343 float64
	_ = v1343
	var v1346 float64
	_ = v1346
	var v1349 float64
	_ = v1349
	var v1353 float64
	_ = v1353
	var v1362 float64
	_ = v1362
	var v1366 float64
	_ = v1366
	var v1367 float64
	_ = v1367
	var v1373 float64
	_ = v1373
	var v1381 float64
	_ = v1381
	var v1385 float64
	_ = v1385
	var v1388 float64
	_ = v1388
	var v1394 float64
	_ = v1394
	var v1395 float64
	_ = v1395
	var v1396 float64
	_ = v1396
	var v1404 float64
	_ = v1404
	var v1408 float64
	_ = v1408
	var v1409 int32
	_ = v1409
	var v1410 float64
	_ = v1410
	var v1411 int32
	_ = v1411
	var v1412 float64
	_ = v1412
	var v1414 int32
	_ = v1414
	var v1415 int32
	_ = v1415
	var v1416 int32
	_ = v1416
	var v1417 int32
	_ = v1417
	var v1418 int32
	_ = v1418
	var v1419 int64
	_ = v1419
	var v1439 int32
	_ = v1439
	var v1440 float64
	_ = v1440
	var v1441 int32
	_ = v1441
	var v1447 int32
	_ = v1447
	var v1455 float64
	_ = v1455
	var v1480 int32
	_ = v1480
	var v1484 int32
	_ = v1484
	var v1485 int32
	_ = v1485
	var v1489 float64
	_ = v1489
	var v1490 int32
	_ = v1490
	var v1491 float64
	_ = v1491
	var v1493 int32
	_ = v1493
	var v1494 int32
	_ = v1494
	var v1507 float64
	_ = v1507
	var v1534 float64
	_ = v1534
	var v1536 float64
	_ = v1536
	var v1544 float64
	_ = v1544
	var v1548 float64
	_ = v1548
	var v1562 float64
	_ = v1562
	var v1586 float64
	_ = v1586
	var v1588 float64
	_ = v1588
	var v1589 int32
	_ = v1589
	var v1590 float64
	_ = v1590
	var v1602 float64
	_ = v1602
	var v1606 float64
	_ = v1606
	var v1607 float64
	_ = v1607
	var v1609 float64
	_ = v1609
	v12 = float64(0)
	v37 = m.G0
	v39 = v37 - int32(16)
	m.G0 = v39
	*(*int32)(unsafe.Add(mBase, uint32(v39)+12)) = l8
	v43 = F_palloc0(m, int32(112))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+8)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v43))) = int64(1541893259564)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v43)+12)) = v50
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v55 = F_get_joinrel_parampathinfo(m, l0, l1, l5, l6, v52, l9, v39+int32(12))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+16)) = v55
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	v59 = l7 & v58
	*(*uint8)(unsafe.Add(mBase, uint32(v43)+20)) = uint8(v59)
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	if v62 != int32(1) {
		v70 = int32(0)
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v72 = v70 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v43)+21)) = uint8(v72)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l5)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v43)+72)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v43)+64)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v43)+24)) = v74
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+8)))
	*(*int32)(unsafe.Add(mBase, uint32(v43)+84)) = l6
	*(*int32)(unsafe.Add(mBase, uint32(v43)+80)) = l5
	*(*uint8)(unsafe.Add(mBase, uint32(v43)+76)) = uint8(v79)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v43)+96)) = l10
	*(*int32)(unsafe.Add(mBase, uint32(v43)+88)) = v83
	v86 = m.G0
	v88 = v86 + int32(-64)
	m.G0 = v88
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l3)+84))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l3)+80))
	v92 = *(*float64)(unsafe.Add(mBase, uint32(l3)+24))
	v93 = *(*float64)(unsafe.Add(mBase, uint32(l3)+8))
	v95 = *(*float64)(unsafe.Add(mBase, uint32(l3)+88))
	v97 = *(*float64)(unsafe.Add(mBase, uint32(l6)+32))
	v99 = *(*float64)(unsafe.Add(mBase, uint32(l5)+32))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(v43)+40)) = v100
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v43)+16))
	if v102 != 0 {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5)+21)))
	if v66 != int32(1) {
		v70 = int32(0)
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6)+21)))
	v70 = v69
	goto L4
L7:
	;
	v109 = *(*float64)(unsafe.Add(mBase, uint32(v108)))
	*(*float64)(unsafe.Add(mBase, uint32(v43)+32)) = v109
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v43)+24))
	if int32(0) < v111 {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	v108 = v102 + int32(8)
	goto L7
L9:
	;
	goto L10
L10:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
	v108 = v105 + int32(16)
	goto L7
L11:
	;
	v114 = base.F64_convert_i32_u(v111)
	v116 = int32(*(*uint8)(unsafe.Add(mBase, _consts[586])))
	if v116 == int32(1) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v43)+104)) = v95
	*(*int32)(unsafe.Add(mBase, uint32(v43)+100)) = v90
	v151 = base.F64_mul(base.F64_convert_i32_s(v91), base.F64_convert_i32_s(v90))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	if v152 == int32(295) {
		goto L25
	} else {
		goto L26
	}
L14:
	;
	v122 = base.F64_add(base.F64_mul(v114, float64(-0.3)), float64(1))
	if base.F64_gt(v122, float64(0)) != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v128 = v114
	goto L16
L16:
	;
	v130 = float64(1e+100)
	v131 = base.F64_div(v109, v128)
	if base.F64_gt(v131, v130) != 0 {
		v143 = v130
		goto L20
	} else {
		goto L21
	}
L17:
	;
	v126 = v122
	goto L19
L18:
	;
	v126 = math.Float64frombits(uint64(0x8000000000000000))
	goto L19
L19:
	;
	v128 = base.F64_add(v126, v114)
	goto L16
L20:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v43)+32)) = v143
	goto L13
L21:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v131)&int64(9223372036854775807)) {
		v143 = v130
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v139 = float64(1)
	if base.F64_le(v131, v139) != 0 {
		v143 = v139
		goto L20
	} else {
		goto L23
	}
L23:
	;
	v143 = base.F64_nearest(v131)
	goto L20
L24:
	;
	v1087 = float64(1e+100)
	v1088 = base.F64_mul(v97, v1062)
	if base.F64_gt(v1088, v1087) != 0 {
		v1100 = v1087
		goto L194
	} else {
		goto L195
	}
L25:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v88))) = base.F64_div(float64(1), v151)
	v1062 = float64(0)
	goto L24
L26:
	;
	goto L27
L27:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v88))) = int64(4607182418800017408)
	v162 = m.G0
	v164 = v162 - int32(16)
	m.G0 = v164
	v166 = F_list_copy(m, l10)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	if l10 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	m.G0 = v164 + int32(16)
	if v846 == int32(0) {
		goto L149
	} else {
		goto L150
	}
L30:
	;
	v846 = int32(0)
	goto L29
L31:
	;
	goto L32
L32:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l10)+4))
	if v171 < int32(2) {
		v846 = l10
		goto L29
	} else {
		goto L33
	}
L33:
	;
	if v166 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v88))) = base.F64_div(float64(1), v824)
	v846 = v807
	goto L29
L35:
	;
	v807 = int32(0)
	v824 = float64(1)
	goto L34
L36:
	;
	goto L37
L37:
	;
	v180 = int32(0)
	v197 = float64(1)
	v211 = v166
	goto L38
L38:
	;
	v216 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v164)+12)) = v216
	v218 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v164)+8)) = v218
	v224 = v180
	v225 = v216
	v226 = v218
	v230 = v218
	v255 = v211
	v256 = v218
	goto L40
L39:
	;
	v807 = v771
	v824 = v788
	goto L34
L40:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v255)+4))
	if v230 < v260 {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v164)+8))
	if v612 != 0 {
		goto L105
	} else {
		goto L106
	}
L42:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v255)+12))
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v262+v230<<(uint(int32(2))%32))))
	v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v266)+120)))
	if v269 != 0 {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	v576 = v224
	v578 = v226
	v607 = v255
	v608 = v256
	goto L44
L44:
	;
	goto L41
L45:
	;
	v270 = int32(48)
	goto L47
L46:
	;
	v270 = int32(44)
	goto L47
L47:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v266+v270)))
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v266)+4))
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v273)+28))
	if v269 != 0 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v291 = int32(0)
	if v272 == v291 {
		goto L62
	} else {
		goto L63
	}
L49:
	;
	v275 = int32(0)
	if v274 == v275 {
		v288 = v275
		goto L48
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	if v274 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v274)+4))
	if v278 < int32(2) {
		v288 = v275
		goto L48
	} else {
		goto L53
	}
L53:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v274)+12))
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v281)+4))
	v288 = v282
	goto L48
L54:
	;
	v288 = int32(0)
	goto L48
L55:
	;
	goto L56
L56:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v274)+12))
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v286)))
	v288 = v287
	goto L48
L57:
	;
	if v569 != 0 {
		v224 = v538
		v225 = v539
		v226 = v540
		v230 = v545 + int32(1)
		v255 = v569
		v256 = v570
		goto L40
	} else {
		goto L102
	}
L58:
	;
	v536 = F_list_delete_nth_cell(m, v255, v230)
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L1
	} else {
		goto L101
	}
L59:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v387 = F_remove_nulling_relids(m, v288, v385, int32(0))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L1
	} else {
		goto L87
	}
L60:
	;
	v380 = F_lappend(m, v224, v266)
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L1
	} else {
		goto L86
	}
L61:
	;
	if v344 == int32(0) {
		goto L60
	} else {
		goto L77
	}
L62:
	;
	v344 = int32(0)
	goto L61
L63:
	;
	goto L64
L64:
	;
	v299 = int32(1)
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v272)+4))
	if v300 <= v299 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v303 = v299
	goto L67
L66:
	;
	v303 = v300
	goto L67
L67:
	;
	v308 = int32(0)
	v311 = int32(-1)
	goto L69
L68:
	;
	v344 = v336
	goto L61
L69:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v272+int32(8)+v308<<(uint(int32(2))%32))))
	if v318 != 0 {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v164+int32(12)))) = v328
	v336 = int32(1)
	goto L68
L71:
	;
	if int32(0) <= v311 {
		v336 = v291
		goto L68
	} else {
		goto L74
	}
L72:
	;
	v328 = v311
	goto L73
L73:
	;
	v330 = v308 + int32(1)
	if v330 != v303 {
		v308 = v330
		v311 = v328
		goto L69
	} else {
		goto L76
	}
L74:
	;
	if base.Ui32(int32(1)) < base.Ui32(base.I32_popcnt(v318)) {
		v336 = v291
		goto L68
	} else {
		goto L75
	}
L75:
	;
	v328 = base.I32_ctz(v318) | v308<<(uint(int32(5))%32)
	goto L73
L76:
	;
	goto L70
L77:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v164)+12))
	v349 = v347 << (uint(int32(2)) % 32)
	v350 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v349+v350)))
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v352)+112))
	if v353 == int32(0) {
		goto L60
	} else {
		goto L78
	}
L78:
	;
	if v225 < int32(0) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v358+v349)))
	if v360 == int32(0) {
		goto L60
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	if base.B2i32(v225 != v347) == int32(0) {
		v382 = v226
		v383 = v225
		goto L59
	} else {
		goto L85
	}
L82:
	;
	v363 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v360)+21)))
	v365 = v363 - int32(102)
	if base.Ui32(int32(12)) < base.Ui32(v365) {
		goto L60
	} else {
		goto L83
	}
L83:
	;
	if int32(1)<<(uint(v365)%32)&int32(5249) == int32(0) {
		goto L60
	} else {
		goto L84
	}
L84:
	;
	v382 = v352
	v383 = v347
	goto L59
L85:
	;
	v538 = v224
	v539 = v225
	v540 = v226
	v545 = v230
	v569 = v255
	v570 = v256
	goto L57
L86:
	;
	v498 = v380
	v499 = v225
	v500 = v226
	v530 = v256
	goto L58
L87:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v164)+8))
	if v389 == int32(0) {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v483 = F_palloc0(m, int32(24))
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L1
	} else {
		goto L98
	}
L89:
	;
	v392 = int32(0)
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v389)+4))
	if v393 <= v392 {
		goto L88
	} else {
		goto L90
	}
L90:
	;
	v423 = v392
	goto L91
L91:
	;
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v389)+12))
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v432+v423<<(uint(int32(2))%32))))
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v436)))
	v438 = F_equal(m, v387, v437)
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L1
	} else {
		goto L93
	}
L92:
	;
	v538 = v224
	v539 = v383
	v540 = v382
	v545 = v230
	v569 = v255
	v570 = v256
	goto L57
L93:
	;
	if v438 == int32(0) {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v443 = v423 + int32(1)
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v389)+4))
	if v443 < v444 {
		v423 = v443
		goto L91
	} else {
		goto L97
	}
L95:
	;
	goto L96
L96:
	;
	goto L92
L97:
	;
	goto L88
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v483))) = v387
	v486 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v164)+12))
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v486+v487<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v483)+4)) = v491
	v493 = F_lappend(m, v389, v483)
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v164)+8)) = v493
	v496 = F_lappend(m, v256, v266)
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	v498 = v224
	v499 = v383
	v500 = v382
	v530 = v496
	goto L58
L101:
	;
	v538 = v498
	v539 = v499
	v540 = v500
	v545 = v230 - int32(1)
	v569 = v536
	v570 = v530
	goto L57
L102:
	;
	v576 = v538
	v578 = v540
	v607 = v569
	v608 = v570
	goto L44
L103:
	;
	if v607 != 0 {
		v180 = v771
		v197 = v788
		v211 = v607
		goto L38
	} else {
		goto L148
	}
L104:
	;
	v640 = v197
	goto L112
L105:
	;
	v613 = *(*int32)(unsafe.Add(mBase, uint32(v612)+4))
	if int32(2) <= v613 {
		goto L104
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	v616 = F_list_concat(m, v576, v608)
	mBase = m.M
	v617 = m.ExcPending
	if v617 != 0 {
		goto L1
	} else {
		goto L109
	}
L108:
	;
	goto L107
L109:
	;
	v618 = *(*int32)(unsafe.Add(mBase, uint32(v164)+8))
	F_list_free_deep(m, v618)
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	F_list_free(m, v608)
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	v771 = v616
	v788 = v197
	goto L103
L112:
	;
	v661 = F_estimate_multivariate_ndistinct(m, l0, v578, v164+int32(8), v164)
	mBase = m.M
	v662 = m.ExcPending
	if v662 != 0 {
		goto L1
	} else {
		goto L114
	}
L113:
	;
	v672 = v576
	v673 = int32(0)
	goto L119
L114:
	;
	v663 = *(*float64)(unsafe.Add(mBase, uint32(v164)))
	if base.F64_lt(v640, v663) != 0 {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v665 = v663
	goto L117
L116:
	;
	v665 = v640
	goto L117
L117:
	;
	if v661 != 0 {
		v640 = v665
		goto L112
	} else {
		goto L118
	}
L118:
	;
	goto L113
L119:
	;
	v704 = *(*int32)(unsafe.Add(mBase, uint32(v612)+4))
	if v673 < v704 {
		goto L121
	} else {
		goto L122
	}
L120:
	;
	v771 = v672
	v788 = v640
	goto L103
L121:
	;
	v706 = *(*int32)(unsafe.Add(mBase, uint32(v612)+12))
	v710 = v706 + v673<<(uint(int32(2))%32)
	goto L123
L122:
	;
	v710 = int32(0)
	goto L123
L123:
	;
	if v608 == int32(0) {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v771 = v576
	v788 = v640
	goto L103
L125:
	;
	goto L126
L126:
	;
	v713 = *(*int32)(unsafe.Add(mBase, uint32(v608)+4))
	if v713 <= v673 {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	goto L120
L128:
	;
	if v710 == int32(0) {
		goto L127
	} else {
		goto L129
	}
L129:
	;
	v717 = *(*int32)(unsafe.Add(mBase, uint32(v608)+12))
	v720 = v717 + v673<<(uint(int32(2))%32)
	if v720 == int32(0) {
		goto L127
	} else {
		goto L130
	}
L130:
	;
	v723 = *(*int32)(unsafe.Add(mBase, uint32(v164)+8))
	v724 = *(*int32)(unsafe.Add(mBase, uint32(v710)))
	v725 = int32(0)
	if v723 == v725 {
		goto L132
	} else {
		goto L133
	}
L131:
	;
	if v763 != 0 {
		goto L144
	} else {
		goto L145
	}
L132:
	;
	v763 = int32(0)
	goto L131
L133:
	;
	goto L134
L134:
	;
	v731 = *(*int32)(unsafe.Add(mBase, uint32(v723)+4))
	if v731 <= int32(0) {
		v756 = v725
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v763 = v756
	goto L131
L136:
	;
	v734 = int32(0)
	if v734 < v731 {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v737 = v731
	goto L139
L138:
	;
	v737 = v734
	goto L139
L139:
	;
	v738 = *(*int32)(unsafe.Add(mBase, uint32(v723)+12))
	v740 = int32(0)
	goto L140
L140:
	;
	v748 = *(*int32)(unsafe.Add(mBase, uint32(v738+v740<<(uint(int32(2))%32))))
	v749 = base.B2i32(v748 == v724)
	if v748 == v724 {
		v756 = v749
		goto L135
	} else {
		goto L142
	}
L141:
	;
	v756 = v749
	goto L135
L142:
	;
	v751 = v740 + int32(1)
	if v751 != v737 {
		v740 = v751
		goto L140
	} else {
		goto L143
	}
L143:
	;
	goto L141
L144:
	;
	v764 = *(*int32)(unsafe.Add(mBase, uint32(v720)))
	v765 = F_lappend(m, v672, v764)
	mBase = m.M
	v766 = m.ExcPending
	if v766 != 0 {
		goto L1
	} else {
		goto L147
	}
L145:
	;
	v767 = v672
	goto L146
L146:
	;
	v672 = v767
	v673 = v673 + int32(1)
	goto L119
L147:
	;
	v767 = v765
	goto L146
L148:
	;
	goto L39
L149:
	;
	v1062 = float64(1)
	goto L24
L150:
	;
	goto L151
L151:
	;
	v888 = float64(1)
	v889 = *(*int32)(unsafe.Add(mBase, uint32(v846)+4))
	if v889 <= int32(0) {
		v1062 = v888
		goto L24
	} else {
		goto L152
	}
L152:
	;
	v898 = int32(0)
	v904 = v888
	goto L153
L153:
	;
	v929 = *(*int32)(unsafe.Add(mBase, uint32(v846)+12))
	v933 = *(*int32)(unsafe.Add(mBase, uint32(v929+v898<<(uint(int32(2))%32))))
	v934 = *(*int32)(unsafe.Add(mBase, uint32(v933)+48))
	v935 = *(*int32)(unsafe.Add(mBase, uint32(l6)+8))
	v936 = *(*int32)(unsafe.Add(mBase, uint32(v935)+8))
	v937 = int32(0)
	if v934 == v937 {
		goto L157
	} else {
		goto L158
	}
L154:
	;
	v1062 = v1046
	goto L24
L155:
	;
	v1041 = *(*float64)(unsafe.Add(mBase, uint32(v1036+v933)))
	v1042 = *(*float64)(unsafe.Add(mBase, uint32(v88)))
	if base.F64_lt(v1039, v1042) != 0 {
		goto L187
	} else {
		goto L188
	}
L156:
	;
	if v990 != 0 {
		goto L170
	} else {
		goto L171
	}
L157:
	;
	v990 = int32(1)
	goto L156
L158:
	;
	goto L159
L159:
	;
	if v936 == int32(0) {
		v981 = v937
		goto L160
	} else {
		goto L161
	}
L160:
	;
	v990 = v981
	goto L156
L161:
	;
	v946 = *(*int32)(unsafe.Add(mBase, uint32(v934)+4))
	v947 = *(*int32)(unsafe.Add(mBase, uint32(v936)+4))
	if v947 < v946 {
		v981 = v937
		goto L160
	} else {
		goto L162
	}
L162:
	;
	v949 = int32(1)
	if v946 <= v949 {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	v952 = v949
	goto L165
L164:
	;
	v952 = v946
	goto L165
L165:
	;
	v953 = int32(8)
	v958 = int32(0)
	goto L166
L166:
	;
	v965 = v958 << (uint(int32(2)) % 32)
	v967 = *(*int32)(unsafe.Add(mBase, uint32(v934+v953+v965)))
	v969 = *(*int32)(unsafe.Add(mBase, uint32(v965+(v936+v953))))
	v972 = v967 & (v969 ^ int32(-1))
	v974 = base.B2i32(v972 == int32(0))
	if v972 != 0 {
		v981 = v974
		goto L160
	} else {
		goto L168
	}
L167:
	;
	v981 = v974
	goto L160
L168:
	;
	v976 = v958 + int32(1)
	if v976 != v952 {
		v958 = v976
		goto L166
	} else {
		goto L169
	}
L169:
	;
	goto L167
L170:
	;
	v991 = *(*float64)(unsafe.Add(mBase, uint32(v933)+136))
	if base.F64_lt(v991, float64(0)) == int32(0) {
		goto L173
	} else {
		goto L174
	}
L171:
	;
	goto L172
L172:
	;
	v1016 = *(*float64)(unsafe.Add(mBase, uint32(v933)+128))
	if base.F64_lt(v1016, float64(0)) == int32(0) {
		goto L180
	} else {
		goto L181
	}
L173:
	;
	v1036 = int32(152)
	v1039 = v991
	goto L155
L174:
	;
	goto L175
L175:
	;
	v999 = int32(0)
	v1000 = *(*int32)(unsafe.Add(mBase, uint32(v933)+4))
	v1001 = *(*int32)(unsafe.Add(mBase, uint32(v1000)+28))
	if v1001 == v999 {
		v1009 = v999
		goto L176
	} else {
		goto L177
	}
L176:
	;
	v1010 = int32(152)
	F_estimate_hash_bucket_stats(m, l0, v1009, v151, v933+v1010, v933+int32(136))
	mBase = m.M
	v1014 = m.ExcPending
	if v1014 != 0 {
		goto L1
	} else {
		goto L179
	}
L177:
	;
	v1004 = *(*int32)(unsafe.Add(mBase, uint32(v1001)+4))
	if v1004 < int32(2) {
		v1009 = v999
		goto L176
	} else {
		goto L178
	}
L178:
	;
	v1007 = *(*int32)(unsafe.Add(mBase, uint32(v1001)+12))
	v1008 = *(*int32)(unsafe.Add(mBase, uint32(v1007)+4))
	v1009 = v1008
	goto L176
L179:
	;
	v1015 = *(*float64)(unsafe.Add(mBase, uint32(v933)+136))
	v1036 = v1010
	v1039 = v1015
	goto L155
L180:
	;
	v1036 = int32(144)
	v1039 = v1016
	goto L155
L181:
	;
	goto L182
L182:
	;
	v1024 = *(*int32)(unsafe.Add(mBase, uint32(v933)+4))
	v1025 = *(*int32)(unsafe.Add(mBase, uint32(v1024)+28))
	if v1025 != 0 {
		goto L183
	} else {
		goto L184
	}
L183:
	;
	v1026 = *(*int32)(unsafe.Add(mBase, uint32(v1025)+12))
	v1027 = *(*int32)(unsafe.Add(mBase, uint32(v1026)))
	v1029 = v1027
	goto L185
L184:
	;
	v1029 = int32(0)
	goto L185
L185:
	;
	v1030 = int32(144)
	F_estimate_hash_bucket_stats(m, l0, v1029, v151, v933+v1030, v933+int32(128))
	mBase = m.M
	v1034 = m.ExcPending
	if v1034 != 0 {
		goto L1
	} else {
		goto L186
	}
L186:
	;
	v1035 = *(*float64)(unsafe.Add(mBase, uint32(v933)+128))
	v1036 = v1030
	v1039 = v1035
	goto L155
L187:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v88))) = v1039
	goto L189
L188:
	;
	goto L189
L189:
	;
	if base.F64_gt(v904, v1041) != 0 {
		goto L190
	} else {
		goto L191
	}
L190:
	;
	v1046 = v1041
	goto L192
L191:
	;
	v1046 = v904
	goto L192
L192:
	;
	v1048 = v898 + int32(1)
	v1049 = *(*int32)(unsafe.Add(mBase, uint32(v846)+4))
	if v1048 < v1049 {
		v898 = v1048
		v904 = v1046
		goto L153
	} else {
		goto L193
	}
L193:
	;
	goto L154
L194:
	;
	v1101 = *(*int32)(unsafe.Add(mBase, uint32(l6)+12))
	v1102 = *(*int32)(unsafe.Add(mBase, uint32(v1101)+32))
	v1106 = *(*float64)(unsafe.Add(mBase, _consts[528]))
	v1108 = *(*int32)(unsafe.Add(mBase, _consts[523]))
	v1112 = base.F64_mul(base.F64_mul(v1106, base.F64_convert_i32_s(v1108)), float64(1024))
	v1113 = float64(4.294967295e+09)
	if base.F64_lt(v1112, v1113) != 0 {
		goto L199
	} else {
		goto L200
	}
L195:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1088)&int64(9223372036854775807)) {
		v1100 = v1087
		goto L194
	} else {
		goto L196
	}
L196:
	;
	v1096 = float64(1)
	if base.F64_le(v1088, v1096) != 0 {
		v1100 = v1096
		goto L194
	} else {
		goto L197
	}
L197:
	;
	v1100 = base.F64_nearest(v1088)
	goto L194
L198:
	;
	v1126 = *(*float64)(unsafe.Add(mBase, _consts[602]))
	v1128 = v88 + int32(24)
	v1129 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1128))) = v1129
	*(*int64)(unsafe.Add(mBase, uint32(v88)+16)) = v1129
	*(*int32)(unsafe.Add(mBase, uint32(v88)+8)) = l0
	v1134 = float64(0)
	if l10 == int32(0) {
		v1208 = v1134
		v1229 = v1134
		goto L205
	} else {
		goto L206
	}
L199:
	;
	v1116 = v1112
	goto L201
L200:
	;
	v1116 = v1113
	goto L201
L201:
	;
	if base.F64_lt(v1116, float64(4.294967296e+09))&base.F64_ge(v1116, float64(0)) != 0 {
		goto L202
	} else {
		goto L203
	}
L202:
	;
	v1122 = base.I32_trunc_f64_u(v1116)
	v1124 = v1122
	goto L198
L203:
	;
	goto L204
L204:
	;
	v1124 = int32(0)
	goto L198
L205:
	;
	v1230 = *(*int32)(unsafe.Add(mBase, uint32(v43)+88))
	v1231 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1128))) = v1231
	*(*int64)(unsafe.Add(mBase, uint32(v88)+16)) = v1231
	*(*int32)(unsafe.Add(mBase, uint32(v88)+8)) = l0
	if v1230 == int32(0) {
		v1319 = v12
		v1331 = float64(0)
		goto L212
	} else {
		goto L213
	}
L206:
	;
	v1139 = *(*int32)(unsafe.Add(mBase, uint32(l10)+4))
	if v1139 <= int32(0) {
		v1208 = v1134
		v1229 = float64(0)
		goto L205
	} else {
		goto L207
	}
L207:
	;
	v1145 = int32(0)
	goto L208
L208:
	;
	v1178 = *(*int32)(unsafe.Add(mBase, uint32(l10)+12))
	v1182 = *(*int32)(unsafe.Add(mBase, uint32(v1178+v1145<<(uint(int32(2))%32))))
	v1185 = F_cost_qual_eval_walker(m, v1182, v88+int32(8))
	mBase = m.M
	v1186 = m.ExcPending
	if v1186 != 0 {
		goto L1
	} else {
		goto L210
	}
L209:
	;
	v1191 = *(*float64)(unsafe.Add(mBase, uint32(v88)+16))
	v1192 = *(*float64)(unsafe.Add(mBase, uint32(v88)+24))
	v1208 = v1191
	v1229 = v1192
	goto L205
L210:
	;
	v1188 = v1145 + int32(1)
	v1189 = *(*int32)(unsafe.Add(mBase, uint32(l10)+4))
	if v1188 < v1189 {
		v1145 = v1188
		goto L208
	} else {
		goto L211
	}
L211:
	;
	goto L209
L212:
	;
	v1332 = *(*int32)(unsafe.Add(mBase, uint32(v43)+72))
	if v1332&int32(-2) != int32(4) {
		goto L221
	} else {
		goto L222
	}
L213:
	;
	v1240 = *(*int32)(unsafe.Add(mBase, uint32(v1230)+4))
	if v1240 <= int32(0) {
		v1319 = v12
		v1331 = float64(0)
		goto L212
	} else {
		goto L214
	}
L214:
	;
	v1247 = int32(0)
	goto L215
L215:
	;
	v1280 = *(*int32)(unsafe.Add(mBase, uint32(v1230)+12))
	v1284 = *(*int32)(unsafe.Add(mBase, uint32(v1280+v1247<<(uint(int32(2))%32))))
	v1287 = F_cost_qual_eval_walker(m, v1284, v88+int32(8))
	mBase = m.M
	v1288 = m.ExcPending
	if v1288 != 0 {
		goto L1
	} else {
		goto L217
	}
L216:
	;
	v1293 = *(*float64)(unsafe.Add(mBase, uint32(v88)+24))
	v1294 = *(*float64)(unsafe.Add(mBase, uint32(v88)+16))
	v1319 = v1293
	v1331 = v1294
	goto L212
L217:
	;
	v1290 = v1247 + int32(1)
	v1291 = *(*int32)(unsafe.Add(mBase, uint32(v1230)+4))
	if v1290 < v1291 {
		v1247 = v1290
		goto L215
	} else {
		goto L218
	}
L218:
	;
	goto L216
L219:
	;
	v1588 = *(*float64)(unsafe.Add(mBase, _consts[589]))
	v1589 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
	v1590 = *(*float64)(unsafe.Add(mBase, uint32(v1589)+24))
	if base.F64_lt(base.F64_convert_i32_u(v1124), base.F64_mul(v1100, base.F64_convert_i32_u((v1102+int32(7))&int32(-8)+int32(24)))) != 0 {
		goto L254
	} else {
		goto L255
	}
L220:
	;
	v1394 = float64(1e+100)
	v1395 = *(*float64)(unsafe.Add(mBase, uint32(v88)))
	v1396 = base.F64_mul(v97, v1395)
	if base.F64_gt(v1396, v1394) != 0 {
		v1408 = v1394
		goto L236
	} else {
		goto L237
	}
L221:
	;
	v1337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+8)))
	if v1337 != int32(1) {
		goto L220
	} else {
		goto L224
	}
L222:
	;
	goto L223
L223:
	;
	v1340 = float64(1e+100)
	v1341 = *(*float64)(unsafe.Add(mBase, uint32(l4)+16))
	v1343 = base.F64_nearest(base.F64_mul(v99, v1341))
	v1346 = *(*float64)(unsafe.Add(mBase, uint32(v88)))
	v1349 = *(*float64)(unsafe.Add(mBase, uint32(l4)+24))
	v1353 = base.F64_mul(base.F64_mul(v97, v1346), base.F64_div(float64(2), base.F64_add(v1349, float64(1))))
	if base.F64_gt(v1353, v1340) != 0 {
		v1366 = v1340
		goto L225
	} else {
		goto L226
	}
L224:
	;
	goto L223
L225:
	;
	v1367 = base.F64_sub(v99, v1343)
	v1373 = base.F64_div(v97, v151)
	if base.F64_gt(v1373, float64(1e+100)) != 0 {
		v1385 = v1340
		goto L229
	} else {
		goto L230
	}
L226:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1353)&int64(9223372036854775807)) {
		v1366 = float64(1e+100)
		goto L225
	} else {
		goto L227
	}
L227:
	;
	v1362 = float64(1)
	if base.F64_le(v1353, v1362) != 0 {
		v1366 = v1362
		goto L225
	} else {
		goto L228
	}
L228:
	;
	v1366 = base.F64_nearest(v1353)
	goto L225
L229:
	;
	if v1332 == int32(5) {
		goto L233
	} else {
		goto L234
	}
L230:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1373)&int64(9223372036854775807)) {
		v1385 = v1340
		goto L229
	} else {
		goto L231
	}
L231:
	;
	v1381 = float64(1)
	if base.F64_le(v1373, v1381) != 0 {
		v1385 = v1381
		goto L229
	} else {
		goto L232
	}
L232:
	;
	v1385 = base.F64_nearest(v1373)
	goto L229
L233:
	;
	v1388 = v1367
	goto L235
L234:
	;
	v1388 = v1343
	goto L235
L235:
	;
	v1562 = v1388
	v1586 = base.F64_add(base.F64_mul(base.F64_mul(base.F64_mul(v1229, v1367), v1385), float64(0.05)), base.F64_add(base.F64_mul(base.F64_mul(base.F64_mul(v1229, v1343), v1366), float64(0.5)), v92))
	goto L219
L236:
	;
	v1409 = *(*int32)(unsafe.Add(mBase, uint32(v43)+84))
	v1410 = *(*float64)(unsafe.Add(mBase, uint32(v1409)+32))
	v1411 = *(*int32)(unsafe.Add(mBase, uint32(v43)+80))
	v1412 = *(*float64)(unsafe.Add(mBase, uint32(v1411)+32))
	v1414 = v88 + int32(8)
	v1415 = *(*int32)(unsafe.Add(mBase, uint32(v1411)+8))
	v1416 = *(*int32)(unsafe.Add(mBase, uint32(v1415)+8))
	v1417 = *(*int32)(unsafe.Add(mBase, uint32(v1409)+8))
	v1418 = *(*int32)(unsafe.Add(mBase, uint32(v1417)+8))
	v1419 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1414)+48)) = v1419
	*(*int32)(unsafe.Add(mBase, uint32(v1414)+16)) = v1418
	*(*int32)(unsafe.Add(mBase, uint32(v1414)+12)) = v1416
	*(*int32)(unsafe.Add(mBase, uint32(v1414)+8)) = v1418
	*(*int32)(unsafe.Add(mBase, uint32(v1414)+4)) = v1416
	*(*int32)(unsafe.Add(mBase, uint32(v1414))) = int32(320)
	*(*int64)(unsafe.Add(mBase, uint32(v1414)+20)) = v1419
	*(*int64)(unsafe.Add(mBase, uint32(v1414)+28)) = v1419
	*(*int64)(unsafe.Add(mBase, uint32(v1414)+36)) = v1419
	*(*int32)(unsafe.Add(mBase, uint32(v1414)+43)) = int32(0)
	goto L240
L237:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1396)&int64(9223372036854775807)) {
		v1408 = v1394
		goto L236
	} else {
		goto L238
	}
L238:
	;
	v1404 = float64(1)
	if base.F64_le(v1396, v1404) != 0 {
		v1408 = v1404
		goto L236
	} else {
		goto L239
	}
L239:
	;
	v1408 = base.F64_nearest(v1396)
	goto L236
L240:
	;
	if l10 == int32(0) {
		goto L242
	} else {
		goto L243
	}
L241:
	;
	v1534 = float64(1e+100)
	v1536 = base.F64_mul(v1410, base.F64_mul(v1412, v1507))
	if base.F64_gt(v1536, v1534) != 0 {
		v1548 = v1534
		goto L250
	} else {
		goto L251
	}
L242:
	;
	v1507 = float64(1)
	goto L241
L243:
	;
	goto L244
L244:
	;
	v1439 = int32(0)
	v1440 = float64(1)
	v1441 = *(*int32)(unsafe.Add(mBase, uint32(l10)+4))
	if v1441 <= v1439 {
		v1507 = v1440
		goto L241
	} else {
		goto L245
	}
L245:
	;
	v1447 = v1439
	v1455 = v1440
	goto L246
L246:
	;
	v1480 = *(*int32)(unsafe.Add(mBase, uint32(l10)+12))
	v1484 = *(*int32)(unsafe.Add(mBase, uint32(v1480+v1447<<(uint(int32(2))%32))))
	v1485 = int32(0)
	v1489 = F_clause_selectivity(m, l0, v1484, v1485, v1485, v88+int32(8))
	mBase = m.M
	v1490 = m.ExcPending
	if v1490 != 0 {
		goto L1
	} else {
		goto L248
	}
L247:
	;
	v1507 = v1491
	goto L241
L248:
	;
	v1491 = base.F64_mul(v1455, v1489)
	v1493 = v1447 + int32(1)
	v1494 = *(*int32)(unsafe.Add(mBase, uint32(l10)+4))
	if v1493 < v1494 {
		v1447 = v1493
		v1455 = v1491
		goto L246
	} else {
		goto L249
	}
L249:
	;
	goto L247
L250:
	;
	v1562 = v1548
	v1586 = base.F64_add(base.F64_mul(base.F64_mul(base.F64_mul(v99, v1229), v1408), float64(0.5)), v92)
	goto L219
L251:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1536)&int64(9223372036854775807)) {
		v1548 = v1534
		goto L250
	} else {
		goto L252
	}
L252:
	;
	v1544 = float64(1)
	if base.F64_le(v1536, v1544) != 0 {
		v1548 = v1544
		goto L250
	} else {
		goto L253
	}
L253:
	;
	v1548 = base.F64_nearest(v1536)
	goto L250
L254:
	;
	v1602 = base.F64_add(v93, v1126)
	goto L256
L255:
	;
	v1602 = v93
	goto L256
L256:
	;
	v1606 = *(*float64)(unsafe.Add(mBase, uint32(v1589)+16))
	v1607 = base.F64_add(base.F64_add(base.F64_add(v1602, v1208), base.F64_sub(v1331, v1208)), v1606)
	*(*float64)(unsafe.Add(mBase, uint32(v43)+48)) = v1607
	v1609 = *(*float64)(unsafe.Add(mBase, uint32(v43)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v43)+56)) = base.F64_add(v1607, base.F64_add(base.F64_mul(v1590, v1609), base.F64_add(base.F64_mul(base.F64_add(v1588, base.F64_sub(v1319, v1229)), v1562), v1586)))
	m.G0 = v88 - int32(-64)
	m.G0 = v39 + int32(16)
	return v43
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
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
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
	var v136 float64
	_ = v136
	var v140 float64
	_ = v140
	var v145 int32
	_ = v145
	var v146 int64
	_ = v146
	var v151 float64
	_ = v151
	var v156 int32
	_ = v156
	var v162 int32
	_ = v162
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 float64
	_ = v205
	var v206 float64
	_ = v206
	var v223 float64
	_ = v223
	var v239 float64
	_ = v239
	var v240 int32
	_ = v240
	var v241 int64
	_ = v241
	var v250 int32
	_ = v250
	var v256 int32
	_ = v256
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
	var v299 float64
	_ = v299
	var v300 float64
	_ = v300
	var v324 float64
	_ = v324
	var v333 float64
	_ = v333
	var v334 int32
	_ = v334
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v360 float64
	_ = v360
	var v361 int32
	_ = v361
	var v362 float64
	_ = v362
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int64
	_ = v369
	var v388 int32
	_ = v388
	var v389 float64
	_ = v389
	var v390 int32
	_ = v390
	var v395 int32
	_ = v395
	var v407 float64
	_ = v407
	var v425 int32
	_ = v425
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v434 float64
	_ = v434
	var v435 int32
	_ = v435
	var v436 float64
	_ = v436
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v455 float64
	_ = v455
	var v475 float64
	_ = v475
	var v477 float64
	_ = v477
	var v485 float64
	_ = v485
	var v489 float64
	_ = v489
	var v491 float64
	_ = v491
	var v492 int32
	_ = v492
	var v493 float64
	_ = v493
	var v494 int32
	_ = v494
	var v499 float64
	_ = v499
	var v505 float64
	_ = v505
	var v508 float64
	_ = v508
	var v509 float64
	_ = v509
	var v510 int32
	_ = v510
	var v514 int32
	_ = v514
	var v518 float64
	_ = v518
	var v521 float64
	_ = v521
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v559 int32
	_ = v559
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v569 int32
	_ = v569
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v576 int32
	_ = v576
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v584 int32
	_ = v584
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v591 int32
	_ = v591
	var v597 int32
	_ = v597
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v648 int32
	_ = v648
	var v663 float64
	_ = v663
	var v678 int32
	_ = v678
	var v679 float64
	_ = v679
	var v681 float64
	_ = v681
	var v689 float64
	_ = v689
	var v690 float64
	_ = v690
	var v692 float64
	_ = v692
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
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v39)+16))
	if v96 != 0 {
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
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v39)+104))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v39)+96))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v39)+80))
	v106 = *(*float64)(unsafe.Add(mBase, uint32(v102)))
	*(*float64)(unsafe.Add(mBase, uint32(v39)+32)) = v106
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v39)+24))
	if int32(0) < v108 {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	v102 = v96 + int32(8)
	goto L7
L9:
	;
	goto L10
L10:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v39)+8))
	v102 = v99 + int32(16)
	goto L7
L11:
	;
	v111 = base.F64_convert_i32_u(v108)
	v113 = int32(*(*uint8)(unsafe.Add(mBase, _consts[586])))
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
	v145 = v81 + int32(-40)
	v146 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v145))) = v146
	*(*int64)(unsafe.Add(mBase, uint32(v83)+16)) = v146
	*(*int32)(unsafe.Add(mBase, uint32(v83)+8)) = l0
	v151 = float64(0)
	if v104 == int32(0) {
		v223 = v151
		v239 = v151
		goto L24
	} else {
		goto L25
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
	if base.F64_gt(v128, v127) != 0 {
		v140 = v127
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
	*(*float64)(unsafe.Add(mBase, uint32(v39)+32)) = v140
	goto L13
L21:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v128)&int64(9223372036854775807)) {
		v140 = v127
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v136 = float64(1)
	if base.F64_le(v128, v136) != 0 {
		v140 = v136
		goto L20
	} else {
		goto L23
	}
L23:
	;
	v140 = base.F64_nearest(v128)
	goto L20
L24:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v39)+88))
	v241 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v145))) = v241
	*(*int64)(unsafe.Add(mBase, uint32(v83)+16)) = v241
	*(*int32)(unsafe.Add(mBase, uint32(v83)+8)) = l0
	if v240 == int32(0) {
		v324 = v15
		v333 = float64(0)
		goto L31
	} else {
		goto L32
	}
L25:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v104)+4))
	if v156 <= int32(0) {
		v223 = v151
		v239 = float64(0)
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v162 = int32(0)
	goto L27
L27:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v104)+12))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v192+v162<<(uint(int32(2))%32))))
	v199 = F_cost_qual_eval_walker(m, v196, v81+int32(-56))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L29
	}
L28:
	;
	v205 = *(*float64)(unsafe.Add(mBase, uint32(v83)+24))
	v206 = *(*float64)(unsafe.Add(mBase, uint32(v83)+16))
	v223 = v205
	v239 = v206
	goto L24
L29:
	;
	v202 = v162 + int32(1)
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v104)+4))
	if v202 < v203 {
		v162 = v202
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v39)+72))
	if v334&int32(-2) != int32(4) {
		goto L40
	} else {
		goto L41
	}
L32:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v240)+4))
	if v250 <= int32(0) {
		v324 = v15
		v333 = float64(0)
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v256 = int32(0)
	goto L34
L34:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v240)+12))
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v286+v256<<(uint(int32(2))%32))))
	v293 = F_cost_qual_eval_walker(m, v290, v81+int32(-56))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L1
	} else {
		goto L36
	}
L35:
	;
	v299 = *(*float64)(unsafe.Add(mBase, uint32(v83)+24))
	v300 = *(*float64)(unsafe.Add(mBase, uint32(v83)+16))
	v324 = v299
	v333 = v300
	goto L31
L36:
	;
	v296 = v256 + int32(1)
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v240)+4))
	if v296 < v297 {
		v256 = v296
		goto L34
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v39)+112)) = uint8(v357)
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v39)+84))
	v360 = *(*float64)(unsafe.Add(mBase, uint32(v359)+32))
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v39)+80))
	v362 = *(*float64)(unsafe.Add(mBase, uint32(v361)+32))
	v364 = v81 + int32(-56)
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v361)+8))
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v365)+8))
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v359)+8))
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v367)+8))
	v369 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v364)+48)) = v369
	*(*int32)(unsafe.Add(mBase, uint32(v364)+16)) = v368
	*(*int32)(unsafe.Add(mBase, uint32(v364)+12)) = v366
	*(*int32)(unsafe.Add(mBase, uint32(v364)+8)) = v368
	*(*int32)(unsafe.Add(mBase, uint32(v364)+4)) = v366
	*(*int32)(unsafe.Add(mBase, uint32(v364))) = int32(320)
	*(*int64)(unsafe.Add(mBase, uint32(v364)+20)) = v369
	*(*int64)(unsafe.Add(mBase, uint32(v364)+28)) = v369
	*(*int64)(unsafe.Add(mBase, uint32(v364)+36)) = v369
	*(*int32)(unsafe.Add(mBase, uint32(v364)+43)) = int32(0)
	goto L51
L39:
	;
	v357 = int32(0)
	goto L38
L40:
	;
	v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+8)))
	if v339 != int32(1) {
		goto L39
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v39)+88))
	if v343 != 0 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	goto L42
L44:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v343)+4))
	v345 = v344
	goto L46
L45:
	;
	v345 = int32(0)
	goto L46
L46:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v39)+96))
	if v347 != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v347)+4))
	v350 = v348
	goto L49
L48:
	;
	v350 = int32(0)
	goto L49
L49:
	;
	if v350 == v345 {
		v357 = int32(1)
		goto L38
	} else {
		goto L50
	}
L50:
	;
	goto L39
L51:
	;
	if v104 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v475 = float64(1e+100)
	v477 = base.F64_mul(v360, base.F64_mul(v362, v455))
	if base.F64_gt(v477, v475) != 0 {
		v489 = v475
		goto L61
	} else {
		goto L62
	}
L53:
	;
	v455 = float64(1)
	goto L52
L54:
	;
	goto L55
L55:
	;
	v388 = int32(0)
	v389 = float64(1)
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v104)+4))
	if v390 <= v388 {
		v455 = v389
		goto L52
	} else {
		goto L56
	}
L56:
	;
	v395 = v388
	v407 = v389
	goto L57
L57:
	;
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v104)+12))
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v425+v395<<(uint(int32(2))%32))))
	v430 = int32(0)
	v434 = F_clause_selectivity(m, l0, v429, v430, v430, v81+int32(-56))
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L1
	} else {
		goto L59
	}
L58:
	;
	v455 = v436
	goto L52
L59:
	;
	v436 = base.F64_mul(v407, v434)
	v438 = v395 + int32(1)
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v104)+4))
	if v438 < v439 {
		v395 = v438
		v407 = v436
		goto L57
	} else {
		goto L60
	}
L60:
	;
	goto L58
L61:
	;
	if base.F64_le(v93, float64(0)) != 0 {
		goto L65
	} else {
		goto L66
	}
L62:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v477)&int64(9223372036854775807)) {
		v489 = v475
		goto L61
	} else {
		goto L63
	}
L63:
	;
	v485 = float64(1)
	if base.F64_le(v477, v485) != 0 {
		v489 = v485
		goto L61
	} else {
		goto L64
	}
L64:
	;
	v489 = base.F64_nearest(v477)
	goto L61
L65:
	;
	v491 = float64(1)
	goto L67
L66:
	;
	v491 = v93
	goto L67
L67:
	;
	v492 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+112)))
	v493 = float64(0)
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
	if v494 == int32(295) {
		v505 = v493
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v508 = base.F64_add(base.F64_div(v505, v87), float64(1))
	v509 = base.F64_mul(v89, v508)
	v510 = int32(1)
	if v492&v510 != 0 {
		goto L73
	} else {
		goto L74
	}
L69:
	;
	if v492&int32(1) != 0 {
		v505 = v493
		goto L68
	} else {
		goto L70
	}
L70:
	;
	v499 = base.F64_sub(v489, v491)
	if base.F64_lt(v499, float64(0)) == int32(0) {
		v505 = v499
		goto L68
	} else {
		goto L71
	}
L71:
	;
	v505 = float64(0)
	goto L68
L72:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v39)+113)) = uint8(v648)
	v678 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
	v679 = *(*float64)(unsafe.Add(mBase, uint32(v678)+24))
	v681 = *(*float64)(unsafe.Add(mBase, _consts[589]))
	v689 = *(*float64)(unsafe.Add(mBase, uint32(v678)+16))
	v690 = base.F64_add(base.F64_add(base.F64_sub(v333, v239), base.F64_add(base.F64_mul(v223, base.F64_add(base.F64_mul(v85, v508), v86)), base.F64_add(v91, v239))), v689)
	*(*float64)(unsafe.Add(mBase, uint32(v39)+48)) = v690
	v692 = *(*float64)(unsafe.Add(mBase, uint32(v39)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v39)+56)) = base.F64_add(v690, base.F64_add(base.F64_mul(v679, v692), base.F64_add(base.F64_mul(base.F64_add(v681, base.F64_sub(v324, v223)), v489), base.F64_add(base.F64_mul(v223, base.F64_add(base.F64_mul(base.F64_sub(v87, v85), v508), base.F64_sub(v88, v86))), base.F64_add(v90, v663)))))
	m.G0 = v83 - int32(-64)
	m.G0 = v35 + int32(16)
	return v39
L73:
	;
	v648 = int32(0)
	v663 = v509
	goto L72
L74:
	;
	v514 = int32(*(*uint8)(unsafe.Add(mBase, _consts[601])))
	v518 = *(*float64)(unsafe.Add(mBase, _consts[588]))
	v521 = base.F64_add(base.F64_mul(base.F64_mul(v87, v518), v508), v89)
	if base.B2i32(v514 == int32(1))&base.F64_gt(v509, v521) != 0 {
		v648 = v510
		v663 = v521
		goto L72
	} else {
		goto L75
	}
L75:
	;
	if v103 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v526 = int32(0)
	v527 = l6
	goto L80
L77:
	;
	goto L78
L78:
	;
	if v514 == int32(0) {
		goto L73
	} else {
		goto L94
	}
L79:
	;
	if v591&int32(1) != 0 {
		goto L73
	} else {
		goto L93
	}
L80:
	;
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v527)+4))
	switch v559 - int32(331) {
	case 0:
		goto L85
	default:
		v591 = v526
		goto L79
	case 3:
		goto L84
	case 4:
		goto L83
	case 10, 11:
		goto L87
	case 24:
		goto L86
	case 29, 31:
		goto L82
	}
L81:
	;
	v591 = int32(1)
	goto L79
L82:
	;
	goto L81
L83:
	;
	v581 = *(*int32)(unsafe.Add(mBase, uint32(v527)+72))
	if v581 == int32(0) {
		v591 = v526
		goto L79
	} else {
		goto L91
	}
L84:
	;
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v527)+72))
	if v573 == int32(0) {
		v591 = v526
		goto L79
	} else {
		goto L89
	}
L85:
	;
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v527)))
	if v569 != int32(301) {
		v591 = v526
		goto L79
	} else {
		goto L88
	}
L86:
	;
	v564 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v527)+72)))
	v591 = int32(base.Ui32(v564&int32(2)) >> (uint(int32(1)) % 32))
	goto L79
L87:
	;
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v527)+72))
	v563 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v562)+112)))
	v591 = v563
	goto L79
L88:
	;
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v527)+72))
	v527 = v572
	goto L80
L89:
	;
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v573)+4))
	if v576 != int32(1) {
		v591 = v526
		goto L79
	} else {
		goto L90
	}
L90:
	;
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v573)+12))
	v580 = *(*int32)(unsafe.Add(mBase, uint32(v579)))
	v527 = v580
	goto L80
L91:
	;
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v581)+4))
	if v584 != int32(1) {
		v591 = v526
		goto L79
	} else {
		goto L92
	}
L92:
	;
	v587 = *(*int32)(unsafe.Add(mBase, uint32(v581)+12))
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v587)))
	v527 = v588
	goto L80
L93:
	;
	v648 = v510
	v663 = v521
	goto L72
L94:
	;
	v597 = *(*int32)(unsafe.Add(mBase, _consts[523]))
	v601 = *(*int32)(unsafe.Add(mBase, uint32(l6)+12))
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v601)+32))
	if base.F64_lt(base.F64_convert_i32_u(v597<<(uint(int32(10))%32)), base.F64_mul(v491, base.F64_convert_i32_u((v602+int32(7))&int32(-8)+int32(24)))) != 0 {
		v648 = v510
		v663 = v521
		goto L72
	} else {
		goto L95
	}
L95:
	;
	goto L73
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
	var v21 int32
	_ = v21
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
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
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
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
	v21 = base.B2i32(v17 != v18)
	if v13&int32(3) == v18 {
		v47 = v13
		v49 = v17
		v50 = v21
		goto L11
	} else {
		goto L12
	}
L5:
	;
	return v163
L6:
	;
	goto L5
L7:
	;
	if base.Ui32(v10) <= base.Ui32(v134) {
		v157 = v133
		v158 = v134
		goto L43
	} else {
		goto L44
	}
L8:
	;
	if v120 != 0 {
		goto L34
	} else {
		goto L35
	}
L9:
	;
	v120 = int32(0)
	goto L8
L10:
	;
	v98 = v91
	v100 = v93
	goto L28
L11:
	;
	if v50 == int32(0) {
		goto L9
	} else {
		goto L19
	}
L12:
	;
	if v17 == int32(0) {
		v47 = v13
		v49 = v17
		v50 = v21
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v30 = v13
	v32 = v17
	goto L14
L14:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
	if v35 == int32(10) {
		v91 = v30
		v93 = v32
		goto L10
	} else {
		goto L16
	}
L15:
	;
	v47 = v42
	v49 = v38
	v50 = v40
	goto L11
L16:
	;
	v37 = int32(1)
	v38 = v32 - v37
	v39 = int32(0)
	v40 = base.B2i32(v38 != v39)
	v42 = v30 + v37
	if v42&int32(3) == v39 {
		v47 = v42
		v49 = v38
		v50 = v40
		goto L11
	} else {
		goto L17
	}
L17:
	;
	if v38 != 0 {
		v30 = v42
		v32 = v38
		goto L14
	} else {
		goto L18
	}
L18:
	;
	goto L15
L19:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47))))
	if v54 == int32(10) {
		v84 = v47
		v86 = v49
		goto L20
	} else {
		goto L21
	}
L20:
	;
	if v86 == int32(0) {
		goto L9
	} else {
		goto L27
	}
L21:
	;
	if base.Ui32(v49) < base.Ui32(int32(4)) {
		v84 = v47
		v86 = v49
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v64 = v47
	v66 = v49
	goto L23
L23:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	v71 = v70 ^ int32(168430090)
	v74 = int32(-2139062144)
	if (int32(16843008)-v71|v71)&v74 != v74 {
		v91 = v64
		v93 = v66
		goto L10
	} else {
		goto L25
	}
L24:
	;
	v84 = v79
	v86 = v81
	goto L20
L25:
	;
	v78 = int32(4)
	v79 = v64 + v78
	v81 = v66 - v78
	if base.Ui32(int32(3)) < base.Ui32(v81) {
		v64 = v79
		v66 = v81
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	v91 = v84
	v93 = v86
	goto L10
L28:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98))))
	if int32(10) == v103 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	goto L9
L30:
	;
	v120 = v98
	goto L8
L31:
	;
	goto L32
L32:
	;
	v105 = int32(1)
	v108 = v100 - v105
	if v108 != 0 {
		v98 = v98 + v105
		v100 = v108
		goto L28
	} else {
		goto L33
	}
L33:
	;
	goto L29
L34:
	;
	v121 = v120
	goto L36
L35:
	;
	v121 = v10
	goto L36
L36:
	;
	v122 = v121 - v13
	if v122 <= int32(0) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v133 = int32(0)
	v134 = v13
	goto L7
L38:
	;
	goto L39
L39:
	;
	v126 = F_pushf_write(m, l0, v13, v122)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	return int32(0)
L41:
	;
	if v126 < int32(0) {
		v163 = v126
		goto L6
	} else {
		goto L42
	}
L42:
	;
	v133 = v126
	v134 = v13 + v122
	goto L7
L43:
	;
	if base.Ui32(v158) < base.Ui32(v10) {
		v13 = v158
		goto L4
	} else {
		goto L54
	}
L44:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134))))
	if v136 != int32(10) {
		v157 = v133
		v158 = v134
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v141 = v134
	goto L46
L46:
	;
	v146 = F_pushf_write(m, l0, int32(4019137), int32(2))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L40
	} else {
		goto L48
	}
L47:
	;
	v157 = v146
	v158 = v10
	goto L43
L48:
	;
	if v146 < int32(0) {
		v157 = v146
		v158 = v141
		goto L43
	} else {
		goto L49
	}
L49:
	;
	v151 = v141 + int32(1)
	if base.Ui32(v151) < base.Ui32(v10) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151))))
	if v153 != int32(10) {
		v157 = v146
		v158 = v151
		goto L43
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	goto L47
L53:
	;
	v141 = v151
	goto L46
L54:
	;
	v163 = v157
	goto L6
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
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
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
	var v282 int32
	_ = v282
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
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
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v339 int32
	_ = v339
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v359 int32
	_ = v359
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
	v19 = *(*int32)(unsafe.Add(mBase, _consts[3]))
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
		goto L13
	default:
		v34 = int32(41)
		goto L8
	case 10:
		goto L12
	case 29:
		goto L9
	case 36:
		goto L10
	case 45:
		goto L11
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
	v34 = int32(18)
	goto L8
L10:
	;
	v36 = int32(23)
	goto L7
L11:
	;
	v36 = int32(51)
	goto L7
L12:
	;
	v36 = int32(37)
	goto L7
L13:
	;
	v36 = int32(20)
	goto L7
L14:
	;
	goto L6
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L1
	} else {
		goto L102
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L1
	} else {
		goto L98
	}
L17:
	;
	m.G0 = v10 + int32(16)
	return v315
L18:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v288
	v290 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v13)+4)) = uint16(v290)
	v292 = F_GetLatestSnapshot(m)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L1
	} else {
		goto L92
	}
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L1
	} else {
		goto L87
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
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if v141 != 0 {
		goto L49
	} else {
		goto L50
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
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L44
	}
L25:
	;
	v65 = v46 + v47<<(uint(int32(4))%32) + int32(20) + v58*int32(100)
	v67 = v65 + int32(4)
	v68 = int32(412248)
	v71 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1036])))
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67))))
	if v72 == int32(0) {
		v91 = v71
		v92 = v72
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L24
L27:
	;
	if v92-v91 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L28:
	;
	goto L27
L29:
	;
	if v71 != v72 {
		v91 = v71
		v92 = v72
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v76 = v67
	v77 = v68
	goto L31
L31:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+1)))
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+1)))
	if v81 == int32(0) {
		v91 = v80
		v92 = v81
		goto L28
	} else {
		goto L33
	}
L32:
	;
	v91 = v80
	v92 = v81
	goto L28
L33:
	;
	v84 = int32(1)
	if v80 == v81 {
		v76 = v76 + v84
		v77 = v77 + v84
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v65)+68))
	if v96 == int32(27) {
		goto L21
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v116 = v58 + int32(1)
	if v116 != v47 {
		v58 = v116
		goto L25
	} else {
		goto L43
	}
L38:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	F_errmsg(m, int32(518864), int32(0))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(476320), int32(356), int32(29856))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L43:
	;
	goto L26
L44:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	F_errmsg(m, int32(518823), int32(0))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(476320), int32(364), int32(29856))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L1
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
	v163 = int32(0)
	goto L57
L49:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v141)))
	if v142 <= int32(0) {
		goto L15
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L53
	}
L52:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v141)+4))
	goto L48
L53:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	F_errmsg(m, int32(153731), int32(0))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(476320), int32(369), int32(29856))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L57:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v145+v163<<(uint(int32(2))%32))))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)+4))
	if v174 != int32(1) {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v173)+12))
	if v180 == int32(0) {
		goto L16
	} else {
		goto L63
	}
L59:
	;
	v178 = v163 + int32(1)
	if v142 != v178 {
		v163 = v178
		goto L57
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	goto L58
L62:
	;
	goto L15
L63:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v180)+4))
	if v183 != int32(1) {
		goto L16
	} else {
		goto L64
	}
L64:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v180)+12))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v186)))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v187)+76))
	if v188 != 0 {
		goto L67
	} else {
		goto L68
	}
L65:
	;
	if v229 == int32(0) {
		goto L15
	} else {
		goto L78
	}
L66:
	;
	goto L65
L67:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v188)+4))
	if v195 <= int32(0) {
		v229 = int32(0)
		goto L66
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v229 = int32(0)
	goto L66
L70:
	;
	v198 = int32(0)
	if v198 < v195 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v201 = v195
	goto L73
L72:
	;
	v201 = v198
	goto L73
L73:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v188)+12))
	v206 = int32(0)
	goto L74
L74:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v202+v206<<(uint(int32(2))%32))))
	v215 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v214)+8)))
	if v215 == base.I32_extend16_s(v58+int32(1))&int32(65535) {
		v229 = v214
		goto L66
	} else {
		goto L76
	}
L75:
	;
	goto L69
L76:
	;
	v218 = v206 + int32(1)
	if v218 != v201 {
		v206 = v218
		goto L74
	} else {
		goto L77
	}
L77:
	;
	goto L75
L78:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v229)+4))
	if v233 == int32(0) {
		goto L15
	} else {
		goto L79
	}
L79:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v233)))
	if v236 != int32(6) {
		goto L15
	} else {
		goto L80
	}
L80:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v233)+4))
	if v239 < int32(0) {
		goto L15
	} else {
		goto L81
	}
L81:
	;
	v242 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v233)+8)))
	if v242 != int32(65535) {
		goto L15
	} else {
		goto L82
	}
L82:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v187)+52))
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v245)+12))
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v246+v239<<(uint(int32(2))%32)-int32(4))))
	if v252 == int32(0) {
		goto L15
	} else {
		goto L83
	}
L83:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v252)+16))
	v257 = F_table_open(m, v255, int32(1))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	v259 = F_currtid_internal(m, v257, l1)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	F_sequence_close(m, v257, int32(1))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	v315 = v259
	goto L17
L87:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v271)+68))
	v273 = F_get_namespace_name(m, v272)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v273
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v275 + int32(4)
	F_errmsg(m, int32(649330), v10)
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	F_errfinish(m, int32(476320), int32(319), int32(296839))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L92:
	;
	v294 = F_RegisterSnapshot(m, v292)
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	v296 = int32(0)
	v300 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v300)+8))
	v302 = m.T0[v301].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, l0, v294, v296, v296, v296, int32(8))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	F_table_tuple_get_latest_tid(m, v302, v13)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v302)))
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v306)+188))
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v307)+12))
	m.T0[v308].(func(*base.Module, int32))(m, v302)
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	F_UnregisterSnapshot(m, v294)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	v315 = v13
	goto L17
L98:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	F_errmsg(m, int32(106330), int32(0))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	F_errfinish(m, int32(476320), int32(381), int32(29856))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L102:
	;
	F_errmsg_internal(m, int32(29908), int32(0))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	F_errfinish(m, int32(476320), int32(408), int32(29856))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
