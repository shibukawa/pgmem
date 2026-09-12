package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_domainAddCheckConstraint(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
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
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v100 int32
	_ = v100
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	if v14 != 0 {
		v16 = F_ConstraintNameIsUsed(m, int32(1), l0, v14)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			if v16 == int32(0) {
				v47 = F_make_parsestate(m, int32(0))
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return int32(0)
				} else {
					v50 = F_palloc0(m, int32(20))
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v50)+8)) = l3
						*(*int32)(unsafe.Add(mBase, uint32(v50)+4)) = l2
						*(*int32)(unsafe.Add(mBase, uint32(v50))) = int32(56)
						v56 = F_get_typcollation(m, l2)
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v50)+16)) = int32(-1)
							*(*int32)(unsafe.Add(mBase, uint32(v50)+12)) = v56
							*(*int32)(unsafe.Add(mBase, uint32(v47)+120)) = v50
							*(*int32)(unsafe.Add(mBase, uint32(v47)+104)) = int32(582)
							v64 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
							v66 = F_transformExpr(m, v47, v64, int32(29))
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return int32(0)
							} else {
								v69 = F_coerce_to_boolean(m, v47, v66, int32(546860))
								mBase = m.M
								v70 = m.ExcPending
								if v70 != 0 {
									return int32(0)
								} else {
									F_assign_expr_collations(m, v47, v69)
									mBase = m.M
									v72 = m.ExcPending
									if v72 != 0 {
										return int32(0)
									} else {
										v73 = *(*int32)(unsafe.Add(mBase, uint32(v47)+8))
										if v73 != 0 {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v125 = m.ExcPending
											if v125 != 0 {
												return int32(0)
											} else {
												F_errcode(m, int32(393348))
												mBase = m.M
												v128 = m.ExcPending
												if v128 != 0 {
													return int32(0)
												} else {
													F_errmsg(m, int32(91803), int32(0))
													mBase = m.M
													v132 = m.ExcPending
													if v132 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(505906), int32(3578), int32(92493))
														mBase = m.M
														v137 = m.ExcPending
														if v137 != 0 {
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
											v74 = F_contain_var_clause(m, v69)
											mBase = m.M
											v75 = m.ExcPending
											if v75 != 0 {
												return int32(0)
											} else {
												if v74 != 0 {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v125 = m.ExcPending
													if v125 != 0 {
														return int32(0)
													} else {
														F_errcode(m, int32(393348))
														mBase = m.M
														v128 = m.ExcPending
														if v128 != 0 {
															return int32(0)
														} else {
															F_errmsg(m, int32(91803), int32(0))
															mBase = m.M
															v132 = m.ExcPending
															if v132 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(505906), int32(3578), int32(92493))
																mBase = m.M
																v137 = m.ExcPending
																if v137 != 0 {
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
													v76 = F_nodeToString(m, v69)
													mBase = m.M
													v77 = m.ExcPending
													if v77 != 0 {
														return int32(0)
													} else {
														v78 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
														v80 = int32(0)
														v82 = int32(1)
														v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+15)))
														v100 = int32(32)
														v111 = F_CreateConstraintEntry(m, v78, l1, int32(99), v80, v80, v82, (v83^int32(-1))&v82, v80, v80, v80, v80, v80, l0, v80, v80, v80, v80, v80, v80, v80, v100, v100, v80, v80, v100, v80, v69, v76, v82, v80, v80, v80, v80)
														mBase = m.M
														v112 = m.ExcPending
														if v112 != 0 {
															return int32(0)
														} else {
															if l6 != 0 {
																*(*int32)(unsafe.Add(mBase, uint32(l6)+8)) = int32(0)
																*(*int32)(unsafe.Add(mBase, uint32(l6)+4)) = v111
																*(*int32)(unsafe.Add(mBase, uint32(l6))) = int32(2606)
															} else {
															}
															m.G0 = v12 + int32(16)
															return v76
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
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(290948))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						v29 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = l5
						*(*int32)(unsafe.Add(mBase, uint32(v12))) = v29
						F_errmsg(m, int32(119402), v12)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(505906), int32(3528), int32(92493))
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
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
		v40 = int32(0)
		v43 = F_ChooseConstraintName(m, l5, v40, int32(324940), l1, v40)
		mBase = m.M
		v44 = m.ExcPending
		if v44 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l4)+8)) = v43
			v47 = F_make_parsestate(m, int32(0))
			mBase = m.M
			v48 = m.ExcPending
			if v48 != 0 {
				return int32(0)
			} else {
				v50 = F_palloc0(m, int32(20))
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v50)+8)) = l3
					*(*int32)(unsafe.Add(mBase, uint32(v50)+4)) = l2
					*(*int32)(unsafe.Add(mBase, uint32(v50))) = int32(56)
					v56 = F_get_typcollation(m, l2)
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v50)+16)) = int32(-1)
						*(*int32)(unsafe.Add(mBase, uint32(v50)+12)) = v56
						*(*int32)(unsafe.Add(mBase, uint32(v47)+120)) = v50
						*(*int32)(unsafe.Add(mBase, uint32(v47)+104)) = int32(582)
						v64 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
						v66 = F_transformExpr(m, v47, v64, int32(29))
						mBase = m.M
						v67 = m.ExcPending
						if v67 != 0 {
							return int32(0)
						} else {
							v69 = F_coerce_to_boolean(m, v47, v66, int32(546860))
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return int32(0)
							} else {
								F_assign_expr_collations(m, v47, v69)
								mBase = m.M
								v72 = m.ExcPending
								if v72 != 0 {
									return int32(0)
								} else {
									v73 = *(*int32)(unsafe.Add(mBase, uint32(v47)+8))
									if v73 != 0 {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v125 = m.ExcPending
										if v125 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(393348))
											mBase = m.M
											v128 = m.ExcPending
											if v128 != 0 {
												return int32(0)
											} else {
												F_errmsg(m, int32(91803), int32(0))
												mBase = m.M
												v132 = m.ExcPending
												if v132 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(505906), int32(3578), int32(92493))
													mBase = m.M
													v137 = m.ExcPending
													if v137 != 0 {
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
										v74 = F_contain_var_clause(m, v69)
										mBase = m.M
										v75 = m.ExcPending
										if v75 != 0 {
											return int32(0)
										} else {
											if v74 != 0 {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v125 = m.ExcPending
												if v125 != 0 {
													return int32(0)
												} else {
													F_errcode(m, int32(393348))
													mBase = m.M
													v128 = m.ExcPending
													if v128 != 0 {
														return int32(0)
													} else {
														F_errmsg(m, int32(91803), int32(0))
														mBase = m.M
														v132 = m.ExcPending
														if v132 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(505906), int32(3578), int32(92493))
															mBase = m.M
															v137 = m.ExcPending
															if v137 != 0 {
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
												v76 = F_nodeToString(m, v69)
												mBase = m.M
												v77 = m.ExcPending
												if v77 != 0 {
													return int32(0)
												} else {
													v78 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
													v80 = int32(0)
													v82 = int32(1)
													v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+15)))
													v100 = int32(32)
													v111 = F_CreateConstraintEntry(m, v78, l1, int32(99), v80, v80, v82, (v83^int32(-1))&v82, v80, v80, v80, v80, v80, l0, v80, v80, v80, v80, v80, v80, v80, v100, v100, v80, v80, v100, v80, v69, v76, v82, v80, v80, v80, v80)
													mBase = m.M
													v112 = m.ExcPending
													if v112 != 0 {
														return int32(0)
													} else {
														if l6 != 0 {
															*(*int32)(unsafe.Add(mBase, uint32(l6)+8)) = int32(0)
															*(*int32)(unsafe.Add(mBase, uint32(l6)+4)) = v111
															*(*int32)(unsafe.Add(mBase, uint32(l6))) = int32(2606)
														} else {
														}
														m.G0 = v12 + int32(16)
														return v76
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
func F_domain_in(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
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
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
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
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	v2 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v12 == v2 {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v16 = v15
	} else {
		v16 = v2
	}
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v17 == int32(1) {
		v20 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v20)
		v66 = int32(0)
		m.G0 = v10 + int32(16)
		return v66
	} else {
		v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
		if v26 != 0 {
			v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
			if v27 == v24 {
				v37 = v26
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
				v41 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
				v44 = F_InputFunctionCallSafe(m, v37+int32(16), v16, v40, v41, v23, v10+int32(12))
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return int32(0)
				} else {
					if v44 == int32(0) {
						v48 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v48)
						v66 = int32(0)
						m.G0 = v10 + int32(16)
						return v66
					} else {
						v51 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
						F_domain_check_input(m, v51, base.B2i32(v16 == int32(0)), v37, v23)
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int32(0)
						} else {
							if v16 == int32(0) {
								v58 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v58)
								v66 = int32(0)
							} else {
								v61 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
								v66 = v61
							}
							m.G0 = v10 + int32(16)
							return v66
						}
					}
				}
			} else {
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
				v31 = F_domain_state_setup(m, v24, int32(0), v30)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					*(*int32)(unsafe.Add(mBase, uint32(v35)+16)) = v31
					v37 = v31
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
					v41 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
					v44 = F_InputFunctionCallSafe(m, v37+int32(16), v16, v40, v41, v23, v10+int32(12))
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return int32(0)
					} else {
						if v44 == int32(0) {
							v48 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v48)
							v66 = int32(0)
							m.G0 = v10 + int32(16)
							return v66
						} else {
							v51 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
							F_domain_check_input(m, v51, base.B2i32(v16 == int32(0)), v37, v23)
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
								return int32(0)
							} else {
								if v16 == int32(0) {
									v58 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v58)
									v66 = int32(0)
								} else {
									v61 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
									v66 = v61
								}
								m.G0 = v10 + int32(16)
								return v66
							}
						}
					}
				}
			}
		} else {
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
			v31 = F_domain_state_setup(m, v24, int32(0), v30)
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				*(*int32)(unsafe.Add(mBase, uint32(v35)+16)) = v31
				v37 = v31
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
				v41 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
				v44 = F_InputFunctionCallSafe(m, v37+int32(16), v16, v40, v41, v23, v10+int32(12))
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return int32(0)
				} else {
					if v44 == int32(0) {
						v48 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v48)
						v66 = int32(0)
						m.G0 = v10 + int32(16)
						return v66
					} else {
						v51 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
						F_domain_check_input(m, v51, base.B2i32(v16 == int32(0)), v37, v23)
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int32(0)
						} else {
							if v16 == int32(0) {
								v58 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v58)
								v66 = int32(0)
							} else {
								v61 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
								v66 = v61
							}
							m.G0 = v10 + int32(16)
							return v66
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
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
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
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	v9 = m.G0
	v11 = v9 - int32(160)
	m.G0 = v11
	v15 = F_table_open(m, int32(2606), int32(1))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		F_ScanKeyInit(m, v11+int32(16), int32(9), int32(3), int32(184), int32(0))
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return int32(0)
		} else {
			F_ScanKeyInit(m, v11-int32(-64), int32(10), int32(3), int32(184), l0)
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				F_ScanKeyInit(m, v11+int32(112), int32(2), int32(3), int32(62), l1)
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return int32(0)
				} else {
					v47 = F_systable_beginscan(m, v15, int32(2665), int32(1), int32(0), int32(3), v11+int32(16))
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
								if l2 != 0 {
									F_sequence_close(m, v15, int32(1))
									mBase = m.M
									v79 = m.ExcPending
									if v79 != 0 {
										return int32(0)
									} else {
										m.G0 = v11 + int32(160)
										return v55
									}
								} else {
									if v55 != 0 {
										F_sequence_close(m, v15, int32(1))
										mBase = m.M
										v79 = m.ExcPending
										if v79 != 0 {
											return int32(0)
										} else {
											m.G0 = v11 + int32(160)
											return v55
										}
									} else {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v61 = m.ExcPending
										if v61 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(67137668))
											mBase = m.M
											v64 = m.ExcPending
											if v64 != 0 {
												return int32(0)
											} else {
												v65 = F_format_type_be(m, l0)
												mBase = m.M
												v66 = m.ExcPending
												if v66 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v65
													*(*int32)(unsafe.Add(mBase, uint32(v11))) = l1
													F_errmsg(m, int32(71363), v11)
													mBase = m.M
													v71 = m.ExcPending
													if v71 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(504071), int32(1428), int32(443457))
														mBase = m.M
														v76 = m.ExcPending
														if v76 != 0 {
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
