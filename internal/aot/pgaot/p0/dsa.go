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
	var v54 int32
	_ = v54
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
	var v99 int32
	_ = v99
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
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v192 int32
	_ = v192
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v234 int32
	_ = v234
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v308 int32
	_ = v308
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v355 int32
	_ = v355
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
						v355 = v31
						m.G0 = v13 + int32(80)
						return v355
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
									v54 = F_errdetail(m, int32(_a_F_dsa_allocate_extended_3), v13)
									mBase = m.M
									v55 = m.ExcPending
									if v55 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_dsa_allocate_extended_4), int32(724), int32(_a_F_dsa_allocate_extended_5))
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
										v249 = m.ExcPending
										if v249 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v70
											F_errmsg_internal(m, int32(_a_F_dsa_allocate_extended_6), v13+int32(32))
											mBase = m.M
											v255 = m.ExcPending
											if v255 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_dsa_allocate_extended_4), int32(759), int32(_a_F_dsa_allocate_extended_5))
												mBase = m.M
												v260 = m.ExcPending
												if v260 != 0 {
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
															v355 = v137
															m.G0 = v13 + int32(80)
															return v355
														} else {
															if v137 != 0 {
																v156 = int32(0)
																v159 = base.AtomicRmwOr32(m, v156, int32(_a_F_dsa_allocate_extended_7), v156)
																v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
																v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
																v162 = *(*int32)(unsafe.Add(mBase, uint32(v161)+1468))
																if v160 != v162 {
																	v167 = F_LWLockAcquire(m, v161+int32(1476), int32(0))
																	mBase = m.M
																	v168 = m.ExcPending
																	if v168 != 0 {
																		return int32(0)
																	} else {
																		F_check_for_freed_segments_locked(m, l0)
																		mBase = m.M
																		v170 = m.ExcPending
																		if v170 != 0 {
																			return int32(0)
																		} else {
																			v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
																			F_LWLockRelease(m, v171+int32(1476))
																			mBase = m.M
																			v175 = m.ExcPending
																			if v175 != 0 {
																				return int32(0)
																			} else {
																				v177 = int32(base.Ui32(v137) >> (uint(int32(27)) % 32))
																				v180 = l0 + v177*int32(20)
																				v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)+12))
																				if v181 != 0 {
																					v185 = v181
																					v192 = v185 + v129&int32(134213632)
																					if l1 == int32(0) {
																						v355 = v137
																					} else {
																						base.MemoryFill(m, v192, int32(0), l1)
																						v355 = v137
																					}
																					m.G0 = v13 + int32(80)
																					return v355
																				} else {
																					v182 = F_get_segment_by_index(m, l0, v177)
																					mBase = m.M
																					v183 = m.ExcPending
																					if v183 != 0 {
																						return int32(0)
																					} else {
																						v184 = *(*int32)(unsafe.Add(mBase, uint32(v180)+12))
																						v185 = v184
																						v192 = v185 + v129&int32(134213632)
																						if l1 == int32(0) {
																							v355 = v137
																						} else {
																							base.MemoryFill(m, v192, int32(0), l1)
																							v355 = v137
																						}
																						m.G0 = v13 + int32(80)
																						return v355
																					}
																				}
																			}
																		}
																	}
																} else {
																	v177 = int32(base.Ui32(v137) >> (uint(int32(27)) % 32))
																	v180 = l0 + v177*int32(20)
																	v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)+12))
																	if v181 != 0 {
																		v185 = v181
																		v192 = v185 + v129&int32(134213632)
																		if l1 == int32(0) {
																			v355 = v137
																		} else {
																			base.MemoryFill(m, v192, int32(0), l1)
																			v355 = v137
																		}
																		m.G0 = v13 + int32(80)
																		return v355
																	} else {
																		v182 = F_get_segment_by_index(m, l0, v177)
																		mBase = m.M
																		v183 = m.ExcPending
																		if v183 != 0 {
																			return int32(0)
																		} else {
																			v184 = *(*int32)(unsafe.Add(mBase, uint32(v180)+12))
																			v185 = v184
																			v192 = v185 + v129&int32(134213632)
																			if l1 == int32(0) {
																				v355 = v137
																			} else {
																				base.MemoryFill(m, v192, int32(0), l1)
																				v355 = v137
																			}
																			m.G0 = v13 + int32(80)
																			return v355
																		}
																	}
																}
															} else {
																v192 = v4
																if l1 == int32(0) {
																	v355 = v137
																} else {
																	base.MemoryFill(m, v192, int32(0), l1)
																	v355 = v137
																}
																m.G0 = v13 + int32(80)
																return v355
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
												v249 = m.ExcPending
												if v249 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v70
													F_errmsg_internal(m, int32(_a_F_dsa_allocate_extended_6), v13+int32(32))
													mBase = m.M
													v255 = m.ExcPending
													if v255 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_dsa_allocate_extended_4), int32(759), int32(_a_F_dsa_allocate_extended_5))
														mBase = m.M
														v260 = m.ExcPending
														if v260 != 0 {
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
																	v355 = v137
																	m.G0 = v13 + int32(80)
																	return v355
																} else {
																	if v137 != 0 {
																		v156 = int32(0)
																		v159 = base.AtomicRmwOr32(m, v156, int32(_a_F_dsa_allocate_extended_7), v156)
																		v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
																		v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
																		v162 = *(*int32)(unsafe.Add(mBase, uint32(v161)+1468))
																		if v160 != v162 {
																			v167 = F_LWLockAcquire(m, v161+int32(1476), int32(0))
																			mBase = m.M
																			v168 = m.ExcPending
																			if v168 != 0 {
																				return int32(0)
																			} else {
																				F_check_for_freed_segments_locked(m, l0)
																				mBase = m.M
																				v170 = m.ExcPending
																				if v170 != 0 {
																					return int32(0)
																				} else {
																					v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
																					F_LWLockRelease(m, v171+int32(1476))
																					mBase = m.M
																					v175 = m.ExcPending
																					if v175 != 0 {
																						return int32(0)
																					} else {
																						v177 = int32(base.Ui32(v137) >> (uint(int32(27)) % 32))
																						v180 = l0 + v177*int32(20)
																						v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)+12))
																						if v181 != 0 {
																							v185 = v181
																							v192 = v185 + v129&int32(134213632)
																							if l1 == int32(0) {
																								v355 = v137
																							} else {
																								base.MemoryFill(m, v192, int32(0), l1)
																								v355 = v137
																							}
																							m.G0 = v13 + int32(80)
																							return v355
																						} else {
																							v182 = F_get_segment_by_index(m, l0, v177)
																							mBase = m.M
																							v183 = m.ExcPending
																							if v183 != 0 {
																								return int32(0)
																							} else {
																								v184 = *(*int32)(unsafe.Add(mBase, uint32(v180)+12))
																								v185 = v184
																								v192 = v185 + v129&int32(134213632)
																								if l1 == int32(0) {
																									v355 = v137
																								} else {
																									base.MemoryFill(m, v192, int32(0), l1)
																									v355 = v137
																								}
																								m.G0 = v13 + int32(80)
																								return v355
																							}
																						}
																					}
																				}
																			}
																		} else {
																			v177 = int32(base.Ui32(v137) >> (uint(int32(27)) % 32))
																			v180 = l0 + v177*int32(20)
																			v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)+12))
																			if v181 != 0 {
																				v185 = v181
																				v192 = v185 + v129&int32(134213632)
																				if l1 == int32(0) {
																					v355 = v137
																				} else {
																					base.MemoryFill(m, v192, int32(0), l1)
																					v355 = v137
																				}
																				m.G0 = v13 + int32(80)
																				return v355
																			} else {
																				v182 = F_get_segment_by_index(m, l0, v177)
																				mBase = m.M
																				v183 = m.ExcPending
																				if v183 != 0 {
																					return int32(0)
																				} else {
																					v184 = *(*int32)(unsafe.Add(mBase, uint32(v180)+12))
																					v185 = v184
																					v192 = v185 + v129&int32(134213632)
																					if l1 == int32(0) {
																						v355 = v137
																					} else {
																						base.MemoryFill(m, v192, int32(0), l1)
																						v355 = v137
																					}
																					m.G0 = v13 + int32(80)
																					return v355
																				}
																			}
																		}
																	} else {
																		v192 = v4
																		if l1 == int32(0) {
																			v355 = v137
																		} else {
																			base.MemoryFill(m, v192, int32(0), l1)
																			v355 = v137
																		}
																		m.G0 = v13 + int32(80)
																		return v355
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
													v355 = v31
													m.G0 = v13 + int32(80)
													return v355
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
																v99 = F_errdetail(m, int32(_a_F_dsa_allocate_extended_3), v13+int32(16))
																mBase = m.M
																v100 = m.ExcPending
																if v100 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(_a_F_dsa_allocate_extended_4), int32(746), int32(_a_F_dsa_allocate_extended_5))
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
				v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(l1+int32(7))>>(uint(int32(3))%32)))+uint32(_c_F_dsa_allocate_extended[0]))))
				v271 = v267
			} else {
				v204 = int32(25)
				v206 = int32(37)
				for {
					v211 = int32(_a_F_dsa_allocate_extended_8)
					v215 = v206&v211 + v204&v211
					v216 = int32(1)
					v217 = int32(base.Ui32(v215) >> (uint(v216) % 32))
					v222 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v215&int32(_a_F_dsa_allocate_extended_9))+uint32(_c_F_dsa_allocate_extended[1]))))
					v223 = base.B2i32(base.Ui32(v222) < base.Ui32(l1))
					if base.Ui32(v222) < base.Ui32(l1) {
						v224 = v217 + v216
					} else {
						v224 = v204
					}
					if base.Ui32(v222) < base.Ui32(l1) {
						v227 = v206
					} else {
						v227 = v217
					}
					if base.Ui32(v224&int32(_a_F_dsa_allocate_extended_8)) < base.Ui32(v227&int32(_a_F_dsa_allocate_extended_8)) {
						v204 = v224
						v206 = v227
						continue
					} else {
						break
					}
					break
				}
				v271 = v224
			}
			v280 = F_alloc_object(m, l0, v271&int32(_a_F_dsa_allocate_extended_8))
			mBase = m.M
			v281 = m.ExcPending
			if v281 != 0 {
				return int32(0)
			} else {
				if v280 == int32(0) {
					if l2&int32(2) != 0 {
						v355 = int32(0)
						m.G0 = v13 + int32(80)
						return v355
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v290 = m.ExcPending
						if v290 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(_a_F_dsa_allocate_extended_1))
							mBase = m.M
							v293 = m.ExcPending
							if v293 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(_a_F_dsa_allocate_extended_2), int32(0))
								mBase = m.M
								v297 = m.ExcPending
								if v297 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = l1
									v302 = F_errdetail(m, int32(_a_F_dsa_allocate_extended_3), v13+int32(48))
									mBase = m.M
									v303 = m.ExcPending
									if v303 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_dsa_allocate_extended_4), int32(826), int32(_a_F_dsa_allocate_extended_5))
										mBase = m.M
										v308 = m.ExcPending
										if v308 != 0 {
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
						v355 = v280
						m.G0 = v13 + int32(80)
						return v355
					} else {
						v313 = int32(0)
						v316 = base.AtomicRmwOr32(m, v313, int32(_a_F_dsa_allocate_extended_7), v313)
						v317 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
						v318 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v319 = *(*int32)(unsafe.Add(mBase, uint32(v318)+1468))
						if v317 != v319 {
							v324 = F_LWLockAcquire(m, v318+int32(1476), int32(0))
							mBase = m.M
							v325 = m.ExcPending
							if v325 != 0 {
								return int32(0)
							} else {
								F_check_for_freed_segments_locked(m, l0)
								mBase = m.M
								v327 = m.ExcPending
								if v327 != 0 {
									return int32(0)
								} else {
									v328 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									F_LWLockRelease(m, v328+int32(1476))
									mBase = m.M
									v332 = m.ExcPending
									if v332 != 0 {
										return int32(0)
									} else {
										v334 = int32(base.Ui32(v280) >> (uint(int32(27)) % 32))
										v337 = l0 + v334*int32(20)
										v338 = *(*int32)(unsafe.Add(mBase, uint32(v337)+12))
										if v338 == int32(0) {
											v341 = F_get_segment_by_index(m, l0, v334)
											mBase = m.M
											v342 = m.ExcPending
											if v342 != 0 {
												return int32(0)
											} else {
												v343 = *(*int32)(unsafe.Add(mBase, uint32(v337)+12))
												v344 = v343
												if l1 == int32(0) {
													v355 = v280
												} else {
													base.MemoryFill(m, v344+v280&int32(134217727), int32(0), l1)
													v355 = v280
												}
												m.G0 = v13 + int32(80)
												return v355
											}
										} else {
											v344 = v338
											if l1 == int32(0) {
												v355 = v280
											} else {
												base.MemoryFill(m, v344+v280&int32(134217727), int32(0), l1)
												v355 = v280
											}
											m.G0 = v13 + int32(80)
											return v355
										}
									}
								}
							}
						} else {
							v334 = int32(base.Ui32(v280) >> (uint(int32(27)) % 32))
							v337 = l0 + v334*int32(20)
							v338 = *(*int32)(unsafe.Add(mBase, uint32(v337)+12))
							if v338 == int32(0) {
								v341 = F_get_segment_by_index(m, l0, v334)
								mBase = m.M
								v342 = m.ExcPending
								if v342 != 0 {
									return int32(0)
								} else {
									v343 = *(*int32)(unsafe.Add(mBase, uint32(v337)+12))
									v344 = v343
									if l1 == int32(0) {
										v355 = v280
									} else {
										base.MemoryFill(m, v344+v280&int32(134217727), int32(0), l1)
										v355 = v280
									}
									m.G0 = v13 + int32(80)
									return v355
								}
							} else {
								v344 = v338
								if l1 == int32(0) {
									v355 = v280
								} else {
									base.MemoryFill(m, v344+v280&int32(134217727), int32(0), l1)
									v355 = v280
								}
								m.G0 = v13 + int32(80)
								return v355
							}
						}
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v234 = m.ExcPending
		if v234 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v13)+64)) = l1
			F_errmsg_internal(m, int32(_a_F_dsa_allocate_extended_10), v13-int32(-64))
			mBase = m.M
			v240 = m.ExcPending
			if v240 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_dsa_allocate_extended_4), int32(698), int32(_a_F_dsa_allocate_extended_5))
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
}
func F_dsa_get_address(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
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
	var v43 int32
	_ = v43
	if l1 == int32(0) {
		return int32(0)
	} else {
		v10 = int32(0)
		v13 = base.AtomicRmwOr32(m, v10, int32(_a_F_dsa_get_address_0), v10)
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+1468))
		if v14 != v16 {
			v21 = F_LWLockAcquire(m, v15+int32(1476), int32(0))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				F_check_for_freed_segments_locked(m, l0)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					F_LWLockRelease(m, v27+int32(1476))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						v35 = int32(base.Ui32(l1) >> (uint(int32(27)) % 32))
						v38 = l0 + v35*int32(20)
						v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+12))
						if v39 != 0 {
							v43 = v39
							return v43 + l1&int32(134217727)
						} else {
							v40 = F_get_segment_by_index(m, l0, v35)
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return int32(0)
							} else {
								v42 = *(*int32)(unsafe.Add(mBase, uint32(v38)+12))
								v43 = v42
								return v43 + l1&int32(134217727)
							}
						}
					}
				}
			}
		} else {
			v35 = int32(base.Ui32(l1) >> (uint(int32(27)) % 32))
			v38 = l0 + v35*int32(20)
			v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+12))
			if v39 != 0 {
				v43 = v39
				return v43 + l1&int32(134217727)
			} else {
				v40 = F_get_segment_by_index(m, l0, v35)
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return int32(0)
				} else {
					v42 = *(*int32)(unsafe.Add(mBase, uint32(v38)+12))
					v43 = v42
					return v43 + l1&int32(134217727)
				}
			}
		}
	}
}
func F_dsa_get_total_size_from_handle(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	v2 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_dsa_get_total_size_from_handle[0]))
	if base.B2i32(v7 == v2)|base.B2i32(v7 == int32(_a_F_dsa_get_total_size_from_handle_0)) == v2 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	if v55 != 0 {
		goto L27
	} else {
		goto L28
	}
L2:
	;
	if v27 != 0 {
		goto L12
	} else {
		goto L13
	}
L3:
	;
	v16 = v7
	goto L6
L4:
	;
	goto L5
L5:
	;
	v27 = int32(0)
	goto L2
L6:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	if l0 == v17 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	goto L5
L8:
	;
	v27 = v16
	goto L2
L9:
	;
	goto L10
L10:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v19 != int32(_a_F_dsa_get_total_size_from_handle_0) {
		v16 = v19
		goto L6
	} else {
		goto L11
	}
L11:
	;
	goto L7
L12:
	;
	v28 = int32(0)
	v30 = *(*int32)(unsafe.Add(mBase, _c_F_dsa_get_total_size_from_handle[0]))
	if base.B2i32(v30 == v28)|base.B2i32(v30 == int32(_a_F_dsa_get_total_size_from_handle_0)) == v28 {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	goto L14
L14:
	;
	v51 = F_dsm_attach(m, l0)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L25
	} else {
		goto L26
	}
L15:
	;
	v55 = v50
	goto L1
L16:
	;
	v39 = v30
	goto L19
L17:
	;
	goto L18
L18:
	;
	v50 = int32(0)
	goto L15
L19:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
	if l0 == v40 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	goto L18
L21:
	;
	v50 = v39
	goto L15
L22:
	;
	goto L23
L23:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	if v42 != int32(_a_F_dsa_get_total_size_from_handle_0) {
		v39 = v42
		goto L19
	} else {
		goto L24
	}
L24:
	;
	goto L20
L25:
	;
	return int32(0)
L26:
	;
	v55 = v51
	goto L1
L27:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+24))
	v58 = v56 + int32(1476)
	v60 = F_LWLockAcquire(m, v58, int32(1))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L25
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
	v73 = m.ExcPending
	if v73 != 0 {
		goto L25
	} else {
		goto L36
	}
L30:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v56)+1448))
	F_LWLockRelease(m, v58)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L25
	} else {
		goto L31
	}
L31:
	;
	if v27 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	F_dsm_detach(m, v55)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L25
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	return v62
L35:
	;
	goto L34
L36:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L25
	} else {
		goto L37
	}
L37:
	;
	F_errmsg(m, int32(_a_F_dsa_get_total_size_from_handle_1), int32(0))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L25
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(_a_F_dsa_get_total_size_from_handle_2), int32(1074), int32(_a_F_dsa_get_total_size_from_handle_3))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L25
	} else {
		goto L39
	}
L39:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
