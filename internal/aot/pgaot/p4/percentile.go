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
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int64
	_ = v37
	var v40 int32
	_ = v40
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
	var v55 int64
	_ = v55
	var v59 int64
	_ = v59
	var v62 int32
	_ = v62
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
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v11 == int32(1) {
		v14 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v14)
		v92 = int32(0)
		m.G0 = v9 + int32(16)
		return v92
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v18 = *(*float64)(unsafe.Add(mBase, uint32(v17)))
		if base.F64_lt(v18, float64(0))|base.F64_gt(v18, float64(1))|base.B2i32(base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v18)&int64(9223372036854775807))) != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v100 = m.ExcPending
			if v100 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50331778))
				mBase = m.M
				v103 = m.ExcPending
				if v103 != 0 {
					return int32(0)
				} else {
					*(*float64)(unsafe.Add(mBase, uint32(v9))) = v18
					F_errmsg(m, int32(_a_F_percentile_disc_final_0), v9)
					mBase = m.M
					v107 = m.ExcPending
					if v107 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_percentile_disc_final_1), int32(447), int32(_a_F_percentile_disc_final_2))
						mBase = m.M
						v112 = m.ExcPending
						if v112 != 0 {
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
			v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
			if v30 == int32(1) {
				v33 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v33)
				v92 = int32(0)
				m.G0 = v9 + int32(16)
				return v92
			} else {
				v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v37 = *(*int64)(unsafe.Add(mBase, uint32(v36)+16))
				if v37 == int64(0) {
					v40 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v40)
					v92 = int32(0)
					m.G0 = v9 + int32(16)
					return v92
				} else {
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v36)+8))
					v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+24)))
					if v44 == int32(0) {
						F_tuplesort_performsort(m, v43)
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return int32(0)
						} else {
							v51 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v36)+24)) = uint8(v51)
							v55 = *(*int64)(unsafe.Add(mBase, uint32(v36)+16))
							v59 = base.I64_trunc_sat_f64_s(base.F64_ceil(base.F64_mul(v18, base.F64_convert_i64_s(v55))))
							if int64(2) <= v59 {
								v62 = *(*int32)(unsafe.Add(mBase, uint32(v36)+8))
								v65 = F_tuplesort_skiptuples(m, v62, v59-int64(1))
								mBase = m.M
								v66 = m.ExcPending
								if v66 != 0 {
									return int32(0)
								} else {
									if v65 == int32(0) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v116 = m.ExcPending
										if v116 != 0 {
											return int32(0)
										} else {
											F_errmsg_internal(m, int32(_a_F_percentile_disc_final_3), int32(0))
											mBase = m.M
											v120 = m.ExcPending
											if v120 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_percentile_disc_final_1), int32(480), int32(_a_F_percentile_disc_final_2))
												mBase = m.M
												v125 = m.ExcPending
												if v125 != 0 {
													return int32(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									} else {
										v69 = *(*int32)(unsafe.Add(mBase, uint32(v36)+8))
										v70 = int32(1)
										v77 = F_tuplesort_getdatum(m, v69, v70, v70, v9+int32(12), v9+int32(11), int32(0))
										mBase = m.M
										v78 = m.ExcPending
										if v78 != 0 {
											return int32(0)
										} else {
											if v77 == int32(0) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v129 = m.ExcPending
												if v129 != 0 {
													return int32(0)
												} else {
													F_errmsg_internal(m, int32(_a_F_percentile_disc_final_3), int32(0))
													mBase = m.M
													v133 = m.ExcPending
													if v133 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_percentile_disc_final_1), int32(485), int32(_a_F_percentile_disc_final_2))
														mBase = m.M
														v138 = m.ExcPending
														if v138 != 0 {
															return int32(0)
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											} else {
												v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+11)))
												if v81 == int32(1) {
													v84 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v84)
													v92 = int32(0)
												} else {
													v87 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
													v92 = v87
												}
												m.G0 = v9 + int32(16)
												return v92
											}
										}
									}
								}
							} else {
								v69 = *(*int32)(unsafe.Add(mBase, uint32(v36)+8))
								v70 = int32(1)
								v77 = F_tuplesort_getdatum(m, v69, v70, v70, v9+int32(12), v9+int32(11), int32(0))
								mBase = m.M
								v78 = m.ExcPending
								if v78 != 0 {
									return int32(0)
								} else {
									if v77 == int32(0) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v129 = m.ExcPending
										if v129 != 0 {
											return int32(0)
										} else {
											F_errmsg_internal(m, int32(_a_F_percentile_disc_final_3), int32(0))
											mBase = m.M
											v133 = m.ExcPending
											if v133 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_percentile_disc_final_1), int32(485), int32(_a_F_percentile_disc_final_2))
												mBase = m.M
												v138 = m.ExcPending
												if v138 != 0 {
													return int32(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									} else {
										v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+11)))
										if v81 == int32(1) {
											v84 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v84)
											v92 = int32(0)
										} else {
											v87 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
											v92 = v87
										}
										m.G0 = v9 + int32(16)
										return v92
									}
								}
							}
						}
					} else {
						F_tuplesort_rescan(m, v43)
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return int32(0)
						} else {
							v55 = *(*int64)(unsafe.Add(mBase, uint32(v36)+16))
							v59 = base.I64_trunc_sat_f64_s(base.F64_ceil(base.F64_mul(v18, base.F64_convert_i64_s(v55))))
							if int64(2) <= v59 {
								v62 = *(*int32)(unsafe.Add(mBase, uint32(v36)+8))
								v65 = F_tuplesort_skiptuples(m, v62, v59-int64(1))
								mBase = m.M
								v66 = m.ExcPending
								if v66 != 0 {
									return int32(0)
								} else {
									if v65 == int32(0) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v116 = m.ExcPending
										if v116 != 0 {
											return int32(0)
										} else {
											F_errmsg_internal(m, int32(_a_F_percentile_disc_final_3), int32(0))
											mBase = m.M
											v120 = m.ExcPending
											if v120 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_percentile_disc_final_1), int32(480), int32(_a_F_percentile_disc_final_2))
												mBase = m.M
												v125 = m.ExcPending
												if v125 != 0 {
													return int32(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									} else {
										v69 = *(*int32)(unsafe.Add(mBase, uint32(v36)+8))
										v70 = int32(1)
										v77 = F_tuplesort_getdatum(m, v69, v70, v70, v9+int32(12), v9+int32(11), int32(0))
										mBase = m.M
										v78 = m.ExcPending
										if v78 != 0 {
											return int32(0)
										} else {
											if v77 == int32(0) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v129 = m.ExcPending
												if v129 != 0 {
													return int32(0)
												} else {
													F_errmsg_internal(m, int32(_a_F_percentile_disc_final_3), int32(0))
													mBase = m.M
													v133 = m.ExcPending
													if v133 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_percentile_disc_final_1), int32(485), int32(_a_F_percentile_disc_final_2))
														mBase = m.M
														v138 = m.ExcPending
														if v138 != 0 {
															return int32(0)
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											} else {
												v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+11)))
												if v81 == int32(1) {
													v84 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v84)
													v92 = int32(0)
												} else {
													v87 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
													v92 = v87
												}
												m.G0 = v9 + int32(16)
												return v92
											}
										}
									}
								}
							} else {
								v69 = *(*int32)(unsafe.Add(mBase, uint32(v36)+8))
								v70 = int32(1)
								v77 = F_tuplesort_getdatum(m, v69, v70, v70, v9+int32(12), v9+int32(11), int32(0))
								mBase = m.M
								v78 = m.ExcPending
								if v78 != 0 {
									return int32(0)
								} else {
									if v77 == int32(0) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v129 = m.ExcPending
										if v129 != 0 {
											return int32(0)
										} else {
											F_errmsg_internal(m, int32(_a_F_percentile_disc_final_3), int32(0))
											mBase = m.M
											v133 = m.ExcPending
											if v133 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_percentile_disc_final_1), int32(485), int32(_a_F_percentile_disc_final_2))
												mBase = m.M
												v138 = m.ExcPending
												if v138 != 0 {
													return int32(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									} else {
										v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+11)))
										if v81 == int32(1) {
											v84 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v84)
											v92 = int32(0)
										} else {
											v87 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
											v92 = v87
										}
										m.G0 = v9 + int32(16)
										return v92
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
	var v67 int32
	_ = v67
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
	var v75 int32
	_ = v75
	var v78 int64
	_ = v78
	var v81 int32
	_ = v81
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v115 int64
	_ = v115
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v156 int32
	_ = v156
	var v157 int64
	_ = v157
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int64
	_ = v163
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
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
	var v188 int32
	_ = v188
	var v189 int64
	_ = v189
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v268 int32
	_ = v268
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
	v259 = m.ExcPending
	if v259 != 0 {
		goto L13
	} else {
		goto L54
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L13
	} else {
		goto L51
	}
L3:
	;
	m.G0 = v16 + int32(32)
	return v228
L4:
	;
	v25 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v25)
	v228 = v2
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
	v228 = v2
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
	v228 = v2
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
	v228 = v57
	goto L3
L20:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v16)+20))
	v70 = F_palloc(m, v67<<(uint(int32(2))%32))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L13
	} else {
		goto L21
	}
L21:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v16)+20))
	v73 = F_palloc(m, v72)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L13
	} else {
		goto L22
	}
L22:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v16)+20))
	if v75 <= int32(0) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	v215 = v39 + int32(16)
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v219)+52))
	v221 = int32(*(*int16)(unsafe.Add(mBase, uint32(v219)+56)))
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219)+58)))
	v223 = int32(*(*int8)(unsafe.Add(mBase, uint32(v219)+59)))
	v224 = F_construct_md_array(m, v70, v73, v213, v215, v215+v213<<(uint(int32(2))%32), v220, v221, v222, v223)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L13
	} else {
		goto L50
	}
L24:
	;
	v78 = *(*int64)(unsafe.Add(mBase, uint32(v64)))
	if v78 <= int64(0) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v81 = v59
	goto L28
L26:
	;
	v118 = v59
	v121 = int32(1)
	goto L27
L27:
	;
	if v121 == int32(0) {
		goto L23
	} else {
		goto L32
	}
L28:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v64+v81<<(uint(int32(5))%32))+24))
	v101 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v70+v97<<(uint(int32(2))%32)))) = v101
	v104 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v97+v73))) = uint8(v104)
	v107 = v81 + v104
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v16)+20))
	v109 = base.B2i32(v107 < v108)
	if v109 == v101 {
		goto L23
	} else {
		goto L30
	}
L29:
	;
	v118 = v107
	v121 = v109
	goto L27
L30:
	;
	v115 = *(*int64)(unsafe.Add(mBase, uint32(v64+v107<<(uint(int32(5))%32))))
	if v115 <= int64(0) {
		v81 = v107
		goto L28
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+24)))
	if v134 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v16)+20))
	if v143 <= v118 {
		goto L23
	} else {
		goto L39
	}
L34:
	;
	F_tuplesort_performsort(m, v133)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L13
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	F_tuplesort_rescan(m, v133)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L13
	} else {
		goto L38
	}
L37:
	;
	v139 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+24)) = uint8(v139)
	goto L33
L38:
	;
	goto L33
L39:
	;
	v146 = v118
	v149 = int32(1)
	v156 = v2
	v157 = int64(0)
	goto L40
L40:
	;
	v161 = v64 + v146<<(uint(int32(5))%32)
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v161)+24))
	v163 = *(*int64)(unsafe.Add(mBase, uint32(v161)))
	if v157 < v163 {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	goto L23
L42:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	v169 = F_tuplesort_skiptuples(m, v165, v163+(v157^int64(-1)))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L13
	} else {
		goto L45
	}
L43:
	;
	v187 = v149
	v188 = v156
	v189 = v157
	goto L44
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v70+v162<<(uint(int32(2))%32)))) = v188
	*(*uint8)(unsafe.Add(mBase, uint32(v73+v162))) = uint8(v187)
	v197 = v146 + int32(1)
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v16)+20))
	if v197 < v198 {
		v146 = v197
		v149 = v187
		v156 = v188
		v157 = v189
		goto L40
	} else {
		goto L49
	}
L45:
	;
	if v169 == int32(0) {
		goto L2
	} else {
		goto L46
	}
L46:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	v174 = int32(1)
	v181 = F_tuplesort_getdatum(m, v173, v174, v174, v16+int32(16), v16+int32(15), int32(0))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L13
	} else {
		goto L47
	}
L47:
	;
	if v181 == int32(0) {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+15)))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
	v187 = v185
	v188 = v186
	v189 = v163
	goto L44
L49:
	;
	goto L41
L50:
	;
	v228 = v224
	goto L3
L51:
	;
	F_errmsg_internal(m, int32(_a_F_percentile_disc_multi_final_0), int32(0))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L13
	} else {
		goto L52
	}
L52:
	;
	F_errfinish(m, int32(_a_F_percentile_disc_multi_final_1), int32(819), int32(_a_F_percentile_disc_multi_final_2))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L13
	} else {
		goto L53
	}
L53:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L54:
	;
	F_errmsg_internal(m, int32(_a_F_percentile_disc_multi_final_0), int32(0))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L13
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(_a_F_percentile_disc_multi_final_1), int32(823), int32(_a_F_percentile_disc_multi_final_2))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L13
	} else {
		goto L56
	}
L56:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
