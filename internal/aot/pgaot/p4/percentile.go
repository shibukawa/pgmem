package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_percentile_disc_final(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 float64
	_ = v18
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int64
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int64
	_ = v53
	var v56 float64
	_ = v56
	var v60 int64
	_ = v60
	var v62 int64
	_ = v62
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v11 == int32(1) {
		v14 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v14)
		v95 = int32(0)
		m.G0 = v9 + int32(16)
		return v95
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v18 = *(*float64)(unsafe.Add(mBase, uint32(v17)))
		if base.F64_lt(v18, float64(0)) != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v103 = m.ExcPending
			if v103 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50331778))
				mBase = m.M
				v106 = m.ExcPending
				if v106 != 0 {
					return int32(0)
				} else {
					*(*float64)(unsafe.Add(mBase, uint32(v9))) = v18
					F_errmsg(m, int32(544717), v9)
					mBase = m.M
					v110 = m.ExcPending
					if v110 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(485213), int32(447), int32(308026))
						mBase = m.M
						v115 = m.ExcPending
						if v115 != 0 {
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
			if base.F64_gt(v18, float64(1)) != 0 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v103 = m.ExcPending
				if v103 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(50331778))
					mBase = m.M
					v106 = m.ExcPending
					if v106 != 0 {
						return int32(0)
					} else {
						*(*float64)(unsafe.Add(mBase, uint32(v9))) = v18
						F_errmsg(m, int32(544717), v9)
						mBase = m.M
						v110 = m.ExcPending
						if v110 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(485213), int32(447), int32(308026))
							mBase = m.M
							v115 = m.ExcPending
							if v115 != 0 {
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
				if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v18)&int64(9223372036854775807)) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v103 = m.ExcPending
					if v103 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(50331778))
						mBase = m.M
						v106 = m.ExcPending
						if v106 != 0 {
							return int32(0)
						} else {
							*(*float64)(unsafe.Add(mBase, uint32(v9))) = v18
							F_errmsg(m, int32(544717), v9)
							mBase = m.M
							v110 = m.ExcPending
							if v110 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(485213), int32(447), int32(308026))
								mBase = m.M
								v115 = m.ExcPending
								if v115 != 0 {
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
					v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
					if v28 == int32(1) {
						v31 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v31)
						v95 = int32(0)
						m.G0 = v9 + int32(16)
						return v95
					} else {
						v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						v35 = *(*int64)(unsafe.Add(mBase, uint32(v34)+16))
						if v35 == int64(0) {
							v38 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v38)
							v95 = int32(0)
							m.G0 = v9 + int32(16)
							return v95
						} else {
							v41 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
							v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+24)))
							if v42 == int32(0) {
								F_tuplesort_performsort(m, v41)
								mBase = m.M
								v48 = m.ExcPending
								if v48 != 0 {
									return int32(0)
								} else {
									v49 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v34)+24)) = uint8(v49)
									v53 = *(*int64)(unsafe.Add(mBase, uint32(v34)+16))
									v56 = base.F64_ceil(base.F64_mul(v18, base.F64_convert_i64_s(v53)))
									if base.F64_lt(base.F64_abs(v56), float64(9.223372036854776e+18)) != 0 {
										v60 = base.I64_trunc_f64_s(v56)
										v62 = v60
									} else {
										v62 = int64(-9223372036854775807 - 1)
									}
									if int64(2) <= v62 {
										v65 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
										v68 = F_tuplesort_skiptuples(m, v65, v62-int64(1))
										mBase = m.M
										v69 = m.ExcPending
										if v69 != 0 {
											return int32(0)
										} else {
											if v68 == int32(0) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v119 = m.ExcPending
												if v119 != 0 {
													return int32(0)
												} else {
													F_errmsg_internal(m, int32(479614), int32(0))
													mBase = m.M
													v123 = m.ExcPending
													if v123 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(485213), int32(480), int32(308026))
														mBase = m.M
														v128 = m.ExcPending
														if v128 != 0 {
															return int32(0)
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											} else {
												v72 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
												v73 = int32(1)
												v80 = F_tuplesort_getdatum(m, v72, v73, v73, v9+int32(12), v9+int32(11), int32(0))
												mBase = m.M
												v81 = m.ExcPending
												if v81 != 0 {
													return int32(0)
												} else {
													if v80 == int32(0) {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v132 = m.ExcPending
														if v132 != 0 {
															return int32(0)
														} else {
															F_errmsg_internal(m, int32(479614), int32(0))
															mBase = m.M
															v136 = m.ExcPending
															if v136 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(485213), int32(485), int32(308026))
																mBase = m.M
																v141 = m.ExcPending
																if v141 != 0 {
																	return int32(0)
																} else {
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															}
														}
													} else {
														v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+11)))
														if v84 == int32(1) {
															v87 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v87)
															v95 = int32(0)
														} else {
															v90 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
															v95 = v90
														}
														m.G0 = v9 + int32(16)
														return v95
													}
												}
											}
										}
									} else {
										v72 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
										v73 = int32(1)
										v80 = F_tuplesort_getdatum(m, v72, v73, v73, v9+int32(12), v9+int32(11), int32(0))
										mBase = m.M
										v81 = m.ExcPending
										if v81 != 0 {
											return int32(0)
										} else {
											if v80 == int32(0) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v132 = m.ExcPending
												if v132 != 0 {
													return int32(0)
												} else {
													F_errmsg_internal(m, int32(479614), int32(0))
													mBase = m.M
													v136 = m.ExcPending
													if v136 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(485213), int32(485), int32(308026))
														mBase = m.M
														v141 = m.ExcPending
														if v141 != 0 {
															return int32(0)
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											} else {
												v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+11)))
												if v84 == int32(1) {
													v87 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v87)
													v95 = int32(0)
												} else {
													v90 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
													v95 = v90
												}
												m.G0 = v9 + int32(16)
												return v95
											}
										}
									}
								}
							} else {
								F_tuplesort_rescan(m, v41)
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
									return int32(0)
								} else {
									v53 = *(*int64)(unsafe.Add(mBase, uint32(v34)+16))
									v56 = base.F64_ceil(base.F64_mul(v18, base.F64_convert_i64_s(v53)))
									if base.F64_lt(base.F64_abs(v56), float64(9.223372036854776e+18)) != 0 {
										v60 = base.I64_trunc_f64_s(v56)
										v62 = v60
									} else {
										v62 = int64(-9223372036854775807 - 1)
									}
									if int64(2) <= v62 {
										v65 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
										v68 = F_tuplesort_skiptuples(m, v65, v62-int64(1))
										mBase = m.M
										v69 = m.ExcPending
										if v69 != 0 {
											return int32(0)
										} else {
											if v68 == int32(0) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v119 = m.ExcPending
												if v119 != 0 {
													return int32(0)
												} else {
													F_errmsg_internal(m, int32(479614), int32(0))
													mBase = m.M
													v123 = m.ExcPending
													if v123 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(485213), int32(480), int32(308026))
														mBase = m.M
														v128 = m.ExcPending
														if v128 != 0 {
															return int32(0)
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											} else {
												v72 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
												v73 = int32(1)
												v80 = F_tuplesort_getdatum(m, v72, v73, v73, v9+int32(12), v9+int32(11), int32(0))
												mBase = m.M
												v81 = m.ExcPending
												if v81 != 0 {
													return int32(0)
												} else {
													if v80 == int32(0) {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v132 = m.ExcPending
														if v132 != 0 {
															return int32(0)
														} else {
															F_errmsg_internal(m, int32(479614), int32(0))
															mBase = m.M
															v136 = m.ExcPending
															if v136 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(485213), int32(485), int32(308026))
																mBase = m.M
																v141 = m.ExcPending
																if v141 != 0 {
																	return int32(0)
																} else {
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															}
														}
													} else {
														v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+11)))
														if v84 == int32(1) {
															v87 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v87)
															v95 = int32(0)
														} else {
															v90 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
															v95 = v90
														}
														m.G0 = v9 + int32(16)
														return v95
													}
												}
											}
										}
									} else {
										v72 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
										v73 = int32(1)
										v80 = F_tuplesort_getdatum(m, v72, v73, v73, v9+int32(12), v9+int32(11), int32(0))
										mBase = m.M
										v81 = m.ExcPending
										if v81 != 0 {
											return int32(0)
										} else {
											if v80 == int32(0) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v132 = m.ExcPending
												if v132 != 0 {
													return int32(0)
												} else {
													F_errmsg_internal(m, int32(479614), int32(0))
													mBase = m.M
													v136 = m.ExcPending
													if v136 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(485213), int32(485), int32(308026))
														mBase = m.M
														v141 = m.ExcPending
														if v141 != 0 {
															return int32(0)
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											} else {
												v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+11)))
												if v84 == int32(1) {
													v87 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v87)
													v95 = int32(0)
												} else {
													v90 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
													v95 = v90
												}
												m.G0 = v9 + int32(16)
												return v95
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
func F_percentile_disc_multi_final(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int64
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
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
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int64
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int64
	_ = v77
	var v80 int32
	_ = v80
	var v96 int32
	_ = v96
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v112 int64
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int64
	_ = v153
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int64
	_ = v159
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int64
	_ = v185
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	v2 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(32)
	m.G0 = v16
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v2
	v20 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+15)) = uint8(v20)
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v22 == v20 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L13
	} else {
		goto L53
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L13
	} else {
		goto L50
	}
L3:
	;
	m.G0 = v16 + int32(32)
	return v223
L4:
	;
	v25 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v25)
	v223 = v2
	goto L3
L5:
	;
	goto L6
L6:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v28 = *(*int64)(unsafe.Add(mBase, uint32(v27)+16))
	if v28 == int64(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v31 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v31)
	v223 = v2
	goto L3
L8:
	;
	goto L9
L9:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v33 == int32(1) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v36 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v36)
	v223 = v2
	goto L3
L11:
	;
	goto L12
L12:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v39 = F_pg_detoast_datum(m, v38)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return int32(0)
L14:
	;
	F_deconstruct_array_builtin(m, v39, int32(701), v16+int32(28), v16+int32(24), v16+int32(20))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v16)+20))
	if v52 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+52))
	v57 = F_construct_empty_array(m, v56)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L13
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v59 = int32(0)
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v16)+28))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
	v62 = *(*int64)(unsafe.Add(mBase, uint32(v27)+16))
	v64 = F_setup_pct_info(m, v52, v60, v61, v62, v59)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L13
	} else {
		goto L20
	}
L19:
	;
	v223 = v57
	goto L3
L20:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v16)+20))
	v69 = F_palloc(m, v66<<(uint(int32(2))%32))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L13
	} else {
		goto L21
	}
L21:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v16)+20))
	v72 = F_palloc(m, v71)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L13
	} else {
		goto L22
	}
L22:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v16)+20))
	if v74 <= int32(0) {
		v115 = v59
		v116 = v74
		goto L23
	} else {
		goto L24
	}
L23:
	;
	if v116 <= v115 {
		goto L30
	} else {
		goto L31
	}
L24:
	;
	v77 = *(*int64)(unsafe.Add(mBase, uint32(v64)))
	if int64(0) < v77 {
		v115 = v59
		v116 = v74
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v80 = v59
	goto L26
L26:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v64+v80<<(uint(int32(5))%32))+24))
	*(*int32)(unsafe.Add(mBase, uint32(v69+v96<<(uint(int32(2))%32)))) = int32(0)
	v103 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v96+v72))) = uint8(v103)
	v106 = v80 + v103
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v16)+20))
	if v107 <= v106 {
		v115 = v106
		v116 = v107
		goto L23
	} else {
		goto L28
	}
L27:
	;
	v115 = v106
	v116 = v107
	goto L23
L28:
	;
	v112 = *(*int64)(unsafe.Add(mBase, uint32(v64+v106<<(uint(int32(5))%32))))
	if v112 <= int64(0) {
		v80 = v106
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	v211 = v39 + int32(16)
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v215)+52))
	v217 = int32(*(*int16)(unsafe.Add(mBase, uint32(v215)+56)))
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215)+58)))
	v219 = int32(*(*int8)(unsafe.Add(mBase, uint32(v215)+59)))
	v220 = F_construct_md_array(m, v69, v72, v209, v211, v211+v209<<(uint(int32(2))%32), v216, v217, v218, v219)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L13
	} else {
		goto L49
	}
L31:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+24)))
	if v130 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v16)+20))
	if v139 <= v115 {
		goto L30
	} else {
		goto L38
	}
L33:
	;
	F_tuplesort_performsort(m, v129)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L13
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	F_tuplesort_rescan(m, v129)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L13
	} else {
		goto L37
	}
L36:
	;
	v135 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+24)) = uint8(v135)
	goto L32
L37:
	;
	goto L32
L38:
	;
	v142 = v115
	v150 = int32(1)
	v151 = v2
	v153 = int64(0)
	goto L39
L39:
	;
	v157 = v64 + v142<<(uint(int32(5))%32)
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)+24))
	v159 = *(*int64)(unsafe.Add(mBase, uint32(v157)))
	if v153 < v159 {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	goto L30
L41:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	v165 = F_tuplesort_skiptuples(m, v161, v159+(v153^int64(-1)))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L13
	} else {
		goto L44
	}
L42:
	;
	v183 = v150
	v184 = v151
	v185 = v153
	goto L43
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v69+v158<<(uint(int32(2))%32)))) = v184
	*(*uint8)(unsafe.Add(mBase, uint32(v158+v72))) = uint8(v183)
	v193 = v142 + int32(1)
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v16)+20))
	if v193 < v194 {
		v142 = v193
		v150 = v183
		v151 = v184
		v153 = v185
		goto L39
	} else {
		goto L48
	}
L44:
	;
	if v165 == int32(0) {
		goto L2
	} else {
		goto L45
	}
L45:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	v170 = int32(1)
	v177 = F_tuplesort_getdatum(m, v169, v170, v170, v16+int32(16), v16+int32(15), int32(0))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L13
	} else {
		goto L46
	}
L46:
	;
	if v177 == int32(0) {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+15)))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
	v183 = v181
	v184 = v182
	v185 = v159
	goto L43
L48:
	;
	goto L40
L49:
	;
	v223 = v220
	goto L3
L50:
	;
	F_errmsg_internal(m, int32(479614), int32(0))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L13
	} else {
		goto L51
	}
L51:
	;
	F_errfinish(m, int32(485213), int32(819), int32(307952))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L13
	} else {
		goto L52
	}
L52:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L53:
	;
	F_errmsg_internal(m, int32(479614), int32(0))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L13
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(485213), int32(823), int32(307952))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L13
	} else {
		goto L55
	}
L55:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
