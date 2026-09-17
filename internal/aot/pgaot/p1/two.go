package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ReadTwoPhaseFile(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int64
	_ = v12
	var v15 int32
	_ = v15
	var v21 int64
	_ = v21
	var v28 int64
	_ = v28
	var v33 int64
	_ = v33
	var v36 int64
	_ = v36
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int64
	_ = v84
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int64
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v263 int32
	_ = v263
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v283 int32
	_ = v283
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v303 int32
	_ = v303
	var v308 int32
	_ = v308
	v8 = m.G0
	v10 = v8 - int32(1296)
	m.G0 = v10
	v12 = F_ReadNextFullTransactionId(m)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		if base.Ui32(l0) <= base.Ui32(int32(2)) {
			v33 = base.I64_extend_i32_u(l0)
		} else {
			v21 = int64(base.Ui64(v12) >> (uint(int64(32)) % 64))
			if base.Ui32(base.I32_wrap_i64(v12)) < base.Ui32(l0) {
				v28 = (v21 - int64(1)) & int64(4294967295)
			} else {
				v28 = v21
			}
			v33 = base.I64_extend_i32_u(l0) | v28<<(uint(int64(32))%64)
		}
		*(*uint32)(unsafe.Add(mBase, uint32(v10)+164)) = uint32(v33)
		v36 = int64(base.Ui64(v33) >> (uint(int64(32)) % 64))
		*(*uint32)(unsafe.Add(mBase, uint32(v10)+160)) = uint32(v36)
		v39 = v10 + int32(272)
		v44 = F_pg_snprintf(m, v39, int32(1024), int32(_a_F_ReadTwoPhaseFile_0), v10+int32(160))
		mBase = m.M
		v45 = m.ExcPending
		if v45 != 0 {
			return int32(0)
		} else {
			v47 = F_OpenTransientFile(m, v39, int32(0))
			mBase = m.M
			v48 = m.ExcPending
			if v48 != 0 {
				return int32(0)
			} else {
				if v47 < int32(0) {
					if l1 != 0 {
						v53 = *(*int32)(unsafe.Add(mBase, _c_F_ReadTwoPhaseFile[0]))
						if v53 == int32(44) {
							v145 = int32(0)
							m.G0 = v10 + int32(1296)
							return v145
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return int32(0)
							} else {
								F_errcode_for_file_access(m)
								mBase = m.M
								v62 = m.ExcPending
								if v62 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v10))) = v10 + int32(272)
									F_errmsg(m, int32(_a_F_ReadTwoPhaseFile_1), v10)
									mBase = m.M
									v68 = m.ExcPending
									if v68 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_ReadTwoPhaseFile_2), int32(1309), int32(_a_F_ReadTwoPhaseFile_3))
										mBase = m.M
										v73 = m.ExcPending
										if v73 != 0 {
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
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
							return int32(0)
						} else {
							F_errcode_for_file_access(m)
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10))) = v10 + int32(272)
								F_errmsg(m, int32(_a_F_ReadTwoPhaseFile_1), v10)
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_ReadTwoPhaseFile_2), int32(1309), int32(_a_F_ReadTwoPhaseFile_3))
									mBase = m.M
									v73 = m.ExcPending
									if v73 != 0 {
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
					if v47 < int32(0) {
						v79 = F___syscall_ret(m, int32(-8))
						mBase = m.M
						v83 = v79
					} else {
						v82 = F___fstatat(m, v47, int32(_a_F_ReadTwoPhaseFile_4), v10+int32(176), int32(_a_F_ReadTwoPhaseFile_5))
						mBase = m.M
						v83 = v82
					}
					if v83 != 0 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v156 = m.ExcPending
						if v156 != 0 {
							return int32(0)
						} else {
							F_errcode_for_file_access(m)
							mBase = m.M
							v158 = m.ExcPending
							if v158 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10)+144)) = v10 + int32(272)
								F_errmsg(m, int32(_a_F_ReadTwoPhaseFile_6), v10+int32(144))
								mBase = m.M
								v166 = m.ExcPending
								if v166 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_ReadTwoPhaseFile_2), int32(1321), int32(_a_F_ReadTwoPhaseFile_3))
									mBase = m.M
									v171 = m.ExcPending
									if v171 != 0 {
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
						v84 = *(*int64)(unsafe.Add(mBase, uint32(v10)+200))
						if base.Ui64(v84-int64(1073741824)) <= base.Ui64(int64(-1073741741)) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v175 = m.ExcPending
							if v175 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(16779816))
								mBase = m.M
								v178 = m.ExcPending
								if v178 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v10 + int32(272)
									*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = v84
									F_errmsg_plural(m, int32(_a_F_ReadTwoPhaseFile_7), int32(_a_F_ReadTwoPhaseFile_8), base.I32_wrap_i64(v84), v10+int32(16))
									mBase = m.M
									v189 = m.ExcPending
									if v189 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_ReadTwoPhaseFile_2), int32(1332), int32(_a_F_ReadTwoPhaseFile_3))
										mBase = m.M
										v194 = m.ExcPending
										if v194 != 0 {
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
							v89 = base.I32_wrap_i64(v84)
							v91 = v89 - int32(4)
							if v91 != (v89+int32(3))&int32(2147483640) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v198 = m.ExcPending
								if v198 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(16779816))
									mBase = m.M
									v201 = m.ExcPending
									if v201 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v10)+128)) = v10 + int32(272)
										F_errmsg(m, int32(_a_F_ReadTwoPhaseFile_9), v10+int32(128))
										mBase = m.M
										v209 = m.ExcPending
										if v209 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_ReadTwoPhaseFile_2), int32(1339), int32(_a_F_ReadTwoPhaseFile_3))
											mBase = m.M
											v214 = m.ExcPending
											if v214 != 0 {
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
								v97 = F_palloc(m, v89)
								mBase = m.M
								v98 = m.ExcPending
								if v98 != 0 {
									return int32(0)
								} else {
									v100 = *(*int32)(unsafe.Add(mBase, _c_F_ReadTwoPhaseFile[1]))
									*(*int32)(unsafe.Add(mBase, uint32(v100))) = int32(167772222)
									v103 = F_read(m, v47, v97, v89)
									mBase = m.M
									if base.I64_extend_i32_s(v103) != v84 {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v109 = m.ExcPending
										if v109 != 0 {
											return int32(0)
										} else {
											if v103 < int32(0) {
												F_errcode_for_file_access(m)
												mBase = m.M
												v216 = m.ExcPending
												if v216 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v10)+96)) = v10 + int32(272)
													F_errmsg(m, int32(_a_F_ReadTwoPhaseFile_10), v10+int32(96))
													mBase = m.M
													v224 = m.ExcPending
													if v224 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_ReadTwoPhaseFile_2), int32(1353), int32(_a_F_ReadTwoPhaseFile_3))
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
											} else {
												*(*int64)(unsafe.Add(mBase, uint32(v10)+120)) = v84
												*(*int32)(unsafe.Add(mBase, uint32(v10)+116)) = v103
												*(*int32)(unsafe.Add(mBase, uint32(v10)+112)) = v10 + int32(272)
												F_errmsg(m, int32(_a_F_ReadTwoPhaseFile_11), v10+int32(112))
												mBase = m.M
												v121 = m.ExcPending
												if v121 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_ReadTwoPhaseFile_2), int32(1357), int32(_a_F_ReadTwoPhaseFile_3))
													mBase = m.M
													v126 = m.ExcPending
													if v126 != 0 {
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
										v128 = *(*int32)(unsafe.Add(mBase, _c_F_ReadTwoPhaseFile[1]))
										*(*int32)(unsafe.Add(mBase, uint32(v128))) = int32(0)
										v131 = F_CloseTransientFile(m, v47)
										mBase = m.M
										v132 = m.ExcPending
										if v132 != 0 {
											return int32(0)
										} else {
											if v131 != 0 {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v233 = m.ExcPending
												if v233 != 0 {
													return int32(0)
												} else {
													F_errcode_for_file_access(m)
													mBase = m.M
													v235 = m.ExcPending
													if v235 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v10)+80)) = v10 + int32(272)
														F_errmsg(m, int32(_a_F_ReadTwoPhaseFile_12), v10+int32(80))
														mBase = m.M
														v243 = m.ExcPending
														if v243 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_ReadTwoPhaseFile_2), int32(1365), int32(_a_F_ReadTwoPhaseFile_3))
															mBase = m.M
															v248 = m.ExcPending
															if v248 != 0 {
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
												v133 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
												if v133 != int32(1475953972) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v252 = m.ExcPending
													if v252 != 0 {
														return int32(0)
													} else {
														F_errcode(m, int32(16779816))
														mBase = m.M
														v255 = m.ExcPending
														if v255 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v10)+64)) = v10 + int32(272)
															F_errmsg(m, int32(_a_F_ReadTwoPhaseFile_13), v10-int32(-64))
															mBase = m.M
															v263 = m.ExcPending
															if v263 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(_a_F_ReadTwoPhaseFile_2), int32(1372), int32(_a_F_ReadTwoPhaseFile_3))
																mBase = m.M
																v268 = m.ExcPending
																if v268 != 0 {
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
													v136 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v97)+4)))
													if v84 != v136 {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v272 = m.ExcPending
														if v272 != 0 {
															return int32(0)
														} else {
															F_errcode(m, int32(16779816))
															mBase = m.M
															v275 = m.ExcPending
															if v275 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = v10 + int32(272)
																F_errmsg(m, int32(_a_F_ReadTwoPhaseFile_14), v10+int32(48))
																mBase = m.M
																v283 = m.ExcPending
																if v283 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(_a_F_ReadTwoPhaseFile_2), int32(1378), int32(_a_F_ReadTwoPhaseFile_3))
																	mBase = m.M
																	v288 = m.ExcPending
																	if v288 != 0 {
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
														v138 = int32(-1)
														v139 = m.Env.Pgmem_crc32c(m, v138, v97, v91)
														mBase = m.M
														v141 = *(*int32)(unsafe.Add(mBase, uint32(v97+v91)))
														if v139^v141 != v138 {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v292 = m.ExcPending
															if v292 != 0 {
																return int32(0)
															} else {
																F_errcode(m, int32(16779816))
																mBase = m.M
																v295 = m.ExcPending
																if v295 != 0 {
																	return int32(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v10 + int32(272)
																	F_errmsg(m, int32(_a_F_ReadTwoPhaseFile_15), v10+int32(32))
																	mBase = m.M
																	v303 = m.ExcPending
																	if v303 != 0 {
																		return int32(0)
																	} else {
																		F_errfinish(m, int32(_a_F_ReadTwoPhaseFile_2), int32(1390), int32(_a_F_ReadTwoPhaseFile_3))
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
														} else {
															v145 = v97
															m.G0 = v10 + int32(1296)
															return v145
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
func F_RemoveTwoPhaseFile(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int64
	_ = v10
	var v11 int32
	_ = v11
	var v17 int64
	_ = v17
	var v24 int64
	_ = v24
	var v29 int64
	_ = v29
	var v32 int64
	_ = v32
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	v6 = m.G0
	v8 = v6 - int32(1056)
	m.G0 = v8
	v10 = F_ReadNextFullTransactionId(m)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return
	} else {
		if base.Ui32(l0) <= base.Ui32(int32(2)) {
			v29 = base.I64_extend_i32_u(l0)
		} else {
			v17 = int64(base.Ui64(v10) >> (uint(int64(32)) % 64))
			if base.Ui32(base.I32_wrap_i64(v10)) < base.Ui32(l0) {
				v24 = (v17 - int64(1)) & int64(4294967295)
			} else {
				v24 = v17
			}
			v29 = base.I64_extend_i32_u(l0) | v24<<(uint(int64(32))%64)
		}
		*(*uint32)(unsafe.Add(mBase, uint32(v8)+20)) = uint32(v29)
		v32 = int64(base.Ui64(v29) >> (uint(int64(32)) % 64))
		*(*uint32)(unsafe.Add(mBase, uint32(v8)+16)) = uint32(v32)
		v35 = v8 + int32(32)
		v40 = F_pg_snprintf(m, v35, int32(1024), int32(_a_F_RemoveTwoPhaseFile_0), v8+int32(16))
		mBase = m.M
		v41 = m.ExcPending
		if v41 != 0 {
			return
		} else {
			v42 = F_unlink(m, v35)
			mBase = m.M
			if v42 == int32(0) {
				m.G0 = v8 + int32(1056)
				return
			} else {
				if l1 == int32(0) {
					v48 = *(*int32)(unsafe.Add(mBase, _c_F_RemoveTwoPhaseFile[0]))
					if v48 == int32(44) {
						m.G0 = v8 + int32(1056)
						return
					} else {
						v53 = F_errstart(m, int32(19), int32(0))
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return
						} else {
							if v53 == int32(0) {
								m.G0 = v8 + int32(1056)
								return
							} else {
								F_errcode_for_file_access(m)
								mBase = m.M
								v58 = m.ExcPending
								if v58 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v8))) = v8 + int32(32)
									F_errmsg(m, int32(_a_F_RemoveTwoPhaseFile_1), v8)
									mBase = m.M
									v64 = m.ExcPending
									if v64 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_RemoveTwoPhaseFile_2), int32(1716), int32(_a_F_RemoveTwoPhaseFile_3))
										mBase = m.M
										v69 = m.ExcPending
										if v69 != 0 {
											return
										} else {
											m.G0 = v8 + int32(1056)
											return
										}
									}
								}
							}
						}
					}
				} else {
					v53 = F_errstart(m, int32(19), int32(0))
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return
					} else {
						if v53 == int32(0) {
							m.G0 = v8 + int32(1056)
							return
						} else {
							F_errcode_for_file_access(m)
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v8))) = v8 + int32(32)
								F_errmsg(m, int32(_a_F_RemoveTwoPhaseFile_1), v8)
								mBase = m.M
								v64 = m.ExcPending
								if v64 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_RemoveTwoPhaseFile_2), int32(1716), int32(_a_F_RemoveTwoPhaseFile_3))
									mBase = m.M
									v69 = m.ExcPending
									if v69 != 0 {
										return
									} else {
										m.G0 = v8 + int32(1056)
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
func F_UpdateTwoPhaseState(m *base.Module, l0 int32) {
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int64
	_ = v37
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	v5 = m.G0
	v7 = v5 - int32(160)
	m.G0 = v7
	v11 = F_table_open(m, int32(_a_F_UpdateTwoPhaseState_0), int32(3))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		v15 = F_SearchSysCacheCopy(m, int32(67), l0, int32(0))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			if v15 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
					F_errmsg_internal(m, int32(_a_F_UpdateTwoPhaseState_1), v7)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_UpdateTwoPhaseState_2), int32(1809), int32(_a_F_UpdateTwoPhaseState_3))
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v33 = v7 + int32(16)
				v34 = int32(0)
				base.MemoryFill(m, v33, v34, int32(72))
				v37 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v7)+104)) = v37
				*(*uint16)(unsafe.Add(mBase, uint32(v7)+144)) = uint16(v34)
				*(*int64)(unsafe.Add(mBase, uint32(v7)+136)) = v37
				*(*int64)(unsafe.Add(mBase, uint32(v7)+128)) = v37
				*(*int64)(unsafe.Add(mBase, uint32(v7)+96)) = v37
				*(*uint16)(unsafe.Add(mBase, uint32(v7)+112)) = uint16(v34)
				*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = int32(101)
				v51 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(v7)+104)) = uint8(v51)
				v53 = *(*int32)(unsafe.Add(mBase, uint32(v11)+52))
				v58 = F_heap_modify_tuple(m, v15, v53, v33, v7+int32(128), v7+int32(96))
				mBase = m.M
				v59 = m.ExcPending
				if v59 != 0 {
					return
				} else {
					F_CatalogTupleUpdate(m, v11, v58+int32(4), v58)
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return
					} else {
						F_pfree(m, v58)
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
							return
						} else {
							F_relation_close(m, v11, int32(3))
							mBase = m.M
							v68 = m.ExcPending
							if v68 != 0 {
								return
							} else {
								m.G0 = v7 + int32(160)
								return
							}
						}
					}
				}
			}
		}
	}
}
