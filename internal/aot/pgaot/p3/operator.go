package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ValidateOperatorReference(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v14 = F_LookupOperName(m, l0, l1, l2, int32(1))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		if v14 != 0 {
			v18 = F_get_opcode(m, v14)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				v23 = base.B2i32(v18 != int32(0))
				*(*uint8)(unsafe.Add(mBase, uint32(v9+int32(31)))) = uint8(v23)
				if v14 != 0 {
					v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+31)))
					if v25 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(52461700))
							mBase = m.M
							v69 = m.ExcPending
							if v69 != 0 {
								return int32(0)
							} else {
								v70 = F_op_signature_string(m, l0, l1, l2)
								mBase = m.M
								v71 = m.ExcPending
								if v71 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v70
									F_errmsg(m, int32(212249), v9+int32(16))
									mBase = m.M
									v77 = m.ExcPending
									if v77 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(516443), int32(432), int32(435464))
										mBase = m.M
										v82 = m.ExcPending
										if v82 != 0 {
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
						v30 = *(*int32)(unsafe.Add(mBase, _consts[3]))
						v31 = F_object_ownercheck(m, int32(2617), v14, v30)
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return int32(0)
						} else {
							if v31 == int32(0) {
								v37 = F_NameListToString(m, l0)
								mBase = m.M
								v38 = m.ExcPending
								if v38 != 0 {
									return int32(0)
								} else {
									F_aclcheck_error(m, int32(2), int32(25), v37)
									mBase = m.M
									v40 = m.ExcPending
									if v40 != 0 {
										return int32(0)
									} else {
										m.G0 = v9 + int32(32)
										return v14
									}
								}
							} else {
								m.G0 = v9 + int32(32)
								return v14
							}
						}
					}
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(52461700))
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return int32(0)
						} else {
							v52 = F_op_signature_string(m, l0, l1, l2)
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v52
								F_errmsg(m, int32(209123), v9)
								mBase = m.M
								v57 = m.ExcPending
								if v57 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(516443), int32(424), int32(435464))
									mBase = m.M
									v62 = m.ExcPending
									if v62 != 0 {
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
		} else {
			v23 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v9+int32(31)))) = uint8(v23)
			if v14 != 0 {
				v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+31)))
				if v25 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v66 = m.ExcPending
					if v66 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(52461700))
						mBase = m.M
						v69 = m.ExcPending
						if v69 != 0 {
							return int32(0)
						} else {
							v70 = F_op_signature_string(m, l0, l1, l2)
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v70
								F_errmsg(m, int32(212249), v9+int32(16))
								mBase = m.M
								v77 = m.ExcPending
								if v77 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(516443), int32(432), int32(435464))
									mBase = m.M
									v82 = m.ExcPending
									if v82 != 0 {
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
					v30 = *(*int32)(unsafe.Add(mBase, _consts[3]))
					v31 = F_object_ownercheck(m, int32(2617), v14, v30)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						if v31 == int32(0) {
							v37 = F_NameListToString(m, l0)
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return int32(0)
							} else {
								F_aclcheck_error(m, int32(2), int32(25), v37)
								mBase = m.M
								v40 = m.ExcPending
								if v40 != 0 {
									return int32(0)
								} else {
									m.G0 = v9 + int32(32)
									return v14
								}
							}
						} else {
							m.G0 = v9 + int32(32)
							return v14
						}
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(52461700))
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return int32(0)
					} else {
						v52 = F_op_signature_string(m, l0, l1, l2)
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v52
							F_errmsg(m, int32(209123), v9)
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(516443), int32(424), int32(435464))
								mBase = m.M
								v62 = m.ExcPending
								if v62 != 0 {
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
func F_operator_predicate_proof(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
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
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
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
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
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
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v219 int32
	_ = v219
	v5 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	if l0 == v5 {
		v219 = v5
		m.G0 = v16 + int32(16)
		return v219
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v20 != int32(17) {
			v219 = v5
			m.G0 = v16 + int32(16)
			return v219
		} else {
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			if v23 == int32(0) {
				v219 = v5
				m.G0 = v16 + int32(16)
				return v219
			} else {
				if l1 == int32(0) {
					v219 = v5
					m.G0 = v16 + int32(16)
					return v219
				} else {
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
					if v28 != int32(2) {
						v219 = v5
						m.G0 = v16 + int32(16)
						return v219
					} else {
						v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						if v31 != int32(17) {
							v219 = v5
							m.G0 = v16 + int32(16)
							return v219
						} else {
							v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
							if v34 == int32(0) {
								v219 = v5
								m.G0 = v16 + int32(16)
								return v219
							} else {
								v37 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
								if v37 != int32(2) {
									v219 = v5
									m.G0 = v16 + int32(16)
									return v219
								} else {
									v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
									v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
									if v40 != v41 {
										v219 = v5
										m.G0 = v16 + int32(16)
										return v219
									} else {
										v43 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
										v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
										v45 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
										v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
										v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
										v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
										v49 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
										v50 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
										v51 = F_equal(m, v49, v50)
										mBase = m.M
										v54 = m.ExcPending
										if v54 != 0 {
											return int32(0)
										} else {
											v55 = F_equal(m, v46, v44)
											mBase = m.M
											v56 = m.ExcPending
											if v56 != 0 {
												return int32(0)
											} else {
												if v51 != 0 {
													if v55 != 0 {
														v57 = F_operator_same_subexprs_proof(m, v48, v47, l2)
														mBase = m.M
														v58 = m.ExcPending
														if v58 != 0 {
															return int32(0)
														} else {
															v219 = v57
															m.G0 = v16 + int32(16)
															return v219
														}
													} else {
														if v46 == int32(0) {
															v219 = v5
															m.G0 = v16 + int32(16)
															return v219
														} else {
															v61 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
															if v61 != int32(7) {
																v219 = v5
																m.G0 = v16 + int32(16)
																return v219
															} else {
																if v44 == int32(0) {
																	v219 = v5
																	m.G0 = v16 + int32(16)
																	return v219
																} else {
																	v66 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
																	if v66 == int32(7) {
																		v126 = v46
																		v127 = v44
																		v128 = v48
																		v129 = v47
																		v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+24)))
																		if v130 == int32(1) {
																			v133 = F_op_strict(m, v129)
																			mBase = m.M
																			v134 = m.ExcPending
																			if v134 != 0 {
																				return int32(0)
																			} else {
																				if v133 == int32(0) {
																					v219 = v5
																					m.G0 = v16 + int32(16)
																					return v219
																				} else {
																					v137 = int32(1)
																					if l2 != 0 {
																						v219 = v137
																						m.G0 = v16 + int32(16)
																						return v219
																					} else {
																						if l3 == int32(0) {
																							v219 = v137
																							m.G0 = v16 + int32(16)
																							return v219
																						} else {
																							v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126)+24)))
																							if v140 == int32(1) {
																								v143 = F_op_strict(m, v128)
																								mBase = m.M
																								v144 = m.ExcPending
																								if v144 != 0 {
																									return int32(0)
																								} else {
																									if v143 != 0 {
																										v219 = v137
																									} else {
																										v219 = int32(0)
																									}
																									m.G0 = v16 + int32(16)
																									return v219
																								}
																							} else {
																								v219 = int32(0)
																								m.G0 = v16 + int32(16)
																								return v219
																							}
																						}
																					}
																				}
																			}
																		} else {
																			v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126)+24)))
																			if v146 == int32(1) {
																				if l3 != 0 {
																					v150 = F_op_strict(m, v128)
																					mBase = m.M
																					v151 = m.ExcPending
																					if v151 != 0 {
																						return int32(0)
																					} else {
																						if v150 != 0 {
																							v219 = int32(1)
																						} else {
																							v219 = int32(0)
																						}
																						m.G0 = v16 + int32(16)
																						return v219
																					}
																				} else {
																					v219 = int32(0)
																					m.G0 = v16 + int32(16)
																					return v219
																				}
																			} else {
																				v154 = F_lookup_proof_cache(m, v128, v129, l2)
																				mBase = m.M
																				v155 = m.ExcPending
																				if v155 != 0 {
																					return int32(0)
																				} else {
																					if l2 != 0 {
																						v158 = int32(16)
																					} else {
																						v158 = int32(12)
																					}
																					v160 = *(*int32)(unsafe.Add(mBase, uint32(v154+v158)))
																					if v160 == int32(0) {
																						v219 = v5
																						m.G0 = v16 + int32(16)
																						return v219
																					} else {
																						v163 = F_CreateExecutorState(m)
																						mBase = m.M
																						v164 = m.ExcPending
																						if v164 != 0 {
																							return int32(0)
																						} else {
																							v165 = int32(4553888)
																							v166 = *(*int32)(unsafe.Add(mBase, _consts[0]))
																							v168 = *(*int32)(unsafe.Add(mBase, uint32(v163)+100))
																							*(*int32)(unsafe.Add(mBase, _consts[0])) = v168
																							v170 = F_make_opclause(m, v160, v126, v127, v40)
																							mBase = m.M
																							v171 = m.ExcPending
																							if v171 != 0 {
																								return int32(0)
																							} else {
																								F_fix_opfuncids(m, v170)
																								mBase = m.M
																								v173 = m.ExcPending
																								if v173 != 0 {
																									return int32(0)
																								} else {
																									v175 = F_ExecInitExpr(m, v170, int32(0))
																									mBase = m.M
																									v176 = m.ExcPending
																									if v176 != 0 {
																										return int32(0)
																									} else {
																										v177 = *(*int32)(unsafe.Add(mBase, uint32(v163)+152))
																										if v177 == int32(0) {
																											v180 = F_MakePerTupleExprContext(m, v163)
																											mBase = m.M
																											v181 = m.ExcPending
																											if v181 != 0 {
																												return int32(0)
																											} else {
																												v182 = v180
																												v184 = *(*int32)(unsafe.Add(mBase, uint32(v182)+20))
																												*(*int32)(unsafe.Add(mBase, _consts[0])) = v184
																												v188 = *(*int32)(unsafe.Add(mBase, uint32(v175)+20))
																												v189 = m.T0[v188].(func(*base.Module, int32, int32, int32) int32)(m, v175, v182, v16+int32(15))
																												mBase = m.M
																												v190 = m.ExcPending
																												if v190 != 0 {
																													return int32(0)
																												} else {
																													*(*int32)(unsafe.Add(mBase, _consts[0])) = v166
																													F_FreeExecutorState(m, v163)
																													mBase = m.M
																													v194 = m.ExcPending
																													if v194 != 0 {
																														return int32(0)
																													} else {
																														v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+15)))
																														if v195 == int32(1) {
																															v198 = int32(0)
																															v201 = F_errstart(m, int32(13), v198)
																															mBase = m.M
																															v202 = m.ExcPending
																															if v202 != 0 {
																																return int32(0)
																															} else {
																																if v201 == int32(0) {
																																	v219 = v198
																																	m.G0 = v16 + int32(16)
																																	return v219
																																} else {
																																	F_errmsg_internal(m, int32(104197), int32(0))
																																	mBase = m.M
																																	v208 = m.ExcPending
																																	if v208 != 0 {
																																		return int32(0)
																																	} else {
																																		F_errfinish(m, int32(514226), int32(2017), int32(354118))
																																		mBase = m.M
																																		v213 = m.ExcPending
																																		if v213 != 0 {
																																			return int32(0)
																																		} else {
																																			v219 = v198
																																			m.G0 = v16 + int32(16)
																																			return v219
																																		}
																																	}
																																}
																															}
																														} else {
																															v219 = base.B2i32(v189 != int32(0))
																															m.G0 = v16 + int32(16)
																															return v219
																														}
																													}
																												}
																											}
																										} else {
																											v182 = v177
																											v184 = *(*int32)(unsafe.Add(mBase, uint32(v182)+20))
																											*(*int32)(unsafe.Add(mBase, _consts[0])) = v184
																											v188 = *(*int32)(unsafe.Add(mBase, uint32(v175)+20))
																											v189 = m.T0[v188].(func(*base.Module, int32, int32, int32) int32)(m, v175, v182, v16+int32(15))
																											mBase = m.M
																											v190 = m.ExcPending
																											if v190 != 0 {
																												return int32(0)
																											} else {
																												*(*int32)(unsafe.Add(mBase, _consts[0])) = v166
																												F_FreeExecutorState(m, v163)
																												mBase = m.M
																												v194 = m.ExcPending
																												if v194 != 0 {
																													return int32(0)
																												} else {
																													v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+15)))
																													if v195 == int32(1) {
																														v198 = int32(0)
																														v201 = F_errstart(m, int32(13), v198)
																														mBase = m.M
																														v202 = m.ExcPending
																														if v202 != 0 {
																															return int32(0)
																														} else {
																															if v201 == int32(0) {
																																v219 = v198
																																m.G0 = v16 + int32(16)
																																return v219
																															} else {
																																F_errmsg_internal(m, int32(104197), int32(0))
																																mBase = m.M
																																v208 = m.ExcPending
																																if v208 != 0 {
																																	return int32(0)
																																} else {
																																	F_errfinish(m, int32(514226), int32(2017), int32(354118))
																																	mBase = m.M
																																	v213 = m.ExcPending
																																	if v213 != 0 {
																																		return int32(0)
																																	} else {
																																		v219 = v198
																																		m.G0 = v16 + int32(16)
																																		return v219
																																	}
																																}
																															}
																														}
																													} else {
																														v219 = base.B2i32(v189 != int32(0))
																														m.G0 = v16 + int32(16)
																														return v219
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
																		v219 = v5
																		m.G0 = v16 + int32(16)
																		return v219
																	}
																}
															}
														}
													}
												} else {
													if v55 != 0 {
														if v49 == int32(0) {
															v219 = v5
															m.G0 = v16 + int32(16)
															return v219
														} else {
															if v50 == int32(0) {
																v219 = v5
																m.G0 = v16 + int32(16)
																return v219
															} else {
																v73 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
																if v73 != int32(7) {
																	v219 = v5
																	m.G0 = v16 + int32(16)
																	return v219
																} else {
																	v76 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
																	if v76 != int32(7) {
																		v219 = v5
																		m.G0 = v16 + int32(16)
																		return v219
																	} else {
																		v79 = F_get_commutator(m, v48)
																		mBase = m.M
																		v80 = m.ExcPending
																		if v80 != 0 {
																			return int32(0)
																		} else {
																			if v79 == int32(0) {
																				v219 = v5
																				m.G0 = v16 + int32(16)
																				return v219
																			} else {
																				v83 = F_get_commutator(m, v47)
																				mBase = m.M
																				v84 = m.ExcPending
																				if v84 != 0 {
																					return int32(0)
																				} else {
																					if v83 != 0 {
																						v126 = v49
																						v127 = v50
																						v128 = v79
																						v129 = v83
																						v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+24)))
																						if v130 == int32(1) {
																							v133 = F_op_strict(m, v129)
																							mBase = m.M
																							v134 = m.ExcPending
																							if v134 != 0 {
																								return int32(0)
																							} else {
																								if v133 == int32(0) {
																									v219 = v5
																									m.G0 = v16 + int32(16)
																									return v219
																								} else {
																									v137 = int32(1)
																									if l2 != 0 {
																										v219 = v137
																										m.G0 = v16 + int32(16)
																										return v219
																									} else {
																										if l3 == int32(0) {
																											v219 = v137
																											m.G0 = v16 + int32(16)
																											return v219
																										} else {
																											v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126)+24)))
																											if v140 == int32(1) {
																												v143 = F_op_strict(m, v128)
																												mBase = m.M
																												v144 = m.ExcPending
																												if v144 != 0 {
																													return int32(0)
																												} else {
																													if v143 != 0 {
																														v219 = v137
																													} else {
																														v219 = int32(0)
																													}
																													m.G0 = v16 + int32(16)
																													return v219
																												}
																											} else {
																												v219 = int32(0)
																												m.G0 = v16 + int32(16)
																												return v219
																											}
																										}
																									}
																								}
																							}
																						} else {
																							v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126)+24)))
																							if v146 == int32(1) {
																								if l3 != 0 {
																									v150 = F_op_strict(m, v128)
																									mBase = m.M
																									v151 = m.ExcPending
																									if v151 != 0 {
																										return int32(0)
																									} else {
																										if v150 != 0 {
																											v219 = int32(1)
																										} else {
																											v219 = int32(0)
																										}
																										m.G0 = v16 + int32(16)
																										return v219
																									}
																								} else {
																									v219 = int32(0)
																									m.G0 = v16 + int32(16)
																									return v219
																								}
																							} else {
																								v154 = F_lookup_proof_cache(m, v128, v129, l2)
																								mBase = m.M
																								v155 = m.ExcPending
																								if v155 != 0 {
																									return int32(0)
																								} else {
																									if l2 != 0 {
																										v158 = int32(16)
																									} else {
																										v158 = int32(12)
																									}
																									v160 = *(*int32)(unsafe.Add(mBase, uint32(v154+v158)))
																									if v160 == int32(0) {
																										v219 = v5
																										m.G0 = v16 + int32(16)
																										return v219
																									} else {
																										v163 = F_CreateExecutorState(m)
																										mBase = m.M
																										v164 = m.ExcPending
																										if v164 != 0 {
																											return int32(0)
																										} else {
																											v165 = int32(4553888)
																											v166 = *(*int32)(unsafe.Add(mBase, _consts[0]))
																											v168 = *(*int32)(unsafe.Add(mBase, uint32(v163)+100))
																											*(*int32)(unsafe.Add(mBase, _consts[0])) = v168
																											v170 = F_make_opclause(m, v160, v126, v127, v40)
																											mBase = m.M
																											v171 = m.ExcPending
																											if v171 != 0 {
																												return int32(0)
																											} else {
																												F_fix_opfuncids(m, v170)
																												mBase = m.M
																												v173 = m.ExcPending
																												if v173 != 0 {
																													return int32(0)
																												} else {
																													v175 = F_ExecInitExpr(m, v170, int32(0))
																													mBase = m.M
																													v176 = m.ExcPending
																													if v176 != 0 {
																														return int32(0)
																													} else {
																														v177 = *(*int32)(unsafe.Add(mBase, uint32(v163)+152))
																														if v177 == int32(0) {
																															v180 = F_MakePerTupleExprContext(m, v163)
																															mBase = m.M
																															v181 = m.ExcPending
																															if v181 != 0 {
																																return int32(0)
																															} else {
																																v182 = v180
																																v184 = *(*int32)(unsafe.Add(mBase, uint32(v182)+20))
																																*(*int32)(unsafe.Add(mBase, _consts[0])) = v184
																																v188 = *(*int32)(unsafe.Add(mBase, uint32(v175)+20))
																																v189 = m.T0[v188].(func(*base.Module, int32, int32, int32) int32)(m, v175, v182, v16+int32(15))
																																mBase = m.M
																																v190 = m.ExcPending
																																if v190 != 0 {
																																	return int32(0)
																																} else {
																																	*(*int32)(unsafe.Add(mBase, _consts[0])) = v166
																																	F_FreeExecutorState(m, v163)
																																	mBase = m.M
																																	v194 = m.ExcPending
																																	if v194 != 0 {
																																		return int32(0)
																																	} else {
																																		v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+15)))
																																		if v195 == int32(1) {
																																			v198 = int32(0)
																																			v201 = F_errstart(m, int32(13), v198)
																																			mBase = m.M
																																			v202 = m.ExcPending
																																			if v202 != 0 {
																																				return int32(0)
																																			} else {
																																				if v201 == int32(0) {
																																					v219 = v198
																																					m.G0 = v16 + int32(16)
																																					return v219
																																				} else {
																																					F_errmsg_internal(m, int32(104197), int32(0))
																																					mBase = m.M
																																					v208 = m.ExcPending
																																					if v208 != 0 {
																																						return int32(0)
																																					} else {
																																						F_errfinish(m, int32(514226), int32(2017), int32(354118))
																																						mBase = m.M
																																						v213 = m.ExcPending
																																						if v213 != 0 {
																																							return int32(0)
																																						} else {
																																							v219 = v198
																																							m.G0 = v16 + int32(16)
																																							return v219
																																						}
																																					}
																																				}
																																			}
																																		} else {
																																			v219 = base.B2i32(v189 != int32(0))
																																			m.G0 = v16 + int32(16)
																																			return v219
																																		}
																																	}
																																}
																															}
																														} else {
																															v182 = v177
																															v184 = *(*int32)(unsafe.Add(mBase, uint32(v182)+20))
																															*(*int32)(unsafe.Add(mBase, _consts[0])) = v184
																															v188 = *(*int32)(unsafe.Add(mBase, uint32(v175)+20))
																															v189 = m.T0[v188].(func(*base.Module, int32, int32, int32) int32)(m, v175, v182, v16+int32(15))
																															mBase = m.M
																															v190 = m.ExcPending
																															if v190 != 0 {
																																return int32(0)
																															} else {
																																*(*int32)(unsafe.Add(mBase, _consts[0])) = v166
																																F_FreeExecutorState(m, v163)
																																mBase = m.M
																																v194 = m.ExcPending
																																if v194 != 0 {
																																	return int32(0)
																																} else {
																																	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+15)))
																																	if v195 == int32(1) {
																																		v198 = int32(0)
																																		v201 = F_errstart(m, int32(13), v198)
																																		mBase = m.M
																																		v202 = m.ExcPending
																																		if v202 != 0 {
																																			return int32(0)
																																		} else {
																																			if v201 == int32(0) {
																																				v219 = v198
																																				m.G0 = v16 + int32(16)
																																				return v219
																																			} else {
																																				F_errmsg_internal(m, int32(104197), int32(0))
																																				mBase = m.M
																																				v208 = m.ExcPending
																																				if v208 != 0 {
																																					return int32(0)
																																				} else {
																																					F_errfinish(m, int32(514226), int32(2017), int32(354118))
																																					mBase = m.M
																																					v213 = m.ExcPending
																																					if v213 != 0 {
																																						return int32(0)
																																					} else {
																																						v219 = v198
																																						m.G0 = v16 + int32(16)
																																						return v219
																																					}
																																				}
																																			}
																																		}
																																	} else {
																																		v219 = base.B2i32(v189 != int32(0))
																																		m.G0 = v16 + int32(16)
																																		return v219
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
																						v219 = v5
																						m.G0 = v16 + int32(16)
																						return v219
																					}
																				}
																			}
																		}
																	}
																}
															}
														}
													} else {
														v85 = F_equal(m, v49, v44)
														mBase = m.M
														v86 = m.ExcPending
														if v86 != 0 {
															return int32(0)
														} else {
															v87 = F_equal(m, v46, v50)
															mBase = m.M
															v88 = m.ExcPending
															if v88 != 0 {
																return int32(0)
															} else {
																if v85 != 0 {
																	if v87 != 0 {
																		v89 = F_get_commutator(m, v48)
																		mBase = m.M
																		v90 = m.ExcPending
																		if v90 != 0 {
																			return int32(0)
																		} else {
																			if v89 == int32(0) {
																				v219 = v5
																				m.G0 = v16 + int32(16)
																				return v219
																			} else {
																				v93 = F_operator_same_subexprs_proof(m, v89, v47, l2)
																				mBase = m.M
																				v94 = m.ExcPending
																				if v94 != 0 {
																					return int32(0)
																				} else {
																					v219 = v93
																					m.G0 = v16 + int32(16)
																					return v219
																				}
																			}
																		}
																	} else {
																		if v46 == int32(0) {
																			v219 = v5
																			m.G0 = v16 + int32(16)
																			return v219
																		} else {
																			if v50 == int32(0) {
																				v219 = v5
																				m.G0 = v16 + int32(16)
																				return v219
																			} else {
																				v99 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
																				if v99 != int32(7) {
																					v219 = v5
																					m.G0 = v16 + int32(16)
																					return v219
																				} else {
																					v102 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
																					if v102 != int32(7) {
																						v219 = v5
																						m.G0 = v16 + int32(16)
																						return v219
																					} else {
																						v105 = F_get_commutator(m, v47)
																						mBase = m.M
																						v106 = m.ExcPending
																						if v106 != 0 {
																							return int32(0)
																						} else {
																							if v105 != 0 {
																								v126 = v46
																								v127 = v50
																								v128 = v48
																								v129 = v105
																								v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+24)))
																								if v130 == int32(1) {
																									v133 = F_op_strict(m, v129)
																									mBase = m.M
																									v134 = m.ExcPending
																									if v134 != 0 {
																										return int32(0)
																									} else {
																										if v133 == int32(0) {
																											v219 = v5
																											m.G0 = v16 + int32(16)
																											return v219
																										} else {
																											v137 = int32(1)
																											if l2 != 0 {
																												v219 = v137
																												m.G0 = v16 + int32(16)
																												return v219
																											} else {
																												if l3 == int32(0) {
																													v219 = v137
																													m.G0 = v16 + int32(16)
																													return v219
																												} else {
																													v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126)+24)))
																													if v140 == int32(1) {
																														v143 = F_op_strict(m, v128)
																														mBase = m.M
																														v144 = m.ExcPending
																														if v144 != 0 {
																															return int32(0)
																														} else {
																															if v143 != 0 {
																																v219 = v137
																															} else {
																																v219 = int32(0)
																															}
																															m.G0 = v16 + int32(16)
																															return v219
																														}
																													} else {
																														v219 = int32(0)
																														m.G0 = v16 + int32(16)
																														return v219
																													}
																												}
																											}
																										}
																									}
																								} else {
																									v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126)+24)))
																									if v146 == int32(1) {
																										if l3 != 0 {
																											v150 = F_op_strict(m, v128)
																											mBase = m.M
																											v151 = m.ExcPending
																											if v151 != 0 {
																												return int32(0)
																											} else {
																												if v150 != 0 {
																													v219 = int32(1)
																												} else {
																													v219 = int32(0)
																												}
																												m.G0 = v16 + int32(16)
																												return v219
																											}
																										} else {
																											v219 = int32(0)
																											m.G0 = v16 + int32(16)
																											return v219
																										}
																									} else {
																										v154 = F_lookup_proof_cache(m, v128, v129, l2)
																										mBase = m.M
																										v155 = m.ExcPending
																										if v155 != 0 {
																											return int32(0)
																										} else {
																											if l2 != 0 {
																												v158 = int32(16)
																											} else {
																												v158 = int32(12)
																											}
																											v160 = *(*int32)(unsafe.Add(mBase, uint32(v154+v158)))
																											if v160 == int32(0) {
																												v219 = v5
																												m.G0 = v16 + int32(16)
																												return v219
																											} else {
																												v163 = F_CreateExecutorState(m)
																												mBase = m.M
																												v164 = m.ExcPending
																												if v164 != 0 {
																													return int32(0)
																												} else {
																													v165 = int32(4553888)
																													v166 = *(*int32)(unsafe.Add(mBase, _consts[0]))
																													v168 = *(*int32)(unsafe.Add(mBase, uint32(v163)+100))
																													*(*int32)(unsafe.Add(mBase, _consts[0])) = v168
																													v170 = F_make_opclause(m, v160, v126, v127, v40)
																													mBase = m.M
																													v171 = m.ExcPending
																													if v171 != 0 {
																														return int32(0)
																													} else {
																														F_fix_opfuncids(m, v170)
																														mBase = m.M
																														v173 = m.ExcPending
																														if v173 != 0 {
																															return int32(0)
																														} else {
																															v175 = F_ExecInitExpr(m, v170, int32(0))
																															mBase = m.M
																															v176 = m.ExcPending
																															if v176 != 0 {
																																return int32(0)
																															} else {
																																v177 = *(*int32)(unsafe.Add(mBase, uint32(v163)+152))
																																if v177 == int32(0) {
																																	v180 = F_MakePerTupleExprContext(m, v163)
																																	mBase = m.M
																																	v181 = m.ExcPending
																																	if v181 != 0 {
																																		return int32(0)
																																	} else {
																																		v182 = v180
																																		v184 = *(*int32)(unsafe.Add(mBase, uint32(v182)+20))
																																		*(*int32)(unsafe.Add(mBase, _consts[0])) = v184
																																		v188 = *(*int32)(unsafe.Add(mBase, uint32(v175)+20))
																																		v189 = m.T0[v188].(func(*base.Module, int32, int32, int32) int32)(m, v175, v182, v16+int32(15))
																																		mBase = m.M
																																		v190 = m.ExcPending
																																		if v190 != 0 {
																																			return int32(0)
																																		} else {
																																			*(*int32)(unsafe.Add(mBase, _consts[0])) = v166
																																			F_FreeExecutorState(m, v163)
																																			mBase = m.M
																																			v194 = m.ExcPending
																																			if v194 != 0 {
																																				return int32(0)
																																			} else {
																																				v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+15)))
																																				if v195 == int32(1) {
																																					v198 = int32(0)
																																					v201 = F_errstart(m, int32(13), v198)
																																					mBase = m.M
																																					v202 = m.ExcPending
																																					if v202 != 0 {
																																						return int32(0)
																																					} else {
																																						if v201 == int32(0) {
																																							v219 = v198
																																							m.G0 = v16 + int32(16)
																																							return v219
																																						} else {
																																							F_errmsg_internal(m, int32(104197), int32(0))
																																							mBase = m.M
																																							v208 = m.ExcPending
																																							if v208 != 0 {
																																								return int32(0)
																																							} else {
																																								F_errfinish(m, int32(514226), int32(2017), int32(354118))
																																								mBase = m.M
																																								v213 = m.ExcPending
																																								if v213 != 0 {
																																									return int32(0)
																																								} else {
																																									v219 = v198
																																									m.G0 = v16 + int32(16)
																																									return v219
																																								}
																																							}
																																						}
																																					}
																																				} else {
																																					v219 = base.B2i32(v189 != int32(0))
																																					m.G0 = v16 + int32(16)
																																					return v219
																																				}
																																			}
																																		}
																																	}
																																} else {
																																	v182 = v177
																																	v184 = *(*int32)(unsafe.Add(mBase, uint32(v182)+20))
																																	*(*int32)(unsafe.Add(mBase, _consts[0])) = v184
																																	v188 = *(*int32)(unsafe.Add(mBase, uint32(v175)+20))
																																	v189 = m.T0[v188].(func(*base.Module, int32, int32, int32) int32)(m, v175, v182, v16+int32(15))
																																	mBase = m.M
																																	v190 = m.ExcPending
																																	if v190 != 0 {
																																		return int32(0)
																																	} else {
																																		*(*int32)(unsafe.Add(mBase, _consts[0])) = v166
																																		F_FreeExecutorState(m, v163)
																																		mBase = m.M
																																		v194 = m.ExcPending
																																		if v194 != 0 {
																																			return int32(0)
																																		} else {
																																			v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+15)))
																																			if v195 == int32(1) {
																																				v198 = int32(0)
																																				v201 = F_errstart(m, int32(13), v198)
																																				mBase = m.M
																																				v202 = m.ExcPending
																																				if v202 != 0 {
																																					return int32(0)
																																				} else {
																																					if v201 == int32(0) {
																																						v219 = v198
																																						m.G0 = v16 + int32(16)
																																						return v219
																																					} else {
																																						F_errmsg_internal(m, int32(104197), int32(0))
																																						mBase = m.M
																																						v208 = m.ExcPending
																																						if v208 != 0 {
																																							return int32(0)
																																						} else {
																																							F_errfinish(m, int32(514226), int32(2017), int32(354118))
																																							mBase = m.M
																																							v213 = m.ExcPending
																																							if v213 != 0 {
																																								return int32(0)
																																							} else {
																																								v219 = v198
																																								m.G0 = v16 + int32(16)
																																								return v219
																																							}
																																						}
																																					}
																																				}
																																			} else {
																																				v219 = base.B2i32(v189 != int32(0))
																																				m.G0 = v16 + int32(16)
																																				return v219
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
																								v219 = v5
																								m.G0 = v16 + int32(16)
																								return v219
																							}
																						}
																					}
																				}
																			}
																		}
																	}
																} else {
																	if base.B2i32(v49 == int32(0))|(v87^int32(1)) != 0 {
																		v219 = v5
																		m.G0 = v16 + int32(16)
																		return v219
																	} else {
																		v112 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
																		if v112 != int32(7) {
																			v219 = v5
																			m.G0 = v16 + int32(16)
																			return v219
																		} else {
																			if v44 == int32(0) {
																				v219 = v5
																				m.G0 = v16 + int32(16)
																				return v219
																			} else {
																				v117 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
																				if v117 != int32(7) {
																					v219 = v5
																					m.G0 = v16 + int32(16)
																					return v219
																				} else {
																					v120 = F_get_commutator(m, v48)
																					mBase = m.M
																					v121 = m.ExcPending
																					if v121 != 0 {
																						return int32(0)
																					} else {
																						if v120 == int32(0) {
																							v219 = v5
																							m.G0 = v16 + int32(16)
																							return v219
																						} else {
																							v126 = v49
																							v127 = v44
																							v128 = v120
																							v129 = v47
																							v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+24)))
																							if v130 == int32(1) {
																								v133 = F_op_strict(m, v129)
																								mBase = m.M
																								v134 = m.ExcPending
																								if v134 != 0 {
																									return int32(0)
																								} else {
																									if v133 == int32(0) {
																										v219 = v5
																										m.G0 = v16 + int32(16)
																										return v219
																									} else {
																										v137 = int32(1)
																										if l2 != 0 {
																											v219 = v137
																											m.G0 = v16 + int32(16)
																											return v219
																										} else {
																											if l3 == int32(0) {
																												v219 = v137
																												m.G0 = v16 + int32(16)
																												return v219
																											} else {
																												v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126)+24)))
																												if v140 == int32(1) {
																													v143 = F_op_strict(m, v128)
																													mBase = m.M
																													v144 = m.ExcPending
																													if v144 != 0 {
																														return int32(0)
																													} else {
																														if v143 != 0 {
																															v219 = v137
																														} else {
																															v219 = int32(0)
																														}
																														m.G0 = v16 + int32(16)
																														return v219
																													}
																												} else {
																													v219 = int32(0)
																													m.G0 = v16 + int32(16)
																													return v219
																												}
																											}
																										}
																									}
																								}
																							} else {
																								v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126)+24)))
																								if v146 == int32(1) {
																									if l3 != 0 {
																										v150 = F_op_strict(m, v128)
																										mBase = m.M
																										v151 = m.ExcPending
																										if v151 != 0 {
																											return int32(0)
																										} else {
																											if v150 != 0 {
																												v219 = int32(1)
																											} else {
																												v219 = int32(0)
																											}
																											m.G0 = v16 + int32(16)
																											return v219
																										}
																									} else {
																										v219 = int32(0)
																										m.G0 = v16 + int32(16)
																										return v219
																									}
																								} else {
																									v154 = F_lookup_proof_cache(m, v128, v129, l2)
																									mBase = m.M
																									v155 = m.ExcPending
																									if v155 != 0 {
																										return int32(0)
																									} else {
																										if l2 != 0 {
																											v158 = int32(16)
																										} else {
																											v158 = int32(12)
																										}
																										v160 = *(*int32)(unsafe.Add(mBase, uint32(v154+v158)))
																										if v160 == int32(0) {
																											v219 = v5
																											m.G0 = v16 + int32(16)
																											return v219
																										} else {
																											v163 = F_CreateExecutorState(m)
																											mBase = m.M
																											v164 = m.ExcPending
																											if v164 != 0 {
																												return int32(0)
																											} else {
																												v165 = int32(4553888)
																												v166 = *(*int32)(unsafe.Add(mBase, _consts[0]))
																												v168 = *(*int32)(unsafe.Add(mBase, uint32(v163)+100))
																												*(*int32)(unsafe.Add(mBase, _consts[0])) = v168
																												v170 = F_make_opclause(m, v160, v126, v127, v40)
																												mBase = m.M
																												v171 = m.ExcPending
																												if v171 != 0 {
																													return int32(0)
																												} else {
																													F_fix_opfuncids(m, v170)
																													mBase = m.M
																													v173 = m.ExcPending
																													if v173 != 0 {
																														return int32(0)
																													} else {
																														v175 = F_ExecInitExpr(m, v170, int32(0))
																														mBase = m.M
																														v176 = m.ExcPending
																														if v176 != 0 {
																															return int32(0)
																														} else {
																															v177 = *(*int32)(unsafe.Add(mBase, uint32(v163)+152))
																															if v177 == int32(0) {
																																v180 = F_MakePerTupleExprContext(m, v163)
																																mBase = m.M
																																v181 = m.ExcPending
																																if v181 != 0 {
																																	return int32(0)
																																} else {
																																	v182 = v180
																																	v184 = *(*int32)(unsafe.Add(mBase, uint32(v182)+20))
																																	*(*int32)(unsafe.Add(mBase, _consts[0])) = v184
																																	v188 = *(*int32)(unsafe.Add(mBase, uint32(v175)+20))
																																	v189 = m.T0[v188].(func(*base.Module, int32, int32, int32) int32)(m, v175, v182, v16+int32(15))
																																	mBase = m.M
																																	v190 = m.ExcPending
																																	if v190 != 0 {
																																		return int32(0)
																																	} else {
																																		*(*int32)(unsafe.Add(mBase, _consts[0])) = v166
																																		F_FreeExecutorState(m, v163)
																																		mBase = m.M
																																		v194 = m.ExcPending
																																		if v194 != 0 {
																																			return int32(0)
																																		} else {
																																			v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+15)))
																																			if v195 == int32(1) {
																																				v198 = int32(0)
																																				v201 = F_errstart(m, int32(13), v198)
																																				mBase = m.M
																																				v202 = m.ExcPending
																																				if v202 != 0 {
																																					return int32(0)
																																				} else {
																																					if v201 == int32(0) {
																																						v219 = v198
																																						m.G0 = v16 + int32(16)
																																						return v219
																																					} else {
																																						F_errmsg_internal(m, int32(104197), int32(0))
																																						mBase = m.M
																																						v208 = m.ExcPending
																																						if v208 != 0 {
																																							return int32(0)
																																						} else {
																																							F_errfinish(m, int32(514226), int32(2017), int32(354118))
																																							mBase = m.M
																																							v213 = m.ExcPending
																																							if v213 != 0 {
																																								return int32(0)
																																							} else {
																																								v219 = v198
																																								m.G0 = v16 + int32(16)
																																								return v219
																																							}
																																						}
																																					}
																																				}
																																			} else {
																																				v219 = base.B2i32(v189 != int32(0))
																																				m.G0 = v16 + int32(16)
																																				return v219
																																			}
																																		}
																																	}
																																}
																															} else {
																																v182 = v177
																																v184 = *(*int32)(unsafe.Add(mBase, uint32(v182)+20))
																																*(*int32)(unsafe.Add(mBase, _consts[0])) = v184
																																v188 = *(*int32)(unsafe.Add(mBase, uint32(v175)+20))
																																v189 = m.T0[v188].(func(*base.Module, int32, int32, int32) int32)(m, v175, v182, v16+int32(15))
																																mBase = m.M
																																v190 = m.ExcPending
																																if v190 != 0 {
																																	return int32(0)
																																} else {
																																	*(*int32)(unsafe.Add(mBase, _consts[0])) = v166
																																	F_FreeExecutorState(m, v163)
																																	mBase = m.M
																																	v194 = m.ExcPending
																																	if v194 != 0 {
																																		return int32(0)
																																	} else {
																																		v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+15)))
																																		if v195 == int32(1) {
																																			v198 = int32(0)
																																			v201 = F_errstart(m, int32(13), v198)
																																			mBase = m.M
																																			v202 = m.ExcPending
																																			if v202 != 0 {
																																				return int32(0)
																																			} else {
																																				if v201 == int32(0) {
																																					v219 = v198
																																					m.G0 = v16 + int32(16)
																																					return v219
																																				} else {
																																					F_errmsg_internal(m, int32(104197), int32(0))
																																					mBase = m.M
																																					v208 = m.ExcPending
																																					if v208 != 0 {
																																						return int32(0)
																																					} else {
																																						F_errfinish(m, int32(514226), int32(2017), int32(354118))
																																						mBase = m.M
																																						v213 = m.ExcPending
																																						if v213 != 0 {
																																							return int32(0)
																																						} else {
																																							v219 = v198
																																							m.G0 = v16 + int32(16)
																																							return v219
																																						}
																																					}
																																				}
																																			}
																																		} else {
																																			v219 = base.B2i32(v189 != int32(0))
																																			m.G0 = v16 + int32(16)
																																			return v219
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
