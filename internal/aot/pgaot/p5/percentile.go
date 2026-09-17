package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_percentile_cont_final_common(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 float64
	_ = v20
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
	var v57 int32
	_ = v57
	var v58 int64
	_ = v58
	var v62 float64
	_ = v62
	var v64 int64
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int64
	_ = v111
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v13 == int32(1) {
		v16 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v16)
		v125 = int32(0)
		m.G0 = v11 + int32(32)
		return v125
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v20 = *(*float64)(unsafe.Add(mBase, uint32(v19)))
		if base.F64_lt(v20, float64(0))|base.F64_gt(v20, float64(1))|base.B2i32(base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v20)&int64(9223372036854775807))) != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v133 = m.ExcPending
			if v133 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50331778))
				mBase = m.M
				v136 = m.ExcPending
				if v136 != 0 {
					return int32(0)
				} else {
					*(*float64)(unsafe.Add(mBase, uint32(v11))) = v20
					F_errmsg(m, int32(_a_F_percentile_cont_final_common_0), v11)
					mBase = m.M
					v140 = m.ExcPending
					if v140 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_percentile_cont_final_common_1), int32(552), int32(_a_F_percentile_cont_final_common_2))
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
			v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
			if v32 == int32(1) {
				v35 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v35)
				v125 = int32(0)
				m.G0 = v11 + int32(32)
				return v125
			} else {
				v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v39 = *(*int64)(unsafe.Add(mBase, uint32(v38)+16))
				if v39 == int64(0) {
					v42 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v42)
					v125 = int32(0)
					m.G0 = v11 + int32(32)
					return v125
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
							v57 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
							v58 = *(*int64)(unsafe.Add(mBase, uint32(v38)+16))
							v62 = base.F64_mul(v20, base.F64_convert_i64_s(v58-int64(1)))
							v64 = base.I64_trunc_sat_f64_s(base.F64_floor(v62))
							v65 = F_tuplesort_skiptuples(m, v57, v64)
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return int32(0)
							} else {
								if v65 == int32(0) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v149 = m.ExcPending
									if v149 != 0 {
										return int32(0)
									} else {
										F_errmsg_internal(m, int32(_a_F_percentile_cont_final_common_3), int32(0))
										mBase = m.M
										v153 = m.ExcPending
										if v153 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_percentile_cont_final_common_1), int32(581), int32(_a_F_percentile_cont_final_common_2))
											mBase = m.M
											v158 = m.ExcPending
											if v158 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								} else {
									v69 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
									v70 = int32(1)
									v77 = F_tuplesort_getdatum(m, v69, v70, v70, v11+int32(28), v11+int32(23), int32(0))
									mBase = m.M
									v78 = m.ExcPending
									if v78 != 0 {
										return int32(0)
									} else {
										if v77 == int32(0) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v162 = m.ExcPending
											if v162 != 0 {
												return int32(0)
											} else {
												F_errmsg_internal(m, int32(_a_F_percentile_cont_final_common_3), int32(0))
												mBase = m.M
												v166 = m.ExcPending
												if v166 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_percentile_cont_final_common_1), int32(585), int32(_a_F_percentile_cont_final_common_2))
													mBase = m.M
													v171 = m.ExcPending
													if v171 != 0 {
														return int32(0)
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											}
										} else {
											v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+23)))
											if v81 == int32(1) {
												v84 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v84)
												v125 = int32(0)
												m.G0 = v11 + int32(32)
												return v125
											} else {
												if base.I64_trunc_sat_f64_s(base.F64_ceil(v62)) == v64 {
													v90 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
													v125 = v90
													m.G0 = v11 + int32(32)
													return v125
												} else {
													v91 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
													v92 = int32(1)
													v99 = F_tuplesort_getdatum(m, v91, v92, v92, v11+int32(24), v11+int32(23), int32(0))
													mBase = m.M
													v100 = m.ExcPending
													if v100 != 0 {
														return int32(0)
													} else {
														if v99 == int32(0) {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v175 = m.ExcPending
															if v175 != 0 {
																return int32(0)
															} else {
																F_errmsg_internal(m, int32(_a_F_percentile_cont_final_common_3), int32(0))
																mBase = m.M
																v179 = m.ExcPending
																if v179 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(_a_F_percentile_cont_final_common_1), int32(597), int32(_a_F_percentile_cont_final_common_2))
																	mBase = m.M
																	v184 = m.ExcPending
																	if v184 != 0 {
																		return int32(0)
																	} else {
																		base.Wasm_trap_unreachable()
																		for {
																		}
																	}
																}
															}
														} else {
															v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+23)))
															if v103 == int32(1) {
																v106 = int32(1)
																*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v106)
																v125 = int32(0)
																m.G0 = v11 + int32(32)
																return v125
															} else {
																v109 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
																v110 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
																v111 = *(*int64)(unsafe.Add(mBase, uint32(v38)+16))
																v118 = m.T0[l1].(func(*base.Module, int32, int32, float64) int32)(m, v109, v110, base.F64_sub(base.F64_mul(v20, base.F64_convert_i64_s(v111-int64(1))), base.F64_convert_i64_s(v64)))
																mBase = m.M
																v119 = m.ExcPending
																if v119 != 0 {
																	return int32(0)
																} else {
																	v125 = v118
																	m.G0 = v11 + int32(32)
																	return v125
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
							v57 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
							v58 = *(*int64)(unsafe.Add(mBase, uint32(v38)+16))
							v62 = base.F64_mul(v20, base.F64_convert_i64_s(v58-int64(1)))
							v64 = base.I64_trunc_sat_f64_s(base.F64_floor(v62))
							v65 = F_tuplesort_skiptuples(m, v57, v64)
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return int32(0)
							} else {
								if v65 == int32(0) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v149 = m.ExcPending
									if v149 != 0 {
										return int32(0)
									} else {
										F_errmsg_internal(m, int32(_a_F_percentile_cont_final_common_3), int32(0))
										mBase = m.M
										v153 = m.ExcPending
										if v153 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_percentile_cont_final_common_1), int32(581), int32(_a_F_percentile_cont_final_common_2))
											mBase = m.M
											v158 = m.ExcPending
											if v158 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								} else {
									v69 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
									v70 = int32(1)
									v77 = F_tuplesort_getdatum(m, v69, v70, v70, v11+int32(28), v11+int32(23), int32(0))
									mBase = m.M
									v78 = m.ExcPending
									if v78 != 0 {
										return int32(0)
									} else {
										if v77 == int32(0) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v162 = m.ExcPending
											if v162 != 0 {
												return int32(0)
											} else {
												F_errmsg_internal(m, int32(_a_F_percentile_cont_final_common_3), int32(0))
												mBase = m.M
												v166 = m.ExcPending
												if v166 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_percentile_cont_final_common_1), int32(585), int32(_a_F_percentile_cont_final_common_2))
													mBase = m.M
													v171 = m.ExcPending
													if v171 != 0 {
														return int32(0)
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											}
										} else {
											v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+23)))
											if v81 == int32(1) {
												v84 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v84)
												v125 = int32(0)
												m.G0 = v11 + int32(32)
												return v125
											} else {
												if base.I64_trunc_sat_f64_s(base.F64_ceil(v62)) == v64 {
													v90 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
													v125 = v90
													m.G0 = v11 + int32(32)
													return v125
												} else {
													v91 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
													v92 = int32(1)
													v99 = F_tuplesort_getdatum(m, v91, v92, v92, v11+int32(24), v11+int32(23), int32(0))
													mBase = m.M
													v100 = m.ExcPending
													if v100 != 0 {
														return int32(0)
													} else {
														if v99 == int32(0) {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v175 = m.ExcPending
															if v175 != 0 {
																return int32(0)
															} else {
																F_errmsg_internal(m, int32(_a_F_percentile_cont_final_common_3), int32(0))
																mBase = m.M
																v179 = m.ExcPending
																if v179 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(_a_F_percentile_cont_final_common_1), int32(597), int32(_a_F_percentile_cont_final_common_2))
																	mBase = m.M
																	v184 = m.ExcPending
																	if v184 != 0 {
																		return int32(0)
																	} else {
																		base.Wasm_trap_unreachable()
																		for {
																		}
																	}
																}
															}
														} else {
															v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+23)))
															if v103 == int32(1) {
																v106 = int32(1)
																*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v106)
																v125 = int32(0)
																m.G0 = v11 + int32(32)
																return v125
															} else {
																v109 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
																v110 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
																v111 = *(*int64)(unsafe.Add(mBase, uint32(v38)+16))
																v118 = m.T0[l1].(func(*base.Module, int32, int32, float64) int32)(m, v109, v110, base.F64_sub(base.F64_mul(v20, base.F64_convert_i64_s(v111-int64(1))), base.F64_convert_i64_s(v64)))
																mBase = m.M
																v119 = m.ExcPending
																if v119 != 0 {
																	return int32(0)
																} else {
																	v125 = v118
																	m.G0 = v11 + int32(32)
																	return v125
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
	v3 = F_percentile_cont_final_common(m, l0, int32(1457))
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
	var v64 int32
	_ = v64
	var v65 int64
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
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
	var v80 int32
	_ = v80
	var v81 int64
	_ = v81
	var v84 int32
	_ = v84
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v121 int64
	_ = v121
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v167 int64
	_ = v167
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int64
	_ = v174
	var v175 int64
	_ = v175
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v220 int32
	_ = v220
	var v221 int64
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v226 int64
	_ = v226
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v245 int64
	_ = v245
	var v246 int32
	_ = v246
	var v251 int32
	_ = v251
	var v252 float64
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v295 int32
	_ = v295
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v322 int32
	_ = v322
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v335 int32
	_ = v335
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
	v326 = m.ExcPending
	if v326 != 0 {
		goto L13
	} else {
		goto L72
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L13
	} else {
		goto L69
	}
L3:
	;
	m.G0 = v19 + int32(32)
	return v295
L4:
	;
	v28 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v28)
	v295 = v5
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
	v295 = v5
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
	v295 = v5
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
	v62 = int32(1)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v19)+28))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v19)+24))
	v65 = *(*int64)(unsafe.Add(mBase, uint32(v30)+16))
	v67 = F_setup_pct_info(m, v55, v63, v64, v65, v62)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L13
	} else {
		goto L20
	}
L19:
	;
	v295 = v60
	goto L3
L20:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	v72 = F_palloc(m, v69<<(uint(int32(2))%32))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L13
	} else {
		goto L21
	}
L21:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	v75 = F_palloc(m, v74)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L13
	} else {
		goto L22
	}
L22:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	if v77 <= int32(0) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	v282 = v42 + int32(16)
	v288 = F_construct_md_array(m, v72, v75, v280, v282, v282+v280<<(uint(int32(2))%32), l1, l2, int32(0), int32(100))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L13
	} else {
		goto L68
	}
L24:
	;
	v80 = int32(0)
	v81 = *(*int64)(unsafe.Add(mBase, uint32(v67)))
	if v81 <= int64(0) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v84 = v80
	goto L28
L26:
	;
	v124 = v80
	v129 = v62
	goto L27
L27:
	;
	if v129 == int32(0) {
		goto L23
	} else {
		goto L32
	}
L28:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v67+v84<<(uint(int32(5))%32))+24))
	v107 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v72+v103<<(uint(int32(2))%32)))) = v107
	v110 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v103+v75))) = uint8(v110)
	v113 = v84 + v110
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	v115 = base.B2i32(v113 < v114)
	if v115 == v107 {
		goto L23
	} else {
		goto L30
	}
L29:
	;
	v124 = v113
	v129 = v115
	goto L27
L30:
	;
	v121 = *(*int64)(unsafe.Add(mBase, uint32(v67+v113<<(uint(int32(5))%32))))
	if v121 <= int64(0) {
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
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+24)))
	if v143 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	if v152 <= v124 {
		goto L23
	} else {
		goto L39
	}
L34:
	;
	F_tuplesort_performsort(m, v142)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L13
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	F_tuplesort_rescan(m, v142)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L13
	} else {
		goto L38
	}
L37:
	;
	v148 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v30)+24)) = uint8(v148)
	goto L33
L38:
	;
	goto L33
L39:
	;
	v154 = v124
	v167 = int64(0)
	goto L40
L40:
	;
	v172 = v67 + v154<<(uint(int32(5))%32)
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v172)+24))
	v174 = *(*int64)(unsafe.Add(mBase, uint32(v172)+8))
	v175 = *(*int64)(unsafe.Add(mBase, uint32(v172)))
	if v167 < v175 {
		goto L44
	} else {
		goto L45
	}
L41:
	;
	goto L23
L42:
	;
	if v226 < v174 {
		goto L57
	} else {
		goto L58
	}
L43:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v220)))
	*(*int32)(unsafe.Add(mBase, uint32(v222))) = v223
	v226 = v221
	goto L42
L44:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
	v181 = F_tuplesort_skiptuples(m, v177, v175+(v167^int64(-1)))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L13
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	if v167 != v175 {
		v226 = v167
		goto L42
	} else {
		goto L56
	}
L47:
	;
	if v181 == int32(0) {
		goto L2
	} else {
		goto L48
	}
L48:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
	v186 = int32(1)
	v189 = v19 + int32(16)
	v193 = F_tuplesort_getdatum(m, v185, v186, v186, v189, v19+int32(11), int32(0))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L13
	} else {
		goto L50
	}
L49:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L13
	} else {
		goto L53
	}
L50:
	;
	if v193 == int32(0) {
		goto L49
	} else {
		goto L51
	}
L51:
	;
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+11)))
	if v197&int32(1) != 0 {
		goto L49
	} else {
		goto L52
	}
L52:
	;
	v220 = v189
	v221 = v175
	v222 = v19 + int32(12)
	goto L43
L53:
	;
	F_errmsg_internal(m, int32(_a_F_percentile_cont_multi_final_common_0), int32(0))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L13
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(_a_F_percentile_cont_multi_final_common_1), int32(952), int32(_a_F_percentile_cont_multi_final_common_2))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L13
	} else {
		goto L55
	}
L55:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L56:
	;
	v220 = v19 + int32(12)
	v221 = v167
	v222 = v19 + int32(16)
	goto L43
L57:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
	v229 = int32(1)
	v236 = F_tuplesort_getdatum(m, v228, v229, v229, v19+int32(12), v19+int32(11), int32(0))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L13
	} else {
		goto L60
	}
L58:
	;
	v245 = v226
	goto L59
L59:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
	if v175 < v174 {
		goto L63
	} else {
		goto L64
	}
L60:
	;
	if v236 == int32(0) {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+11)))
	if v240&int32(1) != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	v245 = v226 + int64(1)
	goto L59
L63:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	v252 = *(*float64)(unsafe.Add(mBase, uint32(v172)+16))
	v253 = m.T0[l3].(func(*base.Module, int32, int32, float64) int32)(m, v246, v251, v252)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L13
	} else {
		goto L66
	}
L64:
	;
	v255 = v246
	goto L65
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v72+v173<<(uint(int32(2))%32)))) = v255
	v258 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v75+v173))) = uint8(v258)
	v261 = v154 + int32(1)
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	if v261 < v262 {
		v154 = v261
		v167 = v245
		goto L40
	} else {
		goto L67
	}
L66:
	;
	v255 = v253
	goto L65
L67:
	;
	goto L41
L68:
	;
	v295 = v288
	goto L3
L69:
	;
	F_errmsg_internal(m, int32(_a_F_percentile_cont_multi_final_common_0), int32(0))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L13
	} else {
		goto L70
	}
L70:
	;
	F_errfinish(m, int32(_a_F_percentile_cont_multi_final_common_1), int32(948), int32(_a_F_percentile_cont_multi_final_common_2))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L13
	} else {
		goto L71
	}
L71:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L72:
	;
	F_errmsg_internal(m, int32(_a_F_percentile_cont_multi_final_common_0), int32(0))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L13
	} else {
		goto L73
	}
L73:
	;
	F_errfinish(m, int32(_a_F_percentile_cont_multi_final_common_1), int32(973), int32(_a_F_percentile_cont_multi_final_common_2))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L13
	} else {
		goto L74
	}
L74:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
