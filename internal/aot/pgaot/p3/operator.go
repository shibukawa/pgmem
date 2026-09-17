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
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
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
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
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
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
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
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v228 int32
	_ = v228
	v5 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	if l0 == v5 {
		v228 = v5
		m.G0 = v16 + int32(16)
		return v228
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v20 != int32(17) {
			v228 = v5
			m.G0 = v16 + int32(16)
			return v228
		} else {
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			v24 = int32(0)
			if base.B2i32(v23 == v24)|base.B2i32(l1 == v24) != 0 {
				v228 = v5
				m.G0 = v16 + int32(16)
				return v228
			} else {
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
				if v29 != int32(2) {
					v228 = v5
					m.G0 = v16 + int32(16)
					return v228
				} else {
					v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					if v32 != int32(17) {
						v228 = v5
						m.G0 = v16 + int32(16)
						return v228
					} else {
						v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
						if v35 == int32(0) {
							v228 = v5
							m.G0 = v16 + int32(16)
							return v228
						} else {
							v38 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
							if v38 != int32(2) {
								v228 = v5
								m.G0 = v16 + int32(16)
								return v228
							} else {
								v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
								v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
								if v41 != v42 {
									v228 = v5
									m.G0 = v16 + int32(16)
									return v228
								} else {
									v44 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
									v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									v46 = *(*int32)(unsafe.Add(mBase, uint32(v35)+12))
									v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
									v48 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
									v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
									v50 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
									v51 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
									v52 = F_equal(m, v50, v51)
									mBase = m.M
									v55 = m.ExcPending
									if v55 != 0 {
										return int32(0)
									} else {
										v56 = F_equal(m, v49, v47)
										mBase = m.M
										v57 = m.ExcPending
										if v57 != 0 {
											return int32(0)
										} else {
											if v52 != 0 {
												if v56 != 0 {
													v58 = F_operator_same_subexprs_proof(m, v45, v44, l2)
													mBase = m.M
													v59 = m.ExcPending
													if v59 != 0 {
														return int32(0)
													} else {
														v228 = v58
														m.G0 = v16 + int32(16)
														return v228
													}
												} else {
													if v49 == int32(0) {
														v228 = v5
														m.G0 = v16 + int32(16)
														return v228
													} else {
														v64 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
														if base.B2i32(v47 == int32(0))|base.B2i32(v64 != int32(7)) != 0 {
															v228 = v5
															m.G0 = v16 + int32(16)
															return v228
														} else {
															v68 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
															if v68 == int32(7) {
																v129 = v49
																v130 = v47
																v131 = v45
																v132 = v44
																v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130)+24)))
																if v134 == int32(1) {
																	v137 = F_op_strict(m, v132)
																	mBase = m.M
																	v138 = m.ExcPending
																	if v138 != 0 {
																		return int32(0)
																	} else {
																		if v137 == int32(0) {
																			v228 = v5
																			m.G0 = v16 + int32(16)
																			return v228
																		} else {
																			v141 = int32(1)
																			if l2|base.B2i32(l3 == int32(0)) != 0 {
																				v228 = v141
																				m.G0 = v16 + int32(16)
																				return v228
																			} else {
																				v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+24)))
																				if v145 == int32(1) {
																					v148 = F_op_strict(m, v131)
																					mBase = m.M
																					v149 = m.ExcPending
																					if v149 != 0 {
																						return int32(0)
																					} else {
																						if v148 != 0 {
																							v228 = v141
																						} else {
																							v228 = int32(0)
																						}
																						m.G0 = v16 + int32(16)
																						return v228
																					}
																				} else {
																					v228 = int32(0)
																					m.G0 = v16 + int32(16)
																					return v228
																				}
																			}
																		}
																	}
																} else {
																	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+24)))
																	if v151 == int32(1) {
																		if l3 != 0 {
																			v155 = F_op_strict(m, v131)
																			mBase = m.M
																			v156 = m.ExcPending
																			if v156 != 0 {
																				return int32(0)
																			} else {
																				if v155 != 0 {
																					v228 = int32(1)
																				} else {
																					v228 = int32(0)
																				}
																				m.G0 = v16 + int32(16)
																				return v228
																			}
																		} else {
																			v228 = int32(0)
																			m.G0 = v16 + int32(16)
																			return v228
																		}
																	} else {
																		v159 = F_lookup_proof_cache(m, v131, v132, l2)
																		mBase = m.M
																		v160 = m.ExcPending
																		if v160 != 0 {
																			return int32(0)
																		} else {
																			if l2 != 0 {
																				v163 = int32(16)
																			} else {
																				v163 = int32(12)
																			}
																			v165 = *(*int32)(unsafe.Add(mBase, uint32(v159+v163)))
																			if v165 == int32(0) {
																				v228 = v5
																				m.G0 = v16 + int32(16)
																				return v228
																			} else {
																				v168 = F_CreateExecutorState(m)
																				mBase = m.M
																				v169 = m.ExcPending
																				if v169 != 0 {
																					return int32(0)
																				} else {
																					v170 = int32(_a_F_operator_predicate_proof_0)
																					v171 = *(*int32)(unsafe.Add(mBase, _c_F_operator_predicate_proof[0]))
																					v173 = *(*int32)(unsafe.Add(mBase, uint32(v168)+100))
																					*(*int32)(unsafe.Add(mBase, _c_F_operator_predicate_proof[0])) = v173
																					v175 = F_make_opclause(m, v165, v129, v130, v41)
																					mBase = m.M
																					v176 = m.ExcPending
																					if v176 != 0 {
																						return int32(0)
																					} else {
																						F_fix_opfuncids(m, v175)
																						mBase = m.M
																						v178 = m.ExcPending
																						if v178 != 0 {
																							return int32(0)
																						} else {
																							v180 = F_ExecInitExpr(m, v175, int32(0))
																							mBase = m.M
																							v181 = m.ExcPending
																							if v181 != 0 {
																								return int32(0)
																							} else {
																								v182 = *(*int32)(unsafe.Add(mBase, uint32(v168)+152))
																								if v182 == int32(0) {
																									v185 = F_MakePerTupleExprContext(m, v168)
																									mBase = m.M
																									v186 = m.ExcPending
																									if v186 != 0 {
																										return int32(0)
																									} else {
																										v187 = v185
																										v189 = *(*int32)(unsafe.Add(mBase, uint32(v187)+20))
																										*(*int32)(unsafe.Add(mBase, _c_F_operator_predicate_proof[0])) = v189
																										v193 = *(*int32)(unsafe.Add(mBase, uint32(v180)+20))
																										v194 = m.T0[v193].(func(*base.Module, int32, int32, int32) int32)(m, v180, v187, v16+int32(15))
																										mBase = m.M
																										v195 = m.ExcPending
																										if v195 != 0 {
																											return int32(0)
																										} else {
																											*(*int32)(unsafe.Add(mBase, _c_F_operator_predicate_proof[0])) = v171
																											F_FreeExecutorState(m, v168)
																											mBase = m.M
																											v199 = m.ExcPending
																											if v199 != 0 {
																												return int32(0)
																											} else {
																												v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+15)))
																												if v200 == int32(1) {
																													v205 = F_errstart(m, int32(13), int32(0))
																													mBase = m.M
																													v206 = m.ExcPending
																													if v206 != 0 {
																														return int32(0)
																													} else {
																														if v205 == int32(0) {
																															v228 = v5
																															m.G0 = v16 + int32(16)
																															return v228
																														} else {
																															F_errmsg_internal(m, int32(_a_F_operator_predicate_proof_1), int32(0))
																															mBase = m.M
																															v212 = m.ExcPending
																															if v212 != 0 {
																																return int32(0)
																															} else {
																																F_errfinish(m, int32(_a_F_operator_predicate_proof_2), int32(2017), int32(_a_F_operator_predicate_proof_3))
																																mBase = m.M
																																v217 = m.ExcPending
																																if v217 != 0 {
																																	return int32(0)
																																} else {
																																	v228 = v5
																																	m.G0 = v16 + int32(16)
																																	return v228
																																}
																															}
																														}
																													}
																												} else {
																													v228 = base.B2i32(v194 != int32(0))
																													m.G0 = v16 + int32(16)
																													return v228
																												}
																											}
																										}
																									}
																								} else {
																									v187 = v182
																									v189 = *(*int32)(unsafe.Add(mBase, uint32(v187)+20))
																									*(*int32)(unsafe.Add(mBase, _c_F_operator_predicate_proof[0])) = v189
																									v193 = *(*int32)(unsafe.Add(mBase, uint32(v180)+20))
																									v194 = m.T0[v193].(func(*base.Module, int32, int32, int32) int32)(m, v180, v187, v16+int32(15))
																									mBase = m.M
																									v195 = m.ExcPending
																									if v195 != 0 {
																										return int32(0)
																									} else {
																										*(*int32)(unsafe.Add(mBase, _c_F_operator_predicate_proof[0])) = v171
																										F_FreeExecutorState(m, v168)
																										mBase = m.M
																										v199 = m.ExcPending
																										if v199 != 0 {
																											return int32(0)
																										} else {
																											v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+15)))
																											if v200 == int32(1) {
																												v205 = F_errstart(m, int32(13), int32(0))
																												mBase = m.M
																												v206 = m.ExcPending
																												if v206 != 0 {
																													return int32(0)
																												} else {
																													if v205 == int32(0) {
																														v228 = v5
																														m.G0 = v16 + int32(16)
																														return v228
																													} else {
																														F_errmsg_internal(m, int32(_a_F_operator_predicate_proof_1), int32(0))
																														mBase = m.M
																														v212 = m.ExcPending
																														if v212 != 0 {
																															return int32(0)
																														} else {
																															F_errfinish(m, int32(_a_F_operator_predicate_proof_2), int32(2017), int32(_a_F_operator_predicate_proof_3))
																															mBase = m.M
																															v217 = m.ExcPending
																															if v217 != 0 {
																																return int32(0)
																															} else {
																																v228 = v5
																																m.G0 = v16 + int32(16)
																																return v228
																															}
																														}
																													}
																												}
																											} else {
																												v228 = base.B2i32(v194 != int32(0))
																												m.G0 = v16 + int32(16)
																												return v228
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
																v228 = v5
																m.G0 = v16 + int32(16)
																return v228
															}
														}
													}
												}
											} else {
												if v56 != 0 {
													v71 = int32(0)
													if base.B2i32(v50 == v71)|base.B2i32(v51 == v71) != 0 {
														v228 = v5
														m.G0 = v16 + int32(16)
														return v228
													} else {
														v76 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
														if v76 != int32(7) {
															v228 = v5
															m.G0 = v16 + int32(16)
															return v228
														} else {
															v79 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
															if v79 != int32(7) {
																v228 = v5
																m.G0 = v16 + int32(16)
																return v228
															} else {
																v82 = F_get_commutator(m, v45)
																mBase = m.M
																v83 = m.ExcPending
																if v83 != 0 {
																	return int32(0)
																} else {
																	if v82 == int32(0) {
																		v228 = v5
																		m.G0 = v16 + int32(16)
																		return v228
																	} else {
																		v86 = F_get_commutator(m, v44)
																		mBase = m.M
																		v87 = m.ExcPending
																		if v87 != 0 {
																			return int32(0)
																		} else {
																			if v86 != 0 {
																				v129 = v50
																				v130 = v51
																				v131 = v82
																				v132 = v86
																				v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130)+24)))
																				if v134 == int32(1) {
																					v137 = F_op_strict(m, v132)
																					mBase = m.M
																					v138 = m.ExcPending
																					if v138 != 0 {
																						return int32(0)
																					} else {
																						if v137 == int32(0) {
																							v228 = v5
																							m.G0 = v16 + int32(16)
																							return v228
																						} else {
																							v141 = int32(1)
																							if l2|base.B2i32(l3 == int32(0)) != 0 {
																								v228 = v141
																								m.G0 = v16 + int32(16)
																								return v228
																							} else {
																								v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+24)))
																								if v145 == int32(1) {
																									v148 = F_op_strict(m, v131)
																									mBase = m.M
																									v149 = m.ExcPending
																									if v149 != 0 {
																										return int32(0)
																									} else {
																										if v148 != 0 {
																											v228 = v141
																										} else {
																											v228 = int32(0)
																										}
																										m.G0 = v16 + int32(16)
																										return v228
																									}
																								} else {
																									v228 = int32(0)
																									m.G0 = v16 + int32(16)
																									return v228
																								}
																							}
																						}
																					}
																				} else {
																					v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+24)))
																					if v151 == int32(1) {
																						if l3 != 0 {
																							v155 = F_op_strict(m, v131)
																							mBase = m.M
																							v156 = m.ExcPending
																							if v156 != 0 {
																								return int32(0)
																							} else {
																								if v155 != 0 {
																									v228 = int32(1)
																								} else {
																									v228 = int32(0)
																								}
																								m.G0 = v16 + int32(16)
																								return v228
																							}
																						} else {
																							v228 = int32(0)
																							m.G0 = v16 + int32(16)
																							return v228
																						}
																					} else {
																						v159 = F_lookup_proof_cache(m, v131, v132, l2)
																						mBase = m.M
																						v160 = m.ExcPending
																						if v160 != 0 {
																							return int32(0)
																						} else {
																							if l2 != 0 {
																								v163 = int32(16)
																							} else {
																								v163 = int32(12)
																							}
																							v165 = *(*int32)(unsafe.Add(mBase, uint32(v159+v163)))
																							if v165 == int32(0) {
																								v228 = v5
																								m.G0 = v16 + int32(16)
																								return v228
																							} else {
																								v168 = F_CreateExecutorState(m)
																								mBase = m.M
																								v169 = m.ExcPending
																								if v169 != 0 {
																									return int32(0)
																								} else {
																									v170 = int32(_a_F_operator_predicate_proof_0)
																									v171 = *(*int32)(unsafe.Add(mBase, _c_F_operator_predicate_proof[0]))
																									v173 = *(*int32)(unsafe.Add(mBase, uint32(v168)+100))
																									*(*int32)(unsafe.Add(mBase, _c_F_operator_predicate_proof[0])) = v173
																									v175 = F_make_opclause(m, v165, v129, v130, v41)
																									mBase = m.M
																									v176 = m.ExcPending
																									if v176 != 0 {
																										return int32(0)
																									} else {
																										F_fix_opfuncids(m, v175)
																										mBase = m.M
																										v178 = m.ExcPending
																										if v178 != 0 {
																											return int32(0)
																										} else {
																											v180 = F_ExecInitExpr(m, v175, int32(0))
																											mBase = m.M
																											v181 = m.ExcPending
																											if v181 != 0 {
																												return int32(0)
																											} else {
																												v182 = *(*int32)(unsafe.Add(mBase, uint32(v168)+152))
																												if v182 == int32(0) {
																													v185 = F_MakePerTupleExprContext(m, v168)
																													mBase = m.M
																													v186 = m.ExcPending
																													if v186 != 0 {
																														return int32(0)
																													} else {
																														v187 = v185
																														v189 = *(*int32)(unsafe.Add(mBase, uint32(v187)+20))
																														*(*int32)(unsafe.Add(mBase, _c_F_operator_predicate_proof[0])) = v189
																														v193 = *(*int32)(unsafe.Add(mBase, uint32(v180)+20))
																														v194 = m.T0[v193].(func(*base.Module, int32, int32, int32) int32)(m, v180, v187, v16+int32(15))
																														mBase = m.M
																														v195 = m.ExcPending
																														if v195 != 0 {
																															return int32(0)
																														} else {
																															*(*int32)(unsafe.Add(mBase, _c_F_operator_predicate_proof[0])) = v171
																															F_FreeExecutorState(m, v168)
																															mBase = m.M
																															v199 = m.ExcPending
																															if v199 != 0 {
																																return int32(0)
																															} else {
																																v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+15)))
																																if v200 == int32(1) {
																																	v205 = F_errstart(m, int32(13), int32(0))
																																	mBase = m.M
																																	v206 = m.ExcPending
																																	if v206 != 0 {
																																		return int32(0)
																																	} else {
																																		if v205 == int32(0) {
																																			v228 = v5
																																			m.G0 = v16 + int32(16)
																																			return v228
																																		} else {
																																			F_errmsg_internal(m, int32(_a_F_operator_predicate_proof_1), int32(0))
																																			mBase = m.M
																																			v212 = m.ExcPending
																																			if v212 != 0 {
																																				return int32(0)
																																			} else {
																																				F_errfinish(m, int32(_a_F_operator_predicate_proof_2), int32(2017), int32(_a_F_operator_predicate_proof_3))
																																				mBase = m.M
																																				v217 = m.ExcPending
																																				if v217 != 0 {
																																					return int32(0)
																																				} else {
																																					v228 = v5
																																					m.G0 = v16 + int32(16)
																																					return v228
																																				}
																																			}
																																		}
																																	}
																																} else {
																																	v228 = base.B2i32(v194 != int32(0))
																																	m.G0 = v16 + int32(16)
																																	return v228
																																}
																															}
																														}
																													}
																												} else {
																													v187 = v182
																													v189 = *(*int32)(unsafe.Add(mBase, uint32(v187)+20))
																													*(*int32)(unsafe.Add(mBase, _c_F_operator_predicate_proof[0])) = v189
																													v193 = *(*int32)(unsafe.Add(mBase, uint32(v180)+20))
																													v194 = m.T0[v193].(func(*base.Module, int32, int32, int32) int32)(m, v180, v187, v16+int32(15))
																													mBase = m.M
																													v195 = m.ExcPending
																													if v195 != 0 {
																														return int32(0)
																													} else {
																														*(*int32)(unsafe.Add(mBase, _c_F_operator_predicate_proof[0])) = v171
																														F_FreeExecutorState(m, v168)
																														mBase = m.M
																														v199 = m.ExcPending
																														if v199 != 0 {
																															return int32(0)
																														} else {
																															v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+15)))
																															if v200 == int32(1) {
																																v205 = F_errstart(m, int32(13), int32(0))
																																mBase = m.M
																																v206 = m.ExcPending
																																if v206 != 0 {
																																	return int32(0)
																																} else {
																																	if v205 == int32(0) {
																																		v228 = v5
																																		m.G0 = v16 + int32(16)
																																		return v228
																																	} else {
																																		F_errmsg_internal(m, int32(_a_F_operator_predicate_proof_1), int32(0))
																																		mBase = m.M
																																		v212 = m.ExcPending
																																		if v212 != 0 {
																																			return int32(0)
																																		} else {
																																			F_errfinish(m, int32(_a_F_operator_predicate_proof_2), int32(2017), int32(_a_F_operator_predicate_proof_3))
																																			mBase = m.M
																																			v217 = m.ExcPending
																																			if v217 != 0 {
																																				return int32(0)
																																			} else {
																																				v228 = v5
																																				m.G0 = v16 + int32(16)
																																				return v228
																																			}
																																		}
																																	}
																																}
																															} else {
																																v228 = base.B2i32(v194 != int32(0))
																																m.G0 = v16 + int32(16)
																																return v228
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
																				v228 = v5
																				m.G0 = v16 + int32(16)
																				return v228
																			}
																		}
																	}
																}
															}
														}
													}
												} else {
													v88 = F_equal(m, v50, v47)
													mBase = m.M
													v89 = m.ExcPending
													if v89 != 0 {
														return int32(0)
													} else {
														v90 = F_equal(m, v49, v51)
														mBase = m.M
														v91 = m.ExcPending
														if v91 != 0 {
															return int32(0)
														} else {
															if v88 != 0 {
																if v90 != 0 {
																	v92 = F_get_commutator(m, v45)
																	mBase = m.M
																	v93 = m.ExcPending
																	if v93 != 0 {
																		return int32(0)
																	} else {
																		if v92 == int32(0) {
																			v228 = v5
																			m.G0 = v16 + int32(16)
																			return v228
																		} else {
																			v96 = F_operator_same_subexprs_proof(m, v92, v44, l2)
																			mBase = m.M
																			v97 = m.ExcPending
																			if v97 != 0 {
																				return int32(0)
																			} else {
																				v228 = v96
																				m.G0 = v16 + int32(16)
																				return v228
																			}
																		}
																	}
																} else {
																	v98 = int32(0)
																	if base.B2i32(v49 == v98)|base.B2i32(v51 == v98) != 0 {
																		v228 = v5
																		m.G0 = v16 + int32(16)
																		return v228
																	} else {
																		v103 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
																		if v103 != int32(7) {
																			v228 = v5
																			m.G0 = v16 + int32(16)
																			return v228
																		} else {
																			v106 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
																			if v106 != int32(7) {
																				v228 = v5
																				m.G0 = v16 + int32(16)
																				return v228
																			} else {
																				v109 = F_get_commutator(m, v44)
																				mBase = m.M
																				v110 = m.ExcPending
																				if v110 != 0 {
																					return int32(0)
																				} else {
																					if v109 != 0 {
																						v129 = v49
																						v130 = v51
																						v131 = v45
																						v132 = v109
																						v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130)+24)))
																						if v134 == int32(1) {
																							v137 = F_op_strict(m, v132)
																							mBase = m.M
																							v138 = m.ExcPending
																							if v138 != 0 {
																								return int32(0)
																							} else {
																								if v137 == int32(0) {
																									v228 = v5
																									m.G0 = v16 + int32(16)
																									return v228
																								} else {
																									v141 = int32(1)
																									if l2|base.B2i32(l3 == int32(0)) != 0 {
																										v228 = v141
																										m.G0 = v16 + int32(16)
																										return v228
																									} else {
																										v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+24)))
																										if v145 == int32(1) {
																											v148 = F_op_strict(m, v131)
																											mBase = m.M
																											v149 = m.ExcPending
																											if v149 != 0 {
																												return int32(0)
																											} else {
																												if v148 != 0 {
																													v228 = v141
																												} else {
																													v228 = int32(0)
																												}
																												m.G0 = v16 + int32(16)
																												return v228
																											}
																										} else {
																											v228 = int32(0)
																											m.G0 = v16 + int32(16)
																											return v228
																										}
																									}
																								}
																							}
																						} else {
																							v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+24)))
																							if v151 == int32(1) {
																								if l3 != 0 {
																									v155 = F_op_strict(m, v131)
																									mBase = m.M
																									v156 = m.ExcPending
																									if v156 != 0 {
																										return int32(0)
																									} else {
																										if v155 != 0 {
																											v228 = int32(1)
																										} else {
																											v228 = int32(0)
																										}
																										m.G0 = v16 + int32(16)
																										return v228
																									}
																								} else {
																									v228 = int32(0)
																									m.G0 = v16 + int32(16)
																									return v228
																								}
																							} else {
																								v159 = F_lookup_proof_cache(m, v131, v132, l2)
																								mBase = m.M
																								v160 = m.ExcPending
																								if v160 != 0 {
																									return int32(0)
																								} else {
																									if l2 != 0 {
																										v163 = int32(16)
																									} else {
																										v163 = int32(12)
																									}
																									v165 = *(*int32)(unsafe.Add(mBase, uint32(v159+v163)))
																									if v165 == int32(0) {
																										v228 = v5
																										m.G0 = v16 + int32(16)
																										return v228
																									} else {
																										v168 = F_CreateExecutorState(m)
																										mBase = m.M
																										v169 = m.ExcPending
																										if v169 != 0 {
																											return int32(0)
																										} else {
																											v170 = int32(_a_F_operator_predicate_proof_0)
																											v171 = *(*int32)(unsafe.Add(mBase, _c_F_operator_predicate_proof[0]))
																											v173 = *(*int32)(unsafe.Add(mBase, uint32(v168)+100))
																											*(*int32)(unsafe.Add(mBase, _c_F_operator_predicate_proof[0])) = v173
																											v175 = F_make_opclause(m, v165, v129, v130, v41)
																											mBase = m.M
																											v176 = m.ExcPending
																											if v176 != 0 {
																												return int32(0)
																											} else {
																												F_fix_opfuncids(m, v175)
																												mBase = m.M
																												v178 = m.ExcPending
																												if v178 != 0 {
																													return int32(0)
																												} else {
																													v180 = F_ExecInitExpr(m, v175, int32(0))
																													mBase = m.M
																													v181 = m.ExcPending
																													if v181 != 0 {
																														return int32(0)
																													} else {
																														v182 = *(*int32)(unsafe.Add(mBase, uint32(v168)+152))
																														if v182 == int32(0) {
																															v185 = F_MakePerTupleExprContext(m, v168)
																															mBase = m.M
																															v186 = m.ExcPending
																															if v186 != 0 {
																																return int32(0)
																															} else {
																																v187 = v185
																																v189 = *(*int32)(unsafe.Add(mBase, uint32(v187)+20))
																																*(*int32)(unsafe.Add(mBase, _c_F_operator_predicate_proof[0])) = v189
																																v193 = *(*int32)(unsafe.Add(mBase, uint32(v180)+20))
																																v194 = m.T0[v193].(func(*base.Module, int32, int32, int32) int32)(m, v180, v187, v16+int32(15))
																																mBase = m.M
																																v195 = m.ExcPending
																																if v195 != 0 {
																																	return int32(0)
																																} else {
																																	*(*int32)(unsafe.Add(mBase, _c_F_operator_predicate_proof[0])) = v171
																																	F_FreeExecutorState(m, v168)
																																	mBase = m.M
																																	v199 = m.ExcPending
																																	if v199 != 0 {
																																		return int32(0)
																																	} else {
																																		v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+15)))
																																		if v200 == int32(1) {
																																			v205 = F_errstart(m, int32(13), int32(0))
																																			mBase = m.M
																																			v206 = m.ExcPending
																																			if v206 != 0 {
																																				return int32(0)
																																			} else {
																																				if v205 == int32(0) {
																																					v228 = v5
																																					m.G0 = v16 + int32(16)
																																					return v228
																																				} else {
																																					F_errmsg_internal(m, int32(_a_F_operator_predicate_proof_1), int32(0))
																																					mBase = m.M
																																					v212 = m.ExcPending
																																					if v212 != 0 {
																																						return int32(0)
																																					} else {
																																						F_errfinish(m, int32(_a_F_operator_predicate_proof_2), int32(2017), int32(_a_F_operator_predicate_proof_3))
																																						mBase = m.M
																																						v217 = m.ExcPending
																																						if v217 != 0 {
																																							return int32(0)
																																						} else {
																																							v228 = v5
																																							m.G0 = v16 + int32(16)
																																							return v228
																																						}
																																					}
																																				}
																																			}
																																		} else {
																																			v228 = base.B2i32(v194 != int32(0))
																																			m.G0 = v16 + int32(16)
																																			return v228
																																		}
																																	}
																																}
																															}
																														} else {
																															v187 = v182
																															v189 = *(*int32)(unsafe.Add(mBase, uint32(v187)+20))
																															*(*int32)(unsafe.Add(mBase, _c_F_operator_predicate_proof[0])) = v189
																															v193 = *(*int32)(unsafe.Add(mBase, uint32(v180)+20))
																															v194 = m.T0[v193].(func(*base.Module, int32, int32, int32) int32)(m, v180, v187, v16+int32(15))
																															mBase = m.M
																															v195 = m.ExcPending
																															if v195 != 0 {
																																return int32(0)
																															} else {
																																*(*int32)(unsafe.Add(mBase, _c_F_operator_predicate_proof[0])) = v171
																																F_FreeExecutorState(m, v168)
																																mBase = m.M
																																v199 = m.ExcPending
																																if v199 != 0 {
																																	return int32(0)
																																} else {
																																	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+15)))
																																	if v200 == int32(1) {
																																		v205 = F_errstart(m, int32(13), int32(0))
																																		mBase = m.M
																																		v206 = m.ExcPending
																																		if v206 != 0 {
																																			return int32(0)
																																		} else {
																																			if v205 == int32(0) {
																																				v228 = v5
																																				m.G0 = v16 + int32(16)
																																				return v228
																																			} else {
																																				F_errmsg_internal(m, int32(_a_F_operator_predicate_proof_1), int32(0))
																																				mBase = m.M
																																				v212 = m.ExcPending
																																				if v212 != 0 {
																																					return int32(0)
																																				} else {
																																					F_errfinish(m, int32(_a_F_operator_predicate_proof_2), int32(2017), int32(_a_F_operator_predicate_proof_3))
																																					mBase = m.M
																																					v217 = m.ExcPending
																																					if v217 != 0 {
																																						return int32(0)
																																					} else {
																																						v228 = v5
																																						m.G0 = v16 + int32(16)
																																						return v228
																																					}
																																				}
																																			}
																																		}
																																	} else {
																																		v228 = base.B2i32(v194 != int32(0))
																																		m.G0 = v16 + int32(16)
																																		return v228
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
																						v228 = v5
																						m.G0 = v16 + int32(16)
																						return v228
																					}
																				}
																			}
																		}
																	}
																}
															} else {
																if base.B2i32(v50 == int32(0))|(v90^int32(1)) != 0 {
																	v228 = v5
																	m.G0 = v16 + int32(16)
																	return v228
																} else {
																	v118 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
																	if base.B2i32(v47 == int32(0))|base.B2i32(v118 != int32(7)) != 0 {
																		v228 = v5
																		m.G0 = v16 + int32(16)
																		return v228
																	} else {
																		v122 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
																		if v122 != int32(7) {
																			v228 = v5
																			m.G0 = v16 + int32(16)
																			return v228
																		} else {
																			v125 = F_get_commutator(m, v45)
																			mBase = m.M
																			v126 = m.ExcPending
																			if v126 != 0 {
																				return int32(0)
																			} else {
																				if v125 == int32(0) {
																					v228 = v5
																					m.G0 = v16 + int32(16)
																					return v228
																				} else {
																					v129 = v50
																					v130 = v47
																					v131 = v125
																					v132 = v44
																					v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130)+24)))
																					if v134 == int32(1) {
																						v137 = F_op_strict(m, v132)
																						mBase = m.M
																						v138 = m.ExcPending
																						if v138 != 0 {
																							return int32(0)
																						} else {
																							if v137 == int32(0) {
																								v228 = v5
																								m.G0 = v16 + int32(16)
																								return v228
																							} else {
																								v141 = int32(1)
																								if l2|base.B2i32(l3 == int32(0)) != 0 {
																									v228 = v141
																									m.G0 = v16 + int32(16)
																									return v228
																								} else {
																									v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+24)))
																									if v145 == int32(1) {
																										v148 = F_op_strict(m, v131)
																										mBase = m.M
																										v149 = m.ExcPending
																										if v149 != 0 {
																											return int32(0)
																										} else {
																											if v148 != 0 {
																												v228 = v141
																											} else {
																												v228 = int32(0)
																											}
																											m.G0 = v16 + int32(16)
																											return v228
																										}
																									} else {
																										v228 = int32(0)
																										m.G0 = v16 + int32(16)
																										return v228
																									}
																								}
																							}
																						}
																					} else {
																						v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+24)))
																						if v151 == int32(1) {
																							if l3 != 0 {
																								v155 = F_op_strict(m, v131)
																								mBase = m.M
																								v156 = m.ExcPending
																								if v156 != 0 {
																									return int32(0)
																								} else {
																									if v155 != 0 {
																										v228 = int32(1)
																									} else {
																										v228 = int32(0)
																									}
																									m.G0 = v16 + int32(16)
																									return v228
																								}
																							} else {
																								v228 = int32(0)
																								m.G0 = v16 + int32(16)
																								return v228
																							}
																						} else {
																							v159 = F_lookup_proof_cache(m, v131, v132, l2)
																							mBase = m.M
																							v160 = m.ExcPending
																							if v160 != 0 {
																								return int32(0)
																							} else {
																								if l2 != 0 {
																									v163 = int32(16)
																								} else {
																									v163 = int32(12)
																								}
																								v165 = *(*int32)(unsafe.Add(mBase, uint32(v159+v163)))
																								if v165 == int32(0) {
																									v228 = v5
																									m.G0 = v16 + int32(16)
																									return v228
																								} else {
																									v168 = F_CreateExecutorState(m)
																									mBase = m.M
																									v169 = m.ExcPending
																									if v169 != 0 {
																										return int32(0)
																									} else {
																										v170 = int32(_a_F_operator_predicate_proof_0)
																										v171 = *(*int32)(unsafe.Add(mBase, _c_F_operator_predicate_proof[0]))
																										v173 = *(*int32)(unsafe.Add(mBase, uint32(v168)+100))
																										*(*int32)(unsafe.Add(mBase, _c_F_operator_predicate_proof[0])) = v173
																										v175 = F_make_opclause(m, v165, v129, v130, v41)
																										mBase = m.M
																										v176 = m.ExcPending
																										if v176 != 0 {
																											return int32(0)
																										} else {
																											F_fix_opfuncids(m, v175)
																											mBase = m.M
																											v178 = m.ExcPending
																											if v178 != 0 {
																												return int32(0)
																											} else {
																												v180 = F_ExecInitExpr(m, v175, int32(0))
																												mBase = m.M
																												v181 = m.ExcPending
																												if v181 != 0 {
																													return int32(0)
																												} else {
																													v182 = *(*int32)(unsafe.Add(mBase, uint32(v168)+152))
																													if v182 == int32(0) {
																														v185 = F_MakePerTupleExprContext(m, v168)
																														mBase = m.M
																														v186 = m.ExcPending
																														if v186 != 0 {
																															return int32(0)
																														} else {
																															v187 = v185
																															v189 = *(*int32)(unsafe.Add(mBase, uint32(v187)+20))
																															*(*int32)(unsafe.Add(mBase, _c_F_operator_predicate_proof[0])) = v189
																															v193 = *(*int32)(unsafe.Add(mBase, uint32(v180)+20))
																															v194 = m.T0[v193].(func(*base.Module, int32, int32, int32) int32)(m, v180, v187, v16+int32(15))
																															mBase = m.M
																															v195 = m.ExcPending
																															if v195 != 0 {
																																return int32(0)
																															} else {
																																*(*int32)(unsafe.Add(mBase, _c_F_operator_predicate_proof[0])) = v171
																																F_FreeExecutorState(m, v168)
																																mBase = m.M
																																v199 = m.ExcPending
																																if v199 != 0 {
																																	return int32(0)
																																} else {
																																	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+15)))
																																	if v200 == int32(1) {
																																		v205 = F_errstart(m, int32(13), int32(0))
																																		mBase = m.M
																																		v206 = m.ExcPending
																																		if v206 != 0 {
																																			return int32(0)
																																		} else {
																																			if v205 == int32(0) {
																																				v228 = v5
																																				m.G0 = v16 + int32(16)
																																				return v228
																																			} else {
																																				F_errmsg_internal(m, int32(_a_F_operator_predicate_proof_1), int32(0))
																																				mBase = m.M
																																				v212 = m.ExcPending
																																				if v212 != 0 {
																																					return int32(0)
																																				} else {
																																					F_errfinish(m, int32(_a_F_operator_predicate_proof_2), int32(2017), int32(_a_F_operator_predicate_proof_3))
																																					mBase = m.M
																																					v217 = m.ExcPending
																																					if v217 != 0 {
																																						return int32(0)
																																					} else {
																																						v228 = v5
																																						m.G0 = v16 + int32(16)
																																						return v228
																																					}
																																				}
																																			}
																																		}
																																	} else {
																																		v228 = base.B2i32(v194 != int32(0))
																																		m.G0 = v16 + int32(16)
																																		return v228
																																	}
																																}
																															}
																														}
																													} else {
																														v187 = v182
																														v189 = *(*int32)(unsafe.Add(mBase, uint32(v187)+20))
																														*(*int32)(unsafe.Add(mBase, _c_F_operator_predicate_proof[0])) = v189
																														v193 = *(*int32)(unsafe.Add(mBase, uint32(v180)+20))
																														v194 = m.T0[v193].(func(*base.Module, int32, int32, int32) int32)(m, v180, v187, v16+int32(15))
																														mBase = m.M
																														v195 = m.ExcPending
																														if v195 != 0 {
																															return int32(0)
																														} else {
																															*(*int32)(unsafe.Add(mBase, _c_F_operator_predicate_proof[0])) = v171
																															F_FreeExecutorState(m, v168)
																															mBase = m.M
																															v199 = m.ExcPending
																															if v199 != 0 {
																																return int32(0)
																															} else {
																																v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+15)))
																																if v200 == int32(1) {
																																	v205 = F_errstart(m, int32(13), int32(0))
																																	mBase = m.M
																																	v206 = m.ExcPending
																																	if v206 != 0 {
																																		return int32(0)
																																	} else {
																																		if v205 == int32(0) {
																																			v228 = v5
																																			m.G0 = v16 + int32(16)
																																			return v228
																																		} else {
																																			F_errmsg_internal(m, int32(_a_F_operator_predicate_proof_1), int32(0))
																																			mBase = m.M
																																			v212 = m.ExcPending
																																			if v212 != 0 {
																																				return int32(0)
																																			} else {
																																				F_errfinish(m, int32(_a_F_operator_predicate_proof_2), int32(2017), int32(_a_F_operator_predicate_proof_3))
																																				mBase = m.M
																																				v217 = m.ExcPending
																																				if v217 != 0 {
																																					return int32(0)
																																				} else {
																																					v228 = v5
																																					m.G0 = v16 + int32(16)
																																					return v228
																																				}
																																			}
																																		}
																																	}
																																} else {
																																	v228 = base.B2i32(v194 != int32(0))
																																	m.G0 = v16 + int32(16)
																																	return v228
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
