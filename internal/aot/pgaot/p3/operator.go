package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ValidateOperatorReference(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v13 = F_LookupOperName(m, l0, l1, l2, int32(1))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		if v13 != 0 {
			v17 = F_get_opcode(m, v13)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				v22 = base.B2i32(v17 != int32(0))
				*(*uint8)(unsafe.Add(mBase, uint32(v8+int32(31)))) = uint8(v22)
				if v13 != 0 {
					v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+31)))
					if v24 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(52461700))
							mBase = m.M
							v68 = m.ExcPending
							if v68 != 0 {
								return int32(0)
							} else {
								v69 = F_op_signature_string(m, l0, l1, l2)
								mBase = m.M
								v70 = m.ExcPending
								if v70 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v69
									F_errmsg(m, int32(_a_F_ValidateOperatorReference_0), v8+int32(16))
									mBase = m.M
									v76 = m.ExcPending
									if v76 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_ValidateOperatorReference_1), int32(432), int32(_a_F_ValidateOperatorReference_2))
										mBase = m.M
										v81 = m.ExcPending
										if v81 != 0 {
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
						v29 = *(*int32)(unsafe.Add(mBase, _c_F_ValidateOperatorReference[0]))
						v30 = F_object_ownercheck(m, int32(2617), v13, v29)
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return int32(0)
						} else {
							if v30 == int32(0) {
								v36 = F_NameListToString(m, l0)
								mBase = m.M
								v37 = m.ExcPending
								if v37 != 0 {
									return int32(0)
								} else {
									F_aclcheck_error(m, int32(2), int32(25), v36)
									mBase = m.M
									v39 = m.ExcPending
									if v39 != 0 {
										return int32(0)
									} else {
										m.G0 = v8 + int32(32)
										return v13
									}
								}
							} else {
								m.G0 = v8 + int32(32)
								return v13
							}
						}
					}
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(52461700))
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return int32(0)
						} else {
							v51 = F_op_signature_string(m, l0, l1, l2)
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v8))) = v51
								F_errmsg(m, int32(_a_F_ValidateOperatorReference_3), v8)
								mBase = m.M
								v56 = m.ExcPending
								if v56 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_ValidateOperatorReference_1), int32(424), int32(_a_F_ValidateOperatorReference_2))
									mBase = m.M
									v61 = m.ExcPending
									if v61 != 0 {
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
			v22 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v8+int32(31)))) = uint8(v22)
			if v13 != 0 {
				v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+31)))
				if v24 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v65 = m.ExcPending
					if v65 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(52461700))
						mBase = m.M
						v68 = m.ExcPending
						if v68 != 0 {
							return int32(0)
						} else {
							v69 = F_op_signature_string(m, l0, l1, l2)
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v69
								F_errmsg(m, int32(_a_F_ValidateOperatorReference_0), v8+int32(16))
								mBase = m.M
								v76 = m.ExcPending
								if v76 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_ValidateOperatorReference_1), int32(432), int32(_a_F_ValidateOperatorReference_2))
									mBase = m.M
									v81 = m.ExcPending
									if v81 != 0 {
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
					v29 = *(*int32)(unsafe.Add(mBase, _c_F_ValidateOperatorReference[0]))
					v30 = F_object_ownercheck(m, int32(2617), v13, v29)
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						if v30 == int32(0) {
							v36 = F_NameListToString(m, l0)
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return int32(0)
							} else {
								F_aclcheck_error(m, int32(2), int32(25), v36)
								mBase = m.M
								v39 = m.ExcPending
								if v39 != 0 {
									return int32(0)
								} else {
									m.G0 = v8 + int32(32)
									return v13
								}
							}
						} else {
							m.G0 = v8 + int32(32)
							return v13
						}
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(52461700))
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return int32(0)
					} else {
						v51 = F_op_signature_string(m, l0, l1, l2)
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = v51
							F_errmsg(m, int32(_a_F_ValidateOperatorReference_3), v8)
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_ValidateOperatorReference_1), int32(424), int32(_a_F_ValidateOperatorReference_2))
								mBase = m.M
								v61 = m.ExcPending
								if v61 != 0 {
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
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
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
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v195 int64
	_ = v195
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v229 int32
	_ = v229
	v5 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	if l0 == v5 {
		v229 = v5
		m.G0 = v17 + int32(16)
		return v229
	} else {
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v21 != int32(17) {
			v229 = v5
			m.G0 = v17 + int32(16)
			return v229
		} else {
			v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			v25 = int32(0)
			if base.B2i32(v24 == v25)|base.B2i32(l1 == v25) != 0 {
				v229 = v5
				m.G0 = v17 + int32(16)
				return v229
			} else {
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
				if v30 != int32(2) {
					v229 = v5
					m.G0 = v17 + int32(16)
					return v229
				} else {
					v33 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					if v33 != int32(17) {
						v229 = v5
						m.G0 = v17 + int32(16)
						return v229
					} else {
						v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
						if v36 == int32(0) {
							v229 = v5
							m.G0 = v17 + int32(16)
							return v229
						} else {
							v39 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
							if v39 != int32(2) {
								v229 = v5
								m.G0 = v17 + int32(16)
								return v229
							} else {
								v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
								v43 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
								if v42 != v43 {
									v229 = v5
									m.G0 = v17 + int32(16)
									return v229
								} else {
									v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
									v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									v47 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
									v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
									v49 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
									v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
									v51 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
									v52 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
									v53 = F_equal(m, v51, v52)
									mBase = m.M
									v56 = m.ExcPending
									if v56 != 0 {
										return int32(0)
									} else {
										v57 = F_equal(m, v50, v48)
										mBase = m.M
										v58 = m.ExcPending
										if v58 != 0 {
											return int32(0)
										} else {
											if v53 != 0 {
												if v57 != 0 {
													v59 = F_operator_same_subexprs_proof(m, v46, v45, l2)
													mBase = m.M
													v60 = m.ExcPending
													if v60 != 0 {
														return int32(0)
													} else {
														v229 = v59
														m.G0 = v17 + int32(16)
														return v229
													}
												} else {
													if v50 == int32(0) {
														v229 = v5
														m.G0 = v17 + int32(16)
														return v229
													} else {
														v65 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
														if base.B2i32(v48 == int32(0))|base.B2i32(v65 != int32(7)) != 0 {
															v229 = v5
															m.G0 = v17 + int32(16)
															return v229
														} else {
															v69 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
															if v69 == int32(7) {
																v130 = v50
																v131 = v48
																v132 = v46
																v133 = v45
																v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131)+32)))
																if v135 == int32(1) {
																	v138 = F_op_strict(m, v133)
																	mBase = m.M
																	v139 = m.ExcPending
																	if v139 != 0 {
																		return int32(0)
																	} else {
																		if v138 == int32(0) {
																			v229 = v5
																			m.G0 = v17 + int32(16)
																			return v229
																		} else {
																			v142 = int32(1)
																			if l2|base.B2i32(l3 == int32(0)) != 0 {
																				v229 = v142
																				m.G0 = v17 + int32(16)
																				return v229
																			} else {
																				v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130)+32)))
																				if v146 == int32(1) {
																					v149 = F_op_strict(m, v132)
																					mBase = m.M
																					v150 = m.ExcPending
																					if v150 != 0 {
																						return int32(0)
																					} else {
																						if v149 != 0 {
																							v229 = v142
																						} else {
																							v229 = int32(0)
																						}
																						m.G0 = v17 + int32(16)
																						return v229
																					}
																				} else {
																					v229 = int32(0)
																					m.G0 = v17 + int32(16)
																					return v229
																				}
																			}
																		}
																	}
																} else {
																	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130)+32)))
																	if v152 == int32(1) {
																		if l3 != 0 {
																			v156 = F_op_strict(m, v132)
																			mBase = m.M
																			v157 = m.ExcPending
																			if v157 != 0 {
																				return int32(0)
																			} else {
																				if v156 != 0 {
																					v229 = int32(1)
																				} else {
																					v229 = int32(0)
																				}
																				m.G0 = v17 + int32(16)
																				return v229
																			}
																		} else {
																			v229 = int32(0)
																			m.G0 = v17 + int32(16)
																			return v229
																		}
																	} else {
																		v160 = F_lookup_proof_cache(m, v132, v133, l2)
																		mBase = m.M
																		v161 = m.ExcPending
																		if v161 != 0 {
																			return int32(0)
																		} else {
																			if l2 != 0 {
																				v164 = int32(16)
																			} else {
																				v164 = int32(12)
																			}
																			v166 = *(*int32)(unsafe.Add(mBase, uint32(v160+v164)))
																			if v166 == int32(0) {
																				v229 = v5
																				m.G0 = v17 + int32(16)
																				return v229
																			} else {
																				v169 = F_CreateExecutorState(m)
																				mBase = m.M
																				v170 = m.ExcPending
																				if v170 != 0 {
																					return int32(0)
																				} else {
																					v171 = int32(_a_F_operator_predicate_proof_0)
																					v172 = *(*int32)(unsafe.Add(mBase, _c_F_operator_predicate_proof[0]))
																					v174 = *(*int32)(unsafe.Add(mBase, uint32(v169)+100))
																					*(*int32)(unsafe.Add(mBase, _c_F_operator_predicate_proof[0])) = v174
																					v176 = F_make_opclause(m, v166, v130, v131, v42)
																					mBase = m.M
																					v177 = m.ExcPending
																					if v177 != 0 {
																						return int32(0)
																					} else {
																						F_fix_opfuncids(m, v176)
																						mBase = m.M
																						v179 = m.ExcPending
																						if v179 != 0 {
																							return int32(0)
																						} else {
																							v181 = F_ExecInitExpr(m, v176, int32(0))
																							mBase = m.M
																							v182 = m.ExcPending
																							if v182 != 0 {
																								return int32(0)
																							} else {
																								v183 = *(*int32)(unsafe.Add(mBase, uint32(v169)+152))
																								if v183 == int32(0) {
																									v186 = F_MakePerTupleExprContext(m, v169)
																									mBase = m.M
																									v187 = m.ExcPending
																									if v187 != 0 {
																										return int32(0)
																									} else {
																										v188 = v186
																										v190 = *(*int32)(unsafe.Add(mBase, uint32(v188)+20))
																										*(*int32)(unsafe.Add(mBase, _c_F_operator_predicate_proof[0])) = v190
																										v194 = *(*int32)(unsafe.Add(mBase, uint32(v181)+24))
																										v195 = m.T0[v194].(func(*base.Module, int32, int32, int32) int64)(m, v181, v188, v17+int32(15))
																										mBase = m.M
																										v196 = m.ExcPending
																										if v196 != 0 {
																											return int32(0)
																										} else {
																											*(*int32)(unsafe.Add(mBase, _c_F_operator_predicate_proof[0])) = v172
																											F_FreeExecutorState(m, v169)
																											mBase = m.M
																											v200 = m.ExcPending
																											if v200 != 0 {
																												return int32(0)
																											} else {
																												v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+15)))
																												if v201 == int32(1) {
																													v206 = F_errstart(m, int32(13), int32(0))
																													mBase = m.M
																													v207 = m.ExcPending
																													if v207 != 0 {
																														return int32(0)
																													} else {
																														if v206 == int32(0) {
																															v229 = v5
																															m.G0 = v17 + int32(16)
																															return v229
																														} else {
																															F_errmsg_internal(m, int32(_a_F_operator_predicate_proof_1), int32(0))
																															mBase = m.M
																															v213 = m.ExcPending
																															if v213 != 0 {
																																return int32(0)
																															} else {
																																F_errfinish(m, int32(_a_F_operator_predicate_proof_2), int32(2019), int32(_a_F_operator_predicate_proof_3))
																																mBase = m.M
																																v218 = m.ExcPending
																																if v218 != 0 {
																																	return int32(0)
																																} else {
																																	v229 = v5
																																	m.G0 = v17 + int32(16)
																																	return v229
																																}
																															}
																														}
																													}
																												} else {
																													v229 = base.B2i32(v195 != int64(0))
																													m.G0 = v17 + int32(16)
																													return v229
																												}
																											}
																										}
																									}
																								} else {
																									v188 = v183
																									v190 = *(*int32)(unsafe.Add(mBase, uint32(v188)+20))
																									*(*int32)(unsafe.Add(mBase, _c_F_operator_predicate_proof[0])) = v190
																									v194 = *(*int32)(unsafe.Add(mBase, uint32(v181)+24))
																									v195 = m.T0[v194].(func(*base.Module, int32, int32, int32) int64)(m, v181, v188, v17+int32(15))
																									mBase = m.M
																									v196 = m.ExcPending
																									if v196 != 0 {
																										return int32(0)
																									} else {
																										*(*int32)(unsafe.Add(mBase, _c_F_operator_predicate_proof[0])) = v172
																										F_FreeExecutorState(m, v169)
																										mBase = m.M
																										v200 = m.ExcPending
																										if v200 != 0 {
																											return int32(0)
																										} else {
																											v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+15)))
																											if v201 == int32(1) {
																												v206 = F_errstart(m, int32(13), int32(0))
																												mBase = m.M
																												v207 = m.ExcPending
																												if v207 != 0 {
																													return int32(0)
																												} else {
																													if v206 == int32(0) {
																														v229 = v5
																														m.G0 = v17 + int32(16)
																														return v229
																													} else {
																														F_errmsg_internal(m, int32(_a_F_operator_predicate_proof_1), int32(0))
																														mBase = m.M
																														v213 = m.ExcPending
																														if v213 != 0 {
																															return int32(0)
																														} else {
																															F_errfinish(m, int32(_a_F_operator_predicate_proof_2), int32(2019), int32(_a_F_operator_predicate_proof_3))
																															mBase = m.M
																															v218 = m.ExcPending
																															if v218 != 0 {
																																return int32(0)
																															} else {
																																v229 = v5
																																m.G0 = v17 + int32(16)
																																return v229
																															}
																														}
																													}
																												}
																											} else {
																												v229 = base.B2i32(v195 != int64(0))
																												m.G0 = v17 + int32(16)
																												return v229
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
																v229 = v5
																m.G0 = v17 + int32(16)
																return v229
															}
														}
													}
												}
											} else {
												if v57 != 0 {
													v72 = int32(0)
													if base.B2i32(v51 == v72)|base.B2i32(v52 == v72) != 0 {
														v229 = v5
														m.G0 = v17 + int32(16)
														return v229
													} else {
														v77 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
														if v77 != int32(7) {
															v229 = v5
															m.G0 = v17 + int32(16)
															return v229
														} else {
															v80 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
															if v80 != int32(7) {
																v229 = v5
																m.G0 = v17 + int32(16)
																return v229
															} else {
																v83 = F_get_commutator(m, v46)
																mBase = m.M
																v84 = m.ExcPending
																if v84 != 0 {
																	return int32(0)
																} else {
																	if v83 == int32(0) {
																		v229 = v5
																		m.G0 = v17 + int32(16)
																		return v229
																	} else {
																		v87 = F_get_commutator(m, v45)
																		mBase = m.M
																		v88 = m.ExcPending
																		if v88 != 0 {
																			return int32(0)
																		} else {
																			if v87 != 0 {
																				v130 = v51
																				v131 = v52
																				v132 = v83
																				v133 = v87
																				v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131)+32)))
																				if v135 == int32(1) {
																					v138 = F_op_strict(m, v133)
																					mBase = m.M
																					v139 = m.ExcPending
																					if v139 != 0 {
																						return int32(0)
																					} else {
																						if v138 == int32(0) {
																							v229 = v5
																							m.G0 = v17 + int32(16)
																							return v229
																						} else {
																							v142 = int32(1)
																							if l2|base.B2i32(l3 == int32(0)) != 0 {
																								v229 = v142
																								m.G0 = v17 + int32(16)
																								return v229
																							} else {
																								v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130)+32)))
																								if v146 == int32(1) {
																									v149 = F_op_strict(m, v132)
																									mBase = m.M
																									v150 = m.ExcPending
																									if v150 != 0 {
																										return int32(0)
																									} else {
																										if v149 != 0 {
																											v229 = v142
																										} else {
																											v229 = int32(0)
																										}
																										m.G0 = v17 + int32(16)
																										return v229
																									}
																								} else {
																									v229 = int32(0)
																									m.G0 = v17 + int32(16)
																									return v229
																								}
																							}
																						}
																					}
																				} else {
																					v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130)+32)))
																					if v152 == int32(1) {
																						if l3 != 0 {
																							v156 = F_op_strict(m, v132)
																							mBase = m.M
																							v157 = m.ExcPending
																							if v157 != 0 {
																								return int32(0)
																							} else {
																								if v156 != 0 {
																									v229 = int32(1)
																								} else {
																									v229 = int32(0)
																								}
																								m.G0 = v17 + int32(16)
																								return v229
																							}
																						} else {
																							v229 = int32(0)
																							m.G0 = v17 + int32(16)
																							return v229
																						}
																					} else {
																						v160 = F_lookup_proof_cache(m, v132, v133, l2)
																						mBase = m.M
																						v161 = m.ExcPending
																						if v161 != 0 {
																							return int32(0)
																						} else {
																							if l2 != 0 {
																								v164 = int32(16)
																							} else {
																								v164 = int32(12)
																							}
																							v166 = *(*int32)(unsafe.Add(mBase, uint32(v160+v164)))
																							if v166 == int32(0) {
																								v229 = v5
																								m.G0 = v17 + int32(16)
																								return v229
																							} else {
																								v169 = F_CreateExecutorState(m)
																								mBase = m.M
																								v170 = m.ExcPending
																								if v170 != 0 {
																									return int32(0)
																								} else {
																									v171 = int32(_a_F_operator_predicate_proof_0)
																									v172 = *(*int32)(unsafe.Add(mBase, _c_F_operator_predicate_proof[0]))
																									v174 = *(*int32)(unsafe.Add(mBase, uint32(v169)+100))
																									*(*int32)(unsafe.Add(mBase, _c_F_operator_predicate_proof[0])) = v174
																									v176 = F_make_opclause(m, v166, v130, v131, v42)
																									mBase = m.M
																									v177 = m.ExcPending
																									if v177 != 0 {
																										return int32(0)
																									} else {
																										F_fix_opfuncids(m, v176)
																										mBase = m.M
																										v179 = m.ExcPending
																										if v179 != 0 {
																											return int32(0)
																										} else {
																											v181 = F_ExecInitExpr(m, v176, int32(0))
																											mBase = m.M
																											v182 = m.ExcPending
																											if v182 != 0 {
																												return int32(0)
																											} else {
																												v183 = *(*int32)(unsafe.Add(mBase, uint32(v169)+152))
																												if v183 == int32(0) {
																													v186 = F_MakePerTupleExprContext(m, v169)
																													mBase = m.M
																													v187 = m.ExcPending
																													if v187 != 0 {
																														return int32(0)
																													} else {
																														v188 = v186
																														v190 = *(*int32)(unsafe.Add(mBase, uint32(v188)+20))
																														*(*int32)(unsafe.Add(mBase, _c_F_operator_predicate_proof[0])) = v190
																														v194 = *(*int32)(unsafe.Add(mBase, uint32(v181)+24))
																														v195 = m.T0[v194].(func(*base.Module, int32, int32, int32) int64)(m, v181, v188, v17+int32(15))
																														mBase = m.M
																														v196 = m.ExcPending
																														if v196 != 0 {
																															return int32(0)
																														} else {
																															*(*int32)(unsafe.Add(mBase, _c_F_operator_predicate_proof[0])) = v172
																															F_FreeExecutorState(m, v169)
																															mBase = m.M
																															v200 = m.ExcPending
																															if v200 != 0 {
																																return int32(0)
																															} else {
																																v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+15)))
																																if v201 == int32(1) {
																																	v206 = F_errstart(m, int32(13), int32(0))
																																	mBase = m.M
																																	v207 = m.ExcPending
																																	if v207 != 0 {
																																		return int32(0)
																																	} else {
																																		if v206 == int32(0) {
																																			v229 = v5
																																			m.G0 = v17 + int32(16)
																																			return v229
																																		} else {
																																			F_errmsg_internal(m, int32(_a_F_operator_predicate_proof_1), int32(0))
																																			mBase = m.M
																																			v213 = m.ExcPending
																																			if v213 != 0 {
																																				return int32(0)
																																			} else {
																																				F_errfinish(m, int32(_a_F_operator_predicate_proof_2), int32(2019), int32(_a_F_operator_predicate_proof_3))
																																				mBase = m.M
																																				v218 = m.ExcPending
																																				if v218 != 0 {
																																					return int32(0)
																																				} else {
																																					v229 = v5
																																					m.G0 = v17 + int32(16)
																																					return v229
																																				}
																																			}
																																		}
																																	}
																																} else {
																																	v229 = base.B2i32(v195 != int64(0))
																																	m.G0 = v17 + int32(16)
																																	return v229
																																}
																															}
																														}
																													}
																												} else {
																													v188 = v183
																													v190 = *(*int32)(unsafe.Add(mBase, uint32(v188)+20))
																													*(*int32)(unsafe.Add(mBase, _c_F_operator_predicate_proof[0])) = v190
																													v194 = *(*int32)(unsafe.Add(mBase, uint32(v181)+24))
																													v195 = m.T0[v194].(func(*base.Module, int32, int32, int32) int64)(m, v181, v188, v17+int32(15))
																													mBase = m.M
																													v196 = m.ExcPending
																													if v196 != 0 {
																														return int32(0)
																													} else {
																														*(*int32)(unsafe.Add(mBase, _c_F_operator_predicate_proof[0])) = v172
																														F_FreeExecutorState(m, v169)
																														mBase = m.M
																														v200 = m.ExcPending
																														if v200 != 0 {
																															return int32(0)
																														} else {
																															v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+15)))
																															if v201 == int32(1) {
																																v206 = F_errstart(m, int32(13), int32(0))
																																mBase = m.M
																																v207 = m.ExcPending
																																if v207 != 0 {
																																	return int32(0)
																																} else {
																																	if v206 == int32(0) {
																																		v229 = v5
																																		m.G0 = v17 + int32(16)
																																		return v229
																																	} else {
																																		F_errmsg_internal(m, int32(_a_F_operator_predicate_proof_1), int32(0))
																																		mBase = m.M
																																		v213 = m.ExcPending
																																		if v213 != 0 {
																																			return int32(0)
																																		} else {
																																			F_errfinish(m, int32(_a_F_operator_predicate_proof_2), int32(2019), int32(_a_F_operator_predicate_proof_3))
																																			mBase = m.M
																																			v218 = m.ExcPending
																																			if v218 != 0 {
																																				return int32(0)
																																			} else {
																																				v229 = v5
																																				m.G0 = v17 + int32(16)
																																				return v229
																																			}
																																		}
																																	}
																																}
																															} else {
																																v229 = base.B2i32(v195 != int64(0))
																																m.G0 = v17 + int32(16)
																																return v229
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
																				v229 = v5
																				m.G0 = v17 + int32(16)
																				return v229
																			}
																		}
																	}
																}
															}
														}
													}
												} else {
													v89 = F_equal(m, v51, v48)
													mBase = m.M
													v90 = m.ExcPending
													if v90 != 0 {
														return int32(0)
													} else {
														v91 = F_equal(m, v50, v52)
														mBase = m.M
														v92 = m.ExcPending
														if v92 != 0 {
															return int32(0)
														} else {
															if v89 != 0 {
																if v91 != 0 {
																	v93 = F_get_commutator(m, v46)
																	mBase = m.M
																	v94 = m.ExcPending
																	if v94 != 0 {
																		return int32(0)
																	} else {
																		if v93 == int32(0) {
																			v229 = v5
																			m.G0 = v17 + int32(16)
																			return v229
																		} else {
																			v97 = F_operator_same_subexprs_proof(m, v93, v45, l2)
																			mBase = m.M
																			v98 = m.ExcPending
																			if v98 != 0 {
																				return int32(0)
																			} else {
																				v229 = v97
																				m.G0 = v17 + int32(16)
																				return v229
																			}
																		}
																	}
																} else {
																	v99 = int32(0)
																	if base.B2i32(v50 == v99)|base.B2i32(v52 == v99) != 0 {
																		v229 = v5
																		m.G0 = v17 + int32(16)
																		return v229
																	} else {
																		v104 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
																		if v104 != int32(7) {
																			v229 = v5
																			m.G0 = v17 + int32(16)
																			return v229
																		} else {
																			v107 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
																			if v107 != int32(7) {
																				v229 = v5
																				m.G0 = v17 + int32(16)
																				return v229
																			} else {
																				v110 = F_get_commutator(m, v45)
																				mBase = m.M
																				v111 = m.ExcPending
																				if v111 != 0 {
																					return int32(0)
																				} else {
																					if v110 != 0 {
																						v130 = v50
																						v131 = v52
																						v132 = v46
																						v133 = v110
																						v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131)+32)))
																						if v135 == int32(1) {
																							v138 = F_op_strict(m, v133)
																							mBase = m.M
																							v139 = m.ExcPending
																							if v139 != 0 {
																								return int32(0)
																							} else {
																								if v138 == int32(0) {
																									v229 = v5
																									m.G0 = v17 + int32(16)
																									return v229
																								} else {
																									v142 = int32(1)
																									if l2|base.B2i32(l3 == int32(0)) != 0 {
																										v229 = v142
																										m.G0 = v17 + int32(16)
																										return v229
																									} else {
																										v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130)+32)))
																										if v146 == int32(1) {
																											v149 = F_op_strict(m, v132)
																											mBase = m.M
																											v150 = m.ExcPending
																											if v150 != 0 {
																												return int32(0)
																											} else {
																												if v149 != 0 {
																													v229 = v142
																												} else {
																													v229 = int32(0)
																												}
																												m.G0 = v17 + int32(16)
																												return v229
																											}
																										} else {
																											v229 = int32(0)
																											m.G0 = v17 + int32(16)
																											return v229
																										}
																									}
																								}
																							}
																						} else {
																							v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130)+32)))
																							if v152 == int32(1) {
																								if l3 != 0 {
																									v156 = F_op_strict(m, v132)
																									mBase = m.M
																									v157 = m.ExcPending
																									if v157 != 0 {
																										return int32(0)
																									} else {
																										if v156 != 0 {
																											v229 = int32(1)
																										} else {
																											v229 = int32(0)
																										}
																										m.G0 = v17 + int32(16)
																										return v229
																									}
																								} else {
																									v229 = int32(0)
																									m.G0 = v17 + int32(16)
																									return v229
																								}
																							} else {
																								v160 = F_lookup_proof_cache(m, v132, v133, l2)
																								mBase = m.M
																								v161 = m.ExcPending
																								if v161 != 0 {
																									return int32(0)
																								} else {
																									if l2 != 0 {
																										v164 = int32(16)
																									} else {
																										v164 = int32(12)
																									}
																									v166 = *(*int32)(unsafe.Add(mBase, uint32(v160+v164)))
																									if v166 == int32(0) {
																										v229 = v5
																										m.G0 = v17 + int32(16)
																										return v229
																									} else {
																										v169 = F_CreateExecutorState(m)
																										mBase = m.M
																										v170 = m.ExcPending
																										if v170 != 0 {
																											return int32(0)
																										} else {
																											v171 = int32(_a_F_operator_predicate_proof_0)
																											v172 = *(*int32)(unsafe.Add(mBase, _c_F_operator_predicate_proof[0]))
																											v174 = *(*int32)(unsafe.Add(mBase, uint32(v169)+100))
																											*(*int32)(unsafe.Add(mBase, _c_F_operator_predicate_proof[0])) = v174
																											v176 = F_make_opclause(m, v166, v130, v131, v42)
																											mBase = m.M
																											v177 = m.ExcPending
																											if v177 != 0 {
																												return int32(0)
																											} else {
																												F_fix_opfuncids(m, v176)
																												mBase = m.M
																												v179 = m.ExcPending
																												if v179 != 0 {
																													return int32(0)
																												} else {
																													v181 = F_ExecInitExpr(m, v176, int32(0))
																													mBase = m.M
																													v182 = m.ExcPending
																													if v182 != 0 {
																														return int32(0)
																													} else {
																														v183 = *(*int32)(unsafe.Add(mBase, uint32(v169)+152))
																														if v183 == int32(0) {
																															v186 = F_MakePerTupleExprContext(m, v169)
																															mBase = m.M
																															v187 = m.ExcPending
																															if v187 != 0 {
																																return int32(0)
																															} else {
																																v188 = v186
																																v190 = *(*int32)(unsafe.Add(mBase, uint32(v188)+20))
																																*(*int32)(unsafe.Add(mBase, _c_F_operator_predicate_proof[0])) = v190
																																v194 = *(*int32)(unsafe.Add(mBase, uint32(v181)+24))
																																v195 = m.T0[v194].(func(*base.Module, int32, int32, int32) int64)(m, v181, v188, v17+int32(15))
																																mBase = m.M
																																v196 = m.ExcPending
																																if v196 != 0 {
																																	return int32(0)
																																} else {
																																	*(*int32)(unsafe.Add(mBase, _c_F_operator_predicate_proof[0])) = v172
																																	F_FreeExecutorState(m, v169)
																																	mBase = m.M
																																	v200 = m.ExcPending
																																	if v200 != 0 {
																																		return int32(0)
																																	} else {
																																		v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+15)))
																																		if v201 == int32(1) {
																																			v206 = F_errstart(m, int32(13), int32(0))
																																			mBase = m.M
																																			v207 = m.ExcPending
																																			if v207 != 0 {
																																				return int32(0)
																																			} else {
																																				if v206 == int32(0) {
																																					v229 = v5
																																					m.G0 = v17 + int32(16)
																																					return v229
																																				} else {
																																					F_errmsg_internal(m, int32(_a_F_operator_predicate_proof_1), int32(0))
																																					mBase = m.M
																																					v213 = m.ExcPending
																																					if v213 != 0 {
																																						return int32(0)
																																					} else {
																																						F_errfinish(m, int32(_a_F_operator_predicate_proof_2), int32(2019), int32(_a_F_operator_predicate_proof_3))
																																						mBase = m.M
																																						v218 = m.ExcPending
																																						if v218 != 0 {
																																							return int32(0)
																																						} else {
																																							v229 = v5
																																							m.G0 = v17 + int32(16)
																																							return v229
																																						}
																																					}
																																				}
																																			}
																																		} else {
																																			v229 = base.B2i32(v195 != int64(0))
																																			m.G0 = v17 + int32(16)
																																			return v229
																																		}
																																	}
																																}
																															}
																														} else {
																															v188 = v183
																															v190 = *(*int32)(unsafe.Add(mBase, uint32(v188)+20))
																															*(*int32)(unsafe.Add(mBase, _c_F_operator_predicate_proof[0])) = v190
																															v194 = *(*int32)(unsafe.Add(mBase, uint32(v181)+24))
																															v195 = m.T0[v194].(func(*base.Module, int32, int32, int32) int64)(m, v181, v188, v17+int32(15))
																															mBase = m.M
																															v196 = m.ExcPending
																															if v196 != 0 {
																																return int32(0)
																															} else {
																																*(*int32)(unsafe.Add(mBase, _c_F_operator_predicate_proof[0])) = v172
																																F_FreeExecutorState(m, v169)
																																mBase = m.M
																																v200 = m.ExcPending
																																if v200 != 0 {
																																	return int32(0)
																																} else {
																																	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+15)))
																																	if v201 == int32(1) {
																																		v206 = F_errstart(m, int32(13), int32(0))
																																		mBase = m.M
																																		v207 = m.ExcPending
																																		if v207 != 0 {
																																			return int32(0)
																																		} else {
																																			if v206 == int32(0) {
																																				v229 = v5
																																				m.G0 = v17 + int32(16)
																																				return v229
																																			} else {
																																				F_errmsg_internal(m, int32(_a_F_operator_predicate_proof_1), int32(0))
																																				mBase = m.M
																																				v213 = m.ExcPending
																																				if v213 != 0 {
																																					return int32(0)
																																				} else {
																																					F_errfinish(m, int32(_a_F_operator_predicate_proof_2), int32(2019), int32(_a_F_operator_predicate_proof_3))
																																					mBase = m.M
																																					v218 = m.ExcPending
																																					if v218 != 0 {
																																						return int32(0)
																																					} else {
																																						v229 = v5
																																						m.G0 = v17 + int32(16)
																																						return v229
																																					}
																																				}
																																			}
																																		}
																																	} else {
																																		v229 = base.B2i32(v195 != int64(0))
																																		m.G0 = v17 + int32(16)
																																		return v229
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
																						v229 = v5
																						m.G0 = v17 + int32(16)
																						return v229
																					}
																				}
																			}
																		}
																	}
																}
															} else {
																if base.B2i32(v51 == int32(0))|(v91^int32(1)) != 0 {
																	v229 = v5
																	m.G0 = v17 + int32(16)
																	return v229
																} else {
																	v119 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
																	if base.B2i32(v48 == int32(0))|base.B2i32(v119 != int32(7)) != 0 {
																		v229 = v5
																		m.G0 = v17 + int32(16)
																		return v229
																	} else {
																		v123 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
																		if v123 != int32(7) {
																			v229 = v5
																			m.G0 = v17 + int32(16)
																			return v229
																		} else {
																			v126 = F_get_commutator(m, v46)
																			mBase = m.M
																			v127 = m.ExcPending
																			if v127 != 0 {
																				return int32(0)
																			} else {
																				if v126 == int32(0) {
																					v229 = v5
																					m.G0 = v17 + int32(16)
																					return v229
																				} else {
																					v130 = v51
																					v131 = v48
																					v132 = v126
																					v133 = v45
																					v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131)+32)))
																					if v135 == int32(1) {
																						v138 = F_op_strict(m, v133)
																						mBase = m.M
																						v139 = m.ExcPending
																						if v139 != 0 {
																							return int32(0)
																						} else {
																							if v138 == int32(0) {
																								v229 = v5
																								m.G0 = v17 + int32(16)
																								return v229
																							} else {
																								v142 = int32(1)
																								if l2|base.B2i32(l3 == int32(0)) != 0 {
																									v229 = v142
																									m.G0 = v17 + int32(16)
																									return v229
																								} else {
																									v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130)+32)))
																									if v146 == int32(1) {
																										v149 = F_op_strict(m, v132)
																										mBase = m.M
																										v150 = m.ExcPending
																										if v150 != 0 {
																											return int32(0)
																										} else {
																											if v149 != 0 {
																												v229 = v142
																											} else {
																												v229 = int32(0)
																											}
																											m.G0 = v17 + int32(16)
																											return v229
																										}
																									} else {
																										v229 = int32(0)
																										m.G0 = v17 + int32(16)
																										return v229
																									}
																								}
																							}
																						}
																					} else {
																						v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130)+32)))
																						if v152 == int32(1) {
																							if l3 != 0 {
																								v156 = F_op_strict(m, v132)
																								mBase = m.M
																								v157 = m.ExcPending
																								if v157 != 0 {
																									return int32(0)
																								} else {
																									if v156 != 0 {
																										v229 = int32(1)
																									} else {
																										v229 = int32(0)
																									}
																									m.G0 = v17 + int32(16)
																									return v229
																								}
																							} else {
																								v229 = int32(0)
																								m.G0 = v17 + int32(16)
																								return v229
																							}
																						} else {
																							v160 = F_lookup_proof_cache(m, v132, v133, l2)
																							mBase = m.M
																							v161 = m.ExcPending
																							if v161 != 0 {
																								return int32(0)
																							} else {
																								if l2 != 0 {
																									v164 = int32(16)
																								} else {
																									v164 = int32(12)
																								}
																								v166 = *(*int32)(unsafe.Add(mBase, uint32(v160+v164)))
																								if v166 == int32(0) {
																									v229 = v5
																									m.G0 = v17 + int32(16)
																									return v229
																								} else {
																									v169 = F_CreateExecutorState(m)
																									mBase = m.M
																									v170 = m.ExcPending
																									if v170 != 0 {
																										return int32(0)
																									} else {
																										v171 = int32(_a_F_operator_predicate_proof_0)
																										v172 = *(*int32)(unsafe.Add(mBase, _c_F_operator_predicate_proof[0]))
																										v174 = *(*int32)(unsafe.Add(mBase, uint32(v169)+100))
																										*(*int32)(unsafe.Add(mBase, _c_F_operator_predicate_proof[0])) = v174
																										v176 = F_make_opclause(m, v166, v130, v131, v42)
																										mBase = m.M
																										v177 = m.ExcPending
																										if v177 != 0 {
																											return int32(0)
																										} else {
																											F_fix_opfuncids(m, v176)
																											mBase = m.M
																											v179 = m.ExcPending
																											if v179 != 0 {
																												return int32(0)
																											} else {
																												v181 = F_ExecInitExpr(m, v176, int32(0))
																												mBase = m.M
																												v182 = m.ExcPending
																												if v182 != 0 {
																													return int32(0)
																												} else {
																													v183 = *(*int32)(unsafe.Add(mBase, uint32(v169)+152))
																													if v183 == int32(0) {
																														v186 = F_MakePerTupleExprContext(m, v169)
																														mBase = m.M
																														v187 = m.ExcPending
																														if v187 != 0 {
																															return int32(0)
																														} else {
																															v188 = v186
																															v190 = *(*int32)(unsafe.Add(mBase, uint32(v188)+20))
																															*(*int32)(unsafe.Add(mBase, _c_F_operator_predicate_proof[0])) = v190
																															v194 = *(*int32)(unsafe.Add(mBase, uint32(v181)+24))
																															v195 = m.T0[v194].(func(*base.Module, int32, int32, int32) int64)(m, v181, v188, v17+int32(15))
																															mBase = m.M
																															v196 = m.ExcPending
																															if v196 != 0 {
																																return int32(0)
																															} else {
																																*(*int32)(unsafe.Add(mBase, _c_F_operator_predicate_proof[0])) = v172
																																F_FreeExecutorState(m, v169)
																																mBase = m.M
																																v200 = m.ExcPending
																																if v200 != 0 {
																																	return int32(0)
																																} else {
																																	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+15)))
																																	if v201 == int32(1) {
																																		v206 = F_errstart(m, int32(13), int32(0))
																																		mBase = m.M
																																		v207 = m.ExcPending
																																		if v207 != 0 {
																																			return int32(0)
																																		} else {
																																			if v206 == int32(0) {
																																				v229 = v5
																																				m.G0 = v17 + int32(16)
																																				return v229
																																			} else {
																																				F_errmsg_internal(m, int32(_a_F_operator_predicate_proof_1), int32(0))
																																				mBase = m.M
																																				v213 = m.ExcPending
																																				if v213 != 0 {
																																					return int32(0)
																																				} else {
																																					F_errfinish(m, int32(_a_F_operator_predicate_proof_2), int32(2019), int32(_a_F_operator_predicate_proof_3))
																																					mBase = m.M
																																					v218 = m.ExcPending
																																					if v218 != 0 {
																																						return int32(0)
																																					} else {
																																						v229 = v5
																																						m.G0 = v17 + int32(16)
																																						return v229
																																					}
																																				}
																																			}
																																		}
																																	} else {
																																		v229 = base.B2i32(v195 != int64(0))
																																		m.G0 = v17 + int32(16)
																																		return v229
																																	}
																																}
																															}
																														}
																													} else {
																														v188 = v183
																														v190 = *(*int32)(unsafe.Add(mBase, uint32(v188)+20))
																														*(*int32)(unsafe.Add(mBase, _c_F_operator_predicate_proof[0])) = v190
																														v194 = *(*int32)(unsafe.Add(mBase, uint32(v181)+24))
																														v195 = m.T0[v194].(func(*base.Module, int32, int32, int32) int64)(m, v181, v188, v17+int32(15))
																														mBase = m.M
																														v196 = m.ExcPending
																														if v196 != 0 {
																															return int32(0)
																														} else {
																															*(*int32)(unsafe.Add(mBase, _c_F_operator_predicate_proof[0])) = v172
																															F_FreeExecutorState(m, v169)
																															mBase = m.M
																															v200 = m.ExcPending
																															if v200 != 0 {
																																return int32(0)
																															} else {
																																v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+15)))
																																if v201 == int32(1) {
																																	v206 = F_errstart(m, int32(13), int32(0))
																																	mBase = m.M
																																	v207 = m.ExcPending
																																	if v207 != 0 {
																																		return int32(0)
																																	} else {
																																		if v206 == int32(0) {
																																			v229 = v5
																																			m.G0 = v17 + int32(16)
																																			return v229
																																		} else {
																																			F_errmsg_internal(m, int32(_a_F_operator_predicate_proof_1), int32(0))
																																			mBase = m.M
																																			v213 = m.ExcPending
																																			if v213 != 0 {
																																				return int32(0)
																																			} else {
																																				F_errfinish(m, int32(_a_F_operator_predicate_proof_2), int32(2019), int32(_a_F_operator_predicate_proof_3))
																																				mBase = m.M
																																				v218 = m.ExcPending
																																				if v218 != 0 {
																																					return int32(0)
																																				} else {
																																					v229 = v5
																																					m.G0 = v17 + int32(16)
																																					return v229
																																				}
																																			}
																																		}
																																	}
																																} else {
																																	v229 = base.B2i32(v195 != int64(0))
																																	m.G0 = v17 + int32(16)
																																	return v229
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
