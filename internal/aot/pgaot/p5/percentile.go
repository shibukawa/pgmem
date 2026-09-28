package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_percentile_cont_final_common(m *base.Module, l0 int32, l1 int32) int64 {
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
	var v18 int64
	_ = v18
	var v19 float64
	_ = v19
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int64
	_ = v36
	var v37 int32
	_ = v37
	var v38 int64
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int64
	_ = v56
	var v60 float64
	_ = v60
	var v62 int64
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v88 int64
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v107 int64
	_ = v107
	var v108 int64
	_ = v108
	var v109 int64
	_ = v109
	var v116 int64
	_ = v116
	var v117 int32
	_ = v117
	var v120 int64
	_ = v120
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
	if v13 == int32(1) {
		v16 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v16)
		v120 = int64(0)
		m.G0 = v11 + int32(32)
		return v120
	} else {
		v18 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
		v19 = base.F64_reinterpret_i64(v18)
		if base.F64_lt(v19, float64(0))|base.F64_gt(v19, float64(1))|base.B2i32(base.Ui64(int64(9218868437227405313)) <= base.Ui64(v18&int64(9223372036854775807))) != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v130 = m.ExcPending
			if v130 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(50331778))
				mBase = m.M
				v133 = m.ExcPending
				if v133 != 0 {
					return int64(0)
				} else {
					*(*float64)(unsafe.Add(mBase, uint32(v11))) = v19
					F_errmsg(m, int32(_a_F_percentile_cont_final_common_0), v11)
					mBase = m.M
					v137 = m.ExcPending
					if v137 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_percentile_cont_final_common_1), int32(553), int32(_a_F_percentile_cont_final_common_2))
						mBase = m.M
						v142 = m.ExcPending
						if v142 != 0 {
							return int64(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
			if v30 == int32(1) {
				v33 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v33)
				v120 = int64(0)
				m.G0 = v11 + int32(32)
				return v120
			} else {
				v36 = int64(0)
				v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				v38 = *(*int64)(unsafe.Add(mBase, uint32(v37)+16))
				if v38 == v36 {
					v41 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v41)
					v120 = v36
					m.G0 = v11 + int32(32)
					return v120
				} else {
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
					v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+24)))
					if v44 == int32(0) {
						F_tuplesort_performsort(m, v43)
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return int64(0)
						} else {
							v51 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v37)+24)) = uint8(v51)
							v55 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
							v56 = *(*int64)(unsafe.Add(mBase, uint32(v37)+16))
							v60 = base.F64_mul(v19, base.F64_convert_i64_s(v56-int64(1)))
							v62 = base.I64_trunc_sat_f64_s(base.F64_floor(v60))
							v63 = F_tuplesort_skiptuples(m, v55, v62)
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
								return int64(0)
							} else {
								if v63 == int32(0) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v146 = m.ExcPending
									if v146 != 0 {
										return int64(0)
									} else {
										F_errmsg_internal(m, int32(_a_F_percentile_cont_final_common_3), int32(0))
										mBase = m.M
										v150 = m.ExcPending
										if v150 != 0 {
											return int64(0)
										} else {
											F_errfinish(m, int32(_a_F_percentile_cont_final_common_1), int32(582), int32(_a_F_percentile_cont_final_common_2))
											mBase = m.M
											v155 = m.ExcPending
											if v155 != 0 {
												return int64(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								} else {
									v67 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
									v68 = int32(1)
									v75 = F_tuplesort_getdatum(m, v67, v68, v68, v11+int32(24), v11+int32(15), int32(0))
									mBase = m.M
									v76 = m.ExcPending
									if v76 != 0 {
										return int64(0)
									} else {
										if v75 == int32(0) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v159 = m.ExcPending
											if v159 != 0 {
												return int64(0)
											} else {
												F_errmsg_internal(m, int32(_a_F_percentile_cont_final_common_3), int32(0))
												mBase = m.M
												v163 = m.ExcPending
												if v163 != 0 {
													return int64(0)
												} else {
													F_errfinish(m, int32(_a_F_percentile_cont_final_common_1), int32(586), int32(_a_F_percentile_cont_final_common_2))
													mBase = m.M
													v168 = m.ExcPending
													if v168 != 0 {
														return int64(0)
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											}
										} else {
											v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)))
											if v79 == int32(1) {
												v82 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v82)
												v120 = int64(0)
												m.G0 = v11 + int32(32)
												return v120
											} else {
												if base.I64_trunc_sat_f64_s(base.F64_ceil(v60)) == v62 {
													v88 = *(*int64)(unsafe.Add(mBase, uint32(v11)+24))
													v120 = v88
													m.G0 = v11 + int32(32)
													return v120
												} else {
													v89 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
													v90 = int32(1)
													v97 = F_tuplesort_getdatum(m, v89, v90, v90, v11+int32(16), v11+int32(15), int32(0))
													mBase = m.M
													v98 = m.ExcPending
													if v98 != 0 {
														return int64(0)
													} else {
														if v97 == int32(0) {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v172 = m.ExcPending
															if v172 != 0 {
																return int64(0)
															} else {
																F_errmsg_internal(m, int32(_a_F_percentile_cont_final_common_3), int32(0))
																mBase = m.M
																v176 = m.ExcPending
																if v176 != 0 {
																	return int64(0)
																} else {
																	F_errfinish(m, int32(_a_F_percentile_cont_final_common_1), int32(598), int32(_a_F_percentile_cont_final_common_2))
																	mBase = m.M
																	v181 = m.ExcPending
																	if v181 != 0 {
																		return int64(0)
																	} else {
																		base.Wasm_trap_unreachable()
																		for {
																		}
																	}
																}
															}
														} else {
															v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)))
															if v101 == int32(1) {
																v104 = int32(1)
																*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v104)
																v120 = int64(0)
																m.G0 = v11 + int32(32)
																return v120
															} else {
																v107 = *(*int64)(unsafe.Add(mBase, uint32(v11)+24))
																v108 = *(*int64)(unsafe.Add(mBase, uint32(v11)+16))
																v109 = *(*int64)(unsafe.Add(mBase, uint32(v37)+16))
																v116 = m.T0[l1].(func(*base.Module, int64, int64, float64) int64)(m, v107, v108, base.F64_sub(base.F64_mul(v19, base.F64_convert_i64_s(v109-int64(1))), base.F64_convert_i64_s(v62)))
																mBase = m.M
																v117 = m.ExcPending
																if v117 != 0 {
																	return int64(0)
																} else {
																	v120 = v116
																	m.G0 = v11 + int32(32)
																	return v120
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
						F_tuplesort_rescan(m, v43)
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return int64(0)
						} else {
							v55 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
							v56 = *(*int64)(unsafe.Add(mBase, uint32(v37)+16))
							v60 = base.F64_mul(v19, base.F64_convert_i64_s(v56-int64(1)))
							v62 = base.I64_trunc_sat_f64_s(base.F64_floor(v60))
							v63 = F_tuplesort_skiptuples(m, v55, v62)
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
								return int64(0)
							} else {
								if v63 == int32(0) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v146 = m.ExcPending
									if v146 != 0 {
										return int64(0)
									} else {
										F_errmsg_internal(m, int32(_a_F_percentile_cont_final_common_3), int32(0))
										mBase = m.M
										v150 = m.ExcPending
										if v150 != 0 {
											return int64(0)
										} else {
											F_errfinish(m, int32(_a_F_percentile_cont_final_common_1), int32(582), int32(_a_F_percentile_cont_final_common_2))
											mBase = m.M
											v155 = m.ExcPending
											if v155 != 0 {
												return int64(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								} else {
									v67 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
									v68 = int32(1)
									v75 = F_tuplesort_getdatum(m, v67, v68, v68, v11+int32(24), v11+int32(15), int32(0))
									mBase = m.M
									v76 = m.ExcPending
									if v76 != 0 {
										return int64(0)
									} else {
										if v75 == int32(0) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v159 = m.ExcPending
											if v159 != 0 {
												return int64(0)
											} else {
												F_errmsg_internal(m, int32(_a_F_percentile_cont_final_common_3), int32(0))
												mBase = m.M
												v163 = m.ExcPending
												if v163 != 0 {
													return int64(0)
												} else {
													F_errfinish(m, int32(_a_F_percentile_cont_final_common_1), int32(586), int32(_a_F_percentile_cont_final_common_2))
													mBase = m.M
													v168 = m.ExcPending
													if v168 != 0 {
														return int64(0)
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											}
										} else {
											v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)))
											if v79 == int32(1) {
												v82 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v82)
												v120 = int64(0)
												m.G0 = v11 + int32(32)
												return v120
											} else {
												if base.I64_trunc_sat_f64_s(base.F64_ceil(v60)) == v62 {
													v88 = *(*int64)(unsafe.Add(mBase, uint32(v11)+24))
													v120 = v88
													m.G0 = v11 + int32(32)
													return v120
												} else {
													v89 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
													v90 = int32(1)
													v97 = F_tuplesort_getdatum(m, v89, v90, v90, v11+int32(16), v11+int32(15), int32(0))
													mBase = m.M
													v98 = m.ExcPending
													if v98 != 0 {
														return int64(0)
													} else {
														if v97 == int32(0) {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v172 = m.ExcPending
															if v172 != 0 {
																return int64(0)
															} else {
																F_errmsg_internal(m, int32(_a_F_percentile_cont_final_common_3), int32(0))
																mBase = m.M
																v176 = m.ExcPending
																if v176 != 0 {
																	return int64(0)
																} else {
																	F_errfinish(m, int32(_a_F_percentile_cont_final_common_1), int32(598), int32(_a_F_percentile_cont_final_common_2))
																	mBase = m.M
																	v181 = m.ExcPending
																	if v181 != 0 {
																		return int64(0)
																	} else {
																		base.Wasm_trap_unreachable()
																		for {
																		}
																	}
																}
															}
														} else {
															v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)))
															if v101 == int32(1) {
																v104 = int32(1)
																*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v104)
																v120 = int64(0)
																m.G0 = v11 + int32(32)
																return v120
															} else {
																v107 = *(*int64)(unsafe.Add(mBase, uint32(v11)+24))
																v108 = *(*int64)(unsafe.Add(mBase, uint32(v11)+16))
																v109 = *(*int64)(unsafe.Add(mBase, uint32(v37)+16))
																v116 = m.T0[l1].(func(*base.Module, int64, int64, float64) int64)(m, v107, v108, base.F64_sub(base.F64_mul(v19, base.F64_convert_i64_s(v109-int64(1))), base.F64_convert_i64_s(v62)))
																mBase = m.M
																v117 = m.ExcPending
																if v117 != 0 {
																	return int64(0)
																} else {
																	v120 = v116
																	m.G0 = v11 + int32(32)
																	return v120
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
func F_percentile_cont_interval_final(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_percentile_cont_final_common(m, l0, int32(1603))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_percentile_cont_multi_final_common(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int64 {
	mBase := m.M
	_ = mBase
	var v15 int64
	_ = v15
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int64
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int64
	_ = v70
	var v72 int32
	_ = v72
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
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int64
	_ = v86
	var v89 int32
	_ = v89
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v128 int64
	_ = v128
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v151 int32
	_ = v151
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
	var v163 int32
	_ = v163
	var v177 int64
	_ = v177
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int64
	_ = v185
	var v186 int64
	_ = v186
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v225 int32
	_ = v225
	var v231 int32
	_ = v231
	var v232 int64
	_ = v232
	var v233 int32
	_ = v233
	var v234 int64
	_ = v234
	var v237 int64
	_ = v237
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v256 int64
	_ = v256
	var v257 int64
	_ = v257
	var v262 int64
	_ = v262
	var v263 float64
	_ = v263
	var v264 int64
	_ = v264
	var v265 int32
	_ = v265
	var v266 int64
	_ = v266
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v320 int32
	_ = v320
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v338 int32
	_ = v338
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v351 int32
	_ = v351
	v15 = int64(0)
	v19 = m.G0
	v21 = v19 - int32(48)
	m.G0 = v21
	*(*int64)(unsafe.Add(mBase, uint32(v21)+24)) = v15
	*(*int64)(unsafe.Add(mBase, uint32(v21)+16)) = v15
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v27 == int32(1) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L13
	} else {
		goto L72
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L13
	} else {
		goto L69
	}
L3:
	;
	m.G0 = v21 + int32(48)
	return base.I64_extend_i32_u(v320)
L4:
	;
	v30 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v30)
	v320 = int32(0)
	goto L3
L5:
	;
	goto L6
L6:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v34 = *(*int64)(unsafe.Add(mBase, uint32(v33)+16))
	if v34 == int64(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v37 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v37)
	v320 = int32(0)
	goto L3
L8:
	;
	goto L9
L9:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
	if v40 == int32(1) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v43 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v43)
	v320 = int32(0)
	goto L3
L11:
	;
	goto L12
L12:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v47 = F_pg_detoast_datum(m, v46)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return int64(0)
L14:
	;
	F_deconstruct_array_builtin(m, v47, int32(701), v21+int32(44), v21+int32(40), v21+int32(36))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v21)+36))
	if v60 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+52))
	v65 = F_construct_empty_array(m, v64)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L13
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v67 = int32(1)
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v21)+44))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v21)+40))
	v70 = *(*int64)(unsafe.Add(mBase, uint32(v33)+16))
	v72 = F_setup_pct_info(m, v60, v68, v69, v70, v67)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L13
	} else {
		goto L20
	}
L19:
	;
	v320 = v65
	goto L3
L20:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v21)+36))
	v77 = F_palloc(m, v74<<(uint(int32(3))%32))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L13
	} else {
		goto L21
	}
L21:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v21)+36))
	v80 = F_palloc(m, v79)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L13
	} else {
		goto L22
	}
L22:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v21)+36))
	if v82 <= int32(0) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v295 = v47 + int32(16)
	v300 = F_construct_md_array(m, v77, v80, v293, v295, v295+v293<<(uint(int32(2))%32), l1, l2, l3, int32(100))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L13
	} else {
		goto L68
	}
L24:
	;
	v85 = int32(0)
	v86 = *(*int64)(unsafe.Add(mBase, uint32(v72)))
	if v86 <= int64(0) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v89 = v85
	goto L28
L26:
	;
	v131 = v85
	v137 = v67
	goto L27
L27:
	;
	if v137 == int32(0) {
		goto L23
	} else {
		goto L32
	}
L28:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v72+v89<<(uint(int32(5))%32))+24))
	*(*int64)(unsafe.Add(mBase, uint32(v77+v110<<(uint(int32(3))%32)))) = int64(0)
	v117 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v110+v80))) = uint8(v117)
	v120 = v89 + v117
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v21)+36))
	v122 = base.B2i32(v120 < v121)
	if v122 == int32(0) {
		goto L23
	} else {
		goto L30
	}
L29:
	;
	v131 = v120
	v137 = v122
	goto L27
L30:
	;
	v128 = *(*int64)(unsafe.Add(mBase, uint32(v72+v120<<(uint(int32(5))%32))))
	if v128 <= int64(0) {
		v89 = v120
		goto L28
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+24)))
	if v152 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v21)+36))
	if v161 <= v131 {
		goto L23
	} else {
		goto L39
	}
L34:
	;
	F_tuplesort_performsort(m, v151)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L13
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	F_tuplesort_rescan(m, v151)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L13
	} else {
		goto L38
	}
L37:
	;
	v157 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v33)+24)) = uint8(v157)
	goto L33
L38:
	;
	goto L33
L39:
	;
	v163 = v131
	v177 = v15
	goto L40
L40:
	;
	v183 = v72 + v163<<(uint(int32(5))%32)
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v183)+24))
	v185 = *(*int64)(unsafe.Add(mBase, uint32(v183)+8))
	v186 = *(*int64)(unsafe.Add(mBase, uint32(v183)))
	if v177 < v186 {
		goto L44
	} else {
		goto L45
	}
L41:
	;
	goto L23
L42:
	;
	if v237 < v185 {
		goto L57
	} else {
		goto L58
	}
L43:
	;
	v234 = *(*int64)(unsafe.Add(mBase, uint32(v231)))
	*(*int64)(unsafe.Add(mBase, uint32(v233))) = v234
	v237 = v232
	goto L42
L44:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
	v192 = F_tuplesort_skiptuples(m, v188, v186+(v177^int64(-1)))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L13
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	if v177 != v186 {
		v237 = v177
		goto L42
	} else {
		goto L56
	}
L47:
	;
	if v192 == int32(0) {
		goto L2
	} else {
		goto L48
	}
L48:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
	v197 = int32(1)
	v200 = v21 + int32(24)
	v204 = F_tuplesort_getdatum(m, v196, v197, v197, v200, v21+int32(15), int32(0))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L13
	} else {
		goto L50
	}
L49:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L13
	} else {
		goto L53
	}
L50:
	;
	if v204 == int32(0) {
		goto L49
	} else {
		goto L51
	}
L51:
	;
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+15)))
	if v208&int32(1) != 0 {
		goto L49
	} else {
		goto L52
	}
L52:
	;
	v231 = v200
	v232 = v186
	v233 = v21 + int32(16)
	goto L43
L53:
	;
	F_errmsg_internal(m, int32(_a_F_percentile_cont_multi_final_common_0), int32(0))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L13
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(_a_F_percentile_cont_multi_final_common_1), int32(953), int32(_a_F_percentile_cont_multi_final_common_2))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
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
	v231 = v21 + int32(16)
	v232 = v177
	v233 = v21 + int32(24)
	goto L43
L57:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
	v240 = int32(1)
	v247 = F_tuplesort_getdatum(m, v239, v240, v240, v21+int32(16), v21+int32(15), int32(0))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L13
	} else {
		goto L60
	}
L58:
	;
	v256 = v237
	goto L59
L59:
	;
	v257 = *(*int64)(unsafe.Add(mBase, uint32(v21)+24))
	if v186 < v185 {
		goto L63
	} else {
		goto L64
	}
L60:
	;
	if v247 == int32(0) {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+15)))
	if v251&int32(1) != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	v256 = v237 + int64(1)
	goto L59
L63:
	;
	v262 = *(*int64)(unsafe.Add(mBase, uint32(v21)+16))
	v263 = *(*float64)(unsafe.Add(mBase, uint32(v183)+16))
	v264 = m.T0[l4].(func(*base.Module, int64, int64, float64) int64)(m, v257, v262, v263)
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L13
	} else {
		goto L66
	}
L64:
	;
	v266 = v257
	goto L65
L65:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v77+v184<<(uint(int32(3))%32)))) = v266
	v269 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v80+v184))) = uint8(v269)
	v272 = v163 + int32(1)
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v21)+36))
	if v272 < v273 {
		v163 = v272
		v177 = v256
		goto L40
	} else {
		goto L67
	}
L66:
	;
	v266 = v264
	goto L65
L67:
	;
	goto L41
L68:
	;
	v320 = v300
	goto L3
L69:
	;
	F_errmsg_internal(m, int32(_a_F_percentile_cont_multi_final_common_0), int32(0))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L13
	} else {
		goto L70
	}
L70:
	;
	F_errfinish(m, int32(_a_F_percentile_cont_multi_final_common_1), int32(949), int32(_a_F_percentile_cont_multi_final_common_2))
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
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
	v346 = m.ExcPending
	if v346 != 0 {
		goto L13
	} else {
		goto L73
	}
L73:
	;
	F_errfinish(m, int32(_a_F_percentile_cont_multi_final_common_1), int32(974), int32(_a_F_percentile_cont_multi_final_common_2))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
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
