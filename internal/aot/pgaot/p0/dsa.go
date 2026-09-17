package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_dsa_allocate_extended(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
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
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v188 int32
	_ = v188
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v230 int32
	_ = v230
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v299 int32
	_ = v299
	var v304 int32
	_ = v304
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v347 int32
	_ = v347
	v4 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(80)
	m.G0 = v13
	v16 = l2 & int32(1)
	if v16&base.B2i32(l1 < v4)|base.B2i32(v16 == v4)&base.B2i32(base.Ui32(int32(1073741824)) <= base.Ui32(l1)) == v4 {
		if base.Ui32(int32(_a_F_dsa_allocate_extended_0)) <= base.Ui32(l1) {
			v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v31 = int32(0)
			v33 = F_alloc_object(m, l0, v31)
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return int32(0)
			} else {
				if v33 == int32(0) {
					if l2&int32(2) != 0 {
						v347 = v31
						m.G0 = v13 + int32(80)
						return v347
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(_a_F_dsa_allocate_extended_1))
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(_a_F_dsa_allocate_extended_2), int32(0))
								mBase = m.M
								v51 = m.ExcPending
								if v51 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v13))) = l1
									F_errdetail(m, int32(_a_F_dsa_allocate_extended_3), v13)
									mBase = m.M
									v55 = m.ExcPending
									if v55 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_dsa_allocate_extended_4), int32(709), int32(_a_F_dsa_allocate_extended_5))
										mBase = m.M
										v60 = m.ExcPending
										if v60 != 0 {
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
				} else {
					v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v65 = F_LWLockAcquire(m, v61+int32(1476), int32(0))
					mBase = m.M
					v66 = m.ExcPending
					if v66 != 0 {
						return int32(0)
					} else {
						v70 = int32(base.Ui32(l1+int32(4095)) >> (uint(int32(12)) % 32))
						v71 = F_get_best_segment(m, l0, v70)
						mBase = m.M
						v72 = m.ExcPending
						if v72 != 0 {
							return int32(0)
						} else {
							if v71 != 0 {
								v106 = v71
								v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)+12))
								v110 = F_FreePageManagerGet(m, v107, v70, v13+int32(76))
								mBase = m.M
								v111 = m.ExcPending
								if v111 != 0 {
									return int32(0)
								} else {
									if v110 == int32(0) {
										F_errstart_cold(m, int32(22), int32(0))
										mBase = m.M
										v245 = m.ExcPending
										if v245 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v70
											F_errmsg_internal(m, int32(_a_F_dsa_allocate_extended_6), v13+int32(32))
											mBase = m.M
											v251 = m.ExcPending
											if v251 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_dsa_allocate_extended_4), int32(744), int32(_a_F_dsa_allocate_extended_5))
												mBase = m.M
												v256 = m.ExcPending
												if v256 != 0 {
													return int32(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									} else {
										v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										F_LWLockRelease(m, v114+int32(1476))
										mBase = m.M
										v118 = m.ExcPending
										if v118 != 0 {
											return int32(0)
										} else {
											v119 = *(*int32)(unsafe.Add(mBase, uint32(v13)+76))
											v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											v124 = F_LWLockAcquire(m, v120+int32(256), int32(0))
											mBase = m.M
											v125 = m.ExcPending
											if v125 != 0 {
												return int32(0)
											} else {
												v129 = v119 << (uint(int32(12)) % 32)
												v134 = base.I32_div_s(v106-l0-int32(8), int32(20))
												v137 = v129 | v134<<(uint(int32(27))%32)
												F_init_span(m, l0, v33, v30+int32(256), v137, v70, int32(1))
												mBase = m.M
												v140 = m.ExcPending
												if v140 != 0 {
													return int32(0)
												} else {
													v141 = *(*int32)(unsafe.Add(mBase, uint32(v106)+16))
													v142 = *(*int32)(unsafe.Add(mBase, uint32(v13)+76))
													*(*int32)(unsafe.Add(mBase, uint32(v141+v142<<(uint(int32(2))%32)))) = v33
													v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
													F_LWLockRelease(m, v147+int32(256))
													mBase = m.M
													v151 = m.ExcPending
													if v151 != 0 {
														return int32(0)
													} else {
														if l2&int32(4) == int32(0) {
															v347 = v137
															m.G0 = v13 + int32(80)
															return v347
														} else {
															if v137 != 0 {
																v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
																v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
																v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)+1468))
																if v156 != v158 {
																	v163 = F_LWLockAcquire(m, v157+int32(1476), int32(0))
																	mBase = m.M
																	v164 = m.ExcPending
																	if v164 != 0 {
																		return int32(0)
																	} else {
																		F_check_for_freed_segments_locked(m, l0)
																		mBase = m.M
																		v166 = m.ExcPending
																		if v166 != 0 {
																			return int32(0)
																		} else {
																			v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
																			F_LWLockRelease(m, v167+int32(1476))
																			mBase = m.M
																			v171 = m.ExcPending
																			if v171 != 0 {
																				return int32(0)
																			} else {
																				v173 = int32(base.Ui32(v137) >> (uint(int32(27)) % 32))
																				v176 = l0 + v173*int32(20)
																				v177 = *(*int32)(unsafe.Add(mBase, uint32(v176)+12))
																				if v177 != 0 {
																					v181 = v177
																					v188 = v181 + v129&int32(134213632)
																					if l1 == int32(0) {
																						v347 = v137
																					} else {
																						base.MemoryFill(m, v188, int32(0), l1)
																						v347 = v137
																					}
																					m.G0 = v13 + int32(80)
																					return v347
																				} else {
																					v178 = F_get_segment_by_index(m, l0, v173)
																					mBase = m.M
																					v179 = m.ExcPending
																					if v179 != 0 {
																						return int32(0)
																					} else {
																						v180 = *(*int32)(unsafe.Add(mBase, uint32(v176)+12))
																						v181 = v180
																						v188 = v181 + v129&int32(134213632)
																						if l1 == int32(0) {
																							v347 = v137
																						} else {
																							base.MemoryFill(m, v188, int32(0), l1)
																							v347 = v137
																						}
																						m.G0 = v13 + int32(80)
																						return v347
																					}
																				}
																			}
																		}
																	}
																} else {
																	v173 = int32(base.Ui32(v137) >> (uint(int32(27)) % 32))
																	v176 = l0 + v173*int32(20)
																	v177 = *(*int32)(unsafe.Add(mBase, uint32(v176)+12))
																	if v177 != 0 {
																		v181 = v177
																		v188 = v181 + v129&int32(134213632)
																		if l1 == int32(0) {
																			v347 = v137
																		} else {
																			base.MemoryFill(m, v188, int32(0), l1)
																			v347 = v137
																		}
																		m.G0 = v13 + int32(80)
																		return v347
																	} else {
																		v178 = F_get_segment_by_index(m, l0, v173)
																		mBase = m.M
																		v179 = m.ExcPending
																		if v179 != 0 {
																			return int32(0)
																		} else {
																			v180 = *(*int32)(unsafe.Add(mBase, uint32(v176)+12))
																			v181 = v180
																			v188 = v181 + v129&int32(134213632)
																			if l1 == int32(0) {
																				v347 = v137
																			} else {
																				base.MemoryFill(m, v188, int32(0), l1)
																				v347 = v137
																			}
																			m.G0 = v13 + int32(80)
																			return v347
																		}
																	}
																}
															} else {
																v188 = v4
																if l1 == int32(0) {
																	v347 = v137
																} else {
																	base.MemoryFill(m, v188, int32(0), l1)
																	v347 = v137
																}
																m.G0 = v13 + int32(80)
																return v347
															}
														}
													}
												}
											}
										}
									}
								}
							} else {
								v73 = F_make_new_segment(m, l0, v70)
								mBase = m.M
								v74 = m.ExcPending
								if v74 != 0 {
									return int32(0)
								} else {
									if v73 != 0 {
										v106 = v73
										v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)+12))
										v110 = F_FreePageManagerGet(m, v107, v70, v13+int32(76))
										mBase = m.M
										v111 = m.ExcPending
										if v111 != 0 {
											return int32(0)
										} else {
											if v110 == int32(0) {
												F_errstart_cold(m, int32(22), int32(0))
												mBase = m.M
												v245 = m.ExcPending
												if v245 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v70
													F_errmsg_internal(m, int32(_a_F_dsa_allocate_extended_6), v13+int32(32))
													mBase = m.M
													v251 = m.ExcPending
													if v251 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_dsa_allocate_extended_4), int32(744), int32(_a_F_dsa_allocate_extended_5))
														mBase = m.M
														v256 = m.ExcPending
														if v256 != 0 {
															return int32(0)
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											} else {
												v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
												F_LWLockRelease(m, v114+int32(1476))
												mBase = m.M
												v118 = m.ExcPending
												if v118 != 0 {
													return int32(0)
												} else {
													v119 = *(*int32)(unsafe.Add(mBase, uint32(v13)+76))
													v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
													v124 = F_LWLockAcquire(m, v120+int32(256), int32(0))
													mBase = m.M
													v125 = m.ExcPending
													if v125 != 0 {
														return int32(0)
													} else {
														v129 = v119 << (uint(int32(12)) % 32)
														v134 = base.I32_div_s(v106-l0-int32(8), int32(20))
														v137 = v129 | v134<<(uint(int32(27))%32)
														F_init_span(m, l0, v33, v30+int32(256), v137, v70, int32(1))
														mBase = m.M
														v140 = m.ExcPending
														if v140 != 0 {
															return int32(0)
														} else {
															v141 = *(*int32)(unsafe.Add(mBase, uint32(v106)+16))
															v142 = *(*int32)(unsafe.Add(mBase, uint32(v13)+76))
															*(*int32)(unsafe.Add(mBase, uint32(v141+v142<<(uint(int32(2))%32)))) = v33
															v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
															F_LWLockRelease(m, v147+int32(256))
															mBase = m.M
															v151 = m.ExcPending
															if v151 != 0 {
																return int32(0)
															} else {
																if l2&int32(4) == int32(0) {
																	v347 = v137
																	m.G0 = v13 + int32(80)
																	return v347
																} else {
																	if v137 != 0 {
																		v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
																		v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
																		v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)+1468))
																		if v156 != v158 {
																			v163 = F_LWLockAcquire(m, v157+int32(1476), int32(0))
																			mBase = m.M
																			v164 = m.ExcPending
																			if v164 != 0 {
																				return int32(0)
																			} else {
																				F_check_for_freed_segments_locked(m, l0)
																				mBase = m.M
																				v166 = m.ExcPending
																				if v166 != 0 {
																					return int32(0)
																				} else {
																					v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
																					F_LWLockRelease(m, v167+int32(1476))
																					mBase = m.M
																					v171 = m.ExcPending
																					if v171 != 0 {
																						return int32(0)
																					} else {
																						v173 = int32(base.Ui32(v137) >> (uint(int32(27)) % 32))
																						v176 = l0 + v173*int32(20)
																						v177 = *(*int32)(unsafe.Add(mBase, uint32(v176)+12))
																						if v177 != 0 {
																							v181 = v177
																							v188 = v181 + v129&int32(134213632)
																							if l1 == int32(0) {
																								v347 = v137
																							} else {
																								base.MemoryFill(m, v188, int32(0), l1)
																								v347 = v137
																							}
																							m.G0 = v13 + int32(80)
																							return v347
																						} else {
																							v178 = F_get_segment_by_index(m, l0, v173)
																							mBase = m.M
																							v179 = m.ExcPending
																							if v179 != 0 {
																								return int32(0)
																							} else {
																								v180 = *(*int32)(unsafe.Add(mBase, uint32(v176)+12))
																								v181 = v180
																								v188 = v181 + v129&int32(134213632)
																								if l1 == int32(0) {
																									v347 = v137
																								} else {
																									base.MemoryFill(m, v188, int32(0), l1)
																									v347 = v137
																								}
																								m.G0 = v13 + int32(80)
																								return v347
																							}
																						}
																					}
																				}
																			}
																		} else {
																			v173 = int32(base.Ui32(v137) >> (uint(int32(27)) % 32))
																			v176 = l0 + v173*int32(20)
																			v177 = *(*int32)(unsafe.Add(mBase, uint32(v176)+12))
																			if v177 != 0 {
																				v181 = v177
																				v188 = v181 + v129&int32(134213632)
																				if l1 == int32(0) {
																					v347 = v137
																				} else {
																					base.MemoryFill(m, v188, int32(0), l1)
																					v347 = v137
																				}
																				m.G0 = v13 + int32(80)
																				return v347
																			} else {
																				v178 = F_get_segment_by_index(m, l0, v173)
																				mBase = m.M
																				v179 = m.ExcPending
																				if v179 != 0 {
																					return int32(0)
																				} else {
																					v180 = *(*int32)(unsafe.Add(mBase, uint32(v176)+12))
																					v181 = v180
																					v188 = v181 + v129&int32(134213632)
																					if l1 == int32(0) {
																						v347 = v137
																					} else {
																						base.MemoryFill(m, v188, int32(0), l1)
																						v347 = v137
																					}
																					m.G0 = v13 + int32(80)
																					return v347
																				}
																			}
																		}
																	} else {
																		v188 = v4
																		if l1 == int32(0) {
																			v347 = v137
																		} else {
																			base.MemoryFill(m, v188, int32(0), l1)
																			v347 = v137
																		}
																		m.G0 = v13 + int32(80)
																		return v347
																	}
																}
															}
														}
													}
												}
											}
										}
									} else {
										v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										F_LWLockRelease(m, v75+int32(1476))
										mBase = m.M
										v79 = m.ExcPending
										if v79 != 0 {
											return int32(0)
										} else {
											F_dsa_free(m, l0, v33)
											mBase = m.M
											v81 = m.ExcPending
											if v81 != 0 {
												return int32(0)
											} else {
												if l2&int32(2) != 0 {
													v347 = v31
													m.G0 = v13 + int32(80)
													return v347
												} else {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v87 = m.ExcPending
													if v87 != 0 {
														return int32(0)
													} else {
														F_errcode(m, int32(_a_F_dsa_allocate_extended_1))
														mBase = m.M
														v90 = m.ExcPending
														if v90 != 0 {
															return int32(0)
														} else {
															F_errmsg(m, int32(_a_F_dsa_allocate_extended_2), int32(0))
															mBase = m.M
															v94 = m.ExcPending
															if v94 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = l1
																F_errdetail(m, int32(_a_F_dsa_allocate_extended_3), v13+int32(16))
																mBase = m.M
																v100 = m.ExcPending
																if v100 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(_a_F_dsa_allocate_extended_4), int32(731), int32(_a_F_dsa_allocate_extended_5))
																	mBase = m.M
																	v105 = m.ExcPending
																	if v105 != 0 {
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
				}
			}
		} else {
			if base.Ui32(l1) <= base.Ui32(int32(1023)) {
				v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(l1+int32(7))>>(uint(int32(3))%32)))+uint32(_c_F_dsa_allocate_extended[0]))))
				v267 = v263
			} else {
				v200 = int32(25)
				v202 = int32(37)
				for {
					v207 = int32(_a_F_dsa_allocate_extended_7)
					v211 = v202&v207 + v200&v207
					v212 = int32(1)
					v213 = int32(base.Ui32(v211) >> (uint(v212) % 32))
					v218 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v211&int32(_a_F_dsa_allocate_extended_8))+uint32(_c_F_dsa_allocate_extended[1]))))
					v219 = base.B2i32(base.Ui32(v218) < base.Ui32(l1))
					if base.Ui32(v218) < base.Ui32(l1) {
						v220 = v213 + v212
					} else {
						v220 = v200
					}
					if base.Ui32(v218) < base.Ui32(l1) {
						v223 = v202
					} else {
						v223 = v213
					}
					if base.Ui32(v220&int32(_a_F_dsa_allocate_extended_7)) < base.Ui32(v223&int32(_a_F_dsa_allocate_extended_7)) {
						v200 = v220
						v202 = v223
						continue
					} else {
						break
					}
					break
				}
				v267 = v220
			}
			v276 = F_alloc_object(m, l0, v267&int32(_a_F_dsa_allocate_extended_7))
			mBase = m.M
			v277 = m.ExcPending
			if v277 != 0 {
				return int32(0)
			} else {
				if v276 == int32(0) {
					if l2&int32(2) != 0 {
						v347 = int32(0)
						m.G0 = v13 + int32(80)
						return v347
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v286 = m.ExcPending
						if v286 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(_a_F_dsa_allocate_extended_1))
							mBase = m.M
							v289 = m.ExcPending
							if v289 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(_a_F_dsa_allocate_extended_2), int32(0))
								mBase = m.M
								v293 = m.ExcPending
								if v293 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = l1
									F_errdetail(m, int32(_a_F_dsa_allocate_extended_3), v13+int32(48))
									mBase = m.M
									v299 = m.ExcPending
									if v299 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_dsa_allocate_extended_4), int32(811), int32(_a_F_dsa_allocate_extended_5))
										mBase = m.M
										v304 = m.ExcPending
										if v304 != 0 {
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
				} else {
					if l2&int32(4) == int32(0) {
						v347 = v276
						m.G0 = v13 + int32(80)
						return v347
					} else {
						v309 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
						v310 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v311 = *(*int32)(unsafe.Add(mBase, uint32(v310)+1468))
						if v309 != v311 {
							v316 = F_LWLockAcquire(m, v310+int32(1476), int32(0))
							mBase = m.M
							v317 = m.ExcPending
							if v317 != 0 {
								return int32(0)
							} else {
								F_check_for_freed_segments_locked(m, l0)
								mBase = m.M
								v319 = m.ExcPending
								if v319 != 0 {
									return int32(0)
								} else {
									v320 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									F_LWLockRelease(m, v320+int32(1476))
									mBase = m.M
									v324 = m.ExcPending
									if v324 != 0 {
										return int32(0)
									} else {
										v326 = int32(base.Ui32(v276) >> (uint(int32(27)) % 32))
										v329 = l0 + v326*int32(20)
										v330 = *(*int32)(unsafe.Add(mBase, uint32(v329)+12))
										if v330 == int32(0) {
											v333 = F_get_segment_by_index(m, l0, v326)
											mBase = m.M
											v334 = m.ExcPending
											if v334 != 0 {
												return int32(0)
											} else {
												v335 = *(*int32)(unsafe.Add(mBase, uint32(v329)+12))
												v336 = v335
												if l1 == int32(0) {
													v347 = v276
												} else {
													base.MemoryFill(m, v336+v276&int32(134217727), int32(0), l1)
													v347 = v276
												}
												m.G0 = v13 + int32(80)
												return v347
											}
										} else {
											v336 = v330
											if l1 == int32(0) {
												v347 = v276
											} else {
												base.MemoryFill(m, v336+v276&int32(134217727), int32(0), l1)
												v347 = v276
											}
											m.G0 = v13 + int32(80)
											return v347
										}
									}
								}
							}
						} else {
							v326 = int32(base.Ui32(v276) >> (uint(int32(27)) % 32))
							v329 = l0 + v326*int32(20)
							v330 = *(*int32)(unsafe.Add(mBase, uint32(v329)+12))
							if v330 == int32(0) {
								v333 = F_get_segment_by_index(m, l0, v326)
								mBase = m.M
								v334 = m.ExcPending
								if v334 != 0 {
									return int32(0)
								} else {
									v335 = *(*int32)(unsafe.Add(mBase, uint32(v329)+12))
									v336 = v335
									if l1 == int32(0) {
										v347 = v276
									} else {
										base.MemoryFill(m, v336+v276&int32(134217727), int32(0), l1)
										v347 = v276
									}
									m.G0 = v13 + int32(80)
									return v347
								}
							} else {
								v336 = v330
								if l1 == int32(0) {
									v347 = v276
								} else {
									base.MemoryFill(m, v336+v276&int32(134217727), int32(0), l1)
									v347 = v276
								}
								m.G0 = v13 + int32(80)
								return v347
							}
						}
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v230 = m.ExcPending
		if v230 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v13)+64)) = l1
			F_errmsg_internal(m, int32(_a_F_dsa_allocate_extended_9), v13-int32(-64))
			mBase = m.M
			v236 = m.ExcPending
			if v236 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_dsa_allocate_extended_4), int32(683), int32(_a_F_dsa_allocate_extended_5))
				mBase = m.M
				v241 = m.ExcPending
				if v241 != 0 {
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
func F_dsa_get_address(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
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
	if l1 == int32(0) {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+1468))
		if v10 != v12 {
			v17 = F_LWLockAcquire(m, v11+int32(1476), int32(0))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				F_check_for_freed_segments_locked(m, l0)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					F_LWLockRelease(m, v23+int32(1476))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						v31 = int32(base.Ui32(l1) >> (uint(int32(27)) % 32))
						v34 = l0 + v31*int32(20)
						v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
						if v35 != 0 {
							v39 = v35
							return v39 + l1&int32(134217727)
						} else {
							v36 = F_get_segment_by_index(m, l0, v31)
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return int32(0)
							} else {
								v38 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
								v39 = v38
								return v39 + l1&int32(134217727)
							}
						}
					}
				}
			}
		} else {
			v31 = int32(base.Ui32(l1) >> (uint(int32(27)) % 32))
			v34 = l0 + v31*int32(20)
			v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
			if v35 != 0 {
				v39 = v35
				return v39 + l1&int32(134217727)
			} else {
				v36 = F_get_segment_by_index(m, l0, v31)
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return int32(0)
				} else {
					v38 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
					v39 = v38
					return v39 + l1&int32(134217727)
				}
			}
		}
	}
}
