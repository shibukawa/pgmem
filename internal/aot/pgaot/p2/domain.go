package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_domainAddCheckConstraint(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
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
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v108 int32
	_ = v108
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	if v15 != 0 {
		v17 = F_ConstraintNameIsUsed(m, int32(1), l0, v15)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			if v17 == int32(0) {
				v48 = F_make_parsestate(m, int32(0))
				mBase = m.M
				v49 = m.ExcPending
				if v49 != 0 {
					return int32(0)
				} else {
					v51 = F_palloc0(m, int32(20))
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v51)+8)) = l3
						*(*int32)(unsafe.Add(mBase, uint32(v51)+4)) = l2
						*(*int32)(unsafe.Add(mBase, uint32(v51))) = int32(56)
						v57 = F_get_typcollation(m, l2)
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v51)+16)) = int32(-1)
							*(*int32)(unsafe.Add(mBase, uint32(v51)+12)) = v57
							*(*int32)(unsafe.Add(mBase, uint32(v48)+116)) = v51
							*(*int32)(unsafe.Add(mBase, uint32(v48)+100)) = int32(627)
							v65 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
							v67 = F_transformExpr(m, v48, v65, int32(29))
							mBase = m.M
							v68 = m.ExcPending
							if v68 != 0 {
								return int32(0)
							} else {
								v70 = F_coerce_to_boolean(m, v48, v67, int32(_a_F_domainAddCheckConstraint_0))
								mBase = m.M
								v71 = m.ExcPending
								if v71 != 0 {
									return int32(0)
								} else {
									F_assign_expr_collations(m, v48, v70)
									mBase = m.M
									v73 = m.ExcPending
									if v73 != 0 {
										return int32(0)
									} else {
										if l7 == int32(0) {
											v78 = *(*int32)(unsafe.Add(mBase, _c_F_domainAddCheckConstraint[0]))
											F_CheckUsageOnTypesInExpr(m, v70, int32(0), v78)
											mBase = m.M
											v80 = m.ExcPending
											if v80 != 0 {
												return int32(0)
											} else {
												v81 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
												if v81 != 0 {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v133 = m.ExcPending
													if v133 != 0 {
														return int32(0)
													} else {
														F_errcode(m, int32(_a_F_domainAddCheckConstraint_1))
														mBase = m.M
														v136 = m.ExcPending
														if v136 != 0 {
															return int32(0)
														} else {
															F_errmsg(m, int32(_a_F_domainAddCheckConstraint_2), int32(0))
															mBase = m.M
															v140 = m.ExcPending
															if v140 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(_a_F_domainAddCheckConstraint_3), int32(3646), int32(_a_F_domainAddCheckConstraint_4))
																mBase = m.M
																v145 = m.ExcPending
																if v145 != 0 {
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
													v82 = F_contain_var_clause(m, v70)
													mBase = m.M
													v83 = m.ExcPending
													if v83 != 0 {
														return int32(0)
													} else {
														if v82 != 0 {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v133 = m.ExcPending
															if v133 != 0 {
																return int32(0)
															} else {
																F_errcode(m, int32(_a_F_domainAddCheckConstraint_1))
																mBase = m.M
																v136 = m.ExcPending
																if v136 != 0 {
																	return int32(0)
																} else {
																	F_errmsg(m, int32(_a_F_domainAddCheckConstraint_2), int32(0))
																	mBase = m.M
																	v140 = m.ExcPending
																	if v140 != 0 {
																		return int32(0)
																	} else {
																		F_errfinish(m, int32(_a_F_domainAddCheckConstraint_3), int32(3646), int32(_a_F_domainAddCheckConstraint_4))
																		mBase = m.M
																		v145 = m.ExcPending
																		if v145 != 0 {
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
															v84 = F_nodeToString(m, v70)
															mBase = m.M
															v85 = m.ExcPending
															if v85 != 0 {
																return int32(0)
															} else {
																v86 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
																v88 = int32(0)
																v90 = int32(1)
																v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+15)))
																v108 = int32(32)
																v119 = F_CreateConstraintEntry(m, v86, l1, int32(99), v88, v88, v90, (v91^int32(-1))&v90, v88, v88, v88, v88, v88, l0, v88, v88, v88, v88, v88, v88, v88, v108, v108, v88, v88, v108, v88, v70, v84, v90, v88, v88, v88, v88)
																mBase = m.M
																v120 = m.ExcPending
																if v120 != 0 {
																	return int32(0)
																} else {
																	if l6 != 0 {
																		*(*int32)(unsafe.Add(mBase, uint32(l6)+8)) = int32(0)
																		*(*int32)(unsafe.Add(mBase, uint32(l6)+4)) = v119
																		*(*int32)(unsafe.Add(mBase, uint32(l6))) = int32(2606)
																	} else {
																	}
																	m.G0 = v13 + int32(16)
																	return v84
																}
															}
														}
													}
												}
											}
										} else {
											v81 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
											if v81 != 0 {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v133 = m.ExcPending
												if v133 != 0 {
													return int32(0)
												} else {
													F_errcode(m, int32(_a_F_domainAddCheckConstraint_1))
													mBase = m.M
													v136 = m.ExcPending
													if v136 != 0 {
														return int32(0)
													} else {
														F_errmsg(m, int32(_a_F_domainAddCheckConstraint_2), int32(0))
														mBase = m.M
														v140 = m.ExcPending
														if v140 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_domainAddCheckConstraint_3), int32(3646), int32(_a_F_domainAddCheckConstraint_4))
															mBase = m.M
															v145 = m.ExcPending
															if v145 != 0 {
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
												v82 = F_contain_var_clause(m, v70)
												mBase = m.M
												v83 = m.ExcPending
												if v83 != 0 {
													return int32(0)
												} else {
													if v82 != 0 {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v133 = m.ExcPending
														if v133 != 0 {
															return int32(0)
														} else {
															F_errcode(m, int32(_a_F_domainAddCheckConstraint_1))
															mBase = m.M
															v136 = m.ExcPending
															if v136 != 0 {
																return int32(0)
															} else {
																F_errmsg(m, int32(_a_F_domainAddCheckConstraint_2), int32(0))
																mBase = m.M
																v140 = m.ExcPending
																if v140 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(_a_F_domainAddCheckConstraint_3), int32(3646), int32(_a_F_domainAddCheckConstraint_4))
																	mBase = m.M
																	v145 = m.ExcPending
																	if v145 != 0 {
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
														v84 = F_nodeToString(m, v70)
														mBase = m.M
														v85 = m.ExcPending
														if v85 != 0 {
															return int32(0)
														} else {
															v86 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
															v88 = int32(0)
															v90 = int32(1)
															v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+15)))
															v108 = int32(32)
															v119 = F_CreateConstraintEntry(m, v86, l1, int32(99), v88, v88, v90, (v91^int32(-1))&v90, v88, v88, v88, v88, v88, l0, v88, v88, v88, v88, v88, v88, v88, v108, v108, v88, v88, v108, v88, v70, v84, v90, v88, v88, v88, v88)
															mBase = m.M
															v120 = m.ExcPending
															if v120 != 0 {
																return int32(0)
															} else {
																if l6 != 0 {
																	*(*int32)(unsafe.Add(mBase, uint32(l6)+8)) = int32(0)
																	*(*int32)(unsafe.Add(mBase, uint32(l6)+4)) = v119
																	*(*int32)(unsafe.Add(mBase, uint32(l6))) = int32(2606)
																} else {
																}
																m.G0 = v13 + int32(16)
																return v84
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
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(_a_F_domainAddCheckConstraint_5))
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int32(0)
					} else {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = l5
						*(*int32)(unsafe.Add(mBase, uint32(v13))) = v30
						F_errmsg(m, int32(_a_F_domainAddCheckConstraint_6), v13)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_domainAddCheckConstraint_3), int32(3589), int32(_a_F_domainAddCheckConstraint_4))
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
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
	} else {
		v41 = int32(0)
		v44 = F_ChooseConstraintName(m, l5, v41, int32(_a_F_domainAddCheckConstraint_7), l1, v41)
		mBase = m.M
		v45 = m.ExcPending
		if v45 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l4)+8)) = v44
			v48 = F_make_parsestate(m, int32(0))
			mBase = m.M
			v49 = m.ExcPending
			if v49 != 0 {
				return int32(0)
			} else {
				v51 = F_palloc0(m, int32(20))
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v51)+8)) = l3
					*(*int32)(unsafe.Add(mBase, uint32(v51)+4)) = l2
					*(*int32)(unsafe.Add(mBase, uint32(v51))) = int32(56)
					v57 = F_get_typcollation(m, l2)
					mBase = m.M
					v58 = m.ExcPending
					if v58 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v51)+16)) = int32(-1)
						*(*int32)(unsafe.Add(mBase, uint32(v51)+12)) = v57
						*(*int32)(unsafe.Add(mBase, uint32(v48)+116)) = v51
						*(*int32)(unsafe.Add(mBase, uint32(v48)+100)) = int32(627)
						v65 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
						v67 = F_transformExpr(m, v48, v65, int32(29))
						mBase = m.M
						v68 = m.ExcPending
						if v68 != 0 {
							return int32(0)
						} else {
							v70 = F_coerce_to_boolean(m, v48, v67, int32(_a_F_domainAddCheckConstraint_0))
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
								return int32(0)
							} else {
								F_assign_expr_collations(m, v48, v70)
								mBase = m.M
								v73 = m.ExcPending
								if v73 != 0 {
									return int32(0)
								} else {
									if l7 == int32(0) {
										v78 = *(*int32)(unsafe.Add(mBase, _c_F_domainAddCheckConstraint[0]))
										F_CheckUsageOnTypesInExpr(m, v70, int32(0), v78)
										mBase = m.M
										v80 = m.ExcPending
										if v80 != 0 {
											return int32(0)
										} else {
											v81 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
											if v81 != 0 {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v133 = m.ExcPending
												if v133 != 0 {
													return int32(0)
												} else {
													F_errcode(m, int32(_a_F_domainAddCheckConstraint_1))
													mBase = m.M
													v136 = m.ExcPending
													if v136 != 0 {
														return int32(0)
													} else {
														F_errmsg(m, int32(_a_F_domainAddCheckConstraint_2), int32(0))
														mBase = m.M
														v140 = m.ExcPending
														if v140 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_domainAddCheckConstraint_3), int32(3646), int32(_a_F_domainAddCheckConstraint_4))
															mBase = m.M
															v145 = m.ExcPending
															if v145 != 0 {
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
												v82 = F_contain_var_clause(m, v70)
												mBase = m.M
												v83 = m.ExcPending
												if v83 != 0 {
													return int32(0)
												} else {
													if v82 != 0 {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v133 = m.ExcPending
														if v133 != 0 {
															return int32(0)
														} else {
															F_errcode(m, int32(_a_F_domainAddCheckConstraint_1))
															mBase = m.M
															v136 = m.ExcPending
															if v136 != 0 {
																return int32(0)
															} else {
																F_errmsg(m, int32(_a_F_domainAddCheckConstraint_2), int32(0))
																mBase = m.M
																v140 = m.ExcPending
																if v140 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(_a_F_domainAddCheckConstraint_3), int32(3646), int32(_a_F_domainAddCheckConstraint_4))
																	mBase = m.M
																	v145 = m.ExcPending
																	if v145 != 0 {
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
														v84 = F_nodeToString(m, v70)
														mBase = m.M
														v85 = m.ExcPending
														if v85 != 0 {
															return int32(0)
														} else {
															v86 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
															v88 = int32(0)
															v90 = int32(1)
															v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+15)))
															v108 = int32(32)
															v119 = F_CreateConstraintEntry(m, v86, l1, int32(99), v88, v88, v90, (v91^int32(-1))&v90, v88, v88, v88, v88, v88, l0, v88, v88, v88, v88, v88, v88, v88, v108, v108, v88, v88, v108, v88, v70, v84, v90, v88, v88, v88, v88)
															mBase = m.M
															v120 = m.ExcPending
															if v120 != 0 {
																return int32(0)
															} else {
																if l6 != 0 {
																	*(*int32)(unsafe.Add(mBase, uint32(l6)+8)) = int32(0)
																	*(*int32)(unsafe.Add(mBase, uint32(l6)+4)) = v119
																	*(*int32)(unsafe.Add(mBase, uint32(l6))) = int32(2606)
																} else {
																}
																m.G0 = v13 + int32(16)
																return v84
															}
														}
													}
												}
											}
										}
									} else {
										v81 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
										if v81 != 0 {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v133 = m.ExcPending
											if v133 != 0 {
												return int32(0)
											} else {
												F_errcode(m, int32(_a_F_domainAddCheckConstraint_1))
												mBase = m.M
												v136 = m.ExcPending
												if v136 != 0 {
													return int32(0)
												} else {
													F_errmsg(m, int32(_a_F_domainAddCheckConstraint_2), int32(0))
													mBase = m.M
													v140 = m.ExcPending
													if v140 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_domainAddCheckConstraint_3), int32(3646), int32(_a_F_domainAddCheckConstraint_4))
														mBase = m.M
														v145 = m.ExcPending
														if v145 != 0 {
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
											v82 = F_contain_var_clause(m, v70)
											mBase = m.M
											v83 = m.ExcPending
											if v83 != 0 {
												return int32(0)
											} else {
												if v82 != 0 {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v133 = m.ExcPending
													if v133 != 0 {
														return int32(0)
													} else {
														F_errcode(m, int32(_a_F_domainAddCheckConstraint_1))
														mBase = m.M
														v136 = m.ExcPending
														if v136 != 0 {
															return int32(0)
														} else {
															F_errmsg(m, int32(_a_F_domainAddCheckConstraint_2), int32(0))
															mBase = m.M
															v140 = m.ExcPending
															if v140 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(_a_F_domainAddCheckConstraint_3), int32(3646), int32(_a_F_domainAddCheckConstraint_4))
																mBase = m.M
																v145 = m.ExcPending
																if v145 != 0 {
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
													v84 = F_nodeToString(m, v70)
													mBase = m.M
													v85 = m.ExcPending
													if v85 != 0 {
														return int32(0)
													} else {
														v86 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
														v88 = int32(0)
														v90 = int32(1)
														v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+15)))
														v108 = int32(32)
														v119 = F_CreateConstraintEntry(m, v86, l1, int32(99), v88, v88, v90, (v91^int32(-1))&v90, v88, v88, v88, v88, v88, l0, v88, v88, v88, v88, v88, v88, v88, v108, v108, v88, v88, v108, v88, v70, v84, v90, v88, v88, v88, v88)
														mBase = m.M
														v120 = m.ExcPending
														if v120 != 0 {
															return int32(0)
														} else {
															if l6 != 0 {
																*(*int32)(unsafe.Add(mBase, uint32(l6)+8)) = int32(0)
																*(*int32)(unsafe.Add(mBase, uint32(l6)+4)) = v119
																*(*int32)(unsafe.Add(mBase, uint32(l6))) = int32(2606)
															} else {
															}
															m.G0 = v13 + int32(16)
															return v84
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
func F_domain_in(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
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
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int64
	_ = v52
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int64
	_ = v62
	var v67 int64
	_ = v67
	v2 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v13 == v2 {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		v17 = v16
	} else {
		v17 = v2
	}
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
	if v18 == int32(1) {
		v21 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v21)
		v67 = int64(0)
		m.G0 = v11 + int32(16)
		return v67
	} else {
		v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
		if v27 != 0 {
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
			if v28 == v25 {
				v38 = v27
				v41 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
				v42 = *(*int32)(unsafe.Add(mBase, uint32(v38)+12))
				v45 = F_InputFunctionCallSafe(m, v38+int32(16), v17, v41, v42, v24, v11+int32(8))
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return int64(0)
				} else {
					if v45 == int32(0) {
						v49 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v49)
						v67 = int64(0)
						m.G0 = v11 + int32(16)
						return v67
					} else {
						v52 = *(*int64)(unsafe.Add(mBase, uint32(v11)+8))
						F_domain_check_input(m, v52, base.B2i32(v17 == int32(0)), v38, v24)
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return int64(0)
						} else {
							if v17 == int32(0) {
								v59 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v59)
								v67 = int64(0)
							} else {
								v62 = *(*int64)(unsafe.Add(mBase, uint32(v11)+8))
								v67 = v62
							}
							m.G0 = v11 + int32(16)
							return v67
						}
					}
				}
			} else {
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
				v32 = F_domain_state_setup(m, v25, int32(0), v31)
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return int64(0)
				} else {
					v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					*(*int32)(unsafe.Add(mBase, uint32(v36)+16)) = v32
					v38 = v32
					v41 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
					v42 = *(*int32)(unsafe.Add(mBase, uint32(v38)+12))
					v45 = F_InputFunctionCallSafe(m, v38+int32(16), v17, v41, v42, v24, v11+int32(8))
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return int64(0)
					} else {
						if v45 == int32(0) {
							v49 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v49)
							v67 = int64(0)
							m.G0 = v11 + int32(16)
							return v67
						} else {
							v52 = *(*int64)(unsafe.Add(mBase, uint32(v11)+8))
							F_domain_check_input(m, v52, base.B2i32(v17 == int32(0)), v38, v24)
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return int64(0)
							} else {
								if v17 == int32(0) {
									v59 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v59)
									v67 = int64(0)
								} else {
									v62 = *(*int64)(unsafe.Add(mBase, uint32(v11)+8))
									v67 = v62
								}
								m.G0 = v11 + int32(16)
								return v67
							}
						}
					}
				}
			}
		} else {
			v31 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
			v32 = F_domain_state_setup(m, v25, int32(0), v31)
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return int64(0)
			} else {
				v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				*(*int32)(unsafe.Add(mBase, uint32(v36)+16)) = v32
				v38 = v32
				v41 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
				v42 = *(*int32)(unsafe.Add(mBase, uint32(v38)+12))
				v45 = F_InputFunctionCallSafe(m, v38+int32(16), v17, v41, v42, v24, v11+int32(8))
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return int64(0)
				} else {
					if v45 == int32(0) {
						v49 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v49)
						v67 = int64(0)
						m.G0 = v11 + int32(16)
						return v67
					} else {
						v52 = *(*int64)(unsafe.Add(mBase, uint32(v11)+8))
						F_domain_check_input(m, v52, base.B2i32(v17 == int32(0)), v38, v24)
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return int64(0)
						} else {
							if v17 == int32(0) {
								v59 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v59)
								v67 = int64(0)
							} else {
								v62 = *(*int64)(unsafe.Add(mBase, uint32(v11)+8))
								v67 = v62
							}
							m.G0 = v11 + int32(16)
							return v67
						}
					}
				}
			}
		}
	}
}
func F_get_domain_constraint_oid(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
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
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	v9 = m.G0
	v11 = v9 - int32(192)
	m.G0 = v11
	v15 = F_table_open(m, int32(2606), int32(1))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v20 = v11 + int32(16)
		F_ScanKeyInit(m, v20, int32(9), int32(3), int32(184), int64(0))
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return int32(0)
		} else {
			F_ScanKeyInit(m, v11+int32(72), int32(10), int32(3), int32(184), base.I64_extend_i32_u(l0))
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				F_ScanKeyInit(m, v11+int32(128), int32(2), int32(3), int32(62), base.I64_extend_i32_u(l1))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					v47 = F_systable_beginscan(m, v15, int32(2665), int32(1), int32(0), int32(3), v20)
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return int32(0)
					} else {
						v49 = F_systable_getnext(m, v47)
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return int32(0)
						} else {
							if v49 != 0 {
								v51 = *(*int32)(unsafe.Add(mBase, uint32(v49)+16))
								v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+22)))
								v54 = *(*int32)(unsafe.Add(mBase, uint32(v51+v52)))
								v55 = v54
							} else {
								v55 = int32(0)
							}
							F_systable_endscan(m, v47)
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return int32(0)
							} else {
								if l2|v55 == int32(0) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v64 = m.ExcPending
									if v64 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(67137668))
										mBase = m.M
										v67 = m.ExcPending
										if v67 != 0 {
											return int32(0)
										} else {
											v68 = F_format_type_be(m, l0)
											mBase = m.M
											v69 = m.ExcPending
											if v69 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v68
												*(*int32)(unsafe.Add(mBase, uint32(v11))) = l1
												F_errmsg(m, int32(_a_F_get_domain_constraint_oid_0), v11)
												mBase = m.M
												v74 = m.ExcPending
												if v74 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_get_domain_constraint_oid_1), int32(1433), int32(_a_F_get_domain_constraint_oid_2))
													mBase = m.M
													v79 = m.ExcPending
													if v79 != 0 {
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
									F_relation_close(m, v15, int32(1))
									mBase = m.M
									v82 = m.ExcPending
									if v82 != 0 {
										return int32(0)
									} else {
										m.G0 = v11 + int32(192)
										return v55
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
