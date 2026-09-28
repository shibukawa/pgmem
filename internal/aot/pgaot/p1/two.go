package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ReadTwoPhaseFile(m *base.Module, l0 int64, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v1 int64
	_ = v1
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int64
	_ = v13
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
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int64
	_ = v63
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int64
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v282 int32
	_ = v282
	var v287 int32
	_ = v287
	v1 = l0
	v7 = m.G0
	v9 = v7 - int32(1296)
	m.G0 = v9
	*(*uint32)(unsafe.Add(mBase, uint32(v9)+164)) = uint32(v1)
	v13 = int64(base.Ui64(v1) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v9)+160)) = uint32(v13)
	v16 = v9 + int32(272)
	v21 = F_pg_snprintf(m, v16, int32(1024), int32(_a_F_ReadTwoPhaseFile_0), v9+int32(160))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		return int32(0)
	} else {
		v26 = F_OpenTransientFile(m, v16, int32(0))
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int32(0)
		} else {
			if v26 < int32(0) {
				if l1 != 0 {
					v32 = *(*int32)(unsafe.Add(mBase, _c_F_ReadTwoPhaseFile[0]))
					if v32 == int32(44) {
						v125 = int32(0)
						m.G0 = v9 + int32(1296)
						return v125
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int32(0)
						} else {
							F_errcode_for_file_access(m)
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v9 + int32(272)
								F_errmsg(m, int32(_a_F_ReadTwoPhaseFile_1), v9)
								mBase = m.M
								v47 = m.ExcPending
								if v47 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_ReadTwoPhaseFile_2), int32(1323), int32(_a_F_ReadTwoPhaseFile_3))
									mBase = m.M
									v52 = m.ExcPending
									if v52 != 0 {
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
					v39 = m.ExcPending
					if v39 != 0 {
						return int32(0)
					} else {
						F_errcode_for_file_access(m)
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v9 + int32(272)
							F_errmsg(m, int32(_a_F_ReadTwoPhaseFile_1), v9)
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_ReadTwoPhaseFile_2), int32(1323), int32(_a_F_ReadTwoPhaseFile_3))
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
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
				if v26 < int32(0) {
					v58 = F___syscall_ret(m, int32(-8))
					mBase = m.M
					v62 = v58
				} else {
					v61 = F___fstatat(m, v26, int32(_a_F_ReadTwoPhaseFile_4), v9+int32(176), int32(_a_F_ReadTwoPhaseFile_5))
					mBase = m.M
					v62 = v61
				}
				if v62 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v135 = m.ExcPending
					if v135 != 0 {
						return int32(0)
					} else {
						F_errcode_for_file_access(m)
						mBase = m.M
						v137 = m.ExcPending
						if v137 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9)+144)) = v9 + int32(272)
							F_errmsg(m, int32(_a_F_ReadTwoPhaseFile_6), v9+int32(144))
							mBase = m.M
							v145 = m.ExcPending
							if v145 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_ReadTwoPhaseFile_2), int32(1335), int32(_a_F_ReadTwoPhaseFile_3))
								mBase = m.M
								v150 = m.ExcPending
								if v150 != 0 {
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
					v63 = *(*int64)(unsafe.Add(mBase, uint32(v9)+200))
					if base.Ui64(v63-int64(1073741824)) <= base.Ui64(int64(-1073741741)) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v154 = m.ExcPending
						if v154 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(16779816))
							mBase = m.M
							v157 = m.ExcPending
							if v157 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v9 + int32(272)
								*(*int64)(unsafe.Add(mBase, uint32(v9)+24)) = v63
								F_errmsg_plural(m, int32(_a_F_ReadTwoPhaseFile_7), int32(_a_F_ReadTwoPhaseFile_8), base.I32_wrap_i64(v63), v9+int32(16))
								mBase = m.M
								v168 = m.ExcPending
								if v168 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_ReadTwoPhaseFile_2), int32(1346), int32(_a_F_ReadTwoPhaseFile_3))
									mBase = m.M
									v173 = m.ExcPending
									if v173 != 0 {
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
						v68 = base.I32_wrap_i64(v63)
						v70 = v68 - int32(4)
						if v70 != (v68+int32(3))&int32(2147483640) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v177 = m.ExcPending
							if v177 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(16779816))
								mBase = m.M
								v180 = m.ExcPending
								if v180 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v9)+128)) = v9 + int32(272)
									F_errmsg(m, int32(_a_F_ReadTwoPhaseFile_9), v9+int32(128))
									mBase = m.M
									v188 = m.ExcPending
									if v188 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_ReadTwoPhaseFile_2), int32(1353), int32(_a_F_ReadTwoPhaseFile_3))
										mBase = m.M
										v193 = m.ExcPending
										if v193 != 0 {
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
							v76 = F_palloc(m, v68)
							mBase = m.M
							v77 = m.ExcPending
							if v77 != 0 {
								return int32(0)
							} else {
								v79 = *(*int32)(unsafe.Add(mBase, _c_F_ReadTwoPhaseFile[1]))
								*(*int32)(unsafe.Add(mBase, uint32(v79))) = int32(167772224)
								v82 = F_read(m, v26, v76, v68)
								mBase = m.M
								if base.I64_extend_i32_s(v82) != v63 {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v88 = m.ExcPending
									if v88 != 0 {
										return int32(0)
									} else {
										if v82 < int32(0) {
											F_errcode_for_file_access(m)
											mBase = m.M
											v195 = m.ExcPending
											if v195 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v9)+96)) = v9 + int32(272)
												F_errmsg(m, int32(_a_F_ReadTwoPhaseFile_10), v9+int32(96))
												mBase = m.M
												v203 = m.ExcPending
												if v203 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_ReadTwoPhaseFile_2), int32(1367), int32(_a_F_ReadTwoPhaseFile_3))
													mBase = m.M
													v208 = m.ExcPending
													if v208 != 0 {
														return int32(0)
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											}
										} else {
											*(*int64)(unsafe.Add(mBase, uint32(v9)+120)) = v63
											*(*int32)(unsafe.Add(mBase, uint32(v9)+116)) = v82
											*(*int32)(unsafe.Add(mBase, uint32(v9)+112)) = v9 + int32(272)
											F_errmsg(m, int32(_a_F_ReadTwoPhaseFile_11), v9+int32(112))
											mBase = m.M
											v100 = m.ExcPending
											if v100 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_ReadTwoPhaseFile_2), int32(1371), int32(_a_F_ReadTwoPhaseFile_3))
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
								} else {
									v107 = *(*int32)(unsafe.Add(mBase, _c_F_ReadTwoPhaseFile[1]))
									*(*int32)(unsafe.Add(mBase, uint32(v107))) = int32(0)
									v110 = F_CloseTransientFile(m, v26)
									mBase = m.M
									v111 = m.ExcPending
									if v111 != 0 {
										return int32(0)
									} else {
										if v110 != 0 {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v212 = m.ExcPending
											if v212 != 0 {
												return int32(0)
											} else {
												F_errcode_for_file_access(m)
												mBase = m.M
												v214 = m.ExcPending
												if v214 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v9)+80)) = v9 + int32(272)
													F_errmsg(m, int32(_a_F_ReadTwoPhaseFile_12), v9+int32(80))
													mBase = m.M
													v222 = m.ExcPending
													if v222 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_ReadTwoPhaseFile_2), int32(1379), int32(_a_F_ReadTwoPhaseFile_3))
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
											v112 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
											if v112 != int32(1475953972) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v231 = m.ExcPending
												if v231 != 0 {
													return int32(0)
												} else {
													F_errcode(m, int32(16779816))
													mBase = m.M
													v234 = m.ExcPending
													if v234 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v9)+64)) = v9 + int32(272)
														F_errmsg(m, int32(_a_F_ReadTwoPhaseFile_13), v9-int32(-64))
														mBase = m.M
														v242 = m.ExcPending
														if v242 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_ReadTwoPhaseFile_2), int32(1386), int32(_a_F_ReadTwoPhaseFile_3))
															mBase = m.M
															v247 = m.ExcPending
															if v247 != 0 {
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
												v115 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v76)+4)))
												if v63 != v115 {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v251 = m.ExcPending
													if v251 != 0 {
														return int32(0)
													} else {
														F_errcode(m, int32(16779816))
														mBase = m.M
														v254 = m.ExcPending
														if v254 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v9 + int32(272)
															F_errmsg(m, int32(_a_F_ReadTwoPhaseFile_14), v9+int32(48))
															mBase = m.M
															v262 = m.ExcPending
															if v262 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(_a_F_ReadTwoPhaseFile_2), int32(1392), int32(_a_F_ReadTwoPhaseFile_3))
																mBase = m.M
																v267 = m.ExcPending
																if v267 != 0 {
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
													v117 = int32(-1)
													v118 = m.Env.Pgmem_crc32c(m, v117, v76, v70)
													mBase = m.M
													v120 = *(*int32)(unsafe.Add(mBase, uint32(v76+v70)))
													if v118^v120 != v117 {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v271 = m.ExcPending
														if v271 != 0 {
															return int32(0)
														} else {
															F_errcode(m, int32(16779816))
															mBase = m.M
															v274 = m.ExcPending
															if v274 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v9 + int32(272)
																F_errmsg(m, int32(_a_F_ReadTwoPhaseFile_15), v9+int32(32))
																mBase = m.M
																v282 = m.ExcPending
																if v282 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(_a_F_ReadTwoPhaseFile_2), int32(1404), int32(_a_F_ReadTwoPhaseFile_3))
																	mBase = m.M
																	v287 = m.ExcPending
																	if v287 != 0 {
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
														v125 = v76
														m.G0 = v9 + int32(1296)
														return v125
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
func F_RemoveTwoPhaseFile(m *base.Module, l0 int64, l1 int32) {
	mBase := m.M
	_ = mBase
	var v1 int64
	_ = v1
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int64
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	v1 = l0
	v5 = m.G0
	v7 = v5 - int32(1056)
	m.G0 = v7
	*(*uint32)(unsafe.Add(mBase, uint32(v7)+20)) = uint32(v1)
	v11 = int64(base.Ui64(v1) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v7)+16)) = uint32(v11)
	v14 = v7 + int32(32)
	v19 = F_pg_snprintf(m, v14, int32(1024), int32(_a_F_RemoveTwoPhaseFile_0), v7+int32(16))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return
	} else {
		v21 = F_unlink(m, v14)
		mBase = m.M
		if v21 == int32(0) {
			m.G0 = v7 + int32(1056)
			return
		} else {
			if l1 == int32(0) {
				v27 = *(*int32)(unsafe.Add(mBase, _c_F_RemoveTwoPhaseFile[0]))
				if v27 == int32(44) {
					m.G0 = v7 + int32(1056)
					return
				} else {
					v32 = F_errstart(m, int32(19), int32(0))
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return
					} else {
						if v32 == int32(0) {
							m.G0 = v7 + int32(1056)
							return
						} else {
							F_errcode_for_file_access(m)
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v7))) = v7 + int32(32)
								F_errmsg(m, int32(_a_F_RemoveTwoPhaseFile_1), v7)
								mBase = m.M
								v43 = m.ExcPending
								if v43 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_RemoveTwoPhaseFile_2), int32(1738), int32(_a_F_RemoveTwoPhaseFile_3))
									mBase = m.M
									v48 = m.ExcPending
									if v48 != 0 {
										return
									} else {
										m.G0 = v7 + int32(1056)
										return
									}
								}
							}
						}
					}
				}
			} else {
				v32 = F_errstart(m, int32(19), int32(0))
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return
				} else {
					if v32 == int32(0) {
						m.G0 = v7 + int32(1056)
						return
					} else {
						F_errcode_for_file_access(m)
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = v7 + int32(32)
							F_errmsg(m, int32(_a_F_RemoveTwoPhaseFile_1), v7)
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_RemoveTwoPhaseFile_2), int32(1738), int32(_a_F_RemoveTwoPhaseFile_3))
								mBase = m.M
								v48 = m.ExcPending
								if v48 != 0 {
									return
								} else {
									m.G0 = v7 + int32(1056)
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
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v38 int64
	_ = v38
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	v5 = m.G0
	v7 = v5 - int32(272)
	m.G0 = v7
	v11 = F_table_open(m, int32(_a_F_UpdateTwoPhaseState_0), int32(3))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		v16 = F_SearchSysCacheCopy(m, int32(67), base.I64_extend_i32_u(l0), int64(0))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			if v16 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
					F_errmsg_internal(m, int32(_a_F_UpdateTwoPhaseState_1), v7)
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_UpdateTwoPhaseState_2), int32(1703), int32(_a_F_UpdateTwoPhaseState_3))
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v34 = v7 + int32(16)
				base.MemoryFill(m, v34, int32(0), int32(184))
				v38 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v7)+216)) = v38
				*(*int64)(unsafe.Add(mBase, uint32(v7)+255)) = v38
				*(*int64)(unsafe.Add(mBase, uint32(v7)+248)) = v38
				*(*int64)(unsafe.Add(mBase, uint32(v7)+240)) = v38
				*(*int64)(unsafe.Add(mBase, uint32(v7)+208)) = v38
				*(*int64)(unsafe.Add(mBase, uint32(v7)+223)) = v38
				*(*int64)(unsafe.Add(mBase, uint32(v7)+80)) = int64(101)
				v52 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(v7)+216)) = uint8(v52)
				v54 = *(*int32)(unsafe.Add(mBase, uint32(v11)+52))
				v59 = F_heap_modify_tuple(m, v16, v54, v34, v7+int32(240), v7+int32(208))
				mBase = m.M
				v60 = m.ExcPending
				if v60 != 0 {
					return
				} else {
					F_CatalogTupleUpdate(m, v11, v59+int32(4), v59)
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
						return
					} else {
						F_pfree(m, v59)
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return
						} else {
							F_relation_close(m, v11, int32(3))
							mBase = m.M
							v69 = m.ExcPending
							if v69 != 0 {
								return
							} else {
								m.G0 = v7 + int32(272)
								return
							}
						}
					}
				}
			}
		}
	}
}
