package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_create_tablespace_directories(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	v7 = m.G0
	v9 = v7 - int32(272)
	m.G0 = v9
	*(*int32)(unsafe.Add(mBase, uint32(v9)+160)) = int32(_a_F_create_tablespace_directories_0)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+164)) = l1
	v17 = F_psprintf(m, int32(_a_F_create_tablespace_directories_1), v9+int32(160))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return
	} else {
		v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
		if v19 == int32(0) {
			v23 = *(*int32)(unsafe.Add(mBase, _c_F_create_tablespace_directories[0]))
			v24 = F_mkdir(m, v17, v23)
			mBase = m.M
			if v24 < int32(0) {
				v28 = *(*int32)(unsafe.Add(mBase, _c_F_create_tablespace_directories[1]))
				if v28 != int32(20) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v98 = m.ExcPending
					if v98 != 0 {
						return
					} else {
						F_errcode_for_file_access(m)
						mBase = m.M
						v100 = m.ExcPending
						if v100 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9)+144)) = v17
							F_errmsg(m, int32(_a_F_create_tablespace_directories_2), v9+int32(144))
							mBase = m.M
							v106 = m.ExcPending
							if v106 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_create_tablespace_directories_3), int32(594), int32(_a_F_create_tablespace_directories_4))
								mBase = m.M
								v111 = m.ExcPending
								if v111 != 0 {
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
					*(*int32)(unsafe.Add(mBase, uint32(v9)+132)) = int32(_a_F_create_tablespace_directories_5)
					*(*int32)(unsafe.Add(mBase, uint32(v9)+128)) = v17
					v37 = F_psprintf(m, int32(_a_F_create_tablespace_directories_6), v9+int32(128))
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return
					} else {
						v112 = v37
						v117 = F___fstatat(m, int32(-100), v112, v9+int32(176), int32(0))
						mBase = m.M
						if v117 < int32(0) {
							v121 = *(*int32)(unsafe.Add(mBase, _c_F_create_tablespace_directories[1]))
							if v121 != int32(44) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v183 = m.ExcPending
								if v183 != 0 {
									return
								} else {
									F_errcode_for_file_access(m)
									mBase = m.M
									v185 = m.ExcPending
									if v185 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v112
										F_errmsg(m, int32(_a_F_create_tablespace_directories_7), v9+int32(32))
										mBase = m.M
										v191 = m.ExcPending
										if v191 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_create_tablespace_directories_3), int32(634), int32(_a_F_create_tablespace_directories_4))
											mBase = m.M
											v196 = m.ExcPending
											if v196 != 0 {
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
								v125 = *(*int32)(unsafe.Add(mBase, _c_F_create_tablespace_directories[0]))
								v126 = F_mkdir(m, v112, v125)
								mBase = m.M
								if int32(0) <= v126 {
									v130 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_tablespace_directories[2])))
									v156 = v130
									v157 = int32(0)
									if base.B2i32(v19 == v157)|base.B2i32(v156&int32(1) == v157) == v157 {
										F_remove_tablespace_symlink(m, v17)
										mBase = m.M
										v167 = m.ExcPending
										if v167 != 0 {
											return
										} else {
											v170 = F_symlink(m, l0, v17)
											mBase = m.M
											if v170 < int32(0) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v236 = m.ExcPending
												if v236 != 0 {
													return
												} else {
													F_errcode_for_file_access(m)
													mBase = m.M
													v238 = m.ExcPending
													if v238 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v17
														F_errmsg(m, int32(_a_F_create_tablespace_directories_8), v9+int32(16))
														mBase = m.M
														v244 = m.ExcPending
														if v244 != 0 {
															return
														} else {
															F_errfinish(m, int32(_a_F_create_tablespace_directories_3), int32(665), int32(_a_F_create_tablespace_directories_4))
															mBase = m.M
															v249 = m.ExcPending
															if v249 != 0 {
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
												F_pfree(m, v17)
												mBase = m.M
												v174 = m.ExcPending
												if v174 != 0 {
													return
												} else {
													F_pfree(m, v112)
													mBase = m.M
													v176 = m.ExcPending
													if v176 != 0 {
														return
													} else {
														m.G0 = v9 + int32(272)
														return
													}
												}
											}
										}
									} else {
										if v19 == int32(0) {
											F_pfree(m, v17)
											mBase = m.M
											v174 = m.ExcPending
											if v174 != 0 {
												return
											} else {
												F_pfree(m, v112)
												mBase = m.M
												v176 = m.ExcPending
												if v176 != 0 {
													return
												} else {
													m.G0 = v9 + int32(272)
													return
												}
											}
										} else {
											v170 = F_symlink(m, l0, v17)
											mBase = m.M
											if v170 < int32(0) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v236 = m.ExcPending
												if v236 != 0 {
													return
												} else {
													F_errcode_for_file_access(m)
													mBase = m.M
													v238 = m.ExcPending
													if v238 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v17
														F_errmsg(m, int32(_a_F_create_tablespace_directories_8), v9+int32(16))
														mBase = m.M
														v244 = m.ExcPending
														if v244 != 0 {
															return
														} else {
															F_errfinish(m, int32(_a_F_create_tablespace_directories_3), int32(665), int32(_a_F_create_tablespace_directories_4))
															mBase = m.M
															v249 = m.ExcPending
															if v249 != 0 {
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
												F_pfree(m, v17)
												mBase = m.M
												v174 = m.ExcPending
												if v174 != 0 {
													return
												} else {
													F_pfree(m, v112)
													mBase = m.M
													v176 = m.ExcPending
													if v176 != 0 {
														return
													} else {
														m.G0 = v9 + int32(272)
														return
													}
												}
											}
										}
									}
								} else {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v134 = m.ExcPending
									if v134 != 0 {
										return
									} else {
										F_errcode_for_file_access(m)
										mBase = m.M
										v136 = m.ExcPending
										if v136 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v9))) = v112
											F_errmsg(m, int32(_a_F_create_tablespace_directories_2), v9)
											mBase = m.M
											v140 = m.ExcPending
											if v140 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_create_tablespace_directories_3), int32(639), int32(_a_F_create_tablespace_directories_4))
												mBase = m.M
												v145 = m.ExcPending
												if v145 != 0 {
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
							v146 = *(*int32)(unsafe.Add(mBase, uint32(v9)+180))
							if v146&int32(_a_F_create_tablespace_directories_9) != int32(_a_F_create_tablespace_directories_10) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v200 = m.ExcPending
								if v200 != 0 {
									return
								} else {
									F_errcode(m, int32(151027844))
									mBase = m.M
									v203 = m.ExcPending
									if v203 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v9)+64)) = v112
										F_errmsg(m, int32(_a_F_create_tablespace_directories_11), v9-int32(-64))
										mBase = m.M
										v209 = m.ExcPending
										if v209 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_create_tablespace_directories_3), int32(645), int32(_a_F_create_tablespace_directories_4))
											mBase = m.M
											v214 = m.ExcPending
											if v214 != 0 {
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
								v153 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_tablespace_directories[2])))
								if v153 == int32(0) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v218 = m.ExcPending
									if v218 != 0 {
										return
									} else {
										F_errcode(m, int32(100663621))
										mBase = m.M
										v221 = m.ExcPending
										if v221 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v112
											F_errmsg(m, int32(_a_F_create_tablespace_directories_12), v9+int32(48))
											mBase = m.M
											v227 = m.ExcPending
											if v227 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_create_tablespace_directories_3), int32(650), int32(_a_F_create_tablespace_directories_4))
												mBase = m.M
												v232 = m.ExcPending
												if v232 != 0 {
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
									v156 = int32(1)
									v157 = int32(0)
									if base.B2i32(v19 == v157)|base.B2i32(v156&int32(1) == v157) == v157 {
										F_remove_tablespace_symlink(m, v17)
										mBase = m.M
										v167 = m.ExcPending
										if v167 != 0 {
											return
										} else {
											v170 = F_symlink(m, l0, v17)
											mBase = m.M
											if v170 < int32(0) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v236 = m.ExcPending
												if v236 != 0 {
													return
												} else {
													F_errcode_for_file_access(m)
													mBase = m.M
													v238 = m.ExcPending
													if v238 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v17
														F_errmsg(m, int32(_a_F_create_tablespace_directories_8), v9+int32(16))
														mBase = m.M
														v244 = m.ExcPending
														if v244 != 0 {
															return
														} else {
															F_errfinish(m, int32(_a_F_create_tablespace_directories_3), int32(665), int32(_a_F_create_tablespace_directories_4))
															mBase = m.M
															v249 = m.ExcPending
															if v249 != 0 {
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
												F_pfree(m, v17)
												mBase = m.M
												v174 = m.ExcPending
												if v174 != 0 {
													return
												} else {
													F_pfree(m, v112)
													mBase = m.M
													v176 = m.ExcPending
													if v176 != 0 {
														return
													} else {
														m.G0 = v9 + int32(272)
														return
													}
												}
											}
										}
									} else {
										if v19 == int32(0) {
											F_pfree(m, v17)
											mBase = m.M
											v174 = m.ExcPending
											if v174 != 0 {
												return
											} else {
												F_pfree(m, v112)
												mBase = m.M
												v176 = m.ExcPending
												if v176 != 0 {
													return
												} else {
													m.G0 = v9 + int32(272)
													return
												}
											}
										} else {
											v170 = F_symlink(m, l0, v17)
											mBase = m.M
											if v170 < int32(0) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v236 = m.ExcPending
												if v236 != 0 {
													return
												} else {
													F_errcode_for_file_access(m)
													mBase = m.M
													v238 = m.ExcPending
													if v238 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v17
														F_errmsg(m, int32(_a_F_create_tablespace_directories_8), v9+int32(16))
														mBase = m.M
														v244 = m.ExcPending
														if v244 != 0 {
															return
														} else {
															F_errfinish(m, int32(_a_F_create_tablespace_directories_3), int32(665), int32(_a_F_create_tablespace_directories_4))
															mBase = m.M
															v249 = m.ExcPending
															if v249 != 0 {
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
												F_pfree(m, v17)
												mBase = m.M
												v174 = m.ExcPending
												if v174 != 0 {
													return
												} else {
													F_pfree(m, v112)
													mBase = m.M
													v176 = m.ExcPending
													if v176 != 0 {
														return
													} else {
														m.G0 = v9 + int32(272)
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
				*(*int32)(unsafe.Add(mBase, uint32(v9)+132)) = int32(_a_F_create_tablespace_directories_5)
				*(*int32)(unsafe.Add(mBase, uint32(v9)+128)) = v17
				v37 = F_psprintf(m, int32(_a_F_create_tablespace_directories_6), v9+int32(128))
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return
				} else {
					v112 = v37
					v117 = F___fstatat(m, int32(-100), v112, v9+int32(176), int32(0))
					mBase = m.M
					if v117 < int32(0) {
						v121 = *(*int32)(unsafe.Add(mBase, _c_F_create_tablespace_directories[1]))
						if v121 != int32(44) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v183 = m.ExcPending
							if v183 != 0 {
								return
							} else {
								F_errcode_for_file_access(m)
								mBase = m.M
								v185 = m.ExcPending
								if v185 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v112
									F_errmsg(m, int32(_a_F_create_tablespace_directories_7), v9+int32(32))
									mBase = m.M
									v191 = m.ExcPending
									if v191 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_create_tablespace_directories_3), int32(634), int32(_a_F_create_tablespace_directories_4))
										mBase = m.M
										v196 = m.ExcPending
										if v196 != 0 {
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
							v125 = *(*int32)(unsafe.Add(mBase, _c_F_create_tablespace_directories[0]))
							v126 = F_mkdir(m, v112, v125)
							mBase = m.M
							if int32(0) <= v126 {
								v130 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_tablespace_directories[2])))
								v156 = v130
								v157 = int32(0)
								if base.B2i32(v19 == v157)|base.B2i32(v156&int32(1) == v157) == v157 {
									F_remove_tablespace_symlink(m, v17)
									mBase = m.M
									v167 = m.ExcPending
									if v167 != 0 {
										return
									} else {
										v170 = F_symlink(m, l0, v17)
										mBase = m.M
										if v170 < int32(0) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v236 = m.ExcPending
											if v236 != 0 {
												return
											} else {
												F_errcode_for_file_access(m)
												mBase = m.M
												v238 = m.ExcPending
												if v238 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v17
													F_errmsg(m, int32(_a_F_create_tablespace_directories_8), v9+int32(16))
													mBase = m.M
													v244 = m.ExcPending
													if v244 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_create_tablespace_directories_3), int32(665), int32(_a_F_create_tablespace_directories_4))
														mBase = m.M
														v249 = m.ExcPending
														if v249 != 0 {
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
											F_pfree(m, v17)
											mBase = m.M
											v174 = m.ExcPending
											if v174 != 0 {
												return
											} else {
												F_pfree(m, v112)
												mBase = m.M
												v176 = m.ExcPending
												if v176 != 0 {
													return
												} else {
													m.G0 = v9 + int32(272)
													return
												}
											}
										}
									}
								} else {
									if v19 == int32(0) {
										F_pfree(m, v17)
										mBase = m.M
										v174 = m.ExcPending
										if v174 != 0 {
											return
										} else {
											F_pfree(m, v112)
											mBase = m.M
											v176 = m.ExcPending
											if v176 != 0 {
												return
											} else {
												m.G0 = v9 + int32(272)
												return
											}
										}
									} else {
										v170 = F_symlink(m, l0, v17)
										mBase = m.M
										if v170 < int32(0) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v236 = m.ExcPending
											if v236 != 0 {
												return
											} else {
												F_errcode_for_file_access(m)
												mBase = m.M
												v238 = m.ExcPending
												if v238 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v17
													F_errmsg(m, int32(_a_F_create_tablespace_directories_8), v9+int32(16))
													mBase = m.M
													v244 = m.ExcPending
													if v244 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_create_tablespace_directories_3), int32(665), int32(_a_F_create_tablespace_directories_4))
														mBase = m.M
														v249 = m.ExcPending
														if v249 != 0 {
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
											F_pfree(m, v17)
											mBase = m.M
											v174 = m.ExcPending
											if v174 != 0 {
												return
											} else {
												F_pfree(m, v112)
												mBase = m.M
												v176 = m.ExcPending
												if v176 != 0 {
													return
												} else {
													m.G0 = v9 + int32(272)
													return
												}
											}
										}
									}
								}
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v134 = m.ExcPending
								if v134 != 0 {
									return
								} else {
									F_errcode_for_file_access(m)
									mBase = m.M
									v136 = m.ExcPending
									if v136 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v9))) = v112
										F_errmsg(m, int32(_a_F_create_tablespace_directories_2), v9)
										mBase = m.M
										v140 = m.ExcPending
										if v140 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_create_tablespace_directories_3), int32(639), int32(_a_F_create_tablespace_directories_4))
											mBase = m.M
											v145 = m.ExcPending
											if v145 != 0 {
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
						v146 = *(*int32)(unsafe.Add(mBase, uint32(v9)+180))
						if v146&int32(_a_F_create_tablespace_directories_9) != int32(_a_F_create_tablespace_directories_10) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v200 = m.ExcPending
							if v200 != 0 {
								return
							} else {
								F_errcode(m, int32(151027844))
								mBase = m.M
								v203 = m.ExcPending
								if v203 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v9)+64)) = v112
									F_errmsg(m, int32(_a_F_create_tablespace_directories_11), v9-int32(-64))
									mBase = m.M
									v209 = m.ExcPending
									if v209 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_create_tablespace_directories_3), int32(645), int32(_a_F_create_tablespace_directories_4))
										mBase = m.M
										v214 = m.ExcPending
										if v214 != 0 {
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
							v153 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_tablespace_directories[2])))
							if v153 == int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v218 = m.ExcPending
								if v218 != 0 {
									return
								} else {
									F_errcode(m, int32(100663621))
									mBase = m.M
									v221 = m.ExcPending
									if v221 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v112
										F_errmsg(m, int32(_a_F_create_tablespace_directories_12), v9+int32(48))
										mBase = m.M
										v227 = m.ExcPending
										if v227 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_create_tablespace_directories_3), int32(650), int32(_a_F_create_tablespace_directories_4))
											mBase = m.M
											v232 = m.ExcPending
											if v232 != 0 {
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
								v156 = int32(1)
								v157 = int32(0)
								if base.B2i32(v19 == v157)|base.B2i32(v156&int32(1) == v157) == v157 {
									F_remove_tablespace_symlink(m, v17)
									mBase = m.M
									v167 = m.ExcPending
									if v167 != 0 {
										return
									} else {
										v170 = F_symlink(m, l0, v17)
										mBase = m.M
										if v170 < int32(0) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v236 = m.ExcPending
											if v236 != 0 {
												return
											} else {
												F_errcode_for_file_access(m)
												mBase = m.M
												v238 = m.ExcPending
												if v238 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v17
													F_errmsg(m, int32(_a_F_create_tablespace_directories_8), v9+int32(16))
													mBase = m.M
													v244 = m.ExcPending
													if v244 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_create_tablespace_directories_3), int32(665), int32(_a_F_create_tablespace_directories_4))
														mBase = m.M
														v249 = m.ExcPending
														if v249 != 0 {
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
											F_pfree(m, v17)
											mBase = m.M
											v174 = m.ExcPending
											if v174 != 0 {
												return
											} else {
												F_pfree(m, v112)
												mBase = m.M
												v176 = m.ExcPending
												if v176 != 0 {
													return
												} else {
													m.G0 = v9 + int32(272)
													return
												}
											}
										}
									}
								} else {
									if v19 == int32(0) {
										F_pfree(m, v17)
										mBase = m.M
										v174 = m.ExcPending
										if v174 != 0 {
											return
										} else {
											F_pfree(m, v112)
											mBase = m.M
											v176 = m.ExcPending
											if v176 != 0 {
												return
											} else {
												m.G0 = v9 + int32(272)
												return
											}
										}
									} else {
										v170 = F_symlink(m, l0, v17)
										mBase = m.M
										if v170 < int32(0) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v236 = m.ExcPending
											if v236 != 0 {
												return
											} else {
												F_errcode_for_file_access(m)
												mBase = m.M
												v238 = m.ExcPending
												if v238 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v17
													F_errmsg(m, int32(_a_F_create_tablespace_directories_8), v9+int32(16))
													mBase = m.M
													v244 = m.ExcPending
													if v244 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_create_tablespace_directories_3), int32(665), int32(_a_F_create_tablespace_directories_4))
														mBase = m.M
														v249 = m.ExcPending
														if v249 != 0 {
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
											F_pfree(m, v17)
											mBase = m.M
											v174 = m.ExcPending
											if v174 != 0 {
												return
											} else {
												F_pfree(m, v112)
												mBase = m.M
												v176 = m.ExcPending
												if v176 != 0 {
													return
												} else {
													m.G0 = v9 + int32(272)
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
			*(*int32)(unsafe.Add(mBase, uint32(v9)+116)) = int32(_a_F_create_tablespace_directories_5)
			*(*int32)(unsafe.Add(mBase, uint32(v9)+112)) = l0
			v45 = F_psprintf(m, int32(_a_F_create_tablespace_directories_6), v9+int32(112))
			mBase = m.M
			v46 = m.ExcPending
			if v46 != 0 {
				return
			} else {
				v48 = *(*int32)(unsafe.Add(mBase, _c_F_create_tablespace_directories[0]))
				v49 = F_chmod(m, l0, v48)
				mBase = m.M
				if v49 == int32(0) {
					v112 = v45
					v117 = F___fstatat(m, int32(-100), v112, v9+int32(176), int32(0))
					mBase = m.M
					if v117 < int32(0) {
						v121 = *(*int32)(unsafe.Add(mBase, _c_F_create_tablespace_directories[1]))
						if v121 != int32(44) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v183 = m.ExcPending
							if v183 != 0 {
								return
							} else {
								F_errcode_for_file_access(m)
								mBase = m.M
								v185 = m.ExcPending
								if v185 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v112
									F_errmsg(m, int32(_a_F_create_tablespace_directories_7), v9+int32(32))
									mBase = m.M
									v191 = m.ExcPending
									if v191 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_create_tablespace_directories_3), int32(634), int32(_a_F_create_tablespace_directories_4))
										mBase = m.M
										v196 = m.ExcPending
										if v196 != 0 {
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
							v125 = *(*int32)(unsafe.Add(mBase, _c_F_create_tablespace_directories[0]))
							v126 = F_mkdir(m, v112, v125)
							mBase = m.M
							if int32(0) <= v126 {
								v130 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_tablespace_directories[2])))
								v156 = v130
								v157 = int32(0)
								if base.B2i32(v19 == v157)|base.B2i32(v156&int32(1) == v157) == v157 {
									F_remove_tablespace_symlink(m, v17)
									mBase = m.M
									v167 = m.ExcPending
									if v167 != 0 {
										return
									} else {
										v170 = F_symlink(m, l0, v17)
										mBase = m.M
										if v170 < int32(0) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v236 = m.ExcPending
											if v236 != 0 {
												return
											} else {
												F_errcode_for_file_access(m)
												mBase = m.M
												v238 = m.ExcPending
												if v238 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v17
													F_errmsg(m, int32(_a_F_create_tablespace_directories_8), v9+int32(16))
													mBase = m.M
													v244 = m.ExcPending
													if v244 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_create_tablespace_directories_3), int32(665), int32(_a_F_create_tablespace_directories_4))
														mBase = m.M
														v249 = m.ExcPending
														if v249 != 0 {
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
											F_pfree(m, v17)
											mBase = m.M
											v174 = m.ExcPending
											if v174 != 0 {
												return
											} else {
												F_pfree(m, v112)
												mBase = m.M
												v176 = m.ExcPending
												if v176 != 0 {
													return
												} else {
													m.G0 = v9 + int32(272)
													return
												}
											}
										}
									}
								} else {
									if v19 == int32(0) {
										F_pfree(m, v17)
										mBase = m.M
										v174 = m.ExcPending
										if v174 != 0 {
											return
										} else {
											F_pfree(m, v112)
											mBase = m.M
											v176 = m.ExcPending
											if v176 != 0 {
												return
											} else {
												m.G0 = v9 + int32(272)
												return
											}
										}
									} else {
										v170 = F_symlink(m, l0, v17)
										mBase = m.M
										if v170 < int32(0) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v236 = m.ExcPending
											if v236 != 0 {
												return
											} else {
												F_errcode_for_file_access(m)
												mBase = m.M
												v238 = m.ExcPending
												if v238 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v17
													F_errmsg(m, int32(_a_F_create_tablespace_directories_8), v9+int32(16))
													mBase = m.M
													v244 = m.ExcPending
													if v244 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_create_tablespace_directories_3), int32(665), int32(_a_F_create_tablespace_directories_4))
														mBase = m.M
														v249 = m.ExcPending
														if v249 != 0 {
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
											F_pfree(m, v17)
											mBase = m.M
											v174 = m.ExcPending
											if v174 != 0 {
												return
											} else {
												F_pfree(m, v112)
												mBase = m.M
												v176 = m.ExcPending
												if v176 != 0 {
													return
												} else {
													m.G0 = v9 + int32(272)
													return
												}
											}
										}
									}
								}
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v134 = m.ExcPending
								if v134 != 0 {
									return
								} else {
									F_errcode_for_file_access(m)
									mBase = m.M
									v136 = m.ExcPending
									if v136 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v9))) = v112
										F_errmsg(m, int32(_a_F_create_tablespace_directories_2), v9)
										mBase = m.M
										v140 = m.ExcPending
										if v140 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_create_tablespace_directories_3), int32(639), int32(_a_F_create_tablespace_directories_4))
											mBase = m.M
											v145 = m.ExcPending
											if v145 != 0 {
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
						v146 = *(*int32)(unsafe.Add(mBase, uint32(v9)+180))
						if v146&int32(_a_F_create_tablespace_directories_9) != int32(_a_F_create_tablespace_directories_10) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v200 = m.ExcPending
							if v200 != 0 {
								return
							} else {
								F_errcode(m, int32(151027844))
								mBase = m.M
								v203 = m.ExcPending
								if v203 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v9)+64)) = v112
									F_errmsg(m, int32(_a_F_create_tablespace_directories_11), v9-int32(-64))
									mBase = m.M
									v209 = m.ExcPending
									if v209 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_create_tablespace_directories_3), int32(645), int32(_a_F_create_tablespace_directories_4))
										mBase = m.M
										v214 = m.ExcPending
										if v214 != 0 {
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
							v153 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_tablespace_directories[2])))
							if v153 == int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v218 = m.ExcPending
								if v218 != 0 {
									return
								} else {
									F_errcode(m, int32(100663621))
									mBase = m.M
									v221 = m.ExcPending
									if v221 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v112
										F_errmsg(m, int32(_a_F_create_tablespace_directories_12), v9+int32(48))
										mBase = m.M
										v227 = m.ExcPending
										if v227 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_create_tablespace_directories_3), int32(650), int32(_a_F_create_tablespace_directories_4))
											mBase = m.M
											v232 = m.ExcPending
											if v232 != 0 {
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
								v156 = int32(1)
								v157 = int32(0)
								if base.B2i32(v19 == v157)|base.B2i32(v156&int32(1) == v157) == v157 {
									F_remove_tablespace_symlink(m, v17)
									mBase = m.M
									v167 = m.ExcPending
									if v167 != 0 {
										return
									} else {
										v170 = F_symlink(m, l0, v17)
										mBase = m.M
										if v170 < int32(0) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v236 = m.ExcPending
											if v236 != 0 {
												return
											} else {
												F_errcode_for_file_access(m)
												mBase = m.M
												v238 = m.ExcPending
												if v238 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v17
													F_errmsg(m, int32(_a_F_create_tablespace_directories_8), v9+int32(16))
													mBase = m.M
													v244 = m.ExcPending
													if v244 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_create_tablespace_directories_3), int32(665), int32(_a_F_create_tablespace_directories_4))
														mBase = m.M
														v249 = m.ExcPending
														if v249 != 0 {
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
											F_pfree(m, v17)
											mBase = m.M
											v174 = m.ExcPending
											if v174 != 0 {
												return
											} else {
												F_pfree(m, v112)
												mBase = m.M
												v176 = m.ExcPending
												if v176 != 0 {
													return
												} else {
													m.G0 = v9 + int32(272)
													return
												}
											}
										}
									}
								} else {
									if v19 == int32(0) {
										F_pfree(m, v17)
										mBase = m.M
										v174 = m.ExcPending
										if v174 != 0 {
											return
										} else {
											F_pfree(m, v112)
											mBase = m.M
											v176 = m.ExcPending
											if v176 != 0 {
												return
											} else {
												m.G0 = v9 + int32(272)
												return
											}
										}
									} else {
										v170 = F_symlink(m, l0, v17)
										mBase = m.M
										if v170 < int32(0) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v236 = m.ExcPending
											if v236 != 0 {
												return
											} else {
												F_errcode_for_file_access(m)
												mBase = m.M
												v238 = m.ExcPending
												if v238 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v17
													F_errmsg(m, int32(_a_F_create_tablespace_directories_8), v9+int32(16))
													mBase = m.M
													v244 = m.ExcPending
													if v244 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_create_tablespace_directories_3), int32(665), int32(_a_F_create_tablespace_directories_4))
														mBase = m.M
														v249 = m.ExcPending
														if v249 != 0 {
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
											F_pfree(m, v17)
											mBase = m.M
											v174 = m.ExcPending
											if v174 != 0 {
												return
											} else {
												F_pfree(m, v112)
												mBase = m.M
												v176 = m.ExcPending
												if v176 != 0 {
													return
												} else {
													m.G0 = v9 + int32(272)
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
					v53 = *(*int32)(unsafe.Add(mBase, _c_F_create_tablespace_directories[1]))
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
						return
					} else {
						if v53 == int32(44) {
							F_errcode(m, int32(16908805))
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9)+80)) = l0
								F_errmsg(m, int32(_a_F_create_tablespace_directories_13), v9+int32(80))
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
									return
								} else {
									v70 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_tablespace_directories[2])))
									if v70 == int32(1) {
										F_errhint(m, int32(_a_F_create_tablespace_directories_14), int32(0))
										mBase = m.M
										v76 = m.ExcPending
										if v76 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_create_tablespace_directories_3), int32(613), int32(_a_F_create_tablespace_directories_4))
											mBase = m.M
											v81 = m.ExcPending
											if v81 != 0 {
												return
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									} else {
										F_errfinish(m, int32(_a_F_create_tablespace_directories_3), int32(613), int32(_a_F_create_tablespace_directories_4))
										mBase = m.M
										v81 = m.ExcPending
										if v81 != 0 {
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
							v83 = m.ExcPending
							if v83 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9)+96)) = l0
								F_errmsg(m, int32(_a_F_create_tablespace_directories_15), v9+int32(96))
								mBase = m.M
								v89 = m.ExcPending
								if v89 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_create_tablespace_directories_3), int32(618), int32(_a_F_create_tablespace_directories_4))
									mBase = m.M
									v94 = m.ExcPending
									if v94 != 0 {
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
	var v253 int32
	_ = v253
	var v259 int32
	_ = v259
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v289 int32
	_ = v289
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v317 int32
	_ = v317
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v392 int32
	_ = v392
	var v397 int32
	_ = v397
	v3 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(208)
	m.G0 = v12
	*(*int32)(unsafe.Add(mBase, uint32(v12)+96)) = int32(_a_F_destroy_tablespace_directories_0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+100)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v12)+104)) = int32(_a_F_destroy_tablespace_directories_1)
	v22 = F_psprintf(m, int32(_a_F_destroy_tablespace_directories_2), v12+int32(96))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_pfree(m, v387)
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L4
	} else {
		goto L153
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
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_destroy_tablespace_directories[0]))
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
	F_errmsg(m, int32(_a_F_destroy_tablespace_directories_3), v12+int32(16))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L4
	} else {
		goto L15
	}
L15:
	;
	F_errfinish(m, int32(_a_F_destroy_tablespace_directories_4), int32(729), int32(_a_F_destroy_tablespace_directories_5))
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
	v387 = v22
	v392 = v3
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
	F_errmsg(m, int32(_a_F_destroy_tablespace_directories_3), v12+int32(32))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(_a_F_destroy_tablespace_directories_4), int32(739), int32(_a_F_destroy_tablespace_directories_5))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L4
	} else {
		goto L24
	}
L24:
	;
	v387 = v22
	v392 = v3
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
	v105 = F_psprintf(m, int32(_a_F_destroy_tablespace_directories_6), v12+int32(80))
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
	v387 = v22
	v392 = int32(0)
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
	F_errmsg(m, int32(_a_F_destroy_tablespace_directories_7), v12-int32(-64))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L4
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(_a_F_destroy_tablespace_directories_4), int32(768), int32(_a_F_destroy_tablespace_directories_5))
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
	v387 = v22
	v392 = v202
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
	F_errmsg(m, int32(_a_F_destroy_tablespace_directories_7), v12+int32(48))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L4
	} else {
		goto L79
	}
L79:
	;
	F_errfinish(m, int32(_a_F_destroy_tablespace_directories_4), int32(781), int32(_a_F_destroy_tablespace_directories_5))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L4
	} else {
		goto L80
	}
L80:
	;
	v387 = v22
	v392 = v202
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
	v289 = F___fstatat(m, int32(-100), v233, v12+int32(112), int32(256))
	mBase = m.M
	goto L106
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
	if base.B2i32(v247 == int32(47))&base.B2i32(base.Ui32(v233) < base.Ui32(v246)) != 0 {
		v242 = v246
		goto L86
	} else {
		goto L88
	}
L87:
	;
	v253 = v246
	goto L89
L88:
	;
	goto L87
L89:
	;
	if base.Ui32(v233) < base.Ui32(v253) {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	v265 = v253
	goto L95
L91:
	;
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v253))))
	if v259 != int32(47) {
		v253 = v253 - int32(1)
		goto L89
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	goto L90
L94:
	;
	goto L93
L95:
	;
	if base.Ui32(v233) < base.Ui32(v265) {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	if v233 == v265 {
		goto L101
	} else {
		goto L102
	}
L97:
	;
	v269 = v265 - int32(1)
	v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v269))))
	if v270 == int32(47) {
		v265 = v269
		goto L95
	} else {
		goto L100
	}
L98:
	;
	goto L99
L99:
	;
	goto L96
L100:
	;
	goto L99
L101:
	;
	v278 = v233 + base.B2i32(v238 == int32(47))
	goto L103
L102:
	;
	v278 = v265
	goto L103
L103:
	;
	v279 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v278))) = uint8(v279)
	goto L85
L104:
	;
	F_pfree(m, v22)
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L4
	} else {
		goto L152
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v233
	F_errmsg(m, v372, v12)
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L4
	} else {
		goto L150
	}
L106:
	;
	if v289 < int32(0) {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v296 = *(*int32)(unsafe.Add(mBase, _c_F_destroy_tablespace_directories[0]))
	if v296 == int32(44) {
		goto L110
	} else {
		goto L111
	}
L108:
	;
	goto L109
L109:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v12)+116))
	v312 = v310 & int32(_a_F_destroy_tablespace_directories_8)
	if v312 != int32(_a_F_destroy_tablespace_directories_9) {
		goto L120
	} else {
		goto L121
	}
L110:
	;
	v299 = int32(19)
	goto L112
L111:
	;
	v299 = int32(21)
	goto L112
L112:
	;
	if l1 != 0 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v300 = int32(15)
	goto L115
L114:
	;
	v300 = v299
	goto L115
L115:
	;
	v302 = F_errstart(m, v300, int32(0))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L4
	} else {
		goto L116
	}
L116:
	;
	if v302 == int32(0) {
		goto L104
	} else {
		goto L117
	}
L117:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L4
	} else {
		goto L118
	}
L118:
	;
	v372 = int32(_a_F_destroy_tablespace_directories_10)
	v374 = int32(805)
	goto L105
L119:
	;
	if l1 != 0 {
		goto L144
	} else {
		goto L145
	}
L120:
	;
	if v312 != int32(_a_F_destroy_tablespace_directories_11) {
		goto L119
	} else {
		goto L123
	}
L121:
	;
	goto L122
L122:
	;
	v338 = F_unlink(m, v233)
	mBase = m.M
	if int32(0) <= v338 {
		goto L104
	} else {
		goto L134
	}
L123:
	;
	v317 = F_rmdir(m, v233)
	mBase = m.M
	if int32(0) <= v317 {
		goto L104
	} else {
		goto L124
	}
L124:
	;
	v324 = *(*int32)(unsafe.Add(mBase, _c_F_destroy_tablespace_directories[0]))
	if v324 == int32(44) {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v327 = int32(19)
	goto L127
L126:
	;
	v327 = int32(21)
	goto L127
L127:
	;
	if l1 != 0 {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v328 = int32(15)
	goto L130
L129:
	;
	v328 = v327
	goto L130
L130:
	;
	v330 = F_errstart(m, v328, int32(0))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L4
	} else {
		goto L131
	}
L131:
	;
	if v330 == int32(0) {
		goto L104
	} else {
		goto L132
	}
L132:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L4
	} else {
		goto L133
	}
L133:
	;
	v372 = int32(_a_F_destroy_tablespace_directories_7)
	v374 = int32(816)
	goto L105
L134:
	;
	v345 = *(*int32)(unsafe.Add(mBase, _c_F_destroy_tablespace_directories[0]))
	if v345 == int32(44) {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v348 = int32(19)
	goto L137
L136:
	;
	v348 = int32(21)
	goto L137
L137:
	;
	if l1 != 0 {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	v349 = int32(15)
	goto L140
L139:
	;
	v349 = v348
	goto L140
L140:
	;
	v351 = F_errstart(m, v349, int32(0))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L4
	} else {
		goto L141
	}
L141:
	;
	if v351 == int32(0) {
		goto L104
	} else {
		goto L142
	}
L142:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L4
	} else {
		goto L143
	}
L143:
	;
	v372 = int32(_a_F_destroy_tablespace_directories_12)
	v374 = int32(828)
	goto L105
L144:
	;
	v361 = int32(15)
	goto L146
L145:
	;
	v361 = int32(21)
	goto L146
L146:
	;
	v363 = F_errstart(m, v361, int32(0))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L4
	} else {
		goto L147
	}
L147:
	;
	if v363 == int32(0) {
		goto L104
	} else {
		goto L148
	}
L148:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L4
	} else {
		goto L149
	}
L149:
	;
	v372 = int32(_a_F_destroy_tablespace_directories_13)
	v374 = int32(837)
	goto L105
L150:
	;
	F_errfinish(m, int32(_a_F_destroy_tablespace_directories_4), v374, int32(_a_F_destroy_tablespace_directories_5))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L4
	} else {
		goto L151
	}
L151:
	;
	goto L104
L152:
	;
	v387 = v233
	v392 = int32(1)
	goto L1
L153:
	;
	m.G0 = v12 + int32(208)
	return v392
}
func F_has_tablespace_privilege_id(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13913(m, l0, int32(_a_F_has_tablespace_privilege_id_0), int32(1213))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
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
			v13 = *(*int32)(unsafe.Add(mBase, _c_F_has_tablespace_privilege_name[0]))
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
					v21 = F_convert_any_priv_string(m, v10, int32(_a_F_has_tablespace_privilege_name_0))
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
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13915(m, l0, int32(_a_F_has_tablespace_privilege_name_id_0), int32(1213))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
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
						v22 = F_convert_any_priv_string(m, v11, int32(_a_F_has_tablespace_privilege_name_name_0))
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
