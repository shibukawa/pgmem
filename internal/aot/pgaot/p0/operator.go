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
					F_errmsg(m, int32(130092), int32(0))
					v26 = m.ExcPending
					if v26 != 0 {
						return
					} else {
						F_errfinish(m, int32(490103), int32(572), int32(149738))
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
						F_errmsg(m, int32(9989), int32(0))
						v42 = m.ExcPending
						if v42 != 0 {
							return
						} else {
							F_errfinish(m, int32(490103), int32(576), int32(149738))
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
							F_errmsg(m, int32(273379), int32(0))
							v58 = m.ExcPending
							if v58 != 0 {
								return
							} else {
								F_errfinish(m, int32(490103), int32(580), int32(149738))
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
								F_errmsg(m, int32(320176), int32(0))
								v74 = m.ExcPending
								if v74 != 0 {
									return
								} else {
									F_errfinish(m, int32(490103), int32(584), int32(149738))
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
										F_errmsg(m, int32(130328), int32(0))
										v90 = m.ExcPending
										if v90 != 0 {
											return
										} else {
											F_errfinish(m, int32(490103), int32(593), int32(149738))
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
											F_errmsg(m, int32(9933), int32(0))
											v106 = m.ExcPending
											if v106 != 0 {
												return
											} else {
												F_errfinish(m, int32(490103), int32(597), int32(149738))
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
												F_errmsg(m, int32(10037), int32(0))
												v122 = m.ExcPending
												if v122 != 0 {
													return
												} else {
													F_errfinish(m, int32(490103), int32(601), int32(149738))
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
													F_errmsg(m, int32(273416), int32(0))
													v138 = m.ExcPending
													if v138 != 0 {
														return
													} else {
														F_errfinish(m, int32(490103), int32(605), int32(149738))
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
														F_errmsg(m, int32(320207), int32(0))
														v154 = m.ExcPending
														if v154 != 0 {
															return
														} else {
															F_errfinish(m, int32(490103), int32(609), int32(149738))
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
						F_errmsg(m, int32(130328), int32(0))
						v90 = m.ExcPending
						if v90 != 0 {
							return
						} else {
							F_errfinish(m, int32(490103), int32(593), int32(149738))
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
							F_errmsg(m, int32(9933), int32(0))
							v106 = m.ExcPending
							if v106 != 0 {
								return
							} else {
								F_errfinish(m, int32(490103), int32(597), int32(149738))
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
								F_errmsg(m, int32(10037), int32(0))
								v122 = m.ExcPending
								if v122 != 0 {
									return
								} else {
									F_errfinish(m, int32(490103), int32(601), int32(149738))
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
									F_errmsg(m, int32(273416), int32(0))
									v138 = m.ExcPending
									if v138 != 0 {
										return
									} else {
										F_errfinish(m, int32(490103), int32(605), int32(149738))
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
										F_errmsg(m, int32(320207), int32(0))
										v154 = m.ExcPending
										if v154 != 0 {
											return
										} else {
											F_errfinish(m, int32(490103), int32(609), int32(149738))
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
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
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
	var v44 int32
	_ = v44
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
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
	var v89 int32
	_ = v89
	var v97 int32
	_ = v97
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
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	v7 = m.G0
	v9 = v7 - int32(112)
	m.G0 = v9
	v12 = F_SearchSysCache1(m, int32(40), l0)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		if v12 != 0 {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
			v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+22)))
			v18 = v16 + v17
			v20 = v18 + int32(4)
			F_initStringInfo(m, v9+int32(96))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				v26 = l1 & int32(2)
				if v26 == int32(0) {
					v30 = F_OperatorIsVisibleExt(m, l0, int32(0))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						if v30 != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v20
							F_appendStringInfo(m, v9+int32(96), int32(664818), v9+int32(48))
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return int32(0)
							} else {
								v68 = *(*int32)(unsafe.Add(mBase, uint32(v18)+80))
								if v68 != 0 {
									v85 = v68
									v86 = F_format_type_be(m, v85)
									mBase = m.M
									v87 = m.ExcPending
									if v87 != 0 {
										return int32(0)
									} else {
										v89 = v86
										*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v89
										F_appendStringInfo(m, v9+int32(96), int32(649293), v9+int32(32))
										mBase = m.M
										v97 = m.ExcPending
										if v97 != 0 {
											return int32(0)
										} else {
											v99 = *(*int32)(unsafe.Add(mBase, uint32(v18)+84))
											if v99 != 0 {
												if v26 != 0 {
													v100 = F_format_type_be_qualified(m, v99)
													mBase = m.M
													v101 = m.ExcPending
													if v101 != 0 {
														return int32(0)
													} else {
														v104 = v100
														*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v104
														F_appendStringInfo(m, v9+int32(96), int32(655025), v9+int32(16))
														mBase = m.M
														v112 = m.ExcPending
														if v112 != 0 {
															return int32(0)
														} else {
															v118 = *(*int32)(unsafe.Add(mBase, uint32(v9)+96))
															F_ReleaseCatCache(m, v12)
															mBase = m.M
															v120 = m.ExcPending
															if v120 != 0 {
																return int32(0)
															} else {
																v123 = v118
																m.G0 = v9 + int32(112)
																return v123
															}
														}
													}
												} else {
													v102 = F_format_type_be(m, v99)
													mBase = m.M
													v103 = m.ExcPending
													if v103 != 0 {
														return int32(0)
													} else {
														v104 = v102
														*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v104
														F_appendStringInfo(m, v9+int32(96), int32(655025), v9+int32(16))
														mBase = m.M
														v112 = m.ExcPending
														if v112 != 0 {
															return int32(0)
														} else {
															v118 = *(*int32)(unsafe.Add(mBase, uint32(v9)+96))
															F_ReleaseCatCache(m, v12)
															mBase = m.M
															v120 = m.ExcPending
															if v120 != 0 {
																return int32(0)
															} else {
																v123 = v118
																m.G0 = v9 + int32(112)
																return v123
															}
														}
													}
												}
											} else {
												F_appendStringInfoString(m, v9+int32(96), int32(661660))
												mBase = m.M
												v117 = m.ExcPending
												if v117 != 0 {
													return int32(0)
												} else {
													v118 = *(*int32)(unsafe.Add(mBase, uint32(v9)+96))
													F_ReleaseCatCache(m, v12)
													mBase = m.M
													v120 = m.ExcPending
													if v120 != 0 {
														return int32(0)
													} else {
														v123 = v118
														m.G0 = v9 + int32(112)
														return v123
													}
												}
											}
										}
									}
								} else {
									F_appendStringInfoString(m, v9+int32(96), int32(649297))
									mBase = m.M
									v74 = m.ExcPending
									if v74 != 0 {
										return int32(0)
									} else {
										v99 = *(*int32)(unsafe.Add(mBase, uint32(v18)+84))
										if v99 != 0 {
											if v26 != 0 {
												v100 = F_format_type_be_qualified(m, v99)
												mBase = m.M
												v101 = m.ExcPending
												if v101 != 0 {
													return int32(0)
												} else {
													v104 = v100
													*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v104
													F_appendStringInfo(m, v9+int32(96), int32(655025), v9+int32(16))
													mBase = m.M
													v112 = m.ExcPending
													if v112 != 0 {
														return int32(0)
													} else {
														v118 = *(*int32)(unsafe.Add(mBase, uint32(v9)+96))
														F_ReleaseCatCache(m, v12)
														mBase = m.M
														v120 = m.ExcPending
														if v120 != 0 {
															return int32(0)
														} else {
															v123 = v118
															m.G0 = v9 + int32(112)
															return v123
														}
													}
												}
											} else {
												v102 = F_format_type_be(m, v99)
												mBase = m.M
												v103 = m.ExcPending
												if v103 != 0 {
													return int32(0)
												} else {
													v104 = v102
													*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v104
													F_appendStringInfo(m, v9+int32(96), int32(655025), v9+int32(16))
													mBase = m.M
													v112 = m.ExcPending
													if v112 != 0 {
														return int32(0)
													} else {
														v118 = *(*int32)(unsafe.Add(mBase, uint32(v9)+96))
														F_ReleaseCatCache(m, v12)
														mBase = m.M
														v120 = m.ExcPending
														if v120 != 0 {
															return int32(0)
														} else {
															v123 = v118
															m.G0 = v9 + int32(112)
															return v123
														}
													}
												}
											}
										} else {
											F_appendStringInfoString(m, v9+int32(96), int32(661660))
											mBase = m.M
											v117 = m.ExcPending
											if v117 != 0 {
												return int32(0)
											} else {
												v118 = *(*int32)(unsafe.Add(mBase, uint32(v9)+96))
												F_ReleaseCatCache(m, v12)
												mBase = m.M
												v120 = m.ExcPending
												if v120 != 0 {
													return int32(0)
												} else {
													v123 = v118
													m.G0 = v9 + int32(112)
													return v123
												}
											}
										}
									}
								}
							}
						} else {
							v32 = *(*int32)(unsafe.Add(mBase, uint32(v18)+68))
							v33 = F_get_namespace_name(m, v32)
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return int32(0)
							} else {
								v35 = F_quote_identifier(m, v33)
								mBase = m.M
								v36 = m.ExcPending
								if v36 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v9)+80)) = v35
									F_appendStringInfo(m, v9+int32(96), int32(586069), v9+int32(80))
									mBase = m.M
									v44 = m.ExcPending
									if v44 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v9)+64)) = v20
										F_appendStringInfo(m, v9+int32(96), int32(664818), v9-int32(-64))
										mBase = m.M
										v52 = m.ExcPending
										if v52 != 0 {
											return int32(0)
										} else {
											v53 = *(*int32)(unsafe.Add(mBase, uint32(v18)+80))
											if v53 == int32(0) {
												F_appendStringInfoString(m, v9+int32(96), int32(649297))
												mBase = m.M
												v74 = m.ExcPending
												if v74 != 0 {
													return int32(0)
												} else {
													v99 = *(*int32)(unsafe.Add(mBase, uint32(v18)+84))
													if v99 != 0 {
														if v26 != 0 {
															v100 = F_format_type_be_qualified(m, v99)
															mBase = m.M
															v101 = m.ExcPending
															if v101 != 0 {
																return int32(0)
															} else {
																v104 = v100
																*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v104
																F_appendStringInfo(m, v9+int32(96), int32(655025), v9+int32(16))
																mBase = m.M
																v112 = m.ExcPending
																if v112 != 0 {
																	return int32(0)
																} else {
																	v118 = *(*int32)(unsafe.Add(mBase, uint32(v9)+96))
																	F_ReleaseCatCache(m, v12)
																	mBase = m.M
																	v120 = m.ExcPending
																	if v120 != 0 {
																		return int32(0)
																	} else {
																		v123 = v118
																		m.G0 = v9 + int32(112)
																		return v123
																	}
																}
															}
														} else {
															v102 = F_format_type_be(m, v99)
															mBase = m.M
															v103 = m.ExcPending
															if v103 != 0 {
																return int32(0)
															} else {
																v104 = v102
																*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v104
																F_appendStringInfo(m, v9+int32(96), int32(655025), v9+int32(16))
																mBase = m.M
																v112 = m.ExcPending
																if v112 != 0 {
																	return int32(0)
																} else {
																	v118 = *(*int32)(unsafe.Add(mBase, uint32(v9)+96))
																	F_ReleaseCatCache(m, v12)
																	mBase = m.M
																	v120 = m.ExcPending
																	if v120 != 0 {
																		return int32(0)
																	} else {
																		v123 = v118
																		m.G0 = v9 + int32(112)
																		return v123
																	}
																}
															}
														}
													} else {
														F_appendStringInfoString(m, v9+int32(96), int32(661660))
														mBase = m.M
														v117 = m.ExcPending
														if v117 != 0 {
															return int32(0)
														} else {
															v118 = *(*int32)(unsafe.Add(mBase, uint32(v9)+96))
															F_ReleaseCatCache(m, v12)
															mBase = m.M
															v120 = m.ExcPending
															if v120 != 0 {
																return int32(0)
															} else {
																v123 = v118
																m.G0 = v9 + int32(112)
																return v123
															}
														}
													}
												}
											} else {
												if v26 == int32(0) {
													v85 = v53
													v86 = F_format_type_be(m, v85)
													mBase = m.M
													v87 = m.ExcPending
													if v87 != 0 {
														return int32(0)
													} else {
														v89 = v86
														*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v89
														F_appendStringInfo(m, v9+int32(96), int32(649293), v9+int32(32))
														mBase = m.M
														v97 = m.ExcPending
														if v97 != 0 {
															return int32(0)
														} else {
															v99 = *(*int32)(unsafe.Add(mBase, uint32(v18)+84))
															if v99 != 0 {
																if v26 != 0 {
																	v100 = F_format_type_be_qualified(m, v99)
																	mBase = m.M
																	v101 = m.ExcPending
																	if v101 != 0 {
																		return int32(0)
																	} else {
																		v104 = v100
																		*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v104
																		F_appendStringInfo(m, v9+int32(96), int32(655025), v9+int32(16))
																		mBase = m.M
																		v112 = m.ExcPending
																		if v112 != 0 {
																			return int32(0)
																		} else {
																			v118 = *(*int32)(unsafe.Add(mBase, uint32(v9)+96))
																			F_ReleaseCatCache(m, v12)
																			mBase = m.M
																			v120 = m.ExcPending
																			if v120 != 0 {
																				return int32(0)
																			} else {
																				v123 = v118
																				m.G0 = v9 + int32(112)
																				return v123
																			}
																		}
																	}
																} else {
																	v102 = F_format_type_be(m, v99)
																	mBase = m.M
																	v103 = m.ExcPending
																	if v103 != 0 {
																		return int32(0)
																	} else {
																		v104 = v102
																		*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v104
																		F_appendStringInfo(m, v9+int32(96), int32(655025), v9+int32(16))
																		mBase = m.M
																		v112 = m.ExcPending
																		if v112 != 0 {
																			return int32(0)
																		} else {
																			v118 = *(*int32)(unsafe.Add(mBase, uint32(v9)+96))
																			F_ReleaseCatCache(m, v12)
																			mBase = m.M
																			v120 = m.ExcPending
																			if v120 != 0 {
																				return int32(0)
																			} else {
																				v123 = v118
																				m.G0 = v9 + int32(112)
																				return v123
																			}
																		}
																	}
																}
															} else {
																F_appendStringInfoString(m, v9+int32(96), int32(661660))
																mBase = m.M
																v117 = m.ExcPending
																if v117 != 0 {
																	return int32(0)
																} else {
																	v118 = *(*int32)(unsafe.Add(mBase, uint32(v9)+96))
																	F_ReleaseCatCache(m, v12)
																	mBase = m.M
																	v120 = m.ExcPending
																	if v120 != 0 {
																		return int32(0)
																	} else {
																		v123 = v118
																		m.G0 = v9 + int32(112)
																		return v123
																	}
																}
															}
														}
													}
												} else {
													v58 = F_format_type_be_qualified(m, v53)
													mBase = m.M
													v59 = m.ExcPending
													if v59 != 0 {
														return int32(0)
													} else {
														v89 = v58
														*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v89
														F_appendStringInfo(m, v9+int32(96), int32(649293), v9+int32(32))
														mBase = m.M
														v97 = m.ExcPending
														if v97 != 0 {
															return int32(0)
														} else {
															v99 = *(*int32)(unsafe.Add(mBase, uint32(v18)+84))
															if v99 != 0 {
																if v26 != 0 {
																	v100 = F_format_type_be_qualified(m, v99)
																	mBase = m.M
																	v101 = m.ExcPending
																	if v101 != 0 {
																		return int32(0)
																	} else {
																		v104 = v100
																		*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v104
																		F_appendStringInfo(m, v9+int32(96), int32(655025), v9+int32(16))
																		mBase = m.M
																		v112 = m.ExcPending
																		if v112 != 0 {
																			return int32(0)
																		} else {
																			v118 = *(*int32)(unsafe.Add(mBase, uint32(v9)+96))
																			F_ReleaseCatCache(m, v12)
																			mBase = m.M
																			v120 = m.ExcPending
																			if v120 != 0 {
																				return int32(0)
																			} else {
																				v123 = v118
																				m.G0 = v9 + int32(112)
																				return v123
																			}
																		}
																	}
																} else {
																	v102 = F_format_type_be(m, v99)
																	mBase = m.M
																	v103 = m.ExcPending
																	if v103 != 0 {
																		return int32(0)
																	} else {
																		v104 = v102
																		*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v104
																		F_appendStringInfo(m, v9+int32(96), int32(655025), v9+int32(16))
																		mBase = m.M
																		v112 = m.ExcPending
																		if v112 != 0 {
																			return int32(0)
																		} else {
																			v118 = *(*int32)(unsafe.Add(mBase, uint32(v9)+96))
																			F_ReleaseCatCache(m, v12)
																			mBase = m.M
																			v120 = m.ExcPending
																			if v120 != 0 {
																				return int32(0)
																			} else {
																				v123 = v118
																				m.G0 = v9 + int32(112)
																				return v123
																			}
																		}
																	}
																}
															} else {
																F_appendStringInfoString(m, v9+int32(96), int32(661660))
																mBase = m.M
																v117 = m.ExcPending
																if v117 != 0 {
																	return int32(0)
																} else {
																	v118 = *(*int32)(unsafe.Add(mBase, uint32(v9)+96))
																	F_ReleaseCatCache(m, v12)
																	mBase = m.M
																	v120 = m.ExcPending
																	if v120 != 0 {
																		return int32(0)
																	} else {
																		v123 = v118
																		m.G0 = v9 + int32(112)
																		return v123
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
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v18)+68))
					v33 = F_get_namespace_name(m, v32)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int32(0)
					} else {
						v35 = F_quote_identifier(m, v33)
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9)+80)) = v35
							F_appendStringInfo(m, v9+int32(96), int32(586069), v9+int32(80))
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9)+64)) = v20
								F_appendStringInfo(m, v9+int32(96), int32(664818), v9-int32(-64))
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
									return int32(0)
								} else {
									v53 = *(*int32)(unsafe.Add(mBase, uint32(v18)+80))
									if v53 == int32(0) {
										F_appendStringInfoString(m, v9+int32(96), int32(649297))
										mBase = m.M
										v74 = m.ExcPending
										if v74 != 0 {
											return int32(0)
										} else {
											v99 = *(*int32)(unsafe.Add(mBase, uint32(v18)+84))
											if v99 != 0 {
												if v26 != 0 {
													v100 = F_format_type_be_qualified(m, v99)
													mBase = m.M
													v101 = m.ExcPending
													if v101 != 0 {
														return int32(0)
													} else {
														v104 = v100
														*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v104
														F_appendStringInfo(m, v9+int32(96), int32(655025), v9+int32(16))
														mBase = m.M
														v112 = m.ExcPending
														if v112 != 0 {
															return int32(0)
														} else {
															v118 = *(*int32)(unsafe.Add(mBase, uint32(v9)+96))
															F_ReleaseCatCache(m, v12)
															mBase = m.M
															v120 = m.ExcPending
															if v120 != 0 {
																return int32(0)
															} else {
																v123 = v118
																m.G0 = v9 + int32(112)
																return v123
															}
														}
													}
												} else {
													v102 = F_format_type_be(m, v99)
													mBase = m.M
													v103 = m.ExcPending
													if v103 != 0 {
														return int32(0)
													} else {
														v104 = v102
														*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v104
														F_appendStringInfo(m, v9+int32(96), int32(655025), v9+int32(16))
														mBase = m.M
														v112 = m.ExcPending
														if v112 != 0 {
															return int32(0)
														} else {
															v118 = *(*int32)(unsafe.Add(mBase, uint32(v9)+96))
															F_ReleaseCatCache(m, v12)
															mBase = m.M
															v120 = m.ExcPending
															if v120 != 0 {
																return int32(0)
															} else {
																v123 = v118
																m.G0 = v9 + int32(112)
																return v123
															}
														}
													}
												}
											} else {
												F_appendStringInfoString(m, v9+int32(96), int32(661660))
												mBase = m.M
												v117 = m.ExcPending
												if v117 != 0 {
													return int32(0)
												} else {
													v118 = *(*int32)(unsafe.Add(mBase, uint32(v9)+96))
													F_ReleaseCatCache(m, v12)
													mBase = m.M
													v120 = m.ExcPending
													if v120 != 0 {
														return int32(0)
													} else {
														v123 = v118
														m.G0 = v9 + int32(112)
														return v123
													}
												}
											}
										}
									} else {
										if v26 == int32(0) {
											v85 = v53
											v86 = F_format_type_be(m, v85)
											mBase = m.M
											v87 = m.ExcPending
											if v87 != 0 {
												return int32(0)
											} else {
												v89 = v86
												*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v89
												F_appendStringInfo(m, v9+int32(96), int32(649293), v9+int32(32))
												mBase = m.M
												v97 = m.ExcPending
												if v97 != 0 {
													return int32(0)
												} else {
													v99 = *(*int32)(unsafe.Add(mBase, uint32(v18)+84))
													if v99 != 0 {
														if v26 != 0 {
															v100 = F_format_type_be_qualified(m, v99)
															mBase = m.M
															v101 = m.ExcPending
															if v101 != 0 {
																return int32(0)
															} else {
																v104 = v100
																*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v104
																F_appendStringInfo(m, v9+int32(96), int32(655025), v9+int32(16))
																mBase = m.M
																v112 = m.ExcPending
																if v112 != 0 {
																	return int32(0)
																} else {
																	v118 = *(*int32)(unsafe.Add(mBase, uint32(v9)+96))
																	F_ReleaseCatCache(m, v12)
																	mBase = m.M
																	v120 = m.ExcPending
																	if v120 != 0 {
																		return int32(0)
																	} else {
																		v123 = v118
																		m.G0 = v9 + int32(112)
																		return v123
																	}
																}
															}
														} else {
															v102 = F_format_type_be(m, v99)
															mBase = m.M
															v103 = m.ExcPending
															if v103 != 0 {
																return int32(0)
															} else {
																v104 = v102
																*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v104
																F_appendStringInfo(m, v9+int32(96), int32(655025), v9+int32(16))
																mBase = m.M
																v112 = m.ExcPending
																if v112 != 0 {
																	return int32(0)
																} else {
																	v118 = *(*int32)(unsafe.Add(mBase, uint32(v9)+96))
																	F_ReleaseCatCache(m, v12)
																	mBase = m.M
																	v120 = m.ExcPending
																	if v120 != 0 {
																		return int32(0)
																	} else {
																		v123 = v118
																		m.G0 = v9 + int32(112)
																		return v123
																	}
																}
															}
														}
													} else {
														F_appendStringInfoString(m, v9+int32(96), int32(661660))
														mBase = m.M
														v117 = m.ExcPending
														if v117 != 0 {
															return int32(0)
														} else {
															v118 = *(*int32)(unsafe.Add(mBase, uint32(v9)+96))
															F_ReleaseCatCache(m, v12)
															mBase = m.M
															v120 = m.ExcPending
															if v120 != 0 {
																return int32(0)
															} else {
																v123 = v118
																m.G0 = v9 + int32(112)
																return v123
															}
														}
													}
												}
											}
										} else {
											v58 = F_format_type_be_qualified(m, v53)
											mBase = m.M
											v59 = m.ExcPending
											if v59 != 0 {
												return int32(0)
											} else {
												v89 = v58
												*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v89
												F_appendStringInfo(m, v9+int32(96), int32(649293), v9+int32(32))
												mBase = m.M
												v97 = m.ExcPending
												if v97 != 0 {
													return int32(0)
												} else {
													v99 = *(*int32)(unsafe.Add(mBase, uint32(v18)+84))
													if v99 != 0 {
														if v26 != 0 {
															v100 = F_format_type_be_qualified(m, v99)
															mBase = m.M
															v101 = m.ExcPending
															if v101 != 0 {
																return int32(0)
															} else {
																v104 = v100
																*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v104
																F_appendStringInfo(m, v9+int32(96), int32(655025), v9+int32(16))
																mBase = m.M
																v112 = m.ExcPending
																if v112 != 0 {
																	return int32(0)
																} else {
																	v118 = *(*int32)(unsafe.Add(mBase, uint32(v9)+96))
																	F_ReleaseCatCache(m, v12)
																	mBase = m.M
																	v120 = m.ExcPending
																	if v120 != 0 {
																		return int32(0)
																	} else {
																		v123 = v118
																		m.G0 = v9 + int32(112)
																		return v123
																	}
																}
															}
														} else {
															v102 = F_format_type_be(m, v99)
															mBase = m.M
															v103 = m.ExcPending
															if v103 != 0 {
																return int32(0)
															} else {
																v104 = v102
																*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v104
																F_appendStringInfo(m, v9+int32(96), int32(655025), v9+int32(16))
																mBase = m.M
																v112 = m.ExcPending
																if v112 != 0 {
																	return int32(0)
																} else {
																	v118 = *(*int32)(unsafe.Add(mBase, uint32(v9)+96))
																	F_ReleaseCatCache(m, v12)
																	mBase = m.M
																	v120 = m.ExcPending
																	if v120 != 0 {
																		return int32(0)
																	} else {
																		v123 = v118
																		m.G0 = v9 + int32(112)
																		return v123
																	}
																}
															}
														}
													} else {
														F_appendStringInfoString(m, v9+int32(96), int32(661660))
														mBase = m.M
														v117 = m.ExcPending
														if v117 != 0 {
															return int32(0)
														} else {
															v118 = *(*int32)(unsafe.Add(mBase, uint32(v9)+96))
															F_ReleaseCatCache(m, v12)
															mBase = m.M
															v120 = m.ExcPending
															if v120 != 0 {
																return int32(0)
															} else {
																v123 = v118
																m.G0 = v9 + int32(112)
																return v123
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
				v123 = int32(0)
				m.G0 = v9 + int32(112)
				return v123
			} else {
				v78 = F_palloc(m, int32(64))
				mBase = m.M
				v79 = m.ExcPending
				if v79 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
					v83 = F_pg_snprintf(m, v78, int32(64), int32(59182), v9)
					mBase = m.M
					v84 = m.ExcPending
					if v84 != 0 {
						return int32(0)
					} else {
						v123 = v78
						m.G0 = v9 + int32(112)
						return v123
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
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
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
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	v9 = m.G0
	v11 = v9 - int32(80)
	m.G0 = v11
	F_initStringInfo(m, v11-int32(-64))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v20 = F_SearchSysCache1(m, int32(40), l0)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			if v20 != 0 {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
				v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+22)))
				v24 = v22 + v23
				v26 = v24 + int32(4)
				v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+76)))
				switch v27 - int32(98) {
				case 0:
					v60 = F_makeString(m, v26)
					mBase = m.M
					v61 = m.ExcPending
					if v61 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v60
						*(*int32)(unsafe.Add(mBase, uint32(v11)+60)) = v60
						v68 = F_list_make1_impl(m, int32(1), v11+int32(48))
						mBase = m.M
						v69 = m.ExcPending
						if v69 != 0 {
							return int32(0)
						} else {
							v72 = F_oper(m, int32(0), v68, l1, l2, int32(1), int32(-1))
							mBase = m.M
							v73 = m.ExcPending
							if v73 != 0 {
								return int32(0)
							} else {
								v76 = v72
								if v76 != 0 {
									v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+16))
									v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+22)))
									v80 = *(*int32)(unsafe.Add(mBase, uint32(v77+v78)))
									if l0 == v80 {
										F_appendStringInfoString(m, v11-int32(-64), v26)
										mBase = m.M
										v123 = m.ExcPending
										if v123 != 0 {
											return int32(0)
										} else {
											F_ReleaseCatCache(m, v76)
											mBase = m.M
											v127 = m.ExcPending
											if v127 != 0 {
												return int32(0)
											} else {
												F_ReleaseCatCache(m, v20)
												mBase = m.M
												v131 = m.ExcPending
												if v131 != 0 {
													return int32(0)
												} else {
													v132 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
													m.G0 = v11 + int32(80)
													return v132
												}
											}
										}
									} else {
										v83 = *(*int32)(unsafe.Add(mBase, uint32(v24)+68))
										v84 = F_get_namespace_name_or_temp(m, v83)
										mBase = m.M
										v85 = m.ExcPending
										if v85 != 0 {
											return int32(0)
										} else {
											v86 = F_quote_identifier(m, v84)
											mBase = m.M
											v87 = m.ExcPending
											if v87 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v86
												F_appendStringInfo(m, v11-int32(-64), int32(582415), v11+int32(32))
												mBase = m.M
												v95 = m.ExcPending
												if v95 != 0 {
													return int32(0)
												} else {
													F_appendStringInfoString(m, v11-int32(-64), v26)
													mBase = m.M
													v99 = m.ExcPending
													if v99 != 0 {
														return int32(0)
													} else {
														if v84 != 0 {
															F_appendStringInfoChar(m, v11-int32(-64), int32(41))
															mBase = m.M
															v104 = m.ExcPending
															if v104 != 0 {
																return int32(0)
															} else {
																if v76 == int32(0) {
																	F_ReleaseCatCache(m, v20)
																	mBase = m.M
																	v131 = m.ExcPending
																	if v131 != 0 {
																		return int32(0)
																	} else {
																		v132 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
																		m.G0 = v11 + int32(80)
																		return v132
																	}
																} else {
																	F_ReleaseCatCache(m, v76)
																	mBase = m.M
																	v127 = m.ExcPending
																	if v127 != 0 {
																		return int32(0)
																	} else {
																		F_ReleaseCatCache(m, v20)
																		mBase = m.M
																		v131 = m.ExcPending
																		if v131 != 0 {
																			return int32(0)
																		} else {
																			v132 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
																			m.G0 = v11 + int32(80)
																			return v132
																		}
																	}
																}
															}
														} else {
															if v76 == int32(0) {
																F_ReleaseCatCache(m, v20)
																mBase = m.M
																v131 = m.ExcPending
																if v131 != 0 {
																	return int32(0)
																} else {
																	v132 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
																	m.G0 = v11 + int32(80)
																	return v132
																}
															} else {
																F_ReleaseCatCache(m, v76)
																mBase = m.M
																v127 = m.ExcPending
																if v127 != 0 {
																	return int32(0)
																} else {
																	F_ReleaseCatCache(m, v20)
																	mBase = m.M
																	v131 = m.ExcPending
																	if v131 != 0 {
																		return int32(0)
																	} else {
																		v132 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
																		m.G0 = v11 + int32(80)
																		return v132
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
									v83 = *(*int32)(unsafe.Add(mBase, uint32(v24)+68))
									v84 = F_get_namespace_name_or_temp(m, v83)
									mBase = m.M
									v85 = m.ExcPending
									if v85 != 0 {
										return int32(0)
									} else {
										v86 = F_quote_identifier(m, v84)
										mBase = m.M
										v87 = m.ExcPending
										if v87 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v86
											F_appendStringInfo(m, v11-int32(-64), int32(582415), v11+int32(32))
											mBase = m.M
											v95 = m.ExcPending
											if v95 != 0 {
												return int32(0)
											} else {
												F_appendStringInfoString(m, v11-int32(-64), v26)
												mBase = m.M
												v99 = m.ExcPending
												if v99 != 0 {
													return int32(0)
												} else {
													if v84 != 0 {
														F_appendStringInfoChar(m, v11-int32(-64), int32(41))
														mBase = m.M
														v104 = m.ExcPending
														if v104 != 0 {
															return int32(0)
														} else {
															if v76 == int32(0) {
																F_ReleaseCatCache(m, v20)
																mBase = m.M
																v131 = m.ExcPending
																if v131 != 0 {
																	return int32(0)
																} else {
																	v132 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
																	m.G0 = v11 + int32(80)
																	return v132
																}
															} else {
																F_ReleaseCatCache(m, v76)
																mBase = m.M
																v127 = m.ExcPending
																if v127 != 0 {
																	return int32(0)
																} else {
																	F_ReleaseCatCache(m, v20)
																	mBase = m.M
																	v131 = m.ExcPending
																	if v131 != 0 {
																		return int32(0)
																	} else {
																		v132 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
																		m.G0 = v11 + int32(80)
																		return v132
																	}
																}
															}
														}
													} else {
														if v76 == int32(0) {
															F_ReleaseCatCache(m, v20)
															mBase = m.M
															v131 = m.ExcPending
															if v131 != 0 {
																return int32(0)
															} else {
																v132 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
																m.G0 = v11 + int32(80)
																return v132
															}
														} else {
															F_ReleaseCatCache(m, v76)
															mBase = m.M
															v127 = m.ExcPending
															if v127 != 0 {
																return int32(0)
															} else {
																F_ReleaseCatCache(m, v20)
																mBase = m.M
																v131 = m.ExcPending
																if v131 != 0 {
																	return int32(0)
																} else {
																	v132 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
																	m.G0 = v11 + int32(80)
																	return v132
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
					v47 = m.ExcPending
					if v47 != 0 {
						return int32(0)
					} else {
						v48 = int32(*(*int8)(unsafe.Add(mBase, uint32(v24)+76)))
						*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v48
						F_errmsg_internal(m, int32(482298), v11+int32(16))
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(488965), int32(13395), int32(374927))
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				case 10:
					v30 = F_makeString(m, v26)
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = v30
						*(*int32)(unsafe.Add(mBase, uint32(v11)+56)) = v30
						v38 = F_list_make1_impl(m, int32(1), v11+int32(52))
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int32(0)
						} else {
							v42 = F_left_oper(m, int32(0), v38, l2, int32(1), int32(-1))
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
								return int32(0)
							} else {
								v76 = v42
								if v76 != 0 {
									v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+16))
									v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+22)))
									v80 = *(*int32)(unsafe.Add(mBase, uint32(v77+v78)))
									if l0 == v80 {
										F_appendStringInfoString(m, v11-int32(-64), v26)
										mBase = m.M
										v123 = m.ExcPending
										if v123 != 0 {
											return int32(0)
										} else {
											F_ReleaseCatCache(m, v76)
											mBase = m.M
											v127 = m.ExcPending
											if v127 != 0 {
												return int32(0)
											} else {
												F_ReleaseCatCache(m, v20)
												mBase = m.M
												v131 = m.ExcPending
												if v131 != 0 {
													return int32(0)
												} else {
													v132 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
													m.G0 = v11 + int32(80)
													return v132
												}
											}
										}
									} else {
										v83 = *(*int32)(unsafe.Add(mBase, uint32(v24)+68))
										v84 = F_get_namespace_name_or_temp(m, v83)
										mBase = m.M
										v85 = m.ExcPending
										if v85 != 0 {
											return int32(0)
										} else {
											v86 = F_quote_identifier(m, v84)
											mBase = m.M
											v87 = m.ExcPending
											if v87 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v86
												F_appendStringInfo(m, v11-int32(-64), int32(582415), v11+int32(32))
												mBase = m.M
												v95 = m.ExcPending
												if v95 != 0 {
													return int32(0)
												} else {
													F_appendStringInfoString(m, v11-int32(-64), v26)
													mBase = m.M
													v99 = m.ExcPending
													if v99 != 0 {
														return int32(0)
													} else {
														if v84 != 0 {
															F_appendStringInfoChar(m, v11-int32(-64), int32(41))
															mBase = m.M
															v104 = m.ExcPending
															if v104 != 0 {
																return int32(0)
															} else {
																if v76 == int32(0) {
																	F_ReleaseCatCache(m, v20)
																	mBase = m.M
																	v131 = m.ExcPending
																	if v131 != 0 {
																		return int32(0)
																	} else {
																		v132 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
																		m.G0 = v11 + int32(80)
																		return v132
																	}
																} else {
																	F_ReleaseCatCache(m, v76)
																	mBase = m.M
																	v127 = m.ExcPending
																	if v127 != 0 {
																		return int32(0)
																	} else {
																		F_ReleaseCatCache(m, v20)
																		mBase = m.M
																		v131 = m.ExcPending
																		if v131 != 0 {
																			return int32(0)
																		} else {
																			v132 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
																			m.G0 = v11 + int32(80)
																			return v132
																		}
																	}
																}
															}
														} else {
															if v76 == int32(0) {
																F_ReleaseCatCache(m, v20)
																mBase = m.M
																v131 = m.ExcPending
																if v131 != 0 {
																	return int32(0)
																} else {
																	v132 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
																	m.G0 = v11 + int32(80)
																	return v132
																}
															} else {
																F_ReleaseCatCache(m, v76)
																mBase = m.M
																v127 = m.ExcPending
																if v127 != 0 {
																	return int32(0)
																} else {
																	F_ReleaseCatCache(m, v20)
																	mBase = m.M
																	v131 = m.ExcPending
																	if v131 != 0 {
																		return int32(0)
																	} else {
																		v132 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
																		m.G0 = v11 + int32(80)
																		return v132
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
									v83 = *(*int32)(unsafe.Add(mBase, uint32(v24)+68))
									v84 = F_get_namespace_name_or_temp(m, v83)
									mBase = m.M
									v85 = m.ExcPending
									if v85 != 0 {
										return int32(0)
									} else {
										v86 = F_quote_identifier(m, v84)
										mBase = m.M
										v87 = m.ExcPending
										if v87 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v86
											F_appendStringInfo(m, v11-int32(-64), int32(582415), v11+int32(32))
											mBase = m.M
											v95 = m.ExcPending
											if v95 != 0 {
												return int32(0)
											} else {
												F_appendStringInfoString(m, v11-int32(-64), v26)
												mBase = m.M
												v99 = m.ExcPending
												if v99 != 0 {
													return int32(0)
												} else {
													if v84 != 0 {
														F_appendStringInfoChar(m, v11-int32(-64), int32(41))
														mBase = m.M
														v104 = m.ExcPending
														if v104 != 0 {
															return int32(0)
														} else {
															if v76 == int32(0) {
																F_ReleaseCatCache(m, v20)
																mBase = m.M
																v131 = m.ExcPending
																if v131 != 0 {
																	return int32(0)
																} else {
																	v132 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
																	m.G0 = v11 + int32(80)
																	return v132
																}
															} else {
																F_ReleaseCatCache(m, v76)
																mBase = m.M
																v127 = m.ExcPending
																if v127 != 0 {
																	return int32(0)
																} else {
																	F_ReleaseCatCache(m, v20)
																	mBase = m.M
																	v131 = m.ExcPending
																	if v131 != 0 {
																		return int32(0)
																	} else {
																		v132 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
																		m.G0 = v11 + int32(80)
																		return v132
																	}
																}
															}
														}
													} else {
														if v76 == int32(0) {
															F_ReleaseCatCache(m, v20)
															mBase = m.M
															v131 = m.ExcPending
															if v131 != 0 {
																return int32(0)
															} else {
																v132 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
																m.G0 = v11 + int32(80)
																return v132
															}
														} else {
															F_ReleaseCatCache(m, v76)
															mBase = m.M
															v127 = m.ExcPending
															if v127 != 0 {
																return int32(0)
															} else {
																F_ReleaseCatCache(m, v20)
																mBase = m.M
																v131 = m.ExcPending
																if v131 != 0 {
																	return int32(0)
																} else {
																	v132 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
																	m.G0 = v11 + int32(80)
																	return v132
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
				v110 = m.ExcPending
				if v110 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v11))) = l0
					F_errmsg_internal(m, int32(42856), v11)
					mBase = m.M
					v114 = m.ExcPending
					if v114 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(488965), int32(13375), int32(374927))
						mBase = m.M
						v119 = m.ExcPending
						if v119 != 0 {
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
