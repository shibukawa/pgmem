package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CallerFInfoFunctionCall2(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int64) int64 {
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
	var v27 int64
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
	v9 = v7 + int32(-64)
	m.G0 = v9
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+56)) = uint8(v6)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+48)) = l4
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+40)) = uint8(v6)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+32)) = l3
	v17 = int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v9)+26)) = uint16(v17)
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+24)) = uint8(v6)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = l2
	*(*int64)(unsafe.Add(mBase, uint32(v9)+12)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = l1
	v27 = m.T0[l0].(func(*base.Module, int32) int64)(m, v7+int32(-56))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		return int64(0)
	} else {
		v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+24)))
		if v31 == int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return int64(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
				F_errmsg_internal(m, int32(_a_F_CallerFInfoFunctionCall2_0), v9)
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_CallerFInfoFunctionCall2_1), int32(1103), int32(_a_F_CallerFInfoFunctionCall2_2))
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return int64(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			m.G0 = v9 - int32(-64)
			return v27
		}
	}
}
func F_CastCreate(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int64
	_ = v28
	var v29 int64
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v95 int32
	_ = v95
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	v10 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(80)
	m.G0 = v17
	*(*uint16)(unsafe.Add(mBase, uint32(v17)+28)) = uint16(v10)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+24)) = v10
	v25 = F_table_open(m, int32(2605), int32(3))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		return
	} else {
		v28 = base.I64_extend_i32_u(l1)
		v29 = base.I64_extend_i32_u(l2)
		v30 = F_SearchSysCache2(m, int32(12), v28, v29)
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return
		} else {
			if v30 == int32(0) {
				v36 = F_GetNewOidWithIndex(m, v25, int32(2660), int32(1))
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v17)+72)) = base.I64_extend_i32_s(l7)
					*(*int64)(unsafe.Add(mBase, uint32(v17)+64)) = base.I64_extend_i32_s(l6)
					*(*int64)(unsafe.Add(mBase, uint32(v17)+56)) = base.I64_extend_i32_u(l3)
					*(*int64)(unsafe.Add(mBase, uint32(v17)+48)) = v29
					*(*int64)(unsafe.Add(mBase, uint32(v17)+40)) = v28
					*(*int64)(unsafe.Add(mBase, uint32(v17)+32)) = base.I64_extend_i32_u(v36)
					v48 = *(*int32)(unsafe.Add(mBase, uint32(v25)+52))
					v53 = F_heap_form_tuple(m, v48, v17+int32(32), v17+int32(24))
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return
					} else {
						F_CatalogTupleInsert(m, v25, v53)
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return
						} else {
							v57 = F_new_object_addresses(m)
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return
							} else {
								v59 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v59
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v36
								*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(2605)
								*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = v59
								*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = l1
								*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = int32(1247)
								v70 = v17 + int32(12)
								F_add_exact_object_address(m, v70, v57)
								mBase = m.M
								v72 = m.ExcPending
								if v72 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = l2
									*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = int32(1247)
									F_add_exact_object_address(m, v70, v57)
									mBase = m.M
									v79 = m.ExcPending
									if v79 != 0 {
										return
									} else {
										if l3 != 0 {
											*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = l3
											*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = int32(1255)
											F_add_exact_object_address(m, v70, v57)
											mBase = m.M
											v86 = m.ExcPending
											if v86 != 0 {
												return
											} else {
												if l4 != 0 {
													*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = int32(0)
													*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = l4
													*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = int32(2605)
													F_add_exact_object_address(m, v17+int32(12), v57)
													mBase = m.M
													v95 = m.ExcPending
													if v95 != 0 {
														return
													} else {
														if l5 != 0 {
															*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = int32(0)
															*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = l5
															*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = int32(2605)
															F_add_exact_object_address(m, v17+int32(12), v57)
															mBase = m.M
															v104 = m.ExcPending
															if v104 != 0 {
																return
															} else {
																F_record_object_address_dependencies(m, l0, v57, l8)
																mBase = m.M
																v106 = m.ExcPending
																if v106 != 0 {
																	return
																} else {
																	F_free_object_addresses(m, v57)
																	mBase = m.M
																	v108 = m.ExcPending
																	if v108 != 0 {
																		return
																	} else {
																		F_recordDependencyOnCurrentExtension(m, l0, int32(0))
																		mBase = m.M
																		v111 = m.ExcPending
																		if v111 != 0 {
																			return
																		} else {
																			v113 = *(*int32)(unsafe.Add(mBase, _c_F_CastCreate[0]))
																			if v113 != 0 {
																				v115 = int32(0)
																				F_RunObjectPostCreateHook(m, int32(2605), v36, v115, v115)
																				mBase = m.M
																				v118 = m.ExcPending
																				if v118 != 0 {
																					return
																				} else {
																					F_pfree(m, v53)
																					mBase = m.M
																					v120 = m.ExcPending
																					if v120 != 0 {
																						return
																					} else {
																						F_relation_close(m, v25, int32(3))
																						mBase = m.M
																						v123 = m.ExcPending
																						if v123 != 0 {
																							return
																						} else {
																							m.G0 = v17 + int32(80)
																							return
																						}
																					}
																				}
																			} else {
																				F_pfree(m, v53)
																				mBase = m.M
																				v120 = m.ExcPending
																				if v120 != 0 {
																					return
																				} else {
																					F_relation_close(m, v25, int32(3))
																					mBase = m.M
																					v123 = m.ExcPending
																					if v123 != 0 {
																						return
																					} else {
																						m.G0 = v17 + int32(80)
																						return
																					}
																				}
																			}
																		}
																	}
																}
															}
														} else {
															F_record_object_address_dependencies(m, l0, v57, l8)
															mBase = m.M
															v106 = m.ExcPending
															if v106 != 0 {
																return
															} else {
																F_free_object_addresses(m, v57)
																mBase = m.M
																v108 = m.ExcPending
																if v108 != 0 {
																	return
																} else {
																	F_recordDependencyOnCurrentExtension(m, l0, int32(0))
																	mBase = m.M
																	v111 = m.ExcPending
																	if v111 != 0 {
																		return
																	} else {
																		v113 = *(*int32)(unsafe.Add(mBase, _c_F_CastCreate[0]))
																		if v113 != 0 {
																			v115 = int32(0)
																			F_RunObjectPostCreateHook(m, int32(2605), v36, v115, v115)
																			mBase = m.M
																			v118 = m.ExcPending
																			if v118 != 0 {
																				return
																			} else {
																				F_pfree(m, v53)
																				mBase = m.M
																				v120 = m.ExcPending
																				if v120 != 0 {
																					return
																				} else {
																					F_relation_close(m, v25, int32(3))
																					mBase = m.M
																					v123 = m.ExcPending
																					if v123 != 0 {
																						return
																					} else {
																						m.G0 = v17 + int32(80)
																						return
																					}
																				}
																			}
																		} else {
																			F_pfree(m, v53)
																			mBase = m.M
																			v120 = m.ExcPending
																			if v120 != 0 {
																				return
																			} else {
																				F_relation_close(m, v25, int32(3))
																				mBase = m.M
																				v123 = m.ExcPending
																				if v123 != 0 {
																					return
																				} else {
																					m.G0 = v17 + int32(80)
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
														*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = int32(0)
														*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = l5
														*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = int32(2605)
														F_add_exact_object_address(m, v17+int32(12), v57)
														mBase = m.M
														v104 = m.ExcPending
														if v104 != 0 {
															return
														} else {
															F_record_object_address_dependencies(m, l0, v57, l8)
															mBase = m.M
															v106 = m.ExcPending
															if v106 != 0 {
																return
															} else {
																F_free_object_addresses(m, v57)
																mBase = m.M
																v108 = m.ExcPending
																if v108 != 0 {
																	return
																} else {
																	F_recordDependencyOnCurrentExtension(m, l0, int32(0))
																	mBase = m.M
																	v111 = m.ExcPending
																	if v111 != 0 {
																		return
																	} else {
																		v113 = *(*int32)(unsafe.Add(mBase, _c_F_CastCreate[0]))
																		if v113 != 0 {
																			v115 = int32(0)
																			F_RunObjectPostCreateHook(m, int32(2605), v36, v115, v115)
																			mBase = m.M
																			v118 = m.ExcPending
																			if v118 != 0 {
																				return
																			} else {
																				F_pfree(m, v53)
																				mBase = m.M
																				v120 = m.ExcPending
																				if v120 != 0 {
																					return
																				} else {
																					F_relation_close(m, v25, int32(3))
																					mBase = m.M
																					v123 = m.ExcPending
																					if v123 != 0 {
																						return
																					} else {
																						m.G0 = v17 + int32(80)
																						return
																					}
																				}
																			}
																		} else {
																			F_pfree(m, v53)
																			mBase = m.M
																			v120 = m.ExcPending
																			if v120 != 0 {
																				return
																			} else {
																				F_relation_close(m, v25, int32(3))
																				mBase = m.M
																				v123 = m.ExcPending
																				if v123 != 0 {
																					return
																				} else {
																					m.G0 = v17 + int32(80)
																					return
																				}
																			}
																		}
																	}
																}
															}
														}
													} else {
														F_record_object_address_dependencies(m, l0, v57, l8)
														mBase = m.M
														v106 = m.ExcPending
														if v106 != 0 {
															return
														} else {
															F_free_object_addresses(m, v57)
															mBase = m.M
															v108 = m.ExcPending
															if v108 != 0 {
																return
															} else {
																F_recordDependencyOnCurrentExtension(m, l0, int32(0))
																mBase = m.M
																v111 = m.ExcPending
																if v111 != 0 {
																	return
																} else {
																	v113 = *(*int32)(unsafe.Add(mBase, _c_F_CastCreate[0]))
																	if v113 != 0 {
																		v115 = int32(0)
																		F_RunObjectPostCreateHook(m, int32(2605), v36, v115, v115)
																		mBase = m.M
																		v118 = m.ExcPending
																		if v118 != 0 {
																			return
																		} else {
																			F_pfree(m, v53)
																			mBase = m.M
																			v120 = m.ExcPending
																			if v120 != 0 {
																				return
																			} else {
																				F_relation_close(m, v25, int32(3))
																				mBase = m.M
																				v123 = m.ExcPending
																				if v123 != 0 {
																					return
																				} else {
																					m.G0 = v17 + int32(80)
																					return
																				}
																			}
																		}
																	} else {
																		F_pfree(m, v53)
																		mBase = m.M
																		v120 = m.ExcPending
																		if v120 != 0 {
																			return
																		} else {
																			F_relation_close(m, v25, int32(3))
																			mBase = m.M
																			v123 = m.ExcPending
																			if v123 != 0 {
																				return
																			} else {
																				m.G0 = v17 + int32(80)
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
												*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = int32(0)
												*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = l4
												*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = int32(2605)
												F_add_exact_object_address(m, v17+int32(12), v57)
												mBase = m.M
												v95 = m.ExcPending
												if v95 != 0 {
													return
												} else {
													if l5 != 0 {
														*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = int32(0)
														*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = l5
														*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = int32(2605)
														F_add_exact_object_address(m, v17+int32(12), v57)
														mBase = m.M
														v104 = m.ExcPending
														if v104 != 0 {
															return
														} else {
															F_record_object_address_dependencies(m, l0, v57, l8)
															mBase = m.M
															v106 = m.ExcPending
															if v106 != 0 {
																return
															} else {
																F_free_object_addresses(m, v57)
																mBase = m.M
																v108 = m.ExcPending
																if v108 != 0 {
																	return
																} else {
																	F_recordDependencyOnCurrentExtension(m, l0, int32(0))
																	mBase = m.M
																	v111 = m.ExcPending
																	if v111 != 0 {
																		return
																	} else {
																		v113 = *(*int32)(unsafe.Add(mBase, _c_F_CastCreate[0]))
																		if v113 != 0 {
																			v115 = int32(0)
																			F_RunObjectPostCreateHook(m, int32(2605), v36, v115, v115)
																			mBase = m.M
																			v118 = m.ExcPending
																			if v118 != 0 {
																				return
																			} else {
																				F_pfree(m, v53)
																				mBase = m.M
																				v120 = m.ExcPending
																				if v120 != 0 {
																					return
																				} else {
																					F_relation_close(m, v25, int32(3))
																					mBase = m.M
																					v123 = m.ExcPending
																					if v123 != 0 {
																						return
																					} else {
																						m.G0 = v17 + int32(80)
																						return
																					}
																				}
																			}
																		} else {
																			F_pfree(m, v53)
																			mBase = m.M
																			v120 = m.ExcPending
																			if v120 != 0 {
																				return
																			} else {
																				F_relation_close(m, v25, int32(3))
																				mBase = m.M
																				v123 = m.ExcPending
																				if v123 != 0 {
																					return
																				} else {
																					m.G0 = v17 + int32(80)
																					return
																				}
																			}
																		}
																	}
																}
															}
														}
													} else {
														F_record_object_address_dependencies(m, l0, v57, l8)
														mBase = m.M
														v106 = m.ExcPending
														if v106 != 0 {
															return
														} else {
															F_free_object_addresses(m, v57)
															mBase = m.M
															v108 = m.ExcPending
															if v108 != 0 {
																return
															} else {
																F_recordDependencyOnCurrentExtension(m, l0, int32(0))
																mBase = m.M
																v111 = m.ExcPending
																if v111 != 0 {
																	return
																} else {
																	v113 = *(*int32)(unsafe.Add(mBase, _c_F_CastCreate[0]))
																	if v113 != 0 {
																		v115 = int32(0)
																		F_RunObjectPostCreateHook(m, int32(2605), v36, v115, v115)
																		mBase = m.M
																		v118 = m.ExcPending
																		if v118 != 0 {
																			return
																		} else {
																			F_pfree(m, v53)
																			mBase = m.M
																			v120 = m.ExcPending
																			if v120 != 0 {
																				return
																			} else {
																				F_relation_close(m, v25, int32(3))
																				mBase = m.M
																				v123 = m.ExcPending
																				if v123 != 0 {
																					return
																				} else {
																					m.G0 = v17 + int32(80)
																					return
																				}
																			}
																		}
																	} else {
																		F_pfree(m, v53)
																		mBase = m.M
																		v120 = m.ExcPending
																		if v120 != 0 {
																			return
																		} else {
																			F_relation_close(m, v25, int32(3))
																			mBase = m.M
																			v123 = m.ExcPending
																			if v123 != 0 {
																				return
																			} else {
																				m.G0 = v17 + int32(80)
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
													*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = int32(0)
													*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = l5
													*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = int32(2605)
													F_add_exact_object_address(m, v17+int32(12), v57)
													mBase = m.M
													v104 = m.ExcPending
													if v104 != 0 {
														return
													} else {
														F_record_object_address_dependencies(m, l0, v57, l8)
														mBase = m.M
														v106 = m.ExcPending
														if v106 != 0 {
															return
														} else {
															F_free_object_addresses(m, v57)
															mBase = m.M
															v108 = m.ExcPending
															if v108 != 0 {
																return
															} else {
																F_recordDependencyOnCurrentExtension(m, l0, int32(0))
																mBase = m.M
																v111 = m.ExcPending
																if v111 != 0 {
																	return
																} else {
																	v113 = *(*int32)(unsafe.Add(mBase, _c_F_CastCreate[0]))
																	if v113 != 0 {
																		v115 = int32(0)
																		F_RunObjectPostCreateHook(m, int32(2605), v36, v115, v115)
																		mBase = m.M
																		v118 = m.ExcPending
																		if v118 != 0 {
																			return
																		} else {
																			F_pfree(m, v53)
																			mBase = m.M
																			v120 = m.ExcPending
																			if v120 != 0 {
																				return
																			} else {
																				F_relation_close(m, v25, int32(3))
																				mBase = m.M
																				v123 = m.ExcPending
																				if v123 != 0 {
																					return
																				} else {
																					m.G0 = v17 + int32(80)
																					return
																				}
																			}
																		}
																	} else {
																		F_pfree(m, v53)
																		mBase = m.M
																		v120 = m.ExcPending
																		if v120 != 0 {
																			return
																		} else {
																			F_relation_close(m, v25, int32(3))
																			mBase = m.M
																			v123 = m.ExcPending
																			if v123 != 0 {
																				return
																			} else {
																				m.G0 = v17 + int32(80)
																				return
																			}
																		}
																	}
																}
															}
														}
													}
												} else {
													F_record_object_address_dependencies(m, l0, v57, l8)
													mBase = m.M
													v106 = m.ExcPending
													if v106 != 0 {
														return
													} else {
														F_free_object_addresses(m, v57)
														mBase = m.M
														v108 = m.ExcPending
														if v108 != 0 {
															return
														} else {
															F_recordDependencyOnCurrentExtension(m, l0, int32(0))
															mBase = m.M
															v111 = m.ExcPending
															if v111 != 0 {
																return
															} else {
																v113 = *(*int32)(unsafe.Add(mBase, _c_F_CastCreate[0]))
																if v113 != 0 {
																	v115 = int32(0)
																	F_RunObjectPostCreateHook(m, int32(2605), v36, v115, v115)
																	mBase = m.M
																	v118 = m.ExcPending
																	if v118 != 0 {
																		return
																	} else {
																		F_pfree(m, v53)
																		mBase = m.M
																		v120 = m.ExcPending
																		if v120 != 0 {
																			return
																		} else {
																			F_relation_close(m, v25, int32(3))
																			mBase = m.M
																			v123 = m.ExcPending
																			if v123 != 0 {
																				return
																			} else {
																				m.G0 = v17 + int32(80)
																				return
																			}
																		}
																	}
																} else {
																	F_pfree(m, v53)
																	mBase = m.M
																	v120 = m.ExcPending
																	if v120 != 0 {
																		return
																	} else {
																		F_relation_close(m, v25, int32(3))
																		mBase = m.M
																		v123 = m.ExcPending
																		if v123 != 0 {
																			return
																		} else {
																			m.G0 = v17 + int32(80)
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
				v130 = m.ExcPending
				if v130 != 0 {
					return
				} else {
					F_errcode(m, int32(_a_F_CastCreate_0))
					mBase = m.M
					v133 = m.ExcPending
					if v133 != 0 {
						return
					} else {
						v134 = F_format_type_be(m, l1)
						mBase = m.M
						v135 = m.ExcPending
						if v135 != 0 {
							return
						} else {
							v136 = F_format_type_be(m, l2)
							mBase = m.M
							v137 = m.ExcPending
							if v137 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v136
								*(*int32)(unsafe.Add(mBase, uint32(v17))) = v134
								F_errmsg(m, int32(_a_F_CastCreate_1), v17)
								mBase = m.M
								v142 = m.ExcPending
								if v142 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_CastCreate_2), int32(77), int32(_a_F_CastCreate_3))
									mBase = m.M
									v147 = m.ExcPending
									if v147 != 0 {
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
	F_aclcheck_error(m, v23, int32(43), v28)
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
func F_cannotCastJsonbValue(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	switch l0 {
	case 0:
		v30 = int32(_a_F_cannotCastJsonbValue_0)
		v31 = F_errsave_start(m, l2)
		mBase = m.M
		v32 = m.ExcPending
		if v32 != 0 {
			return
		} else {
			if v31 != 0 {
				F_errcode(m, int32(50856066))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l1
					v37 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
					F_errmsg(m, v37, v8+int32(16))
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return
					} else {
						F_errsave_finish(m, l2, int32(_a_F_cannotCastJsonbValue_1), int32(1812), int32(_a_F_cannotCastJsonbValue_2))
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return
						} else {
							m.G0 = v8 + int32(32)
							return
						}
					}
				}
			} else {
				m.G0 = v8 + int32(32)
				return
			}
		}
	case 1:
		v30 = int32(_a_F_cannotCastJsonbValue_3)
		v31 = F_errsave_start(m, l2)
		mBase = m.M
		v32 = m.ExcPending
		if v32 != 0 {
			return
		} else {
			if v31 != 0 {
				F_errcode(m, int32(50856066))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l1
					v37 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
					F_errmsg(m, v37, v8+int32(16))
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return
					} else {
						F_errsave_finish(m, l2, int32(_a_F_cannotCastJsonbValue_1), int32(1812), int32(_a_F_cannotCastJsonbValue_2))
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return
						} else {
							m.G0 = v8 + int32(32)
							return
						}
					}
				}
			} else {
				m.G0 = v8 + int32(32)
				return
			}
		}
	case 2:
		v30 = int32(_a_F_cannotCastJsonbValue_4)
		v31 = F_errsave_start(m, l2)
		mBase = m.M
		v32 = m.ExcPending
		if v32 != 0 {
			return
		} else {
			if v31 != 0 {
				F_errcode(m, int32(50856066))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l1
					v37 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
					F_errmsg(m, v37, v8+int32(16))
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return
					} else {
						F_errsave_finish(m, l2, int32(_a_F_cannotCastJsonbValue_1), int32(1812), int32(_a_F_cannotCastJsonbValue_2))
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return
						} else {
							m.G0 = v8 + int32(32)
							return
						}
					}
				}
			} else {
				m.G0 = v8 + int32(32)
				return
			}
		}
	case 3:
		v30 = int32(_a_F_cannotCastJsonbValue_5)
		v31 = F_errsave_start(m, l2)
		mBase = m.M
		v32 = m.ExcPending
		if v32 != 0 {
			return
		} else {
			if v31 != 0 {
				F_errcode(m, int32(50856066))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l1
					v37 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
					F_errmsg(m, v37, v8+int32(16))
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return
					} else {
						F_errsave_finish(m, l2, int32(_a_F_cannotCastJsonbValue_1), int32(1812), int32(_a_F_cannotCastJsonbValue_2))
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return
						} else {
							m.G0 = v8 + int32(32)
							return
						}
					}
				}
			} else {
				m.G0 = v8 + int32(32)
				return
			}
		}
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
			F_errmsg_internal(m, int32(_a_F_cannotCastJsonbValue_6), v8)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_cannotCastJsonbValue_1), int32(1815), int32(_a_F_cannotCastJsonbValue_2))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	case 16:
		v30 = int32(_a_F_cannotCastJsonbValue_7)
		v31 = F_errsave_start(m, l2)
		mBase = m.M
		v32 = m.ExcPending
		if v32 != 0 {
			return
		} else {
			if v31 != 0 {
				F_errcode(m, int32(50856066))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l1
					v37 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
					F_errmsg(m, v37, v8+int32(16))
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return
					} else {
						F_errsave_finish(m, l2, int32(_a_F_cannotCastJsonbValue_1), int32(1812), int32(_a_F_cannotCastJsonbValue_2))
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return
						} else {
							m.G0 = v8 + int32(32)
							return
						}
					}
				}
			} else {
				m.G0 = v8 + int32(32)
				return
			}
		}
	case 17:
		v30 = int32(_a_F_cannotCastJsonbValue_8)
		v31 = F_errsave_start(m, l2)
		mBase = m.M
		v32 = m.ExcPending
		if v32 != 0 {
			return
		} else {
			if v31 != 0 {
				F_errcode(m, int32(50856066))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l1
					v37 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
					F_errmsg(m, v37, v8+int32(16))
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return
					} else {
						F_errsave_finish(m, l2, int32(_a_F_cannotCastJsonbValue_1), int32(1812), int32(_a_F_cannotCastJsonbValue_2))
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return
						} else {
							m.G0 = v8 + int32(32)
							return
						}
					}
				}
			} else {
				m.G0 = v8 + int32(32)
				return
			}
		}
	case 18:
		v30 = int32(_a_F_cannotCastJsonbValue_9)
		v31 = F_errsave_start(m, l2)
		mBase = m.M
		v32 = m.ExcPending
		if v32 != 0 {
			return
		} else {
			if v31 != 0 {
				F_errcode(m, int32(50856066))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l1
					v37 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
					F_errmsg(m, v37, v8+int32(16))
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return
					} else {
						F_errsave_finish(m, l2, int32(_a_F_cannotCastJsonbValue_1), int32(1812), int32(_a_F_cannotCastJsonbValue_2))
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return
						} else {
							m.G0 = v8 + int32(32)
							return
						}
					}
				}
			} else {
				m.G0 = v8 + int32(32)
				return
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
func F_cashsmaller(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v5 int64
	_ = v5
	var v7 int64
	_ = v7
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	if v4 < v5 {
		v7 = v4
	} else {
		v7 = v5
	}
	return v7
}
