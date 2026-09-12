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
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
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
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
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
	var v180 int32
	_ = v180
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
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
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	v4 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(80)
	m.G0 = v13
	v16 = l2 & int32(1)
	if v16&base.B2i32(l1 < v4) != 0 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v230 = m.ExcPending
		if v230 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v13)+64)) = l1
			F_errmsg_internal(m, int32(36749), v13-int32(-64))
			mBase = m.M
			v236 = m.ExcPending
			if v236 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(491521), int32(683), int32(453536))
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
	} else {
		if base.B2i32(v16 == int32(0))&base.B2i32(base.Ui32(int32(1073741824)) <= base.Ui32(l1)) != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v230 = m.ExcPending
			if v230 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v13)+64)) = l1
				F_errmsg_internal(m, int32(36749), v13-int32(-64))
				mBase = m.M
				v236 = m.ExcPending
				if v236 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(491521), int32(683), int32(453536))
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
		} else {
			if base.Ui32(int32(8193)) <= base.Ui32(l1) {
				v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v28 = int32(0)
				v30 = F_alloc_object(m, l0, v28)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int32(0)
				} else {
					if v30 == int32(0) {
						if l2&int32(2) != 0 {
							v346 = v28
							m.G0 = v13 + int32(80)
							return v346
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(8389))
								mBase = m.M
								v44 = m.ExcPending
								if v44 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(13796), int32(0))
									mBase = m.M
									v48 = m.ExcPending
									if v48 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v13))) = l1
										F_errdetail(m, int32(552728), v13)
										mBase = m.M
										v52 = m.ExcPending
										if v52 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(491521), int32(709), int32(453536))
											mBase = m.M
											v57 = m.ExcPending
											if v57 != 0 {
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
						v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v62 = F_LWLockAcquire(m, v58+int32(1476), int32(0))
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return int32(0)
						} else {
							v67 = int32(base.Ui32(l1+int32(4095)) >> (uint(int32(12)) % 32))
							v68 = F_get_best_segment(m, l0, v67)
							mBase = m.M
							v69 = m.ExcPending
							if v69 != 0 {
								return int32(0)
							} else {
								if v68 != 0 {
									v103 = v68
									v104 = *(*int32)(unsafe.Add(mBase, uint32(v103)+12))
									v107 = F_FreePageManagerGet(m, v104, v67, v13+int32(76))
									mBase = m.M
									v108 = m.ExcPending
									if v108 != 0 {
										return int32(0)
									} else {
										if v107 == int32(0) {
											F_errstart_cold(m, int32(22), int32(0))
											mBase = m.M
											v245 = m.ExcPending
											if v245 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v67
												F_errmsg_internal(m, int32(168156), v13+int32(32))
												mBase = m.M
												v251 = m.ExcPending
												if v251 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(491521), int32(744), int32(453536))
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
											v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											F_LWLockRelease(m, v111+int32(1476))
											mBase = m.M
											v115 = m.ExcPending
											if v115 != 0 {
												return int32(0)
											} else {
												v116 = *(*int32)(unsafe.Add(mBase, uint32(v13)+76))
												v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
												v121 = F_LWLockAcquire(m, v117+int32(256), int32(0))
												mBase = m.M
												v122 = m.ExcPending
												if v122 != 0 {
													return int32(0)
												} else {
													v126 = v116 << (uint(int32(12)) % 32)
													v131 = base.I32_div_s(v103-l0-int32(8), int32(20))
													v134 = v126 | v131<<(uint(int32(27))%32)
													F_init_span(m, l0, v30, v27+int32(256), v134, v67, int32(1))
													mBase = m.M
													v137 = m.ExcPending
													if v137 != 0 {
														return int32(0)
													} else {
														v138 = *(*int32)(unsafe.Add(mBase, uint32(v103)+16))
														v139 = *(*int32)(unsafe.Add(mBase, uint32(v13)+76))
														*(*int32)(unsafe.Add(mBase, uint32(v138+v139<<(uint(int32(2))%32)))) = v30
														v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
														F_LWLockRelease(m, v144+int32(256))
														mBase = m.M
														v148 = m.ExcPending
														if v148 != 0 {
															return int32(0)
														} else {
															if l2&int32(4) == int32(0) {
																v346 = v134
																m.G0 = v13 + int32(80)
																return v346
															} else {
																if v134 != 0 {
																	v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
																	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
																	v155 = *(*int32)(unsafe.Add(mBase, uint32(v154)+1468))
																	if v153 != v155 {
																		v160 = F_LWLockAcquire(m, v154+int32(1476), int32(0))
																		mBase = m.M
																		v161 = m.ExcPending
																		if v161 != 0 {
																			return int32(0)
																		} else {
																			F_check_for_freed_segments_locked(m, l0)
																			mBase = m.M
																			v163 = m.ExcPending
																			if v163 != 0 {
																				return int32(0)
																			} else {
																				v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
																				F_LWLockRelease(m, v164+int32(1476))
																				mBase = m.M
																				v168 = m.ExcPending
																				if v168 != 0 {
																					return int32(0)
																				} else {
																					v170 = int32(base.Ui32(v134) >> (uint(int32(27)) % 32))
																					v175 = l0 + v170*int32(20) + int32(12)
																					v176 = *(*int32)(unsafe.Add(mBase, uint32(v175)))
																					if v176 != 0 {
																						v180 = v176
																						v187 = v180 + v126&int32(134213632)
																						v190 = F__emscripten_memset_bulkmem(m, v187, base.I32_extend8_s(int32(0)), l1)
																						mBase = m.M
																						v346 = v134
																						m.G0 = v13 + int32(80)
																						return v346
																					} else {
																						v177 = F_get_segment_by_index(m, l0, v170)
																						mBase = m.M
																						v178 = m.ExcPending
																						if v178 != 0 {
																							return int32(0)
																						} else {
																							v179 = *(*int32)(unsafe.Add(mBase, uint32(v175)))
																							v180 = v179
																							v187 = v180 + v126&int32(134213632)
																							v190 = F__emscripten_memset_bulkmem(m, v187, base.I32_extend8_s(int32(0)), l1)
																							mBase = m.M
																							v346 = v134
																							m.G0 = v13 + int32(80)
																							return v346
																						}
																					}
																				}
																			}
																		}
																	} else {
																		v170 = int32(base.Ui32(v134) >> (uint(int32(27)) % 32))
																		v175 = l0 + v170*int32(20) + int32(12)
																		v176 = *(*int32)(unsafe.Add(mBase, uint32(v175)))
																		if v176 != 0 {
																			v180 = v176
																			v187 = v180 + v126&int32(134213632)
																			v190 = F__emscripten_memset_bulkmem(m, v187, base.I32_extend8_s(int32(0)), l1)
																			mBase = m.M
																			v346 = v134
																			m.G0 = v13 + int32(80)
																			return v346
																		} else {
																			v177 = F_get_segment_by_index(m, l0, v170)
																			mBase = m.M
																			v178 = m.ExcPending
																			if v178 != 0 {
																				return int32(0)
																			} else {
																				v179 = *(*int32)(unsafe.Add(mBase, uint32(v175)))
																				v180 = v179
																				v187 = v180 + v126&int32(134213632)
																				v190 = F__emscripten_memset_bulkmem(m, v187, base.I32_extend8_s(int32(0)), l1)
																				mBase = m.M
																				v346 = v134
																				m.G0 = v13 + int32(80)
																				return v346
																			}
																		}
																	}
																} else {
																	v187 = v4
																	v190 = F__emscripten_memset_bulkmem(m, v187, base.I32_extend8_s(int32(0)), l1)
																	mBase = m.M
																	v346 = v134
																	m.G0 = v13 + int32(80)
																	return v346
																}
															}
														}
													}
												}
											}
										}
									}
								} else {
									v70 = F_make_new_segment(m, l0, v67)
									mBase = m.M
									v71 = m.ExcPending
									if v71 != 0 {
										return int32(0)
									} else {
										if v70 != 0 {
											v103 = v70
											v104 = *(*int32)(unsafe.Add(mBase, uint32(v103)+12))
											v107 = F_FreePageManagerGet(m, v104, v67, v13+int32(76))
											mBase = m.M
											v108 = m.ExcPending
											if v108 != 0 {
												return int32(0)
											} else {
												if v107 == int32(0) {
													F_errstart_cold(m, int32(22), int32(0))
													mBase = m.M
													v245 = m.ExcPending
													if v245 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v67
														F_errmsg_internal(m, int32(168156), v13+int32(32))
														mBase = m.M
														v251 = m.ExcPending
														if v251 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(491521), int32(744), int32(453536))
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
													v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
													F_LWLockRelease(m, v111+int32(1476))
													mBase = m.M
													v115 = m.ExcPending
													if v115 != 0 {
														return int32(0)
													} else {
														v116 = *(*int32)(unsafe.Add(mBase, uint32(v13)+76))
														v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
														v121 = F_LWLockAcquire(m, v117+int32(256), int32(0))
														mBase = m.M
														v122 = m.ExcPending
														if v122 != 0 {
															return int32(0)
														} else {
															v126 = v116 << (uint(int32(12)) % 32)
															v131 = base.I32_div_s(v103-l0-int32(8), int32(20))
															v134 = v126 | v131<<(uint(int32(27))%32)
															F_init_span(m, l0, v30, v27+int32(256), v134, v67, int32(1))
															mBase = m.M
															v137 = m.ExcPending
															if v137 != 0 {
																return int32(0)
															} else {
																v138 = *(*int32)(unsafe.Add(mBase, uint32(v103)+16))
																v139 = *(*int32)(unsafe.Add(mBase, uint32(v13)+76))
																*(*int32)(unsafe.Add(mBase, uint32(v138+v139<<(uint(int32(2))%32)))) = v30
																v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
																F_LWLockRelease(m, v144+int32(256))
																mBase = m.M
																v148 = m.ExcPending
																if v148 != 0 {
																	return int32(0)
																} else {
																	if l2&int32(4) == int32(0) {
																		v346 = v134
																		m.G0 = v13 + int32(80)
																		return v346
																	} else {
																		if v134 != 0 {
																			v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
																			v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
																			v155 = *(*int32)(unsafe.Add(mBase, uint32(v154)+1468))
																			if v153 != v155 {
																				v160 = F_LWLockAcquire(m, v154+int32(1476), int32(0))
																				mBase = m.M
																				v161 = m.ExcPending
																				if v161 != 0 {
																					return int32(0)
																				} else {
																					F_check_for_freed_segments_locked(m, l0)
																					mBase = m.M
																					v163 = m.ExcPending
																					if v163 != 0 {
																						return int32(0)
																					} else {
																						v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
																						F_LWLockRelease(m, v164+int32(1476))
																						mBase = m.M
																						v168 = m.ExcPending
																						if v168 != 0 {
																							return int32(0)
																						} else {
																							v170 = int32(base.Ui32(v134) >> (uint(int32(27)) % 32))
																							v175 = l0 + v170*int32(20) + int32(12)
																							v176 = *(*int32)(unsafe.Add(mBase, uint32(v175)))
																							if v176 != 0 {
																								v180 = v176
																								v187 = v180 + v126&int32(134213632)
																								v190 = F__emscripten_memset_bulkmem(m, v187, base.I32_extend8_s(int32(0)), l1)
																								mBase = m.M
																								v346 = v134
																								m.G0 = v13 + int32(80)
																								return v346
																							} else {
																								v177 = F_get_segment_by_index(m, l0, v170)
																								mBase = m.M
																								v178 = m.ExcPending
																								if v178 != 0 {
																									return int32(0)
																								} else {
																									v179 = *(*int32)(unsafe.Add(mBase, uint32(v175)))
																									v180 = v179
																									v187 = v180 + v126&int32(134213632)
																									v190 = F__emscripten_memset_bulkmem(m, v187, base.I32_extend8_s(int32(0)), l1)
																									mBase = m.M
																									v346 = v134
																									m.G0 = v13 + int32(80)
																									return v346
																								}
																							}
																						}
																					}
																				}
																			} else {
																				v170 = int32(base.Ui32(v134) >> (uint(int32(27)) % 32))
																				v175 = l0 + v170*int32(20) + int32(12)
																				v176 = *(*int32)(unsafe.Add(mBase, uint32(v175)))
																				if v176 != 0 {
																					v180 = v176
																					v187 = v180 + v126&int32(134213632)
																					v190 = F__emscripten_memset_bulkmem(m, v187, base.I32_extend8_s(int32(0)), l1)
																					mBase = m.M
																					v346 = v134
																					m.G0 = v13 + int32(80)
																					return v346
																				} else {
																					v177 = F_get_segment_by_index(m, l0, v170)
																					mBase = m.M
																					v178 = m.ExcPending
																					if v178 != 0 {
																						return int32(0)
																					} else {
																						v179 = *(*int32)(unsafe.Add(mBase, uint32(v175)))
																						v180 = v179
																						v187 = v180 + v126&int32(134213632)
																						v190 = F__emscripten_memset_bulkmem(m, v187, base.I32_extend8_s(int32(0)), l1)
																						mBase = m.M
																						v346 = v134
																						m.G0 = v13 + int32(80)
																						return v346
																					}
																				}
																			}
																		} else {
																			v187 = v4
																			v190 = F__emscripten_memset_bulkmem(m, v187, base.I32_extend8_s(int32(0)), l1)
																			mBase = m.M
																			v346 = v134
																			m.G0 = v13 + int32(80)
																			return v346
																		}
																	}
																}
															}
														}
													}
												}
											}
										} else {
											v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											F_LWLockRelease(m, v72+int32(1476))
											mBase = m.M
											v76 = m.ExcPending
											if v76 != 0 {
												return int32(0)
											} else {
												F_dsa_free(m, l0, v30)
												mBase = m.M
												v78 = m.ExcPending
												if v78 != 0 {
													return int32(0)
												} else {
													if l2&int32(2) != 0 {
														v346 = v28
														m.G0 = v13 + int32(80)
														return v346
													} else {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v84 = m.ExcPending
														if v84 != 0 {
															return int32(0)
														} else {
															F_errcode(m, int32(8389))
															mBase = m.M
															v87 = m.ExcPending
															if v87 != 0 {
																return int32(0)
															} else {
																F_errmsg(m, int32(13796), int32(0))
																mBase = m.M
																v91 = m.ExcPending
																if v91 != 0 {
																	return int32(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = l1
																	F_errdetail(m, int32(552728), v13+int32(16))
																	mBase = m.M
																	v97 = m.ExcPending
																	if v97 != 0 {
																		return int32(0)
																	} else {
																		F_errfinish(m, int32(491521), int32(731), int32(453536))
																		mBase = m.M
																		v102 = m.ExcPending
																		if v102 != 0 {
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
					v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(l1+int32(7))>>(uint(int32(3))%32)))+uint32(_consts[979]))))
					v267 = v263
				} else {
					v198 = int32(25)
					v199 = int32(37)
					for {
						v205 = int32(65535)
						v209 = v199&v205 + v198&v205
						v210 = int32(1)
						v211 = int32(base.Ui32(v209) >> (uint(v210) % 32))
						v218 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v209&int32(131070))+uint32(_consts[980]))))
						v219 = base.B2i32(base.Ui32(v218) < base.Ui32(l1))
						if base.Ui32(v218) < base.Ui32(l1) {
							v220 = v211 + v210
						} else {
							v220 = v198
						}
						if base.Ui32(v218) < base.Ui32(l1) {
							v223 = v199
						} else {
							v223 = v211
						}
						if base.Ui32(v220&int32(65535)) < base.Ui32(v223&int32(65535)) {
							v198 = v220
							v199 = v223
							continue
						} else {
							break
						}
						break
					}
					v267 = v220
				}
				v276 = F_alloc_object(m, l0, v267&int32(65535))
				mBase = m.M
				v277 = m.ExcPending
				if v277 != 0 {
					return int32(0)
				} else {
					if v276 == int32(0) {
						if l2&int32(2) != 0 {
							v346 = int32(0)
							m.G0 = v13 + int32(80)
							return v346
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v286 = m.ExcPending
							if v286 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(8389))
								mBase = m.M
								v289 = m.ExcPending
								if v289 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(13796), int32(0))
									mBase = m.M
									v293 = m.ExcPending
									if v293 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = l1
										F_errdetail(m, int32(552728), v13+int32(48))
										mBase = m.M
										v299 = m.ExcPending
										if v299 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(491521), int32(811), int32(453536))
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
							v346 = v276
							m.G0 = v13 + int32(80)
							return v346
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
											v331 = l0 + v326*int32(20) + int32(12)
											v332 = *(*int32)(unsafe.Add(mBase, uint32(v331)))
											if v332 != 0 {
												v336 = v332
												v342 = F__emscripten_memset_bulkmem(m, v336+v276&int32(134217727), base.I32_extend8_s(int32(0)), l1)
												mBase = m.M
												v346 = v276
												m.G0 = v13 + int32(80)
												return v346
											} else {
												v333 = F_get_segment_by_index(m, l0, v326)
												mBase = m.M
												v334 = m.ExcPending
												if v334 != 0 {
													return int32(0)
												} else {
													v335 = *(*int32)(unsafe.Add(mBase, uint32(v331)))
													v336 = v335
													v342 = F__emscripten_memset_bulkmem(m, v336+v276&int32(134217727), base.I32_extend8_s(int32(0)), l1)
													mBase = m.M
													v346 = v276
													m.G0 = v13 + int32(80)
													return v346
												}
											}
										}
									}
								}
							} else {
								v326 = int32(base.Ui32(v276) >> (uint(int32(27)) % 32))
								v331 = l0 + v326*int32(20) + int32(12)
								v332 = *(*int32)(unsafe.Add(mBase, uint32(v331)))
								if v332 != 0 {
									v336 = v332
									v342 = F__emscripten_memset_bulkmem(m, v336+v276&int32(134217727), base.I32_extend8_s(int32(0)), l1)
									mBase = m.M
									v346 = v276
									m.G0 = v13 + int32(80)
									return v346
								} else {
									v333 = F_get_segment_by_index(m, l0, v326)
									mBase = m.M
									v334 = m.ExcPending
									if v334 != 0 {
										return int32(0)
									} else {
										v335 = *(*int32)(unsafe.Add(mBase, uint32(v331)))
										v336 = v335
										v342 = F__emscripten_memset_bulkmem(m, v336+v276&int32(134217727), base.I32_extend8_s(int32(0)), l1)
										mBase = m.M
										v346 = v276
										m.G0 = v13 + int32(80)
										return v346
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
						v36 = l0 + v31*int32(20) + int32(12)
						v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
						if v37 != 0 {
							v41 = v37
							return v41 + l1&int32(134217727)
						} else {
							v38 = F_get_segment_by_index(m, l0, v31)
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
								return int32(0)
							} else {
								v40 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
								v41 = v40
								return v41 + l1&int32(134217727)
							}
						}
					}
				}
			}
		} else {
			v31 = int32(base.Ui32(l1) >> (uint(int32(27)) % 32))
			v36 = l0 + v31*int32(20) + int32(12)
			v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
			if v37 != 0 {
				v41 = v37
				return v41 + l1&int32(134217727)
			} else {
				v38 = F_get_segment_by_index(m, l0, v31)
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int32(0)
				} else {
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
					v41 = v40
					return v41 + l1&int32(134217727)
				}
			}
		}
	}
}
