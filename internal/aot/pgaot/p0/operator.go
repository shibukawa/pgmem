package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_OperatorValidateParams(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) {
	var v11 int32
	_ = v11
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	if l1 != 0 {
		v11 = l0
	} else {
		v11 = int32(0)
	}
	if v11 == int32(0) {
		if l3 != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				F_errcode(m, int32(50724996))
				v22 = m.ExcPending
				if v22 != 0 {
					return
				} else {
					F_errmsg(m, int32(_a_F_OperatorValidateParams_0), int32(0))
					v26 = m.ExcPending
					if v26 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_OperatorValidateParams_1), int32(572), int32(_a_F_OperatorValidateParams_2))
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
			}
		} else {
			if l6 != 0 {
				F_errstart_cold(m, int32(21), int32(0))
				v35 = m.ExcPending
				if v35 != 0 {
					return
				} else {
					F_errcode(m, int32(50724996))
					v38 = m.ExcPending
					if v38 != 0 {
						return
					} else {
						F_errmsg(m, int32(_a_F_OperatorValidateParams_3), int32(0))
						v42 = m.ExcPending
						if v42 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_OperatorValidateParams_1), int32(576), int32(_a_F_OperatorValidateParams_2))
							v47 = m.ExcPending
							if v47 != 0 {
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
				if l7 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					v51 = m.ExcPending
					if v51 != 0 {
						return
					} else {
						F_errcode(m, int32(50724996))
						v54 = m.ExcPending
						if v54 != 0 {
							return
						} else {
							F_errmsg(m, int32(_a_F_OperatorValidateParams_4), int32(0))
							v58 = m.ExcPending
							if v58 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_OperatorValidateParams_1), int32(580), int32(_a_F_OperatorValidateParams_2))
								v63 = m.ExcPending
								if v63 != 0 {
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
					if l8 != 0 {
						F_errstart_cold(m, int32(21), int32(0))
						v67 = m.ExcPending
						if v67 != 0 {
							return
						} else {
							F_errcode(m, int32(50724996))
							v70 = m.ExcPending
							if v70 != 0 {
								return
							} else {
								F_errmsg(m, int32(_a_F_OperatorValidateParams_5), int32(0))
								v74 = m.ExcPending
								if v74 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_OperatorValidateParams_1), int32(584), int32(_a_F_OperatorValidateParams_2))
									v79 = m.ExcPending
									if v79 != 0 {
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
						if l2 != int32(16) {
							if l4 != 0 {
								F_errstart_cold(m, int32(21), int32(0))
								v83 = m.ExcPending
								if v83 != 0 {
									return
								} else {
									F_errcode(m, int32(50724996))
									v86 = m.ExcPending
									if v86 != 0 {
										return
									} else {
										F_errmsg(m, int32(_a_F_OperatorValidateParams_6), int32(0))
										v90 = m.ExcPending
										if v90 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_OperatorValidateParams_1), int32(593), int32(_a_F_OperatorValidateParams_2))
											v95 = m.ExcPending
											if v95 != 0 {
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
								if l5 != 0 {
									F_errstart_cold(m, int32(21), int32(0))
									v99 = m.ExcPending
									if v99 != 0 {
										return
									} else {
										F_errcode(m, int32(50724996))
										v102 = m.ExcPending
										if v102 != 0 {
											return
										} else {
											F_errmsg(m, int32(_a_F_OperatorValidateParams_7), int32(0))
											v106 = m.ExcPending
											if v106 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_OperatorValidateParams_1), int32(597), int32(_a_F_OperatorValidateParams_2))
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
									if l6 != 0 {
										F_errstart_cold(m, int32(21), int32(0))
										v115 = m.ExcPending
										if v115 != 0 {
											return
										} else {
											F_errcode(m, int32(50724996))
											v118 = m.ExcPending
											if v118 != 0 {
												return
											} else {
												F_errmsg(m, int32(_a_F_OperatorValidateParams_8), int32(0))
												v122 = m.ExcPending
												if v122 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_OperatorValidateParams_1), int32(601), int32(_a_F_OperatorValidateParams_2))
													v127 = m.ExcPending
													if v127 != 0 {
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
										if l7 != 0 {
											F_errstart_cold(m, int32(21), int32(0))
											v131 = m.ExcPending
											if v131 != 0 {
												return
											} else {
												F_errcode(m, int32(50724996))
												v134 = m.ExcPending
												if v134 != 0 {
													return
												} else {
													F_errmsg(m, int32(_a_F_OperatorValidateParams_9), int32(0))
													v138 = m.ExcPending
													if v138 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_OperatorValidateParams_1), int32(605), int32(_a_F_OperatorValidateParams_2))
														v143 = m.ExcPending
														if v143 != 0 {
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
											if l8 != 0 {
												F_errstart_cold(m, int32(21), int32(0))
												v147 = m.ExcPending
												if v147 != 0 {
													return
												} else {
													F_errcode(m, int32(50724996))
													v150 = m.ExcPending
													if v150 != 0 {
														return
													} else {
														F_errmsg(m, int32(_a_F_OperatorValidateParams_10), int32(0))
														v154 = m.ExcPending
														if v154 != 0 {
															return
														} else {
															F_errfinish(m, int32(_a_F_OperatorValidateParams_1), int32(609), int32(_a_F_OperatorValidateParams_2))
															v159 = m.ExcPending
															if v159 != 0 {
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
												return
											}
										}
									}
								}
							}
						} else {
							return
						}
					}
				}
			}
		}
	} else {
		if l2 != int32(16) {
			if l4 != 0 {
				F_errstart_cold(m, int32(21), int32(0))
				v83 = m.ExcPending
				if v83 != 0 {
					return
				} else {
					F_errcode(m, int32(50724996))
					v86 = m.ExcPending
					if v86 != 0 {
						return
					} else {
						F_errmsg(m, int32(_a_F_OperatorValidateParams_6), int32(0))
						v90 = m.ExcPending
						if v90 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_OperatorValidateParams_1), int32(593), int32(_a_F_OperatorValidateParams_2))
							v95 = m.ExcPending
							if v95 != 0 {
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
				if l5 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					v99 = m.ExcPending
					if v99 != 0 {
						return
					} else {
						F_errcode(m, int32(50724996))
						v102 = m.ExcPending
						if v102 != 0 {
							return
						} else {
							F_errmsg(m, int32(_a_F_OperatorValidateParams_7), int32(0))
							v106 = m.ExcPending
							if v106 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_OperatorValidateParams_1), int32(597), int32(_a_F_OperatorValidateParams_2))
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
					if l6 != 0 {
						F_errstart_cold(m, int32(21), int32(0))
						v115 = m.ExcPending
						if v115 != 0 {
							return
						} else {
							F_errcode(m, int32(50724996))
							v118 = m.ExcPending
							if v118 != 0 {
								return
							} else {
								F_errmsg(m, int32(_a_F_OperatorValidateParams_8), int32(0))
								v122 = m.ExcPending
								if v122 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_OperatorValidateParams_1), int32(601), int32(_a_F_OperatorValidateParams_2))
									v127 = m.ExcPending
									if v127 != 0 {
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
						if l7 != 0 {
							F_errstart_cold(m, int32(21), int32(0))
							v131 = m.ExcPending
							if v131 != 0 {
								return
							} else {
								F_errcode(m, int32(50724996))
								v134 = m.ExcPending
								if v134 != 0 {
									return
								} else {
									F_errmsg(m, int32(_a_F_OperatorValidateParams_9), int32(0))
									v138 = m.ExcPending
									if v138 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_OperatorValidateParams_1), int32(605), int32(_a_F_OperatorValidateParams_2))
										v143 = m.ExcPending
										if v143 != 0 {
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
							if l8 != 0 {
								F_errstart_cold(m, int32(21), int32(0))
								v147 = m.ExcPending
								if v147 != 0 {
									return
								} else {
									F_errcode(m, int32(50724996))
									v150 = m.ExcPending
									if v150 != 0 {
										return
									} else {
										F_errmsg(m, int32(_a_F_OperatorValidateParams_10), int32(0))
										v154 = m.ExcPending
										if v154 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_OperatorValidateParams_1), int32(609), int32(_a_F_OperatorValidateParams_2))
											v159 = m.ExcPending
											if v159 != 0 {
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
								return
							}
						}
					}
				}
			}
		} else {
			return
		}
	}
}
func F_format_operator_extended(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
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
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
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
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
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
	var v96 int32
	_ = v96
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
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	v7 = m.G0
	v9 = v7 - int32(112)
	m.G0 = v9
	v13 = F_SearchSysCache1(m, int32(40), base.I64_extend_i32_u(l0))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		if v13 != 0 {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
			v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+22)))
			v19 = v17 + v18
			v21 = v19 + int32(4)
			F_initStringInfo(m, v9+int32(96))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				v27 = l1 & int32(2)
				if v27 == int32(0) {
					v31 = F_OperatorIsVisibleExt(m, l0, int32(0))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						if v31 != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v21
							F_appendStringInfo(m, v9+int32(96), int32(_a_F_format_operator_extended_0), v9+int32(48))
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return int32(0)
							} else {
								v67 = *(*int32)(unsafe.Add(mBase, uint32(v19)+80))
								if v67 != 0 {
									v84 = v67
									v85 = F_format_type_be(m, v84)
									mBase = m.M
									v86 = m.ExcPending
									if v86 != 0 {
										return int32(0)
									} else {
										v88 = v85
										*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v88
										F_appendStringInfo(m, v9+int32(96), int32(_a_F_format_operator_extended_1), v9+int32(32))
										mBase = m.M
										v96 = m.ExcPending
										if v96 != 0 {
											return int32(0)
										} else {
											v98 = *(*int32)(unsafe.Add(mBase, uint32(v19)+84))
											if v98 != 0 {
												if v27 != 0 {
													v99 = F_format_type_be_qualified(m, v98)
													mBase = m.M
													v100 = m.ExcPending
													if v100 != 0 {
														return int32(0)
													} else {
														v103 = v99
														*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v103
														F_appendStringInfo(m, v9+int32(96), int32(_a_F_format_operator_extended_2), v9+int32(16))
														mBase = m.M
														v111 = m.ExcPending
														if v111 != 0 {
															return int32(0)
														} else {
															v117 = *(*int32)(unsafe.Add(mBase, uint32(v9)+96))
															F_ReleaseCatCache(m, v13)
															mBase = m.M
															v119 = m.ExcPending
															if v119 != 0 {
																return int32(0)
															} else {
																v122 = v117
																m.G0 = v9 + int32(112)
																return v122
															}
														}
													}
												} else {
													v101 = F_format_type_be(m, v98)
													mBase = m.M
													v102 = m.ExcPending
													if v102 != 0 {
														return int32(0)
													} else {
														v103 = v101
														*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v103
														F_appendStringInfo(m, v9+int32(96), int32(_a_F_format_operator_extended_2), v9+int32(16))
														mBase = m.M
														v111 = m.ExcPending
														if v111 != 0 {
															return int32(0)
														} else {
															v117 = *(*int32)(unsafe.Add(mBase, uint32(v9)+96))
															F_ReleaseCatCache(m, v13)
															mBase = m.M
															v119 = m.ExcPending
															if v119 != 0 {
																return int32(0)
															} else {
																v122 = v117
																m.G0 = v9 + int32(112)
																return v122
															}
														}
													}
												}
											} else {
												F_appendStringInfoString(m, v9+int32(96), int32(_a_F_format_operator_extended_3))
												mBase = m.M
												v116 = m.ExcPending
												if v116 != 0 {
													return int32(0)
												} else {
													v117 = *(*int32)(unsafe.Add(mBase, uint32(v9)+96))
													F_ReleaseCatCache(m, v13)
													mBase = m.M
													v119 = m.ExcPending
													if v119 != 0 {
														return int32(0)
													} else {
														v122 = v117
														m.G0 = v9 + int32(112)
														return v122
													}
												}
											}
										}
									}
								} else {
									F_appendStringInfoString(m, v9+int32(96), int32(_a_F_format_operator_extended_4))
									mBase = m.M
									v73 = m.ExcPending
									if v73 != 0 {
										return int32(0)
									} else {
										v98 = *(*int32)(unsafe.Add(mBase, uint32(v19)+84))
										if v98 != 0 {
											if v27 != 0 {
												v99 = F_format_type_be_qualified(m, v98)
												mBase = m.M
												v100 = m.ExcPending
												if v100 != 0 {
													return int32(0)
												} else {
													v103 = v99
													*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v103
													F_appendStringInfo(m, v9+int32(96), int32(_a_F_format_operator_extended_2), v9+int32(16))
													mBase = m.M
													v111 = m.ExcPending
													if v111 != 0 {
														return int32(0)
													} else {
														v117 = *(*int32)(unsafe.Add(mBase, uint32(v9)+96))
														F_ReleaseCatCache(m, v13)
														mBase = m.M
														v119 = m.ExcPending
														if v119 != 0 {
															return int32(0)
														} else {
															v122 = v117
															m.G0 = v9 + int32(112)
															return v122
														}
													}
												}
											} else {
												v101 = F_format_type_be(m, v98)
												mBase = m.M
												v102 = m.ExcPending
												if v102 != 0 {
													return int32(0)
												} else {
													v103 = v101
													*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v103
													F_appendStringInfo(m, v9+int32(96), int32(_a_F_format_operator_extended_2), v9+int32(16))
													mBase = m.M
													v111 = m.ExcPending
													if v111 != 0 {
														return int32(0)
													} else {
														v117 = *(*int32)(unsafe.Add(mBase, uint32(v9)+96))
														F_ReleaseCatCache(m, v13)
														mBase = m.M
														v119 = m.ExcPending
														if v119 != 0 {
															return int32(0)
														} else {
															v122 = v117
															m.G0 = v9 + int32(112)
															return v122
														}
													}
												}
											}
										} else {
											F_appendStringInfoString(m, v9+int32(96), int32(_a_F_format_operator_extended_3))
											mBase = m.M
											v116 = m.ExcPending
											if v116 != 0 {
												return int32(0)
											} else {
												v117 = *(*int32)(unsafe.Add(mBase, uint32(v9)+96))
												F_ReleaseCatCache(m, v13)
												mBase = m.M
												v119 = m.ExcPending
												if v119 != 0 {
													return int32(0)
												} else {
													v122 = v117
													m.G0 = v9 + int32(112)
													return v122
												}
											}
										}
									}
								}
							}
						} else {
							v33 = *(*int32)(unsafe.Add(mBase, uint32(v19)+68))
							v34 = F_get_namespace_name(m, v33)
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return int32(0)
							} else {
								v36 = F_quote_identifier(m, v34)
								mBase = m.M
								v37 = m.ExcPending
								if v37 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v9)+80)) = v36
									v40 = v9 + int32(96)
									F_appendStringInfo(m, v40, int32(_a_F_format_operator_extended_5), v9+int32(80))
									mBase = m.M
									v45 = m.ExcPending
									if v45 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v9)+64)) = v21
										F_appendStringInfo(m, v40, int32(_a_F_format_operator_extended_0), v9-int32(-64))
										mBase = m.M
										v51 = m.ExcPending
										if v51 != 0 {
											return int32(0)
										} else {
											v52 = *(*int32)(unsafe.Add(mBase, uint32(v19)+80))
											if v52 == int32(0) {
												F_appendStringInfoString(m, v9+int32(96), int32(_a_F_format_operator_extended_4))
												mBase = m.M
												v73 = m.ExcPending
												if v73 != 0 {
													return int32(0)
												} else {
													v98 = *(*int32)(unsafe.Add(mBase, uint32(v19)+84))
													if v98 != 0 {
														if v27 != 0 {
															v99 = F_format_type_be_qualified(m, v98)
															mBase = m.M
															v100 = m.ExcPending
															if v100 != 0 {
																return int32(0)
															} else {
																v103 = v99
																*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v103
																F_appendStringInfo(m, v9+int32(96), int32(_a_F_format_operator_extended_2), v9+int32(16))
																mBase = m.M
																v111 = m.ExcPending
																if v111 != 0 {
																	return int32(0)
																} else {
																	v117 = *(*int32)(unsafe.Add(mBase, uint32(v9)+96))
																	F_ReleaseCatCache(m, v13)
																	mBase = m.M
																	v119 = m.ExcPending
																	if v119 != 0 {
																		return int32(0)
																	} else {
																		v122 = v117
																		m.G0 = v9 + int32(112)
																		return v122
																	}
																}
															}
														} else {
															v101 = F_format_type_be(m, v98)
															mBase = m.M
															v102 = m.ExcPending
															if v102 != 0 {
																return int32(0)
															} else {
																v103 = v101
																*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v103
																F_appendStringInfo(m, v9+int32(96), int32(_a_F_format_operator_extended_2), v9+int32(16))
																mBase = m.M
																v111 = m.ExcPending
																if v111 != 0 {
																	return int32(0)
																} else {
																	v117 = *(*int32)(unsafe.Add(mBase, uint32(v9)+96))
																	F_ReleaseCatCache(m, v13)
																	mBase = m.M
																	v119 = m.ExcPending
																	if v119 != 0 {
																		return int32(0)
																	} else {
																		v122 = v117
																		m.G0 = v9 + int32(112)
																		return v122
																	}
																}
															}
														}
													} else {
														F_appendStringInfoString(m, v9+int32(96), int32(_a_F_format_operator_extended_3))
														mBase = m.M
														v116 = m.ExcPending
														if v116 != 0 {
															return int32(0)
														} else {
															v117 = *(*int32)(unsafe.Add(mBase, uint32(v9)+96))
															F_ReleaseCatCache(m, v13)
															mBase = m.M
															v119 = m.ExcPending
															if v119 != 0 {
																return int32(0)
															} else {
																v122 = v117
																m.G0 = v9 + int32(112)
																return v122
															}
														}
													}
												}
											} else {
												if v27 == int32(0) {
													v84 = v52
													v85 = F_format_type_be(m, v84)
													mBase = m.M
													v86 = m.ExcPending
													if v86 != 0 {
														return int32(0)
													} else {
														v88 = v85
														*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v88
														F_appendStringInfo(m, v9+int32(96), int32(_a_F_format_operator_extended_1), v9+int32(32))
														mBase = m.M
														v96 = m.ExcPending
														if v96 != 0 {
															return int32(0)
														} else {
															v98 = *(*int32)(unsafe.Add(mBase, uint32(v19)+84))
															if v98 != 0 {
																if v27 != 0 {
																	v99 = F_format_type_be_qualified(m, v98)
																	mBase = m.M
																	v100 = m.ExcPending
																	if v100 != 0 {
																		return int32(0)
																	} else {
																		v103 = v99
																		*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v103
																		F_appendStringInfo(m, v9+int32(96), int32(_a_F_format_operator_extended_2), v9+int32(16))
																		mBase = m.M
																		v111 = m.ExcPending
																		if v111 != 0 {
																			return int32(0)
																		} else {
																			v117 = *(*int32)(unsafe.Add(mBase, uint32(v9)+96))
																			F_ReleaseCatCache(m, v13)
																			mBase = m.M
																			v119 = m.ExcPending
																			if v119 != 0 {
																				return int32(0)
																			} else {
																				v122 = v117
																				m.G0 = v9 + int32(112)
																				return v122
																			}
																		}
																	}
																} else {
																	v101 = F_format_type_be(m, v98)
																	mBase = m.M
																	v102 = m.ExcPending
																	if v102 != 0 {
																		return int32(0)
																	} else {
																		v103 = v101
																		*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v103
																		F_appendStringInfo(m, v9+int32(96), int32(_a_F_format_operator_extended_2), v9+int32(16))
																		mBase = m.M
																		v111 = m.ExcPending
																		if v111 != 0 {
																			return int32(0)
																		} else {
																			v117 = *(*int32)(unsafe.Add(mBase, uint32(v9)+96))
																			F_ReleaseCatCache(m, v13)
																			mBase = m.M
																			v119 = m.ExcPending
																			if v119 != 0 {
																				return int32(0)
																			} else {
																				v122 = v117
																				m.G0 = v9 + int32(112)
																				return v122
																			}
																		}
																	}
																}
															} else {
																F_appendStringInfoString(m, v9+int32(96), int32(_a_F_format_operator_extended_3))
																mBase = m.M
																v116 = m.ExcPending
																if v116 != 0 {
																	return int32(0)
																} else {
																	v117 = *(*int32)(unsafe.Add(mBase, uint32(v9)+96))
																	F_ReleaseCatCache(m, v13)
																	mBase = m.M
																	v119 = m.ExcPending
																	if v119 != 0 {
																		return int32(0)
																	} else {
																		v122 = v117
																		m.G0 = v9 + int32(112)
																		return v122
																	}
																}
															}
														}
													}
												} else {
													v57 = F_format_type_be_qualified(m, v52)
													mBase = m.M
													v58 = m.ExcPending
													if v58 != 0 {
														return int32(0)
													} else {
														v88 = v57
														*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v88
														F_appendStringInfo(m, v9+int32(96), int32(_a_F_format_operator_extended_1), v9+int32(32))
														mBase = m.M
														v96 = m.ExcPending
														if v96 != 0 {
															return int32(0)
														} else {
															v98 = *(*int32)(unsafe.Add(mBase, uint32(v19)+84))
															if v98 != 0 {
																if v27 != 0 {
																	v99 = F_format_type_be_qualified(m, v98)
																	mBase = m.M
																	v100 = m.ExcPending
																	if v100 != 0 {
																		return int32(0)
																	} else {
																		v103 = v99
																		*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v103
																		F_appendStringInfo(m, v9+int32(96), int32(_a_F_format_operator_extended_2), v9+int32(16))
																		mBase = m.M
																		v111 = m.ExcPending
																		if v111 != 0 {
																			return int32(0)
																		} else {
																			v117 = *(*int32)(unsafe.Add(mBase, uint32(v9)+96))
																			F_ReleaseCatCache(m, v13)
																			mBase = m.M
																			v119 = m.ExcPending
																			if v119 != 0 {
																				return int32(0)
																			} else {
																				v122 = v117
																				m.G0 = v9 + int32(112)
																				return v122
																			}
																		}
																	}
																} else {
																	v101 = F_format_type_be(m, v98)
																	mBase = m.M
																	v102 = m.ExcPending
																	if v102 != 0 {
																		return int32(0)
																	} else {
																		v103 = v101
																		*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v103
																		F_appendStringInfo(m, v9+int32(96), int32(_a_F_format_operator_extended_2), v9+int32(16))
																		mBase = m.M
																		v111 = m.ExcPending
																		if v111 != 0 {
																			return int32(0)
																		} else {
																			v117 = *(*int32)(unsafe.Add(mBase, uint32(v9)+96))
																			F_ReleaseCatCache(m, v13)
																			mBase = m.M
																			v119 = m.ExcPending
																			if v119 != 0 {
																				return int32(0)
																			} else {
																				v122 = v117
																				m.G0 = v9 + int32(112)
																				return v122
																			}
																		}
																	}
																}
															} else {
																F_appendStringInfoString(m, v9+int32(96), int32(_a_F_format_operator_extended_3))
																mBase = m.M
																v116 = m.ExcPending
																if v116 != 0 {
																	return int32(0)
																} else {
																	v117 = *(*int32)(unsafe.Add(mBase, uint32(v9)+96))
																	F_ReleaseCatCache(m, v13)
																	mBase = m.M
																	v119 = m.ExcPending
																	if v119 != 0 {
																		return int32(0)
																	} else {
																		v122 = v117
																		m.G0 = v9 + int32(112)
																		return v122
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
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v19)+68))
					v34 = F_get_namespace_name(m, v33)
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						v36 = F_quote_identifier(m, v34)
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9)+80)) = v36
							v40 = v9 + int32(96)
							F_appendStringInfo(m, v40, int32(_a_F_format_operator_extended_5), v9+int32(80))
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9)+64)) = v21
								F_appendStringInfo(m, v40, int32(_a_F_format_operator_extended_0), v9-int32(-64))
								mBase = m.M
								v51 = m.ExcPending
								if v51 != 0 {
									return int32(0)
								} else {
									v52 = *(*int32)(unsafe.Add(mBase, uint32(v19)+80))
									if v52 == int32(0) {
										F_appendStringInfoString(m, v9+int32(96), int32(_a_F_format_operator_extended_4))
										mBase = m.M
										v73 = m.ExcPending
										if v73 != 0 {
											return int32(0)
										} else {
											v98 = *(*int32)(unsafe.Add(mBase, uint32(v19)+84))
											if v98 != 0 {
												if v27 != 0 {
													v99 = F_format_type_be_qualified(m, v98)
													mBase = m.M
													v100 = m.ExcPending
													if v100 != 0 {
														return int32(0)
													} else {
														v103 = v99
														*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v103
														F_appendStringInfo(m, v9+int32(96), int32(_a_F_format_operator_extended_2), v9+int32(16))
														mBase = m.M
														v111 = m.ExcPending
														if v111 != 0 {
															return int32(0)
														} else {
															v117 = *(*int32)(unsafe.Add(mBase, uint32(v9)+96))
															F_ReleaseCatCache(m, v13)
															mBase = m.M
															v119 = m.ExcPending
															if v119 != 0 {
																return int32(0)
															} else {
																v122 = v117
																m.G0 = v9 + int32(112)
																return v122
															}
														}
													}
												} else {
													v101 = F_format_type_be(m, v98)
													mBase = m.M
													v102 = m.ExcPending
													if v102 != 0 {
														return int32(0)
													} else {
														v103 = v101
														*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v103
														F_appendStringInfo(m, v9+int32(96), int32(_a_F_format_operator_extended_2), v9+int32(16))
														mBase = m.M
														v111 = m.ExcPending
														if v111 != 0 {
															return int32(0)
														} else {
															v117 = *(*int32)(unsafe.Add(mBase, uint32(v9)+96))
															F_ReleaseCatCache(m, v13)
															mBase = m.M
															v119 = m.ExcPending
															if v119 != 0 {
																return int32(0)
															} else {
																v122 = v117
																m.G0 = v9 + int32(112)
																return v122
															}
														}
													}
												}
											} else {
												F_appendStringInfoString(m, v9+int32(96), int32(_a_F_format_operator_extended_3))
												mBase = m.M
												v116 = m.ExcPending
												if v116 != 0 {
													return int32(0)
												} else {
													v117 = *(*int32)(unsafe.Add(mBase, uint32(v9)+96))
													F_ReleaseCatCache(m, v13)
													mBase = m.M
													v119 = m.ExcPending
													if v119 != 0 {
														return int32(0)
													} else {
														v122 = v117
														m.G0 = v9 + int32(112)
														return v122
													}
												}
											}
										}
									} else {
										if v27 == int32(0) {
											v84 = v52
											v85 = F_format_type_be(m, v84)
											mBase = m.M
											v86 = m.ExcPending
											if v86 != 0 {
												return int32(0)
											} else {
												v88 = v85
												*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v88
												F_appendStringInfo(m, v9+int32(96), int32(_a_F_format_operator_extended_1), v9+int32(32))
												mBase = m.M
												v96 = m.ExcPending
												if v96 != 0 {
													return int32(0)
												} else {
													v98 = *(*int32)(unsafe.Add(mBase, uint32(v19)+84))
													if v98 != 0 {
														if v27 != 0 {
															v99 = F_format_type_be_qualified(m, v98)
															mBase = m.M
															v100 = m.ExcPending
															if v100 != 0 {
																return int32(0)
															} else {
																v103 = v99
																*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v103
																F_appendStringInfo(m, v9+int32(96), int32(_a_F_format_operator_extended_2), v9+int32(16))
																mBase = m.M
																v111 = m.ExcPending
																if v111 != 0 {
																	return int32(0)
																} else {
																	v117 = *(*int32)(unsafe.Add(mBase, uint32(v9)+96))
																	F_ReleaseCatCache(m, v13)
																	mBase = m.M
																	v119 = m.ExcPending
																	if v119 != 0 {
																		return int32(0)
																	} else {
																		v122 = v117
																		m.G0 = v9 + int32(112)
																		return v122
																	}
																}
															}
														} else {
															v101 = F_format_type_be(m, v98)
															mBase = m.M
															v102 = m.ExcPending
															if v102 != 0 {
																return int32(0)
															} else {
																v103 = v101
																*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v103
																F_appendStringInfo(m, v9+int32(96), int32(_a_F_format_operator_extended_2), v9+int32(16))
																mBase = m.M
																v111 = m.ExcPending
																if v111 != 0 {
																	return int32(0)
																} else {
																	v117 = *(*int32)(unsafe.Add(mBase, uint32(v9)+96))
																	F_ReleaseCatCache(m, v13)
																	mBase = m.M
																	v119 = m.ExcPending
																	if v119 != 0 {
																		return int32(0)
																	} else {
																		v122 = v117
																		m.G0 = v9 + int32(112)
																		return v122
																	}
																}
															}
														}
													} else {
														F_appendStringInfoString(m, v9+int32(96), int32(_a_F_format_operator_extended_3))
														mBase = m.M
														v116 = m.ExcPending
														if v116 != 0 {
															return int32(0)
														} else {
															v117 = *(*int32)(unsafe.Add(mBase, uint32(v9)+96))
															F_ReleaseCatCache(m, v13)
															mBase = m.M
															v119 = m.ExcPending
															if v119 != 0 {
																return int32(0)
															} else {
																v122 = v117
																m.G0 = v9 + int32(112)
																return v122
															}
														}
													}
												}
											}
										} else {
											v57 = F_format_type_be_qualified(m, v52)
											mBase = m.M
											v58 = m.ExcPending
											if v58 != 0 {
												return int32(0)
											} else {
												v88 = v57
												*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v88
												F_appendStringInfo(m, v9+int32(96), int32(_a_F_format_operator_extended_1), v9+int32(32))
												mBase = m.M
												v96 = m.ExcPending
												if v96 != 0 {
													return int32(0)
												} else {
													v98 = *(*int32)(unsafe.Add(mBase, uint32(v19)+84))
													if v98 != 0 {
														if v27 != 0 {
															v99 = F_format_type_be_qualified(m, v98)
															mBase = m.M
															v100 = m.ExcPending
															if v100 != 0 {
																return int32(0)
															} else {
																v103 = v99
																*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v103
																F_appendStringInfo(m, v9+int32(96), int32(_a_F_format_operator_extended_2), v9+int32(16))
																mBase = m.M
																v111 = m.ExcPending
																if v111 != 0 {
																	return int32(0)
																} else {
																	v117 = *(*int32)(unsafe.Add(mBase, uint32(v9)+96))
																	F_ReleaseCatCache(m, v13)
																	mBase = m.M
																	v119 = m.ExcPending
																	if v119 != 0 {
																		return int32(0)
																	} else {
																		v122 = v117
																		m.G0 = v9 + int32(112)
																		return v122
																	}
																}
															}
														} else {
															v101 = F_format_type_be(m, v98)
															mBase = m.M
															v102 = m.ExcPending
															if v102 != 0 {
																return int32(0)
															} else {
																v103 = v101
																*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v103
																F_appendStringInfo(m, v9+int32(96), int32(_a_F_format_operator_extended_2), v9+int32(16))
																mBase = m.M
																v111 = m.ExcPending
																if v111 != 0 {
																	return int32(0)
																} else {
																	v117 = *(*int32)(unsafe.Add(mBase, uint32(v9)+96))
																	F_ReleaseCatCache(m, v13)
																	mBase = m.M
																	v119 = m.ExcPending
																	if v119 != 0 {
																		return int32(0)
																	} else {
																		v122 = v117
																		m.G0 = v9 + int32(112)
																		return v122
																	}
																}
															}
														}
													} else {
														F_appendStringInfoString(m, v9+int32(96), int32(_a_F_format_operator_extended_3))
														mBase = m.M
														v116 = m.ExcPending
														if v116 != 0 {
															return int32(0)
														} else {
															v117 = *(*int32)(unsafe.Add(mBase, uint32(v9)+96))
															F_ReleaseCatCache(m, v13)
															mBase = m.M
															v119 = m.ExcPending
															if v119 != 0 {
																return int32(0)
															} else {
																v122 = v117
																m.G0 = v9 + int32(112)
																return v122
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
			if l1&int32(1) != 0 {
				v122 = int32(0)
				m.G0 = v9 + int32(112)
				return v122
			} else {
				v77 = F_palloc(m, int32(64))
				mBase = m.M
				v78 = m.ExcPending
				if v78 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
					v82 = F_pg_snprintf(m, v77, int32(64), int32(_a_F_format_operator_extended_6), v9)
					mBase = m.M
					v83 = m.ExcPending
					if v83 != 0 {
						return int32(0)
					} else {
						v122 = v77
						m.G0 = v9 + int32(112)
						return v122
					}
				}
			}
		}
	}
}
func F_generate_operator_name(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
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
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
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
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	v9 = m.G0
	v11 = v9 - int32(80)
	m.G0 = v11
	F_initStringInfo(m, v11-int32(-64))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v21 = F_SearchSysCache1(m, int32(40), base.I64_extend_i32_u(l0))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			if v21 != 0 {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
				v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+22)))
				v25 = v23 + v24
				v27 = v25 + int32(4)
				v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+76)))
				switch v28 - int32(98) {
				case 0:
					v61 = F_makeString(m, v27)
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v61
						*(*int32)(unsafe.Add(mBase, uint32(v11)+60)) = v61
						v69 = F_list_make1_impl(m, int32(1), v11+int32(48))
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
							return int32(0)
						} else {
							v73 = F_oper(m, int32(0), v69, l1, l2, int32(1), int32(-1))
							mBase = m.M
							v74 = m.ExcPending
							if v74 != 0 {
								return int32(0)
							} else {
								v77 = v73
								if v77 != 0 {
									v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+16))
									v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+22)))
									v81 = *(*int32)(unsafe.Add(mBase, uint32(v78+v79)))
									if l0 == v81 {
										F_appendStringInfoString(m, v11-int32(-64), v27)
										mBase = m.M
										v120 = m.ExcPending
										if v120 != 0 {
											return int32(0)
										} else {
											F_ReleaseCatCache(m, v77)
											mBase = m.M
											v124 = m.ExcPending
											if v124 != 0 {
												return int32(0)
											} else {
												F_ReleaseCatCache(m, v21)
												mBase = m.M
												v128 = m.ExcPending
												if v128 != 0 {
													return int32(0)
												} else {
													v129 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
													m.G0 = v11 + int32(80)
													return v129
												}
											}
										}
									} else {
										v84 = *(*int32)(unsafe.Add(mBase, uint32(v25)+68))
										v85 = F_get_namespace_name_or_temp(m, v84)
										mBase = m.M
										v86 = m.ExcPending
										if v86 != 0 {
											return int32(0)
										} else {
											v87 = F_quote_identifier(m, v85)
											mBase = m.M
											v88 = m.ExcPending
											if v88 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v87
												v91 = v11 - int32(-64)
												F_appendStringInfo(m, v91, int32(_a_F_generate_operator_name_0), v11+int32(32))
												mBase = m.M
												v96 = m.ExcPending
												if v96 != 0 {
													return int32(0)
												} else {
													F_appendStringInfoString(m, v91, v27)
													mBase = m.M
													v98 = m.ExcPending
													if v98 != 0 {
														return int32(0)
													} else {
														if v85 != 0 {
															F_appendStringInfoChar(m, v91, int32(41))
															mBase = m.M
															v101 = m.ExcPending
															if v101 != 0 {
																return int32(0)
															} else {
																if v77 == int32(0) {
																	F_ReleaseCatCache(m, v21)
																	mBase = m.M
																	v128 = m.ExcPending
																	if v128 != 0 {
																		return int32(0)
																	} else {
																		v129 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
																		m.G0 = v11 + int32(80)
																		return v129
																	}
																} else {
																	F_ReleaseCatCache(m, v77)
																	mBase = m.M
																	v124 = m.ExcPending
																	if v124 != 0 {
																		return int32(0)
																	} else {
																		F_ReleaseCatCache(m, v21)
																		mBase = m.M
																		v128 = m.ExcPending
																		if v128 != 0 {
																			return int32(0)
																		} else {
																			v129 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
																			m.G0 = v11 + int32(80)
																			return v129
																		}
																	}
																}
															}
														} else {
															if v77 == int32(0) {
																F_ReleaseCatCache(m, v21)
																mBase = m.M
																v128 = m.ExcPending
																if v128 != 0 {
																	return int32(0)
																} else {
																	v129 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
																	m.G0 = v11 + int32(80)
																	return v129
																}
															} else {
																F_ReleaseCatCache(m, v77)
																mBase = m.M
																v124 = m.ExcPending
																if v124 != 0 {
																	return int32(0)
																} else {
																	F_ReleaseCatCache(m, v21)
																	mBase = m.M
																	v128 = m.ExcPending
																	if v128 != 0 {
																		return int32(0)
																	} else {
																		v129 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
																		m.G0 = v11 + int32(80)
																		return v129
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
									v84 = *(*int32)(unsafe.Add(mBase, uint32(v25)+68))
									v85 = F_get_namespace_name_or_temp(m, v84)
									mBase = m.M
									v86 = m.ExcPending
									if v86 != 0 {
										return int32(0)
									} else {
										v87 = F_quote_identifier(m, v85)
										mBase = m.M
										v88 = m.ExcPending
										if v88 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v87
											v91 = v11 - int32(-64)
											F_appendStringInfo(m, v91, int32(_a_F_generate_operator_name_0), v11+int32(32))
											mBase = m.M
											v96 = m.ExcPending
											if v96 != 0 {
												return int32(0)
											} else {
												F_appendStringInfoString(m, v91, v27)
												mBase = m.M
												v98 = m.ExcPending
												if v98 != 0 {
													return int32(0)
												} else {
													if v85 != 0 {
														F_appendStringInfoChar(m, v91, int32(41))
														mBase = m.M
														v101 = m.ExcPending
														if v101 != 0 {
															return int32(0)
														} else {
															if v77 == int32(0) {
																F_ReleaseCatCache(m, v21)
																mBase = m.M
																v128 = m.ExcPending
																if v128 != 0 {
																	return int32(0)
																} else {
																	v129 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
																	m.G0 = v11 + int32(80)
																	return v129
																}
															} else {
																F_ReleaseCatCache(m, v77)
																mBase = m.M
																v124 = m.ExcPending
																if v124 != 0 {
																	return int32(0)
																} else {
																	F_ReleaseCatCache(m, v21)
																	mBase = m.M
																	v128 = m.ExcPending
																	if v128 != 0 {
																		return int32(0)
																	} else {
																		v129 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
																		m.G0 = v11 + int32(80)
																		return v129
																	}
																}
															}
														}
													} else {
														if v77 == int32(0) {
															F_ReleaseCatCache(m, v21)
															mBase = m.M
															v128 = m.ExcPending
															if v128 != 0 {
																return int32(0)
															} else {
																v129 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
																m.G0 = v11 + int32(80)
																return v129
															}
														} else {
															F_ReleaseCatCache(m, v77)
															mBase = m.M
															v124 = m.ExcPending
															if v124 != 0 {
																return int32(0)
															} else {
																F_ReleaseCatCache(m, v21)
																mBase = m.M
																v128 = m.ExcPending
																if v128 != 0 {
																	return int32(0)
																} else {
																	v129 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
																	m.G0 = v11 + int32(80)
																	return v129
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
				default:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return int32(0)
					} else {
						v49 = int32(*(*int8)(unsafe.Add(mBase, uint32(v25)+76)))
						*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v49
						F_errmsg_internal(m, int32(_a_F_generate_operator_name_1), v11+int32(16))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_generate_operator_name_2), int32(_a_F_generate_operator_name_3), int32(_a_F_generate_operator_name_4))
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				case 10:
					v31 = F_makeString(m, v27)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = v31
						*(*int32)(unsafe.Add(mBase, uint32(v11)+56)) = v31
						v39 = F_list_make1_impl(m, int32(1), v11+int32(52))
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int32(0)
						} else {
							v43 = F_left_oper(m, int32(0), v39, l2, int32(1), int32(-1))
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return int32(0)
							} else {
								v77 = v43
								if v77 != 0 {
									v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+16))
									v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+22)))
									v81 = *(*int32)(unsafe.Add(mBase, uint32(v78+v79)))
									if l0 == v81 {
										F_appendStringInfoString(m, v11-int32(-64), v27)
										mBase = m.M
										v120 = m.ExcPending
										if v120 != 0 {
											return int32(0)
										} else {
											F_ReleaseCatCache(m, v77)
											mBase = m.M
											v124 = m.ExcPending
											if v124 != 0 {
												return int32(0)
											} else {
												F_ReleaseCatCache(m, v21)
												mBase = m.M
												v128 = m.ExcPending
												if v128 != 0 {
													return int32(0)
												} else {
													v129 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
													m.G0 = v11 + int32(80)
													return v129
												}
											}
										}
									} else {
										v84 = *(*int32)(unsafe.Add(mBase, uint32(v25)+68))
										v85 = F_get_namespace_name_or_temp(m, v84)
										mBase = m.M
										v86 = m.ExcPending
										if v86 != 0 {
											return int32(0)
										} else {
											v87 = F_quote_identifier(m, v85)
											mBase = m.M
											v88 = m.ExcPending
											if v88 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v87
												v91 = v11 - int32(-64)
												F_appendStringInfo(m, v91, int32(_a_F_generate_operator_name_0), v11+int32(32))
												mBase = m.M
												v96 = m.ExcPending
												if v96 != 0 {
													return int32(0)
												} else {
													F_appendStringInfoString(m, v91, v27)
													mBase = m.M
													v98 = m.ExcPending
													if v98 != 0 {
														return int32(0)
													} else {
														if v85 != 0 {
															F_appendStringInfoChar(m, v91, int32(41))
															mBase = m.M
															v101 = m.ExcPending
															if v101 != 0 {
																return int32(0)
															} else {
																if v77 == int32(0) {
																	F_ReleaseCatCache(m, v21)
																	mBase = m.M
																	v128 = m.ExcPending
																	if v128 != 0 {
																		return int32(0)
																	} else {
																		v129 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
																		m.G0 = v11 + int32(80)
																		return v129
																	}
																} else {
																	F_ReleaseCatCache(m, v77)
																	mBase = m.M
																	v124 = m.ExcPending
																	if v124 != 0 {
																		return int32(0)
																	} else {
																		F_ReleaseCatCache(m, v21)
																		mBase = m.M
																		v128 = m.ExcPending
																		if v128 != 0 {
																			return int32(0)
																		} else {
																			v129 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
																			m.G0 = v11 + int32(80)
																			return v129
																		}
																	}
																}
															}
														} else {
															if v77 == int32(0) {
																F_ReleaseCatCache(m, v21)
																mBase = m.M
																v128 = m.ExcPending
																if v128 != 0 {
																	return int32(0)
																} else {
																	v129 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
																	m.G0 = v11 + int32(80)
																	return v129
																}
															} else {
																F_ReleaseCatCache(m, v77)
																mBase = m.M
																v124 = m.ExcPending
																if v124 != 0 {
																	return int32(0)
																} else {
																	F_ReleaseCatCache(m, v21)
																	mBase = m.M
																	v128 = m.ExcPending
																	if v128 != 0 {
																		return int32(0)
																	} else {
																		v129 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
																		m.G0 = v11 + int32(80)
																		return v129
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
									v84 = *(*int32)(unsafe.Add(mBase, uint32(v25)+68))
									v85 = F_get_namespace_name_or_temp(m, v84)
									mBase = m.M
									v86 = m.ExcPending
									if v86 != 0 {
										return int32(0)
									} else {
										v87 = F_quote_identifier(m, v85)
										mBase = m.M
										v88 = m.ExcPending
										if v88 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v87
											v91 = v11 - int32(-64)
											F_appendStringInfo(m, v91, int32(_a_F_generate_operator_name_0), v11+int32(32))
											mBase = m.M
											v96 = m.ExcPending
											if v96 != 0 {
												return int32(0)
											} else {
												F_appendStringInfoString(m, v91, v27)
												mBase = m.M
												v98 = m.ExcPending
												if v98 != 0 {
													return int32(0)
												} else {
													if v85 != 0 {
														F_appendStringInfoChar(m, v91, int32(41))
														mBase = m.M
														v101 = m.ExcPending
														if v101 != 0 {
															return int32(0)
														} else {
															if v77 == int32(0) {
																F_ReleaseCatCache(m, v21)
																mBase = m.M
																v128 = m.ExcPending
																if v128 != 0 {
																	return int32(0)
																} else {
																	v129 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
																	m.G0 = v11 + int32(80)
																	return v129
																}
															} else {
																F_ReleaseCatCache(m, v77)
																mBase = m.M
																v124 = m.ExcPending
																if v124 != 0 {
																	return int32(0)
																} else {
																	F_ReleaseCatCache(m, v21)
																	mBase = m.M
																	v128 = m.ExcPending
																	if v128 != 0 {
																		return int32(0)
																	} else {
																		v129 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
																		m.G0 = v11 + int32(80)
																		return v129
																	}
																}
															}
														}
													} else {
														if v77 == int32(0) {
															F_ReleaseCatCache(m, v21)
															mBase = m.M
															v128 = m.ExcPending
															if v128 != 0 {
																return int32(0)
															} else {
																v129 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
																m.G0 = v11 + int32(80)
																return v129
															}
														} else {
															F_ReleaseCatCache(m, v77)
															mBase = m.M
															v124 = m.ExcPending
															if v124 != 0 {
																return int32(0)
															} else {
																F_ReleaseCatCache(m, v21)
																mBase = m.M
																v128 = m.ExcPending
																if v128 != 0 {
																	return int32(0)
																} else {
																	v129 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
																	m.G0 = v11 + int32(80)
																	return v129
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
				v107 = m.ExcPending
				if v107 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v11))) = l0
					F_errmsg_internal(m, int32(_a_F_generate_operator_name_5), v11)
					mBase = m.M
					v111 = m.ExcPending
					if v111 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_generate_operator_name_2), int32(_a_F_generate_operator_name_6), int32(_a_F_generate_operator_name_4))
						mBase = m.M
						v116 = m.ExcPending
						if v116 != 0 {
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
func F_pushOperator(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	v2 = l1
	v6 = F_palloc0(m, int32(8))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		*(*uint8)(unsafe.Add(mBase, uint32(v6)+1)) = uint8(v2)
		v9 = int32(2)
		*(*uint8)(unsafe.Add(mBase, uint32(v6))) = uint8(v9)
		if v2 == int32(4) {
			v14 = l2
		} else {
			v14 = int32(0)
		}
		*(*uint16)(unsafe.Add(mBase, uint32(v6)+2)) = uint16(v14)
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v17 = F_lcons(m, v6, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v17
			return
		}
	}
}
