package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_float8_avg(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 float64
	_ = v25
	var v28 int32
	_ = v28
	var v31 float64
	_ = v31
	var v34 int64
	_ = v34
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int64(0)
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
		if v15 != int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v42 = m.ExcPending
			if v42 != 0 {
				return int64(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = int32(3)
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(_a_F_float8_avg_0)
				F_errmsg_internal(m, int32(_a_F_float8_avg_1), v8)
				mBase = m.M
				v49 = m.ExcPending
				if v49 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_float8_avg_2), int32(2985), int32(_a_F_float8_avg_3))
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return int64(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
			if v18 != int32(3) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int64(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = int32(3)
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(_a_F_float8_avg_0)
					F_errmsg_internal(m, int32(_a_F_float8_avg_1), v8)
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_float8_avg_2), int32(2985), int32(_a_F_float8_avg_3))
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return int64(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
				if v21 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = int32(3)
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(_a_F_float8_avg_0)
						F_errmsg_internal(m, int32(_a_F_float8_avg_1), v8)
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_float8_avg_2), int32(2985), int32(_a_F_float8_avg_3))
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return int64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
					if v22 != int32(701) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = int32(3)
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(_a_F_float8_avg_0)
							F_errmsg_internal(m, int32(_a_F_float8_avg_1), v8)
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_float8_avg_2), int32(2985), int32(_a_F_float8_avg_3))
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return int64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v25 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
						if base.F64_eq(v25, float64(0)) != 0 {
							v28 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v28)
							v34 = int64(0)
						} else {
							v31 = *(*float64)(unsafe.Add(mBase, uint32(v11)+32))
							v34 = base.I64_reinterpret_f64(base.F64_div(v31, v25))
						}
						m.G0 = v8 + int32(16)
						return v34
					}
				}
			}
		}
	}
}
func F_float8_cmp_internal(m *base.Module, l0 float64, l1 float64) int32 {
	var v7 int64
	_ = v7
	var v8 int64
	_ = v8
	var v11 int64
	_ = v11
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v30 int32
	_ = v30
	v7 = int64(9223372036854775807)
	v8 = base.I64_reinterpret_f64(l0) & v7
	v11 = base.I64_reinterpret_f64(l1) & v7
	if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v11) {
		v21 = base.B2i32(base.Ui64(v8) < base.Ui64(int64(9218868437227405313)))
		v30 = int32(0) - v21&(base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v11))|base.F64_lt(l0, l1))
	} else {
		v16 = int32(1)
		if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v8))|base.F64_gt(l0, l1) != 0 {
			v30 = v16
		} else {
			v21 = v16
			v30 = int32(0) - v21&(base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v11))|base.F64_lt(l0, l1))
		}
	}
	return v30
}
func F_float8_regr_combine(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
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
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 float64
	_ = v56
	var v57 float64
	_ = v57
	var v58 float64
	_ = v58
	var v59 float64
	_ = v59
	var v60 float64
	_ = v60
	var v61 float64
	_ = v61
	var v62 float64
	_ = v62
	var v63 float64
	_ = v63
	var v64 float64
	_ = v64
	var v67 float64
	_ = v67
	var v68 float64
	_ = v68
	var v69 float64
	_ = v69
	var v70 float64
	_ = v70
	var v71 float64
	_ = v71
	var v72 float64
	_ = v72
	var v73 float64
	_ = v73
	var v76 float64
	_ = v76
	var v78 float64
	_ = v78
	var v80 float64
	_ = v80
	var v92 float64
	_ = v92
	var v93 int32
	_ = v93
	var v95 float64
	_ = v95
	var v97 float64
	_ = v97
	var v102 float64
	_ = v102
	var v103 float64
	_ = v103
	var v104 float64
	_ = v104
	var v107 float64
	_ = v107
	var v119 float64
	_ = v119
	var v121 float64
	_ = v121
	var v133 float64
	_ = v133
	var v134 int32
	_ = v134
	var v136 float64
	_ = v136
	var v138 float64
	_ = v138
	var v143 float64
	_ = v143
	var v147 float64
	_ = v147
	var v159 float64
	_ = v159
	var v164 float64
	_ = v164
	var v176 int64
	_ = v176
	var v177 int64
	_ = v177
	var v189 float64
	_ = v189
	var v191 int64
	_ = v191
	var v192 int64
	_ = v192
	var v205 float64
	_ = v205
	var v206 float64
	_ = v206
	var v207 float64
	_ = v207
	var v208 float64
	_ = v208
	var v211 float64
	_ = v211
	var v212 float64
	_ = v212
	var v213 float64
	_ = v213
	var v214 float64
	_ = v214
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v248 int32
	_ = v248
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v280 int32
	_ = v280
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v305 int32
	_ = v305
	var v310 int32
	_ = v310
	var v316 int32
	_ = v316
	v24 = m.G0
	v26 = v24 - int32(96)
	m.G0 = v26
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v29 = F_pg_detoast_datum(m, v28)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		return int64(0)
	} else {
		v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v34 = F_pg_detoast_datum(m, v33)
		mBase = m.M
		v35 = m.ExcPending
		if v35 != 0 {
			return int64(0)
		} else {
			v36 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
			if v36 != int32(1) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v280 = m.ExcPending
				if v280 != 0 {
					return int64(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v26)+20)) = int32(8)
					*(*int32)(unsafe.Add(mBase, uint32(v26)+16)) = int32(_a_F_float8_regr_combine_0)
					F_errmsg_internal(m, int32(_a_F_float8_regr_combine_1), v26+int32(16))
					mBase = m.M
					v289 = m.ExcPending
					if v289 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_float8_regr_combine_2), int32(2985), int32(_a_F_float8_regr_combine_3))
						mBase = m.M
						v294 = m.ExcPending
						if v294 != 0 {
							return int64(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v39 = *(*int32)(unsafe.Add(mBase, uint32(v29)+16))
				if v39 != int32(8) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v280 = m.ExcPending
					if v280 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v26)+20)) = int32(8)
						*(*int32)(unsafe.Add(mBase, uint32(v26)+16)) = int32(_a_F_float8_regr_combine_0)
						F_errmsg_internal(m, int32(_a_F_float8_regr_combine_1), v26+int32(16))
						mBase = m.M
						v289 = m.ExcPending
						if v289 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_float8_regr_combine_2), int32(2985), int32(_a_F_float8_regr_combine_3))
							mBase = m.M
							v294 = m.ExcPending
							if v294 != 0 {
								return int64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v42 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
					if v42 != 0 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v280 = m.ExcPending
						if v280 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v26)+20)) = int32(8)
							*(*int32)(unsafe.Add(mBase, uint32(v26)+16)) = int32(_a_F_float8_regr_combine_0)
							F_errmsg_internal(m, int32(_a_F_float8_regr_combine_1), v26+int32(16))
							mBase = m.M
							v289 = m.ExcPending
							if v289 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_float8_regr_combine_2), int32(2985), int32(_a_F_float8_regr_combine_3))
								mBase = m.M
								v294 = m.ExcPending
								if v294 != 0 {
									return int64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v43 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
						if v43 != int32(701) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v280 = m.ExcPending
							if v280 != 0 {
								return int64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v26)+20)) = int32(8)
								*(*int32)(unsafe.Add(mBase, uint32(v26)+16)) = int32(_a_F_float8_regr_combine_0)
								F_errmsg_internal(m, int32(_a_F_float8_regr_combine_1), v26+int32(16))
								mBase = m.M
								v289 = m.ExcPending
								if v289 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_float8_regr_combine_2), int32(2985), int32(_a_F_float8_regr_combine_3))
									mBase = m.M
									v294 = m.ExcPending
									if v294 != 0 {
										return int64(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v46 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
							if v46 != int32(1) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v298 = m.ExcPending
								if v298 != 0 {
									return int64(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v26)+4)) = int32(8)
									*(*int32)(unsafe.Add(mBase, uint32(v26))) = int32(_a_F_float8_regr_combine_0)
									F_errmsg_internal(m, int32(_a_F_float8_regr_combine_1), v26)
									mBase = m.M
									v305 = m.ExcPending
									if v305 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(_a_F_float8_regr_combine_2), int32(2985), int32(_a_F_float8_regr_combine_3))
										mBase = m.M
										v310 = m.ExcPending
										if v310 != 0 {
											return int64(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							} else {
								v49 = *(*int32)(unsafe.Add(mBase, uint32(v34)+16))
								if v49 != int32(8) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v298 = m.ExcPending
									if v298 != 0 {
										return int64(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v26)+4)) = int32(8)
										*(*int32)(unsafe.Add(mBase, uint32(v26))) = int32(_a_F_float8_regr_combine_0)
										F_errmsg_internal(m, int32(_a_F_float8_regr_combine_1), v26)
										mBase = m.M
										v305 = m.ExcPending
										if v305 != 0 {
											return int64(0)
										} else {
											F_errfinish(m, int32(_a_F_float8_regr_combine_2), int32(2985), int32(_a_F_float8_regr_combine_3))
											mBase = m.M
											v310 = m.ExcPending
											if v310 != 0 {
												return int64(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								} else {
									v52 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
									if v52 != 0 {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v298 = m.ExcPending
										if v298 != 0 {
											return int64(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v26)+4)) = int32(8)
											*(*int32)(unsafe.Add(mBase, uint32(v26))) = int32(_a_F_float8_regr_combine_0)
											F_errmsg_internal(m, int32(_a_F_float8_regr_combine_1), v26)
											mBase = m.M
											v305 = m.ExcPending
											if v305 != 0 {
												return int64(0)
											} else {
												F_errfinish(m, int32(_a_F_float8_regr_combine_2), int32(2985), int32(_a_F_float8_regr_combine_3))
												mBase = m.M
												v310 = m.ExcPending
												if v310 != 0 {
													return int64(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									} else {
										v53 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
										if v53 != int32(701) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v298 = m.ExcPending
											if v298 != 0 {
												return int64(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v26)+4)) = int32(8)
												*(*int32)(unsafe.Add(mBase, uint32(v26))) = int32(_a_F_float8_regr_combine_0)
												F_errmsg_internal(m, int32(_a_F_float8_regr_combine_1), v26)
												mBase = m.M
												v305 = m.ExcPending
												if v305 != 0 {
													return int64(0)
												} else {
													F_errfinish(m, int32(_a_F_float8_regr_combine_2), int32(2985), int32(_a_F_float8_regr_combine_3))
													mBase = m.M
													v310 = m.ExcPending
													if v310 != 0 {
														return int64(0)
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											}
										} else {
											v56 = *(*float64)(unsafe.Add(mBase, uint32(v34)+80))
											v57 = *(*float64)(unsafe.Add(mBase, uint32(v34)+72))
											v58 = *(*float64)(unsafe.Add(mBase, uint32(v34)+64))
											v59 = *(*float64)(unsafe.Add(mBase, uint32(v34)+56))
											v60 = *(*float64)(unsafe.Add(mBase, uint32(v34)+48))
											v61 = *(*float64)(unsafe.Add(mBase, uint32(v34)+40))
											v62 = *(*float64)(unsafe.Add(mBase, uint32(v34)+32))
											v63 = *(*float64)(unsafe.Add(mBase, uint32(v34)+24))
											v64 = *(*float64)(unsafe.Add(mBase, uint32(v29)+24))
											if base.F64_eq(v64, float64(0)) != 0 {
												v205 = v60
												v206 = v56
												v207 = v61
												v208 = v63
												v211 = v62
												v212 = v59
												v213 = v58
												v214 = v57
												v220 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
												if v220 == int32(0) {
													v248 = int32(0)
												} else {
													v223 = *(*int32)(unsafe.Add(mBase, uint32(v220)))
													switch v223 - int32(435) {
													case 0:
														v248 = int32(1)
													case 1:
														v248 = int32(2)
													default:
														v248 = int32(0)
													}
												}
												if v248 != 0 {
													*(*float64)(unsafe.Add(mBase, uint32(v29)+80)) = v206
													*(*float64)(unsafe.Add(mBase, uint32(v29)+72)) = v214
													*(*float64)(unsafe.Add(mBase, uint32(v29)+64)) = v213
													*(*float64)(unsafe.Add(mBase, uint32(v29)+56)) = v212
													*(*float64)(unsafe.Add(mBase, uint32(v29)+48)) = v205
													*(*float64)(unsafe.Add(mBase, uint32(v29)+40)) = v207
													*(*float64)(unsafe.Add(mBase, uint32(v29)+32)) = v211
													*(*float64)(unsafe.Add(mBase, uint32(v29)+24)) = v208
													v271 = v29
													m.G0 = v26 + int32(96)
													return base.I64_extend_i32_u(v271)
												} else {
													*(*float64)(unsafe.Add(mBase, uint32(v26)+88)) = v206
													*(*float64)(unsafe.Add(mBase, uint32(v26)+80)) = v214
													*(*float64)(unsafe.Add(mBase, uint32(v26)+72)) = v213
													*(*float64)(unsafe.Add(mBase, uint32(v26)+64)) = v212
													*(*float64)(unsafe.Add(mBase, uint32(v26)+56)) = v205
													*(*float64)(unsafe.Add(mBase, uint32(v26)+48)) = v207
													*(*float64)(unsafe.Add(mBase, uint32(v26)+40)) = v211
													*(*float64)(unsafe.Add(mBase, uint32(v26)+32)) = v208
													v269 = F_construct_array_builtin(m, v26+int32(32), int32(8), int32(701))
													mBase = m.M
													v270 = m.ExcPending
													if v270 != 0 {
														return int64(0)
													} else {
														v271 = v269
														m.G0 = v26 + int32(96)
														return base.I64_extend_i32_u(v271)
													}
												}
											} else {
												v67 = *(*float64)(unsafe.Add(mBase, uint32(v29)+80))
												v68 = *(*float64)(unsafe.Add(mBase, uint32(v29)+72))
												v69 = *(*float64)(unsafe.Add(mBase, uint32(v29)+64))
												v70 = *(*float64)(unsafe.Add(mBase, uint32(v29)+56))
												v71 = *(*float64)(unsafe.Add(mBase, uint32(v29)+48))
												v72 = *(*float64)(unsafe.Add(mBase, uint32(v29)+40))
												v73 = *(*float64)(unsafe.Add(mBase, uint32(v29)+32))
												if base.F64_eq(v63, float64(0)) != 0 {
													v205 = v71
													v206 = v67
													v207 = v72
													v208 = v64
													v211 = v73
													v212 = v70
													v213 = v69
													v214 = v68
													v220 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
													if v220 == int32(0) {
														v248 = int32(0)
													} else {
														v223 = *(*int32)(unsafe.Add(mBase, uint32(v220)))
														switch v223 - int32(435) {
														case 0:
															v248 = int32(1)
														case 1:
															v248 = int32(2)
														default:
															v248 = int32(0)
														}
													}
													if v248 != 0 {
														*(*float64)(unsafe.Add(mBase, uint32(v29)+80)) = v206
														*(*float64)(unsafe.Add(mBase, uint32(v29)+72)) = v214
														*(*float64)(unsafe.Add(mBase, uint32(v29)+64)) = v213
														*(*float64)(unsafe.Add(mBase, uint32(v29)+56)) = v212
														*(*float64)(unsafe.Add(mBase, uint32(v29)+48)) = v205
														*(*float64)(unsafe.Add(mBase, uint32(v29)+40)) = v207
														*(*float64)(unsafe.Add(mBase, uint32(v29)+32)) = v211
														*(*float64)(unsafe.Add(mBase, uint32(v29)+24)) = v208
														v271 = v29
														m.G0 = v26 + int32(96)
														return base.I64_extend_i32_u(v271)
													} else {
														*(*float64)(unsafe.Add(mBase, uint32(v26)+88)) = v206
														*(*float64)(unsafe.Add(mBase, uint32(v26)+80)) = v214
														*(*float64)(unsafe.Add(mBase, uint32(v26)+72)) = v213
														*(*float64)(unsafe.Add(mBase, uint32(v26)+64)) = v212
														*(*float64)(unsafe.Add(mBase, uint32(v26)+56)) = v205
														*(*float64)(unsafe.Add(mBase, uint32(v26)+48)) = v207
														*(*float64)(unsafe.Add(mBase, uint32(v26)+40)) = v211
														*(*float64)(unsafe.Add(mBase, uint32(v26)+32)) = v208
														v269 = F_construct_array_builtin(m, v26+int32(32), int32(8), int32(701))
														mBase = m.M
														v270 = m.ExcPending
														if v270 != 0 {
															return int64(0)
														} else {
															v271 = v269
															m.G0 = v26 + int32(96)
															return base.I64_extend_i32_u(v271)
														}
													}
												} else {
													v76 = base.F64_add(v64, v63)
													v78 = math.Float64frombits(uint64(0x7ff0000000000000))
													v80 = base.F64_add(v73, v62)
													if base.F64_eq(base.F64_abs(v73), v78)|base.F64_ne(base.F64_abs(v80), v78)|base.F64_eq(base.F64_abs(v62), v78) == int32(0) {
														v92 = F_float_overflow_error_ext(m, int32(0))
														mBase = m.M
														v93 = m.ExcPending
														if v93 != 0 {
															return int64(0)
														} else {
															v95 = float64(0)
															v97 = math.Float64frombits(uint64(0x7ff0000000000000))
															v102 = base.F64_sub(base.F64_div(v73, v64), base.F64_div(v62, v63))
															v103 = base.F64_mul(v64, v63)
															v104 = base.F64_mul(v102, v103)
															v107 = base.F64_add(base.F64_add(v72, v61), base.F64_div(base.F64_mul(v102, v104), v76))
															if base.B2i32(base.F64_eq(base.F64_abs(v72), v97)|base.F64_ne(base.F64_abs(v107), v97) == int32(0))&base.F64_ne(base.F64_abs(v61), v97) != 0 {
																F_float_overflow_error(m)
																mBase = m.M
																v316 = m.ExcPending
																if v316 != 0 {
																	return int64(0)
																} else {
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															} else {
																v119 = math.Float64frombits(uint64(0x7ff0000000000000))
																v121 = base.F64_add(v71, v60)
																if base.F64_eq(base.F64_abs(v71), v119)|base.F64_ne(base.F64_abs(v121), v119)|base.F64_eq(base.F64_abs(v60), v119) == int32(0) {
																	v133 = F_float_overflow_error_ext(m, int32(0))
																	mBase = m.M
																	v134 = m.ExcPending
																	if v134 != 0 {
																		return int64(0)
																	} else {
																		v136 = float64(0)
																		v138 = math.Float64frombits(uint64(0x7ff0000000000000))
																		v143 = base.F64_sub(base.F64_div(v71, v64), base.F64_div(v60, v63))
																		v147 = base.F64_add(base.F64_add(v70, v59), base.F64_div(base.F64_mul(v143, base.F64_mul(v103, v143)), v76))
																		if base.B2i32(base.F64_eq(base.F64_abs(v70), v138)|base.F64_ne(base.F64_abs(v147), v138) == int32(0))&base.F64_ne(base.F64_abs(v59), v138) != 0 {
																			F_float_overflow_error(m)
																			mBase = m.M
																			v316 = m.ExcPending
																			if v316 != 0 {
																				return int64(0)
																			} else {
																				base.Wasm_trap_unreachable()
																				for {
																				}
																			}
																		} else {
																			v159 = math.Float64frombits(uint64(0x7ff0000000000000))
																			v164 = base.F64_add(base.F64_add(v69, v58), base.F64_div(base.F64_mul(v104, v143), v76))
																			if base.B2i32(base.F64_eq(base.F64_abs(v69), v159)|base.F64_ne(base.F64_abs(v164), v159) == int32(0))&base.F64_ne(base.F64_abs(v58), v159) != 0 {
																				F_float_overflow_error(m)
																				mBase = m.M
																				v316 = m.ExcPending
																				if v316 != 0 {
																					return int64(0)
																				} else {
																					base.Wasm_trap_unreachable()
																					for {
																					}
																				}
																			} else {
																				v176 = int64(9223372036854775807)
																				v177 = base.I64_reinterpret_f64(v57) & v176
																				if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v68)&v176) {
																					if base.Ui64(v177) <= base.Ui64(int64(9218868437227405312)) {
																						v189 = math.Float64frombits(uint64(0x7ff8000000000000))
																					} else {
																						v189 = v68
																					}
																				} else {
																					if base.F64_ne(v68, v57) != 0 {
																						v189 = math.Float64frombits(uint64(0x7ff8000000000000))
																					} else {
																						if base.Ui64(v177) < base.Ui64(int64(9218868437227405313)) {
																							v189 = v68
																						} else {
																							v189 = math.Float64frombits(uint64(0x7ff8000000000000))
																						}
																					}
																				}
																				v191 = int64(9223372036854775807)
																				v192 = base.I64_reinterpret_f64(v56) & v191
																				if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v67)&v191) {
																					if base.Ui64(v192) <= base.Ui64(int64(9218868437227405312)) {
																						v205 = v136
																						v206 = math.Float64frombits(uint64(0x7ff8000000000000))
																						v207 = v107
																						v208 = v76
																						v211 = v95
																						v212 = v147
																						v213 = v164
																						v214 = v189
																					} else {
																						v205 = v136
																						v206 = v67
																						v207 = v107
																						v208 = v76
																						v211 = v95
																						v212 = v147
																						v213 = v164
																						v214 = v189
																					}
																				} else {
																					if base.B2i32(base.Ui64(int64(9218868437227405313)) <= base.Ui64(v192))|base.F64_ne(v56, v67) != 0 {
																						v205 = v136
																						v206 = math.Float64frombits(uint64(0x7ff8000000000000))
																						v207 = v107
																						v208 = v76
																						v211 = v95
																						v212 = v147
																						v213 = v164
																						v214 = v189
																					} else {
																						v205 = v136
																						v206 = v67
																						v207 = v107
																						v208 = v76
																						v211 = v95
																						v212 = v147
																						v213 = v164
																						v214 = v189
																					}
																				}
																				v220 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
																				if v220 == int32(0) {
																					v248 = int32(0)
																				} else {
																					v223 = *(*int32)(unsafe.Add(mBase, uint32(v220)))
																					switch v223 - int32(435) {
																					case 0:
																						v248 = int32(1)
																					case 1:
																						v248 = int32(2)
																					default:
																						v248 = int32(0)
																					}
																				}
																				if v248 != 0 {
																					*(*float64)(unsafe.Add(mBase, uint32(v29)+80)) = v206
																					*(*float64)(unsafe.Add(mBase, uint32(v29)+72)) = v214
																					*(*float64)(unsafe.Add(mBase, uint32(v29)+64)) = v213
																					*(*float64)(unsafe.Add(mBase, uint32(v29)+56)) = v212
																					*(*float64)(unsafe.Add(mBase, uint32(v29)+48)) = v205
																					*(*float64)(unsafe.Add(mBase, uint32(v29)+40)) = v207
																					*(*float64)(unsafe.Add(mBase, uint32(v29)+32)) = v211
																					*(*float64)(unsafe.Add(mBase, uint32(v29)+24)) = v208
																					v271 = v29
																					m.G0 = v26 + int32(96)
																					return base.I64_extend_i32_u(v271)
																				} else {
																					*(*float64)(unsafe.Add(mBase, uint32(v26)+88)) = v206
																					*(*float64)(unsafe.Add(mBase, uint32(v26)+80)) = v214
																					*(*float64)(unsafe.Add(mBase, uint32(v26)+72)) = v213
																					*(*float64)(unsafe.Add(mBase, uint32(v26)+64)) = v212
																					*(*float64)(unsafe.Add(mBase, uint32(v26)+56)) = v205
																					*(*float64)(unsafe.Add(mBase, uint32(v26)+48)) = v207
																					*(*float64)(unsafe.Add(mBase, uint32(v26)+40)) = v211
																					*(*float64)(unsafe.Add(mBase, uint32(v26)+32)) = v208
																					v269 = F_construct_array_builtin(m, v26+int32(32), int32(8), int32(701))
																					mBase = m.M
																					v270 = m.ExcPending
																					if v270 != 0 {
																						return int64(0)
																					} else {
																						v271 = v269
																						m.G0 = v26 + int32(96)
																						return base.I64_extend_i32_u(v271)
																					}
																				}
																			}
																		}
																	}
																} else {
																	v136 = v121
																	v138 = math.Float64frombits(uint64(0x7ff0000000000000))
																	v143 = base.F64_sub(base.F64_div(v71, v64), base.F64_div(v60, v63))
																	v147 = base.F64_add(base.F64_add(v70, v59), base.F64_div(base.F64_mul(v143, base.F64_mul(v103, v143)), v76))
																	if base.B2i32(base.F64_eq(base.F64_abs(v70), v138)|base.F64_ne(base.F64_abs(v147), v138) == int32(0))&base.F64_ne(base.F64_abs(v59), v138) != 0 {
																		F_float_overflow_error(m)
																		mBase = m.M
																		v316 = m.ExcPending
																		if v316 != 0 {
																			return int64(0)
																		} else {
																			base.Wasm_trap_unreachable()
																			for {
																			}
																		}
																	} else {
																		v159 = math.Float64frombits(uint64(0x7ff0000000000000))
																		v164 = base.F64_add(base.F64_add(v69, v58), base.F64_div(base.F64_mul(v104, v143), v76))
																		if base.B2i32(base.F64_eq(base.F64_abs(v69), v159)|base.F64_ne(base.F64_abs(v164), v159) == int32(0))&base.F64_ne(base.F64_abs(v58), v159) != 0 {
																			F_float_overflow_error(m)
																			mBase = m.M
																			v316 = m.ExcPending
																			if v316 != 0 {
																				return int64(0)
																			} else {
																				base.Wasm_trap_unreachable()
																				for {
																				}
																			}
																		} else {
																			v176 = int64(9223372036854775807)
																			v177 = base.I64_reinterpret_f64(v57) & v176
																			if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v68)&v176) {
																				if base.Ui64(v177) <= base.Ui64(int64(9218868437227405312)) {
																					v189 = math.Float64frombits(uint64(0x7ff8000000000000))
																				} else {
																					v189 = v68
																				}
																			} else {
																				if base.F64_ne(v68, v57) != 0 {
																					v189 = math.Float64frombits(uint64(0x7ff8000000000000))
																				} else {
																					if base.Ui64(v177) < base.Ui64(int64(9218868437227405313)) {
																						v189 = v68
																					} else {
																						v189 = math.Float64frombits(uint64(0x7ff8000000000000))
																					}
																				}
																			}
																			v191 = int64(9223372036854775807)
																			v192 = base.I64_reinterpret_f64(v56) & v191
																			if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v67)&v191) {
																				if base.Ui64(v192) <= base.Ui64(int64(9218868437227405312)) {
																					v205 = v136
																					v206 = math.Float64frombits(uint64(0x7ff8000000000000))
																					v207 = v107
																					v208 = v76
																					v211 = v95
																					v212 = v147
																					v213 = v164
																					v214 = v189
																				} else {
																					v205 = v136
																					v206 = v67
																					v207 = v107
																					v208 = v76
																					v211 = v95
																					v212 = v147
																					v213 = v164
																					v214 = v189
																				}
																			} else {
																				if base.B2i32(base.Ui64(int64(9218868437227405313)) <= base.Ui64(v192))|base.F64_ne(v56, v67) != 0 {
																					v205 = v136
																					v206 = math.Float64frombits(uint64(0x7ff8000000000000))
																					v207 = v107
																					v208 = v76
																					v211 = v95
																					v212 = v147
																					v213 = v164
																					v214 = v189
																				} else {
																					v205 = v136
																					v206 = v67
																					v207 = v107
																					v208 = v76
																					v211 = v95
																					v212 = v147
																					v213 = v164
																					v214 = v189
																				}
																			}
																			v220 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
																			if v220 == int32(0) {
																				v248 = int32(0)
																			} else {
																				v223 = *(*int32)(unsafe.Add(mBase, uint32(v220)))
																				switch v223 - int32(435) {
																				case 0:
																					v248 = int32(1)
																				case 1:
																					v248 = int32(2)
																				default:
																					v248 = int32(0)
																				}
																			}
																			if v248 != 0 {
																				*(*float64)(unsafe.Add(mBase, uint32(v29)+80)) = v206
																				*(*float64)(unsafe.Add(mBase, uint32(v29)+72)) = v214
																				*(*float64)(unsafe.Add(mBase, uint32(v29)+64)) = v213
																				*(*float64)(unsafe.Add(mBase, uint32(v29)+56)) = v212
																				*(*float64)(unsafe.Add(mBase, uint32(v29)+48)) = v205
																				*(*float64)(unsafe.Add(mBase, uint32(v29)+40)) = v207
																				*(*float64)(unsafe.Add(mBase, uint32(v29)+32)) = v211
																				*(*float64)(unsafe.Add(mBase, uint32(v29)+24)) = v208
																				v271 = v29
																				m.G0 = v26 + int32(96)
																				return base.I64_extend_i32_u(v271)
																			} else {
																				*(*float64)(unsafe.Add(mBase, uint32(v26)+88)) = v206
																				*(*float64)(unsafe.Add(mBase, uint32(v26)+80)) = v214
																				*(*float64)(unsafe.Add(mBase, uint32(v26)+72)) = v213
																				*(*float64)(unsafe.Add(mBase, uint32(v26)+64)) = v212
																				*(*float64)(unsafe.Add(mBase, uint32(v26)+56)) = v205
																				*(*float64)(unsafe.Add(mBase, uint32(v26)+48)) = v207
																				*(*float64)(unsafe.Add(mBase, uint32(v26)+40)) = v211
																				*(*float64)(unsafe.Add(mBase, uint32(v26)+32)) = v208
																				v269 = F_construct_array_builtin(m, v26+int32(32), int32(8), int32(701))
																				mBase = m.M
																				v270 = m.ExcPending
																				if v270 != 0 {
																					return int64(0)
																				} else {
																					v271 = v269
																					m.G0 = v26 + int32(96)
																					return base.I64_extend_i32_u(v271)
																				}
																			}
																		}
																	}
																}
															}
														}
													} else {
														v95 = v80
														v97 = math.Float64frombits(uint64(0x7ff0000000000000))
														v102 = base.F64_sub(base.F64_div(v73, v64), base.F64_div(v62, v63))
														v103 = base.F64_mul(v64, v63)
														v104 = base.F64_mul(v102, v103)
														v107 = base.F64_add(base.F64_add(v72, v61), base.F64_div(base.F64_mul(v102, v104), v76))
														if base.B2i32(base.F64_eq(base.F64_abs(v72), v97)|base.F64_ne(base.F64_abs(v107), v97) == int32(0))&base.F64_ne(base.F64_abs(v61), v97) != 0 {
															F_float_overflow_error(m)
															mBase = m.M
															v316 = m.ExcPending
															if v316 != 0 {
																return int64(0)
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														} else {
															v119 = math.Float64frombits(uint64(0x7ff0000000000000))
															v121 = base.F64_add(v71, v60)
															if base.F64_eq(base.F64_abs(v71), v119)|base.F64_ne(base.F64_abs(v121), v119)|base.F64_eq(base.F64_abs(v60), v119) == int32(0) {
																v133 = F_float_overflow_error_ext(m, int32(0))
																mBase = m.M
																v134 = m.ExcPending
																if v134 != 0 {
																	return int64(0)
																} else {
																	v136 = float64(0)
																	v138 = math.Float64frombits(uint64(0x7ff0000000000000))
																	v143 = base.F64_sub(base.F64_div(v71, v64), base.F64_div(v60, v63))
																	v147 = base.F64_add(base.F64_add(v70, v59), base.F64_div(base.F64_mul(v143, base.F64_mul(v103, v143)), v76))
																	if base.B2i32(base.F64_eq(base.F64_abs(v70), v138)|base.F64_ne(base.F64_abs(v147), v138) == int32(0))&base.F64_ne(base.F64_abs(v59), v138) != 0 {
																		F_float_overflow_error(m)
																		mBase = m.M
																		v316 = m.ExcPending
																		if v316 != 0 {
																			return int64(0)
																		} else {
																			base.Wasm_trap_unreachable()
																			for {
																			}
																		}
																	} else {
																		v159 = math.Float64frombits(uint64(0x7ff0000000000000))
																		v164 = base.F64_add(base.F64_add(v69, v58), base.F64_div(base.F64_mul(v104, v143), v76))
																		if base.B2i32(base.F64_eq(base.F64_abs(v69), v159)|base.F64_ne(base.F64_abs(v164), v159) == int32(0))&base.F64_ne(base.F64_abs(v58), v159) != 0 {
																			F_float_overflow_error(m)
																			mBase = m.M
																			v316 = m.ExcPending
																			if v316 != 0 {
																				return int64(0)
																			} else {
																				base.Wasm_trap_unreachable()
																				for {
																				}
																			}
																		} else {
																			v176 = int64(9223372036854775807)
																			v177 = base.I64_reinterpret_f64(v57) & v176
																			if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v68)&v176) {
																				if base.Ui64(v177) <= base.Ui64(int64(9218868437227405312)) {
																					v189 = math.Float64frombits(uint64(0x7ff8000000000000))
																				} else {
																					v189 = v68
																				}
																			} else {
																				if base.F64_ne(v68, v57) != 0 {
																					v189 = math.Float64frombits(uint64(0x7ff8000000000000))
																				} else {
																					if base.Ui64(v177) < base.Ui64(int64(9218868437227405313)) {
																						v189 = v68
																					} else {
																						v189 = math.Float64frombits(uint64(0x7ff8000000000000))
																					}
																				}
																			}
																			v191 = int64(9223372036854775807)
																			v192 = base.I64_reinterpret_f64(v56) & v191
																			if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v67)&v191) {
																				if base.Ui64(v192) <= base.Ui64(int64(9218868437227405312)) {
																					v205 = v136
																					v206 = math.Float64frombits(uint64(0x7ff8000000000000))
																					v207 = v107
																					v208 = v76
																					v211 = v95
																					v212 = v147
																					v213 = v164
																					v214 = v189
																				} else {
																					v205 = v136
																					v206 = v67
																					v207 = v107
																					v208 = v76
																					v211 = v95
																					v212 = v147
																					v213 = v164
																					v214 = v189
																				}
																			} else {
																				if base.B2i32(base.Ui64(int64(9218868437227405313)) <= base.Ui64(v192))|base.F64_ne(v56, v67) != 0 {
																					v205 = v136
																					v206 = math.Float64frombits(uint64(0x7ff8000000000000))
																					v207 = v107
																					v208 = v76
																					v211 = v95
																					v212 = v147
																					v213 = v164
																					v214 = v189
																				} else {
																					v205 = v136
																					v206 = v67
																					v207 = v107
																					v208 = v76
																					v211 = v95
																					v212 = v147
																					v213 = v164
																					v214 = v189
																				}
																			}
																			v220 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
																			if v220 == int32(0) {
																				v248 = int32(0)
																			} else {
																				v223 = *(*int32)(unsafe.Add(mBase, uint32(v220)))
																				switch v223 - int32(435) {
																				case 0:
																					v248 = int32(1)
																				case 1:
																					v248 = int32(2)
																				default:
																					v248 = int32(0)
																				}
																			}
																			if v248 != 0 {
																				*(*float64)(unsafe.Add(mBase, uint32(v29)+80)) = v206
																				*(*float64)(unsafe.Add(mBase, uint32(v29)+72)) = v214
																				*(*float64)(unsafe.Add(mBase, uint32(v29)+64)) = v213
																				*(*float64)(unsafe.Add(mBase, uint32(v29)+56)) = v212
																				*(*float64)(unsafe.Add(mBase, uint32(v29)+48)) = v205
																				*(*float64)(unsafe.Add(mBase, uint32(v29)+40)) = v207
																				*(*float64)(unsafe.Add(mBase, uint32(v29)+32)) = v211
																				*(*float64)(unsafe.Add(mBase, uint32(v29)+24)) = v208
																				v271 = v29
																				m.G0 = v26 + int32(96)
																				return base.I64_extend_i32_u(v271)
																			} else {
																				*(*float64)(unsafe.Add(mBase, uint32(v26)+88)) = v206
																				*(*float64)(unsafe.Add(mBase, uint32(v26)+80)) = v214
																				*(*float64)(unsafe.Add(mBase, uint32(v26)+72)) = v213
																				*(*float64)(unsafe.Add(mBase, uint32(v26)+64)) = v212
																				*(*float64)(unsafe.Add(mBase, uint32(v26)+56)) = v205
																				*(*float64)(unsafe.Add(mBase, uint32(v26)+48)) = v207
																				*(*float64)(unsafe.Add(mBase, uint32(v26)+40)) = v211
																				*(*float64)(unsafe.Add(mBase, uint32(v26)+32)) = v208
																				v269 = F_construct_array_builtin(m, v26+int32(32), int32(8), int32(701))
																				mBase = m.M
																				v270 = m.ExcPending
																				if v270 != 0 {
																					return int64(0)
																				} else {
																					v271 = v269
																					m.G0 = v26 + int32(96)
																					return base.I64_extend_i32_u(v271)
																				}
																			}
																		}
																	}
																}
															} else {
																v136 = v121
																v138 = math.Float64frombits(uint64(0x7ff0000000000000))
																v143 = base.F64_sub(base.F64_div(v71, v64), base.F64_div(v60, v63))
																v147 = base.F64_add(base.F64_add(v70, v59), base.F64_div(base.F64_mul(v143, base.F64_mul(v103, v143)), v76))
																if base.B2i32(base.F64_eq(base.F64_abs(v70), v138)|base.F64_ne(base.F64_abs(v147), v138) == int32(0))&base.F64_ne(base.F64_abs(v59), v138) != 0 {
																	F_float_overflow_error(m)
																	mBase = m.M
																	v316 = m.ExcPending
																	if v316 != 0 {
																		return int64(0)
																	} else {
																		base.Wasm_trap_unreachable()
																		for {
																		}
																	}
																} else {
																	v159 = math.Float64frombits(uint64(0x7ff0000000000000))
																	v164 = base.F64_add(base.F64_add(v69, v58), base.F64_div(base.F64_mul(v104, v143), v76))
																	if base.B2i32(base.F64_eq(base.F64_abs(v69), v159)|base.F64_ne(base.F64_abs(v164), v159) == int32(0))&base.F64_ne(base.F64_abs(v58), v159) != 0 {
																		F_float_overflow_error(m)
																		mBase = m.M
																		v316 = m.ExcPending
																		if v316 != 0 {
																			return int64(0)
																		} else {
																			base.Wasm_trap_unreachable()
																			for {
																			}
																		}
																	} else {
																		v176 = int64(9223372036854775807)
																		v177 = base.I64_reinterpret_f64(v57) & v176
																		if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v68)&v176) {
																			if base.Ui64(v177) <= base.Ui64(int64(9218868437227405312)) {
																				v189 = math.Float64frombits(uint64(0x7ff8000000000000))
																			} else {
																				v189 = v68
																			}
																		} else {
																			if base.F64_ne(v68, v57) != 0 {
																				v189 = math.Float64frombits(uint64(0x7ff8000000000000))
																			} else {
																				if base.Ui64(v177) < base.Ui64(int64(9218868437227405313)) {
																					v189 = v68
																				} else {
																					v189 = math.Float64frombits(uint64(0x7ff8000000000000))
																				}
																			}
																		}
																		v191 = int64(9223372036854775807)
																		v192 = base.I64_reinterpret_f64(v56) & v191
																		if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v67)&v191) {
																			if base.Ui64(v192) <= base.Ui64(int64(9218868437227405312)) {
																				v205 = v136
																				v206 = math.Float64frombits(uint64(0x7ff8000000000000))
																				v207 = v107
																				v208 = v76
																				v211 = v95
																				v212 = v147
																				v213 = v164
																				v214 = v189
																			} else {
																				v205 = v136
																				v206 = v67
																				v207 = v107
																				v208 = v76
																				v211 = v95
																				v212 = v147
																				v213 = v164
																				v214 = v189
																			}
																		} else {
																			if base.B2i32(base.Ui64(int64(9218868437227405313)) <= base.Ui64(v192))|base.F64_ne(v56, v67) != 0 {
																				v205 = v136
																				v206 = math.Float64frombits(uint64(0x7ff8000000000000))
																				v207 = v107
																				v208 = v76
																				v211 = v95
																				v212 = v147
																				v213 = v164
																				v214 = v189
																			} else {
																				v205 = v136
																				v206 = v67
																				v207 = v107
																				v208 = v76
																				v211 = v95
																				v212 = v147
																				v213 = v164
																				v214 = v189
																			}
																		}
																		v220 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
																		if v220 == int32(0) {
																			v248 = int32(0)
																		} else {
																			v223 = *(*int32)(unsafe.Add(mBase, uint32(v220)))
																			switch v223 - int32(435) {
																			case 0:
																				v248 = int32(1)
																			case 1:
																				v248 = int32(2)
																			default:
																				v248 = int32(0)
																			}
																		}
																		if v248 != 0 {
																			*(*float64)(unsafe.Add(mBase, uint32(v29)+80)) = v206
																			*(*float64)(unsafe.Add(mBase, uint32(v29)+72)) = v214
																			*(*float64)(unsafe.Add(mBase, uint32(v29)+64)) = v213
																			*(*float64)(unsafe.Add(mBase, uint32(v29)+56)) = v212
																			*(*float64)(unsafe.Add(mBase, uint32(v29)+48)) = v205
																			*(*float64)(unsafe.Add(mBase, uint32(v29)+40)) = v207
																			*(*float64)(unsafe.Add(mBase, uint32(v29)+32)) = v211
																			*(*float64)(unsafe.Add(mBase, uint32(v29)+24)) = v208
																			v271 = v29
																			m.G0 = v26 + int32(96)
																			return base.I64_extend_i32_u(v271)
																		} else {
																			*(*float64)(unsafe.Add(mBase, uint32(v26)+88)) = v206
																			*(*float64)(unsafe.Add(mBase, uint32(v26)+80)) = v214
																			*(*float64)(unsafe.Add(mBase, uint32(v26)+72)) = v213
																			*(*float64)(unsafe.Add(mBase, uint32(v26)+64)) = v212
																			*(*float64)(unsafe.Add(mBase, uint32(v26)+56)) = v205
																			*(*float64)(unsafe.Add(mBase, uint32(v26)+48)) = v207
																			*(*float64)(unsafe.Add(mBase, uint32(v26)+40)) = v211
																			*(*float64)(unsafe.Add(mBase, uint32(v26)+32)) = v208
																			v269 = F_construct_array_builtin(m, v26+int32(32), int32(8), int32(701))
																			mBase = m.M
																			v270 = m.ExcPending
																			if v270 != 0 {
																				return int64(0)
																			} else {
																				v271 = v269
																				m.G0 = v26 + int32(96)
																				return base.I64_extend_i32_u(v271)
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
}
func F_float8_regr_sxx(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 float64
	_ = v24
	var v27 int32
	_ = v27
	var v30 int64
	_ = v30
	var v31 int64
	_ = v31
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = F_pg_detoast_datum(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int64(0)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
		if v14 != int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return int64(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(8)
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(_a_F_float8_regr_sxx_0)
				F_errmsg_internal(m, int32(_a_F_float8_regr_sxx_1), v7)
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_float8_regr_sxx_2), int32(2985), int32(_a_F_float8_regr_sxx_3))
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return int64(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
			if v17 != int32(8) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int64(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(8)
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(_a_F_float8_regr_sxx_0)
					F_errmsg_internal(m, int32(_a_F_float8_regr_sxx_1), v7)
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_float8_regr_sxx_2), int32(2985), int32(_a_F_float8_regr_sxx_3))
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return int64(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
				if v20 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(8)
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(_a_F_float8_regr_sxx_0)
						F_errmsg_internal(m, int32(_a_F_float8_regr_sxx_1), v7)
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_float8_regr_sxx_2), int32(2985), int32(_a_F_float8_regr_sxx_3))
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return int64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v21 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
					if v21 != int32(701) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(8)
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(_a_F_float8_regr_sxx_0)
							F_errmsg_internal(m, int32(_a_F_float8_regr_sxx_1), v7)
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_float8_regr_sxx_2), int32(2985), int32(_a_F_float8_regr_sxx_3))
								mBase = m.M
								v51 = m.ExcPending
								if v51 != 0 {
									return int64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v24 = *(*float64)(unsafe.Add(mBase, uint32(v10)+24))
						if base.F64_lt(v24, float64(1)) != 0 {
							v27 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v27)
							v31 = int64(0)
						} else {
							v30 = *(*int64)(unsafe.Add(mBase, uint32(v10)+40))
							v31 = v30
						}
						m.G0 = v7 + int32(16)
						return v31
					}
				}
			}
		}
	}
}
func F_float8_regr_syy(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 float64
	_ = v24
	var v27 int32
	_ = v27
	var v30 int64
	_ = v30
	var v31 int64
	_ = v31
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = F_pg_detoast_datum(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int64(0)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
		if v14 != int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return int64(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(8)
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(_a_F_float8_regr_syy_0)
				F_errmsg_internal(m, int32(_a_F_float8_regr_syy_1), v7)
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_float8_regr_syy_2), int32(2985), int32(_a_F_float8_regr_syy_3))
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return int64(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
			if v17 != int32(8) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int64(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(8)
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(_a_F_float8_regr_syy_0)
					F_errmsg_internal(m, int32(_a_F_float8_regr_syy_1), v7)
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_float8_regr_syy_2), int32(2985), int32(_a_F_float8_regr_syy_3))
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return int64(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
				if v20 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(8)
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(_a_F_float8_regr_syy_0)
						F_errmsg_internal(m, int32(_a_F_float8_regr_syy_1), v7)
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_float8_regr_syy_2), int32(2985), int32(_a_F_float8_regr_syy_3))
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return int64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v21 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
					if v21 != int32(701) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(8)
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(_a_F_float8_regr_syy_0)
							F_errmsg_internal(m, int32(_a_F_float8_regr_syy_1), v7)
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_float8_regr_syy_2), int32(2985), int32(_a_F_float8_regr_syy_3))
								mBase = m.M
								v51 = m.ExcPending
								if v51 != 0 {
									return int64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v24 = *(*float64)(unsafe.Add(mBase, uint32(v10)+24))
						if base.F64_lt(v24, float64(1)) != 0 {
							v27 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v27)
							v31 = int64(0)
						} else {
							v30 = *(*int64)(unsafe.Add(mBase, uint32(v10)+56))
							v31 = v30
						}
						m.G0 = v7 + int32(16)
						return v31
					}
				}
			}
		}
	}
}
func F_float8_stddev_pop(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 float64
	_ = v25
	var v28 int32
	_ = v28
	var v31 float64
	_ = v31
	var v35 int64
	_ = v35
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int64(0)
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
		if v15 != int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return int64(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = int32(3)
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(_a_F_float8_stddev_pop_0)
				F_errmsg_internal(m, int32(_a_F_float8_stddev_pop_1), v8)
				mBase = m.M
				v50 = m.ExcPending
				if v50 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_float8_stddev_pop_2), int32(2985), int32(_a_F_float8_stddev_pop_3))
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return int64(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
			if v18 != int32(3) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return int64(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = int32(3)
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(_a_F_float8_stddev_pop_0)
					F_errmsg_internal(m, int32(_a_F_float8_stddev_pop_1), v8)
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_float8_stddev_pop_2), int32(2985), int32(_a_F_float8_stddev_pop_3))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int64(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
				if v21 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = int32(3)
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(_a_F_float8_stddev_pop_0)
						F_errmsg_internal(m, int32(_a_F_float8_stddev_pop_1), v8)
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_float8_stddev_pop_2), int32(2985), int32(_a_F_float8_stddev_pop_3))
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
								return int64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
					if v22 != int32(701) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = int32(3)
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(_a_F_float8_stddev_pop_0)
							F_errmsg_internal(m, int32(_a_F_float8_stddev_pop_1), v8)
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_float8_stddev_pop_2), int32(2985), int32(_a_F_float8_stddev_pop_3))
								mBase = m.M
								v55 = m.ExcPending
								if v55 != 0 {
									return int64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v25 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
						if base.F64_eq(v25, float64(0)) != 0 {
							v28 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v28)
							v35 = int64(0)
						} else {
							v31 = *(*float64)(unsafe.Add(mBase, uint32(v11)+40))
							v35 = base.I64_reinterpret_f64(base.F64_sqrt(base.F64_div(v31, v25)))
						}
						m.G0 = v8 + int32(16)
						return v35
					}
				}
			}
		}
	}
}
func F_float8_stddev_samp(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 float64
	_ = v25
	var v28 int32
	_ = v28
	var v31 float64
	_ = v31
	var v37 int64
	_ = v37
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int64(0)
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
		if v15 != int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v45 = m.ExcPending
			if v45 != 0 {
				return int64(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = int32(3)
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(_a_F_float8_stddev_samp_0)
				F_errmsg_internal(m, int32(_a_F_float8_stddev_samp_1), v8)
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_float8_stddev_samp_2), int32(2985), int32(_a_F_float8_stddev_samp_3))
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
						return int64(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
			if v18 != int32(3) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return int64(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = int32(3)
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(_a_F_float8_stddev_samp_0)
					F_errmsg_internal(m, int32(_a_F_float8_stddev_samp_1), v8)
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_float8_stddev_samp_2), int32(2985), int32(_a_F_float8_stddev_samp_3))
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
							return int64(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
				if v21 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = int32(3)
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(_a_F_float8_stddev_samp_0)
						F_errmsg_internal(m, int32(_a_F_float8_stddev_samp_1), v8)
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_float8_stddev_samp_2), int32(2985), int32(_a_F_float8_stddev_samp_3))
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return int64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
					if v22 != int32(701) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = int32(3)
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(_a_F_float8_stddev_samp_0)
							F_errmsg_internal(m, int32(_a_F_float8_stddev_samp_1), v8)
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_float8_stddev_samp_2), int32(2985), int32(_a_F_float8_stddev_samp_3))
								mBase = m.M
								v57 = m.ExcPending
								if v57 != 0 {
									return int64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v25 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
						if base.F64_le(v25, float64(1)) != 0 {
							v28 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v28)
							v37 = int64(0)
						} else {
							v31 = *(*float64)(unsafe.Add(mBase, uint32(v11)+40))
							v37 = base.I64_reinterpret_f64(base.F64_sqrt(base.F64_div(v31, base.F64_add(v25, float64(-1)))))
						}
						m.G0 = v8 + int32(16)
						return v37
					}
				}
			}
		}
	}
}
func F_float8_to_char(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int64
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 float64
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 float64
	_ = v76
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v147 float64
	_ = v147
	var v150 float64
	_ = v150
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v244 int32
	_ = v244
	v2 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(96)
	m.G0 = v14
	v16 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v18 = F_pg_detoast_datum_packed(m, v17)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return int64(0)
	} else {
		v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
		if v22 == int32(1) {
			v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+1)))
			if v28 == int32(18) {
				v31 = int32(16)
			} else {
				v31 = int32(0)
			}
			if base.Ui32((v28-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				v38 = int32(4)
			} else {
				v38 = v31
			}
			v51 = v38
		} else {
			v39 = int32(1)
			if v22&v39 != 0 {
				v51 = int32(base.Ui32(v22)>>(uint(v39)%32)) - v39
			} else {
				v45 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
				v51 = int32(base.Ui32(v45)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		if base.Ui32(int32(268435454)) <= base.Ui32(v51-int32(1)) {
			v57 = F_cstring_to_text(m, int32(_a_F_float8_to_char_0))
			mBase = m.M
			v58 = m.ExcPending
			if v58 != 0 {
				return int64(0)
			} else {
				v244 = v57
				m.G0 = v14 + int32(96)
				return base.I64_extend_i32_u(v244)
			}
		} else {
			v59 = base.F64_reinterpret_i64(v16)
			v64 = F_palloc0(m, v51<<(uint(int32(3))%32)|int32(5))
			mBase = m.M
			v65 = m.ExcPending
			if v65 != 0 {
				return int64(0)
			} else {
				v70 = F_NUM_cache(m, v51, v14+int32(60), v18, v14+int32(59))
				mBase = m.M
				v71 = m.ExcPending
				if v71 != 0 {
					return int64(0)
				} else {
					v72 = *(*int32)(unsafe.Add(mBase, uint32(v14)+72))
					if v72&int32(1024) != 0 {
						v76 = base.F64_nearest(v59)
						if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v76)&int64(9223372036854775807)) {
							v83 = int32(2147483647)
						} else {
							v83 = base.I32_trunc_sat_f64_s(v76)
						}
						if base.F64_ge(v76, float64(-2.147483648e+09)) != 0 {
							v87 = v83
						} else {
							v87 = int32(2147483647)
						}
						if base.F64_lt(v76, float64(2.147483648e+09)) != 0 {
							v91 = v87
						} else {
							v91 = int32(2147483647)
						}
						v92 = F_int_to_roman(m, v91)
						mBase = m.M
						v93 = m.ExcPending
						if v93 != 0 {
							return int64(0)
						} else {
							v214 = v92
							v217 = v2
							v219 = v2
							v224 = v64 + int32(4)
							F_NUM_processor(m, v70, v14+int32(60), v224, v214, int32(0), v217, v219, int32(1))
							mBase = m.M
							v228 = m.ExcPending
							if v228 != 0 {
								return int64(0)
							} else {
								v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+59)))
								if v229 == int32(1) {
									F_pfree(m, v70)
									mBase = m.M
									v233 = m.ExcPending
									if v233 != 0 {
										return int64(0)
									} else {
										v234 = F_strlen(m, v224)
										mBase = m.M
										*(*int32)(unsafe.Add(mBase, uint32(v64))) = v234<<(uint(int32(2))%32) + int32(16)
										v244 = v64
										m.G0 = v14 + int32(96)
										return base.I64_extend_i32_u(v244)
									}
								} else {
									v234 = F_strlen(m, v224)
									mBase = m.M
									*(*int32)(unsafe.Add(mBase, uint32(v64))) = v234<<(uint(int32(2))%32) + int32(16)
									v244 = v64
									m.G0 = v14 + int32(96)
									return base.I64_extend_i32_u(v244)
								}
							}
						}
					} else {
						if v72&int32(_a_F_float8_to_char_1) != 0 {
							if base.B2i32(base.Ui64(v16&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)))&base.F64_ne(base.F64_abs(v59), math.Float64frombits(uint64(0x7ff0000000000000))) == int32(0) {
								v106 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
								v107 = *(*int32)(unsafe.Add(mBase, uint32(v14)+64))
								v108 = v106 + v107
								v111 = F_palloc(m, v108+int32(7))
								mBase = m.M
								v112 = m.ExcPending
								if v112 != 0 {
									return int64(0)
								} else {
									v114 = v108 + int32(6)
									if v114 != 0 {
										base.MemoryFill(m, v111, int32(35), v114)
									} else {
									}
									v117 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v111+v114))) = uint8(v117)
									v121 = int32(32)
									*(*uint8)(unsafe.Add(mBase, uint32(v111))) = uint8(v121)
									v124 = int32(46)
									*(*uint8)(unsafe.Add(mBase, uint32(v111+v106)+1)) = uint8(v124)
									v214 = v111
									v217 = v117
									v219 = v2
									v224 = v64 + int32(4)
									F_NUM_processor(m, v70, v14+int32(60), v224, v214, int32(0), v217, v219, int32(1))
									mBase = m.M
									v228 = m.ExcPending
									if v228 != 0 {
										return int64(0)
									} else {
										v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+59)))
										if v229 == int32(1) {
											F_pfree(m, v70)
											mBase = m.M
											v233 = m.ExcPending
											if v233 != 0 {
												return int64(0)
											} else {
												v234 = F_strlen(m, v224)
												mBase = m.M
												*(*int32)(unsafe.Add(mBase, uint32(v64))) = v234<<(uint(int32(2))%32) + int32(16)
												v244 = v64
												m.G0 = v14 + int32(96)
												return base.I64_extend_i32_u(v244)
											}
										} else {
											v234 = F_strlen(m, v224)
											mBase = m.M
											*(*int32)(unsafe.Add(mBase, uint32(v64))) = v234<<(uint(int32(2))%32) + int32(16)
											v244 = v64
											m.G0 = v14 + int32(96)
											return base.I64_extend_i32_u(v244)
										}
									}
								}
							} else {
								v126 = *(*int32)(unsafe.Add(mBase, uint32(v14)+64))
								*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v126
								*(*float64)(unsafe.Add(mBase, uint32(v14)+40)) = v59
								v132 = F_psprintf(m, int32(_a_F_float8_to_char_2), v14+int32(32))
								mBase = m.M
								v133 = m.ExcPending
								if v133 != 0 {
									return int64(0)
								} else {
									v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132))))
									if v134 != int32(43) {
										v214 = v132
										v217 = v2
										v219 = v2
									} else {
										v137 = int32(32)
										*(*uint8)(unsafe.Add(mBase, uint32(v132))) = uint8(v137)
										v214 = v132
										v217 = v2
										v219 = v2
									}
									v224 = v64 + int32(4)
									F_NUM_processor(m, v70, v14+int32(60), v224, v214, int32(0), v217, v219, int32(1))
									mBase = m.M
									v228 = m.ExcPending
									if v228 != 0 {
										return int64(0)
									} else {
										v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+59)))
										if v229 == int32(1) {
											F_pfree(m, v70)
											mBase = m.M
											v233 = m.ExcPending
											if v233 != 0 {
												return int64(0)
											} else {
												v234 = F_strlen(m, v224)
												mBase = m.M
												*(*int32)(unsafe.Add(mBase, uint32(v64))) = v234<<(uint(int32(2))%32) + int32(16)
												v244 = v64
												m.G0 = v14 + int32(96)
												return base.I64_extend_i32_u(v244)
											}
										} else {
											v234 = F_strlen(m, v224)
											mBase = m.M
											*(*int32)(unsafe.Add(mBase, uint32(v64))) = v234<<(uint(int32(2))%32) + int32(16)
											v244 = v64
											m.G0 = v14 + int32(96)
											return base.I64_extend_i32_u(v244)
										}
									}
								}
							}
						} else {
							if v72&int32(2048) != 0 {
								v141 = *(*int32)(unsafe.Add(mBase, uint32(v14)+80))
								v142 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
								*(*int32)(unsafe.Add(mBase, uint32(v14)+60)) = v141 + v142
								v147 = F_pow(m, float64(10), base.F64_convert_i32_s(v141))
								mBase = m.M
								v150 = base.F64_mul(v147, v59)
							} else {
								v150 = v59
							}
							*(*float64)(unsafe.Add(mBase, uint32(v14)+16)) = base.F64_abs(v150)
							v156 = F_psprintf(m, int32(_a_F_float8_to_char_3), v14+int32(16))
							mBase = m.M
							v157 = m.ExcPending
							if v157 != 0 {
								return int64(0)
							} else {
								v158 = F_strlen(m, v156)
								mBase = m.M
								if base.Ui32(v158) <= base.Ui32(int32(14)) {
									v161 = *(*int32)(unsafe.Add(mBase, uint32(v14)+64))
									if base.Ui32(v161+v158) < base.Ui32(int32(16)) {
										v169 = v161
									} else {
										v167 = int32(15) - v158
										*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = v167
										v169 = v167
									}
								} else {
									v167 = v2
									*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = v167
									v169 = v167
								}
								*(*float64)(unsafe.Add(mBase, uint32(v14)+8)) = v150
								*(*int32)(unsafe.Add(mBase, uint32(v14))) = v169
								v173 = F_psprintf(m, int32(_a_F_float8_to_char_4), v14)
								mBase = m.M
								v174 = m.ExcPending
								if v174 != 0 {
									return int64(0)
								} else {
									v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173))))
									v177 = base.B2i32(v175 == int32(45))
									v178 = v173 + v177
									v179 = int32(46)
									v180 = F___strchrnul(m, v178, v179)
									mBase = m.M
									v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180))))
									if v182 == v179 {
										v186 = v180
									} else {
										v186 = int32(0)
									}
									if v186 != 0 {
										v189 = v186 - v178
									} else {
										v188 = F_strlen(m, v178)
										mBase = m.M
										v189 = v188
									}
									if v175 == int32(45) {
										v192 = int32(45)
									} else {
										v192 = int32(43)
									}
									v193 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
									if base.Ui32(v189) < base.Ui32(v193) {
										v214 = v178
										v217 = v193 - v189
										v219 = v192
										v224 = v64 + int32(4)
										F_NUM_processor(m, v70, v14+int32(60), v224, v214, int32(0), v217, v219, int32(1))
										mBase = m.M
										v228 = m.ExcPending
										if v228 != 0 {
											return int64(0)
										} else {
											v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+59)))
											if v229 == int32(1) {
												F_pfree(m, v70)
												mBase = m.M
												v233 = m.ExcPending
												if v233 != 0 {
													return int64(0)
												} else {
													v234 = F_strlen(m, v224)
													mBase = m.M
													*(*int32)(unsafe.Add(mBase, uint32(v64))) = v234<<(uint(int32(2))%32) + int32(16)
													v244 = v64
													m.G0 = v14 + int32(96)
													return base.I64_extend_i32_u(v244)
												}
											} else {
												v234 = F_strlen(m, v224)
												mBase = m.M
												*(*int32)(unsafe.Add(mBase, uint32(v64))) = v234<<(uint(int32(2))%32) + int32(16)
												v244 = v64
												m.G0 = v14 + int32(96)
												return base.I64_extend_i32_u(v244)
											}
										}
									} else {
										if base.Ui32(v189) <= base.Ui32(v193) {
											v214 = v178
											v217 = int32(0)
											v219 = v192
											v224 = v64 + int32(4)
											F_NUM_processor(m, v70, v14+int32(60), v224, v214, int32(0), v217, v219, int32(1))
											mBase = m.M
											v228 = m.ExcPending
											if v228 != 0 {
												return int64(0)
											} else {
												v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+59)))
												if v229 == int32(1) {
													F_pfree(m, v70)
													mBase = m.M
													v233 = m.ExcPending
													if v233 != 0 {
														return int64(0)
													} else {
														v234 = F_strlen(m, v224)
														mBase = m.M
														*(*int32)(unsafe.Add(mBase, uint32(v64))) = v234<<(uint(int32(2))%32) + int32(16)
														v244 = v64
														m.G0 = v14 + int32(96)
														return base.I64_extend_i32_u(v244)
													}
												} else {
													v234 = F_strlen(m, v224)
													mBase = m.M
													*(*int32)(unsafe.Add(mBase, uint32(v64))) = v234<<(uint(int32(2))%32) + int32(16)
													v244 = v64
													m.G0 = v14 + int32(96)
													return base.I64_extend_i32_u(v244)
												}
											}
										} else {
											v198 = v169 + v193
											v201 = F_palloc(m, v198+int32(2))
											mBase = m.M
											v202 = m.ExcPending
											if v202 != 0 {
												return int64(0)
											} else {
												v204 = v198 + int32(1)
												if v204 != 0 {
													base.MemoryFill(m, v201, int32(35), v204)
												} else {
												}
												v207 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(v201+v204))) = uint8(v207)
												v212 = int32(46)
												*(*uint8)(unsafe.Add(mBase, uint32(v201+v193))) = uint8(v212)
												v214 = v201
												v217 = v207
												v219 = v192
												v224 = v64 + int32(4)
												F_NUM_processor(m, v70, v14+int32(60), v224, v214, int32(0), v217, v219, int32(1))
												mBase = m.M
												v228 = m.ExcPending
												if v228 != 0 {
													return int64(0)
												} else {
													v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+59)))
													if v229 == int32(1) {
														F_pfree(m, v70)
														mBase = m.M
														v233 = m.ExcPending
														if v233 != 0 {
															return int64(0)
														} else {
															v234 = F_strlen(m, v224)
															mBase = m.M
															*(*int32)(unsafe.Add(mBase, uint32(v64))) = v234<<(uint(int32(2))%32) + int32(16)
															v244 = v64
															m.G0 = v14 + int32(96)
															return base.I64_extend_i32_u(v244)
														}
													} else {
														v234 = F_strlen(m, v224)
														mBase = m.M
														*(*int32)(unsafe.Add(mBase, uint32(v64))) = v234<<(uint(int32(2))%32) + int32(16)
														v244 = v64
														m.G0 = v14 + int32(96)
														return base.I64_extend_i32_u(v244)
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
