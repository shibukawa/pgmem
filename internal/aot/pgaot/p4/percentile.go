package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_percentile_disc_final(m *base.Module, l0 int32) int64 {
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
	var v16 int64
	_ = v16
	var v17 float64
	_ = v17
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int64
	_ = v34
	var v35 int32
	_ = v35
	var v36 int64
	_ = v36
	var v39 int32
	_ = v39
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
	var v57 int64
	_ = v57
	var v60 int32
	_ = v60
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
	var v85 int64
	_ = v85
	var v86 int64
	_ = v86
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
	if v11 == int32(1) {
		v14 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v14)
		v86 = int64(0)
		m.G0 = v9 + int32(32)
		return v86
	} else {
		v16 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
		v17 = base.F64_reinterpret_i64(v16)
		if base.F64_lt(v17, float64(0))|base.F64_gt(v17, float64(1))|base.B2i32(base.Ui64(int64(9218868437227405313)) <= base.Ui64(v16&int64(9223372036854775807))) != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v97 = m.ExcPending
			if v97 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(50331778))
				mBase = m.M
				v100 = m.ExcPending
				if v100 != 0 {
					return int64(0)
				} else {
					*(*float64)(unsafe.Add(mBase, uint32(v9))) = v17
					F_errmsg(m, int32(_a_F_percentile_disc_final_0), v9)
					mBase = m.M
					v104 = m.ExcPending
					if v104 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_percentile_disc_final_1), int32(448), int32(_a_F_percentile_disc_final_2))
						mBase = m.M
						v109 = m.ExcPending
						if v109 != 0 {
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
			v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
			if v28 == int32(1) {
				v31 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v31)
				v86 = int64(0)
				m.G0 = v9 + int32(32)
				return v86
			} else {
				v34 = int64(0)
				v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				v36 = *(*int64)(unsafe.Add(mBase, uint32(v35)+16))
				if v36 == v34 {
					v39 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v39)
					v86 = v34
					m.G0 = v9 + int32(32)
					return v86
				} else {
					v41 = *(*int32)(unsafe.Add(mBase, uint32(v35)+8))
					v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+24)))
					if v42 == int32(0) {
						F_tuplesort_performsort(m, v41)
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return int64(0)
						} else {
							v49 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v35)+24)) = uint8(v49)
							v53 = *(*int64)(unsafe.Add(mBase, uint32(v35)+16))
							v57 = base.I64_trunc_sat_f64_s(base.F64_ceil(base.F64_mul(v17, base.F64_convert_i64_s(v53))))
							if int64(2) <= v57 {
								v60 = *(*int32)(unsafe.Add(mBase, uint32(v35)+8))
								v63 = F_tuplesort_skiptuples(m, v60, v57-int64(1))
								mBase = m.M
								v64 = m.ExcPending
								if v64 != 0 {
									return int64(0)
								} else {
									if v63 == int32(0) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v113 = m.ExcPending
										if v113 != 0 {
											return int64(0)
										} else {
											F_errmsg_internal(m, int32(_a_F_percentile_disc_final_3), int32(0))
											mBase = m.M
											v117 = m.ExcPending
											if v117 != 0 {
												return int64(0)
											} else {
												F_errfinish(m, int32(_a_F_percentile_disc_final_1), int32(481), int32(_a_F_percentile_disc_final_2))
												mBase = m.M
												v122 = m.ExcPending
												if v122 != 0 {
													return int64(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									} else {
										v67 = *(*int32)(unsafe.Add(mBase, uint32(v35)+8))
										v68 = int32(1)
										v75 = F_tuplesort_getdatum(m, v67, v68, v68, v9+int32(24), v9+int32(23), int32(0))
										mBase = m.M
										v76 = m.ExcPending
										if v76 != 0 {
											return int64(0)
										} else {
											if v75 == int32(0) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v126 = m.ExcPending
												if v126 != 0 {
													return int64(0)
												} else {
													F_errmsg_internal(m, int32(_a_F_percentile_disc_final_3), int32(0))
													mBase = m.M
													v130 = m.ExcPending
													if v130 != 0 {
														return int64(0)
													} else {
														F_errfinish(m, int32(_a_F_percentile_disc_final_1), int32(486), int32(_a_F_percentile_disc_final_2))
														mBase = m.M
														v135 = m.ExcPending
														if v135 != 0 {
															return int64(0)
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											} else {
												v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+23)))
												if v79 == int32(1) {
													v82 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v82)
													v86 = int64(0)
												} else {
													v85 = *(*int64)(unsafe.Add(mBase, uint32(v9)+24))
													v86 = v85
												}
												m.G0 = v9 + int32(32)
												return v86
											}
										}
									}
								}
							} else {
								v67 = *(*int32)(unsafe.Add(mBase, uint32(v35)+8))
								v68 = int32(1)
								v75 = F_tuplesort_getdatum(m, v67, v68, v68, v9+int32(24), v9+int32(23), int32(0))
								mBase = m.M
								v76 = m.ExcPending
								if v76 != 0 {
									return int64(0)
								} else {
									if v75 == int32(0) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v126 = m.ExcPending
										if v126 != 0 {
											return int64(0)
										} else {
											F_errmsg_internal(m, int32(_a_F_percentile_disc_final_3), int32(0))
											mBase = m.M
											v130 = m.ExcPending
											if v130 != 0 {
												return int64(0)
											} else {
												F_errfinish(m, int32(_a_F_percentile_disc_final_1), int32(486), int32(_a_F_percentile_disc_final_2))
												mBase = m.M
												v135 = m.ExcPending
												if v135 != 0 {
													return int64(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									} else {
										v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+23)))
										if v79 == int32(1) {
											v82 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v82)
											v86 = int64(0)
										} else {
											v85 = *(*int64)(unsafe.Add(mBase, uint32(v9)+24))
											v86 = v85
										}
										m.G0 = v9 + int32(32)
										return v86
									}
								}
							}
						}
					} else {
						F_tuplesort_rescan(m, v41)
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return int64(0)
						} else {
							v53 = *(*int64)(unsafe.Add(mBase, uint32(v35)+16))
							v57 = base.I64_trunc_sat_f64_s(base.F64_ceil(base.F64_mul(v17, base.F64_convert_i64_s(v53))))
							if int64(2) <= v57 {
								v60 = *(*int32)(unsafe.Add(mBase, uint32(v35)+8))
								v63 = F_tuplesort_skiptuples(m, v60, v57-int64(1))
								mBase = m.M
								v64 = m.ExcPending
								if v64 != 0 {
									return int64(0)
								} else {
									if v63 == int32(0) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v113 = m.ExcPending
										if v113 != 0 {
											return int64(0)
										} else {
											F_errmsg_internal(m, int32(_a_F_percentile_disc_final_3), int32(0))
											mBase = m.M
											v117 = m.ExcPending
											if v117 != 0 {
												return int64(0)
											} else {
												F_errfinish(m, int32(_a_F_percentile_disc_final_1), int32(481), int32(_a_F_percentile_disc_final_2))
												mBase = m.M
												v122 = m.ExcPending
												if v122 != 0 {
													return int64(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									} else {
										v67 = *(*int32)(unsafe.Add(mBase, uint32(v35)+8))
										v68 = int32(1)
										v75 = F_tuplesort_getdatum(m, v67, v68, v68, v9+int32(24), v9+int32(23), int32(0))
										mBase = m.M
										v76 = m.ExcPending
										if v76 != 0 {
											return int64(0)
										} else {
											if v75 == int32(0) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v126 = m.ExcPending
												if v126 != 0 {
													return int64(0)
												} else {
													F_errmsg_internal(m, int32(_a_F_percentile_disc_final_3), int32(0))
													mBase = m.M
													v130 = m.ExcPending
													if v130 != 0 {
														return int64(0)
													} else {
														F_errfinish(m, int32(_a_F_percentile_disc_final_1), int32(486), int32(_a_F_percentile_disc_final_2))
														mBase = m.M
														v135 = m.ExcPending
														if v135 != 0 {
															return int64(0)
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											} else {
												v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+23)))
												if v79 == int32(1) {
													v82 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v82)
													v86 = int64(0)
												} else {
													v85 = *(*int64)(unsafe.Add(mBase, uint32(v9)+24))
													v86 = v85
												}
												m.G0 = v9 + int32(32)
												return v86
											}
										}
									}
								}
							} else {
								v67 = *(*int32)(unsafe.Add(mBase, uint32(v35)+8))
								v68 = int32(1)
								v75 = F_tuplesort_getdatum(m, v67, v68, v68, v9+int32(24), v9+int32(23), int32(0))
								mBase = m.M
								v76 = m.ExcPending
								if v76 != 0 {
									return int64(0)
								} else {
									if v75 == int32(0) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v126 = m.ExcPending
										if v126 != 0 {
											return int64(0)
										} else {
											F_errmsg_internal(m, int32(_a_F_percentile_disc_final_3), int32(0))
											mBase = m.M
											v130 = m.ExcPending
											if v130 != 0 {
												return int64(0)
											} else {
												F_errfinish(m, int32(_a_F_percentile_disc_final_1), int32(486), int32(_a_F_percentile_disc_final_2))
												mBase = m.M
												v135 = m.ExcPending
												if v135 != 0 {
													return int64(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									} else {
										v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+23)))
										if v79 == int32(1) {
											v82 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v82)
											v86 = int64(0)
										} else {
											v85 = *(*int64)(unsafe.Add(mBase, uint32(v9)+24))
											v86 = v85
										}
										m.G0 = v9 + int32(32)
										return v86
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
func F_percentile_disc_multi_final(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v11 int64
	_ = v11
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
	var v28 int32
	_ = v28
	var v29 int64
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
	var v70 int32
	_ = v70
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
	var v81 int64
	_ = v81
	var v84 int32
	_ = v84
	var v100 int32
	_ = v100
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v118 int64
	_ = v118
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v159 int64
	_ = v159
	var v161 int64
	_ = v161
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int64
	_ = v166
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int64
	_ = v189
	var v190 int32
	_ = v190
	var v191 int64
	_ = v191
	var v192 int64
	_ = v192
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v242 int32
	_ = v242
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v273 int32
	_ = v273
	v11 = int64(0)
	v14 = m.G0
	v16 = v14 - int32(32)
	m.G0 = v16
	*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v11
	v20 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+7)) = uint8(v20)
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v22 == v20 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L13
	} else {
		goto L54
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L13
	} else {
		goto L51
	}
L3:
	;
	m.G0 = v16 + int32(32)
	return base.I64_extend_i32_u(v242)
L4:
	;
	v25 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v25)
	v242 = int32(0)
	goto L3
L5:
	;
	goto L6
L6:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v29 = *(*int64)(unsafe.Add(mBase, uint32(v28)+16))
	if v29 == int64(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v32 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v32)
	v242 = int32(0)
	goto L3
L8:
	;
	goto L9
L9:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
	if v35 == int32(1) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v38 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v38)
	v242 = int32(0)
	goto L3
L11:
	;
	goto L12
L12:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
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
	return int64(0)
L14:
	;
	F_deconstruct_array_builtin(m, v42, int32(701), v16+int32(28), v16+int32(24), v16+int32(20))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v16)+20))
	if v55 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
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
	v62 = int32(0)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v16)+28))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
	v65 = *(*int64)(unsafe.Add(mBase, uint32(v28)+16))
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
	v242 = v60
	goto L3
L20:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v16)+20))
	v73 = F_palloc(m, v70<<(uint(int32(3))%32))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L13
	} else {
		goto L21
	}
L21:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v16)+20))
	v76 = F_palloc(m, v75)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L13
	} else {
		goto L22
	}
L22:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v16)+20))
	if v78 <= int32(0) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	v218 = v42 + int32(16)
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v222)+52))
	v224 = int32(*(*int16)(unsafe.Add(mBase, uint32(v222)+56)))
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222)+58)))
	v226 = int32(*(*int8)(unsafe.Add(mBase, uint32(v222)+59)))
	v227 = F_construct_md_array(m, v73, v76, v216, v218, v218+v216<<(uint(int32(2))%32), v223, v224, v225, v226)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L13
	} else {
		goto L50
	}
L24:
	;
	v81 = *(*int64)(unsafe.Add(mBase, uint32(v67)))
	if v81 <= int64(0) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v84 = v62
	goto L28
L26:
	;
	v121 = v62
	v123 = int32(1)
	goto L27
L27:
	;
	if v123 == int32(0) {
		goto L23
	} else {
		goto L32
	}
L28:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v67+v84<<(uint(int32(5))%32))+24))
	*(*int64)(unsafe.Add(mBase, uint32(v73+v100<<(uint(int32(3))%32)))) = int64(0)
	v107 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v100+v76))) = uint8(v107)
	v110 = v84 + v107
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v16)+20))
	v112 = base.B2i32(v110 < v111)
	if v112 == int32(0) {
		goto L23
	} else {
		goto L30
	}
L29:
	;
	v121 = v110
	v123 = v112
	goto L27
L30:
	;
	v118 = *(*int64)(unsafe.Add(mBase, uint32(v67+v110<<(uint(int32(5))%32))))
	if v118 <= int64(0) {
		v84 = v110
		goto L28
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+24)))
	if v137 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v16)+20))
	if v146 <= v121 {
		goto L23
	} else {
		goto L39
	}
L34:
	;
	F_tuplesort_performsort(m, v136)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L13
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	F_tuplesort_rescan(m, v136)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L13
	} else {
		goto L38
	}
L37:
	;
	v142 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+24)) = uint8(v142)
	goto L33
L38:
	;
	goto L33
L39:
	;
	v149 = v121
	v151 = int32(1)
	v159 = v11
	v161 = v11
	goto L40
L40:
	;
	v164 = v67 + v149<<(uint(int32(5))%32)
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v164)+24))
	v166 = *(*int64)(unsafe.Add(mBase, uint32(v164)))
	if v159 < v166 {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	goto L23
L42:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
	v172 = F_tuplesort_skiptuples(m, v168, v166+(v159^int64(-1)))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L13
	} else {
		goto L45
	}
L43:
	;
	v190 = v151
	v191 = v159
	v192 = v161
	goto L44
L44:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v73+v165<<(uint(int32(3))%32)))) = v192
	*(*uint8)(unsafe.Add(mBase, uint32(v76+v165))) = uint8(v190)
	v200 = v149 + int32(1)
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v16)+20))
	if v200 < v201 {
		v149 = v200
		v151 = v190
		v159 = v191
		v161 = v192
		goto L40
	} else {
		goto L49
	}
L45:
	;
	if v172 == int32(0) {
		goto L2
	} else {
		goto L46
	}
L46:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
	v177 = int32(1)
	v184 = F_tuplesort_getdatum(m, v176, v177, v177, v16+int32(8), v16+int32(7), int32(0))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L13
	} else {
		goto L47
	}
L47:
	;
	if v184 == int32(0) {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+7)))
	v189 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
	v190 = v188
	v191 = v166
	v192 = v189
	goto L44
L49:
	;
	goto L41
L50:
	;
	v242 = v227
	goto L3
L51:
	;
	F_errmsg_internal(m, int32(_a_F_percentile_disc_multi_final_0), int32(0))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L13
	} else {
		goto L52
	}
L52:
	;
	F_errfinish(m, int32(_a_F_percentile_disc_multi_final_1), int32(820), int32(_a_F_percentile_disc_multi_final_2))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
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
	v268 = m.ExcPending
	if v268 != 0 {
		goto L13
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(_a_F_percentile_disc_multi_final_1), int32(824), int32(_a_F_percentile_disc_multi_final_2))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
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
