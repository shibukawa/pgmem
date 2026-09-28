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
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
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
				v28 = v14
				v29 = v15
				*(*int32)(unsafe.Add(mBase, uint32(v28+v29-int32(4)))) = l2
				v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
				if v36 == int32(2) {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = int32(0)
				} else {
				}
				return
			}
		} else {
			if l2 == int32(0) {
				v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
				if v36 == int32(2) {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = int32(0)
				} else {
				}
				return
			} else {
				if v6 != 0 {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
					v22 = v20
				} else {
					v22 = int32(0)
				}
				v24 = v22 << (uint(int32(2)) % 32)
				v25 = F_palloc0(m, v24)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v25
					v28 = v24
					v29 = v25
					*(*int32)(unsafe.Add(mBase, uint32(v28+v29-int32(4)))) = l2
					v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
					if v36 == int32(2) {
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
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v84 int64
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v98 int64
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
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
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int64
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	v13 = m.G0
	v15 = v13 - int32(32)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v24 = v17 + v18<<(uint(int32(3))%32) + l1*int32(100)
	v26 = v24 - int32(72)
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+17)))
	if v27 != 0 {
		v29 = F_palloc0(m, int32(12))
		mBase = m.M
		v32 = m.ExcPending
		if v32 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v29))) = int32(59)
			v37 = F_getIdentitySequence(m, l0, base.I32_extend16_s(l1), int32(0))
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v29)+4)) = v37
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v26)+68))
				*(*int32)(unsafe.Add(mBase, uint32(v29)+8)) = v40
				v198 = v29
				m.G0 = v15 + int32(32)
				return v198
			}
		}
	} else {
		v42 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
		v43 = *(*int32)(unsafe.Add(mBase, uint32(v26)+68))
		v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+87)))
		if v44 == int32(1) {
			v48 = F_TupleDescGetDefault(m, v17, base.I32_extend16_s(l1))
			mBase = m.M
			v49 = m.ExcPending
			if v49 != 0 {
				return int32(0)
			} else {
				if v48 != 0 {
					v151 = v48
					v158 = F_exprType(m, v151)
					mBase = m.M
					v159 = m.ExcPending
					if v159 != 0 {
						return int32(0)
					} else {
						v163 = F_coerce_to_target_type(m, int32(0), v151, v158, v43, v42, int32(1), int32(2), int32(-1))
						mBase = m.M
						v164 = m.ExcPending
						if v164 != 0 {
							return int32(0)
						} else {
							if v163 != 0 {
								v198 = v163
								m.G0 = v15 + int32(32)
								return v198
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v168 = m.ExcPending
								if v168 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(67141764))
									mBase = m.M
									v171 = m.ExcPending
									if v171 != 0 {
										return int32(0)
									} else {
										v172 = F_format_type_be(m, v43)
										mBase = m.M
										v173 = m.ExcPending
										if v173 != 0 {
											return int32(0)
										} else {
											v174 = F_format_type_be(m, v158)
											mBase = m.M
											v175 = m.ExcPending
											if v175 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = v174
												*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v172
												*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v26 + int32(4)
												F_errmsg(m, int32(_a_F_build_column_default_0), v15+int32(16))
												mBase = m.M
												v185 = m.ExcPending
												if v185 != 0 {
													return int32(0)
												} else {
													F_errhint(m, int32(_a_F_build_column_default_1), int32(0))
													mBase = m.M
													v189 = m.ExcPending
													if v189 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_build_column_default_2), int32(1340), int32(_a_F_build_column_default_3))
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
										}
									}
								}
							}
						}
					}
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return int32(0)
					} else {
						v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
						*(*int32)(unsafe.Add(mBase, uint32(v15))) = l1
						*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v54 + int32(4)
						F_errmsg_internal(m, int32(_a_F_build_column_default_4), v15)
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_build_column_default_2), int32(1304), int32(_a_F_build_column_default_3))
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
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
			v67 = int32(0)
			v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+90)))
			if v68 != 0 {
				v198 = v67
				m.G0 = v15 + int32(32)
				return v198
			} else {
				v70 = m.G0
				v72 = v70 - int32(16)
				m.G0 = v72
				v76 = F_SearchSysCache1(m, int32(82), base.I64_extend_i32_u(v43))
				mBase = m.M
				v77 = m.ExcPending
				if v77 != 0 {
					return int32(0)
				} else {
					if v76 != 0 {
						v78 = *(*int32)(unsafe.Add(mBase, uint32(v76)+16))
						v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+22)))
						v84 = F_SysCacheGetAttr(m, int32(82), v76, int32(30), v72+int32(15))
						mBase = m.M
						v85 = m.ExcPending
						if v85 != 0 {
							return int32(0)
						} else {
							v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+15)))
							if v86 == int32(0) {
								v90 = F_text_to_cstring(m, base.I32_wrap_i64(v84))
								mBase = m.M
								v91 = m.ExcPending
								if v91 != 0 {
									return int32(0)
								} else {
									v92 = F_stringToNode(m, v90)
									mBase = m.M
									v93 = m.ExcPending
									if v93 != 0 {
										return int32(0)
									} else {
										v124 = v92
										F_ReleaseCatCache(m, v76)
										mBase = m.M
										v130 = m.ExcPending
										if v130 != 0 {
											return int32(0)
										} else {
											m.G0 = v72 + int32(16)
											if v124 == int32(0) {
												v198 = v67
												m.G0 = v15 + int32(32)
												return v198
											} else {
												v151 = v124
												v158 = F_exprType(m, v151)
												mBase = m.M
												v159 = m.ExcPending
												if v159 != 0 {
													return int32(0)
												} else {
													v163 = F_coerce_to_target_type(m, int32(0), v151, v158, v43, v42, int32(1), int32(2), int32(-1))
													mBase = m.M
													v164 = m.ExcPending
													if v164 != 0 {
														return int32(0)
													} else {
														if v163 != 0 {
															v198 = v163
															m.G0 = v15 + int32(32)
															return v198
														} else {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v168 = m.ExcPending
															if v168 != 0 {
																return int32(0)
															} else {
																F_errcode(m, int32(67141764))
																mBase = m.M
																v171 = m.ExcPending
																if v171 != 0 {
																	return int32(0)
																} else {
																	v172 = F_format_type_be(m, v43)
																	mBase = m.M
																	v173 = m.ExcPending
																	if v173 != 0 {
																		return int32(0)
																	} else {
																		v174 = F_format_type_be(m, v158)
																		mBase = m.M
																		v175 = m.ExcPending
																		if v175 != 0 {
																			return int32(0)
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = v174
																			*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v172
																			*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v26 + int32(4)
																			F_errmsg(m, int32(_a_F_build_column_default_0), v15+int32(16))
																			mBase = m.M
																			v185 = m.ExcPending
																			if v185 != 0 {
																				return int32(0)
																			} else {
																				F_errhint(m, int32(_a_F_build_column_default_1), int32(0))
																				mBase = m.M
																				v189 = m.ExcPending
																				if v189 != 0 {
																					return int32(0)
																				} else {
																					F_errfinish(m, int32(_a_F_build_column_default_2), int32(1340), int32(_a_F_build_column_default_3))
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
								v98 = F_SysCacheGetAttr(m, int32(82), v76, int32(31), v72+int32(15))
								mBase = m.M
								v99 = m.ExcPending
								if v99 != 0 {
									return int32(0)
								} else {
									v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+15)))
									if v100 != 0 {
										v124 = int32(0)
										F_ReleaseCatCache(m, v76)
										mBase = m.M
										v130 = m.ExcPending
										if v130 != 0 {
											return int32(0)
										} else {
											m.G0 = v72 + int32(16)
											if v124 == int32(0) {
												v198 = v67
												m.G0 = v15 + int32(32)
												return v198
											} else {
												v151 = v124
												v158 = F_exprType(m, v151)
												mBase = m.M
												v159 = m.ExcPending
												if v159 != 0 {
													return int32(0)
												} else {
													v163 = F_coerce_to_target_type(m, int32(0), v151, v158, v43, v42, int32(1), int32(2), int32(-1))
													mBase = m.M
													v164 = m.ExcPending
													if v164 != 0 {
														return int32(0)
													} else {
														if v163 != 0 {
															v198 = v163
															m.G0 = v15 + int32(32)
															return v198
														} else {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v168 = m.ExcPending
															if v168 != 0 {
																return int32(0)
															} else {
																F_errcode(m, int32(67141764))
																mBase = m.M
																v171 = m.ExcPending
																if v171 != 0 {
																	return int32(0)
																} else {
																	v172 = F_format_type_be(m, v43)
																	mBase = m.M
																	v173 = m.ExcPending
																	if v173 != 0 {
																		return int32(0)
																	} else {
																		v174 = F_format_type_be(m, v158)
																		mBase = m.M
																		v175 = m.ExcPending
																		if v175 != 0 {
																			return int32(0)
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = v174
																			*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v172
																			*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v26 + int32(4)
																			F_errmsg(m, int32(_a_F_build_column_default_0), v15+int32(16))
																			mBase = m.M
																			v185 = m.ExcPending
																			if v185 != 0 {
																				return int32(0)
																			} else {
																				F_errhint(m, int32(_a_F_build_column_default_1), int32(0))
																				mBase = m.M
																				v189 = m.ExcPending
																				if v189 != 0 {
																					return int32(0)
																				} else {
																					F_errfinish(m, int32(_a_F_build_column_default_2), int32(1340), int32(_a_F_build_column_default_3))
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
																	}
																}
															}
														}
													}
												}
											}
										}
									} else {
										v102 = F_text_to_cstring(m, base.I32_wrap_i64(v98))
										mBase = m.M
										v103 = m.ExcPending
										if v103 != 0 {
											return int32(0)
										} else {
											v104 = v78 + v79
											v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)+100))
											v106 = *(*int32)(unsafe.Add(mBase, uint32(v76)+16))
											v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+22)))
											v108 = v106 + v107
											v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)+92))
											if v109 != 0 {
												v111 = v109
											} else {
												v110 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
												v111 = v110
											}
											v113 = F_OidInputFunctionCall(m, v105, v102, v111, int32(-1))
											mBase = m.M
											v114 = m.ExcPending
											if v114 != 0 {
												return int32(0)
											} else {
												v116 = *(*int32)(unsafe.Add(mBase, uint32(v104)+144))
												v117 = int32(*(*int16)(unsafe.Add(mBase, uint32(v104)+76)))
												v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+78)))
												v120 = F_makeConst(m, v43, int32(-1), v116, v117, v113, int32(0), v119)
												mBase = m.M
												v121 = m.ExcPending
												if v121 != 0 {
													return int32(0)
												} else {
													F_pfree(m, v102)
													mBase = m.M
													v123 = m.ExcPending
													if v123 != 0 {
														return int32(0)
													} else {
														v124 = v120
														F_ReleaseCatCache(m, v76)
														mBase = m.M
														v130 = m.ExcPending
														if v130 != 0 {
															return int32(0)
														} else {
															m.G0 = v72 + int32(16)
															if v124 == int32(0) {
																v198 = v67
																m.G0 = v15 + int32(32)
																return v198
															} else {
																v151 = v124
																v158 = F_exprType(m, v151)
																mBase = m.M
																v159 = m.ExcPending
																if v159 != 0 {
																	return int32(0)
																} else {
																	v163 = F_coerce_to_target_type(m, int32(0), v151, v158, v43, v42, int32(1), int32(2), int32(-1))
																	mBase = m.M
																	v164 = m.ExcPending
																	if v164 != 0 {
																		return int32(0)
																	} else {
																		if v163 != 0 {
																			v198 = v163
																			m.G0 = v15 + int32(32)
																			return v198
																		} else {
																			F_errstart_cold(m, int32(21), int32(0))
																			mBase = m.M
																			v168 = m.ExcPending
																			if v168 != 0 {
																				return int32(0)
																			} else {
																				F_errcode(m, int32(67141764))
																				mBase = m.M
																				v171 = m.ExcPending
																				if v171 != 0 {
																					return int32(0)
																				} else {
																					v172 = F_format_type_be(m, v43)
																					mBase = m.M
																					v173 = m.ExcPending
																					if v173 != 0 {
																						return int32(0)
																					} else {
																						v174 = F_format_type_be(m, v158)
																						mBase = m.M
																						v175 = m.ExcPending
																						if v175 != 0 {
																							return int32(0)
																						} else {
																							*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = v174
																							*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v172
																							*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v26 + int32(4)
																							F_errmsg(m, int32(_a_F_build_column_default_0), v15+int32(16))
																							mBase = m.M
																							v185 = m.ExcPending
																							if v185 != 0 {
																								return int32(0)
																							} else {
																								F_errhint(m, int32(_a_F_build_column_default_1), int32(0))
																								mBase = m.M
																								v189 = m.ExcPending
																								if v189 != 0 {
																									return int32(0)
																								} else {
																									F_errfinish(m, int32(_a_F_build_column_default_2), int32(1340), int32(_a_F_build_column_default_3))
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
						v137 = m.ExcPending
						if v137 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v72))) = v43
							F_errmsg_internal(m, int32(_a_F_build_column_default_5), v72)
							mBase = m.M
							v141 = m.ExcPending
							if v141 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_build_column_default_6), int32(2772), int32(_a_F_build_column_default_7))
								mBase = m.M
								v146 = m.ExcPending
								if v146 != 0 {
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
func F_has_column_privilege_id_id_attnum(m *base.Module, l0 int32) int64 {
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
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v43 int64
	_ = v43
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+56)))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v15 = F_pg_detoast_datum_packed(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int64(0)
	} else {
		v20 = F_convert_any_priv_string(m, v15, int32(_a_F_has_column_privilege_id_id_attnum_0))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int64(0)
		} else {
			v22 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)) = uint8(v22)
			if v11 == v22 {
				v38 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v38)
				v43 = int64(0)
				m.G0 = v9 + int32(16)
				return v43
			} else {
				v27 = v9 + int32(15)
				v28 = F_pg_attribute_aclcheck_ext(m, v12, v11, v13, v20, v27)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int64(0)
				} else {
					if v28 != 0 {
						v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
						if v30 != 0 {
							v38 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v38)
							v43 = int64(0)
							m.G0 = v9 + int32(16)
							return v43
						} else {
							v31 = F_pg_class_aclcheck_ext(m, v12, v13, v20, v27)
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
								return int64(0)
							} else {
								if v31 != 0 {
									v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
									if v34 == int32(0) {
									} else {
										v38 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v38)
									}
									v43 = int64(0)
								} else {
									v43 = int64(1)
								}
								m.G0 = v9 + int32(16)
								return v43
							}
						}
					} else {
						v43 = int64(1)
						m.G0 = v9 + int32(16)
						return v43
					}
				}
			}
		}
	}
}
func F_has_column_privilege_name_attnum(m *base.Module, l0 int32) int64 {
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
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v57 int64
	_ = v57
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = F_pg_detoast_datum_packed(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int64(0)
	} else {
		v17 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+40)))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
		v19 = F_pg_detoast_datum_packed(m, v18)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int64(0)
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, _c_F_has_column_privilege_name_attnum[0]))
			v23 = F_textToQualifiedNameList(m, v13)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int64(0)
			} else {
				v25 = F_makeRangeVarFromNameList(m, v23)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int64(0)
				} else {
					v27 = int32(0)
					v31 = F_RangeVarGetRelidExtended(m, v25, v27, v27, v27, v27)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int64(0)
					} else {
						v34 = F_convert_any_priv_string(m, v19, int32(_a_F_has_column_privilege_name_attnum_0))
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int64(0)
						} else {
							v36 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v36)
							if v17 == v36 {
								v52 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v52)
								v57 = int64(0)
								m.G0 = v10 + int32(16)
								return v57
							} else {
								v41 = v10 + int32(15)
								v42 = F_pg_attribute_aclcheck_ext(m, v31, v17, v22, v34, v41)
								mBase = m.M
								v43 = m.ExcPending
								if v43 != 0 {
									return int64(0)
								} else {
									if v42 != 0 {
										v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
										if v44 != 0 {
											v52 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v52)
											v57 = int64(0)
											m.G0 = v10 + int32(16)
											return v57
										} else {
											v45 = F_pg_class_aclcheck_ext(m, v31, v22, v34, v41)
											mBase = m.M
											v46 = m.ExcPending
											if v46 != 0 {
												return int64(0)
											} else {
												if v45 != 0 {
													v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
													if v48 == int32(0) {
													} else {
														v52 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v52)
													}
													v57 = int64(0)
												} else {
													v57 = int64(1)
												}
												m.G0 = v10 + int32(16)
												return v57
											}
										}
									} else {
										v57 = int64(1)
										m.G0 = v10 + int32(16)
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
func F_has_column_privilege_name_name_attnum(m *base.Module, l0 int32) int64 {
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
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v58 int64
	_ = v58
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v14 = F_pg_detoast_datum_packed(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int64(0)
	} else {
		v18 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+56)))
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
		v20 = F_pg_detoast_datum_packed(m, v19)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int64(0)
		} else {
			v22 = F_get_role_oid_or_public(m, v12)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int64(0)
			} else {
				v24 = F_textToQualifiedNameList(m, v14)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int64(0)
				} else {
					v26 = F_makeRangeVarFromNameList(m, v24)
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int64(0)
					} else {
						v28 = int32(0)
						v32 = F_RangeVarGetRelidExtended(m, v26, v28, v28, v28, v28)
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int64(0)
						} else {
							v35 = F_convert_any_priv_string(m, v20, int32(_a_F_has_column_privilege_name_name_attnum_0))
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return int64(0)
							} else {
								v37 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v37)
								if v18 == v37 {
									v53 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v53)
									v58 = int64(0)
									m.G0 = v10 + int32(16)
									return v58
								} else {
									v42 = v10 + int32(15)
									v43 = F_pg_attribute_aclcheck_ext(m, v32, v18, v22, v35, v42)
									mBase = m.M
									v44 = m.ExcPending
									if v44 != 0 {
										return int64(0)
									} else {
										if v43 != 0 {
											v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
											if v45 != 0 {
												v53 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v53)
												v58 = int64(0)
												m.G0 = v10 + int32(16)
												return v58
											} else {
												v46 = F_pg_class_aclcheck_ext(m, v32, v22, v35, v42)
												mBase = m.M
												v47 = m.ExcPending
												if v47 != 0 {
													return int64(0)
												} else {
													if v46 != 0 {
														v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
														if v49 == int32(0) {
														} else {
															v53 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v53)
														}
														v58 = int64(0)
													} else {
														v58 = int64(1)
													}
													m.G0 = v10 + int32(16)
													return v58
												}
											}
										} else {
											v58 = int64(1)
											m.G0 = v10 + int32(16)
											return v58
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
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = l1
	v18 = F_SearchSysCache1(m, int32(82), base.I64_extend_i32_u(l1))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return
	} else {
		if v18 != 0 {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
			v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+22)))
			v22 = v20 + v21
			v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+79)))
			if v23 == int32(100) {
				*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = l2
				v29 = F_getBaseTypeAndTypmod(m, l1, v12+int32(12))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return
				} else {
					v31 = F_get_typtype(m, v29)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return
					} else {
						if v31 == int32(99) {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v29
							v36 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v36
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(67)
							v40 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v36
							v53 = v40
							v54 = v36
							*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v53
							*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v54
							if l4 == int32(0) {
								F_ReleaseCatCache(m, v18)
								mBase = m.M
								v122 = m.ExcPending
								if v122 != 0 {
									return
								} else {
									m.G0 = v12 + int32(16)
									return
								}
							} else {
								F_getTypeInputInfo(m, l1, v12+int32(8), l0+int32(12))
								mBase = m.M
								v113 = m.ExcPending
								if v113 != 0 {
									return
								} else {
									v114 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
									F_fmgr_info_cxt(m, v114, l0+int32(16), l3)
									mBase = m.M
									v118 = m.ExcPending
									if v118 != 0 {
										return
									} else {
										F_ReleaseCatCache(m, v18)
										mBase = m.M
										v122 = m.ExcPending
										if v122 != 0 {
											return
										} else {
											m.G0 = v12 + int32(16)
											return
										}
									}
								}
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v29
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(100)
							v47 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v47
							v51 = F_MemoryContextAllocZero(m, l3, int32(64))
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return
							} else {
								v53 = int32(0)
								v54 = v51
								*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v53
								*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v54
								if l4 == int32(0) {
									F_ReleaseCatCache(m, v18)
									mBase = m.M
									v122 = m.ExcPending
									if v122 != 0 {
										return
									} else {
										m.G0 = v12 + int32(16)
										return
									}
								} else {
									F_getTypeInputInfo(m, l1, v12+int32(8), l0+int32(12))
									mBase = m.M
									v113 = m.ExcPending
									if v113 != 0 {
										return
									} else {
										v114 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
										F_fmgr_info_cxt(m, v114, l0+int32(16), l3)
										mBase = m.M
										v118 = m.ExcPending
										if v118 != 0 {
											return
										} else {
											F_ReleaseCatCache(m, v18)
											mBase = m.M
											v122 = m.ExcPending
											if v122 != 0 {
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
				if base.B2i32(l1 != int32(2249))&base.B2i32(v23 != int32(99)) == int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = l2
					*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = l1
					*(*int64)(unsafe.Add(mBase, uint32(l0)+44)) = int64(0)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(99)
					if l4 == int32(0) {
						F_ReleaseCatCache(m, v18)
						mBase = m.M
						v122 = m.ExcPending
						if v122 != 0 {
							return
						} else {
							m.G0 = v12 + int32(16)
							return
						}
					} else {
						F_getTypeInputInfo(m, l1, v12+int32(8), l0+int32(12))
						mBase = m.M
						v113 = m.ExcPending
						if v113 != 0 {
							return
						} else {
							v114 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
							F_fmgr_info_cxt(m, v114, l0+int32(16), l3)
							mBase = m.M
							v118 = m.ExcPending
							if v118 != 0 {
								return
							} else {
								F_ReleaseCatCache(m, v18)
								mBase = m.M
								v122 = m.ExcPending
								if v122 != 0 {
									return
								} else {
									m.G0 = v12 + int32(16)
									return
								}
							}
						}
					}
				} else {
					v72 = *(*int32)(unsafe.Add(mBase, uint32(v22)+92))
					if v72 == int32(0) {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(115)
						F_getTypeInputInfo(m, l1, v12+int32(8), l0+int32(12))
						mBase = m.M
						v113 = m.ExcPending
						if v113 != 0 {
							return
						} else {
							v114 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
							F_fmgr_info_cxt(m, v114, l0+int32(16), l3)
							mBase = m.M
							v118 = m.ExcPending
							if v118 != 0 {
								return
							} else {
								F_ReleaseCatCache(m, v18)
								mBase = m.M
								v122 = m.ExcPending
								if v122 != 0 {
									return
								} else {
									m.G0 = v12 + int32(16)
									return
								}
							}
						}
					} else {
						v75 = *(*int32)(unsafe.Add(mBase, uint32(v22)+88))
						if v75 != int32(_a_F_prepare_column_cache_0) {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(115)
							F_getTypeInputInfo(m, l1, v12+int32(8), l0+int32(12))
							mBase = m.M
							v113 = m.ExcPending
							if v113 != 0 {
								return
							} else {
								v114 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
								F_fmgr_info_cxt(m, v114, l0+int32(16), l3)
								mBase = m.M
								v118 = m.ExcPending
								if v118 != 0 {
									return
								} else {
									F_ReleaseCatCache(m, v18)
									mBase = m.M
									v122 = m.ExcPending
									if v122 != 0 {
										return
									} else {
										m.G0 = v12 + int32(16)
										return
									}
								}
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(97)
							v81 = F_MemoryContextAllocZero(m, l3, int32(64))
							mBase = m.M
							v82 = m.ExcPending
							if v82 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v81
								v84 = *(*int32)(unsafe.Add(mBase, uint32(v22)+92))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = l2
								*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v84
								if l4 == int32(0) {
									F_ReleaseCatCache(m, v18)
									mBase = m.M
									v122 = m.ExcPending
									if v122 != 0 {
										return
									} else {
										m.G0 = v12 + int32(16)
										return
									}
								} else {
									F_getTypeInputInfo(m, l1, v12+int32(8), l0+int32(12))
									mBase = m.M
									v113 = m.ExcPending
									if v113 != 0 {
										return
									} else {
										v114 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
										F_fmgr_info_cxt(m, v114, l0+int32(16), l3)
										mBase = m.M
										v118 = m.ExcPending
										if v118 != 0 {
											return
										} else {
											F_ReleaseCatCache(m, v18)
											mBase = m.M
											v122 = m.ExcPending
											if v122 != 0 {
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
			v94 = m.ExcPending
			if v94 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v12))) = l1
				F_errmsg_internal(m, int32(_a_F_prepare_column_cache_1), v12)
				mBase = m.M
				v98 = m.ExcPending
				if v98 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_prepare_column_cache_2), int32(3265), int32(_a_F_prepare_column_cache_3))
					mBase = m.M
					v103 = m.ExcPending
					if v103 != 0 {
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
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
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
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
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
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
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
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v85 int32
	_ = v85
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v180 int32
	_ = v180
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v198 int32
	_ = v198
	var v212 int32
	_ = v212
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v244 int32
	_ = v244
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v314 int32
	_ = v314
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v358 int32
	_ = v358
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v388 int32
	_ = v388
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v400 int32
	_ = v400
	var v406 int32
	_ = v406
	v5 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(32)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v19 == v5 {
		v406 = v5
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v17 + int32(32)
	return v406
L2:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+528))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	switch v24 - int32(1) {
	case 0:
		goto L7
	case 1:
		goto L6
	case 2:
		goto L5
	default:
		v406 = v5
		goto L1
	}
L3:
	;
	v73 = int32(0)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v77 = v17 + int32(28)
	if v74 == v73 {
		goto L18
	} else {
		goto L19
	}
L4:
	;
	v65 = v59
	v67 = v61
	v68 = v62
	v69 = v5
	v70 = v5
	v71 = v63
	v72 = int32(0)
	goto L3
L5:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	if v51 == int32(77) {
		goto L11
	} else {
		goto L12
	}
L6:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	if v36 == int32(77) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	v30 = int32(1)
	v65 = v29
	v67 = v30
	v68 = v5
	v69 = v5
	v70 = v5
	v71 = v5
	v72 = v30
	goto L3
L8:
	;
	v59 = v34
	v61 = int32(1)
	v62 = int32(_a_F_resolve_column_ref_0)
	v63 = v5
	goto L4
L9:
	;
	goto L10
L10:
	;
	v41 = int32(2)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	v65 = v34
	v67 = v41
	v68 = v43
	v69 = v43
	v70 = int32(1)
	v71 = v5
	v72 = v41
	goto L3
L11:
	;
	v59 = v49
	v61 = int32(2)
	v62 = v47
	v63 = int32(_a_F_resolve_column_ref_0)
	goto L4
L12:
	;
	goto L13
L13:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	v65 = v49
	v67 = v5
	v68 = v47
	v69 = v57
	v70 = int32(2)
	v71 = v57
	v72 = int32(0)
	goto L3
L14:
	;
	if v222 == int32(0) {
		v406 = v73
		goto L1
	} else {
		goto L50
	}
L15:
	;
	v222 = v212
	goto L14
L16:
	;
	v212 = v198
	goto L15
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v77))) = v190
	v198 = v188
	goto L16
L18:
	;
	v180 = int32(0)
	if v77 == v180 {
		v212 = v180
		goto L15
	} else {
		goto L49
	}
L19:
	;
	v85 = v74
	goto L20
L20:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	if v94 != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	goto L18
L22:
	;
	v99 = v85
	v101 = v94
	goto L25
L23:
	;
	v124 = v85
	goto L24
L24:
	;
	if v68 == int32(0) {
		v165 = v124
		goto L32
	} else {
		goto L33
	}
L25:
	;
	v106 = F_strcmp(m, v99+int32(12), v65)
	mBase = m.M
	v109 = int32(0)
	if v106|base.B2i32(v101 == int32(1))&base.B2i32(v68 != v109) == v109 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v124 = v118
	goto L24
L27:
	;
	if v77 == int32(0) {
		v198 = v99
		goto L16
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v99)+8))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v118)))
	if v119 != 0 {
		v99 = v118
		v101 = v119
		goto L25
	} else {
		goto L31
	}
L30:
	;
	v188 = v99
	v190 = int32(1)
	goto L17
L31:
	;
	goto L26
L32:
	;
	goto L47
L33:
	;
	v133 = F_strcmp(m, v124+int32(12), v65)
	mBase = m.M
	if v133 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v134 = v124
	goto L36
L35:
	;
	v134 = v85
	goto L36
L36:
	;
	if v133|base.B2i32(v94 == int32(0)) != 0 {
		v165 = v134
		goto L32
	} else {
		goto L37
	}
L37:
	;
	v138 = v85
	v145 = v94
	goto L38
L38:
	;
	v149 = F_strcmp(m, v138+int32(12), v68)
	mBase = m.M
	if v149|base.B2i32(v71 != int32(0))&base.B2i32(v145 == int32(1)) == int32(0) {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v165 = v159
	goto L32
L40:
	;
	if v77 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	goto L42
L42:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v138)+8))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v159)))
	if v160 != 0 {
		v138 = v159
		v145 = v160
		goto L38
	} else {
		goto L46
	}
L43:
	;
	v222 = v138
	goto L14
L44:
	;
	goto L45
L45:
	;
	v188 = v138
	v190 = int32(2)
	goto L17
L46:
	;
	goto L39
L47:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v165)+8))
	if v170 != 0 {
		v85 = v170
		goto L20
	} else {
		goto L48
	}
L48:
	;
	goto L21
L49:
	;
	v188 = v180
	v190 = v180
	goto L17
L50:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v222)))
	switch v225 - int32(1) {
	case 0:
		goto L54
	case 1:
		goto L56
	default:
		goto L55
	}
L51:
	;
	if l3 == int32(0) {
		v406 = v73
		goto L1
	} else {
		goto L79
	}
L52:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v324)+528))
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v325)+76))
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v326+v314<<(uint(int32(2))%32))))
	v331 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v332 = int32(_a_F_resolve_column_ref_1)
	v333 = *(*int32)(unsafe.Add(mBase, _c_F_resolve_column_ref[0]))
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v324)+48))
	*(*int32)(unsafe.Add(mBase, _c_F_resolve_column_ref[0])) = v335
	v337 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v338 = F_bms_add_member(m, v337, v314)
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L71
	} else {
		goto L76
	}
L53:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v222)+4))
	v314 = v309
	goto L52
L54:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v17)+28))
	if v306 != v72 {
		v406 = v73
		goto L1
	} else {
		goto L75
	}
L55:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_resolve_column_ref_2))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L71
	} else {
		goto L72
	}
L56:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v17)+28))
	if v67 == v228 {
		goto L53
	} else {
		goto L57
	}
L57:
	;
	if v228 != v70 {
		v406 = v73
		goto L1
	} else {
		goto L58
	}
L58:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v23)+76))
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v222)+4))
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v231+v232<<(uint(int32(2))%32))))
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v236)+32))
	if v237 < int32(0) {
		goto L51
	} else {
		goto L59
	}
L59:
	;
	v244 = v237
	goto L60
L60:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v231+v244<<(uint(int32(2))%32))))
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v257)+8))
	v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258))))
	v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69))))
	if base.B2i32(v261 == int32(0))|base.B2i32(v261 != v264) != 0 {
		v282 = v261
		v283 = v264
		goto L63
	} else {
		goto L64
	}
L61:
	;
	goto L51
L62:
	;
	if v282-v283 == int32(0) {
		v314 = v244
		goto L52
	} else {
		goto L69
	}
L63:
	;
	goto L62
L64:
	;
	v267 = v258
	v268 = v69
	goto L65
L65:
	;
	v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v268)+1)))
	v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v267)+1)))
	if v272 == int32(0) {
		v282 = v272
		v283 = v271
		goto L63
	} else {
		goto L67
	}
L66:
	;
	v282 = v272
	v283 = v271
	goto L63
L67:
	;
	v275 = int32(1)
	if v272 == v271 {
		v267 = v267 + v275
		v268 = v268 + v275
		goto L65
	} else {
		goto L68
	}
L68:
	;
	goto L66
L69:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v257)+16))
	if int32(0) <= v287 {
		v244 = v287
		goto L60
	} else {
		goto L70
	}
L70:
	;
	goto L61
L71:
	;
	return int32(0)
L72:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v222)))
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v296
	F_errmsg_internal(m, int32(_a_F_resolve_column_ref_3), v17)
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L71
	} else {
		goto L73
	}
L73:
	;
	F_errfinish(m, int32(_a_F_resolve_column_ref_4), int32(1238), int32(_a_F_resolve_column_ref_5))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L71
	} else {
		goto L74
	}
L74:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L75:
	;
	goto L53
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v338
	*(*int32)(unsafe.Add(mBase, _c_F_resolve_column_ref[0])) = v333
	v344 = F_palloc0(m, int32(28))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L71
	} else {
		goto L77
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v344)+8)) = v314 + int32(1)
	*(*int64)(unsafe.Add(mBase, uint32(v344))) = int64(8)
	F_plpgsql_exec_get_datum_type_info(m, v325, v330, v344+int32(12), v344+int32(16), v344+int32(20))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L71
	} else {
		goto L78
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v344)+24)) = v331
	v406 = v344
	goto L1
L79:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_resolve_column_ref_2))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L71
	} else {
		goto L80
	}
L80:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L71
	} else {
		goto L81
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v69
	F_errmsg(m, int32(_a_F_resolve_column_ref_6), v17+int32(16))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L71
	} else {
		goto L82
	}
L82:
	;
	F_errhint(m, int32(_a_F_resolve_column_ref_7), int32(0))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L71
	} else {
		goto L83
	}
L83:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	F_parser_errposition(m, l0, v393)
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L71
	} else {
		goto L84
	}
L84:
	;
	F_errfinish(m, int32(_a_F_resolve_column_ref_4), int32(1234), int32(_a_F_resolve_column_ref_5))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L71
	} else {
		goto L85
	}
L85:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
