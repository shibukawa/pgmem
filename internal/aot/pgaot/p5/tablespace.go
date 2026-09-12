package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_create_tablespace_directories(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
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
	var v27 int32
	_ = v27
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	v6 = m.G0
	v8 = v6 - int32(272)
	m.G0 = v8
	*(*int32)(unsafe.Add(mBase, uint32(v8)+160)) = int32(509709)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+164)) = l1
	v16 = F_psprintf(m, int32(41688), v8+int32(160))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return
	} else {
		v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
		if v18 == int32(0) {
			v22 = *(*int32)(unsafe.Add(mBase, _consts[307]))
			v23 = F_mkdir(m, v16, v22)
			mBase = m.M
			if v23 < int32(0) {
				v27 = *(*int32)(unsafe.Add(mBase, _consts[87]))
				if v27 != int32(20) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v97 = m.ExcPending
					if v97 != 0 {
						return
					} else {
						F_errcode_for_file_access(m)
						mBase = m.M
						v99 = m.ExcPending
						if v99 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+96)) = v16
							F_errmsg(m, int32(309722), v8+int32(96))
							mBase = m.M
							v105 = m.ExcPending
							if v105 != 0 {
								return
							} else {
								F_errfinish(m, int32(522409), int32(594), int32(177675))
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
					*(*int32)(unsafe.Add(mBase, uint32(v8)+84)) = int32(586979)
					*(*int32)(unsafe.Add(mBase, uint32(v8)+80)) = v16
					v36 = F_psprintf(m, int32(186923), v8+int32(80))
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return
					} else {
						v111 = v36
						v116 = F___fstatat(m, int32(-100), v111, v8+int32(176), int32(0))
						mBase = m.M
						if v116 < int32(0) {
							v120 = *(*int32)(unsafe.Add(mBase, _consts[87]))
							if v120 != int32(44) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v171 = m.ExcPending
								if v171 != 0 {
									return
								} else {
									F_errcode_for_file_access(m)
									mBase = m.M
									v173 = m.ExcPending
									if v173 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v111
										F_errmsg(m, int32(309351), v8+int32(32))
										mBase = m.M
										v179 = m.ExcPending
										if v179 != 0 {
											return
										} else {
											F_errfinish(m, int32(522409), int32(634), int32(177675))
											mBase = m.M
											v184 = m.ExcPending
											if v184 != 0 {
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
								v124 = *(*int32)(unsafe.Add(mBase, _consts[307]))
								v125 = F_mkdir(m, v111, v124)
								mBase = m.M
								if int32(0) <= v125 {
									if v18 != 0 {
										v153 = int32(*(*uint8)(unsafe.Add(mBase, _consts[308])))
										if v153 == int32(1) {
											F_remove_tablespace_symlink(m, v16)
											mBase = m.M
											v157 = m.ExcPending
											if v157 != 0 {
												return
											} else {
												v158 = F_symlink(m, l0, v16)
												mBase = m.M
												if v158 < int32(0) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v224 = m.ExcPending
													if v224 != 0 {
														return
													} else {
														F_errcode_for_file_access(m)
														mBase = m.M
														v226 = m.ExcPending
														if v226 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v16
															F_errmsg(m, int32(310519), v8+int32(16))
															mBase = m.M
															v232 = m.ExcPending
															if v232 != 0 {
																return
															} else {
																F_errfinish(m, int32(522409), int32(665), int32(177675))
																mBase = m.M
																v237 = m.ExcPending
																if v237 != 0 {
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
													F_pfree(m, v16)
													mBase = m.M
													v162 = m.ExcPending
													if v162 != 0 {
														return
													} else {
														F_pfree(m, v111)
														mBase = m.M
														v164 = m.ExcPending
														if v164 != 0 {
															return
														} else {
															m.G0 = v8 + int32(272)
															return
														}
													}
												}
											}
										} else {
											v158 = F_symlink(m, l0, v16)
											mBase = m.M
											if v158 < int32(0) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v224 = m.ExcPending
												if v224 != 0 {
													return
												} else {
													F_errcode_for_file_access(m)
													mBase = m.M
													v226 = m.ExcPending
													if v226 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v16
														F_errmsg(m, int32(310519), v8+int32(16))
														mBase = m.M
														v232 = m.ExcPending
														if v232 != 0 {
															return
														} else {
															F_errfinish(m, int32(522409), int32(665), int32(177675))
															mBase = m.M
															v237 = m.ExcPending
															if v237 != 0 {
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
												F_pfree(m, v16)
												mBase = m.M
												v162 = m.ExcPending
												if v162 != 0 {
													return
												} else {
													F_pfree(m, v111)
													mBase = m.M
													v164 = m.ExcPending
													if v164 != 0 {
														return
													} else {
														m.G0 = v8 + int32(272)
														return
													}
												}
											}
										}
									} else {
										F_pfree(m, v16)
										mBase = m.M
										v162 = m.ExcPending
										if v162 != 0 {
											return
										} else {
											F_pfree(m, v111)
											mBase = m.M
											v164 = m.ExcPending
											if v164 != 0 {
												return
											} else {
												m.G0 = v8 + int32(272)
												return
											}
										}
									}
								} else {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v131 = m.ExcPending
									if v131 != 0 {
										return
									} else {
										F_errcode_for_file_access(m)
										mBase = m.M
										v133 = m.ExcPending
										if v133 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v8))) = v111
											F_errmsg(m, int32(309722), v8)
											mBase = m.M
											v137 = m.ExcPending
											if v137 != 0 {
												return
											} else {
												F_errfinish(m, int32(522409), int32(639), int32(177675))
												mBase = m.M
												v142 = m.ExcPending
												if v142 != 0 {
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
							v143 = *(*int32)(unsafe.Add(mBase, uint32(v8)+180))
							if v143&int32(61440) != int32(16384) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v188 = m.ExcPending
								if v188 != 0 {
									return
								} else {
									F_errcode(m, int32(151027844))
									mBase = m.M
									v191 = m.ExcPending
									if v191 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v8)+64)) = v111
										F_errmsg(m, int32(13592), v8-int32(-64))
										mBase = m.M
										v197 = m.ExcPending
										if v197 != 0 {
											return
										} else {
											F_errfinish(m, int32(522409), int32(645), int32(177675))
											mBase = m.M
											v202 = m.ExcPending
											if v202 != 0 {
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
								v149 = int32(*(*uint8)(unsafe.Add(mBase, _consts[308])))
								if v149 == int32(0) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v206 = m.ExcPending
									if v206 != 0 {
										return
									} else {
										F_errcode(m, int32(100663621))
										mBase = m.M
										v209 = m.ExcPending
										if v209 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = v111
											F_errmsg(m, int32(438159), v8+int32(48))
											mBase = m.M
											v215 = m.ExcPending
											if v215 != 0 {
												return
											} else {
												F_errfinish(m, int32(522409), int32(650), int32(177675))
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
									if v18 != 0 {
										v153 = int32(*(*uint8)(unsafe.Add(mBase, _consts[308])))
										if v153 == int32(1) {
											F_remove_tablespace_symlink(m, v16)
											mBase = m.M
											v157 = m.ExcPending
											if v157 != 0 {
												return
											} else {
												v158 = F_symlink(m, l0, v16)
												mBase = m.M
												if v158 < int32(0) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v224 = m.ExcPending
													if v224 != 0 {
														return
													} else {
														F_errcode_for_file_access(m)
														mBase = m.M
														v226 = m.ExcPending
														if v226 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v16
															F_errmsg(m, int32(310519), v8+int32(16))
															mBase = m.M
															v232 = m.ExcPending
															if v232 != 0 {
																return
															} else {
																F_errfinish(m, int32(522409), int32(665), int32(177675))
																mBase = m.M
																v237 = m.ExcPending
																if v237 != 0 {
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
													F_pfree(m, v16)
													mBase = m.M
													v162 = m.ExcPending
													if v162 != 0 {
														return
													} else {
														F_pfree(m, v111)
														mBase = m.M
														v164 = m.ExcPending
														if v164 != 0 {
															return
														} else {
															m.G0 = v8 + int32(272)
															return
														}
													}
												}
											}
										} else {
											v158 = F_symlink(m, l0, v16)
											mBase = m.M
											if v158 < int32(0) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v224 = m.ExcPending
												if v224 != 0 {
													return
												} else {
													F_errcode_for_file_access(m)
													mBase = m.M
													v226 = m.ExcPending
													if v226 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v16
														F_errmsg(m, int32(310519), v8+int32(16))
														mBase = m.M
														v232 = m.ExcPending
														if v232 != 0 {
															return
														} else {
															F_errfinish(m, int32(522409), int32(665), int32(177675))
															mBase = m.M
															v237 = m.ExcPending
															if v237 != 0 {
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
												F_pfree(m, v16)
												mBase = m.M
												v162 = m.ExcPending
												if v162 != 0 {
													return
												} else {
													F_pfree(m, v111)
													mBase = m.M
													v164 = m.ExcPending
													if v164 != 0 {
														return
													} else {
														m.G0 = v8 + int32(272)
														return
													}
												}
											}
										}
									} else {
										F_pfree(m, v16)
										mBase = m.M
										v162 = m.ExcPending
										if v162 != 0 {
											return
										} else {
											F_pfree(m, v111)
											mBase = m.M
											v164 = m.ExcPending
											if v164 != 0 {
												return
											} else {
												m.G0 = v8 + int32(272)
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
				*(*int32)(unsafe.Add(mBase, uint32(v8)+84)) = int32(586979)
				*(*int32)(unsafe.Add(mBase, uint32(v8)+80)) = v16
				v36 = F_psprintf(m, int32(186923), v8+int32(80))
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return
				} else {
					v111 = v36
					v116 = F___fstatat(m, int32(-100), v111, v8+int32(176), int32(0))
					mBase = m.M
					if v116 < int32(0) {
						v120 = *(*int32)(unsafe.Add(mBase, _consts[87]))
						if v120 != int32(44) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v171 = m.ExcPending
							if v171 != 0 {
								return
							} else {
								F_errcode_for_file_access(m)
								mBase = m.M
								v173 = m.ExcPending
								if v173 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v111
									F_errmsg(m, int32(309351), v8+int32(32))
									mBase = m.M
									v179 = m.ExcPending
									if v179 != 0 {
										return
									} else {
										F_errfinish(m, int32(522409), int32(634), int32(177675))
										mBase = m.M
										v184 = m.ExcPending
										if v184 != 0 {
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
							v124 = *(*int32)(unsafe.Add(mBase, _consts[307]))
							v125 = F_mkdir(m, v111, v124)
							mBase = m.M
							if int32(0) <= v125 {
								if v18 != 0 {
									v153 = int32(*(*uint8)(unsafe.Add(mBase, _consts[308])))
									if v153 == int32(1) {
										F_remove_tablespace_symlink(m, v16)
										mBase = m.M
										v157 = m.ExcPending
										if v157 != 0 {
											return
										} else {
											v158 = F_symlink(m, l0, v16)
											mBase = m.M
											if v158 < int32(0) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v224 = m.ExcPending
												if v224 != 0 {
													return
												} else {
													F_errcode_for_file_access(m)
													mBase = m.M
													v226 = m.ExcPending
													if v226 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v16
														F_errmsg(m, int32(310519), v8+int32(16))
														mBase = m.M
														v232 = m.ExcPending
														if v232 != 0 {
															return
														} else {
															F_errfinish(m, int32(522409), int32(665), int32(177675))
															mBase = m.M
															v237 = m.ExcPending
															if v237 != 0 {
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
												F_pfree(m, v16)
												mBase = m.M
												v162 = m.ExcPending
												if v162 != 0 {
													return
												} else {
													F_pfree(m, v111)
													mBase = m.M
													v164 = m.ExcPending
													if v164 != 0 {
														return
													} else {
														m.G0 = v8 + int32(272)
														return
													}
												}
											}
										}
									} else {
										v158 = F_symlink(m, l0, v16)
										mBase = m.M
										if v158 < int32(0) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v224 = m.ExcPending
											if v224 != 0 {
												return
											} else {
												F_errcode_for_file_access(m)
												mBase = m.M
												v226 = m.ExcPending
												if v226 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v16
													F_errmsg(m, int32(310519), v8+int32(16))
													mBase = m.M
													v232 = m.ExcPending
													if v232 != 0 {
														return
													} else {
														F_errfinish(m, int32(522409), int32(665), int32(177675))
														mBase = m.M
														v237 = m.ExcPending
														if v237 != 0 {
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
											F_pfree(m, v16)
											mBase = m.M
											v162 = m.ExcPending
											if v162 != 0 {
												return
											} else {
												F_pfree(m, v111)
												mBase = m.M
												v164 = m.ExcPending
												if v164 != 0 {
													return
												} else {
													m.G0 = v8 + int32(272)
													return
												}
											}
										}
									}
								} else {
									F_pfree(m, v16)
									mBase = m.M
									v162 = m.ExcPending
									if v162 != 0 {
										return
									} else {
										F_pfree(m, v111)
										mBase = m.M
										v164 = m.ExcPending
										if v164 != 0 {
											return
										} else {
											m.G0 = v8 + int32(272)
											return
										}
									}
								}
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v131 = m.ExcPending
								if v131 != 0 {
									return
								} else {
									F_errcode_for_file_access(m)
									mBase = m.M
									v133 = m.ExcPending
									if v133 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v8))) = v111
										F_errmsg(m, int32(309722), v8)
										mBase = m.M
										v137 = m.ExcPending
										if v137 != 0 {
											return
										} else {
											F_errfinish(m, int32(522409), int32(639), int32(177675))
											mBase = m.M
											v142 = m.ExcPending
											if v142 != 0 {
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
						v143 = *(*int32)(unsafe.Add(mBase, uint32(v8)+180))
						if v143&int32(61440) != int32(16384) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v188 = m.ExcPending
							if v188 != 0 {
								return
							} else {
								F_errcode(m, int32(151027844))
								mBase = m.M
								v191 = m.ExcPending
								if v191 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v8)+64)) = v111
									F_errmsg(m, int32(13592), v8-int32(-64))
									mBase = m.M
									v197 = m.ExcPending
									if v197 != 0 {
										return
									} else {
										F_errfinish(m, int32(522409), int32(645), int32(177675))
										mBase = m.M
										v202 = m.ExcPending
										if v202 != 0 {
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
							v149 = int32(*(*uint8)(unsafe.Add(mBase, _consts[308])))
							if v149 == int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v206 = m.ExcPending
								if v206 != 0 {
									return
								} else {
									F_errcode(m, int32(100663621))
									mBase = m.M
									v209 = m.ExcPending
									if v209 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = v111
										F_errmsg(m, int32(438159), v8+int32(48))
										mBase = m.M
										v215 = m.ExcPending
										if v215 != 0 {
											return
										} else {
											F_errfinish(m, int32(522409), int32(650), int32(177675))
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
								if v18 != 0 {
									v153 = int32(*(*uint8)(unsafe.Add(mBase, _consts[308])))
									if v153 == int32(1) {
										F_remove_tablespace_symlink(m, v16)
										mBase = m.M
										v157 = m.ExcPending
										if v157 != 0 {
											return
										} else {
											v158 = F_symlink(m, l0, v16)
											mBase = m.M
											if v158 < int32(0) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v224 = m.ExcPending
												if v224 != 0 {
													return
												} else {
													F_errcode_for_file_access(m)
													mBase = m.M
													v226 = m.ExcPending
													if v226 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v16
														F_errmsg(m, int32(310519), v8+int32(16))
														mBase = m.M
														v232 = m.ExcPending
														if v232 != 0 {
															return
														} else {
															F_errfinish(m, int32(522409), int32(665), int32(177675))
															mBase = m.M
															v237 = m.ExcPending
															if v237 != 0 {
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
												F_pfree(m, v16)
												mBase = m.M
												v162 = m.ExcPending
												if v162 != 0 {
													return
												} else {
													F_pfree(m, v111)
													mBase = m.M
													v164 = m.ExcPending
													if v164 != 0 {
														return
													} else {
														m.G0 = v8 + int32(272)
														return
													}
												}
											}
										}
									} else {
										v158 = F_symlink(m, l0, v16)
										mBase = m.M
										if v158 < int32(0) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v224 = m.ExcPending
											if v224 != 0 {
												return
											} else {
												F_errcode_for_file_access(m)
												mBase = m.M
												v226 = m.ExcPending
												if v226 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v16
													F_errmsg(m, int32(310519), v8+int32(16))
													mBase = m.M
													v232 = m.ExcPending
													if v232 != 0 {
														return
													} else {
														F_errfinish(m, int32(522409), int32(665), int32(177675))
														mBase = m.M
														v237 = m.ExcPending
														if v237 != 0 {
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
											F_pfree(m, v16)
											mBase = m.M
											v162 = m.ExcPending
											if v162 != 0 {
												return
											} else {
												F_pfree(m, v111)
												mBase = m.M
												v164 = m.ExcPending
												if v164 != 0 {
													return
												} else {
													m.G0 = v8 + int32(272)
													return
												}
											}
										}
									}
								} else {
									F_pfree(m, v16)
									mBase = m.M
									v162 = m.ExcPending
									if v162 != 0 {
										return
									} else {
										F_pfree(m, v111)
										mBase = m.M
										v164 = m.ExcPending
										if v164 != 0 {
											return
										} else {
											m.G0 = v8 + int32(272)
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
			*(*int32)(unsafe.Add(mBase, uint32(v8)+148)) = int32(586979)
			*(*int32)(unsafe.Add(mBase, uint32(v8)+144)) = l0
			v44 = F_psprintf(m, int32(186923), v8+int32(144))
			mBase = m.M
			v45 = m.ExcPending
			if v45 != 0 {
				return
			} else {
				v47 = *(*int32)(unsafe.Add(mBase, _consts[307]))
				v48 = F_chmod(m, l0, v47)
				mBase = m.M
				if v48 == int32(0) {
					v111 = v44
					v116 = F___fstatat(m, int32(-100), v111, v8+int32(176), int32(0))
					mBase = m.M
					if v116 < int32(0) {
						v120 = *(*int32)(unsafe.Add(mBase, _consts[87]))
						if v120 != int32(44) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v171 = m.ExcPending
							if v171 != 0 {
								return
							} else {
								F_errcode_for_file_access(m)
								mBase = m.M
								v173 = m.ExcPending
								if v173 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v111
									F_errmsg(m, int32(309351), v8+int32(32))
									mBase = m.M
									v179 = m.ExcPending
									if v179 != 0 {
										return
									} else {
										F_errfinish(m, int32(522409), int32(634), int32(177675))
										mBase = m.M
										v184 = m.ExcPending
										if v184 != 0 {
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
							v124 = *(*int32)(unsafe.Add(mBase, _consts[307]))
							v125 = F_mkdir(m, v111, v124)
							mBase = m.M
							if int32(0) <= v125 {
								if v18 != 0 {
									v153 = int32(*(*uint8)(unsafe.Add(mBase, _consts[308])))
									if v153 == int32(1) {
										F_remove_tablespace_symlink(m, v16)
										mBase = m.M
										v157 = m.ExcPending
										if v157 != 0 {
											return
										} else {
											v158 = F_symlink(m, l0, v16)
											mBase = m.M
											if v158 < int32(0) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v224 = m.ExcPending
												if v224 != 0 {
													return
												} else {
													F_errcode_for_file_access(m)
													mBase = m.M
													v226 = m.ExcPending
													if v226 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v16
														F_errmsg(m, int32(310519), v8+int32(16))
														mBase = m.M
														v232 = m.ExcPending
														if v232 != 0 {
															return
														} else {
															F_errfinish(m, int32(522409), int32(665), int32(177675))
															mBase = m.M
															v237 = m.ExcPending
															if v237 != 0 {
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
												F_pfree(m, v16)
												mBase = m.M
												v162 = m.ExcPending
												if v162 != 0 {
													return
												} else {
													F_pfree(m, v111)
													mBase = m.M
													v164 = m.ExcPending
													if v164 != 0 {
														return
													} else {
														m.G0 = v8 + int32(272)
														return
													}
												}
											}
										}
									} else {
										v158 = F_symlink(m, l0, v16)
										mBase = m.M
										if v158 < int32(0) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v224 = m.ExcPending
											if v224 != 0 {
												return
											} else {
												F_errcode_for_file_access(m)
												mBase = m.M
												v226 = m.ExcPending
												if v226 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v16
													F_errmsg(m, int32(310519), v8+int32(16))
													mBase = m.M
													v232 = m.ExcPending
													if v232 != 0 {
														return
													} else {
														F_errfinish(m, int32(522409), int32(665), int32(177675))
														mBase = m.M
														v237 = m.ExcPending
														if v237 != 0 {
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
											F_pfree(m, v16)
											mBase = m.M
											v162 = m.ExcPending
											if v162 != 0 {
												return
											} else {
												F_pfree(m, v111)
												mBase = m.M
												v164 = m.ExcPending
												if v164 != 0 {
													return
												} else {
													m.G0 = v8 + int32(272)
													return
												}
											}
										}
									}
								} else {
									F_pfree(m, v16)
									mBase = m.M
									v162 = m.ExcPending
									if v162 != 0 {
										return
									} else {
										F_pfree(m, v111)
										mBase = m.M
										v164 = m.ExcPending
										if v164 != 0 {
											return
										} else {
											m.G0 = v8 + int32(272)
											return
										}
									}
								}
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v131 = m.ExcPending
								if v131 != 0 {
									return
								} else {
									F_errcode_for_file_access(m)
									mBase = m.M
									v133 = m.ExcPending
									if v133 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v8))) = v111
										F_errmsg(m, int32(309722), v8)
										mBase = m.M
										v137 = m.ExcPending
										if v137 != 0 {
											return
										} else {
											F_errfinish(m, int32(522409), int32(639), int32(177675))
											mBase = m.M
											v142 = m.ExcPending
											if v142 != 0 {
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
						v143 = *(*int32)(unsafe.Add(mBase, uint32(v8)+180))
						if v143&int32(61440) != int32(16384) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v188 = m.ExcPending
							if v188 != 0 {
								return
							} else {
								F_errcode(m, int32(151027844))
								mBase = m.M
								v191 = m.ExcPending
								if v191 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v8)+64)) = v111
									F_errmsg(m, int32(13592), v8-int32(-64))
									mBase = m.M
									v197 = m.ExcPending
									if v197 != 0 {
										return
									} else {
										F_errfinish(m, int32(522409), int32(645), int32(177675))
										mBase = m.M
										v202 = m.ExcPending
										if v202 != 0 {
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
							v149 = int32(*(*uint8)(unsafe.Add(mBase, _consts[308])))
							if v149 == int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v206 = m.ExcPending
								if v206 != 0 {
									return
								} else {
									F_errcode(m, int32(100663621))
									mBase = m.M
									v209 = m.ExcPending
									if v209 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = v111
										F_errmsg(m, int32(438159), v8+int32(48))
										mBase = m.M
										v215 = m.ExcPending
										if v215 != 0 {
											return
										} else {
											F_errfinish(m, int32(522409), int32(650), int32(177675))
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
								if v18 != 0 {
									v153 = int32(*(*uint8)(unsafe.Add(mBase, _consts[308])))
									if v153 == int32(1) {
										F_remove_tablespace_symlink(m, v16)
										mBase = m.M
										v157 = m.ExcPending
										if v157 != 0 {
											return
										} else {
											v158 = F_symlink(m, l0, v16)
											mBase = m.M
											if v158 < int32(0) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v224 = m.ExcPending
												if v224 != 0 {
													return
												} else {
													F_errcode_for_file_access(m)
													mBase = m.M
													v226 = m.ExcPending
													if v226 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v16
														F_errmsg(m, int32(310519), v8+int32(16))
														mBase = m.M
														v232 = m.ExcPending
														if v232 != 0 {
															return
														} else {
															F_errfinish(m, int32(522409), int32(665), int32(177675))
															mBase = m.M
															v237 = m.ExcPending
															if v237 != 0 {
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
												F_pfree(m, v16)
												mBase = m.M
												v162 = m.ExcPending
												if v162 != 0 {
													return
												} else {
													F_pfree(m, v111)
													mBase = m.M
													v164 = m.ExcPending
													if v164 != 0 {
														return
													} else {
														m.G0 = v8 + int32(272)
														return
													}
												}
											}
										}
									} else {
										v158 = F_symlink(m, l0, v16)
										mBase = m.M
										if v158 < int32(0) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v224 = m.ExcPending
											if v224 != 0 {
												return
											} else {
												F_errcode_for_file_access(m)
												mBase = m.M
												v226 = m.ExcPending
												if v226 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v16
													F_errmsg(m, int32(310519), v8+int32(16))
													mBase = m.M
													v232 = m.ExcPending
													if v232 != 0 {
														return
													} else {
														F_errfinish(m, int32(522409), int32(665), int32(177675))
														mBase = m.M
														v237 = m.ExcPending
														if v237 != 0 {
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
											F_pfree(m, v16)
											mBase = m.M
											v162 = m.ExcPending
											if v162 != 0 {
												return
											} else {
												F_pfree(m, v111)
												mBase = m.M
												v164 = m.ExcPending
												if v164 != 0 {
													return
												} else {
													m.G0 = v8 + int32(272)
													return
												}
											}
										}
									}
								} else {
									F_pfree(m, v16)
									mBase = m.M
									v162 = m.ExcPending
									if v162 != 0 {
										return
									} else {
										F_pfree(m, v111)
										mBase = m.M
										v164 = m.ExcPending
										if v164 != 0 {
											return
										} else {
											m.G0 = v8 + int32(272)
											return
										}
									}
								}
							}
						}
					}
				} else {
					v52 = *(*int32)(unsafe.Add(mBase, _consts[87]))
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return
					} else {
						if v52 == int32(44) {
							F_errcode(m, int32(16908805))
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v8)+112)) = l0
								F_errmsg(m, int32(76198), v8+int32(112))
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return
								} else {
									v69 = int32(*(*uint8)(unsafe.Add(mBase, _consts[308])))
									if v69 == int32(1) {
										F_errhint(m, int32(635108), int32(0))
										mBase = m.M
										v75 = m.ExcPending
										if v75 != 0 {
											return
										} else {
											F_errfinish(m, int32(522409), int32(613), int32(177675))
											mBase = m.M
											v80 = m.ExcPending
											if v80 != 0 {
												return
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									} else {
										F_errfinish(m, int32(522409), int32(613), int32(177675))
										mBase = m.M
										v80 = m.ExcPending
										if v80 != 0 {
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
							F_errcode_for_file_access(m)
							mBase = m.M
							v82 = m.ExcPending
							if v82 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v8)+128)) = l0
								F_errmsg(m, int32(309511), v8+int32(128))
								mBase = m.M
								v88 = m.ExcPending
								if v88 != 0 {
									return
								} else {
									F_errfinish(m, int32(522409), int32(618), int32(177675))
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
					}
				}
			}
		}
	}
}
func F_destroy_tablespace_directories(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
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
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v252 int32
	_ = v252
	var v258 int32
	_ = v258
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v288 int32
	_ = v288
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v316 int32
	_ = v316
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v390 int32
	_ = v390
	var v396 int32
	_ = v396
	v3 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(208)
	m.G0 = v12
	*(*int32)(unsafe.Add(mBase, uint32(v12)+96)) = int32(509709)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+100)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v12)+104)) = int32(586979)
	v22 = F_psprintf(m, int32(186797), v12+int32(96))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_pfree(m, v386)
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L4
	} else {
		goto L156
	}
L2:
	;
	v233 = F_pstrdup(m, v22)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L4
	} else {
		goto L81
	}
L3:
	;
	v72 = F_ReadDir(m, v26, v22)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L4
	} else {
		goto L25
	}
L4:
	;
	return int32(0)
L5:
	;
	v26 = F_AllocateDir(m, v22)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	if v26 != 0 {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	v29 = *(*int32)(unsafe.Add(mBase, _consts[87]))
	if v29 == int32(44) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	if l1 != 0 {
		goto L2
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	if l1 == int32(0) {
		goto L3
	} else {
		goto L17
	}
L11:
	;
	v34 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L4
	} else {
		goto L12
	}
L12:
	;
	if v34 == int32(0) {
		goto L2
	} else {
		goto L13
	}
L13:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v22
	F_errmsg(m, int32(309559), v12+int32(16))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L4
	} else {
		goto L15
	}
L15:
	;
	F_errfinish(m, int32(522409), int32(729), int32(177644))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L4
	} else {
		goto L16
	}
L16:
	;
	goto L2
L17:
	;
	v55 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L4
	} else {
		goto L18
	}
L18:
	;
	if v55 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v386 = v22
	v390 = v3
	goto L1
L20:
	;
	goto L21
L21:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v22
	F_errmsg(m, int32(309559), v12+int32(32))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(522409), int32(739), int32(177644))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L4
	} else {
		goto L24
	}
L24:
	;
	v386 = v22
	v390 = v3
	goto L1
L25:
	;
	if v72 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	if l1 != 0 {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	goto L28
L28:
	;
	F_FreeDir(m, v26)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L4
	} else {
		goto L69
	}
L29:
	;
	v76 = int32(15)
	goto L31
L30:
	;
	v76 = int32(21)
	goto L31
L31:
	;
	v77 = v72
	goto L32
L32:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+19)))
	if v86 != int32(46) {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	goto L28
L34:
	;
	v186 = F_ReadDir(m, v26, v22)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L4
	} else {
		goto L67
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+84)) = v77 + int32(19)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+80)) = v22
	v105 = F_psprintf(m, int32(186923), v12+int32(80))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L4
	} else {
		goto L40
	}
L36:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+20)))
	if v89 == int32(0) {
		goto L34
	} else {
		goto L37
	}
L37:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+20)))
	if v92 != int32(46) {
		goto L35
	} else {
		goto L38
	}
L38:
	;
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+21)))
	if v95 == int32(0) {
		goto L34
	} else {
		goto L39
	}
L39:
	;
	goto L35
L40:
	;
	if l1 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v109 = F_AllocateDir(m, v105)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L4
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v154 = F_rmdir(m, v105)
	mBase = m.M
	if int32(0) <= v154 {
		goto L59
	} else {
		goto L60
	}
L44:
	;
	goto L46
L45:
	;
	F_FreeDir(m, v109)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L4
	} else {
		goto L58
	}
L46:
	;
	v120 = F_ReadDir(m, v109, v105)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L4
	} else {
		goto L48
	}
L47:
	;
	F_FreeDir(m, v109)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L4
	} else {
		goto L55
	}
L48:
	;
	if v120 == int32(0) {
		goto L45
	} else {
		goto L49
	}
L49:
	;
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120)+19)))
	if v124 != int32(46) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	goto L47
L51:
	;
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120)+20)))
	if v127 == int32(0) {
		goto L46
	} else {
		goto L52
	}
L52:
	;
	if v127 != int32(46) {
		goto L50
	} else {
		goto L53
	}
L53:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120)+21)))
	if v132 == int32(0) {
		goto L46
	} else {
		goto L54
	}
L54:
	;
	goto L50
L55:
	;
	F_FreeDir(m, v26)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L4
	} else {
		goto L56
	}
L56:
	;
	F_pfree(m, v105)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L4
	} else {
		goto L57
	}
L57:
	;
	v386 = v22
	v390 = int32(0)
	goto L1
L58:
	;
	goto L43
L59:
	;
	F_pfree(m, v105)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L4
	} else {
		goto L66
	}
L60:
	;
	v158 = F_errstart(m, v76, int32(0))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L4
	} else {
		goto L61
	}
L61:
	;
	if v158 == int32(0) {
		goto L59
	} else {
		goto L62
	}
L62:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L4
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+64)) = v105
	F_errmsg(m, int32(309686), v12-int32(-64))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L4
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(522409), int32(768), int32(177644))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L4
	} else {
		goto L65
	}
L65:
	;
	goto L59
L66:
	;
	goto L34
L67:
	;
	if v186 != 0 {
		v77 = v186
		goto L32
	} else {
		goto L68
	}
L68:
	;
	goto L33
L69:
	;
	v199 = F_rmdir(m, v22)
	mBase = m.M
	if int32(0) <= v199 {
		goto L2
	} else {
		goto L70
	}
L70:
	;
	v202 = int32(0)
	if l1 != 0 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v205 = int32(15)
	goto L73
L72:
	;
	v205 = int32(21)
	goto L73
L73:
	;
	v207 = F_errstart(m, v205, int32(0))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L4
	} else {
		goto L74
	}
L74:
	;
	if v207 == int32(0) {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v386 = v22
	v390 = v202
	goto L1
L76:
	;
	goto L77
L77:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L4
	} else {
		goto L78
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v22
	F_errmsg(m, int32(309686), v12+int32(48))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L4
	} else {
		goto L79
	}
L79:
	;
	F_errfinish(m, int32(522409), int32(781), int32(177644))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L4
	} else {
		goto L80
	}
L80:
	;
	v386 = v22
	v390 = v202
	goto L1
L81:
	;
	v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233))))
	if v238 != 0 {
		goto L83
	} else {
		goto L84
	}
L82:
	;
	v288 = F___fstatat(m, int32(-100), v233, v12+int32(112), int32(256))
	mBase = m.M
	goto L109
L83:
	;
	v239 = F_strlen(m, v233)
	mBase = m.M
	v242 = v239 + v233
	goto L86
L84:
	;
	goto L85
L85:
	;
	goto L82
L86:
	;
	v246 = v242 - int32(1)
	v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v246))))
	if v247 == int32(47) {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	v252 = v246
	goto L92
L88:
	;
	if base.Ui32(v233) < base.Ui32(v246) {
		v242 = v246
		goto L86
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	goto L87
L91:
	;
	goto L90
L92:
	;
	if base.Ui32(v233) < base.Ui32(v252) {
		goto L94
	} else {
		goto L95
	}
L93:
	;
	v264 = v252
	goto L98
L94:
	;
	v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v252))))
	if v258 != int32(47) {
		v252 = v252 - int32(1)
		goto L92
	} else {
		goto L97
	}
L95:
	;
	goto L96
L96:
	;
	goto L93
L97:
	;
	goto L96
L98:
	;
	if base.Ui32(v233) < base.Ui32(v264) {
		goto L100
	} else {
		goto L101
	}
L99:
	;
	if v233 == v264 {
		goto L104
	} else {
		goto L105
	}
L100:
	;
	v268 = v264 - int32(1)
	v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v268))))
	if v269 == int32(47) {
		v264 = v268
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
	v277 = v233 + base.B2i32(v238 == int32(47))
	goto L106
L105:
	;
	v277 = v264
	goto L106
L106:
	;
	v278 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v277))) = uint8(v278)
	goto L85
L107:
	;
	F_pfree(m, v22)
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L4
	} else {
		goto L155
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v233
	F_errmsg(m, v372, v12)
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L4
	} else {
		goto L153
	}
L109:
	;
	if v288 < int32(0) {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v295 = *(*int32)(unsafe.Add(mBase, _consts[87]))
	if v295 == int32(44) {
		goto L113
	} else {
		goto L114
	}
L111:
	;
	goto L112
L112:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v12)+116))
	v311 = v309 & int32(61440)
	if v311 != int32(40960) {
		goto L123
	} else {
		goto L124
	}
L113:
	;
	v298 = int32(19)
	goto L115
L114:
	;
	v298 = int32(21)
	goto L115
L115:
	;
	if l1 != 0 {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v299 = int32(15)
	goto L118
L117:
	;
	v299 = v298
	goto L118
L118:
	;
	v301 = F_errstart(m, v299, int32(0))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L4
	} else {
		goto L119
	}
L119:
	;
	if v301 == int32(0) {
		goto L107
	} else {
		goto L120
	}
L120:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L4
	} else {
		goto L121
	}
L121:
	;
	v372 = int32(310831)
	v373 = int32(805)
	goto L108
L122:
	;
	if l1 != 0 {
		goto L147
	} else {
		goto L148
	}
L123:
	;
	if v311 != int32(16384) {
		goto L122
	} else {
		goto L126
	}
L124:
	;
	goto L125
L125:
	;
	v337 = F_unlink(m, v233)
	mBase = m.M
	if int32(0) <= v337 {
		goto L107
	} else {
		goto L137
	}
L126:
	;
	v316 = F_rmdir(m, v233)
	mBase = m.M
	if int32(0) <= v316 {
		goto L107
	} else {
		goto L127
	}
L127:
	;
	v323 = *(*int32)(unsafe.Add(mBase, _consts[87]))
	if v323 == int32(44) {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v326 = int32(19)
	goto L130
L129:
	;
	v326 = int32(21)
	goto L130
L130:
	;
	if l1 != 0 {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	v327 = int32(15)
	goto L133
L132:
	;
	v327 = v326
	goto L133
L133:
	;
	v329 = F_errstart(m, v327, int32(0))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L4
	} else {
		goto L134
	}
L134:
	;
	if v329 == int32(0) {
		goto L107
	} else {
		goto L135
	}
L135:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L4
	} else {
		goto L136
	}
L136:
	;
	v372 = int32(309686)
	v373 = int32(816)
	goto L108
L137:
	;
	v344 = *(*int32)(unsafe.Add(mBase, _consts[87]))
	if v344 == int32(44) {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	v347 = int32(19)
	goto L140
L139:
	;
	v347 = int32(21)
	goto L140
L140:
	;
	if l1 != 0 {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	v348 = int32(15)
	goto L143
L142:
	;
	v348 = v347
	goto L143
L143:
	;
	v350 = F_errstart(m, v348, int32(0))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L4
	} else {
		goto L144
	}
L144:
	;
	if v350 == int32(0) {
		goto L107
	} else {
		goto L145
	}
L145:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L4
	} else {
		goto L146
	}
L146:
	;
	v372 = int32(310479)
	v373 = int32(828)
	goto L108
L147:
	;
	v360 = int32(15)
	goto L149
L148:
	;
	v360 = int32(21)
	goto L149
L149:
	;
	v362 = F_errstart(m, v360, int32(0))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L4
	} else {
		goto L150
	}
L150:
	;
	if v362 == int32(0) {
		goto L107
	} else {
		goto L151
	}
L151:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L4
	} else {
		goto L152
	}
L152:
	;
	v372 = int32(329933)
	v373 = int32(837)
	goto L108
L153:
	;
	F_errfinish(m, int32(522409), v373, int32(177644))
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L4
	} else {
		goto L154
	}
L154:
	;
	goto L107
L155:
	;
	v386 = v233
	v390 = int32(1)
	goto L1
L156:
	;
	m.G0 = v12 + int32(208)
	return v390
}
func F_has_tablespace_privilege_id(m *base.Module, l0 int32) int32 {
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v22 int64
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v12 = F_pg_detoast_datum_packed(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)) = uint8(v16)
		v20 = *(*int32)(unsafe.Add(mBase, _consts[237]))
		v22 = F_convert_any_priv_string(m, v12, int32(1690096))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			v26 = F_object_aclcheck_ext(m, int32(1213), v10, v20, v22, v8+int32(15))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
				if v28 == int32(1) {
					v31 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v31)
					v35 = int32(0)
				} else {
					v35 = base.B2i32(v26 == int32(0))
				}
				m.G0 = v8 + int32(16)
				return v35
			}
		}
	}
}
func F_has_tablespace_privilege_name(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int64
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_pg_detoast_datum_packed(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v10 = F_pg_detoast_datum_packed(m, v9)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, _consts[237]))
			v15 = F_text_to_cstring(m, v5)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				v18 = F_get_tablespace_oid(m, v15, int32(0))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					v21 = F_convert_any_priv_string(m, v10, int32(1690096))
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return int32(0)
					} else {
						v23 = F_object_aclcheck(m, int32(1213), v18, v13, v21)
						mBase = m.M
						v24 = m.ExcPending
						if v24 != 0 {
							return int32(0)
						} else {
							return base.B2i32(v23 == int32(0))
						}
					}
				}
			}
		}
	}
}
func F_has_tablespace_privilege_name_id(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int64
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v14 = F_pg_detoast_datum_packed(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v18 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)) = uint8(v18)
		v21 = F_get_role_oid_or_public(m, v12)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			v24 = F_convert_any_priv_string(m, v14, int32(1690096))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				v28 = F_object_aclcheck_ext(m, int32(1213), v11, v21, v24, v9+int32(15))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
					if v30 == int32(1) {
						v33 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v33)
						v37 = int32(0)
					} else {
						v37 = base.B2i32(v28 == int32(0))
					}
					m.G0 = v9 + int32(16)
					return v37
				}
			}
		}
	}
}
func F_has_tablespace_privilege_name_name(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
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
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int64
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = F_pg_detoast_datum_packed(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v11 = F_pg_detoast_datum_packed(m, v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			v13 = F_get_role_oid_or_public(m, v4)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				v16 = F_text_to_cstring(m, v6)
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return int32(0)
				} else {
					v19 = F_get_tablespace_oid(m, v16, int32(0))
					mBase = m.M
					v20 = m.ExcPending
					if v20 != 0 {
						return int32(0)
					} else {
						v22 = F_convert_any_priv_string(m, v11, int32(1690096))
						mBase = m.M
						v23 = m.ExcPending
						if v23 != 0 {
							return int32(0)
						} else {
							v24 = F_object_aclcheck(m, int32(1213), v19, v13, v22)
							mBase = m.M
							v25 = m.ExcPending
							if v25 != 0 {
								return int32(0)
							} else {
								return base.B2i32(v24 == int32(0))
							}
						}
					}
				}
			}
		}
	}
}
