package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_IncrBufferRefCount(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_IncrBufferRefCount[0]))
	F_ResourceOwnerEnlarge(m, v4)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		if l0 < int32(0) {
			v10 = *(*int32)(unsafe.Add(mBase, _c_F_IncrBufferRefCount[1]))
			v15 = v10 + (l0^int32(-1))<<(uint(int32(2))%32)
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
			*(*int32)(unsafe.Add(mBase, uint32(v15))) = v16 + int32(1)
			v28 = *(*int32)(unsafe.Add(mBase, _c_F_IncrBufferRefCount[0]))
			F_ResourceOwnerRemember(m, v28, l0, int32(_a_F_IncrBufferRefCount_0))
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return
			} else {
				return
			}
		} else {
			v20 = F_GetPrivateRefCountEntry(m, l0)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v22 + int32(1)
				v28 = *(*int32)(unsafe.Add(mBase, _c_F_IncrBufferRefCount[0]))
				F_ResourceOwnerRemember(m, v28, l0, int32(_a_F_IncrBufferRefCount_0))
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return
				} else {
					return
				}
			}
		}
	}
}
func F_InitBuildState_2(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
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
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int64
	_ = v44
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
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
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
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
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = l1
	v13 = F_IvfflatGetTypeInfo(m, l2)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v13
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l2)+52))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v16
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l2)+180))
		if v18 == int32(0) {
			v23 = int32(100)
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
			v23 = v22
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v23
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l2)+52))
		v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
		v27 = int32(4)
		v30 = *(*int32)(unsafe.Add(mBase, uint32(v25+v26<<(uint(v27)%32))+96))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v30
		v32 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
		v36 = *(*int32)(unsafe.Add(mBase, uint32(v25+v32<<(uint(v27)%32))+88))
		if v36 != int32(1562) {
			if v30 < int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v174 = m.ExcPending
				if v174 != 0 {
					return
				} else {
					F_errcode(m, int32(50856066))
					mBase = m.M
					v177 = m.ExcPending
					if v177 != 0 {
						return
					} else {
						F_errmsg(m, int32(_a_F_InitBuildState_2_0), int32(0))
						mBase = m.M
						v181 = m.ExcPending
						if v181 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_InitBuildState_2_1), int32(358), int32(_a_F_InitBuildState_2_2))
							mBase = m.M
							v186 = m.ExcPending
							if v186 != 0 {
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
				v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
				if v42 < v30 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v190 = m.ExcPending
					if v190 != 0 {
						return
					} else {
						F_errcode(m, int32(261))
						mBase = m.M
						v193 = m.ExcPending
						if v193 != 0 {
							return
						} else {
							v194 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							v195 = *(*int32)(unsafe.Add(mBase, uint32(v194)))
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = v195
							F_errmsg(m, int32(_a_F_InitBuildState_2_3), v8)
							mBase = m.M
							v199 = m.ExcPending
							if v199 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_InitBuildState_2_1), int32(363), int32(_a_F_InitBuildState_2_2))
								mBase = m.M
								v204 = m.ExcPending
								if v204 != 0 {
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
					v44 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v44
					*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = v44
					v48 = int32(1)
					v50 = F_index_getprocinfo(m, l2, v48, v48)
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v50
						v54 = F_HnswOptionalProcInfo(m, l2, int32(2))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v54
							v58 = F_HnswOptionalProcInfo(m, l2, int32(4))
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v58
								v61 = *(*int32)(unsafe.Add(mBase, uint32(l2)+248))
								v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v62
								if v58 != 0 {
									v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
									if v64 == int32(1) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v208 = m.ExcPending
										if v208 != 0 {
											return
										} else {
											F_errcode(m, int32(50856066))
											mBase = m.M
											v211 = m.ExcPending
											if v211 != 0 {
												return
											} else {
												F_errmsg(m, int32(_a_F_InitBuildState_2_4), int32(0))
												mBase = m.M
												v215 = m.ExcPending
												if v215 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_InitBuildState_2_1), int32(378), int32(_a_F_InitBuildState_2_2))
													mBase = m.M
													v220 = m.ExcPending
													if v220 != 0 {
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
										v68 = F_CreateTemplateTupleDesc(m, int32(3))
										mBase = m.M
										v69 = m.ExcPending
										if v69 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+156)) = v68
											F_TupleDescInitEntry(m, v68, int32(1), int32(_a_F_InitBuildState_2_5), int32(23), int32(-1), int32(0))
											mBase = m.M
											v77 = m.ExcPending
											if v77 != 0 {
												return
											} else {
												v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
												F_TupleDescInitEntry(m, v78, int32(2), int32(_a_F_InitBuildState_2_6), int32(27), int32(-1), int32(0))
												mBase = m.M
												v85 = m.ExcPending
												if v85 != 0 {
													return
												} else {
													v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
													v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
													v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
													v94 = *(*int32)(unsafe.Add(mBase, uint32(v89+v90<<(uint(int32(4))%32))+88))
													F_TupleDescInitEntry(m, v86, int32(3), int32(_a_F_InitBuildState_2_7), v94, int32(-1), int32(0))
													mBase = m.M
													v98 = m.ExcPending
													if v98 != 0 {
														return
													} else {
														v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
														v101 = F_MakeSingleTupleTableSlot(m, v99, int32(_a_F_InitBuildState_2_8))
														mBase = m.M
														v102 = m.ExcPending
														if v102 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = int32(0)
															*(*int32)(unsafe.Add(mBase, uint32(l0)+160)) = v101
															v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
															v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
															v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)+8))
															v109 = m.T0[v108].(func(*base.Module, int32) int32)(m, v106)
															mBase = m.M
															v110 = m.ExcPending
															if v110 != 0 {
																return
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(l0)+76)) = v109
																v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
																v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
																v119 = F_mul_size(m, v114, (v109+int32(7))&int32(-8))
																mBase = m.M
																v120 = m.ExcPending
																if v120 != 0 {
																	return
																} else {
																	v121 = F_add_size(m, int32(20), v119)
																	mBase = m.M
																	v122 = m.ExcPending
																	if v122 != 0 {
																		return
																	} else {
																		v123 = F_add_size(m, v112, v121)
																		mBase = m.M
																		v124 = m.ExcPending
																		if v124 != 0 {
																			return
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = v123
																			F_IvfflatCheckMemoryUsage(m, v123)
																			mBase = m.M
																			v127 = m.ExcPending
																			if v127 != 0 {
																				return
																			} else {
																				v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
																				v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																				v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
																				v131 = F_VectorArrayInit(m, v128, v129, v130)
																				mBase = m.M
																				v132 = m.ExcPending
																				if v132 != 0 {
																					return
																				} else {
																					*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v131
																					v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
																					v136 = F_mul_size(m, int32(8), v135)
																					mBase = m.M
																					v137 = m.ExcPending
																					if v137 != 0 {
																						return
																					} else {
																						v138 = F_palloc(m, v136)
																						mBase = m.M
																						v139 = m.ExcPending
																						if v139 != 0 {
																							return
																						} else {
																							*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v138
																							v142 = *(*int32)(unsafe.Add(mBase, _c_F_InitBuildState_2[0]))
																							v147 = F_AllocSetContextCreateInternal(m, v142, int32(_a_F_InitBuildState_2_9), int32(0), int32(_a_F_InitBuildState_2_10), int32(_a_F_InitBuildState_2_11))
																							mBase = m.M
																							v148 = m.ExcPending
																							if v148 != 0 {
																								return
																							} else {
																								*(*int32)(unsafe.Add(mBase, uint32(l0)+172)) = int32(0)
																								*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v147
																								m.G0 = v8 + int32(16)
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
									v68 = F_CreateTemplateTupleDesc(m, int32(3))
									mBase = m.M
									v69 = m.ExcPending
									if v69 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+156)) = v68
										F_TupleDescInitEntry(m, v68, int32(1), int32(_a_F_InitBuildState_2_5), int32(23), int32(-1), int32(0))
										mBase = m.M
										v77 = m.ExcPending
										if v77 != 0 {
											return
										} else {
											v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
											F_TupleDescInitEntry(m, v78, int32(2), int32(_a_F_InitBuildState_2_6), int32(27), int32(-1), int32(0))
											mBase = m.M
											v85 = m.ExcPending
											if v85 != 0 {
												return
											} else {
												v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
												v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
												v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
												v94 = *(*int32)(unsafe.Add(mBase, uint32(v89+v90<<(uint(int32(4))%32))+88))
												F_TupleDescInitEntry(m, v86, int32(3), int32(_a_F_InitBuildState_2_7), v94, int32(-1), int32(0))
												mBase = m.M
												v98 = m.ExcPending
												if v98 != 0 {
													return
												} else {
													v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
													v101 = F_MakeSingleTupleTableSlot(m, v99, int32(_a_F_InitBuildState_2_8))
													mBase = m.M
													v102 = m.ExcPending
													if v102 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = int32(0)
														*(*int32)(unsafe.Add(mBase, uint32(l0)+160)) = v101
														v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
														v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
														v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)+8))
														v109 = m.T0[v108].(func(*base.Module, int32) int32)(m, v106)
														mBase = m.M
														v110 = m.ExcPending
														if v110 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(l0)+76)) = v109
															v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
															v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
															v119 = F_mul_size(m, v114, (v109+int32(7))&int32(-8))
															mBase = m.M
															v120 = m.ExcPending
															if v120 != 0 {
																return
															} else {
																v121 = F_add_size(m, int32(20), v119)
																mBase = m.M
																v122 = m.ExcPending
																if v122 != 0 {
																	return
																} else {
																	v123 = F_add_size(m, v112, v121)
																	mBase = m.M
																	v124 = m.ExcPending
																	if v124 != 0 {
																		return
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = v123
																		F_IvfflatCheckMemoryUsage(m, v123)
																		mBase = m.M
																		v127 = m.ExcPending
																		if v127 != 0 {
																			return
																		} else {
																			v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
																			v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																			v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
																			v131 = F_VectorArrayInit(m, v128, v129, v130)
																			mBase = m.M
																			v132 = m.ExcPending
																			if v132 != 0 {
																				return
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v131
																				v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
																				v136 = F_mul_size(m, int32(8), v135)
																				mBase = m.M
																				v137 = m.ExcPending
																				if v137 != 0 {
																					return
																				} else {
																					v138 = F_palloc(m, v136)
																					mBase = m.M
																					v139 = m.ExcPending
																					if v139 != 0 {
																						return
																					} else {
																						*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v138
																						v142 = *(*int32)(unsafe.Add(mBase, _c_F_InitBuildState_2[0]))
																						v147 = F_AllocSetContextCreateInternal(m, v142, int32(_a_F_InitBuildState_2_9), int32(0), int32(_a_F_InitBuildState_2_10), int32(_a_F_InitBuildState_2_11))
																						mBase = m.M
																						v148 = m.ExcPending
																						if v148 != 0 {
																							return
																						} else {
																							*(*int32)(unsafe.Add(mBase, uint32(l0)+172)) = int32(0)
																							*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v147
																							m.G0 = v8 + int32(16)
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
							}
						}
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v158 = m.ExcPending
			if v158 != 0 {
				return
			} else {
				F_errcode(m, int32(1088))
				mBase = m.M
				v161 = m.ExcPending
				if v161 != 0 {
					return
				} else {
					F_errmsg(m, int32(_a_F_InitBuildState_2_12), int32(0))
					mBase = m.M
					v165 = m.ExcPending
					if v165 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_InitBuildState_2_1), int32(352), int32(_a_F_InitBuildState_2_2))
						mBase = m.M
						v170 = m.ExcPending
						if v170 != 0 {
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
func F_InitHalfVector(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	v1 = l0
	v6 = F_mul_size(m, int32(2), v1)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = F_add_size(m, int32(8), v6)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			v12 = F_palloc0(m, v10)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				*(*uint16)(unsafe.Add(mBase, uint32(v12)+4)) = uint16(v1)
				*(*int32)(unsafe.Add(mBase, uint32(v12))) = v10 << (uint(int32(2)) % 32)
				return v12
			}
		}
	}
}
func F_InitLatch(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v6 int32
	_ = v6
	v1 = int32(0)
	*(*int64)(unsafe.Add(mBase, _c_F_InitLatch[0])) = int64(0)
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_InitLatch[1]))
	*(*uint8)(unsafe.Add(mBase, _c_F_InitLatch[2])) = uint8(v1)
	*(*int32)(unsafe.Add(mBase, _c_F_InitLatch[3])) = v6
	return
}
func F_InitMaterializedSRF(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
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
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v10 == int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v81 = m.ExcPending
		if v81 != 0 {
			return
		} else {
			F_errcode(m, int32(1088))
			mBase = m.M
			v84 = m.ExcPending
			if v84 != 0 {
				return
			} else {
				F_errmsg(m, int32(_a_F_InitMaterializedSRF_0), int32(0))
				mBase = m.M
				v88 = m.ExcPending
				if v88 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_InitMaterializedSRF_1), int32(89), int32(_a_F_InitMaterializedSRF_2))
					mBase = m.M
					v93 = m.ExcPending
					if v93 != 0 {
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
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
		if v13 != int32(383) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v81 = m.ExcPending
			if v81 != 0 {
				return
			} else {
				F_errcode(m, int32(1088))
				mBase = m.M
				v84 = m.ExcPending
				if v84 != 0 {
					return
				} else {
					F_errmsg(m, int32(_a_F_InitMaterializedSRF_0), int32(0))
					mBase = m.M
					v88 = m.ExcPending
					if v88 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_InitMaterializedSRF_1), int32(89), int32(_a_F_InitMaterializedSRF_2))
						mBase = m.M
						v93 = m.ExcPending
						if v93 != 0 {
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
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+12)))
			if v16&int32(2) == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v98 = m.ExcPending
				if v98 != 0 {
					return
				} else {
					F_errcode(m, int32(1088))
					mBase = m.M
					v101 = m.ExcPending
					if v101 != 0 {
						return
					} else {
						F_errmsg(m, int32(_a_F_InitMaterializedSRF_3), int32(0))
						mBase = m.M
						v105 = m.ExcPending
						if v105 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_InitMaterializedSRF_1), int32(94), int32(_a_F_InitMaterializedSRF_2))
							mBase = m.M
							v110 = m.ExcPending
							if v110 != 0 {
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
				if l1&int32(1) != 0 {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
					if v23 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v98 = m.ExcPending
						if v98 != 0 {
							return
						} else {
							F_errcode(m, int32(1088))
							mBase = m.M
							v101 = m.ExcPending
							if v101 != 0 {
								return
							} else {
								F_errmsg(m, int32(_a_F_InitMaterializedSRF_3), int32(0))
								mBase = m.M
								v105 = m.ExcPending
								if v105 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_InitMaterializedSRF_1), int32(94), int32(_a_F_InitMaterializedSRF_2))
									mBase = m.M
									v110 = m.ExcPending
									if v110 != 0 {
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
						v26 = int32(_a_F_InitMaterializedSRF_4)
						v27 = *(*int32)(unsafe.Add(mBase, _c_F_InitMaterializedSRF[0]))
						v29 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
						v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+16))
						*(*int32)(unsafe.Add(mBase, _c_F_InitMaterializedSRF[0])) = v30
						v32 = F_CreateTupleDescCopy(m, v23)
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v32
							v52 = v27
							if l1&int32(2) != 0 {
								v55 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
								v56 = F_BlessTupleDesc(m, v55)
								mBase = m.M
								v57 = m.ExcPending
								if v57 != 0 {
									return
								} else {
									v58 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
									v65 = *(*int32)(unsafe.Add(mBase, _c_F_InitMaterializedSRF[1]))
									v66 = F_tuplestore_begin_heap(m, int32(base.Ui32(v58&int32(4))>>(uint(int32(2))%32)), int32(0), v65)
									mBase = m.M
									v67 = m.ExcPending
									if v67 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v66
										*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = int32(2)
										v71 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
										*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = v71
										*(*int32)(unsafe.Add(mBase, _c_F_InitMaterializedSRF[0])) = v52
										m.G0 = v8 + int32(16)
										return
									}
								}
							} else {
								v58 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
								v65 = *(*int32)(unsafe.Add(mBase, _c_F_InitMaterializedSRF[1]))
								v66 = F_tuplestore_begin_heap(m, int32(base.Ui32(v58&int32(4))>>(uint(int32(2))%32)), int32(0), v65)
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v66
									*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = int32(2)
									v71 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
									*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = v71
									*(*int32)(unsafe.Add(mBase, _c_F_InitMaterializedSRF[0])) = v52
									m.G0 = v8 + int32(16)
									return
								}
							}
						}
					}
				} else {
					v35 = int32(_a_F_InitMaterializedSRF_4)
					v36 = *(*int32)(unsafe.Add(mBase, _c_F_InitMaterializedSRF[0]))
					v38 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+16))
					*(*int32)(unsafe.Add(mBase, _c_F_InitMaterializedSRF[0])) = v39
					v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v41)+24))
					v47 = F_internal_get_result_type(m, v42, v43, v10, int32(0), v8+int32(12))
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return
					} else {
						if v47 != int32(1) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v114 = m.ExcPending
							if v114 != 0 {
								return
							} else {
								F_errmsg_internal(m, int32(_a_F_InitMaterializedSRF_5), int32(0))
								mBase = m.M
								v118 = m.ExcPending
								if v118 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_InitMaterializedSRF_1), int32(109), int32(_a_F_InitMaterializedSRF_2))
									mBase = m.M
									v123 = m.ExcPending
									if v123 != 0 {
										return
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v52 = v36
							if l1&int32(2) != 0 {
								v55 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
								v56 = F_BlessTupleDesc(m, v55)
								mBase = m.M
								v57 = m.ExcPending
								if v57 != 0 {
									return
								} else {
									v58 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
									v65 = *(*int32)(unsafe.Add(mBase, _c_F_InitMaterializedSRF[1]))
									v66 = F_tuplestore_begin_heap(m, int32(base.Ui32(v58&int32(4))>>(uint(int32(2))%32)), int32(0), v65)
									mBase = m.M
									v67 = m.ExcPending
									if v67 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v66
										*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = int32(2)
										v71 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
										*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = v71
										*(*int32)(unsafe.Add(mBase, _c_F_InitMaterializedSRF[0])) = v52
										m.G0 = v8 + int32(16)
										return
									}
								}
							} else {
								v58 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
								v65 = *(*int32)(unsafe.Add(mBase, _c_F_InitMaterializedSRF[1]))
								v66 = F_tuplestore_begin_heap(m, int32(base.Ui32(v58&int32(4))>>(uint(int32(2))%32)), int32(0), v65)
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v66
									*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = int32(2)
									v71 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
									*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = v71
									*(*int32)(unsafe.Add(mBase, _c_F_InitMaterializedSRF[0])) = v52
									m.G0 = v8 + int32(16)
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
func F_InitStandaloneProcess(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v15 int64
	_ = v15
	var v18 int64
	_ = v18
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int64
	_ = v27
	var v29 int64
	_ = v29
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v66 int64
	_ = v66
	var v67 int32
	_ = v67
	var v76 int64
	_ = v76
	var v78 int64
	_ = v78
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v156 int32
	_ = v156
	var v162 int32
	_ = v162
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	v3 = m.G0
	v4 = int32(16)
	v5 = v3 - v4
	m.G0 = v5
	*(*int32)(unsafe.Add(mBase, _c_F_InitStandaloneProcess[0])) = int32(8)
	v15 = F_GetCurrentTimestamp(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, _c_F_InitStandaloneProcess[1])) = v15
	v18 = F_timestamptz_to_time_t(m, v15)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, _c_F_InitStandaloneProcess[2])) = v18
	v22 = F_pg_strong_random(m, int32(_a_F_InitStandaloneProcess_0), v4)
	mBase = m.M
	if v22 != 0 {
		v24 = F_pg_prng_seed_check(m, int32(_a_F_InitStandaloneProcess_0))
		mBase = m.M
		if v24 != 0 {
		} else {
			v27 = int64(*(*int32)(unsafe.Add(mBase, _c_F_InitStandaloneProcess[3])))
			v29 = *(*int64)(unsafe.Add(mBase, _c_F_InitStandaloneProcess[1]))
			F_pg_prng_seed(m, int32(_a_F_InitStandaloneProcess_0), v27^v29<<(uint(int64(12))%64)^int64(base.Ui64(v29)>>(uint(int64(20))%64)))
			mBase = m.M
		}
	} else {
		v27 = int64(*(*int32)(unsafe.Add(mBase, _c_F_InitStandaloneProcess[3])))
		v29 = *(*int64)(unsafe.Add(mBase, _c_F_InitStandaloneProcess[1]))
		F_pg_prng_seed(m, int32(_a_F_InitStandaloneProcess_0), v27^v29<<(uint(int64(12))%64)^int64(base.Ui64(v29)>>(uint(int64(20))%64)))
		mBase = m.M
	}
	v38 = F_pg_prng_uint32(m)
	mBase = m.M
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_InitStandaloneProcess[4]))
	if v40 == int32(0) {
		v44 = *(*int32)(unsafe.Add(mBase, _c_F_InitStandaloneProcess[5]))
		*(*int32)(unsafe.Add(mBase, uint32(v44))) = v38
	} else {
		v47 = int32(3)
		if v40 == int32(7) {
			v52 = v47
		} else {
			v52 = int32(1)
		}
		if v40 == int32(31) {
			v55 = v47
		} else {
			v55 = v52
		}
		*(*int32)(unsafe.Add(mBase, _c_F_InitStandaloneProcess[6])) = v55
		v58 = int32(0)
		*(*int32)(unsafe.Add(mBase, _c_F_InitStandaloneProcess[7])) = v58
		v61 = *(*int32)(unsafe.Add(mBase, _c_F_InitStandaloneProcess[5]))
		if v58 < v40 {
			v66 = base.I64_extend_i32_u(v38)
			v67 = int32(0)
			for {
				v76 = v66*int64(6364136223846793005) + int64(1)
				v78 = int64(base.Ui64(v76) >> (uint(int64(32)) % 64))
				*(*uint32)(unsafe.Add(mBase, uint32(v61+v67<<(uint(int32(2))%32)))) = uint32(v78)
				v81 = v67 + int32(1)
				if v81 != v40 {
					v66 = v76
					v67 = v81
					continue
				} else {
					break
				}
				break
			}
		} else {
		}
		v87 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
		*(*int32)(unsafe.Add(mBase, uint32(v61))) = v87 | int32(1)
	}
	F_InitializeWaitEventSupport(m)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		return
	} else {
		v98 = int32(_a_F_InitStandaloneProcess_1)
		*(*int32)(unsafe.Add(mBase, _c_F_InitStandaloneProcess[8])) = v98
		v100 = int32(0)
		*(*int64)(unsafe.Add(mBase, _c_F_InitStandaloneProcess[9])) = int64(0)
		v105 = *(*int32)(unsafe.Add(mBase, _c_F_InitStandaloneProcess[3]))
		*(*uint8)(unsafe.Add(mBase, _c_F_InitStandaloneProcess[10])) = uint8(v100)
		*(*int32)(unsafe.Add(mBase, _c_F_InitStandaloneProcess[11])) = v105
		F_InitializeLatchWaitSet(m)
		mBase = m.M
		v112 = m.ExcPending
		if v112 != 0 {
			return
		} else {
			F_sigemptyset(m, int32(_a_F_InitStandaloneProcess_2))
			mBase = m.M
			v115 = int32(_a_F_InitStandaloneProcess_3)
			F_sigfillset(m, v115)
			mBase = m.M
			v117 = int32(_a_F_InitStandaloneProcess_4)
			F_sigfillset(m, v117)
			mBase = m.M
			v120 = int32(5)
			F_sigdelset(m, v115, v120)
			mBase = m.M
			F_sigdelset(m, v117, v120)
			mBase = m.M
			v126 = int32(6)
			F_sigdelset(m, v115, v126)
			mBase = m.M
			F_sigdelset(m, v117, v126)
			mBase = m.M
			v132 = int32(4)
			F_sigdelset(m, v115, v132)
			mBase = m.M
			F_sigdelset(m, v117, v132)
			mBase = m.M
			v138 = int32(8)
			F_sigdelset(m, v115, v138)
			mBase = m.M
			F_sigdelset(m, v117, v138)
			mBase = m.M
			v144 = int32(11)
			F_sigdelset(m, v115, v144)
			mBase = m.M
			F_sigdelset(m, v117, v144)
			mBase = m.M
			v150 = int32(7)
			F_sigdelset(m, v115, v150)
			mBase = m.M
			F_sigdelset(m, v117, v150)
			mBase = m.M
			v156 = int32(31)
			F_sigdelset(m, v115, v156)
			mBase = m.M
			F_sigdelset(m, v117, v156)
			mBase = m.M
			v162 = int32(18)
			F_sigdelset(m, v115, v162)
			mBase = m.M
			F_sigdelset(m, v117, v162)
			mBase = m.M
			F_sigdelset(m, v117, int32(3))
			mBase = m.M
			F_sigdelset(m, v117, int32(15))
			mBase = m.M
			F_sigdelset(m, v117, int32(14))
			mBase = m.M
			F_sigprocmask(m, int32(_a_F_InitStandaloneProcess_3), int32(0))
			mBase = m.M
			v179 = m.ExcPending
			if v179 != 0 {
				return
			} else {
				v181 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitStandaloneProcess[12])))
				if v181 == int32(0) {
					v185 = F_find_my_exec(m, l0, int32(_a_F_InitStandaloneProcess_5))
					mBase = m.M
					v186 = m.ExcPending
					if v186 != 0 {
						return
					} else {
						if v185 < int32(0) {
							F_errstart_cold(m, int32(22), int32(0))
							mBase = m.M
							v202 = m.ExcPending
							if v202 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v5))) = l0
								F_errmsg_internal(m, int32(_a_F_InitStandaloneProcess_6), v5)
								mBase = m.M
								v206 = m.ExcPending
								if v206 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_InitStandaloneProcess_7), int32(207), int32(_a_F_InitStandaloneProcess_8))
									mBase = m.M
									v211 = m.ExcPending
									if v211 != 0 {
										return
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v190 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitStandaloneProcess[13])))
							if v190 == int32(0) {
								F_get_pkglib_path(m, int32(_a_F_InitStandaloneProcess_9))
								mBase = m.M
								v195 = m.ExcPending
								if v195 != 0 {
									return
								} else {
									m.G0 = v5 + int32(16)
									return
								}
							} else {
								m.G0 = v5 + int32(16)
								return
							}
						}
					}
				} else {
					v190 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitStandaloneProcess[13])))
					if v190 == int32(0) {
						F_get_pkglib_path(m, int32(_a_F_InitStandaloneProcess_9))
						mBase = m.M
						v195 = m.ExcPending
						if v195 != 0 {
							return
						} else {
							m.G0 = v5 + int32(16)
							return
						}
					} else {
						m.G0 = v5 + int32(16)
						return
					}
				}
			}
		}
	}
}
func F_InitializeOneGUCOption(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int64
	_ = v10
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
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
	var v77 float64
	_ = v77
	var v79 int32
	_ = v79
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 float64
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 float64
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
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
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 float64
	_ = v205
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v249 int32
	_ = v249
	var v254 int32
	_ = v254
	v2 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(96)
	m.G0 = v8
	v10 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+28)) = v10
	*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v2
	*(*int64)(unsafe.Add(mBase, uint32(l0)+80)) = v10
	*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v2
	*(*int64)(unsafe.Add(mBase, uint32(l0)+52)) = int64(10)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+36)) = v10
	*(*int64)(unsafe.Add(mBase, uint32(l0)+44)) = int64(42949672960)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	switch v24 {
	case 0:
		v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+96)))
		*(*uint8)(unsafe.Add(mBase, uint32(v8)+92)) = uint8(v25)
		v27 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v8)+80)) = v27
		v35 = F_call_bool_check_hook(m, l0, v8+int32(92), v8+int32(80), v27, int32(15))
		mBase = m.M
		v36 = m.ExcPending
		if v36 != 0 {
			return
		} else {
			if v35 == int32(0) {
				F_errstart_cold(m, int32(22), int32(0))
				mBase = m.M
				v168 = m.ExcPending
				if v168 != 0 {
					return
				} else {
					v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = v169
					v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+92)))
					*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v171
					F_errmsg_internal(m, int32(_a_F_InitializeOneGUCOption_0), v8)
					mBase = m.M
					v175 = m.ExcPending
					if v175 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_InitializeOneGUCOption_1), int32(1670), int32(_a_F_InitializeOneGUCOption_2))
						mBase = m.M
						v180 = m.ExcPending
						if v180 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
				if v39 != 0 {
					v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+92)))
					v41 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
					m.T0[v39].(func(*base.Module, int32, int32))(m, v40, v41)
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return
					} else {
						v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+92)))
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+112)) = uint8(v44)
						v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
						*(*uint8)(unsafe.Add(mBase, uint32(v46))) = uint8(v44)
						v48 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v48
						*(*int32)(unsafe.Add(mBase, uint32(l0)+116)) = v48
						m.G0 = v8 + int32(96)
						return
					}
				} else {
					v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+92)))
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+112)) = uint8(v44)
					v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
					*(*uint8)(unsafe.Add(mBase, uint32(v46))) = uint8(v44)
					v48 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v48
					*(*int32)(unsafe.Add(mBase, uint32(l0)+116)) = v48
					m.G0 = v8 + int32(96)
					return
				}
			}
		}
	case 1:
		v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
		*(*int32)(unsafe.Add(mBase, uint32(v8)+80)) = v51
		v53 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v8)+92)) = v53
		v61 = F_call_int_check_hook(m, l0, v8+int32(80), v8+int32(92), v53, int32(15))
		mBase = m.M
		v62 = m.ExcPending
		if v62 != 0 {
			return
		} else {
			if v61 == int32(0) {
				F_errstart_cold(m, int32(22), int32(0))
				mBase = m.M
				v184 = m.ExcPending
				if v184 != 0 {
					return
				} else {
					v185 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v185
					v187 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
					*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v187
					F_errmsg_internal(m, int32(_a_F_InitializeOneGUCOption_0), v8+int32(16))
					mBase = m.M
					v193 = m.ExcPending
					if v193 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_InitializeOneGUCOption_1), int32(1688), int32(_a_F_InitializeOneGUCOption_2))
						mBase = m.M
						v198 = m.ExcPending
						if v198 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
				if v65 != 0 {
					v66 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
					v67 = *(*int32)(unsafe.Add(mBase, uint32(v8)+92))
					m.T0[v65].(func(*base.Module, int32, int32))(m, v66, v67)
					mBase = m.M
					v69 = m.ExcPending
					if v69 != 0 {
						return
					} else {
						v70 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v70
						v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
						*(*int32)(unsafe.Add(mBase, uint32(v72))) = v70
						v74 = *(*int32)(unsafe.Add(mBase, uint32(v8)+92))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v74
						*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = v74
						m.G0 = v8 + int32(96)
						return
					}
				} else {
					v70 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v70
					v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
					*(*int32)(unsafe.Add(mBase, uint32(v72))) = v70
					v74 = *(*int32)(unsafe.Add(mBase, uint32(v8)+92))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v74
					*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = v74
					m.G0 = v8 + int32(96)
					return
				}
			}
		}
	case 2:
		v77 = *(*float64)(unsafe.Add(mBase, uint32(l0)+96))
		*(*float64)(unsafe.Add(mBase, uint32(v8)+80)) = v77
		v79 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v8)+92)) = v79
		v87 = F_call_real_check_hook(m, l0, v8+int32(80), v8+int32(92), v79, int32(15))
		mBase = m.M
		v88 = m.ExcPending
		if v88 != 0 {
			return
		} else {
			if v87 == int32(0) {
				F_errstart_cold(m, int32(22), int32(0))
				mBase = m.M
				v202 = m.ExcPending
				if v202 != 0 {
					return
				} else {
					v203 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v203
					v205 = *(*float64)(unsafe.Add(mBase, uint32(v8)+80))
					*(*float64)(unsafe.Add(mBase, uint32(v8)+40)) = v205
					F_errmsg_internal(m, int32(_a_F_InitializeOneGUCOption_3), v8+int32(32))
					mBase = m.M
					v211 = m.ExcPending
					if v211 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_InitializeOneGUCOption_1), int32(1706), int32(_a_F_InitializeOneGUCOption_2))
						mBase = m.M
						v216 = m.ExcPending
						if v216 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
				if v91 != 0 {
					v92 = *(*float64)(unsafe.Add(mBase, uint32(v8)+80))
					v93 = *(*int32)(unsafe.Add(mBase, uint32(v8)+92))
					m.T0[v91].(func(*base.Module, float64, int32))(m, v92, v93)
					mBase = m.M
					v95 = m.ExcPending
					if v95 != 0 {
						return
					} else {
						v96 = *(*float64)(unsafe.Add(mBase, uint32(v8)+80))
						*(*float64)(unsafe.Add(mBase, uint32(l0)+136)) = v96
						v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
						*(*float64)(unsafe.Add(mBase, uint32(v98))) = v96
						v100 = *(*int32)(unsafe.Add(mBase, uint32(v8)+92))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v100
						*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = v100
						m.G0 = v8 + int32(96)
						return
					}
				} else {
					v96 = *(*float64)(unsafe.Add(mBase, uint32(v8)+80))
					*(*float64)(unsafe.Add(mBase, uint32(l0)+136)) = v96
					v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
					*(*float64)(unsafe.Add(mBase, uint32(v98))) = v96
					v100 = *(*int32)(unsafe.Add(mBase, uint32(v8)+92))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v100
					*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = v100
					m.G0 = v8 + int32(96)
					return
				}
			}
		}
	case 3:
		*(*int32)(unsafe.Add(mBase, uint32(v8)+92)) = int32(0)
		v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
		if v105 != 0 {
			v107 = F_guc_strdup(m, int32(22), v105)
			mBase = m.M
			v108 = m.ExcPending
			if v108 != 0 {
				return
			} else {
				v109 = v107
				*(*int32)(unsafe.Add(mBase, uint32(v8)+80)) = v109
				v117 = F_call_string_check_hook(m, l0, v8+int32(80), v8+int32(92), int32(0), int32(15))
				mBase = m.M
				v118 = m.ExcPending
				if v118 != 0 {
					return
				} else {
					if v117 == int32(0) {
						F_errstart_cold(m, int32(22), int32(0))
						mBase = m.M
						v220 = m.ExcPending
						if v220 != 0 {
							return
						} else {
							v221 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = v221
							v223 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
							if v223 != 0 {
								v225 = v223
							} else {
								v225 = int32(_a_F_InitializeOneGUCOption_4)
							}
							*(*int32)(unsafe.Add(mBase, uint32(v8)+52)) = v225
							F_errmsg_internal(m, int32(_a_F_InitializeOneGUCOption_5), v8+int32(48))
							mBase = m.M
							v231 = m.ExcPending
							if v231 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_InitializeOneGUCOption_1), int32(1728), int32(_a_F_InitializeOneGUCOption_2))
								mBase = m.M
								v236 = m.ExcPending
								if v236 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
						if v121 != 0 {
							v122 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
							v123 = *(*int32)(unsafe.Add(mBase, uint32(v8)+92))
							m.T0[v121].(func(*base.Module, int32, int32))(m, v122, v123)
							mBase = m.M
							v125 = m.ExcPending
							if v125 != 0 {
								return
							} else {
								v126 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v126
								v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
								*(*int32)(unsafe.Add(mBase, uint32(v128))) = v126
								v130 = *(*int32)(unsafe.Add(mBase, uint32(v8)+92))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v130
								*(*int32)(unsafe.Add(mBase, uint32(l0)+116)) = v130
								m.G0 = v8 + int32(96)
								return
							}
						} else {
							v126 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v126
							v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
							*(*int32)(unsafe.Add(mBase, uint32(v128))) = v126
							v130 = *(*int32)(unsafe.Add(mBase, uint32(v8)+92))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v130
							*(*int32)(unsafe.Add(mBase, uint32(l0)+116)) = v130
							m.G0 = v8 + int32(96)
							return
						}
					}
				}
			}
		} else {
			v109 = v2
			*(*int32)(unsafe.Add(mBase, uint32(v8)+80)) = v109
			v117 = F_call_string_check_hook(m, l0, v8+int32(80), v8+int32(92), int32(0), int32(15))
			mBase = m.M
			v118 = m.ExcPending
			if v118 != 0 {
				return
			} else {
				if v117 == int32(0) {
					F_errstart_cold(m, int32(22), int32(0))
					mBase = m.M
					v220 = m.ExcPending
					if v220 != 0 {
						return
					} else {
						v221 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = v221
						v223 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
						if v223 != 0 {
							v225 = v223
						} else {
							v225 = int32(_a_F_InitializeOneGUCOption_4)
						}
						*(*int32)(unsafe.Add(mBase, uint32(v8)+52)) = v225
						F_errmsg_internal(m, int32(_a_F_InitializeOneGUCOption_5), v8+int32(48))
						mBase = m.M
						v231 = m.ExcPending
						if v231 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_InitializeOneGUCOption_1), int32(1728), int32(_a_F_InitializeOneGUCOption_2))
							mBase = m.M
							v236 = m.ExcPending
							if v236 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
					if v121 != 0 {
						v122 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
						v123 = *(*int32)(unsafe.Add(mBase, uint32(v8)+92))
						m.T0[v121].(func(*base.Module, int32, int32))(m, v122, v123)
						mBase = m.M
						v125 = m.ExcPending
						if v125 != 0 {
							return
						} else {
							v126 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v126
							v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
							*(*int32)(unsafe.Add(mBase, uint32(v128))) = v126
							v130 = *(*int32)(unsafe.Add(mBase, uint32(v8)+92))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v130
							*(*int32)(unsafe.Add(mBase, uint32(l0)+116)) = v130
							m.G0 = v8 + int32(96)
							return
						}
					} else {
						v126 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v126
						v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
						*(*int32)(unsafe.Add(mBase, uint32(v128))) = v126
						v130 = *(*int32)(unsafe.Add(mBase, uint32(v8)+92))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v130
						*(*int32)(unsafe.Add(mBase, uint32(l0)+116)) = v130
						m.G0 = v8 + int32(96)
						return
					}
				}
			}
		}
	case 4:
		v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
		*(*int32)(unsafe.Add(mBase, uint32(v8)+80)) = v133
		v135 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v8)+92)) = v135
		v143 = F_call_enum_check_hook(m, l0, v8+int32(80), v8+int32(92), v135, int32(15))
		mBase = m.M
		v144 = m.ExcPending
		if v144 != 0 {
			return
		} else {
			if v143 == int32(0) {
				F_errstart_cold(m, int32(22), int32(0))
				mBase = m.M
				v240 = m.ExcPending
				if v240 != 0 {
					return
				} else {
					v241 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					*(*int32)(unsafe.Add(mBase, uint32(v8)+64)) = v241
					v243 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
					*(*int32)(unsafe.Add(mBase, uint32(v8)+68)) = v243
					F_errmsg_internal(m, int32(_a_F_InitializeOneGUCOption_0), v8-int32(-64))
					mBase = m.M
					v249 = m.ExcPending
					if v249 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_InitializeOneGUCOption_1), int32(1744), int32(_a_F_InitializeOneGUCOption_2))
						mBase = m.M
						v254 = m.ExcPending
						if v254 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
				if v147 != 0 {
					v148 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
					v149 = *(*int32)(unsafe.Add(mBase, uint32(v8)+92))
					m.T0[v147].(func(*base.Module, int32, int32))(m, v148, v149)
					mBase = m.M
					v151 = m.ExcPending
					if v151 != 0 {
						return
					} else {
						v152 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+116)) = v152
						v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
						*(*int32)(unsafe.Add(mBase, uint32(v154))) = v152
						v156 = *(*int32)(unsafe.Add(mBase, uint32(v8)+92))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v156
						*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v156
						m.G0 = v8 + int32(96)
						return
					}
				} else {
					v152 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+116)) = v152
					v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
					*(*int32)(unsafe.Add(mBase, uint32(v154))) = v152
					v156 = *(*int32)(unsafe.Add(mBase, uint32(v8)+92))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v156
					*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v156
					m.G0 = v8 + int32(96)
					return
				}
			}
		}
	default:
		m.G0 = v8 + int32(96)
		return
	}
}
func F_InitializeSessionUserId(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
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
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v198 int32
	_ = v198
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v231 int32
	_ = v231
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v295 int32
	_ = v295
	var v300 int32
	_ = v300
	v15 = m.G0
	v17 = v15 + int32(-64)
	m.G0 = v17
	v20 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitializeSessionUserId[0])))
	if v20 == int32(0) {
		F_ReceiveSharedInvalidMessages(m)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return
		} else {
			if l0 != 0 {
				v26 = F_SearchSysCache1(m, int32(10), l0)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return
				} else {
					if v26 != 0 {
						v51 = v26
						v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+16))
						v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+22)))
						v54 = v52 + v53
						v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+68)))
						v57 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
						*(*int32)(unsafe.Add(mBase, _c_F_InitializeSessionUserId[1])) = v57
						v60 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeSessionUserId[2]))
						*(*int32)(unsafe.Add(mBase, uint32(v60)+64)) = v57
						v63 = int32(4)
						v64 = v54 + v63
						F_SetConfigOption(m, int32(_a_F_InitializeSessionUserId_0), v64, v63, int32(10))
						mBase = m.M
						v68 = m.ExcPending
						if v68 != 0 {
							return
						} else {
							v70 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitializeSessionUserId[3])))
							if v70 != int32(1) {
								F_ReleaseCatCache(m, v51)
								mBase = m.M
								v231 = m.ExcPending
								if v231 != 0 {
									return
								} else {
									m.G0 = v17 - int32(-64)
									return
								}
							} else {
								if l2 == int32(0) {
									v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+72)))
									if v75 == int32(0) {
										F_errstart_cold(m, int32(22), int32(0))
										mBase = m.M
										v268 = m.ExcPending
										if v268 != 0 {
											return
										} else {
											F_errcode(m, int32(514))
											mBase = m.M
											v271 = m.ExcPending
											if v271 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v64
												F_errmsg(m, int32(_a_F_InitializeSessionUserId_1), v15+int32(-32))
												mBase = m.M
												v277 = m.ExcPending
												if v277 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_InitializeSessionUserId_2), int32(859), int32(_a_F_InitializeSessionUserId_3))
													mBase = m.M
													v282 = m.ExcPending
													if v282 != 0 {
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
										v78 = *(*int32)(unsafe.Add(mBase, uint32(v54)+76))
										if v78 < int32(0) {
											F_ReleaseCatCache(m, v51)
											mBase = m.M
											v231 = m.ExcPending
											if v231 != 0 {
												return
											} else {
												m.G0 = v17 - int32(-64)
												return
											}
										} else {
											v82 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeSessionUserId[4]))
											if v82 != int32(1) {
												F_ReleaseCatCache(m, v51)
												mBase = m.M
												v231 = m.ExcPending
												if v231 != 0 {
													return
												} else {
													m.G0 = v17 - int32(-64)
													return
												}
											} else {
												if v55&int32(1) != 0 {
													F_ReleaseCatCache(m, v51)
													mBase = m.M
													v231 = m.ExcPending
													if v231 != 0 {
														return
													} else {
														m.G0 = v17 - int32(-64)
														return
													}
												} else {
													v87 = int32(0)
													v89 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeSessionUserId[5]))
													v91 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeSessionUserId[6]))
													v95 = F_LWLockAcquire(m, v91+int32(512), int32(1))
													mBase = m.M
													v96 = m.ExcPending
													if v96 != 0 {
														return
													} else {
														v97 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
														if v97 <= int32(0) {
															v198 = v87
														} else {
															v100 = int32(1)
															v103 = v89 + int32(36)
															v105 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeSessionUserId[7]))
															v106 = int32(0)
															if v97 != v100 {
																v114 = v106
																v116 = v87
																v117 = int32(0)
																for {
																	v128 = v103 + v114<<(uint(int32(2))%32)
																	v129 = *(*int32)(unsafe.Add(mBase, uint32(v128)))
																	v132 = v105 + v129*int32(640)
																	v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)+44))
																	if v133 == int32(0) {
																		v142 = v116
																	} else {
																		v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132)+72)))
																		if v136 != int32(1) {
																			v142 = v116
																		} else {
																			v139 = *(*int32)(unsafe.Add(mBase, uint32(v132)+64))
																			v142 = v116 + base.B2i32(v139 == v57)
																		}
																	}
																	v143 = *(*int32)(unsafe.Add(mBase, uint32(v128)+4))
																	v146 = v105 + v143*int32(640)
																	v147 = *(*int32)(unsafe.Add(mBase, uint32(v146)+44))
																	if v147 == int32(0) {
																		v156 = v142
																	} else {
																		v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+72)))
																		if v150 != int32(1) {
																			v156 = v142
																		} else {
																			v153 = *(*int32)(unsafe.Add(mBase, uint32(v146)+64))
																			v156 = v142 + base.B2i32(v153 == v57)
																		}
																	}
																	v157 = int32(2)
																	v158 = v114 + v157
																	v160 = v117 + v157
																	if v160 != v97&int32(2147483646) {
																		v114 = v158
																		v116 = v156
																		v117 = v160
																		continue
																	} else {
																		break
																	}
																	break
																}
																v164 = v158
																v166 = v156
															} else {
																v164 = v106
																v166 = v87
															}
															if v97&v100 == int32(0) {
																v198 = v166
															} else {
																v181 = *(*int32)(unsafe.Add(mBase, uint32(v103+v164<<(uint(int32(2))%32))))
																v184 = v105 + v181*int32(640)
																v185 = *(*int32)(unsafe.Add(mBase, uint32(v184)+44))
																if v185 == int32(0) {
																	v198 = v166
																} else {
																	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184)+72)))
																	if v188 != int32(1) {
																		v198 = v166
																	} else {
																		v191 = *(*int32)(unsafe.Add(mBase, uint32(v184)+64))
																		v198 = v166 + base.B2i32(v191 == v57)
																	}
																}
															}
														}
														v209 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeSessionUserId[6]))
														F_LWLockRelease(m, v209+int32(512))
														mBase = m.M
														v213 = m.ExcPending
														if v213 != 0 {
															return
														} else {
															v214 = *(*int32)(unsafe.Add(mBase, uint32(v54)+76))
															if v214 < v198 {
																F_errstart_cold(m, int32(22), int32(0))
																mBase = m.M
																v286 = m.ExcPending
																if v286 != 0 {
																	return
																} else {
																	F_errcode(m, int32(_a_F_InitializeSessionUserId_4))
																	mBase = m.M
																	v289 = m.ExcPending
																	if v289 != 0 {
																		return
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v64
																		F_errmsg(m, int32(_a_F_InitializeSessionUserId_5), v15+int32(-48))
																		mBase = m.M
																		v295 = m.ExcPending
																		if v295 != 0 {
																			return
																		} else {
																			F_errfinish(m, int32(_a_F_InitializeSessionUserId_2), int32(880), int32(_a_F_InitializeSessionUserId_3))
																			mBase = m.M
																			v300 = m.ExcPending
																			if v300 != 0 {
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
																F_ReleaseCatCache(m, v51)
																mBase = m.M
																v231 = m.ExcPending
																if v231 != 0 {
																	return
																} else {
																	m.G0 = v17 - int32(-64)
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
									v78 = *(*int32)(unsafe.Add(mBase, uint32(v54)+76))
									if v78 < int32(0) {
										F_ReleaseCatCache(m, v51)
										mBase = m.M
										v231 = m.ExcPending
										if v231 != 0 {
											return
										} else {
											m.G0 = v17 - int32(-64)
											return
										}
									} else {
										v82 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeSessionUserId[4]))
										if v82 != int32(1) {
											F_ReleaseCatCache(m, v51)
											mBase = m.M
											v231 = m.ExcPending
											if v231 != 0 {
												return
											} else {
												m.G0 = v17 - int32(-64)
												return
											}
										} else {
											if v55&int32(1) != 0 {
												F_ReleaseCatCache(m, v51)
												mBase = m.M
												v231 = m.ExcPending
												if v231 != 0 {
													return
												} else {
													m.G0 = v17 - int32(-64)
													return
												}
											} else {
												v87 = int32(0)
												v89 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeSessionUserId[5]))
												v91 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeSessionUserId[6]))
												v95 = F_LWLockAcquire(m, v91+int32(512), int32(1))
												mBase = m.M
												v96 = m.ExcPending
												if v96 != 0 {
													return
												} else {
													v97 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
													if v97 <= int32(0) {
														v198 = v87
													} else {
														v100 = int32(1)
														v103 = v89 + int32(36)
														v105 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeSessionUserId[7]))
														v106 = int32(0)
														if v97 != v100 {
															v114 = v106
															v116 = v87
															v117 = int32(0)
															for {
																v128 = v103 + v114<<(uint(int32(2))%32)
																v129 = *(*int32)(unsafe.Add(mBase, uint32(v128)))
																v132 = v105 + v129*int32(640)
																v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)+44))
																if v133 == int32(0) {
																	v142 = v116
																} else {
																	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132)+72)))
																	if v136 != int32(1) {
																		v142 = v116
																	} else {
																		v139 = *(*int32)(unsafe.Add(mBase, uint32(v132)+64))
																		v142 = v116 + base.B2i32(v139 == v57)
																	}
																}
																v143 = *(*int32)(unsafe.Add(mBase, uint32(v128)+4))
																v146 = v105 + v143*int32(640)
																v147 = *(*int32)(unsafe.Add(mBase, uint32(v146)+44))
																if v147 == int32(0) {
																	v156 = v142
																} else {
																	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+72)))
																	if v150 != int32(1) {
																		v156 = v142
																	} else {
																		v153 = *(*int32)(unsafe.Add(mBase, uint32(v146)+64))
																		v156 = v142 + base.B2i32(v153 == v57)
																	}
																}
																v157 = int32(2)
																v158 = v114 + v157
																v160 = v117 + v157
																if v160 != v97&int32(2147483646) {
																	v114 = v158
																	v116 = v156
																	v117 = v160
																	continue
																} else {
																	break
																}
																break
															}
															v164 = v158
															v166 = v156
														} else {
															v164 = v106
															v166 = v87
														}
														if v97&v100 == int32(0) {
															v198 = v166
														} else {
															v181 = *(*int32)(unsafe.Add(mBase, uint32(v103+v164<<(uint(int32(2))%32))))
															v184 = v105 + v181*int32(640)
															v185 = *(*int32)(unsafe.Add(mBase, uint32(v184)+44))
															if v185 == int32(0) {
																v198 = v166
															} else {
																v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184)+72)))
																if v188 != int32(1) {
																	v198 = v166
																} else {
																	v191 = *(*int32)(unsafe.Add(mBase, uint32(v184)+64))
																	v198 = v166 + base.B2i32(v191 == v57)
																}
															}
														}
													}
													v209 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeSessionUserId[6]))
													F_LWLockRelease(m, v209+int32(512))
													mBase = m.M
													v213 = m.ExcPending
													if v213 != 0 {
														return
													} else {
														v214 = *(*int32)(unsafe.Add(mBase, uint32(v54)+76))
														if v214 < v198 {
															F_errstart_cold(m, int32(22), int32(0))
															mBase = m.M
															v286 = m.ExcPending
															if v286 != 0 {
																return
															} else {
																F_errcode(m, int32(_a_F_InitializeSessionUserId_4))
																mBase = m.M
																v289 = m.ExcPending
																if v289 != 0 {
																	return
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v64
																	F_errmsg(m, int32(_a_F_InitializeSessionUserId_5), v15+int32(-48))
																	mBase = m.M
																	v295 = m.ExcPending
																	if v295 != 0 {
																		return
																	} else {
																		F_errfinish(m, int32(_a_F_InitializeSessionUserId_2), int32(880), int32(_a_F_InitializeSessionUserId_3))
																		mBase = m.M
																		v300 = m.ExcPending
																		if v300 != 0 {
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
															F_ReleaseCatCache(m, v51)
															mBase = m.M
															v231 = m.ExcPending
															if v231 != 0 {
																return
															} else {
																m.G0 = v17 - int32(-64)
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
					} else {
						F_errstart_cold(m, int32(22), int32(0))
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return
						} else {
							F_errcode(m, int32(514))
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v17)+48)) = l0
								F_errmsg(m, int32(_a_F_InitializeSessionUserId_6), v15+int32(-16))
								mBase = m.M
								v40 = m.ExcPending
								if v40 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_InitializeSessionUserId_2), int32(804), int32(_a_F_InitializeSessionUserId_3))
									mBase = m.M
									v45 = m.ExcPending
									if v45 != 0 {
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
				v47 = F_SearchSysCache1(m, int32(11), l1)
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return
				} else {
					if v47 == int32(0) {
						F_errstart_cold(m, int32(22), int32(0))
						mBase = m.M
						v252 = m.ExcPending
						if v252 != 0 {
							return
						} else {
							F_errcode(m, int32(514))
							mBase = m.M
							v255 = m.ExcPending
							if v255 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v17))) = l1
								F_errmsg(m, int32(_a_F_InitializeSessionUserId_7), v17)
								mBase = m.M
								v259 = m.ExcPending
								if v259 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_InitializeSessionUserId_2), int32(812), int32(_a_F_InitializeSessionUserId_3))
									mBase = m.M
									v264 = m.ExcPending
									if v264 != 0 {
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
						v51 = v47
						v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+16))
						v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+22)))
						v54 = v52 + v53
						v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+68)))
						v57 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
						*(*int32)(unsafe.Add(mBase, _c_F_InitializeSessionUserId[1])) = v57
						v60 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeSessionUserId[2]))
						*(*int32)(unsafe.Add(mBase, uint32(v60)+64)) = v57
						v63 = int32(4)
						v64 = v54 + v63
						F_SetConfigOption(m, int32(_a_F_InitializeSessionUserId_0), v64, v63, int32(10))
						mBase = m.M
						v68 = m.ExcPending
						if v68 != 0 {
							return
						} else {
							v70 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitializeSessionUserId[3])))
							if v70 != int32(1) {
								F_ReleaseCatCache(m, v51)
								mBase = m.M
								v231 = m.ExcPending
								if v231 != 0 {
									return
								} else {
									m.G0 = v17 - int32(-64)
									return
								}
							} else {
								if l2 == int32(0) {
									v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+72)))
									if v75 == int32(0) {
										F_errstart_cold(m, int32(22), int32(0))
										mBase = m.M
										v268 = m.ExcPending
										if v268 != 0 {
											return
										} else {
											F_errcode(m, int32(514))
											mBase = m.M
											v271 = m.ExcPending
											if v271 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v64
												F_errmsg(m, int32(_a_F_InitializeSessionUserId_1), v15+int32(-32))
												mBase = m.M
												v277 = m.ExcPending
												if v277 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_InitializeSessionUserId_2), int32(859), int32(_a_F_InitializeSessionUserId_3))
													mBase = m.M
													v282 = m.ExcPending
													if v282 != 0 {
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
										v78 = *(*int32)(unsafe.Add(mBase, uint32(v54)+76))
										if v78 < int32(0) {
											F_ReleaseCatCache(m, v51)
											mBase = m.M
											v231 = m.ExcPending
											if v231 != 0 {
												return
											} else {
												m.G0 = v17 - int32(-64)
												return
											}
										} else {
											v82 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeSessionUserId[4]))
											if v82 != int32(1) {
												F_ReleaseCatCache(m, v51)
												mBase = m.M
												v231 = m.ExcPending
												if v231 != 0 {
													return
												} else {
													m.G0 = v17 - int32(-64)
													return
												}
											} else {
												if v55&int32(1) != 0 {
													F_ReleaseCatCache(m, v51)
													mBase = m.M
													v231 = m.ExcPending
													if v231 != 0 {
														return
													} else {
														m.G0 = v17 - int32(-64)
														return
													}
												} else {
													v87 = int32(0)
													v89 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeSessionUserId[5]))
													v91 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeSessionUserId[6]))
													v95 = F_LWLockAcquire(m, v91+int32(512), int32(1))
													mBase = m.M
													v96 = m.ExcPending
													if v96 != 0 {
														return
													} else {
														v97 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
														if v97 <= int32(0) {
															v198 = v87
														} else {
															v100 = int32(1)
															v103 = v89 + int32(36)
															v105 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeSessionUserId[7]))
															v106 = int32(0)
															if v97 != v100 {
																v114 = v106
																v116 = v87
																v117 = int32(0)
																for {
																	v128 = v103 + v114<<(uint(int32(2))%32)
																	v129 = *(*int32)(unsafe.Add(mBase, uint32(v128)))
																	v132 = v105 + v129*int32(640)
																	v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)+44))
																	if v133 == int32(0) {
																		v142 = v116
																	} else {
																		v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132)+72)))
																		if v136 != int32(1) {
																			v142 = v116
																		} else {
																			v139 = *(*int32)(unsafe.Add(mBase, uint32(v132)+64))
																			v142 = v116 + base.B2i32(v139 == v57)
																		}
																	}
																	v143 = *(*int32)(unsafe.Add(mBase, uint32(v128)+4))
																	v146 = v105 + v143*int32(640)
																	v147 = *(*int32)(unsafe.Add(mBase, uint32(v146)+44))
																	if v147 == int32(0) {
																		v156 = v142
																	} else {
																		v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+72)))
																		if v150 != int32(1) {
																			v156 = v142
																		} else {
																			v153 = *(*int32)(unsafe.Add(mBase, uint32(v146)+64))
																			v156 = v142 + base.B2i32(v153 == v57)
																		}
																	}
																	v157 = int32(2)
																	v158 = v114 + v157
																	v160 = v117 + v157
																	if v160 != v97&int32(2147483646) {
																		v114 = v158
																		v116 = v156
																		v117 = v160
																		continue
																	} else {
																		break
																	}
																	break
																}
																v164 = v158
																v166 = v156
															} else {
																v164 = v106
																v166 = v87
															}
															if v97&v100 == int32(0) {
																v198 = v166
															} else {
																v181 = *(*int32)(unsafe.Add(mBase, uint32(v103+v164<<(uint(int32(2))%32))))
																v184 = v105 + v181*int32(640)
																v185 = *(*int32)(unsafe.Add(mBase, uint32(v184)+44))
																if v185 == int32(0) {
																	v198 = v166
																} else {
																	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184)+72)))
																	if v188 != int32(1) {
																		v198 = v166
																	} else {
																		v191 = *(*int32)(unsafe.Add(mBase, uint32(v184)+64))
																		v198 = v166 + base.B2i32(v191 == v57)
																	}
																}
															}
														}
														v209 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeSessionUserId[6]))
														F_LWLockRelease(m, v209+int32(512))
														mBase = m.M
														v213 = m.ExcPending
														if v213 != 0 {
															return
														} else {
															v214 = *(*int32)(unsafe.Add(mBase, uint32(v54)+76))
															if v214 < v198 {
																F_errstart_cold(m, int32(22), int32(0))
																mBase = m.M
																v286 = m.ExcPending
																if v286 != 0 {
																	return
																} else {
																	F_errcode(m, int32(_a_F_InitializeSessionUserId_4))
																	mBase = m.M
																	v289 = m.ExcPending
																	if v289 != 0 {
																		return
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v64
																		F_errmsg(m, int32(_a_F_InitializeSessionUserId_5), v15+int32(-48))
																		mBase = m.M
																		v295 = m.ExcPending
																		if v295 != 0 {
																			return
																		} else {
																			F_errfinish(m, int32(_a_F_InitializeSessionUserId_2), int32(880), int32(_a_F_InitializeSessionUserId_3))
																			mBase = m.M
																			v300 = m.ExcPending
																			if v300 != 0 {
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
																F_ReleaseCatCache(m, v51)
																mBase = m.M
																v231 = m.ExcPending
																if v231 != 0 {
																	return
																} else {
																	m.G0 = v17 - int32(-64)
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
									v78 = *(*int32)(unsafe.Add(mBase, uint32(v54)+76))
									if v78 < int32(0) {
										F_ReleaseCatCache(m, v51)
										mBase = m.M
										v231 = m.ExcPending
										if v231 != 0 {
											return
										} else {
											m.G0 = v17 - int32(-64)
											return
										}
									} else {
										v82 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeSessionUserId[4]))
										if v82 != int32(1) {
											F_ReleaseCatCache(m, v51)
											mBase = m.M
											v231 = m.ExcPending
											if v231 != 0 {
												return
											} else {
												m.G0 = v17 - int32(-64)
												return
											}
										} else {
											if v55&int32(1) != 0 {
												F_ReleaseCatCache(m, v51)
												mBase = m.M
												v231 = m.ExcPending
												if v231 != 0 {
													return
												} else {
													m.G0 = v17 - int32(-64)
													return
												}
											} else {
												v87 = int32(0)
												v89 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeSessionUserId[5]))
												v91 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeSessionUserId[6]))
												v95 = F_LWLockAcquire(m, v91+int32(512), int32(1))
												mBase = m.M
												v96 = m.ExcPending
												if v96 != 0 {
													return
												} else {
													v97 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
													if v97 <= int32(0) {
														v198 = v87
													} else {
														v100 = int32(1)
														v103 = v89 + int32(36)
														v105 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeSessionUserId[7]))
														v106 = int32(0)
														if v97 != v100 {
															v114 = v106
															v116 = v87
															v117 = int32(0)
															for {
																v128 = v103 + v114<<(uint(int32(2))%32)
																v129 = *(*int32)(unsafe.Add(mBase, uint32(v128)))
																v132 = v105 + v129*int32(640)
																v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)+44))
																if v133 == int32(0) {
																	v142 = v116
																} else {
																	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132)+72)))
																	if v136 != int32(1) {
																		v142 = v116
																	} else {
																		v139 = *(*int32)(unsafe.Add(mBase, uint32(v132)+64))
																		v142 = v116 + base.B2i32(v139 == v57)
																	}
																}
																v143 = *(*int32)(unsafe.Add(mBase, uint32(v128)+4))
																v146 = v105 + v143*int32(640)
																v147 = *(*int32)(unsafe.Add(mBase, uint32(v146)+44))
																if v147 == int32(0) {
																	v156 = v142
																} else {
																	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+72)))
																	if v150 != int32(1) {
																		v156 = v142
																	} else {
																		v153 = *(*int32)(unsafe.Add(mBase, uint32(v146)+64))
																		v156 = v142 + base.B2i32(v153 == v57)
																	}
																}
																v157 = int32(2)
																v158 = v114 + v157
																v160 = v117 + v157
																if v160 != v97&int32(2147483646) {
																	v114 = v158
																	v116 = v156
																	v117 = v160
																	continue
																} else {
																	break
																}
																break
															}
															v164 = v158
															v166 = v156
														} else {
															v164 = v106
															v166 = v87
														}
														if v97&v100 == int32(0) {
															v198 = v166
														} else {
															v181 = *(*int32)(unsafe.Add(mBase, uint32(v103+v164<<(uint(int32(2))%32))))
															v184 = v105 + v181*int32(640)
															v185 = *(*int32)(unsafe.Add(mBase, uint32(v184)+44))
															if v185 == int32(0) {
																v198 = v166
															} else {
																v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184)+72)))
																if v188 != int32(1) {
																	v198 = v166
																} else {
																	v191 = *(*int32)(unsafe.Add(mBase, uint32(v184)+64))
																	v198 = v166 + base.B2i32(v191 == v57)
																}
															}
														}
													}
													v209 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeSessionUserId[6]))
													F_LWLockRelease(m, v209+int32(512))
													mBase = m.M
													v213 = m.ExcPending
													if v213 != 0 {
														return
													} else {
														v214 = *(*int32)(unsafe.Add(mBase, uint32(v54)+76))
														if v214 < v198 {
															F_errstart_cold(m, int32(22), int32(0))
															mBase = m.M
															v286 = m.ExcPending
															if v286 != 0 {
																return
															} else {
																F_errcode(m, int32(_a_F_InitializeSessionUserId_4))
																mBase = m.M
																v289 = m.ExcPending
																if v289 != 0 {
																	return
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v64
																	F_errmsg(m, int32(_a_F_InitializeSessionUserId_5), v15+int32(-48))
																	mBase = m.M
																	v295 = m.ExcPending
																	if v295 != 0 {
																		return
																	} else {
																		F_errfinish(m, int32(_a_F_InitializeSessionUserId_2), int32(880), int32(_a_F_InitializeSessionUserId_3))
																		mBase = m.M
																		v300 = m.ExcPending
																		if v300 != 0 {
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
															F_ReleaseCatCache(m, v51)
															mBase = m.M
															v231 = m.ExcPending
															if v231 != 0 {
																return
															} else {
																m.G0 = v17 - int32(-64)
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
	} else {
		m.G0 = v17 - int32(-64)
		return
	}
}
func F_InitializeShmemGUCs(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	v5 = m.G0
	v7 = v5 - int32(112)
	m.G0 = v7
	v11 = F_CalculateShmemSize(m, v7+int32(40))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		v14 = F_add_size(m, v11, int32(_a_F_InitializeShmemGUCs_0))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = int32(base.Ui32(v14) >> (uint(int32(20)) % 32))
			v24 = F_pg_sprintf(m, v7+int32(48), int32(_a_F_InitializeShmemGUCs_1), v7+int32(32))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return
			} else {
				F_SetConfigOption(m, int32(_a_F_InitializeShmemGUCs_2), v7+int32(48), int32(0), int32(1))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return
				} else {
					v34 = v7 + int32(44)
					v36 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeShmemGUCs[0]))
					if v36 != 0 {
						v40 = v36 << (uint(int32(10)) % 32)
					} else {
						v40 = int32(_a_F_InitializeShmemGUCs_3)
					}
					if v40 != 0 {
					} else {
					}
					if v34 != 0 {
						*(*int32)(unsafe.Add(mBase, uint32(v34))) = v40
					} else {
					}
					v53 = *(*int32)(unsafe.Add(mBase, uint32(v7)+44))
					if v53 != 0 {
						v54 = base.I32_div_u_s(v11, v53)
						v56 = F_add_size(m, v54, int32(1))
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v56
							v64 = F_pg_sprintf(m, v7+int32(48), int32(_a_F_InitializeShmemGUCs_1), v7+int32(16))
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return
							} else {
								F_SetConfigOption(m, int32(_a_F_InitializeShmemGUCs_4), v7+int32(48), int32(0), int32(1))
								mBase = m.M
								v72 = m.ExcPending
								if v72 != 0 {
									return
								} else {
									v73 = *(*int32)(unsafe.Add(mBase, uint32(v7)+40))
									*(*int32)(unsafe.Add(mBase, uint32(v7))) = v73
									v78 = F_pg_sprintf(m, v7+int32(48), int32(_a_F_InitializeShmemGUCs_5), v7)
									mBase = m.M
									v79 = m.ExcPending
									if v79 != 0 {
										return
									} else {
										F_SetConfigOption(m, int32(_a_F_InitializeShmemGUCs_6), v7+int32(48), int32(0), int32(1))
										mBase = m.M
										v86 = m.ExcPending
										if v86 != 0 {
											return
										} else {
											m.G0 = v7 + int32(112)
											return
										}
									}
								}
							}
						}
					} else {
						v73 = *(*int32)(unsafe.Add(mBase, uint32(v7)+40))
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = v73
						v78 = F_pg_sprintf(m, v7+int32(48), int32(_a_F_InitializeShmemGUCs_5), v7)
						mBase = m.M
						v79 = m.ExcPending
						if v79 != 0 {
							return
						} else {
							F_SetConfigOption(m, int32(_a_F_InitializeShmemGUCs_6), v7+int32(48), int32(0), int32(1))
							mBase = m.M
							v86 = m.ExcPending
							if v86 != 0 {
								return
							} else {
								m.G0 = v7 + int32(112)
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_Int64GetDatum(m *base.Module, l0 int64) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_palloc(m, int32(8))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		*(*int64)(unsafe.Add(mBase, uint32(v4))) = l0
		return v4
	}
}
func F_IsThereFunctionInNamespace(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
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
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v12 = F_SearchSysCacheExists(m, int32(46), l0, l2, l3, int32(0))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		if v12 != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return
			} else {
				F_errcode(m, int32(50884740))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return
				} else {
					v24 = F_funcname_signature_string(m, l0, l1, int32(0), l2+int32(24))
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return
					} else {
						v26 = F_get_namespace_name(m, l3)
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v26
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = v24
							F_errmsg(m, int32(_a_F_IsThereFunctionInNamespace_0), v8)
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_IsThereFunctionInNamespace_1), int32(2074), int32(_a_F_IsThereFunctionInNamespace_2))
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
		} else {
			m.G0 = v8 + int32(16)
			return
		}
	}
}
func F__intbig_contains(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v10 = F_ArrayGetNItems(m, v7, l1+int32(16))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v14 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L18
	}
L4:
	;
	if v10 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L5:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v26 = (v17<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L4
L6:
	;
	goto L7
L7:
	;
	v24 = F_array_contains_nulls(m, l1)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	if v24 != 0 {
		goto L3
	} else {
		goto L9
	}
L9:
	;
	v26 = v14
	goto L4
L10:
	;
	return int32(1)
L11:
	;
	goto L12
L12:
	;
	v37 = l1 + v26
	v40 = v10
	goto L13
L13:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	v44 = base.I32_rem_u_s(v43, l2<<(uint(int32(3))%32))
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(8)+int32(base.Ui32(v44)>>(uint(int32(3))%32))))))
	v52 = int32(1) << (uint(v44&int32(7)) % 32) & v51
	v53 = int32(0)
	if v52 == v53 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	return base.B2i32(v52 != v53)
L15:
	;
	goto L14
L16:
	;
	v60 = v40 - int32(1)
	if v60 != 0 {
		v37 = v37 + int32(4)
		v40 = v60
		goto L13
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	F_errmsg(m, int32(_a_F__intbig_contains_0), int32(0))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	F_errfinish(m, int32(_a_F__intbig_contains_1), int32(100), int32(_a_F__intbig_contains_2))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_i4tof(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	return base.I32_reinterpret_f32(base.F32_convert_i32_s(v2))
}
func F_icu_unicode_version(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v2)
	return int32(0)
}
func F_indexam_property(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
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
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v83 int32
	_ = v83
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v106 int32
	_ = v106
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v128 int32
	_ = v128
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v151 int32
	_ = v151
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v173 int32
	_ = v173
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v196 int32
	_ = v196
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v218 int32
	_ = v218
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v241 int32
	_ = v241
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v263 int32
	_ = v263
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v286 int32
	_ = v286
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v308 int32
	_ = v308
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v331 int32
	_ = v331
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v353 int32
	_ = v353
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v376 int32
	_ = v376
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v398 int32
	_ = v398
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v421 int32
	_ = v421
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v443 int32
	_ = v443
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v466 int32
	_ = v466
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v488 int32
	_ = v488
	var v497 int32
	_ = v497
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v511 int32
	_ = v511
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v533 int32
	_ = v533
	var v542 int32
	_ = v542
	var v545 int32
	_ = v545
	var v547 int32
	_ = v547
	var v556 int32
	_ = v556
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v578 int32
	_ = v578
	var v587 int32
	_ = v587
	var v590 int32
	_ = v590
	var v592 int32
	_ = v592
	var v601 int32
	_ = v601
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v623 int32
	_ = v623
	var v632 int32
	_ = v632
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v646 int32
	_ = v646
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v668 int32
	_ = v668
	var v677 int32
	_ = v677
	var v680 int32
	_ = v680
	var v682 int32
	_ = v682
	var v691 int32
	_ = v691
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v713 int32
	_ = v713
	var v722 int32
	_ = v722
	var v725 int32
	_ = v725
	var v727 int32
	_ = v727
	var v736 int32
	_ = v736
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v758 int32
	_ = v758
	var v767 int32
	_ = v767
	var v770 int32
	_ = v770
	var v772 int32
	_ = v772
	var v781 int32
	_ = v781
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v802 int32
	_ = v802
	var v811 int32
	_ = v811
	var v814 int32
	_ = v814
	var v816 int32
	_ = v816
	var v825 int32
	_ = v825
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v831 int32
	_ = v831
	var v834 int32
	_ = v834
	var v837 int32
	_ = v837
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v857 int32
	_ = v857
	var v859 int32
	_ = v859
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v866 int32
	_ = v866
	var v868 int32
	_ = v868
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v879 int32
	_ = v879
	var v882 int32
	_ = v882
	var v884 int32
	_ = v884
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v890 int32
	_ = v890
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v897 int32
	_ = v897
	var v901 int32
	_ = v901
	var v903 int32
	_ = v903
	var v909 int32
	_ = v909
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v919 int32
	_ = v919
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v931 int32
	_ = v931
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v941 int32
	_ = v941
	var v944 int32
	_ = v944
	var v945 int32
	_ = v945
	var v951 int32
	_ = v951
	var v956 int32
	_ = v956
	var v957 int32
	_ = v957
	var v958 int32
	_ = v958
	var v961 int32
	_ = v961
	var v966 int32
	_ = v966
	var v967 int32
	_ = v967
	var v973 int32
	_ = v973
	var v978 int32
	_ = v978
	var v979 int32
	_ = v979
	var v983 int32
	_ = v983
	var v988 int32
	_ = v988
	var v989 int32
	_ = v989
	var v993 int32
	_ = v993
	var v994 int32
	_ = v994
	var v995 int32
	_ = v995
	var v998 int32
	_ = v998
	var v999 int32
	_ = v999
	var v1003 int32
	_ = v1003
	var v1007 int32
	_ = v1007
	var v1011 int32
	_ = v1011
	var v1012 int32
	_ = v1012
	var v1013 int32
	_ = v1013
	var v1014 int32
	_ = v1014
	var v1018 int32
	_ = v1018
	var v1021 int32
	_ = v1021
	var v1022 int32
	_ = v1022
	var v1027 int32
	_ = v1027
	var v1028 int32
	_ = v1028
	var v1033 int32
	_ = v1033
	var v1034 int32
	_ = v1034
	var v1037 int32
	_ = v1037
	var v1038 int32
	_ = v1038
	var v1042 int32
	_ = v1042
	var v1043 int32
	_ = v1043
	var v1046 int32
	_ = v1046
	var v1049 int32
	_ = v1049
	var v1050 int32
	_ = v1050
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	var v1056 int32
	_ = v1056
	var v1057 int32
	_ = v1057
	var v1060 int32
	_ = v1060
	var v1061 int32
	_ = v1061
	var v1076 int32
	_ = v1076
	v6 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+15)) = uint8(v6)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+14)) = uint8(v6)
	v23 = int32(_a_F_indexam_property_0)
	v24 = l1
	goto L4
L1:
	;
	if l3 != 0 {
		goto L257
	} else {
		goto L258
	}
L2:
	;
	v828 = *(*int32)(unsafe.Add(mBase, uint32(v827)+4))
	v829 = v828
	goto L1
L3:
	;
	if v61 == int32(0) {
		v827 = int32(_a_F_indexam_property_1)
		goto L2
	} else {
		goto L16
	}
L4:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	if v27 == v28 {
		v50 = v27
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v61 = int32(0)
	goto L3
L6:
	;
	v52 = int32(1)
	if v50 != 0 {
		v23 = v23 + v52
		v24 = v24 + v52
		goto L4
	} else {
		goto L15
	}
L7:
	;
	if base.Ui32((v27-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v38 = v27 | int32(32)
	goto L10
L9:
	;
	v38 = v27
	goto L10
L10:
	;
	if base.Ui32((v28-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v47 = v28 | int32(32)
	goto L13
L12:
	;
	v47 = v28
	goto L13
L13:
	;
	if v38 == v47 {
		v50 = v38
		goto L6
	} else {
		goto L14
	}
L14:
	;
	v61 = v38 - v47
	goto L3
L15:
	;
	goto L5
L16:
	;
	v68 = int32(_a_F_indexam_property_2)
	v69 = l1
	goto L18
L17:
	;
	if v106 == int32(0) {
		v827 = int32(_a_F_indexam_property_3)
		goto L2
	} else {
		goto L30
	}
L18:
	;
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68))))
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69))))
	if v72 == v73 {
		v95 = v72
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v106 = int32(0)
	goto L17
L20:
	;
	v97 = int32(1)
	if v95 != 0 {
		v68 = v68 + v97
		v69 = v69 + v97
		goto L18
	} else {
		goto L29
	}
L21:
	;
	if base.Ui32((v72-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v83 = v72 | int32(32)
	goto L24
L23:
	;
	v83 = v72
	goto L24
L24:
	;
	if base.Ui32((v73-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v92 = v73 | int32(32)
	goto L27
L26:
	;
	v92 = v73
	goto L27
L27:
	;
	if v83 == v92 {
		v95 = v83
		goto L20
	} else {
		goto L28
	}
L28:
	;
	v106 = v83 - v92
	goto L17
L29:
	;
	goto L19
L30:
	;
	v113 = int32(_a_F_indexam_property_4)
	v114 = l1
	goto L32
L31:
	;
	if v151 == int32(0) {
		v827 = int32(_a_F_indexam_property_5)
		goto L2
	} else {
		goto L44
	}
L32:
	;
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113))))
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114))))
	if v117 == v118 {
		v140 = v117
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v151 = int32(0)
	goto L31
L34:
	;
	v142 = int32(1)
	if v140 != 0 {
		v113 = v113 + v142
		v114 = v114 + v142
		goto L32
	} else {
		goto L43
	}
L35:
	;
	if base.Ui32((v117-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v128 = v117 | int32(32)
	goto L38
L37:
	;
	v128 = v117
	goto L38
L38:
	;
	if base.Ui32((v118-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v137 = v118 | int32(32)
	goto L41
L40:
	;
	v137 = v118
	goto L41
L41:
	;
	if v128 == v137 {
		v140 = v128
		goto L34
	} else {
		goto L42
	}
L42:
	;
	v151 = v128 - v137
	goto L31
L43:
	;
	goto L33
L44:
	;
	v158 = int32(_a_F_indexam_property_6)
	v159 = l1
	goto L46
L45:
	;
	if v196 == int32(0) {
		v827 = int32(_a_F_indexam_property_7)
		goto L2
	} else {
		goto L58
	}
L46:
	;
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158))))
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159))))
	if v162 == v163 {
		v185 = v162
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v196 = int32(0)
	goto L45
L48:
	;
	v187 = int32(1)
	if v185 != 0 {
		v158 = v158 + v187
		v159 = v159 + v187
		goto L46
	} else {
		goto L57
	}
L49:
	;
	if base.Ui32((v162-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v173 = v162 | int32(32)
	goto L52
L51:
	;
	v173 = v162
	goto L52
L52:
	;
	if base.Ui32((v163-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v182 = v163 | int32(32)
	goto L55
L54:
	;
	v182 = v163
	goto L55
L55:
	;
	if v173 == v182 {
		v185 = v173
		goto L48
	} else {
		goto L56
	}
L56:
	;
	v196 = v173 - v182
	goto L45
L57:
	;
	goto L47
L58:
	;
	v203 = int32(_a_F_indexam_property_8)
	v204 = l1
	goto L60
L59:
	;
	if v241 == int32(0) {
		v827 = int32(_a_F_indexam_property_9)
		goto L2
	} else {
		goto L72
	}
L60:
	;
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v203))))
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v204))))
	if v207 == v208 {
		v230 = v207
		goto L62
	} else {
		goto L63
	}
L61:
	;
	v241 = int32(0)
	goto L59
L62:
	;
	v232 = int32(1)
	if v230 != 0 {
		v203 = v203 + v232
		v204 = v204 + v232
		goto L60
	} else {
		goto L71
	}
L63:
	;
	if base.Ui32((v207-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v218 = v207 | int32(32)
	goto L66
L65:
	;
	v218 = v207
	goto L66
L66:
	;
	if base.Ui32((v208-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v227 = v208 | int32(32)
	goto L69
L68:
	;
	v227 = v208
	goto L69
L69:
	;
	if v218 == v227 {
		v230 = v218
		goto L62
	} else {
		goto L70
	}
L70:
	;
	v241 = v218 - v227
	goto L59
L71:
	;
	goto L61
L72:
	;
	v248 = int32(_a_F_indexam_property_10)
	v249 = l1
	goto L74
L73:
	;
	if v286 == int32(0) {
		v827 = int32(_a_F_indexam_property_11)
		goto L2
	} else {
		goto L86
	}
L74:
	;
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v248))))
	v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v249))))
	if v252 == v253 {
		v275 = v252
		goto L76
	} else {
		goto L77
	}
L75:
	;
	v286 = int32(0)
	goto L73
L76:
	;
	v277 = int32(1)
	if v275 != 0 {
		v248 = v248 + v277
		v249 = v249 + v277
		goto L74
	} else {
		goto L85
	}
L77:
	;
	if base.Ui32((v252-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v263 = v252 | int32(32)
	goto L80
L79:
	;
	v263 = v252
	goto L80
L80:
	;
	if base.Ui32((v253-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v272 = v253 | int32(32)
	goto L83
L82:
	;
	v272 = v253
	goto L83
L83:
	;
	if v263 == v272 {
		v275 = v263
		goto L76
	} else {
		goto L84
	}
L84:
	;
	v286 = v263 - v272
	goto L73
L85:
	;
	goto L75
L86:
	;
	v293 = int32(_a_F_indexam_property_12)
	v294 = l1
	goto L88
L87:
	;
	if v331 == int32(0) {
		v827 = int32(_a_F_indexam_property_13)
		goto L2
	} else {
		goto L100
	}
L88:
	;
	v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v293))))
	v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v294))))
	if v297 == v298 {
		v320 = v297
		goto L90
	} else {
		goto L91
	}
L89:
	;
	v331 = int32(0)
	goto L87
L90:
	;
	v322 = int32(1)
	if v320 != 0 {
		v293 = v293 + v322
		v294 = v294 + v322
		goto L88
	} else {
		goto L99
	}
L91:
	;
	if base.Ui32((v297-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v308 = v297 | int32(32)
	goto L94
L93:
	;
	v308 = v297
	goto L94
L94:
	;
	if base.Ui32((v298-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v317 = v298 | int32(32)
	goto L97
L96:
	;
	v317 = v298
	goto L97
L97:
	;
	if v308 == v317 {
		v320 = v308
		goto L90
	} else {
		goto L98
	}
L98:
	;
	v331 = v308 - v317
	goto L87
L99:
	;
	goto L89
L100:
	;
	v338 = int32(_a_F_indexam_property_14)
	v339 = l1
	goto L102
L101:
	;
	if v376 == int32(0) {
		v827 = int32(_a_F_indexam_property_15)
		goto L2
	} else {
		goto L114
	}
L102:
	;
	v342 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v338))))
	v343 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v339))))
	if v342 == v343 {
		v365 = v342
		goto L104
	} else {
		goto L105
	}
L103:
	;
	v376 = int32(0)
	goto L101
L104:
	;
	v367 = int32(1)
	if v365 != 0 {
		v338 = v338 + v367
		v339 = v339 + v367
		goto L102
	} else {
		goto L113
	}
L105:
	;
	if base.Ui32((v342-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v353 = v342 | int32(32)
	goto L108
L107:
	;
	v353 = v342
	goto L108
L108:
	;
	if base.Ui32((v343-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v362 = v343 | int32(32)
	goto L111
L110:
	;
	v362 = v343
	goto L111
L111:
	;
	if v353 == v362 {
		v365 = v353
		goto L104
	} else {
		goto L112
	}
L112:
	;
	v376 = v353 - v362
	goto L101
L113:
	;
	goto L103
L114:
	;
	v383 = int32(_a_F_indexam_property_16)
	v384 = l1
	goto L116
L115:
	;
	if v421 == int32(0) {
		v827 = int32(_a_F_indexam_property_17)
		goto L2
	} else {
		goto L128
	}
L116:
	;
	v387 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v383))))
	v388 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v384))))
	if v387 == v388 {
		v410 = v387
		goto L118
	} else {
		goto L119
	}
L117:
	;
	v421 = int32(0)
	goto L115
L118:
	;
	v412 = int32(1)
	if v410 != 0 {
		v383 = v383 + v412
		v384 = v384 + v412
		goto L116
	} else {
		goto L127
	}
L119:
	;
	if base.Ui32((v387-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v398 = v387 | int32(32)
	goto L122
L121:
	;
	v398 = v387
	goto L122
L122:
	;
	if base.Ui32((v388-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v407 = v388 | int32(32)
	goto L125
L124:
	;
	v407 = v388
	goto L125
L125:
	;
	if v398 == v407 {
		v410 = v398
		goto L118
	} else {
		goto L126
	}
L126:
	;
	v421 = v398 - v407
	goto L115
L127:
	;
	goto L117
L128:
	;
	v428 = int32(_a_F_indexam_property_18)
	v429 = l1
	goto L130
L129:
	;
	if v466 == int32(0) {
		v827 = int32(_a_F_indexam_property_19)
		goto L2
	} else {
		goto L142
	}
L130:
	;
	v432 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v428))))
	v433 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v429))))
	if v432 == v433 {
		v455 = v432
		goto L132
	} else {
		goto L133
	}
L131:
	;
	v466 = int32(0)
	goto L129
L132:
	;
	v457 = int32(1)
	if v455 != 0 {
		v428 = v428 + v457
		v429 = v429 + v457
		goto L130
	} else {
		goto L141
	}
L133:
	;
	if base.Ui32((v432-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v443 = v432 | int32(32)
	goto L136
L135:
	;
	v443 = v432
	goto L136
L136:
	;
	if base.Ui32((v433-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v452 = v433 | int32(32)
	goto L139
L138:
	;
	v452 = v433
	goto L139
L139:
	;
	if v443 == v452 {
		v455 = v443
		goto L132
	} else {
		goto L140
	}
L140:
	;
	v466 = v443 - v452
	goto L129
L141:
	;
	goto L131
L142:
	;
	v473 = int32(_a_F_indexam_property_20)
	v474 = l1
	goto L144
L143:
	;
	if v511 == int32(0) {
		v827 = int32(_a_F_indexam_property_21)
		goto L2
	} else {
		goto L156
	}
L144:
	;
	v477 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v473))))
	v478 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v474))))
	if v477 == v478 {
		v500 = v477
		goto L146
	} else {
		goto L147
	}
L145:
	;
	v511 = int32(0)
	goto L143
L146:
	;
	v502 = int32(1)
	if v500 != 0 {
		v473 = v473 + v502
		v474 = v474 + v502
		goto L144
	} else {
		goto L155
	}
L147:
	;
	if base.Ui32((v477-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	v488 = v477 | int32(32)
	goto L150
L149:
	;
	v488 = v477
	goto L150
L150:
	;
	if base.Ui32((v478-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L151
	} else {
		goto L152
	}
L151:
	;
	v497 = v478 | int32(32)
	goto L153
L152:
	;
	v497 = v478
	goto L153
L153:
	;
	if v488 == v497 {
		v500 = v488
		goto L146
	} else {
		goto L154
	}
L154:
	;
	v511 = v488 - v497
	goto L143
L155:
	;
	goto L145
L156:
	;
	v518 = int32(_a_F_indexam_property_22)
	v519 = l1
	goto L158
L157:
	;
	if v556 == int32(0) {
		v827 = int32(_a_F_indexam_property_23)
		goto L2
	} else {
		goto L170
	}
L158:
	;
	v522 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v518))))
	v523 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v519))))
	if v522 == v523 {
		v545 = v522
		goto L160
	} else {
		goto L161
	}
L159:
	;
	v556 = int32(0)
	goto L157
L160:
	;
	v547 = int32(1)
	if v545 != 0 {
		v518 = v518 + v547
		v519 = v519 + v547
		goto L158
	} else {
		goto L169
	}
L161:
	;
	if base.Ui32((v522-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L162
	} else {
		goto L163
	}
L162:
	;
	v533 = v522 | int32(32)
	goto L164
L163:
	;
	v533 = v522
	goto L164
L164:
	;
	if base.Ui32((v523-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L165
	} else {
		goto L166
	}
L165:
	;
	v542 = v523 | int32(32)
	goto L167
L166:
	;
	v542 = v523
	goto L167
L167:
	;
	if v533 == v542 {
		v545 = v533
		goto L160
	} else {
		goto L168
	}
L168:
	;
	v556 = v533 - v542
	goto L157
L169:
	;
	goto L159
L170:
	;
	v563 = int32(_a_F_indexam_property_24)
	v564 = l1
	goto L172
L171:
	;
	if v601 == int32(0) {
		v827 = int32(_a_F_indexam_property_25)
		goto L2
	} else {
		goto L184
	}
L172:
	;
	v567 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v563))))
	v568 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v564))))
	if v567 == v568 {
		v590 = v567
		goto L174
	} else {
		goto L175
	}
L173:
	;
	v601 = int32(0)
	goto L171
L174:
	;
	v592 = int32(1)
	if v590 != 0 {
		v563 = v563 + v592
		v564 = v564 + v592
		goto L172
	} else {
		goto L183
	}
L175:
	;
	if base.Ui32((v567-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L176
	} else {
		goto L177
	}
L176:
	;
	v578 = v567 | int32(32)
	goto L178
L177:
	;
	v578 = v567
	goto L178
L178:
	;
	if base.Ui32((v568-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L179
	} else {
		goto L180
	}
L179:
	;
	v587 = v568 | int32(32)
	goto L181
L180:
	;
	v587 = v568
	goto L181
L181:
	;
	if v578 == v587 {
		v590 = v578
		goto L174
	} else {
		goto L182
	}
L182:
	;
	v601 = v578 - v587
	goto L171
L183:
	;
	goto L173
L184:
	;
	v608 = int32(_a_F_indexam_property_26)
	v609 = l1
	goto L186
L185:
	;
	if v646 == int32(0) {
		v827 = int32(_a_F_indexam_property_27)
		goto L2
	} else {
		goto L198
	}
L186:
	;
	v612 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v608))))
	v613 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v609))))
	if v612 == v613 {
		v635 = v612
		goto L188
	} else {
		goto L189
	}
L187:
	;
	v646 = int32(0)
	goto L185
L188:
	;
	v637 = int32(1)
	if v635 != 0 {
		v608 = v608 + v637
		v609 = v609 + v637
		goto L186
	} else {
		goto L197
	}
L189:
	;
	if base.Ui32((v612-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L190
	} else {
		goto L191
	}
L190:
	;
	v623 = v612 | int32(32)
	goto L192
L191:
	;
	v623 = v612
	goto L192
L192:
	;
	if base.Ui32((v613-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L193
	} else {
		goto L194
	}
L193:
	;
	v632 = v613 | int32(32)
	goto L195
L194:
	;
	v632 = v613
	goto L195
L195:
	;
	if v623 == v632 {
		v635 = v623
		goto L188
	} else {
		goto L196
	}
L196:
	;
	v646 = v623 - v632
	goto L185
L197:
	;
	goto L187
L198:
	;
	v653 = int32(_a_F_indexam_property_28)
	v654 = l1
	goto L200
L199:
	;
	if v691 == int32(0) {
		v827 = int32(_a_F_indexam_property_29)
		goto L2
	} else {
		goto L212
	}
L200:
	;
	v657 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v653))))
	v658 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v654))))
	if v657 == v658 {
		v680 = v657
		goto L202
	} else {
		goto L203
	}
L201:
	;
	v691 = int32(0)
	goto L199
L202:
	;
	v682 = int32(1)
	if v680 != 0 {
		v653 = v653 + v682
		v654 = v654 + v682
		goto L200
	} else {
		goto L211
	}
L203:
	;
	if base.Ui32((v657-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L204
	} else {
		goto L205
	}
L204:
	;
	v668 = v657 | int32(32)
	goto L206
L205:
	;
	v668 = v657
	goto L206
L206:
	;
	if base.Ui32((v658-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L207
	} else {
		goto L208
	}
L207:
	;
	v677 = v658 | int32(32)
	goto L209
L208:
	;
	v677 = v658
	goto L209
L209:
	;
	if v668 == v677 {
		v680 = v668
		goto L202
	} else {
		goto L210
	}
L210:
	;
	v691 = v668 - v677
	goto L199
L211:
	;
	goto L201
L212:
	;
	v698 = int32(_a_F_indexam_property_30)
	v699 = l1
	goto L214
L213:
	;
	if v736 == int32(0) {
		v827 = int32(_a_F_indexam_property_31)
		goto L2
	} else {
		goto L226
	}
L214:
	;
	v702 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v698))))
	v703 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v699))))
	if v702 == v703 {
		v725 = v702
		goto L216
	} else {
		goto L217
	}
L215:
	;
	v736 = int32(0)
	goto L213
L216:
	;
	v727 = int32(1)
	if v725 != 0 {
		v698 = v698 + v727
		v699 = v699 + v727
		goto L214
	} else {
		goto L225
	}
L217:
	;
	if base.Ui32((v702-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L218
	} else {
		goto L219
	}
L218:
	;
	v713 = v702 | int32(32)
	goto L220
L219:
	;
	v713 = v702
	goto L220
L220:
	;
	if base.Ui32((v703-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L221
	} else {
		goto L222
	}
L221:
	;
	v722 = v703 | int32(32)
	goto L223
L222:
	;
	v722 = v703
	goto L223
L223:
	;
	if v713 == v722 {
		v725 = v713
		goto L216
	} else {
		goto L224
	}
L224:
	;
	v736 = v713 - v722
	goto L213
L225:
	;
	goto L215
L226:
	;
	v743 = int32(_a_F_indexam_property_32)
	v744 = l1
	goto L228
L227:
	;
	if v781 == int32(0) {
		v827 = int32(_a_F_indexam_property_33)
		goto L2
	} else {
		goto L240
	}
L228:
	;
	v747 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v743))))
	v748 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744))))
	if v747 == v748 {
		v770 = v747
		goto L230
	} else {
		goto L231
	}
L229:
	;
	v781 = int32(0)
	goto L227
L230:
	;
	v772 = int32(1)
	if v770 != 0 {
		v743 = v743 + v772
		v744 = v744 + v772
		goto L228
	} else {
		goto L239
	}
L231:
	;
	if base.Ui32((v747-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L232
	} else {
		goto L233
	}
L232:
	;
	v758 = v747 | int32(32)
	goto L234
L233:
	;
	v758 = v747
	goto L234
L234:
	;
	if base.Ui32((v748-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L235
	} else {
		goto L236
	}
L235:
	;
	v767 = v748 | int32(32)
	goto L237
L236:
	;
	v767 = v748
	goto L237
L237:
	;
	if v758 == v767 {
		v770 = v758
		goto L230
	} else {
		goto L238
	}
L238:
	;
	v781 = v758 - v767
	goto L227
L239:
	;
	goto L229
L240:
	;
	v787 = int32(_a_F_indexam_property_34)
	v788 = l1
	goto L242
L241:
	;
	if v825 != 0 {
		v829 = v6
		goto L1
	} else {
		goto L254
	}
L242:
	;
	v791 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v787))))
	v792 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v788))))
	if v791 == v792 {
		v814 = v791
		goto L244
	} else {
		goto L245
	}
L243:
	;
	v825 = int32(0)
	goto L241
L244:
	;
	v816 = int32(1)
	if v814 != 0 {
		v787 = v787 + v816
		v788 = v788 + v816
		goto L242
	} else {
		goto L253
	}
L245:
	;
	if base.Ui32((v791-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L246
	} else {
		goto L247
	}
L246:
	;
	v802 = v791 | int32(32)
	goto L248
L247:
	;
	v802 = v791
	goto L248
L248:
	;
	if base.Ui32((v792-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L249
	} else {
		goto L250
	}
L249:
	;
	v811 = v792 | int32(32)
	goto L251
L250:
	;
	v811 = v792
	goto L251
L251:
	;
	if v802 == v811 {
		v814 = v802
		goto L244
	} else {
		goto L252
	}
L252:
	;
	v825 = v802 - v811
	goto L241
L253:
	;
	goto L243
L254:
	;
	v827 = int32(_a_F_indexam_property_35)
	goto L2
L255:
	;
	m.G0 = v13 + int32(16)
	return v1076
L256:
	;
	v1076 = int32(0)
	goto L255
L257:
	;
	v831 = F_SearchSysCache1(m, int32(57), l3)
	mBase = m.M
	v834 = m.ExcPending
	if v834 != 0 {
		goto L260
	} else {
		goto L261
	}
L258:
	;
	v855 = l2
	v857 = v6
	goto L259
L259:
	;
	if v857 < l4 {
		goto L270
	} else {
		goto L271
	}
L260:
	;
	return int32(0)
L261:
	;
	if v831 == int32(0) {
		goto L262
	} else {
		goto L263
	}
L262:
	;
	v837 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v837)
	goto L256
L263:
	;
	goto L264
L264:
	;
	v839 = *(*int32)(unsafe.Add(mBase, uint32(v831)+16))
	v840 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v839)+22)))
	v841 = v839 + v840
	v842 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v841)+119)))
	if v842|int32(32) != int32(105) {
		goto L265
	} else {
		goto L266
	}
L265:
	;
	F_ReleaseCatCache(m, v831)
	mBase = m.M
	v848 = m.ExcPending
	if v848 != 0 {
		goto L260
	} else {
		goto L268
	}
L266:
	;
	goto L267
L267:
	;
	v851 = int32(*(*int16)(unsafe.Add(mBase, uint32(v841)+120)))
	v852 = *(*int32)(unsafe.Add(mBase, uint32(v841)+84))
	F_ReleaseCatCache(m, v831)
	mBase = m.M
	v854 = m.ExcPending
	if v854 != 0 {
		goto L260
	} else {
		goto L269
	}
L268:
	;
	v849 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v849)
	goto L256
L269:
	;
	v855 = v852
	v857 = v851
	goto L259
L270:
	;
	v859 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v859)
	goto L256
L271:
	;
	goto L272
L272:
	;
	v862 = F_GetIndexAmRoutineByAmId(m, v855, int32(1))
	mBase = m.M
	v863 = m.ExcPending
	if v863 != 0 {
		goto L260
	} else {
		goto L273
	}
L273:
	;
	if v862 == int32(0) {
		goto L274
	} else {
		goto L275
	}
L274:
	;
	v866 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v866)
	goto L256
L275:
	;
	goto L276
L276:
	;
	v868 = *(*int32)(unsafe.Add(mBase, uint32(v862)+76))
	if v868 == int32(0) {
		goto L277
	} else {
		goto L278
	}
L277:
	;
	if l4 != 0 {
		goto L284
	} else {
		goto L285
	}
L278:
	;
	v875 = m.T0[v868].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, l3, l4, v829, l1, v13+int32(15), v13+int32(14))
	mBase = m.M
	v876 = m.ExcPending
	if v876 != 0 {
		goto L260
	} else {
		goto L279
	}
L279:
	;
	if v875 == int32(0) {
		goto L277
	} else {
		goto L280
	}
L280:
	;
	v879 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+14)))
	if v879 == int32(1) {
		goto L281
	} else {
		goto L282
	}
L281:
	;
	v882 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v882)
	goto L256
L282:
	;
	goto L283
L283:
	;
	v884 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+15)))
	v1076 = v884
	goto L255
L284:
	;
	v886 = F_SearchSysCache1(m, int32(34), l3)
	mBase = m.M
	v887 = m.ExcPending
	if v887 != 0 {
		goto L260
	} else {
		goto L287
	}
L285:
	;
	goto L286
L286:
	;
	if l3 != 0 {
		goto L341
	} else {
		goto L342
	}
L287:
	;
	if v886 == int32(0) {
		goto L288
	} else {
		goto L289
	}
L288:
	;
	v890 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v890)
	goto L256
L289:
	;
	goto L290
L290:
	;
	v892 = *(*int32)(unsafe.Add(mBase, uint32(v886)+16))
	v893 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v892)+22)))
	v894 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+14)) = uint8(v894)
	v897 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v862)+26)))
	if v897 == v894 {
		goto L291
	} else {
		goto L292
	}
L291:
	;
	v901 = int32(*(*int16)(unsafe.Add(mBase, uint32(v892+v893)+10)))
	v903 = base.B2i32(l4 <= v901)
	goto L293
L292:
	;
	v903 = v894
	goto L293
L293:
	;
	switch v829 - int32(1) {
	case 0:
		goto L303
	case 1:
		goto L302
	case 2:
		goto L301
	case 3:
		goto L300
	case 4:
		goto L299
	case 5:
		goto L298
	case 6:
		goto L297
	case 7:
		goto L296
	case 8:
		goto L295
	default:
		goto L294
	}
L294:
	;
	F_ReleaseCatCache(m, v886)
	mBase = m.M
	v1033 = m.ExcPending
	if v1033 != 0 {
		goto L260
	} else {
		goto L337
	}
L295:
	;
	if v903 == int32(0) {
		goto L294
	} else {
		goto L336
	}
L296:
	;
	if v903 == int32(0) {
		goto L294
	} else {
		goto L335
	}
L297:
	;
	v1003 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+15)) = uint8(v1003)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+14)) = uint8(v1003)
	v1007 = *(*int32)(unsafe.Add(mBase, uint32(v862)+60))
	if v1007 == v1003 {
		goto L294
	} else {
		goto L331
	}
L298:
	;
	if v903 != 0 {
		goto L327
	} else {
		goto L328
	}
L299:
	;
	if v903 != 0 {
		goto L324
	} else {
		goto L325
	}
L300:
	;
	if v903 == int32(0) {
		goto L294
	} else {
		goto L319
	}
L301:
	;
	if v903 == int32(0) {
		goto L294
	} else {
		goto L314
	}
L302:
	;
	if v903 == int32(0) {
		goto L294
	} else {
		goto L309
	}
L303:
	;
	if v903 == int32(0) {
		goto L294
	} else {
		goto L304
	}
L304:
	;
	v909 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v862)+10)))
	if v909 == int32(1) {
		goto L305
	} else {
		goto L306
	}
L305:
	;
	v914 = F_SysCacheGetAttrNotNull(m, int32(34), v886, int32(19))
	mBase = m.M
	v915 = m.ExcPending
	if v915 != 0 {
		goto L260
	} else {
		goto L308
	}
L306:
	;
	v924 = int32(0)
	goto L307
L307:
	;
	v925 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+14)) = uint8(v925)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+15)) = uint8(v924)
	goto L294
L308:
	;
	v916 = int32(1)
	v919 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v914+l4<<(uint(v916)%32))+22)))
	v924 = (v919 ^ int32(-1)) & v916
	goto L307
L309:
	;
	v931 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v862)+10)))
	if v931 == int32(1) {
		goto L310
	} else {
		goto L311
	}
L310:
	;
	v936 = F_SysCacheGetAttrNotNull(m, int32(34), v886, int32(19))
	mBase = m.M
	v937 = m.ExcPending
	if v937 != 0 {
		goto L260
	} else {
		goto L313
	}
L311:
	;
	v944 = int32(0)
	goto L312
L312:
	;
	v945 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+14)) = uint8(v945)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+15)) = uint8(v944)
	goto L294
L313:
	;
	v938 = int32(1)
	v941 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v936+l4<<(uint(v938)%32))+22)))
	v944 = v941 & v938
	goto L312
L314:
	;
	v951 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v862)+10)))
	if v951 == int32(1) {
		goto L315
	} else {
		goto L316
	}
L315:
	;
	v956 = F_SysCacheGetAttrNotNull(m, int32(34), v886, int32(19))
	mBase = m.M
	v957 = m.ExcPending
	if v957 != 0 {
		goto L260
	} else {
		goto L318
	}
L316:
	;
	v966 = int32(0)
	goto L317
L317:
	;
	v967 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+14)) = uint8(v967)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+15)) = uint8(v966)
	goto L294
L318:
	;
	v958 = int32(1)
	v961 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v956+l4<<(uint(v958)%32))+22)))
	v966 = int32(base.Ui32(v961)>>(uint(v958)%32)) & v958
	goto L317
L319:
	;
	v973 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v862)+10)))
	if v973 == int32(1) {
		goto L320
	} else {
		goto L321
	}
L320:
	;
	v978 = F_SysCacheGetAttrNotNull(m, int32(34), v886, int32(19))
	mBase = m.M
	v979 = m.ExcPending
	if v979 != 0 {
		goto L260
	} else {
		goto L323
	}
L321:
	;
	v988 = int32(0)
	goto L322
L322:
	;
	v989 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+14)) = uint8(v989)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+15)) = uint8(v988)
	goto L294
L323:
	;
	v983 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v978+l4<<(uint(int32(1))%32))+22)))
	v988 = base.B2i32(v983&int32(2) == int32(0))
	goto L322
L324:
	;
	v993 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v862)+10)))
	v994 = v993
	goto L326
L325:
	;
	v994 = int32(0)
	goto L326
L326:
	;
	v995 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+14)) = uint8(v995)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+15)) = uint8(v994)
	goto L294
L327:
	;
	v998 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v862)+11)))
	if v998 != 0 {
		goto L294
	} else {
		goto L330
	}
L328:
	;
	goto L329
L329:
	;
	v999 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+14)) = uint8(v999)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+15)) = uint8(v999)
	goto L294
L330:
	;
	goto L329
L331:
	;
	v1011 = F_index_open(m, l3, int32(1))
	mBase = m.M
	v1012 = m.ExcPending
	if v1012 != 0 {
		goto L260
	} else {
		goto L332
	}
L332:
	;
	v1013 = F_index_can_return(m, v1011, l4)
	mBase = m.M
	v1014 = m.ExcPending
	if v1014 != 0 {
		goto L260
	} else {
		goto L333
	}
L333:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+15)) = uint8(v1013)
	F_relation_close(m, v1011, int32(1))
	mBase = m.M
	v1018 = m.ExcPending
	if v1018 != 0 {
		goto L260
	} else {
		goto L334
	}
L334:
	;
	goto L294
L335:
	;
	v1021 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v862)+19)))
	v1022 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+14)) = uint8(v1022)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+15)) = uint8(v1021)
	goto L294
L336:
	;
	v1027 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v862)+20)))
	v1028 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+14)) = uint8(v1028)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+15)) = uint8(v1027)
	goto L294
L337:
	;
	v1034 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+14)))
	if v1034 == int32(0) {
		goto L338
	} else {
		goto L339
	}
L338:
	;
	v1037 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+15)))
	v1076 = v1037
	goto L255
L339:
	;
	goto L340
L340:
	;
	v1038 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v1038)
	goto L256
L341:
	;
	switch v829 - int32(10) {
	case 0:
		goto L348
	case 1:
		goto L347
	case 2:
		goto L346
	case 3:
		goto L345
	default:
		goto L344
	}
L342:
	;
	goto L343
L343:
	;
	switch v829 - int32(14) {
	case 0:
		goto L354
	case 1:
		goto L353
	case 2:
		goto L352
	case 3:
		goto L351
	case 4:
		goto L350
	default:
		goto L349
	}
L344:
	;
	v1050 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v1050)
	goto L256
L345:
	;
	v1049 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v862)+15)))
	v1076 = v1049
	goto L255
L346:
	;
	v1046 = *(*int32)(unsafe.Add(mBase, uint32(v862)+104))
	v1076 = base.B2i32(v1046 != int32(0))
	goto L255
L347:
	;
	v1043 = *(*int32)(unsafe.Add(mBase, uint32(v862)+100))
	v1076 = base.B2i32(v1043 != int32(0))
	goto L255
L348:
	;
	v1042 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v862)+22)))
	v1076 = v1042
	goto L255
L349:
	;
	v1061 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v1061)
	goto L256
L350:
	;
	v1060 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v862)+26)))
	v1076 = v1060
	goto L255
L351:
	;
	v1057 = *(*int32)(unsafe.Add(mBase, uint32(v862)+100))
	v1076 = base.B2i32(v1057 != int32(0))
	goto L255
L352:
	;
	v1056 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v862)+17)))
	v1076 = v1056
	goto L255
L353:
	;
	v1055 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v862)+16)))
	v1076 = v1055
	goto L255
L354:
	;
	v1054 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v862)+10)))
	v1076 = v1054
	goto L255
}
func F_inetmi(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int64
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
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
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v56 int64
	_ = v56
	var v57 int64
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v80 int64
	_ = v80
	var v86 int64
	_ = v86
	var v98 int64
	_ = v98
	var v102 int64
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	v8 = int64(0)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum_packed(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v17 = F_pg_detoast_datum_packed(m, v16)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v19 = int32(1)
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v22&v19 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L40
	}
L5:
	;
	v25 = v19
	goto L7
L6:
	;
	v25 = int32(4)
	goto L7
L7:
	;
	v26 = v12 + v25
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	v28 = int32(1)
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if v30&v28 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v33 = v28
	goto L10
L9:
	;
	v33 = int32(4)
	goto L10
L10:
	;
	v34 = v17 + v33
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34))))
	if v27 == v35 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v37 = int32(2)
	if v27 == v37 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L36
	}
L14:
	;
	v45 = int32(3)
	goto L16
L15:
	;
	v45 = int32(15)
	goto L16
L16:
	;
	v49 = v45
	v51 = v19
	v56 = v8
	v57 = v8
	goto L17
L17:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49+(v26+v37)))))
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49+(v34+v37)))))
	v64 = int32(255)
	v66 = v51 + v60 + (v63 ^ v64)
	v68 = v66 & v64
	if base.Ui64(v56) <= base.Ui64(int64(7)) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	if v27 == int32(2) {
		goto L29
	} else {
		goto L30
	}
L19:
	;
	v86 = v56 + int64(1)
	if v86 != base.I64_extend_i32_u(v45+int32(1)) {
		v49 = v49 - int32(1)
		v51 = int32(base.Ui32(v66) >> (uint(int32(8)) % 32))
		v56 = v86
		v57 = v80
		goto L17
	} else {
		goto L28
	}
L20:
	;
	v80 = base.I64_extend_i32_u(v68)<<(uint(v56<<(uint(int64(3))%64))%64) | v57
	goto L19
L21:
	;
	goto L22
L22:
	;
	if v57 < int64(0) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	if v68 == int32(255) {
		v80 = v57
		goto L19
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	if v68 != 0 {
		goto L4
	} else {
		goto L27
	}
L26:
	;
	goto L4
L27:
	;
	v80 = v57
	goto L19
L28:
	;
	goto L18
L29:
	;
	v98 = int64(-1) << (uint(base.I64_extend_i32_u(v45<<(uint(int32(3))%32)+int32(8))) % 64)
	goto L31
L30:
	;
	v98 = int64(0)
	goto L31
L31:
	;
	if base.Ui32(v66) < base.Ui32(int32(256)) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v102 = v98
	goto L34
L33:
	;
	v102 = int64(0)
	goto L34
L34:
	;
	v104 = F_Int64GetDatum(m, v102|v80)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	return v104
L36:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	F_errmsg(m, int32(_a_F_inetmi_0), int32(0))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(_a_F_inetmi_1), int32(1992), int32(_a_F_inetmi_2))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L40:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	F_errmsg(m, int32(_a_F_inetmi_3), int32(0))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	F_errfinish(m, int32(_a_F_inetmi_1), int32(2027), int32(_a_F_inetmi_2))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_inetpl(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int64
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum_packed(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v8 = *(*int64)(unsafe.Add(mBase, uint32(v7)))
		v9 = F_internal_inetpl(m, v3, v8)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			return v9
		}
	}
}
func F_infix_3(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
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
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v218 int32
	_ = v218
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
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
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v274 int32
	_ = v274
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
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
	var v322 int32
	_ = v322
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
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
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v410 int32
	_ = v410
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	v11 = m.G0
	v13 = v11 - int32(48)
	m.G0 = v13
	v16 = l1
	goto L3
L1:
	;
	m.G0 = v13 + int32(48)
	return
L2:
	;
	v233 = v27 - int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v233
	v237 = v16 | base.B2i32(v83 != int32(124))
	if v237&int32(1) == int32(0) {
		goto L44
	} else {
		goto L45
	}
L3:
	;
	F_check_stack_depth(m)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v141 = v139 - v140
	v143 = v141 + int32(3)
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v144 <= v143 {
		goto L27
	} else {
		goto L28
	}
L5:
	;
	return
L6:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v28 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v27))))
	if v28 == int32(2) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v33 = v31 - v32
	v35 = v33 + int32(12)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v36 <= v35 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	if v83 != int32(33) {
		goto L2
	} else {
		goto L18
	}
L10:
	;
	v39 = v32
	v42 = v36
	goto L13
L11:
	;
	v64 = v31
	v69 = v27
	goto L12
L12:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v70
	v73 = F_pg_sprintf(m, v64, int32(_a_F_infix_3_0), v13)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L5
	} else {
		goto L17
	}
L13:
	;
	v49 = v42 << (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v49
	v51 = F_repalloc(m, v39, v49)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L5
	} else {
		goto L15
	}
L14:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v64 = v54
	v69 = v58
	goto L12
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v51
	v54 = v51 + v33
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v54
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v56 <= v35 {
		v39 = v51
		v42 = v56
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v76 = F_strlen(m, v75)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v76 + v75
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v79 - int32(8)
	goto L1
L18:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v88 = v86 - v87
	v90 = v88 + int32(2)
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v91 <= v90 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v94 = v91
	v95 = v87
	goto L22
L20:
	;
	v117 = v86
	goto L21
L21:
	;
	v123 = int32(33)
	*(*uint8)(unsafe.Add(mBase, uint32(v117))) = uint8(v123)
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v125 + int32(1)
	v129 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v125)+1)) = uint8(v129)
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v134 = v132 - int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v134
	v136 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v134))))
	if v136 != int32(3) {
		v16 = v129
		goto L3
	} else {
		goto L26
	}
L22:
	;
	v104 = v94 << (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v104
	v106 = F_repalloc(m, v95, v104)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L5
	} else {
		goto L24
	}
L23:
	;
	v117 = v109
	goto L21
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v106
	v109 = v106 + v88
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v109
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v111 <= v90 {
		v94 = v111
		v95 = v106
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	goto L4
L27:
	;
	v147 = v144
	v148 = v140
	goto L30
L28:
	;
	v170 = v139
	goto L29
L29:
	;
	v178 = F_pg_sprintf(m, v170, int32(_a_F_infix_3_1), int32(0))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L5
	} else {
		goto L34
	}
L30:
	;
	v157 = v147 << (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v157
	v159 = F_repalloc(m, v148, v157)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L5
	} else {
		goto L32
	}
L31:
	;
	v170 = v162
	goto L29
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v159
	v162 = v159 + v141
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v162
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v164 <= v143 {
		v147 = v164
		v148 = v159
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v181 = F_strlen(m, v180)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v181 + v180
	F_infix_3(m, l0, int32(1))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L5
	} else {
		goto L35
	}
L35:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v189 = v187 - v188
	v191 = v189 + int32(3)
	v192 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v192 <= v191 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v195 = v192
	v196 = v188
	goto L39
L37:
	;
	v218 = v187
	goto L38
L38:
	;
	v226 = F_pg_sprintf(m, v218, int32(_a_F_infix_3_2), int32(0))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L5
	} else {
		goto L43
	}
L39:
	;
	v205 = v195 << (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v205
	v207 = F_repalloc(m, v196, v205)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L5
	} else {
		goto L41
	}
L40:
	;
	v218 = v210
	goto L38
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v207
	v210 = v207 + v189
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v210
	v212 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v212 <= v191 {
		v195 = v212
		v196 = v207
		goto L39
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v229 = F_strlen(m, v228)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v229 + v228
	goto L1
L44:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v244 = v242 - v243
	v246 = v244 + int32(3)
	v247 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v247 <= v246 {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	v290 = v233
	goto L46
L46:
	;
	v298 = int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+44)) = v298
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v290
	v302 = F_palloc(m, v298)
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L5
	} else {
		goto L55
	}
L47:
	;
	v250 = v247
	v251 = v243
	goto L50
L48:
	;
	v274 = v242
	goto L49
L49:
	;
	v281 = F_pg_sprintf(m, v274, int32(_a_F_infix_3_1), int32(0))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L5
	} else {
		goto L54
	}
L50:
	;
	v260 = v250 << (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v260
	v262 = F_repalloc(m, v251, v260)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L5
	} else {
		goto L52
	}
L51:
	;
	v274 = v265
	goto L49
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v262
	v265 = v262 + v244
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v265
	v267 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v267 <= v246 {
		v250 = v267
		v251 = v262
		goto L50
	} else {
		goto L53
	}
L53:
	;
	goto L51
L54:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v284 = F_strlen(m, v283)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v284 + v283
	v287 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v290 = v287
	goto L46
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+40)) = v302
	*(*int32)(unsafe.Add(mBase, uint32(v13)+36)) = v302
	F_infix_3(m, v13+int32(32), int32(0))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L5
	} else {
		goto L56
	}
L56:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v13)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v311
	F_infix_3(m, l0, int32(0))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L5
	} else {
		goto L57
	}
L57:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v317 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v318 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v319 = v317 - v318
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v13)+40))
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v13)+36))
	if v316 <= v319+v320-v322+int32(4) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v330 = v316
	v331 = v318
	goto L61
L59:
	;
	v358 = v317
	v359 = v322
	goto L60
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v359
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v83
	v368 = F_pg_sprintf(m, v358, int32(_a_F_infix_3_3), v13+int32(16))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L5
	} else {
		goto L65
	}
L61:
	;
	v340 = v330 << (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v340
	v342 = F_repalloc(m, v331, v340)
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L5
	} else {
		goto L63
	}
L62:
	;
	v358 = v345
	v359 = v350
	goto L60
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v342
	v345 = v342 + v319
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v345
	v347 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v13)+40))
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v13)+36))
	if v347 <= v319+int32(4)+v348-v350 {
		v330 = v347
		v331 = v342
		goto L61
	} else {
		goto L64
	}
L64:
	;
	goto L62
L65:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v371 = F_strlen(m, v370)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v371 + v370
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v13)+36))
	F_pfree(m, v374)
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L5
	} else {
		goto L66
	}
L66:
	;
	if v237&int32(1) != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	v379 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v380 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v381 = v379 - v380
	v383 = v381 + int32(3)
	v384 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v384 <= v383 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v387 = v384
	v388 = v380
	goto L71
L69:
	;
	v410 = v379
	goto L70
L70:
	;
	v418 = F_pg_sprintf(m, v410, int32(_a_F_infix_3_2), int32(0))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L5
	} else {
		goto L75
	}
L71:
	;
	v397 = v387 << (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v397
	v399 = F_repalloc(m, v388, v397)
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L5
	} else {
		goto L73
	}
L72:
	;
	v410 = v402
	goto L70
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v399
	v402 = v399 + v381
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v402
	v404 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v404 <= v383 {
		v387 = v404
		v388 = v399
		goto L71
	} else {
		goto L74
	}
L74:
	;
	goto L72
L75:
	;
	v420 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v421 = F_strlen(m, v420)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v421 + v420
	goto L1
}
func F_initialize_aggregate(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
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
	var v32 int32
	_ = v32
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
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
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
	var v75 int32
	_ = v75
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)))
	if v5 == int32(1) {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+208))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v8+v9<<(uint(int32(2))%32))))
		if v13 != 0 {
			F_tuplesort_end(m, v13)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+196))
				v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				if v17 == int32(1) {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
					v24 = *(*int32)(unsafe.Add(mBase, uint32(v16+v20<<(uint(int32(4))%32))+88))
					v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+132))
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
					v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+136))
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
					v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+140))
					v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
					v32 = *(*int32)(unsafe.Add(mBase, _c_F_initialize_aggregate[0]))
					v34 = F_tuplesort_begin_datum(m, v24, v26, v28, v30, v32, int32(0))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return
					} else {
						v47 = v34
						v48 = *(*int32)(unsafe.Add(mBase, uint32(l1)+208))
						v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
						*(*int32)(unsafe.Add(mBase, uint32(v48+v49<<(uint(int32(2))%32)))) = v47
						v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+180)))
						if v55 == int32(1) {
							v58 = *(*int32)(unsafe.Add(mBase, uint32(l1)+176))
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v58
							v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+180)))
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v75)
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v75)
							return
						} else {
							v60 = int32(_a_F_initialize_aggregate_0)
							v61 = *(*int32)(unsafe.Add(mBase, _c_F_initialize_aggregate[1]))
							v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
							v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+20))
							*(*int32)(unsafe.Add(mBase, _c_F_initialize_aggregate[1])) = v64
							v66 = *(*int32)(unsafe.Add(mBase, uint32(l1)+176))
							v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+187)))
							v68 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+184)))
							v69 = F_datumCopy(m, v66, v67, v68)
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l2))) = v69
								*(*int32)(unsafe.Add(mBase, _c_F_initialize_aggregate[1])) = v61
								v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+180)))
								*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v75)
								*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v75)
								return
							}
						}
					}
				} else {
					v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)+120))
					v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+128))
					v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)+132))
					v39 = *(*int32)(unsafe.Add(mBase, uint32(l1)+136))
					v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+140))
					v42 = *(*int32)(unsafe.Add(mBase, _c_F_initialize_aggregate[0]))
					v43 = int32(0)
					v45 = F_tuplesort_begin_heap(m, v16, v36, v37, v38, v39, v40, v42, v43, v43)
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return
					} else {
						v47 = v45
						v48 = *(*int32)(unsafe.Add(mBase, uint32(l1)+208))
						v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
						*(*int32)(unsafe.Add(mBase, uint32(v48+v49<<(uint(int32(2))%32)))) = v47
						v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+180)))
						if v55 == int32(1) {
							v58 = *(*int32)(unsafe.Add(mBase, uint32(l1)+176))
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v58
							v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+180)))
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v75)
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v75)
							return
						} else {
							v60 = int32(_a_F_initialize_aggregate_0)
							v61 = *(*int32)(unsafe.Add(mBase, _c_F_initialize_aggregate[1]))
							v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
							v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+20))
							*(*int32)(unsafe.Add(mBase, _c_F_initialize_aggregate[1])) = v64
							v66 = *(*int32)(unsafe.Add(mBase, uint32(l1)+176))
							v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+187)))
							v68 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+184)))
							v69 = F_datumCopy(m, v66, v67, v68)
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l2))) = v69
								*(*int32)(unsafe.Add(mBase, _c_F_initialize_aggregate[1])) = v61
								v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+180)))
								*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v75)
								*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v75)
								return
							}
						}
					}
				}
			}
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+196))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			if v17 == int32(1) {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v16+v20<<(uint(int32(4))%32))+88))
				v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+132))
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
				v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+136))
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
				v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+140))
				v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
				v32 = *(*int32)(unsafe.Add(mBase, _c_F_initialize_aggregate[0]))
				v34 = F_tuplesort_begin_datum(m, v24, v26, v28, v30, v32, int32(0))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return
				} else {
					v47 = v34
					v48 = *(*int32)(unsafe.Add(mBase, uint32(l1)+208))
					v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
					*(*int32)(unsafe.Add(mBase, uint32(v48+v49<<(uint(int32(2))%32)))) = v47
					v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+180)))
					if v55 == int32(1) {
						v58 = *(*int32)(unsafe.Add(mBase, uint32(l1)+176))
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = v58
						v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+180)))
						*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v75)
						*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v75)
						return
					} else {
						v60 = int32(_a_F_initialize_aggregate_0)
						v61 = *(*int32)(unsafe.Add(mBase, _c_F_initialize_aggregate[1]))
						v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
						v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+20))
						*(*int32)(unsafe.Add(mBase, _c_F_initialize_aggregate[1])) = v64
						v66 = *(*int32)(unsafe.Add(mBase, uint32(l1)+176))
						v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+187)))
						v68 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+184)))
						v69 = F_datumCopy(m, v66, v67, v68)
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v69
							*(*int32)(unsafe.Add(mBase, _c_F_initialize_aggregate[1])) = v61
							v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+180)))
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v75)
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v75)
							return
						}
					}
				}
			} else {
				v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)+120))
				v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+128))
				v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)+132))
				v39 = *(*int32)(unsafe.Add(mBase, uint32(l1)+136))
				v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+140))
				v42 = *(*int32)(unsafe.Add(mBase, _c_F_initialize_aggregate[0]))
				v43 = int32(0)
				v45 = F_tuplesort_begin_heap(m, v16, v36, v37, v38, v39, v40, v42, v43, v43)
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return
				} else {
					v47 = v45
					v48 = *(*int32)(unsafe.Add(mBase, uint32(l1)+208))
					v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
					*(*int32)(unsafe.Add(mBase, uint32(v48+v49<<(uint(int32(2))%32)))) = v47
					v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+180)))
					if v55 == int32(1) {
						v58 = *(*int32)(unsafe.Add(mBase, uint32(l1)+176))
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = v58
						v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+180)))
						*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v75)
						*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v75)
						return
					} else {
						v60 = int32(_a_F_initialize_aggregate_0)
						v61 = *(*int32)(unsafe.Add(mBase, _c_F_initialize_aggregate[1]))
						v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
						v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+20))
						*(*int32)(unsafe.Add(mBase, _c_F_initialize_aggregate[1])) = v64
						v66 = *(*int32)(unsafe.Add(mBase, uint32(l1)+176))
						v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+187)))
						v68 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+184)))
						v69 = F_datumCopy(m, v66, v67, v68)
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v69
							*(*int32)(unsafe.Add(mBase, _c_F_initialize_aggregate[1])) = v61
							v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+180)))
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v75)
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v75)
							return
						}
					}
				}
			}
		}
	} else {
		v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+180)))
		if v55 == int32(1) {
			v58 = *(*int32)(unsafe.Add(mBase, uint32(l1)+176))
			*(*int32)(unsafe.Add(mBase, uint32(l2))) = v58
			v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+180)))
			*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v75)
			*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v75)
			return
		} else {
			v60 = int32(_a_F_initialize_aggregate_0)
			v61 = *(*int32)(unsafe.Add(mBase, _c_F_initialize_aggregate[1]))
			v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
			v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+20))
			*(*int32)(unsafe.Add(mBase, _c_F_initialize_aggregate[1])) = v64
			v66 = *(*int32)(unsafe.Add(mBase, uint32(l1)+176))
			v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+187)))
			v68 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+184)))
			v69 = F_datumCopy(m, v66, v67, v68)
			mBase = m.M
			v70 = m.ExcPending
			if v70 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l2))) = v69
				*(*int32)(unsafe.Add(mBase, _c_F_initialize_aggregate[1])) = v61
				v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+180)))
				*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v75)
				*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v75)
				return
			}
		}
	}
}
func F_initialize_mergeclause_eclasses(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
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
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
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
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	v3 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	F_op_input_types(m, v13, v10+int32(12), v10+int32(8))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
		if v20 != 0 {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
			v23 = v22
		} else {
			v23 = v3
		}
		v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
		v25 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
		v26 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
		v27 = int32(0)
		v30 = F_get_eclass_for_sort_expr(m, l0, v23, v24, v25, v26, v27, v27, int32(1))
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l1)+100)) = v30
			v33 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
			if v33 == int32(0) {
				v41 = v3
			} else {
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
				if v36 < int32(2) {
					v41 = v3
				} else {
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
					v41 = v40
				}
			}
			v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
			v43 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
			v44 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
			v45 = int32(0)
			v48 = F_get_eclass_for_sort_expr(m, l0, v41, v42, v43, v44, v45, v45, int32(1))
			mBase = m.M
			v49 = m.ExcPending
			if v49 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l1)+104)) = v48
				m.G0 = v10 + int32(16)
				return
			}
		}
	}
}
func F_inline_set_returning_function(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
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
	var v38 int32
	_ = v38
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
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
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
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
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
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
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
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v287 int32
	_ = v287
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	v3 = int32(0)
	v13 = m.G0
	v15 = v13 + int32(-64)
	m.G0 = v15
	F_check_stack_depth(m)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+72)))
	if v21 != 0 {
		v309 = v3
		goto L3
	} else {
		goto L4
	}
L3:
	;
	m.G0 = v15 - int32(-64)
	return v309
L4:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	if v22 == int32(0) {
		v309 = v3
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if v25 != int32(1) {
		v309 = v3
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	if v31 != int32(15) {
		v309 = v3
		goto L3
	} else {
		goto L7
	}
L7:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+12)))
	if v34 != int32(1) {
		v309 = v3
		goto L3
	} else {
		goto L8
	}
L8:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v30)+28))
	v40 = F_contain_volatile_functions_walker(m, v38, int32(0))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	if v40 != 0 {
		v309 = v3
		goto L3
	} else {
		goto L10
	}
L10:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v30)+28))
	if v42 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	if base.Ui32(v43-int32(22)) < base.Ui32(int32(3)) {
		v309 = v3
		goto L3
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v54 = *(*int32)(unsafe.Add(mBase, _c_F_inline_set_returning_function[0]))
	v56 = F_object_aclcheck(m, int32(1255), v37, v54, int64(128))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L17
	}
L14:
	;
	v50 = F_expression_tree_walker_impl(m, v42, int32(857), int32(0))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	if v50 != 0 {
		v309 = v3
		goto L3
	} else {
		goto L16
	}
L16:
	;
	goto L13
L17:
	;
	if v56 != 0 {
		v309 = v3
		goto L3
	} else {
		goto L18
	}
L18:
	;
	v59 = *(*int32)(unsafe.Add(mBase, _c_F_inline_set_returning_function[1]))
	if v59 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v60 = m.T0[v59].(func(*base.Module, int32) int32)(m, v37)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v63 = F_SearchSysCache1(m, int32(47), v37)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L26
	}
L22:
	;
	if v60 != 0 {
		v309 = v3
		goto L3
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	F_ReleaseCatCache(m, v63)
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L1
	} else {
		goto L92
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_inline_set_returning_function[2])) = v106
	F_MemoryContextDelete(m, v103)
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L1
	} else {
		goto L91
	}
L26:
	;
	if v63 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v63)+16))
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+22)))
	v67 = v65 + v66
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+76))
	if v68 != int32(14) {
		v300 = v3
		goto L24
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L1
	} else {
		goto L88
	}
L30:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+96)))
	if v71 != int32(102) {
		v300 = v3
		goto L24
	} else {
		goto L31
	}
L31:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+99)))
	if v74 != 0 {
		v300 = v3
		goto L24
	} else {
		goto L32
	}
L32:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+101)))
	if v75 == int32(118) {
		v300 = v3
		goto L24
	} else {
		goto L33
	}
L33:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v67)+108))
	if v78 == int32(2278) {
		v300 = v3
		goto L24
	} else {
		goto L34
	}
L34:
	;
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+97)))
	if v81 != 0 {
		v300 = v3
		goto L24
	} else {
		goto L35
	}
L35:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+100)))
	if v82 != int32(1) {
		v300 = v3
		goto L24
	} else {
		goto L36
	}
L36:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v30)+28))
	if v85 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
	v88 = v86
	goto L39
L38:
	;
	v88 = int32(0)
	goto L39
L39:
	;
	v89 = int32(*(*int16)(unsafe.Add(mBase, uint32(v67)+104)))
	if v88 != v89 {
		v300 = v3
		goto L24
	} else {
		goto L40
	}
L40:
	;
	v93 = F_heap_attisnull(m, v63, int32(29), int32(0))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	if v93 == int32(0) {
		v300 = v3
		goto L24
	} else {
		goto L42
	}
L42:
	;
	v98 = *(*int32)(unsafe.Add(mBase, _c_F_inline_set_returning_function[2]))
	v103 = F_AllocSetContextCreateInternal(m, v98, int32(_a_F_inline_set_returning_function_0), int32(0), int32(_a_F_inline_set_returning_function_1), int32(_a_F_inline_set_returning_function_2))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v105 = int32(_a_F_inline_set_returning_function_3)
	v106 = *(*int32)(unsafe.Add(mBase, _c_F_inline_set_returning_function[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_inline_set_returning_function[2])) = v103
	v111 = F_SysCacheGetAttrNotNull(m, int32(47), v63, int32(26))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v113 = F_text_to_cstring(m, v111)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+44)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v15)+40)) = v67 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = int32(873)
	v121 = int32(_a_F_inline_set_returning_function_4)
	v122 = *(*int32)(unsafe.Add(mBase, _c_F_inline_set_returning_function[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_inline_set_returning_function[3])) = v13 + int32(-36)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = v13 + int32(-24)
	v135 = F_SysCacheGetAttr(m, int32(47), v63, int32(28), v13+int32(-13))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+51)))
	if v137 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v196)+12))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v198)))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	if v200 != 0 {
		goto L72
	} else {
		goto L73
	}
L48:
	;
	v140 = F_text_to_cstring(m, v135)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L52
	}
L49:
	;
	goto L50
L50:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v30)+24))
	v176 = F_prepare_sql_fn_parse_info(m, v63, v30, v175)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L1
	} else {
		goto L64
	}
L51:
	;
	if v156 == int32(0) {
		goto L25
	} else {
		goto L58
	}
L52:
	;
	v142 = F_stringToNode(m, v140)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	if v144 == int32(1) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v142)+12))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	v156 = v148
	goto L51
L55:
	;
	goto L56
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v142
	v154 = F_list_make1_impl(m, int32(1), v13+int32(-52))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	v156 = v154
	goto L51
L58:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v156)+4))
	if v159 != int32(1) {
		goto L25
	} else {
		goto L59
	}
L59:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v156)+12))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v162)))
	F_AcquireRewriteLocks(m, v163, int32(1), int32(0))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	v168 = F_pg_rewrite_query(m, v163)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	if v168 == int32(0) {
		goto L25
	} else {
		goto L62
	}
L62:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v168)+4))
	if v172 == int32(1) {
		v196 = v168
		goto L47
	} else {
		goto L63
	}
L63:
	;
	goto L25
L64:
	;
	v178 = F_pg_parse_query(m, v113)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	if v178 == int32(0) {
		goto L25
	} else {
		goto L66
	}
L66:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v178)+4))
	if v182 != int32(1) {
		goto L25
	} else {
		goto L67
	}
L67:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v178)+12))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v185)))
	v189 = F_pg_analyze_and_rewrite_withcb(m, v186, v113, int32(474), v176, int32(0))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	if v189 == int32(0) {
		goto L25
	} else {
		goto L69
	}
L69:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v189)+4))
	if v193 != int32(1) {
		goto L25
	} else {
		goto L70
	}
L70:
	;
	v196 = v189
	goto L47
L71:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v199)))
	if v218 != int32(67) {
		goto L25
	} else {
		goto L77
	}
L72:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v29)+16))
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v29)+20))
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v29)+24))
	v204 = F_BuildDescFromLists(m, v200, v201, v202, v203)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L1
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	v211 = F_get_expr_result_type(m, v30, int32(0), v13+int32(-40))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L1
	} else {
		goto L76
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = v204
	v217 = int32(0)
	goto L71
L76:
	;
	v217 = base.B2i32(base.Ui32(v211-int32(4)) < base.Ui32(int32(-3)))
	goto L71
L77:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v199)+4))
	if v221 != int32(1) {
		goto L25
	} else {
		goto L78
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v196
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v196
	v229 = F_list_make1_impl(m, int32(1), v13+int32(-56))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
	v233 = int32(*(*int8)(unsafe.Add(mBase, uint32(v67)+96)))
	v235 = F_check_sql_fn_retval(m, v229, v231, v232, v233, int32(1))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	if (v235|v217)&int32(1) == int32(0) {
		goto L25
	} else {
		goto L81
	}
L81:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v196)+12))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v242)))
	v244 = int32(*(*int16)(unsafe.Add(mBase, uint32(v67)+104)))
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v30)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+56)) = v245
	*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v244
	*(*int32)(unsafe.Add(mBase, uint32(v15)+60)) = int32(1)
	v254 = F_query_tree_mutator_impl(m, v243, int32(875), v13+int32(-12), int32(0))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_inline_set_returning_function[2])) = v106
	v258 = F_copyObjectImpl(m, v254)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	F_MemoryContextDelete(m, v103)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
	*(*int32)(unsafe.Add(mBase, _c_F_inline_set_returning_function[3])) = v263
	F_ReleaseCatCache(m, v63)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	F_record_plan_function_dependency(m, l0, v37)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258)+44)))
	if v269 != int32(1) {
		v309 = v258
		goto L3
	} else {
		goto L87
	}
L87:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v273 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v272)+81)) = uint8(v273)
	v309 = v258
	goto L3
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v37
	F_errmsg_internal(m, int32(_a_F_inline_set_returning_function_5), v15)
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	F_errfinish(m, int32(_a_F_inline_set_returning_function_6), int32(_a_F_inline_set_returning_function_7), int32(_a_F_inline_set_returning_function_0))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
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
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
	*(*int32)(unsafe.Add(mBase, _c_F_inline_set_returning_function[3])) = v298
	v300 = int32(0)
	goto L24
L92:
	;
	v309 = v300
	goto L3
}
func F_innerrel_is_unique(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	v7 = int32(0)
	v9 = F_innerrel_is_unique_ext(m, l0, l1, l2, l3, l4, l5, v7, v7)
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		return v9
	}
}
func F_insert_v(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l3-int32(4))))
	v17 = F_replace_s(m, l0, l1, l2, v14, l3, v9+int32(12))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int32(0)
	} else {
		if v17 != 0 {
			v33 = int32(-1)
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			if l1 <= v21 {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v23 + v21
			} else {
			}
			v26 = int32(0)
			v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v27 < l1 {
				v33 = v26
			} else {
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v29 + v27
				v33 = v26
			}
		}
		m.G0 = v9 + int32(16)
		return v33
	}
}
func F_int24gt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v3 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+20)))
	return base.B2i32(v2 < v3)
}
func F_int28gt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	var v4 int64
	_ = v4
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(v2)))
	v4 = int64(*(*int16)(unsafe.Add(mBase, uint32(l0)+20)))
	return base.B2i32(v3 < v4)
}
func F_int28ne(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	var v4 int64
	_ = v4
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(v2)))
	v4 = int64(*(*int16)(unsafe.Add(mBase, uint32(l0)+20)))
	return base.B2i32(v3 != v4)
}
func F_int2and(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	return base.I32_extend16_s(v2 & v3)
}
func F_int2le(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+20)))
	v3 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+28)))
	return base.B2i32(v2 <= v3)
}
func F_int2vectorout(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	if v7 != int32(1) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v94 = m.ExcPending
		if v94 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(67141764))
			mBase = m.M
			v97 = m.ExcPending
			if v97 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_int2vectorout_0), int32(0))
				mBase = m.M
				v101 = m.ExcPending
				if v101 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_int2vectorout_1), int32(158), int32(_a_F_int2vectorout_2))
					mBase = m.M
					v106 = m.ExcPending
					if v106 != 0 {
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
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
		if v10 != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v94 = m.ExcPending
			if v94 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(67141764))
				mBase = m.M
				v97 = m.ExcPending
				if v97 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_int2vectorout_0), int32(0))
					mBase = m.M
					v101 = m.ExcPending
					if v101 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_int2vectorout_1), int32(158), int32(_a_F_int2vectorout_2))
						mBase = m.M
						v106 = m.ExcPending
						if v106 != 0 {
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
			v11 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
			if v11 != int32(21) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v94 = m.ExcPending
				if v94 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(67141764))
					mBase = m.M
					v97 = m.ExcPending
					if v97 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_int2vectorout_0), int32(0))
						mBase = m.M
						v101 = m.ExcPending
						if v101 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_int2vectorout_1), int32(158), int32(_a_F_int2vectorout_2))
							mBase = m.M
							v106 = m.ExcPending
							if v106 != 0 {
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
				v14 = int32(1)
				v15 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
				v20 = F_palloc(m, v15*int32(7)+v14)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					if v15 <= int32(0) {
						v83 = v20
					} else {
						v26 = int32(*(*int16)(unsafe.Add(mBase, uint32(v6)+24)))
						if int32(0) <= v26 {
							v36 = v26
							v37 = int32(0)
						} else {
							v31 = int32(45)
							*(*uint8)(unsafe.Add(mBase, uint32(v20))) = uint8(v31)
							v36 = int32(0) - v26
							v37 = int32(1)
						}
						v39 = F_pg_ultoa_n(m, v36, v20+v37)
						mBase = m.M
						v40 = v39 + v37
						v42 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v20+v40))) = uint8(v42)
						v44 = v40 + v20
						if v15 == int32(1) {
							v83 = v44
						} else {
							v49 = v44
							v51 = v14
							for {
								v54 = int32(32)
								*(*uint8)(unsafe.Add(mBase, uint32(v49))) = uint8(v54)
								v56 = int32(1)
								v59 = int32(*(*int16)(unsafe.Add(mBase, uint32(v6+int32(24)+v51<<(uint(v56)%32)))))
								v61 = v49 + v56
								if int32(0) <= v59 {
									v71 = v59
									v72 = int32(0)
								} else {
									v66 = int32(45)
									*(*uint8)(unsafe.Add(mBase, uint32(v61))) = uint8(v66)
									v71 = int32(0) - v59
									v72 = int32(1)
								}
								v74 = F_pg_ultoa_n(m, v71, v61+v72)
								mBase = m.M
								v75 = v74 + v72
								v77 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v61+v75))) = uint8(v77)
								v79 = v75 + v61
								v81 = v51 + int32(1)
								if v81 != v15 {
									v49 = v79
									v51 = v81
									continue
								} else {
									break
								}
								break
							}
							v83 = v79
						}
					}
					v88 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v83))) = uint8(v88)
					return v20
				}
			}
		}
	}
}
func F_int2vectorrecv(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int64
	_ = v10
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	v2 = int32(0)
	v4 = m.G0
	v6 = v4 - int32(48)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v10 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v6)+13)) = v10
	*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = v10
	*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v9
	*(*uint8)(unsafe.Add(mBase, uint32(v6)+44)) = uint8(v2)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+40)) = int32(-1)
	*(*uint8)(unsafe.Add(mBase, uint32(v6)+36)) = uint8(v2)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+32)) = int32(21)
	*(*uint8)(unsafe.Add(mBase, uint32(v6)+28)) = uint8(v2)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+24)) = v8
	v26 = int32(3)
	*(*uint16)(unsafe.Add(mBase, uint32(v6)+22)) = uint16(v26)
	v30 = F_array_recv(m, v6+int32(4))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		return int32(0)
	} else {
		v34 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
		if v34 != int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v47 = m.ExcPending
			if v47 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50462850))
				mBase = m.M
				v50 = m.ExcPending
				if v50 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_int2vectorrecv_0), int32(0))
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_int2vectorrecv_1), int32(292), int32(_a_F_int2vectorrecv_2))
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
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
			v37 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
			if v37 != 0 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(50462850))
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_int2vectorrecv_0), int32(0))
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_int2vectorrecv_1), int32(292), int32(_a_F_int2vectorrecv_2))
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
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
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
				if v38 != int32(21) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(50462850))
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(_a_F_int2vectorrecv_0), int32(0))
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_int2vectorrecv_1), int32(292), int32(_a_F_int2vectorrecv_2))
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
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
					v41 = *(*int32)(unsafe.Add(mBase, uint32(v30)+20))
					if v41 == int32(0) {
						m.G0 = v6 + int32(48)
						return v30
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(50462850))
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(_a_F_int2vectorrecv_0), int32(0))
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_int2vectorrecv_1), int32(292), int32(_a_F_int2vectorrecv_2))
									mBase = m.M
									v59 = m.ExcPending
									if v59 != 0 {
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
func F_int2xor(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	return base.I32_extend16_s(v2 ^ v3)
}
func F_int42div(m *base.Module, l0 int32) int32 {
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
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = int32(_a_F_int42div_0)
	v7 = v5 & v6
	if v7 != v6 {
		if v7 != 0 {
			v34 = base.I32_div_s(v4, base.I32_extend16_s(v5))
			return v34
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(33816706))
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_int42div_1), int32(0))
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_int42div_2), int32(1130), int32(_a_F_int42div_3))
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
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
		if v4 == int32(-2147483648) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50331778))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_int42div_4), int32(0))
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_int42div_2), int32(1146), int32(_a_F_int42div_3))
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
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
			return int32(0) - v4
		}
	}
}
func F_int48le(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	var v4 int64
	_ = v4
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(v2)))
	v4 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+20)))
	return base.B2i32(v4 <= v3)
}
func F_int4in(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v4 = F_pg_strtoint32_safe(m, v2, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_int4not(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	return v2 ^ int32(-1)
}
func F_int4random(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int64
	_ = v63
	var v66 int64
	_ = v66
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int64
	_ = v83
	var v84 int64
	_ = v84
	var v94 int64
	_ = v94
	var v97 int64
	_ = v97
	var v101 int64
	_ = v101
	var v102 int64
	_ = v102
	var v105 int64
	_ = v105
	var v106 int64
	_ = v106
	var v107 int64
	_ = v107
	var v110 int64
	_ = v110
	var v111 int64
	_ = v111
	var v112 int64
	_ = v112
	var v117 int64
	_ = v117
	var v122 int64
	_ = v122
	var v127 int64
	_ = v127
	var v140 int32
	_ = v140
	var v142 int64
	_ = v142
	var v143 int64
	_ = v143
	var v150 int64
	_ = v150
	var v153 int64
	_ = v153
	var v155 int64
	_ = v155
	var v157 int64
	_ = v157
	var v159 int64
	_ = v159
	var v163 int64
	_ = v163
	var v165 int64
	_ = v165
	var v173 int64
	_ = v173
	var v178 int64
	_ = v178
	var v192 int64
	_ = v192
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v3 <= v4 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v7 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_int4random[0])))
	if v7 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L41
	} else {
		goto L42
	}
L4:
	;
	v11 = int32(16)
	v12 = int32(0)
	v16 = m.G0
	v18 = v16 - v11
	m.G0 = v18
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v12
	v24 = F_open(m, int32(_a_F_int4random_0), v12, v18)
	mBase = m.M
	if v24 != int32(-1) {
		goto L9
	} else {
		goto L10
	}
L5:
	;
	goto L6
L6:
	;
	v142 = base.I64_extend_i32_s(v3)
	v143 = base.I64_extend_i32_s(v4)
	if v142 < v143 {
		goto L35
	} else {
		goto L36
	}
L7:
	;
	v140 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_int4random[0])) = uint8(v140)
	goto L6
L8:
	;
	if v57 != 0 {
		goto L21
	} else {
		goto L22
	}
L9:
	;
	goto L13
L10:
	;
	v57 = v12
	goto L11
L11:
	;
	m.G0 = v18 + int32(16)
	goto L8
L12:
	;
	v52 = F_close(m, v24)
	mBase = m.M
	v57 = v50
	goto L11
L13:
	;
	v30 = int32(_a_F_int4random_1)
	v31 = v11
	goto L14
L14:
	;
	v36 = F_read(m, v24, v30, v31)
	mBase = m.M
	if v36 <= int32(0) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v50 = int32(1)
	goto L12
L16:
	;
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_int4random[1]))
	if v40 == int32(27) {
		goto L14
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v45 = v31 - v36
	if v45 != 0 {
		v30 = v30 + v36
		v31 = v45
		goto L14
	} else {
		goto L20
	}
L19:
	;
	v50 = int32(0)
	goto L12
L20:
	;
	goto L15
L21:
	;
	v62 = int32(_a_F_int4random_1)
	v63 = *(*int64)(unsafe.Add(mBase, _c_F_int4random[2]))
	if v63 != int64(0) {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	goto L23
L23:
	;
	v74 = int32(_a_F_int4random_1)
	v78 = m.G0
	v79 = int32(16)
	v80 = v78 - v79
	m.G0 = v80
	F___gettimeofday(m, v80)
	mBase = m.M
	v83 = *(*int64)(unsafe.Add(mBase, uint32(v80)))
	v84 = int64(*(*int32)(unsafe.Add(mBase, uint32(v80)+8)))
	m.G0 = v80 + v79
	goto L29
L24:
	;
	goto L7
L25:
	;
	goto L24
L26:
	;
	v66 = *(*int64)(unsafe.Add(mBase, _c_F_int4random[3]))
	if v66 != int64(0) {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_int4random[3])) = int64(1442695040888963407)
	*(*int64)(unsafe.Add(mBase, _c_F_int4random[2])) = int64(6364136223846793005)
	goto L25
L29:
	;
	v94 = int64(*(*uint32)(unsafe.Add(mBase, _c_F_int4random[4])))
	v97 = v84 + v83*int64(1000000) - int64(946684800000000) ^ v94<<(uint(int64(32))%64)
	v101 = v97 + int64(4354685564936845354)
	v102 = int64(30)
	v105 = int64(-4658895280553007687)
	v106 = (int64(base.Ui64(v101)>>(uint(v102)%64)) ^ v101) * v105
	v107 = int64(27)
	v110 = int64(-7723592293110705685)
	v111 = (int64(base.Ui64(v106)>>(uint(v107)%64)) ^ v106) * v110
	v112 = int64(31)
	*(*int64)(unsafe.Add(mBase, _c_F_int4random[3])) = int64(base.Ui64(v111)>>(uint(v112)%64)) ^ v111
	v117 = v97 - int64(7046029254386353131)
	v122 = (int64(base.Ui64(v117)>>(uint(v102)%64)) ^ v117) * v105
	v127 = (int64(base.Ui64(v122)>>(uint(v107)%64)) ^ v122) * v110
	*(*int64)(unsafe.Add(mBase, _c_F_int4random[2])) = int64(base.Ui64(v127)>>(uint(v112)%64)) ^ v127
	if v117|v101 == int64(0) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	goto L7
L31:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_int4random[3])) = int64(1442695040888963407)
	*(*int64)(unsafe.Add(mBase, _c_F_int4random[2])) = int64(6364136223846793005)
	goto L33
L32:
	;
	goto L33
L33:
	;
	goto L30
L34:
	;
	return base.I32_wrap_i64(v192)
L35:
	;
	v150 = v143 - v142
	v153 = *(*int64)(unsafe.Add(mBase, _c_F_int4random[3]))
	v155 = *(*int64)(unsafe.Add(mBase, _c_F_int4random[2]))
	v157 = v155
	v159 = v153
	goto L38
L36:
	;
	v192 = v142
	goto L37
L37:
	;
	goto L34
L38:
	;
	v163 = v157 ^ v159
	v165 = base.I64_rotl(v163, int64(37))
	v173 = v163 ^ (v163<<(uint(int64(16))%64) ^ base.I64_rotl(v157, int64(24)))
	v178 = int64(base.Ui64(base.I64_rotl(v157*int64(5), int64(7))*int64(9)) >> (uint(base.I64_clz(v150)) % 64))
	if base.Ui64(v150) < base.Ui64(v178) {
		v157 = v173
		v159 = v165
		goto L38
	} else {
		goto L40
	}
L39:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_int4random[3])) = v165
	*(*int64)(unsafe.Add(mBase, _c_F_int4random[2])) = v173
	v192 = v142 + v178
	goto L37
L40:
	;
	goto L39
L41:
	;
	return int32(0)
L42:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L41
	} else {
		goto L43
	}
L43:
	;
	F_errmsg(m, int32(_a_F_int4random_2), int32(0))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L41
	} else {
		goto L44
	}
L44:
	;
	F_errfinish(m, int32(_a_F_int4random_3), int32(135), int32(_a_F_int4random_4))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L41
	} else {
		goto L45
	}
L45:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_int4shl(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	return v2 << (uint(v3) % 32)
}
func F_int82gt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	var v4 int64
	_ = v4
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(v2)))
	v4 = int64(*(*int16)(unsafe.Add(mBase, uint32(l0)+28)))
	return base.B2i32(v4 < v3)
}
func F_int84div(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int64
	_ = v4
	var v5 int32
	_ = v5
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int64
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int64)(unsafe.Add(mBase, uint32(v3)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	switch v5 + int32(1) {
	case 0:
		if v4 == int64(-9223372036854775807-1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50331778))
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_int84div_0), int32(0))
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_int84div_1), int32(958), int32(_a_F_int84div_2))
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
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
			v30 = F_Int64GetDatum(m, int64(0)-v4)
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int32(0)
			} else {
				return v30
			}
		}
	case 1:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(33816706))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_int84div_3), int32(0))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_int84div_1), int32(942), int32(_a_F_int84div_2))
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	default:
		v34 = base.I64_div_s(v4, base.I64_extend_i32_s(v5))
		v35 = F_Int64GetDatum(m, v34)
		mBase = m.M
		v36 = m.ExcPending
		if v36 != 0 {
			return int32(0)
		} else {
			return v35
		}
	}
}
func F_int84le(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	var v4 int64
	_ = v4
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(v2)))
	v4 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+28)))
	return base.B2i32(v3 <= v4)
}
func F_int8and(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	var v4 int32
	_ = v4
	var v5 int64
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(v2)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(v4)))
	v7 = F_Int64GetDatum(m, v3&v5)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_int8not(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(v2)))
	v6 = F_Int64GetDatum(m, v3^int64(-1))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_int8random(m *base.Module, l0 int32) int32 {
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
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v66 int64
	_ = v66
	var v69 int64
	_ = v69
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int64
	_ = v86
	var v87 int64
	_ = v87
	var v97 int64
	_ = v97
	var v100 int64
	_ = v100
	var v104 int64
	_ = v104
	var v105 int64
	_ = v105
	var v108 int64
	_ = v108
	var v109 int64
	_ = v109
	var v110 int64
	_ = v110
	var v113 int64
	_ = v113
	var v114 int64
	_ = v114
	var v115 int64
	_ = v115
	var v120 int64
	_ = v120
	var v125 int64
	_ = v125
	var v130 int64
	_ = v130
	var v143 int32
	_ = v143
	var v151 int64
	_ = v151
	var v154 int64
	_ = v154
	var v156 int64
	_ = v156
	var v158 int64
	_ = v158
	var v160 int64
	_ = v160
	var v164 int64
	_ = v164
	var v166 int64
	_ = v166
	var v174 int64
	_ = v174
	var v179 int64
	_ = v179
	var v193 int64
	_ = v193
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(v4)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
	if v5 <= v7 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v10 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_int8random[0])))
	if v10 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L41
	} else {
		goto L43
	}
L4:
	;
	v14 = int32(16)
	v15 = int32(0)
	v19 = m.G0
	v21 = v19 - v14
	m.G0 = v21
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v15
	v27 = F_open(m, int32(_a_F_int8random_0), v15, v21)
	mBase = m.M
	if v27 != int32(-1) {
		goto L9
	} else {
		goto L10
	}
L5:
	;
	goto L6
L6:
	;
	if v5 < v7 {
		goto L35
	} else {
		goto L36
	}
L7:
	;
	v143 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_int8random[0])) = uint8(v143)
	goto L6
L8:
	;
	if v60 != 0 {
		goto L21
	} else {
		goto L22
	}
L9:
	;
	goto L13
L10:
	;
	v60 = v15
	goto L11
L11:
	;
	m.G0 = v21 + int32(16)
	goto L8
L12:
	;
	v55 = F_close(m, v27)
	mBase = m.M
	v60 = v53
	goto L11
L13:
	;
	v33 = int32(_a_F_int8random_1)
	v34 = v14
	goto L14
L14:
	;
	v39 = F_read(m, v27, v33, v34)
	mBase = m.M
	if v39 <= int32(0) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v53 = int32(1)
	goto L12
L16:
	;
	v43 = *(*int32)(unsafe.Add(mBase, _c_F_int8random[1]))
	if v43 == int32(27) {
		goto L14
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v48 = v34 - v39
	if v48 != 0 {
		v33 = v33 + v39
		v34 = v48
		goto L14
	} else {
		goto L20
	}
L19:
	;
	v53 = int32(0)
	goto L12
L20:
	;
	goto L15
L21:
	;
	v65 = int32(_a_F_int8random_1)
	v66 = *(*int64)(unsafe.Add(mBase, _c_F_int8random[2]))
	if v66 != int64(0) {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	goto L23
L23:
	;
	v77 = int32(_a_F_int8random_1)
	v81 = m.G0
	v82 = int32(16)
	v83 = v81 - v82
	m.G0 = v83
	F___gettimeofday(m, v83)
	mBase = m.M
	v86 = *(*int64)(unsafe.Add(mBase, uint32(v83)))
	v87 = int64(*(*int32)(unsafe.Add(mBase, uint32(v83)+8)))
	m.G0 = v83 + v82
	goto L29
L24:
	;
	goto L7
L25:
	;
	goto L24
L26:
	;
	v69 = *(*int64)(unsafe.Add(mBase, _c_F_int8random[3]))
	if v69 != int64(0) {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_int8random[3])) = int64(1442695040888963407)
	*(*int64)(unsafe.Add(mBase, _c_F_int8random[2])) = int64(6364136223846793005)
	goto L25
L29:
	;
	v97 = int64(*(*uint32)(unsafe.Add(mBase, _c_F_int8random[4])))
	v100 = v87 + v86*int64(1000000) - int64(946684800000000) ^ v97<<(uint(int64(32))%64)
	v104 = v100 + int64(4354685564936845354)
	v105 = int64(30)
	v108 = int64(-4658895280553007687)
	v109 = (int64(base.Ui64(v104)>>(uint(v105)%64)) ^ v104) * v108
	v110 = int64(27)
	v113 = int64(-7723592293110705685)
	v114 = (int64(base.Ui64(v109)>>(uint(v110)%64)) ^ v109) * v113
	v115 = int64(31)
	*(*int64)(unsafe.Add(mBase, _c_F_int8random[3])) = int64(base.Ui64(v114)>>(uint(v115)%64)) ^ v114
	v120 = v100 - int64(7046029254386353131)
	v125 = (int64(base.Ui64(v120)>>(uint(v105)%64)) ^ v120) * v108
	v130 = (int64(base.Ui64(v125)>>(uint(v110)%64)) ^ v125) * v113
	*(*int64)(unsafe.Add(mBase, _c_F_int8random[2])) = int64(base.Ui64(v130)>>(uint(v115)%64)) ^ v130
	if v120|v104 == int64(0) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	goto L7
L31:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_int8random[3])) = int64(1442695040888963407)
	*(*int64)(unsafe.Add(mBase, _c_F_int8random[2])) = int64(6364136223846793005)
	goto L33
L32:
	;
	goto L33
L33:
	;
	goto L30
L34:
	;
	v194 = F_Int64GetDatum(m, v193)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L41
	} else {
		goto L42
	}
L35:
	;
	v151 = v7 - v5
	v154 = *(*int64)(unsafe.Add(mBase, _c_F_int8random[3]))
	v156 = *(*int64)(unsafe.Add(mBase, _c_F_int8random[2]))
	v158 = v156
	v160 = v154
	goto L38
L36:
	;
	v193 = v5
	goto L37
L37:
	;
	goto L34
L38:
	;
	v164 = v158 ^ v160
	v166 = base.I64_rotl(v164, int64(37))
	v174 = v164 ^ (v164<<(uint(int64(16))%64) ^ base.I64_rotl(v158, int64(24)))
	v179 = int64(base.Ui64(base.I64_rotl(v158*int64(5), int64(7))*int64(9)) >> (uint(base.I64_clz(v151)) % 64))
	if base.Ui64(v151) < base.Ui64(v179) {
		v158 = v174
		v160 = v166
		goto L38
	} else {
		goto L40
	}
L39:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_int8random[3])) = v166
	*(*int64)(unsafe.Add(mBase, _c_F_int8random[2])) = v174
	v193 = v5 + v179
	goto L37
L40:
	;
	goto L39
L41:
	;
	return int32(0)
L42:
	;
	return v194
L43:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L41
	} else {
		goto L44
	}
L44:
	;
	F_errmsg(m, int32(_a_F_int8random_2), int32(0))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L41
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(_a_F_int8random_3), int32(159), int32(_a_F_int8random_4))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L41
	} else {
		goto L46
	}
L46:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_int8xor(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	var v4 int32
	_ = v4
	var v5 int64
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(v2)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(v4)))
	v7 = F_Int64GetDatum(m, v3^v5)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_intarray_concat_arrays(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
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
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v136 int32
	_ = v136
	var v143 int32
	_ = v143
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v11 = F_ArrayGetNItems(m, v8, l0+int32(16))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		v18 = F_ArrayGetNItems(m, v15, l1+int32(16))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			if v20 != 0 {
				v21 = F_array_contains_nulls(m, l0)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					if v21 != 0 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v107 = m.ExcPending
						if v107 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(67108994))
							mBase = m.M
							v110 = m.ExcPending
							if v110 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(_a_F_intarray_concat_arrays_0), int32(0))
								mBase = m.M
								v116 = m.ExcPending
								if v116 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_intarray_concat_arrays_1), int32(377), int32(_a_F_intarray_concat_arrays_2))
									mBase = m.M
									v123 = m.ExcPending
									if v123 != 0 {
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
						v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
						if v23 != 0 {
							v24 = F_array_contains_nulls(m, l1)
							mBase = m.M
							v25 = m.ExcPending
							if v25 != 0 {
								return int32(0)
							} else {
								if v24 != 0 {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v127 = m.ExcPending
									if v127 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(67108994))
										mBase = m.M
										v130 = m.ExcPending
										if v130 != 0 {
											return int32(0)
										} else {
											F_errmsg(m, int32(_a_F_intarray_concat_arrays_0), int32(0))
											mBase = m.M
											v136 = m.ExcPending
											if v136 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_intarray_concat_arrays_1), int32(378), int32(_a_F_intarray_concat_arrays_2))
												mBase = m.M
												v143 = m.ExcPending
												if v143 != 0 {
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
									v26 = v11 + v18
									if v26 <= int32(0) {
										v30 = F_construct_empty_array(m, int32(23))
										mBase = m.M
										v31 = m.ExcPending
										if v31 != 0 {
											return int32(0)
										} else {
											v48 = v30
											if v11 != 0 {
												v50 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
												if v50 != 0 {
													v58 = v50
												} else {
													v51 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
													v58 = (v51<<(uint(int32(3))%32) + int32(23)) & int32(-8)
												}
												v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												if v60 != 0 {
													v68 = v60
												} else {
													v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
													v68 = (v61<<(uint(int32(3))%32) + int32(23)) & int32(-8)
												}
												v71 = v11 << (uint(int32(2)) % 32)
												if v71 != 0 {
													v72 = F__emscripten_memcpy_bulkmem(m, v58+v48, v68+l0, v71)
													mBase = m.M
												} else {
												}
											} else {
											}
											if v18 != 0 {
												v75 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
												if v75 != 0 {
													v83 = v75
												} else {
													v76 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
													v83 = (v76<<(uint(int32(3))%32) + int32(23)) & int32(-8)
												}
												v88 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
												if v88 != 0 {
													v96 = v88
												} else {
													v89 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
													v96 = (v89<<(uint(int32(3))%32) + int32(23)) & int32(-8)
												}
												v99 = v18 << (uint(int32(2)) % 32)
												if v99 != 0 {
													v100 = F__emscripten_memcpy_bulkmem(m, v83+v48+v11<<(uint(int32(2))%32), v96+l1, v99)
													mBase = m.M
												} else {
												}
											} else {
											}
											return v48
										}
									} else {
										v35 = v26<<(uint(int32(2))%32) + int32(24)
										v36 = F_palloc0(m, v35)
										mBase = m.M
										v37 = m.ExcPending
										if v37 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v36)+20)) = int32(1)
											*(*int32)(unsafe.Add(mBase, uint32(v36)+16)) = v26
											*(*int32)(unsafe.Add(mBase, uint32(v36)+12)) = int32(23)
											*(*int64)(unsafe.Add(mBase, uint32(v36)+4)) = int64(1)
											*(*int32)(unsafe.Add(mBase, uint32(v36))) = v35 << (uint(int32(2)) % 32)
											v48 = v36
											if v11 != 0 {
												v50 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
												if v50 != 0 {
													v58 = v50
												} else {
													v51 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
													v58 = (v51<<(uint(int32(3))%32) + int32(23)) & int32(-8)
												}
												v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												if v60 != 0 {
													v68 = v60
												} else {
													v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
													v68 = (v61<<(uint(int32(3))%32) + int32(23)) & int32(-8)
												}
												v71 = v11 << (uint(int32(2)) % 32)
												if v71 != 0 {
													v72 = F__emscripten_memcpy_bulkmem(m, v58+v48, v68+l0, v71)
													mBase = m.M
												} else {
												}
											} else {
											}
											if v18 != 0 {
												v75 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
												if v75 != 0 {
													v83 = v75
												} else {
													v76 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
													v83 = (v76<<(uint(int32(3))%32) + int32(23)) & int32(-8)
												}
												v88 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
												if v88 != 0 {
													v96 = v88
												} else {
													v89 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
													v96 = (v89<<(uint(int32(3))%32) + int32(23)) & int32(-8)
												}
												v99 = v18 << (uint(int32(2)) % 32)
												if v99 != 0 {
													v100 = F__emscripten_memcpy_bulkmem(m, v83+v48+v11<<(uint(int32(2))%32), v96+l1, v99)
													mBase = m.M
												} else {
												}
											} else {
											}
											return v48
										}
									}
								}
							}
						} else {
							v26 = v11 + v18
							if v26 <= int32(0) {
								v30 = F_construct_empty_array(m, int32(23))
								mBase = m.M
								v31 = m.ExcPending
								if v31 != 0 {
									return int32(0)
								} else {
									v48 = v30
									if v11 != 0 {
										v50 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
										if v50 != 0 {
											v58 = v50
										} else {
											v51 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
											v58 = (v51<<(uint(int32(3))%32) + int32(23)) & int32(-8)
										}
										v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										if v60 != 0 {
											v68 = v60
										} else {
											v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
											v68 = (v61<<(uint(int32(3))%32) + int32(23)) & int32(-8)
										}
										v71 = v11 << (uint(int32(2)) % 32)
										if v71 != 0 {
											v72 = F__emscripten_memcpy_bulkmem(m, v58+v48, v68+l0, v71)
											mBase = m.M
										} else {
										}
									} else {
									}
									if v18 != 0 {
										v75 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
										if v75 != 0 {
											v83 = v75
										} else {
											v76 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
											v83 = (v76<<(uint(int32(3))%32) + int32(23)) & int32(-8)
										}
										v88 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
										if v88 != 0 {
											v96 = v88
										} else {
											v89 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
											v96 = (v89<<(uint(int32(3))%32) + int32(23)) & int32(-8)
										}
										v99 = v18 << (uint(int32(2)) % 32)
										if v99 != 0 {
											v100 = F__emscripten_memcpy_bulkmem(m, v83+v48+v11<<(uint(int32(2))%32), v96+l1, v99)
											mBase = m.M
										} else {
										}
									} else {
									}
									return v48
								}
							} else {
								v35 = v26<<(uint(int32(2))%32) + int32(24)
								v36 = F_palloc0(m, v35)
								mBase = m.M
								v37 = m.ExcPending
								if v37 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v36)+20)) = int32(1)
									*(*int32)(unsafe.Add(mBase, uint32(v36)+16)) = v26
									*(*int32)(unsafe.Add(mBase, uint32(v36)+12)) = int32(23)
									*(*int64)(unsafe.Add(mBase, uint32(v36)+4)) = int64(1)
									*(*int32)(unsafe.Add(mBase, uint32(v36))) = v35 << (uint(int32(2)) % 32)
									v48 = v36
									if v11 != 0 {
										v50 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
										if v50 != 0 {
											v58 = v50
										} else {
											v51 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
											v58 = (v51<<(uint(int32(3))%32) + int32(23)) & int32(-8)
										}
										v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										if v60 != 0 {
											v68 = v60
										} else {
											v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
											v68 = (v61<<(uint(int32(3))%32) + int32(23)) & int32(-8)
										}
										v71 = v11 << (uint(int32(2)) % 32)
										if v71 != 0 {
											v72 = F__emscripten_memcpy_bulkmem(m, v58+v48, v68+l0, v71)
											mBase = m.M
										} else {
										}
									} else {
									}
									if v18 != 0 {
										v75 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
										if v75 != 0 {
											v83 = v75
										} else {
											v76 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
											v83 = (v76<<(uint(int32(3))%32) + int32(23)) & int32(-8)
										}
										v88 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
										if v88 != 0 {
											v96 = v88
										} else {
											v89 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
											v96 = (v89<<(uint(int32(3))%32) + int32(23)) & int32(-8)
										}
										v99 = v18 << (uint(int32(2)) % 32)
										if v99 != 0 {
											v100 = F__emscripten_memcpy_bulkmem(m, v83+v48+v11<<(uint(int32(2))%32), v96+l1, v99)
											mBase = m.M
										} else {
										}
									} else {
									}
									return v48
								}
							}
						}
					}
				}
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				if v23 != 0 {
					v24 = F_array_contains_nulls(m, l1)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						if v24 != 0 {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v127 = m.ExcPending
							if v127 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(67108994))
								mBase = m.M
								v130 = m.ExcPending
								if v130 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(_a_F_intarray_concat_arrays_0), int32(0))
									mBase = m.M
									v136 = m.ExcPending
									if v136 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_intarray_concat_arrays_1), int32(378), int32(_a_F_intarray_concat_arrays_2))
										mBase = m.M
										v143 = m.ExcPending
										if v143 != 0 {
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
							v26 = v11 + v18
							if v26 <= int32(0) {
								v30 = F_construct_empty_array(m, int32(23))
								mBase = m.M
								v31 = m.ExcPending
								if v31 != 0 {
									return int32(0)
								} else {
									v48 = v30
									if v11 != 0 {
										v50 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
										if v50 != 0 {
											v58 = v50
										} else {
											v51 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
											v58 = (v51<<(uint(int32(3))%32) + int32(23)) & int32(-8)
										}
										v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										if v60 != 0 {
											v68 = v60
										} else {
											v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
											v68 = (v61<<(uint(int32(3))%32) + int32(23)) & int32(-8)
										}
										v71 = v11 << (uint(int32(2)) % 32)
										if v71 != 0 {
											v72 = F__emscripten_memcpy_bulkmem(m, v58+v48, v68+l0, v71)
											mBase = m.M
										} else {
										}
									} else {
									}
									if v18 != 0 {
										v75 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
										if v75 != 0 {
											v83 = v75
										} else {
											v76 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
											v83 = (v76<<(uint(int32(3))%32) + int32(23)) & int32(-8)
										}
										v88 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
										if v88 != 0 {
											v96 = v88
										} else {
											v89 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
											v96 = (v89<<(uint(int32(3))%32) + int32(23)) & int32(-8)
										}
										v99 = v18 << (uint(int32(2)) % 32)
										if v99 != 0 {
											v100 = F__emscripten_memcpy_bulkmem(m, v83+v48+v11<<(uint(int32(2))%32), v96+l1, v99)
											mBase = m.M
										} else {
										}
									} else {
									}
									return v48
								}
							} else {
								v35 = v26<<(uint(int32(2))%32) + int32(24)
								v36 = F_palloc0(m, v35)
								mBase = m.M
								v37 = m.ExcPending
								if v37 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v36)+20)) = int32(1)
									*(*int32)(unsafe.Add(mBase, uint32(v36)+16)) = v26
									*(*int32)(unsafe.Add(mBase, uint32(v36)+12)) = int32(23)
									*(*int64)(unsafe.Add(mBase, uint32(v36)+4)) = int64(1)
									*(*int32)(unsafe.Add(mBase, uint32(v36))) = v35 << (uint(int32(2)) % 32)
									v48 = v36
									if v11 != 0 {
										v50 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
										if v50 != 0 {
											v58 = v50
										} else {
											v51 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
											v58 = (v51<<(uint(int32(3))%32) + int32(23)) & int32(-8)
										}
										v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										if v60 != 0 {
											v68 = v60
										} else {
											v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
											v68 = (v61<<(uint(int32(3))%32) + int32(23)) & int32(-8)
										}
										v71 = v11 << (uint(int32(2)) % 32)
										if v71 != 0 {
											v72 = F__emscripten_memcpy_bulkmem(m, v58+v48, v68+l0, v71)
											mBase = m.M
										} else {
										}
									} else {
									}
									if v18 != 0 {
										v75 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
										if v75 != 0 {
											v83 = v75
										} else {
											v76 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
											v83 = (v76<<(uint(int32(3))%32) + int32(23)) & int32(-8)
										}
										v88 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
										if v88 != 0 {
											v96 = v88
										} else {
											v89 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
											v96 = (v89<<(uint(int32(3))%32) + int32(23)) & int32(-8)
										}
										v99 = v18 << (uint(int32(2)) % 32)
										if v99 != 0 {
											v100 = F__emscripten_memcpy_bulkmem(m, v83+v48+v11<<(uint(int32(2))%32), v96+l1, v99)
											mBase = m.M
										} else {
										}
									} else {
									}
									return v48
								}
							}
						}
					}
				} else {
					v26 = v11 + v18
					if v26 <= int32(0) {
						v30 = F_construct_empty_array(m, int32(23))
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return int32(0)
						} else {
							v48 = v30
							if v11 != 0 {
								v50 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
								if v50 != 0 {
									v58 = v50
								} else {
									v51 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
									v58 = (v51<<(uint(int32(3))%32) + int32(23)) & int32(-8)
								}
								v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								if v60 != 0 {
									v68 = v60
								} else {
									v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									v68 = (v61<<(uint(int32(3))%32) + int32(23)) & int32(-8)
								}
								v71 = v11 << (uint(int32(2)) % 32)
								if v71 != 0 {
									v72 = F__emscripten_memcpy_bulkmem(m, v58+v48, v68+l0, v71)
									mBase = m.M
								} else {
								}
							} else {
							}
							if v18 != 0 {
								v75 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
								if v75 != 0 {
									v83 = v75
								} else {
									v76 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
									v83 = (v76<<(uint(int32(3))%32) + int32(23)) & int32(-8)
								}
								v88 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
								if v88 != 0 {
									v96 = v88
								} else {
									v89 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
									v96 = (v89<<(uint(int32(3))%32) + int32(23)) & int32(-8)
								}
								v99 = v18 << (uint(int32(2)) % 32)
								if v99 != 0 {
									v100 = F__emscripten_memcpy_bulkmem(m, v83+v48+v11<<(uint(int32(2))%32), v96+l1, v99)
									mBase = m.M
								} else {
								}
							} else {
							}
							return v48
						}
					} else {
						v35 = v26<<(uint(int32(2))%32) + int32(24)
						v36 = F_palloc0(m, v35)
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v36)+20)) = int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(v36)+16)) = v26
							*(*int32)(unsafe.Add(mBase, uint32(v36)+12)) = int32(23)
							*(*int64)(unsafe.Add(mBase, uint32(v36)+4)) = int64(1)
							*(*int32)(unsafe.Add(mBase, uint32(v36))) = v35 << (uint(int32(2)) % 32)
							v48 = v36
							if v11 != 0 {
								v50 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
								if v50 != 0 {
									v58 = v50
								} else {
									v51 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
									v58 = (v51<<(uint(int32(3))%32) + int32(23)) & int32(-8)
								}
								v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								if v60 != 0 {
									v68 = v60
								} else {
									v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									v68 = (v61<<(uint(int32(3))%32) + int32(23)) & int32(-8)
								}
								v71 = v11 << (uint(int32(2)) % 32)
								if v71 != 0 {
									v72 = F__emscripten_memcpy_bulkmem(m, v58+v48, v68+l0, v71)
									mBase = m.M
								} else {
								}
							} else {
							}
							if v18 != 0 {
								v75 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
								if v75 != 0 {
									v83 = v75
								} else {
									v76 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
									v83 = (v76<<(uint(int32(3))%32) + int32(23)) & int32(-8)
								}
								v88 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
								if v88 != 0 {
									v96 = v88
								} else {
									v89 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
									v96 = (v89<<(uint(int32(3))%32) + int32(23)) & int32(-8)
								}
								v99 = v18 << (uint(int32(2)) % 32)
								if v99 != 0 {
									v100 = F__emscripten_memcpy_bulkmem(m, v83+v48+v11<<(uint(int32(2))%32), v96+l1, v99)
									mBase = m.M
								} else {
								}
							} else {
							}
							return v48
						}
					}
				}
			}
		}
	}
}
func F_interpret_func_parallel(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)+4))
	v4 = int32(_a_F_interpret_func_parallel_0)
	v7 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_interpret_func_parallel[0])))
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3))))
	if v8 == int32(0) {
		v27 = v7
		v28 = v8
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v28-v27 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L2:
	;
	goto L1
L3:
	;
	if v7 != v8 {
		v27 = v7
		v28 = v8
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v12 = v3
	v13 = v4
	goto L5
L5:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+1)))
	if v17 == int32(0) {
		v27 = v16
		v28 = v17
		goto L2
	} else {
		goto L7
	}
L6:
	;
	v27 = v16
	v28 = v17
	goto L2
L7:
	;
	v20 = int32(1)
	if v16 == v17 {
		v12 = v12 + v20
		v13 = v13 + v20
		goto L5
	} else {
		goto L8
	}
L8:
	;
	goto L6
L9:
	;
	return int32(115)
L10:
	;
	goto L11
L11:
	;
	v34 = int32(_a_F_interpret_func_parallel_1)
	v37 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_interpret_func_parallel[1])))
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3))))
	if v38 == int32(0) {
		v57 = v37
		v58 = v38
		goto L13
	} else {
		goto L14
	}
L12:
	;
	if v58-v57 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L13:
	;
	goto L12
L14:
	;
	if v37 != v38 {
		v57 = v37
		v58 = v38
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v42 = v3
	v43 = v34
	goto L16
L16:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+1)))
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+1)))
	if v47 == int32(0) {
		v57 = v46
		v58 = v47
		goto L13
	} else {
		goto L18
	}
L17:
	;
	v57 = v46
	v58 = v47
	goto L13
L18:
	;
	v50 = int32(1)
	if v46 == v47 {
		v42 = v42 + v50
		v43 = v43 + v50
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	return int32(117)
L21:
	;
	goto L22
L22:
	;
	v64 = int32(_a_F_interpret_func_parallel_2)
	v67 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_interpret_func_parallel[2])))
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3))))
	if v68 == int32(0) {
		v87 = v67
		v88 = v68
		goto L24
	} else {
		goto L25
	}
L23:
	;
	if v88-v87 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L24:
	;
	goto L23
L25:
	;
	if v67 != v68 {
		v87 = v67
		v88 = v68
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v72 = v3
	v73 = v64
	goto L27
L27:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+1)))
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+1)))
	if v77 == int32(0) {
		v87 = v76
		v88 = v77
		goto L24
	} else {
		goto L29
	}
L28:
	;
	v87 = v76
	v88 = v77
	goto L24
L29:
	;
	v80 = int32(1)
	if v76 == v77 {
		v72 = v72 + v80
		v73 = v73 + v80
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	return int32(114)
L32:
	;
	goto L33
L33:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	return int32(0)
L35:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L34
	} else {
		goto L36
	}
L36:
	;
	F_errmsg(m, int32(_a_F_interpret_func_parallel_3), int32(0))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L34
	} else {
		goto L37
	}
L37:
	;
	F_errfinish(m, int32(_a_F_interpret_func_parallel_4), int32(649), int32(_a_F_interpret_func_parallel_5))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L34
	} else {
		goto L38
	}
L38:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_intset_add_member(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v22 int32
	_ = v22
	var v25 int64
	_ = v25
	var v27 int64
	_ = v27
	var v30 int32
	_ = v30
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
	var v42 int64
	_ = v42
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int64
	_ = v89
	var v90 int64
	_ = v90
	var v93 int64
	_ = v93
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v110 int64
	_ = v110
	var v111 int64
	_ = v111
	var v119 int32
	_ = v119
	var __phi119 int32
	_ = __phi119
	var v132 int64
	_ = v132
	var __phi132 int64
	_ = __phi132
	var v133 int64
	_ = v133
	var __phi133 int64
	_ = __phi133
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v160 int64
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v172 int64
	_ = v172
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int64
	_ = v181
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v204 int64
	_ = v204
	var v211 int32
	_ = v211
	var v213 int64
	_ = v213
	var v215 int64
	_ = v215
	var v216 int64
	_ = v216
	var v220 int64
	_ = v220
	var v227 int64
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v248 int64
	_ = v248
	var v257 int32
	_ = v257
	var v259 int64
	_ = v259
	var v261 int64
	_ = v261
	var v282 int64
	_ = v282
	var v310 int64
	_ = v310
	var v334 int64
	_ = v334
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int64
	_ = v344
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v374 int32
	_ = v374
	var v379 int32
	_ = v379
	var v385 int32
	_ = v385
	var v386 int64
	_ = v386
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int64
	_ = v393
	var v400 int32
	_ = v400
	var v407 int32
	_ = v407
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v420 int64
	_ = v420
	var v427 int32
	_ = v427
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v475 int32
	_ = v475
	var v480 int32
	_ = v480
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v492 int32
	_ = v492
	var v496 int32
	_ = v496
	var v503 int32
	_ = v503
	var v506 int32
	_ = v506
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v526 int32
	_ = v526
	var v533 int32
	_ = v533
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v542 int32
	_ = v542
	var v548 int32
	_ = v548
	var v555 int32
	_ = v555
	var v559 int32
	_ = v559
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
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v596 int32
	_ = v596
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v601 int32
	_ = v601
	var v603 int32
	_ = v603
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v618 int32
	_ = v618
	var v620 int32
	_ = v620
	var v625 int32
	_ = v625
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v643 int32
	_ = v643
	var v667 int32
	_ = v667
	var v671 int64
	_ = v671
	var v678 int32
	_ = v678
	var v682 int32
	_ = v682
	var v687 int32
	_ = v687
	var v691 int32
	_ = v691
	var v695 int32
	_ = v695
	var v700 int32
	_ = v700
	var v704 int32
	_ = v704
	var v708 int32
	_ = v708
	var v713 int32
	_ = v713
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+3948)))
	if v22 != int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v25 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	if base.Ui64(l1) <= base.Ui64(v25) {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	goto L3
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v704 = m.ExcPending
	if v704 != 0 {
		goto L16
	} else {
		goto L121
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v691 = m.ExcPending
	if v691 != 0 {
		goto L16
	} else {
		goto L118
	}
L5:
	;
	v27 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	if v27 != int64(0) {
		goto L4
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+3944))
	if int32(482) <= v30 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L7
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v678 = m.ExcPending
	if v678 != 0 {
		goto L16
	} else {
		goto L115
	}
L10:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v33 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v643 = v30
	goto L12
L12:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0+v643<<(uint(int32(3))%32))+88)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = l1
	v667 = *(*int32)(unsafe.Add(mBase, uint32(l0)+3944))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+3944)) = v667 + int32(1)
	v671 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v671 + int64(1)
	return
L13:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v38 = F_MemoryContextAlloc(m, v36, int32(1032))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v54 = v33
	goto L15
L15:
	;
	v56 = l0 + int32(40)
	v58 = l0 + int32(88)
	v65 = v54
	v69 = int32(0)
	goto L19
L16:
	;
	return
L17:
	;
	v40 = F_GetMemoryChunkSpace(m, v38)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v42 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v42 + base.I64_extend_i32_u(v40)
	*(*int64)(unsafe.Add(mBase, uint32(v38))) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = int32(1)
	v54 = v38
	goto L15
L19:
	;
	v82 = int32(0)
	v86 = v58 + v69<<(uint(int32(3))%32)
	v88 = v86 + int32(8)
	v89 = *(*int64)(unsafe.Add(mBase, uint32(v88)))
	v90 = *(*int64)(unsafe.Add(mBase, uint32(v86)))
	v93 = v89 + (v90 ^ int64(-1))
	v97 = v82
	v99 = v82
	v102 = v82
	v103 = int32(240)
	v110 = v93
	v111 = v89
	goto L21
L20:
	;
	v485 = *(*int32)(unsafe.Add(mBase, uint32(l0)+3944))
	if v480 < v485 {
		goto L66
	} else {
		goto L67
	}
L21:
	;
	__phi119 = v97
	__phi132 = v110
	__phi133 = v111
	v119 = __phi119
	v132 = __phi132
	v133 = __phi133
	goto L23
L22:
	;
	if v165 != 0 {
		goto L31
	} else {
		goto L32
	}
L23:
	;
	if int64(base.Ui64(v132)>>(uint(base.I64_extend_i32_u(v102))%64)) != int64(0) {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	goto L22
L25:
	;
	goto L24
L26:
	;
	v141 = int32(1)
	v142 = v99 + v141
	v144 = v142 << (uint(v141) % 32)
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144)+uint32(_c_F_intset_add_member[0]))))
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144)+uint32(_c_F_intset_add_member[1]))))
	if v119 < v150 {
		v97 = v119
		v99 = v142
		v102 = v147
		v103 = v150
		v110 = v132
		v111 = v133
		goto L21
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v153 = v119 + int32(1)
	if v103 <= v153 {
		v163 = v99
		v164 = v102
		v165 = v103
		goto L25
	} else {
		goto L30
	}
L29:
	;
	v163 = v142
	v164 = v147
	v165 = v150
	goto L25
L30:
	;
	v160 = *(*int64)(unsafe.Add(mBase, uint32(v88+v153<<(uint(int32(3))%32))))
	__phi119 = v153
	__phi132 = v133 ^ int64(-1) + v160
	__phi133 = v160
	v119 = __phi119
	v132 = __phi132
	v133 = __phi133
	goto L23
L31:
	;
	if v164 <= int32(0) {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	v334 = int64(1152921504606846975)
	goto L33
L33:
	;
	v335 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v65)+2)))
	if base.Ui32(int32(64)) <= base.Ui32(v335) {
		goto L46
	} else {
		goto L47
	}
L34:
	;
	v310 = int64(0)
	goto L36
L35:
	;
	v172 = int64(0)
	if v165 < int32(2) {
		v282 = v172
		goto L37
	} else {
		goto L38
	}
L36:
	;
	v334 = base.I64_extend_i32_u(v163)<<(uint(int64(60))%64) | v310
	goto L33
L37:
	;
	v310 = v282 | v93
	goto L36
L38:
	;
	v176 = v86 - int32(8)
	v177 = int32(1)
	v178 = v165 - v177
	v181 = base.I64_extend_i32_u(v164)
	if v165 != int32(2) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v191 = v165
	v196 = int32(0)
	v204 = v172
	goto L42
L40:
	;
	v235 = v165
	v248 = v172
	goto L41
L41:
	;
	if v178&v177 == int32(0) {
		v282 = v248
		goto L37
	} else {
		goto L45
	}
L42:
	;
	v211 = v191 << (uint(int32(3)) % 32)
	v213 = *(*int64)(unsafe.Add(mBase, uint32(v176+v211)))
	v215 = *(*int64)(unsafe.Add(mBase, uint32(v211+(v176-int32(8)))))
	v216 = int64(-1)
	v220 = *(*int64)(unsafe.Add(mBase, uint32(v211+v86)))
	v227 = (v213 + (v215 ^ v216) | (v220+(v213^v216)|v204)<<(uint(v181)%64)) << (uint(v181) % 64)
	v228 = int32(2)
	v229 = v191 - v228
	v231 = v196 + v228
	if v231 != v178&int32(-2) {
		v191 = v229
		v196 = v231
		v204 = v227
		goto L42
	} else {
		goto L44
	}
L43:
	;
	v235 = v229
	v248 = v227
	goto L41
L44:
	;
	goto L43
L45:
	;
	v257 = v235 << (uint(int32(3)) % 32)
	v259 = *(*int64)(unsafe.Add(mBase, uint32(v86+v257)))
	v261 = *(*int64)(unsafe.Add(mBase, uint32(v257+v176)))
	v282 = (v259 + (v261 ^ int64(-1)) | v248) << (uint(v181) % 64)
	goto L37
L46:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v340 = F_MemoryContextAlloc(m, v338, int32(1032))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L16
	} else {
		goto L49
	}
L47:
	;
	v449 = v335
	v452 = v65
	goto L48
L48:
	;
	v468 = int32(1)
	v469 = v449 + v468
	*(*uint16)(unsafe.Add(mBase, uint32(v452)+2)) = uint16(v469)
	v475 = v452 + v449&int32(_a_F_intset_add_member_0)<<(uint(int32(4))%32)
	*(*int64)(unsafe.Add(mBase, uint32(v475)+16)) = v334
	*(*int64)(unsafe.Add(mBase, uint32(v475)+8)) = v90
	v480 = v165 + v69 + v468
	if base.Ui64(int64(240)) < base.Ui64(base.I64_extend_i32_u(v30)-base.I64_extend_i32_s(v480)) {
		v65 = v452
		v69 = v480
		goto L19
	} else {
		goto L65
	}
L49:
	;
	v342 = F_GetMemoryChunkSpace(m, v340)
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L16
	} else {
		goto L50
	}
L50:
	;
	v344 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v344 + base.I64_extend_i32_u(v342)
	*(*int64)(unsafe.Add(mBase, uint32(v340))) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v65)+4)) = v340
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v340
	v356 = int32(1)
	v360 = v340
	goto L51
L51:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v356 < v374 {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v407+v411<<(uint(int32(3))%32))+8)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v407+v411<<(uint(int32(2))%32))+520)) = v360
	v444 = v411 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v407)+2)) = uint16(v444)
	v446 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v340)+2)))
	v449 = v446
	v452 = v340
	goto L48
L53:
	;
	v411 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v407)+2)))
	if base.Ui32(int32(64)) <= base.Ui32(v411) {
		goto L60
	} else {
		goto L61
	}
L54:
	;
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v56+v356<<(uint(int32(2))%32))))
	v407 = v379
	goto L53
L55:
	;
	goto L56
L56:
	;
	if v374 == int32(11) {
		goto L9
	} else {
		goto L57
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v374 + int32(1)
	v385 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v386 = *(*int64)(unsafe.Add(mBase, uint32(v385)+8))
	v387 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v389 = F_MemoryContextAlloc(m, v387, int32(776))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L16
	} else {
		goto L58
	}
L58:
	;
	v391 = F_GetMemoryChunkSpace(m, v389)
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L16
	} else {
		goto L59
	}
L59:
	;
	v393 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v393 + base.I64_extend_i32_u(v391)
	*(*int32)(unsafe.Add(mBase, uint32(v389)+520)) = v385
	*(*int64)(unsafe.Add(mBase, uint32(v389)+8)) = v386
	*(*uint16)(unsafe.Add(mBase, uint32(v389))) = uint16(v356)
	v400 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v389)+2)) = uint16(v400)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v56+v356<<(uint(int32(2))%32)))) = v389
	v407 = v389
	goto L53
L60:
	;
	v414 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v416 = F_MemoryContextAlloc(m, v414, int32(776))
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L16
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	goto L52
L63:
	;
	v418 = F_GetMemoryChunkSpace(m, v416)
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L16
	} else {
		goto L64
	}
L64:
	;
	v420 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v420 + base.I64_extend_i32_u(v418)
	*(*int32)(unsafe.Add(mBase, uint32(v416)+520)) = v360
	*(*int64)(unsafe.Add(mBase, uint32(v416)+8)) = v90
	*(*uint16)(unsafe.Add(mBase, uint32(v416))) = uint16(v356)
	v427 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v416)+2)) = uint16(v427)
	*(*int32)(unsafe.Add(mBase, uint32(v56+v356<<(uint(int32(2))%32)))) = v416
	v356 = v356 + v427
	v360 = v416
	goto L51
L65:
	;
	goto L20
L66:
	;
	v487 = int32(3)
	v489 = v58 + v480<<(uint(v487)%32)
	v492 = (v485 - v480) << (uint(v487) % 32)
	if v58 == v489 {
		goto L70
	} else {
		goto L71
	}
L67:
	;
	v638 = v485
	goto L68
L68:
	;
	v639 = v638 - v480
	*(*int32)(unsafe.Add(mBase, uint32(l0)+3944)) = v639
	v643 = v639
	goto L12
L69:
	;
	v637 = *(*int32)(unsafe.Add(mBase, uint32(l0)+3944))
	v638 = v637
	goto L68
L70:
	;
	goto L69
L71:
	;
	v496 = v58 + v492
	if base.Ui32(v489-v496) <= base.Ui32(int32(0)-v492<<(uint(int32(1))%32)) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v503 = F___memcpy(m, v58, v489, v492)
	mBase = m.M
	goto L69
L73:
	;
	goto L74
L74:
	;
	v506 = (v58 ^ v489) & int32(3)
	if base.Ui32(v58) < base.Ui32(v489) {
		goto L77
	} else {
		goto L78
	}
L75:
	;
	if v608 == int32(0) {
		goto L70
	} else {
		goto L111
	}
L76:
	;
	if base.Ui32(v586) <= base.Ui32(int32(3)) {
		v607 = v585
		v608 = v586
		v609 = v587
		goto L75
	} else {
		goto L107
	}
L77:
	;
	if v506 != 0 {
		goto L80
	} else {
		goto L81
	}
L78:
	;
	goto L79
L79:
	;
	if v506 != 0 {
		v568 = v492
		goto L90
	} else {
		goto L91
	}
L80:
	;
	v607 = v489
	v608 = v492
	v609 = v58
	goto L75
L81:
	;
	goto L82
L82:
	;
	if v58&int32(3) == int32(0) {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v585 = v489
	v586 = v492
	v587 = v58
	goto L76
L84:
	;
	goto L85
L85:
	;
	v513 = v489
	v514 = v492
	v515 = v58
	goto L86
L86:
	;
	if v514 == int32(0) {
		goto L70
	} else {
		goto L88
	}
L87:
	;
	v585 = v522
	v586 = v524
	v587 = v526
	goto L76
L88:
	;
	v519 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v513))))
	*(*uint8)(unsafe.Add(mBase, uint32(v515))) = uint8(v519)
	v521 = int32(1)
	v522 = v513 + v521
	v524 = v514 - v521
	v526 = v515 + v521
	if v526&int32(3) != 0 {
		v513 = v522
		v514 = v524
		v515 = v526
		goto L86
	} else {
		goto L89
	}
L89:
	;
	goto L87
L90:
	;
	if v568 == int32(0) {
		goto L70
	} else {
		goto L103
	}
L91:
	;
	if v496&int32(3) != 0 {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v533 = v492
	goto L95
L93:
	;
	v548 = v492
	goto L94
L94:
	;
	if base.Ui32(v548) <= base.Ui32(int32(3)) {
		v568 = v548
		goto L90
	} else {
		goto L99
	}
L95:
	;
	if v533 == int32(0) {
		goto L70
	} else {
		goto L97
	}
L96:
	;
	v548 = v539
	goto L94
L97:
	;
	v539 = v533 - int32(1)
	v540 = v58 + v539
	v542 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489+v539))))
	*(*uint8)(unsafe.Add(mBase, uint32(v540))) = uint8(v542)
	if v540&int32(3) != 0 {
		v533 = v539
		goto L95
	} else {
		goto L98
	}
L98:
	;
	goto L96
L99:
	;
	v555 = v548
	goto L100
L100:
	;
	v559 = v555 - int32(4)
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v489+v559)))
	*(*int32)(unsafe.Add(mBase, uint32(v58+v559))) = v562
	if base.Ui32(int32(3)) < base.Ui32(v559) {
		v555 = v559
		goto L100
	} else {
		goto L102
	}
L101:
	;
	v568 = v559
	goto L90
L102:
	;
	goto L101
L103:
	;
	v575 = v568
	goto L104
L104:
	;
	v579 = v575 - int32(1)
	v582 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489+v579))))
	*(*uint8)(unsafe.Add(mBase, uint32(v58+v579))) = uint8(v582)
	if v579 != 0 {
		v575 = v579
		goto L104
	} else {
		goto L106
	}
L105:
	;
	goto L70
L106:
	;
	goto L105
L107:
	;
	v592 = v585
	v593 = v586
	v594 = v587
	goto L108
L108:
	;
	v596 = *(*int32)(unsafe.Add(mBase, uint32(v592)))
	*(*int32)(unsafe.Add(mBase, uint32(v594))) = v596
	v598 = int32(4)
	v599 = v592 + v598
	v601 = v594 + v598
	v603 = v593 - v598
	if base.Ui32(int32(3)) < base.Ui32(v603) {
		v592 = v599
		v593 = v603
		v594 = v601
		goto L108
	} else {
		goto L110
	}
L109:
	;
	v607 = v599
	v608 = v603
	v609 = v601
	goto L75
L110:
	;
	goto L109
L111:
	;
	v614 = v607
	v615 = v608
	v616 = v609
	goto L112
L112:
	;
	v618 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v614))))
	*(*uint8)(unsafe.Add(mBase, uint32(v616))) = uint8(v618)
	v620 = int32(1)
	v625 = v615 - v620
	if v625 != 0 {
		v614 = v614 + v620
		v615 = v625
		v616 = v616 + v620
		goto L112
	} else {
		goto L114
	}
L113:
	;
	goto L70
L114:
	;
	goto L113
L115:
	;
	F_errmsg_internal(m, int32(_a_F_intset_add_member_1), int32(0))
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
		goto L16
	} else {
		goto L116
	}
L116:
	;
	F_errfinish(m, int32(_a_F_intset_add_member_2), int32(497), int32(_a_F_intset_add_member_3))
	mBase = m.M
	v687 = m.ExcPending
	if v687 != 0 {
		goto L16
	} else {
		goto L117
	}
L117:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L118:
	;
	F_errmsg_internal(m, int32(_a_F_intset_add_member_4), int32(0))
	mBase = m.M
	v695 = m.ExcPending
	if v695 != 0 {
		goto L16
	} else {
		goto L119
	}
L119:
	;
	F_errfinish(m, int32(_a_F_intset_add_member_2), int32(375), int32(_a_F_intset_add_member_5))
	mBase = m.M
	v700 = m.ExcPending
	if v700 != 0 {
		goto L16
	} else {
		goto L120
	}
L120:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L121:
	;
	F_errmsg_internal(m, int32(_a_F_intset_add_member_6), int32(0))
	mBase = m.M
	v708 = m.ExcPending
	if v708 != 0 {
		goto L16
	} else {
		goto L122
	}
L122:
	;
	F_errfinish(m, int32(_a_F_intset_add_member_2), int32(372), int32(_a_F_intset_add_member_5))
	mBase = m.M
	v713 = m.ExcPending
	if v713 != 0 {
		goto L16
	} else {
		goto L123
	}
L123:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_intset_create(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v24 int32
	_ = v24
	var v31 int64
	_ = v31
	v4 = F_palloc(m, int32(_a_F_intset_create_0))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, _c_F_intset_create[0]))
		*(*int32)(unsafe.Add(mBase, uint32(v4))) = v9
		v11 = F_GetMemoryChunkSpace(m, v4)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			v13 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v4)+3948)) = uint8(v13)
			*(*int32)(unsafe.Add(mBase, uint32(v4)+3944)) = v13
			*(*int64)(unsafe.Add(mBase, uint32(v4)+8)) = base.I64_extend_i32_u(v11)
			v24 = F__emscripten_memset_bulkmem(m, v4+int32(16), base.I32_extend8_s(v13), int32(72))
			mBase = m.M
			*(*int32)(unsafe.Add(mBase, uint32(v4+int32(3968)))) = int32(0)
			v31 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v4+int32(3960)))) = v31
			*(*int64)(unsafe.Add(mBase, uint32(v4)+3952)) = v31
			return v4
		}
	}
}
func F_intset_subtract(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
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
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v212 int32
	_ = v212
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v232 int32
	_ = v232
	var v239 int32
	_ = v239
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v20 = F_pg_detoast_datum_copy(m, v19)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		return int32(0)
	} else {
		v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v25 = F_pg_detoast_datum_copy(m, v24)
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return int32(0)
		} else {
			v27 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
			if v27 != 0 {
				v28 = F_array_contains_nulls(m, v20)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					if v28 != 0 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v203 = m.ExcPending
						if v203 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(67108994))
							mBase = m.M
							v206 = m.ExcPending
							if v206 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(_a_F_intset_subtract_0), int32(0))
								mBase = m.M
								v212 = m.ExcPending
								if v212 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_intset_subtract_1), int32(406), int32(_a_F_intset_subtract_2))
									mBase = m.M
									v219 = m.ExcPending
									if v219 != 0 {
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
						v30 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
						if v30 != 0 {
							v31 = F_array_contains_nulls(m, v25)
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
								return int32(0)
							} else {
								if v31 != 0 {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v223 = m.ExcPending
									if v223 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(67108994))
										mBase = m.M
										v226 = m.ExcPending
										if v226 != 0 {
											return int32(0)
										} else {
											F_errmsg(m, int32(_a_F_intset_subtract_0), int32(0))
											mBase = m.M
											v232 = m.ExcPending
											if v232 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_intset_subtract_1), int32(407), int32(_a_F_intset_subtract_2))
												mBase = m.M
												v239 = m.ExcPending
												if v239 != 0 {
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
									v33 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
									v36 = F_ArrayGetNItems(m, v33, v20+int32(16))
									mBase = m.M
									v37 = m.ExcPending
									if v37 != 0 {
										return int32(0)
									} else {
										v38 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v17)+15)) = uint8(v38)
										v40 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
										if v40 != 0 {
											v48 = v40
										} else {
											v41 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
											v48 = (v41<<(uint(int32(3))%32) + int32(23)) & int32(-8)
										}
										F_isort(m, v48+v20, v36, v17+int32(15))
										mBase = m.M
										v53 = F__int_unique(m, v20)
										mBase = m.M
										v54 = m.ExcPending
										if v54 != 0 {
											return int32(0)
										} else {
											v55 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
											v58 = F_ArrayGetNItems(m, v55, v53+int32(16))
											mBase = m.M
											v59 = m.ExcPending
											if v59 != 0 {
												return int32(0)
											} else {
												v60 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
												v63 = F_ArrayGetNItems(m, v60, v25+int32(16))
												mBase = m.M
												v64 = m.ExcPending
												if v64 != 0 {
													return int32(0)
												} else {
													v65 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v17)+14)) = uint8(v65)
													v67 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
													if v67 != 0 {
														v75 = v67
													} else {
														v68 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
														v75 = (v68<<(uint(int32(3))%32) + int32(23)) & int32(-8)
													}
													F_isort(m, v75+v25, v63, v17+int32(14))
													mBase = m.M
													v80 = F__int_unique(m, v25)
													mBase = m.M
													v81 = m.ExcPending
													if v81 != 0 {
														return int32(0)
													} else {
														v82 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
														v85 = F_ArrayGetNItems(m, v82, v80+int32(16))
														mBase = m.M
														v86 = m.ExcPending
														if v86 != 0 {
															return int32(0)
														} else {
															v87 = F_new_intArrayType(m, v58)
															mBase = m.M
															v88 = m.ExcPending
															if v88 != 0 {
																return int32(0)
															} else {
																v89 = *(*int32)(unsafe.Add(mBase, uint32(v53)+8))
																if v89 == int32(0) {
																	v92 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
																	v99 = (v92<<(uint(int32(3))%32) + int32(23)) & int32(-8)
																} else {
																	v99 = v89
																}
																v100 = *(*int32)(unsafe.Add(mBase, uint32(v80)+8))
																if v100 == int32(0) {
																	v103 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
																	v110 = (v103<<(uint(int32(3))%32) + int32(23)) & int32(-8)
																} else {
																	v110 = v100
																}
																v111 = *(*int32)(unsafe.Add(mBase, uint32(v87)+8))
																if v111 == int32(0) {
																	v114 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
																	v121 = (v114<<(uint(int32(3))%32) + int32(23)) & int32(-8)
																} else {
																	v121 = v111
																}
																v122 = int32(0)
																if v58 <= v122 {
																	v178 = int32(0)
																} else {
																	v129 = int32(0)
																	v132 = v129
																	v133 = v129
																	v134 = v122
																	for {
																		v148 = *(*int32)(unsafe.Add(mBase, uint32(v99+v53+v132<<(uint(int32(2))%32))))
																		if v134 != v85 {
																			v153 = *(*int32)(unsafe.Add(mBase, uint32(v110+v80+v134<<(uint(int32(2))%32))))
																			if v153 <= v148 {
																				if v148 == v153 {
																					v165 = int32(1)
																					v171 = v132 + v165
																					v172 = v133
																					v173 = v134 + v165
																				} else {
																					v171 = v132
																					v172 = v133
																					v173 = v134 + int32(1)
																				}
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(v121+v87+v133<<(uint(int32(2))%32)))) = v148
																				v160 = int32(1)
																				v171 = v132 + v160
																				v172 = v133 + v160
																				v173 = v134
																			}
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v121+v87+v133<<(uint(int32(2))%32)))) = v148
																			v160 = int32(1)
																			v171 = v132 + v160
																			v172 = v133 + v160
																			v173 = v134
																		}
																		if v171 < v58 {
																			v132 = v171
																			v133 = v172
																			v134 = v173
																			continue
																		} else {
																			break
																		}
																		break
																	}
																	v178 = v172
																}
																v190 = F_resize_intArrayType(m, v87, v178)
																mBase = m.M
																v191 = m.ExcPending
																if v191 != 0 {
																	return int32(0)
																} else {
																	F_pfree(m, v53)
																	mBase = m.M
																	v193 = m.ExcPending
																	if v193 != 0 {
																		return int32(0)
																	} else {
																		F_pfree(m, v80)
																		mBase = m.M
																		v195 = m.ExcPending
																		if v195 != 0 {
																			return int32(0)
																		} else {
																			m.G0 = v17 + int32(16)
																			return v190
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
							v33 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
							v36 = F_ArrayGetNItems(m, v33, v20+int32(16))
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return int32(0)
							} else {
								v38 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v17)+15)) = uint8(v38)
								v40 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
								if v40 != 0 {
									v48 = v40
								} else {
									v41 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
									v48 = (v41<<(uint(int32(3))%32) + int32(23)) & int32(-8)
								}
								F_isort(m, v48+v20, v36, v17+int32(15))
								mBase = m.M
								v53 = F__int_unique(m, v20)
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return int32(0)
								} else {
									v55 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
									v58 = F_ArrayGetNItems(m, v55, v53+int32(16))
									mBase = m.M
									v59 = m.ExcPending
									if v59 != 0 {
										return int32(0)
									} else {
										v60 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
										v63 = F_ArrayGetNItems(m, v60, v25+int32(16))
										mBase = m.M
										v64 = m.ExcPending
										if v64 != 0 {
											return int32(0)
										} else {
											v65 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v17)+14)) = uint8(v65)
											v67 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
											if v67 != 0 {
												v75 = v67
											} else {
												v68 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
												v75 = (v68<<(uint(int32(3))%32) + int32(23)) & int32(-8)
											}
											F_isort(m, v75+v25, v63, v17+int32(14))
											mBase = m.M
											v80 = F__int_unique(m, v25)
											mBase = m.M
											v81 = m.ExcPending
											if v81 != 0 {
												return int32(0)
											} else {
												v82 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
												v85 = F_ArrayGetNItems(m, v82, v80+int32(16))
												mBase = m.M
												v86 = m.ExcPending
												if v86 != 0 {
													return int32(0)
												} else {
													v87 = F_new_intArrayType(m, v58)
													mBase = m.M
													v88 = m.ExcPending
													if v88 != 0 {
														return int32(0)
													} else {
														v89 = *(*int32)(unsafe.Add(mBase, uint32(v53)+8))
														if v89 == int32(0) {
															v92 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
															v99 = (v92<<(uint(int32(3))%32) + int32(23)) & int32(-8)
														} else {
															v99 = v89
														}
														v100 = *(*int32)(unsafe.Add(mBase, uint32(v80)+8))
														if v100 == int32(0) {
															v103 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
															v110 = (v103<<(uint(int32(3))%32) + int32(23)) & int32(-8)
														} else {
															v110 = v100
														}
														v111 = *(*int32)(unsafe.Add(mBase, uint32(v87)+8))
														if v111 == int32(0) {
															v114 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
															v121 = (v114<<(uint(int32(3))%32) + int32(23)) & int32(-8)
														} else {
															v121 = v111
														}
														v122 = int32(0)
														if v58 <= v122 {
															v178 = int32(0)
														} else {
															v129 = int32(0)
															v132 = v129
															v133 = v129
															v134 = v122
															for {
																v148 = *(*int32)(unsafe.Add(mBase, uint32(v99+v53+v132<<(uint(int32(2))%32))))
																if v134 != v85 {
																	v153 = *(*int32)(unsafe.Add(mBase, uint32(v110+v80+v134<<(uint(int32(2))%32))))
																	if v153 <= v148 {
																		if v148 == v153 {
																			v165 = int32(1)
																			v171 = v132 + v165
																			v172 = v133
																			v173 = v134 + v165
																		} else {
																			v171 = v132
																			v172 = v133
																			v173 = v134 + int32(1)
																		}
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v121+v87+v133<<(uint(int32(2))%32)))) = v148
																		v160 = int32(1)
																		v171 = v132 + v160
																		v172 = v133 + v160
																		v173 = v134
																	}
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v121+v87+v133<<(uint(int32(2))%32)))) = v148
																	v160 = int32(1)
																	v171 = v132 + v160
																	v172 = v133 + v160
																	v173 = v134
																}
																if v171 < v58 {
																	v132 = v171
																	v133 = v172
																	v134 = v173
																	continue
																} else {
																	break
																}
																break
															}
															v178 = v172
														}
														v190 = F_resize_intArrayType(m, v87, v178)
														mBase = m.M
														v191 = m.ExcPending
														if v191 != 0 {
															return int32(0)
														} else {
															F_pfree(m, v53)
															mBase = m.M
															v193 = m.ExcPending
															if v193 != 0 {
																return int32(0)
															} else {
																F_pfree(m, v80)
																mBase = m.M
																v195 = m.ExcPending
																if v195 != 0 {
																	return int32(0)
																} else {
																	m.G0 = v17 + int32(16)
																	return v190
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
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
				if v30 != 0 {
					v31 = F_array_contains_nulls(m, v25)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						if v31 != 0 {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v223 = m.ExcPending
							if v223 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(67108994))
								mBase = m.M
								v226 = m.ExcPending
								if v226 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(_a_F_intset_subtract_0), int32(0))
									mBase = m.M
									v232 = m.ExcPending
									if v232 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_intset_subtract_1), int32(407), int32(_a_F_intset_subtract_2))
										mBase = m.M
										v239 = m.ExcPending
										if v239 != 0 {
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
							v33 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
							v36 = F_ArrayGetNItems(m, v33, v20+int32(16))
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return int32(0)
							} else {
								v38 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v17)+15)) = uint8(v38)
								v40 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
								if v40 != 0 {
									v48 = v40
								} else {
									v41 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
									v48 = (v41<<(uint(int32(3))%32) + int32(23)) & int32(-8)
								}
								F_isort(m, v48+v20, v36, v17+int32(15))
								mBase = m.M
								v53 = F__int_unique(m, v20)
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return int32(0)
								} else {
									v55 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
									v58 = F_ArrayGetNItems(m, v55, v53+int32(16))
									mBase = m.M
									v59 = m.ExcPending
									if v59 != 0 {
										return int32(0)
									} else {
										v60 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
										v63 = F_ArrayGetNItems(m, v60, v25+int32(16))
										mBase = m.M
										v64 = m.ExcPending
										if v64 != 0 {
											return int32(0)
										} else {
											v65 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v17)+14)) = uint8(v65)
											v67 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
											if v67 != 0 {
												v75 = v67
											} else {
												v68 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
												v75 = (v68<<(uint(int32(3))%32) + int32(23)) & int32(-8)
											}
											F_isort(m, v75+v25, v63, v17+int32(14))
											mBase = m.M
											v80 = F__int_unique(m, v25)
											mBase = m.M
											v81 = m.ExcPending
											if v81 != 0 {
												return int32(0)
											} else {
												v82 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
												v85 = F_ArrayGetNItems(m, v82, v80+int32(16))
												mBase = m.M
												v86 = m.ExcPending
												if v86 != 0 {
													return int32(0)
												} else {
													v87 = F_new_intArrayType(m, v58)
													mBase = m.M
													v88 = m.ExcPending
													if v88 != 0 {
														return int32(0)
													} else {
														v89 = *(*int32)(unsafe.Add(mBase, uint32(v53)+8))
														if v89 == int32(0) {
															v92 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
															v99 = (v92<<(uint(int32(3))%32) + int32(23)) & int32(-8)
														} else {
															v99 = v89
														}
														v100 = *(*int32)(unsafe.Add(mBase, uint32(v80)+8))
														if v100 == int32(0) {
															v103 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
															v110 = (v103<<(uint(int32(3))%32) + int32(23)) & int32(-8)
														} else {
															v110 = v100
														}
														v111 = *(*int32)(unsafe.Add(mBase, uint32(v87)+8))
														if v111 == int32(0) {
															v114 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
															v121 = (v114<<(uint(int32(3))%32) + int32(23)) & int32(-8)
														} else {
															v121 = v111
														}
														v122 = int32(0)
														if v58 <= v122 {
															v178 = int32(0)
														} else {
															v129 = int32(0)
															v132 = v129
															v133 = v129
															v134 = v122
															for {
																v148 = *(*int32)(unsafe.Add(mBase, uint32(v99+v53+v132<<(uint(int32(2))%32))))
																if v134 != v85 {
																	v153 = *(*int32)(unsafe.Add(mBase, uint32(v110+v80+v134<<(uint(int32(2))%32))))
																	if v153 <= v148 {
																		if v148 == v153 {
																			v165 = int32(1)
																			v171 = v132 + v165
																			v172 = v133
																			v173 = v134 + v165
																		} else {
																			v171 = v132
																			v172 = v133
																			v173 = v134 + int32(1)
																		}
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v121+v87+v133<<(uint(int32(2))%32)))) = v148
																		v160 = int32(1)
																		v171 = v132 + v160
																		v172 = v133 + v160
																		v173 = v134
																	}
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v121+v87+v133<<(uint(int32(2))%32)))) = v148
																	v160 = int32(1)
																	v171 = v132 + v160
																	v172 = v133 + v160
																	v173 = v134
																}
																if v171 < v58 {
																	v132 = v171
																	v133 = v172
																	v134 = v173
																	continue
																} else {
																	break
																}
																break
															}
															v178 = v172
														}
														v190 = F_resize_intArrayType(m, v87, v178)
														mBase = m.M
														v191 = m.ExcPending
														if v191 != 0 {
															return int32(0)
														} else {
															F_pfree(m, v53)
															mBase = m.M
															v193 = m.ExcPending
															if v193 != 0 {
																return int32(0)
															} else {
																F_pfree(m, v80)
																mBase = m.M
																v195 = m.ExcPending
																if v195 != 0 {
																	return int32(0)
																} else {
																	m.G0 = v17 + int32(16)
																	return v190
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
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
					v36 = F_ArrayGetNItems(m, v33, v20+int32(16))
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return int32(0)
					} else {
						v38 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v17)+15)) = uint8(v38)
						v40 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
						if v40 != 0 {
							v48 = v40
						} else {
							v41 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
							v48 = (v41<<(uint(int32(3))%32) + int32(23)) & int32(-8)
						}
						F_isort(m, v48+v20, v36, v17+int32(15))
						mBase = m.M
						v53 = F__int_unique(m, v20)
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return int32(0)
						} else {
							v55 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
							v58 = F_ArrayGetNItems(m, v55, v53+int32(16))
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
								return int32(0)
							} else {
								v60 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
								v63 = F_ArrayGetNItems(m, v60, v25+int32(16))
								mBase = m.M
								v64 = m.ExcPending
								if v64 != 0 {
									return int32(0)
								} else {
									v65 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v17)+14)) = uint8(v65)
									v67 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
									if v67 != 0 {
										v75 = v67
									} else {
										v68 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
										v75 = (v68<<(uint(int32(3))%32) + int32(23)) & int32(-8)
									}
									F_isort(m, v75+v25, v63, v17+int32(14))
									mBase = m.M
									v80 = F__int_unique(m, v25)
									mBase = m.M
									v81 = m.ExcPending
									if v81 != 0 {
										return int32(0)
									} else {
										v82 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
										v85 = F_ArrayGetNItems(m, v82, v80+int32(16))
										mBase = m.M
										v86 = m.ExcPending
										if v86 != 0 {
											return int32(0)
										} else {
											v87 = F_new_intArrayType(m, v58)
											mBase = m.M
											v88 = m.ExcPending
											if v88 != 0 {
												return int32(0)
											} else {
												v89 = *(*int32)(unsafe.Add(mBase, uint32(v53)+8))
												if v89 == int32(0) {
													v92 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
													v99 = (v92<<(uint(int32(3))%32) + int32(23)) & int32(-8)
												} else {
													v99 = v89
												}
												v100 = *(*int32)(unsafe.Add(mBase, uint32(v80)+8))
												if v100 == int32(0) {
													v103 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
													v110 = (v103<<(uint(int32(3))%32) + int32(23)) & int32(-8)
												} else {
													v110 = v100
												}
												v111 = *(*int32)(unsafe.Add(mBase, uint32(v87)+8))
												if v111 == int32(0) {
													v114 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
													v121 = (v114<<(uint(int32(3))%32) + int32(23)) & int32(-8)
												} else {
													v121 = v111
												}
												v122 = int32(0)
												if v58 <= v122 {
													v178 = int32(0)
												} else {
													v129 = int32(0)
													v132 = v129
													v133 = v129
													v134 = v122
													for {
														v148 = *(*int32)(unsafe.Add(mBase, uint32(v99+v53+v132<<(uint(int32(2))%32))))
														if v134 != v85 {
															v153 = *(*int32)(unsafe.Add(mBase, uint32(v110+v80+v134<<(uint(int32(2))%32))))
															if v153 <= v148 {
																if v148 == v153 {
																	v165 = int32(1)
																	v171 = v132 + v165
																	v172 = v133
																	v173 = v134 + v165
																} else {
																	v171 = v132
																	v172 = v133
																	v173 = v134 + int32(1)
																}
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v121+v87+v133<<(uint(int32(2))%32)))) = v148
																v160 = int32(1)
																v171 = v132 + v160
																v172 = v133 + v160
																v173 = v134
															}
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v121+v87+v133<<(uint(int32(2))%32)))) = v148
															v160 = int32(1)
															v171 = v132 + v160
															v172 = v133 + v160
															v173 = v134
														}
														if v171 < v58 {
															v132 = v171
															v133 = v172
															v134 = v173
															continue
														} else {
															break
														}
														break
													}
													v178 = v172
												}
												v190 = F_resize_intArrayType(m, v87, v178)
												mBase = m.M
												v191 = m.ExcPending
												if v191 != 0 {
													return int32(0)
												} else {
													F_pfree(m, v53)
													mBase = m.M
													v193 = m.ExcPending
													if v193 != 0 {
														return int32(0)
													} else {
														F_pfree(m, v80)
														mBase = m.M
														v195 = m.ExcPending
														if v195 != 0 {
															return int32(0)
														} else {
															m.G0 = v17 + int32(16)
															return v190
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
func F_inv_create(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v13 = F_table_open(m, int32(2995), int32(3))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v17 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v9)+2)) = uint8(v17)
		*(*uint16)(unsafe.Add(mBase, uint32(v9))) = uint16(v17)
		*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v17
		if l0 == v17 {
			v27 = F_GetNewOidWithIndex(m, v13, int32(2996), int32(1))
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				v29 = v27
				v32 = *(*int32)(unsafe.Add(mBase, _c_F_inv_create[0]))
				v34 = F_get_user_default_acl(m, int32(22), v32, int32(0))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v32
					*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v29
					if v34 != 0 {
						*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v34
					} else {
						v39 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v9)+2)) = uint8(v39)
					}
					v41 = *(*int32)(unsafe.Add(mBase, uint32(v13)+52))
					v44 = F_heap_form_tuple(m, v41, v9+int32(4), v9)
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return int32(0)
					} else {
						F_CatalogTupleInsert(m, v13, v44)
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return int32(0)
						} else {
							F_pfree(m, v44)
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return int32(0)
							} else {
								F_sequence_close(m, v13, int32(3))
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
									return int32(0)
								} else {
									F_recordDependencyOnNewAcl(m, int32(2613), v29, v32, v34)
									mBase = m.M
									v55 = m.ExcPending
									if v55 != 0 {
										return int32(0)
									} else {
										m.G0 = v9 + int32(16)
										v61 = *(*int32)(unsafe.Add(mBase, _c_F_inv_create[0]))
										F_recordDependencyOnOwner(m, int32(2613), v29, v61)
										mBase = m.M
										v63 = m.ExcPending
										if v63 != 0 {
											return int32(0)
										} else {
											v65 = *(*int32)(unsafe.Add(mBase, _c_F_inv_create[1]))
											if v65 != 0 {
												v67 = int32(0)
												F_RunObjectPostCreateHook(m, int32(2613), v29, v67, v67)
												mBase = m.M
												v70 = m.ExcPending
												if v70 != 0 {
													return int32(0)
												} else {
													F_CommandCounterIncrement(m)
													mBase = m.M
													v72 = m.ExcPending
													if v72 != 0 {
														return int32(0)
													} else {
														return v29
													}
												}
											} else {
												F_CommandCounterIncrement(m)
												mBase = m.M
												v72 = m.ExcPending
												if v72 != 0 {
													return int32(0)
												} else {
													return v29
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
			v29 = l0
			v32 = *(*int32)(unsafe.Add(mBase, _c_F_inv_create[0]))
			v34 = F_get_user_default_acl(m, int32(22), v32, int32(0))
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v32
				*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v29
				if v34 != 0 {
					*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v34
				} else {
					v39 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v9)+2)) = uint8(v39)
				}
				v41 = *(*int32)(unsafe.Add(mBase, uint32(v13)+52))
				v44 = F_heap_form_tuple(m, v41, v9+int32(4), v9)
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return int32(0)
				} else {
					F_CatalogTupleInsert(m, v13, v44)
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return int32(0)
					} else {
						F_pfree(m, v44)
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return int32(0)
						} else {
							F_sequence_close(m, v13, int32(3))
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return int32(0)
							} else {
								F_recordDependencyOnNewAcl(m, int32(2613), v29, v32, v34)
								mBase = m.M
								v55 = m.ExcPending
								if v55 != 0 {
									return int32(0)
								} else {
									m.G0 = v9 + int32(16)
									v61 = *(*int32)(unsafe.Add(mBase, _c_F_inv_create[0]))
									F_recordDependencyOnOwner(m, int32(2613), v29, v61)
									mBase = m.M
									v63 = m.ExcPending
									if v63 != 0 {
										return int32(0)
									} else {
										v65 = *(*int32)(unsafe.Add(mBase, _c_F_inv_create[1]))
										if v65 != 0 {
											v67 = int32(0)
											F_RunObjectPostCreateHook(m, int32(2613), v29, v67, v67)
											mBase = m.M
											v70 = m.ExcPending
											if v70 != 0 {
												return int32(0)
											} else {
												F_CommandCounterIncrement(m)
												mBase = m.M
												v72 = m.ExcPending
												if v72 != 0 {
													return int32(0)
												} else {
													return v29
												}
											}
										} else {
											F_CommandCounterIncrement(m)
											mBase = m.M
											v72 = m.ExcPending
											if v72 != 0 {
												return int32(0)
											} else {
												return v29
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
func F_inv_open(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	v7 = m.G0
	v9 = v7 + int32(-64)
	m.G0 = v9
	v21 = l1<<(uint(int32(14))%32)>>(uint(int32(31))%32)&int32(3) | int32(base.Ui32(l1)>>(uint(int32(18))%32))&int32(1)
	if v21 != 0 {
		v23 = l1 & int32(_a_F_inv_open_0)
		if v23 == int32(0) {
			v27 = *(*int32)(unsafe.Add(mBase, _c_F_inv_open[0]))
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
			v29 = v28
		} else {
			v29 = int32(0)
		}
		v30 = F_LargeObjectExistsWithSnapshot(m, l0, v29)
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return int32(0)
		} else {
			if v30 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v89 = m.ExcPending
				if v89 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(67137668))
					mBase = m.M
					v92 = m.ExcPending
					if v92 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = l0
						F_errmsg(m, int32(_a_F_inv_open_1), v7+int32(-16))
						mBase = m.M
						v98 = m.ExcPending
						if v98 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_inv_open_2), int32(247), int32(_a_F_inv_open_3))
							mBase = m.M
							v103 = m.ExcPending
							if v103 != 0 {
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
				if v21&int32(1) == int32(0) {
					if v23 == int32(0) {
						v57 = F_MemoryContextAlloc(m, l2, int32(32))
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v57)+24)) = v21
							*(*int64)(unsafe.Add(mBase, uint32(v57)+16)) = int64(0)
							*(*int32)(unsafe.Add(mBase, uint32(v57))) = l0
							*(*int32)(unsafe.Add(mBase, uint32(v57)+8)) = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v57)+4)) = v29
							m.G0 = v9 - int32(-64)
							return v57
						}
					} else {
						v50 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_inv_open[1])))
						if v50 != 0 {
							v57 = F_MemoryContextAlloc(m, l2, int32(32))
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v57)+24)) = v21
								*(*int64)(unsafe.Add(mBase, uint32(v57)+16)) = int64(0)
								*(*int32)(unsafe.Add(mBase, uint32(v57))) = l0
								*(*int32)(unsafe.Add(mBase, uint32(v57)+8)) = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v57)+4)) = v29
								m.G0 = v9 - int32(-64)
								return v57
							}
						} else {
							v52 = *(*int32)(unsafe.Add(mBase, _c_F_inv_open[2]))
							v54 = F_pg_largeobject_aclcheck_snapshot(m, l0, v52, int64(4), v29)
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
								return int32(0)
							} else {
								if v54 != 0 {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v125 = m.ExcPending
									if v125 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(16797828))
										mBase = m.M
										v128 = m.ExcPending
										if v128 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l0
											F_errmsg(m, int32(_a_F_inv_open_4), v7+int32(-48))
											mBase = m.M
											v134 = m.ExcPending
											if v134 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_inv_open_2), int32(272), int32(_a_F_inv_open_3))
												mBase = m.M
												v139 = m.ExcPending
												if v139 != 0 {
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
									v57 = F_MemoryContextAlloc(m, l2, int32(32))
									mBase = m.M
									v58 = m.ExcPending
									if v58 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v57)+24)) = v21
										*(*int64)(unsafe.Add(mBase, uint32(v57)+16)) = int64(0)
										*(*int32)(unsafe.Add(mBase, uint32(v57))) = l0
										*(*int32)(unsafe.Add(mBase, uint32(v57)+8)) = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v57)+4)) = v29
										m.G0 = v9 - int32(-64)
										return v57
									}
								}
							}
						}
					}
				} else {
					v41 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_inv_open[1])))
					if v41 != 0 {
						if v23 == int32(0) {
							v57 = F_MemoryContextAlloc(m, l2, int32(32))
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v57)+24)) = v21
								*(*int64)(unsafe.Add(mBase, uint32(v57)+16)) = int64(0)
								*(*int32)(unsafe.Add(mBase, uint32(v57))) = l0
								*(*int32)(unsafe.Add(mBase, uint32(v57)+8)) = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v57)+4)) = v29
								m.G0 = v9 - int32(-64)
								return v57
							}
						} else {
							v50 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_inv_open[1])))
							if v50 != 0 {
								v57 = F_MemoryContextAlloc(m, l2, int32(32))
								mBase = m.M
								v58 = m.ExcPending
								if v58 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v57)+24)) = v21
									*(*int64)(unsafe.Add(mBase, uint32(v57)+16)) = int64(0)
									*(*int32)(unsafe.Add(mBase, uint32(v57))) = l0
									*(*int32)(unsafe.Add(mBase, uint32(v57)+8)) = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v57)+4)) = v29
									m.G0 = v9 - int32(-64)
									return v57
								}
							} else {
								v52 = *(*int32)(unsafe.Add(mBase, _c_F_inv_open[2]))
								v54 = F_pg_largeobject_aclcheck_snapshot(m, l0, v52, int64(4), v29)
								mBase = m.M
								v55 = m.ExcPending
								if v55 != 0 {
									return int32(0)
								} else {
									if v54 != 0 {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v125 = m.ExcPending
										if v125 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(16797828))
											mBase = m.M
											v128 = m.ExcPending
											if v128 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l0
												F_errmsg(m, int32(_a_F_inv_open_4), v7+int32(-48))
												mBase = m.M
												v134 = m.ExcPending
												if v134 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_inv_open_2), int32(272), int32(_a_F_inv_open_3))
													mBase = m.M
													v139 = m.ExcPending
													if v139 != 0 {
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
										v57 = F_MemoryContextAlloc(m, l2, int32(32))
										mBase = m.M
										v58 = m.ExcPending
										if v58 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v57)+24)) = v21
											*(*int64)(unsafe.Add(mBase, uint32(v57)+16)) = int64(0)
											*(*int32)(unsafe.Add(mBase, uint32(v57))) = l0
											*(*int32)(unsafe.Add(mBase, uint32(v57)+8)) = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v57)+4)) = v29
											m.G0 = v9 - int32(-64)
											return v57
										}
									}
								}
							}
						}
					} else {
						v43 = *(*int32)(unsafe.Add(mBase, _c_F_inv_open[2]))
						v45 = F_pg_largeobject_aclcheck_snapshot(m, l0, v43, int64(2), v29)
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return int32(0)
						} else {
							if v45 != 0 {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v107 = m.ExcPending
								if v107 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(16797828))
									mBase = m.M
									v110 = m.ExcPending
									if v110 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = l0
										F_errmsg(m, int32(_a_F_inv_open_4), v7+int32(-32))
										mBase = m.M
										v116 = m.ExcPending
										if v116 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_inv_open_2), int32(260), int32(_a_F_inv_open_3))
											mBase = m.M
											v121 = m.ExcPending
											if v121 != 0 {
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
								if v23 == int32(0) {
									v57 = F_MemoryContextAlloc(m, l2, int32(32))
									mBase = m.M
									v58 = m.ExcPending
									if v58 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v57)+24)) = v21
										*(*int64)(unsafe.Add(mBase, uint32(v57)+16)) = int64(0)
										*(*int32)(unsafe.Add(mBase, uint32(v57))) = l0
										*(*int32)(unsafe.Add(mBase, uint32(v57)+8)) = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v57)+4)) = v29
										m.G0 = v9 - int32(-64)
										return v57
									}
								} else {
									v50 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_inv_open[1])))
									if v50 != 0 {
										v57 = F_MemoryContextAlloc(m, l2, int32(32))
										mBase = m.M
										v58 = m.ExcPending
										if v58 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v57)+24)) = v21
											*(*int64)(unsafe.Add(mBase, uint32(v57)+16)) = int64(0)
											*(*int32)(unsafe.Add(mBase, uint32(v57))) = l0
											*(*int32)(unsafe.Add(mBase, uint32(v57)+8)) = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v57)+4)) = v29
											m.G0 = v9 - int32(-64)
											return v57
										}
									} else {
										v52 = *(*int32)(unsafe.Add(mBase, _c_F_inv_open[2]))
										v54 = F_pg_largeobject_aclcheck_snapshot(m, l0, v52, int64(4), v29)
										mBase = m.M
										v55 = m.ExcPending
										if v55 != 0 {
											return int32(0)
										} else {
											if v54 != 0 {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v125 = m.ExcPending
												if v125 != 0 {
													return int32(0)
												} else {
													F_errcode(m, int32(16797828))
													mBase = m.M
													v128 = m.ExcPending
													if v128 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l0
														F_errmsg(m, int32(_a_F_inv_open_4), v7+int32(-48))
														mBase = m.M
														v134 = m.ExcPending
														if v134 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_inv_open_2), int32(272), int32(_a_F_inv_open_3))
															mBase = m.M
															v139 = m.ExcPending
															if v139 != 0 {
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
												v57 = F_MemoryContextAlloc(m, l2, int32(32))
												mBase = m.M
												v58 = m.ExcPending
												if v58 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v57)+24)) = v21
													*(*int64)(unsafe.Add(mBase, uint32(v57)+16)) = int64(0)
													*(*int32)(unsafe.Add(mBase, uint32(v57))) = l0
													*(*int32)(unsafe.Add(mBase, uint32(v57)+8)) = int32(0)
													*(*int32)(unsafe.Add(mBase, uint32(v57)+4)) = v29
													m.G0 = v9 - int32(-64)
													return v57
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
		v73 = m.ExcPending
		if v73 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50856066))
			mBase = m.M
			v76 = m.ExcPending
			if v76 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v9))) = l1
				F_errmsg(m, int32(_a_F_inv_open_5), v9)
				mBase = m.M
				v80 = m.ExcPending
				if v80 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_inv_open_2), int32(235), int32(_a_F_inv_open_3))
					mBase = m.M
					v85 = m.ExcPending
					if v85 != 0 {
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
func F_inv_read(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v26 int64
	_ = v26
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int64
	_ = v121
	var v123 int64
	_ = v123
	var v124 int64
	_ = v124
	var v126 int64
	_ = v126
	var v128 int64
	_ = v128
	var v130 int64
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v158 int64
	_ = v158
	var v159 int64
	_ = v159
	var v163 int32
	_ = v163
	var v167 int64
	_ = v167
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v187 int64
	_ = v187
	var v188 int64
	_ = v188
	var v189 int64
	_ = v189
	var v196 int64
	_ = v196
	var v198 int64
	_ = v198
	var v200 int64
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int64
	_ = v204
	var v208 int32
	_ = v208
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v227 int32
	_ = v227
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v294 int64
	_ = v294
	var v301 int32
	_ = v301
	var v306 int32
	_ = v306
	v17 = m.G0
	v19 = v17 - int32(128)
	m.G0 = v19
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v21&int32(1) != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L19
	} else {
		goto L79
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L19
	} else {
		goto L76
	}
L3:
	;
	if l2 <= int32(0) {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	goto L5
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L19
	} else {
		goto L72
	}
L6:
	;
	m.G0 = v19 + int32(128)
	return v241
L7:
	;
	v241 = int32(0)
	goto L6
L8:
	;
	goto L9
L9:
	;
	v26 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_inv_read[0]))
	v34 = *(*int32)(unsafe.Add(mBase, _c_F_inv_read[1]))
	if v34 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v35 = v31
	goto L12
L11:
	;
	v35 = int32(0)
	goto L12
L12:
	;
	if v35 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v38 = int32(_a_F_inv_read_0)
	v39 = *(*int32)(unsafe.Add(mBase, _c_F_inv_read[2]))
	v42 = *(*int32)(unsafe.Add(mBase, _c_F_inv_read[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_inv_read[2])) = v42
	if v31 != 0 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_ScanKeyInit(m, v19+int32(32), int32(1), int32(3), int32(184), v71)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L19
	} else {
		goto L25
	}
L16:
	;
	v54 = v34
	goto L18
L17:
	;
	v47 = F_table_open(m, int32(2613), int32(3))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	if v54 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	return int32(0)
L20:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_inv_read[0])) = v47
	v53 = *(*int32)(unsafe.Add(mBase, _c_F_inv_read[1]))
	v54 = v53
	goto L18
L21:
	;
	v60 = F_index_open(m, int32(2683), int32(3))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L19
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_inv_read[2])) = v39
	goto L15
L24:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_inv_read[1])) = v60
	goto L23
L25:
	;
	F_ScanKeyInit(m, v19+int32(80), int32(2), int32(4), int32(150), base.I32_wrap_i64(int64(base.Ui64(v26)>>(uint(int64(11))%64))))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L19
	} else {
		goto L26
	}
L26:
	;
	v86 = *(*int32)(unsafe.Add(mBase, _c_F_inv_read[0]))
	v88 = *(*int32)(unsafe.Add(mBase, _c_F_inv_read[1]))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v93 = F_systable_beginscan_ordered(m, v86, v88, v89, int32(2), v19+int32(32))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L19
	} else {
		goto L27
	}
L27:
	;
	v100 = int32(0)
	goto L28
L28:
	;
	v113 = F_systable_getnext_ordered(m, v93, int32(1))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L19
	} else {
		goto L30
	}
L29:
	;
	F_systable_endscan_ordered(m, v93)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L19
	} else {
		goto L71
	}
L30:
	;
	if v113 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v113)+16))
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+20)))
	if v116&int32(1) != 0 {
		goto L2
	} else {
		goto L34
	}
L32:
	;
	v227 = v100
	goto L33
L33:
	;
	goto L29
L34:
	;
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+22)))
	v120 = v115 + v119
	v121 = int64(*(*int32)(unsafe.Add(mBase, uint32(v120)+4)))
	v123 = v121 << (uint(int64(11)) % 64)
	v124 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	if base.Ui64(v124) < base.Ui64(v123) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v126 = v123 - v124
	v128 = base.I64_extend_i32_s(l2 - v100)
	if v126 < v128 {
		goto L40
	} else {
		goto L41
	}
L36:
	;
	v163 = v100
	v167 = v124
	goto L37
L37:
	;
	if l2 <= v163 {
		v217 = v163
		goto L51
	} else {
		goto L52
	}
L38:
	;
	v158 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	v159 = v158 + v130
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v159
	v163 = v131 + v100
	v167 = v159
	goto L37
L39:
	;
	v155 = F__emscripten_memset_bulkmem(m, v132, base.I32_extend8_s(int32(0)), v152)
	mBase = m.M
	goto L50
L40:
	;
	v130 = v126
	goto L42
L41:
	;
	v130 = v128
	goto L42
L42:
	;
	v131 = base.I32_wrap_i64(v130)
	v132 = l1 + v100
	if v132&int32(3) != 0 {
		v152 = v131
		goto L39
	} else {
		goto L43
	}
L43:
	;
	if base.Ui32(int32(1024)) < base.Ui32(v131) {
		v152 = v131
		goto L39
	} else {
		goto L44
	}
L44:
	;
	if v131&int32(3) != 0 {
		v152 = v131
		goto L39
	} else {
		goto L45
	}
L45:
	;
	v139 = v131 + v132
	if base.Ui32(v139) <= base.Ui32(v132) {
		goto L38
	} else {
		goto L46
	}
L46:
	;
	v141 = v100 + (l1 + int32(4))
	if base.Ui32(v141) < base.Ui32(v139) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v143 = v139
	goto L49
L48:
	;
	v143 = v141
	goto L49
L49:
	;
	v152 = (v143+(l1^int32(-1))-v100)&int32(-4) + int32(4)
	goto L39
L50:
	;
	goto L38
L51:
	;
	if v217 < l2 {
		v100 = v217
		goto L28
	} else {
		goto L70
	}
L52:
	;
	v171 = v120 + int32(8)
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120)+8)))
	v174 = v172 & int32(3)
	if v174 != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v175 = F_detoast_attr(m, v171)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L19
	} else {
		goto L56
	}
L54:
	;
	v177 = v171
	goto L55
L55:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v177)))
	v182 = int32(base.Ui32(v178)>>(uint(int32(2))%32)) - int32(4)
	if base.Ui32(v178-int32(_a_F_inv_read_1)) <= base.Ui32(int32(-8197)) {
		goto L1
	} else {
		goto L57
	}
L56:
	;
	v177 = v175
	goto L55
L57:
	;
	v187 = base.I64_extend_i32_s(v182)
	v188 = v167 - v123
	v189 = base.I64_extend32_s(v188)
	if v189 < v187 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v196 = v187 - v189
	v198 = base.I64_extend_i32_s(l2 - v163)
	if v196 < v198 {
		goto L61
	} else {
		goto L62
	}
L59:
	;
	v208 = v163
	goto L60
L60:
	;
	if v174 == int32(0) {
		v217 = v208
		goto L51
	} else {
		goto L68
	}
L61:
	;
	v200 = v196
	goto L63
L62:
	;
	v200 = v198
	goto L63
L63:
	;
	v201 = base.I32_wrap_i64(v200)
	if v201 != 0 {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	v204 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v204 + v200
	v208 = v163 + v201
	goto L60
L65:
	;
	v202 = F__emscripten_memcpy_bulkmem(m, l1+v163, v177+base.I32_wrap_i64(v188)+int32(4), v201)
	mBase = m.M
	goto L67
L66:
	;
	goto L67
L67:
	;
	goto L64
L68:
	;
	F_pfree(m, v177)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L19
	} else {
		goto L69
	}
L69:
	;
	v217 = v208
	goto L51
L70:
	;
	v227 = v217
	goto L33
L71:
	;
	v241 = v227
	goto L6
L72:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L19
	} else {
		goto L73
	}
L73:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v264
	F_errmsg(m, int32(_a_F_inv_read_2), v19)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L19
	} else {
		goto L74
	}
L74:
	;
	F_errfinish(m, int32(_a_F_inv_read_3), int32(469), int32(_a_F_inv_read_4))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L19
	} else {
		goto L75
	}
L75:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L76:
	;
	F_errmsg_internal(m, int32(_a_F_inv_read_5), int32(0))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L19
	} else {
		goto L77
	}
L77:
	;
	F_errfinish(m, int32(_a_F_inv_read_3), int32(496), int32(_a_F_inv_read_4))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L19
	} else {
		goto L78
	}
L78:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L79:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L19
	} else {
		goto L80
	}
L80:
	;
	v294 = *(*int64)(unsafe.Add(mBase, uint32(v120)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+24)) = v182
	*(*int64)(unsafe.Add(mBase, uint32(v19)+16)) = v294
	F_errmsg(m, int32(_a_F_inv_read_6), v19+int32(16))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L19
	} else {
		goto L81
	}
L81:
	;
	F_errfinish(m, int32(_a_F_inv_read_3), int32(153), int32(_a_F_inv_read_7))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L19
	} else {
		goto L82
	}
L82:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_irish_ISO_8859_1_stem(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v125 int32
	_ = v125
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v175 int32
	_ = v175
	var v184 int32
	_ = v184
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v241 int32
	_ = v241
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v273 int32
	_ = v273
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v289 int32
	_ = v289
	var v298 int32
	_ = v298
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v316 int32
	_ = v316
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v345 int32
	_ = v345
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v404 int32
	_ = v404
	var v408 int32
	_ = v408
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v4
	v8 = F_find_among(m, l0, int32(_a_F_irish_ISO_8859_1_stem_0), int32(24))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v448
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v4
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v78)+4)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(v78)+8)) = v79
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v78))) = v82
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v93 < v84 {
		goto L38
	} else {
		goto L39
	}
L3:
	;
	return int32(0)
L4:
	;
	if v8 == int32(0) {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v14
	switch v8 - int32(1) {
	case 0:
		goto L15
	case 1:
		goto L14
	case 2:
		goto L13
	case 3:
		goto L12
	case 4:
		goto L11
	case 5:
		goto L10
	case 6:
		goto L9
	case 7:
		goto L8
	case 8:
		goto L7
	case 9:
		goto L6
	default:
		goto L2
	}
L6:
	;
	v72 = F_slice_from_s(m, l0, int32(1), int32(_a_F_irish_ISO_8859_1_stem_1))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L3
	} else {
		goto L34
	}
L7:
	;
	v66 = F_slice_from_s(m, l0, int32(1), int32(_a_F_irish_ISO_8859_1_stem_2))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L3
	} else {
		goto L32
	}
L8:
	;
	v60 = F_slice_from_s(m, l0, int32(1), int32(_a_F_irish_ISO_8859_1_stem_3))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L3
	} else {
		goto L30
	}
L9:
	;
	v54 = F_slice_from_s(m, l0, int32(1), int32(_a_F_irish_ISO_8859_1_stem_4))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L3
	} else {
		goto L28
	}
L10:
	;
	v48 = F_slice_from_s(m, l0, int32(1), int32(_a_F_irish_ISO_8859_1_stem_5))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L3
	} else {
		goto L26
	}
L11:
	;
	v42 = F_slice_from_s(m, l0, int32(1), int32(_a_F_irish_ISO_8859_1_stem_6))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L3
	} else {
		goto L24
	}
L12:
	;
	v36 = F_slice_from_s(m, l0, int32(1), int32(_a_F_irish_ISO_8859_1_stem_7))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L3
	} else {
		goto L22
	}
L13:
	;
	v30 = F_slice_from_s(m, l0, int32(1), int32(_a_F_irish_ISO_8859_1_stem_8))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L3
	} else {
		goto L20
	}
L14:
	;
	v24 = F_slice_from_s(m, l0, int32(1), int32(_a_F_irish_ISO_8859_1_stem_9))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L3
	} else {
		goto L18
	}
L15:
	;
	v18 = F_slice_del(m, l0)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L3
	} else {
		goto L16
	}
L16:
	;
	if int32(0) <= v18 {
		goto L2
	} else {
		goto L17
	}
L17:
	;
	v448 = v18
	goto L1
L18:
	;
	if int32(0) <= v24 {
		goto L2
	} else {
		goto L19
	}
L19:
	;
	v448 = v24
	goto L1
L20:
	;
	if int32(0) <= v30 {
		goto L2
	} else {
		goto L21
	}
L21:
	;
	v448 = v30
	goto L1
L22:
	;
	if int32(0) <= v36 {
		goto L2
	} else {
		goto L23
	}
L23:
	;
	v448 = v36
	goto L1
L24:
	;
	if int32(0) <= v42 {
		goto L2
	} else {
		goto L25
	}
L25:
	;
	v448 = v42
	goto L1
L26:
	;
	if int32(0) <= v48 {
		goto L2
	} else {
		goto L27
	}
L27:
	;
	v448 = v48
	goto L1
L28:
	;
	if int32(0) <= v54 {
		goto L2
	} else {
		goto L29
	}
L29:
	;
	v448 = v54
	goto L1
L30:
	;
	if int32(0) <= v60 {
		goto L2
	} else {
		goto L31
	}
L31:
	;
	v448 = v60
	goto L1
L32:
	;
	if int32(0) <= v66 {
		goto L2
	} else {
		goto L33
	}
L33:
	;
	v448 = v66
	goto L1
L34:
	;
	if v72 < int32(0) {
		v448 = v72
		goto L1
	} else {
		goto L35
	}
L35:
	;
	goto L2
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v84
	v316 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v316
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v316
	v321 = F_find_among_b(m, l0, int32(_a_F_irish_ISO_8859_1_stem_10), int32(16))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L3
	} else {
		goto L100
	}
L37:
	;
	if v133 < int32(0) {
		goto L36
	} else {
		goto L52
	}
L38:
	;
	v95 = v84
	goto L40
L39:
	;
	v95 = v93
	goto L40
L40:
	;
	v102 = v84
	goto L42
L41:
	;
	v133 = v113
	goto L37
L42:
	;
	if v102 == v95 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v133 = int32(-1)
	goto L37
L45:
	;
	goto L46
L46:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106+v102))))
	if int32(250) < v108 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v125 = v102 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v125
	v102 = v125
	goto L42
L48:
	;
	v110 = v108 - int32(97)
	if v110 < int32(0) {
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v113 = int32(1)
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v110)>>(uint(int32(3))%32)))+uint32(_c_F_irish_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v117)>>(uint(v110&int32(7))%32))&v113 != 0 {
		goto L41
	} else {
		goto L50
	}
L50:
	;
	goto L47
L52:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v137 = v136 + v133
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v137
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v139)+8)) = v137
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v150 < v149 {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	if v193 < int32(0) {
		goto L36
	} else {
		goto L67
	}
L54:
	;
	v152 = v149
	goto L56
L55:
	;
	v152 = v150
	goto L56
L56:
	;
	v159 = v149
	goto L58
L57:
	;
	v193 = int32(1)
	goto L53
L58:
	;
	if v159 == v152 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v193 = int32(-1)
	goto L53
L61:
	;
	goto L62
L62:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165+v159))))
	if int32(250) < v167 {
		goto L57
	} else {
		goto L63
	}
L63:
	;
	v169 = v167 - int32(97)
	if v169 < int32(0) {
		goto L57
	} else {
		goto L64
	}
L64:
	;
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v169)>>(uint(int32(3))%32)))+uint32(_c_F_irish_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v175)>>(uint(v169&int32(7))%32))&int32(1) == int32(0) {
		goto L57
	} else {
		goto L65
	}
L65:
	;
	v184 = v159 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v184
	v159 = v184
	goto L58
L67:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v197 = v196 + v193
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v197
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v199)+4)) = v197
	v208 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v209 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v209 < v208 {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	if v249 < int32(0) {
		goto L36
	} else {
		goto L83
	}
L69:
	;
	v211 = v208
	goto L71
L70:
	;
	v211 = v209
	goto L71
L71:
	;
	v218 = v208
	goto L73
L72:
	;
	v249 = v229
	goto L68
L73:
	;
	if v218 == v211 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v249 = int32(-1)
	goto L68
L76:
	;
	goto L77
L77:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222+v218))))
	if int32(250) < v224 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v241 = v218 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v241
	v218 = v241
	goto L73
L79:
	;
	v226 = v224 - int32(97)
	if v226 < int32(0) {
		goto L78
	} else {
		goto L80
	}
L80:
	;
	v229 = int32(1)
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v226)>>(uint(int32(3))%32)))+uint32(_c_F_irish_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v233)>>(uint(v226&int32(7))%32))&v229 != 0 {
		goto L72
	} else {
		goto L81
	}
L81:
	;
	goto L78
L83:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v253 = v252 + v249
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v253
	v264 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v264 < v253 {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	if v307 < int32(0) {
		goto L36
	} else {
		goto L98
	}
L85:
	;
	v266 = v253
	goto L87
L86:
	;
	v266 = v264
	goto L87
L87:
	;
	v273 = v253
	goto L89
L88:
	;
	v307 = int32(1)
	goto L84
L89:
	;
	if v273 == v266 {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v307 = int32(-1)
	goto L84
L92:
	;
	goto L93
L93:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v279+v273))))
	if int32(250) < v281 {
		goto L88
	} else {
		goto L94
	}
L94:
	;
	v283 = v281 - int32(97)
	if v283 < int32(0) {
		goto L88
	} else {
		goto L95
	}
L95:
	;
	v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v283)>>(uint(int32(3))%32)))+uint32(_c_F_irish_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v289)>>(uint(v283&int32(7))%32))&int32(1) == int32(0) {
		goto L88
	} else {
		goto L96
	}
L96:
	;
	v298 = v273 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v298
	v273 = v298
	goto L89
L98:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v311 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v310))) = v311 + v307
	goto L36
L99:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v345
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v345
	v350 = F_find_among_b(m, l0, int32(_a_F_irish_ISO_8859_1_stem_11), int32(25))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L3
	} else {
		goto L111
	}
L100:
	;
	if v321 == int32(0) {
		goto L99
	} else {
		goto L101
	}
L101:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v325
	switch v321 - int32(1) {
	case 0:
		goto L103
	case 1:
		goto L102
	default:
		goto L99
	}
L102:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v336)))
	if v325 < v337 {
		goto L99
	} else {
		goto L107
	}
L103:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v329)+4))
	if v325 < v330 {
		goto L99
	} else {
		goto L104
	}
L104:
	;
	v332 = F_slice_del(m, l0)
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L3
	} else {
		goto L105
	}
L105:
	;
	if int32(0) <= v332 {
		goto L99
	} else {
		goto L106
	}
L106:
	;
	v448 = v332
	goto L1
L107:
	;
	v339 = F_slice_del(m, l0)
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L3
	} else {
		goto L108
	}
L108:
	;
	if v339 < int32(0) {
		v448 = v339
		goto L1
	} else {
		goto L109
	}
L109:
	;
	goto L99
L110:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v397
	v400 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v397-int32(2) <= v400 {
		goto L132
	} else {
		goto L133
	}
L111:
	;
	if v350 == int32(0) {
		goto L110
	} else {
		goto L112
	}
L112:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v354
	switch v350 - int32(1) {
	case 0:
		goto L118
	case 1:
		goto L117
	case 2:
		goto L116
	case 3:
		goto L115
	case 4:
		goto L114
	case 5:
		goto L113
	default:
		goto L110
	}
L113:
	;
	v391 = F_slice_from_s(m, l0, int32(3), int32(_a_F_irish_ISO_8859_1_stem_12))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L3
	} else {
		goto L130
	}
L114:
	;
	v385 = F_slice_from_s(m, l0, int32(5), int32(_a_F_irish_ISO_8859_1_stem_13))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L3
	} else {
		goto L128
	}
L115:
	;
	v379 = F_slice_from_s(m, l0, int32(4), int32(_a_F_irish_ISO_8859_1_stem_14))
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L3
	} else {
		goto L126
	}
L116:
	;
	v373 = F_slice_from_s(m, l0, int32(3), int32(_a_F_irish_ISO_8859_1_stem_15))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L3
	} else {
		goto L124
	}
L117:
	;
	v367 = F_slice_from_s(m, l0, int32(3), int32(_a_F_irish_ISO_8859_1_stem_16))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L3
	} else {
		goto L122
	}
L118:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v358)))
	if v354 < v359 {
		goto L110
	} else {
		goto L119
	}
L119:
	;
	v361 = F_slice_del(m, l0)
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L3
	} else {
		goto L120
	}
L120:
	;
	if int32(0) <= v361 {
		goto L110
	} else {
		goto L121
	}
L121:
	;
	v448 = v361
	goto L1
L122:
	;
	if int32(0) <= v367 {
		goto L110
	} else {
		goto L123
	}
L123:
	;
	v448 = v367
	goto L1
L124:
	;
	if int32(0) <= v373 {
		goto L110
	} else {
		goto L125
	}
L125:
	;
	v448 = v373
	goto L1
L126:
	;
	if int32(0) <= v379 {
		goto L110
	} else {
		goto L127
	}
L127:
	;
	v448 = v379
	goto L1
L128:
	;
	if int32(0) <= v385 {
		goto L110
	} else {
		goto L129
	}
L129:
	;
	v448 = v385
	goto L1
L130:
	;
	if v391 < int32(0) {
		v448 = v391
		goto L1
	} else {
		goto L131
	}
L131:
	;
	goto L110
L132:
	;
	v445 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v445
	v448 = int32(1)
	goto L1
L133:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v408 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v404+v397-int32(1)))))
	if v408&int32(224) != int32(96) {
		goto L132
	} else {
		goto L134
	}
L134:
	;
	if int32(1)<<(uint(v408)%32)&int32(_a_F_irish_ISO_8859_1_stem_17) == int32(0) {
		goto L132
	} else {
		goto L135
	}
L135:
	;
	v421 = F_find_among_b(m, l0, int32(_a_F_irish_ISO_8859_1_stem_18), int32(12))
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L3
	} else {
		goto L136
	}
L136:
	;
	if v421 == int32(0) {
		goto L132
	} else {
		goto L137
	}
L137:
	;
	v425 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v425
	switch v421 - int32(1) {
	case 0:
		goto L139
	case 1:
		goto L138
	default:
		goto L132
	}
L138:
	;
	v436 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v436)+4))
	if v425 < v437 {
		goto L132
	} else {
		goto L143
	}
L139:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v429)+8))
	if v425 < v430 {
		goto L132
	} else {
		goto L140
	}
L140:
	;
	v432 = F_slice_del(m, l0)
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L3
	} else {
		goto L141
	}
L141:
	;
	if int32(0) <= v432 {
		goto L132
	} else {
		goto L142
	}
L142:
	;
	v448 = v432
	goto L1
L143:
	;
	v439 = F_slice_del(m, l0)
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L3
	} else {
		goto L144
	}
L144:
	;
	if v439 < int32(0) {
		v448 = v439
		goto L1
	} else {
		goto L145
	}
L145:
	;
	goto L132
}
func F_is_redundant_with_indexclauses(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v53 int32
	_ = v53
	if l1 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v53
L2:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v8 <= int32(0) {
		v53 = int32(0)
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	v53 = int32(0)
	goto L1
L5:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v12 = int32(0)
	if v12 < v8 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v15 = v8
	goto L8
L7:
	;
	v15 = v12
	goto L8
L8:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v19 = int32(0)
	goto L9
L9:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v16+v19<<(uint(int32(2))%32))))
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+12)))
	if v29 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L4
L11:
	;
	v40 = v19 + int32(1)
	if v40 != v15 {
		v19 = v40
		goto L9
	} else {
		goto L16
	}
L12:
	;
	v30 = int32(1)
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	if l0 == v31 {
		v53 = v30
		goto L1
	} else {
		goto L13
	}
L13:
	;
	if v11 == int32(0) {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v31)+60))
	if v35 == v11 {
		v53 = v30
		goto L1
	} else {
		goto L15
	}
L15:
	;
	goto L11
L16:
	;
	goto L10
}
func F_issue_xlog_fsync(m *base.Module, l0 int32, l1 int64, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v29 int64
	_ = v29
	var v30 int64
	_ = v30
	var v34 int64
	_ = v34
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v92 int32
	_ = v92
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v112 int32
	_ = v112
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v166 int64
	_ = v166
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v178 int64
	_ = v178
	var v179 int64
	_ = v179
	var v183 int64
	_ = v183
	var v235 int32
	_ = v235
	var v236 int64
	_ = v236
	var v240 int32
	_ = v240
	var v258 int32
	_ = v258
	var v259 int64
	_ = v259
	var v263 int32
	_ = v263
	var v280 int32
	_ = v280
	var v281 int64
	_ = v281
	var v286 int32
	_ = v286
	var v287 int64
	_ = v287
	var v292 int32
	_ = v292
	v7 = m.G0
	v9 = v7 - int32(96)
	m.G0 = v9
	v12 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_issue_xlog_fsync[0])))
	if v12 != int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v9 + int32(96)
	return
L2:
	;
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_issue_xlog_fsync[1]))
	switch v16 - int32(2) {
	case 0, 2:
		goto L1
	default:
		goto L3
	}
L3:
	;
	v20 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_issue_xlog_fsync[2])))
	v23 = m.G0
	v25 = v23 - int32(16)
	m.G0 = v25
	if v20 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v39 = *(*int32)(unsafe.Add(mBase, _c_F_issue_xlog_fsync[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v39))) = int32(167772238)
	v43 = *(*int32)(unsafe.Add(mBase, _c_F_issue_xlog_fsync[1]))
	switch v43 {
	case 0:
		goto L12
	case 1:
		goto L10
	case 2, 4:
		goto L8
	default:
		goto L11
	}
L5:
	;
	F___clock_gettime(m, int32(1), v25)
	mBase = m.M
	v29 = int64(*(*int32)(unsafe.Add(mBase, uint32(v25)+8)))
	v30 = *(*int64)(unsafe.Add(mBase, uint32(v25)))
	v34 = v29 + v30*int64(1000000000)
	goto L7
L6:
	;
	v34 = int64(0)
	goto L7
L7:
	;
	m.G0 = v25 + int32(16)
	goto L4
L8:
	;
	v159 = *(*int32)(unsafe.Add(mBase, _c_F_issue_xlog_fsync[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v159))) = int32(0)
	v164 = int32(1)
	v166 = int64(0)
	v170 = m.G0
	v172 = v170 - int32(16)
	m.G0 = v172
	if v34 != v166 {
		goto L43
	} else {
		goto L44
	}
L9:
	;
	v125 = *(*int32)(unsafe.Add(mBase, _c_F_issue_xlog_fsync[4]))
	v129 = *(*int32)(unsafe.Add(mBase, _c_F_issue_xlog_fsync[5]))
	F_XLogFileName(m, v9+int32(32), l2, l1, v129)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L21
	} else {
		goto L37
	}
L10:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_issue_xlog_fsync[0])))
	if v83 != int32(1) {
		v112 = int32(0)
		goto L26
	} else {
		goto L27
	}
L11:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L21
	} else {
		goto L22
	}
L12:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_issue_xlog_fsync[0])))
	if v46 != int32(1) {
		v60 = int32(0)
		goto L14
	} else {
		goto L15
	}
L13:
	;
	if v60 == int32(0) {
		goto L8
	} else {
		goto L20
	}
L14:
	;
	goto L13
L15:
	;
	goto L16
L16:
	;
	v51 = F_fsync(m, l0)
	mBase = m.M
	if v51 != int32(-1) {
		v60 = v51
		goto L14
	} else {
		goto L18
	}
L17:
	;
	v60 = int32(-1)
	goto L14
L18:
	;
	v55 = *(*int32)(unsafe.Add(mBase, _c_F_issue_xlog_fsync[4]))
	if v55 == int32(27) {
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	v123 = int32(_a_F_issue_xlog_fsync_0)
	goto L9
L21:
	;
	return
L22:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v72 = *(*int32)(unsafe.Add(mBase, _c_F_issue_xlog_fsync[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v72
	F_errmsg_internal(m, int32(_a_F_issue_xlog_fsync_1), v9)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L21
	} else {
		goto L24
	}
L24:
	;
	F_errfinish(m, int32(_a_F_issue_xlog_fsync_2), int32(_a_F_issue_xlog_fsync_3), int32(_a_F_issue_xlog_fsync_4))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L21
	} else {
		goto L25
	}
L25:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L26:
	;
	if v112 == int32(0) {
		goto L8
	} else {
		goto L36
	}
L27:
	;
	goto L28
L28:
	;
	v92 = m.Env.X__syscall_fdatasync(m, l0)
	mBase = m.M
	if base.Ui32(int32(-4095)) <= base.Ui32(v92) {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	v112 = int32(-1)
	goto L26
L30:
	;
	if v100 != int32(-1) {
		v112 = v100
		goto L26
	} else {
		goto L34
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_issue_xlog_fsync[4])) = int32(0) - v92
	v100 = int32(-1)
	goto L33
L32:
	;
	v100 = v92
	goto L33
L33:
	;
	goto L30
L34:
	;
	v104 = *(*int32)(unsafe.Add(mBase, _c_F_issue_xlog_fsync[4]))
	if v104 == int32(27) {
		goto L28
	} else {
		goto L35
	}
L35:
	;
	goto L29
L36:
	;
	v123 = int32(_a_F_issue_xlog_fsync_5)
	goto L9
L37:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_issue_xlog_fsync[4])) = v125
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L21
	} else {
		goto L38
	}
L38:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L21
	} else {
		goto L39
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v9 + int32(32)
	F_errmsg(m, v123, v9+int32(16))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L21
	} else {
		goto L40
	}
L40:
	;
	F_errfinish(m, int32(_a_F_issue_xlog_fsync_2), int32(_a_F_issue_xlog_fsync_6), int32(_a_F_issue_xlog_fsync_4))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L21
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
	goto L1
L43:
	;
	F___clock_gettime(m, int32(1), v172)
	mBase = m.M
	v178 = int64(*(*int32)(unsafe.Add(mBase, uint32(v172)+8)))
	v179 = *(*int64)(unsafe.Add(mBase, uint32(v172)))
	v183 = v178 + (v179*int64(1000000000) - v34)
	goto L46
L44:
	;
	goto L45
L45:
	;
	v280 = int32(_a_F_issue_xlog_fsync_7)
	v281 = *(*int64)(unsafe.Add(mBase, _c_F_issue_xlog_fsync[6]))
	*(*int64)(unsafe.Add(mBase, _c_F_issue_xlog_fsync[6])) = v281 + base.I64_extend_i32_u(v164)
	v286 = int32(_a_F_issue_xlog_fsync_8)
	v287 = *(*int64)(unsafe.Add(mBase, _c_F_issue_xlog_fsync[7]))
	*(*int64)(unsafe.Add(mBase, _c_F_issue_xlog_fsync[7])) = v287 + v166
	F_pgstat_count_backend_io_op(m, int32(2), int32(3), v164, v164, v166)
	mBase = m.M
	v292 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_issue_xlog_fsync[8])) = uint8(v292)
	*(*uint8)(unsafe.Add(mBase, _c_F_issue_xlog_fsync[9])) = uint8(v292)
	m.G0 = v172 + int32(16)
	goto L42
L46:
	;
	v235 = int32(_a_F_issue_xlog_fsync_9)
	v236 = *(*int64)(unsafe.Add(mBase, _c_F_issue_xlog_fsync[10]))
	*(*int64)(unsafe.Add(mBase, _c_F_issue_xlog_fsync[10])) = v236 + v183
	v240 = *(*int32)(unsafe.Add(mBase, _c_F_issue_xlog_fsync[11]))
	if base.Ui32(int32(16)) < base.Ui32(v240) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	goto L45
L57:
	;
	if int32(1)<<(uint(v240)%32)&int32(_a_F_issue_xlog_fsync_10) == int32(0) {
		goto L56
	} else {
		goto L58
	}
L58:
	;
	v258 = int32(_a_F_issue_xlog_fsync_11)
	v259 = *(*int64)(unsafe.Add(mBase, _c_F_issue_xlog_fsync[12]))
	*(*int64)(unsafe.Add(mBase, _c_F_issue_xlog_fsync[12])) = v259 + v183
	v263 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_issue_xlog_fsync[8])) = uint8(v263)
	*(*uint8)(unsafe.Add(mBase, _c_F_issue_xlog_fsync[13])) = uint8(v263)
	goto L56
}
func F_isupper(m *base.Module, l0 int32) int32 {
	return base.B2i32(base.Ui32(l0-int32(65)) < base.Ui32(int32(26)))
}
func F_iswlower(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	v3 = F_casemap(m, l0, int32(1))
	return base.B2i32(v3 != l0)
}
func F_iswprint(m *base.Module, l0 int32) int32 {
	var v12 int32
	_ = v12
	var v23 int32
	_ = v23
	var v32 int32
	_ = v32
	if base.Ui32(l0) <= base.Ui32(int32(254)) {
		return base.B2i32(base.Ui32(int32(32)) < base.Ui32((l0+int32(1))&int32(127)))
	} else {
		v12 = int32(1)
		if base.Ui32(l0-int32(_a_F_iswprint_0)) < base.Ui32(int32(_a_F_iswprint_1)) {
			v32 = v12
		} else {
			if base.Ui32(l0) < base.Ui32(int32(_a_F_iswprint_2)) {
				v32 = v12
			} else {
				if base.Ui32(l0-int32(_a_F_iswprint_3)) < base.Ui32(int32(_a_F_iswprint_4)) {
					v32 = v12
				} else {
					v23 = int32(_a_F_iswprint_5)
					v32 = base.B2i32(l0&v23 != v23) & base.B2i32(base.Ui32(l0-int32(_a_F_iswprint_6)) < base.Ui32(int32(_a_F_iswprint_7)))
				}
			}
		}
		return v32
	}
}
func F_iswxdigit(m *base.Module, l0 int32) int32 {
	return base.B2i32(base.Ui32(l0-int32(48)) < base.Ui32(int32(10))) | base.B2i32(base.Ui32(l0|int32(32)-int32(97)) < base.Ui32(int32(6)))
}
func F_itemptr_comparator(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	v6 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+2)))
	v7 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0))))
	v8 = int32(16)
	v10 = v6 | v7<<(uint(v8)%32)
	v11 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+2)))
	v12 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
	v15 = v11 | v12<<(uint(v8)%32)
	if base.Ui32(v10) < base.Ui32(v15) {
		v26 = int32(-1)
	} else {
		if base.Ui32(v15) < base.Ui32(v10) {
			v26 = int32(1)
		} else {
			v20 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
			v21 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
			if base.Ui32(v20) < base.Ui32(v21) {
				v26 = int32(-1)
			} else {
				v26 = base.B2i32(base.Ui32(v21) < base.Ui32(v20))
			}
		}
	}
	return v26
}
func F_iterate_jsonb_values(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v14 = F_JsonbIteratorInit(m, l0+int32(4))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = v14
	goto L3
L3:
	;
	v37 = F_JsonbIteratorNext(m, v10+int32(28), v10+int32(8), int32(0))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L8
	}
L4:
	;
	m.G0 = v10 + int32(32)
	return
L5:
	;
	goto L4
L6:
	;
	if v37&int32(-2) != int32(2) {
		goto L3
	} else {
		goto L11
	}
L7:
	;
	if l1&int32(1) == int32(0) {
		goto L3
	} else {
		goto L9
	}
L8:
	;
	switch v37 {
	case 0:
		goto L5
	case 1:
		goto L7
	default:
		goto L6
	}
L9:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	F_add_to_tsvector(m, l2, v41, v42)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	goto L3
L11:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	switch v49 - int32(1) {
	case 0:
		goto L14
	case 1:
		goto L13
	case 2:
		goto L12
	default:
		goto L3
	}
L12:
	;
	if l1&int32(8) == int32(0) {
		goto L3
	} else {
		goto L21
	}
L13:
	;
	if l1&int32(4) == int32(0) {
		goto L3
	} else {
		goto L17
	}
L14:
	;
	if l1&int32(2) == int32(0) {
		goto L3
	} else {
		goto L15
	}
L15:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	F_add_to_tsvector(m, l2, v54, v55)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	goto L3
L17:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	v63 = F_DirectFunctionCall1Coll(m, int32(618), int32(0), v62)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v65 = F_strlen(m, v63)
	mBase = m.M
	F_add_to_tsvector(m, l2, v63, v65)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	F_pfree(m, v63)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	goto L3
L21:
	;
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+12)))
	if v72 == int32(1) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	F_add_to_tsvector(m, l2, int32(_a_F_iterate_jsonb_values_0), int32(4))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	F_add_to_tsvector(m, l2, int32(_a_F_iterate_jsonb_values_1), int32(5))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L26
	}
L25:
	;
	goto L3
L26:
	;
	goto L3
}
func F_ivfflatbuildphasename(m *base.Module, l0 int64) int32 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	v4 = l0 - int64(1)
	if base.Ui64(v4) <= base.Ui64(int64(3)) {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v4)<<(uint(int32(2))%32))+uint32(_c_F_ivfflatbuildphasename[0])))
		v13 = v12
	} else {
		v13 = int32(0)
	}
	return v13
}
func F_ivfflatendscan(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+28))
	F_tuplesort_end(m, v4)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v3)+24))
		F_MemoryContextDelete(m, v7)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			F_pfree(m, v3)
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = int32(0)
				return
			}
		}
	}
}
func F_ivfflatgettuple(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int64
	_ = v33
	var v38 int32
	_ = v38
	var v39 int64
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v101 float64
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v155 float64
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 float64
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 float64
	_ = v206
	var v208 int32
	_ = v208
	var v210 float64
	_ = v210
	var v219 int32
	_ = v219
	var v227 float64
	_ = v227
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v255 int32
	_ = v255
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	v3 = int32(0)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+16)))
	if v16 == int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+272))
	if v20 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	goto L3
L3:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
	v312 = int32(0)
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v15)+40))
	v315 = F_tuplesort_gettupleslot(m, v310, int32(1), v312, v313, v312)
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L9
	} else {
		goto L67
	}
L4:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v38 != 0 {
		goto L11
	} else {
		goto L12
	}
L5:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+268)))
	if v23 != int32(1) {
		goto L4
	} else {
		goto L8
	}
L6:
	;
	v32 = v20
	goto L7
L7:
	;
	v33 = *(*int64)(unsafe.Add(mBase, uint32(v32)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v32)+16)) = v33 + int64(1)
	goto L4
L8:
	;
	F_pgstat_assoc_relation(m, v19)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return int32(0)
L10:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+272))
	v32 = v31
	goto L7
L11:
	;
	v39 = *(*int64)(unsafe.Add(mBase, uint32(v38)))
	*(*int64)(unsafe.Add(mBase, uint32(v38))) = v39 + int64(1)
	goto L13
L12:
	;
	goto L13
L13:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v43 != 0 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	if int32(0) < v219 {
		goto L59
	} else {
		goto L60
	}
L15:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	switch v45 {
	case 0, 5:
		goto L18
	default:
		goto L19
	}
L16:
	;
	goto L17
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L9
	} else {
		goto L56
	}
L18:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
	if v60&int32(1) != 0 {
		goto L24
	} else {
		goto L25
	}
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L9
	} else {
		goto L20
	}
L20:
	;
	F_errmsg_internal(m, int32(_a_F_ivfflatgettuple_0), int32(0))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L9
	} else {
		goto L21
	}
L21:
	;
	F_errfinish(m, int32(_a_F_ivfflatgettuple_1), int32(391), int32(_a_F_ivfflatgettuple_2))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L9
	} else {
		goto L22
	}
L22:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L23:
	;
	v91 = int32(1)
	v93 = v3
	v101 = float64(1.7976931348623157e+308)
	goto L29
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v59)+60)) = int32(_a_F_ivfflatgettuple_3)
	v83 = v59
	v85 = v3
	goto L23
L25:
	;
	goto L26
L26:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v43)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v59)+60)) = int32(_a_F_ivfflatgettuple_4)
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v59)+52))
	if v68 == int32(0) {
		v83 = v59
		v85 = v65
		goto L23
	} else {
		goto L27
	}
L27:
	;
	v71 = int32(_a_F_ivfflatgettuple_5)
	v72 = *(*int32)(unsafe.Add(mBase, _c_F_ivfflatgettuple[0]))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v59)+24))
	*(*int32)(unsafe.Add(mBase, _c_F_ivfflatgettuple[0])) = v74
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v59)+56))
	v78 = F_HnswNormValue(m, v76, v77, v65)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L9
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ivfflatgettuple[0])) = v72
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v83 = v82
	v85 = v78
	goto L23
L29:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v103 = F_ReadBuffer(m, v102, v91)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L9
	} else {
		goto L31
	}
L30:
	;
	goto L14
L31:
	;
	F_LockBuffer(m, v103, int32(1))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L9
	} else {
		goto L32
	}
L32:
	;
	if v103 < int32(0) {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	v228 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v125)+16)))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v125+v228)))
	F_UnlockReleaseBuffer(m, v103)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L9
	} else {
		goto L54
	}
L34:
	;
	v126 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v125)+12)))
	if base.Ui32(v126) < base.Ui32(int32(25)) {
		v219 = v93
		v227 = v101
		goto L33
	} else {
		goto L38
	}
L35:
	;
	v111 = *(*int32)(unsafe.Add(mBase, _c_F_ivfflatgettuple[1]))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v111+(v103^int32(-1))<<(uint(int32(2))%32))))
	v125 = v117
	goto L34
L36:
	;
	goto L37
L37:
	;
	v119 = *(*int32)(unsafe.Add(mBase, _c_F_ivfflatgettuple[2]))
	v125 = v119 + v103<<(uint(int32(13))%32) + int32(-8192)
	goto L34
L38:
	;
	v130 = v126 + int32(_a_F_ivfflatgettuple_6)
	if v130&int32(_a_F_ivfflatgettuple_7) == int32(0) {
		v219 = v93
		v227 = v101
		goto L33
	} else {
		goto L39
	}
L39:
	;
	v143 = int32(1)
	v147 = v93
	v155 = v101
	goto L40
L40:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v83)+48))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v83)+56))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v143<<(uint(int32(2))%32)+(v125+int32(24))-int32(4))))
	v166 = v125 + v163&int32(_a_F_ivfflatgettuple_8)
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v83)+60))
	v170 = m.T0[v169].(func(*base.Module, int32, int32, int32, int32) int32)(m, v156, v157, v166+int32(8), v85)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L9
	} else {
		goto L42
	}
L41:
	;
	v219 = v208
	v227 = v210
	goto L33
L42:
	;
	v172 = *(*float64)(unsafe.Add(mBase, uint32(v170)))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v83)+8))
	if v147 < v173 {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	if v143 != int32(base.Ui32(v130)>>(uint(int32(2))%32))&int32(_a_F_ivfflatgettuple_9) {
		v143 = v143 + int32(1)
		v147 = v208
		v155 = v210
		goto L40
	} else {
		goto L53
	}
L44:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v83)+64))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v204)+8))
	v206 = *(*float64)(unsafe.Add(mBase, uint32(v205)+16))
	v208 = v202
	v210 = v206
	goto L43
L45:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v166)))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v83)+76))
	v179 = v176 + v147*int32(24)
	*(*float64)(unsafe.Add(mBase, uint32(v179)+16)) = v172
	*(*int32)(unsafe.Add(mBase, uint32(v179)+12)) = v175
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v83)+64))
	F_pairingheap_add(m, v182, v179)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L9
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	if base.F64_lt(v172, v155) == int32(0) {
		v208 = v147
		v210 = v155
		goto L43
	} else {
		goto L50
	}
L48:
	;
	v186 = v147 + int32(1)
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v83)+8))
	if v186 == v187 {
		v202 = v186
		goto L44
	} else {
		goto L49
	}
L49:
	;
	v208 = v186
	v210 = v155
	goto L43
L50:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v83)+64))
	v193 = F_pairingheap_remove_first(m, v192)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L9
	} else {
		goto L51
	}
L51:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v166)))
	*(*float64)(unsafe.Add(mBase, uint32(v193)+16)) = v172
	*(*int32)(unsafe.Add(mBase, uint32(v193)+12)) = v195
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v83)+64))
	F_pairingheap_add(m, v198, v193)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L9
	} else {
		goto L52
	}
L52:
	;
	v202 = v147
	goto L44
L53:
	;
	goto L41
L54:
	;
	if v230 != int32(-1) {
		v91 = v230
		v93 = v219
		v101 = v227
		goto L29
	} else {
		goto L55
	}
L55:
	;
	goto L30
L56:
	;
	F_errmsg_internal(m, int32(_a_F_ivfflatgettuple_10), int32(0))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L9
	} else {
		goto L57
	}
L57:
	;
	F_errfinish(m, int32(_a_F_ivfflatgettuple_1), int32(386), int32(_a_F_ivfflatgettuple_2))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L9
	} else {
		goto L58
	}
L58:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L59:
	;
	v255 = v219
	goto L62
L60:
	;
	goto L61
L61:
	;
	F_GetScanItems(m, l0, v85)
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L9
	} else {
		goto L66
	}
L62:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v83)+64))
	v265 = F_pairingheap_remove_first(m, v264)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L9
	} else {
		goto L64
	}
L63:
	;
	goto L61
L64:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v83)+68))
	v268 = int32(1)
	v269 = v255 - v268
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v265)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v267+v269<<(uint(int32(2))%32)))) = v273
	if base.Ui32(v268) < base.Ui32(v255) {
		v255 = v269
		goto L62
	} else {
		goto L65
	}
L65:
	;
	goto L63
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v85
	v294 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+16)) = uint8(v294)
	goto L3
L67:
	;
	if v315 == int32(0) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	goto L71
L69:
	;
	goto L70
L70:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v15)+40))
	v365 = int32(*(*int16)(unsafe.Add(mBase, uint32(v364)+6)))
	if v365 <= int32(1) {
		goto L79
	} else {
		goto L80
	}
L71:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v15)+72))
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	if v333 == v334 {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	goto L70
L73:
	;
	return int32(0)
L74:
	;
	goto L75
L75:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
	F_GetScanItems(m, l0, v338)
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L9
	} else {
		goto L76
	}
L76:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
	v343 = int32(0)
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v15)+40))
	v346 = F_tuplesort_gettupleslot(m, v341, int32(1), v343, v344, v343)
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L9
	} else {
		goto L77
	}
L77:
	;
	if v346 == int32(0) {
		goto L71
	} else {
		goto L78
	}
L78:
	;
	goto L72
L79:
	;
	F_slot_getsomeattrs_int(m, v364, int32(2))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L9
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v364)+16))
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v371)+4))
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v372)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v373
	v375 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v372)+4)))
	v376 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+72)) = uint8(v376)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+84)) = uint8(v376)
	*(*uint16)(unsafe.Add(mBase, uint32(l0-int32(-64)))) = uint16(v375)
	return int32(1)
L82:
	;
	goto L81
}
