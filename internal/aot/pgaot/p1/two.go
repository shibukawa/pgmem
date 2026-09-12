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
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int64
	_ = v86
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int64
	_ = v138
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v285 int32
	_ = v285
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v305 int32
	_ = v305
	var v310 int32
	_ = v310
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
		v44 = F_pg_snprintf(m, v10+int32(272), int32(1024), int32(537489), v10+int32(160))
		mBase = m.M
		v45 = m.ExcPending
		if v45 != 0 {
			return int32(0)
		} else {
			v49 = F_OpenTransientFile(m, v10+int32(272), int32(0))
			mBase = m.M
			v50 = m.ExcPending
			if v50 != 0 {
				return int32(0)
			} else {
				if v49 < int32(0) {
					if l1 != 0 {
						v55 = *(*int32)(unsafe.Add(mBase, _consts[137]))
						if v55 == int32(44) {
							v147 = int32(0)
							m.G0 = v10 + int32(1296)
							return v147
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return int32(0)
							} else {
								F_errcode_for_file_access(m)
								mBase = m.M
								v64 = m.ExcPending
								if v64 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v10))) = v10 + int32(272)
									F_errmsg(m, int32(313722), v10)
									mBase = m.M
									v70 = m.ExcPending
									if v70 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(525010), int32(1309), int32(410188))
										mBase = m.M
										v75 = m.ExcPending
										if v75 != 0 {
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
						v62 = m.ExcPending
						if v62 != 0 {
							return int32(0)
						} else {
							F_errcode_for_file_access(m)
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10))) = v10 + int32(272)
								F_errmsg(m, int32(313722), v10)
								mBase = m.M
								v70 = m.ExcPending
								if v70 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(525010), int32(1309), int32(410188))
									mBase = m.M
									v75 = m.ExcPending
									if v75 != 0 {
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
					if v49 < int32(0) {
						v81 = F___syscall_ret(m, int32(-8))
						mBase = m.M
						v85 = v81
					} else {
						v84 = F___fstatat(m, v49, int32(794587), v10+int32(176), int32(4096))
						mBase = m.M
						v85 = v84
					}
					if v85 != 0 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v158 = m.ExcPending
						if v158 != 0 {
							return int32(0)
						} else {
							F_errcode_for_file_access(m)
							mBase = m.M
							v160 = m.ExcPending
							if v160 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10)+144)) = v10 + int32(272)
								F_errmsg(m, int32(312824), v10+int32(144))
								mBase = m.M
								v168 = m.ExcPending
								if v168 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(525010), int32(1321), int32(410188))
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
						v86 = *(*int64)(unsafe.Add(mBase, uint32(v10)+200))
						if base.Ui64(v86-int64(1073741824)) <= base.Ui64(int64(-1073741741)) {
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
									*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v10 + int32(272)
									*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = v86
									F_errmsg_plural(m, int32(367274), int32(169715), base.I32_wrap_i64(v86), v10+int32(16))
									mBase = m.M
									v191 = m.ExcPending
									if v191 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(525010), int32(1332), int32(410188))
										mBase = m.M
										v196 = m.ExcPending
										if v196 != 0 {
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
							v91 = base.I32_wrap_i64(v86)
							v93 = v91 - int32(4)
							if v93 != (v91+int32(3))&int32(2147483640) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v200 = m.ExcPending
								if v200 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(16779816))
									mBase = m.M
									v203 = m.ExcPending
									if v203 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v10)+128)) = v10 + int32(272)
										F_errmsg(m, int32(752925), v10+int32(128))
										mBase = m.M
										v211 = m.ExcPending
										if v211 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(525010), int32(1339), int32(410188))
											mBase = m.M
											v216 = m.ExcPending
											if v216 != 0 {
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
								v99 = F_palloc(m, v91)
								mBase = m.M
								v100 = m.ExcPending
								if v100 != 0 {
									return int32(0)
								} else {
									v102 = *(*int32)(unsafe.Add(mBase, _consts[127]))
									*(*int32)(unsafe.Add(mBase, uint32(v102))) = int32(167772222)
									v105 = F_read(m, v49, v99, v91)
									mBase = m.M
									if base.I64_extend_i32_s(v105) != v86 {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v111 = m.ExcPending
										if v111 != 0 {
											return int32(0)
										} else {
											if v105 < int32(0) {
												F_errcode_for_file_access(m)
												mBase = m.M
												v218 = m.ExcPending
												if v218 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v10)+96)) = v10 + int32(272)
													F_errmsg(m, int32(314641), v10+int32(96))
													mBase = m.M
													v226 = m.ExcPending
													if v226 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(525010), int32(1353), int32(410188))
														mBase = m.M
														v231 = m.ExcPending
														if v231 != 0 {
															return int32(0)
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											} else {
												*(*int64)(unsafe.Add(mBase, uint32(v10)+120)) = v86
												*(*int32)(unsafe.Add(mBase, uint32(v10)+116)) = v105
												*(*int32)(unsafe.Add(mBase, uint32(v10)+112)) = v10 + int32(272)
												F_errmsg(m, int32(452073), v10+int32(112))
												mBase = m.M
												v123 = m.ExcPending
												if v123 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(525010), int32(1357), int32(410188))
													mBase = m.M
													v128 = m.ExcPending
													if v128 != 0 {
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
										v130 = *(*int32)(unsafe.Add(mBase, _consts[127]))
										*(*int32)(unsafe.Add(mBase, uint32(v130))) = int32(0)
										v133 = F_CloseTransientFile(m, v49)
										mBase = m.M
										v134 = m.ExcPending
										if v134 != 0 {
											return int32(0)
										} else {
											if v133 != 0 {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v235 = m.ExcPending
												if v235 != 0 {
													return int32(0)
												} else {
													F_errcode_for_file_access(m)
													mBase = m.M
													v237 = m.ExcPending
													if v237 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v10)+80)) = v10 + int32(272)
														F_errmsg(m, int32(314465), v10+int32(80))
														mBase = m.M
														v245 = m.ExcPending
														if v245 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(525010), int32(1365), int32(410188))
															mBase = m.M
															v250 = m.ExcPending
															if v250 != 0 {
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
												v135 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
												if v135 != int32(1475953972) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v254 = m.ExcPending
													if v254 != 0 {
														return int32(0)
													} else {
														F_errcode(m, int32(16779816))
														mBase = m.M
														v257 = m.ExcPending
														if v257 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v10)+64)) = v10 + int32(272)
															F_errmsg(m, int32(753229), v10-int32(-64))
															mBase = m.M
															v265 = m.ExcPending
															if v265 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(525010), int32(1372), int32(410188))
																mBase = m.M
																v270 = m.ExcPending
																if v270 != 0 {
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
													v138 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v99)+4)))
													if v86 != v138 {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v274 = m.ExcPending
														if v274 != 0 {
															return int32(0)
														} else {
															F_errcode(m, int32(16779816))
															mBase = m.M
															v277 = m.ExcPending
															if v277 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = v10 + int32(272)
																F_errmsg(m, int32(753270), v10+int32(48))
																mBase = m.M
																v285 = m.ExcPending
																if v285 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(525010), int32(1378), int32(410188))
																	mBase = m.M
																	v290 = m.ExcPending
																	if v290 != 0 {
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
														v140 = int32(-1)
														v141 = m.Env.Pgmem_crc32c(m, v140, v99, v93)
														mBase = m.M
														v143 = *(*int32)(unsafe.Add(mBase, uint32(v99+v93)))
														if v141^v143 != v140 {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v294 = m.ExcPending
															if v294 != 0 {
																return int32(0)
															} else {
																F_errcode(m, int32(16779816))
																mBase = m.M
																v297 = m.ExcPending
																if v297 != 0 {
																	return int32(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v10 + int32(272)
																	F_errmsg(m, int32(753303), v10+int32(32))
																	mBase = m.M
																	v305 = m.ExcPending
																	if v305 != 0 {
																		return int32(0)
																	} else {
																		F_errfinish(m, int32(525010), int32(1390), int32(410188))
																		mBase = m.M
																		v310 = m.ExcPending
																		if v310 != 0 {
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
															v147 = v99
															m.G0 = v10 + int32(1296)
															return v147
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
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
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
		v40 = F_pg_snprintf(m, v8+int32(32), int32(1024), int32(537489), v8+int32(16))
		mBase = m.M
		v41 = m.ExcPending
		if v41 != 0 {
			return
		} else {
			v44 = F_unlink(m, v8+int32(32))
			mBase = m.M
			if v44 == int32(0) {
				m.G0 = v8 + int32(1056)
				return
			} else {
				if l1 == int32(0) {
					v50 = *(*int32)(unsafe.Add(mBase, _consts[137]))
					if v50 == int32(44) {
						m.G0 = v8 + int32(1056)
						return
					} else {
						v55 = F_errstart(m, int32(19), int32(0))
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return
						} else {
							if v55 == int32(0) {
								m.G0 = v8 + int32(1056)
								return
							} else {
								F_errcode_for_file_access(m)
								mBase = m.M
								v60 = m.ExcPending
								if v60 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v8))) = v8 + int32(32)
									F_errmsg(m, int32(314276), v8)
									mBase = m.M
									v66 = m.ExcPending
									if v66 != 0 {
										return
									} else {
										F_errfinish(m, int32(525010), int32(1716), int32(410148))
										mBase = m.M
										v71 = m.ExcPending
										if v71 != 0 {
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
					v55 = F_errstart(m, int32(19), int32(0))
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return
					} else {
						if v55 == int32(0) {
							m.G0 = v8 + int32(1056)
							return
						} else {
							F_errcode_for_file_access(m)
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v8))) = v8 + int32(32)
								F_errmsg(m, int32(314276), v8)
								mBase = m.M
								v66 = m.ExcPending
								if v66 != 0 {
									return
								} else {
									F_errfinish(m, int32(525010), int32(1716), int32(410148))
									mBase = m.M
									v71 = m.ExcPending
									if v71 != 0 {
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
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int64
	_ = v42
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	v5 = m.G0
	v7 = v5 - int32(160)
	m.G0 = v7
	v11 = F_table_open(m, int32(6100), int32(3))
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
					F_errmsg_internal(m, int32(59361), v7)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return
					} else {
						F_errfinish(m, int32(526547), int32(1809), int32(372915))
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
				v37 = F__emscripten_memset_bulkmem(m, v7+int32(16), base.I32_extend8_s(int32(0)), int32(72))
				mBase = m.M
				v38 = int32(0)
				*(*uint16)(unsafe.Add(mBase, uint32(v7)+144)) = uint16(v38)
				*(*uint16)(unsafe.Add(mBase, uint32(v7)+112)) = uint16(v38)
				v42 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v7)+104)) = v42
				*(*int64)(unsafe.Add(mBase, uint32(v7)+136)) = v42
				*(*int64)(unsafe.Add(mBase, uint32(v7)+128)) = v42
				*(*int64)(unsafe.Add(mBase, uint32(v7)+96)) = v42
				*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = int32(101)
				v52 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(v7)+104)) = uint8(v52)
				v54 = *(*int32)(unsafe.Add(mBase, uint32(v11)+52))
				v61 = F_heap_modify_tuple(m, v15, v54, v7+int32(16), v7+int32(128), v7+int32(96))
				mBase = m.M
				v62 = m.ExcPending
				if v62 != 0 {
					return
				} else {
					F_CatalogTupleUpdate(m, v11, v61+int32(4), v61)
					mBase = m.M
					v66 = m.ExcPending
					if v66 != 0 {
						return
					} else {
						F_pfree(m, v61)
						mBase = m.M
						v68 = m.ExcPending
						if v68 != 0 {
							return
						} else {
							F_sequence_close(m, v11, int32(3))
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
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
