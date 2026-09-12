package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_percentile_cont_final_common(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 float64
	_ = v22
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int64
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int64
	_ = v57
	var v61 float64
	_ = v61
	var v62 float64
	_ = v62
	var v66 int64
	_ = v66
	var v68 int64
	_ = v68
	var v69 int32
	_ = v69
	var v70 float64
	_ = v70
	var v74 int64
	_ = v74
	var v76 int64
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int64
	_ = v121
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v137 int32
	_ = v137
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v15 == int32(1) {
		v18 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v18)
		v137 = int32(0)
		m.G0 = v13 + int32(32)
		return v137
	} else {
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v22 = *(*float64)(unsafe.Add(mBase, uint32(v21)))
		if base.F64_lt(v22, float64(0)) != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v145 = m.ExcPending
			if v145 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50331778))
				mBase = m.M
				v148 = m.ExcPending
				if v148 != 0 {
					return int32(0)
				} else {
					*(*float64)(unsafe.Add(mBase, uint32(v13))) = v22
					F_errmsg(m, int32(536015), v13)
					mBase = m.M
					v152 = m.ExcPending
					if v152 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(477321), int32(552), int32(236656))
						mBase = m.M
						v157 = m.ExcPending
						if v157 != 0 {
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
			if base.F64_gt(v22, float64(1)) != 0 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v145 = m.ExcPending
				if v145 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(50331778))
					mBase = m.M
					v148 = m.ExcPending
					if v148 != 0 {
						return int32(0)
					} else {
						*(*float64)(unsafe.Add(mBase, uint32(v13))) = v22
						F_errmsg(m, int32(536015), v13)
						mBase = m.M
						v152 = m.ExcPending
						if v152 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(477321), int32(552), int32(236656))
							mBase = m.M
							v157 = m.ExcPending
							if v157 != 0 {
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
				if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v22)&int64(9223372036854775807)) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v145 = m.ExcPending
					if v145 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(50331778))
						mBase = m.M
						v148 = m.ExcPending
						if v148 != 0 {
							return int32(0)
						} else {
							*(*float64)(unsafe.Add(mBase, uint32(v13))) = v22
							F_errmsg(m, int32(536015), v13)
							mBase = m.M
							v152 = m.ExcPending
							if v152 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(477321), int32(552), int32(236656))
								mBase = m.M
								v157 = m.ExcPending
								if v157 != 0 {
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
					v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
					if v32 == int32(1) {
						v35 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v35)
						v137 = int32(0)
						m.G0 = v13 + int32(32)
						return v137
					} else {
						v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						v39 = *(*int64)(unsafe.Add(mBase, uint32(v38)+16))
						if v39 == int64(0) {
							v42 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v42)
							v137 = int32(0)
							m.G0 = v13 + int32(32)
							return v137
						} else {
							v45 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
							v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+24)))
							if v46 == int32(0) {
								F_tuplesort_performsort(m, v45)
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
									return int32(0)
								} else {
									v53 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v38)+24)) = uint8(v53)
									v57 = *(*int64)(unsafe.Add(mBase, uint32(v38)+16))
									v61 = base.F64_mul(v22, base.F64_convert_i64_s(v57-int64(1)))
									v62 = base.F64_ceil(v61)
									if base.F64_lt(base.F64_abs(v62), float64(9.223372036854776e+18)) != 0 {
										v66 = base.I64_trunc_f64_s(v62)
										v68 = v66
									} else {
										v68 = int64(-9223372036854775807 - 1)
									}
									v69 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
									v70 = base.F64_floor(v61)
									if base.F64_lt(base.F64_abs(v70), float64(9.223372036854776e+18)) != 0 {
										v74 = base.I64_trunc_f64_s(v70)
										v76 = v74
									} else {
										v76 = int64(-9223372036854775807 - 1)
									}
									v77 = F_tuplesort_skiptuples(m, v69, v76)
									mBase = m.M
									v78 = m.ExcPending
									if v78 != 0 {
										return int32(0)
									} else {
										if v77 == int32(0) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v161 = m.ExcPending
											if v161 != 0 {
												return int32(0)
											} else {
												F_errmsg_internal(m, int32(84748), int32(0))
												mBase = m.M
												v165 = m.ExcPending
												if v165 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(477321), int32(581), int32(236656))
													mBase = m.M
													v170 = m.ExcPending
													if v170 != 0 {
														return int32(0)
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											}
										} else {
											v81 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
											v82 = int32(1)
											v89 = F_tuplesort_getdatum(m, v81, v82, v82, v13+int32(28), v13+int32(23), int32(0))
											mBase = m.M
											v90 = m.ExcPending
											if v90 != 0 {
												return int32(0)
											} else {
												if v89 == int32(0) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v174 = m.ExcPending
													if v174 != 0 {
														return int32(0)
													} else {
														F_errmsg_internal(m, int32(84748), int32(0))
														mBase = m.M
														v178 = m.ExcPending
														if v178 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(477321), int32(585), int32(236656))
															mBase = m.M
															v183 = m.ExcPending
															if v183 != 0 {
																return int32(0)
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												} else {
													v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+23)))
													if v93 == int32(1) {
														v96 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v96)
														v137 = int32(0)
														m.G0 = v13 + int32(32)
														return v137
													} else {
														if v76 == v68 {
															v100 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
															v137 = v100
															m.G0 = v13 + int32(32)
															return v137
														} else {
															v101 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
															v102 = int32(1)
															v109 = F_tuplesort_getdatum(m, v101, v102, v102, v13+int32(24), v13+int32(23), int32(0))
															mBase = m.M
															v110 = m.ExcPending
															if v110 != 0 {
																return int32(0)
															} else {
																if v109 == int32(0) {
																	F_errstart_cold(m, int32(21), int32(0))
																	mBase = m.M
																	v187 = m.ExcPending
																	if v187 != 0 {
																		return int32(0)
																	} else {
																		F_errmsg_internal(m, int32(84748), int32(0))
																		mBase = m.M
																		v191 = m.ExcPending
																		if v191 != 0 {
																			return int32(0)
																		} else {
																			F_errfinish(m, int32(477321), int32(597), int32(236656))
																			mBase = m.M
																			v196 = m.ExcPending
																			if v196 != 0 {
																				return int32(0)
																			} else {
																				base.Wasm_trap_unreachable()
																				for {
																				}
																			}
																		}
																	}
																} else {
																	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+23)))
																	if v113 == int32(1) {
																		v116 = int32(1)
																		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v116)
																		v137 = int32(0)
																		m.G0 = v13 + int32(32)
																		return v137
																	} else {
																		v119 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
																		v120 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
																		v121 = *(*int64)(unsafe.Add(mBase, uint32(v38)+16))
																		v128 = m.T0[l1].(func(*base.Module, int32, int32, float64) int32)(m, v119, v120, base.F64_sub(base.F64_mul(v22, base.F64_convert_i64_s(v121-int64(1))), base.F64_convert_i64_s(v76)))
																		mBase = m.M
																		v129 = m.ExcPending
																		if v129 != 0 {
																			return int32(0)
																		} else {
																			v137 = v128
																			m.G0 = v13 + int32(32)
																			return v137
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
								F_tuplesort_rescan(m, v45)
								mBase = m.M
								v56 = m.ExcPending
								if v56 != 0 {
									return int32(0)
								} else {
									v57 = *(*int64)(unsafe.Add(mBase, uint32(v38)+16))
									v61 = base.F64_mul(v22, base.F64_convert_i64_s(v57-int64(1)))
									v62 = base.F64_ceil(v61)
									if base.F64_lt(base.F64_abs(v62), float64(9.223372036854776e+18)) != 0 {
										v66 = base.I64_trunc_f64_s(v62)
										v68 = v66
									} else {
										v68 = int64(-9223372036854775807 - 1)
									}
									v69 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
									v70 = base.F64_floor(v61)
									if base.F64_lt(base.F64_abs(v70), float64(9.223372036854776e+18)) != 0 {
										v74 = base.I64_trunc_f64_s(v70)
										v76 = v74
									} else {
										v76 = int64(-9223372036854775807 - 1)
									}
									v77 = F_tuplesort_skiptuples(m, v69, v76)
									mBase = m.M
									v78 = m.ExcPending
									if v78 != 0 {
										return int32(0)
									} else {
										if v77 == int32(0) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v161 = m.ExcPending
											if v161 != 0 {
												return int32(0)
											} else {
												F_errmsg_internal(m, int32(84748), int32(0))
												mBase = m.M
												v165 = m.ExcPending
												if v165 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(477321), int32(581), int32(236656))
													mBase = m.M
													v170 = m.ExcPending
													if v170 != 0 {
														return int32(0)
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											}
										} else {
											v81 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
											v82 = int32(1)
											v89 = F_tuplesort_getdatum(m, v81, v82, v82, v13+int32(28), v13+int32(23), int32(0))
											mBase = m.M
											v90 = m.ExcPending
											if v90 != 0 {
												return int32(0)
											} else {
												if v89 == int32(0) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v174 = m.ExcPending
													if v174 != 0 {
														return int32(0)
													} else {
														F_errmsg_internal(m, int32(84748), int32(0))
														mBase = m.M
														v178 = m.ExcPending
														if v178 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(477321), int32(585), int32(236656))
															mBase = m.M
															v183 = m.ExcPending
															if v183 != 0 {
																return int32(0)
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												} else {
													v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+23)))
													if v93 == int32(1) {
														v96 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v96)
														v137 = int32(0)
														m.G0 = v13 + int32(32)
														return v137
													} else {
														if v76 == v68 {
															v100 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
															v137 = v100
															m.G0 = v13 + int32(32)
															return v137
														} else {
															v101 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
															v102 = int32(1)
															v109 = F_tuplesort_getdatum(m, v101, v102, v102, v13+int32(24), v13+int32(23), int32(0))
															mBase = m.M
															v110 = m.ExcPending
															if v110 != 0 {
																return int32(0)
															} else {
																if v109 == int32(0) {
																	F_errstart_cold(m, int32(21), int32(0))
																	mBase = m.M
																	v187 = m.ExcPending
																	if v187 != 0 {
																		return int32(0)
																	} else {
																		F_errmsg_internal(m, int32(84748), int32(0))
																		mBase = m.M
																		v191 = m.ExcPending
																		if v191 != 0 {
																			return int32(0)
																		} else {
																			F_errfinish(m, int32(477321), int32(597), int32(236656))
																			mBase = m.M
																			v196 = m.ExcPending
																			if v196 != 0 {
																				return int32(0)
																			} else {
																				base.Wasm_trap_unreachable()
																				for {
																				}
																			}
																		}
																	}
																} else {
																	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+23)))
																	if v113 == int32(1) {
																		v116 = int32(1)
																		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v116)
																		v137 = int32(0)
																		m.G0 = v13 + int32(32)
																		return v137
																	} else {
																		v119 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
																		v120 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
																		v121 = *(*int64)(unsafe.Add(mBase, uint32(v38)+16))
																		v128 = m.T0[l1].(func(*base.Module, int32, int32, float64) int32)(m, v119, v120, base.F64_sub(base.F64_mul(v22, base.F64_convert_i64_s(v121-int64(1))), base.F64_convert_i64_s(v76)))
																		mBase = m.M
																		v129 = m.ExcPending
																		if v129 != 0 {
																			return int32(0)
																		} else {
																			v137 = v128
																			m.G0 = v13 + int32(32)
																			return v137
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
func F_percentile_cont_interval_final(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_percentile_cont_final_common(m, l0, int32(1472))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_percentile_cont_multi_final_common(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int64
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int64
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
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
	var v80 int32
	_ = v80
	var v81 int64
	_ = v81
	var v84 int32
	_ = v84
	var v103 int32
	_ = v103
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v119 int64
	_ = v119
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v164 int64
	_ = v164
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int64
	_ = v171
	var v172 int64
	_ = v172
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v219 int32
	_ = v219
	var v220 int64
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v225 int64
	_ = v225
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v244 int64
	_ = v244
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v251 float64
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v294 int32
	_ = v294
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v334 int32
	_ = v334
	v5 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(32)
	m.G0 = v19
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v5
	*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = v5
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v25 == int32(1) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L13
	} else {
		goto L73
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L13
	} else {
		goto L70
	}
L3:
	;
	m.G0 = v19 + int32(32)
	return v294
L4:
	;
	v28 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v28)
	v294 = v5
	goto L3
L5:
	;
	goto L6
L6:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v31 = *(*int64)(unsafe.Add(mBase, uint32(v30)+16))
	if v31 == int64(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v34 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v34)
	v294 = v5
	goto L3
L8:
	;
	goto L9
L9:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v36 == int32(1) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v39 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v39)
	v294 = v5
	goto L3
L11:
	;
	goto L12
L12:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v42 = F_pg_detoast_datum(m, v41)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return int32(0)
L14:
	;
	F_deconstruct_array_builtin(m, v42, int32(701), v19+int32(28), v19+int32(24), v19+int32(20))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	if v55 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+52))
	v60 = F_construct_empty_array(m, v59)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L13
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v19)+28))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v19)+24))
	v64 = *(*int64)(unsafe.Add(mBase, uint32(v30)+16))
	v66 = F_setup_pct_info(m, v55, v62, v63, v64, int32(1))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L13
	} else {
		goto L20
	}
L19:
	;
	v294 = v60
	goto L3
L20:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	v71 = F_palloc(m, v68<<(uint(int32(2))%32))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L13
	} else {
		goto L21
	}
L21:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	v74 = F_palloc(m, v73)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L13
	} else {
		goto L22
	}
L22:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	if v76 <= int32(0) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	if v127 <= v122 {
		goto L32
	} else {
		goto L33
	}
L24:
	;
	v122 = int32(0)
	v127 = v76
	goto L23
L25:
	;
	goto L26
L26:
	;
	v80 = int32(0)
	v81 = *(*int64)(unsafe.Add(mBase, uint32(v66)))
	if int64(0) < v81 {
		v122 = v80
		v127 = v76
		goto L23
	} else {
		goto L27
	}
L27:
	;
	v84 = v80
	goto L28
L28:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v66+v84<<(uint(int32(5))%32))+24))
	*(*int32)(unsafe.Add(mBase, uint32(v71+v103<<(uint(int32(2))%32)))) = int32(0)
	v110 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v103+v74))) = uint8(v110)
	v113 = v84 + v110
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	if v114 <= v113 {
		v122 = v113
		v127 = v114
		goto L23
	} else {
		goto L30
	}
L29:
	;
	v122 = v113
	v127 = v114
	goto L23
L30:
	;
	v119 = *(*int64)(unsafe.Add(mBase, uint32(v66+v113<<(uint(int32(5))%32))))
	if v119 <= int64(0) {
		v84 = v113
		goto L28
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	v281 = v42 + int32(16)
	v287 = F_construct_md_array(m, v71, v74, v279, v281, v281+v279<<(uint(int32(2))%32), l1, l2, int32(0), int32(100))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L13
	} else {
		goto L69
	}
L33:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+24)))
	if v140 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	if v149 <= v122 {
		goto L32
	} else {
		goto L40
	}
L35:
	;
	F_tuplesort_performsort(m, v139)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L13
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	F_tuplesort_rescan(m, v139)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L13
	} else {
		goto L39
	}
L38:
	;
	v145 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v30)+24)) = uint8(v145)
	goto L34
L39:
	;
	goto L34
L40:
	;
	v151 = v122
	v164 = int64(0)
	goto L41
L41:
	;
	v169 = v66 + v151<<(uint(int32(5))%32)
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v169)+24))
	v171 = *(*int64)(unsafe.Add(mBase, uint32(v169)+8))
	v172 = *(*int64)(unsafe.Add(mBase, uint32(v169)))
	if v164 < v172 {
		goto L45
	} else {
		goto L46
	}
L42:
	;
	goto L32
L43:
	;
	if v225 < v171 {
		goto L58
	} else {
		goto L59
	}
L44:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v219)))
	*(*int32)(unsafe.Add(mBase, uint32(v221))) = v222
	v225 = v220
	goto L43
L45:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
	v178 = F_tuplesort_skiptuples(m, v174, v172+(v164^int64(-1)))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L13
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	if v164 != v172 {
		v225 = v164
		goto L43
	} else {
		goto L57
	}
L48:
	;
	if v178 == int32(0) {
		goto L2
	} else {
		goto L49
	}
L49:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
	v183 = int32(1)
	v190 = F_tuplesort_getdatum(m, v182, v183, v183, v19+int32(16), v19+int32(11), int32(0))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L13
	} else {
		goto L51
	}
L50:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L13
	} else {
		goto L54
	}
L51:
	;
	if v190 == int32(0) {
		goto L50
	} else {
		goto L52
	}
L52:
	;
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+11)))
	if v194 == int32(1) {
		goto L50
	} else {
		goto L53
	}
L53:
	;
	v219 = v19 + int32(16)
	v220 = v172
	v221 = v19 + int32(12)
	goto L44
L54:
	;
	F_errmsg_internal(m, int32(84748), int32(0))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L13
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(477321), int32(952), int32(236685))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L13
	} else {
		goto L56
	}
L56:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L57:
	;
	v219 = v19 + int32(12)
	v220 = v164
	v221 = v19 + int32(16)
	goto L44
L58:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
	v228 = int32(1)
	v235 = F_tuplesort_getdatum(m, v227, v228, v228, v19+int32(12), v19+int32(11), int32(0))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L13
	} else {
		goto L61
	}
L59:
	;
	v244 = v225
	goto L60
L60:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
	if v172 < v171 {
		goto L64
	} else {
		goto L65
	}
L61:
	;
	if v235 == int32(0) {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+11)))
	if v239 == int32(1) {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	v244 = v225 + int64(1)
	goto L60
L64:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	v251 = *(*float64)(unsafe.Add(mBase, uint32(v169)+16))
	v252 = m.T0[l3].(func(*base.Module, int32, int32, float64) int32)(m, v245, v250, v251)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L13
	} else {
		goto L67
	}
L65:
	;
	v254 = v245
	goto L66
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v71+v170<<(uint(int32(2))%32)))) = v254
	v257 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v74+v170))) = uint8(v257)
	v260 = v151 + int32(1)
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	if v260 < v261 {
		v151 = v260
		v164 = v244
		goto L41
	} else {
		goto L68
	}
L67:
	;
	v254 = v252
	goto L66
L68:
	;
	goto L42
L69:
	;
	v294 = v287
	goto L3
L70:
	;
	F_errmsg_internal(m, int32(84748), int32(0))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L13
	} else {
		goto L71
	}
L71:
	;
	F_errfinish(m, int32(477321), int32(948), int32(236685))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L13
	} else {
		goto L72
	}
L72:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L73:
	;
	F_errmsg_internal(m, int32(84748), int32(0))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L13
	} else {
		goto L74
	}
L74:
	;
	F_errfinish(m, int32(477321), int32(973), int32(236685))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L13
	} else {
		goto L75
	}
L75:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
