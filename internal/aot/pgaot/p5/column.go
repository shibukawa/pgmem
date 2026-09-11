package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_add_column_to_pathtarget(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v6 = F_lappend(m, v5, l1)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v6
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		if v9 != 0 {
			if v6 != 0 {
				v10 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
				v12 = v10
			} else {
				v12 = int32(0)
			}
			v14 = v12 << (uint(int32(2)) % 32)
			v15 = F_repalloc(m, v9, v14)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v15
				v32 = v14 + v15
				*(*int32)(unsafe.Add(mBase, uint32(v32-int32(4)))) = l2
				v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
				if v38 == int32(2) {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = int32(0)
				} else {
				}
				return
			}
		} else {
			if l2 == int32(0) {
				v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
				if v38 == int32(2) {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = int32(0)
				} else {
				}
				return
			} else {
				if v6 != 0 {
					v21 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
					v23 = v21
				} else {
					v23 = int32(0)
				}
				v25 = v23 << (uint(int32(2)) % 32)
				v26 = F_palloc0(m, v25)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v26
					v32 = v25 + v26
					*(*int32)(unsafe.Add(mBase, uint32(v32-int32(4)))) = l2
					v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
					if v38 == int32(2) {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = int32(0)
					} else {
					}
					return
				}
			}
		}
	}
}
func F_build_column_default(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
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
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
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
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	v12 = m.G0
	v14 = v12 - int32(32)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v23 = v16 + v17<<(uint(int32(4))%32) + l1*int32(100)
	v25 = v23 - int32(80)
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+9)))
	if v26 != 0 {
		v28 = F_palloc0(m, int32(12))
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v28))) = int32(59)
			v36 = F_getIdentitySequence(m, l0, base.I32_extend16_s(l1), int32(0))
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v28)+4)) = v36
				v39 = *(*int32)(unsafe.Add(mBase, uint32(v25)+68))
				*(*int32)(unsafe.Add(mBase, uint32(v28)+8)) = v39
				v192 = v28
				m.G0 = v14 + int32(32)
				return v192
			}
		}
	} else {
		v41 = *(*int32)(unsafe.Add(mBase, uint32(v25)+76))
		v42 = *(*int32)(unsafe.Add(mBase, uint32(v25)+68))
		v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+87)))
		if v43 == int32(1) {
			v47 = F_TupleDescGetDefault(m, v16, base.I32_extend16_s(l1))
			mBase = m.M
			v48 = m.ExcPending
			if v48 != 0 {
				return int32(0)
			} else {
				if v47 != 0 {
					v146 = v47
					v152 = F_exprType(m, v146)
					mBase = m.M
					v153 = m.ExcPending
					if v153 != 0 {
						return int32(0)
					} else {
						v157 = F_coerce_to_target_type(m, int32(0), v146, v152, v42, v41, int32(1), int32(2), int32(-1))
						mBase = m.M
						v158 = m.ExcPending
						if v158 != 0 {
							return int32(0)
						} else {
							if v157 != 0 {
								v192 = v157
								m.G0 = v14 + int32(32)
								return v192
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v162 = m.ExcPending
								if v162 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(67141764))
									mBase = m.M
									v165 = m.ExcPending
									if v165 != 0 {
										return int32(0)
									} else {
										v166 = F_format_type_be(m, v42)
										mBase = m.M
										v167 = m.ExcPending
										if v167 != 0 {
											return int32(0)
										} else {
											v168 = F_format_type_be(m, v152)
											mBase = m.M
											v169 = m.ExcPending
											if v169 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v168
												*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v166
												*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v25 + int32(4)
												F_errmsg(m, int32(_a_F_build_column_default_0), v14+int32(16))
												mBase = m.M
												v179 = m.ExcPending
												if v179 != 0 {
													return int32(0)
												} else {
													F_errhint(m, int32(_a_F_build_column_default_1), int32(0))
													mBase = m.M
													v183 = m.ExcPending
													if v183 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_build_column_default_2), int32(1293), int32(_a_F_build_column_default_3))
														mBase = m.M
														v188 = m.ExcPending
														if v188 != 0 {
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
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return int32(0)
					} else {
						v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
						*(*int32)(unsafe.Add(mBase, uint32(v14))) = l1
						*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v53 + int32(4)
						F_errmsg_internal(m, int32(_a_F_build_column_default_4), v14)
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_build_column_default_2), int32(1257), int32(_a_F_build_column_default_3))
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
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
			v66 = int32(0)
			v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+90)))
			if v67 != 0 {
				v192 = v66
				m.G0 = v14 + int32(32)
				return v192
			} else {
				v68 = m.G0
				v70 = v68 - int32(16)
				m.G0 = v70
				v73 = F_SearchSysCache1(m, int32(82), v42)
				mBase = m.M
				v74 = m.ExcPending
				if v74 != 0 {
					return int32(0)
				} else {
					if v73 != 0 {
						v75 = *(*int32)(unsafe.Add(mBase, uint32(v73)+16))
						v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+22)))
						v81 = F_SysCacheGetAttr(m, int32(82), v73, int32(30), v70+int32(15))
						mBase = m.M
						v82 = m.ExcPending
						if v82 != 0 {
							return int32(0)
						} else {
							v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+15)))
							if v83 == int32(0) {
								v86 = F_text_to_cstring(m, v81)
								mBase = m.M
								v87 = m.ExcPending
								if v87 != 0 {
									return int32(0)
								} else {
									v88 = F_stringToNode(m, v86)
									mBase = m.M
									v89 = m.ExcPending
									if v89 != 0 {
										return int32(0)
									} else {
										v120 = v88
										F_ReleaseCatCache(m, v73)
										mBase = m.M
										v125 = m.ExcPending
										if v125 != 0 {
											return int32(0)
										} else {
											m.G0 = v70 + int32(16)
											if v120 == int32(0) {
												v192 = v66
												m.G0 = v14 + int32(32)
												return v192
											} else {
												v146 = v120
												v152 = F_exprType(m, v146)
												mBase = m.M
												v153 = m.ExcPending
												if v153 != 0 {
													return int32(0)
												} else {
													v157 = F_coerce_to_target_type(m, int32(0), v146, v152, v42, v41, int32(1), int32(2), int32(-1))
													mBase = m.M
													v158 = m.ExcPending
													if v158 != 0 {
														return int32(0)
													} else {
														if v157 != 0 {
															v192 = v157
															m.G0 = v14 + int32(32)
															return v192
														} else {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v162 = m.ExcPending
															if v162 != 0 {
																return int32(0)
															} else {
																F_errcode(m, int32(67141764))
																mBase = m.M
																v165 = m.ExcPending
																if v165 != 0 {
																	return int32(0)
																} else {
																	v166 = F_format_type_be(m, v42)
																	mBase = m.M
																	v167 = m.ExcPending
																	if v167 != 0 {
																		return int32(0)
																	} else {
																		v168 = F_format_type_be(m, v152)
																		mBase = m.M
																		v169 = m.ExcPending
																		if v169 != 0 {
																			return int32(0)
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v168
																			*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v166
																			*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v25 + int32(4)
																			F_errmsg(m, int32(_a_F_build_column_default_0), v14+int32(16))
																			mBase = m.M
																			v179 = m.ExcPending
																			if v179 != 0 {
																				return int32(0)
																			} else {
																				F_errhint(m, int32(_a_F_build_column_default_1), int32(0))
																				mBase = m.M
																				v183 = m.ExcPending
																				if v183 != 0 {
																					return int32(0)
																				} else {
																					F_errfinish(m, int32(_a_F_build_column_default_2), int32(1293), int32(_a_F_build_column_default_3))
																					mBase = m.M
																					v188 = m.ExcPending
																					if v188 != 0 {
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
							} else {
								v95 = F_SysCacheGetAttr(m, int32(82), v73, int32(31), v70+int32(15))
								mBase = m.M
								v96 = m.ExcPending
								if v96 != 0 {
									return int32(0)
								} else {
									v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+15)))
									if v97 != 0 {
										v120 = int32(0)
										F_ReleaseCatCache(m, v73)
										mBase = m.M
										v125 = m.ExcPending
										if v125 != 0 {
											return int32(0)
										} else {
											m.G0 = v70 + int32(16)
											if v120 == int32(0) {
												v192 = v66
												m.G0 = v14 + int32(32)
												return v192
											} else {
												v146 = v120
												v152 = F_exprType(m, v146)
												mBase = m.M
												v153 = m.ExcPending
												if v153 != 0 {
													return int32(0)
												} else {
													v157 = F_coerce_to_target_type(m, int32(0), v146, v152, v42, v41, int32(1), int32(2), int32(-1))
													mBase = m.M
													v158 = m.ExcPending
													if v158 != 0 {
														return int32(0)
													} else {
														if v157 != 0 {
															v192 = v157
															m.G0 = v14 + int32(32)
															return v192
														} else {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v162 = m.ExcPending
															if v162 != 0 {
																return int32(0)
															} else {
																F_errcode(m, int32(67141764))
																mBase = m.M
																v165 = m.ExcPending
																if v165 != 0 {
																	return int32(0)
																} else {
																	v166 = F_format_type_be(m, v42)
																	mBase = m.M
																	v167 = m.ExcPending
																	if v167 != 0 {
																		return int32(0)
																	} else {
																		v168 = F_format_type_be(m, v152)
																		mBase = m.M
																		v169 = m.ExcPending
																		if v169 != 0 {
																			return int32(0)
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v168
																			*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v166
																			*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v25 + int32(4)
																			F_errmsg(m, int32(_a_F_build_column_default_0), v14+int32(16))
																			mBase = m.M
																			v179 = m.ExcPending
																			if v179 != 0 {
																				return int32(0)
																			} else {
																				F_errhint(m, int32(_a_F_build_column_default_1), int32(0))
																				mBase = m.M
																				v183 = m.ExcPending
																				if v183 != 0 {
																					return int32(0)
																				} else {
																					F_errfinish(m, int32(_a_F_build_column_default_2), int32(1293), int32(_a_F_build_column_default_3))
																					mBase = m.M
																					v188 = m.ExcPending
																					if v188 != 0 {
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
									} else {
										v98 = F_text_to_cstring(m, v95)
										mBase = m.M
										v99 = m.ExcPending
										if v99 != 0 {
											return int32(0)
										} else {
											v100 = v75 + v76
											v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+100))
											v102 = *(*int32)(unsafe.Add(mBase, uint32(v73)+16))
											v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102)+22)))
											v104 = v102 + v103
											v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)+92))
											if v105 != 0 {
												v107 = v105
											} else {
												v106 = *(*int32)(unsafe.Add(mBase, uint32(v104)))
												v107 = v106
											}
											v109 = F_OidInputFunctionCall(m, v101, v98, v107, int32(-1))
											mBase = m.M
											v110 = m.ExcPending
											if v110 != 0 {
												return int32(0)
											} else {
												v112 = *(*int32)(unsafe.Add(mBase, uint32(v100)+144))
												v113 = int32(*(*int16)(unsafe.Add(mBase, uint32(v100)+76)))
												v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+78)))
												v116 = F_makeConst(m, v42, int32(-1), v112, v113, v109, int32(0), v115)
												mBase = m.M
												v117 = m.ExcPending
												if v117 != 0 {
													return int32(0)
												} else {
													F_pfree(m, v98)
													mBase = m.M
													v119 = m.ExcPending
													if v119 != 0 {
														return int32(0)
													} else {
														v120 = v116
														F_ReleaseCatCache(m, v73)
														mBase = m.M
														v125 = m.ExcPending
														if v125 != 0 {
															return int32(0)
														} else {
															m.G0 = v70 + int32(16)
															if v120 == int32(0) {
																v192 = v66
																m.G0 = v14 + int32(32)
																return v192
															} else {
																v146 = v120
																v152 = F_exprType(m, v146)
																mBase = m.M
																v153 = m.ExcPending
																if v153 != 0 {
																	return int32(0)
																} else {
																	v157 = F_coerce_to_target_type(m, int32(0), v146, v152, v42, v41, int32(1), int32(2), int32(-1))
																	mBase = m.M
																	v158 = m.ExcPending
																	if v158 != 0 {
																		return int32(0)
																	} else {
																		if v157 != 0 {
																			v192 = v157
																			m.G0 = v14 + int32(32)
																			return v192
																		} else {
																			F_errstart_cold(m, int32(21), int32(0))
																			mBase = m.M
																			v162 = m.ExcPending
																			if v162 != 0 {
																				return int32(0)
																			} else {
																				F_errcode(m, int32(67141764))
																				mBase = m.M
																				v165 = m.ExcPending
																				if v165 != 0 {
																					return int32(0)
																				} else {
																					v166 = F_format_type_be(m, v42)
																					mBase = m.M
																					v167 = m.ExcPending
																					if v167 != 0 {
																						return int32(0)
																					} else {
																						v168 = F_format_type_be(m, v152)
																						mBase = m.M
																						v169 = m.ExcPending
																						if v169 != 0 {
																							return int32(0)
																						} else {
																							*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v168
																							*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v166
																							*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v25 + int32(4)
																							F_errmsg(m, int32(_a_F_build_column_default_0), v14+int32(16))
																							mBase = m.M
																							v179 = m.ExcPending
																							if v179 != 0 {
																								return int32(0)
																							} else {
																								F_errhint(m, int32(_a_F_build_column_default_1), int32(0))
																								mBase = m.M
																								v183 = m.ExcPending
																								if v183 != 0 {
																									return int32(0)
																								} else {
																									F_errfinish(m, int32(_a_F_build_column_default_2), int32(1293), int32(_a_F_build_column_default_3))
																									mBase = m.M
																									v188 = m.ExcPending
																									if v188 != 0 {
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
										}
									}
								}
							}
						}
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v132 = m.ExcPending
						if v132 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v70))) = v42
							F_errmsg_internal(m, int32(_a_F_build_column_default_5), v70)
							mBase = m.M
							v136 = m.ExcPending
							if v136 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_build_column_default_6), int32(2598), int32(_a_F_build_column_default_7))
								mBase = m.M
								v141 = m.ExcPending
								if v141 != 0 {
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
func F_has_column_privilege_id_id_attnum(m *base.Module, l0 int32) int32 {
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
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int64
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
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
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v15 = F_pg_detoast_datum_packed(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v20 = F_convert_any_priv_string(m, v15, int32(_a_F_has_column_privilege_id_id_attnum_0))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			v22 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)) = uint8(v22)
			if v12&int32(_a_F_has_column_privilege_id_id_attnum_1) == v22 {
				v42 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v42)
				v45 = int32(0)
				m.G0 = v9 + int32(16)
				return v45
			} else {
				v31 = F_pg_attribute_aclcheck_ext(m, v11, base.I32_extend16_s(v12), v13, v20, v9+int32(15))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					if v31 != 0 {
						v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
						if v33 != 0 {
							v42 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v42)
							v45 = int32(0)
							m.G0 = v9 + int32(16)
							return v45
						} else {
							v36 = F_pg_class_aclcheck_ext(m, v11, v13, v20, v9+int32(15))
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return int32(0)
							} else {
								if v36 != 0 {
									v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
									if v39 == int32(0) {
									} else {
										v42 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v42)
									}
									v45 = int32(0)
								} else {
									v45 = int32(1)
								}
								m.G0 = v9 + int32(16)
								return v45
							}
						}
					} else {
						v45 = int32(1)
						m.G0 = v9 + int32(16)
						return v45
					}
				}
			}
		}
	}
}
func F_has_column_privilege_name_attnum(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
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
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int64
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = F_pg_detoast_datum_packed(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v19 = F_pg_detoast_datum_packed(m, v18)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, _c_F_has_column_privilege_name_attnum[0]))
			v23 = F_textToQualifiedNameList(m, v13)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				v25 = F_makeRangeVarFromNameList(m, v23)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					v27 = int32(0)
					v31 = F_RangeVarGetRelidExtended(m, v25, v27, v27, v27, v27)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						v34 = F_convert_any_priv_string(m, v19, int32(_a_F_has_column_privilege_name_attnum_0))
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int32(0)
						} else {
							v36 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v36)
							if v17&int32(_a_F_has_column_privilege_name_attnum_1) == v36 {
								v56 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v56)
								v59 = int32(0)
								m.G0 = v10 + int32(16)
								return v59
							} else {
								v45 = F_pg_attribute_aclcheck_ext(m, v31, base.I32_extend16_s(v17), v22, v34, v10+int32(15))
								mBase = m.M
								v46 = m.ExcPending
								if v46 != 0 {
									return int32(0)
								} else {
									if v45 != 0 {
										v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
										if v47 != 0 {
											v56 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v56)
											v59 = int32(0)
											m.G0 = v10 + int32(16)
											return v59
										} else {
											v50 = F_pg_class_aclcheck_ext(m, v31, v22, v34, v10+int32(15))
											mBase = m.M
											v51 = m.ExcPending
											if v51 != 0 {
												return int32(0)
											} else {
												if v50 != 0 {
													v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
													if v53 == int32(0) {
													} else {
														v56 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v56)
													}
													v59 = int32(0)
												} else {
													v59 = int32(1)
												}
												m.G0 = v10 + int32(16)
												return v59
											}
										}
									} else {
										v59 = int32(1)
										m.G0 = v10 + int32(16)
										return v59
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
func F_has_column_privilege_name_name_attnum(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
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
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int64
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v14 = F_pg_detoast_datum_packed(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
		v20 = F_pg_detoast_datum_packed(m, v19)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			v22 = F_get_role_oid_or_public(m, v12)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				v24 = F_textToQualifiedNameList(m, v14)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					v26 = F_makeRangeVarFromNameList(m, v24)
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						v28 = int32(0)
						v32 = F_RangeVarGetRelidExtended(m, v26, v28, v28, v28, v28)
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int32(0)
						} else {
							v35 = F_convert_any_priv_string(m, v20, int32(_a_F_has_column_privilege_name_name_attnum_0))
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return int32(0)
							} else {
								v37 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v37)
								if v18&int32(_a_F_has_column_privilege_name_name_attnum_1) == v37 {
									v57 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v57)
									v60 = int32(0)
									m.G0 = v10 + int32(16)
									return v60
								} else {
									v46 = F_pg_attribute_aclcheck_ext(m, v32, base.I32_extend16_s(v18), v22, v35, v10+int32(15))
									mBase = m.M
									v47 = m.ExcPending
									if v47 != 0 {
										return int32(0)
									} else {
										if v46 != 0 {
											v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
											if v48 != 0 {
												v57 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v57)
												v60 = int32(0)
												m.G0 = v10 + int32(16)
												return v60
											} else {
												v51 = F_pg_class_aclcheck_ext(m, v32, v22, v35, v10+int32(15))
												mBase = m.M
												v52 = m.ExcPending
												if v52 != 0 {
													return int32(0)
												} else {
													if v51 != 0 {
														v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
														if v54 == int32(0) {
														} else {
															v57 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v57)
														}
														v60 = int32(0)
													} else {
														v60 = int32(1)
													}
													m.G0 = v10 + int32(16)
													return v60
												}
											}
										} else {
											v60 = int32(1)
											m.G0 = v10 + int32(16)
											return v60
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
func F_makeColumnDef(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v27 int64
	_ = v27
	var v30 int32
	_ = v30
	v7 = F_palloc0(m, int32(68))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(90)
		v13 = F_pstrdup(m, l0)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v13
			v17 = F_palloc0(m, int32(32))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				v19 = int32(-1)
				*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = v19
				*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = l2
				*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v17))) = int32(68)
				*(*int32)(unsafe.Add(mBase, uint32(v7)+64)) = v19
				v27 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v7)+56)) = v27
				*(*int32)(unsafe.Add(mBase, uint32(v7)+52)) = l3
				v30 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = v30
				*(*int64)(unsafe.Add(mBase, uint32(v7)+28)) = v27
				*(*int32)(unsafe.Add(mBase, uint32(v7)+18)) = int32(1)
				*(*uint16)(unsafe.Add(mBase, uint32(v7)+16)) = uint16(v30)
				*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v17
				return v7
			}
		}
	}
}
func F_prepare_column_cache(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
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
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = l1
	v17 = F_SearchSysCache1(m, int32(82), l1)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return
	} else {
		if v17 != 0 {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
			v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+22)))
			v21 = v19 + v20
			v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+79)))
			if v22 == int32(100) {
				*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = l2
				v28 = F_getBaseTypeAndTypmod(m, l1, v12+int32(12))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return
				} else {
					v30 = F_get_typtype(m, v28)
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return
					} else {
						if v30 == int32(99) {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v28
							v35 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v35
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(67)
							v39 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v35
							v52 = v39
							v53 = v35
							*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v52
							*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v53
							if l4 == int32(0) {
								F_ReleaseCatCache(m, v17)
								mBase = m.M
								v121 = m.ExcPending
								if v121 != 0 {
									return
								} else {
									m.G0 = v12 + int32(16)
									return
								}
							} else {
								F_getTypeInputInfo(m, l1, v12+int32(8), l0+int32(12))
								mBase = m.M
								v112 = m.ExcPending
								if v112 != 0 {
									return
								} else {
									v113 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
									F_fmgr_info_cxt(m, v113, l0+int32(16), l3)
									mBase = m.M
									v117 = m.ExcPending
									if v117 != 0 {
										return
									} else {
										F_ReleaseCatCache(m, v17)
										mBase = m.M
										v121 = m.ExcPending
										if v121 != 0 {
											return
										} else {
											m.G0 = v12 + int32(16)
											return
										}
									}
								}
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v28
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(100)
							v46 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v46
							v50 = F_MemoryContextAllocZero(m, l3, int32(64))
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return
							} else {
								v52 = int32(0)
								v53 = v50
								*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v52
								*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v53
								if l4 == int32(0) {
									F_ReleaseCatCache(m, v17)
									mBase = m.M
									v121 = m.ExcPending
									if v121 != 0 {
										return
									} else {
										m.G0 = v12 + int32(16)
										return
									}
								} else {
									F_getTypeInputInfo(m, l1, v12+int32(8), l0+int32(12))
									mBase = m.M
									v112 = m.ExcPending
									if v112 != 0 {
										return
									} else {
										v113 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
										F_fmgr_info_cxt(m, v113, l0+int32(16), l3)
										mBase = m.M
										v117 = m.ExcPending
										if v117 != 0 {
											return
										} else {
											F_ReleaseCatCache(m, v17)
											mBase = m.M
											v121 = m.ExcPending
											if v121 != 0 {
												return
											} else {
												m.G0 = v12 + int32(16)
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
				if base.B2i32(l1 != int32(2249))&base.B2i32(v22 != int32(99)) == int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = l2
					*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = l1
					*(*int64)(unsafe.Add(mBase, uint32(l0)+44)) = int64(0)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(99)
					if l4 == int32(0) {
						F_ReleaseCatCache(m, v17)
						mBase = m.M
						v121 = m.ExcPending
						if v121 != 0 {
							return
						} else {
							m.G0 = v12 + int32(16)
							return
						}
					} else {
						F_getTypeInputInfo(m, l1, v12+int32(8), l0+int32(12))
						mBase = m.M
						v112 = m.ExcPending
						if v112 != 0 {
							return
						} else {
							v113 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
							F_fmgr_info_cxt(m, v113, l0+int32(16), l3)
							mBase = m.M
							v117 = m.ExcPending
							if v117 != 0 {
								return
							} else {
								F_ReleaseCatCache(m, v17)
								mBase = m.M
								v121 = m.ExcPending
								if v121 != 0 {
									return
								} else {
									m.G0 = v12 + int32(16)
									return
								}
							}
						}
					}
				} else {
					v71 = *(*int32)(unsafe.Add(mBase, uint32(v21)+92))
					if v71 == int32(0) {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(115)
						F_getTypeInputInfo(m, l1, v12+int32(8), l0+int32(12))
						mBase = m.M
						v112 = m.ExcPending
						if v112 != 0 {
							return
						} else {
							v113 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
							F_fmgr_info_cxt(m, v113, l0+int32(16), l3)
							mBase = m.M
							v117 = m.ExcPending
							if v117 != 0 {
								return
							} else {
								F_ReleaseCatCache(m, v17)
								mBase = m.M
								v121 = m.ExcPending
								if v121 != 0 {
									return
								} else {
									m.G0 = v12 + int32(16)
									return
								}
							}
						}
					} else {
						v74 = *(*int32)(unsafe.Add(mBase, uint32(v21)+88))
						if v74 != int32(_a_F_prepare_column_cache_0) {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(115)
							F_getTypeInputInfo(m, l1, v12+int32(8), l0+int32(12))
							mBase = m.M
							v112 = m.ExcPending
							if v112 != 0 {
								return
							} else {
								v113 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
								F_fmgr_info_cxt(m, v113, l0+int32(16), l3)
								mBase = m.M
								v117 = m.ExcPending
								if v117 != 0 {
									return
								} else {
									F_ReleaseCatCache(m, v17)
									mBase = m.M
									v121 = m.ExcPending
									if v121 != 0 {
										return
									} else {
										m.G0 = v12 + int32(16)
										return
									}
								}
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(97)
							v80 = F_MemoryContextAllocZero(m, l3, int32(64))
							mBase = m.M
							v81 = m.ExcPending
							if v81 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v80
								v83 = *(*int32)(unsafe.Add(mBase, uint32(v21)+92))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = l2
								*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v83
								if l4 == int32(0) {
									F_ReleaseCatCache(m, v17)
									mBase = m.M
									v121 = m.ExcPending
									if v121 != 0 {
										return
									} else {
										m.G0 = v12 + int32(16)
										return
									}
								} else {
									F_getTypeInputInfo(m, l1, v12+int32(8), l0+int32(12))
									mBase = m.M
									v112 = m.ExcPending
									if v112 != 0 {
										return
									} else {
										v113 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
										F_fmgr_info_cxt(m, v113, l0+int32(16), l3)
										mBase = m.M
										v117 = m.ExcPending
										if v117 != 0 {
											return
										} else {
											F_ReleaseCatCache(m, v17)
											mBase = m.M
											v121 = m.ExcPending
											if v121 != 0 {
												return
											} else {
												m.G0 = v12 + int32(16)
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
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v93 = m.ExcPending
			if v93 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v12))) = l1
				F_errmsg_internal(m, int32(_a_F_prepare_column_cache_1), v12)
				mBase = m.M
				v97 = m.ExcPending
				if v97 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_prepare_column_cache_2), int32(3265), int32(_a_F_prepare_column_cache_3))
					mBase = m.M
					v102 = m.ExcPending
					if v102 != 0 {
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
func F_resolve_column_ref(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
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
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
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
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v98 int32
	_ = v98
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v186 int32
	_ = v186
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v204 int32
	_ = v204
	var v218 int32
	_ = v218
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v251 int32
	_ = v251
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v309 int32
	_ = v309
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v326 int32
	_ = v326
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v368 int32
	_ = v368
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v413 int32
	_ = v413
	var v418 int32
	_ = v418
	v5 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(32)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v20 == v5 {
		v418 = v5
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v18 + int32(32)
	return v418
L2:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+528))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	switch v25 - int32(1) {
	case 0:
		goto L8
	case 1:
		goto L7
	case 2:
		goto L6
	default:
		v418 = v5
		goto L1
	}
L3:
	;
	v86 = int32(0)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v90 = v18 + int32(28)
	if v87 == v86 {
		goto L19
	} else {
		goto L20
	}
L4:
	;
	v78 = v70
	v79 = v71
	v80 = v72
	v81 = v73
	v82 = v74
	v83 = v75
	v84 = v76
	v85 = v73
	goto L3
L5:
	;
	v70 = v64
	v71 = v65
	v72 = v66
	v73 = v67
	v74 = v5
	v75 = v5
	v76 = int32(0)
	goto L4
L6:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	if v54 == int32(77) {
		goto L12
	} else {
		goto L13
	}
L7:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	if v38 == int32(77) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	v31 = int32(1)
	v78 = v30
	v79 = v5
	v80 = v31
	v81 = v5
	v82 = v5
	v83 = v5
	v84 = v31
	v85 = int32(0)
	goto L3
L9:
	;
	v64 = v36
	v65 = v5
	v66 = int32(1)
	v67 = int32(_a_F_resolve_column_ref_0)
	goto L5
L10:
	;
	goto L11
L11:
	;
	v44 = int32(2)
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	v78 = v36
	v79 = v5
	v80 = v44
	v81 = v36
	v82 = int32(1)
	v83 = v47
	v84 = v44
	v85 = v47
	goto L3
L12:
	;
	v64 = v52
	v65 = int32(_a_F_resolve_column_ref_0)
	v66 = int32(2)
	v67 = v50
	goto L5
L13:
	;
	goto L14
L14:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	v70 = v52
	v71 = v61
	v72 = v5
	v73 = v50
	v74 = int32(2)
	v75 = v61
	v76 = int32(0)
	goto L4
L15:
	;
	if v228 == int32(0) {
		v418 = v86
		goto L1
	} else {
		goto L52
	}
L16:
	;
	v228 = v218
	goto L15
L17:
	;
	v218 = v204
	goto L16
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v90))) = v196
	v204 = v194
	goto L17
L19:
	;
	v186 = int32(0)
	if v90 == v186 {
		v218 = v186
		goto L16
	} else {
		goto L51
	}
L20:
	;
	v98 = v87
	goto L21
L21:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	if v107 != 0 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	goto L19
L23:
	;
	v112 = v98
	v114 = v107
	goto L26
L24:
	;
	v134 = v98
	goto L25
L25:
	;
	if v85 == int32(0) {
		v171 = v134
		goto L33
	} else {
		goto L34
	}
L26:
	;
	v119 = F_strcmp(m, v112+int32(12), v78)
	mBase = m.M
	if v119 != 0 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v134 = v128
	goto L25
L28:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v112)+8))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v128)))
	if v129 != 0 {
		v112 = v128
		v114 = v129
		goto L26
	} else {
		goto L32
	}
L29:
	;
	if base.B2i32(v114 == int32(1))&base.B2i32(v85 != int32(0)) != 0 {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	if v90 == int32(0) {
		v204 = v112
		goto L17
	} else {
		goto L31
	}
L31:
	;
	v194 = v112
	v196 = int32(1)
	goto L18
L32:
	;
	goto L27
L33:
	;
	goto L49
L34:
	;
	v143 = F_strcmp(m, v134+int32(12), v78)
	mBase = m.M
	if v143 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v144 = v134
	goto L37
L36:
	;
	v144 = v98
	goto L37
L37:
	;
	if v143 != 0 {
		v171 = v144
		goto L33
	} else {
		goto L38
	}
L38:
	;
	if v107 == int32(0) {
		v171 = v144
		goto L33
	} else {
		goto L39
	}
L39:
	;
	v147 = v98
	v154 = v107
	goto L40
L40:
	;
	v158 = F_strcmp(m, v147+int32(12), v85)
	mBase = m.M
	if v158 != 0 {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v171 = v165
	goto L33
L42:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v147)+8))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v165)))
	if v166 != 0 {
		v147 = v165
		v154 = v166
		goto L40
	} else {
		goto L48
	}
L43:
	;
	if base.B2i32(v79 != int32(0))&base.B2i32(v154 == int32(1)) != 0 {
		goto L42
	} else {
		goto L44
	}
L44:
	;
	if v90 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v228 = v147
	goto L15
L46:
	;
	goto L47
L47:
	;
	v194 = v147
	v196 = int32(2)
	goto L18
L48:
	;
	goto L41
L49:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v171)+8))
	if v176 != 0 {
		v98 = v176
		goto L21
	} else {
		goto L50
	}
L50:
	;
	goto L22
L51:
	;
	v194 = v186
	v196 = v186
	goto L18
L52:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v228)))
	switch v231 - int32(1) {
	case 0:
		goto L56
	case 1:
		goto L58
	default:
		goto L57
	}
L53:
	;
	if l3 == int32(0) {
		v418 = v86
		goto L1
	} else {
		goto L82
	}
L54:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v336)+528))
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v337)+68))
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v338+v326<<(uint(int32(2))%32))))
	v343 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v344 = int32(_a_F_resolve_column_ref_1)
	v345 = *(*int32)(unsafe.Add(mBase, _c_F_resolve_column_ref[0]))
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v336)+48))
	*(*int32)(unsafe.Add(mBase, _c_F_resolve_column_ref[0])) = v346
	v348 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v349 = F_bms_add_member(m, v348, v326)
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L74
	} else {
		goto L79
	}
L55:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v228)+4))
	v326 = v320
	goto L54
L56:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v18)+28))
	if v317 != v84 {
		v418 = v86
		goto L1
	} else {
		goto L78
	}
L57:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_resolve_column_ref_2))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L74
	} else {
		goto L75
	}
L58:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v18)+28))
	if v234 == v80 {
		goto L55
	} else {
		goto L59
	}
L59:
	;
	if v234 != v82 {
		v418 = v86
		goto L1
	} else {
		goto L60
	}
L60:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v24)+68))
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v228)+4))
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v237+v238<<(uint(int32(2))%32))))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v242)+32))
	if v243 < int32(0) {
		goto L53
	} else {
		goto L61
	}
L61:
	;
	v251 = v243
	goto L62
L62:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v237+v251<<(uint(int32(2))%32))))
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v264)+8))
	v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83))))
	v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v265))))
	if v269 == int32(0) {
		v288 = v268
		v289 = v269
		goto L65
	} else {
		goto L66
	}
L63:
	;
	goto L53
L64:
	;
	if v289-v288 == int32(0) {
		v326 = v251
		goto L54
	} else {
		goto L72
	}
L65:
	;
	goto L64
L66:
	;
	if v268 != v269 {
		v288 = v268
		v289 = v269
		goto L65
	} else {
		goto L67
	}
L67:
	;
	v273 = v265
	v274 = v83
	goto L68
L68:
	;
	v277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v274)+1)))
	v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v273)+1)))
	if v278 == int32(0) {
		v288 = v277
		v289 = v278
		goto L65
	} else {
		goto L70
	}
L69:
	;
	v288 = v277
	v289 = v278
	goto L65
L70:
	;
	v281 = int32(1)
	if v277 == v278 {
		v273 = v273 + v281
		v274 = v274 + v281
		goto L68
	} else {
		goto L71
	}
L71:
	;
	goto L69
L72:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v264)+16))
	if int32(0) <= v293 {
		v251 = v293
		goto L62
	} else {
		goto L73
	}
L73:
	;
	goto L63
L74:
	;
	return int32(0)
L75:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v228)))
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v304
	F_errmsg_internal(m, int32(_a_F_resolve_column_ref_3), v18)
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L74
	} else {
		goto L76
	}
L76:
	;
	F_errfinish(m, int32(_a_F_resolve_column_ref_4), int32(1229), int32(_a_F_resolve_column_ref_5))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L74
	} else {
		goto L77
	}
L77:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L78:
	;
	goto L55
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v349
	*(*int32)(unsafe.Add(mBase, _c_F_resolve_column_ref[0])) = v345
	v354 = F_palloc0(m, int32(28))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L74
	} else {
		goto L80
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v354)+8)) = v326 + int32(1)
	*(*int64)(unsafe.Add(mBase, uint32(v354))) = int64(8)
	F_plpgsql_exec_get_datum_type_info(m, v337, v342, v354+int32(12), v354+int32(16), v354+int32(20))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L74
	} else {
		goto L81
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v354)+24)) = v343
	v418 = v354
	goto L1
L82:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_resolve_column_ref_2))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L74
	} else {
		goto L83
	}
L83:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L74
	} else {
		goto L84
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v83
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v81
	F_errmsg(m, int32(_a_F_resolve_column_ref_6), v18+int32(16))
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L74
	} else {
		goto L85
	}
L85:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	F_parser_errposition(m, l0, v404)
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L74
	} else {
		goto L86
	}
L86:
	;
	F_errfinish(m, int32(_a_F_resolve_column_ref_4), int32(1225), int32(_a_F_resolve_column_ref_5))
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L74
	} else {
		goto L87
	}
L87:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
